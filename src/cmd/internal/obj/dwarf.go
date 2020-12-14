// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Writes dwarf information to object files.

package obj

import (
	"cmd/internal/dwarf"
	"cmd/internal/objabi"
	"cmd/internal/src"
	"fmt"
	"sort"
	"sync"
)

// Generate a sequence of opcodes that is as short as possible.
// See section 6.2.5
const (
	LINE_BASE   = -4
	LINE_RANGE  = 10
	PC_RANGE    = (255 - OPCODE_BASE) / LINE_RANGE
	OPCODE_BASE = 11
)

// generateDebugLinesSymbol fills the debug lines symbol of a given function.
//
// It's worth noting that this function doesn't generate the full debug_lines
// DWARF section, saving that for the linker. This function just generates the
// state machine part of debug_lines. The full table is generated by the
// linker.  Also, we use the file numbers from the full package (not just the
// function in question) when generating the state machine. We do this so we
// don't have to do a fixup on the indices when writing the full section.
func (ctxt *Link) generateDebugLinesSymbol(s, lines *LSym) {
	dctxt := dwCtxt{ctxt}

	// Emit a LNE_set_address extended opcode, so as to establish the
	// starting text address of this function.
	dctxt.AddUint8(lines, 0)
	dwarf.Uleb128put(dctxt, lines, 1+int64(ctxt.Arch.PtrSize))
	dctxt.AddUint8(lines, dwarf.DW_LNE_set_address)
	dctxt.AddAddress(lines, s, 0)

	// Set up the debug_lines state machine to the default values
	// we expect at the start of a new sequence.
	stmt := true
	line := int64(1)
	pc := s.Func().Text.Pc
	var lastpc int64 // last PC written to line table, not last PC in func
	name := ""
	prologue, wrotePrologue := false, false
	// Walk the progs, generating the DWARF table.
	for p := s.Func().Text; p != nil; p = p.Link {
		prologue = prologue || (p.Pos.Xlogue() == src.PosPrologueEnd)
		// If we're not at a real instruction, keep looping!
		if p.Pos.Line() == 0 || (p.Link != nil && p.Link.Pc == p.Pc) {
			continue
		}
		newStmt := p.Pos.IsStmt() != src.PosNotStmt
		newName, newLine := linkgetlineFromPos(ctxt, p.Pos)

		// Output debug info.
		wrote := false
		if name != newName {
			newFile := ctxt.PosTable.FileIndex(newName) + 1 // 1 indexing for the table.
			dctxt.AddUint8(lines, dwarf.DW_LNS_set_file)
			dwarf.Uleb128put(dctxt, lines, int64(newFile))
			name = newName
			wrote = true
		}
		if prologue && !wrotePrologue {
			dctxt.AddUint8(lines, uint8(dwarf.DW_LNS_set_prologue_end))
			wrotePrologue = true
			wrote = true
		}
		if stmt != newStmt {
			dctxt.AddUint8(lines, uint8(dwarf.DW_LNS_negate_stmt))
			stmt = newStmt
			wrote = true
		}

		if line != int64(newLine) || wrote {
			pcdelta := p.Pc - pc
			lastpc = p.Pc
			putpclcdelta(ctxt, dctxt, lines, uint64(pcdelta), int64(newLine)-line)
			line, pc = int64(newLine), p.Pc
		}
	}

	// Because these symbols will be concatenated together by the
	// linker, we need to reset the state machine that controls the
	// debug symbols. Do this using an end-of-sequence operator.
	//
	// Note: at one point in time, Delve did not support multiple end
	// sequence ops within a compilation unit (bug for this:
	// https://github.com/go-delve/delve/issues/1694), however the bug
	// has since been fixed (Oct 2019).
	//
	// Issue 38192: the DWARF standard specifies that when you issue
	// an end-sequence op, the PC value should be one past the last
	// text address in the translation unit, so apply a delta to the
	// text address before the end sequence op. If this isn't done,
	// GDB will assign a line number of zero the last row in the line
	// table, which we don't want.
	lastlen := uint64(s.Size - (lastpc - s.Func().Text.Pc))
	dctxt.AddUint8(lines, dwarf.DW_LNS_advance_pc)
	dwarf.Uleb128put(dctxt, lines, int64(lastlen))
	dctxt.AddUint8(lines, 0) // start extended opcode
	dwarf.Uleb128put(dctxt, lines, 1)
	dctxt.AddUint8(lines, dwarf.DW_LNE_end_sequence)
}

func putpclcdelta(linkctxt *Link, dctxt dwCtxt, s *LSym, deltaPC uint64, deltaLC int64) {
	// Choose a special opcode that minimizes the number of bytes needed to
	// encode the remaining PC delta and LC delta.
	var opcode int64
	if deltaLC < LINE_BASE {
		if deltaPC >= PC_RANGE {
			opcode = OPCODE_BASE + (LINE_RANGE * PC_RANGE)
		} else {
			opcode = OPCODE_BASE + (LINE_RANGE * int64(deltaPC))
		}
	} else if deltaLC < LINE_BASE+LINE_RANGE {
		if deltaPC >= PC_RANGE {
			opcode = OPCODE_BASE + (deltaLC - LINE_BASE) + (LINE_RANGE * PC_RANGE)
			if opcode > 255 {
				opcode -= LINE_RANGE
			}
		} else {
			opcode = OPCODE_BASE + (deltaLC - LINE_BASE) + (LINE_RANGE * int64(deltaPC))
		}
	} else {
		if deltaPC <= PC_RANGE {
			opcode = OPCODE_BASE + (LINE_RANGE - 1) + (LINE_RANGE * int64(deltaPC))
			if opcode > 255 {
				opcode = 255
			}
		} else {
			// Use opcode 249 (pc+=23, lc+=5) or 255 (pc+=24, lc+=1).
			//
			// Let x=deltaPC-PC_RANGE.  If we use opcode 255, x will be the remaining
			// deltaPC that we need to encode separately before emitting 255.  If we
			// use opcode 249, we will need to encode x+1.  If x+1 takes one more
			// byte to encode than x, then we use opcode 255.
			//
			// In all other cases x and x+1 take the same number of bytes to encode,
			// so we use opcode 249, which may save us a byte in encoding deltaLC,
			// for similar reasons.
			switch deltaPC - PC_RANGE {
			// PC_RANGE is the largest deltaPC we can encode in one byte, using
			// DW_LNS_const_add_pc.
			//
			// (1<<16)-1 is the largest deltaPC we can encode in three bytes, using
			// DW_LNS_fixed_advance_pc.
			//
			// (1<<(7n))-1 is the largest deltaPC we can encode in n+1 bytes for
			// n=1,3,4,5,..., using DW_LNS_advance_pc.
			case PC_RANGE, (1 << 7) - 1, (1 << 16) - 1, (1 << 21) - 1, (1 << 28) - 1,
				(1 << 35) - 1, (1 << 42) - 1, (1 << 49) - 1, (1 << 56) - 1, (1 << 63) - 1:
				opcode = 255
			default:
				opcode = OPCODE_BASE + LINE_RANGE*PC_RANGE - 1 // 249
			}
		}
	}
	if opcode < OPCODE_BASE || opcode > 255 {
		panic(fmt.Sprintf("produced invalid special opcode %d", opcode))
	}

	// Subtract from deltaPC and deltaLC the amounts that the opcode will add.
	deltaPC -= uint64((opcode - OPCODE_BASE) / LINE_RANGE)
	deltaLC -= (opcode-OPCODE_BASE)%LINE_RANGE + LINE_BASE

	// Encode deltaPC.
	if deltaPC != 0 {
		if deltaPC <= PC_RANGE {
			// Adjust the opcode so that we can use the 1-byte DW_LNS_const_add_pc
			// instruction.
			opcode -= LINE_RANGE * int64(PC_RANGE-deltaPC)
			if opcode < OPCODE_BASE {
				panic(fmt.Sprintf("produced invalid special opcode %d", opcode))
			}
			dctxt.AddUint8(s, dwarf.DW_LNS_const_add_pc)
		} else if (1<<14) <= deltaPC && deltaPC < (1<<16) {
			dctxt.AddUint8(s, dwarf.DW_LNS_fixed_advance_pc)
			dctxt.AddUint16(s, uint16(deltaPC))
		} else {
			dctxt.AddUint8(s, dwarf.DW_LNS_advance_pc)
			dwarf.Uleb128put(dctxt, s, int64(deltaPC))
		}
	}

	// Encode deltaLC.
	if deltaLC != 0 {
		dctxt.AddUint8(s, dwarf.DW_LNS_advance_line)
		dwarf.Sleb128put(dctxt, s, deltaLC)
	}

	// Output the special opcode.
	dctxt.AddUint8(s, uint8(opcode))
}

// implement dwarf.Context
type dwCtxt struct{ *Link }

func (c dwCtxt) PtrSize() int {
	return c.Arch.PtrSize
}
func (c dwCtxt) AddInt(s dwarf.Sym, size int, i int64) {
	ls := s.(*LSym)
	ls.WriteInt(c.Link, ls.Size, size, i)
}
func (c dwCtxt) AddUint16(s dwarf.Sym, i uint16) {
	c.AddInt(s, 2, int64(i))
}
func (c dwCtxt) AddUint8(s dwarf.Sym, i uint8) {
	b := []byte{byte(i)}
	c.AddBytes(s, b)
}
func (c dwCtxt) AddBytes(s dwarf.Sym, b []byte) {
	ls := s.(*LSym)
	ls.WriteBytes(c.Link, ls.Size, b)
}
func (c dwCtxt) AddString(s dwarf.Sym, v string) {
	ls := s.(*LSym)
	ls.WriteString(c.Link, ls.Size, len(v), v)
	ls.WriteInt(c.Link, ls.Size, 1, 0)
}
func (c dwCtxt) AddAddress(s dwarf.Sym, data interface{}, value int64) {
	ls := s.(*LSym)
	size := c.PtrSize()
	if data != nil {
		rsym := data.(*LSym)
		ls.WriteAddr(c.Link, ls.Size, size, rsym, value)
	} else {
		ls.WriteInt(c.Link, ls.Size, size, value)
	}
}
func (c dwCtxt) AddCURelativeAddress(s dwarf.Sym, data interface{}, value int64) {
	ls := s.(*LSym)
	rsym := data.(*LSym)
	ls.WriteCURelativeAddr(c.Link, ls.Size, rsym, value)
}
func (c dwCtxt) AddSectionOffset(s dwarf.Sym, size int, t interface{}, ofs int64) {
	panic("should be used only in the linker")
}
func (c dwCtxt) AddDWARFAddrSectionOffset(s dwarf.Sym, t interface{}, ofs int64) {
	size := 4
	if isDwarf64(c.Link) {
		size = 8
	}

	ls := s.(*LSym)
	rsym := t.(*LSym)
	ls.WriteAddr(c.Link, ls.Size, size, rsym, ofs)
	r := &ls.R[len(ls.R)-1]
	r.Type = objabi.R_DWARFSECREF
}

func (c dwCtxt) AddFileRef(s dwarf.Sym, f interface{}) {
	ls := s.(*LSym)
	rsym := f.(*LSym)
	fidx := c.Link.PosTable.FileIndex(rsym.Name)
	// Note the +1 here -- the value we're writing is going to be an
	// index into the DWARF line table file section, whose entries
	// are numbered starting at 1, not 0.
	ls.WriteInt(c.Link, ls.Size, 4, int64(fidx+1))
}

func (c dwCtxt) CurrentOffset(s dwarf.Sym) int64 {
	ls := s.(*LSym)
	return ls.Size
}

// Here "from" is a symbol corresponding to an inlined or concrete
// function, "to" is the symbol for the corresponding abstract
// function, and "dclIdx" is the index of the symbol of interest with
// respect to the Dcl slice of the original pre-optimization version
// of the inlined function.
func (c dwCtxt) RecordDclReference(from dwarf.Sym, to dwarf.Sym, dclIdx int, inlIndex int) {
	ls := from.(*LSym)
	tls := to.(*LSym)
	ridx := len(ls.R) - 1
	c.Link.DwFixups.ReferenceChildDIE(ls, ridx, tls, dclIdx, inlIndex)
}

func (c dwCtxt) RecordChildDieOffsets(s dwarf.Sym, vars []*dwarf.Var, offsets []int32) {
	ls := s.(*LSym)
	c.Link.DwFixups.RegisterChildDIEOffsets(ls, vars, offsets)
}

func (c dwCtxt) Logf(format string, args ...interface{}) {
	c.Link.Logf(format, args...)
}

func isDwarf64(ctxt *Link) bool {
	return ctxt.Headtype == objabi.Haix
}

func (ctxt *Link) dwarfSym(s *LSym) (dwarfInfoSym, dwarfLocSym, dwarfRangesSym, dwarfAbsFnSym, dwarfDebugLines *LSym) {
	if s.Type != objabi.STEXT {
		ctxt.Diag("dwarfSym of non-TEXT %v", s)
	}
	fn := s.Func()
	if fn.dwarfInfoSym == nil {
		fn.dwarfInfoSym = &LSym{
			Type: objabi.SDWARFFCN,
		}
		if ctxt.Flag_locationlists {
			fn.dwarfLocSym = &LSym{
				Type: objabi.SDWARFLOC,
			}
		}
		fn.dwarfRangesSym = &LSym{
			Type: objabi.SDWARFRANGE,
		}
		fn.dwarfDebugLinesSym = &LSym{
			Type: objabi.SDWARFLINES,
		}
		if s.WasInlined() {
			fn.dwarfAbsFnSym = ctxt.DwFixups.AbsFuncDwarfSym(s)
		}
	}
	return fn.dwarfInfoSym, fn.dwarfLocSym, fn.dwarfRangesSym, fn.dwarfAbsFnSym, fn.dwarfDebugLinesSym
}

func (s *LSym) Length(dwarfContext interface{}) int64 {
	return s.Size
}

// fileSymbol returns a symbol corresponding to the source file of the
// first instruction (prog) of the specified function. This will
// presumably be the file in which the function is defined.
func (ctxt *Link) fileSymbol(fn *LSym) *LSym {
	p := fn.Func().Text
	if p != nil {
		f, _ := linkgetlineFromPos(ctxt, p.Pos)
		fsym := ctxt.Lookup(f)
		return fsym
	}
	return nil
}

// populateDWARF fills in the DWARF Debugging Information Entries for
// TEXT symbol 's'. The various DWARF symbols must already have been
// initialized in InitTextSym.
func (ctxt *Link) populateDWARF(curfn interface{}, s *LSym, myimportpath string) {
	info, loc, ranges, absfunc, lines := ctxt.dwarfSym(s)
	if info.Size != 0 {
		ctxt.Diag("makeFuncDebugEntry double process %v", s)
	}
	var scopes []dwarf.Scope
	var inlcalls dwarf.InlCalls
	if ctxt.DebugInfo != nil {
		scopes, inlcalls = ctxt.DebugInfo(s, info, curfn)
	}
	var err error
	dwctxt := dwCtxt{ctxt}
	filesym := ctxt.fileSymbol(s)
	fnstate := &dwarf.FnState{
		Name:          s.Name,
		Importpath:    myimportpath,
		Info:          info,
		Filesym:       filesym,
		Loc:           loc,
		Ranges:        ranges,
		Absfn:         absfunc,
		StartPC:       s,
		Size:          s.Size,
		External:      !s.Static(),
		Scopes:        scopes,
		InlCalls:      inlcalls,
		UseBASEntries: ctxt.UseBASEntries,
	}
	if absfunc != nil {
		err = dwarf.PutAbstractFunc(dwctxt, fnstate)
		if err != nil {
			ctxt.Diag("emitting DWARF for %s failed: %v", s.Name, err)
		}
		err = dwarf.PutConcreteFunc(dwctxt, fnstate)
	} else {
		err = dwarf.PutDefaultFunc(dwctxt, fnstate)
	}
	if err != nil {
		ctxt.Diag("emitting DWARF for %s failed: %v", s.Name, err)
	}
	// Fill in the debug lines symbol.
	ctxt.generateDebugLinesSymbol(s, lines)
}

// DwarfIntConst creates a link symbol for an integer constant with the
// given name, type and value.
func (ctxt *Link) DwarfIntConst(myimportpath, name, typename string, val int64) {
	if myimportpath == "" {
		return
	}
	s := ctxt.LookupInit(dwarf.ConstInfoPrefix+myimportpath, func(s *LSym) {
		s.Type = objabi.SDWARFCONST
		ctxt.Data = append(ctxt.Data, s)
	})
	dwarf.PutIntConst(dwCtxt{ctxt}, s, ctxt.Lookup(dwarf.InfoPrefix+typename), myimportpath+"."+name, val)
}

func (ctxt *Link) DwarfAbstractFunc(curfn interface{}, s *LSym, myimportpath string) {
	absfn := ctxt.DwFixups.AbsFuncDwarfSym(s)
	if absfn.Size != 0 {
		ctxt.Diag("internal error: DwarfAbstractFunc double process %v", s)
	}
	if s.Func() == nil {
		s.NewFuncInfo()
	}
	scopes, _ := ctxt.DebugInfo(s, absfn, curfn)
	dwctxt := dwCtxt{ctxt}
	filesym := ctxt.fileSymbol(s)
	fnstate := dwarf.FnState{
		Name:          s.Name,
		Importpath:    myimportpath,
		Info:          absfn,
		Filesym:       filesym,
		Absfn:         absfn,
		External:      !s.Static(),
		Scopes:        scopes,
		UseBASEntries: ctxt.UseBASEntries,
	}
	if err := dwarf.PutAbstractFunc(dwctxt, &fnstate); err != nil {
		ctxt.Diag("emitting DWARF for %s failed: %v", s.Name, err)
	}
}

// This table is designed to aid in the creation of references between
// DWARF subprogram DIEs.
//
// In most cases when one DWARF DIE has to refer to another DWARF DIE,
// the target of the reference has an LSym, which makes it easy to use
// the existing relocation mechanism. For DWARF inlined routine DIEs,
// however, the subprogram DIE has to refer to a child
// parameter/variable DIE of the abstract subprogram. This child DIE
// doesn't have an LSym, and also of interest is the fact that when
// DWARF generation is happening for inlined function F within caller
// G, it's possible that DWARF generation hasn't happened yet for F,
// so there is no way to know the offset of a child DIE within F's
// abstract function. Making matters more complex, each inlined
// instance of F may refer to a subset of the original F's variables
// (depending on what happens with optimization, some vars may be
// eliminated).
//
// The fixup table below helps overcome this hurdle. At the point
// where a parameter/variable reference is made (via a call to
// "ReferenceChildDIE"), a fixup record is generate that records
// the relocation that is targeting that child variable. At a later
// point when the abstract function DIE is emitted, there will be
// a call to "RegisterChildDIEOffsets", at which point the offsets
// needed to apply fixups are captured. Finally, once the parallel
// portion of the compilation is done, fixups can actually be applied
// during the "Finalize" method (this can't be done during the
// parallel portion of the compile due to the possibility of data
// races).
//
// This table is also used to record the "precursor" function node for
// each function that is the target of an inline -- child DIE references
// have to be made with respect to the original pre-optimization
// version of the function (to allow for the fact that each inlined
// body may be optimized differently).
type DwarfFixupTable struct {
	ctxt      *Link
	mu        sync.Mutex
	symtab    map[*LSym]int // maps abstract fn LSYM to index in svec
	svec      []symFixups
	precursor map[*LSym]fnState // maps fn Lsym to precursor Node, absfn sym
}

type symFixups struct {
	fixups   []relFixup
	doffsets []declOffset
	inlIndex int32
	defseen  bool
}

type declOffset struct {
	// Index of variable within DCL list of pre-optimization function
	dclIdx int32
	// Offset of var's child DIE with respect to containing subprogram DIE
	offset int32
}

type relFixup struct {
	refsym *LSym
	relidx int32
	dclidx int32
}

type fnState struct {
	// precursor function (really *gc.Node)
	precursor interface{}
	// abstract function symbol
	absfn *LSym
}

func NewDwarfFixupTable(ctxt *Link) *DwarfFixupTable {
	return &DwarfFixupTable{
		ctxt:      ctxt,
		symtab:    make(map[*LSym]int),
		precursor: make(map[*LSym]fnState),
	}
}

func (ft *DwarfFixupTable) GetPrecursorFunc(s *LSym) interface{} {
	if fnstate, found := ft.precursor[s]; found {
		return fnstate.precursor
	}
	return nil
}

func (ft *DwarfFixupTable) SetPrecursorFunc(s *LSym, fn interface{}) {
	if _, found := ft.precursor[s]; found {
		ft.ctxt.Diag("internal error: DwarfFixupTable.SetPrecursorFunc double call on %v", s)
	}

	// initialize abstract function symbol now. This is done here so
	// as to avoid data races later on during the parallel portion of
	// the back end.
	absfn := ft.ctxt.LookupDerived(s, dwarf.InfoPrefix+s.Name+dwarf.AbstractFuncSuffix)
	absfn.Set(AttrDuplicateOK, true)
	absfn.Type = objabi.SDWARFABSFCN
	ft.ctxt.Data = append(ft.ctxt.Data, absfn)

	// In the case of "late" inlining (inlines that happen during
	// wrapper generation as opposed to the main inlining phase) it's
	// possible that we didn't cache the abstract function sym for the
	// text symbol -- do so now if needed. See issue 38068.
	if fn := s.Func(); fn != nil && fn.dwarfAbsFnSym == nil {
		fn.dwarfAbsFnSym = absfn
	}

	ft.precursor[s] = fnState{precursor: fn, absfn: absfn}
}

// Make a note of a child DIE reference: relocation 'ridx' within symbol 's'
// is targeting child 'c' of DIE with symbol 'tgt'.
func (ft *DwarfFixupTable) ReferenceChildDIE(s *LSym, ridx int, tgt *LSym, dclidx int, inlIndex int) {
	// Protect against concurrent access if multiple backend workers
	ft.mu.Lock()
	defer ft.mu.Unlock()

	// Create entry for symbol if not already present.
	idx, found := ft.symtab[tgt]
	if !found {
		ft.svec = append(ft.svec, symFixups{inlIndex: int32(inlIndex)})
		idx = len(ft.svec) - 1
		ft.symtab[tgt] = idx
	}

	// Do we have child DIE offsets available? If so, then apply them,
	// otherwise create a fixup record.
	sf := &ft.svec[idx]
	if len(sf.doffsets) > 0 {
		found := false
		for _, do := range sf.doffsets {
			if do.dclIdx == int32(dclidx) {
				off := do.offset
				s.R[ridx].Add += int64(off)
				found = true
				break
			}
		}
		if !found {
			ft.ctxt.Diag("internal error: DwarfFixupTable.ReferenceChildDIE unable to locate child DIE offset for dclIdx=%d src=%v tgt=%v", dclidx, s, tgt)
		}
	} else {
		sf.fixups = append(sf.fixups, relFixup{s, int32(ridx), int32(dclidx)})
	}
}

// Called once DWARF generation is complete for a given abstract function,
// whose children might have been referenced via a call above. Stores
// the offsets for any child DIEs (vars, params) so that they can be
// consumed later in on DwarfFixupTable.Finalize, which applies any
// outstanding fixups.
func (ft *DwarfFixupTable) RegisterChildDIEOffsets(s *LSym, vars []*dwarf.Var, coffsets []int32) {
	// Length of these two slices should agree
	if len(vars) != len(coffsets) {
		ft.ctxt.Diag("internal error: RegisterChildDIEOffsets vars/offsets length mismatch")
		return
	}

	// Generate the slice of declOffset's based in vars/coffsets
	doffsets := make([]declOffset, len(coffsets))
	for i := range coffsets {
		doffsets[i].dclIdx = vars[i].ChildIndex
		doffsets[i].offset = coffsets[i]
	}

	ft.mu.Lock()
	defer ft.mu.Unlock()

	// Store offsets for this symbol.
	idx, found := ft.symtab[s]
	if !found {
		sf := symFixups{inlIndex: -1, defseen: true, doffsets: doffsets}
		ft.svec = append(ft.svec, sf)
		ft.symtab[s] = len(ft.svec) - 1
	} else {
		sf := &ft.svec[idx]
		sf.doffsets = doffsets
		sf.defseen = true
	}
}

func (ft *DwarfFixupTable) processFixups(slot int, s *LSym) {
	sf := &ft.svec[slot]
	for _, f := range sf.fixups {
		dfound := false
		for _, doffset := range sf.doffsets {
			if doffset.dclIdx == f.dclidx {
				f.refsym.R[f.relidx].Add += int64(doffset.offset)
				dfound = true
				break
			}
		}
		if !dfound {
			ft.ctxt.Diag("internal error: DwarfFixupTable has orphaned fixup on %v targeting %v relidx=%d dclidx=%d", f.refsym, s, f.relidx, f.dclidx)
		}
	}
}

// return the LSym corresponding to the 'abstract subprogram' DWARF
// info entry for a function.
func (ft *DwarfFixupTable) AbsFuncDwarfSym(fnsym *LSym) *LSym {
	// Protect against concurrent access if multiple backend workers
	ft.mu.Lock()
	defer ft.mu.Unlock()

	if fnstate, found := ft.precursor[fnsym]; found {
		return fnstate.absfn
	}
	ft.ctxt.Diag("internal error: AbsFuncDwarfSym requested for %v, not seen during inlining", fnsym)
	return nil
}

// Called after all functions have been compiled; the main job of this
// function is to identify cases where there are outstanding fixups.
// This scenario crops up when we have references to variables of an
// inlined routine, but that routine is defined in some other package.
// This helper walks through and locate these fixups, then invokes a
// helper to create an abstract subprogram DIE for each one.
func (ft *DwarfFixupTable) Finalize(myimportpath string, trace bool) {
	if trace {
		ft.ctxt.Logf("DwarfFixupTable.Finalize invoked for %s\n", myimportpath)
	}

	// Collect up the keys from the precursor map, then sort the
	// resulting list (don't want to rely on map ordering here).
	fns := make([]*LSym, len(ft.precursor))
	idx := 0
	for fn := range ft.precursor {
		fns[idx] = fn
		idx++
	}
	sort.Sort(BySymName(fns))

	// Should not be called during parallel portion of compilation.
	if ft.ctxt.InParallel {
		ft.ctxt.Diag("internal error: DwarfFixupTable.Finalize call during parallel backend")
	}

	// Generate any missing abstract functions.
	for _, s := range fns {
		absfn := ft.AbsFuncDwarfSym(s)
		slot, found := ft.symtab[absfn]
		if !found || !ft.svec[slot].defseen {
			ft.ctxt.GenAbstractFunc(s)
		}
	}

	// Apply fixups.
	for _, s := range fns {
		absfn := ft.AbsFuncDwarfSym(s)
		slot, found := ft.symtab[absfn]
		if !found {
			ft.ctxt.Diag("internal error: DwarfFixupTable.Finalize orphan abstract function for %v", s)
		} else {
			ft.processFixups(slot, s)
		}
	}
}

type BySymName []*LSym

func (s BySymName) Len() int           { return len(s) }
func (s BySymName) Less(i, j int) bool { return s[i].Name < s[j].Name }
func (s BySymName) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
