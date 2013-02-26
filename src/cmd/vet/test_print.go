// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build vet_test

// This file contains tests for the printf checker.

package main

import (
	"fmt"
	"unsafe" // just for test case printing unsafe.Pointer
)

func UnsafePointerPrintfTest() {
	var up unsafe.Pointer
	fmt.Printf("%p, %x %X", up, up, up)
}

// Error methods that do not satisfy the Error interface and should be checked.
type errorTest1 int

func (errorTest1) Error(...interface{}) string {
	return "hi"
}

type errorTest2 int // Analogous to testing's *T type.
func (errorTest2) Error(...interface{}) {
}

type errorTest3 int

func (errorTest3) Error() { // No return value.
}

type errorTest4 int

func (errorTest4) Error() int { // Different return type.
	return 3
}

type errorTest5 int

func (errorTest5) error() { // niladic; don't complain if no args (was bug)
}

// This function never executes, but it serves as a simple test for the program.
// Test with make test.
func PrintfTests() {
	var b bool
	var i int
	var r rune
	var s string
	var x float64
	var p *int
	// Some good format/argtypes
	fmt.Printf("")
	fmt.Printf("%b %b", 3, i)
	fmt.Printf("%c %c %c %c", 3, i, 'x', r)
	fmt.Printf("%d %d", 3, i)
	fmt.Printf("%e %e %e", 3, 3e9, x)
	fmt.Printf("%E %E %E", 3, 3e9, x)
	fmt.Printf("%f %f %f", 3, 3e9, x)
	fmt.Printf("%F %F %F", 3, 3e9, x)
	fmt.Printf("%g %g %g", 3, 3e9, x)
	fmt.Printf("%G %G %G", 3, 3e9, x)
	fmt.Printf("%o %o", 3, i)
	fmt.Printf("%p %p", p, nil)
	fmt.Printf("%q %q %q %q", 3, i, 'x', r)
	fmt.Printf("%s %s", "hi", s)
	fmt.Printf("%t %t", true, b)
	fmt.Printf("%T %T", 3, i)
	fmt.Printf("%U %U", 3, i)
	fmt.Printf("%v %v", 3, i)
	fmt.Printf("%x %x %x %x", 3, i, "hi", s)
	fmt.Printf("%X %X %X %X", 3, i, "hi", s)
	fmt.Printf("%.*s %d %g", 3, "hi", 23, 2.3)
	// Some bad format/argTypes
	fmt.Printf("%b", 2.3)                      // ERROR "arg for printf verb %b of wrong type"
	fmt.Printf("%c", 2.3)                      // ERROR "arg for printf verb %c of wrong type"
	fmt.Printf("%d", 2.3)                      // ERROR "arg for printf verb %d of wrong type"
	fmt.Printf("%e", "hi")                     // ERROR "arg for printf verb %e of wrong type"
	fmt.Printf("%E", true)                     // ERROR "arg for printf verb %E of wrong type"
	fmt.Printf("%f", "hi")                     // ERROR "arg for printf verb %f of wrong type"
	fmt.Printf("%F", 'x')                      // ERROR "arg for printf verb %F of wrong type"
	fmt.Printf("%g", "hi")                     // ERROR "arg for printf verb %g of wrong type"
	fmt.Printf("%G", i)                        // ERROR "arg for printf verb %G of wrong type"
	fmt.Printf("%o", x)                        // ERROR "arg for printf verb %o of wrong type"
	fmt.Printf("%p", 23)                       // ERROR "arg for printf verb %p of wrong type"
	fmt.Printf("%q", x)                        // ERROR "arg for printf verb %q of wrong type"
	fmt.Printf("%s", b)                        // ERROR "arg for printf verb %s of wrong type"
	fmt.Printf("%t", 23)                       // ERROR "arg for printf verb %t of wrong type"
	fmt.Printf("%U", x)                        // ERROR "arg for printf verb %U of wrong type"
	fmt.Printf("%x", nil)                      // ERROR "arg for printf verb %x of wrong type"
	fmt.Printf("%X", 2.3)                      // ERROR "arg for printf verb %X of wrong type"
	fmt.Printf("%.*s %d %g", 3, "hi", 23, 'x') // ERROR "arg for printf verb %g of wrong type"
	// TODO
	fmt.Println()                      // not an error
	fmt.Println("%s", "hi")            // ERROR "possible formatting directive in Println call"
	fmt.Printf("%s", "hi", 3)          // ERROR "wrong number of args for format in Printf call"
	fmt.Printf("%"+("s"), "hi", 3)     // ERROR "wrong number of args for format in Printf call"
	fmt.Printf("%s%%%d", "hi", 3)      // correct
	fmt.Printf("%08s", "woo")          // correct
	fmt.Printf("% 8s", "woo")          // correct
	fmt.Printf("%.*d", 3, 3)           // correct
	fmt.Printf("%.*d", 3, 3, 3)        // ERROR "wrong number of args for format in Printf call"
	fmt.Printf("%.*d", "hi", 3)        // ERROR "arg for \* in printf format not of type int"
	fmt.Printf("%.*d", i, 3)           // correct
	fmt.Printf("%.*d", s, 3)           // ERROR "arg for \* in printf format not of type int"
	fmt.Printf("%q %q", multi()...)    // ok
	fmt.Printf("%#q", `blah`)          // ok
	printf("now is the time", "buddy") // ERROR "no formatting directive"
	Printf("now is the time", "buddy") // ERROR "no formatting directive"
	Printf("hi")                       // ok
	const format = "%s %s\n"
	Printf(format, "hi", "there")
	Printf(format, "hi") // ERROR "wrong number of args for format in Printf call"
	f := new(File)
	f.Warn(0, "%s", "hello", 3)  // ERROR "possible formatting directive in Warn call"
	f.Warnf(0, "%s", "hello", 3) // ERROR "wrong number of args for format in Warnf call"
	f.Warnf(0, "%r", "hello")    // ERROR "unrecognized printf verb"
	f.Warnf(0, "%#s", "hello")   // ERROR "unrecognized printf flag"
	// Something that satisfies the error interface.
	var e error
	fmt.Println(e.Error()) // ok
	// Something that looks like an error interface but isn't, such as the (*T).Error method
	// in the testing package.
	var et1 errorTest1
	fmt.Println(et1.Error())        // ERROR "no args in Error call"
	fmt.Println(et1.Error("hi"))    // ok
	fmt.Println(et1.Error("%d", 3)) // ERROR "possible formatting directive in Error call"
	var et2 errorTest2
	et2.Error()        // ERROR "no args in Error call"
	et2.Error("hi")    // ok, not an error method.
	et2.Error("%d", 3) // ERROR "possible formatting directive in Error call"
	var et3 errorTest3
	et3.Error() // ok, not an error method.
	var et4 errorTest4
	et4.Error() // ok, not an error method.
	var et5 errorTest5
	et5.error() // ok, not an error method.
}

// printf is used by the test.
func printf(format string, args ...interface{}) {
	panic("don't call - testing only")
}

// multi is used by the test.
func multi() []interface{} {
	panic("don't call - testing only")
}
