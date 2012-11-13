// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// go-specific code shared across loaders (5l, 6l, 8l).

#include	"l.h"
#include	"../ld/lib.h"

// accumulate all type information from .6 files.
// check for inconsistencies.

// TODO:
//	generate debugging section in binary.
//	once the dust settles, try to move some code to
//		libmach, so that other linkers and ar can share.

/*
 *	package import data
 */
typedef struct Import Import;
struct Import
{
	Import *hash;	// next in hash table
	char *prefix;	// "type", "var", "func", "const"
	char *name;
	char *def;
	char *file;
};
enum {
	NIHASH = 1024
};
static Import *ihash[NIHASH];
static int nimport;
static void imported(char *pkg, char *import);

static int
hashstr(char *name)
{
	int h;
	char *cp;

	h = 0;
	for(cp = name; *cp; h += *cp++)
		h *= 1119;
	// not if(h < 0) h = ~h, because gcc 4.3 -O2 miscompiles it.
	h &= 0xffffff;
	return h;
}

static Import *
ilookup(char *name)
{
	int h;
	Import *x;

	h = hashstr(name) % NIHASH;
	for(x=ihash[h]; x; x=x->hash)
		if(x->name[0] == name[0] && strcmp(x->name, name) == 0)
			return x;
	x = mal(sizeof *x);
	x->name = strdup(name);
	x->hash = ihash[h];
	ihash[h] = x;
	nimport++;
	return x;
}

static void loadpkgdata(char*, char*, char*, int);
static void loaddynimport(char*, char*, char*, int);
static void loaddynexport(char*, char*, char*, int);
static void loaddynlinker(char*, char*, char*, int);
static int parsemethod(char**, char*, char**);
static int parsepkgdata(char*, char*, char**, char*, char**, char**, char**);

static Sym **dynexp;

void
ldpkg(Biobuf *f, char *pkg, int64 len, char *filename, int whence)
{
	char *data, *p0, *p1, *name;

	if(debug['g'])
		return;

	if((int)len != len) {
		fprint(2, "%s: too much pkg data in %s\n", argv0, filename);
		if(debug['u'])
			errorexit();
		return;
	}
	data = mal(len+1);
	if(Bread(f, data, len) != len) {
		fprint(2, "%s: short pkg read %s\n", argv0, filename);
		if(debug['u'])
			errorexit();
		return;
	}
	data[len] = '\0';

	// first \n$$ marks beginning of exports - skip rest of line
	p0 = strstr(data, "\n$$");
	if(p0 == nil) {
		if(debug['u'] && whence != ArchiveObj) {
			fprint(2, "%s: cannot find export data in %s\n", argv0, filename);
			errorexit();
		}
		return;
	}
	p0 += 3;
	while(*p0 != '\n' && *p0 != '\0')
		p0++;

	// second marks end of exports / beginning of local data
	p1 = strstr(p0, "\n$$");
	if(p1 == nil) {
		fprint(2, "%s: cannot find end of exports in %s\n", argv0, filename);
		if(debug['u'])
			errorexit();
		return;
	}
	while(p0 < p1 && (*p0 == ' ' || *p0 == '\t' || *p0 == '\n'))
		p0++;
	if(p0 < p1) {
		if(strncmp(p0, "package ", 8) != 0) {
			fprint(2, "%s: bad package section in %s - %s\n", argv0, filename, p0);
			if(debug['u'])
				errorexit();
			return;
		}
		p0 += 8;
		while(p0 < p1 && (*p0 == ' ' || *p0 == '\t' || *p0 == '\n'))
			p0++;
		name = p0;
		while(p0 < p1 && *p0 != ' ' && *p0 != '\t' && *p0 != '\n')
			p0++;
		if(debug['u'] && whence != ArchiveObj &&
		   (p0+6 > p1 || memcmp(p0, " safe\n", 6) != 0)) {
			fprint(2, "%s: load of unsafe package %s\n", argv0, filename);
			nerrors++;
			errorexit();
		}
		if(p0 < p1) {
			if(*p0 == '\n')
				*p0++ = '\0';
			else {
				*p0++ = '\0';
				while(p0 < p1 && *p0++ != '\n')
					;
			}
		}
		if(strcmp(pkg, "main") == 0 && strcmp(name, "main") != 0) {
			fprint(2, "%s: %s: not package main (package %s)\n", argv0, filename, name);
			nerrors++;
			errorexit();
		}
		loadpkgdata(filename, pkg, p0, p1 - p0);
	}

	// The __.PKGDEF archive summary has no local types.
	if(whence == Pkgdef)
		return;

	// local types begin where exports end.
	// skip rest of line after $$ we found above
	p0 = p1 + 3;
	while(*p0 != '\n' && *p0 != '\0')
		p0++;

	// local types end at next \n$$.
	p1 = strstr(p0, "\n$$");
	if(p1 == nil) {
		fprint(2, "%s: cannot find end of local types in %s\n", argv0, filename);
		if(debug['u'])
			errorexit();
		return;
	}

	loadpkgdata(filename, pkg, p0, p1 - p0);

	// look for dynimport section
	p0 = strstr(p1, "\n$$  // dynimport");
	if(p0 != nil) {
		p0 = strchr(p0+1, '\n');
		if(p0 == nil) {
			fprint(2, "%s: found $$ // dynimport but no newline in %s\n", argv0, filename);
			if(debug['u'])
				errorexit();
			return;
		}
		p1 = strstr(p0, "\n$$");
		if(p1 == nil)
			p1 = strstr(p0, "\n!\n");
		if(p1 == nil) {
			fprint(2, "%s: cannot find end of // dynimport section in %s\n", argv0, filename);
			if(debug['u'])
				errorexit();
			return;
		}
		loaddynimport(filename, pkg, p0 + 1, p1 - (p0+1));
	}

	// look for dynexp section
	p0 = strstr(p1, "\n$$  // dynexport");
	if(p0 != nil) {
		p0 = strchr(p0+1, '\n');
		if(p0 == nil) {
			fprint(2, "%s: found $$ // dynexport but no newline in %s\n", argv0, filename);
			if(debug['u'])
				errorexit();
			return;
		}
		p1 = strstr(p0, "\n$$");
		if(p1 == nil)
			p1 = strstr(p0, "\n!\n");
		if(p1 == nil) {
			fprint(2, "%s: cannot find end of // dynexport section in %s\n", argv0, filename);
			if(debug['u'])
				errorexit();
			return;
		}
		loaddynexport(filename, pkg, p0 + 1, p1 - (p0+1));
	}

	p0 = strstr(p1, "\n$$  // dynlinker");
	if(p0 != nil) {
		p0 = strchr(p0+1, '\n');
		if(p0 == nil) {
			fprint(2, "%s: found $$ // dynlinker but no newline in %s\n", argv0, filename);
			if(debug['u'])
				errorexit();
			return;
		}
		p1 = strstr(p0, "\n$$");
		if(p1 == nil)
			p1 = strstr(p0, "\n!\n");
		if(p1 == nil) {
			fprint(2, "%s: cannot find end of // dynlinker section in %s\n", argv0, filename);
			if(debug['u'])
				errorexit();
			return;
		}
		loaddynlinker(filename, pkg, p0 + 1, p1 - (p0+1));
	}
}

static void
loadpkgdata(char *file, char *pkg, char *data, int len)
{
	char *p, *ep, *prefix, *name, *def;
	Import *x;

	file = strdup(file);
	p = data;
	ep = data + len;
	while(parsepkgdata(file, pkg, &p, ep, &prefix, &name, &def) > 0) {
		x = ilookup(name);
		if(x->prefix == nil) {
			x->prefix = prefix;
			x->def = strdup(def);
			x->file = file;
		} else if(strcmp(x->prefix, prefix) != 0) {
			fprint(2, "%s: conflicting definitions for %s\n", argv0, name);
			fprint(2, "%s:\t%s %s ...\n", x->file, x->prefix, name);
			fprint(2, "%s:\t%s %s ...\n", file, prefix, name);
			nerrors++;
		} else if(strcmp(x->def, def) != 0) {
			fprint(2, "%s: conflicting definitions for %s\n", argv0, name);
			fprint(2, "%s:\t%s %s %s\n", x->file, x->prefix, name, x->def);
			fprint(2, "%s:\t%s %s %s\n", file, prefix, name, def);
			nerrors++;
		}
		free(name);
		free(def);
	}
	free(file);
}

// replace all "". with pkg.
char*
expandpkg(char *t0, char *pkg)
{
	int n;
	char *p;
	char *w, *w0, *t;

	n = 0;
	for(p=t0; (p=strstr(p, "\"\".")) != nil; p+=3)
		n++;

	if(n == 0)
		return strdup(t0);

	// use malloc, not mal, so that caller can free
	w0 = malloc(strlen(t0) + strlen(pkg)*n);
	if(w0 == nil) {
		diag("out of memory");
		errorexit();
	}
	w = w0;
	for(p=t=t0; (p=strstr(p, "\"\".")) != nil; p=t) {
		memmove(w, t, p - t);
		w += p-t;
		strcpy(w, pkg);
		w += strlen(pkg);
		t = p+2;
	}
	strcpy(w, t);
	return w0;
}

static int
parsepkgdata(char *file, char *pkg, char **pp, char *ep, char **prefixp, char **namep, char **defp)
{
	char *p, *prefix, *name, *def, *edef, *meth;
	int n, inquote;

	// skip white space
	p = *pp;
loop:
	while(p < ep && (*p == ' ' || *p == '\t' || *p == '\n'))
		p++;
	if(p == ep || strncmp(p, "$$\n", 3) == 0)
		return 0;

	// prefix: (var|type|func|const)
	prefix = p;
	if(p + 7 > ep)
		return -1;
	if(strncmp(p, "var ", 4) == 0)
		p += 4;
	else if(strncmp(p, "type ", 5) == 0)
		p += 5;
	else if(strncmp(p, "func ", 5) == 0)
		p += 5;
	else if(strncmp(p, "const ", 6) == 0)
		p += 6;
	else if(strncmp(p, "import ", 7) == 0) {
		p += 7;
		while(p < ep && *p != ' ')
			p++;
		p++;
		name = p;
		while(p < ep && *p != '\n')
			p++;
		if(p >= ep) {
			fprint(2, "%s: %s: confused in import line\n", argv0, file);
			nerrors++;
			return -1;
		}
		*p++ = '\0';
		imported(pkg, name);
		goto loop;
	}
	else {
		fprint(2, "%s: %s: confused in pkg data near <<%.40s>>\n", argv0, file, prefix);
		nerrors++;
		return -1;
	}
	p[-1] = '\0';

	// name: a.b followed by space
	name = p;
	inquote = 0;
	while(p < ep) {
		if (*p == ' ' && !inquote)
			break;

		if(*p == '\\')
			p++;
		else if(*p == '"')
			inquote = !inquote;

		p++;
	}

	if(p >= ep)
		return -1;
	*p++ = '\0';

	// def: free form to new line
	def = p;
	while(p < ep && *p != '\n')
		p++;
	if(p >= ep)
		return -1;
	edef = p;
	*p++ = '\0';

	// include methods on successive lines in def of named type
	while(parsemethod(&p, ep, &meth) > 0) {
		*edef++ = '\n';	// overwrites '\0'
		if(edef+1 > meth) {
			// We want to indent methods with a single \t.
			// 6g puts at least one char of indent before all method defs,
			// so there will be room for the \t.  If the method def wasn't
			// indented we could do something more complicated,
			// but for now just diagnose the problem and assume
			// 6g will keep indenting for us.
			fprint(2, "%s: %s: expected methods to be indented %p %p %.10s\n", argv0,
				file, edef, meth, meth);
			nerrors++;
			return -1;
		}
		*edef++ = '\t';
		n = strlen(meth);
		memmove(edef, meth, n);
		edef += n;
	}

	name = expandpkg(name, pkg);
	def = expandpkg(def, pkg);

	// done
	*pp = p;
	*prefixp = prefix;
	*namep = name;
	*defp = def;
	return 1;
}

static int
parsemethod(char **pp, char *ep, char **methp)
{
	char *p;

	// skip white space
	p = *pp;
	while(p < ep && (*p == ' ' || *p == '\t'))
		p++;
	if(p == ep)
		return 0;

	// might be a comment about the method
	if(p + 2 < ep && strncmp(p, "//", 2) == 0)
		goto useline;
	
	// if it says "func (", it's a method
	if(p + 6 < ep && strncmp(p, "func (", 6) == 0)
		goto useline;
	return 0;

useline:
	// definition to end of line
	*methp = p;
	while(p < ep && *p != '\n')
		p++;
	if(p >= ep) {
		fprint(2, "%s: lost end of line in method definition\n", argv0);
		*pp = ep;
		return -1;
	}
	*p++ = '\0';
	*pp = p;
	return 1;
}

static void
loaddynimport(char *file, char *pkg, char *p, int n)
{
	char *pend, *next, *name, *def, *p0, *lib, *q;
	Sym *s;

	USED(file);
	pend = p + n;
	for(; p<pend; p=next) {
		next = strchr(p, '\n');
		if(next == nil)
			next = "";
		else
			*next++ = '\0';
		p0 = p;
		if(strncmp(p, "dynimport ", 10) != 0)
			goto err;
		p += 10;
		name = p;
		p = strchr(name, ' ');
		if(p == nil)
			goto err;
		while(*p == ' ')
			p++;
		def = p;
		p = strchr(def, ' ');
		if(p == nil)
			goto err;
		while(*p == ' ')
			p++;
		lib = p;

		// successful parse: now can edit the line
		*strchr(name, ' ') = 0;
		*strchr(def, ' ') = 0;
		
		if(debug['d']) {
			fprint(2, "%s: %s: cannot use dynamic imports with -d flag\n", argv0, file);
			nerrors++;
			return;
		}
		
		if(strcmp(name, "_") == 0 && strcmp(def, "_") == 0) {
			// allow #pragma dynimport _ _ "foo.so"
			// to force a link of foo.so.
			havedynamic = 1;
			adddynlib(lib);
			continue;
		}

		name = expandpkg(name, pkg);
		q = strchr(def, '@');
		if(q)
			*q++ = '\0';
		s = lookup(name, 0);
		free(name);
		if(s->type == 0 || s->type == SXREF) {
			s->dynimplib = lib;
			s->dynimpname = def;
			s->dynimpvers = q;
			s->type = SDYNIMPORT;
			havedynamic = 1;
		}
	}
	return;

err:
	fprint(2, "%s: %s: invalid dynimport line: %s\n", argv0, file, p0);
	nerrors++;
}

static void
loaddynexport(char *file, char *pkg, char *p, int n)
{
	char *pend, *next, *local, *elocal, *remote, *p0;
	Sym *s;

	USED(file);
	pend = p + n;
	for(; p<pend; p=next) {
		next = strchr(p, '\n');
		if(next == nil)
			next = "";
		else
			*next++ = '\0';
		p0 = p;
		if(strncmp(p, "dynexport ", 10) != 0)
			goto err;
		p += 10;
		local = p;
		p = strchr(local, ' ');
		if(p == nil)
			goto err;
		while(*p == ' ')
			p++;
		remote = p;

		// successful parse: now can edit the line
		*strchr(local, ' ') = 0;

		elocal = expandpkg(local, pkg);

		s = lookup(elocal, 0);
		if(s->dynimplib != nil) {
			fprint(2, "%s: symbol is both dynimport and dynexport %s\n", argv0, local);
			nerrors++;
		}
		s->dynimpname = remote;
		s->dynexport = 1;

		if(ndynexp%32 == 0)
			dynexp = realloc(dynexp, (ndynexp+32)*sizeof dynexp[0]);
		dynexp[ndynexp++] = s;

		if (elocal != local)
			free(elocal);
	}
	return;

err:
	fprint(2, "%s: invalid dynexport line: %s\n", argv0, p0);
	nerrors++;
}

static void
loaddynlinker(char *file, char *pkg, char *p, int n)
{
	char *pend, *next, *dynlinker, *p0;

	USED(file);
	USED(pkg);
	pend = p + n;
	for(; p<pend; p=next) {
		next = strchr(p, '\n');
		if(next == nil)
			next = "";
		else
			*next++ = '\0';
		p0 = p;
		if(strncmp(p, "dynlinker ", 10) != 0)
			goto err;
		p += 10;
		dynlinker = p;

		if(*dynlinker == '\0')
			goto err;
		if(!debug['I']) { // not overrided by cmdline
			if(interpreter != nil && strcmp(interpreter, dynlinker) != 0) {
				fprint(2, "%s: conflict dynlinker: %s and %s\n", argv0, interpreter, dynlinker);
				nerrors++;
				return;
			}
			free(interpreter);
			interpreter = strdup(dynlinker);
		}
	}
	return;

err:
	fprint(2, "%s: invalid dynlinker line: %s\n", argv0, p0);
	nerrors++;
}

static Sym *markq;
static Sym *emarkq;

static void
mark1(Sym *s, Sym *parent)
{
	if(s == S || s->reachable)
		return;
	if(strncmp(s->name, "go.weak.", 8) == 0)
		return;
	s->reachable = 1;
	s->reachparent = parent;
	if(markq == nil)
		markq = s;
	else
		emarkq->queue = s;
	emarkq = s;
}

void
mark(Sym *s)
{
	mark1(s, nil);
}

static void
markflood(void)
{
	Auto *a;
	Prog *p;
	Sym *s;
	int i;
	
	for(s=markq; s!=S; s=s->queue) {
		if(s->text) {
			if(debug['v'] > 1)
				Bprint(&bso, "marktext %s\n", s->name);
			for(a=s->autom; a; a=a->link)
				mark1(a->gotype, s);
			for(p=s->text; p != P; p=p->link) {
				mark1(p->from.sym, s);
				mark1(p->to.sym, s);
			}
		}
		for(i=0; i<s->nr; i++)
			mark1(s->r[i].sym, s);
		mark1(s->gotype, s);
		mark1(s->sub, s);
		mark1(s->outer, s);
	}
}

static char*
morename[] =
{
	"runtime.morestack",
	"runtime.morestackx",

	"runtime.morestack00",
	"runtime.morestack10",
	"runtime.morestack01",
	"runtime.morestack11",

	"runtime.morestack8",
	"runtime.morestack16",
	"runtime.morestack24",
	"runtime.morestack32",
	"runtime.morestack40",
	"runtime.morestack48",
};

static int
isz(Auto *a)
{
	for(; a; a=a->link)
		if(a->type == D_FILE || a->type == D_FILE1)
			return 1;
	return 0;
}

static void
addz(Sym *s, Auto *z)
{
	Auto *a, *last;

	// strip out non-z
	last = nil;
	for(a = z; a != nil; a = a->link) {
		if(a->type == D_FILE || a->type == D_FILE1) {
			if(last == nil)
				z = a;
			else
				last->link = a;
			last = a;
		}
	}
	if(last) {
		last->link = s->autom;
		s->autom = z;
	}
}

void
deadcode(void)
{
	int i;
	Sym *s, *last, *p;
	Auto *z;
	Fmt fmt;

	if(debug['v'])
		Bprint(&bso, "%5.2f deadcode\n", cputime());

	mark(lookup(INITENTRY, 0));
	for(i=0; i<nelem(morename); i++)
		mark(lookup(morename[i], 0));

	for(i=0; i<ndynexp; i++)
		mark(dynexp[i]);

	markflood();
	
	// keep each beginning with 'typelink.' if the symbol it points at is being kept.
	for(s = allsym; s != S; s = s->allsym) {
		if(strncmp(s->name, "go.typelink.", 12) == 0)
			s->reachable = s->nr==1 && s->r[0].sym->reachable;
	}

	// remove dead text but keep file information (z symbols).
	last = nil;
	z = nil;
	for(s = textp; s != nil; s = s->next) {
		if(!s->reachable) {
			if(isz(s->autom))
				z = s->autom;
			continue;
		}
		if(last == nil)
			textp = s;
		else
			last->next = s;
		last = s;
		if(z != nil) {
			if(!isz(s->autom))
				addz(s, z);
			z = nil;
		}
	}
	if(last == nil)
		textp = nil;
	else
		last->next = nil;
	
	for(s = allsym; s != S; s = s->allsym)
		if(strncmp(s->name, "go.weak.", 8) == 0) {
			s->special = 1;  // do not lay out in data segment
			s->reachable = 1;
			s->hide = 1;
		}
	
	// record field tracking references
	fmtstrinit(&fmt);
	for(s = allsym; s != S; s = s->allsym) {
		if(strncmp(s->name, "go.track.", 9) == 0) {
			s->special = 1;  // do not lay out in data segment
			s->hide = 1;
			if(s->reachable) {
				fmtprint(&fmt, "%s", s->name+9);
				for(p=s->reachparent; p; p=p->reachparent)
					fmtprint(&fmt, "\t%s", p->name);
				fmtprint(&fmt, "\n");
			}
			s->type = SCONST;
			s->value = 0;
		}
	}
	if(tracksym == nil)
		return;
	s = lookup(tracksym, 0);
	if(!s->reachable)
		return;
	addstrdata(tracksym, fmtstrflush(&fmt));
}

void
doweak(void)
{
	Sym *s, *t;

	// resolve weak references only if
	// target symbol will be in binary anyway.
	for(s = allsym; s != S; s = s->allsym) {
		if(strncmp(s->name, "go.weak.", 8) == 0) {
			t = rlookup(s->name+8, s->version);
			if(t && t->type != 0 && t->reachable) {
				s->value = t->value;
				s->type = t->type;
			} else {
				s->type = SCONST;
				s->value = 0;
			}
			continue;
		}
	}
}

void
addexport(void)
{
	int i;
	
	for(i=0; i<ndynexp; i++)
		adddynsym(dynexp[i]);
}

/* %Z from gc, for quoting import paths */
int
Zconv(Fmt *fp)
{
	Rune r;
	char *s, *se;
	int n;

	s = va_arg(fp->args, char*);
	if(s == nil)
		return fmtstrcpy(fp, "<nil>");

	se = s + strlen(s);
	while(s < se) {
		n = chartorune(&r, s);
		s += n;
		switch(r) {
		case Runeerror:
			if(n == 1) {
				fmtprint(fp, "\\x%02x", (uchar)*(s-1));
				break;
			}
			// fall through
		default:
			if(r < ' ') {
				fmtprint(fp, "\\x%02x", r);
				break;
			}
			fmtrune(fp, r);
			break;
		case '\t':
			fmtstrcpy(fp, "\\t");
			break;
		case '\n':
			fmtstrcpy(fp, "\\n");
			break;
		case '\"':
		case '\\':
			fmtrune(fp, '\\');
			fmtrune(fp, r);
			break;
		}
	}
	return 0;
}


typedef struct Pkg Pkg;
struct Pkg
{
	uchar mark;
	uchar checked;
	Pkg *next;
	char *path;
	Pkg **impby;
	int nimpby;
	int mimpby;
	Pkg *all;
};

static Pkg *phash[1024];
static Pkg *pkgall;

static Pkg*
getpkg(char *path)
{
	Pkg *p;
	int h;
	
	h = hashstr(path) % nelem(phash);
	for(p=phash[h]; p; p=p->next)
		if(strcmp(p->path, path) == 0)
			return p;
	p = mal(sizeof *p);
	p->path = strdup(path);
	p->next = phash[h];
	phash[h] = p;
	p->all = pkgall;
	pkgall = p;
	return p;
}

static void
imported(char *pkg, char *import)
{
	Pkg *p, *i;
	
	// everyone imports runtime, even runtime.
	if(strcmp(import, "\"runtime\"") == 0)
		return;

	pkg = smprint("\"%Z\"", pkg);  // turn pkg path into quoted form, freed below
	p = getpkg(pkg);
	i = getpkg(import);
	if(i->nimpby >= i->mimpby) {
		i->mimpby *= 2;
		if(i->mimpby == 0)
			i->mimpby = 16;
		i->impby = realloc(i->impby, i->mimpby*sizeof i->impby[0]);
	}
	i->impby[i->nimpby++] = p;
	free(pkg);
}

static Pkg*
cycle(Pkg *p)
{
	int i;
	Pkg *bad;

	if(p->checked)
		return 0;

	if(p->mark) {
		nerrors++;
		print("import cycle:\n");
		print("\t%s\n", p->path);
		return p;
	}
	p->mark = 1;
	for(i=0; i<p->nimpby; i++) {
		if((bad = cycle(p->impby[i])) != nil) {
			p->mark = 0;
			p->checked = 1;
			print("\timports %s\n", p->path);
			if(bad == p)
				return nil;
			return bad;
		}
	}
	p->checked = 1;
	p->mark = 0;
	return 0;
}

void
importcycles(void)
{
	Pkg *p;
	
	for(p=pkgall; p; p=p->all)
		cycle(p);
}

static int
scmp(const void *p1, const void *p2)
{
	Sym *s1, *s2;

	s1 = *(Sym**)p1;
	s2 = *(Sym**)p2;
	return strcmp(s1->dynimpname, s2->dynimpname);
}
void
sortdynexp(void)
{
	int i;

	// On Mac OS X Mountain Lion, we must sort exported symbols
	// So we sort them here and pre-allocate dynid for them
	// See http://golang.org/issue/4029
	if(HEADTYPE != Hdarwin)
		return;
	qsort(dynexp, ndynexp, sizeof dynexp[0], scmp);
	for(i=0; i<ndynexp; i++) {
		dynexp[i]->dynid = -i-100; // also known to [68]l/asm.c:^adddynsym
	}
}
