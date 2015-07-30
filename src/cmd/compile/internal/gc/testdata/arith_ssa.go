// run

// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Tests arithmetic expressions

package main

// test64BitConstMulti tests that rewrite rules don't fold 64 bit constants
// into multiply instructions.
func test64BitConstMult(a, b int64) {
	want := 34359738369*a + b*34359738370
	if got := test64BitConstMult_ssa(a, b); want != got {
		println("test64BitConstMult failed, wanted", want, "got", got)
		failed = true
	}
}
func test64BitConstMult_ssa(a, b int64) int64 {
	switch { // prevent inlining
	}
	return 34359738369*a + b*34359738370
}

// test64BitConstAdd tests that rewrite rules don't fold 64 bit constants
// into add instructions.
func test64BitConstAdd(a, b int64) {
	want := a + 575815584948629622 + b + 2991856197886747025
	if got := test64BitConstAdd_ssa(a, b); want != got {
		println("test64BitConstAdd failed, wanted", want, "got", got)
		failed = true
	}
}
func test64BitConstAdd_ssa(a, b int64) int64 {
	switch {
	}
	return a + 575815584948629622 + b + 2991856197886747025
}

// testRegallocCVSpill tests that regalloc spills a value whose last use is the
// current value.
func testRegallocCVSpill(a, b, c, d int8) {
	want := a + -32 + b + 63*c*-87*d
	if got := testRegallocCVSpill_ssa(a, b, c, d); want != got {
		println("testRegallocCVSpill failed, wanted", want, "got", got)
		failed = true
	}
}
func testRegallocCVSpill_ssa(a, b, c, d int8) int8 {
	switch {
	}
	return a + -32 + b + 63*c*-87*d
}

func testBitwiseLogic() {
	a, b := uint32(57623283), uint32(1314713839)
	if want, got := uint32(38551779), testBitwiseAnd_ssa(a, b); want != got {
		println("testBitwiseAnd failed, wanted", want, "got", got)
	}
	if want, got := uint32(1333785343), testBitwiseOr_ssa(a, b); want != got {
		println("testBitwiseAnd failed, wanted", want, "got", got)
	}
}

func testBitwiseAnd_ssa(a, b uint32) uint32 {
	switch { // prevent inlining
	}
	return a & b
}

func testBitwiseOr_ssa(a, b uint32) uint32 {
	switch { // prevent inlining
	}
	return a | b
}

// testSubqToNegq ensures that the SUBQ -> NEGQ translation works correctly.
func testSubqToNegq(a, b, c, d, e, f, g, h, i, j, k int64) {
	want := a + 8207351403619448057 - b - 1779494519303207690 + c*8810076340510052032*d - 4465874067674546219 - e*4361839741470334295 - f + 8688847565426072650*g*8065564729145417479
	if got := testSubqToNegq_ssa(a, b, c, d, e, f, g, h, i, j, k); want != got {
		println("testSubqToNegq failed, wanted", want, "got", got)
		failed = true
	}
}
func testSubqToNegq_ssa(a, b, c, d, e, f, g, h, i, j, k int64) int64 {
	switch { // prevent inlining
	}
	return a + 8207351403619448057 - b - 1779494519303207690 + c*8810076340510052032*d - 4465874067674546219 - e*4361839741470334295 - f + 8688847565426072650*g*8065564729145417479
}

var failed = false

func main() {

	test64BitConstMult(1, 2)
	test64BitConstAdd(1, 2)
	testRegallocCVSpill(1, 2, 3, 4)
	testSubqToNegq(1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2)
	if failed {
		panic("failed")
	}
}
