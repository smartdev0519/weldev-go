// asmcheck

// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package codegen

import "unsafe"

// This file contains code generation tests related to the comparison
// operators.

// -------------- //
//    Equality    //
// -------------- //

// Check that compare to constant string use 2/4/8 byte compares

func CompareString1(s string) bool {
	// amd64:`CMPW\t\(.*\), [$]`
	return s == "xx"
}

func CompareString2(s string) bool {
	// amd64:`CMPL\t\(.*\), [$]`
	return s == "xxxx"
}

func CompareString3(s string) bool {
	// amd64:`CMPQ\t\(.*\), [A-Z]`
	return s == "xxxxxxxx"
}

// Check that arrays compare use 2/4/8 byte compares

func CompareArray1(a, b [2]byte) bool {
	// amd64:`CMPW\t""[.+_a-z0-9]+\(SP\), [A-Z]`
	return a == b
}

func CompareArray2(a, b [3]uint16) bool {
	// amd64:`CMPL\t""[.+_a-z0-9]+\(SP\), [A-Z]`
	// amd64:`CMPW\t""[.+_a-z0-9]+\(SP\), [A-Z]`
	return a == b
}

func CompareArray3(a, b [3]int16) bool {
	// amd64:`CMPL\t""[.+_a-z0-9]+\(SP\), [A-Z]`
	// amd64:`CMPW\t""[.+_a-z0-9]+\(SP\), [A-Z]`
	return a == b
}

func CompareArray4(a, b [12]int8) bool {
	// amd64:`CMPQ\t""[.+_a-z0-9]+\(SP\), [A-Z]`
	// amd64:`CMPL\t""[.+_a-z0-9]+\(SP\), [A-Z]`
	return a == b
}

func CompareArray5(a, b [15]byte) bool {
	// amd64:`CMPQ\t""[.+_a-z0-9]+\(SP\), [A-Z]`
	return a == b
}

// This was a TODO in mapaccess1_faststr
func CompareArray6(a, b unsafe.Pointer) bool {
	// amd64:`CMPL\t\(.*\), [A-Z]`
	return *((*[4]byte)(a)) != *((*[4]byte)(b))
}
