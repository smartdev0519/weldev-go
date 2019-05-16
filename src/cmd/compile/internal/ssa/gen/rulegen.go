// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gen

// This program generates Go code that applies rewrite rules to a Value.
// The generated code implements a function of type func (v *Value) bool
// which reports whether if did something.
// Ideas stolen from Swift: http://www.hpl.hp.com/techreports/Compaq-DEC/WRL-2000-2.html

package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"strings"
)

// rule syntax:
//  sexpr [&& extra conditions] -> [@block] sexpr
//
// sexpr are s-expressions (lisp-like parenthesized groupings)
// sexpr ::= [variable:](opcode sexpr*)
//         | variable
//         | <type>
//         | [auxint]
//         | {aux}
//
// aux      ::= variable | {code}
// type     ::= variable | {code}
// variable ::= some token
// opcode   ::= one of the opcodes from the *Ops.go files

// extra conditions is just a chunk of Go that evaluates to a boolean. It may use
// variables declared in the matching sexpr. The variable "v" is predefined to be
// the value matched by the entire rule.

// If multiple rules match, the first one in file order is selected.

var genLog = flag.Bool("log", false, "generate code that logs; for debugging only")

type Rule struct {
	rule string
	loc  string // file name & line number
}

func (r Rule) String() string {
	return fmt.Sprintf("rule %q at %s", r.rule, r.loc)
}

func normalizeSpaces(s string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(s)), " ")
}

// parse returns the matching part of the rule, additional conditions, and the result.
func (r Rule) parse() (match, cond, result string) {
	s := strings.Split(r.rule, "->")
	if len(s) != 2 {
		log.Fatalf("no arrow in %s", r)
	}
	match = normalizeSpaces(s[0])
	result = normalizeSpaces(s[1])
	cond = ""
	if i := strings.Index(match, "&&"); i >= 0 {
		cond = normalizeSpaces(match[i+2:])
		match = normalizeSpaces(match[:i])
	}
	return match, cond, result
}

func genRules(arch arch)          { genRulesSuffix(arch, "") }
func genSplitLoadRules(arch arch) { genRulesSuffix(arch, "splitload") }

func genRulesSuffix(arch arch, suff string) {
	// Open input file.
	text, err := os.Open(arch.name + suff + ".rules")
	if err != nil {
		if suff == "" {
			// All architectures must have a plain rules file.
			log.Fatalf("can't read rule file: %v", err)
		}
		// Some architectures have bonus rules files that others don't share. That's fine.
		return
	}

	// oprules contains a list of rules for each block and opcode
	blockrules := map[string][]Rule{}
	oprules := map[string][]Rule{}

	// read rule file
	scanner := bufio.NewScanner(text)
	rule := ""
	var lineno int
	var ruleLineno int // line number of "->"
	for scanner.Scan() {
		lineno++
		line := scanner.Text()
		if i := strings.Index(line, "//"); i >= 0 {
			// Remove comments. Note that this isn't string safe, so
			// it will truncate lines with // inside strings. Oh well.
			line = line[:i]
		}
		rule += " " + line
		rule = strings.TrimSpace(rule)
		if rule == "" {
			continue
		}
		if !strings.Contains(rule, "->") {
			continue
		}
		if ruleLineno == 0 {
			ruleLineno = lineno
		}
		if strings.HasSuffix(rule, "->") {
			continue
		}
		if unbalanced(rule) {
			continue
		}

		loc := fmt.Sprintf("%s%s.rules:%d", arch.name, suff, ruleLineno)
		for _, rule2 := range expandOr(rule) {
			for _, rule3 := range commute(rule2, arch) {
				r := Rule{rule: rule3, loc: loc}
				if rawop := strings.Split(rule3, " ")[0][1:]; isBlock(rawop, arch) {
					blockrules[rawop] = append(blockrules[rawop], r)
					continue
				}
				// Do fancier value op matching.
				match, _, _ := r.parse()
				op, oparch, _, _, _, _ := parseValue(match, arch, loc)
				opname := fmt.Sprintf("Op%s%s", oparch, op.name)
				oprules[opname] = append(oprules[opname], r)
			}
		}
		rule = ""
		ruleLineno = 0
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner failed: %v\n", err)
	}
	if unbalanced(rule) {
		log.Fatalf("%s.rules:%d: unbalanced rule: %v\n", arch.name, lineno, rule)
	}

	// Order all the ops.
	var ops []string
	for op := range oprules {
		ops = append(ops, op)
	}
	sort.Strings(ops)

	file := &File{arch: arch, suffix: suff}
	const chunkSize = 10
	// Main rewrite routine is a switch on v.Op.
	fn := &Func{kind: "Value"}

	sw := &Switch{expr: exprf("v.Op")}
	for _, op := range ops {
		var ors []string
		for chunk := 0; chunk < len(oprules[op]); chunk += chunkSize {
			ors = append(ors, fmt.Sprintf("rewriteValue%s%s_%s_%d(v)", arch.name, suff, op, chunk))
		}
		swc := &Case{expr: exprf(op)}
		swc.add(stmtf("return %s", strings.Join(ors, " || ")))
		sw.add(swc)
	}
	fn.add(sw)
	fn.add(stmtf("return false"))
	file.add(fn)

	// Generate a routine per op. Note that we don't make one giant routine
	// because it is too big for some compilers.
	for _, op := range ops {
		rules := oprules[op]
		// rr is kept between chunks, so that a following chunk checks
		// that the previous one ended with a rule that wasn't
		// unconditional.
		var rr *RuleRewrite
		for chunk := 0; chunk < len(rules); chunk += chunkSize {
			endchunk := chunk + chunkSize
			if endchunk > len(rules) {
				endchunk = len(rules)
			}
			fn := &Func{
				kind:   "Value",
				suffix: fmt.Sprintf("_%s_%d", op, chunk),
			}
			var rewrites bodyBase
			for _, rule := range rules[chunk:endchunk] {
				if rr != nil && !rr.canFail {
					log.Fatalf("unconditional rule %s is followed by other rules", rr.match)
				}
				rr = &RuleRewrite{loc: rule.loc}
				rr.match, rr.cond, rr.result = rule.parse()
				pos, _ := genMatch(rr, arch, rr.match)
				if pos == "" {
					pos = "v.Pos"
				}
				if rr.cond != "" {
					rr.add(breakf("!(%s)", rr.cond))
				}
				genResult(rr, arch, rr.result, pos)
				if *genLog {
					rr.add(stmtf("logRule(%q)", rule.loc))
				}
				rewrites.add(rr)
			}

			// TODO(mvdan): remove unused vars later instead
			uses := make(map[string]int)
			walk(&rewrites, func(node Node) {
				switch node := node.(type) {
				case *Declare:
					// work around var shadowing
					// TODO(mvdan): forbid it instead.
					uses[node.name] = math.MinInt32
				case *ast.Ident:
					uses[node.Name]++
				}
			})
			if uses["b"]+uses["config"]+uses["fe"]+uses["typ"] > 0 {
				fn.add(declf("b", "v.Block"))
			}
			if uses["config"] > 0 {
				fn.add(declf("config", "b.Func.Config"))
			}
			if uses["fe"] > 0 {
				fn.add(declf("fe", "b.Func.fe"))
			}
			if uses["typ"] > 0 {
				fn.add(declf("typ", "&b.Func.Config.Types"))
			}
			fn.add(rewrites.list...)
			if rr.canFail {
				fn.add(stmtf("return false"))
			}
			file.add(fn)
		}
	}

	// Generate block rewrite function. There are only a few block types
	// so we can make this one function with a switch.
	fn = &Func{kind: "Block"}
	fn.add(declf("config", "b.Func.Config"))
	// TODO(mvdan): declare these only if needed
	fn.add(declf("typ", "&config.Types"))
	fn.add(stmtf("_ = typ"))
	fn.add(declf("v", "b.Control"))
	fn.add(stmtf("_ = v"))

	sw = &Switch{expr: exprf("b.Kind")}
	ops = ops[:0]
	for op := range blockrules {
		ops = append(ops, op)
	}
	sort.Strings(ops)
	for _, op := range ops {
		swc := &Case{expr: exprf("%s", blockName(op, arch))}
		for _, rule := range blockrules[op] {
			swc.add(genBlockRewrite(rule, arch))
		}
		sw.add(swc)
	}
	fn.add(sw)
	fn.add(stmtf("return false"))
	file.add(fn)

	// gofmt result
	buf := new(bytes.Buffer)
	fprint(buf, file)
	b := buf.Bytes()
	src, err := format.Source(b)
	if err != nil {
		fmt.Printf("%s\n", b)
		panic(err)
	}
	// Write to file
	if err := ioutil.WriteFile("../rewrite"+arch.name+suff+".go", src, 0666); err != nil {
		log.Fatalf("can't write output: %v\n", err)
	}
}

func walk(node Node, fn func(Node)) {
	fn(node)
	switch node := node.(type) {
	case *bodyBase:
	case *File:
	case *Func:
	case *Switch:
		walk(node.expr, fn)
	case *Case:
		walk(node.expr, fn)
	case *RuleRewrite:
	case *Declare:
		walk(node.value, fn)
	case *CondBreak:
		walk(node.expr, fn)
	case ast.Node:
		ast.Inspect(node, func(node ast.Node) bool {
			fn(node)
			return true
		})
	default:
		log.Fatalf("cannot walk %T", node)
	}
	if wb, ok := node.(interface{ body() []Statement }); ok {
		for _, node := range wb.body() {
			walk(node, fn)
		}
	}
	fn(nil)
}

func fprint(w io.Writer, n Node) {
	switch n := n.(type) {
	case *File:
		fmt.Fprintf(w, "// Code generated from gen/%s%s.rules; DO NOT EDIT.\n", n.arch.name, n.suffix)
		fmt.Fprintf(w, "// generated with: cd gen; go run *.go\n")
		fmt.Fprintf(w, "\npackage ssa\n")
		// TODO(mvdan): keep the needed imports only
		fmt.Fprintln(w, "import \"fmt\"")
		fmt.Fprintln(w, "import \"math\"")
		fmt.Fprintln(w, "import \"cmd/internal/obj\"")
		fmt.Fprintln(w, "import \"cmd/internal/objabi\"")
		fmt.Fprintln(w, "import \"cmd/compile/internal/types\"")
		fmt.Fprintln(w, "var _ = fmt.Println   // in case not otherwise used")
		fmt.Fprintln(w, "var _ = math.MinInt8  // in case not otherwise used")
		fmt.Fprintln(w, "var _ = obj.ANOP      // in case not otherwise used")
		fmt.Fprintln(w, "var _ = objabi.GOROOT // in case not otherwise used")
		fmt.Fprintln(w, "var _ = types.TypeMem // in case not otherwise used")
		fmt.Fprintln(w)
		for _, f := range n.list {
			f := f.(*Func)
			fmt.Fprintf(w, "func rewrite%s%s%s%s(", f.kind, n.arch.name, n.suffix, f.suffix)
			fmt.Fprintf(w, "%c *%s) bool {\n", strings.ToLower(f.kind)[0], f.kind)
			for _, n := range f.list {
				fprint(w, n)
			}
			fmt.Fprintf(w, "}\n")
		}
	case *Switch:
		fmt.Fprintf(w, "switch ")
		fprint(w, n.expr)
		fmt.Fprintf(w, " {\n")
		for _, n := range n.list {
			fprint(w, n)
		}
		fmt.Fprintf(w, "}\n")
	case *Case:
		fmt.Fprintf(w, "case ")
		fprint(w, n.expr)
		fmt.Fprintf(w, ":\n")
		for _, n := range n.list {
			fprint(w, n)
		}
	case *RuleRewrite:
		fmt.Fprintf(w, "// match: %s\n", n.match)
		fmt.Fprintf(w, "// cond: %s\n", n.cond)
		fmt.Fprintf(w, "// result: %s\n", n.result)
		if n.checkOp != "" {
			fmt.Fprintf(w, "for v.Op == %s {\n", n.checkOp)
		} else {
			fmt.Fprintf(w, "for {\n")
		}
		for _, n := range n.list {
			fprint(w, n)
		}
		fmt.Fprintf(w, "return true\n}\n")
	case *Declare:
		fmt.Fprintf(w, "%s := ", n.name)
		fprint(w, n.value)
		fmt.Fprintln(w)
	case *CondBreak:
		fmt.Fprintf(w, "if ")
		fprint(w, n.expr)
		fmt.Fprintf(w, " {\nbreak\n}\n")
	case ast.Node:
		printer.Fprint(w, emptyFset, n)
		if _, ok := n.(ast.Stmt); ok {
			fmt.Fprintln(w)
		}
	default:
		log.Fatalf("cannot print %T", n)
	}
}

var emptyFset = token.NewFileSet()

// Node can be a Statement or an ast.Expr.
type Node interface{}

// Statement can be one of our high-level statement struct types, or an
// ast.Stmt under some limited circumstances.
type Statement interface{}

// bodyBase is shared by all of our statement psuedo-node types which can
// contain other statements.
type bodyBase struct {
	list    []Statement
	canFail bool
}

func (w *bodyBase) body() []Statement { return w.list }
func (w *bodyBase) add(nodes ...Statement) {
	w.list = append(w.list, nodes...)
	for _, node := range nodes {
		if _, ok := node.(*CondBreak); ok {
			w.canFail = true
		}
	}
}

// declared reports if the body contains a Declare with the given name.
func (w *bodyBase) declared(name string) bool {
	for _, s := range w.list {
		if decl, ok := s.(*Declare); ok && decl.name == name {
			return true
		}
	}
	return false
}

// These types define some high-level statement struct types, which can be used
// as a Statement. This allows us to keep some node structs simpler, and have
// higher-level nodes such as an entire rule rewrite.
//
// Note that ast.Expr is always used as-is; we don't declare our own expression
// nodes.
type (
	File struct {
		bodyBase // []*Func
		arch     arch
		suffix   string
	}
	Func struct {
		bodyBase
		kind   string // "Value" or "Block"
		suffix string
	}
	Switch struct {
		bodyBase // []*Case
		expr     ast.Expr
	}
	Case struct {
		bodyBase
		expr ast.Expr
	}
	RuleRewrite struct {
		bodyBase
		match, cond, result string // top comments
		checkOp             string

		alloc int    // for unique var names
		loc   string // file name & line number of the original rule
	}
	Declare struct {
		name  string
		value ast.Expr
	}
	CondBreak struct {
		expr ast.Expr
	}
)

// exprf parses a Go expression generated from fmt.Sprintf, panicking if an
// error occurs.
func exprf(format string, a ...interface{}) ast.Expr {
	src := fmt.Sprintf(format, a...)
	expr, err := parser.ParseExpr(src)
	if err != nil {
		log.Fatalf("expr parse error on %q: %v", src, err)
	}
	return expr
}

// stmtf parses a Go statement generated from fmt.Sprintf. This function is only
// meant for simple statements that don't have a custom Statement node declared
// in this package, such as ast.ReturnStmt or ast.ExprStmt.
func stmtf(format string, a ...interface{}) Statement {
	src := fmt.Sprintf(format, a...)
	fsrc := "package p\nfunc _() {\n" + src + "\n}\n"
	file, err := parser.ParseFile(token.NewFileSet(), "", fsrc, 0)
	if err != nil {
		log.Fatalf("stmt parse error on %q: %v", src, err)
	}
	return file.Decls[0].(*ast.FuncDecl).Body.List[0]
}

// declf constructs a simple "name := value" declaration, using exprf for its
// value.
func declf(name, format string, a ...interface{}) *Declare {
	return &Declare{name, exprf(format, a...)}
}

// breakf constructs a simple "if cond { break }" statement, using exprf for its
// condition.
func breakf(format string, a ...interface{}) *CondBreak {
	return &CondBreak{exprf(format, a...)}
}

func genBlockRewrite(rule Rule, arch arch) *RuleRewrite {
	rr := &RuleRewrite{loc: rule.loc}
	rr.match, rr.cond, rr.result = rule.parse()
	_, _, _, aux, s := extract(rr.match) // remove parens, then split

	// check match of control value
	pos := ""
	if s[0] != "nil" {
		if strings.Contains(s[0], "(") {
			pos, rr.checkOp = genMatch0(rr, arch, s[0], "v")
		} else {
			rr.add(declf(s[0], "b.Control"))
		}
	}
	if aux != "" {
		rr.add(declf(aux, "b.Aux"))
	}
	if rr.cond != "" {
		rr.add(breakf("!(%s)", rr.cond))
	}

	// Rule matches. Generate result.
	outop, _, _, aux, t := extract(rr.result) // remove parens, then split
	newsuccs := t[1:]

	// Check if newsuccs is the same set as succs.
	succs := s[1:]
	m := map[string]bool{}
	for _, succ := range succs {
		if m[succ] {
			log.Fatalf("can't have a repeat successor name %s in %s", succ, rule)
		}
		m[succ] = true
	}
	for _, succ := range newsuccs {
		if !m[succ] {
			log.Fatalf("unknown successor %s in %s", succ, rule)
		}
		delete(m, succ)
	}
	if len(m) != 0 {
		log.Fatalf("unmatched successors %v in %s", m, rule)
	}

	rr.add(stmtf("b.Kind = %s", blockName(outop, arch)))
	if t[0] == "nil" {
		rr.add(stmtf("b.SetControl(nil)"))
	} else {
		if pos == "" {
			pos = "v.Pos"
		}
		v := genResult0(rr, arch, t[0], false, false, pos)
		rr.add(stmtf("b.SetControl(%s)", v))
	}
	if aux != "" {
		rr.add(stmtf("b.Aux = %s", aux))
	} else {
		rr.add(stmtf("b.Aux = nil"))
	}

	succChanged := false
	for i := 0; i < len(succs); i++ {
		if succs[i] != newsuccs[i] {
			succChanged = true
		}
	}
	if succChanged {
		if len(succs) != 2 {
			log.Fatalf("changed successors, len!=2 in %s", rule)
		}
		if succs[0] != newsuccs[1] || succs[1] != newsuccs[0] {
			log.Fatalf("can only handle swapped successors in %s", rule)
		}
		rr.add(stmtf("b.swapSuccessors()"))
	}

	if *genLog {
		rr.add(stmtf("logRule(%q)", rule.loc))
	}
	return rr
}

// genMatch returns the variable whose source position should be used for the
// result (or "" if no opinion), and a boolean that reports whether the match can fail.
func genMatch(rr *RuleRewrite, arch arch, match string) (pos, checkOp string) {
	return genMatch0(rr, arch, match, "v")
}

func genMatch0(rr *RuleRewrite, arch arch, match, v string) (pos, checkOp string) {
	if match[0] != '(' || match[len(match)-1] != ')' {
		log.Fatalf("non-compound expr in genMatch0: %q", match)
	}
	op, oparch, typ, auxint, aux, args := parseValue(match, arch, rr.loc)

	checkOp = fmt.Sprintf("Op%s%s", oparch, op.name)

	if op.faultOnNilArg0 || op.faultOnNilArg1 {
		// Prefer the position of an instruction which could fault.
		pos = v + ".Pos"
	}

	if typ != "" {
		if !token.IsIdentifier(typ) || rr.declared(typ) {
			// code or variable
			rr.add(breakf("%s.Type != %s", v, typ))
		} else {
			rr.add(declf(typ, "%s.Type", v))
		}
	}
	if auxint != "" {
		if !token.IsIdentifier(auxint) || rr.declared(auxint) {
			// code or variable
			rr.add(breakf("%s.AuxInt != %s", v, auxint))
		} else {
			rr.add(declf(auxint, "%s.AuxInt", v))
		}
	}
	if aux != "" {
		if !token.IsIdentifier(aux) || rr.declared(aux) {
			// code or variable
			rr.add(breakf("%s.Aux != %s", v, aux))
		} else {
			rr.add(declf(aux, "%s.Aux", v))
		}
	}

	// Access last argument first to minimize bounds checks.
	if n := len(args); n > 1 {
		a := args[n-1]
		if a != "_" && !rr.declared(a) && token.IsIdentifier(a) {
			rr.add(declf(a, "%s.Args[%d]", v, n-1))

			// delete the last argument so it is not reprocessed
			args = args[:n-1]
		} else {
			rr.add(stmtf("_ = %s.Args[%d]", v, n-1))
		}
	}
	for i, arg := range args {
		if arg == "_" {
			continue
		}
		if !strings.Contains(arg, "(") {
			// leaf variable
			if rr.declared(arg) {
				// variable already has a definition. Check whether
				// the old definition and the new definition match.
				// For example, (add x x).  Equality is just pointer equality
				// on Values (so cse is important to do before lowering).
				rr.add(breakf("%s != %s.Args[%d]", arg, v, i))
			} else {
				rr.add(declf(arg, "%s.Args[%d]", v, i))
			}
			continue
		}
		// compound sexpr
		argname := fmt.Sprintf("%s_%d", v, i)
		colon := strings.Index(arg, ":")
		openparen := strings.Index(arg, "(")
		if colon >= 0 && openparen >= 0 && colon < openparen {
			// rule-specified name
			argname = arg[:colon]
			arg = arg[colon+1:]
		}
		if argname == "b" {
			log.Fatalf("don't name args 'b', it is ambiguous with blocks")
		}

		rr.add(declf(argname, "%s.Args[%d]", v, i))
		bexpr := exprf("%s.Op != addLater", argname)
		rr.add(&CondBreak{expr: bexpr})
		rr.canFail = true // since we're not using breakf
		argPos, argCheckOp := genMatch0(rr, arch, arg, argname)
		bexpr.(*ast.BinaryExpr).Y.(*ast.Ident).Name = argCheckOp

		if argPos != "" {
			// Keep the argument in preference to the parent, as the
			// argument is normally earlier in program flow.
			// Keep the argument in preference to an earlier argument,
			// as that prefers the memory argument which is also earlier
			// in the program flow.
			pos = argPos
		}
	}

	if op.argLength == -1 {
		rr.add(breakf("len(%s.Args) != %d", v, len(args)))
	}
	return pos, checkOp
}

func genResult(rr *RuleRewrite, arch arch, result, pos string) {
	move := result[0] == '@'
	if move {
		// parse @block directive
		s := strings.SplitN(result[1:], " ", 2)
		rr.add(stmtf("b = %s", s[0]))
		result = s[1]
	}
	genResult0(rr, arch, result, true, move, pos)
}

func genResult0(rr *RuleRewrite, arch arch, result string, top, move bool, pos string) string {
	// TODO: when generating a constant result, use f.constVal to avoid
	// introducing copies just to clean them up again.
	if result[0] != '(' {
		// variable
		if top {
			// It in not safe in general to move a variable between blocks
			// (and particularly not a phi node).
			// Introduce a copy.
			rr.add(stmtf("v.reset(OpCopy)"))
			rr.add(stmtf("v.Type = %s.Type", result))
			rr.add(stmtf("v.AddArg(%s)", result))
		}
		return result
	}

	op, oparch, typ, auxint, aux, args := parseValue(result, arch, rr.loc)

	// Find the type of the variable.
	typeOverride := typ != ""
	if typ == "" && op.typ != "" {
		typ = typeName(op.typ)
	}

	v := "v"
	if top && !move {
		rr.add(stmtf("v.reset(Op%s%s)", oparch, op.name))
		if typeOverride {
			rr.add(stmtf("v.Type = %s", typ))
		}
	} else {
		if typ == "" {
			log.Fatalf("sub-expression %s (op=Op%s%s) at %s must have a type", result, oparch, op.name, rr.loc)
		}
		v = fmt.Sprintf("v%d", rr.alloc)
		rr.alloc++
		rr.add(declf(v, "b.NewValue0(%s, Op%s%s, %s)", pos, oparch, op.name, typ))
		if move && top {
			// Rewrite original into a copy
			rr.add(stmtf("v.reset(OpCopy)"))
			rr.add(stmtf("v.AddArg(%s)", v))
		}
	}

	if auxint != "" {
		rr.add(stmtf("%s.AuxInt = %s", v, auxint))
	}
	if aux != "" {
		rr.add(stmtf("%s.Aux = %s", v, aux))
	}
	for _, arg := range args {
		x := genResult0(rr, arch, arg, false, move, pos)
		rr.add(stmtf("%s.AddArg(%s)", v, x))
	}

	return v
}

func split(s string) []string {
	var r []string

outer:
	for s != "" {
		d := 0               // depth of ({[<
		var open, close byte // opening and closing markers ({[< or )}]>
		nonsp := false       // found a non-space char so far
		for i := 0; i < len(s); i++ {
			switch {
			case d == 0 && s[i] == '(':
				open, close = '(', ')'
				d++
			case d == 0 && s[i] == '<':
				open, close = '<', '>'
				d++
			case d == 0 && s[i] == '[':
				open, close = '[', ']'
				d++
			case d == 0 && s[i] == '{':
				open, close = '{', '}'
				d++
			case d == 0 && (s[i] == ' ' || s[i] == '\t'):
				if nonsp {
					r = append(r, strings.TrimSpace(s[:i]))
					s = s[i:]
					continue outer
				}
			case d > 0 && s[i] == open:
				d++
			case d > 0 && s[i] == close:
				d--
			default:
				nonsp = true
			}
		}
		if d != 0 {
			log.Fatalf("imbalanced expression: %q", s)
		}
		if nonsp {
			r = append(r, strings.TrimSpace(s))
		}
		break
	}
	return r
}

// isBlock reports whether this op is a block opcode.
func isBlock(name string, arch arch) bool {
	for _, b := range genericBlocks {
		if b.name == name {
			return true
		}
	}
	for _, b := range arch.blocks {
		if b.name == name {
			return true
		}
	}
	return false
}

func extract(val string) (op, typ, auxint, aux string, args []string) {
	val = val[1 : len(val)-1] // remove ()

	// Split val up into regions.
	// Split by spaces/tabs, except those contained in (), {}, [], or <>.
	s := split(val)

	// Extract restrictions and args.
	op = s[0]
	for _, a := range s[1:] {
		switch a[0] {
		case '<':
			typ = a[1 : len(a)-1] // remove <>
		case '[':
			auxint = a[1 : len(a)-1] // remove []
		case '{':
			aux = a[1 : len(a)-1] // remove {}
		default:
			args = append(args, a)
		}
	}
	return
}

// parseValue parses a parenthesized value from a rule.
// The value can be from the match or the result side.
// It returns the op and unparsed strings for typ, auxint, and aux restrictions and for all args.
// oparch is the architecture that op is located in, or "" for generic.
func parseValue(val string, arch arch, loc string) (op opData, oparch, typ, auxint, aux string, args []string) {
	// Resolve the op.
	var s string
	s, typ, auxint, aux, args = extract(val)

	// match reports whether x is a good op to select.
	// If strict is true, rule generation might succeed.
	// If strict is false, rule generation has failed,
	// but we're trying to generate a useful error.
	// Doing strict=true then strict=false allows
	// precise op matching while retaining good error messages.
	match := func(x opData, strict bool, archname string) bool {
		if x.name != s {
			return false
		}
		if x.argLength != -1 && int(x.argLength) != len(args) {
			if strict {
				return false
			}
			log.Printf("%s: op %s (%s) should have %d args, has %d", loc, s, archname, x.argLength, len(args))
		}
		return true
	}

	for _, x := range genericOps {
		if match(x, true, "generic") {
			op = x
			break
		}
	}
	for _, x := range arch.ops {
		if arch.name != "generic" && match(x, true, arch.name) {
			if op.name != "" {
				log.Fatalf("%s: matches for op %s found in both generic and %s", loc, op.name, arch.name)
			}
			op = x
			oparch = arch.name
			break
		}
	}

	if op.name == "" {
		// Failed to find the op.
		// Run through everything again with strict=false
		// to generate useful diagnosic messages before failing.
		for _, x := range genericOps {
			match(x, false, "generic")
		}
		for _, x := range arch.ops {
			match(x, false, arch.name)
		}
		log.Fatalf("%s: unknown op %s", loc, s)
	}

	// Sanity check aux, auxint.
	if auxint != "" {
		switch op.aux {
		case "Bool", "Int8", "Int16", "Int32", "Int64", "Int128", "Float32", "Float64", "SymOff", "SymValAndOff", "SymInt32", "TypSize":
		default:
			log.Fatalf("%s: op %s %s can't have auxint", loc, op.name, op.aux)
		}
	}
	if aux != "" {
		switch op.aux {
		case "String", "Sym", "SymOff", "SymValAndOff", "SymInt32", "Typ", "TypSize", "CCop":
		default:
			log.Fatalf("%s: op %s %s can't have aux", loc, op.name, op.aux)
		}
	}
	return
}

func blockName(name string, arch arch) string {
	for _, b := range genericBlocks {
		if b.name == name {
			return "Block" + name
		}
	}
	return "Block" + arch.name + name
}

// typeName returns the string to use to generate a type.
func typeName(typ string) string {
	if typ[0] == '(' {
		ts := strings.Split(typ[1:len(typ)-1], ",")
		if len(ts) != 2 {
			log.Fatalf("Tuple expect 2 arguments")
		}
		return "types.NewTuple(" + typeName(ts[0]) + ", " + typeName(ts[1]) + ")"
	}
	switch typ {
	case "Flags", "Mem", "Void", "Int128":
		return "types.Type" + typ
	default:
		return "typ." + typ
	}
}

// unbalanced reports whether there aren't the same number of ( and ) in the string.
func unbalanced(s string) bool {
	balance := 0
	for _, c := range s {
		if c == '(' {
			balance++
		} else if c == ')' {
			balance--
		}
	}
	return balance != 0
}

// findAllOpcode is a function to find the opcode portion of s-expressions.
var findAllOpcode = regexp.MustCompile(`[(](\w+[|])+\w+[)]`).FindAllStringIndex

// excludeFromExpansion reports whether the substring s[idx[0]:idx[1]] in a rule
// should be disregarded as a candidate for | expansion.
// It uses simple syntactic checks to see whether the substring
// is inside an AuxInt expression or inside the && conditions.
func excludeFromExpansion(s string, idx []int) bool {
	left := s[:idx[0]]
	if strings.LastIndexByte(left, '[') > strings.LastIndexByte(left, ']') {
		// Inside an AuxInt expression.
		return true
	}
	right := s[idx[1]:]
	if strings.Contains(left, "&&") && strings.Contains(right, "->") {
		// Inside && conditions.
		return true
	}
	return false
}

// expandOr converts a rule into multiple rules by expanding | ops.
func expandOr(r string) []string {
	// Find every occurrence of |-separated things.
	// They look like MOV(B|W|L|Q|SS|SD)load or MOV(Q|L)loadidx(1|8).
	// Generate rules selecting one case from each |-form.

	// Count width of |-forms.  They must match.
	n := 1
	for _, idx := range findAllOpcode(r, -1) {
		if excludeFromExpansion(r, idx) {
			continue
		}
		s := r[idx[0]:idx[1]]
		c := strings.Count(s, "|") + 1
		if c == 1 {
			continue
		}
		if n > 1 && n != c {
			log.Fatalf("'|' count doesn't match in %s: both %d and %d\n", r, n, c)
		}
		n = c
	}
	if n == 1 {
		// No |-form in this rule.
		return []string{r}
	}
	// Build each new rule.
	res := make([]string, n)
	for i := 0; i < n; i++ {
		buf := new(strings.Builder)
		x := 0
		for _, idx := range findAllOpcode(r, -1) {
			if excludeFromExpansion(r, idx) {
				continue
			}
			buf.WriteString(r[x:idx[0]])              // write bytes we've skipped over so far
			s := r[idx[0]+1 : idx[1]-1]               // remove leading "(" and trailing ")"
			buf.WriteString(strings.Split(s, "|")[i]) // write the op component for this rule
			x = idx[1]                                // note that we've written more bytes
		}
		buf.WriteString(r[x:])
		res[i] = buf.String()
	}
	return res
}

// commute returns all equivalent rules to r after applying all possible
// argument swaps to the commutable ops in r.
// Potentially exponential, be careful.
func commute(r string, arch arch) []string {
	match, cond, result := Rule{rule: r}.parse()
	a := commute1(match, varCount(match), arch)
	for i, m := range a {
		if cond != "" {
			m += " && " + cond
		}
		m += " -> " + result
		a[i] = m
	}
	if len(a) == 1 && normalizeWhitespace(r) != normalizeWhitespace(a[0]) {
		fmt.Println(normalizeWhitespace(r))
		fmt.Println(normalizeWhitespace(a[0]))
		log.Fatalf("commute() is not the identity for noncommuting rule")
	}
	if false && len(a) > 1 {
		fmt.Println(r)
		for _, x := range a {
			fmt.Println("  " + x)
		}
	}
	return a
}

func commute1(m string, cnt map[string]int, arch arch) []string {
	if m[0] == '<' || m[0] == '[' || m[0] == '{' || token.IsIdentifier(m) {
		return []string{m}
	}
	// Split up input.
	var prefix string
	if i := strings.Index(m, ":"); i >= 0 && token.IsIdentifier(m[:i]) {
		prefix = m[:i+1]
		m = m[i+1:]
	}
	if m[0] != '(' || m[len(m)-1] != ')' {
		log.Fatalf("non-compound expr in commute1: %q", m)
	}
	s := split(m[1 : len(m)-1])
	op := s[0]

	// Figure out if the op is commutative or not.
	commutative := false
	for _, x := range genericOps {
		if op == x.name {
			if x.commutative {
				commutative = true
			}
			break
		}
	}
	if arch.name != "generic" {
		for _, x := range arch.ops {
			if op == x.name {
				if x.commutative {
					commutative = true
				}
				break
			}
		}
	}
	var idx0, idx1 int
	if commutative {
		// Find indexes of two args we can swap.
		for i, arg := range s {
			if i == 0 || arg[0] == '<' || arg[0] == '[' || arg[0] == '{' {
				continue
			}
			if idx0 == 0 {
				idx0 = i
				continue
			}
			if idx1 == 0 {
				idx1 = i
				break
			}
		}
		if idx1 == 0 {
			log.Fatalf("couldn't find first two args of commutative op %q", s[0])
		}
		if cnt[s[idx0]] == 1 && cnt[s[idx1]] == 1 || s[idx0] == s[idx1] && cnt[s[idx0]] == 2 {
			// When we have (Add x y) with no other uses of x and y in the matching rule,
			// then we can skip the commutative match (Add y x).
			commutative = false
		}
	}

	// Recursively commute arguments.
	a := make([][]string, len(s))
	for i, arg := range s {
		a[i] = commute1(arg, cnt, arch)
	}

	// Choose all possibilities from all args.
	r := crossProduct(a)

	// If commutative, do that again with its two args reversed.
	if commutative {
		a[idx0], a[idx1] = a[idx1], a[idx0]
		r = append(r, crossProduct(a)...)
	}

	// Construct result.
	for i, x := range r {
		r[i] = prefix + "(" + x + ")"
	}
	return r
}

// varCount returns a map which counts the number of occurrences of
// Value variables in m.
func varCount(m string) map[string]int {
	cnt := map[string]int{}
	varCount1(m, cnt)
	return cnt
}

func varCount1(m string, cnt map[string]int) {
	if m[0] == '<' || m[0] == '[' || m[0] == '{' {
		return
	}
	if token.IsIdentifier(m) {
		cnt[m]++
		return
	}
	// Split up input.
	if i := strings.Index(m, ":"); i >= 0 && token.IsIdentifier(m[:i]) {
		cnt[m[:i]]++
		m = m[i+1:]
	}
	if m[0] != '(' || m[len(m)-1] != ')' {
		log.Fatalf("non-compound expr in commute1: %q", m)
	}
	s := split(m[1 : len(m)-1])
	for _, arg := range s[1:] {
		varCount1(arg, cnt)
	}
}

// crossProduct returns all possible values
// x[0][i] + " " + x[1][j] + " " + ... + " " + x[len(x)-1][k]
// for all valid values of i, j, ..., k.
func crossProduct(x [][]string) []string {
	if len(x) == 1 {
		return x[0]
	}
	var r []string
	for _, tail := range crossProduct(x[1:]) {
		for _, first := range x[0] {
			r = append(r, first+" "+tail)
		}
	}
	return r
}

// normalizeWhitespace replaces 2+ whitespace sequences with a single space.
func normalizeWhitespace(x string) string {
	x = strings.Join(strings.Fields(x), " ")
	x = strings.Replace(x, "( ", "(", -1)
	x = strings.Replace(x, " )", ")", -1)
	x = strings.Replace(x, ")->", ") ->", -1)
	return x
}
