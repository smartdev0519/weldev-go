// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package riscv64

import (
	"cmd/internal/objabi"
	"cmd/internal/sys"
	"cmd/link/internal/ld"
	"cmd/link/internal/sym"
	"fmt"
	"log"
)

func gentext(ctxt *ld.Link) {
}

func adddynrela(ctxt *ld.Link, rel *sym.Symbol, s *sym.Symbol, r *sym.Reloc) {
	log.Fatalf("adddynrela not implemented")
}

func adddynrel(ctxt *ld.Link, s *sym.Symbol, r *sym.Reloc) bool {
	log.Fatalf("adddynrel not implemented")
	return false
}

func elfreloc1(ctxt *ld.Link, r *sym.Reloc, sectoff int64) bool {
	log.Fatalf("elfreloc1")
	return false
}

func elfsetupplt(ctxt *ld.Link) {
	log.Fatalf("elfsetuplt")
}

func machoreloc1(arch *sys.Arch, out *ld.OutBuf, s *sym.Symbol, r *sym.Reloc, sectoff int64) bool {
	log.Fatalf("machoreloc1 not implemented")
	return false
}

func archreloc(ctxt *ld.Link, r *sym.Reloc, s *sym.Symbol, val int64) (int64, bool) {
	// TODO(jsing): Implement.
	log.Fatalf("archreloc not implemented")
	return val, false
}

func archrelocvariant(ctxt *ld.Link, r *sym.Reloc, s *sym.Symbol, t int64) int64 {
	log.Fatalf("archrelocvariant")
	return -1
}

func asmb(ctxt *ld.Link) {
	if ctxt.IsELF {
		ld.Asmbelfsetup()
	}

	sect := ld.Segtext.Sections[0]
	ctxt.Out.SeekSet(int64(sect.Vaddr - ld.Segtext.Vaddr + ld.Segtext.Fileoff))
	ld.Codeblk(ctxt, int64(sect.Vaddr), int64(sect.Length))
	for _, sect = range ld.Segtext.Sections[1:] {
		ctxt.Out.SeekSet(int64(sect.Vaddr - ld.Segtext.Vaddr + ld.Segtext.Fileoff))
		ld.Datblk(ctxt, int64(sect.Vaddr), int64(sect.Length))
	}

	if ld.Segrodata.Filelen > 0 {
		ctxt.Out.SeekSet(int64(ld.Segrodata.Fileoff))
		ld.Datblk(ctxt, int64(ld.Segrodata.Vaddr), int64(ld.Segrodata.Filelen))
	}
	if ld.Segrelrodata.Filelen > 0 {
		ctxt.Out.SeekSet(int64(ld.Segrelrodata.Fileoff))
		ld.Datblk(ctxt, int64(ld.Segrelrodata.Vaddr), int64(ld.Segrelrodata.Filelen))
	}

	ctxt.Out.SeekSet(int64(ld.Segdata.Fileoff))
	ld.Datblk(ctxt, int64(ld.Segdata.Vaddr), int64(ld.Segdata.Filelen))

	ctxt.Out.SeekSet(int64(ld.Segdwarf.Fileoff))
	ld.Dwarfblk(ctxt, int64(ld.Segdwarf.Vaddr), int64(ld.Segdwarf.Filelen))
}

func asmb2(ctxt *ld.Link) {
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
		ctxt.Out.Flush()
		ctxt.Out.Write(ld.Elfstrdat)

		if ctxt.LinkMode == ld.LinkExternal {
			ld.Elfemitreloc(ctxt)
		}
	}

	ctxt.Out.SeekSet(0)
	switch ctxt.HeadType {
	case objabi.Hlinux:
		ld.Asmbelf(ctxt, int64(symo))
	default:
		ld.Errorf(nil, "unsupported operating system")
	}
	ctxt.Out.Flush()

	if *ld.FlagC {
		fmt.Printf("textsize=%d\n", ld.Segtext.Filelen)
		fmt.Printf("datsize=%d\n", ld.Segdata.Filelen)
		fmt.Printf("bsssize=%d\n", ld.Segdata.Length-ld.Segdata.Filelen)
		fmt.Printf("symsize=%d\n", ld.Symsize)
		fmt.Printf("lcsize=%d\n", ld.Lcsize)
		fmt.Printf("total=%d\n", ld.Segtext.Filelen+ld.Segdata.Length+uint64(ld.Symsize)+uint64(ld.Lcsize))
	}
}
