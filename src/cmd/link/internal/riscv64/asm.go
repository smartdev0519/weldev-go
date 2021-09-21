// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package riscv64

import (
	"cmd/internal/obj/riscv"
	"cmd/internal/objabi"
	"cmd/internal/sys"
	"cmd/link/internal/ld"
	"cmd/link/internal/loader"
	"cmd/link/internal/sym"
	"debug/elf"
	"fmt"
	"log"
	"sort"
)

// fakeLabelName matches the RISCV_FAKE_LABEL_NAME from binutils.
const fakeLabelName = ".L0 "

func gentext(ctxt *ld.Link, ldr *loader.Loader) {
}

func genSymsLate(ctxt *ld.Link, ldr *loader.Loader) {
	if ctxt.LinkMode != ld.LinkExternal {
		return
	}

	// Generate a local text symbol for each relocation target, as the
	// R_RISCV_PCREL_LO12_* relocations generated by elfreloc1 need it.
	if ctxt.Textp == nil {
		log.Fatal("genSymsLate called before Textp has been assigned")
	}
	var hi20Syms []loader.Sym
	for _, s := range ctxt.Textp {
		relocs := ldr.Relocs(s)
		for ri := 0; ri < relocs.Count(); ri++ {
			r := relocs.At(ri)
			if r.Type() != objabi.R_RISCV_PCREL_ITYPE && r.Type() != objabi.R_RISCV_PCREL_STYPE &&
				r.Type() != objabi.R_RISCV_TLS_IE_ITYPE && r.Type() != objabi.R_RISCV_TLS_IE_STYPE {
				continue
			}
			if r.Off() == 0 && ldr.SymType(s) == sym.STEXT {
				// Use the symbol for the function instead of creating
				// an overlapping symbol.
				continue
			}

			// TODO(jsing): Consider generating ELF symbols without needing
			// loader symbols, in order to reduce memory consumption. This
			// would require changes to genelfsym so that it called
			// putelfsym and putelfsyment as appropriate.
			sb := ldr.MakeSymbolBuilder(fakeLabelName)
			sb.SetType(sym.STEXT)
			sb.SetValue(ldr.SymValue(s) + int64(r.Off()))
			sb.SetLocal(true)
			sb.SetReachable(true)
			sb.SetVisibilityHidden(true)
			sb.SetSect(ldr.SymSect(s))
			if outer := ldr.OuterSym(s); outer != 0 {
				ldr.AddInteriorSym(outer, sb.Sym())
			}
			hi20Syms = append(hi20Syms, sb.Sym())
		}
	}
	ctxt.Textp = append(ctxt.Textp, hi20Syms...)
	ldr.SortSyms(ctxt.Textp)
}

func findHI20Symbol(ctxt *ld.Link, ldr *loader.Loader, val int64) loader.Sym {
	idx := sort.Search(len(ctxt.Textp), func(i int) bool { return ldr.SymValue(ctxt.Textp[i]) >= val })
	if idx >= len(ctxt.Textp) {
		return 0
	}
	if s := ctxt.Textp[idx]; ldr.SymValue(s) == val && ldr.SymType(s) == sym.STEXT {
		return s
	}
	return 0
}

func elfreloc1(ctxt *ld.Link, out *ld.OutBuf, ldr *loader.Loader, s loader.Sym, r loader.ExtReloc, ri int, sectoff int64) bool {
	elfsym := ld.ElfSymForReloc(ctxt, r.Xsym)
	switch r.Type {
	case objabi.R_ADDR, objabi.R_DWARFSECREF:
		out.Write64(uint64(sectoff))
		switch r.Size {
		case 4:
			out.Write64(uint64(elf.R_RISCV_32) | uint64(elfsym)<<32)
		case 8:
			out.Write64(uint64(elf.R_RISCV_64) | uint64(elfsym)<<32)
		default:
			ld.Errorf(nil, "unknown size %d for %v relocation", r.Size, r.Type)
			return false
		}
		out.Write64(uint64(r.Xadd))

	case objabi.R_CALLRISCV:
		// Call relocations are currently handled via R_RISCV_PCREL_ITYPE.
		// TODO(jsing): Consider generating elf.R_RISCV_CALL instead of a
		// HI20/LO12_I pair.

	case objabi.R_RISCV_PCREL_ITYPE, objabi.R_RISCV_PCREL_STYPE, objabi.R_RISCV_TLS_IE_ITYPE, objabi.R_RISCV_TLS_IE_STYPE:
		// Find the text symbol for the AUIPC instruction targeted
		// by this relocation.
		relocs := ldr.Relocs(s)
		offset := int64(relocs.At(ri).Off())
		hi20Sym := findHI20Symbol(ctxt, ldr, ldr.SymValue(s)+offset)
		if hi20Sym == 0 {
			ld.Errorf(nil, "failed to find text symbol for HI20 relocation at %d (%x)", sectoff, ldr.SymValue(s)+offset)
			return false
		}
		hi20ElfSym := ld.ElfSymForReloc(ctxt, hi20Sym)

		// Emit two relocations - a R_RISCV_PCREL_HI20 relocation and a
		// corresponding R_RISCV_PCREL_LO12_I or R_RISCV_PCREL_LO12_S relocation.
		// Note that the LO12 relocation must point to a target that has a valid
		// HI20 PC-relative relocation text symbol, which in turn points to the
		// given symbol. For further details see the ELF specification for RISC-V:
		//
		//   https://github.com/riscv/riscv-elf-psabi-doc/blob/master/riscv-elf.md#pc-relative-symbol-addresses
		//
		var hiRel, loRel elf.R_RISCV
		switch r.Type {
		case objabi.R_RISCV_PCREL_ITYPE:
			hiRel, loRel = elf.R_RISCV_PCREL_HI20, elf.R_RISCV_PCREL_LO12_I
		case objabi.R_RISCV_PCREL_STYPE:
			hiRel, loRel = elf.R_RISCV_PCREL_HI20, elf.R_RISCV_PCREL_LO12_S
		case objabi.R_RISCV_TLS_IE_ITYPE:
			hiRel, loRel = elf.R_RISCV_TLS_GOT_HI20, elf.R_RISCV_PCREL_LO12_I
		case objabi.R_RISCV_TLS_IE_STYPE:
			hiRel, loRel = elf.R_RISCV_TLS_GOT_HI20, elf.R_RISCV_PCREL_LO12_S
		}
		out.Write64(uint64(sectoff))
		out.Write64(uint64(hiRel) | uint64(elfsym)<<32)
		out.Write64(uint64(r.Xadd))
		out.Write64(uint64(sectoff + 4))
		out.Write64(uint64(loRel) | uint64(hi20ElfSym)<<32)
		out.Write64(uint64(0))

	default:
		return false
	}

	return true
}

func elfsetupplt(ctxt *ld.Link, plt, gotplt *loader.SymbolBuilder, dynamic loader.Sym) {
	log.Fatalf("elfsetupplt")
}

func machoreloc1(*sys.Arch, *ld.OutBuf, *loader.Loader, loader.Sym, loader.ExtReloc, int64) bool {
	log.Fatalf("machoreloc1 not implemented")
	return false
}

func archreloc(target *ld.Target, ldr *loader.Loader, syms *ld.ArchSyms, r loader.Reloc, s loader.Sym, val int64) (o int64, nExtReloc int, ok bool) {
	if target.IsExternal() {
		switch r.Type() {
		case objabi.R_CALLRISCV:
			return val, 0, true

		case objabi.R_RISCV_PCREL_ITYPE, objabi.R_RISCV_PCREL_STYPE, objabi.R_RISCV_TLS_IE_ITYPE, objabi.R_RISCV_TLS_IE_STYPE:
			return val, 2, true
		}

		return val, 0, false
	}

	rs := r.Sym()

	switch r.Type() {
	case objabi.R_CALLRISCV:
		// Nothing to do.
		return val, 0, true

	case objabi.R_RISCV_TLS_IE_ITYPE, objabi.R_RISCV_TLS_IE_STYPE:
		// TLS relocations are not currently handled for internal linking.
		// For now, TLS is only used when cgo is in use and cgo currently
		// requires external linking. However, we need to accept these
		// relocations so that code containing TLS variables will link,
		// even when they're not being used. For now, replace these
		// instructions with EBREAK to detect accidental use.
		const ebreakIns = 0x00100073
		return ebreakIns<<32 | ebreakIns, 0, true

	case objabi.R_RISCV_PCREL_ITYPE, objabi.R_RISCV_PCREL_STYPE:
		pc := ldr.SymValue(s) + int64(r.Off())
		off := ldr.SymValue(rs) + r.Add() - pc

		// Generate AUIPC and second instruction immediates.
		low, high, err := riscv.Split32BitImmediate(off)
		if err != nil {
			ldr.Errorf(s, "R_RISCV_PCREL_ relocation does not fit in 32 bits: %d", off)
		}

		auipcImm, err := riscv.EncodeUImmediate(high)
		if err != nil {
			ldr.Errorf(s, "cannot encode R_RISCV_PCREL_ AUIPC relocation offset for %s: %v", ldr.SymName(rs), err)
		}

		var secondImm, secondImmMask int64
		switch r.Type() {
		case objabi.R_RISCV_PCREL_ITYPE:
			secondImmMask = riscv.ITypeImmMask
			secondImm, err = riscv.EncodeIImmediate(low)
			if err != nil {
				ldr.Errorf(s, "cannot encode R_RISCV_PCREL_ITYPE I-type instruction relocation offset for %s: %v", ldr.SymName(rs), err)
			}
		case objabi.R_RISCV_PCREL_STYPE:
			secondImmMask = riscv.STypeImmMask
			secondImm, err = riscv.EncodeSImmediate(low)
			if err != nil {
				ldr.Errorf(s, "cannot encode R_RISCV_PCREL_STYPE S-type instruction relocation offset for %s: %v", ldr.SymName(rs), err)
			}
		default:
			panic(fmt.Sprintf("Unknown relocation type: %v", r.Type()))
		}

		auipc := int64(uint32(val))
		second := int64(uint32(val >> 32))

		auipc = (auipc &^ riscv.UTypeImmMask) | int64(uint32(auipcImm))
		second = (second &^ secondImmMask) | int64(uint32(secondImm))

		return second<<32 | auipc, 0, true
	}

	return val, 0, false
}

func archrelocvariant(*ld.Target, *loader.Loader, loader.Reloc, sym.RelocVariant, loader.Sym, int64, []byte) int64 {
	log.Fatalf("archrelocvariant")
	return -1
}

func extreloc(target *ld.Target, ldr *loader.Loader, r loader.Reloc, s loader.Sym) (loader.ExtReloc, bool) {
	switch r.Type() {
	case objabi.R_RISCV_PCREL_ITYPE, objabi.R_RISCV_PCREL_STYPE, objabi.R_RISCV_TLS_IE_ITYPE, objabi.R_RISCV_TLS_IE_STYPE:
		return ld.ExtrelocViaOuterSym(ldr, r, s), true
	}
	return loader.ExtReloc{}, false
}
