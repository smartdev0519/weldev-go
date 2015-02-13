// Derived from Inferno utils/6c/gc.h
// http://code.google.com/p/inferno-os/source/browse/utils/6c/gc.h
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

#define	Z	N
#define	Adr	Addr

#define	BLOAD(r)	band(bnot(r->refbehind), r->refahead)
#define	BSTORE(r)	band(bnot(r->calbehind), r->calahead)
#define	LOAD(r)		(~r->refbehind.b[z] & r->refahead.b[z])
#define	STORE(r)	(~r->calbehind.b[z] & r->calahead.b[z])

enum
{
	CLOAD = 5,
	CREF = 5,
	CINF = 1000,
	LOOP = 3,
};

typedef	struct	Reg	Reg;
typedef	struct	Rgn	Rgn;

/*c2go
extern Node *Z;

uint32 BLOAD(Reg*);
uint32 BSTORE(Reg*);
uint64 LOAD(Reg*);
uint64 STORE(Reg*);
*/

// A Reg is a wrapper around a single Prog (one instruction) that holds
// register optimization information while the optimizer runs.
// r->prog is the instruction.
struct	Reg
{
	Bits	set;  		// regopt variables written by this instruction.
	Bits	use1; 		// regopt variables read by prog->from.
	Bits	use2; 		// regopt variables read by prog->to.

	// refahead/refbehind are the regopt variables whose current
	// value may be used in the following/preceding instructions
	// up to a CALL (or the value is clobbered).
	Bits	refbehind;
	Bits	refahead;
	// calahead/calbehind are similar, but for variables in
	// instructions that are reachable after hitting at least one
	// CALL.
	Bits	calbehind;
	Bits	calahead;
	Bits	regdiff;
	Bits	act;

	uint64	regu;		// register used bitmap
};
#define	R	((Reg*)0)
/*c2go extern Reg *R; */

#define	NRGN	600
/*c2go enum { NRGN = 600 }; */

// A Rgn represents a single regopt variable over a region of code
// where a register could potentially be dedicated to that variable.
// The code encompassed by a Rgn is defined by the flow graph,
// starting at enter, flood-filling forward while varno is refahead
// and backward while varno is refbehind, and following branches.  A
// single variable may be represented by multiple disjoint Rgns and
// each Rgn may choose a different register for that variable.
// Registers are allocated to regions greedily in order of descending
// cost.
struct	Rgn
{
	Flow*	enter;
	short	cost;
	short	varno;
	short	regno;
};

EXTERN	Reg	zreg;
EXTERN	Rgn	region[NRGN];
EXTERN	Rgn*	rgp;
EXTERN	int	nregion;
EXTERN	int	nvar;
EXTERN	uint64	regbits;
EXTERN	Bits	externs;
EXTERN	Bits	params;
EXTERN	Bits	consts;
EXTERN	Bits	addrs;
EXTERN	Bits	ivar;
EXTERN	Bits	ovar;
EXTERN	int	change;
EXTERN	int32	maxnr;

typedef struct OptStats OptStats;
struct OptStats
{
	int32	ncvtreg;
	int32	nspill;
	int32	nreload;
	int32	ndelmov;
	int32	nvar;
	int32	naddr;
};

EXTERN	OptStats ostats;

/*
 * reg.c
 */
void	regopt(Prog*);
void	dumpone(Flow*, int);
void	dumpit(char*, Flow*, int);

/*
 * peep.c
void	peep(Prog*);
void	excise(Flow*);
int	copyu(Prog*, Adr*, Adr*);
 */

/*
 * prog.c

void proginfo(ProgInfo*, Prog*);
 */
