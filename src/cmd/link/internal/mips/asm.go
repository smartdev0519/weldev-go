// Inferno utils/5l/asm.c
// https://bitbucket.org/inferno-os/inferno-os/src/default/utils/5l/asm.c
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
//	Portions Copyright © 2016 The Go Authors. All rights reserved.
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

package mips

import (
	"cmd/internal/objabi"
	"cmd/internal/sys"
	"cmd/link/internal/ld"
	"cmd/link/internal/loader"
	"cmd/link/internal/sym"
	"debug/elf"
	"fmt"
)

func gentext(ctxt *ld.Link, ldr *loader.Loader) {
	return
}

func elfreloc1(ctxt *ld.Link, ldr *loader.Loader, s loader.Sym, r loader.ExtRelocView, sectoff int64) bool {
	ctxt.Out.Write32(uint32(sectoff))

	elfsym := ld.ElfSymForReloc(ctxt, r.Xsym)
	switch r.Type() {
	default:
		return false
	case objabi.R_ADDR, objabi.R_DWARFSECREF:
		if r.Siz() != 4 {
			return false
		}
		ctxt.Out.Write32(uint32(elf.R_MIPS_32) | uint32(elfsym)<<8)
	case objabi.R_ADDRMIPS:
		ctxt.Out.Write32(uint32(elf.R_MIPS_LO16) | uint32(elfsym)<<8)
	case objabi.R_ADDRMIPSU:
		ctxt.Out.Write32(uint32(elf.R_MIPS_HI16) | uint32(elfsym)<<8)
	case objabi.R_ADDRMIPSTLS:
		ctxt.Out.Write32(uint32(elf.R_MIPS_TLS_TPREL_LO16) | uint32(elfsym)<<8)
	case objabi.R_CALLMIPS, objabi.R_JMPMIPS:
		ctxt.Out.Write32(uint32(elf.R_MIPS_26) | uint32(elfsym)<<8)
	}

	return true
}

func elfsetupplt(ctxt *ld.Link, plt, gotplt *loader.SymbolBuilder, dynamic loader.Sym) {
	return
}

func machoreloc1(*sys.Arch, *ld.OutBuf, *loader.Loader, loader.Sym, loader.ExtRelocView, int64) bool {
	return false
}

func applyrel(arch *sys.Arch, ldr *loader.Loader, rt objabi.RelocType, off int32, s loader.Sym, val int64, t int64) int64 {
	o := arch.ByteOrder.Uint32(ldr.OutData(s)[off:])
	switch rt {
	case objabi.R_ADDRMIPS, objabi.R_ADDRMIPSTLS:
		return int64(o&0xffff0000 | uint32(t)&0xffff)
	case objabi.R_ADDRMIPSU:
		return int64(o&0xffff0000 | uint32((t+(1<<15))>>16)&0xffff)
	case objabi.R_CALLMIPS, objabi.R_JMPMIPS:
		return int64(o&0xfc000000 | uint32(t>>2)&^0xfc000000)
	default:
		return val
	}
}

func archreloc(target *ld.Target, ldr *loader.Loader, syms *ld.ArchSyms, r loader.Reloc2, rr *loader.ExtReloc, s loader.Sym, val int64) (o int64, needExtReloc bool, ok bool) {
	rs := r.Sym()
	rs = ldr.ResolveABIAlias(rs)
	if target.IsExternal() {
		switch r.Type() {
		default:
			return val, false, false

		case objabi.R_ADDRMIPS, objabi.R_ADDRMIPSU:
			// set up addend for eventual relocation via outer symbol.
			rs, off := ld.FoldSubSymbolOffset(ldr, rs)
			rr.Xadd = r.Add() + off
			rst := ldr.SymType(rs)
			if rst != sym.SHOSTOBJ && rst != sym.SDYNIMPORT && ldr.SymSect(rs) == nil {
				ldr.Errorf(s, "missing section for %s", ldr.SymName(rs))
			}
			rr.Xsym = rs
			return applyrel(target.Arch, ldr, r.Type(), r.Off(), s, val, rr.Xadd), true, true

		case objabi.R_ADDRMIPSTLS, objabi.R_CALLMIPS, objabi.R_JMPMIPS:
			rr.Xsym = rs
			rr.Xadd = r.Add()
			return applyrel(target.Arch, ldr, r.Type(), r.Off(), s, val, rr.Xadd), true, true
		}
	}

	const isOk = true
	const noExtReloc = false
	switch rt := r.Type(); rt {
	case objabi.R_ADDRMIPS, objabi.R_ADDRMIPSU:
		t := ldr.SymValue(rs) + r.Add()
		return applyrel(target.Arch, ldr, rt, r.Off(), s, val, t), noExtReloc, isOk
	case objabi.R_CALLMIPS, objabi.R_JMPMIPS:
		t := ldr.SymValue(rs) + r.Add()

		if t&3 != 0 {
			ldr.Errorf(s, "direct call is not aligned: %s %x", ldr.SymName(rs), t)
		}

		// check if target address is in the same 256 MB region as the next instruction
		if (ldr.SymValue(s)+int64(r.Off())+4)&0xf0000000 != (t & 0xf0000000) {
			ldr.Errorf(s, "direct call too far: %s %x", ldr.SymName(rs), t)
		}

		return applyrel(target.Arch, ldr, rt, r.Off(), s, val, t), noExtReloc, isOk
	case objabi.R_ADDRMIPSTLS:
		// thread pointer is at 0x7000 offset from the start of TLS data area
		t := ldr.SymValue(rs) + r.Add() - 0x7000
		if t < -32768 || t >= 32678 {
			ldr.Errorf(s, "TLS offset out of range %d", t)
		}
		return applyrel(target.Arch, ldr, rt, r.Off(), s, val, t), noExtReloc, isOk
	}

	return val, false, false
}

func archrelocvariant(*ld.Target, *loader.Loader, loader.Reloc2, sym.RelocVariant, loader.Sym, int64) int64 {
	return -1
}

func asmb2(ctxt *ld.Link, _ *loader.Loader) {
	/* output symbol table */
	ld.Symsize = 0

	ld.Lcsize = 0
	symo := uint32(0)
	if !*ld.FlagS {
		if !ctxt.IsELF {
			ld.Errorf(nil, "unsupported executable format")
		}
		symo = uint32(ld.Segdwarf.Fileoff + ld.Segdwarf.Filelen)
		symo = uint32(ld.Rnd(int64(symo), int64(*ld.FlagRound)))

		ctxt.Out.SeekSet(int64(symo))
		ld.Asmelfsym(ctxt)
		ctxt.Out.Write(ld.Elfstrdat)

		if ctxt.LinkMode == ld.LinkExternal {
			ld.Elfemitreloc(ctxt)
		}
	}

	ctxt.Out.SeekSet(0)
	switch ctxt.HeadType {
	default:
		ld.Errorf(nil, "unsupported operating system")
	case objabi.Hlinux:
		ld.Asmbelf(ctxt, int64(symo))
	}

	if *ld.FlagC {
		fmt.Printf("textsize=%d\n", ld.Segtext.Filelen)
		fmt.Printf("datsize=%d\n", ld.Segdata.Filelen)
		fmt.Printf("bsssize=%d\n", ld.Segdata.Length-ld.Segdata.Filelen)
		fmt.Printf("symsize=%d\n", ld.Symsize)
		fmt.Printf("lcsize=%d\n", ld.Lcsize)
		fmt.Printf("total=%d\n", ld.Segtext.Filelen+ld.Segdata.Length+uint64(ld.Symsize)+uint64(ld.Lcsize))
	}
}
