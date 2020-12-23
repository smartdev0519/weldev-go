// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"cmd/compile/internal/base"
	"cmd/compile/internal/ir"
	"cmd/compile/internal/ssa"
	"cmd/compile/internal/types"
	"cmd/internal/obj"
	"sync"
)

var (
	// maximum size variable which we will allocate on the stack.
	// This limit is for explicit variable declarations like "var x T" or "x := ...".
	// Note: the flag smallframes can update this value.
	maxStackVarSize = int64(10 * 1024 * 1024)

	// maximum size of implicit variables that we will allocate on the stack.
	//   p := new(T)          allocating T on the stack
	//   p := &T{}            allocating T on the stack
	//   s := make([]T, n)    allocating [n]T on the stack
	//   s := []byte("...")   allocating [n]byte on the stack
	// Note: the flag smallframes can update this value.
	maxImplicitStackVarSize = int64(64 * 1024)

	// smallArrayBytes is the maximum size of an array which is considered small.
	// Small arrays will be initialized directly with a sequence of constant stores.
	// Large arrays will be initialized by copying from a static temp.
	// 256 bytes was chosen to minimize generated code + statictmp size.
	smallArrayBytes = int64(256)
)

// isRuntimePkg reports whether p is package runtime.
func isRuntimePkg(p *types.Pkg) bool {
	if base.Flag.CompilingRuntime && p == types.LocalPkg {
		return true
	}
	return p.Path == "runtime"
}

// isReflectPkg reports whether p is package reflect.
func isReflectPkg(p *types.Pkg) bool {
	if p == types.LocalPkg {
		return base.Ctxt.Pkgpath == "reflect"
	}
	return p.Path == "reflect"
}

// Slices in the runtime are represented by three components:
//
// type slice struct {
// 	ptr unsafe.Pointer
// 	len int
// 	cap int
// }
//
// Strings in the runtime are represented by two components:
//
// type string struct {
// 	ptr unsafe.Pointer
// 	len int
// }
//
// These variables are the offsets of fields and sizes of these structs.
var (
	slicePtrOffset int64
	sliceLenOffset int64
	sliceCapOffset int64

	sizeofSlice  int64
	sizeofString int64
)

var pragcgobuf [][]string

var decldepth int32

var inimport bool // set during import

var itabpkg *types.Pkg // fake pkg for itab entries

var itablinkpkg *types.Pkg // fake package for runtime itab entries

var Runtimepkg *types.Pkg // fake package runtime

var racepkg *types.Pkg // package runtime/race

var msanpkg *types.Pkg // package runtime/msan

var unsafepkg *types.Pkg // package unsafe

var trackpkg *types.Pkg // fake package for field tracking

var mappkg *types.Pkg // fake package for map zero value

var gopkg *types.Pkg // pseudo-package for method symbols on anonymous receiver types

var zerosize int64

var simtype [types.NTYPE]types.Kind

var (
	isInt     [types.NTYPE]bool
	isFloat   [types.NTYPE]bool
	isComplex [types.NTYPE]bool
	issimple  [types.NTYPE]bool
)

var (
	okforeq    [types.NTYPE]bool
	okforadd   [types.NTYPE]bool
	okforand   [types.NTYPE]bool
	okfornone  [types.NTYPE]bool
	okforbool  [types.NTYPE]bool
	okforcap   [types.NTYPE]bool
	okforlen   [types.NTYPE]bool
	okforarith [types.NTYPE]bool
)

var okforcmp [types.NTYPE]bool

var (
	okfor [ir.OEND][]bool
	iscmp [ir.OEND]bool
)

var (
	funcsymsmu sync.Mutex // protects funcsyms and associated package lookups (see func funcsym)
	funcsyms   []*types.Sym
)

var dclcontext ir.Class // PEXTERN/PAUTO

var Curfn *ir.Func

var Widthptr int

var Widthreg int

var typecheckok bool

var nodfp *ir.Name

// interface to back end

type Arch struct {
	LinkArch *obj.LinkArch

	REGSP     int
	MAXWIDTH  int64
	SoftFloat bool

	PadFrame func(int64) int64

	// ZeroRange zeroes a range of memory on stack. It is only inserted
	// at function entry, and it is ok to clobber registers.
	ZeroRange func(*Progs, *obj.Prog, int64, int64, *uint32) *obj.Prog

	Ginsnop      func(*Progs) *obj.Prog
	Ginsnopdefer func(*Progs) *obj.Prog // special ginsnop for deferreturn

	// SSAMarkMoves marks any MOVXconst ops that need to avoid clobbering flags.
	SSAMarkMoves func(*SSAGenState, *ssa.Block)

	// SSAGenValue emits Prog(s) for the Value.
	SSAGenValue func(*SSAGenState, *ssa.Value)

	// SSAGenBlock emits end-of-block Progs. SSAGenValue should be called
	// for all values in the block before SSAGenBlock.
	SSAGenBlock func(s *SSAGenState, b, next *ssa.Block)
}

var thearch Arch

var (
	staticuint64s *ir.Name
	zerobase      *ir.Name

	assertE2I,
	assertE2I2,
	assertI2I,
	assertI2I2,
	deferproc,
	deferprocStack,
	Deferreturn,
	Duffcopy,
	Duffzero,
	gcWriteBarrier,
	goschedguarded,
	growslice,
	msanread,
	msanwrite,
	msanmove,
	newobject,
	newproc,
	panicdivide,
	panicshift,
	panicdottypeE,
	panicdottypeI,
	panicnildottype,
	panicoverflow,
	raceread,
	racereadrange,
	racewrite,
	racewriterange,
	x86HasPOPCNT,
	x86HasSSE41,
	x86HasFMA,
	armHasVFPv4,
	arm64HasATOMICS,
	typedmemclr,
	typedmemmove,
	Udiv,
	writeBarrier,
	zerobaseSym *obj.LSym

	BoundsCheckFunc [ssa.BoundsKindCount]*obj.LSym
	ExtendCheckFunc [ssa.BoundsKindCount]*obj.LSym

	// Wasm
	WasmMove,
	WasmZero,
	WasmDiv,
	WasmTruncS,
	WasmTruncU,
	SigPanic *obj.LSym
)

// GCWriteBarrierReg maps from registers to gcWriteBarrier implementation LSyms.
var GCWriteBarrierReg map[int16]*obj.LSym
