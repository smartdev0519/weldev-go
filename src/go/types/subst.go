// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements instantiation of generic types
// through substitution of type parameters by actual
// types.

package types

import (
	"bytes"
	"fmt"
	"go/token"
)

// TODO(rFindley) decide error codes for the errors in this file, and check
//                if error spans can be improved

type substMap struct {
	// The targs field is currently needed for *Named type substitution.
	// TODO(gri) rewrite that code, get rid of this field, and make this
	//           struct just the map (proj)
	targs []Type
	proj  map[*TypeParam]Type
}

// makeSubstMap creates a new substitution map mapping tpars[i] to targs[i].
// If targs[i] is nil, tpars[i] is not substituted.
func makeSubstMap(tpars []*TypeName, targs []Type) *substMap {
	assert(len(tpars) == len(targs))
	proj := make(map[*TypeParam]Type, len(tpars))
	for i, tpar := range tpars {
		// We must expand type arguments otherwise *instance
		// types end up as components in composite types.
		// TODO(gri) explain why this causes problems, if it does
		targ := expand(targs[i]) // possibly nil
		targs[i] = targ
		proj[tpar.typ.(*TypeParam)] = targ
	}
	return &substMap{targs, proj}
}

func (m *substMap) String() string {
	return fmt.Sprintf("%s", m.proj)
}

func (m *substMap) empty() bool {
	return len(m.proj) == 0
}

func (m *substMap) lookup(tpar *TypeParam) Type {
	if t := m.proj[tpar]; t != nil {
		return t
	}
	return tpar
}

// Instantiate instantiates the type typ with the given type arguments
// targs. To check type constraint satisfaction, verify must be set.
// pos and posList correspond to the instantiation and type argument
// positions respectively; posList may be nil or shorter than the number
// of type arguments provided.
// typ must be a *Named or a *Signature type, and its number of type
// parameters must match the number of provided type arguments.
// The receiver (check) may be nil if and only if verify is not set.
// The result is a new, instantiated (not generic) type of the same kind
// (either a *Named or a *Signature).
// Any methods attached to a *Named are simply copied; they are not
// instantiated.
func (check *Checker) Instantiate(pos token.Pos, typ Type, targs []Type, posList []token.Pos, verify bool) (res Type) {
	if verify && check == nil {
		panic("cannot have nil receiver if verify is set")
	}

	if check != nil && trace {
		check.trace(pos, "-- instantiating %s with %s", typ, typeListString(targs))
		check.indent++
		defer func() {
			check.indent--
			var under Type
			if res != nil {
				// Calling under() here may lead to endless instantiations.
				// Test case: type T[P any] T[P]
				// TODO(gri) investigate if that's a bug or to be expected.
				under = res.Underlying()
			}
			check.trace(pos, "=> %s (under = %s)", res, under)
		}()
	}

	assert(len(posList) <= len(targs))

	// TODO(gri) What is better here: work with TypeParams, or work with TypeNames?
	var tparams []*TypeName
	switch t := typ.(type) {
	case *Named:
		tparams = t.TParams()
	case *Signature:
		tparams = t.tparams
		defer func() {
			// If we had an unexpected failure somewhere don't panic below when
			// asserting res.(*Signature). Check for *Signature in case Typ[Invalid]
			// is returned.
			if _, ok := res.(*Signature); !ok {
				return
			}
			// If the signature doesn't use its type parameters, subst
			// will not make a copy. In that case, make a copy now (so
			// we can set tparams to nil w/o causing side-effects).
			if t == res {
				copy := *t
				res = &copy
			}
			// After instantiating a generic signature, it is not generic
			// anymore; we need to set tparams to nil.
			res.(*Signature).tparams = nil
		}()
	default:
		// only types and functions can be generic
		panic(fmt.Sprintf("%v: cannot instantiate %v", pos, typ))
	}

	// the number of supplied types must match the number of type parameters
	if len(targs) != len(tparams) {
		// TODO(gri) provide better error message
		if check != nil {
			check.errorf(atPos(pos), _Todo, "got %d arguments but %d type parameters", len(targs), len(tparams))
			return Typ[Invalid]
		}
		panic(fmt.Sprintf("%v: got %d arguments but %d type parameters", pos, len(targs), len(tparams)))
	}

	if len(tparams) == 0 {
		return typ // nothing to do (minor optimization)
	}

	smap := makeSubstMap(tparams, targs)

	// check bounds
	if verify {
		for i, tname := range tparams {
			// best position for error reporting
			pos := pos
			if i < len(posList) {
				pos = posList[i]
			}

			// stop checking bounds after the first failure
			if !check.satisfies(pos, targs[i], tname.typ.(*TypeParam), smap) {
				break
			}
		}
	}

	return check.subst(pos, typ, smap)
}

// satisfies reports whether the type argument targ satisfies the constraint of type parameter
// parameter tpar (after any of its type parameters have been substituted through smap).
// A suitable error is reported if the result is false.
// TODO(gri) This should be a method of interfaces or type sets.
func (check *Checker) satisfies(pos token.Pos, targ Type, tpar *TypeParam, smap *substMap) bool {
	iface := tpar.Bound()
	if iface.Empty() {
		return true // no type bound
	}

	// The type parameter bound is parameterized with the same type parameters
	// as the instantiated type; before we can use it for bounds checking we
	// need to instantiate it with the type arguments with which we instantiate
	// the parameterized type.
	iface = check.subst(pos, iface, smap).(*Interface)

	// targ must implement iface (methods)
	// - check only if we have methods
	if iface.NumMethods() > 0 {
		// If the type argument is a pointer to a type parameter, the type argument's
		// method set is empty.
		// TODO(gri) is this what we want? (spec question)
		if base, isPtr := deref(targ); isPtr && asTypeParam(base) != nil {
			check.errorf(atPos(pos), 0, "%s has no methods", targ)
			return false
		}
		if m, wrong := check.missingMethod(targ, iface, true); m != nil {
			// TODO(gri) needs to print updated name to avoid major confusion in error message!
			//           (print warning for now)
			// Old warning:
			// check.softErrorf(pos, "%s does not satisfy %s (warning: name not updated) = %s (missing method %s)", targ, tpar.bound, iface, m)
			if m.name == "==" {
				// We don't want to report "missing method ==".
				check.softErrorf(atPos(pos), 0, "%s does not satisfy comparable", targ)
			} else if wrong != nil {
				// TODO(gri) This can still report uninstantiated types which makes the error message
				//           more difficult to read then necessary.
				// TODO(rFindley) should this use parentheses rather than ':' for qualification?
				check.softErrorf(atPos(pos), _Todo,
					"%s does not satisfy %s: wrong method signature\n\tgot  %s\n\twant %s",
					targ, tpar.bound, wrong, m,
				)
			} else {
				check.softErrorf(atPos(pos), 0, "%s does not satisfy %s (missing method %s)", targ, tpar.bound, m.name)
			}
			return false
		}
	}

	// targ's underlying type must also be one of the interface types listed, if any
	if iface.typeSet().types == nil {
		return true // nothing to do
	}

	// If targ is itself a type parameter, each of its possible types, but at least one, must be in the
	// list of iface types (i.e., the targ type list must be a non-empty subset of the iface types).
	if targ := asTypeParam(targ); targ != nil {
		targBound := targ.Bound()
		if targBound.typeSet().types == nil {
			check.softErrorf(atPos(pos), _Todo, "%s does not satisfy %s (%s has no type constraints)", targ, tpar.bound, targ)
			return false
		}
		return iface.is(func(typ Type, tilde bool) bool {
			// TODO(gri) incorporate tilde information!
			if !iface.isSatisfiedBy(typ) {
				// TODO(gri) match this error message with the one below (or vice versa)
				check.softErrorf(atPos(pos), 0, "%s does not satisfy %s (%s type constraint %s not found in %s)", targ, tpar.bound, targ, typ, iface.typeSet().types)
				return false
			}
			return true
		})
	}

	// Otherwise, targ's type or underlying type must also be one of the interface types listed, if any.
	if !iface.isSatisfiedBy(targ) {
		check.softErrorf(atPos(pos), _Todo, "%s does not satisfy %s (%s not found in %s)", targ, tpar.bound, targ, iface.typeSet().types)
		return false
	}

	return true
}

// subst returns the type typ with its type parameters tpars replaced by
// the corresponding type arguments targs, recursively.
// subst is functional in the sense that it doesn't modify the incoming
// type. If a substitution took place, the result type is different from
// from the incoming type.
func (check *Checker) subst(pos token.Pos, typ Type, smap *substMap) Type {
	if smap.empty() {
		return typ
	}

	// common cases
	switch t := typ.(type) {
	case *Basic:
		return typ // nothing to do
	case *TypeParam:
		return smap.lookup(t)
	}

	// general case
	var subst subster
	subst.pos = pos
	subst.smap = smap
	if check != nil {
		subst.check = check
		subst.typMap = check.typMap
	} else {
		// If we don't have a *Checker and its global type map,
		// use a local version. Besides avoiding duplicate work,
		// the type map prevents infinite recursive substitution
		// for recursive types (example: type T[P any] *T[P]).
		subst.typMap = make(map[string]*Named)
	}
	return subst.typ(typ)
}

type subster struct {
	pos    token.Pos
	smap   *substMap
	check  *Checker // nil if called via Instantiate
	typMap map[string]*Named
}

func (subst *subster) typ(typ Type) Type {
	switch t := typ.(type) {
	case nil:
		// Call typOrNil if it's possible that typ is nil.
		panic("nil typ")

	case *Basic, *top:
		// nothing to do

	case *Array:
		elem := subst.typOrNil(t.elem)
		if elem != t.elem {
			return &Array{len: t.len, elem: elem}
		}

	case *Slice:
		elem := subst.typOrNil(t.elem)
		if elem != t.elem {
			return &Slice{elem: elem}
		}

	case *Struct:
		if fields, copied := subst.varList(t.fields); copied {
			return &Struct{fields: fields, tags: t.tags}
		}

	case *Pointer:
		base := subst.typ(t.base)
		if base != t.base {
			return &Pointer{base: base}
		}

	case *Tuple:
		return subst.tuple(t)

	case *Signature:
		// TODO(gri) rethink the recv situation with respect to methods on parameterized types
		// recv := subst.var_(t.recv) // TODO(gri) this causes a stack overflow - explain
		recv := t.recv
		params := subst.tuple(t.params)
		results := subst.tuple(t.results)
		if recv != t.recv || params != t.params || results != t.results {
			return &Signature{
				rparams: t.rparams,
				// TODO(rFindley) why can't we nil out tparams here, rather than in
				//                instantiate above?
				tparams:  t.tparams,
				scope:    t.scope,
				recv:     recv,
				params:   params,
				results:  results,
				variadic: t.variadic,
			}
		}

	case *Union:
		types, copied := subst.typeList(t.types)
		if copied {
			// TODO(gri) Remove duplicates that may have crept in after substitution
			//           (unlikely but possible). This matters for the Identical
			//           predicate on unions.
			return newUnion(types, t.tilde)
		}

	case *Interface:
		methods, mcopied := subst.funcList(t.methods)
		embeddeds, ecopied := subst.typeList(t.embeddeds)
		if mcopied || ecopied {
			iface := &Interface{methods: methods, embeddeds: embeddeds, complete: t.complete}
			if subst.check == nil {
				panic("internal error: cannot instantiate interfaces yet")
			}
			return iface
		}

	case *Map:
		key := subst.typ(t.key)
		elem := subst.typ(t.elem)
		if key != t.key || elem != t.elem {
			return &Map{key: key, elem: elem}
		}

	case *Chan:
		elem := subst.typ(t.elem)
		if elem != t.elem {
			return &Chan{dir: t.dir, elem: elem}
		}

	case *Named:
		// dump is for debugging
		dump := func(string, ...interface{}) {}
		if subst.check != nil && trace {
			subst.check.indent++
			defer func() {
				subst.check.indent--
			}()
			dump = func(format string, args ...interface{}) {
				subst.check.trace(subst.pos, format, args...)
			}
		}

		if t.TParams() == nil {
			dump(">>> %s is not parameterized", t)
			return t // type is not parameterized
		}

		var newTargs []Type

		if len(t.targs) > 0 {
			// already instantiated
			dump(">>> %s already instantiated", t)
			assert(len(t.targs) == len(t.TParams()))
			// For each (existing) type argument targ, determine if it needs
			// to be substituted; i.e., if it is or contains a type parameter
			// that has a type argument for it.
			for i, targ := range t.targs {
				dump(">>> %d targ = %s", i, targ)
				newTarg := subst.typ(targ)
				if newTarg != targ {
					dump(">>> substituted %d targ %s => %s", i, targ, newTarg)
					if newTargs == nil {
						newTargs = make([]Type, len(t.TParams()))
						copy(newTargs, t.targs)
					}
					newTargs[i] = newTarg
				}
			}

			if newTargs == nil {
				dump(">>> nothing to substitute in %s", t)
				return t // nothing to substitute
			}
		} else {
			// not yet instantiated
			dump(">>> first instantiation of %s", t)
			// TODO(rFindley) can we instead subst the tparam types here?
			newTargs = subst.smap.targs
		}

		// before creating a new named type, check if we have this one already
		h := instantiatedHash(t, newTargs)
		dump(">>> new type hash: %s", h)
		if named, found := subst.typMap[h]; found {
			dump(">>> found %s", named)
			return named
		}

		// create a new named type and populate typMap to avoid endless recursion
		tname := NewTypeName(subst.pos, t.obj.pkg, t.obj.name, nil)
		named := subst.check.newNamed(tname, t, t.Underlying(), t.TParams(), t.methods) // method signatures are updated lazily
		named.targs = newTargs
		subst.typMap[h] = named

		// do the substitution
		dump(">>> subst %s with %s (new: %s)", t.underlying, subst.smap, newTargs)
		named.underlying = subst.typOrNil(t.Underlying())
		named.fromRHS = named.underlying // for cycle detection (Checker.validType)

		return named

	case *TypeParam:
		return subst.smap.lookup(t)

	case *instance:
		// TODO(gri) can we avoid the expansion here and just substitute the type parameters?
		return subst.typ(t.expand())

	default:
		panic("unimplemented")
	}

	return typ
}

var instanceHashing = 0

func instantiatedHash(typ *Named, targs []Type) string {
	assert(instanceHashing == 0)
	instanceHashing++
	var buf bytes.Buffer
	writeTypeName(&buf, typ.obj, nil)
	buf.WriteByte('[')
	writeTypeList(&buf, targs, nil, nil)
	buf.WriteByte(']')
	instanceHashing--

	// With respect to the represented type, whether a
	// type is fully expanded or stored as instance
	// does not matter - they are the same types.
	// Remove the instanceMarkers printed for instances.
	res := buf.Bytes()
	i := 0
	for _, b := range res {
		if b != instanceMarker {
			res[i] = b
			i++
		}
	}

	return string(res[:i])
}

func typeListString(list []Type) string {
	var buf bytes.Buffer
	writeTypeList(&buf, list, nil, nil)
	return buf.String()
}

// typOrNil is like typ but if the argument is nil it is replaced with Typ[Invalid].
// A nil type may appear in pathological cases such as type T[P any] []func(_ T([]_))
// where an array/slice element is accessed before it is set up.
func (subst *subster) typOrNil(typ Type) Type {
	if typ == nil {
		return Typ[Invalid]
	}
	return subst.typ(typ)
}

func (subst *subster) var_(v *Var) *Var {
	if v != nil {
		if typ := subst.typ(v.typ); typ != v.typ {
			copy := *v
			copy.typ = typ
			return &copy
		}
	}
	return v
}

func (subst *subster) tuple(t *Tuple) *Tuple {
	if t != nil {
		if vars, copied := subst.varList(t.vars); copied {
			return &Tuple{vars: vars}
		}
	}
	return t
}

func (subst *subster) varList(in []*Var) (out []*Var, copied bool) {
	out = in
	for i, v := range in {
		if w := subst.var_(v); w != v {
			if !copied {
				// first variable that got substituted => allocate new out slice
				// and copy all variables
				new := make([]*Var, len(in))
				copy(new, out)
				out = new
				copied = true
			}
			out[i] = w
		}
	}
	return
}

func (subst *subster) func_(f *Func) *Func {
	if f != nil {
		if typ := subst.typ(f.typ); typ != f.typ {
			copy := *f
			copy.typ = typ
			return &copy
		}
	}
	return f
}

func (subst *subster) funcList(in []*Func) (out []*Func, copied bool) {
	out = in
	for i, f := range in {
		if g := subst.func_(f); g != f {
			if !copied {
				// first function that got substituted => allocate new out slice
				// and copy all functions
				new := make([]*Func, len(in))
				copy(new, out)
				out = new
				copied = true
			}
			out[i] = g
		}
	}
	return
}

func (subst *subster) typeList(in []Type) (out []Type, copied bool) {
	out = in
	for i, t := range in {
		if u := subst.typ(t); u != t {
			if !copied {
				// first function that got substituted => allocate new out slice
				// and copy all functions
				new := make([]Type, len(in))
				copy(new, out)
				out = new
				copied = true
			}
			out[i] = u
		}
	}
	return
}
