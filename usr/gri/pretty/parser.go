// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package Parser

import Scanner "scanner"
import AST "ast"


export type Parser struct {
	verbose bool;
	indent uint;
	scanner *Scanner.Scanner;
	tokchan *<-chan *Scanner.Token;
	
	// Token
	tok int;  // one token look-ahead
	pos int;  // token source position
	val string;  // token value (for IDENT, NUMBER, STRING only)

	// Nesting level
	level int;  // 0 = global scope, -1 = function/struct scope of global functions/structs, etc.
};


// ----------------------------------------------------------------------------
// Support functions

func (P *Parser) PrintIndent() {
	for i := P.indent; i > 0; i-- {
		print(". ");
	}
}


func (P *Parser) Trace(msg string) {
	if P.verbose {
		P.PrintIndent();
		print(msg, " {\n");
	}
	P.indent++;  // always, so proper identation is always checked
}


func (P *Parser) Ecart() {
	P.indent--;  // always, so proper identation is always checked
	if P.verbose {
		P.PrintIndent();
		print("}\n");
	}
}


func (P *Parser) Next() {
	if P.tokchan == nil {
		P.tok, P.pos, P.val = P.scanner.Scan();
	} else {
		t := <-P.tokchan;
		P.tok, P.pos, P.val = t.tok, t.pos, t.val;
	}
	if P.verbose {
		P.PrintIndent();
		print("[", P.pos, "] ", Scanner.TokenName(P.tok), "\n");
	}
}


func (P *Parser) Open(verbose bool, scanner *Scanner.Scanner, tokchan *<-chan *Scanner.Token) {
	P.verbose = verbose;
	P.indent = 0;
	P.scanner = scanner;
	P.tokchan = tokchan;
	P.Next();
	P.level = 0;
}


func (P *Parser) Error(pos int, msg string) {
	P.scanner.Error(pos, msg);
}


func (P *Parser) Expect(tok int) {
	if P.tok != tok {
		P.Error(P.pos, "expected '" + Scanner.TokenName(tok) + "', found '" + Scanner.TokenName(P.tok) + "'");
	}
	P.Next();  // make progress in any case
}


func (P *Parser) Optional(tok int) {
	if P.tok == tok {
		P.Next();
	}
}


// ----------------------------------------------------------------------------
// Scopes

func (P *Parser) OpenScope() {
}


func (P *Parser) CloseScope() {
}


// ----------------------------------------------------------------------------
// Common productions

func (P *Parser) TryType() (AST.Type, bool);
func (P *Parser) ParseExpression() AST.Expr;
func (P *Parser) TryStatement() (AST.Stat, bool);
func (P *Parser) ParseDeclaration() AST.Node;


func (P *Parser) ParseIdent() *AST.Ident {
	P.Trace("Ident");

	ident := new(AST.Ident);
	ident.pos, ident.val = P.pos, "";
	if P.tok == Scanner.IDENT {
		ident.val = P.val;
		if P.verbose {
			P.PrintIndent();
			print("Ident = \"", ident.val, "\"\n");
		}
		P.Next();
	} else {
		P.Expect(Scanner.IDENT);  // use Expect() error handling
	}
	
	P.Ecart();
	return ident;
}


func (P *Parser) ParseIdentList() *AST.List {
	P.Trace("IdentList");

	list := AST.NewList();
	list.Add(P.ParseIdent());
	for P.tok == Scanner.COMMA {
		P.Next();
		list.Add(P.ParseIdent());
	}

	P.Ecart();
	return list;
}


func (P *Parser) ParseQualifiedIdent() AST.Expr {
	P.Trace("QualifiedIdent");

	var x AST.Expr = P.ParseIdent();
	if P.tok == Scanner.PERIOD {
		pos := P.pos;
		P.Next();
		y := P.ParseIdent();

		z := new(AST.Selector);
		z.pos, z.x, z.field = pos, x, y.val;
		x = z;
	}
	
	P.Ecart();
	return x;
}


// ----------------------------------------------------------------------------
// Types

func (P *Parser) ParseType() AST.Type {
	P.Trace("Type");
	
	typ, ok := P.TryType();
	if !ok {
		P.Error(P.pos, "type expected");
	}
	
	P.Ecart();
	return typ;
}


func (P *Parser) ParseVarType() AST.Type {
	P.Trace("VarType");
	
	typ := P.ParseType();
	
	P.Ecart();
	return typ;
}


func (P *Parser) ParseTypeName() AST.Type {
	P.Trace("TypeName");
	
	typ := P.ParseQualifiedIdent();

	P.Ecart();
	return typ;
}


func (P *Parser) ParseArrayType() *AST.ArrayType {
	P.Trace("ArrayType");
	
	typ := new(AST.ArrayType);
	typ.pos = P.pos;
	
	P.Expect(Scanner.LBRACK);
	if P.tok != Scanner.RBRACK {
		// TODO set typ.len
		typ.len_ = P.ParseExpression();
	}
	P.Expect(Scanner.RBRACK);
	typ.elt = P.ParseType();

	P.Ecart();
	return typ;
}


func (P *Parser) ParseChannelType() *AST.ChannelType {
	P.Trace("ChannelType");
	
	typ := new(AST.ChannelType);
	typ.pos = P.pos;
	
	if P.tok == Scanner.CHAN {
		P.Next();
		if P.tok == Scanner.ARROW {
			P.Next();
		}
	} else {
		P.Expect(Scanner.ARROW);
		P.Expect(Scanner.CHAN);
	}
	typ.elt = P.ParseVarType();

	P.Ecart();
	return typ;
}


func (P *Parser) ParseVarDeclList() *AST.VarDeclList {
	P.Trace("VarDeclList");
	
	res := new(AST.VarDeclList);
	res.idents = P.ParseIdentList();
	res.typ = P.ParseVarType();
	
	P.Ecart();
	return res;
}


// Returns a list of AST.VarDeclList
func (P *Parser) ParseParameterList() *AST.List {
	P.Trace("ParameterList");
	
	list := AST.NewList();
	list.Add(P.ParseVarDeclList());
	for P.tok == Scanner.COMMA {
		P.Next();
		list.Add(P.ParseVarDeclList());
	}
	
	P.Ecart();
	return list;
}


// Returns a list of AST.VarDeclList
func (P *Parser) ParseParameters() *AST.List {
	P.Trace("Parameters");
	
	var list *AST.List;
	P.Expect(Scanner.LPAREN);
	if P.tok != Scanner.RPAREN {
		list = P.ParseParameterList();
	}
	P.Expect(Scanner.RPAREN);
	
	P.Ecart();
	return list;
}


func (P *Parser) ParseResultList() {
	P.Trace("ResultList");

	P.ParseType();
	for P.tok == Scanner.COMMA {
		P.Next();
		P.ParseType();
	}
	if P.tok != Scanner.RPAREN {
		P.ParseType();
	}

	P.Ecart();
}


func (P *Parser) ParseResult() *AST.List {
	P.Trace("Result");
	
	if P.tok == Scanner.LPAREN {
		P.Next();
		P.ParseResultList();
		for P.tok == Scanner.COMMA {
			P.Next();
			P.ParseResultList();
		}
		P.Expect(Scanner.RPAREN);

	} else {
		// anonymous result
		P.TryType();
	}

	P.Ecart();
	return nil
}


// Function types
//
// (params)
// (params) type
// (params) (results)

func (P *Parser) ParseFunctionType() *AST.FunctionType {
	P.Trace("FunctionType");
	
	P.OpenScope();
	P.level--;

	typ := new(AST.FunctionType);
	typ.pos = P.pos;
	typ.params = P.ParseParameters();
	typ.result = P.ParseResult();

	P.level++;
	P.CloseScope();
	
	P.Ecart();
	return typ;
}


func (P *Parser) ParseMethodDecl() {
	P.Trace("MethodDecl");
	
	ident := P.ParseIdent();
	P.ParseFunctionType();
	P.Optional(Scanner.SEMICOLON);
	
	P.Ecart();
}


func (P *Parser) ParseInterfaceType() *AST.InterfaceType {
	P.Trace("InterfaceType");
	
	P.Expect(Scanner.INTERFACE);
	P.Expect(Scanner.LBRACE);
	P.OpenScope();
	P.level--;
	for P.tok >= Scanner.IDENT {
		P.ParseMethodDecl();
	}
	P.level++;
	P.CloseScope();
	P.Expect(Scanner.RBRACE);
	
	P.Ecart();
	return nil;
}


func (P *Parser) ParseMapType() *AST.MapType {
	P.Trace("MapType");
	
	typ := new(AST.MapType);
	typ.pos = P.pos;
	
	P.Expect(Scanner.MAP);
	P.Expect(Scanner.LBRACK);
	typ.key = P.ParseVarType();
	P.Expect(Scanner.RBRACK);
	typ.val = P.ParseVarType();
	
	P.Ecart();
	return typ;
}


func (P *Parser) ParseStructType() *AST.StructType {
	P.Trace("StructType");
	
	typ := new(AST.StructType);
	typ.pos = P.pos;
	typ.fields = AST.NewList();
	
	P.Expect(Scanner.STRUCT);
	P.Expect(Scanner.LBRACE);
	P.OpenScope();
	P.level--;
	for P.tok >= Scanner.IDENT {
		typ.fields.Add(P.ParseVarDeclList());
		if P.tok != Scanner.RBRACE {
			P.Expect(Scanner.SEMICOLON);
		}
	}
	P.Optional(Scanner.SEMICOLON);
	P.level++;
	P.CloseScope();
	P.Expect(Scanner.RBRACE);
	
	P.Ecart();
	return typ;
}


func (P *Parser) ParsePointerType() *AST.PointerType {
	P.Trace("PointerType");
	
	typ := new(AST.PointerType);
	typ.pos = P.pos;
	
	P.Expect(Scanner.MUL);
	typ.base = P.ParseType();
	
	P.Ecart();
	return typ;
}


// Returns false if no type was found.
func (P *Parser) TryType() (AST.Type, bool) {
	P.Trace("Type (try)");
	
	var typ AST.Type = AST.NIL;
	found := true;
	switch P.tok {
	case Scanner.IDENT: typ = P.ParseTypeName();
	case Scanner.LBRACK: typ = P.ParseArrayType();
	case Scanner.CHAN, Scanner.ARROW: typ = P.ParseChannelType();
	case Scanner.INTERFACE: typ = P.ParseInterfaceType();
	case Scanner.LPAREN: typ = P.ParseFunctionType();
	case Scanner.MAP: typ = P.ParseMapType();
	case Scanner.STRUCT: typ = P.ParseStructType();
	case Scanner.MUL: typ = P.ParsePointerType();
	default: found = false;
	}

	P.Ecart();
	return typ, found;
}


// ----------------------------------------------------------------------------
// Blocks

func (P *Parser) ParseStatement() AST.Stat {
	P.Trace("Statement");
	
	stat, ok := P.TryStatement();
	if !ok {
		P.Error(P.pos, "statement expected");
		P.Next();  // make progress
	}
	P.Ecart();
	
	return stat;
}


func (P *Parser) ParseStatementList() *AST.List {
	P.Trace("StatementList");
	
	stats := AST.NewList();
	for {
		stat, ok := P.TryStatement();
		if ok {
			stats.Add(stat);
			P.Optional(Scanner.SEMICOLON);
		} else {
			break;
		}
	}
	
	P.Ecart();
	return stats;
}


func (P *Parser) ParseBlock() *AST.Block {
	P.Trace("Block");
	
	pos := P.pos;
	P.Expect(Scanner.LBRACE);
	P.OpenScope();
	
	var stats *AST.List;
	if P.tok != Scanner.RBRACE && P.tok != Scanner.SEMICOLON {
		stats = P.ParseStatementList();
	}
	P.Optional(Scanner.SEMICOLON);
	P.CloseScope();
	P.Expect(Scanner.RBRACE);
	
	P.Ecart();
	
	x := new(AST.Block);
	x.pos, x.stats = pos, stats;
	return x;
}


// ----------------------------------------------------------------------------
// Expressions

func (P *Parser) ParseExpressionList() *AST.List {
	P.Trace("ExpressionList");

	p := AST.NewList();
	p.Add(P.ParseExpression());
	for P.tok == Scanner.COMMA {
		P.Next();
		p.Add(P.ParseExpression());
	}
	
	P.Ecart();
	return p;
}


func (P *Parser) ParseFunctionLit() AST.Expr {
	P.Trace("FunctionLit");
	
	P.Expect(Scanner.FUNC);
	P.ParseFunctionType();
	P.ParseBlock();
	
	P.Ecart();
	var x AST.Expr;
	return x;
}


func (P *Parser) ParseExpressionPair() AST.Expr {
	P.Trace("ExpressionPair");

	x := P.ParseExpression();
	pos := P.pos;
	P.Expect(Scanner.COLON);
	y := P.ParseExpression();
	
	z := new(AST.Pair);
	z.pos, z.x, z.y = pos, x, y;
	
	P.Ecart();
	return z;
}


func (P *Parser) ParseExpressionPairList() *AST.List {
	P.Trace("ExpressionPairList");

	p := AST.NewList();
	p.Add(P.ParseExpressionPair());
	for P.tok == Scanner.COMMA {
		p.Add(P.ParseExpressionPair());
	}
	
	P.Ecart();
	return p;
}


func (P *Parser) ParseCompositeLit() AST.Expr {
	P.Trace("CompositeLit");
	
	P.Expect(Scanner.LBRACE);
	// TODO: should allow trailing ','
	if P.tok != Scanner.RBRACE {
		P.ParseExpression();
		if P.tok == Scanner.COMMA {
			P.Next();
			if P.tok != Scanner.RBRACE {
				P.ParseExpressionList();
			}
		} else if P.tok == Scanner.COLON {
			P.Next();
			P.ParseExpression();
			if P.tok == Scanner.COMMA {
				P.Next();
				if P.tok != Scanner.RBRACE {
					P.ParseExpressionPairList();
				}
			}
		}
	}
	P.Expect(Scanner.RBRACE);

	P.Ecart();
	return nil;
}


func (P *Parser) ParseOperand() AST.Expr {
	P.Trace("Operand");

	var z AST.Expr;
	switch P.tok {
	case Scanner.IDENT:
		z = P.ParseIdent();
		
	case Scanner.LPAREN:
		P.Next();
		z = P.ParseExpression();
		P.Expect(Scanner.RPAREN);

	case Scanner.INT, Scanner.FLOAT, Scanner.STRING:
		x := new(AST.Literal);
		x.pos, x.tok, x.val = P.pos, P.tok, P.val;
		z = x;
		P.Next();

	case Scanner.FUNC:
		z = P.ParseFunctionLit();
		
	case Scanner.HASH:
		P.Next();
		P.ParseType();
		P.ParseCompositeLit();
		z = nil;

	default:
		if P.tok != Scanner.IDENT {
			typ, ok := P.TryType();
			if ok {
				z = P.ParseCompositeLit();
				break;
			}
		}

		P.Error(P.pos, "operand expected");
		P.Next();  // make progress
	}

	P.Ecart();
	return z;
}


func (P *Parser) ParseSelectorOrTypeGuard(x AST.Expr) AST.Expr {
	P.Trace("SelectorOrTypeGuard");

	pos := P.pos;
	P.Expect(Scanner.PERIOD);
	
	if P.tok == Scanner.IDENT {
		ident := P.ParseIdent();
		
		z := new(AST.Selector);
		z.pos, z.x, z.field = pos, x, ident.val;
		x = z;
		
	} else {
		P.Expect(Scanner.LPAREN);
		P.ParseType();
		P.Expect(Scanner.RPAREN);
	}
	
	P.Ecart();
	return x;
}


func (P *Parser) ParseIndexOrSlice(x AST.Expr) AST.Expr {
	P.Trace("IndexOrSlice");
	
	pos := P.pos;
	P.Expect(Scanner.LBRACK);
	i := P.ParseExpression();
	if P.tok == Scanner.COLON {
		P.Next();
		j := P.ParseExpression();
		// TODO: handle this case
	}
	P.Expect(Scanner.RBRACK);

	z := new(AST.Index);
	z.pos, z.x, z.index = pos, x, i;
	
	P.Ecart();
	return z;
}


func (P *Parser) ParseCall(x AST.Expr) AST.Expr {
	P.Trace("Call");

	pos := P.pos;
	var args *AST.List = nil;
	P.Expect(Scanner.LPAREN);
	if P.tok != Scanner.RPAREN {
	   	// first arguments could be a type if the call is to "new"
		// - exclude type names because they could be expression starts
		// - exclude "("'s because function types are not allowed and they indicate an expression
		// - still a problem for "new(*T)" (the "*")
		// - possibility: make "new" a keyword again (or disallow "*" types in new)
		if P.tok != Scanner.IDENT && P.tok != Scanner.LPAREN {
			typ, ok := P.TryType();
			if ok {
				if P.tok == Scanner.COMMA {
					P.Next();
					if P.tok != Scanner.RPAREN {
						args = P.ParseExpressionList();
					}
				}
			} else {
				args = P.ParseExpressionList();
			}
		} else {
			args = P.ParseExpressionList();
		}
	}
	P.Expect(Scanner.RPAREN);
	
	P.Ecart();
	call := new(AST.Call);
	call.pos, call.fun, call.args = pos, x, args;
	return call;
}


func (P *Parser) ParsePrimaryExpr() AST.Expr {
	P.Trace("PrimaryExpr");
	
	x := P.ParseOperand();
	L: for {
		switch P.tok {
		case Scanner.PERIOD: x = P.ParseSelectorOrTypeGuard(x);
		case Scanner.LBRACK: x = P.ParseIndexOrSlice(x);
		case Scanner.LPAREN: x = P.ParseCall(x);
		default: break L;
		}
	}

	P.Ecart();
	return x;
}


func (P *Parser) ParseUnaryExpr() AST.Expr {
	P.Trace("UnaryExpr");
	
	var x AST.Expr = AST.NIL;
	switch P.tok {
	case
		Scanner.ADD, Scanner.SUB,
		Scanner.NOT, Scanner.XOR,
		Scanner.MUL, Scanner.ARROW,
		Scanner.AND:
			pos, tok := P.pos, P.tok;
			P.Next();
			y := P.ParseUnaryExpr();

			z := new(AST.Unary);
			z.pos, z.tok, z.x = pos, tok, y;
			x = z;
			
		default:
			x = P.ParsePrimaryExpr();
	}
	
	P.Ecart();
	return x;
}


func Precedence(tok int) int {
	// TODO should use a map or array here for lookup
	switch tok {
	case Scanner.LOR:
		return 1;
	case Scanner.LAND:
		return 2;
	case Scanner.ARROW:
		return 3;
	case Scanner.EQL, Scanner.NEQ, Scanner.LSS, Scanner.LEQ, Scanner.GTR, Scanner.GEQ:
		return 4;
	case Scanner.ADD, Scanner.SUB, Scanner.OR, Scanner.XOR:
		return 5;
	case Scanner.MUL, Scanner.QUO, Scanner.REM, Scanner.SHL, Scanner.SHR, Scanner.AND:
		return 6;
	}
	return 0;
}


func (P *Parser) ParseBinaryExpr(prec1 int) AST.Expr {
	P.Trace("BinaryExpr");
	
	x := P.ParseUnaryExpr();
	for prec := Precedence(P.tok); prec >= prec1; prec-- {
		for Precedence(P.tok) == prec {
			pos, tok := P.pos, P.tok;
			P.Next();
			y := P.ParseBinaryExpr(prec + 1);
			
			z := new(AST.Binary);
			z.pos, z.tok, z.x, z.y = pos, tok, x, y;
			x = z;
		}
	}
	
	P.Ecart();
	return x;
}


func (P *Parser) ParseExpression() AST.Expr {
	P.Trace("Expression");
	indent := P.indent;
	
	x := P.ParseBinaryExpr(1);
	
	if indent != P.indent {
		panic("imbalanced tracing code (Expression)");
	}

	P.Ecart();
	return x;
}


// ----------------------------------------------------------------------------
// Statements

func (P *Parser) ParseSimpleStat() AST.Stat {
	P.Trace("SimpleStat");
	
	var stat AST.Stat = AST.NIL;
	x := P.ParseExpressionList();
	
	switch P.tok {
	case Scanner.COLON:
		// label declaration
		P.Next();  // consume ":"
		
	case
		Scanner.DEFINE, Scanner.ASSIGN, Scanner.ADD_ASSIGN,
		Scanner.SUB_ASSIGN, Scanner.MUL_ASSIGN, Scanner.QUO_ASSIGN,
		Scanner.REM_ASSIGN, Scanner.AND_ASSIGN, Scanner.OR_ASSIGN,
		Scanner.XOR_ASSIGN, Scanner.SHL_ASSIGN, Scanner.SHR_ASSIGN:
		pos, tok := P.pos, P.tok;
		P.Next();
		y := P.ParseExpressionList();
		asgn := new(AST.Assignment);
		asgn.pos, asgn.tok, asgn.lhs, asgn.rhs = pos, tok, x, y;
		stat = asgn;
		
	default:
		if P.tok == Scanner.INC || P.tok == Scanner.DEC {
			s := new(AST.IncDecStat);
			s.pos, s.tok = P.pos, P.tok;
			if x.len() == 1 {
				s.expr = x.at(0);
			} else {
				P.Error(P.pos, "more then one operand");
			}
			P.Next();
			stat = s;
		} else {
			xstat := new(AST.ExprStat);
			if x != nil && x.len() > 0 {
				xstat.expr = x.at(0);
			} else {
				// this is a syntax error
				xstat.expr = AST.NIL;
			}
			stat = xstat;
		}
	}
	
	P.Ecart();
	return stat;
}


func (P *Parser) ParseGoStat() {
	P.Trace("GoStat");
	
	P.Expect(Scanner.GO);
	P.ParseExpression();
	
	P.Ecart();
}


func (P *Parser) ParseReturnStat() *AST.ReturnStat {
	P.Trace("ReturnStat");
	
	ret := new(AST.ReturnStat);
	ret.pos = P.pos;
	
	P.Expect(Scanner.RETURN);
	if P.tok != Scanner.SEMICOLON && P.tok != Scanner.RBRACE {
		ret.res = P.ParseExpressionList();
	}
	
	P.Ecart();
	return ret;
}


func (P *Parser) ParseControlFlowStat(tok int) {
	P.Trace("ControlFlowStat");
	
	P.Expect(tok);
	if P.tok == Scanner.IDENT {
		P.ParseIdent();
	}
	
	P.Ecart();
}


func (P *Parser) ParseStatHeader(keyword int) (AST.Stat, AST.Expr, AST.Stat) {
	P.Trace("StatHeader");
	
	var (
		init AST.Stat = AST.NIL;
		expr AST.Expr = AST.NIL;
		post AST.Stat = AST.NIL;
	)
	
	has_init, has_expr, has_post := false, false, false;

	P.Expect(keyword);
	if P.tok != Scanner.LBRACE {
		if P.tok != Scanner.SEMICOLON {
			init = P.ParseSimpleStat();
			has_init = true;
		}
		if P.tok == Scanner.SEMICOLON {
			P.Next();
			if keyword == Scanner.FOR {
				if P.tok != Scanner.SEMICOLON {
					expr = P.ParseExpression();
					has_expr = true;
				}
				P.Expect(Scanner.SEMICOLON);
				if P.tok != Scanner.LBRACE {
					post = P.ParseSimpleStat();
					has_post = true;
				}
			} else {
				if P.tok != Scanner.LBRACE {
					expr = P.ParseExpression();
					has_expr = true;
				}
			}
		}
	}

	P.Ecart();
	return init, expr, post;
}


func (P *Parser) ParseIfStat() *AST.IfStat {
	P.Trace("IfStat");

	x := new(AST.IfStat);
	x.pos = P.pos;
	var dummy AST.Stat;
		
	x.init, x.cond, dummy = P.ParseStatHeader(Scanner.IF);
	
	x.then = P.ParseBlock();
	if P.tok == Scanner.ELSE {
		P.Next();
		b := new(AST.Block);
		b.stats = AST.NewList();
		if P.tok == Scanner.IF {
			b.stats.Add(P.ParseIfStat());
		} else {
			// TODO should be P.ParseBlock()
			b.stats.Add(P.ParseStatement());
		}
		x.else_ = b;
	}
	
	P.Ecart();
	return x;
}


func (P *Parser) ParseForStat() *AST.ForStat {
	P.Trace("ForStat");
	
	stat := new(AST.ForStat);
	stat.pos = P.pos;
	
	P.ParseStatHeader(Scanner.FOR);
	stat.body = P.ParseBlock();
	
	P.Ecart();
	return stat;
}


func (P *Parser) ParseCase() *AST.CaseClause {
	P.Trace("Case");
	
	clause := new(AST.CaseClause);
	clause.pos = P.pos;
	
	if P.tok == Scanner.CASE {
		P.Next();
		clause.exprs = P.ParseExpressionList();
	} else {
		P.Expect(Scanner.DEFAULT);
	}
	P.Expect(Scanner.COLON);
	
	P.Ecart();
	return clause;
}


func (P *Parser) ParseCaseClause() *AST.CaseClause {
	P.Trace("CaseClause");

	clause := P.ParseCase();
	if P.tok != Scanner.FALLTHROUGH && P.tok != Scanner.RBRACE {
		clause.stats = P.ParseStatementList();
		P.Optional(Scanner.SEMICOLON);
	}
	if P.tok == Scanner.FALLTHROUGH {
		P.Next();
		clause.falls = true;
		P.Optional(Scanner.SEMICOLON);
	}
	
	P.Ecart();
	return clause;
}


func (P *Parser) ParseSwitchStat() *AST.SwitchStat {
	P.Trace("SwitchStat");
	
	stat := new(AST.SwitchStat);
	stat.pos = P.pos;
	stat.init = AST.NIL;
	stat.cases = AST.NewList();

	P.ParseStatHeader(Scanner.SWITCH);
	
	P.Expect(Scanner.LBRACE);
	for P.tok == Scanner.CASE || P.tok == Scanner.DEFAULT {
		stat.cases.Add(P.ParseCaseClause());
	}
	P.Expect(Scanner.RBRACE);

	P.Ecart();
	return stat;
}


func (P *Parser) ParseCommCase() {
  P.Trace("CommCase");
  
  if P.tok == Scanner.CASE {
	P.Next();
	if P.tok == Scanner.GTR {
		// send
		P.Next();
		P.ParseExpression();
		P.Expect(Scanner.EQL);
		P.ParseExpression();
	} else {
		// receive
		if P.tok != Scanner.LSS {
			P.ParseIdent();
			P.Expect(Scanner.ASSIGN);
		}
		P.Expect(Scanner.LSS);
		P.ParseExpression();
	}
  } else {
	P.Expect(Scanner.DEFAULT);
  }
  P.Expect(Scanner.COLON);
  
  P.Ecart();
}


func (P *Parser) ParseCommClause() {
	P.Trace("CommClause");
	
	P.ParseCommCase();
	if P.tok != Scanner.CASE && P.tok != Scanner.DEFAULT && P.tok != Scanner.RBRACE {
		P.ParseStatementList();
		P.Optional(Scanner.SEMICOLON);
	}
	
	P.Ecart();
}


func (P *Parser) ParseRangeStat() {
	P.Trace("RangeStat");
	
	P.Expect(Scanner.RANGE);
	P.ParseIdentList();
	P.Expect(Scanner.DEFINE);
	P.ParseExpression();
	P.ParseBlock();
	
	P.Ecart();
}


func (P *Parser) ParseSelectStat() {
	P.Trace("SelectStat");
	
	P.Expect(Scanner.SELECT);
	P.Expect(Scanner.LBRACE);
	for P.tok != Scanner.RBRACE && P.tok != Scanner.EOF {
		P.ParseCommClause();
	}
	P.Next();
	
	P.Ecart();
}


func (P *Parser) TryStatement() (AST.Stat, bool) {
	P.Trace("Statement (try)");
	indent := P.indent;

	var stat AST.Stat = AST.NIL;
	res := true;
	switch P.tok {
	case Scanner.CONST, Scanner.TYPE, Scanner.VAR:
		stat = P.ParseDeclaration();
	case Scanner.FUNC:
		// for now we do not allow local function declarations
		fallthrough;
	case Scanner.MUL, Scanner.ARROW, Scanner.IDENT, Scanner.LPAREN:
		stat = P.ParseSimpleStat();
	case Scanner.GO:
		P.ParseGoStat();
	case Scanner.RETURN:
		stat = P.ParseReturnStat();
	case Scanner.BREAK, Scanner.CONTINUE, Scanner.GOTO:
		P.ParseControlFlowStat(P.tok);
	case Scanner.LBRACE:
		stat = P.ParseBlock();
	case Scanner.IF:
		stat = P.ParseIfStat();
	case Scanner.FOR:
		stat = P.ParseForStat();
	case Scanner.SWITCH:
		stat = P.ParseSwitchStat();
	case Scanner.RANGE:
		P.ParseRangeStat();
	case Scanner.SELECT:
		P.ParseSelectStat();
	default:
		// no statement found
		res = false;
	}

	if indent != P.indent {
		panic("imbalanced tracing code (Statement)");
	}
	P.Ecart();
	return stat, res;
}


// ----------------------------------------------------------------------------
// Declarations

func (P *Parser) ParseImportSpec() {
	P.Trace("ImportSpec");
	
	if P.tok == Scanner.PERIOD {
		P.Error(P.pos, `"import ." not yet handled properly`);
		P.Next();
	} else if P.tok == Scanner.IDENT {
		P.ParseIdent();
	}
	
	if P.tok == Scanner.STRING {
		// TODO eventually the scanner should strip the quotes
		P.Next();
	} else {
		P.Expect(Scanner.STRING);  // use Expect() error handling
	}
	
	P.Ecart();
}


func (P *Parser) ParseConstSpec(exported bool) AST.Decl {
	P.Trace("ConstSpec");
	
	decl := new(AST.ConstDecl);
	decl.ident = P.ParseIdent();
	var ok bool;
	decl.typ, ok = P.TryType();
	
	if P.tok == Scanner.ASSIGN {
		P.Next();
		decl.val = P.ParseExpression();
	}
	
	P.Ecart();
	return decl;
}


func (P *Parser) ParseTypeSpec(exported bool) AST.Decl {
	P.Trace("TypeSpec");

	decl := new(AST.TypeDecl);
	decl.ident = P.ParseIdent();
	decl.typ = P.ParseType();
	
	P.Ecart();
	return decl;
}


func (P *Parser) ParseVarSpec(exported bool) AST.Decl {
	P.Trace("VarSpec");
	
	decl := new(AST.VarDecl);
	decl.idents = P.ParseIdentList();
	if P.tok == Scanner.ASSIGN {
		P.Next();
		decl.vals = P.ParseExpressionList();
	} else {
		decl.typ = P.ParseVarType();
		if P.tok == Scanner.ASSIGN {
			P.Next();
			decl.vals = P.ParseExpressionList();
		}
	}
	
	P.Ecart();
	return decl;
}


// TODO With method variables, we wouldn't need this dispatch function.
func (P *Parser) ParseSpec(exported bool, keyword int) AST.Decl {
	var decl AST.Decl = AST.NIL;
	switch keyword {
	case Scanner.IMPORT: P.ParseImportSpec();
	case Scanner.CONST: decl = P.ParseConstSpec(exported);
	case Scanner.TYPE: decl = P.ParseTypeSpec(exported);
	case Scanner.VAR: decl = P.ParseVarSpec(exported);
	default: panic("UNREACHABLE");
	}
	return decl;
}


func (P *Parser) ParseDecl(exported bool, keyword int) *AST.Declaration {
	P.Trace("Decl");
	
	decl := new(AST.Declaration);
	decl.decls = AST.NewList();
	decl.pos, decl.tok = P.pos, P.tok;
	
	P.Expect(keyword);
	if P.tok == Scanner.LPAREN {
		P.Next();
		for P.tok != Scanner.RPAREN {
			decl.decls.Add(P.ParseSpec(exported, keyword));
			if P.tok != Scanner.RPAREN {
				// P.Expect(Scanner.SEMICOLON);
				P.Optional(Scanner.SEMICOLON);  // TODO this seems wrong! (needed for math.go)
			}
		}
		P.Next();  // consume ")"
	} else {
		decl.decls.Add(P.ParseSpec(exported, keyword));
	}
	
	P.Ecart();
	return decl;
}


// Function declarations
//
// func        ident (params)
// func        ident (params) type
// func        ident (params) (results)
// func (recv) ident (params)
// func (recv) ident (params) type
// func (recv) ident (params) (results)

func (P *Parser) ParseFuncDecl(exported bool) *AST.FuncDecl {
	P.Trace("FuncDecl");
	
	fun := new(AST.FuncDecl);
	fun.pos = P.pos;

	P.Expect(Scanner.FUNC);

	P.OpenScope();
	P.level--;

	var recv *AST.VarDeclList;
	if P.tok == Scanner.LPAREN {
		recv_pos := P.pos;
		recv := P.ParseParameters().at(0);
		/*
		if n != 1 {
			P.Error(recv_pos, "must have exactly one receiver");
		}
		*/
	}
	
	fun.ident = P.ParseIdent();
	fun.typ = P.ParseFunctionType();
	fun.typ.recv = recv;
	
	P.level++;
	P.CloseScope();

	if P.tok == Scanner.SEMICOLON {
		// forward declaration
		P.Next();
	} else {
		fun.body = P.ParseBlock();
	}
	
	P.Ecart();
	
	return fun;
}


func (P *Parser) ParseExportDecl() {
	P.Trace("ExportDecl");
	
	// TODO This is deprecated syntax and should go away eventually.
	// (Also at the moment the syntax is everything goes...)
	//P.Expect(Scanner.EXPORT);

	has_paren := false;
	if P.tok == Scanner.LPAREN {
		P.Next();
		has_paren = true;
	}
	for P.tok == Scanner.IDENT {
		ident := P.ParseIdent();
		P.Optional(Scanner.COMMA);  // TODO this seems wrong
	}
	if has_paren {
		P.Expect(Scanner.RPAREN)
	}
	
	P.Ecart();
}


func (P *Parser) ParseDeclaration() AST.Node {
	P.Trace("Declaration");
	indent := P.indent;
	
	var node AST.Node;

	exported := false;
	if P.tok == Scanner.EXPORT {
		if P.level == 0 {
			exported = true;
		} else {
			P.Error(P.pos, "local declarations cannot be exported");
		}
		P.Next();
	}
	
	switch P.tok {
	case Scanner.CONST, Scanner.TYPE, Scanner.VAR:
		node = P.ParseDecl(exported, P.tok);
	case Scanner.FUNC:
		node = P.ParseFuncDecl(exported);
	case Scanner.EXPORT:
		if exported {
			P.Error(P.pos, "cannot mark export declaration for export");
		}
		P.Next();
		P.ParseExportDecl();
	default:
		if exported && (P.tok == Scanner.IDENT || P.tok == Scanner.LPAREN) {
			P.ParseExportDecl();
		} else {
			P.Error(P.pos, "declaration expected");
			P.Next();  // make progress
		}
	}
	
	if indent != P.indent {
		panic("imbalanced tracing code (Declaration)");
	}
	P.Ecart();
	return node;
}


// ----------------------------------------------------------------------------
// Program

func (P *Parser) ParseProgram() *AST.Program {
	P.Trace("Program");
	
	P.OpenScope();
	pos := P.pos;
	P.Expect(Scanner.PACKAGE);
	ident := P.ParseIdent();
	P.Optional(Scanner.SEMICOLON);
	
	decls := AST.NewList();
	{	P.OpenScope();
		if P.level != 0 {
			panic("incorrect scope level");
		}
		
		for P.tok == Scanner.IMPORT {
			P.ParseDecl(false, Scanner.IMPORT);
			P.Optional(Scanner.SEMICOLON);
		}
		
		for P.tok != Scanner.EOF {
			decls.Add(P.ParseDeclaration());
			P.Optional(Scanner.SEMICOLON);
		}
		
		if P.level != 0 {
			panic("incorrect scope level");
		}
		P.CloseScope();
	}
	
	P.CloseScope();
	P.Ecart();
	
	x := new(AST.Program);
	x.pos, x.ident, x.decls = pos, ident, decls;
	return x;
}
