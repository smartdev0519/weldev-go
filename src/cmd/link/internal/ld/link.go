// Derived from Inferno utils/6l/l.h and related files.
// https://bitbucket.org/inferno-os/inferno-os/src/default/utils/6l/l.h
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
//	Portions Copyright © 2009 The Go Authors. All rights reserved.
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

package ld

import (
	"bufio"
	"cmd/internal/objabi"
	"cmd/internal/sys"
	"debug/elf"
	"fmt"
)

// Symbol is an entry in the symbol table.
type Symbol struct {
	Name        string
	Extname     string
	Type        SymKind
	Version     int16
	Attr        Attribute
	Localentry  uint8
	Dynid       int32
	Plt         int32
	Got         int32
	Align       int32
	Elfsym      int32
	LocalElfsym int32
	Value       int64
	Size        int64
	// ElfType is set for symbols read from shared libraries by ldshlibsyms. It
	// is not set for symbols defined by the packages being linked or by symbols
	// read by ldelf (and so is left as elf.STT_NOTYPE).
	ElfType     elf.SymType
	Sub         *Symbol
	Outer       *Symbol
	Gotype      *Symbol
	Reachparent *Symbol
	File        string
	Dynimplib   string
	Dynimpvers  string
	Sect        *Section
	FuncInfo    *FuncInfo
	// P contains the raw symbol data.
	P []byte
	R []Reloc
}

func (s *Symbol) String() string {
	if s.Version == 0 {
		return s.Name
	}
	return fmt.Sprintf("%s<%d>", s.Name, s.Version)
}

func (s *Symbol) ElfsymForReloc() int32 {
	// If putelfsym created a local version of this symbol, use that in all
	// relocations.
	if s.LocalElfsym != 0 {
		return s.LocalElfsym
	} else {
		return s.Elfsym
	}
}

// Attribute is a set of common symbol attributes.
type Attribute int16

const (
	// AttrDuplicateOK marks a symbol that can be present in multiple object
	// files.
	AttrDuplicateOK Attribute = 1 << iota
	// AttrExternal marks function symbols loaded from host object files.
	AttrExternal
	// AttrNoSplit marks functions that cannot split the stack; the linker
	// cares because it checks that there are no call chains of nosplit
	// functions that require more than StackLimit bytes (see
	// lib.go:dostkcheck)
	AttrNoSplit
	// AttrReachable marks symbols that are transitively referenced from the
	// entry points. Unreachable symbols are not written to the output.
	AttrReachable
	// AttrCgoExportDynamic and AttrCgoExportStatic mark symbols referenced
	// by directives written by cgo (in response to //export directives in
	// the source).
	AttrCgoExportDynamic
	AttrCgoExportStatic
	// AttrSpecial marks symbols that do not have their address (i.e. Value)
	// computed by the usual mechanism of data.go:dodata() &
	// data.go:address().
	AttrSpecial
	// AttrStackCheck is used by dostkcheck to only check each NoSplit
	// function's stack usage once.
	AttrStackCheck
	// AttrNotInSymbolTable marks symbols that are not written to the symbol table.
	AttrNotInSymbolTable
	// AttrOnList marks symbols that are on some list (such as the list of
	// all text symbols, or one of the lists of data symbols) and is
	// consulted to avoid bugs where a symbol is put on a list twice.
	AttrOnList
	// AttrLocal marks symbols that are only visible within the module
	// (exectuable or shared library) being linked. Only relevant when
	// dynamically linking Go code.
	AttrLocal
	// AttrReflectMethod marks certain methods from the reflect package that
	// can be used to call arbitrary methods. If no symbol with this bit set
	// is marked as reachable, more dead code elimination can be done.
	AttrReflectMethod
	// AttrMakeTypelink Amarks types that should be added to the typelink
	// table. See typelinks.go:typelinks().
	AttrMakeTypelink
	// AttrShared marks symbols compiled with the -shared option.
	AttrShared
	// 14 attributes defined so far.
)

func (a Attribute) DuplicateOK() bool      { return a&AttrDuplicateOK != 0 }
func (a Attribute) External() bool         { return a&AttrExternal != 0 }
func (a Attribute) NoSplit() bool          { return a&AttrNoSplit != 0 }
func (a Attribute) Reachable() bool        { return a&AttrReachable != 0 }
func (a Attribute) CgoExportDynamic() bool { return a&AttrCgoExportDynamic != 0 }
func (a Attribute) CgoExportStatic() bool  { return a&AttrCgoExportStatic != 0 }
func (a Attribute) Special() bool          { return a&AttrSpecial != 0 }
func (a Attribute) StackCheck() bool       { return a&AttrStackCheck != 0 }
func (a Attribute) NotInSymbolTable() bool { return a&AttrNotInSymbolTable != 0 }
func (a Attribute) OnList() bool           { return a&AttrOnList != 0 }
func (a Attribute) Local() bool            { return a&AttrLocal != 0 }
func (a Attribute) ReflectMethod() bool    { return a&AttrReflectMethod != 0 }
func (a Attribute) MakeTypelink() bool     { return a&AttrMakeTypelink != 0 }
func (a Attribute) Shared() bool           { return a&AttrShared != 0 }

func (a Attribute) CgoExport() bool {
	return a.CgoExportDynamic() || a.CgoExportStatic()
}

func (a *Attribute) Set(flag Attribute, value bool) {
	if value {
		*a |= flag
	} else {
		*a &^= flag
	}
}

// Reloc is a relocation.
//
// The typical Reloc rewrites part of a symbol at offset Off to address Sym.
// A Reloc is stored in a slice on the Symbol it rewrites.
//
// Relocations are generated by the compiler as the type
// cmd/internal/obj.Reloc, which is encoded into the object file wire
// format and decoded by the linker into this type. A separate type is
// used to hold linker-specific state about the relocation.
//
// Some relocations are created by cmd/link.
type Reloc struct {
	Off     int32            // offset to rewrite
	Siz     uint8            // number of bytes to rewrite, 1, 2, or 4
	Done    uint8            // set to 1 when relocation is complete
	Variant RelocVariant     // variation on Type
	Type    objabi.RelocType // the relocation type
	Add     int64            // addend
	Xadd    int64            // addend passed to external linker
	Sym     *Symbol          // symbol the relocation addresses
	Xsym    *Symbol          // symbol passed to external linker
}

type Auto struct {
	Asym    *Symbol
	Gotype  *Symbol
	Aoffset int32
	Name    int16
}

type Shlib struct {
	Path            string
	Hash            []byte
	Deps            []string
	File            *elf.File
	gcdataAddresses map[*Symbol]uint64
}

// Link holds the context for writing object code from a compiler
// or for reading that input into the linker.
type Link struct {
	Syms *Symbols

	Arch      *sys.Arch
	Debugvlog int
	Bso       *bufio.Writer

	Loaded bool // set after all inputs have been loaded as symbols

	Tlsg       *Symbol
	Libdir     []string
	Library    []*Library
	Shlibs     []Shlib
	Tlsoffset  int
	Textp      []*Symbol
	Filesyms   []*Symbol
	Moduledata *Symbol

	tramps []*Symbol // trampolines
}

// The smallest possible offset from the hardware stack pointer to a local
// variable on the stack. Architectures that use a link register save its value
// on the stack in the function prologue and so always have a pointer between
// the hardware stack pointer and the local variable area.
func (ctxt *Link) FixedFrameSize() int64 {
	switch ctxt.Arch.Family {
	case sys.AMD64, sys.I386:
		return 0
	case sys.PPC64:
		// PIC code on ppc64le requires 32 bytes of stack, and it's easier to
		// just use that much stack always on ppc64x.
		return int64(4 * ctxt.Arch.PtrSize)
	default:
		return int64(ctxt.Arch.PtrSize)
	}
}

func (l *Link) Logf(format string, args ...interface{}) {
	fmt.Fprintf(l.Bso, format, args...)
	l.Bso.Flush()
}

type Library struct {
	Objref      string
	Srcref      string
	File        string
	Pkg         string
	Shlib       string
	hash        string
	imports     []*Library
	textp       []*Symbol // text symbols defined in this library
	dupTextSyms []*Symbol // dupok text symbols defined in this library
}

func (l Library) String() string {
	return l.Pkg
}

type FuncInfo struct {
	Args        int32
	Locals      int32
	Autom       []Auto
	Pcsp        Pcdata
	Pcfile      Pcdata
	Pcline      Pcdata
	Pcinline    Pcdata
	Pcdata      []Pcdata
	Funcdata    []*Symbol
	Funcdataoff []int64
	File        []*Symbol
	InlTree     []InlinedCall
}

// InlinedCall is a node in a local inlining tree (FuncInfo.InlTree).
type InlinedCall struct {
	Parent int32   // index of parent in InlTree
	File   *Symbol // file of the inlined call
	Line   int32   // line number of the inlined call
	Func   *Symbol // function that was inlined
}

type Pcdata struct {
	P []byte
}

type Pciter struct {
	d       Pcdata
	p       []byte
	pc      uint32
	nextpc  uint32
	pcscale uint32
	value   int32
	start   int
	done    int
}

// RelocVariant is a linker-internal variation on a relocation.
type RelocVariant uint8

const (
	RV_NONE RelocVariant = iota
	RV_POWER_LO
	RV_POWER_HI
	RV_POWER_HA
	RV_POWER_DS

	// RV_390_DBL is a s390x-specific relocation variant that indicates that
	// the value to be placed into the relocatable field should first be
	// divided by 2.
	RV_390_DBL

	RV_CHECK_OVERFLOW RelocVariant = 1 << 7
	RV_TYPE_MASK      RelocVariant = RV_CHECK_OVERFLOW - 1
)
