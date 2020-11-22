// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"cmd/compile/internal/base"
	"cmd/compile/internal/ir"
	"cmd/compile/internal/types"
	"cmd/internal/sys"
	"unicode/utf8"
)

// range
func typecheckrange(n *ir.Node) {
	// Typechecking order is important here:
	// 0. first typecheck range expression (slice/map/chan),
	//	it is evaluated only once and so logically it is not part of the loop.
	// 1. typecheck produced values,
	//	this part can declare new vars and so it must be typechecked before body,
	//	because body can contain a closure that captures the vars.
	// 2. decldepth++ to denote loop body.
	// 3. typecheck body.
	// 4. decldepth--.
	typecheckrangeExpr(n)

	// second half of dance, the first half being typecheckrangeExpr
	n.SetTypecheck(1)
	ls := n.List().Slice()
	for i1, n1 := range ls {
		if n1.Typecheck() == 0 {
			ls[i1] = typecheck(ls[i1], ctxExpr|ctxAssign)
		}
	}

	decldepth++
	typecheckslice(n.Body().Slice(), ctxStmt)
	decldepth--
}

func typecheckrangeExpr(n *ir.Node) {
	n.SetRight(typecheck(n.Right(), ctxExpr))

	t := n.Right().Type()
	if t == nil {
		return
	}
	// delicate little dance.  see typecheckas2
	ls := n.List().Slice()
	for i1, n1 := range ls {
		if n1.Name() == nil || n1.Name().Defn != n {
			ls[i1] = typecheck(ls[i1], ctxExpr|ctxAssign)
		}
	}

	if t.IsPtr() && t.Elem().IsArray() {
		t = t.Elem()
	}
	n.SetType(t)

	var t1, t2 *types.Type
	toomany := false
	switch t.Etype {
	default:
		base.ErrorfAt(n.Pos(), "cannot range over %L", n.Right())
		return

	case types.TARRAY, types.TSLICE:
		t1 = types.Types[types.TINT]
		t2 = t.Elem()

	case types.TMAP:
		t1 = t.Key()
		t2 = t.Elem()

	case types.TCHAN:
		if !t.ChanDir().CanRecv() {
			base.ErrorfAt(n.Pos(), "invalid operation: range %v (receive from send-only type %v)", n.Right(), n.Right().Type())
			return
		}

		t1 = t.Elem()
		t2 = nil
		if n.List().Len() == 2 {
			toomany = true
		}

	case types.TSTRING:
		t1 = types.Types[types.TINT]
		t2 = types.Runetype
	}

	if n.List().Len() > 2 || toomany {
		base.ErrorfAt(n.Pos(), "too many variables in range")
	}

	var v1, v2 *ir.Node
	if n.List().Len() != 0 {
		v1 = n.List().First()
	}
	if n.List().Len() > 1 {
		v2 = n.List().Second()
	}

	// this is not only an optimization but also a requirement in the spec.
	// "if the second iteration variable is the blank identifier, the range
	// clause is equivalent to the same clause with only the first variable
	// present."
	if ir.IsBlank(v2) {
		if v1 != nil {
			n.PtrList().Set1(v1)
		}
		v2 = nil
	}

	if v1 != nil {
		if v1.Name() != nil && v1.Name().Defn == n {
			v1.SetType(t1)
		} else if v1.Type() != nil {
			if op, why := assignop(t1, v1.Type()); op == ir.OXXX {
				base.ErrorfAt(n.Pos(), "cannot assign type %v to %L in range%s", t1, v1, why)
			}
		}
		checkassign(n, v1)
	}

	if v2 != nil {
		if v2.Name() != nil && v2.Name().Defn == n {
			v2.SetType(t2)
		} else if v2.Type() != nil {
			if op, why := assignop(t2, v2.Type()); op == ir.OXXX {
				base.ErrorfAt(n.Pos(), "cannot assign type %v to %L in range%s", t2, v2, why)
			}
		}
		checkassign(n, v2)
	}
}

func cheapComputableIndex(width int64) bool {
	switch thearch.LinkArch.Family {
	// MIPS does not have R+R addressing
	// Arm64 may lack ability to generate this code in our assembler,
	// but the architecture supports it.
	case sys.PPC64, sys.S390X:
		return width == 1
	case sys.AMD64, sys.I386, sys.ARM64, sys.ARM:
		switch width {
		case 1, 2, 4, 8:
			return true
		}
	}
	return false
}

// walkrange transforms various forms of ORANGE into
// simpler forms.  The result must be assigned back to n.
// Node n may also be modified in place, and may also be
// the returned node.
func walkrange(n *ir.Node) *ir.Node {
	if isMapClear(n) {
		m := n.Right()
		lno := setlineno(m)
		n = mapClear(m)
		base.Pos = lno
		return n
	}

	// variable name conventions:
	//	ohv1, hv1, hv2: hidden (old) val 1, 2
	//	ha, hit: hidden aggregate, iterator
	//	hn, hp: hidden len, pointer
	//	hb: hidden bool
	//	a, v1, v2: not hidden aggregate, val 1, 2

	t := n.Type()

	a := n.Right()
	lno := setlineno(a)
	n.SetRight(nil)

	var v1, v2 *ir.Node
	l := n.List().Len()
	if l > 0 {
		v1 = n.List().First()
	}

	if l > 1 {
		v2 = n.List().Second()
	}

	if ir.IsBlank(v2) {
		v2 = nil
	}

	if ir.IsBlank(v1) && v2 == nil {
		v1 = nil
	}

	if v1 == nil && v2 != nil {
		base.Fatalf("walkrange: v2 != nil while v1 == nil")
	}

	// n.List has no meaning anymore, clear it
	// to avoid erroneous processing by racewalk.
	n.PtrList().Set(nil)

	var ifGuard *ir.Node

	translatedLoopOp := ir.OFOR

	var body []*ir.Node
	var init []*ir.Node
	switch t.Etype {
	default:
		base.Fatalf("walkrange")

	case types.TARRAY, types.TSLICE:
		if arrayClear(n, v1, v2, a) {
			base.Pos = lno
			return n
		}

		// order.stmt arranged for a copy of the array/slice variable if needed.
		ha := a

		hv1 := temp(types.Types[types.TINT])
		hn := temp(types.Types[types.TINT])

		init = append(init, ir.Nod(ir.OAS, hv1, nil))
		init = append(init, ir.Nod(ir.OAS, hn, ir.Nod(ir.OLEN, ha, nil)))

		n.SetLeft(ir.Nod(ir.OLT, hv1, hn))
		n.SetRight(ir.Nod(ir.OAS, hv1, ir.Nod(ir.OADD, hv1, nodintconst(1))))

		// for range ha { body }
		if v1 == nil {
			break
		}

		// for v1 := range ha { body }
		if v2 == nil {
			body = []*ir.Node{ir.Nod(ir.OAS, v1, hv1)}
			break
		}

		// for v1, v2 := range ha { body }
		if cheapComputableIndex(n.Type().Elem().Width) {
			// v1, v2 = hv1, ha[hv1]
			tmp := ir.Nod(ir.OINDEX, ha, hv1)
			tmp.SetBounded(true)
			// Use OAS2 to correctly handle assignments
			// of the form "v1, a[v1] := range".
			a := ir.Nod(ir.OAS2, nil, nil)
			a.PtrList().Set2(v1, v2)
			a.PtrRlist().Set2(hv1, tmp)
			body = []*ir.Node{a}
			break
		}

		// TODO(austin): OFORUNTIL is a strange beast, but is
		// necessary for expressing the control flow we need
		// while also making "break" and "continue" work. It
		// would be nice to just lower ORANGE during SSA, but
		// racewalk needs to see many of the operations
		// involved in ORANGE's implementation. If racewalk
		// moves into SSA, consider moving ORANGE into SSA and
		// eliminating OFORUNTIL.

		// TODO(austin): OFORUNTIL inhibits bounds-check
		// elimination on the index variable (see #20711).
		// Enhance the prove pass to understand this.
		ifGuard = ir.Nod(ir.OIF, nil, nil)
		ifGuard.SetLeft(ir.Nod(ir.OLT, hv1, hn))
		translatedLoopOp = ir.OFORUNTIL

		hp := temp(types.NewPtr(n.Type().Elem()))
		tmp := ir.Nod(ir.OINDEX, ha, nodintconst(0))
		tmp.SetBounded(true)
		init = append(init, ir.Nod(ir.OAS, hp, ir.Nod(ir.OADDR, tmp, nil)))

		// Use OAS2 to correctly handle assignments
		// of the form "v1, a[v1] := range".
		a := ir.Nod(ir.OAS2, nil, nil)
		a.PtrList().Set2(v1, v2)
		a.PtrRlist().Set2(hv1, ir.Nod(ir.ODEREF, hp, nil))
		body = append(body, a)

		// Advance pointer as part of the late increment.
		//
		// This runs *after* the condition check, so we know
		// advancing the pointer is safe and won't go past the
		// end of the allocation.
		a = ir.Nod(ir.OAS, hp, addptr(hp, t.Elem().Width))
		a = typecheck(a, ctxStmt)
		n.PtrList().Set1(a)

	case types.TMAP:
		// order.stmt allocated the iterator for us.
		// we only use a once, so no copy needed.
		ha := a

		hit := prealloc[n]
		th := hit.Type()
		n.SetLeft(nil)
		keysym := th.Field(0).Sym  // depends on layout of iterator struct.  See reflect.go:hiter
		elemsym := th.Field(1).Sym // ditto

		fn := syslook("mapiterinit")

		fn = substArgTypes(fn, t.Key(), t.Elem(), th)
		init = append(init, mkcall1(fn, nil, nil, typename(t), ha, ir.Nod(ir.OADDR, hit, nil)))
		n.SetLeft(ir.Nod(ir.ONE, nodSym(ir.ODOT, hit, keysym), nodnil()))

		fn = syslook("mapiternext")
		fn = substArgTypes(fn, th)
		n.SetRight(mkcall1(fn, nil, nil, ir.Nod(ir.OADDR, hit, nil)))

		key := nodSym(ir.ODOT, hit, keysym)
		key = ir.Nod(ir.ODEREF, key, nil)
		if v1 == nil {
			body = nil
		} else if v2 == nil {
			body = []*ir.Node{ir.Nod(ir.OAS, v1, key)}
		} else {
			elem := nodSym(ir.ODOT, hit, elemsym)
			elem = ir.Nod(ir.ODEREF, elem, nil)
			a := ir.Nod(ir.OAS2, nil, nil)
			a.PtrList().Set2(v1, v2)
			a.PtrRlist().Set2(key, elem)
			body = []*ir.Node{a}
		}

	case types.TCHAN:
		// order.stmt arranged for a copy of the channel variable.
		ha := a

		n.SetLeft(nil)

		hv1 := temp(t.Elem())
		hv1.SetTypecheck(1)
		if t.Elem().HasPointers() {
			init = append(init, ir.Nod(ir.OAS, hv1, nil))
		}
		hb := temp(types.Types[types.TBOOL])

		n.SetLeft(ir.Nod(ir.ONE, hb, nodbool(false)))
		a := ir.Nod(ir.OAS2RECV, nil, nil)
		a.SetTypecheck(1)
		a.PtrList().Set2(hv1, hb)
		a.SetRight(ir.Nod(ir.ORECV, ha, nil))
		n.Left().PtrInit().Set1(a)
		if v1 == nil {
			body = nil
		} else {
			body = []*ir.Node{ir.Nod(ir.OAS, v1, hv1)}
		}
		// Zero hv1. This prevents hv1 from being the sole, inaccessible
		// reference to an otherwise GC-able value during the next channel receive.
		// See issue 15281.
		body = append(body, ir.Nod(ir.OAS, hv1, nil))

	case types.TSTRING:
		// Transform string range statements like "for v1, v2 = range a" into
		//
		// ha := a
		// for hv1 := 0; hv1 < len(ha); {
		//   hv1t := hv1
		//   hv2 := rune(ha[hv1])
		//   if hv2 < utf8.RuneSelf {
		//      hv1++
		//   } else {
		//      hv2, hv1 = decoderune(ha, hv1)
		//   }
		//   v1, v2 = hv1t, hv2
		//   // original body
		// }

		// order.stmt arranged for a copy of the string variable.
		ha := a

		hv1 := temp(types.Types[types.TINT])
		hv1t := temp(types.Types[types.TINT])
		hv2 := temp(types.Runetype)

		// hv1 := 0
		init = append(init, ir.Nod(ir.OAS, hv1, nil))

		// hv1 < len(ha)
		n.SetLeft(ir.Nod(ir.OLT, hv1, ir.Nod(ir.OLEN, ha, nil)))

		if v1 != nil {
			// hv1t = hv1
			body = append(body, ir.Nod(ir.OAS, hv1t, hv1))
		}

		// hv2 := rune(ha[hv1])
		nind := ir.Nod(ir.OINDEX, ha, hv1)
		nind.SetBounded(true)
		body = append(body, ir.Nod(ir.OAS, hv2, conv(nind, types.Runetype)))

		// if hv2 < utf8.RuneSelf
		nif := ir.Nod(ir.OIF, nil, nil)
		nif.SetLeft(ir.Nod(ir.OLT, hv2, nodintconst(utf8.RuneSelf)))

		// hv1++
		nif.PtrBody().Set1(ir.Nod(ir.OAS, hv1, ir.Nod(ir.OADD, hv1, nodintconst(1))))

		// } else {
		eif := ir.Nod(ir.OAS2, nil, nil)
		nif.PtrRlist().Set1(eif)

		// hv2, hv1 = decoderune(ha, hv1)
		eif.PtrList().Set2(hv2, hv1)
		fn := syslook("decoderune")
		eif.PtrRlist().Set1(mkcall1(fn, fn.Type().Results(), nil, ha, hv1))

		body = append(body, nif)

		if v1 != nil {
			if v2 != nil {
				// v1, v2 = hv1t, hv2
				a := ir.Nod(ir.OAS2, nil, nil)
				a.PtrList().Set2(v1, v2)
				a.PtrRlist().Set2(hv1t, hv2)
				body = append(body, a)
			} else {
				// v1 = hv1t
				body = append(body, ir.Nod(ir.OAS, v1, hv1t))
			}
		}
	}

	n.SetOp(translatedLoopOp)
	typecheckslice(init, ctxStmt)

	if ifGuard != nil {
		ifGuard.PtrInit().Append(init...)
		ifGuard = typecheck(ifGuard, ctxStmt)
	} else {
		n.PtrInit().Append(init...)
	}

	typecheckslice(n.Left().Init().Slice(), ctxStmt)

	n.SetLeft(typecheck(n.Left(), ctxExpr))
	n.SetLeft(defaultlit(n.Left(), nil))
	n.SetRight(typecheck(n.Right(), ctxStmt))
	typecheckslice(body, ctxStmt)
	n.PtrBody().Prepend(body...)

	if ifGuard != nil {
		ifGuard.PtrBody().Set1(n)
		n = ifGuard
	}

	n = walkstmt(n)

	base.Pos = lno
	return n
}

// isMapClear checks if n is of the form:
//
// for k := range m {
//   delete(m, k)
// }
//
// where == for keys of map m is reflexive.
func isMapClear(n *ir.Node) bool {
	if base.Flag.N != 0 || instrumenting {
		return false
	}

	if n.Op() != ir.ORANGE || n.Type().Etype != types.TMAP || n.List().Len() != 1 {
		return false
	}

	k := n.List().First()
	if k == nil || ir.IsBlank(k) {
		return false
	}

	// Require k to be a new variable name.
	if k.Name() == nil || k.Name().Defn != n {
		return false
	}

	if n.Body().Len() != 1 {
		return false
	}

	stmt := n.Body().First() // only stmt in body
	if stmt == nil || stmt.Op() != ir.ODELETE {
		return false
	}

	m := n.Right()
	if !samesafeexpr(stmt.List().First(), m) || !samesafeexpr(stmt.List().Second(), k) {
		return false
	}

	// Keys where equality is not reflexive can not be deleted from maps.
	if !isreflexive(m.Type().Key()) {
		return false
	}

	return true
}

// mapClear constructs a call to runtime.mapclear for the map m.
func mapClear(m *ir.Node) *ir.Node {
	t := m.Type()

	// instantiate mapclear(typ *type, hmap map[any]any)
	fn := syslook("mapclear")
	fn = substArgTypes(fn, t.Key(), t.Elem())
	n := mkcall1(fn, nil, nil, typename(t), m)

	n = typecheck(n, ctxStmt)
	n = walkstmt(n)

	return n
}

// Lower n into runtime·memclr if possible, for
// fast zeroing of slices and arrays (issue 5373).
// Look for instances of
//
// for i := range a {
// 	a[i] = zero
// }
//
// in which the evaluation of a is side-effect-free.
//
// Parameters are as in walkrange: "for v1, v2 = range a".
func arrayClear(n, v1, v2, a *ir.Node) bool {
	if base.Flag.N != 0 || instrumenting {
		return false
	}

	if v1 == nil || v2 != nil {
		return false
	}

	if n.Body().Len() != 1 || n.Body().First() == nil {
		return false
	}

	stmt := n.Body().First() // only stmt in body
	if stmt.Op() != ir.OAS || stmt.Left().Op() != ir.OINDEX {
		return false
	}

	if !samesafeexpr(stmt.Left().Left(), a) || !samesafeexpr(stmt.Left().Right(), v1) {
		return false
	}

	elemsize := n.Type().Elem().Width
	if elemsize <= 0 || !isZero(stmt.Right()) {
		return false
	}

	// Convert to
	// if len(a) != 0 {
	// 	hp = &a[0]
	// 	hn = len(a)*sizeof(elem(a))
	// 	memclr{NoHeap,Has}Pointers(hp, hn)
	// 	i = len(a) - 1
	// }
	n.SetOp(ir.OIF)

	n.PtrBody().Set(nil)
	n.SetLeft(ir.Nod(ir.ONE, ir.Nod(ir.OLEN, a, nil), nodintconst(0)))

	// hp = &a[0]
	hp := temp(types.Types[types.TUNSAFEPTR])

	tmp := ir.Nod(ir.OINDEX, a, nodintconst(0))
	tmp.SetBounded(true)
	tmp = ir.Nod(ir.OADDR, tmp, nil)
	tmp = convnop(tmp, types.Types[types.TUNSAFEPTR])
	n.PtrBody().Append(ir.Nod(ir.OAS, hp, tmp))

	// hn = len(a) * sizeof(elem(a))
	hn := temp(types.Types[types.TUINTPTR])

	tmp = ir.Nod(ir.OLEN, a, nil)
	tmp = ir.Nod(ir.OMUL, tmp, nodintconst(elemsize))
	tmp = conv(tmp, types.Types[types.TUINTPTR])
	n.PtrBody().Append(ir.Nod(ir.OAS, hn, tmp))

	var fn *ir.Node
	if a.Type().Elem().HasPointers() {
		// memclrHasPointers(hp, hn)
		Curfn.Func().SetWBPos(stmt.Pos())
		fn = mkcall("memclrHasPointers", nil, nil, hp, hn)
	} else {
		// memclrNoHeapPointers(hp, hn)
		fn = mkcall("memclrNoHeapPointers", nil, nil, hp, hn)
	}

	n.PtrBody().Append(fn)

	// i = len(a) - 1
	v1 = ir.Nod(ir.OAS, v1, ir.Nod(ir.OSUB, ir.Nod(ir.OLEN, a, nil), nodintconst(1)))

	n.PtrBody().Append(v1)

	n.SetLeft(typecheck(n.Left(), ctxExpr))
	n.SetLeft(defaultlit(n.Left(), nil))
	typecheckslice(n.Body().Slice(), ctxStmt)
	n = walkstmt(n)
	return true
}

// addptr returns (*T)(uintptr(p) + n).
func addptr(p *ir.Node, n int64) *ir.Node {
	t := p.Type()

	p = ir.Nod(ir.OCONVNOP, p, nil)
	p.SetType(types.Types[types.TUINTPTR])

	p = ir.Nod(ir.OADD, p, nodintconst(n))

	p = ir.Nod(ir.OCONVNOP, p, nil)
	p.SetType(t)

	return p
}
