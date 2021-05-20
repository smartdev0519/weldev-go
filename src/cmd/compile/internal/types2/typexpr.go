// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements type-checking of identifiers and type expressions.

package types2

import (
	"cmd/compile/internal/syntax"
	"fmt"
	"go/constant"
	"strings"
)

// Disabled by default, but enabled when running tests (via types_test.go).
var acceptMethodTypeParams bool

// ident type-checks identifier e and initializes x with the value or type of e.
// If an error occurred, x.mode is set to invalid.
// For the meaning of def, see Checker.definedType, below.
// If wantType is set, the identifier e is expected to denote a type.
//
func (check *Checker) ident(x *operand, e *syntax.Name, def *Named, wantType bool) {
	x.mode = invalid
	x.expr = e

	// Note that we cannot use check.lookup here because the returned scope
	// may be different from obj.Parent(). See also Scope.LookupParent doc.
	scope, obj := check.scope.LookupParent(e.Value, check.pos)
	if obj == nil {
		if e.Value == "_" {
			check.error(e, "cannot use _ as value or type")
		} else {
			if check.conf.CompilerErrorMessages {
				check.errorf(e, "undefined: %s", e.Value)
			} else {
				check.errorf(e, "undeclared name: %s", e.Value)
			}
		}
		return
	}
	check.recordUse(e, obj)

	// Type-check the object.
	// Only call Checker.objDecl if the object doesn't have a type yet
	// (in which case we must actually determine it) or the object is a
	// TypeName and we also want a type (in which case we might detect
	// a cycle which needs to be reported). Otherwise we can skip the
	// call and avoid a possible cycle error in favor of the more
	// informative "not a type/value" error that this function's caller
	// will issue (see issue #25790).
	typ := obj.Type()
	if _, gotType := obj.(*TypeName); typ == nil || gotType && wantType {
		check.objDecl(obj, def)
		typ = obj.Type() // type must have been assigned by Checker.objDecl
	}
	assert(typ != nil)

	// The object may have been dot-imported.
	// If so, mark the respective package as used.
	// (This code is only needed for dot-imports. Without them,
	// we only have to mark variables, see *Var case below).
	if pkgName := check.dotImportMap[dotImportKey{scope, obj}]; pkgName != nil {
		pkgName.used = true
	}

	switch obj := obj.(type) {
	case *PkgName:
		check.errorf(e, "use of package %s not in selector", obj.name)
		return

	case *Const:
		check.addDeclDep(obj)
		if typ == Typ[Invalid] {
			return
		}
		if obj == universeIota {
			if check.iota == nil {
				check.error(e, "cannot use iota outside constant declaration")
				return
			}
			x.val = check.iota
		} else {
			x.val = obj.val
		}
		assert(x.val != nil)
		x.mode = constant_

	case *TypeName:
		x.mode = typexpr

	case *Var:
		// It's ok to mark non-local variables, but ignore variables
		// from other packages to avoid potential race conditions with
		// dot-imported variables.
		if obj.pkg == check.pkg {
			obj.used = true
		}
		check.addDeclDep(obj)
		if typ == Typ[Invalid] {
			return
		}
		x.mode = variable

	case *Func:
		check.addDeclDep(obj)
		x.mode = value

	case *Builtin:
		x.id = obj.id
		x.mode = builtin

	case *Nil:
		x.mode = nilvalue

	default:
		unreachable()
	}

	x.typ = typ
}

// typ type-checks the type expression e and returns its type, or Typ[Invalid].
// The type must not be an (uninstantiated) generic type.
func (check *Checker) typ(e syntax.Expr) Type {
	return check.definedType(e, nil)
}

// varType type-checks the type expression e and returns its type, or Typ[Invalid].
// The type must not be an (uninstantiated) generic type and it must be ordinary
// (see ordinaryType).
func (check *Checker) varType(e syntax.Expr) Type {
	typ := check.definedType(e, nil)
	check.ordinaryType(syntax.StartPos(e), typ)
	return typ
}

// ordinaryType reports an error if typ is an interface type containing
// type lists or is (or embeds) the predeclared type comparable.
func (check *Checker) ordinaryType(pos syntax.Pos, typ Type) {
	// We don't want to call under() (via Interface) or complete interfaces while we
	// are in the middle of type-checking parameter declarations that might belong to
	// interface methods. Delay this check to the end of type-checking.
	check.later(func() {
		if t := asInterface(typ); t != nil {
			check.completeInterface(pos, t) // TODO(gri) is this the correct position?
			if t.allTypes != nil {
				check.softErrorf(pos, "interface contains type constraints (%s)", t.allTypes)
				return
			}
			if t.IsComparable() {
				check.softErrorf(pos, "interface is (or embeds) comparable")
			}
		}
	})
}

// anyType type-checks the type expression e and returns its type, or Typ[Invalid].
// The type may be generic or instantiated.
func (check *Checker) anyType(e syntax.Expr) Type {
	typ := check.typInternal(e, nil)
	assert(isTyped(typ))
	check.recordTypeAndValue(e, typexpr, typ, nil)
	return typ
}

// definedType is like typ but also accepts a type name def.
// If def != nil, e is the type specification for the defined type def, declared
// in a type declaration, and def.underlying will be set to the type of e before
// any components of e are type-checked.
//
func (check *Checker) definedType(e syntax.Expr, def *Named) Type {
	typ := check.typInternal(e, def)
	assert(isTyped(typ))
	if isGeneric(typ) {
		check.errorf(e, "cannot use generic type %s without instantiation", typ)
		typ = Typ[Invalid]
	}
	check.recordTypeAndValue(e, typexpr, typ, nil)
	return typ
}

// genericType is like typ but the type must be an (uninstantiated) generic type.
func (check *Checker) genericType(e syntax.Expr, reportErr bool) Type {
	typ := check.typInternal(e, nil)
	assert(isTyped(typ))
	if typ != Typ[Invalid] && !isGeneric(typ) {
		if reportErr {
			check.errorf(e, "%s is not a generic type", typ)
		}
		typ = Typ[Invalid]
	}
	// TODO(gri) what is the correct call below?
	check.recordTypeAndValue(e, typexpr, typ, nil)
	return typ
}

// isubst returns an x with identifiers substituted per the substitution map smap.
// isubst only handles the case of (valid) method receiver type expressions correctly.
func isubst(x syntax.Expr, smap map[*syntax.Name]*syntax.Name) syntax.Expr {
	switch n := x.(type) {
	case *syntax.Name:
		if alt := smap[n]; alt != nil {
			return alt
		}
	// case *syntax.StarExpr:
	// 	X := isubst(n.X, smap)
	// 	if X != n.X {
	// 		new := *n
	// 		new.X = X
	// 		return &new
	// 	}
	case *syntax.Operation:
		if n.Op == syntax.Mul && n.Y == nil {
			X := isubst(n.X, smap)
			if X != n.X {
				new := *n
				new.X = X
				return &new
			}
		}
	case *syntax.IndexExpr:
		Index := isubst(n.Index, smap)
		if Index != n.Index {
			new := *n
			new.Index = Index
			return &new
		}
	case *syntax.ListExpr:
		var elems []syntax.Expr
		for i, elem := range n.ElemList {
			new := isubst(elem, smap)
			if new != elem {
				if elems == nil {
					elems = make([]syntax.Expr, len(n.ElemList))
					copy(elems, n.ElemList)
				}
				elems[i] = new
			}
		}
		if elems != nil {
			new := *n
			new.ElemList = elems
			return &new
		}
	case *syntax.ParenExpr:
		return isubst(n.X, smap) // no need to keep parentheses
	default:
		// Other receiver type expressions are invalid.
		// It's fine to ignore those here as they will
		// be checked elsewhere.
	}
	return x
}

// funcType type-checks a function or method type.
func (check *Checker) funcType(sig *Signature, recvPar *syntax.Field, tparams []*syntax.Field, ftyp *syntax.FuncType) {
	check.openScope(ftyp, "function")
	check.scope.isFunc = true
	check.recordScope(ftyp, check.scope)
	sig.scope = check.scope
	defer check.closeScope()

	var recvTyp syntax.Expr // rewritten receiver type; valid if != nil
	if recvPar != nil {
		// collect generic receiver type parameters, if any
		// - a receiver type parameter is like any other type parameter, except that it is declared implicitly
		// - the receiver specification acts as local declaration for its type parameters, which may be blank
		_, rname, rparams := check.unpackRecv(recvPar.Type, true)
		if len(rparams) > 0 {
			// Blank identifiers don't get declared and regular type-checking of the instantiated
			// parameterized receiver type expression fails in Checker.collectParams of receiver.
			// Identify blank type parameters and substitute each with a unique new identifier named
			// "n_" (where n is the parameter index) and which cannot conflict with any user-defined
			// name.
			var smap map[*syntax.Name]*syntax.Name // substitution map from "_" to "!n" identifiers
			for i, p := range rparams {
				if p.Value == "_" {
					new := *p
					new.Value = fmt.Sprintf("%d_", i)
					rparams[i] = &new // use n_ identifier instead of _ so it can be looked up
					if smap == nil {
						smap = make(map[*syntax.Name]*syntax.Name)
					}
					smap[p] = &new
				}
			}
			if smap != nil {
				// blank identifiers were found => use rewritten receiver type
				recvTyp = isubst(recvPar.Type, smap)
			}
			// TODO(gri) rework declareTypeParams
			sig.rparams = nil
			for _, rparam := range rparams {
				sig.rparams = check.declareTypeParam(sig.rparams, rparam)
			}
			// determine receiver type to get its type parameters
			// and the respective type parameter bounds
			var recvTParams []*TypeName
			if rname != nil {
				// recv should be a Named type (otherwise an error is reported elsewhere)
				// Also: Don't report an error via genericType since it will be reported
				//       again when we type-check the signature.
				// TODO(gri) maybe the receiver should be marked as invalid instead?
				if recv := asNamed(check.genericType(rname, false)); recv != nil {
					recvTParams = recv.tparams
				}
			}
			// provide type parameter bounds
			// - only do this if we have the right number (otherwise an error is reported elsewhere)
			if len(sig.rparams) == len(recvTParams) {
				// We have a list of *TypeNames but we need a list of Types.
				list := make([]Type, len(sig.rparams))
				for i, t := range sig.rparams {
					list[i] = t.typ
				}
				smap := makeSubstMap(recvTParams, list)
				for i, tname := range sig.rparams {
					bound := recvTParams[i].typ.(*TypeParam).bound
					// bound is (possibly) parameterized in the context of the
					// receiver type declaration. Substitute parameters for the
					// current context.
					// TODO(gri) should we assume now that bounds always exist?
					//           (no bound == empty interface)
					if bound != nil {
						bound = check.subst(tname.pos, bound, smap)
						tname.typ.(*TypeParam).bound = bound
					}
				}
			}
		}
	}

	if tparams != nil {
		sig.tparams = check.collectTypeParams(tparams)
		// Always type-check method type parameters but complain if they are not enabled.
		// (A separate check is needed when type-checking interface method signatures because
		// they don't have a receiver specification.)
		if recvPar != nil && !acceptMethodTypeParams {
			check.error(ftyp, "methods cannot have type parameters")
		}
	}

	// Value (non-type) parameters' scope starts in the function body. Use a temporary scope for their
	// declarations and then squash that scope into the parent scope (and report any redeclarations at
	// that time).
	scope := NewScope(check.scope, nopos, nopos, "function body (temp. scope)")
	var recvList []*Var // TODO(gri) remove the need for making a list here
	if recvPar != nil {
		recvList, _ = check.collectParams(scope, []*syntax.Field{recvPar}, recvTyp, false) // use rewritten receiver type, if any
	}
	params, variadic := check.collectParams(scope, ftyp.ParamList, nil, true)
	results, _ := check.collectParams(scope, ftyp.ResultList, nil, false)
	scope.Squash(func(obj, alt Object) {
		var err error_
		err.errorf(obj, "%s redeclared in this block", obj.Name())
		err.recordAltDecl(alt)
		check.report(&err)
	})

	if recvPar != nil {
		// recv parameter list present (may be empty)
		// spec: "The receiver is specified via an extra parameter section preceding the
		// method name. That parameter section must declare a single parameter, the receiver."
		var recv *Var
		switch len(recvList) {
		case 0:
			// error reported by resolver
			recv = NewParam(nopos, nil, "", Typ[Invalid]) // ignore recv below
		default:
			// more than one receiver
			check.error(recvList[len(recvList)-1].Pos(), "method must have exactly one receiver")
			fallthrough // continue with first receiver
		case 1:
			recv = recvList[0]
		}

		// TODO(gri) We should delay rtyp expansion to when we actually need the
		//           receiver; thus all checks here should be delayed to later.
		rtyp, _ := deref(recv.typ)
		rtyp = expand(rtyp)

		// spec: "The receiver type must be of the form T or *T where T is a type name."
		// (ignore invalid types - error was reported before)
		if t := rtyp; t != Typ[Invalid] {
			var err string
			if T := asNamed(t); T != nil {
				// spec: "The type denoted by T is called the receiver base type; it must not
				// be a pointer or interface type and it must be declared in the same package
				// as the method."
				if T.obj.pkg != check.pkg {
					err = "type not defined in this package"
					if check.conf.CompilerErrorMessages {
						check.errorf(recv.pos, "cannot define new methods on non-local type %s", recv.typ)
						err = ""
					}
				} else {
					switch u := optype(T).(type) {
					case *Basic:
						// unsafe.Pointer is treated like a regular pointer
						if u.kind == UnsafePointer {
							err = "unsafe.Pointer"
						}
					case *Pointer, *Interface:
						err = "pointer or interface type"
					}
				}
			} else if T := asBasic(t); T != nil {
				err = "basic or unnamed type"
				if check.conf.CompilerErrorMessages {
					check.errorf(recv.pos, "cannot define new methods on non-local type %s", recv.typ)
					err = ""
				}
			} else {
				check.errorf(recv.pos, "invalid receiver type %s", recv.typ)
			}
			if err != "" {
				check.errorf(recv.pos, "invalid receiver type %s (%s)", recv.typ, err)
				// ok to continue
			}
		}
		sig.recv = recv
	}

	sig.params = NewTuple(params...)
	sig.results = NewTuple(results...)
	sig.variadic = variadic
}

// goTypeName returns the Go type name for typ and
// removes any occurrences of "types2." from that name.
func goTypeName(typ Type) string {
	return strings.Replace(fmt.Sprintf("%T", typ), "types2.", "", -1) // strings.ReplaceAll is not available in Go 1.4
}

// typInternal drives type checking of types.
// Must only be called by definedType or genericType.
//
func (check *Checker) typInternal(e0 syntax.Expr, def *Named) (T Type) {
	if check.conf.Trace {
		check.trace(e0.Pos(), "type %s", e0)
		check.indent++
		defer func() {
			check.indent--
			var under Type
			if T != nil {
				// Calling under() here may lead to endless instantiations.
				// Test case: type T[P any] *T[P]
				// TODO(gri) investigate if that's a bug or to be expected
				// (see also analogous comment in Checker.instantiate).
				under = T.Underlying()
			}
			if T == under {
				check.trace(e0.Pos(), "=> %s // %s", T, goTypeName(T))
			} else {
				check.trace(e0.Pos(), "=> %s (under = %s) // %s", T, under, goTypeName(T))
			}
		}()
	}

	switch e := e0.(type) {
	case *syntax.BadExpr:
		// ignore - error reported before

	case *syntax.Name:
		var x operand
		check.ident(&x, e, def, true)

		switch x.mode {
		case typexpr:
			typ := x.typ
			def.setUnderlying(typ)
			return typ
		case invalid:
			// ignore - error reported before
		case novalue:
			check.errorf(&x, "%s used as type", &x)
		default:
			check.errorf(&x, "%s is not a type", &x)
		}

	case *syntax.SelectorExpr:
		var x operand
		check.selector(&x, e)

		switch x.mode {
		case typexpr:
			typ := x.typ
			def.setUnderlying(typ)
			return typ
		case invalid:
			// ignore - error reported before
		case novalue:
			check.errorf(&x, "%s used as type", &x)
		default:
			check.errorf(&x, "%s is not a type", &x)
		}

	case *syntax.IndexExpr:
		return check.instantiatedType(e.X, unpackExpr(e.Index), def)

	case *syntax.ParenExpr:
		// Generic types must be instantiated before they can be used in any form.
		// Consequently, generic types cannot be parenthesized.
		return check.definedType(e.X, def)

	case *syntax.ArrayType:
		typ := new(Array)
		def.setUnderlying(typ)
		if e.Len != nil {
			typ.len = check.arrayLength(e.Len)
		} else {
			// [...]array
			check.error(e, "invalid use of [...] array (outside a composite literal)")
			typ.len = -1
		}
		typ.elem = check.varType(e.Elem)
		if typ.len >= 0 {
			return typ
		}
		// report error if we encountered [...]

	case *syntax.SliceType:
		typ := new(Slice)
		def.setUnderlying(typ)
		typ.elem = check.varType(e.Elem)
		return typ

	case *syntax.DotsType:
		// dots are handled explicitly where they are legal
		// (array composite literals and parameter lists)
		check.error(e, "invalid use of '...'")
		check.use(e.Elem)

	case *syntax.StructType:
		typ := new(Struct)
		def.setUnderlying(typ)
		check.structType(typ, e)
		return typ

	case *syntax.Operation:
		if e.Op == syntax.Mul && e.Y == nil {
			typ := new(Pointer)
			def.setUnderlying(typ)
			typ.base = check.varType(e.X)
			return typ
		}

		check.errorf(e0, "%s is not a type", e0)
		check.use(e0)

	case *syntax.FuncType:
		typ := new(Signature)
		def.setUnderlying(typ)
		check.funcType(typ, nil, nil, e)
		return typ

	case *syntax.InterfaceType:
		typ := new(Interface)
		def.setUnderlying(typ)
		if def != nil {
			typ.obj = def.obj
		}
		check.interfaceType(typ, e, def)
		return typ

	case *syntax.MapType:
		typ := new(Map)
		def.setUnderlying(typ)

		typ.key = check.varType(e.Key)
		typ.elem = check.varType(e.Value)

		// spec: "The comparison operators == and != must be fully defined
		// for operands of the key type; thus the key type must not be a
		// function, map, or slice."
		//
		// Delay this check because it requires fully setup types;
		// it is safe to continue in any case (was issue 6667).
		check.later(func() {
			if !Comparable(typ.key) {
				var why string
				if asTypeParam(typ.key) != nil {
					why = " (missing comparable constraint)"
				}
				check.errorf(e.Key, "invalid map key type %s%s", typ.key, why)
			}
		})

		return typ

	case *syntax.ChanType:
		typ := new(Chan)
		def.setUnderlying(typ)

		dir := SendRecv
		switch e.Dir {
		case 0:
			// nothing to do
		case syntax.SendOnly:
			dir = SendOnly
		case syntax.RecvOnly:
			dir = RecvOnly
		default:
			check.errorf(e, invalidAST+"unknown channel direction %d", e.Dir)
			// ok to continue
		}

		typ.dir = dir
		typ.elem = check.varType(e.Elem)
		return typ

	default:
		check.errorf(e0, "%s is not a type", e0)
		check.use(e0)
	}

	typ := Typ[Invalid]
	def.setUnderlying(typ)
	return typ
}

// typeOrNil type-checks the type expression (or nil value) e
// and returns the type of e, or nil. If e is a type, it must
// not be an (uninstantiated) generic type.
// If e is neither a type nor nil, typeOrNil returns Typ[Invalid].
// TODO(gri) should we also disallow non-var types?
func (check *Checker) typOrNil(e syntax.Expr) Type {
	var x operand
	check.rawExpr(&x, e, nil)
	switch x.mode {
	case invalid:
		// ignore - error reported before
	case novalue:
		check.errorf(&x, "%s used as type", &x)
	case typexpr:
		check.instantiatedOperand(&x)
		return x.typ
	case nilvalue:
		return nil
	default:
		check.errorf(&x, "%s is not a type", &x)
	}
	return Typ[Invalid]
}

func (check *Checker) instantiatedType(x syntax.Expr, targs []syntax.Expr, def *Named) Type {
	b := check.genericType(x, true) // TODO(gri) what about cycles?
	if b == Typ[Invalid] {
		return b // error already reported
	}
	base := asNamed(b)
	if base == nil {
		unreachable() // should have been caught by genericType
	}

	// create a new type instance rather than instantiate the type
	// TODO(gri) should do argument number check here rather than
	//           when instantiating the type?
	typ := new(instance)
	def.setUnderlying(typ)

	typ.check = check
	typ.pos = x.Pos()
	typ.base = base

	// evaluate arguments (always)
	typ.targs = check.typeList(targs)
	if typ.targs == nil {
		def.setUnderlying(Typ[Invalid]) // avoid later errors due to lazy instantiation
		return Typ[Invalid]
	}

	// determine argument positions (for error reporting)
	typ.poslist = make([]syntax.Pos, len(targs))
	for i, arg := range targs {
		typ.poslist[i] = syntax.StartPos(arg)
	}

	// make sure we check instantiation works at least once
	// and that the resulting type is valid
	check.later(func() {
		t := typ.expand()
		check.validType(t, nil)
	})

	return typ
}

// arrayLength type-checks the array length expression e
// and returns the constant length >= 0, or a value < 0
// to indicate an error (and thus an unknown length).
func (check *Checker) arrayLength(e syntax.Expr) int64 {
	var x operand
	check.expr(&x, e)
	if x.mode != constant_ {
		if x.mode != invalid {
			check.errorf(&x, "array length %s must be constant", &x)
		}
		return -1
	}
	if isUntyped(x.typ) || isInteger(x.typ) {
		if val := constant.ToInt(x.val); val.Kind() == constant.Int {
			if representableConst(val, check, Typ[Int], nil) {
				if n, ok := constant.Int64Val(val); ok && n >= 0 {
					return n
				}
				check.errorf(&x, "invalid array length %s", &x)
				return -1
			}
		}
	}
	check.errorf(&x, "array length %s must be integer", &x)
	return -1
}

// typeList provides the list of types corresponding to the incoming expression list.
// If an error occurred, the result is nil, but all list elements were type-checked.
func (check *Checker) typeList(list []syntax.Expr) []Type {
	res := make([]Type, len(list)) // res != nil even if len(list) == 0
	for i, x := range list {
		t := check.varType(x)
		if t == Typ[Invalid] {
			res = nil
		}
		if res != nil {
			res[i] = t
		}
	}
	return res
}

// collectParams declares the parameters of list in scope and returns the corresponding
// variable list. If type0 != nil, it is used instead of the first type in list.
func (check *Checker) collectParams(scope *Scope, list []*syntax.Field, type0 syntax.Expr, variadicOk bool) (params []*Var, variadic bool) {
	if list == nil {
		return
	}

	var named, anonymous bool

	var typ Type
	var prev syntax.Expr
	for i, field := range list {
		ftype := field.Type
		// type-check type of grouped fields only once
		if ftype != prev {
			prev = ftype
			if i == 0 && type0 != nil {
				ftype = type0
			}
			if t, _ := ftype.(*syntax.DotsType); t != nil {
				ftype = t.Elem
				if variadicOk && i == len(list)-1 {
					variadic = true
				} else {
					check.softErrorf(t, "can only use ... with final parameter in list")
					// ignore ... and continue
				}
			}
			typ = check.varType(ftype)
		}
		// The parser ensures that f.Tag is nil and we don't
		// care if a constructed AST contains a non-nil tag.
		if field.Name != nil {
			// named parameter
			name := field.Name.Value
			if name == "" {
				check.error(field.Name, invalidAST+"anonymous parameter")
				// ok to continue
			}
			par := NewParam(field.Name.Pos(), check.pkg, name, typ)
			check.declare(scope, field.Name, par, scope.pos)
			params = append(params, par)
			named = true
		} else {
			// anonymous parameter
			par := NewParam(field.Pos(), check.pkg, "", typ)
			check.recordImplicit(field, par)
			params = append(params, par)
			anonymous = true
		}
	}

	if named && anonymous {
		check.error(list[0], invalidAST+"list contains both named and anonymous parameters")
		// ok to continue
	}

	// For a variadic function, change the last parameter's type from T to []T.
	// Since we type-checked T rather than ...T, we also need to retro-actively
	// record the type for ...T.
	if variadic {
		last := params[len(params)-1]
		last.typ = &Slice{elem: last.typ}
		check.recordTypeAndValue(list[len(list)-1].Type, typexpr, last.typ, nil)
	}

	return
}
