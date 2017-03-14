// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import "testing"

func TestWriteBarrierStoreOrder(t *testing.T) {
	// Make sure writebarrier phase works even StoreWB ops are not in dependency order
	c := testConfig(t)
	ptrType := &TypeImpl{Size_: 8, Ptr: true, Name: "testptr"} // dummy for testing
	fun := Fun(c, "entry",
		Bloc("entry",
			Valu("start", OpInitMem, TypeMem, 0, nil),
			Valu("sb", OpSB, TypeInvalid, 0, nil),
			Valu("sp", OpSP, TypeInvalid, 0, nil),
			Valu("v", OpConstNil, ptrType, 0, nil),
			Valu("addr1", OpAddr, ptrType, 0, nil, "sb"),
			Valu("wb2", OpStore, TypeMem, 0, ptrType, "addr1", "v", "wb1"),
			Valu("wb1", OpStore, TypeMem, 0, ptrType, "addr1", "v", "start"), // wb1 and wb2 are out of order
			Goto("exit")),
		Bloc("exit",
			Exit("wb2")))

	CheckFunc(fun.f)
	writebarrier(fun.f)
	CheckFunc(fun.f)
}

func TestWriteBarrierPhi(t *testing.T) {
	// Make sure writebarrier phase works for single-block loop, where
	// a Phi op takes the store in the same block as argument.
	// See issue #19067.
	c := testConfig(t)
	ptrType := &TypeImpl{Size_: 8, Ptr: true, Name: "testptr"} // dummy for testing
	fun := Fun(c, "entry",
		Bloc("entry",
			Valu("start", OpInitMem, TypeMem, 0, nil),
			Valu("sb", OpSB, TypeInvalid, 0, nil),
			Valu("sp", OpSP, TypeInvalid, 0, nil),
			Goto("loop")),
		Bloc("loop",
			Valu("phi", OpPhi, TypeMem, 0, nil, "start", "wb"),
			Valu("v", OpConstNil, ptrType, 0, nil),
			Valu("addr", OpAddr, ptrType, 0, nil, "sb"),
			Valu("wb", OpStore, TypeMem, 0, ptrType, "addr", "v", "phi"), // has write barrier
			Goto("loop")))

	CheckFunc(fun.f)
	writebarrier(fun.f)
	CheckFunc(fun.f)
}
