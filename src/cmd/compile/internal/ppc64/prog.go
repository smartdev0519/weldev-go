// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ppc64

import (
	"cmd/compile/internal/gc"
	"cmd/internal/obj"
	"cmd/internal/obj/ppc64"
)

const (
	LeftRdwr  uint32 = gc.LeftRead | gc.LeftWrite
	RightRdwr uint32 = gc.RightRead | gc.RightWrite
)

// This table gives the basic information about instruction
// generated by the compiler and processed in the optimizer.
// See opt.h for bit definitions.
//
// Instructions not generated need not be listed.
// As an exception to that rule, we typically write down all the
// size variants of an operation even if we just use a subset.
//
// The table is formatted for 8-space tabs.
var progtable = [ppc64.ALAST]obj.ProgInfo{
	obj.ATYPE:     {Flags: gc.Pseudo | gc.Skip},
	obj.ATEXT:     {Flags: gc.Pseudo},
	obj.AFUNCDATA: {Flags: gc.Pseudo},
	obj.APCDATA:   {Flags: gc.Pseudo},
	obj.AUNDEF:    {Flags: gc.Break},
	obj.AUSEFIELD: {Flags: gc.OK},
	obj.ACHECKNIL: {Flags: gc.LeftRead},
	obj.AVARDEF:   {Flags: gc.Pseudo | gc.RightWrite},
	obj.AVARKILL:  {Flags: gc.Pseudo | gc.RightWrite},

	// NOP is an internal no-op that also stands
	// for USED and SET annotations, not the Power opcode.
	obj.ANOP: {Flags: gc.LeftRead | gc.RightWrite},

	// Integer
	ppc64.AADD:    {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.ASUB:    {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.ANEG:    {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AAND:    {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AOR:     {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AXOR:    {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AMULLD:  {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AMULLW:  {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AMULHD:  {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AMULHDU: {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.ADIVD:   {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.ADIVDU:  {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.ASLD:    {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.ASRD:    {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.ASRAD:   {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.ACMP:    {Flags: gc.SizeQ | gc.LeftRead | gc.RightRead},
	ppc64.ACMPU:   {Flags: gc.SizeQ | gc.LeftRead | gc.RightRead},
	ppc64.ATD:     {Flags: gc.SizeQ | gc.RightRead},

	// Floating point.
	ppc64.AFADD:   {Flags: gc.SizeD | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AFADDS:  {Flags: gc.SizeF | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AFSUB:   {Flags: gc.SizeD | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AFSUBS:  {Flags: gc.SizeF | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AFMUL:   {Flags: gc.SizeD | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AFMULS:  {Flags: gc.SizeF | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AFDIV:   {Flags: gc.SizeD | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AFDIVS:  {Flags: gc.SizeF | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AFCTIDZ: {Flags: gc.SizeF | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AFCFID:  {Flags: gc.SizeF | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AFCMPU:  {Flags: gc.SizeD | gc.LeftRead | gc.RightRead},
	ppc64.AFRSP:   {Flags: gc.SizeD | gc.LeftRead | gc.RightWrite | gc.Conv},

	// Moves
	ppc64.AMOVB:  {Flags: gc.SizeB | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},
	ppc64.AMOVBU: {Flags: gc.SizeB | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv | gc.PostInc},
	ppc64.AMOVBZ: {Flags: gc.SizeB | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},
	ppc64.AMOVH:  {Flags: gc.SizeW | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},
	ppc64.AMOVHU: {Flags: gc.SizeW | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv | gc.PostInc},
	ppc64.AMOVHZ: {Flags: gc.SizeW | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},
	ppc64.AMOVW:  {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},

	// there is no AMOVWU.
	ppc64.AMOVWZU: {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv | gc.PostInc},
	ppc64.AMOVWZ:  {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},
	ppc64.AMOVD:   {Flags: gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Move},
	ppc64.AMOVDU:  {Flags: gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Move | gc.PostInc},
	ppc64.AFMOVS:  {Flags: gc.SizeF | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},
	ppc64.AFMOVD:  {Flags: gc.SizeD | gc.LeftRead | gc.RightWrite | gc.Move},

	// Jumps
	ppc64.ABR:     {Flags: gc.Jump | gc.Break},
	ppc64.ABL:     {Flags: gc.Call},
	ppc64.ABEQ:    {Flags: gc.Cjmp},
	ppc64.ABNE:    {Flags: gc.Cjmp},
	ppc64.ABGE:    {Flags: gc.Cjmp},
	ppc64.ABLT:    {Flags: gc.Cjmp},
	ppc64.ABGT:    {Flags: gc.Cjmp},
	ppc64.ABLE:    {Flags: gc.Cjmp},
	obj.ARET:      {Flags: gc.Break},
	obj.ADUFFZERO: {Flags: gc.Call},
	obj.ADUFFCOPY: {Flags: gc.Call},
}

var initproginfo_initialized int

func initproginfo() {
	var addvariant = []int{V_CC, V_V, V_CC | V_V}

	if initproginfo_initialized != 0 {
		return
	}
	initproginfo_initialized = 1

	// Perform one-time expansion of instructions in progtable to
	// their CC, V, and VCC variants
	var as2 int
	var i int
	var variant int
	for as := int(0); as < len(progtable); as++ {
		if progtable[as].Flags == 0 {
			continue
		}
		variant = as2variant(as)
		for i = 0; i < len(addvariant); i++ {
			as2 = variant2as(as, variant|addvariant[i])
			if as2 != 0 && progtable[as2].Flags == 0 {
				progtable[as2] = progtable[as]
			}
		}
	}
}

func proginfo(p *obj.Prog) {
	initproginfo()

	info := &p.Info
	*info = progtable[p.As]
	if info.Flags == 0 {
		gc.Fatal("proginfo: unknown instruction %v", p)
	}

	if (info.Flags&gc.RegRead != 0) && p.Reg == 0 {
		info.Flags &^= gc.RegRead
		info.Flags |= gc.RightRead /*CanRegRead |*/
	}

	if (p.From.Type == obj.TYPE_MEM || p.From.Type == obj.TYPE_ADDR) && p.From.Reg != 0 {
		info.Regindex |= RtoB(int(p.From.Reg))
		if info.Flags&gc.PostInc != 0 {
			info.Regset |= RtoB(int(p.From.Reg))
		}
	}

	if (p.To.Type == obj.TYPE_MEM || p.To.Type == obj.TYPE_ADDR) && p.To.Reg != 0 {
		info.Regindex |= RtoB(int(p.To.Reg))
		if info.Flags&gc.PostInc != 0 {
			info.Regset |= RtoB(int(p.To.Reg))
		}
	}

	if p.From.Type == obj.TYPE_ADDR && p.From.Sym != nil && (info.Flags&gc.LeftRead != 0) {
		info.Flags &^= gc.LeftRead
		info.Flags |= gc.LeftAddr
	}

	if p.As == obj.ADUFFZERO {
		info.Reguse |= 1<<0 | RtoB(ppc64.REG_R3)
		info.Regset |= RtoB(ppc64.REG_R3)
	}

	if p.As == obj.ADUFFCOPY {
		// TODO(austin) Revisit when duffcopy is implemented
		info.Reguse |= RtoB(ppc64.REG_R3) | RtoB(ppc64.REG_R4) | RtoB(ppc64.REG_R5)

		info.Regset |= RtoB(ppc64.REG_R3) | RtoB(ppc64.REG_R4)
	}
}

// Instruction variants table.  Initially this contains entries only
// for the "base" form of each instruction.  On the first call to
// as2variant or variant2as, we'll add the variants to the table.
var varianttable = [ppc64.ALAST][4]int{
	ppc64.AADD:     [4]int{ppc64.AADD, ppc64.AADDCC, ppc64.AADDV, ppc64.AADDVCC},
	ppc64.AADDC:    [4]int{ppc64.AADDC, ppc64.AADDCCC, ppc64.AADDCV, ppc64.AADDCVCC},
	ppc64.AADDE:    [4]int{ppc64.AADDE, ppc64.AADDECC, ppc64.AADDEV, ppc64.AADDEVCC},
	ppc64.AADDME:   [4]int{ppc64.AADDME, ppc64.AADDMECC, ppc64.AADDMEV, ppc64.AADDMEVCC},
	ppc64.AADDZE:   [4]int{ppc64.AADDZE, ppc64.AADDZECC, ppc64.AADDZEV, ppc64.AADDZEVCC},
	ppc64.AAND:     [4]int{ppc64.AAND, ppc64.AANDCC, 0, 0},
	ppc64.AANDN:    [4]int{ppc64.AANDN, ppc64.AANDNCC, 0, 0},
	ppc64.ACNTLZD:  [4]int{ppc64.ACNTLZD, ppc64.ACNTLZDCC, 0, 0},
	ppc64.ACNTLZW:  [4]int{ppc64.ACNTLZW, ppc64.ACNTLZWCC, 0, 0},
	ppc64.ADIVD:    [4]int{ppc64.ADIVD, ppc64.ADIVDCC, ppc64.ADIVDV, ppc64.ADIVDVCC},
	ppc64.ADIVDU:   [4]int{ppc64.ADIVDU, ppc64.ADIVDUCC, ppc64.ADIVDUV, ppc64.ADIVDUVCC},
	ppc64.ADIVW:    [4]int{ppc64.ADIVW, ppc64.ADIVWCC, ppc64.ADIVWV, ppc64.ADIVWVCC},
	ppc64.ADIVWU:   [4]int{ppc64.ADIVWU, ppc64.ADIVWUCC, ppc64.ADIVWUV, ppc64.ADIVWUVCC},
	ppc64.AEQV:     [4]int{ppc64.AEQV, ppc64.AEQVCC, 0, 0},
	ppc64.AEXTSB:   [4]int{ppc64.AEXTSB, ppc64.AEXTSBCC, 0, 0},
	ppc64.AEXTSH:   [4]int{ppc64.AEXTSH, ppc64.AEXTSHCC, 0, 0},
	ppc64.AEXTSW:   [4]int{ppc64.AEXTSW, ppc64.AEXTSWCC, 0, 0},
	ppc64.AFABS:    [4]int{ppc64.AFABS, ppc64.AFABSCC, 0, 0},
	ppc64.AFADD:    [4]int{ppc64.AFADD, ppc64.AFADDCC, 0, 0},
	ppc64.AFADDS:   [4]int{ppc64.AFADDS, ppc64.AFADDSCC, 0, 0},
	ppc64.AFCFID:   [4]int{ppc64.AFCFID, ppc64.AFCFIDCC, 0, 0},
	ppc64.AFCTID:   [4]int{ppc64.AFCTID, ppc64.AFCTIDCC, 0, 0},
	ppc64.AFCTIDZ:  [4]int{ppc64.AFCTIDZ, ppc64.AFCTIDZCC, 0, 0},
	ppc64.AFCTIW:   [4]int{ppc64.AFCTIW, ppc64.AFCTIWCC, 0, 0},
	ppc64.AFCTIWZ:  [4]int{ppc64.AFCTIWZ, ppc64.AFCTIWZCC, 0, 0},
	ppc64.AFDIV:    [4]int{ppc64.AFDIV, ppc64.AFDIVCC, 0, 0},
	ppc64.AFDIVS:   [4]int{ppc64.AFDIVS, ppc64.AFDIVSCC, 0, 0},
	ppc64.AFMADD:   [4]int{ppc64.AFMADD, ppc64.AFMADDCC, 0, 0},
	ppc64.AFMADDS:  [4]int{ppc64.AFMADDS, ppc64.AFMADDSCC, 0, 0},
	ppc64.AFMOVD:   [4]int{ppc64.AFMOVD, ppc64.AFMOVDCC, 0, 0},
	ppc64.AFMSUB:   [4]int{ppc64.AFMSUB, ppc64.AFMSUBCC, 0, 0},
	ppc64.AFMSUBS:  [4]int{ppc64.AFMSUBS, ppc64.AFMSUBSCC, 0, 0},
	ppc64.AFMUL:    [4]int{ppc64.AFMUL, ppc64.AFMULCC, 0, 0},
	ppc64.AFMULS:   [4]int{ppc64.AFMULS, ppc64.AFMULSCC, 0, 0},
	ppc64.AFNABS:   [4]int{ppc64.AFNABS, ppc64.AFNABSCC, 0, 0},
	ppc64.AFNEG:    [4]int{ppc64.AFNEG, ppc64.AFNEGCC, 0, 0},
	ppc64.AFNMADD:  [4]int{ppc64.AFNMADD, ppc64.AFNMADDCC, 0, 0},
	ppc64.AFNMADDS: [4]int{ppc64.AFNMADDS, ppc64.AFNMADDSCC, 0, 0},
	ppc64.AFNMSUB:  [4]int{ppc64.AFNMSUB, ppc64.AFNMSUBCC, 0, 0},
	ppc64.AFNMSUBS: [4]int{ppc64.AFNMSUBS, ppc64.AFNMSUBSCC, 0, 0},
	ppc64.AFRES:    [4]int{ppc64.AFRES, ppc64.AFRESCC, 0, 0},
	ppc64.AFRSP:    [4]int{ppc64.AFRSP, ppc64.AFRSPCC, 0, 0},
	ppc64.AFRSQRTE: [4]int{ppc64.AFRSQRTE, ppc64.AFRSQRTECC, 0, 0},
	ppc64.AFSEL:    [4]int{ppc64.AFSEL, ppc64.AFSELCC, 0, 0},
	ppc64.AFSQRT:   [4]int{ppc64.AFSQRT, ppc64.AFSQRTCC, 0, 0},
	ppc64.AFSQRTS:  [4]int{ppc64.AFSQRTS, ppc64.AFSQRTSCC, 0, 0},
	ppc64.AFSUB:    [4]int{ppc64.AFSUB, ppc64.AFSUBCC, 0, 0},
	ppc64.AFSUBS:   [4]int{ppc64.AFSUBS, ppc64.AFSUBSCC, 0, 0},
	ppc64.AMTFSB0:  [4]int{ppc64.AMTFSB0, ppc64.AMTFSB0CC, 0, 0},
	ppc64.AMTFSB1:  [4]int{ppc64.AMTFSB1, ppc64.AMTFSB1CC, 0, 0},
	ppc64.AMULHD:   [4]int{ppc64.AMULHD, ppc64.AMULHDCC, 0, 0},
	ppc64.AMULHDU:  [4]int{ppc64.AMULHDU, ppc64.AMULHDUCC, 0, 0},
	ppc64.AMULHW:   [4]int{ppc64.AMULHW, ppc64.AMULHWCC, 0, 0},
	ppc64.AMULHWU:  [4]int{ppc64.AMULHWU, ppc64.AMULHWUCC, 0, 0},
	ppc64.AMULLD:   [4]int{ppc64.AMULLD, ppc64.AMULLDCC, ppc64.AMULLDV, ppc64.AMULLDVCC},
	ppc64.AMULLW:   [4]int{ppc64.AMULLW, ppc64.AMULLWCC, ppc64.AMULLWV, ppc64.AMULLWVCC},
	ppc64.ANAND:    [4]int{ppc64.ANAND, ppc64.ANANDCC, 0, 0},
	ppc64.ANEG:     [4]int{ppc64.ANEG, ppc64.ANEGCC, ppc64.ANEGV, ppc64.ANEGVCC},
	ppc64.ANOR:     [4]int{ppc64.ANOR, ppc64.ANORCC, 0, 0},
	ppc64.AOR:      [4]int{ppc64.AOR, ppc64.AORCC, 0, 0},
	ppc64.AORN:     [4]int{ppc64.AORN, ppc64.AORNCC, 0, 0},
	ppc64.AREM:     [4]int{ppc64.AREM, ppc64.AREMCC, ppc64.AREMV, ppc64.AREMVCC},
	ppc64.AREMD:    [4]int{ppc64.AREMD, ppc64.AREMDCC, ppc64.AREMDV, ppc64.AREMDVCC},
	ppc64.AREMDU:   [4]int{ppc64.AREMDU, ppc64.AREMDUCC, ppc64.AREMDUV, ppc64.AREMDUVCC},
	ppc64.AREMU:    [4]int{ppc64.AREMU, ppc64.AREMUCC, ppc64.AREMUV, ppc64.AREMUVCC},
	ppc64.ARLDC:    [4]int{ppc64.ARLDC, ppc64.ARLDCCC, 0, 0},
	ppc64.ARLDCL:   [4]int{ppc64.ARLDCL, ppc64.ARLDCLCC, 0, 0},
	ppc64.ARLDCR:   [4]int{ppc64.ARLDCR, ppc64.ARLDCRCC, 0, 0},
	ppc64.ARLDMI:   [4]int{ppc64.ARLDMI, ppc64.ARLDMICC, 0, 0},
	ppc64.ARLWMI:   [4]int{ppc64.ARLWMI, ppc64.ARLWMICC, 0, 0},
	ppc64.ARLWNM:   [4]int{ppc64.ARLWNM, ppc64.ARLWNMCC, 0, 0},
	ppc64.ASLD:     [4]int{ppc64.ASLD, ppc64.ASLDCC, 0, 0},
	ppc64.ASLW:     [4]int{ppc64.ASLW, ppc64.ASLWCC, 0, 0},
	ppc64.ASRAD:    [4]int{ppc64.ASRAD, ppc64.ASRADCC, 0, 0},
	ppc64.ASRAW:    [4]int{ppc64.ASRAW, ppc64.ASRAWCC, 0, 0},
	ppc64.ASRD:     [4]int{ppc64.ASRD, ppc64.ASRDCC, 0, 0},
	ppc64.ASRW:     [4]int{ppc64.ASRW, ppc64.ASRWCC, 0, 0},
	ppc64.ASUB:     [4]int{ppc64.ASUB, ppc64.ASUBCC, ppc64.ASUBV, ppc64.ASUBVCC},
	ppc64.ASUBC:    [4]int{ppc64.ASUBC, ppc64.ASUBCCC, ppc64.ASUBCV, ppc64.ASUBCVCC},
	ppc64.ASUBE:    [4]int{ppc64.ASUBE, ppc64.ASUBECC, ppc64.ASUBEV, ppc64.ASUBEVCC},
	ppc64.ASUBME:   [4]int{ppc64.ASUBME, ppc64.ASUBMECC, ppc64.ASUBMEV, ppc64.ASUBMEVCC},
	ppc64.ASUBZE:   [4]int{ppc64.ASUBZE, ppc64.ASUBZECC, ppc64.ASUBZEV, ppc64.ASUBZEVCC},
	ppc64.AXOR:     [4]int{ppc64.AXOR, ppc64.AXORCC, 0, 0},
}

var initvariants_initialized int

func initvariants() {
	if initvariants_initialized != 0 {
		return
	}
	initvariants_initialized = 1

	var j int
	for i := int(0); i < len(varianttable); i++ {
		if varianttable[i][0] == 0 {
			// Instruction has no variants
			varianttable[i][0] = i

			continue
		}

		// Copy base form to other variants
		if varianttable[i][0] == i {
			for j = 0; j < len(varianttable[i]); j++ {
				varianttable[varianttable[i][j]] = varianttable[i]
			}
		}
	}
}

// as2variant returns the variant (V_*) flags of instruction as.
func as2variant(as int) int {
	initvariants()
	for i := int(0); i < len(varianttable[as]); i++ {
		if varianttable[as][i] == as {
			return i
		}
	}
	gc.Fatal("as2variant: instruction %v is not a variant of itself", obj.Aconv(as))
	return 0
}

// variant2as returns the instruction as with the given variant (V_*) flags.
// If no such variant exists, this returns 0.
func variant2as(as int, flags int) int {
	initvariants()
	return varianttable[as][flags]
}
