// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package Parser

import Scanner "scanner"
import Globals "globals"
import Object "object"
import Type "type"
import Universe "universe"
import AST "ast"


// So I can submit and have a running parser for now...
const EnableSemanticTests = false;


export Parser
type Parser struct {
	comp *Globals.Compilation;
	verbose, indent int;
	S *Scanner.Scanner;
	
	// Token
	tok int;  // one token look-ahead
	pos int;  // token source position
	val string;  // token value (for IDENT, NUMBER, STRING only)

	// Semantic analysis
	top_scope *Globals.Scope;
	undef_types *Globals.List;
	exports *Globals.List;
}


// ----------------------------------------------------------------------------
// Support functions

func (P *Parser) PrintIndent() {
	for i := P.indent; i > 0; i-- {
		print ". ";
	}
}


func (P *Parser) Trace(msg string) {
	if P.verbose > 0 {
		P.PrintIndent();
		print msg, " {\n";
		P.indent++;
	}
}


func (P *Parser) Ecart() {
	if P.verbose > 0 {
		P.indent--;
		P.PrintIndent();
		print "}\n";
	}
}


func (P *Parser) Next() {
	P.tok, P.pos, P.val = P.S.Scan();
	if P.verbose > 1 {
		P.PrintIndent();
		print "[", P.pos, "] ", Scanner.TokenName(P.tok), "\n";
	}
}


func (P *Parser) Open(comp *Globals.Compilation, S *Scanner.Scanner, verbose int) {
	P.comp = comp;
	P.verbose = verbose;
	P.indent = 0;
	P.S = S;
	P.Next();
	P.top_scope = Universe.scope;
	P.undef_types = Globals.NewList();
	P.exports = Globals.NewList();
}


func (P *Parser) Error(pos int, msg string) {
	P.S.Error(pos, msg);
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
	P.top_scope = Globals.NewScope(P.top_scope);
}


func (P *Parser) CloseScope() {
	P.top_scope = P.top_scope.parent;
}


func (P *Parser) Lookup(ident string) *Globals.Object {
	for scope := P.top_scope; scope != nil; scope = scope.parent {
		obj := scope.Lookup(ident);
		if obj != nil {
			return obj;
		}
	}
	return nil;
}


func (P *Parser) DeclareInScope(scope *Globals.Scope, obj *Globals.Object) {
	if EnableSemanticTests && scope.Lookup(obj.ident) != nil {
		P.Error(obj.pos, `"` + obj.ident + `" is declared already`);
		return;  // don't insert it into the scope
	}
	scope.Insert(obj);
}


func (P *Parser) Declare(obj *Globals.Object) {
	P.DeclareInScope(P.top_scope, obj);
}


// ----------------------------------------------------------------------------
// Common productions


func (P *Parser) TryType() *Globals.Type;
func (P *Parser) ParseExpression();
func (P *Parser) TryStatement() bool;
func (P *Parser) ParseDeclaration();


func (P *Parser) ParseIdent() string {
	P.Trace("Ident");

	ident := "";
	if P.tok == Scanner.IDENT {
		ident = P.val;
		if P.verbose > 0 {
			P.PrintIndent();
			print "Ident = \"", ident, "\"\n";
		}
		P.Next();
	} else {
		P.Expect(Scanner.IDENT);  // use Expect() error handling
	}
	
	P.Ecart();
	return ident;
}


func (P *Parser) ParseIdentDecl(kind int) *Globals.Object {
	P.Trace("IdentDecl");
	
	pos := P.pos;
	obj := Globals.NewObject(pos, kind, P.ParseIdent());
	P.Declare(obj);
	
	P.Ecart();
	return obj;
}


func (P *Parser) ParseIdentDeclList(kind int) *Globals.List {
	P.Trace("IdentDeclList");
	
	list := Globals.NewList();
	list.AddObj(P.ParseIdentDecl(kind));
	for P.tok == Scanner.COMMA {
		P.Next();
		list.AddObj(P.ParseIdentDecl(kind));
	}
	
	P.Ecart();
	return list;
}


func (P *Parser) ParseIdentList() {
	P.Trace("IdentList");
	P.ParseIdent();
	for P.tok == Scanner.COMMA {
		P.Next();
		P.ParseIdent();
	}
	P.Ecart();
}


func (P *Parser) ParseQualifiedIdent(pos int, ident string) *Globals.Object {
	P.Trace("QualifiedIdent");

	if pos < 0 {
		pos = P.pos;
		ident = P.ParseIdent();
	}
	
	if EnableSemanticTests {
		obj := P.Lookup(ident);
		if obj == nil {
			P.Error(pos, `"` + ident + `" is not declared`);
			obj = Globals.NewObject(pos, Object.BAD, ident);
		}

		if obj.kind == Object.PACKAGE && P.tok == Scanner.PERIOD {
			panic "Qualified ident not complete yet";
			P.Next();
			P.ParseIdent();
		}
		P.Ecart();
		return obj;
		
	} else {
		if P.tok == Scanner.PERIOD {
			P.Next();
			P.ParseIdent();
		}
		P.Ecart();
		return nil;
	}
}


// ----------------------------------------------------------------------------
// Types

func (P *Parser) ParseType() *Globals.Type {
	P.Trace("Type");
	
	typ := P.TryType();
	if typ == nil {
		P.Error(P.pos, "type expected");
		typ = Universe.bad_t;
	}
	
	P.Ecart();
	return typ;
}


func (P *Parser) ParseTypeName() *Globals.Type {
	P.Trace("TypeName");
	
	if EnableSemanticTests {
		pos := P.pos;
		obj := P.ParseQualifiedIdent(-1, "");
		typ := obj.typ;
		if obj.kind != Object.TYPE {
			P.Error(pos, "qualified identifier does not denote a type");
			typ = Universe.bad_t;
		}
		P.Ecart();
		return typ;
	} else {
		P.ParseQualifiedIdent(-1, "");
		P.Ecart();
		return Universe.bad_t;
	}
}


func (P *Parser) ParseArrayType() *Globals.Type {
	P.Trace("ArrayType");
	
	P.Expect(Scanner.LBRACK);
	typ := Globals.NewType(Type.ARRAY);
	if P.tok != Scanner.RBRACK {
		// TODO set typ.len_
		P.ParseExpression();
	}
	P.Expect(Scanner.RBRACK);
	typ.elt = P.ParseType();
	P.Ecart();
	
	return typ;
}


func (P *Parser) ParseChannelType() *Globals.Type {
	P.Trace("ChannelType");
	
	P.Expect(Scanner.CHAN);
	typ := Globals.NewType(Type.CHANNEL);
	switch P.tok {
	case Scanner.SEND:
		typ.flags = Type.SEND;
		P.Next();
	case Scanner.RECV:
		typ.flags = Type.RECV;
		P.Next();
	default:
		typ.flags = Type.SEND + Type.RECV;
	}
	typ.elt = P.ParseType();
	P.Ecart();
	
	return typ;
}


func (P *Parser) ParseVarDeclList() {
	P.Trace("VarDeclList");
	
	list := P.ParseIdentDeclList(Object.VAR);
	typ := P.ParseType();  // TODO should check completeness of types
	for p := list.first; p != nil; p = p.next {
		p.obj.typ = typ;  // TODO should use/have set_type()
	}
	
	P.Ecart();
}


func (P *Parser) ParseParameterSection() {
	P.Trace("ParameterSection");
	P.ParseVarDeclList();
	P.Ecart();
}


func (P *Parser) ParseParameterList() {
	P.Trace("ParameterList");
	
	P.ParseParameterSection();
	for P.tok == Scanner.COMMA {
		P.Next();
		P.ParseParameterSection();
	}
	
	P.Ecart();
}


func (P *Parser) ParseParameters() {
	P.Trace("Parameters");
	
	P.Expect(Scanner.LPAREN);
	if P.tok != Scanner.RPAREN {
		P.ParseParameterList();
	}
	P.Expect(Scanner.RPAREN);
	
	P.Ecart();
}


func (P *Parser) TryResult() bool {
	P.Trace("Result (try)");
	
	res := false;
	if P.tok == Scanner.LPAREN {
		// TODO: here we allow empty returns - should proably fix this
		P.ParseParameters();
		res = true;
	} else {
		res = P.TryType() != nil;
	}
	P.Ecart();
	
	return res;
}


func MakeFunctionType(sig *Globals.Scope, p0, r0 int, check_recv bool) *Globals.Type {
  // Determine if we have a receiver or not.
  if p0 > 0 && check_recv {
    // method
	if p0 != 1 {
		panic "p0 != 1";
	}
  }
  typ := Globals.NewType(Type.FUNCTION);
  if p0 == 0 {
	typ.flags = 0;
  } else {
	typ.flags = Type.RECV;
  }
  typ.len_ = r0 - p0;
  typ.scope = sig;
  return typ;
}


// Anonymous signatures
//
//          (params)
//          (params) type
//          (params) (results)
// (recv) . (params)
// (recv) . (params) type
// (recv) . (params) (results)

func (P *Parser) ParseAnonymousSignature() *Globals.Type {
	P.Trace("AnonymousSignature");
	
	P.OpenScope();
	sig := P.top_scope;
	p0 := 0;
	
	recv_pos := P.pos;
	P.ParseParameters();
	
	if P.tok == Scanner.PERIOD {
		p0 = sig.entries.len_;
		if (p0 != 1) {
			P.Error(recv_pos, "must have exactly one receiver")
			panic "UNIMPLEMENTED";
			// TODO do something useful here
		}
		P.Next();
		P.ParseParameters();
	}
	
	r0 := sig.entries.len_;
	P.TryResult();
	P.CloseScope();
	
	P.Ecart();
	return MakeFunctionType(sig, p0, r0, true);
}


// Named signatures
//
//        name (params)
//        name (params) type
//        name (params) (results)
// (recv) name (params)
// (recv) name (params) type
// (recv) name (params) (results)

func (P *Parser) ParseNamedSignature() (name string, typ *Globals.Type) {
	P.Trace("NamedSignature");
	
	P.OpenScope();
	sig := P.top_scope;
	p0 := 0;

	if P.tok == Scanner.LPAREN {
		recv_pos := P.pos;
		P.ParseParameters();
		p0 = sig.entries.len_;
		if (p0 != 1) {
			print "p0 = ", p0, "\n";
			P.Error(recv_pos, "must have exactly one receiver")
			panic "UNIMPLEMENTED";
			// TODO do something useful here
		}
	}
	
	name = P.ParseIdent();

	P.ParseParameters();
	
	r0 := sig.entries.len_;
	P.TryResult();
	P.CloseScope();
	
	P.Ecart();
	return name, MakeFunctionType(sig, p0, r0, true);
}


func (P *Parser) ParseFunctionType() *Globals.Type {
	P.Trace("FunctionType");
	
	P.Expect(Scanner.FUNC);
	typ := P.ParseAnonymousSignature();
	
	P.Ecart();
	return typ;
}


func (P *Parser) ParseMethodDecl() {
	P.Trace("MethodDecl");
	
	P.ParseIdent();
	P.ParseParameters();
	P.TryResult();
	P.Optional(Scanner.SEMICOLON);
	
	P.Ecart();
}


func (P *Parser) ParseInterfaceType() *Globals.Type {
	P.Trace("InterfaceType");
	
	P.Expect(Scanner.INTERFACE);
	P.Expect(Scanner.LBRACE);
	P.OpenScope();
	typ := Globals.NewType(Type.INTERFACE);
	typ.scope = P.top_scope;
	for P.tok == Scanner.IDENT {
		P.ParseMethodDecl();
	}
	P.CloseScope();
	P.Expect(Scanner.RBRACE);
	
	P.Ecart();
	return typ;
}


func (P *Parser) ParseMapType() *Globals.Type {
	P.Trace("MapType");
	
	P.Expect(Scanner.MAP);
	P.Expect(Scanner.LBRACK);
	typ := Globals.NewType(Type.MAP);
	typ.key = P.ParseType();
	P.Expect(Scanner.RBRACK);
	typ.elt = P.ParseType();
	P.Ecart();
	
	return typ;
}


func (P *Parser) ParseFieldDecl() {
	P.Trace("FieldDecl");
	P.ParseVarDeclList();
	P.Ecart();
}


func (P *Parser) ParseStructType() *Globals.Type {
	P.Trace("StructType");
	
	P.Expect(Scanner.STRUCT);
	P.Expect(Scanner.LBRACE);
	P.OpenScope();
	typ := Globals.NewType(Type.STRUCT);
	typ.scope = P.top_scope;
	for P.tok == Scanner.IDENT {
		P.ParseFieldDecl();
		if P.tok != Scanner.RBRACE {
			P.Expect(Scanner.SEMICOLON);
		}
	}
	P.Optional(Scanner.SEMICOLON);
	P.CloseScope();
	P.Expect(Scanner.RBRACE);
	
	P.Ecart();
	return typ;
}


func (P *Parser) ParsePointerType() *Globals.Type {
	P.Trace("PointerType");
	
	P.Expect(Scanner.MUL);
	typ := Globals.NewType(Type.POINTER);
	
	if EnableSemanticTests {
		if P.tok == Scanner.IDENT {
			if P.Lookup(P.val) == nil {
				// implicit forward declaration
				// TODO very problematic: in which scope should the
				// type object be declared? It's different if this
				// is inside a struct or say in a var declaration.
				// This code is only here for "compatibility" with 6g.
				pos := P.pos;
				obj := Globals.NewObject(pos, Object.TYPE, P.ParseIdent());
				obj.typ = Globals.NewType(Type.UNDEF);
				obj.typ.obj = obj;  // primary type object
				typ.elt = obj.typ;
				// TODO obj should be declared, but scope is not clear
			} else {
				// type name
				// (ParseType() doesn't permit incomplete types,
				// so call ParseTypeName() here)
				typ.elt = P.ParseTypeName();
			}
		} else {
			typ.elt = P.ParseType();
		}
	
		// collect undefined pointer types
		if typ.elt.form == Type.UNDEF {
			P.undef_types.AddTyp(typ);
		}
		
	} else {
		typ.elt = P.ParseType();
	}

	P.Ecart();
	return typ;
}


// Returns nil if no type was found.
func (P *Parser) TryType() *Globals.Type {
	P.Trace("Type (try)");
	
	pos := P.pos;
	var typ *Globals.Type = nil;
	switch P.tok {
	case Scanner.IDENT: typ = P.ParseTypeName();
	case Scanner.LBRACK: typ = P.ParseArrayType();
	case Scanner.CHAN: typ = P.ParseChannelType();
	case Scanner.INTERFACE: typ = P.ParseInterfaceType();
	case Scanner.FUNC: typ = P.ParseFunctionType();
	case Scanner.MAP: typ = P.ParseMapType();
	case Scanner.STRUCT: typ = P.ParseStructType();
	case Scanner.MUL: typ = P.ParsePointerType();
	}

	if typ != nil && typ.form == Type.UNDEF {
		P.Error(pos, "incomplete type");
	}

	P.Ecart();
	return typ;
}


// ----------------------------------------------------------------------------
// Blocks

func (P *Parser) ParseStatement() {
	P.Trace("Statement");
	if !P.TryStatement() {
		P.Error(P.pos, "statement expected");
		P.Next();  // make progress
	}
	P.Ecart();
}


func (P *Parser) ParseStatementList() {
	P.Trace("StatementList");
	for P.TryStatement() {
		P.Optional(Scanner.SEMICOLON);
	}
	P.Ecart();
}


func (P *Parser) ParseBlock() {
	P.Trace("Block");
	
	P.Expect(Scanner.LBRACE);
	P.OpenScope();
	if P.tok != Scanner.RBRACE && P.tok != Scanner.SEMICOLON {
		P.ParseStatementList();
	}
	P.Optional(Scanner.SEMICOLON);
	P.CloseScope();
	P.Expect(Scanner.RBRACE);
	
	P.Ecart();
}


// ----------------------------------------------------------------------------
// Expressions

func (P *Parser) ParseExpressionList() *Globals.List {
	P.Trace("ExpressionList");
	
	list := Globals.NewList();
	P.ParseExpression();
	list.AddInt(0);  // TODO fix this - add correct list element
	for P.tok == Scanner.COMMA {
		P.Next();
		P.ParseExpression();
		list.AddInt(0);  // TODO fix this - add correct list element
	}
	
	P.Ecart();
	return list;
}


func (P *Parser) ParseNew() {
	P.Trace("New");
	
	P.Expect(Scanner.NEW);
	P.Expect(Scanner.LPAREN);
	P.ParseType();
	if P.tok == Scanner.COMMA {
		P.Next();
		P.ParseExpressionList()
	}
	P.Expect(Scanner.RPAREN);
	
	P.Ecart();
}


func (P *Parser) ParseFunctionLit() {
	P.Trace("FunctionLit");
	
	P.ParseFunctionType();
	P.ParseBlock();
	
	P.Ecart();
}


func (P *Parser) ParseExpressionPair() {
	P.Trace("ExpressionPair");

	P.ParseExpression();
	P.Expect(Scanner.COLON);
	P.ParseExpression();
	
	P.Ecart();
}


func (P *Parser) ParseExpressionPairList() {
	P.Trace("ExpressionPairList");

	P.ParseExpressionPair();
	for (P.tok == Scanner.COMMA) {
		P.ParseExpressionPair();
	}
	
	P.Ecart();
}


func (P *Parser) ParseBuiltinCall() {
	P.Trace("BuiltinCall");
	
	P.ParseExpressionList();  // TODO should be optional
	
	P.Ecart();
}


func (P *Parser) ParseCompositeLit(typ *Globals.Type) {
	P.Trace("CompositeLit");
	
	// TODO I think we should use {} instead of () for
	// composite literals to syntactically distinguish
	// them from conversions. For now: allow both.
	var paren int;
	if P.tok == Scanner.LPAREN {
		P.Next();
		paren = Scanner.RPAREN;
	} else {
		P.Expect(Scanner.LBRACE);
		paren = Scanner.RBRACE;
	}
	
	// TODO: should allow trailing ','
	if P.tok != paren {
		P.ParseExpression();
		if P.tok == Scanner.COMMA {
			P.Next();
			if P.tok != paren {
				P.ParseExpressionList();
			}
		} else if P.tok == Scanner.COLON {
			P.Next();
			P.ParseExpression();
			if P.tok == Scanner.COMMA {
				P.Next();
				if P.tok != paren {
					P.ParseExpressionPairList();
				}
			}
		}
	}

	P.Expect(paren);

	P.Ecart();
}


func (P *Parser) ParseOperand(pos int, ident string) {
	P.Trace("Operand");

	if pos < 0 && P.tok == Scanner.IDENT {
		// no look-ahead yet
		pos = P.pos;
		ident = P.val;
		P.Next();
	}
	
	if pos >= 0 {
		// TODO set these up properly in the Universe
		if ident == "panic" || ident == "print" {
			P.ParseBuiltinCall();
			goto exit;
		}
	
		P.ParseQualifiedIdent(pos, ident);
		// TODO enable code below
		/*
		if obj.kind == Object.TYPE {
			P.ParseCompositeLit(obj.typ);
		}
		*/
		goto exit;
	}
	
	switch P.tok {
	case Scanner.IDENT:
		panic "UNREACHABLE";
	case Scanner.LPAREN:
		P.Next();
		P.ParseExpression();
		P.Expect(Scanner.RPAREN);
	case Scanner.STRING: fallthrough;
	case Scanner.NUMBER: fallthrough;
	case Scanner.NIL: fallthrough;
	case Scanner.IOTA: fallthrough;
	case Scanner.TRUE: fallthrough;
	case Scanner.FALSE:
		P.Next();
	case Scanner.FUNC:
		P.ParseFunctionLit();
	case Scanner.NEW:
		P.ParseNew();
	default:
		typ := P.TryType();
		if typ != nil {
			P.ParseCompositeLit(typ);
		} else {
			P.Error(P.pos, "operand expected");
			P.Next();  // make progress
		}
	}
	
exit:
	P.Ecart();
}


func (P *Parser) ParseSelectorOrTypeAssertion() {
	P.Trace("SelectorOrTypeAssertion");
	
	P.Expect(Scanner.PERIOD);
	if P.tok == Scanner.IDENT {
		P.ParseIdent();
	} else {
		P.Expect(Scanner.LPAREN);
		P.ParseType();
		P.Expect(Scanner.RPAREN);
	}
	
	P.Ecart();
}


func (P *Parser) ParseIndexOrSlice() {
	P.Trace("IndexOrSlice");
	
	P.Expect(Scanner.LBRACK);
	P.ParseExpression();
	if P.tok == Scanner.COLON {
		P.Next();
		P.ParseExpression();
	}
	P.Expect(Scanner.RBRACK);
	
	P.Ecart();
}


func (P *Parser) ParseCall() {
	P.Trace("Call");
	
	P.Expect(Scanner.LPAREN);
	if P.tok != Scanner.RPAREN {
		P.ParseExpressionList();
	}
	P.Expect(Scanner.RPAREN);
	
	P.Ecart();
}


func (P *Parser) ParsePrimaryExpr(pos int, ident string) AST.Expr {
	P.Trace("PrimaryExpr");
	
	P.ParseOperand(pos, ident);
	for {
		switch P.tok {
		case Scanner.PERIOD:
			P.ParseSelectorOrTypeAssertion();
		case Scanner.LBRACK:
			P.ParseIndexOrSlice();
		case Scanner.LPAREN:
			P.ParseCall();
		default:
			P.Ecart();
			return nil;
		}
	}
	
	P.Ecart();
	return nil;
}


func (P *Parser) ParsePrimaryExprList() {
	P.Trace("PrimaryExprList");
	
	P.ParsePrimaryExpr(-1, "");
	for P.tok == Scanner.COMMA {
		P.Next();
		P.ParsePrimaryExpr(-1, "");
	}
	
	P.Ecart();
}


func (P *Parser) ParseUnaryExpr() AST.Expr {
	P.Trace("UnaryExpr");
	
	switch P.tok {
	case Scanner.ADD: fallthrough;
	case Scanner.SUB: fallthrough;
	case Scanner.NOT: fallthrough;
	case Scanner.XOR: fallthrough;
	case Scanner.MUL: fallthrough;
	case Scanner.RECV: fallthrough;
	case Scanner.AND:
		P.Next();
		P.ParseUnaryExpr();
		P.Ecart();
		return nil;  // TODO fix this
	}
	P.ParsePrimaryExpr(-1, "");
	
	P.Ecart();
	return nil;  // TODO fix this
}


func Precedence(tok int) int {
	// TODO should use a map or array here for lookup
	switch tok {
	case Scanner.LOR:
		return 1;
	case Scanner.LAND:
		return 2;
	case Scanner.SEND, Scanner.RECV:
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


func (P *Parser) ParseBinaryExpr(pos int, ident string, prec1 int) AST.Expr {
	P.Trace("BinaryExpr");
	
	var x AST.Expr;
	if pos >= 0 {
		x = P.ParsePrimaryExpr(pos, ident);
	} else {
		x = P.ParseUnaryExpr();
	}
	for prec := Precedence(P.tok); prec >= prec1; prec-- {
		for Precedence(P.tok) == prec {
			e := new(AST.BinaryExpr);
			e.typ = Universe.undef_t;  // TODO fix this
			e.op = P.tok;  // TODO should we use tokens or separate operator constants?
			e.x = x;
			P.Next();
			e.y = P.ParseBinaryExpr(-1, "", prec + 1);
			x = e;
		}
	}
	
	P.Ecart();
}


// Expressions where the first token may be an identifier which has already
// been consumed. If the identifier is present, pos is the identifier position,
// otherwise pos must be < 0 (and ident is ignored).
func (P *Parser) ParseIdentExpression(pos int, ident string) {
	P.Trace("IdentExpression");
	indent := P.indent;
	
	P.ParseBinaryExpr(pos, ident, 1);
	
	if indent != P.indent {
		panic "imbalanced tracing code (Expression)";
	}
	P.Ecart();
}


func (P *Parser) ParseExpression() {
	P.Trace("Expression");	
	P.ParseIdentExpression(-1, "");
	P.Ecart();
}


// ----------------------------------------------------------------------------
// Statements

func (P *Parser) ConvertToExprList(pos_list, ident_list, expr_list *Globals.List) {
	for p, q := pos_list.first, ident_list.first; q != nil; p, q = p.next, q.next {
		pos, ident := p.val, q.str;
		if EnableSemanticTests {
			obj := P.Lookup(ident);
			if obj == nil {
				P.Error(pos, `"` + ident + `" is not declared`);
				obj = Globals.NewObject(pos, Object.BAD, ident);
			}
		}
		expr_list.AddInt(0);  // TODO fix this - add correct expression
	}
	ident_list.Clear();
}


func (P *Parser) ParseIdentOrExpr(pos_list, ident_list, expr_list *Globals.List) {
	P.Trace("IdentOrExpr");
	
	pos_list.AddInt(P.pos);
	pos, ident := -1, "";
	just_ident := false;
	if expr_list.len_ == 0 /* only idents so far */ && P.tok == Scanner.IDENT {
		pos, ident = P.pos, P.val;
		P.Next();
		switch P.tok {
		case Scanner.COMMA,
			Scanner.COLON,
			Scanner.DEFINE,
			Scanner.ASSIGN,
			Scanner.ADD_ASSIGN,
			Scanner.SUB_ASSIGN,
			Scanner.MUL_ASSIGN,
			Scanner.QUO_ASSIGN,
			Scanner.REM_ASSIGN,
			Scanner.AND_ASSIGN,
			Scanner.OR_ASSIGN,
			Scanner.XOR_ASSIGN,
			Scanner.SHL_ASSIGN,
			Scanner.SHR_ASSIGN:
			// identifier is not part of a more complicated expression
			just_ident = true;
		}
	}

	if just_ident {
		ident_list.AddStr(ident);
	} else {
		P.ConvertToExprList(pos_list, ident_list, expr_list);
		P.ParseIdentExpression(pos, ident);
		expr_list.AddInt(0);  // TODO fix this - add correct expression
	}
	
	P.Ecart();
}


func (P *Parser) ParseIdentOrExprList() (pos_list, ident_list, expr_list *Globals.List) {
	P.Trace("IdentOrExprList");
	
	pos_list, ident_list, expr_list = Globals.NewList(), Globals.NewList(), Globals.NewList();
	P.ParseIdentOrExpr(pos_list, ident_list, expr_list);
	for P.tok == Scanner.COMMA {
		P.Next();
		P.ParseIdentOrExpr(pos_list, ident_list, expr_list);
	}
	
	P.Ecart();
	return pos_list, ident_list, expr_list;
}


func (P *Parser) ParseSimpleStat() {
	P.Trace("SimpleStat");
	
	// If we see an identifier, we don't know if it's part of a
	// label declaration, (multiple) variable declaration, assignment,
	// or simply an expression, without looking ahead.
	// Strategy: We parse an expression list, but simultaneously, as
	// long as possible, maintain a list of identifiers which is converted
	// into an expression list only if neccessary. The result of
	// ParseIdentOrExprList is a list of ident/expr positions and either
	// a non-empty list of identifiers or a non-empty list of expressions
	// (but not both).
	pos_list, ident_list, expr_list := P.ParseIdentOrExprList();
	
	switch P.tok {
	case Scanner.COLON:
		// label declaration
		if EnableSemanticTests && ident_list.len_ != 1 {
			P.Error(P.pos, "illegal label declaration");
		}
		P.Next();
		
	case Scanner.DEFINE:
		// variable declaration
		if EnableSemanticTests && ident_list.len_ == 0 {
			P.Error(P.pos, "illegal left-hand side for declaration");
		}
		P.Next();
		pos := P.pos;
		val_list := P.ParseExpressionList();
		if EnableSemanticTests && val_list.len_ != ident_list.len_ {
			P.Error(pos, "number of expressions does not match number of variables");
		}
		// declare variables
		if EnableSemanticTests {
			for p, q := pos_list.first, ident_list.first; q != nil; p, q = p.next, q.next {
				obj := Globals.NewObject(p.val, Object.VAR, q.str);
				P.Declare(obj);
				// TODO set correct types
			}
		}
		
	case Scanner.ASSIGN: fallthrough;
	case Scanner.ADD_ASSIGN: fallthrough;
	case Scanner.SUB_ASSIGN: fallthrough;
	case Scanner.MUL_ASSIGN: fallthrough;
	case Scanner.QUO_ASSIGN: fallthrough;
	case Scanner.REM_ASSIGN: fallthrough;
	case Scanner.AND_ASSIGN: fallthrough;
	case Scanner.OR_ASSIGN: fallthrough;
	case Scanner.XOR_ASSIGN: fallthrough;
	case Scanner.SHL_ASSIGN: fallthrough;
	case Scanner.SHR_ASSIGN:
		P.ConvertToExprList(pos_list, ident_list, expr_list);
		P.Next();
		pos := P.pos;
		val_list := P.ParseExpressionList();
		if EnableSemanticTests && val_list.len_ != expr_list.len_ {
			P.Error(pos, "number of expressions does not match number of variables");
		}
		
	default:
		P.ConvertToExprList(pos_list, ident_list, expr_list);
		if EnableSemanticTests && expr_list.len_ != 1 {
			P.Error(P.pos, "no expression list allowed");
		}
		if P.tok == Scanner.INC || P.tok == Scanner.DEC {
			P.Next();
		}
	}
	
	P.Ecart();
}


func (P *Parser) ParseGoStat() {
	P.Trace("GoStat");
	P.Expect(Scanner.GO);
	P.ParseExpression();
	P.Ecart();
}


func (P *Parser) ParseReturnStat() {
	P.Trace("ReturnStat");
	
	P.Expect(Scanner.RETURN);
	if P.tok != Scanner.SEMICOLON && P.tok != Scanner.RBRACE {
		P.ParseExpressionList();
	}
	
	P.Ecart();
}


func (P *Parser) ParseControlFlowStat(tok int) {
	P.Trace("ControlFlowStat");
	
	P.Expect(tok);
	if P.tok == Scanner.IDENT {
		P.ParseIdent();
	}
	
	P.Ecart();
}


func (P *Parser) ParseIfStat() *AST.IfStat {
	P.Trace("IfStat");
	
	P.Expect(Scanner.IF);
	P.OpenScope();
	if P.tok != Scanner.LBRACE {
		if P.tok != Scanner.SEMICOLON {
			P.ParseSimpleStat();
		}
		if P.tok == Scanner.SEMICOLON {
			P.Next();
			if P.tok != Scanner.LBRACE {
				P.ParseExpression();
			}
		}
	}
	P.ParseBlock();
	if P.tok == Scanner.ELSE {
		P.Next();
		if P.tok == Scanner.IF {
			P.ParseIfStat();
		} else {
			// TODO should be P.ParseBlock()
			P.ParseStatement();
		}
	}
	P.CloseScope();
	
	P.Ecart();
}


func (P *Parser) ParseForStat() {
	P.Trace("ForStat");
	
	P.Expect(Scanner.FOR);
	P.OpenScope();
	if P.tok != Scanner.LBRACE {
		if P.tok != Scanner.SEMICOLON {
			P.ParseSimpleStat();
		}
		if P.tok == Scanner.SEMICOLON {
			P.Next();
			if P.tok != Scanner.SEMICOLON {
				P.ParseExpression();
			}
			P.Expect(Scanner.SEMICOLON);
			if P.tok != Scanner.LBRACE {
				P.ParseSimpleStat();
			}
		}
	}
	P.ParseBlock();
	P.CloseScope();
	
	P.Ecart();
}


func (P *Parser) ParseCase() {
	P.Trace("Case");
	
	if P.tok == Scanner.CASE {
		P.Next();
		P.ParseExpressionList();
	} else {
		P.Expect(Scanner.DEFAULT);
	}
	P.Expect(Scanner.COLON);
	
	P.Ecart();
}


func (P *Parser) ParseCaseList() {
	P.Trace("CaseList");
	
	P.ParseCase();
	for P.tok == Scanner.CASE || P.tok == Scanner.DEFAULT {
		P.ParseCase();
	}
	
	P.Ecart();
}


func (P *Parser) ParseCaseClause() {
	P.Trace("CaseClause");
	
	P.ParseCaseList();
	if P.tok != Scanner.FALLTHROUGH && P.tok != Scanner.RBRACE {
		P.ParseStatementList();
		P.Optional(Scanner.SEMICOLON);
	}
	if P.tok == Scanner.FALLTHROUGH {
		P.Next();
		P.Optional(Scanner.SEMICOLON);
	}
	
	P.Ecart();
}


func (P *Parser) ParseSwitchStat() {
	P.Trace("SwitchStat");
	
	P.Expect(Scanner.SWITCH);
	P.OpenScope();
	if P.tok != Scanner.LBRACE {
		if P.tok != Scanner.SEMICOLON {
			P.ParseSimpleStat();
		}
		if P.tok == Scanner.SEMICOLON {
			P.Next();
			if P.tok != Scanner.LBRACE {
				P.ParseExpression();
			}
		}
	}
	P.Expect(Scanner.LBRACE);
	for P.tok == Scanner.CASE || P.tok == Scanner.DEFAULT {
		P.ParseCaseClause();
	}
	P.Expect(Scanner.RBRACE);
	P.CloseScope();
	
	P.Ecart();
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


func (P *Parser) ParseRangeStat() bool {
	P.Trace("RangeStat");
	
	P.Expect(Scanner.RANGE);
	P.ParseIdentList();
	P.Expect(Scanner.DEFINE);
	P.ParseExpression();
	P.ParseBlock();
	
	P.Ecart();
}


func (P *Parser) ParseSelectStat() bool {
	P.Trace("SelectStat");
	
	P.Expect(Scanner.SELECT);
	P.Expect(Scanner.LBRACE);
	for P.tok != Scanner.RBRACE && P.tok != Scanner.EOF {
		P.ParseCommClause();
	}
	P.Next();
	
	P.Ecart();
}


func (P *Parser) TryStatement() bool {
	P.Trace("Statement (try)");
	indent := P.indent;

	res := true;
	switch P.tok {
	case Scanner.CONST: fallthrough;
	case Scanner.TYPE: fallthrough;
	case Scanner.VAR:
		P.ParseDeclaration();
	case Scanner.FUNC:
		// for now we do not allow local function declarations
		fallthrough;
	case Scanner.MUL, Scanner.SEND, Scanner.RECV, Scanner.IDENT:
		P.ParseSimpleStat();
	case Scanner.GO:
		P.ParseGoStat();
	case Scanner.RETURN:
		P.ParseReturnStat();
	case Scanner.BREAK, Scanner.CONTINUE, Scanner.GOTO:
		P.ParseControlFlowStat(P.tok);
	case Scanner.LBRACE:
		P.ParseBlock();
	case Scanner.IF:
		P.ParseIfStat();
	case Scanner.FOR:
		P.ParseForStat();
	case Scanner.SWITCH:
		P.ParseSwitchStat();
	case Scanner.RANGE:
		P.ParseRangeStat();
	case Scanner.SELECT:
		P.ParseSelectStat();
	default:
		// no statement found
		res = false;
	}

	if indent != P.indent {
		panic "imbalanced tracing code (Statement)"
	}
	P.Ecart();
	return res;
}


// ----------------------------------------------------------------------------
// Declarations

func (P *Parser) ParseImportSpec() {
	P.Trace("ImportSpec");
	
	if P.tok == Scanner.PERIOD {
		P.Next();
	} else if P.tok == Scanner.IDENT {
		P.Next();
	}
	P.Expect(Scanner.STRING);
	
	P.Ecart();
}


func (P *Parser) ParseImportDecl() {
	P.Trace("ImportDecl");
	
	P.Expect(Scanner.IMPORT);
	if P.tok == Scanner.LPAREN {
		P.Next();
		for P.tok != Scanner.RPAREN && P.tok != Scanner.EOF {
			P.ParseImportSpec();
			P.Optional(Scanner.SEMICOLON);  // TODO this seems wrong
		}
		P.Next();
	} else {
		P.ParseImportSpec();
	}
	
	P.Ecart();
}


func (P *Parser) ParseConstSpec(exported bool) {
	P.Trace("ConstSpec");
	
	list := P.ParseIdentDeclList(Object.CONST);
	typ := P.TryType();
	if typ != nil {
		for p := list.first; p != nil; p = p.next {
			p.obj.exported = exported;
			p.obj.typ = typ;  // TODO should use/have set_type()!
		}
	}
	if P.tok == Scanner.ASSIGN {
		P.Next();
		P.ParseExpressionList();
	}
	
	P.Ecart();
}


func (P *Parser) ParseTypeSpec(exported bool) {
	P.Trace("TypeSpec");
	
	pos := P.pos;
	ident := P.ParseIdent();
	obj := P.top_scope.Lookup(ident);  // only lookup in top scope!
	if obj != nil {
		// name already declared - ok if forward declared type
		if obj.kind != Object.TYPE || obj.typ.form != Type.UNDEF {
			// TODO use obj.pos to refer to decl pos in error msg!
			P.Error(pos, `"` + ident + `" is declared already`);
		}
	} else {
		obj = Globals.NewObject(pos, Object.TYPE, ident);
		obj.exported = exported;
		obj.typ = Globals.NewType(Type.UNDEF);
		obj.typ.obj = obj;  // primary type object
		P.Declare(obj);
	}
	
	typ := P.TryType();  // nil if we have an explicit forward declaration

	if typ != nil {
		obj.typ = typ;
		if typ.obj == nil {
			typ.obj = obj;  // primary type object
		}
	}
	
	P.Ecart();
}


func (P *Parser) ParseVarSpec(exported bool) {
	P.Trace("VarSpec");
	
	list := P.ParseIdentDeclList(Object.VAR);
	if P.tok == Scanner.ASSIGN {
		P.Next();
		P.ParseExpressionList();
	} else {
		typ := P.ParseType();
		for p := list.first; p != nil; p = p.next {
			p.obj.typ = typ;  // TODO should use/have set_type()!
		}
		if P.tok == Scanner.ASSIGN {
			P.Next();
			P.ParseExpressionList();
		}
	}
	
	P.Ecart();
}


// TODO With method variables, we wouldn't need this dispatch function.
func (P *Parser) ParseSpec(exported bool, keyword int) {
	switch keyword {
	case Scanner.CONST: P.ParseConstSpec(exported);
	case Scanner.TYPE: P.ParseTypeSpec(exported);
	case Scanner.VAR: P.ParseVarSpec(exported);
	default: panic "UNREACHABLE";
	}
}


func (P *Parser) ParseDecl(exported bool, keyword int) {
	P.Trace("Decl");
	
	P.Expect(keyword);
	if P.tok == Scanner.LPAREN {
		P.Next();
		for P.tok == Scanner.IDENT {
			P.ParseSpec(exported, keyword);
			if P.tok != Scanner.RPAREN {
				P.Expect(Scanner.SEMICOLON);
			}
		}
		P.Next();
	} else {
		P.ParseSpec(exported, keyword);
	}
	
	P.Ecart();
}


func (P *Parser) ParseFuncDecl(exported bool) {
	P.Trace("FuncDecl");
	
	P.Expect(Scanner.FUNC);
	P.ParseNamedSignature();
	if P.tok == Scanner.SEMICOLON {
		// forward declaration
		P.Next();
	} else {
		P.ParseBlock();
	}
	
	P.Ecart();
}


func (P *Parser) ParseExportDecl() {
	P.Trace("ExportDecl");
	
	// TODO this needs to be clarified - the current syntax is
	// "everything goes" - sigh...
	//P.Expect(Scanner.EXPORT);
	has_paren := false;
	if P.tok == Scanner.LPAREN {
		P.Next();
		has_paren = true;
	}
	for P.tok == Scanner.IDENT {
		P.exports.AddStr(P.ParseIdent());
		P.Optional(Scanner.COMMA);  // TODO this seems wrong
	}
	if has_paren {
		P.Expect(Scanner.RPAREN)
	}
	
	P.Ecart();
}


func (P *Parser) ParseDeclaration() {
	P.Trace("Declaration");
	indent := P.indent;
	
	exported := false;
	if P.tok == Scanner.EXPORT {
		P.Next();
		exported = true;
	}
	switch P.tok {
	case Scanner.CONST, Scanner.TYPE, Scanner.VAR:
		P.ParseDecl(exported, P.tok);
	case Scanner.FUNC:
		P.ParseFuncDecl(exported);
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
		panic "imbalanced tracing code (Declaration)"
	}
	P.Ecart();
}


// ----------------------------------------------------------------------------
// Program

func (P *Parser) ResolveUndefTypes() {
	if !EnableSemanticTests {
		return;
	}
	
	for p := P.undef_types.first; p != nil; p = p.next {
		typ := p.typ;
		if typ.form != Type.POINTER {
			panic "unresolved types should be pointers only";
		}
		if typ.elt.form != Type.UNDEF {
			panic "unresolved pointer should point to undefined type";
		}
		obj := typ.elt.obj;
		typ.elt = obj.typ;
		if typ.elt.form == Type.UNDEF {
			P.Error(obj.pos, `"` + obj.ident + `" is not declared`);
		}
	}
}


func (P *Parser) MarkExports() {
	if !EnableSemanticTests {
		return;
	}
	
	scope := P.top_scope;
	for p := P.exports.first; p != nil; p = p.next {
		obj := scope.Lookup(p.str);
		if obj != nil {
			obj.exported = true;
			// For now we export deep
			// TODO this should change eventually - we need selective export
			if obj.kind == Object.TYPE {
				typ := obj.typ;
				if typ.form == Type.STRUCT || typ.form == Type.INTERFACE {
					scope := typ.scope;
					for p := scope.entries.first; p != nil; p = p.next {
						p.obj.exported = true;
					}
				}
			}
		} else {
			// TODO need to report proper src position
			P.Error(0, `"` + p.str + `" is not declared - cannot be exported`);
		}
	}
}


func (P *Parser) ParseProgram() {
	P.Trace("Program");
	
	P.OpenScope();
	P.Expect(Scanner.PACKAGE);
	pkg := P.comp.pkgs[0];
	pkg.obj = P.ParseIdentDecl(Object.PACKAGE);
	P.Optional(Scanner.SEMICOLON);
	
	{	P.OpenScope();
		pkg.scope = P.top_scope;
		for P.tok == Scanner.IMPORT {
			P.ParseImportDecl();
			P.Optional(Scanner.SEMICOLON);
		}
		
		for P.tok != Scanner.EOF {
			P.ParseDeclaration();
			P.Optional(Scanner.SEMICOLON);
		}
		
		P.ResolveUndefTypes();
		P.MarkExports();
		P.CloseScope();
	}
	
	P.CloseScope();
	P.Ecart();
}
