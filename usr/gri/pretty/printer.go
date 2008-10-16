// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package Printer

import Scanner "scanner"
import Node "node"


export type Printer struct {
	level int;  // true scope level
	indent int;  // indentation level
	semi bool;  // pending ";"
	newl bool;  // pending "\n"
}


func (P *Printer) String(pos int, s string) {
	if P.semi && P.level > 0 {  // no semicolons at level 0
		print(";");
	}
	if P.newl {
		print("\n");
		for i := P.indent; i > 0; i-- {
			print("\t");
		}
	}
	print(s);
	P.newl, P.semi = false, false;
}


func (P *Printer) Blank() {
	P.String(0, " ");
}


func (P *Printer) Token(pos int, tok int) {
	P.String(pos, Scanner.TokenString(tok));
}


func (P *Printer) NewLine() {  // explicit "\n"
	print("\n");
	P.semi, P.newl = false, true;
}


func (P *Printer) OpenScope(paren string) {
	P.semi, P.newl = false, false;
	P.String(0, paren);
	P.level++;
	P.indent++;
	P.newl = true;
}


func (P *Printer) CloseScope(paren string) {
	P.indent--;
	P.semi = false;
	P.String(0, paren);
	P.level--;
	P.semi, P.newl = false, true;
}


// ----------------------------------------------------------------------------
// Types

func (P *Printer) Type(t *Node.Type)
func (P *Printer) Expr(x *Node.Expr)

func (P *Printer) Parameters(list *Node.List) {
	var prev int;
	for i, n := 0, list.len(); i < n; i++ {
		x := list.at(i).(*Node.Expr);
		if i > 0 {
			if prev == x.tok || prev == Scanner.TYPE {
				P.String(0, ", ");
			} else {
				P.Blank();
			}
		}
		P.Expr(x);
		prev = x.tok;
	}
}


func (P *Printer) Fields(list *Node.List) {
	var prev int;
	for i, n := 0, list.len(); i < n; i++ {
		x := list.at(i).(*Node.Expr);
		if i > 0 {
			if prev == Scanner.TYPE {
				P.String(0, ";");
				P.newl = true;
			} else if prev == x.tok {
				P.String(0, ", ");
			} else {
				P.Blank();
			}
		}
		P.Expr(x);
		prev = x.tok;
	}
	P.newl = true;
}


func (P *Printer) Type(t *Node.Type) {
	if t == nil {  // TODO remove this check
		P.String(0, "<nil type>");
		return;
	}

	switch t.tok {
	case Scanner.IDENT:
		P.Expr(t.expr);

	case Scanner.LBRACK:
		P.String(t.pos, "[");
		if t.expr != nil {
			P.Expr(t.expr);
		}
		P.String(0, "] ");
		P.Type(t.elt);

	case Scanner.STRUCT:
		P.String(t.pos, "struct");
		if t.list != nil {
			P.OpenScope(" {");
			P.Fields(t.list);
			P.CloseScope("}");
		}

	case Scanner.MAP:
		P.String(t.pos, "[");
		P.Type(t.key);
		P.String(0, "] ");
		P.Type(t.elt);

	case Scanner.CHAN:
		var m string;
		switch t.mode {
		case Node.FULL: m = "chan ";
		case Node.RECV: m = "<-chan ";
		case Node.SEND: m = "chan <- ";
		}
		P.String(t.pos, m);
		P.Type(t.elt);

	case Scanner.INTERFACE:
		P.String(t.pos, "interface");
		if t.list != nil {
			P.OpenScope(" {");
			/*
			for i := 0; i < x.methods.len(); i++ {
				P.Print(x.methods.at(i));
				P.newl, P.semi = true, true;
			}
			*/
			P.CloseScope("}");
		}

	case Scanner.MUL:
		P.String(t.pos, "*");
		P.Type(t.elt);

	case Scanner.LPAREN:
		P.String(t.pos, "(");
		P.Parameters(t.list);
		P.String(0, ")");
		if t.elt != nil {
			P.String(0, " (");
			P.Parameters(t.elt.list);
			P.String(0, ")");
		}

	default:
		panic("UNREACHABLE");
	}
}


// ----------------------------------------------------------------------------
// Expressions

func (P *Printer) Expr1(x *Node.Expr, prec1 int) {
	if x == nil {
		return;  // empty expression list
	}

	switch x.tok {
	case Scanner.VAR:
		panic("UNIMPLEMENTED (VAR)");
		
	case Scanner.TYPE:
		P.Type(x.t);

	case Scanner.IDENT, Scanner.INT, Scanner.STRING, Scanner.FLOAT:
		P.String(x.pos, x.s);

	case Scanner.COMMA:
		P.Expr1(x.x, 0);
		P.String(x.pos, ", ");
		P.Expr1(x.y, 0);

	case Scanner.PERIOD:
		P.Expr1(x.x, 8);  // 8 == highest precedence
		P.String(x.pos, ".");
		P.Expr1(x.y, 8);
		
	case Scanner.LBRACK:
		P.Expr1(x.x, 8);
		P.String(x.pos, "[");
		P.Expr1(x.y, 0);
		P.String(0, "]");

	case Scanner.LPAREN:
		P.Expr1(x.x, 8);
		P.String(x.pos, "(");
		P.Expr1(x.y, 0);
		P.String(0, ")");
		
	default:
		if x.x == nil {
			// unary expression
			P.Token(x.pos, x.tok);
			P.Expr1(x.y, 7);  // 7 == unary operator precedence
		} else {
			// binary expression: print ()'s if necessary
			prec := Scanner.Precedence(x.tok);
			if prec < prec1 {
				P.String(0, "(");
			}
			P.Expr1(x.x, prec);
			P.Blank();
			P.Token(x.pos, x.tok);
			P.Blank();
			P.Expr1(x.y, prec);
			if prec < prec1 {
				P.String(0, ")");
			}
		}
	}
}


func (P *Printer) Expr(x *Node.Expr) {
	P.Expr1(x, 0);
}


// ----------------------------------------------------------------------------
// Statements

func (P *Printer) Stat(s *Node.Stat)

func (P *Printer) StatementList(list *Node.List) {
	for i, n := 0, list.len(); i < n; i++ {
		P.Stat(list.at(i).(*Node.Stat));
		P.newl = true;
	}
}


func (P *Printer) Block(list *Node.List, indent bool) {
	P.OpenScope("{");
	if !indent {
		P.indent--;
	}
	P.StatementList(list);
	if !indent {
		P.indent++;
	}
	P.CloseScope("}");
}


func (P *Printer) ControlClause(s *Node.Stat) {
	if s.init != nil {
		P.Blank();
		P.Stat(s.init);
		P.semi = true;
	}
	if s.expr != nil {
		P.Blank();
		P.Expr(s.expr);
		P.semi = false;
	}
	if s.tok == Scanner.FOR && s.post != nil {
		P.semi = true;
		P.Blank();
		P.Stat(s.post);
		P.semi = false;
	}
	P.Blank();
}


func (P *Printer) Declaration(d *Node.Decl, parenthesized bool);

func (P *Printer) Stat(s *Node.Stat) {
	if s == nil {  // TODO remove this check
		P.String(0, "<nil stat>");
		return;
	}

	switch s.tok {
	case 0: // TODO use a real token const
		P.Expr(s.expr);
		P.semi = true;

	case Scanner.CONST, Scanner.TYPE, Scanner.VAR:
		P.Declaration(s.decl, false);

	case Scanner.DEFINE, Scanner.ASSIGN, Scanner.ADD_ASSIGN,
		Scanner.SUB_ASSIGN, Scanner.MUL_ASSIGN, Scanner.QUO_ASSIGN,
		Scanner.REM_ASSIGN, Scanner.AND_ASSIGN, Scanner.OR_ASSIGN,
		Scanner.XOR_ASSIGN, Scanner.SHL_ASSIGN, Scanner.SHR_ASSIGN:
		P.Expr(s.lhs);
		P.Blank();
		P.Token(s.pos, s.tok);
		P.Blank();
		P.Expr(s.expr);
		P.semi = true;

	case Scanner.INC, Scanner.DEC:
		P.Expr(s.expr);
		P.Token(s.pos, s.tok);
		P.semi = true;

	case Scanner.LBRACE:
		P.Block(s.block, true);

	case Scanner.IF:
		P.String(s.pos, "if");
		P.ControlClause(s);
		P.Block(s.block, true);
		if s.post != nil {
			P.newl = false;
			P.String(0, " else ");
			P.Stat(s.post);
		}

	case Scanner.FOR:
		P.String(s.pos, "for");
		P.ControlClause(s);
		P.Block(s.block, true);

	case Scanner.SWITCH, Scanner.SELECT:
		P.Token(s.pos, s.tok);
		P.ControlClause(s);
		P.Block(s.block, false);

	case Scanner.CASE, Scanner.DEFAULT:
		P.Token(s.pos, s.tok);
		if s.expr != nil {
			P.Blank();
			P.Expr(s.expr);
		}
		P.String(0, ":");
		P.OpenScope("");
		P.StatementList(s.block);
		P.CloseScope("");

	case Scanner.GO, Scanner.RETURN, Scanner.BREAK, Scanner.CONTINUE, Scanner.GOTO:
		P.Token(s.pos, s.tok);
		P.Blank();
		P.Expr(s.expr);
		P.semi = true;

	default:
		P.String(s.pos, "<stat>");
		P.semi = true;
	}
}


// ----------------------------------------------------------------------------
// Declarations


/*
func (P *Printer) DoFuncDecl(x *AST.FuncDecl) {
	P.String("func ");
	if x.typ.recv != nil {
		P.String("(");
		P.DoVarDeclList(x.typ.recv);
		P.String(") ");
	}
	P.DoIdent(x.ident);
	P.DoFunctionType(x.typ);
	if x.body != nil {
		P.String(" ");
		P.DoBlock(x.body);
	} else {
		P.String(" ;");
	}
	P.NewLine();
	P.NewLine();

}


func (P *Printer) DoMethodDecl(x *AST.MethodDecl) {
	//P.DoIdent(x.ident);
	//P.DoFunctionType(x.typ);
}
*/


func (P *Printer) Declaration(d *Node.Decl, parenthesized bool) {
	if d == nil {  // TODO remove this check
		P.String(0, "<nil decl>");
		return;
	}

	if !parenthesized {
		if d.exported {
			P.String(0, "export ");
		}
		P.Token(d.pos, d.tok);
		P.Blank();
	}

	if d.ident == nil {
		P.OpenScope("(");
		for i := 0; i < d.list.len(); i++ {
			P.Declaration(d.list.at(i).(*Node.Decl), true);
			P.newl, P.semi = true, true;
		}
		P.CloseScope(")");

	} else {
		P.Expr(d.ident);
		if d.typ != nil {
			P.Blank();
			P.Type(d.typ);
		}
		if d.val != nil {
			if d.tok == Scanner.IMPORT {
				P.Blank();
			} else {
				P.String(0, " = ");
			}
			P.Expr(d.val);
		}
		if d.list != nil {
			if d.tok != Scanner.FUNC {
				panic("must be a func declaration");
			}
			P.Blank();
			P.Block(d.list, true);
		}
	}

	// extra newline at the top level
	if P.level == 0 {
		P.NewLine();
	}

	P.newl = true;
}


// ----------------------------------------------------------------------------
// Program

func (P *Printer) Program(p *Node.Program) {
	P.String(p.pos, "package ");
	P.Expr(p.ident);
	P.NewLine();
	for i := 0; i < p.decls.len(); i++ {
		P.Declaration(p.decls.at(i), false);
	}
	P.newl = true;
	P.String(0, "");  // flush
}
