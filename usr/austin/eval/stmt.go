// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package eval

import (
	"bignum";
	"eval";
	"log";
	"os";
	"go/ast";
	"go/scanner";
	"go/token";
	"strconv";
)

const (
	returnPC = ^uint(0);
	badPC = ^uint(1);
)

/*
 * Statement compiler
 */

type stmtCompiler struct {
	*blockCompiler;
	pos token.Position;
	// This statement's label, or nil if it is not labeled.
	stmtLabel *label;
	// err should be initialized to true before visiting and set
	// to false when the statement is compiled successfully.  The
	// function invoking Visit should or this with
	// blockCompiler.err.  This is less error prone than setting
	// blockCompiler.err on every failure path.
	err bool;
}

func (a *stmtCompiler) diag(format string, args ...) {
	a.diagAt(&a.pos, format, args);
}

/*
 * Flow checker
 */

type flowEnt struct {
	// Whether this flow entry is conditional.  If true, flow can
	// continue to the next PC.
	cond bool;
	// True if this will terminate flow (e.g., a return statement).
	// cond must be false and jumps must be nil if this is true.
	term bool;
	// PC's that can be reached from this flow entry.
	jumps []*uint;
	// Whether this flow entry has been visited by reachesEnd.
	visited bool;
}

type flowBlock struct {
	// If this is a goto, the target label.
	target string;
	// The inner-most block containing definitions.
	block *block;
	// The numVars from each block leading to the root of the
	// scope, starting at block.
	numVars []int;
}

type flowBuf struct {
	cb *codeBuf;
	// ents is a map from PC's to flow entries.  Any PC missing
	// from this map is assumed to reach only PC+1.
	ents map[uint] *flowEnt;
	// gotos is a map from goto positions to information on the
	// block at the point of the goto.
	gotos map[*token.Position] *flowBlock;
	// labels is a map from label name to information on the block
	// at the point of the label.  labels are tracked by name,
	// since mutliple labels at the same PC can have different
	// blocks.
	labels map[string] *flowBlock;
}

func newFlowBuf(cb *codeBuf) *flowBuf {
	return &flowBuf{cb, make(map[uint] *flowEnt), make(map[*token.Position] *flowBlock), make(map[string] *flowBlock)};
}

// put creates a flow control point for the next PC in the code buffer.
// This should be done before pushing the instruction into the code buffer.
func (f *flowBuf) put(cond bool, term bool, jumps []*uint) {
	pc := f.cb.nextPC();
	if ent, ok := f.ents[pc]; ok {
		log.Crashf("Flow entry already exists at PC %d: %+v", pc, ent);
	}
	f.ents[pc] = &flowEnt{cond, term, jumps, false};
}

// putTerm creates a flow control point at the next PC that
// unconditionally terminates execution.
func (f *flowBuf) putTerm() {
	f.put(false, true, nil);
}

// put1 creates a flow control point at the next PC that jumps to one
// PC and, if cond is true, can also continue to the PC following the
// next PC.
func (f *flowBuf) put1(cond bool, jumpPC *uint) {
	f.put(cond, false, []*uint {jumpPC});
}

func newFlowBlock(target string, b *block) *flowBlock {
	// Find the inner-most block containing definitions
	for b.numVars == 0 && b.outer != nil && b.outer.scope == b.scope {
		b = b.outer;
	}

	// Count parents leading to the root of the scope
	n := 0;
	for bp := b; bp.scope == b.scope; bp = bp.outer {
		n++;
	}

	// Capture numVars from each block to the root of the scope
	numVars := make([]int, n);
	i := 0;
	for bp := b; i < n; bp = bp.outer {
		numVars[i] = bp.numVars;
		i++;
	}

	return &flowBlock{target, b, numVars};
}

// putGoto captures the block at a goto statement.  This should be
// called in addition to putting a flow control point.
func (f *flowBuf) putGoto(pos token.Position, target string, b *block) {
	f.gotos[&pos] = newFlowBlock(target, b);
}

// putLabel captures the block at a label.
func (f *flowBuf) putLabel(name string, b *block) {
	f.labels[name] = newFlowBlock("", b);
}

// reachesEnd returns true if the end of f's code buffer can be
// reached from the given program counter.  Error reporting is the
// caller's responsibility.
func (f *flowBuf) reachesEnd(pc uint) bool {
	endPC := f.cb.nextPC();
	if pc > endPC {
		log.Crashf("Reached bad PC %d past end PC %d", pc, endPC);
	}

	for ; pc < endPC; pc++ {
		ent, ok := f.ents[pc];
		if !ok {
			continue;
		}

		if ent.visited {
			return false;
		}
		ent.visited = true;

		if ent.term {
			return false;
		}

		// If anything can reach the end, we can reach the end
		// from pc.
		for _, j := range ent.jumps {
			if f.reachesEnd(*j) {
				return true;
			}
		}
		// If the jump was conditional, we can reach the next
		// PC, so try reaching the end from it.
		if ent.cond {
			continue;
		}
		return false;
	}
	return true;
}

// gotosObeyScopes returns true if no goto statement causes any
// variables to come into scope that were not in scope at the point of
// the goto.  Reports any errors using the given compiler.
func (f *flowBuf) gotosObeyScopes(a *compiler) bool {
	for pos, src := range f.gotos {
		tgt := f.labels[src.target];

		// The target block must be a parent of this block
		numVars := src.numVars;
		b := src.block;
		for len(numVars) > 0 && b != tgt.block {
			b = b.outer;
			numVars = numVars[1:len(numVars)];
		}
		if b != tgt.block {
			// We jumped into a deeper block
			a.diagAt(pos, "goto causes variables to come into scope");
			return false;
		}

		// There must be no variables in the target block that
		// did not exist at the jump
		tgtNumVars := tgt.numVars;
		for i := range numVars {
			if tgtNumVars[i] > numVars[i] {
				a.diagAt(pos, "goto causes variables to come into scope");
				return false;
			}
		}
	}
	return true;
}

/*
 * Statement generation helpers
 */

func (a *stmtCompiler) defineVar(ident *ast.Ident, t Type) *Variable {
	v, prev := a.block.DefineVar(ident.Value, ident.Pos(), t);
	if prev != nil {
		// TODO(austin) It's silly that we have to capture
		// Pos() in a variable.
		pos := prev.Pos();
		if pos.IsValid() {
			a.diagAt(ident, "variable %s redeclared in this block\n\tprevious declaration at %s", ident.Value, &pos);
		} else {
			a.diagAt(ident, "variable %s redeclared in this block", ident.Value);
		}
		return nil;
	}

	// Initialize the variable
	index := v.Index;
	a.push(func(v *vm) {
		v.f.Vars[index] = t.Zero();
	});
	return v;
}

// TODO(austin) Move the real definition
func (a *stmtCompiler) doAssign(lhs []ast.Expr, rhs []ast.Expr, tok token.Token, declTypeExpr ast.Expr)

/*
 * Statement visitors
 */

func (a *stmtCompiler) DoBadStmt(s *ast.BadStmt) {
	// Do nothing.  Already reported by parser.
}

func (a *stmtCompiler) DoDeclStmt(s *ast.DeclStmt) {
	ok := true;

	switch decl := s.Decl.(type) {
	case *ast.BadDecl:
		// Do nothing.  Already reported by parser.
		ok = false;

	case *ast.FuncDecl:
		log.Crash("FuncDecl at statement level");

	case *ast.GenDecl:
		switch decl.Tok {
		case token.IMPORT:
			log.Crash("import at statement level");

		case token.CONST:
			log.Crashf("%v not implemented", decl.Tok);

		case token.TYPE:
			ok = a.compileTypeDecl(a.block, decl);

		case token.VAR:
			for _, spec := range decl.Specs {
				spec := spec.(*ast.ValueSpec);
				if spec.Values == nil {
					// Declaration without assignment
					if spec.Type == nil {
						// Parser should have caught
						log.Crash("Type and Values nil");
					}
					t := a.compileType(a.block, spec.Type);
					if t == nil {
						// Define placeholders
						ok = false;
					}
					for _, n := range spec.Names {
						if a.defineVar(n, t) == nil {
							ok = false;
						}
					}
				} else {
					// Decalaration with assignment
					lhs := make([]ast.Expr, len(spec.Names));
					for i, n := range spec.Names {
						lhs[i] = n;
					}
					a.doAssign(lhs, spec.Values, decl.Tok, spec.Type);
					// TODO(austin) This is rediculous.  doAssign
					// indicates failure by setting a.err.
					if a.err {
						ok = false;
					}
				}
			}
		}
	default:
		log.Crashf("Unexpected Decl type %T", s.Decl);
	}

	if ok {
		a.err = false;
	}
}

func (a *stmtCompiler) DoEmptyStmt(s *ast.EmptyStmt) {
	a.err = false;
}

func (a *stmtCompiler) DoLabeledStmt(s *ast.LabeledStmt) {
	bad := false;

	// Define label
	l, ok := a.labels[s.Label.Value];
	if ok {
		if l.resolved.IsValid() {
			a.diag("label %s redeclared in this block\n\tprevious declaration at %s", s.Label.Value, &l.resolved);
			bad = true;
		}
	} else {
		pc := badPC;
		l = &label{name: s.Label.Value, gotoPC: &pc};
		a.labels[l.name] = l;
	}
	l.desc = "regular label";
	l.resolved = s.Pos();

	// Set goto PC
	*l.gotoPC = a.nextPC();

	// Define flow entry so we can check for jumps over declarations.
	a.flow.putLabel(l.name, a.block);

	// Compile the statement.  Reuse our stmtCompiler for simplicity.
	a.pos = s.Stmt.Pos();
	a.stmtLabel = l;
	s.Stmt.Visit(a);
	if bad {
		a.err = true;
	}
}

func (a *stmtCompiler) DoExprStmt(s *ast.ExprStmt) {
	e := a.compileExpr(a.block, s.X, false);
	if e == nil {
		return;
	}

	if e.exec == nil {
		a.diag("%s cannot be used as expression statement", e.desc);
		return;
	}

	exec := e.exec;
	a.push(func(v *vm) {
		exec(v.f);
	});
	a.err = false;
}

func (a *stmtCompiler) DoIncDecStmt(s *ast.IncDecStmt) {
	// Create temporary block for extractEffect
	bc := a.enterChild();
	defer bc.exit();

	l := a.compileExpr(bc.block, s.X, false);
	if l == nil {
		return;
	}

	if l.evalAddr == nil {
		l.diag("cannot assign to %s", l.desc);
		return;
	}
	if !(l.t.isInteger() || l.t.isFloat()) {
		l.diagOpType(s.Tok, l.t);
		return;
	}

	effect, l := l.extractEffect();

	one := l.copy();
	one.pos = s.Pos();
	one.t = IdealIntType;
	one.evalIdealInt = func() *bignum.Integer { return bignum.Int(1) };

	var op token.Token;
	switch s.Tok {
	case token.INC:
		op = token.ADD;
	case token.DEC:
		op = token.SUB;
	default:
		log.Crashf("Unexpected IncDec token %v", s.Tok);
	}

	binop := l.copy();
	binop.pos = s.Pos();
	binop.doBinaryExpr(op, l, one);
	if binop.t == nil {
		return;
	}

	assign := a.compileAssign(s.Pos(), l.t, []*exprCompiler{binop}, "", "");
	if assign == nil {
		log.Crashf("compileAssign type check failed");
	}

	lf := l.evalAddr;
	a.push(func(v *vm) {
		effect(v.f);
		assign(lf(v.f), v.f);
	});
	a.err = false;
}

func (a *stmtCompiler) doAssign(lhs []ast.Expr, rhs []ast.Expr, tok token.Token, declTypeExpr ast.Expr) {
	bad := false;

	// Compile right side first so we have the types when
	// compiling the left side and so we don't see definitions
	// made on the left side.
	rs := make([]*exprCompiler, len(rhs));
	for i, re := range rhs {
		rs[i] = a.compileExpr(a.block, re, false);
		if rs[i] == nil {
			bad = true;
		}
	}

	errOp := "assignment";
	if tok == token.DEFINE || tok == token.VAR {
		errOp = "declaration";
	}
	ac, ok := a.checkAssign(a.pos, rs, errOp, "value");
	if !ok {
		bad = true;
	}

	// If this is a definition and the LHS is too big, we won't be
	// able to produce the usual error message because we can't
	// begin to infer the types of the LHS.
	if (tok == token.DEFINE || tok == token.VAR) && len(lhs) > len(ac.rmt.Elems) {
		a.diag("not enough values for definition");
		bad = true;
	}

	// Compile left type if there is one
	var declType Type;
	if declTypeExpr != nil {
		declType = a.compileType(a.block, declTypeExpr);
		if declType == nil {
			bad = true;
		}
	}

	// Compile left side
	ls := make([]*exprCompiler, len(lhs));
	nDefs := 0;
	for i, le := range lhs {
		// If this is a definition, get the identifier and its type
		var ident *ast.Ident;
		var lt Type;
		switch tok {
		case token.DEFINE:
			// Check that it's an identifier
			ident, ok = le.(*ast.Ident);
			if !ok {
				a.diagAt(le, "left side of := must be a name");
				bad = true;
				// Suppress new defitions errors
				nDefs++;
				continue;
			}

			// Is this simply an assignment?
			if _, ok := a.block.defs[ident.Value]; ok {
				ident = nil;
				break;
			}
			nDefs++;

		case token.VAR:
			ident = le.(*ast.Ident);
		}

		// If it's a definition, get or infer its type.
		if ident != nil {
			// Compute the identifier's type from the RHS
			// type.  We use the computed MultiType so we
			// don't have to worry about unpacking.
			switch {
			case declTypeExpr != nil:
				// We have a declaration type, use it.
				// If declType is nil, we gave an
				// error when we compiled it.
				lt = declType;

			case i >= len(ac.rmt.Elems):
				// Define a placeholder.  We already
				// gave the "not enough" error above.
				lt = nil;

			case ac.rmt.Elems[i] == nil:
				// We gave the error when we compiled
				// the RHS.
				lt = nil;

			case ac.rmt.Elems[i].isIdeal():
				// If the type is absent and the
				// corresponding expression is a
				// constant expression of ideal
				// integer or ideal float type, the
				// type of the declared variable is
				// int or float respectively.
				switch {
				case ac.rmt.Elems[i].isInteger():
					lt = IntType;
				case ac.rmt.Elems[i].isFloat():
					lt = FloatType;
				default:
					log.Crashf("unexpected ideal type %v", rs[i].t);
				}

			default:
				lt = ac.rmt.Elems[i];
			}
		}

		// If it's a definition, define the identifier
		if ident != nil {
			if a.defineVar(ident, lt) == nil {
				bad = true;
				continue;
			}
		}

		// Compile LHS
		ls[i] = a.compileExpr(a.block, le, false);
		if ls[i] == nil {
			bad = true;
			continue;
		}

		if ls[i].evalAddr == nil {
			ls[i].diag("cannot assign to %s", ls[i].desc);
			bad = true;
			continue;
		}
	}

	// A short variable declaration may redeclare variables
	// provided they were originally declared in the same block
	// with the same type, and at least one of the variables is
	// new.
	if tok == token.DEFINE && nDefs == 0 {
		a.diag("at least one new variable must be declared");
		return;
	}

	if bad {
		return;
	}

	// Create assigner
	var lt Type;
	n := len(lhs);
	if n == 1 {
		lt = ls[0].t;
	} else {
		lts := make([]Type, len(ls));
		for i, l := range ls {
			if l != nil {
				lts[i] = l.t;
			}
		}
		lt = NewMultiType(lts);
	}
	assign := ac.compile(lt);
	if assign == nil {
		return;
	}

	// Compile
	if n == 1 {
		// Don't need temporaries and can avoid []Value.
		lf := ls[0].evalAddr;
		a.push(func(v *vm) { assign(lf(v.f), v.f) });
	} else if tok == token.VAR || (tok == token.DEFINE && nDefs == n) {
		// Don't need temporaries
		lfs := make([]func(*Frame) Value, n);
		for i, l := range ls {
			lfs[i] = l.evalAddr;
		}
		a.push(func(v *vm) {
			dest := make([]Value, n);
			for i, lf := range lfs {
				dest[i] = lf(v.f);
			}
			assign(multiV(dest), v.f);
		});
	} else {
		// Need temporaries
		lmt := lt.(*MultiType);
		lfs := make([]func(*Frame) Value, n);
		for i, l := range ls {
			lfs[i] = l.evalAddr;
		}
		a.push(func(v *vm) {
			temp := lmt.Zero().(multiV);
			assign(temp, v.f);
			// Copy to destination
			for i := 0; i < n; i ++ {
				// TODO(austin) Need to evaluate LHS
				// before RHS
				lfs[i](v.f).Assign(temp[i]);
			}
		});
	}
	a.err = false;
}

var assignOpToOp = map[token.Token] token.Token {
	token.ADD_ASSIGN : token.ADD,
	token.SUB_ASSIGN : token.SUB,
	token.MUL_ASSIGN : token.MUL,
	token.QUO_ASSIGN : token.QUO,
	token.REM_ASSIGN : token.REM,

	token.AND_ASSIGN : token.AND,
	token.OR_ASSIGN  : token.OR,
        token.XOR_ASSIGN : token.XOR,
        token.SHL_ASSIGN : token.SHL,
        token.SHR_ASSIGN : token.SHR,
        token.AND_NOT_ASSIGN : token.AND_NOT,
}

func (a *stmtCompiler) doAssignOp(s *ast.AssignStmt) {
	if len(s.Lhs) != 1 || len(s.Rhs) != 1 {
		a.diag("tuple assignment cannot be combined with an arithmetic operation");
		return;
	}

	// Create temporary block for extractEffect
	bc := a.enterChild();
	defer bc.exit();

	l := a.compileExpr(bc.block, s.Lhs[0], false);
	r := a.compileExpr(bc.block, s.Rhs[0], false);
	if l == nil || r == nil {
		return;
	}

	if l.evalAddr == nil {
		l.diag("cannot assign to %s", l.desc);
		return;
	}

	effect, l := l.extractEffect();

	binop := r.copy();
	binop.pos = s.TokPos;
	binop.doBinaryExpr(assignOpToOp[s.Tok], l, r);
	if binop.t == nil {
		return;
	}

	assign := a.compileAssign(s.Pos(), l.t, []*exprCompiler{binop}, "assignment", "value");
	if assign == nil {
		log.Crashf("compileAssign type check failed");
	}

	lf := l.evalAddr;
	a.push(func(v *vm) {
		effect(v.f);
		assign(lf(v.f), v.f);
	});
	a.err = false;
}

func (a *stmtCompiler) DoAssignStmt(s *ast.AssignStmt) {
	switch s.Tok {
	case token.ASSIGN, token.DEFINE:
		a.doAssign(s.Lhs, s.Rhs, s.Tok, nil);

	default:
		a.doAssignOp(s);
	}
}

func (a *stmtCompiler) DoGoStmt(s *ast.GoStmt) {
	log.Crash("Not implemented");
}

func (a *stmtCompiler) DoDeferStmt(s *ast.DeferStmt) {
	log.Crash("Not implemented");
}

func (a *stmtCompiler) DoReturnStmt(s *ast.ReturnStmt) {
	if a.fnType == nil {
		a.diag("cannot return at the top level");
		return;
	}

	if len(s.Results) == 0 && (len(a.fnType.Out) == 0 || a.outVarsNamed) {
		// Simple case.  Simply exit from the function.
		a.flow.putTerm();
		a.push(func(v *vm) { v.pc = returnPC });
		a.err = false;
		return;
	}

	// Compile expressions
	bad := false;
	rs := make([]*exprCompiler, len(s.Results));
	for i, re := range s.Results {
		rs[i] = a.compileExpr(a.block, re, false);
		if rs[i] == nil {
			bad = true;
		}
	}
	if bad {
		return;
	}

	// Create assigner

	// However, if the expression list in the "return" statement
	// is a single call to a multi-valued function, the values
	// returned from the called function will be returned from
	// this one.
	assign := a.compileAssign(s.Pos(), NewMultiType(a.fnType.Out), rs, "return", "value");
	if assign == nil {
		return;
	}

	// XXX(Spec) "The result types of the current function and the
	// called function must match."  Match is fuzzy.  It should
	// say that they must be assignment compatible.

	// Compile
	start := len(a.fnType.In);
	nout := len(a.fnType.Out);
	a.flow.putTerm();
	a.push(func(v *vm) {
		assign(multiV(v.f.Vars[start:start+nout]), v.f);
		v.pc = returnPC;
	});
	a.err = false;
}

func (a *stmtCompiler) findLexicalLabel(name *ast.Ident, pred func(*label) bool, errOp, errCtx string) *label {
	bc := a.blockCompiler;
	for ; bc != nil; bc = bc.parent {
		if bc.label == nil {
			continue;
		}
		l := bc.label;
		if name == nil && pred(l) {
			return l;
		}
		if name != nil && l.name == name.Value {
			if !pred(l) {
				a.diag("cannot %s to %s %s", errOp, l.desc, l.name);
				return nil;
			}
			return l;
		}
	}
	if name == nil {
		a.diag("%s outside %s", errOp, errCtx);
	} else {
		a.diag("%s label %s not defined", errOp, name.Value);
	}
	return nil;
}

func (a *stmtCompiler) DoBranchStmt(s *ast.BranchStmt) {
	var pc *uint;

	switch s.Tok {
	case token.BREAK:
		l := a.findLexicalLabel(s.Label, func(l *label) bool { return l.breakPC != nil }, "break", "for loop, switch, or select");
		if l == nil {
			return;
		}
		pc = l.breakPC;

	case token.CONTINUE:
		l := a.findLexicalLabel(s.Label, func(l *label) bool { return l.continuePC != nil }, "continue", "for loop");
		if l == nil {
			return;
		}
		pc = l.continuePC;

	case token.GOTO:
		l, ok := a.labels[s.Label.Value];
		if !ok {
			pc := badPC;
			l = &label{name: s.Label.Value, desc: "unresolved label", gotoPC: &pc, used: s.Pos()};
			a.labels[l.name] = l;
		}

		pc = l.gotoPC;
		a.flow.putGoto(s.Pos(), l.name, a.block);

	case token.FALLTHROUGH:
		log.Crash("fallthrough not implemented");

	default:
		log.Crash("Unexpected branch token %v", s.Tok);
	}

	a.flow.put1(false, pc);
	a.push(func(v *vm) { v.pc = *pc });
	a.err = false;
}

func (a *stmtCompiler) DoBlockStmt(s *ast.BlockStmt) {
	bc := a.enterChild();
	bc.compileStmts(s);
	bc.exit();

	a.err = false;
}

func (a *stmtCompiler) DoIfStmt(s *ast.IfStmt) {
	// The scope of any variables declared by [the init] statement
	// extends to the end of the "if" statement and the variables
	// are initialized once before the statement is entered.
	//
	// XXX(Spec) What this really wants to say is that there's an
	// implicit scope wrapping every if, for, and switch
	// statement.  This is subtly different from what it actually
	// says when there's a non-block else clause, because that
	// else claus has to execute in a scope that is *not* the
	// surrounding scope.
	bc := a.enterChild();
	defer bc.exit();

	// Compile init statement, if any
	if s.Init != nil {
		bc.compileStmt(s.Init);
	}

	elsePC := badPC;
	endPC := badPC;

	// Compile condition, if any.  If there is no condition, we
	// fall through to the body.
	bad := false;
	if s.Cond != nil {
		e := bc.compileExpr(bc.block, s.Cond, false);
		switch {
		case e == nil:
			bad = true;
		case !e.t.isBoolean():
			e.diag("'if' condition must be boolean\n\t%v", e.t);
			bad = true;
		default:
			eval := e.asBool();
			a.flow.put1(true, &elsePC);
			a.push(func(v *vm) {
				if !eval(v.f) {
					v.pc = elsePC;
				}
			});
		}
	}

	// Compile body
	body := bc.enterChild();
	body.compileStmts(s.Body);
	body.exit();

	// Compile else
	if s.Else != nil {
		// Skip over else if we executed the body
		a.flow.put1(false, &endPC);
		a.push(func(v *vm) {
			v.pc = endPC;
		});
		elsePC = a.nextPC();
		bc.compileStmt(s.Else);
	} else {
		elsePC = a.nextPC();
	}
	endPC = a.nextPC();

	if !bad {
		a.err = false;
	}
}

func (a *stmtCompiler) DoCaseClause(s *ast.CaseClause) {
	log.Crash("Not implemented");
}

func (a *stmtCompiler) DoSwitchStmt(s *ast.SwitchStmt) {
	log.Crash("Not implemented");
}

func (a *stmtCompiler) DoTypeCaseClause(s *ast.TypeCaseClause) {
	log.Crash("Not implemented");
}

func (a *stmtCompiler) DoTypeSwitchStmt(s *ast.TypeSwitchStmt) {
	log.Crash("Not implemented");
}

func (a *stmtCompiler) DoCommClause(s *ast.CommClause) {
	log.Crash("Not implemented");
}

func (a *stmtCompiler) DoSelectStmt(s *ast.SelectStmt) {
	log.Crash("Not implemented");
}

func (a *stmtCompiler) DoForStmt(s *ast.ForStmt) {
	// Wrap the entire for in a block.
	bc := a.enterChild();
	defer bc.exit();

	// Compile init statement, if any
	if s.Init != nil {
		bc.compileStmt(s.Init);
	}

	bodyPC := badPC;
	postPC := badPC;
	checkPC := badPC;
	endPC := badPC;

	// Jump to condition check.  We generate slightly less code by
	// placing the condition check after the body.
	a.flow.put1(false, &checkPC);
	a.push(func(v *vm) { v.pc = checkPC });

	// Compile body
	bodyPC = a.nextPC();
	body := bc.enterChild();
	if a.stmtLabel != nil {
		body.label = a.stmtLabel;
	} else {
		body.label = &label{resolved: s.Pos()};
	}
	body.label.desc = "for loop";
	body.label.breakPC = &endPC;
	body.label.continuePC = &postPC;
	body.compileStmts(s.Body);
	body.exit();

	// Compile post, if any
	postPC = a.nextPC();
	if s.Post != nil {
		// TODO(austin) Does the parser disallow short
		// declarations in s.Post?
		bc.compileStmt(s.Post);
	}

	// Compile condition check, if any
	bad := false;
	checkPC = a.nextPC();
	if s.Cond == nil {
		// If the condition is absent, it is equivalent to true.
		a.flow.put1(false, &bodyPC);
		a.push(func(v *vm) { v.pc = bodyPC });
	} else {
		e := bc.compileExpr(bc.block, s.Cond, false);
		switch {
		case e == nil:
			bad = true;
		case !e.t.isBoolean():
			a.diag("'for' condition must be boolean\n\t%v", e.t);
			bad = true;
		default:
			eval := e.asBool();
			a.flow.put1(true, &bodyPC);
			a.push(func(v *vm) {
				if eval(v.f) {
					v.pc = bodyPC;
				}
			});
		}
	}

	endPC = a.nextPC();

	if !bad {
		a.err = false;
	}
}

func (a *stmtCompiler) DoRangeStmt(s *ast.RangeStmt) {
	log.Crash("Not implemented");
}

/*
 * Block compiler
 */

func (a *blockCompiler) compileStmt(s ast.Stmt) {
	if a.block.inner != nil {
		log.Crash("Child scope still entered");
	}
	sc := &stmtCompiler{a, s.Pos(), nil, true};
	s.Visit(sc);
	if a.block.inner != nil {
		log.Crash("Forgot to exit child scope");
	}
	a.err = a.err || sc.err;
}

func (a *blockCompiler) compileStmts(block *ast.BlockStmt) {
	for i, sub := range block.List {
		a.compileStmt(sub);
	}
}

func (a *blockCompiler) enterChild() *blockCompiler {
	block := a.block.enterChild();
	return &blockCompiler{
		funcCompiler: a.funcCompiler,
		block: block,
		parent: a,
	};
}

func (a *blockCompiler) exit() {
	a.block.exit();
}

/*
 * Function compiler
 */

func (a *compiler) compileFunc(b *block, decl *FuncDecl, body *ast.BlockStmt) (func (f *Frame) Func) {
	// Create body scope
	//
	// The scope of a parameter or result is the body of the
	// corresponding function.
	bodyScope := b.ChildScope();
	defer bodyScope.exit();
	for i, t := range decl.Type.In {
		if decl.InNames[i] != nil {
			bodyScope.DefineVar(decl.InNames[i].Value, decl.InNames[i].Pos(), t);
		} else {
			bodyScope.DefineSlot(t);
		}
	}
	for i, t := range decl.Type.Out {
		if decl.OutNames[i] != nil {
			bodyScope.DefineVar(decl.OutNames[i].Value, decl.OutNames[i].Pos(), t);
		} else {
			bodyScope.DefineSlot(t);
		}
	}

	// Create block context
	cb := newCodeBuf();
	fc := &funcCompiler{
		compiler: a,
		fnType: decl.Type,
		outVarsNamed: len(decl.OutNames) > 0 && decl.OutNames[0] != nil,
		codeBuf: cb,
		flow: newFlowBuf(cb),
		labels: make(map[string] *label),
		err: false,
	};
	bc := &blockCompiler{
		funcCompiler: fc,
		block: bodyScope.block,
	};

	// Compile body
	bc.compileStmts(body);
	fc.checkLabels();

	if fc.err {
		return nil;
	}

	// Check that the body returned if necessary.  We only check
	// this if there were no errors compiling the body.
	if len(decl.Type.Out) > 0 && fc.flow.reachesEnd(0) {
		// XXX(Spec) Not specified.
		a.diagAt(&body.Rbrace, "function ends without a return statement");
		return nil;
	}

	code := fc.get();
	maxVars := bodyScope.maxVars;
	return func(f *Frame) Func { return &evalFunc{f, maxVars, code} };
}

// Checks that labels were resolved and that all jumps obey scoping
// rules.  Reports an error and set fc.err if any check fails.
func (a *funcCompiler) checkLabels() {
	bad := false;
	for _, l := range a.labels {
		if !l.resolved.IsValid() {
			a.diagAt(&l.used, "label %s not defined", l.name);
			bad = true;
		}
	}
	if bad {
		a.err = true;
		// Don't check scopes if we have unresolved labels
		return;
	}

	// Executing the "goto" statement must not cause any variables
	// to come into scope that were not already in scope at the
	// point of the goto.
	if !a.flow.gotosObeyScopes(a.compiler) {
		a.err = true;
	}
}

/*
 * Testing interface
 */

type Stmt struct {
	f func (f *Frame);
}

func (s *Stmt) Exec(f *Frame) {
	s.f(f);
}

func CompileStmts(scope *Scope, stmts []ast.Stmt) (*Stmt, os.Error) {
	errors := scanner.NewErrorVector();
	cc := &compiler{errors};
	cb := newCodeBuf();
	fc := &funcCompiler{
		compiler: cc,
		fnType: nil,
		outVarsNamed: false,
		codeBuf: cb,
		flow: newFlowBuf(cb),
		labels: make(map[string] *label),
		err: false,
	};
	bc := &blockCompiler{
		funcCompiler: fc,
		block: scope.block,
	};
	out := make([]*Stmt, len(stmts));
	for i, stmt := range stmts {
		bc.compileStmt(stmt);
	}
	fc.checkLabels();
	if fc.err {
		return nil, errors.GetError(scanner.Sorted);
	}
	code := fc.get();
	return &Stmt{func(f *Frame) { code.exec(f); }}, nil;
}
