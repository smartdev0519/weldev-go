// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package walk

import (
	"cmd/compile/internal/base"
	"cmd/compile/internal/ir"
	"cmd/compile/internal/typecheck"
)

// The result of walkStmt MUST be assigned back to n, e.g.
// 	n.Left = walkStmt(n.Left)
func walkStmt(n ir.Node) ir.Node {
	if n == nil {
		return n
	}

	ir.SetPos(n)

	walkStmtList(n.Init())

	switch n.Op() {
	default:
		if n.Op() == ir.ONAME {
			n := n.(*ir.Name)
			base.Errorf("%v is not a top level statement", n.Sym())
		} else {
			base.Errorf("%v is not a top level statement", n.Op())
		}
		ir.Dump("nottop", n)
		return n

	case ir.OAS,
		ir.OASOP,
		ir.OAS2,
		ir.OAS2DOTTYPE,
		ir.OAS2RECV,
		ir.OAS2FUNC,
		ir.OAS2MAPR,
		ir.OCLOSE,
		ir.OCOPY,
		ir.OCALLMETH,
		ir.OCALLINTER,
		ir.OCALL,
		ir.OCALLFUNC,
		ir.ODELETE,
		ir.OSEND,
		ir.OPRINT,
		ir.OPRINTN,
		ir.OPANIC,
		ir.ORECOVER,
		ir.OGETG:
		if n.Typecheck() == 0 {
			base.Fatalf("missing typecheck: %+v", n)
		}
		init := ir.TakeInit(n)
		n = walkExpr(n, &init)
		if n.Op() == ir.ONAME {
			// copy rewrote to a statement list and a temp for the length.
			// Throw away the temp to avoid plain values as statements.
			n = ir.NewBlockStmt(n.Pos(), init)
			init.Set(nil)
		}
		if len(init) > 0 {
			switch n.Op() {
			case ir.OAS, ir.OAS2, ir.OBLOCK:
				n.(ir.InitNode).PtrInit().Prepend(init...)

			default:
				init.Append(n)
				n = ir.NewBlockStmt(n.Pos(), init)
			}
		}
		return n

	// special case for a receive where we throw away
	// the value received.
	case ir.ORECV:
		n := n.(*ir.UnaryExpr)
		return walkRecv(n)

	case ir.OBREAK,
		ir.OCONTINUE,
		ir.OFALL,
		ir.OGOTO,
		ir.OLABEL,
		ir.ODCLCONST,
		ir.ODCLTYPE,
		ir.OCHECKNIL,
		ir.OVARDEF,
		ir.OVARKILL,
		ir.OVARLIVE:
		return n

	case ir.ODCL:
		n := n.(*ir.Decl)
		return walkDecl(n)

	case ir.OBLOCK:
		n := n.(*ir.BlockStmt)
		walkStmtList(n.List)
		return n

	case ir.OCASE:
		base.Errorf("case statement out of place")
		panic("unreachable")

	case ir.ODEFER:
		n := n.(*ir.GoDeferStmt)
		ir.CurFunc.SetHasDefer(true)
		ir.CurFunc.NumDefers++
		if ir.CurFunc.NumDefers > maxOpenDefers {
			// Don't allow open-coded defers if there are more than
			// 8 defers in the function, since we use a single
			// byte to record active defers.
			ir.CurFunc.SetOpenCodedDeferDisallowed(true)
		}
		if n.Esc() != ir.EscNever {
			// If n.Esc is not EscNever, then this defer occurs in a loop,
			// so open-coded defers cannot be used in this function.
			ir.CurFunc.SetOpenCodedDeferDisallowed(true)
		}
		fallthrough
	case ir.OGO:
		n := n.(*ir.GoDeferStmt)
		return walkGoDefer(n)

	case ir.OFOR, ir.OFORUNTIL:
		n := n.(*ir.ForStmt)
		return walkFor(n)

	case ir.OIF:
		n := n.(*ir.IfStmt)
		return walkIf(n)

	case ir.ORETURN:
		n := n.(*ir.ReturnStmt)
		return walkReturn(n)

	case ir.ORETJMP:
		n := n.(*ir.BranchStmt)
		return n

	case ir.OINLMARK:
		n := n.(*ir.InlineMarkStmt)
		return n

	case ir.OSELECT:
		n := n.(*ir.SelectStmt)
		walkSelect(n)
		return n

	case ir.OSWITCH:
		n := n.(*ir.SwitchStmt)
		walkSwitch(n)
		return n

	case ir.ORANGE:
		n := n.(*ir.RangeStmt)
		return walkRange(n)
	}

	// No return! Each case must return (or panic),
	// to avoid confusion about what gets returned
	// in the presence of type assertions.
}

func walkStmtList(s []ir.Node) {
	for i := range s {
		s[i] = walkStmt(s[i])
	}
}

// walkDecl walks an ODCL node.
func walkDecl(n *ir.Decl) ir.Node {
	v := n.X
	if v.Class_ == ir.PAUTOHEAP {
		if base.Flag.CompilingRuntime {
			base.Errorf("%v escapes to heap, not allowed in runtime", v)
		}
		nn := ir.NewAssignStmt(base.Pos, v.Heapaddr, callnew(v.Type()))
		nn.Def = true
		return walkStmt(typecheck.Stmt(nn))
	}
	return n
}

// walkFor walks an OFOR or OFORUNTIL node.
func walkFor(n *ir.ForStmt) ir.Node {
	if n.Cond != nil {
		init := ir.TakeInit(n.Cond)
		walkStmtList(init)
		n.Cond = walkExpr(n.Cond, &init)
		n.Cond = ir.InitExpr(init, n.Cond)
	}

	n.Post = walkStmt(n.Post)
	if n.Op() == ir.OFORUNTIL {
		walkStmtList(n.Late)
	}
	walkStmtList(n.Body)
	return n
}

// walkGoDefer walks an OGO or ODEFER node.
func walkGoDefer(n *ir.GoDeferStmt) ir.Node {
	var init ir.Nodes
	switch call := n.Call; call.Op() {
	case ir.OPRINT, ir.OPRINTN:
		call := call.(*ir.CallExpr)
		n.Call = wrapCall(call, &init)

	case ir.ODELETE:
		call := call.(*ir.CallExpr)
		if mapfast(call.Args[0].Type()) == mapslow {
			n.Call = wrapCall(call, &init)
		} else {
			n.Call = walkExpr(call, &init)
		}

	case ir.OCOPY:
		call := call.(*ir.BinaryExpr)
		n.Call = walkCopy(call, &init, true)

	case ir.OCALLFUNC, ir.OCALLMETH, ir.OCALLINTER:
		call := call.(*ir.CallExpr)
		if len(call.KeepAlive) > 0 {
			n.Call = wrapCall(call, &init)
		} else {
			n.Call = walkExpr(call, &init)
		}

	default:
		n.Call = walkExpr(call, &init)
	}
	if len(init) > 0 {
		init.Append(n)
		return ir.NewBlockStmt(n.Pos(), init)
	}
	return n
}

// walkIf walks an OIF node.
func walkIf(n *ir.IfStmt) ir.Node {
	n.Cond = walkExpr(n.Cond, n.PtrInit())
	walkStmtList(n.Body)
	walkStmtList(n.Else)
	return n
}

// The result of wrapCall MUST be assigned back to n, e.g.
// 	n.Left = wrapCall(n.Left, init)
func wrapCall(n *ir.CallExpr, init *ir.Nodes) ir.Node {
	if len(n.Init()) != 0 {
		walkStmtList(n.Init())
		init.Append(ir.TakeInit(n)...)
	}

	isBuiltinCall := n.Op() != ir.OCALLFUNC && n.Op() != ir.OCALLMETH && n.Op() != ir.OCALLINTER

	// Turn f(a, b, []T{c, d, e}...) back into f(a, b, c, d, e).
	if !isBuiltinCall && n.IsDDD {
		last := len(n.Args) - 1
		if va := n.Args[last]; va.Op() == ir.OSLICELIT {
			va := va.(*ir.CompLitExpr)
			n.Args.Set(append(n.Args[:last], va.List...))
			n.IsDDD = false
		}
	}

	// origArgs keeps track of what argument is uintptr-unsafe/unsafe-uintptr conversion.
	origArgs := make([]ir.Node, len(n.Args))
	var funcArgs []*ir.Field
	for i, arg := range n.Args {
		s := typecheck.LookupNum("a", i)
		if !isBuiltinCall && arg.Op() == ir.OCONVNOP && arg.Type().IsUintptr() && arg.(*ir.ConvExpr).X.Type().IsUnsafePtr() {
			origArgs[i] = arg
			arg = arg.(*ir.ConvExpr).X
			n.Args[i] = arg
		}
		funcArgs = append(funcArgs, ir.NewField(base.Pos, s, nil, arg.Type()))
	}
	t := ir.NewFuncType(base.Pos, nil, funcArgs, nil)

	wrapCall_prgen++
	sym := typecheck.LookupNum("wrap·", wrapCall_prgen)
	fn := typecheck.DeclFunc(sym, t)

	args := ir.ParamNames(t.Type())
	for i, origArg := range origArgs {
		if origArg == nil {
			continue
		}
		args[i] = ir.NewConvExpr(base.Pos, origArg.Op(), origArg.Type(), args[i])
	}
	call := ir.NewCallExpr(base.Pos, n.Op(), n.X, args)
	if !isBuiltinCall {
		call.SetOp(ir.OCALL)
		call.IsDDD = n.IsDDD
	}
	fn.Body = []ir.Node{call}

	typecheck.FinishFuncBody()

	typecheck.Func(fn)
	typecheck.Stmts(fn.Body)
	typecheck.Target.Decls = append(typecheck.Target.Decls, fn)

	call = ir.NewCallExpr(base.Pos, ir.OCALL, fn.Nname, n.Args)
	return walkExpr(typecheck.Stmt(call), init)
}
