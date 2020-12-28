// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reflectdata

import (
	"fmt"
	"sort"

	"cmd/compile/internal/base"
	"cmd/compile/internal/ir"
	"cmd/compile/internal/objw"
	"cmd/compile/internal/typecheck"
	"cmd/compile/internal/types"
	"cmd/internal/obj"
)

// isRegularMemory reports whether t can be compared/hashed as regular memory.
func isRegularMemory(t *types.Type) bool {
	a, _ := types.AlgType(t)
	return a == types.AMEM
}

// eqCanPanic reports whether == on type t could panic (has an interface somewhere).
// t must be comparable.
func eqCanPanic(t *types.Type) bool {
	switch t.Kind() {
	default:
		return false
	case types.TINTER:
		return true
	case types.TARRAY:
		return eqCanPanic(t.Elem())
	case types.TSTRUCT:
		for _, f := range t.FieldSlice() {
			if !f.Sym.IsBlank() && eqCanPanic(f.Type) {
				return true
			}
		}
		return false
	}
}

// AlgType is like algtype1, except it returns the fixed-width AMEMxx variants
// instead of the general AMEM kind when possible.
func AlgType(t *types.Type) types.AlgKind {
	a, _ := types.AlgType(t)
	if a == types.AMEM {
		switch t.Width {
		case 0:
			return types.AMEM0
		case 1:
			return types.AMEM8
		case 2:
			return types.AMEM16
		case 4:
			return types.AMEM32
		case 8:
			return types.AMEM64
		case 16:
			return types.AMEM128
		}
	}

	return a
}

// genhash returns a symbol which is the closure used to compute
// the hash of a value of type t.
// Note: the generated function must match runtime.typehash exactly.
func genhash(t *types.Type) *obj.LSym {
	switch AlgType(t) {
	default:
		// genhash is only called for types that have equality
		base.Fatalf("genhash %v", t)
	case types.AMEM0:
		return sysClosure("memhash0")
	case types.AMEM8:
		return sysClosure("memhash8")
	case types.AMEM16:
		return sysClosure("memhash16")
	case types.AMEM32:
		return sysClosure("memhash32")
	case types.AMEM64:
		return sysClosure("memhash64")
	case types.AMEM128:
		return sysClosure("memhash128")
	case types.ASTRING:
		return sysClosure("strhash")
	case types.AINTER:
		return sysClosure("interhash")
	case types.ANILINTER:
		return sysClosure("nilinterhash")
	case types.AFLOAT32:
		return sysClosure("f32hash")
	case types.AFLOAT64:
		return sysClosure("f64hash")
	case types.ACPLX64:
		return sysClosure("c64hash")
	case types.ACPLX128:
		return sysClosure("c128hash")
	case types.AMEM:
		// For other sizes of plain memory, we build a closure
		// that calls memhash_varlen. The size of the memory is
		// encoded in the first slot of the closure.
		closure := types.TypeSymLookup(fmt.Sprintf(".hashfunc%d", t.Width)).Linksym()
		if len(closure.P) > 0 { // already generated
			return closure
		}
		if memhashvarlen == nil {
			memhashvarlen = typecheck.LookupRuntimeFunc("memhash_varlen")
		}
		ot := 0
		ot = objw.SymPtr(closure, ot, memhashvarlen, 0)
		ot = objw.Uintptr(closure, ot, uint64(t.Width)) // size encoded in closure
		objw.Global(closure, int32(ot), obj.DUPOK|obj.RODATA)
		return closure
	case types.ASPECIAL:
		break
	}

	closure := TypeSymPrefix(".hashfunc", t).Linksym()
	if len(closure.P) > 0 { // already generated
		return closure
	}

	// Generate hash functions for subtypes.
	// There are cases where we might not use these hashes,
	// but in that case they will get dead-code eliminated.
	// (And the closure generated by genhash will also get
	// dead-code eliminated, as we call the subtype hashers
	// directly.)
	switch t.Kind() {
	case types.TARRAY:
		genhash(t.Elem())
	case types.TSTRUCT:
		for _, f := range t.FieldSlice() {
			genhash(f.Type)
		}
	}

	sym := TypeSymPrefix(".hash", t)
	if base.Flag.LowerR != 0 {
		fmt.Printf("genhash %v %v %v\n", closure, sym, t)
	}

	base.Pos = base.AutogeneratedPos // less confusing than end of input
	typecheck.DeclContext = ir.PEXTERN

	// func sym(p *T, h uintptr) uintptr
	args := []*ir.Field{
		ir.NewField(base.Pos, typecheck.Lookup("p"), nil, types.NewPtr(t)),
		ir.NewField(base.Pos, typecheck.Lookup("h"), nil, types.Types[types.TUINTPTR]),
	}
	results := []*ir.Field{ir.NewField(base.Pos, nil, nil, types.Types[types.TUINTPTR])}
	tfn := ir.NewFuncType(base.Pos, nil, args, results)

	fn := typecheck.DeclFunc(sym, tfn)
	np := ir.AsNode(tfn.Type().Params().Field(0).Nname)
	nh := ir.AsNode(tfn.Type().Params().Field(1).Nname)

	switch t.Kind() {
	case types.TARRAY:
		// An array of pure memory would be handled by the
		// standard algorithm, so the element type must not be
		// pure memory.
		hashel := hashfor(t.Elem())

		// for i := 0; i < nelem; i++
		ni := typecheck.Temp(types.Types[types.TINT])
		init := ir.NewAssignStmt(base.Pos, ni, ir.NewInt(0))
		cond := ir.NewBinaryExpr(base.Pos, ir.OLT, ni, ir.NewInt(t.NumElem()))
		post := ir.NewAssignStmt(base.Pos, ni, ir.NewBinaryExpr(base.Pos, ir.OADD, ni, ir.NewInt(1)))
		loop := ir.NewForStmt(base.Pos, nil, cond, post, nil)
		loop.PtrInit().Append(init)

		// h = hashel(&p[i], h)
		call := ir.NewCallExpr(base.Pos, ir.OCALL, hashel, nil)

		nx := ir.NewIndexExpr(base.Pos, np, ni)
		nx.SetBounded(true)
		na := typecheck.NodAddr(nx)
		call.Args.Append(na)
		call.Args.Append(nh)
		loop.Body.Append(ir.NewAssignStmt(base.Pos, nh, call))

		fn.Body.Append(loop)

	case types.TSTRUCT:
		// Walk the struct using memhash for runs of AMEM
		// and calling specific hash functions for the others.
		for i, fields := 0, t.FieldSlice(); i < len(fields); {
			f := fields[i]

			// Skip blank fields.
			if f.Sym.IsBlank() {
				i++
				continue
			}

			// Hash non-memory fields with appropriate hash function.
			if !isRegularMemory(f.Type) {
				hashel := hashfor(f.Type)
				call := ir.NewCallExpr(base.Pos, ir.OCALL, hashel, nil)
				nx := ir.NewSelectorExpr(base.Pos, ir.OXDOT, np, f.Sym) // TODO: fields from other packages?
				na := typecheck.NodAddr(nx)
				call.Args.Append(na)
				call.Args.Append(nh)
				fn.Body.Append(ir.NewAssignStmt(base.Pos, nh, call))
				i++
				continue
			}

			// Otherwise, hash a maximal length run of raw memory.
			size, next := memrun(t, i)

			// h = hashel(&p.first, size, h)
			hashel := hashmem(f.Type)
			call := ir.NewCallExpr(base.Pos, ir.OCALL, hashel, nil)
			nx := ir.NewSelectorExpr(base.Pos, ir.OXDOT, np, f.Sym) // TODO: fields from other packages?
			na := typecheck.NodAddr(nx)
			call.Args.Append(na)
			call.Args.Append(nh)
			call.Args.Append(ir.NewInt(size))
			fn.Body.Append(ir.NewAssignStmt(base.Pos, nh, call))

			i = next
		}
	}

	r := ir.NewReturnStmt(base.Pos, nil)
	r.Results.Append(nh)
	fn.Body.Append(r)

	if base.Flag.LowerR != 0 {
		ir.DumpList("genhash body", fn.Body)
	}

	typecheck.FinishFuncBody()

	fn.SetDupok(true)
	typecheck.Func(fn)

	ir.CurFunc = fn
	typecheck.Stmts(fn.Body)
	ir.CurFunc = nil

	if base.Debug.DclStack != 0 {
		types.CheckDclstack()
	}

	fn.SetNilCheckDisabled(true)
	typecheck.Target.Decls = append(typecheck.Target.Decls, fn)

	// Build closure. It doesn't close over any variables, so
	// it contains just the function pointer.
	objw.SymPtr(closure, 0, sym.Linksym(), 0)
	objw.Global(closure, int32(types.PtrSize), obj.DUPOK|obj.RODATA)

	return closure
}

func hashfor(t *types.Type) ir.Node {
	var sym *types.Sym

	switch a, _ := types.AlgType(t); a {
	case types.AMEM:
		base.Fatalf("hashfor with AMEM type")
	case types.AINTER:
		sym = ir.Pkgs.Runtime.Lookup("interhash")
	case types.ANILINTER:
		sym = ir.Pkgs.Runtime.Lookup("nilinterhash")
	case types.ASTRING:
		sym = ir.Pkgs.Runtime.Lookup("strhash")
	case types.AFLOAT32:
		sym = ir.Pkgs.Runtime.Lookup("f32hash")
	case types.AFLOAT64:
		sym = ir.Pkgs.Runtime.Lookup("f64hash")
	case types.ACPLX64:
		sym = ir.Pkgs.Runtime.Lookup("c64hash")
	case types.ACPLX128:
		sym = ir.Pkgs.Runtime.Lookup("c128hash")
	default:
		// Note: the caller of hashfor ensured that this symbol
		// exists and has a body by calling genhash for t.
		sym = TypeSymPrefix(".hash", t)
	}

	n := typecheck.NewName(sym)
	ir.MarkFunc(n)
	n.SetType(types.NewSignature(types.NoPkg, nil, []*types.Field{
		types.NewField(base.Pos, nil, types.NewPtr(t)),
		types.NewField(base.Pos, nil, types.Types[types.TUINTPTR]),
	}, []*types.Field{
		types.NewField(base.Pos, nil, types.Types[types.TUINTPTR]),
	}))
	return n
}

// sysClosure returns a closure which will call the
// given runtime function (with no closed-over variables).
func sysClosure(name string) *obj.LSym {
	s := typecheck.LookupRuntimeVar(name + "·f")
	if len(s.P) == 0 {
		f := typecheck.LookupRuntimeFunc(name)
		objw.SymPtr(s, 0, f, 0)
		objw.Global(s, int32(types.PtrSize), obj.DUPOK|obj.RODATA)
	}
	return s
}

// geneq returns a symbol which is the closure used to compute
// equality for two objects of type t.
func geneq(t *types.Type) *obj.LSym {
	switch AlgType(t) {
	case types.ANOEQ:
		// The runtime will panic if it tries to compare
		// a type with a nil equality function.
		return nil
	case types.AMEM0:
		return sysClosure("memequal0")
	case types.AMEM8:
		return sysClosure("memequal8")
	case types.AMEM16:
		return sysClosure("memequal16")
	case types.AMEM32:
		return sysClosure("memequal32")
	case types.AMEM64:
		return sysClosure("memequal64")
	case types.AMEM128:
		return sysClosure("memequal128")
	case types.ASTRING:
		return sysClosure("strequal")
	case types.AINTER:
		return sysClosure("interequal")
	case types.ANILINTER:
		return sysClosure("nilinterequal")
	case types.AFLOAT32:
		return sysClosure("f32equal")
	case types.AFLOAT64:
		return sysClosure("f64equal")
	case types.ACPLX64:
		return sysClosure("c64equal")
	case types.ACPLX128:
		return sysClosure("c128equal")
	case types.AMEM:
		// make equality closure. The size of the type
		// is encoded in the closure.
		closure := types.TypeSymLookup(fmt.Sprintf(".eqfunc%d", t.Width)).Linksym()
		if len(closure.P) != 0 {
			return closure
		}
		if memequalvarlen == nil {
			memequalvarlen = typecheck.LookupRuntimeVar("memequal_varlen") // asm func
		}
		ot := 0
		ot = objw.SymPtr(closure, ot, memequalvarlen, 0)
		ot = objw.Uintptr(closure, ot, uint64(t.Width))
		objw.Global(closure, int32(ot), obj.DUPOK|obj.RODATA)
		return closure
	case types.ASPECIAL:
		break
	}

	closure := TypeSymPrefix(".eqfunc", t).Linksym()
	if len(closure.P) > 0 { // already generated
		return closure
	}
	sym := TypeSymPrefix(".eq", t)
	if base.Flag.LowerR != 0 {
		fmt.Printf("geneq %v\n", t)
	}

	// Autogenerate code for equality of structs and arrays.

	base.Pos = base.AutogeneratedPos // less confusing than end of input
	typecheck.DeclContext = ir.PEXTERN

	// func sym(p, q *T) bool
	tfn := ir.NewFuncType(base.Pos, nil,
		[]*ir.Field{ir.NewField(base.Pos, typecheck.Lookup("p"), nil, types.NewPtr(t)), ir.NewField(base.Pos, typecheck.Lookup("q"), nil, types.NewPtr(t))},
		[]*ir.Field{ir.NewField(base.Pos, typecheck.Lookup("r"), nil, types.Types[types.TBOOL])})

	fn := typecheck.DeclFunc(sym, tfn)
	np := ir.AsNode(tfn.Type().Params().Field(0).Nname)
	nq := ir.AsNode(tfn.Type().Params().Field(1).Nname)
	nr := ir.AsNode(tfn.Type().Results().Field(0).Nname)

	// Label to jump to if an equality test fails.
	neq := typecheck.AutoLabel(".neq")

	// We reach here only for types that have equality but
	// cannot be handled by the standard algorithms,
	// so t must be either an array or a struct.
	switch t.Kind() {
	default:
		base.Fatalf("geneq %v", t)

	case types.TARRAY:
		nelem := t.NumElem()

		// checkAll generates code to check the equality of all array elements.
		// If unroll is greater than nelem, checkAll generates:
		//
		// if eq(p[0], q[0]) && eq(p[1], q[1]) && ... {
		// } else {
		//   return
		// }
		//
		// And so on.
		//
		// Otherwise it generates:
		//
		// for i := 0; i < nelem; i++ {
		//   if eq(p[i], q[i]) {
		//   } else {
		//     goto neq
		//   }
		// }
		//
		// TODO(josharian): consider doing some loop unrolling
		// for larger nelem as well, processing a few elements at a time in a loop.
		checkAll := func(unroll int64, last bool, eq func(pi, qi ir.Node) ir.Node) {
			// checkIdx generates a node to check for equality at index i.
			checkIdx := func(i ir.Node) ir.Node {
				// pi := p[i]
				pi := ir.NewIndexExpr(base.Pos, np, i)
				pi.SetBounded(true)
				pi.SetType(t.Elem())
				// qi := q[i]
				qi := ir.NewIndexExpr(base.Pos, nq, i)
				qi.SetBounded(true)
				qi.SetType(t.Elem())
				return eq(pi, qi)
			}

			if nelem <= unroll {
				if last {
					// Do last comparison in a different manner.
					nelem--
				}
				// Generate a series of checks.
				for i := int64(0); i < nelem; i++ {
					// if check {} else { goto neq }
					nif := ir.NewIfStmt(base.Pos, checkIdx(ir.NewInt(i)), nil, nil)
					nif.Else.Append(ir.NewBranchStmt(base.Pos, ir.OGOTO, neq))
					fn.Body.Append(nif)
				}
				if last {
					fn.Body.Append(ir.NewAssignStmt(base.Pos, nr, checkIdx(ir.NewInt(nelem))))
				}
			} else {
				// Generate a for loop.
				// for i := 0; i < nelem; i++
				i := typecheck.Temp(types.Types[types.TINT])
				init := ir.NewAssignStmt(base.Pos, i, ir.NewInt(0))
				cond := ir.NewBinaryExpr(base.Pos, ir.OLT, i, ir.NewInt(nelem))
				post := ir.NewAssignStmt(base.Pos, i, ir.NewBinaryExpr(base.Pos, ir.OADD, i, ir.NewInt(1)))
				loop := ir.NewForStmt(base.Pos, nil, cond, post, nil)
				loop.PtrInit().Append(init)
				// if eq(pi, qi) {} else { goto neq }
				nif := ir.NewIfStmt(base.Pos, checkIdx(i), nil, nil)
				nif.Else.Append(ir.NewBranchStmt(base.Pos, ir.OGOTO, neq))
				loop.Body.Append(nif)
				fn.Body.Append(loop)
				if last {
					fn.Body.Append(ir.NewAssignStmt(base.Pos, nr, ir.NewBool(true)))
				}
			}
		}

		switch t.Elem().Kind() {
		case types.TSTRING:
			// Do two loops. First, check that all the lengths match (cheap).
			// Second, check that all the contents match (expensive).
			// TODO: when the array size is small, unroll the length match checks.
			checkAll(3, false, func(pi, qi ir.Node) ir.Node {
				// Compare lengths.
				eqlen, _ := EqString(pi, qi)
				return eqlen
			})
			checkAll(1, true, func(pi, qi ir.Node) ir.Node {
				// Compare contents.
				_, eqmem := EqString(pi, qi)
				return eqmem
			})
		case types.TFLOAT32, types.TFLOAT64:
			checkAll(2, true, func(pi, qi ir.Node) ir.Node {
				// p[i] == q[i]
				return ir.NewBinaryExpr(base.Pos, ir.OEQ, pi, qi)
			})
		// TODO: pick apart structs, do them piecemeal too
		default:
			checkAll(1, true, func(pi, qi ir.Node) ir.Node {
				// p[i] == q[i]
				return ir.NewBinaryExpr(base.Pos, ir.OEQ, pi, qi)
			})
		}

	case types.TSTRUCT:
		// Build a list of conditions to satisfy.
		// The conditions are a list-of-lists. Conditions are reorderable
		// within each inner list. The outer lists must be evaluated in order.
		var conds [][]ir.Node
		conds = append(conds, []ir.Node{})
		and := func(n ir.Node) {
			i := len(conds) - 1
			conds[i] = append(conds[i], n)
		}

		// Walk the struct using memequal for runs of AMEM
		// and calling specific equality tests for the others.
		for i, fields := 0, t.FieldSlice(); i < len(fields); {
			f := fields[i]

			// Skip blank-named fields.
			if f.Sym.IsBlank() {
				i++
				continue
			}

			// Compare non-memory fields with field equality.
			if !isRegularMemory(f.Type) {
				if eqCanPanic(f.Type) {
					// Enforce ordering by starting a new set of reorderable conditions.
					conds = append(conds, []ir.Node{})
				}
				p := ir.NewSelectorExpr(base.Pos, ir.OXDOT, np, f.Sym)
				q := ir.NewSelectorExpr(base.Pos, ir.OXDOT, nq, f.Sym)
				switch {
				case f.Type.IsString():
					eqlen, eqmem := EqString(p, q)
					and(eqlen)
					and(eqmem)
				default:
					and(ir.NewBinaryExpr(base.Pos, ir.OEQ, p, q))
				}
				if eqCanPanic(f.Type) {
					// Also enforce ordering after something that can panic.
					conds = append(conds, []ir.Node{})
				}
				i++
				continue
			}

			// Find maximal length run of memory-only fields.
			size, next := memrun(t, i)

			// TODO(rsc): All the calls to newname are wrong for
			// cross-package unexported fields.
			if s := fields[i:next]; len(s) <= 2 {
				// Two or fewer fields: use plain field equality.
				for _, f := range s {
					and(eqfield(np, nq, f.Sym))
				}
			} else {
				// More than two fields: use memequal.
				and(eqmem(np, nq, f.Sym, size))
			}
			i = next
		}

		// Sort conditions to put runtime calls last.
		// Preserve the rest of the ordering.
		var flatConds []ir.Node
		for _, c := range conds {
			isCall := func(n ir.Node) bool {
				return n.Op() == ir.OCALL || n.Op() == ir.OCALLFUNC
			}
			sort.SliceStable(c, func(i, j int) bool {
				return !isCall(c[i]) && isCall(c[j])
			})
			flatConds = append(flatConds, c...)
		}

		if len(flatConds) == 0 {
			fn.Body.Append(ir.NewAssignStmt(base.Pos, nr, ir.NewBool(true)))
		} else {
			for _, c := range flatConds[:len(flatConds)-1] {
				// if cond {} else { goto neq }
				n := ir.NewIfStmt(base.Pos, c, nil, nil)
				n.Else.Append(ir.NewBranchStmt(base.Pos, ir.OGOTO, neq))
				fn.Body.Append(n)
			}
			fn.Body.Append(ir.NewAssignStmt(base.Pos, nr, flatConds[len(flatConds)-1]))
		}
	}

	// ret:
	//   return
	ret := typecheck.AutoLabel(".ret")
	fn.Body.Append(ir.NewLabelStmt(base.Pos, ret))
	fn.Body.Append(ir.NewReturnStmt(base.Pos, nil))

	// neq:
	//   r = false
	//   return (or goto ret)
	fn.Body.Append(ir.NewLabelStmt(base.Pos, neq))
	fn.Body.Append(ir.NewAssignStmt(base.Pos, nr, ir.NewBool(false)))
	if eqCanPanic(t) || anyCall(fn) {
		// Epilogue is large, so share it with the equal case.
		fn.Body.Append(ir.NewBranchStmt(base.Pos, ir.OGOTO, ret))
	} else {
		// Epilogue is small, so don't bother sharing.
		fn.Body.Append(ir.NewReturnStmt(base.Pos, nil))
	}
	// TODO(khr): the epilogue size detection condition above isn't perfect.
	// We should really do a generic CL that shares epilogues across
	// the board. See #24936.

	if base.Flag.LowerR != 0 {
		ir.DumpList("geneq body", fn.Body)
	}

	typecheck.FinishFuncBody()

	fn.SetDupok(true)
	typecheck.Func(fn)

	ir.CurFunc = fn
	typecheck.Stmts(fn.Body)
	ir.CurFunc = nil

	if base.Debug.DclStack != 0 {
		types.CheckDclstack()
	}

	// Disable checknils while compiling this code.
	// We are comparing a struct or an array,
	// neither of which can be nil, and our comparisons
	// are shallow.
	fn.SetNilCheckDisabled(true)
	typecheck.Target.Decls = append(typecheck.Target.Decls, fn)

	// Generate a closure which points at the function we just generated.
	objw.SymPtr(closure, 0, sym.Linksym(), 0)
	objw.Global(closure, int32(types.PtrSize), obj.DUPOK|obj.RODATA)
	return closure
}

func anyCall(fn *ir.Func) bool {
	return ir.Any(fn, func(n ir.Node) bool {
		// TODO(rsc): No methods?
		op := n.Op()
		return op == ir.OCALL || op == ir.OCALLFUNC
	})
}

// eqfield returns the node
// 	p.field == q.field
func eqfield(p ir.Node, q ir.Node, field *types.Sym) ir.Node {
	nx := ir.NewSelectorExpr(base.Pos, ir.OXDOT, p, field)
	ny := ir.NewSelectorExpr(base.Pos, ir.OXDOT, q, field)
	ne := ir.NewBinaryExpr(base.Pos, ir.OEQ, nx, ny)
	return ne
}

// EqString returns the nodes
//   len(s) == len(t)
// and
//   memequal(s.ptr, t.ptr, len(s))
// which can be used to construct string equality comparison.
// eqlen must be evaluated before eqmem, and shortcircuiting is required.
func EqString(s, t ir.Node) (eqlen *ir.BinaryExpr, eqmem *ir.CallExpr) {
	s = typecheck.Conv(s, types.Types[types.TSTRING])
	t = typecheck.Conv(t, types.Types[types.TSTRING])
	sptr := ir.NewUnaryExpr(base.Pos, ir.OSPTR, s)
	tptr := ir.NewUnaryExpr(base.Pos, ir.OSPTR, t)
	slen := typecheck.Conv(ir.NewUnaryExpr(base.Pos, ir.OLEN, s), types.Types[types.TUINTPTR])
	tlen := typecheck.Conv(ir.NewUnaryExpr(base.Pos, ir.OLEN, t), types.Types[types.TUINTPTR])

	fn := typecheck.LookupRuntime("memequal")
	fn = typecheck.SubstArgTypes(fn, types.Types[types.TUINT8], types.Types[types.TUINT8])
	call := ir.NewCallExpr(base.Pos, ir.OCALL, fn, []ir.Node{sptr, tptr, ir.Copy(slen)})
	typecheck.Call(call)

	cmp := ir.NewBinaryExpr(base.Pos, ir.OEQ, slen, tlen)
	cmp = typecheck.Expr(cmp).(*ir.BinaryExpr)
	cmp.SetType(types.Types[types.TBOOL])
	return cmp, call
}

// EqInterface returns the nodes
//   s.tab == t.tab (or s.typ == t.typ, as appropriate)
// and
//   ifaceeq(s.tab, s.data, t.data) (or efaceeq(s.typ, s.data, t.data), as appropriate)
// which can be used to construct interface equality comparison.
// eqtab must be evaluated before eqdata, and shortcircuiting is required.
func EqInterface(s, t ir.Node) (eqtab *ir.BinaryExpr, eqdata *ir.CallExpr) {
	if !types.Identical(s.Type(), t.Type()) {
		base.Fatalf("eqinterface %v %v", s.Type(), t.Type())
	}
	// func ifaceeq(tab *uintptr, x, y unsafe.Pointer) (ret bool)
	// func efaceeq(typ *uintptr, x, y unsafe.Pointer) (ret bool)
	var fn ir.Node
	if s.Type().IsEmptyInterface() {
		fn = typecheck.LookupRuntime("efaceeq")
	} else {
		fn = typecheck.LookupRuntime("ifaceeq")
	}

	stab := ir.NewUnaryExpr(base.Pos, ir.OITAB, s)
	ttab := ir.NewUnaryExpr(base.Pos, ir.OITAB, t)
	sdata := ir.NewUnaryExpr(base.Pos, ir.OIDATA, s)
	tdata := ir.NewUnaryExpr(base.Pos, ir.OIDATA, t)
	sdata.SetType(types.Types[types.TUNSAFEPTR])
	tdata.SetType(types.Types[types.TUNSAFEPTR])
	sdata.SetTypecheck(1)
	tdata.SetTypecheck(1)

	call := ir.NewCallExpr(base.Pos, ir.OCALL, fn, []ir.Node{stab, sdata, tdata})
	typecheck.Call(call)

	cmp := ir.NewBinaryExpr(base.Pos, ir.OEQ, stab, ttab)
	cmp = typecheck.Expr(cmp).(*ir.BinaryExpr)
	cmp.SetType(types.Types[types.TBOOL])
	return cmp, call
}

// eqmem returns the node
// 	memequal(&p.field, &q.field [, size])
func eqmem(p ir.Node, q ir.Node, field *types.Sym, size int64) ir.Node {
	nx := typecheck.Expr(typecheck.NodAddr(ir.NewSelectorExpr(base.Pos, ir.OXDOT, p, field)))
	ny := typecheck.Expr(typecheck.NodAddr(ir.NewSelectorExpr(base.Pos, ir.OXDOT, q, field)))

	fn, needsize := eqmemfunc(size, nx.Type().Elem())
	call := ir.NewCallExpr(base.Pos, ir.OCALL, fn, nil)
	call.Args.Append(nx)
	call.Args.Append(ny)
	if needsize {
		call.Args.Append(ir.NewInt(size))
	}

	return call
}

func eqmemfunc(size int64, t *types.Type) (fn *ir.Name, needsize bool) {
	switch size {
	default:
		fn = typecheck.LookupRuntime("memequal")
		needsize = true
	case 1, 2, 4, 8, 16:
		buf := fmt.Sprintf("memequal%d", int(size)*8)
		fn = typecheck.LookupRuntime(buf)
	}

	fn = typecheck.SubstArgTypes(fn, t, t)
	return fn, needsize
}

// memrun finds runs of struct fields for which memory-only algs are appropriate.
// t is the parent struct type, and start is the field index at which to start the run.
// size is the length in bytes of the memory included in the run.
// next is the index just after the end of the memory run.
func memrun(t *types.Type, start int) (size int64, next int) {
	next = start
	for {
		next++
		if next == t.NumFields() {
			break
		}
		// Stop run after a padded field.
		if types.IsPaddedField(t, next-1) {
			break
		}
		// Also, stop before a blank or non-memory field.
		if f := t.Field(next); f.Sym.IsBlank() || !isRegularMemory(f.Type) {
			break
		}
	}
	return t.Field(next-1).End() - t.Field(start).Offset, next
}

func hashmem(t *types.Type) ir.Node {
	sym := ir.Pkgs.Runtime.Lookup("memhash")

	n := typecheck.NewName(sym)
	ir.MarkFunc(n)
	n.SetType(types.NewSignature(types.NoPkg, nil, []*types.Field{
		types.NewField(base.Pos, nil, types.NewPtr(t)),
		types.NewField(base.Pos, nil, types.Types[types.TUINTPTR]),
		types.NewField(base.Pos, nil, types.Types[types.TUINTPTR]),
	}, []*types.Field{
		types.NewField(base.Pos, nil, types.Types[types.TUINTPTR]),
	}))
	return n
}
