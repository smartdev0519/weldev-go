// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ir

import (
	"cmd/compile/internal/base"
	"cmd/compile/internal/types"
	"cmd/internal/objabi"
	"cmd/internal/src"
	"fmt"
	"go/constant"
)

// Name holds Node fields used only by named nodes (ONAME, OTYPE, some OLITERAL).
type Name struct {
	PkgName *PkgName // real package for import . names
	// For a local variable (not param) or extern, the initializing assignment (OAS or OAS2).
	// For a closure var, the ONAME node of the outer captured variable
	Defn Node
	// The ODCLFUNC node (for a static function/method or a closure) in which
	// local variable or param is declared.
	Curfn     Node
	Param     *Param // additional fields for ONAME, OTYPE
	Decldepth int32  // declaration loop depth, increased for every loop or label
	// Unique number for ONAME nodes within a function. Function outputs
	// (results) are numbered starting at one, followed by function inputs
	// (parameters), and then local variables. Vargen is used to distinguish
	// local variables/params with the same name.
	Vargen int32
	flags  bitset16
}

type Param struct {
	Ntype    Node
	Heapaddr Node // temp holding heap address of param

	// ONAME PAUTOHEAP
	Stackcopy Node // the PPARAM/PPARAMOUT on-stack slot (moved func params only)

	// ONAME closure linkage
	// Consider:
	//
	//	func f() {
	//		x := 1 // x1
	//		func() {
	//			use(x) // x2
	//			func() {
	//				use(x) // x3
	//				--- parser is here ---
	//			}()
	//		}()
	//	}
	//
	// There is an original declaration of x and then a chain of mentions of x
	// leading into the current function. Each time x is mentioned in a new closure,
	// we create a variable representing x for use in that specific closure,
	// since the way you get to x is different in each closure.
	//
	// Let's number the specific variables as shown in the code:
	// x1 is the original x, x2 is when mentioned in the closure,
	// and x3 is when mentioned in the closure in the closure.
	//
	// We keep these linked (assume N > 1):
	//
	//   - x1.Defn = original declaration statement for x (like most variables)
	//   - x1.Innermost = current innermost closure x (in this case x3), or nil for none
	//   - x1.IsClosureVar() = false
	//
	//   - xN.Defn = x1, N > 1
	//   - xN.IsClosureVar() = true, N > 1
	//   - x2.Outer = nil
	//   - xN.Outer = x(N-1), N > 2
	//
	//
	// When we look up x in the symbol table, we always get x1.
	// Then we can use x1.Innermost (if not nil) to get the x
	// for the innermost known closure function,
	// but the first reference in a closure will find either no x1.Innermost
	// or an x1.Innermost with .Funcdepth < Funcdepth.
	// In that case, a new xN must be created, linked in with:
	//
	//     xN.Defn = x1
	//     xN.Outer = x1.Innermost
	//     x1.Innermost = xN
	//
	// When we finish the function, we'll process its closure variables
	// and find xN and pop it off the list using:
	//
	//     x1 := xN.Defn
	//     x1.Innermost = xN.Outer
	//
	// We leave x1.Innermost set so that we can still get to the original
	// variable quickly. Not shown here, but once we're
	// done parsing a function and no longer need xN.Outer for the
	// lexical x reference links as described above, funcLit
	// recomputes xN.Outer as the semantic x reference link tree,
	// even filling in x in intermediate closures that might not
	// have mentioned it along the way to inner closures that did.
	// See funcLit for details.
	//
	// During the eventual compilation, then, for closure variables we have:
	//
	//     xN.Defn = original variable
	//     xN.Outer = variable captured in next outward scope
	//                to make closure where xN appears
	//
	// Because of the sharding of pieces of the node, x.Defn means x.Name.Defn
	// and x.Innermost/Outer means x.Name.Param.Innermost/Outer.
	Innermost Node
	Outer     Node

	// OTYPE & ONAME //go:embed info,
	// sharing storage to reduce gc.Param size.
	// Extra is nil, or else *Extra is a *paramType or an *embedFileList.
	Extra *interface{}
}

// NewNameAt returns a new ONAME Node associated with symbol s at position pos.
// The caller is responsible for setting n.Name.Curfn.
func NewNameAt(pos src.XPos, s *types.Sym) Node {
	if s == nil {
		base.Fatalf("newnamel nil")
	}

	var x struct {
		n node
		m Name
		p Param
	}
	n := &x.n
	n.SetName(&x.m)
	n.Name().Param = &x.p

	n.SetOp(ONAME)
	n.SetPos(pos)
	n.SetOrig(n)

	n.SetSym(s)
	return n
}

type paramType struct {
	flag  PragmaFlag
	alias bool
}

// Pragma returns the PragmaFlag for p, which must be for an OTYPE.
func (p *Param) Pragma() PragmaFlag {
	if p.Extra == nil {
		return 0
	}
	return (*p.Extra).(*paramType).flag
}

// SetPragma sets the PragmaFlag for p, which must be for an OTYPE.
func (p *Param) SetPragma(flag PragmaFlag) {
	if p.Extra == nil {
		if flag == 0 {
			return
		}
		p.Extra = new(interface{})
		*p.Extra = &paramType{flag: flag}
		return
	}
	(*p.Extra).(*paramType).flag = flag
}

// Alias reports whether p, which must be for an OTYPE, is a type alias.
func (p *Param) Alias() bool {
	if p.Extra == nil {
		return false
	}
	t, ok := (*p.Extra).(*paramType)
	if !ok {
		return false
	}
	return t.alias
}

// SetAlias sets whether p, which must be for an OTYPE, is a type alias.
func (p *Param) SetAlias(alias bool) {
	if p.Extra == nil {
		if !alias {
			return
		}
		p.Extra = new(interface{})
		*p.Extra = &paramType{alias: alias}
		return
	}
	(*p.Extra).(*paramType).alias = alias
}

type embedFileList []string

// EmbedFiles returns the list of embedded files for p,
// which must be for an ONAME var.
func (p *Param) EmbedFiles() []string {
	if p.Extra == nil {
		return nil
	}
	return *(*p.Extra).(*embedFileList)
}

// SetEmbedFiles sets the list of embedded files for p,
// which must be for an ONAME var.
func (p *Param) SetEmbedFiles(list []string) {
	if p.Extra == nil {
		if len(list) == 0 {
			return
		}
		f := embedFileList(list)
		p.Extra = new(interface{})
		*p.Extra = &f
		return
	}
	*(*p.Extra).(*embedFileList) = list
}

const (
	nameCaptured = 1 << iota // is the variable captured by a closure
	nameReadonly
	nameByval                 // is the variable captured by value or by reference
	nameNeedzero              // if it contains pointers, needs to be zeroed on function entry
	nameAutoTemp              // is the variable a temporary (implies no dwarf info. reset if escapes to heap)
	nameUsed                  // for variable declared and not used error
	nameIsClosureVar          // PAUTOHEAP closure pseudo-variable; original at n.Name.Defn
	nameIsOutputParamHeapAddr // pointer to a result parameter's heap copy
	nameAssigned              // is the variable ever assigned to
	nameAddrtaken             // address taken, even if not moved to heap
	nameInlFormal             // PAUTO created by inliner, derived from callee formal
	nameInlLocal              // PAUTO created by inliner, derived from callee local
	nameOpenDeferSlot         // if temporary var storing info for open-coded defers
	nameLibfuzzerExtraCounter // if PEXTERN should be assigned to __libfuzzer_extra_counters section
)

func (n *Name) Captured() bool              { return n.flags&nameCaptured != 0 }
func (n *Name) Readonly() bool              { return n.flags&nameReadonly != 0 }
func (n *Name) Byval() bool                 { return n.flags&nameByval != 0 }
func (n *Name) Needzero() bool              { return n.flags&nameNeedzero != 0 }
func (n *Name) AutoTemp() bool              { return n.flags&nameAutoTemp != 0 }
func (n *Name) Used() bool                  { return n.flags&nameUsed != 0 }
func (n *Name) IsClosureVar() bool          { return n.flags&nameIsClosureVar != 0 }
func (n *Name) IsOutputParamHeapAddr() bool { return n.flags&nameIsOutputParamHeapAddr != 0 }
func (n *Name) Assigned() bool              { return n.flags&nameAssigned != 0 }
func (n *Name) Addrtaken() bool             { return n.flags&nameAddrtaken != 0 }
func (n *Name) InlFormal() bool             { return n.flags&nameInlFormal != 0 }
func (n *Name) InlLocal() bool              { return n.flags&nameInlLocal != 0 }
func (n *Name) OpenDeferSlot() bool         { return n.flags&nameOpenDeferSlot != 0 }
func (n *Name) LibfuzzerExtraCounter() bool { return n.flags&nameLibfuzzerExtraCounter != 0 }

func (n *Name) SetCaptured(b bool)              { n.flags.set(nameCaptured, b) }
func (n *Name) SetReadonly(b bool)              { n.flags.set(nameReadonly, b) }
func (n *Name) SetByval(b bool)                 { n.flags.set(nameByval, b) }
func (n *Name) SetNeedzero(b bool)              { n.flags.set(nameNeedzero, b) }
func (n *Name) SetAutoTemp(b bool)              { n.flags.set(nameAutoTemp, b) }
func (n *Name) SetUsed(b bool)                  { n.flags.set(nameUsed, b) }
func (n *Name) SetIsClosureVar(b bool)          { n.flags.set(nameIsClosureVar, b) }
func (n *Name) SetIsOutputParamHeapAddr(b bool) { n.flags.set(nameIsOutputParamHeapAddr, b) }
func (n *Name) SetAssigned(b bool)              { n.flags.set(nameAssigned, b) }
func (n *Name) SetAddrtaken(b bool)             { n.flags.set(nameAddrtaken, b) }
func (n *Name) SetInlFormal(b bool)             { n.flags.set(nameInlFormal, b) }
func (n *Name) SetInlLocal(b bool)              { n.flags.set(nameInlLocal, b) }
func (n *Name) SetOpenDeferSlot(b bool)         { n.flags.set(nameOpenDeferSlot, b) }
func (n *Name) SetLibfuzzerExtraCounter(b bool) { n.flags.set(nameLibfuzzerExtraCounter, b) }

// MarkReadonly indicates that n is an ONAME with readonly contents.
func (n *node) MarkReadonly() {
	if n.Op() != ONAME {
		base.Fatalf("Node.MarkReadonly %v", n.Op())
	}
	n.Name().SetReadonly(true)
	// Mark the linksym as readonly immediately
	// so that the SSA backend can use this information.
	// It will be overridden later during dumpglobls.
	n.Sym().Linksym().Type = objabi.SRODATA
}

// Val returns the constant.Value for the node.
func (n *node) Val() constant.Value {
	if !n.HasVal() {
		return constant.MakeUnknown()
	}
	return *n.e.(*constant.Value)
}

// SetVal sets the constant.Value for the node,
// which must not have been used with SetOpt.
func (n *node) SetVal(v constant.Value) {
	if n.hasOpt() {
		base.Flag.LowerH = 1
		Dump("have Opt", n)
		base.Fatalf("have Opt")
	}
	if n.Op() == OLITERAL {
		AssertValidTypeForConst(n.Type(), v)
	}
	n.setHasVal(true)
	n.e = &v
}

// Int64Val returns n as an int64.
// n must be an integer or rune constant.
func (n *node) Int64Val() int64 {
	if !IsConst(n, constant.Int) {
		base.Fatalf("Int64Val(%v)", n)
	}
	x, ok := constant.Int64Val(n.Val())
	if !ok {
		base.Fatalf("Int64Val(%v)", n)
	}
	return x
}

// CanInt64 reports whether it is safe to call Int64Val() on n.
func (n *node) CanInt64() bool {
	if !IsConst(n, constant.Int) {
		return false
	}

	// if the value inside n cannot be represented as an int64, the
	// return value of Int64 is undefined
	_, ok := constant.Int64Val(n.Val())
	return ok
}

// Uint64Val returns n as an uint64.
// n must be an integer or rune constant.
func (n *node) Uint64Val() uint64 {
	if !IsConst(n, constant.Int) {
		base.Fatalf("Uint64Val(%v)", n)
	}
	x, ok := constant.Uint64Val(n.Val())
	if !ok {
		base.Fatalf("Uint64Val(%v)", n)
	}
	return x
}

// BoolVal returns n as a bool.
// n must be a boolean constant.
func (n *node) BoolVal() bool {
	if !IsConst(n, constant.Bool) {
		base.Fatalf("BoolVal(%v)", n)
	}
	return constant.BoolVal(n.Val())
}

// StringVal returns the value of a literal string Node as a string.
// n must be a string constant.
func (n *node) StringVal() string {
	if !IsConst(n, constant.String) {
		base.Fatalf("StringVal(%v)", n)
	}
	return constant.StringVal(n.Val())
}

// The Class of a variable/function describes the "storage class"
// of a variable or function. During parsing, storage classes are
// called declaration contexts.
type Class uint8

//go:generate stringer -type=Class
const (
	Pxxx      Class = iota // no class; used during ssa conversion to indicate pseudo-variables
	PEXTERN                // global variables
	PAUTO                  // local variables
	PAUTOHEAP              // local variables or parameters moved to heap
	PPARAM                 // input arguments
	PPARAMOUT              // output results
	PFUNC                  // global functions

	// Careful: Class is stored in three bits in Node.flags.
	_ = uint((1 << 3) - iota) // static assert for iota <= (1 << 3)
)

// A Pack is an identifier referring to an imported package.
type PkgName struct {
	miniNode
	sym  *types.Sym
	Pkg  *types.Pkg
	Used bool
}

func (p *PkgName) String() string                { return fmt.Sprint(p) }
func (p *PkgName) Format(s fmt.State, verb rune) { FmtNode(p, s, verb) }
func (p *PkgName) RawCopy() Node                 { c := *p; return &c }
func (p *PkgName) Sym() *types.Sym               { return p.sym }

func NewPkgName(pos src.XPos, sym *types.Sym, pkg *types.Pkg) *PkgName {
	p := &PkgName{sym: sym, Pkg: pkg}
	p.op = OPACK
	p.pos = pos
	return p
}
