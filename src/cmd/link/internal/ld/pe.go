// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ld

import (
	"cmd/internal/objabi"
	"cmd/internal/sys"
	"debug/pe"
	"encoding/binary"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type IMAGE_FILE_HEADER struct {
	Machine              uint16
	NumberOfSections     uint16
	TimeDateStamp        uint32
	PointerToSymbolTable uint32
	NumberOfSymbols      uint32
	SizeOfOptionalHeader uint16
	Characteristics      uint16
}

type IMAGE_DATA_DIRECTORY struct {
	VirtualAddress uint32
	Size           uint32
}

type IMAGE_OPTIONAL_HEADER struct {
	Magic                       uint16
	MajorLinkerVersion          uint8
	MinorLinkerVersion          uint8
	SizeOfCode                  uint32
	SizeOfInitializedData       uint32
	SizeOfUninitializedData     uint32
	AddressOfEntryPoint         uint32
	BaseOfCode                  uint32
	BaseOfData                  uint32
	ImageBase                   uint32
	SectionAlignment            uint32
	FileAlignment               uint32
	MajorOperatingSystemVersion uint16
	MinorOperatingSystemVersion uint16
	MajorImageVersion           uint16
	MinorImageVersion           uint16
	MajorSubsystemVersion       uint16
	MinorSubsystemVersion       uint16
	Win32VersionValue           uint32
	SizeOfImage                 uint32
	SizeOfHeaders               uint32
	CheckSum                    uint32
	Subsystem                   uint16
	DllCharacteristics          uint16
	SizeOfStackReserve          uint32
	SizeOfStackCommit           uint32
	SizeOfHeapReserve           uint32
	SizeOfHeapCommit            uint32
	LoaderFlags                 uint32
	NumberOfRvaAndSizes         uint32
	DataDirectory               [16]IMAGE_DATA_DIRECTORY
}

type IMAGE_IMPORT_DESCRIPTOR struct {
	OriginalFirstThunk uint32
	TimeDateStamp      uint32
	ForwarderChain     uint32
	Name               uint32
	FirstThunk         uint32
}

type IMAGE_EXPORT_DIRECTORY struct {
	Characteristics       uint32
	TimeDateStamp         uint32
	MajorVersion          uint16
	MinorVersion          uint16
	Name                  uint32
	Base                  uint32
	NumberOfFunctions     uint32
	NumberOfNames         uint32
	AddressOfFunctions    uint32
	AddressOfNames        uint32
	AddressOfNameOrdinals uint32
}

const (
	PEBASE = 0x00400000
)

var (
	// SectionAlignment must be greater than or equal to FileAlignment.
	// The default is the page size for the architecture.
	PESECTALIGN int64 = 0x1000

	// FileAlignment should be a power of 2 between 512 and 64 K, inclusive.
	// The default is 512. If the SectionAlignment is less than
	// the architecture's page size, then FileAlignment must match SectionAlignment.
	PEFILEALIGN int64 = 2 << 8
)

const (
	IMAGE_FILE_MACHINE_I386              = 0x14c
	IMAGE_FILE_MACHINE_AMD64             = 0x8664
	IMAGE_FILE_RELOCS_STRIPPED           = 0x0001
	IMAGE_FILE_EXECUTABLE_IMAGE          = 0x0002
	IMAGE_FILE_LINE_NUMS_STRIPPED        = 0x0004
	IMAGE_FILE_LARGE_ADDRESS_AWARE       = 0x0020
	IMAGE_FILE_32BIT_MACHINE             = 0x0100
	IMAGE_FILE_DEBUG_STRIPPED            = 0x0200
	IMAGE_SCN_CNT_CODE                   = 0x00000020
	IMAGE_SCN_CNT_INITIALIZED_DATA       = 0x00000040
	IMAGE_SCN_CNT_UNINITIALIZED_DATA     = 0x00000080
	IMAGE_SCN_MEM_EXECUTE                = 0x20000000
	IMAGE_SCN_MEM_READ                   = 0x40000000
	IMAGE_SCN_MEM_WRITE                  = 0x80000000
	IMAGE_SCN_MEM_DISCARDABLE            = 0x2000000
	IMAGE_SCN_LNK_NRELOC_OVFL            = 0x1000000
	IMAGE_SCN_ALIGN_32BYTES              = 0x600000
	IMAGE_DIRECTORY_ENTRY_EXPORT         = 0
	IMAGE_DIRECTORY_ENTRY_IMPORT         = 1
	IMAGE_DIRECTORY_ENTRY_RESOURCE       = 2
	IMAGE_DIRECTORY_ENTRY_EXCEPTION      = 3
	IMAGE_DIRECTORY_ENTRY_SECURITY       = 4
	IMAGE_DIRECTORY_ENTRY_BASERELOC      = 5
	IMAGE_DIRECTORY_ENTRY_DEBUG          = 6
	IMAGE_DIRECTORY_ENTRY_COPYRIGHT      = 7
	IMAGE_DIRECTORY_ENTRY_ARCHITECTURE   = 7
	IMAGE_DIRECTORY_ENTRY_GLOBALPTR      = 8
	IMAGE_DIRECTORY_ENTRY_TLS            = 9
	IMAGE_DIRECTORY_ENTRY_LOAD_CONFIG    = 10
	IMAGE_DIRECTORY_ENTRY_BOUND_IMPORT   = 11
	IMAGE_DIRECTORY_ENTRY_IAT            = 12
	IMAGE_DIRECTORY_ENTRY_DELAY_IMPORT   = 13
	IMAGE_DIRECTORY_ENTRY_COM_DESCRIPTOR = 14
	IMAGE_SUBSYSTEM_WINDOWS_GUI          = 2
	IMAGE_SUBSYSTEM_WINDOWS_CUI          = 3
)

// X64
type PE64_IMAGE_OPTIONAL_HEADER struct {
	Magic                       uint16
	MajorLinkerVersion          uint8
	MinorLinkerVersion          uint8
	SizeOfCode                  uint32
	SizeOfInitializedData       uint32
	SizeOfUninitializedData     uint32
	AddressOfEntryPoint         uint32
	BaseOfCode                  uint32
	ImageBase                   uint64
	SectionAlignment            uint32
	FileAlignment               uint32
	MajorOperatingSystemVersion uint16
	MinorOperatingSystemVersion uint16
	MajorImageVersion           uint16
	MinorImageVersion           uint16
	MajorSubsystemVersion       uint16
	MinorSubsystemVersion       uint16
	Win32VersionValue           uint32
	SizeOfImage                 uint32
	SizeOfHeaders               uint32
	CheckSum                    uint32
	Subsystem                   uint16
	DllCharacteristics          uint16
	SizeOfStackReserve          uint64
	SizeOfStackCommit           uint64
	SizeOfHeapReserve           uint64
	SizeOfHeapCommit            uint64
	LoaderFlags                 uint32
	NumberOfRvaAndSizes         uint32
	DataDirectory               [16]IMAGE_DATA_DIRECTORY
}

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// PE (Portable Executable) file writing
// http://www.microsoft.com/whdc/system/platform/firmware/PECOFF.mspx

// DOS stub that prints out
// "This program cannot be run in DOS mode."
var dosstub = []uint8{
	0x4d,
	0x5a,
	0x90,
	0x00,
	0x03,
	0x00,
	0x04,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0xff,
	0xff,
	0x00,
	0x00,
	0x8b,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x40,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x80,
	0x00,
	0x00,
	0x00,
	0x0e,
	0x1f,
	0xba,
	0x0e,
	0x00,
	0xb4,
	0x09,
	0xcd,
	0x21,
	0xb8,
	0x01,
	0x4c,
	0xcd,
	0x21,
	0x54,
	0x68,
	0x69,
	0x73,
	0x20,
	0x70,
	0x72,
	0x6f,
	0x67,
	0x72,
	0x61,
	0x6d,
	0x20,
	0x63,
	0x61,
	0x6e,
	0x6e,
	0x6f,
	0x74,
	0x20,
	0x62,
	0x65,
	0x20,
	0x72,
	0x75,
	0x6e,
	0x20,
	0x69,
	0x6e,
	0x20,
	0x44,
	0x4f,
	0x53,
	0x20,
	0x6d,
	0x6f,
	0x64,
	0x65,
	0x2e,
	0x0d,
	0x0d,
	0x0a,
	0x24,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
	0x00,
}

var rsrcsym *Symbol

var PESECTHEADR int32

var PEFILEHEADR int32

var pe64 int

var fh IMAGE_FILE_HEADER

var oh IMAGE_OPTIONAL_HEADER

var oh64 PE64_IMAGE_OPTIONAL_HEADER

var dd []IMAGE_DATA_DIRECTORY

type Imp struct {
	s       *Symbol
	off     uint64
	next    *Imp
	argsize int
}

type Dll struct {
	name     string
	nameoff  uint64
	thunkoff uint64
	ms       *Imp
	next     *Dll
}

var dr *Dll

var dexport [1024]*Symbol

var nexport int

// peStringTable is a COFF string table.
type peStringTable struct {
	strings    []string
	stringsLen int
}

// size resturns size of string table t.
func (t *peStringTable) size() int {
	// string table starts with 4-byte length at the beginning
	return t.stringsLen + 4
}

// add adds string str to string table t.
func (t *peStringTable) add(str string) int {
	off := t.size()
	t.strings = append(t.strings, str)
	t.stringsLen += len(str) + 1 // each string will have 0 appended to it
	return off
}

// write writes string table t into the output file.
func (t *peStringTable) write() {
	Lputl(uint32(t.size()))
	for _, s := range t.strings {
		Cwritestring(s)
		Cput(0)
	}
}

// peSection represents section from COFF section table.
type peSection struct {
	name      string
	shortName string
	index     int // one-based index into the Section Table
	// TODO: change all these names to start with small letters
	VirtualSize          uint32
	VirtualAddress       uint32
	SizeOfRawData        uint32
	PointerToRawData     uint32
	PointerToRelocations uint32
	NumberOfRelocations  uint16
	Characteristics      uint32
}

// checkOffset verifies COFF section sect offset in the file.
func (sect *peSection) checkOffset(off int64) {
	if off != int64(sect.PointerToRawData) {
		Errorf(nil, "%s.PointerToRawData = %#x, want %#x", sect.name, uint64(int64(sect.PointerToRawData)), uint64(off))
		errorexit()
	}
}

// checkSegment verifies COFF section sect matches address
// and file offset provided in segment seg.
func (sect *peSection) checkSegment(seg *Segment) {
	if seg.Vaddr-PEBASE != uint64(sect.VirtualAddress) {
		Errorf(nil, "%s.VirtualAddress = %#x, want %#x", sect.name, uint64(int64(sect.VirtualAddress)), uint64(int64(seg.Vaddr-PEBASE)))
		errorexit()
	}
	if seg.Fileoff != uint64(sect.PointerToRawData) {
		Errorf(nil, "%s.PointerToRawData = %#x, want %#x", sect.name, uint64(int64(sect.PointerToRawData)), uint64(int64(seg.Fileoff)))
		errorexit()
	}
}

// pad adds zeros to the section sect. It writes as many bytes
// as necessary to make section sect.SizeOfRawData bytes long.
// It assumes that n bytes are already written to the file.
func (sect *peSection) pad(n uint32) {
	strnput("", int(sect.SizeOfRawData-n))
}

// write writes COFF section sect into the output file.
func (sect *peSection) write() error {
	h := pe.SectionHeader32{
		VirtualSize:          sect.VirtualSize,
		SizeOfRawData:        sect.SizeOfRawData,
		PointerToRawData:     sect.PointerToRawData,
		PointerToRelocations: sect.PointerToRelocations,
		NumberOfRelocations:  sect.NumberOfRelocations,
		Characteristics:      sect.Characteristics,
	}
	if Linkmode != LinkExternal {
		h.VirtualAddress = sect.VirtualAddress
	}
	copy(h.Name[:], sect.shortName)
	return binary.Write(&coutbuf, binary.LittleEndian, h)
}

// peFile is used to build COFF file.
type peFile struct {
	sections       []*peSection
	stringTable    peStringTable
	textSect       *peSection
	dataSect       *peSection
	bssSect        *peSection
	nextSectOffset uint32
	nextFileOffset uint32
}

// addSection adds section to the COFF file f.
func (f *peFile) addSection(name string, sectsize int, filesize int) *peSection {
	sect := &peSection{
		name:             name,
		shortName:        name,
		index:            len(f.sections) + 1,
		VirtualSize:      uint32(sectsize),
		VirtualAddress:   f.nextSectOffset,
		PointerToRawData: f.nextFileOffset,
	}
	f.nextSectOffset = uint32(Rnd(int64(f.nextSectOffset)+int64(sectsize), PESECTALIGN))
	if filesize > 0 {
		sect.SizeOfRawData = uint32(Rnd(int64(filesize), PEFILEALIGN))
		f.nextFileOffset += sect.SizeOfRawData
	}
	f.sections = append(f.sections, sect)
	return sect
}

// addDWARFSection adds DWARF section to the COFF file f.
// This function is similar to addSection, but DWARF section names are
// longer than 8 characters, so they need to be stored in the string table.
func (f *peFile) addDWARFSection(name string, size int) *peSection {
	if size == 0 {
		Exitf("DWARF section %q is empty", name)
	}
	// DWARF section names are longer than 8 characters.
	// PE format requires such names to be stored in string table,
	// and section names replaced with slash (/) followed by
	// correspondent string table index.
	// see http://www.microsoft.com/whdc/system/platform/firmware/PECOFFdwn.mspx
	// for details
	off := f.stringTable.add(name)
	h := f.addSection(name, size, size)
	h.shortName = fmt.Sprintf("/%d", off)
	h.Characteristics = IMAGE_SCN_MEM_READ | IMAGE_SCN_MEM_DISCARDABLE
	return h
}

// addInitArray adds .ctors COFF section to the file f.
func (f *peFile) addInitArray(ctxt *Link) *peSection {
	// The size below was determined by the specification for array relocations,
	// and by observing what GCC writes here. If the initarray section grows to
	// contain more than one constructor entry, the size will need to be 8 * constructor_count.
	// However, the entire Go runtime is initialized from just one function, so it is unlikely
	// that this will need to grow in the future.
	var size int
	switch objabi.GOARCH {
	default:
		Exitf("peFile.addInitArray: unsupported GOARCH=%q\n", objabi.GOARCH)
	case "386":
		size = 4
	case "amd64":
		size = 8
	}
	sect := f.addSection(".ctors", size, size)
	sect.Characteristics = IMAGE_SCN_CNT_INITIALIZED_DATA | IMAGE_SCN_MEM_READ
	sect.SizeOfRawData = uint32(size)
	Cseek(int64(sect.PointerToRawData))
	sect.checkOffset(coutbuf.Offset())

	init_entry := ctxt.Syms.Lookup(*flagEntrySymbol, 0)
	addr := uint64(init_entry.Value) - init_entry.Sect.Vaddr
	switch objabi.GOARCH {
	case "386":
		Lputl(uint32(addr))
	case "amd64":
		Vputl(addr)
	}
	return sect
}

var pefile peFile

func Peinit(ctxt *Link) {
	var l int

	switch SysArch.Family {
	// 64-bit architectures
	case sys.AMD64:
		pe64 = 1

		l = binary.Size(&oh64)
		dd = oh64.DataDirectory[:]

	// 32-bit architectures
	default:
		l = binary.Size(&oh)

		dd = oh.DataDirectory[:]
	}

	if Linkmode == LinkExternal {
		PESECTALIGN = 0
		PEFILEALIGN = 0
	}

	var sh [16]pe.SectionHeader32
	PEFILEHEADR = int32(Rnd(int64(len(dosstub)+binary.Size(&fh)+l+binary.Size(&sh)), PEFILEALIGN))
	if Linkmode != LinkExternal {
		PESECTHEADR = int32(Rnd(int64(PEFILEHEADR), PESECTALIGN))
	} else {
		PESECTHEADR = 0
	}
	pefile.nextSectOffset = uint32(PESECTHEADR)
	pefile.nextFileOffset = uint32(PEFILEHEADR)

	if Linkmode == LinkInternal {
		// some mingw libs depend on this symbol, for example, FindPESectionByName
		ctxt.xdefine("__image_base__", SDATA, PEBASE)
		ctxt.xdefine("_image_base__", SDATA, PEBASE)
	}

	HEADR = PEFILEHEADR
	if *FlagTextAddr == -1 {
		*FlagTextAddr = PEBASE + int64(PESECTHEADR)
	}
	if *FlagDataAddr == -1 {
		*FlagDataAddr = 0
	}
	if *FlagRound == -1 {
		*FlagRound = int(PESECTALIGN)
	}
	if *FlagDataAddr != 0 && *FlagRound != 0 {
		fmt.Printf("warning: -D0x%x is ignored because of -R0x%x\n", uint64(*FlagDataAddr), uint32(*FlagRound))
	}
}

func pewrite() {
	Cseek(0)
	if Linkmode != LinkExternal {
		Cwrite(dosstub)
		strnput("PE", 4)
	}

	binary.Write(&coutbuf, binary.LittleEndian, &fh)

	if pe64 != 0 {
		binary.Write(&coutbuf, binary.LittleEndian, &oh64)
	} else {
		binary.Write(&coutbuf, binary.LittleEndian, &oh)
	}
	for _, sect := range pefile.sections {
		sect.write()
	}
}

func strput(s string) {
	coutbuf.WriteString(s)
	Cput(0)
	// string must be padded to even size
	if (len(s)+1)%2 != 0 {
		Cput(0)
	}
}

func initdynimport(ctxt *Link) *Dll {
	var d *Dll

	dr = nil
	var m *Imp
	for _, s := range ctxt.Syms.Allsym {
		if !s.Attr.Reachable() || s.Type != SDYNIMPORT {
			continue
		}
		for d = dr; d != nil; d = d.next {
			if d.name == s.Dynimplib {
				m = new(Imp)
				break
			}
		}

		if d == nil {
			d = new(Dll)
			d.name = s.Dynimplib
			d.next = dr
			dr = d
			m = new(Imp)
		}

		// Because external link requires properly stdcall decorated name,
		// all external symbols in runtime use %n to denote that the number
		// of uinptrs this function consumes. Store the argsize and discard
		// the %n suffix if any.
		m.argsize = -1
		if i := strings.IndexByte(s.Extname, '%'); i >= 0 {
			var err error
			m.argsize, err = strconv.Atoi(s.Extname[i+1:])
			if err != nil {
				Errorf(s, "failed to parse stdcall decoration: %v", err)
			}
			m.argsize *= SysArch.PtrSize
			s.Extname = s.Extname[:i]
		}

		m.s = s
		m.next = d.ms
		d.ms = m
	}

	if Linkmode == LinkExternal {
		// Add real symbol name
		for d := dr; d != nil; d = d.next {
			for m = d.ms; m != nil; m = m.next {
				m.s.Type = SDATA
				Symgrow(m.s, int64(SysArch.PtrSize))
				dynName := m.s.Extname
				// only windows/386 requires stdcall decoration
				if SysArch.Family == sys.I386 && m.argsize >= 0 {
					dynName += fmt.Sprintf("@%d", m.argsize)
				}
				dynSym := ctxt.Syms.Lookup(dynName, 0)
				dynSym.Attr |= AttrReachable
				dynSym.Type = SHOSTOBJ
				r := Addrel(m.s)
				r.Sym = dynSym
				r.Off = 0
				r.Siz = uint8(SysArch.PtrSize)
				r.Type = objabi.R_ADDR
			}
		}
	} else {
		dynamic := ctxt.Syms.Lookup(".windynamic", 0)
		dynamic.Attr |= AttrReachable
		dynamic.Type = SWINDOWS
		for d := dr; d != nil; d = d.next {
			for m = d.ms; m != nil; m = m.next {
				m.s.Type = SWINDOWS | SSUB
				m.s.Sub = dynamic.Sub
				dynamic.Sub = m.s
				m.s.Value = dynamic.Size
				dynamic.Size += int64(SysArch.PtrSize)
			}

			dynamic.Size += int64(SysArch.PtrSize)
		}
	}

	return dr
}

// peimporteddlls returns the gcc command line argument to link all imported
// DLLs.
func peimporteddlls() []string {
	var dlls []string

	for d := dr; d != nil; d = d.next {
		dlls = append(dlls, "-l"+strings.TrimSuffix(d.name, ".dll"))
	}

	return dlls
}

func addimports(ctxt *Link, datsect *peSection) {
	startoff := coutbuf.Offset()
	dynamic := ctxt.Syms.Lookup(".windynamic", 0)

	// skip import descriptor table (will write it later)
	n := uint64(0)

	for d := dr; d != nil; d = d.next {
		n++
	}
	Cseek(startoff + int64(binary.Size(&IMAGE_IMPORT_DESCRIPTOR{}))*int64(n+1))

	// write dll names
	for d := dr; d != nil; d = d.next {
		d.nameoff = uint64(coutbuf.Offset()) - uint64(startoff)
		strput(d.name)
	}

	// write function names
	var m *Imp
	for d := dr; d != nil; d = d.next {
		for m = d.ms; m != nil; m = m.next {
			m.off = uint64(pefile.nextSectOffset) + uint64(coutbuf.Offset()) - uint64(startoff)
			Wputl(0) // hint
			strput(m.s.Extname)
		}
	}

	// write OriginalFirstThunks
	oftbase := uint64(coutbuf.Offset()) - uint64(startoff)

	n = uint64(coutbuf.Offset())
	for d := dr; d != nil; d = d.next {
		d.thunkoff = uint64(coutbuf.Offset()) - n
		for m = d.ms; m != nil; m = m.next {
			if pe64 != 0 {
				Vputl(m.off)
			} else {
				Lputl(uint32(m.off))
			}
		}

		if pe64 != 0 {
			Vputl(0)
		} else {
			Lputl(0)
		}
	}

	// add pe section and pad it at the end
	n = uint64(coutbuf.Offset()) - uint64(startoff)

	isect := pefile.addSection(".idata", int(n), int(n))
	isect.Characteristics = IMAGE_SCN_CNT_INITIALIZED_DATA | IMAGE_SCN_MEM_READ | IMAGE_SCN_MEM_WRITE
	isect.checkOffset(startoff)
	isect.pad(uint32(n))
	endoff := coutbuf.Offset()

	// write FirstThunks (allocated in .data section)
	ftbase := uint64(dynamic.Value) - uint64(datsect.VirtualAddress) - PEBASE

	Cseek(int64(uint64(datsect.PointerToRawData) + ftbase))
	for d := dr; d != nil; d = d.next {
		for m = d.ms; m != nil; m = m.next {
			if pe64 != 0 {
				Vputl(m.off)
			} else {
				Lputl(uint32(m.off))
			}
		}

		if pe64 != 0 {
			Vputl(0)
		} else {
			Lputl(0)
		}
	}

	// finally write import descriptor table
	Cseek(startoff)

	for d := dr; d != nil; d = d.next {
		Lputl(uint32(uint64(isect.VirtualAddress) + oftbase + d.thunkoff))
		Lputl(0)
		Lputl(0)
		Lputl(uint32(uint64(isect.VirtualAddress) + d.nameoff))
		Lputl(uint32(uint64(datsect.VirtualAddress) + ftbase + d.thunkoff))
	}

	Lputl(0) //end
	Lputl(0)
	Lputl(0)
	Lputl(0)
	Lputl(0)

	// update data directory
	dd[IMAGE_DIRECTORY_ENTRY_IMPORT].VirtualAddress = isect.VirtualAddress
	dd[IMAGE_DIRECTORY_ENTRY_IMPORT].Size = isect.VirtualSize
	dd[IMAGE_DIRECTORY_ENTRY_IAT].VirtualAddress = uint32(dynamic.Value - PEBASE)
	dd[IMAGE_DIRECTORY_ENTRY_IAT].Size = uint32(dynamic.Size)

	Cseek(endoff)
}

type byExtname []*Symbol

func (s byExtname) Len() int           { return len(s) }
func (s byExtname) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s byExtname) Less(i, j int) bool { return s[i].Extname < s[j].Extname }

func initdynexport(ctxt *Link) {
	nexport = 0
	for _, s := range ctxt.Syms.Allsym {
		if !s.Attr.Reachable() || !s.Attr.CgoExportDynamic() {
			continue
		}
		if nexport+1 > len(dexport) {
			Errorf(s, "pe dynexport table is full")
			errorexit()
		}

		dexport[nexport] = s
		nexport++
	}

	sort.Sort(byExtname(dexport[:nexport]))
}

func addexports(ctxt *Link) {
	var e IMAGE_EXPORT_DIRECTORY

	size := binary.Size(&e) + 10*nexport + len(*flagOutfile) + 1
	for i := 0; i < nexport; i++ {
		size += len(dexport[i].Extname) + 1
	}

	if nexport == 0 {
		return
	}

	sect := pefile.addSection(".edata", size, size)
	sect.Characteristics = IMAGE_SCN_CNT_INITIALIZED_DATA | IMAGE_SCN_MEM_READ
	sect.checkOffset(coutbuf.Offset())
	va := int(sect.VirtualAddress)
	dd[IMAGE_DIRECTORY_ENTRY_EXPORT].VirtualAddress = uint32(va)
	dd[IMAGE_DIRECTORY_ENTRY_EXPORT].Size = sect.VirtualSize

	vaName := va + binary.Size(&e) + nexport*4
	vaAddr := va + binary.Size(&e)
	vaNa := va + binary.Size(&e) + nexport*8

	e.Characteristics = 0
	e.MajorVersion = 0
	e.MinorVersion = 0
	e.NumberOfFunctions = uint32(nexport)
	e.NumberOfNames = uint32(nexport)
	e.Name = uint32(va+binary.Size(&e)) + uint32(nexport)*10 // Program names.
	e.Base = 1
	e.AddressOfFunctions = uint32(vaAddr)
	e.AddressOfNames = uint32(vaName)
	e.AddressOfNameOrdinals = uint32(vaNa)

	// put IMAGE_EXPORT_DIRECTORY
	binary.Write(&coutbuf, binary.LittleEndian, &e)

	// put EXPORT Address Table
	for i := 0; i < nexport; i++ {
		Lputl(uint32(dexport[i].Value - PEBASE))
	}

	// put EXPORT Name Pointer Table
	v := int(e.Name + uint32(len(*flagOutfile)) + 1)

	for i := 0; i < nexport; i++ {
		Lputl(uint32(v))
		v += len(dexport[i].Extname) + 1
	}

	// put EXPORT Ordinal Table
	for i := 0; i < nexport; i++ {
		Wputl(uint16(i))
	}

	// put Names
	strnput(*flagOutfile, len(*flagOutfile)+1)

	for i := 0; i < nexport; i++ {
		strnput(dexport[i].Extname, len(dexport[i].Extname)+1)
	}
	sect.pad(uint32(size))
}

// perelocsect relocates symbols from first in section sect, and returns
// the total number of relocations emitted.
func perelocsect(ctxt *Link, sect *Section, syms []*Symbol, base uint64) int {
	// If main section has no bits, nothing to relocate.
	if sect.Vaddr >= sect.Seg.Vaddr+sect.Seg.Filelen {
		return 0
	}

	relocs := 0

	sect.Reloff = uint64(coutbuf.Offset())
	for i, s := range syms {
		if !s.Attr.Reachable() {
			continue
		}
		if uint64(s.Value) >= sect.Vaddr {
			syms = syms[i:]
			break
		}
	}

	eaddr := int32(sect.Vaddr + sect.Length)
	for _, sym := range syms {
		if !sym.Attr.Reachable() {
			continue
		}
		if sym.Value >= int64(eaddr) {
			break
		}
		for ri := 0; ri < len(sym.R); ri++ {
			r := &sym.R[ri]
			if r.Done != 0 {
				continue
			}
			if r.Xsym == nil {
				Errorf(sym, "missing xsym in relocation")
				continue
			}

			if r.Xsym.Dynid < 0 {
				Errorf(sym, "reloc %d (%s) to non-coff symbol %s (outer=%s) %d (%s)", r.Type, RelocName(r.Type), r.Sym.Name, r.Xsym.Name, r.Sym.Type, r.Sym.Type)
			}
			if !Thearch.PEreloc1(sym, r, int64(uint64(sym.Value+int64(r.Off))-base)) {
				Errorf(sym, "unsupported obj reloc %d (%s)/%d to %s", r.Type, RelocName(r.Type), r.Siz, r.Sym.Name)
			}

			relocs++
		}
	}

	sect.Rellen = uint64(coutbuf.Offset()) - sect.Reloff

	return relocs
}

// peemitsectreloc emits the relocation entries for sect.
// The actual relocations are emitted by relocfn.
// This updates the corresponding PE section table entry
// with the relocation offset and count.
func peemitsectreloc(sect *peSection, relocfn func() int) {
	sect.PointerToRelocations = uint32(coutbuf.Offset())
	// first entry: extended relocs
	Lputl(0) // placeholder for number of relocation + 1
	Lputl(0)
	Wputl(0)

	n := relocfn() + 1

	cpos := coutbuf.Offset()
	Cseek(int64(sect.PointerToRelocations))
	Lputl(uint32(n))
	Cseek(cpos)
	if n > 0x10000 {
		n = 0x10000
		sect.Characteristics |= IMAGE_SCN_LNK_NRELOC_OVFL
	} else {
		sect.PointerToRelocations += 10 // skip the extend reloc entry
	}
	sect.NumberOfRelocations = uint16(n - 1)
}

// peemitreloc emits relocation entries for go.o in external linking.
func peemitreloc(ctxt *Link, text, data, ctors *peSection) {
	for coutbuf.Offset()&7 != 0 {
		Cput(0)
	}

	peemitsectreloc(text, func() int {
		n := perelocsect(ctxt, Segtext.Sections[0], ctxt.Textp, Segtext.Vaddr)
		for _, sect := range Segtext.Sections[1:] {
			n += perelocsect(ctxt, sect, datap, Segtext.Vaddr)
		}
		return n
	})

	peemitsectreloc(data, func() int {
		var n int
		for _, sect := range Segdata.Sections {
			n += perelocsect(ctxt, sect, datap, Segdata.Vaddr)
		}
		return n
	})

dwarfLoop:
	for _, sect := range Segdwarf.Sections {
		for _, pesect := range pefile.sections {
			if sect.Name == pesect.name {
				peemitsectreloc(pesect, func() int {
					return perelocsect(ctxt, sect, dwarfp, sect.Vaddr)
				})
				continue dwarfLoop
			}
		}
		Errorf(nil, "peemitsectreloc: could not find %q section", sect.Name)
	}

	peemitsectreloc(ctors, func() int {
		dottext := ctxt.Syms.Lookup(".text", 0)
		Lputl(0)
		Lputl(uint32(dottext.Dynid))
		switch objabi.GOARCH {
		default:
			Errorf(dottext, "unknown architecture for PE: %q\n", objabi.GOARCH)
		case "386":
			Wputl(IMAGE_REL_I386_DIR32)
		case "amd64":
			Wputl(IMAGE_REL_AMD64_ADDR64)
		}
		return 1
	})
}

func (ctxt *Link) dope() {
	/* relocation table */
	rel := ctxt.Syms.Lookup(".rel", 0)

	rel.Attr |= AttrReachable
	rel.Type = SELFROSECT

	initdynimport(ctxt)
	initdynexport(ctxt)
}

// writePESymTableRecords writes all COFF symbol table records.
// It returns number of records written.
func writePESymTableRecords(ctxt *Link) int {
	var symcnt int

	writeOneSymbol := func(s *Symbol, addr int64, sectidx int, typ uint16, class uint8) {
		// write COFF symbol table record
		if len(s.Name) > 8 {
			Lputl(0)
			Lputl(uint32(pefile.stringTable.add(s.Name)))
		} else {
			strnput(s.Name, 8)
		}
		Lputl(uint32(addr))
		Wputl(uint16(sectidx))
		Wputl(typ)
		Cput(class)
		Cput(0) // no aux entries

		s.Dynid = int32(symcnt)

		symcnt++
	}

	put := func(ctxt *Link, s *Symbol, name string, type_ SymbolType, addr int64, gotype *Symbol) {
		if s == nil {
			return
		}
		if s.Sect == nil && type_ != UndefinedSym {
			return
		}
		switch type_ {
		default:
			return
		case DataSym, BSSSym, TextSym, UndefinedSym:
		}

		// Only windows/386 requires underscore prefix on external symbols.
		if SysArch.Family == sys.I386 &&
			Linkmode == LinkExternal &&
			(s.Type == SHOSTOBJ || s.Attr.CgoExport()) {
			s.Name = "_" + s.Name
		}

		typ := uint16(IMAGE_SYM_TYPE_NULL)
		var sect int
		var value int64
		if s.Sect != nil && s.Sect.Seg == &Segdata {
			// Note: although address of runtime.edata (type SDATA) is at the start of .bss section
			// it still belongs to the .data section, not the .bss section.
			if uint64(s.Value) >= Segdata.Vaddr+Segdata.Filelen && s.Type != SDATA && Linkmode == LinkExternal {
				value = int64(uint64(s.Value) - Segdata.Vaddr - Segdata.Filelen)
				sect = pefile.bssSect.index
			} else {
				value = int64(uint64(s.Value) - Segdata.Vaddr)
				sect = pefile.dataSect.index
			}
		} else if s.Sect != nil && s.Sect.Seg == &Segtext {
			value = int64(uint64(s.Value) - Segtext.Vaddr)
			sect = pefile.textSect.index
		} else if type_ == UndefinedSym {
			typ = IMAGE_SYM_DTYPE_FUNCTION
		} else {
			Errorf(s, "addpesym %#x", addr)
		}
		if typ != IMAGE_SYM_TYPE_NULL {
		} else if Linkmode != LinkExternal {
			// TODO: fix IMAGE_SYM_DTYPE_ARRAY value and use following expression, instead of 0x0308
			typ = IMAGE_SYM_DTYPE_ARRAY<<8 + IMAGE_SYM_TYPE_STRUCT
			typ = 0x0308 // "array of structs"
		}
		class := IMAGE_SYM_CLASS_EXTERNAL
		if s.Version != 0 || (s.Type&SHIDDEN != 0) || s.Attr.Local() {
			class = IMAGE_SYM_CLASS_STATIC
		}
		writeOneSymbol(s, value, sect, typ, uint8(class))
	}

	if Linkmode == LinkExternal {
		// Include section symbols as external, because
		// .ctors and .debug_* section relocations refer to it.
		for _, pesect := range pefile.sections {
			sym := ctxt.Syms.Lookup(pesect.name, 0)
			writeOneSymbol(sym, 0, pesect.index, IMAGE_SYM_TYPE_NULL, IMAGE_SYM_CLASS_STATIC)
		}
	}

	genasmsym(ctxt, put)

	return symcnt
}

func addpesymtable(ctxt *Link) {
	symtabStartPos := coutbuf.Offset()

	// write COFF symbol table
	var symcnt int
	if !*FlagS || Linkmode == LinkExternal {
		symcnt = writePESymTableRecords(ctxt)
	}

	// update COFF file header and section table
	size := pefile.stringTable.size() + 18*symcnt
	var h *peSection
	if Linkmode != LinkExternal {
		// We do not really need .symtab for go.o, and if we have one, ld
		// will also include it in the exe, and that will confuse windows.
		h = pefile.addSection(".symtab", size, size)
		h.Characteristics = IMAGE_SCN_MEM_READ | IMAGE_SCN_MEM_DISCARDABLE
		h.checkOffset(symtabStartPos)
	}
	fh.PointerToSymbolTable = uint32(symtabStartPos)
	fh.NumberOfSymbols = uint32(symcnt)

	// write COFF string table
	pefile.stringTable.write()
	if Linkmode != LinkExternal {
		h.pad(uint32(size))
	}
}

func setpersrc(ctxt *Link, sym *Symbol) {
	if rsrcsym != nil {
		Errorf(sym, "too many .rsrc sections")
	}

	rsrcsym = sym
}

func addpersrc(ctxt *Link) {
	if rsrcsym == nil {
		return
	}

	h := pefile.addSection(".rsrc", int(rsrcsym.Size), int(rsrcsym.Size))
	h.Characteristics = IMAGE_SCN_MEM_READ | IMAGE_SCN_MEM_WRITE | IMAGE_SCN_CNT_INITIALIZED_DATA
	h.checkOffset(coutbuf.Offset())

	// relocation
	var p []byte
	var r *Reloc
	var val uint32
	for ri := 0; ri < len(rsrcsym.R); ri++ {
		r = &rsrcsym.R[ri]
		p = rsrcsym.P[r.Off:]
		val = uint32(int64(h.VirtualAddress) + r.Add)

		// 32-bit little-endian
		p[0] = byte(val)

		p[1] = byte(val >> 8)
		p[2] = byte(val >> 16)
		p[3] = byte(val >> 24)
	}

	Cwrite(rsrcsym.P)
	h.pad(uint32(rsrcsym.Size))

	// update data directory
	dd[IMAGE_DIRECTORY_ENTRY_RESOURCE].VirtualAddress = h.VirtualAddress

	dd[IMAGE_DIRECTORY_ENTRY_RESOURCE].Size = h.VirtualSize
}

func Asmbpe(ctxt *Link) {
	switch SysArch.Family {
	default:
		Exitf("unknown PE architecture: %v", SysArch.Family)
	case sys.AMD64:
		fh.Machine = IMAGE_FILE_MACHINE_AMD64
	case sys.I386:
		fh.Machine = IMAGE_FILE_MACHINE_I386
	}

	t := pefile.addSection(".text", int(Segtext.Length), int(Segtext.Length))
	t.Characteristics = IMAGE_SCN_CNT_CODE | IMAGE_SCN_CNT_INITIALIZED_DATA | IMAGE_SCN_MEM_EXECUTE | IMAGE_SCN_MEM_READ
	if Linkmode == LinkExternal {
		// some data symbols (e.g. masks) end up in the .text section, and they normally
		// expect larger alignment requirement than the default text section alignment.
		t.Characteristics |= IMAGE_SCN_ALIGN_32BYTES
	}
	t.checkSegment(&Segtext)
	pefile.textSect = t

	var d *peSection
	var c *peSection
	if Linkmode != LinkExternal {
		d = pefile.addSection(".data", int(Segdata.Length), int(Segdata.Filelen))
		d.Characteristics = IMAGE_SCN_CNT_INITIALIZED_DATA | IMAGE_SCN_MEM_READ | IMAGE_SCN_MEM_WRITE
		d.checkSegment(&Segdata)
		pefile.dataSect = d
	} else {
		d = pefile.addSection(".data", int(Segdata.Filelen), int(Segdata.Filelen))
		d.Characteristics = IMAGE_SCN_CNT_INITIALIZED_DATA | IMAGE_SCN_MEM_READ | IMAGE_SCN_MEM_WRITE | IMAGE_SCN_ALIGN_32BYTES
		d.checkSegment(&Segdata)
		pefile.dataSect = d

		b := pefile.addSection(".bss", int(Segdata.Length-Segdata.Filelen), 0)
		b.Characteristics = IMAGE_SCN_CNT_UNINITIALIZED_DATA | IMAGE_SCN_MEM_READ | IMAGE_SCN_MEM_WRITE | IMAGE_SCN_ALIGN_32BYTES
		b.PointerToRawData = 0
		pefile.bssSect = b
	}

	if !*FlagS {
		dwarfaddpeheaders(ctxt)
	}

	if Linkmode == LinkExternal {
		c = pefile.addInitArray(ctxt)
	}

	Cseek(int64(pefile.nextFileOffset))
	if Linkmode != LinkExternal {
		addimports(ctxt, d)
		addexports(ctxt)
	}
	addpesymtable(ctxt)
	addpersrc(ctxt)
	if Linkmode == LinkExternal {
		peemitreloc(ctxt, t, d, c)
	}

	fh.NumberOfSections = uint16(len(pefile.sections))

	// Being able to produce identical output for identical input is
	// much more beneficial than having build timestamp in the header.
	fh.TimeDateStamp = 0

	if Linkmode == LinkExternal {
		fh.Characteristics = IMAGE_FILE_LINE_NUMS_STRIPPED
	} else {
		fh.Characteristics = IMAGE_FILE_RELOCS_STRIPPED | IMAGE_FILE_EXECUTABLE_IMAGE | IMAGE_FILE_DEBUG_STRIPPED
	}
	if pe64 != 0 {
		fh.SizeOfOptionalHeader = uint16(binary.Size(&oh64))
		fh.Characteristics |= IMAGE_FILE_LARGE_ADDRESS_AWARE
		oh64.Magic = 0x20b // PE32+
	} else {
		fh.SizeOfOptionalHeader = uint16(binary.Size(&oh))
		fh.Characteristics |= IMAGE_FILE_32BIT_MACHINE
		oh.Magic = 0x10b // PE32
		oh.BaseOfData = d.VirtualAddress
	}

	// Fill out both oh64 and oh. We only use one. Oh well.
	oh64.MajorLinkerVersion = 3

	oh.MajorLinkerVersion = 3
	oh64.MinorLinkerVersion = 0
	oh.MinorLinkerVersion = 0
	oh64.SizeOfCode = t.SizeOfRawData
	oh.SizeOfCode = t.SizeOfRawData
	oh64.SizeOfInitializedData = d.SizeOfRawData
	oh.SizeOfInitializedData = d.SizeOfRawData
	oh64.SizeOfUninitializedData = 0
	oh.SizeOfUninitializedData = 0
	if Linkmode != LinkExternal {
		oh64.AddressOfEntryPoint = uint32(Entryvalue(ctxt) - PEBASE)
		oh.AddressOfEntryPoint = uint32(Entryvalue(ctxt) - PEBASE)
	}
	oh64.BaseOfCode = t.VirtualAddress
	oh.BaseOfCode = t.VirtualAddress
	oh64.ImageBase = PEBASE
	oh.ImageBase = PEBASE
	oh64.SectionAlignment = uint32(PESECTALIGN)
	oh.SectionAlignment = uint32(PESECTALIGN)
	oh64.FileAlignment = uint32(PEFILEALIGN)
	oh.FileAlignment = uint32(PEFILEALIGN)
	oh64.MajorOperatingSystemVersion = 4
	oh.MajorOperatingSystemVersion = 4
	oh64.MinorOperatingSystemVersion = 0
	oh.MinorOperatingSystemVersion = 0
	oh64.MajorImageVersion = 1
	oh.MajorImageVersion = 1
	oh64.MinorImageVersion = 0
	oh.MinorImageVersion = 0
	oh64.MajorSubsystemVersion = 4
	oh.MajorSubsystemVersion = 4
	oh64.MinorSubsystemVersion = 0
	oh.MinorSubsystemVersion = 0
	oh64.SizeOfImage = pefile.nextSectOffset
	oh.SizeOfImage = pefile.nextSectOffset
	oh64.SizeOfHeaders = uint32(PEFILEHEADR)
	oh.SizeOfHeaders = uint32(PEFILEHEADR)
	if windowsgui {
		oh64.Subsystem = IMAGE_SUBSYSTEM_WINDOWS_GUI
		oh.Subsystem = IMAGE_SUBSYSTEM_WINDOWS_GUI
	} else {
		oh64.Subsystem = IMAGE_SUBSYSTEM_WINDOWS_CUI
		oh.Subsystem = IMAGE_SUBSYSTEM_WINDOWS_CUI
	}

	// Disable stack growth as we don't want Windows to
	// fiddle with the thread stack limits, which we set
	// ourselves to circumvent the stack checks in the
	// Windows exception dispatcher.
	// Commit size must be strictly less than reserve
	// size otherwise reserve will be rounded up to a
	// larger size, as verified with VMMap.

	// On 64-bit, we always reserve 2MB stacks. "Pure" Go code is
	// okay with much smaller stacks, but the syscall package
	// makes it easy to call into arbitrary C code without cgo,
	// and system calls even in "pure" Go code are actually C
	// calls that may need more stack than we think.
	//
	// The default stack reserve size affects only the main
	// thread, ctrlhandler thread, and profileloop thread. For
	// these, it must be greater than the stack size assumed by
	// externalthreadhandler.
	//
	// For other threads we specify stack size in runtime explicitly.
	// For these, the reserve must match STACKSIZE in
	// runtime/cgo/gcc_windows_{386,amd64}.c and the correspondent
	// CreateThread parameter in runtime.newosproc.
	oh64.SizeOfStackReserve = 0x00200000
	oh64.SizeOfStackCommit = 0x00200000 - 0x2000 // account for 2 guard pages

	// 32-bit is trickier since there much less address space to
	// work with. Here we use large stacks only in cgo binaries as
	// a compromise.
	if !iscgo {
		oh.SizeOfStackReserve = 0x00020000
		oh.SizeOfStackCommit = 0x00001000
	} else {
		oh.SizeOfStackReserve = 0x00100000
		oh.SizeOfStackCommit = 0x00100000 - 0x2000
	}

	oh64.SizeOfHeapReserve = 0x00100000
	oh.SizeOfHeapReserve = 0x00100000
	oh64.SizeOfHeapCommit = 0x00001000
	oh.SizeOfHeapCommit = 0x00001000
	oh64.NumberOfRvaAndSizes = 16
	oh.NumberOfRvaAndSizes = 16

	pewrite()
}
