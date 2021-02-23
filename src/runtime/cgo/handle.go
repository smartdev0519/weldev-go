// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgo

import (
	"sync"
	"sync/atomic"
)

// Handle provides a safe representation of Go values to pass between
// Go and C. The zero value of a handle is not a valid handle, and thus
// is safe to use as a sentinel in C APIs.
//
// The underlying type of Handle is guaranteed to fit in an integer type
// that is large enough to hold the bit pattern of any pointer.
// For instance, on the Go side:
//
// 	package main
//
// 	/*
// 	#include <stdint.h> // for uintptr_t
//
// 	extern void MyGoPrint(uintptr_t handle);
// 	void myprint(uintptr_t handle);
// 	*/
// 	import "C"
// 	import "runtime/cgo"
//
// 	//export MyGoPrint
// 	func MyGoPrint(handle C.uintptr_t) {
// 		h := cgo.Handle(handle)
// 		val := h.Value().(int)
// 		println(val)
// 		h.Delete()
// 	}
//
// 	func main() {
// 		val := 42
// 		C.myprint(C.uintptr_t(cgo.NewHandle(val)))
// 		// Output: 42
// 	}
//
// and on the C side:
//
// 	#include <stdint.h> // for uintptr_t
//
// 	// A Go function
// 	extern void MyGoPrint(uintptr_t handle);
//
// 	// A C function
// 	void myprint(uintptr_t handle) {
// 	    MyGoPrint(handle);
// 	}
type Handle uintptr

// NewHandle returns a handle for a given value.
//
// The handle is valid until the program calls Delete on it. The handle
// uses resources, and this package assumes that C code may hold on to
// the handle, so a program must explicitly call Delete when the handle
// is no longer needed.
//
// The intended use is to pass the returned handle to C code, which
// passes it back to Go, which calls Value.
func NewHandle(v interface{}) Handle {
	h := atomic.AddUintptr(&handleIdx, 1)
	if h == 0 {
		panic("runtime/cgo: ran out of handle space")
	}

	handles.Store(h, v)
	return Handle(h)
}

// Value returns the associated Go value for a valid handle.
//
// The method panics if the handle is invalid.
func (h Handle) Value() interface{} {
	v, ok := handles.Load(uintptr(h))
	if !ok {
		panic("runtime/cgo: misuse of an invalid Handle")
	}
	return v
}

// Delete invalidates a handle. This method should only be called once
// the program no longer needs to pass the handle to C and the C code
// no longer has a copy of the handle value.
//
// The method panics if the handle is invalid.
func (h Handle) Delete() {
	_, ok := handles.LoadAndDelete(uintptr(h))
	if !ok {
		panic("runtime/cgo: misuse of an invalid Handle")
	}
}

var (
	handles   = sync.Map{} // map[Handle]interface{}
	handleIdx uintptr      // atomic
)
