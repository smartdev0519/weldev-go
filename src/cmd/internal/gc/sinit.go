// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"cmd/internal/obj"
	"fmt"
)

/*
 * static initialization
 */
const (
	InitNotStarted = 0
	InitDone       = 1
	InitPending    = 2
)

var initlist *NodeList

// init1 walks the AST starting at n, and accumulates in out
// the list of definitions needing init code in dependency order.
func init1(n *Node, out **NodeList) {
	var l *NodeList
	var nv *Node

	if n == nil {
		return
	}
	init1(n.Left, out)
	init1(n.Right, out)
	for l = n.List; l != nil; l = l.Next {
		init1(l.N, out)
	}

	if n.Left != nil && n.Type != nil && n.Left.Op == OTYPE && n.Class == PFUNC {
		// Methods called as Type.Method(receiver, ...).
		// Definitions for method expressions are stored in type->nname.
		init1(n.Type.Nname, out)
	}

	if n.Op != ONAME {
		return
	}
	switch n.Class {
	case PEXTERN,
		PFUNC:
		break

	default:
		if isblank(n) && n.Curfn == nil && n.Defn != nil && n.Defn.Initorder == InitNotStarted {
			// blank names initialization is part of init() but not
			// when they are inside a function.
			break
		}

		return
	}

	if n.Initorder == InitDone {
		return
	}
	if n.Initorder == InitPending {
		// Since mutually recursive sets of functions are allowed,
		// we don't necessarily raise an error if n depends on a node
		// which is already waiting for its dependencies to be visited.
		//
		// initlist contains a cycle of identifiers referring to each other.
		// If this cycle contains a variable, then this variable refers to itself.
		// Conversely, if there exists an initialization cycle involving
		// a variable in the program, the tree walk will reach a cycle
		// involving that variable.
		if n.Class != PFUNC {
			nv = n
			goto foundinitloop
		}

		for l = initlist; l.N != n; l = l.Next {
			if l.N.Class != PFUNC {
				nv = l.N
				goto foundinitloop
			}
		}

		// The loop involves only functions, ok.
		return

		// if there have already been errors printed,
		// those errors probably confused us and
		// there might not be a loop.  let the user
		// fix those first.
	foundinitloop:
		Flusherrors()

		if nerrors > 0 {
			errorexit()
		}

		// There is a loop involving nv. We know about
		// n and initlist = n1 <- ... <- nv <- ... <- n <- ...
		fmt.Printf("%v: initialization loop:\n", nv.Line())

		// Build back pointers in initlist.
		for l = initlist; l != nil; l = l.Next {
			if l.Next != nil {
				l.Next.End = l
			}
		}

		// Print nv -> ... -> n1 -> n.
		for l = initlist; l.N != nv; l = l.Next {
		}
		for ; l != nil; l = l.End {
			fmt.Printf("\t%v %v refers to\n", l.N.Line(), Sconv(l.N.Sym, 0))
		}

		// Print n -> ... -> nv.
		for l = initlist; l.N != n; l = l.Next {
		}
		for ; l.N != nv; l = l.End {
			fmt.Printf("\t%v %v refers to\n", l.N.Line(), Sconv(l.N.Sym, 0))
		}
		fmt.Printf("\t%v %v\n", nv.Line(), Sconv(nv.Sym, 0))
		errorexit()
	}

	// reached a new unvisited node.
	n.Initorder = InitPending

	l = new(NodeList)
	if l == nil {
		Flusherrors()
		Yyerror("out of memory")
		errorexit()
	}

	l.Next = initlist
	l.N = n
	l.End = nil
	initlist = l

	// make sure that everything n depends on is initialized.
	// n->defn is an assignment to n
	if n.Defn != nil {
		switch n.Defn.Op {
		default:
			goto bad

		case ODCLFUNC:
			init2list(n.Defn.Nbody, out)

		case OAS:
			if n.Defn.Left != n {
				goto bad
			}
			if isblank(n.Defn.Left) && candiscard(n.Defn.Right) {
				n.Defn.Op = OEMPTY
				n.Defn.Left = nil
				n.Defn.Right = nil
				break
			}

			init2(n.Defn.Right, out)
			if Debug['j'] != 0 {
				fmt.Printf("%v\n", Sconv(n.Sym, 0))
			}
			if isblank(n) || !staticinit(n, out) {
				if Debug['%'] != 0 {
					Dump("nonstatic", n.Defn)
				}
				*out = list(*out, n.Defn)
			}

		case OAS2FUNC,
			OAS2MAPR,
			OAS2DOTTYPE,
			OAS2RECV:
			if n.Defn.Initorder != InitNotStarted {
				break
			}
			n.Defn.Initorder = InitDone
			for l = n.Defn.Rlist; l != nil; l = l.Next {
				init1(l.N, out)
			}
			if Debug['%'] != 0 {
				Dump("nonstatic", n.Defn)
			}
			*out = list(*out, n.Defn)
		}
	}

	l = initlist
	initlist = l.Next
	if l.N != n {
		Fatal("bad initlist")
	}

	n.Initorder = InitDone
	return

bad:
	Dump("defn", n.Defn)
	Fatal("init1: bad defn")
}

// recurse over n, doing init1 everywhere.
func init2(n *Node, out **NodeList) {
	if n == nil || n.Initorder == InitDone {
		return
	}

	if n.Op == ONAME && n.Ninit != nil {
		Fatal("name %v with ninit: %v\n", Sconv(n.Sym, 0), Nconv(n, obj.FmtSign))
	}

	init1(n, out)
	init2(n.Left, out)
	init2(n.Right, out)
	init2(n.Ntest, out)
	init2list(n.Ninit, out)
	init2list(n.List, out)
	init2list(n.Rlist, out)
	init2list(n.Nbody, out)
	init2list(n.Nelse, out)

	if n.Op == OCLOSURE {
		init2list(n.Closure.Nbody, out)
	}
	if n.Op == ODOTMETH || n.Op == OCALLPART {
		init2(n.Type.Nname, out)
	}
}

func init2list(l *NodeList, out **NodeList) {
	for ; l != nil; l = l.Next {
		init2(l.N, out)
	}
}

func initreorder(l *NodeList, out **NodeList) {
	var n *Node

	for ; l != nil; l = l.Next {
		n = l.N
		switch n.Op {
		case ODCLFUNC,
			ODCLCONST,
			ODCLTYPE:
			continue
		}

		initreorder(n.Ninit, out)
		n.Ninit = nil
		init1(n, out)
	}
}

// initfix computes initialization order for a list l of top-level
// declarations and outputs the corresponding list of statements
// to include in the init() function body.
func initfix(l *NodeList) *NodeList {
	var lout *NodeList
	var lno int

	lout = nil
	lno = int(lineno)
	initreorder(l, &lout)
	lineno = int32(lno)
	return lout
}

/*
 * compilation of top-level (static) assignments
 * into DATA statements if at all possible.
 */
func staticinit(n *Node, out **NodeList) bool {
	var l *Node
	var r *Node

	if n.Op != ONAME || n.Class != PEXTERN || n.Defn == nil || n.Defn.Op != OAS {
		Fatal("staticinit")
	}

	lineno = n.Lineno
	l = n.Defn.Left
	r = n.Defn.Right
	return staticassign(l, r, out)
}

// like staticassign but we are copying an already
// initialized value r.
func staticcopy(l *Node, r *Node, out **NodeList) bool {
	var i int
	var e *InitEntry
	var p *InitPlan
	var a *Node
	var ll *Node
	var rr *Node
	var orig *Node
	var n1 Node

	if r.Op != ONAME || r.Class != PEXTERN || r.Sym.Pkg != localpkg {
		return false
	}
	if r.Defn == nil { // probably zeroed but perhaps supplied externally and of unknown value
		return false
	}
	if r.Defn.Op != OAS {
		return false
	}
	orig = r
	r = r.Defn.Right

	switch r.Op {
	case ONAME:
		if staticcopy(l, r, out) {
			return true
		}
		*out = list(*out, Nod(OAS, l, r))
		return true

	case OLITERAL:
		if iszero(r) {
			return true
		}
		gdata(l, r, int(l.Type.Width))
		return true

	case OADDR:
		switch r.Left.Op {
		case ONAME:
			gdata(l, r, int(l.Type.Width))
			return true
		}

	case OPTRLIT:
		switch r.Left.Op {
		//dump("not static addr", r);
		default:
			break

			// copy pointer
		case OARRAYLIT,
			OSTRUCTLIT,
			OMAPLIT:
			gdata(l, Nod(OADDR, r.Nname, nil), int(l.Type.Width))

			return true
		}

	case OARRAYLIT:
		if Isslice(r.Type) {
			// copy slice
			a = r.Nname

			n1 = *l
			n1.Xoffset = l.Xoffset + int64(Array_array)
			gdata(&n1, Nod(OADDR, a, nil), Widthptr)
			n1.Xoffset = l.Xoffset + int64(Array_nel)
			gdata(&n1, r.Right, Widthint)
			n1.Xoffset = l.Xoffset + int64(Array_cap)
			gdata(&n1, r.Right, Widthint)
			return true
		}
		fallthrough

		// fall through
	case OSTRUCTLIT:
		p = r.Initplan

		n1 = *l
		for i = 0; i < len(p.E); i++ {
			e = &p.E[i]
			n1.Xoffset = l.Xoffset + e.Xoffset
			n1.Type = e.Expr.Type
			if e.Expr.Op == OLITERAL {
				gdata(&n1, e.Expr, int(n1.Type.Width))
			} else {
				ll = Nod(OXXX, nil, nil)
				*ll = n1
				ll.Orig = ll // completely separate copy
				if !staticassign(ll, e.Expr, out) {
					// Requires computation, but we're
					// copying someone else's computation.
					rr = Nod(OXXX, nil, nil)

					*rr = *orig
					rr.Orig = rr // completely separate copy
					rr.Type = ll.Type
					rr.Xoffset += e.Xoffset
					*out = list(*out, Nod(OAS, ll, rr))
				}
			}
		}

		return true
	}

	return false
}

func staticassign(l *Node, r *Node, out **NodeList) bool {
	var a *Node
	var n1 Node
	var nam Node
	var ta *Type
	var p *InitPlan
	var e *InitEntry
	var i int
	var sval *Strlit

	switch r.Op {
	//dump("not static", r);
	default:
		break

	case ONAME:
		if r.Class == PEXTERN && r.Sym.Pkg == localpkg {
			return staticcopy(l, r, out)
		}

	case OLITERAL:
		if iszero(r) {
			return true
		}
		gdata(l, r, int(l.Type.Width))
		return true

	case OADDR:
		if stataddr(&nam, r.Left) {
			n1 = *r
			n1.Left = &nam
			gdata(l, &n1, int(l.Type.Width))
			return true
		}
		fallthrough

	case OPTRLIT:
		switch r.Left.Op {
		//dump("not static ptrlit", r);
		default:
			break

			// Init pointer.
		case OARRAYLIT,
			OMAPLIT,
			OSTRUCTLIT:
			a = staticname(r.Left.Type, 1)

			r.Nname = a
			gdata(l, Nod(OADDR, a, nil), int(l.Type.Width))

			// Init underlying literal.
			if !staticassign(a, r.Left, out) {
				*out = list(*out, Nod(OAS, a, r.Left))
			}
			return true
		}

	case OSTRARRAYBYTE:
		if l.Class == PEXTERN && r.Left.Op == OLITERAL {
			sval = r.Left.Val.U.Sval
			slicebytes(l, sval.S, len(sval.S))
			return true
		}

	case OARRAYLIT:
		initplan(r)
		if Isslice(r.Type) {
			// Init slice.
			ta = typ(TARRAY)

			ta.Type = r.Type.Type
			ta.Bound = Mpgetfix(r.Right.Val.U.Xval)
			a = staticname(ta, 1)
			r.Nname = a
			n1 = *l
			n1.Xoffset = l.Xoffset + int64(Array_array)
			gdata(&n1, Nod(OADDR, a, nil), Widthptr)
			n1.Xoffset = l.Xoffset + int64(Array_nel)
			gdata(&n1, r.Right, Widthint)
			n1.Xoffset = l.Xoffset + int64(Array_cap)
			gdata(&n1, r.Right, Widthint)

			// Fall through to init underlying array.
			l = a
		}
		fallthrough

		// fall through
	case OSTRUCTLIT:
		initplan(r)

		p = r.Initplan
		n1 = *l
		for i = 0; i < len(p.E); i++ {
			e = &p.E[i]
			n1.Xoffset = l.Xoffset + e.Xoffset
			n1.Type = e.Expr.Type
			if e.Expr.Op == OLITERAL {
				gdata(&n1, e.Expr, int(n1.Type.Width))
			} else {
				a = Nod(OXXX, nil, nil)
				*a = n1
				a.Orig = a // completely separate copy
				if !staticassign(a, e.Expr, out) {
					*out = list(*out, Nod(OAS, a, e.Expr))
				}
			}
		}

		return true

		// TODO: Table-driven map insert.
	case OMAPLIT:
		break
	}

	return false
}

/*
 * from here down is the walk analysis
 * of composite literals.
 * most of the work is to generate
 * data statements for the constant
 * part of the composite literal.
 */
func staticname(t *Type, ctxt int) *Node {
	var n *Node

	namebuf = fmt.Sprintf("statictmp_%.4d", statuniqgen)
	statuniqgen++
	n = newname(Lookup(namebuf))
	if ctxt == 0 {
		n.Readonly = 1
	}
	addvar(n, t, PEXTERN)
	return n
}

func isliteral(n *Node) bool {
	if n.Op == OLITERAL {
		if n.Val.Ctype != CTNIL {
			return true
		}
	}
	return false
}

func simplename(n *Node) bool {
	if n.Op != ONAME {
		goto no
	}
	if n.Addable == 0 {
		goto no
	}
	if n.Class&PHEAP != 0 {
		goto no
	}
	if n.Class == PPARAMREF {
		goto no
	}
	return true

no:
	return false
}

func litas(l *Node, r *Node, init **NodeList) {
	var a *Node

	a = Nod(OAS, l, r)
	typecheck(&a, Etop)
	walkexpr(&a, init)
	*init = list(*init, a)
}

const (
	MODEDYNAM = 1
	MODECONST = 2
)

func getdyn(n *Node, top int) int {
	var nl *NodeList
	var value *Node
	var mode int

	mode = 0
	switch n.Op {
	default:
		if isliteral(n) {
			return MODECONST
		}
		return MODEDYNAM

	case OARRAYLIT:
		if top == 0 && n.Type.Bound < 0 {
			return MODEDYNAM
		}
		fallthrough

	case OSTRUCTLIT:
		break
	}

	for nl = n.List; nl != nil; nl = nl.Next {
		value = nl.N.Right
		mode |= getdyn(value, 0)
		if mode == MODEDYNAM|MODECONST {
			break
		}
	}

	return mode
}

func structlit(ctxt int, pass int, n *Node, var_ *Node, init **NodeList) {
	var r *Node
	var a *Node
	var nl *NodeList
	var index *Node
	var value *Node

	for nl = n.List; nl != nil; nl = nl.Next {
		r = nl.N
		if r.Op != OKEY {
			Fatal("structlit: rhs not OKEY: %v", Nconv(r, 0))
		}
		index = r.Left
		value = r.Right

		switch value.Op {
		case OARRAYLIT:
			if value.Type.Bound < 0 {
				if pass == 1 && ctxt != 0 {
					a = Nod(ODOT, var_, newname(index.Sym))
					slicelit(ctxt, value, a, init)
				} else if pass == 2 && ctxt == 0 {
					a = Nod(ODOT, var_, newname(index.Sym))
					slicelit(ctxt, value, a, init)
				} else if pass == 3 {
					break
				}
				continue
			}

			a = Nod(ODOT, var_, newname(index.Sym))
			arraylit(ctxt, pass, value, a, init)
			continue

		case OSTRUCTLIT:
			a = Nod(ODOT, var_, newname(index.Sym))
			structlit(ctxt, pass, value, a, init)
			continue
		}

		if isliteral(value) {
			if pass == 2 {
				continue
			}
		} else if pass == 1 {
			continue
		}

		// build list of var.field = expr
		a = Nod(ODOT, var_, newname(index.Sym))

		a = Nod(OAS, a, value)
		typecheck(&a, Etop)
		if pass == 1 {
			walkexpr(&a, init) // add any assignments in r to top
			if a.Op != OAS {
				Fatal("structlit: not as")
			}
			a.Dodata = 2
		} else {
			orderstmtinplace(&a)
			walkstmt(&a)
		}

		*init = list(*init, a)
	}
}

func arraylit(ctxt int, pass int, n *Node, var_ *Node, init **NodeList) {
	var r *Node
	var a *Node
	var l *NodeList
	var index *Node
	var value *Node

	for l = n.List; l != nil; l = l.Next {
		r = l.N
		if r.Op != OKEY {
			Fatal("arraylit: rhs not OKEY: %v", Nconv(r, 0))
		}
		index = r.Left
		value = r.Right

		switch value.Op {
		case OARRAYLIT:
			if value.Type.Bound < 0 {
				if pass == 1 && ctxt != 0 {
					a = Nod(OINDEX, var_, index)
					slicelit(ctxt, value, a, init)
				} else if pass == 2 && ctxt == 0 {
					a = Nod(OINDEX, var_, index)
					slicelit(ctxt, value, a, init)
				} else if pass == 3 {
					break
				}
				continue
			}

			a = Nod(OINDEX, var_, index)
			arraylit(ctxt, pass, value, a, init)
			continue

		case OSTRUCTLIT:
			a = Nod(OINDEX, var_, index)
			structlit(ctxt, pass, value, a, init)
			continue
		}

		if isliteral(index) && isliteral(value) {
			if pass == 2 {
				continue
			}
		} else if pass == 1 {
			continue
		}

		// build list of var[index] = value
		a = Nod(OINDEX, var_, index)

		a = Nod(OAS, a, value)
		typecheck(&a, Etop)
		if pass == 1 {
			walkexpr(&a, init)
			if a.Op != OAS {
				Fatal("arraylit: not as")
			}
			a.Dodata = 2
		} else {
			orderstmtinplace(&a)
			walkstmt(&a)
		}

		*init = list(*init, a)
	}
}

func slicelit(ctxt int, n *Node, var_ *Node, init **NodeList) {
	var r *Node
	var a *Node
	var l *NodeList
	var t *Type
	var vstat *Node
	var vauto *Node
	var index *Node
	var value *Node
	var mode int

	// make an array type
	t = shallow(n.Type)

	t.Bound = Mpgetfix(n.Right.Val.U.Xval)
	t.Width = 0
	t.Sym = nil
	t.Haspointers = 0
	dowidth(t)

	if ctxt != 0 {
		// put everything into static array
		vstat = staticname(t, ctxt)

		arraylit(ctxt, 1, n, vstat, init)
		arraylit(ctxt, 2, n, vstat, init)

		// copy static to slice
		a = Nod(OSLICE, vstat, Nod(OKEY, nil, nil))

		a = Nod(OAS, var_, a)
		typecheck(&a, Etop)
		a.Dodata = 2
		*init = list(*init, a)
		return
	}

	// recipe for var = []t{...}
	// 1. make a static array
	//	var vstat [...]t
	// 2. assign (data statements) the constant part
	//	vstat = constpart{}
	// 3. make an auto pointer to array and allocate heap to it
	//	var vauto *[...]t = new([...]t)
	// 4. copy the static array to the auto array
	//	*vauto = vstat
	// 5. assign slice of allocated heap to var
	//	var = [0:]*auto
	// 6. for each dynamic part assign to the slice
	//	var[i] = dynamic part
	//
	// an optimization is done if there is no constant part
	//	3. var vauto *[...]t = new([...]t)
	//	5. var = [0:]*auto
	//	6. var[i] = dynamic part

	// if the literal contains constants,
	// make static initialized array (1),(2)
	vstat = nil

	mode = getdyn(n, 1)
	if mode&MODECONST != 0 {
		vstat = staticname(t, ctxt)
		arraylit(ctxt, 1, n, vstat, init)
	}

	// make new auto *array (3 declare)
	vauto = temp(Ptrto(t))

	// set auto to point at new temp or heap (3 assign)
	if n.Alloc != nil {
		// temp allocated during order.c for dddarg
		n.Alloc.Type = t

		if vstat == nil {
			a = Nod(OAS, n.Alloc, nil)
			typecheck(&a, Etop)
			*init = list(*init, a) // zero new temp
		}

		a = Nod(OADDR, n.Alloc, nil)
	} else if n.Esc == EscNone {
		a = temp(t)
		if vstat == nil {
			a = Nod(OAS, temp(t), nil)
			typecheck(&a, Etop)
			*init = list(*init, a) // zero new temp
			a = a.Left
		}

		a = Nod(OADDR, a, nil)
	} else {
		a = Nod(ONEW, nil, nil)
		a.List = list1(typenod(t))
	}

	a = Nod(OAS, vauto, a)
	typecheck(&a, Etop)
	walkexpr(&a, init)
	*init = list(*init, a)

	if vstat != nil {
		// copy static to heap (4)
		a = Nod(OIND, vauto, nil)

		a = Nod(OAS, a, vstat)
		typecheck(&a, Etop)
		walkexpr(&a, init)
		*init = list(*init, a)
	}

	// make slice out of heap (5)
	a = Nod(OAS, var_, Nod(OSLICE, vauto, Nod(OKEY, nil, nil)))

	typecheck(&a, Etop)
	orderstmtinplace(&a)
	walkstmt(&a)
	*init = list(*init, a)

	// put dynamics into slice (6)
	for l = n.List; l != nil; l = l.Next {
		r = l.N
		if r.Op != OKEY {
			Fatal("slicelit: rhs not OKEY: %v", Nconv(r, 0))
		}
		index = r.Left
		value = r.Right
		a = Nod(OINDEX, var_, index)
		a.Bounded = true

		// TODO need to check bounds?

		switch value.Op {
		case OARRAYLIT:
			if value.Type.Bound < 0 {
				break
			}
			arraylit(ctxt, 2, value, a, init)
			continue

		case OSTRUCTLIT:
			structlit(ctxt, 2, value, a, init)
			continue
		}

		if isliteral(index) && isliteral(value) {
			continue
		}

		// build list of var[c] = expr
		a = Nod(OAS, a, value)

		typecheck(&a, Etop)
		orderstmtinplace(&a)
		walkstmt(&a)
		*init = list(*init, a)
	}
}

func maplit(ctxt int, n *Node, var_ *Node, init **NodeList) {
	var r *Node
	var a *Node
	var l *NodeList
	var nerr int
	var b int64
	var t *Type
	var tk *Type
	var tv *Type
	var t1 *Type
	var vstat *Node
	var index *Node
	var value *Node
	var key *Node
	var val *Node
	var syma *Sym
	var symb *Sym

	ctxt = 0

	// make the map var
	nerr = nerrors

	a = Nod(OMAKE, nil, nil)
	a.List = list1(typenod(n.Type))
	litas(var_, a, init)

	// count the initializers
	b = 0

	for l = n.List; l != nil; l = l.Next {
		r = l.N

		if r.Op != OKEY {
			Fatal("maplit: rhs not OKEY: %v", Nconv(r, 0))
		}
		index = r.Left
		value = r.Right

		if isliteral(index) && isliteral(value) {
			b++
		}
	}

	if b != 0 {
		// build type [count]struct { a Tindex, b Tvalue }
		t = n.Type

		tk = t.Down
		tv = t.Type

		symb = Lookup("b")
		t = typ(TFIELD)
		t.Type = tv
		t.Sym = symb

		syma = Lookup("a")
		t1 = t
		t = typ(TFIELD)
		t.Type = tk
		t.Sym = syma
		t.Down = t1

		t1 = t
		t = typ(TSTRUCT)
		t.Type = t1

		t1 = t
		t = typ(TARRAY)
		t.Bound = b
		t.Type = t1

		dowidth(t)

		// make and initialize static array
		vstat = staticname(t, ctxt)

		b = 0
		for l = n.List; l != nil; l = l.Next {
			r = l.N

			if r.Op != OKEY {
				Fatal("maplit: rhs not OKEY: %v", Nconv(r, 0))
			}
			index = r.Left
			value = r.Right

			if isliteral(index) && isliteral(value) {
				// build vstat[b].a = key;
				a = Nodintconst(b)

				a = Nod(OINDEX, vstat, a)
				a = Nod(ODOT, a, newname(syma))
				a = Nod(OAS, a, index)
				typecheck(&a, Etop)
				walkexpr(&a, init)
				a.Dodata = 2
				*init = list(*init, a)

				// build vstat[b].b = value;
				a = Nodintconst(b)

				a = Nod(OINDEX, vstat, a)
				a = Nod(ODOT, a, newname(symb))
				a = Nod(OAS, a, value)
				typecheck(&a, Etop)
				walkexpr(&a, init)
				a.Dodata = 2
				*init = list(*init, a)

				b++
			}
		}

		// loop adding structure elements to map
		// for i = 0; i < len(vstat); i++ {
		//	map[vstat[i].a] = vstat[i].b
		// }
		index = temp(Types[TINT])

		a = Nod(OINDEX, vstat, index)
		a.Bounded = true
		a = Nod(ODOT, a, newname(symb))

		r = Nod(OINDEX, vstat, index)
		r.Bounded = true
		r = Nod(ODOT, r, newname(syma))
		r = Nod(OINDEX, var_, r)

		r = Nod(OAS, r, a)

		a = Nod(OFOR, nil, nil)
		a.Nbody = list1(r)

		a.Ninit = list1(Nod(OAS, index, Nodintconst(0)))
		a.Ntest = Nod(OLT, index, Nodintconst(t.Bound))
		a.Nincr = Nod(OAS, index, Nod(OADD, index, Nodintconst(1)))

		typecheck(&a, Etop)
		walkstmt(&a)
		*init = list(*init, a)
	}

	// put in dynamic entries one-at-a-time
	key = nil

	val = nil
	for l = n.List; l != nil; l = l.Next {
		r = l.N

		if r.Op != OKEY {
			Fatal("maplit: rhs not OKEY: %v", Nconv(r, 0))
		}
		index = r.Left
		value = r.Right

		if isliteral(index) && isliteral(value) {
			continue
		}

		// build list of var[c] = expr.
		// use temporary so that mapassign1 can have addressable key, val.
		if key == nil {
			key = temp(var_.Type.Down)
			val = temp(var_.Type.Type)
		}

		a = Nod(OAS, key, r.Left)
		typecheck(&a, Etop)
		walkstmt(&a)
		*init = list(*init, a)
		a = Nod(OAS, val, r.Right)
		typecheck(&a, Etop)
		walkstmt(&a)
		*init = list(*init, a)

		a = Nod(OAS, Nod(OINDEX, var_, key), val)
		typecheck(&a, Etop)
		walkstmt(&a)
		*init = list(*init, a)

		if nerr != nerrors {
			break
		}
	}

	if key != nil {
		a = Nod(OVARKILL, key, nil)
		typecheck(&a, Etop)
		*init = list(*init, a)
		a = Nod(OVARKILL, val, nil)
		typecheck(&a, Etop)
		*init = list(*init, a)
	}
}

func anylit(ctxt int, n *Node, var_ *Node, init **NodeList) {
	var t *Type
	var a *Node
	var vstat *Node
	var r *Node

	t = n.Type
	switch n.Op {
	default:
		Fatal("anylit: not lit")

	case OPTRLIT:
		if Isptr[t.Etype] == 0 {
			Fatal("anylit: not ptr")
		}

		if n.Right != nil {
			r = Nod(OADDR, n.Right, nil)
			typecheck(&r, Erv)
		} else {
			r = Nod(ONEW, nil, nil)
			r.Typecheck = 1
			r.Type = t
			r.Esc = n.Esc
		}

		walkexpr(&r, init)
		a = Nod(OAS, var_, r)

		typecheck(&a, Etop)
		*init = list(*init, a)

		var_ = Nod(OIND, var_, nil)
		typecheck(&var_, Erv|Easgn)
		anylit(ctxt, n.Left, var_, init)

	case OSTRUCTLIT:
		if t.Etype != TSTRUCT {
			Fatal("anylit: not struct")
		}

		if simplename(var_) && count(n.List) > 4 {
			if ctxt == 0 {
				// lay out static data
				vstat = staticname(t, ctxt)

				structlit(ctxt, 1, n, vstat, init)

				// copy static to var
				a = Nod(OAS, var_, vstat)

				typecheck(&a, Etop)
				walkexpr(&a, init)
				*init = list(*init, a)

				// add expressions to automatic
				structlit(ctxt, 2, n, var_, init)

				break
			}

			structlit(ctxt, 1, n, var_, init)
			structlit(ctxt, 2, n, var_, init)
			break
		}

		// initialize of not completely specified
		if simplename(var_) || count(n.List) < structcount(t) {
			a = Nod(OAS, var_, nil)
			typecheck(&a, Etop)
			walkexpr(&a, init)
			*init = list(*init, a)
		}

		structlit(ctxt, 3, n, var_, init)

	case OARRAYLIT:
		if t.Etype != TARRAY {
			Fatal("anylit: not array")
		}
		if t.Bound < 0 {
			slicelit(ctxt, n, var_, init)
			break
		}

		if simplename(var_) && count(n.List) > 4 {
			if ctxt == 0 {
				// lay out static data
				vstat = staticname(t, ctxt)

				arraylit(1, 1, n, vstat, init)

				// copy static to automatic
				a = Nod(OAS, var_, vstat)

				typecheck(&a, Etop)
				walkexpr(&a, init)
				*init = list(*init, a)

				// add expressions to automatic
				arraylit(ctxt, 2, n, var_, init)

				break
			}

			arraylit(ctxt, 1, n, var_, init)
			arraylit(ctxt, 2, n, var_, init)
			break
		}

		// initialize of not completely specified
		if simplename(var_) || int64(count(n.List)) < t.Bound {
			a = Nod(OAS, var_, nil)
			typecheck(&a, Etop)
			walkexpr(&a, init)
			*init = list(*init, a)
		}

		arraylit(ctxt, 3, n, var_, init)

	case OMAPLIT:
		if t.Etype != TMAP {
			Fatal("anylit: not map")
		}
		maplit(ctxt, n, var_, init)
	}
}

func oaslit(n *Node, init **NodeList) bool {
	var ctxt int

	if n.Left == nil || n.Right == nil {
		goto no
	}
	if n.Left.Type == nil || n.Right.Type == nil {
		goto no
	}
	if !simplename(n.Left) {
		goto no
	}
	if !Eqtype(n.Left.Type, n.Right.Type) {
		goto no
	}

	// context is init() function.
	// implies generated data executed
	// exactly once and not subject to races.
	ctxt = 0

	//	if(n->dodata == 1)
	//		ctxt = 1;

	switch n.Right.Op {
	default:
		goto no

	case OSTRUCTLIT,
		OARRAYLIT,
		OMAPLIT:
		if vmatch1(n.Left, n.Right) {
			goto no
		}
		anylit(ctxt, n.Right, n.Left, init)
	}

	n.Op = OEMPTY
	return true

	// not a special composit literal assignment
no:
	return false
}

func getlit(lit *Node) int {
	if Smallintconst(lit) {
		return int(Mpgetfix(lit.Val.U.Xval))
	}
	return -1
}

func stataddr(nam *Node, n *Node) bool {
	var l int

	if n == nil {
		goto no
	}

	switch n.Op {
	case ONAME:
		*nam = *n
		return n.Addable != 0

	case ODOT:
		if !stataddr(nam, n.Left) {
			break
		}
		nam.Xoffset += n.Xoffset
		nam.Type = n.Type
		return true

	case OINDEX:
		if n.Left.Type.Bound < 0 {
			break
		}
		if !stataddr(nam, n.Left) {
			break
		}
		l = getlit(n.Right)
		if l < 0 {
			break
		}

		// Check for overflow.
		if n.Type.Width != 0 && Thearch.MAXWIDTH/n.Type.Width <= int64(l) {
			break
		}
		nam.Xoffset += int64(l) * n.Type.Width
		nam.Type = n.Type
		return true
	}

no:
	return false
}

func initplan(n *Node) {
	var p *InitPlan
	var a *Node
	var l *NodeList

	if n.Initplan != nil {
		return
	}
	p = new(InitPlan)
	n.Initplan = p
	switch n.Op {
	default:
		Fatal("initplan")

	case OARRAYLIT:
		for l = n.List; l != nil; l = l.Next {
			a = l.N
			if a.Op != OKEY || !Smallintconst(a.Left) {
				Fatal("initplan arraylit")
			}
			addvalue(p, n.Type.Type.Width*Mpgetfix(a.Left.Val.U.Xval), nil, a.Right)
		}

	case OSTRUCTLIT:
		for l = n.List; l != nil; l = l.Next {
			a = l.N
			if a.Op != OKEY || a.Left.Type == nil {
				Fatal("initplan structlit")
			}
			addvalue(p, a.Left.Type.Width, nil, a.Right)
		}

	case OMAPLIT:
		for l = n.List; l != nil; l = l.Next {
			a = l.N
			if a.Op != OKEY {
				Fatal("initplan maplit")
			}
			addvalue(p, -1, a.Left, a.Right)
		}
	}
}

func addvalue(p *InitPlan, xoffset int64, key *Node, n *Node) {
	var i int
	var q *InitPlan
	var e *InitEntry

	// special case: zero can be dropped entirely
	if iszero(n) {
		p.Zero += n.Type.Width
		return
	}

	// special case: inline struct and array (not slice) literals
	if isvaluelit(n) {
		initplan(n)
		q = n.Initplan
		for i = 0; i < len(q.E); i++ {
			e = entry(p)
			*e = q.E[i]
			e.Xoffset += xoffset
		}

		return
	}

	// add to plan
	if n.Op == OLITERAL {
		p.Lit += n.Type.Width
	} else {
		p.Expr += n.Type.Width
	}

	e = entry(p)
	e.Xoffset = xoffset
	e.Expr = n
}

func iszero(n *Node) bool {
	var l *NodeList

	switch n.Op {
	case OLITERAL:
		switch n.Val.Ctype {
		default:
			Dump("unexpected literal", n)
			Fatal("iszero")

		case CTNIL:
			return true

		case CTSTR:
			return n.Val.U.Sval == nil || len(n.Val.U.Sval.S) == 0

		case CTBOOL:
			return n.Val.U.Bval == 0

		case CTINT,
			CTRUNE:
			return mpcmpfixc(n.Val.U.Xval, 0) == 0

		case CTFLT:
			return mpcmpfltc(n.Val.U.Fval, 0) == 0

		case CTCPLX:
			return mpcmpfltc(&n.Val.U.Cval.Real, 0) == 0 && mpcmpfltc(&n.Val.U.Cval.Imag, 0) == 0
		}

	case OARRAYLIT:
		if Isslice(n.Type) {
			break
		}
		fallthrough

		// fall through
	case OSTRUCTLIT:
		for l = n.List; l != nil; l = l.Next {
			if !iszero(l.N.Right) {
				return false
			}
		}
		return true
	}

	return false
}

func isvaluelit(n *Node) bool {
	return (n.Op == OARRAYLIT && Isfixedarray(n.Type)) || n.Op == OSTRUCTLIT
}

func entry(p *InitPlan) *InitEntry {
	p.E = append(p.E, InitEntry{})
	return &p.E[len(p.E)-1]
}

func gen_as_init(n *Node) bool {
	var nr *Node
	var nl *Node
	var nam Node
	var nod1 Node

	if n.Dodata == 0 {
		goto no
	}

	nr = n.Right
	nl = n.Left
	if nr == nil {
		if !stataddr(&nam, nl) {
			goto no
		}
		if nam.Class != PEXTERN {
			goto no
		}
		goto yes
	}

	if nr.Type == nil || !Eqtype(nl.Type, nr.Type) {
		goto no
	}

	if !stataddr(&nam, nl) {
		goto no
	}

	if nam.Class != PEXTERN {
		goto no
	}

	switch nr.Op {
	default:
		goto no

	case OCONVNOP:
		nr = nr.Left
		if nr == nil || nr.Op != OSLICEARR {
			goto no
		}
		fallthrough

		// fall through
	case OSLICEARR:
		if nr.Right.Op == OKEY && nr.Right.Left == nil && nr.Right.Right == nil {
			nr = nr.Left
			goto slice
		}

		goto no

	case OLITERAL:
		break
	}

	switch nr.Type.Etype {
	default:
		goto no

	case TBOOL,
		TINT8,
		TUINT8,
		TINT16,
		TUINT16,
		TINT32,
		TUINT32,
		TINT64,
		TUINT64,
		TINT,
		TUINT,
		TUINTPTR,
		TPTR32,
		TPTR64,
		TFLOAT32,
		TFLOAT64:
		gdata(&nam, nr, int(nr.Type.Width))

	case TCOMPLEX64,
		TCOMPLEX128:
		gdatacomplex(&nam, nr.Val.U.Cval)

	case TSTRING:
		gdatastring(&nam, nr.Val.U.Sval)
	}

yes:
	return true

slice:
	gused(nil) // in case the data is the dest of a goto
	nl = nr
	if nr == nil || nr.Op != OADDR {
		goto no
	}
	nr = nr.Left
	if nr == nil || nr.Op != ONAME {
		goto no
	}

	// nr is the array being converted to a slice
	if nr.Type == nil || nr.Type.Etype != TARRAY || nr.Type.Bound < 0 {
		goto no
	}

	nam.Xoffset += int64(Array_array)
	gdata(&nam, nl, int(Types[Tptr].Width))

	nam.Xoffset += int64(Array_nel) - int64(Array_array)
	Nodconst(&nod1, Types[TINT], nr.Type.Bound)
	gdata(&nam, &nod1, Widthint)

	nam.Xoffset += int64(Array_cap) - int64(Array_nel)
	gdata(&nam, &nod1, Widthint)

	goto yes

no:
	if n.Dodata == 2 {
		Dump("\ngen_as_init", n)
		Fatal("gen_as_init couldnt make data statement")
	}

	return false
}
