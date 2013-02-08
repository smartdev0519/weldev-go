package interp

// Values
//
// All interpreter values are "boxed" in the empty interface, value.
// The range of possible dynamic types within value are:
//
// - bool
// - numbers (all built-in int/float/complex types are distinguished)
// - string
// - map[value]value --- maps for which  usesBuiltinMap(keyType)
//   *hashmap        --- maps for which !usesBuiltinMap(keyType)
// - chan value
// - []value --- slices
// - iface --- interfaces.
// - structure --- structs.  Fields are ordered and accessed by numeric indices.
// - array --- arrays.
// - *value --- pointers.  Careful: *value is a distinct type from *array etc.
// - *ssa.Function \
//   *ssa.Builtin   } --- functions.
//   *closure      /
// - tuple --- as returned by Ret, Next, "value,ok" modes, etc.
// - iter --- iterators from 'range'.
// - bad --- a poison pill for locals that have gone out of scope.
// - rtype -- the interpreter's concrete implementation of reflect.Type
//
// Note that nil is not on this list.
//
// Pay close attention to whether or not the dynamic type is a pointer.
// The compiler cannot help you since value is an empty interface.

import (
	"bytes"
	"exp/ssa"
	"fmt"
	"go/types"
	"io"
	"reflect"
	"strings"
	"unsafe"
)

type value interface{}

type tuple []value

type array []value

type iface struct {
	t types.Type // never an "untyped" type
	v value
}

type structure []value

// For map, array, *array, slice, string or channel.
type iter interface {
	// next returns a Tuple (key, value, ok).
	// key and value are unaliased, e.g. copies of the sequence element.
	next() tuple
}

type closure struct {
	Fn  *ssa.Function
	Env []value
}

type bad struct{}

type rtype struct {
	t types.Type
}

// Hash functions and equivalence relation:

// hashString computes the FNV hash of s.
func hashString(s string) int {
	var h uint32
	for i := 0; i < len(s); i++ {
		h ^= uint32(s[i])
		h *= 16777619
	}
	return int(h)
}

// hashType returns a hash for t such that
// types.IsIdentical(x, y) => hashType(x) == hashType(y).
func hashType(t types.Type) int {
	return hashString(t.String()) // TODO(gri): provide a better hash
}

// usesBuiltinMap returns true if the built-in hash function and
// equivalence relation for type t are consistent with those of the
// interpreter's representation of type t.  Such types are: all basic
// types (bool, numbers, string), pointers and channels.
//
// usesBuiltinMap returns false for types that require a custom map
// implementation: interfaces, arrays and structs.
//
// Panic ensues if t is an invalid map key type: function, map or slice.
func usesBuiltinMap(t types.Type) bool {
	switch t := t.(type) {
	case *types.Basic, *types.Chan, *types.Pointer:
		return true
	case *types.NamedType:
		return usesBuiltinMap(t.Underlying)
	case *types.Interface, *types.Array, *types.Struct:
		return false
	}
	panic(fmt.Sprintf("invalid map key type: %T", t))
}

func (x array) eq(_y interface{}) bool {
	y := _y.(array)
	for i, xi := range x {
		if !equals(xi, y[i]) {
			return false
		}
	}
	return true
}

func (x array) hash() int {
	h := 0
	for _, xi := range x {
		h += hash(xi)
	}
	return h
}

func (x structure) eq(_y interface{}) bool {
	y := _y.(structure)
	// TODO(adonovan): fix: only non-blank fields should be
	// compared.  This requires that we have type information
	// available from the enclosing == operation or map access;
	// the value is not sufficient.
	for i, xi := range x {
		if !equals(xi, y[i]) {
			return false
		}
	}
	return true
}

func (x structure) hash() int {
	h := 0
	for _, xi := range x {
		h += hash(xi)
	}
	return h
}

func (x iface) eq(_y interface{}) bool {
	y := _y.(iface)
	return types.IsIdentical(x.t, y.t) && (x.t == nil || equals(x.v, y.v))
}

func (x iface) hash() int {
	return hashType(x.t)*8581 + hash(x.v)
}

func (x rtype) hash() int {
	return hashType(x.t)
}

func (x rtype) eq(y interface{}) bool {
	return types.IsIdentical(x.t, y.(rtype).t)
}

// equals returns true iff x and y are equal according to Go's
// linguistic equivalence relation.  In a well-typed program, the
// types of x and y are guaranteed equal.
func equals(x, y value) bool {
	switch x := x.(type) {
	case bool:
		return x == y.(bool)
	case int:
		return x == y.(int)
	case int8:
		return x == y.(int8)
	case int16:
		return x == y.(int16)
	case int32:
		return x == y.(int32)
	case int64:
		return x == y.(int64)
	case uint:
		return x == y.(uint)
	case uint8:
		return x == y.(uint8)
	case uint16:
		return x == y.(uint16)
	case uint32:
		return x == y.(uint32)
	case uint64:
		return x == y.(uint64)
	case uintptr:
		return x == y.(uintptr)
	case float32:
		return x == y.(float32)
	case float64:
		return x == y.(float64)
	case complex64:
		return x == y.(complex64)
	case complex128:
		return x == y.(complex128)
	case string:
		return x == y.(string)
	case *value:
		return x == y.(*value)
	case chan value:
		return x == y.(chan value)
	case structure:
		return x.eq(y)
	case array:
		return x.eq(y)
	case iface:
		return x.eq(y)
	case rtype:
		return x.eq(y)

		// Since the following types don't support comparison,
		// these cases are only reachable if one of x or y is
		// (literally) nil.
	case *hashmap:
		return x == y.(*hashmap)
	case map[value]value:
		return (x != nil) == (y.(map[value]value) != nil)
	case *ssa.Function:
		return x == y.(*ssa.Function)
	case *closure:
		return x == y.(*closure)
	case []value:
		return (x != nil) == (y.([]value) != nil)
	}
	panic(fmt.Sprintf("comparing incomparable type %T", x))
}

// Returns an integer hash of x such that equals(x, y) => hash(x) == hash(y).
func hash(x value) int {
	switch x := x.(type) {
	case bool:
		if x {
			return 1
		}
		return 0
	case int:
		return x
	case int8:
		return int(x)
	case int16:
		return int(x)
	case int32:
		return int(x)
	case int64:
		return int(x)
	case uint:
		return int(x)
	case uint8:
		return int(x)
	case uint16:
		return int(x)
	case uint32:
		return int(x)
	case uint64:
		return int(x)
	case uintptr:
		return int(x)
	case float32:
		return int(x)
	case float64:
		return int(x)
	case complex64:
		return int(real(x))
	case complex128:
		return int(real(x))
	case string:
		return hashString(x)
	case *value:
		return int(uintptr(unsafe.Pointer(x)))
	case chan value:
		return int(uintptr(reflect.ValueOf(x).Pointer()))
	case structure:
		return x.hash()
	case array:
		return x.hash()
	case iface:
		return x.hash()
	case rtype:
		return x.hash()
	}
	panic(fmt.Sprintf("%T is unhashable", x))
}

// copyVal returns a copy of value v.
// TODO(adonovan): add tests of aliasing and mutation.
func copyVal(v value) value {
	if v == nil {
		panic("copyVal(nil)")
	}
	switch v := v.(type) {
	case bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64, complex64, complex128, string, unsafe.Pointer:
		return v
	case map[value]value:
		return v
	case *hashmap:
		return v
	case chan value:
		return v
	case *value:
		return v
	case *ssa.Function, *ssa.Builtin, *closure:
		return v
	case iface:
		return v
	case []value:
		return v
	case structure:
		a := make(structure, len(v))
		copy(a, v)
		return a
	case array:
		a := make(array, len(v))
		copy(a, v)
		return a
	case tuple:
		break
	case rtype:
		return v
	}
	panic(fmt.Sprintf("cannot copy %T", v))
}

// Prints in the style of built-in println.
// (More or less; in gc println is actually a compiler intrinsic and
// can distinguish println(1) from println(interface{}(1)).)
func toWriter(w io.Writer, v value) {
	switch v := v.(type) {
	case nil, bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64, complex64, complex128, string:
		fmt.Fprintf(w, "%v", v)

	case map[value]value:
		io.WriteString(w, "map[")
		sep := " "
		for k, e := range v {
			io.WriteString(w, sep)
			sep = " "
			toWriter(w, k)
			io.WriteString(w, ":")
			toWriter(w, e)
		}
		io.WriteString(w, "]")

	case *hashmap:
		io.WriteString(w, "map[")
		sep := " "
		for _, e := range v.table {
			for e != nil {
				io.WriteString(w, sep)
				sep = " "
				toWriter(w, e.key)
				io.WriteString(w, ":")
				toWriter(w, e.value)
				e = e.next
			}
		}
		io.WriteString(w, "]")

	case chan value:
		fmt.Fprintf(w, "%v", v) // (an address)

	case *value:
		if v == nil {
			io.WriteString(w, "<nil>")
		} else {
			fmt.Fprintf(w, "%p", v)
		}

	case iface:
		toWriter(w, v.v)

	case structure:
		io.WriteString(w, "{")
		for i, e := range v {
			if i > 0 {
				io.WriteString(w, " ")
			}
			toWriter(w, e)
		}
		io.WriteString(w, "}")

	case array:
		io.WriteString(w, "[")
		for i, e := range v {
			if i > 0 {
				io.WriteString(w, " ")
			}
			toWriter(w, e)
		}
		io.WriteString(w, "]")

	case []value:
		io.WriteString(w, "[")
		for i, e := range v {
			if i > 0 {
				io.WriteString(w, " ")
			}
			toWriter(w, e)
		}
		io.WriteString(w, "]")

	case *ssa.Function, *ssa.Builtin, *closure:
		fmt.Fprintf(w, "%p", v) // (an address)

	case rtype:
		io.WriteString(w, v.t.String())

	case tuple:
		// Unreachable in well-formed Go programs
		io.WriteString(w, "(")
		for i, e := range v {
			if i > 0 {
				io.WriteString(w, ", ")
			}
			toWriter(w, e)
		}
		io.WriteString(w, ")")

	default:
		fmt.Fprintf(w, "<%T>", v)
	}
}

// Implements printing of Go values in the style of built-in println.
func toString(v value) string {
	var b bytes.Buffer
	toWriter(&b, v)
	return b.String()
}

// ------------------------------------------------------------------------
// Iterators

type arrayIter struct {
	a array
	i int
}

func (it *arrayIter) next() tuple {
	okv := make(tuple, 3)
	ok := it.i < len(it.a)
	okv[0] = ok
	if ok {
		okv[1] = it.i
		okv[2] = copyVal(it.a[it.i])
	}
	it.i++
	return okv
}

type chanIter chan value

func (it chanIter) next() tuple {
	okv := make(tuple, 3)
	okv[1], okv[0] = <-it
	return okv
}

type stringIter struct {
	*strings.Reader
	i int
}

func (it *stringIter) next() tuple {
	okv := make(tuple, 3)
	ch, n, err := it.ReadRune()
	ok := err != io.EOF
	okv[0] = ok
	if ok {
		okv[1] = it.i
		okv[2] = ch
	}
	it.i += n
	return okv
}

type mapIter chan [2]value

func (it mapIter) next() tuple {
	kv, ok := <-it
	return tuple{ok, kv[0], kv[1]}
}
