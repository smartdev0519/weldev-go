// Derived from Inferno utils/6c/peep.c
// http://code.google.com/p/inferno-os/source/browse/utils/6c/peep.c
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

package main

import (
	"cmd/internal/gc"
	"cmd/internal/obj"
	"cmd/internal/obj/i386"
	"fmt"
)

const (
	REGEXT      = 0
	exregoffset = i386.REG_DI
)

var gactive uint32

// do we need the carry bit
func needc(p *obj.Prog) bool {
	var info gc.ProgInfo

	for p != nil {
		info = proginfo(p)
		if info.Flags&gc.UseCarry != 0 {
			return true
		}
		if info.Flags&(gc.SetCarry|gc.KillCarry) != 0 {
			return false
		}
		p = p.Link
	}

	return false
}

func rnops(r *gc.Flow) *gc.Flow {
	if r != nil {
		var p *obj.Prog
		var r1 *gc.Flow
		for {
			p = r.Prog
			if p.As != obj.ANOP || p.From.Type != obj.TYPE_NONE || p.To.Type != obj.TYPE_NONE {
				break
			}
			r1 = gc.Uniqs(r)
			if r1 == nil {
				break
			}
			r = r1
		}
	}

	return r
}

func peep(firstp *obj.Prog) {
	g := gc.Flowstart(firstp, nil)
	if g == nil {
		return
	}
	gactive = 0

	// byte, word arithmetic elimination.
	elimshortmov(g)

	// constant propagation
	// find MOV $con,R followed by
	// another MOV $con,R without
	// setting R in the interim
	var p *obj.Prog
	for r := g.Start; r != nil; r = r.Link {
		p = r.Prog
		switch p.As {
		case i386.ALEAL:
			if regtyp(&p.To) {
				if p.From.Sym != nil {
					if p.From.Index == i386.REG_NONE {
						conprop(r)
					}
				}
			}

		case i386.AMOVB,
			i386.AMOVW,
			i386.AMOVL,
			i386.AMOVSS,
			i386.AMOVSD:
			if regtyp(&p.To) {
				if p.From.Type == obj.TYPE_CONST || p.From.Type == obj.TYPE_FCONST {
					conprop(r)
				}
			}
		}
	}

	var r1 *gc.Flow
	var p1 *obj.Prog
	var r *gc.Flow
	var t int
loop1:
	if gc.Debug['P'] != 0 && gc.Debug['v'] != 0 {
		gc.Dumpit("loop1", g.Start, 0)
	}

	t = 0
	for r = g.Start; r != nil; r = r.Link {
		p = r.Prog
		switch p.As {
		case i386.AMOVL,
			i386.AMOVSS,
			i386.AMOVSD:
			if regtyp(&p.To) {
				if regtyp(&p.From) {
					if copyprop(g, r) {
						excise(r)
						t++
					} else if subprop(r) && copyprop(g, r) {
						excise(r)
						t++
					}
				}
			}

		case i386.AMOVBLZX,
			i386.AMOVWLZX,
			i386.AMOVBLSX,
			i386.AMOVWLSX:
			if regtyp(&p.To) {
				r1 = rnops(gc.Uniqs(r))
				if r1 != nil {
					p1 = r1.Prog
					if p.As == p1.As && p.To.Type == p1.From.Type && p.To.Reg == p1.From.Reg {
						p1.As = i386.AMOVL
						t++
					}
				}
			}

		case i386.AADDL,
			i386.AADDW:
			if p.From.Type != obj.TYPE_CONST || needc(p.Link) {
				break
			}
			if p.From.Offset == -1 {
				if p.As == i386.AADDL {
					p.As = i386.ADECL
				} else {
					p.As = i386.ADECW
				}
				p.From = obj.Addr{}
				break
			}

			if p.From.Offset == 1 {
				if p.As == i386.AADDL {
					p.As = i386.AINCL
				} else {
					p.As = i386.AINCW
				}
				p.From = obj.Addr{}
				break
			}

		case i386.ASUBL,
			i386.ASUBW:
			if p.From.Type != obj.TYPE_CONST || needc(p.Link) {
				break
			}
			if p.From.Offset == -1 {
				if p.As == i386.ASUBL {
					p.As = i386.AINCL
				} else {
					p.As = i386.AINCW
				}
				p.From = obj.Addr{}
				break
			}

			if p.From.Offset == 1 {
				if p.As == i386.ASUBL {
					p.As = i386.ADECL
				} else {
					p.As = i386.ADECW
				}
				p.From = obj.Addr{}
				break
			}
		}
	}

	if t != 0 {
		goto loop1
	}

	// MOVSD removal.
	// We never use packed registers, so a MOVSD between registers
	// can be replaced by MOVAPD, which moves the pair of float64s
	// instead of just the lower one.  We only use the lower one, but
	// the processor can do better if we do moves using both.
	for r := g.Start; r != nil; r = r.Link {
		p = r.Prog
		if p.As == i386.AMOVSD {
			if regtyp(&p.From) {
				if regtyp(&p.To) {
					p.As = i386.AMOVAPD
				}
			}
		}
	}

	gc.Flowend(g)
}

func excise(r *gc.Flow) {
	p := r.Prog
	if gc.Debug['P'] != 0 && gc.Debug['v'] != 0 {
		fmt.Printf("%v ===delete===\n", p)
	}

	obj.Nopout(p)

	gc.Ostats.Ndelmov++
}

func regtyp(a *obj.Addr) bool {
	return a.Type == obj.TYPE_REG && (i386.REG_AX <= a.Reg && a.Reg <= i386.REG_DI || i386.REG_X0 <= a.Reg && a.Reg <= i386.REG_X7)
}

// movb elimination.
// movb is simulated by the linker
// when a register other than ax, bx, cx, dx
// is used, so rewrite to other instructions
// when possible.  a movb into a register
// can smash the entire 64-bit register without
// causing any trouble.
func elimshortmov(g *gc.Graph) {
	var p *obj.Prog

	for r := g.Start; r != nil; r = r.Link {
		p = r.Prog
		if regtyp(&p.To) {
			switch p.As {
			case i386.AINCB,
				i386.AINCW:
				p.As = i386.AINCL

			case i386.ADECB,
				i386.ADECW:
				p.As = i386.ADECL

			case i386.ANEGB,
				i386.ANEGW:
				p.As = i386.ANEGL

			case i386.ANOTB,
				i386.ANOTW:
				p.As = i386.ANOTL
			}

			if regtyp(&p.From) || p.From.Type == obj.TYPE_CONST {
				// move or artihmetic into partial register.
				// from another register or constant can be movl.
				// we don't switch to 32-bit arithmetic if it can
				// change how the carry bit is set (and the carry bit is needed).
				switch p.As {
				case i386.AMOVB,
					i386.AMOVW:
					p.As = i386.AMOVL

				case i386.AADDB,
					i386.AADDW:
					if !needc(p.Link) {
						p.As = i386.AADDL
					}

				case i386.ASUBB,
					i386.ASUBW:
					if !needc(p.Link) {
						p.As = i386.ASUBL
					}

				case i386.AMULB,
					i386.AMULW:
					p.As = i386.AMULL

				case i386.AIMULB,
					i386.AIMULW:
					p.As = i386.AIMULL

				case i386.AANDB,
					i386.AANDW:
					p.As = i386.AANDL

				case i386.AORB,
					i386.AORW:
					p.As = i386.AORL

				case i386.AXORB,
					i386.AXORW:
					p.As = i386.AXORL

				case i386.ASHLB,
					i386.ASHLW:
					p.As = i386.ASHLL
				}
			} else {
				// explicit zero extension
				switch p.As {
				case i386.AMOVB:
					p.As = i386.AMOVBLZX

				case i386.AMOVW:
					p.As = i386.AMOVWLZX
				}
			}
		}
	}
}

/*
 * the idea is to substitute
 * one register for another
 * from one MOV to another
 *	MOV	a, R0
 *	ADD	b, R0	/ no use of R1
 *	MOV	R0, R1
 * would be converted to
 *	MOV	a, R1
 *	ADD	b, R1
 *	MOV	R1, R0
 * hopefully, then the former or latter MOV
 * will be eliminated by copy propagation.
 */
func subprop(r0 *gc.Flow) bool {
	p := r0.Prog
	v1 := &p.From
	if !regtyp(v1) {
		return false
	}
	v2 := &p.To
	if !regtyp(v2) {
		return false
	}
	var info gc.ProgInfo
	for r := gc.Uniqp(r0); r != nil; r = gc.Uniqp(r) {
		if gc.Debug['P'] != 0 && gc.Debug['v'] != 0 {
			fmt.Printf("\t? %v\n", r.Prog)
		}
		if gc.Uniqs(r) == nil {
			break
		}
		p = r.Prog
		if p.As == obj.AVARDEF || p.As == obj.AVARKILL {
			continue
		}
		info = proginfo(p)
		if info.Flags&gc.Call != 0 {
			return false
		}

		if info.Reguse|info.Regset != 0 {
			return false
		}

		if (info.Flags&gc.Move != 0) && (info.Flags&(gc.SizeL|gc.SizeQ|gc.SizeF|gc.SizeD) != 0) && p.To.Type == v1.Type && p.To.Reg == v1.Reg {
			copysub(&p.To, v1, v2, 1)
			if gc.Debug['P'] != 0 {
				fmt.Printf("gotit: %v->%v\n%v", gc.Ctxt.Dconv(v1), gc.Ctxt.Dconv(v2), r.Prog)
				if p.From.Type == v2.Type && p.From.Reg == v2.Reg {
					fmt.Printf(" excise")
				}
				fmt.Printf("\n")
			}

			for r = gc.Uniqs(r); r != r0; r = gc.Uniqs(r) {
				p = r.Prog
				copysub(&p.From, v1, v2, 1)
				copysub(&p.To, v1, v2, 1)
				if gc.Debug['P'] != 0 {
					fmt.Printf("%v\n", r.Prog)
				}
			}

			t := int(v1.Reg)
			v1.Reg = v2.Reg
			v2.Reg = int16(t)
			if gc.Debug['P'] != 0 {
				fmt.Printf("%v last\n", r.Prog)
			}
			return true
		}

		if copyau(&p.From, v2) || copyau(&p.To, v2) {
			break
		}
		if copysub(&p.From, v1, v2, 0) != 0 || copysub(&p.To, v1, v2, 0) != 0 {
			break
		}
	}

	return false
}

/*
 * The idea is to remove redundant copies.
 *	v1->v2	F=0
 *	(use v2	s/v2/v1/)*
 *	set v1	F=1
 *	use v2	return fail
 *	-----------------
 *	v1->v2	F=0
 *	(use v2	s/v2/v1/)*
 *	set v1	F=1
 *	set v2	return success
 */
func copyprop(g *gc.Graph, r0 *gc.Flow) bool {
	p := r0.Prog
	v1 := &p.From
	v2 := &p.To
	if copyas(v1, v2) {
		return true
	}
	gactive++
	return copy1(v1, v2, r0.S1, 0)
}

func copy1(v1 *obj.Addr, v2 *obj.Addr, r *gc.Flow, f int) bool {
	if uint32(r.Active) == gactive {
		if gc.Debug['P'] != 0 {
			fmt.Printf("act set; return 1\n")
		}
		return true
	}

	r.Active = int32(gactive)
	if gc.Debug['P'] != 0 {
		fmt.Printf("copy %v->%v f=%d\n", gc.Ctxt.Dconv(v1), gc.Ctxt.Dconv(v2), f)
	}
	var t int
	var p *obj.Prog
	for ; r != nil; r = r.S1 {
		p = r.Prog
		if gc.Debug['P'] != 0 {
			fmt.Printf("%v", p)
		}
		if f == 0 && gc.Uniqp(r) == nil {
			f = 1
			if gc.Debug['P'] != 0 {
				fmt.Printf("; merge; f=%d", f)
			}
		}

		t = copyu(p, v2, nil)
		switch t {
		case 2: /* rar, can't split */
			if gc.Debug['P'] != 0 {
				fmt.Printf("; %v rar; return 0\n", gc.Ctxt.Dconv(v2))
			}
			return false

		case 3: /* set */
			if gc.Debug['P'] != 0 {
				fmt.Printf("; %v set; return 1\n", gc.Ctxt.Dconv(v2))
			}
			return true

		case 1, /* used, substitute */
			4: /* use and set */
			if f != 0 {
				if gc.Debug['P'] == 0 {
					return false
				}
				if t == 4 {
					fmt.Printf("; %v used+set and f=%d; return 0\n", gc.Ctxt.Dconv(v2), f)
				} else {
					fmt.Printf("; %v used and f=%d; return 0\n", gc.Ctxt.Dconv(v2), f)
				}
				return false
			}

			if copyu(p, v2, v1) != 0 {
				if gc.Debug['P'] != 0 {
					fmt.Printf("; sub fail; return 0\n")
				}
				return false
			}

			if gc.Debug['P'] != 0 {
				fmt.Printf("; sub %v/%v", gc.Ctxt.Dconv(v2), gc.Ctxt.Dconv(v1))
			}
			if t == 4 {
				if gc.Debug['P'] != 0 {
					fmt.Printf("; %v used+set; return 1\n", gc.Ctxt.Dconv(v2))
				}
				return true
			}
		}

		if f == 0 {
			t = copyu(p, v1, nil)
			if f == 0 && (t == 2 || t == 3 || t == 4) {
				f = 1
				if gc.Debug['P'] != 0 {
					fmt.Printf("; %v set and !f; f=%d", gc.Ctxt.Dconv(v1), f)
				}
			}
		}

		if gc.Debug['P'] != 0 {
			fmt.Printf("\n")
		}
		if r.S2 != nil {
			if !copy1(v1, v2, r.S2, f) {
				return false
			}
		}
	}

	return true
}

/*
 * return
 * 1 if v only used (and substitute),
 * 2 if read-alter-rewrite
 * 3 if set
 * 4 if set and used
 * 0 otherwise (not touched)
 */
func copyu(p *obj.Prog, v *obj.Addr, s *obj.Addr) int {
	switch p.As {
	case obj.AJMP:
		if s != nil {
			if copysub(&p.To, v, s, 1) != 0 {
				return 1
			}
			return 0
		}

		if copyau(&p.To, v) {
			return 1
		}
		return 0

	case obj.ARET:
		if s != nil {
			return 1
		}
		return 3

	case obj.ACALL:
		if REGEXT != 0 /*TypeKind(100016)*/ && v.Type == obj.TYPE_REG && v.Reg <= REGEXT && v.Reg > exregoffset {
			return 2
		}
		if i386.REGARG >= 0 && v.Type == obj.TYPE_REG && v.Reg == i386.REGARG {
			return 2
		}
		if v.Type == p.From.Type && v.Reg == p.From.Reg {
			return 2
		}

		if s != nil {
			if copysub(&p.To, v, s, 1) != 0 {
				return 1
			}
			return 0
		}

		if copyau(&p.To, v) {
			return 4
		}
		return 3

	case obj.ATEXT:
		if i386.REGARG >= 0 && v.Type == obj.TYPE_REG && v.Reg == i386.REGARG {
			return 3
		}
		return 0
	}

	if p.As == obj.AVARDEF || p.As == obj.AVARKILL {
		return 0
	}
	var info gc.ProgInfo
	info = proginfo(p)

	if (info.Reguse|info.Regset)&RtoB(int(v.Reg)) != 0 {
		return 2
	}

	if info.Flags&gc.LeftAddr != 0 {
		if copyas(&p.From, v) {
			return 2
		}
	}

	if info.Flags&(gc.RightRead|gc.RightWrite) == gc.RightRead|gc.RightWrite {
		if copyas(&p.To, v) {
			return 2
		}
	}

	if info.Flags&gc.RightWrite != 0 {
		if copyas(&p.To, v) {
			if s != nil {
				return copysub(&p.From, v, s, 1)
			}
			if copyau(&p.From, v) {
				return 4
			}
			return 3
		}
	}

	if info.Flags&(gc.LeftAddr|gc.LeftRead|gc.LeftWrite|gc.RightAddr|gc.RightRead|gc.RightWrite) != 0 {
		if s != nil {
			if copysub(&p.From, v, s, 1) != 0 {
				return 1
			}
			return copysub(&p.To, v, s, 1)
		}

		if copyau(&p.From, v) {
			return 1
		}
		if copyau(&p.To, v) {
			return 1
		}
	}

	return 0
}

/*
 * direct reference,
 * could be set/use depending on
 * semantics
 */
func copyas(a *obj.Addr, v *obj.Addr) bool {
	if i386.REG_AL <= a.Reg && a.Reg <= i386.REG_BL {
		gc.Fatal("use of byte register")
	}
	if i386.REG_AL <= v.Reg && v.Reg <= i386.REG_BL {
		gc.Fatal("use of byte register")
	}

	if a.Type != v.Type || a.Name != v.Name || a.Reg != v.Reg {
		return false
	}
	if regtyp(v) {
		return true
	}
	if v.Type == obj.TYPE_MEM && (v.Name == obj.NAME_AUTO || v.Name == obj.NAME_PARAM) {
		if v.Offset == a.Offset {
			return true
		}
	}
	return false
}

func sameaddr(a *obj.Addr, v *obj.Addr) bool {
	if a.Type != v.Type || a.Name != v.Name || a.Reg != v.Reg {
		return false
	}
	if regtyp(v) {
		return true
	}
	if v.Type == obj.TYPE_MEM && (v.Name == obj.NAME_AUTO || v.Name == obj.NAME_PARAM) {
		if v.Offset == a.Offset {
			return true
		}
	}
	return false
}

/*
 * either direct or indirect
 */
func copyau(a *obj.Addr, v *obj.Addr) bool {
	if copyas(a, v) {
		return true
	}
	if regtyp(v) {
		if a.Type == obj.TYPE_MEM && a.Reg == v.Reg {
			return true
		}
		if a.Index == v.Reg {
			return true
		}
	}

	return false
}

/*
 * substitute s for v in a
 * return failure to substitute
 */
func copysub(a *obj.Addr, v *obj.Addr, s *obj.Addr, f int) int {
	if copyas(a, v) {
		reg := int(s.Reg)
		if reg >= i386.REG_AX && reg <= i386.REG_DI || reg >= i386.REG_X0 && reg <= i386.REG_X7 {
			if f != 0 {
				a.Reg = int16(reg)
			}
		}

		return 0
	}

	if regtyp(v) {
		reg := int(v.Reg)
		if a.Type == obj.TYPE_MEM && int(a.Reg) == reg {
			if (s.Reg == i386.REG_BP) && a.Index != obj.TYPE_NONE {
				return 1 /* can't use BP-base with index */
			}
			if f != 0 {
				a.Reg = s.Reg
			}
		}

		//			return 0;
		if int(a.Index) == reg {
			if f != 0 {
				a.Index = s.Reg
			}
			return 0
		}

		return 0
	}

	return 0
}

func conprop(r0 *gc.Flow) {
	var p *obj.Prog
	var t int

	p0 := r0.Prog
	v0 := &p0.To
	r := r0

loop:
	r = gc.Uniqs(r)
	if r == nil || r == r0 {
		return
	}
	if gc.Uniqp(r) == nil {
		return
	}

	p = r.Prog
	t = copyu(p, v0, nil)
	switch t {
	case 0, // miss
		1: // use
		goto loop

	case 2, // rar
		4: // use and set
		break

	case 3: // set
		if p.As == p0.As {
			if p.From.Type == p0.From.Type {
				if p.From.Reg == p0.From.Reg {
					if p.From.Node == p0.From.Node {
						if p.From.Offset == p0.From.Offset {
							if p.From.Scale == p0.From.Scale {
								if p.From.Type == obj.TYPE_FCONST && p.From.U.Dval == p0.From.U.Dval {
									if p.From.Index == p0.From.Index {
										excise(r)
										goto loop
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func smallindir(a *obj.Addr, reg *obj.Addr) bool {
	return regtyp(reg) && a.Type == obj.TYPE_MEM && a.Reg == reg.Reg && a.Index == i386.REG_NONE && 0 <= a.Offset && a.Offset < 4096
}

func stackaddr(a *obj.Addr) bool {
	return a.Type == obj.TYPE_REG && a.Reg == i386.REG_SP
}
