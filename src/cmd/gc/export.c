// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include	"go.h"
#include	"y.tab.h"

void
exportsym(Sym *s)
{
	Dcl *d, *r;

	if(s == S)
		return;
	if(s->export != 0)
		return;
	s->export = 1;

	d = mal(sizeof(*d));
	d->dsym = s;
	d->dnode = N;
	d->lineno = lineno;

	r = exportlist;
	d->back = r->back;
	r->back->forw = d;
	r->back = d;
}

void
makeexportsym(Type *t)
{
	Sym *s;

	if(t->sym == S) {
		exportgen++;
		snprint(namebuf, sizeof(namebuf), "_e%s_%.3ld", filename, exportgen);
		s = lookup(namebuf);
		s->lexical = LATYPE;
		s->otype = t;
		t->sym = s;
	}
}

void
reexport(Type *t)
{
	Sym *s;

	if(t == T)
		fatal("reexport: type nil");

	makeexportsym(t);
	s = t->sym;
	dumpexporttype(s);
}

void
dumpexportconst(Sym *s)
{
	Node *n;
	Type *t;

	if(s->exported != 0)
		return;
	s->exported = 1;

	n = s->oconst;
	if(n == N || n->op != OLITERAL)
		fatal("dumpexportconst: oconst nil: %S", s);

	t = n->type;	// may or may not be specified
	if(t != T)
		reexport(t);

	Bprint(bout, "\tconst ");
	if(s->export != 0)
		Bprint(bout, "!");
	Bprint(bout, "%lS ", s);
	if(t != T)
		Bprint(bout, "%lS ", t->sym);

	switch(n->val.ctype) {
	default:
		fatal("dumpexportconst: unknown ctype: %S", s);
	case CTINT:
	case CTSINT:
	case CTUINT:
		Bprint(bout, "%B\n", n->val.u.xval);
		break;
	case CTBOOL:
		Bprint(bout, "0x%llux\n", n->val.u.bval);
		break;
	case CTFLT:
		Bprint(bout, "%.17e\n", mpgetflt(n->val.u.fval));
		break;
	case CTSTR:
		Bprint(bout, "\"%Z\"\n", n->val.u.sval);
		break;
	}
}

void
dumpexportvar(Sym *s)
{
	Node *n;
	Type *t;

	if(s->exported != 0)
		return;
	s->exported = 1;

	n = s->oname;
	if(n == N || n->type == T) {
		yyerror("variable exported but not defined: %S", s);
		return;
	}

	t = n->type;
	reexport(t);

	Bprint(bout, "\tvar ");
	if(s->export != 0)
		Bprint(bout, "!");
	Bprint(bout, "%lS %lS\n", s, t->sym);
}

void
dumpexporttype(Sym *s)
{
	Type *t, *f;
	Sym *ts;
	int et;

	if(s->exported != 0)
		return;
	s->exported = 1;

	t = s->otype;
	if(t == T) {
		yyerror("type exported but not defined: %S", s);
		return;
	}

	if(t->sym != s)
		fatal("dumpexporttype: cross reference: %S", s);

	et = t->etype;
	switch(et) {
	default:
		if(et < 0 || et >= nelem(types) || types[et] == T)
			fatal("dumpexporttype: basic type: %S %E", s, et);
		/* type 5 */
		Bprint(bout, "\ttype ");
		if(s->export != 0)
			Bprint(bout, "!");
		Bprint(bout, "%lS %d\n", s, et);
		break;

	case TARRAY:
		reexport(t->type);

		/* type 2 */
		Bprint(bout, "\ttype ");
		if(s->export != 0)
			Bprint(bout, "!");
		if(t->bound >= 0)
			Bprint(bout, "%lS [%lud] %lS\n", s, t->bound, t->type->sym);
		else
			Bprint(bout, "%lS [] %lS\n", s, t->type->sym);
		break;

	case TPTR32:
	case TPTR64:
		if(t->type == T)
			fatal("dumpexporttype: ptr %S", s);
		makeexportsym(t->type); /* forw declare */

		/* type 6 */
		Bprint(bout, "\ttype ");
		if(s->export != 0)
			Bprint(bout, "!");
		Bprint(bout, "%lS *%lS\n", s, t->type->sym);

		reexport(t->type);
		break;

	case TFUNC:
		for(f=t->type; f!=T; f=f->down) {
			if(f->etype != TSTRUCT)
				fatal("dumpexporttype: funct not field: %T", f);
			reexport(f);
		}

		/* type 3 */
		Bprint(bout, "\ttype ");
		if(s->export != 0)
			Bprint(bout, "!");
		Bprint(bout, "%lS (", s);
		for(f=t->type; f!=T; f=f->down) {
			if(f != t->type)
				Bprint(bout, " ");
			Bprint(bout, "%lS", f->sym);
		}
		Bprint(bout, ")\n");
		break;

	case TSTRUCT:
	case TINTER:
		for(f=t->type; f!=T; f=f->down) {
			if(f->etype != TFIELD)
				fatal("dumpexporttype: funct not field: %lT", f);
			reexport(f->type);
		}

		/* type 4 */
		Bprint(bout, "\ttype ");
		if(s->export)
			Bprint(bout, "!");
		Bprint(bout, "%lS %c", s, (et==TSTRUCT)? '{': '<');
		for(f=t->type; f!=T; f=f->down) {
			ts = f->type->sym;
			if(f != t->type)
				Bprint(bout, " ");
			Bprint(bout, "%s %lS", f->sym->name, ts);
		}
		Bprint(bout, "%c\n", (et==TSTRUCT)? '}': '>');
		break;

	case TMAP:
		reexport(t->type);
		reexport(t->down);

		/* type 1 */
		Bprint(bout, "\ttype ");
		if(s->export != 0)
			Bprint(bout, "!");
		Bprint(bout, "%lS [%lS] %lS\n", s, t->down->sym, t->type->sym);
		break;

	case TCHAN:
		reexport(t->type);

		/* type 8 */
		Bprint(bout, "\ttype ");
		if(s->export != 0)
			Bprint(bout, "!");
		Bprint(bout, "%lS %d %lS\n", s, t->chan, t->type->sym);
		break;
	}
}

void
dumpe(Sym *s)
{
	switch(s->lexical) {
	default:
		yyerror("unknown export symbol: %S", s, s->lexical);
		break;
	case LPACK:
		yyerror("package export symbol: %S", s);
		break;
	case LATYPE:
	case LBASETYPE:
		dumpexporttype(s);
		break;
	case LNAME:
		dumpexportvar(s);
		break;
	case LACONST:
		dumpexportconst(s);
		break;
	}
}

void
dumpm(Sym *s)
{
	Type *t, *f;

	switch(s->lexical) {
	default:
		return;

	case LATYPE:
	case LBASETYPE:
		break;
	}

	t = s->otype;
	if(t == T) {
		yyerror("type exported but not defined: %S", s);
		return;
	}

	for(f=t->method; f!=T; f=f->down) {
		if(f->etype != TFIELD)
			fatal("dumpexporttype: method not field: %lT", f);
		reexport(f->type);
		Bprint(bout, "\tfunc %S %lS\n", f->sym, f->type->sym);
	}
}

void
dumpexport(void)
{
	Dcl *d;
	int32 lno;

	lno = lineno;

	Bprint(bout, "   import\n");
	Bprint(bout, "   ((\n");

	Bprint(bout, "    package %s\n", package);

	// first pass dump vars/types depth first
	for(d=exportlist->forw; d!=D; d=d->forw) {
		lineno = d->lineno;
		dumpe(d->dsym);
	}

	// second pass dump methods
	for(d=exportlist->forw; d!=D; d=d->forw) {
		lineno = d->lineno;
		dumpm(d->dsym);
	}

	Bprint(bout, "   ))\n");

	lineno = lno;
}

/*
 * ******* import *******
 */

void
checkimports(void)
{
	Sym *s;
	Type *t, *t1;
	uint32 h;
	int et;

	for(h=0; h<NHASH; h++)
	for(s = hash[h]; s != S; s = s->link) {
		t = s->otype;
		if(t == T)
			continue;

		et = t->etype;
		switch(t->etype) {
		case TFORW:
			print("ci-1: %S %lT\n", s, t);
			break;

		case TPTR32:
		case TPTR64:
			if(t->type == T) {
				print("ci-2: %S %lT\n", s, t);
				break;
			}

			t1 = t->type;
			if(t1 == T) {
				print("ci-3: %S %lT\n", s, t1);
				break;
			}

			et = t1->etype;
			if(et == TFORW) {
				print("%L: ci-4: %S %lT\n", lineno, s, t);
				break;
			}
			break;
		}
	}
}

void
renamepkg(Node *n)
{
	if(n->psym == pkgimportname)
		if(pkgmyname != S)
			n->psym = pkgmyname;
}

Sym*
getimportsym(Node *ss)
{
	char *pkg;
	Sym *s;

	if(ss->op != OIMPORT)
		fatal("getimportsym: oops1 %N", ss);

	pkg = ss->psym->name;
	s = pkglookup(ss->sym->name, pkg);

	/* botch - need some diagnostic checking for the following assignment */
	s->opackage = ss->osym->name;
	return s;
}

Type*
importlooktype(Node *n)
{
	Sym *s;

	s = getimportsym(n);
	if(s->otype == T)
		fatal("importlooktype: oops2 %S", s);
	return s->otype;
}

Type**
importstotype(Node *fl, Type **t, Type *uber)
{
	Type *f;
	Iter save;
	Node *n;

	n = listfirst(&save, &fl);

loop:
	if(n == N) {
		*t = T;
		return t;
	}
	f = typ(TFIELD);
	f->type = importlooktype(n);

	if(n->fsym != S) {
		f->nname = newname(n->fsym);
	} else {
		vargen++;
		snprint(namebuf, sizeof(namebuf), "_m%.3ld", vargen);
		f->nname = newname(lookup(namebuf));
	}
	f->sym = f->nname->sym;

	*t = f;
	t = &f->down;

	n = listnext(&save);
	goto loop;
}

int
importcount(Type *t)
{
	int i;
	Type *f;

	if(t == T || t->etype != TSTRUCT)
		fatal("importcount: not a struct: %N", t);

	i = 0;
	for(f=t->type; f!=T; f=f->down)
		i = i+1;
	return i;
}

void
importfuncnam(Type *t)
{
	Node *n;
	Type *t1;

	if(t->etype != TFUNC)
		fatal("importfuncnam: not func %T", t);

	if(t->thistuple > 0) {
		t1 = t->type;
		if(t1->sym == S)
			fatal("importfuncnam: no this");
		n = newname(t1->sym);
		vargen++;
		n->vargen = vargen;
		t1->nname = n;
	}
	if(t->outtuple > 0) {
		t1 = t->type->down;
		if(t1->sym == S)
			fatal("importfuncnam: no output");
		n = newname(t1->sym);
		vargen++;
		n->vargen = vargen;
		t1->nname = n;
	}
	if(t->intuple > 0) {
		t1 = t->type->down->down;
		if(t1->sym == S)
			fatal("importfuncnam: no input");
		n = newname(t1->sym);
		vargen++;
		n->vargen = vargen;
		t1->nname = n;
	}
}

void
importaddtyp(Node *ss, Type *t)
{
	Sym *s;

	s = getimportsym(ss);
	if(s->otype != T) {
		// here we should try to discover if
		// the new type is the same as the old type
		if(eqtype(t, s->otype, 0))
			return;
		if(isptrto(t, TFORW))
			return;	// hard part
		warn("redeclare import %S from %lT to %lT",
			s, s->otype, t);
		return;
	}
	addtyp(newtype(s), t, PEXTERN);
}

/*
 * LCONST importsym LITERAL
 * untyped constant
 */
void
doimportc1(Node *ss, Val *v)
{
	Node *n;
	Sym *s;

	n = nod(OLITERAL, N, N);
	n->val = *v;

	s = getimportsym(ss);
	if(s->oconst == N) {
		// botch sould ask if already declared the same
		dodclconst(newname(s), n);
	}
}

/*
 * LCONST importsym importsym LITERAL
 * typed constant
 */
void
doimportc2(Node *ss, Node *st, Val *v)
{
	Node *n;
	Type *t;
	Sym *s;

	n = nod(OLITERAL, N, N);
	n->val = *v;

	t = importlooktype(st);
	n->type = t;

	s = getimportsym(ss);
	if(s->oconst == N) {
		// botch sould ask if already declared the same
		dodclconst(newname(s), n);
	}
}

/*
 * LVAR importsym importsym
 * variable
 */
void
doimportv1(Node *ss, Node *st)
{
	Type *t;
	Sym *s;

	t = importlooktype(st);
	s = getimportsym(ss);
	if(s->oname == N || !eqtype(t, s->oname->type, 0)) {
		addvar(newname(s), t, dclcontext);
	}
}

/*
 * LTYPE importsym [ importsym ] importsym
 * array type
 */
void
doimport1(Node *ss, Node *si, Node *st)
{
	Type *t;
	Sym *s;

	t = typ(TMAP);
	s = pkglookup(si->sym->name, si->psym->name);
	t->down = s->otype;
	s = pkglookup(st->sym->name, st->psym->name);
	t->type = s->otype;

	importaddtyp(ss, t);
}

/*
 * LTYPE importsym [ LLITERAL ] importsym
 * array type
 */
void
doimport2(Node *ss, Val *b, Node *st)
{
	Type *t;
	Sym *s;

	t = typ(TARRAY);
	t->bound = -1;
	if(b != nil)
		t->bound = mpgetfix(b->u.xval);
	s = pkglookup(st->sym->name, st->psym->name);
	t->type = s->otype;

	importaddtyp(ss, t);
}

/*
 * LTYPE importsym '(' importsym_list ')'
 * function/method type
 */
void
doimport3(Node *ss, Node *n)
{
	Type *t;

	t = typ(TFUNC);

	t->type = importlooktype(n->left);
	t->type->down = importlooktype(n->right->left);
	t->type->down->down = importlooktype(n->right->right);

	t->thistuple = importcount(t->type);
	t->outtuple = importcount(t->type->down);
	t->intuple = importcount(t->type->down->down);
	dowidth(t);
	importfuncnam(t);

	importaddtyp(ss, t);
}

/*
 * LTYPE importsym '{' importsym_list '}'
 * structure type
 */
void
doimport4(Node *ss, Node *n)
{
	Type *t;

	t = typ(TSTRUCT);
	importstotype(n, &t->type, t);
	dowidth(t);

	importaddtyp(ss, t);
}

/*
 * LTYPE importsym LLITERAL
 * basic type
 */
void
doimport5(Node *ss, Val *v)
{
	int et;
	Type *t;

	et = mpgetfix(v->u.xval);
	if(et <= 0 || et >= nelem(types) || types[et] == T)
		fatal("doimport5: bad type index: %E", et);

	t = typ(et);
	t->sym = S;

	importaddtyp(ss, t);
}

/*
 * LTYPE importsym * importsym
 * pointer type
 */
void
doimport6(Node *ss, Node *st)
{
	Type *t;
	Sym *s;

	s = pkglookup(st->sym->name, st->psym->name);
	t = s->otype;
	if(t == T)
		t = forwdcl(s);
	else
		t = ptrto(t);

	importaddtyp(ss, t);
}

/*
 * LTYPE importsym '<' importsym '>'
 * interface type
 */
void
doimport7(Node *ss, Node *n)
{
	Type *t;

	t = typ(TINTER);
	importstotype(n, &t->type, t);
	dowidth(t);

	importaddtyp(ss, t);
}

/*
 * LTYPE importsym chdir importsym
 * interface type
 */
void
doimport8(Node *ss, Val *v, Node *st)
{
	Type *t;
	Sym *s;
	int dir;

	s = pkglookup(st->sym->name, st->psym->name);
	dir = mpgetfix(v->u.xval);

	t = typ(TCHAN);
	s = pkglookup(st->sym->name, st->psym->name);
	t->type = s->otype;
	t->chan = dir;

	importaddtyp(ss, t);
}

/*
 * LFUNC importsym sym
 * method type
 */
void
doimport9(Sym *sf, Node *ss)
{
	Sym *sfun;

	sfun = getimportsym(ss);
	addmethod(newname(sf), sfun->otype, 0);
}
