// Inferno utils/8a/a.h
// http://code.google.com/p/inferno-os/source/browse/utils/8a/a.h
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
//	Portions Copyright © 2009 The Go Authors.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

#include <lib9.h>
#include <bio.h>
#include "../8c/8.out.h"


#ifndef	EXTERN
#define	EXTERN	extern
#endif

typedef	struct	Sym	Sym;
typedef	struct	Ref	Ref;
typedef	struct	Gen	Gen;
typedef	struct	Io	Io;
typedef	struct	Hist	Hist;
typedef	struct	Gen2 	Gen2;

#define	MAXALIGN	7
#define	FPCHIP		1
#define	NSYMB		500
#define	BUFSIZ		8192
#define	HISTSZ		20
#define	NINCLUDE	10
#define	NHUNK		10000
#define	EOF		(-1)
#define	IGN		(-2)
#define	GETC()		((--fi.c < 0)? filbuf(): *fi.p++ & 0xff)
#define	NHASH		503
#define	STRINGSZ	200
#define	NMACRO		10

struct	Sym
{
	Sym*	link;
	Ref*	ref;
	char*	macro;
	long	value;
	ushort	type;
	char	*name;
	char	sym;
};
#define	S	((Sym*)0)

struct	Ref
{
	int	class;
};

EXTERN struct
{
	char*	p;
	int	c;
} fi;

struct	Io
{
	Io*	link;
	char	b[BUFSIZ];
	char*	p;
	short	c;
	short	f;
};
#define	I	((Io*)0)

EXTERN struct
{
	Sym*	sym;
	short	type;
} h[NSYM];

struct	Gen
{
	double	dval;
	char	sval[8];
	long	offset;
	Sym*	sym;
	short	type;
	short	index;
	short	scale;
};
struct	Gen2
{
	Gen	from;
	Gen	to;
};

struct	Hist
{
	Hist*	link;
	char*	name;
	long	line;
	long	offset;
};
#define	H	((Hist*)0)

enum
{
	CLAST,
	CMACARG,
	CMACRO,
	CPREPROC,
};


EXTERN	char	debug[256];
EXTERN	Sym*	hash[NHASH];
EXTERN	char*	Dlist[30];
EXTERN	int	nDlist;
EXTERN	Hist*	ehist;
EXTERN	int	newflag;
EXTERN	Hist*	hist;
EXTERN	char*	hunk;
EXTERN	char*	include[NINCLUDE];
EXTERN	Io*	iofree;
EXTERN	Io*	ionext;
EXTERN	Io*	iostack;
EXTERN	long	lineno;
EXTERN	int	nerrors;
EXTERN	long	nhunk;
EXTERN	int	ninclude;
EXTERN	Gen	nullgen;
EXTERN	char*	outfile;
EXTERN	int	pass;
EXTERN	char*	pathname;
EXTERN	long	pc;
EXTERN	int	peekc;
EXTERN	int	sym;
EXTERN	char	symb[NSYMB];
EXTERN	int	thechar;
EXTERN	char*	thestring;
EXTERN	long	thunk;
EXTERN	Biobuf	obuf;

void*	allocn(void*, long, long);
void	errorexit(void);
void	pushio(void);
void	newio(void);
void	newfile(char*, int);
Sym*	slookup(char*);
Sym*	lookup(void);
void	syminit(Sym*);
long	yylex(void);
int	getc(void);
int	getnsc(void);
void	unget(int);
int	escchar(int);
void	cinit(void);
void	checkscale(int);
void	pinit(char*);
void	cclean(void);
int	isreg(Gen*);
void	outcode(int, Gen2*);
void	outhist(void);
void	zaddr(Gen*, int);
void	zname(char*, int, int);
void	ieeedtod(Ieee*, double);
int	filbuf(void);
Sym*	getsym(void);
void	domacro(void);
void	macund(void);
void	macdef(void);
void	macexpand(Sym*, char*);
void	macinc(void);
void	macprag(void);
void	maclin(void);
void	macif(int);
void	macend(void);
void	dodefine(char*);
void	prfile(long);
void	linehist(char*, int);
void	gethunk(void);
void	yyerror(char*, ...);
int	yyparse(void);
void	setinclude(char*);
int	assemble(char*);

/*
 *	Posix.c/Inferno.c/Nt.c
 */
enum	/* keep in synch with ../cc/cc.h */
{
	Plan9	= 1<<0,
	Unix	= 1<<1,
	Windows	= 1<<2
};
int	mywait(int*);
int	mycreat(char*, int);
int	systemtype(int);
int	pathchar(void);
char*	mygetwd(char*, int);
int	myexec(char*, char*[]);
int	mydup(int, int);
int	myfork(void);
int	mypipe(int*);
void*	mysbrk(ulong);
