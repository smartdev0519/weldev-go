// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains tests for the printf checker.

package testdata

import (
	"fmt"
	"io"
	"math"
	"os"
	"unsafe" // just for test case printing unsafe.Pointer

	// For testing printf-like functions from external package.
	"github.com/foobar/externalprintf"
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
	var imap map[int]int
	var fslice []float64
	var c complex64
	// Some good format/argtypes
	fmt.Printf("")
	fmt.Printf("%b %b %b", 3, i, x)
	fmt.Printf("%c %c %c %c", 3, i, 'x', r)
	fmt.Printf("%d %d %d", 3, i, imap)
	fmt.Printf("%e %e %e %e", 3e9, x, fslice, c)
	fmt.Printf("%E %E %E %E", 3e9, x, fslice, c)
	fmt.Printf("%f %f %f %f", 3e9, x, fslice, c)
	fmt.Printf("%F %F %F %F", 3e9, x, fslice, c)
	fmt.Printf("%g %g %g %g", 3e9, x, fslice, c)
	fmt.Printf("%G %G %G %G", 3e9, x, fslice, c)
	fmt.Printf("%b %b %b %b", 3e9, x, fslice, c)
	fmt.Printf("%o %o", 3, i)
	fmt.Printf("%p %p", p, nil)
	fmt.Printf("%q %q %q %q", 3, i, 'x', r)
	fmt.Printf("%s %s %s", "hi", s, []byte{65})
	fmt.Printf("%t %t", true, b)
	fmt.Printf("%T %T", 3, i)
	fmt.Printf("%U %U", 3, i)
	fmt.Printf("%v %v", 3, i)
	fmt.Printf("%x %x %x %x", 3, i, "hi", s)
	fmt.Printf("%X %X %X %X", 3, i, "hi", s)
	fmt.Printf("%.*s %d %g", 3, "hi", 23, 2.3)
	fmt.Printf("%s", &stringerv)
	fmt.Printf("%v", &stringerv)
	fmt.Printf("%T", &stringerv)
	fmt.Printf("%s", &embeddedStringerv)
	fmt.Printf("%v", &embeddedStringerv)
	fmt.Printf("%T", &embeddedStringerv)
	fmt.Printf("%v", notstringerv)
	fmt.Printf("%T", notstringerv)
	fmt.Printf("%q", stringerarrayv)
	fmt.Printf("%v", stringerarrayv)
	fmt.Printf("%s", stringerarrayv)
	fmt.Printf("%v", notstringerarrayv)
	fmt.Printf("%T", notstringerarrayv)
	fmt.Printf("%d", new(Formatter))
	fmt.Printf("%*%", 2)               // Ridiculous but allowed.
	fmt.Printf("%s", interface{}(nil)) // Nothing useful we can say.

	fmt.Printf("%g", 1+2i)
	fmt.Printf("%#e %#E %#f %#F %#g %#G", 1.2, 1.2, 1.2, 1.2, 1.2, 1.2) // OK since Go 1.9
	// Some bad format/argTypes
	fmt.Printf("%b", "hi")                      // ERROR "Printf format %b has arg \x22hi\x22 of wrong type string"
	fmt.Printf("%t", c)                         // ERROR "Printf format %t has arg c of wrong type complex64"
	fmt.Printf("%t", 1+2i)                      // ERROR "Printf format %t has arg 1 \+ 2i of wrong type complex128"
	fmt.Printf("%c", 2.3)                       // ERROR "Printf format %c has arg 2.3 of wrong type float64"
	fmt.Printf("%d", 2.3)                       // ERROR "Printf format %d has arg 2.3 of wrong type float64"
	fmt.Printf("%e", "hi")                      // ERROR "Printf format %e has arg \x22hi\x22 of wrong type string"
	fmt.Printf("%E", true)                      // ERROR "Printf format %E has arg true of wrong type bool"
	fmt.Printf("%f", "hi")                      // ERROR "Printf format %f has arg \x22hi\x22 of wrong type string"
	fmt.Printf("%F", 'x')                       // ERROR "Printf format %F has arg 'x' of wrong type rune"
	fmt.Printf("%g", "hi")                      // ERROR "Printf format %g has arg \x22hi\x22 of wrong type string"
	fmt.Printf("%g", imap)                      // ERROR "Printf format %g has arg imap of wrong type map\[int\]int"
	fmt.Printf("%G", i)                         // ERROR "Printf format %G has arg i of wrong type int"
	fmt.Printf("%o", x)                         // ERROR "Printf format %o has arg x of wrong type float64"
	fmt.Printf("%p", 23)                        // ERROR "Printf format %p has arg 23 of wrong type int"
	fmt.Printf("%q", x)                         // ERROR "Printf format %q has arg x of wrong type float64"
	fmt.Printf("%s", b)                         // ERROR "Printf format %s has arg b of wrong type bool"
	fmt.Printf("%s", byte(65))                  // ERROR "Printf format %s has arg byte\(65\) of wrong type byte"
	fmt.Printf("%t", 23)                        // ERROR "Printf format %t has arg 23 of wrong type int"
	fmt.Printf("%U", x)                         // ERROR "Printf format %U has arg x of wrong type float64"
	fmt.Printf("%x", nil)                       // ERROR "Printf format %x has arg nil of wrong type untyped nil"
	fmt.Printf("%X", 2.3)                       // ERROR "Printf format %X has arg 2.3 of wrong type float64"
	fmt.Printf("%s", stringerv)                 // ERROR "Printf format %s has arg stringerv of wrong type testdata.stringer"
	fmt.Printf("%t", stringerv)                 // ERROR "Printf format %t has arg stringerv of wrong type testdata.stringer"
	fmt.Printf("%s", embeddedStringerv)         // ERROR "Printf format %s has arg embeddedStringerv of wrong type testdata.embeddedStringer"
	fmt.Printf("%t", embeddedStringerv)         // ERROR "Printf format %t has arg embeddedStringerv of wrong type testdata.embeddedStringer"
	fmt.Printf("%q", notstringerv)              // ERROR "Printf format %q has arg notstringerv of wrong type testdata.notstringer"
	fmt.Printf("%t", notstringerv)              // ERROR "Printf format %t has arg notstringerv of wrong type testdata.notstringer"
	fmt.Printf("%t", stringerarrayv)            // ERROR "Printf format %t has arg stringerarrayv of wrong type testdata.stringerarray"
	fmt.Printf("%t", notstringerarrayv)         // ERROR "Printf format %t has arg notstringerarrayv of wrong type testdata.notstringerarray"
	fmt.Printf("%q", notstringerarrayv)         // ERROR "Printf format %q has arg notstringerarrayv of wrong type testdata.notstringerarray"
	fmt.Printf("%d", Formatter(true))           // ERROR "Printf format %d has arg Formatter\(true\) of wrong type testdata.Formatter"
	fmt.Printf("%z", FormatterVal(true))        // correct (the type is responsible for formatting)
	fmt.Printf("%d", FormatterVal(true))        // correct (the type is responsible for formatting)
	fmt.Printf("%s", nonemptyinterface)         // correct (the type is responsible for formatting)
	fmt.Printf("%.*s %d %6g", 3, "hi", 23, 'x') // ERROR "Printf format %6g has arg 'x' of wrong type rune"
	fmt.Println()                               // not an error
	fmt.Println("%s", "hi")                     // ERROR "Println call has possible formatting directive %s"
	fmt.Println("0.0%")                         // correct (trailing % couldn't be a formatting directive)
	fmt.Printf("%s", "hi", 3)                   // ERROR "Printf call needs 1 arg but has 2 args"
	_ = fmt.Sprintf("%"+("s"), "hi", 3)         // ERROR "Sprintf call needs 1 arg but has 2 args"
	fmt.Printf("%s%%%d", "hi", 3)               // correct
	fmt.Printf("%08s", "woo")                   // correct
	fmt.Printf("% 8s", "woo")                   // correct
	fmt.Printf("%.*d", 3, 3)                    // correct
	fmt.Printf("%.*d x", 3, 3, 3, 3)            // ERROR "Printf call needs 2 args but has 4 args"
	fmt.Printf("%.*d x", "hi", 3)               // ERROR "Printf format %.*d uses non-int \x22hi\x22 as argument of \*"
	fmt.Printf("%.*d x", i, 3)                  // correct
	fmt.Printf("%.*d x", s, 3)                  // ERROR "Printf format %.\*d uses non-int s as argument of \*"
	fmt.Printf("%*% x", 0.22)                   // ERROR "Printf format %\*% uses non-int 0.22 as argument of \*"
	fmt.Printf("%q %q", multi()...)             // ok
	fmt.Printf("%#q", `blah`)                   // ok
	printf("now is the time", "buddy")          // ERROR "printf call has arguments but no formatting directives"
	Printf("now is the time", "buddy")          // ERROR "Printf call has arguments but no formatting directives"
	Printf("hi")                                // ok
	const format = "%s %s\n"
	Printf(format, "hi", "there")
	Printf(format, "hi")              // ERROR "Printf format %s reads arg #2, but call has only 1 arg$"
	Printf("%s %d %.3v %q", "str", 4) // ERROR "Printf format %.3v reads arg #3, but call has only 2 args"
	f := new(stringer)
	f.Warn(0, "%s", "hello", 3)  // ERROR "Warn call has possible formatting directive %s"
	f.Warnf(0, "%s", "hello", 3) // ERROR "Warnf call needs 1 arg but has 2 args"
	f.Warnf(0, "%r", "hello")    // ERROR "Warnf format %r has unknown verb r"
	f.Warnf(0, "%#s", "hello")   // ERROR "Warnf format %#s has unrecognized flag #"
	Printf("d%", 2)              // ERROR "Printf format % is missing verb at end of string"
	Printf("%d", percentDV)
	Printf("%d", &percentDV)
	Printf("%d", notPercentDV)  // ERROR "Printf format %d has arg notPercentDV of wrong type testdata.notPercentDStruct"
	Printf("%d", &notPercentDV) // ERROR "Printf format %d has arg &notPercentDV of wrong type \*testdata.notPercentDStruct"
	Printf("%p", &notPercentDV) // Works regardless: we print it as a pointer.
	Printf("%s", percentSV)
	Printf("%s", &percentSV)
	// Good argument reorderings.
	Printf("%[1]d", 3)
	Printf("%[1]*d", 3, 1)
	Printf("%[2]*[1]d", 1, 3)
	Printf("%[2]*.[1]*[3]d", 2, 3, 4)
	fmt.Fprintf(os.Stderr, "%[2]*.[1]*[3]d", 2, 3, 4) // Use Fprintf to make sure we count arguments correctly.
	// Bad argument reorderings.
	Printf("%[xd", 3)                      // ERROR "Printf format %\[xd is missing closing \]"
	Printf("%[x]d x", 3)                   // ERROR "Printf format has invalid argument index \[x\]"
	Printf("%[3]*s x", "hi", 2)            // ERROR "Printf format has invalid argument index \[3\]"
	_ = fmt.Sprintf("%[3]d x", 2)          // ERROR "Sprintf format has invalid argument index \[3\]"
	Printf("%[2]*.[1]*[3]d x", 2, "hi", 4) // ERROR "Printf format %\[2]\*\.\[1\]\*\[3\]d uses non-int \x22hi\x22 as argument of \*"
	Printf("%[0]s x", "arg1")              // ERROR "Printf format has invalid argument index \[0\]"
	Printf("%[0]d x", 1)                   // ERROR "Printf format has invalid argument index \[0\]"
	// Something that satisfies the error interface.
	var e error
	fmt.Println(e.Error()) // ok
	// Something that looks like an error interface but isn't, such as the (*T).Error method
	// in the testing package.
	var et1 errorTest1
	fmt.Println(et1.Error())        // ok
	fmt.Println(et1.Error("hi"))    // ok
	fmt.Println(et1.Error("%d", 3)) // ERROR "Error call has possible formatting directive %d"
	var et2 errorTest2
	et2.Error()        // ok
	et2.Error("hi")    // ok, not an error method.
	et2.Error("%d", 3) // ERROR "Error call has possible formatting directive %d"
	var et3 errorTest3
	et3.Error() // ok, not an error method.
	var et4 errorTest4
	et4.Error() // ok, not an error method.
	var et5 errorTest5
	et5.error() // ok, not an error method.
	// Interfaces can be used with any verb.
	var iface interface {
		ToTheMadness() bool // Method ToTheMadness usually returns false
	}
	fmt.Printf("%f", iface) // ok: fmt treats interfaces as transparent and iface may well have a float concrete type
	// Can't print a function.
	Printf("%d", someFunction) // ERROR "Printf format %d arg someFunction is a func value, not called"
	Printf("%v", someFunction) // ERROR "Printf format %v arg someFunction is a func value, not called"
	Println(someFunction)      // ERROR "Println arg someFunction is a func value, not called"
	Printf("%p", someFunction) // ok: maybe someone wants to see the pointer
	Printf("%T", someFunction) // ok: maybe someone wants to see the type
	// Bug: used to recur forever.
	Printf("%p %x", recursiveStructV, recursiveStructV.next)
	Printf("%p %x", recursiveStruct1V, recursiveStruct1V.next)
	Printf("%p %x", recursiveSliceV, recursiveSliceV)
	Printf("%p %x", recursiveMapV, recursiveMapV)
	// Special handling for Log.
	math.Log(3)  // OK
	Log(3)       // OK
	Log("%d", 3) // ERROR "Log call has possible formatting directive %d"
	Logf("%d", 3)
	Logf("%d", "hi") // ERROR "Logf format %d has arg \x22hi\x22 of wrong type string"

	Errorf(1, "%d", 3)    // OK
	Errorf(1, "%d", "hi") // ERROR "Errorf format %d has arg \x22hi\x22 of wrong type string"

	// Multiple string arguments before variadic args
	errorf("WARNING", "foobar")            // OK
	errorf("INFO", "s=%s, n=%d", "foo", 1) // OK
	errorf("ERROR", "%d")                  // ERROR "errorf format %d reads arg #1, but call has only 0 args"

	// Printf from external package
	externalprintf.Printf("%d", 42) // OK
	externalprintf.Printf("foobar") // OK
	level := 123
	externalprintf.Logf(level, "%d", 42)                        // OK
	externalprintf.Errorf(level, level, "foo %q bar", "foobar") // OK
	externalprintf.Logf(level, "%d")                            // ERROR "Logf format %d reads arg #1, but call has only 0 args"
	var formatStr = "%s %s"
	externalprintf.Sprintf(formatStr, "a", "b")     // OK
	externalprintf.Logf(level, formatStr, "a", "b") // OK

	// user-defined Println-like functions
	ss := &someStruct{}
	ss.Log(someFunction, "foo")          // OK
	ss.Error(someFunction, someFunction) // OK
	ss.Println()                         // OK
	ss.Println(1.234, "foo")             // OK
	ss.Println(1, someFunction)          // ERROR "Println arg someFunction is a func value, not called"
	ss.log(someFunction)                 // OK
	ss.log(someFunction, "bar", 1.33)    // OK
	ss.log(someFunction, someFunction)   // ERROR "log arg someFunction is a func value, not called"

	// indexed arguments
	Printf("%d %[3]d %d %[2]d x", 1, 2, 3, 4)             // OK
	Printf("%d %[0]d %d %[2]d x", 1, 2, 3, 4)             // ERROR "Printf format has invalid argument index \[0\]"
	Printf("%d %[3]d %d %[-2]d x", 1, 2, 3, 4)            // ERROR "Printf format has invalid argument index \[-2\]"
	Printf("%d %[3]d %d %[2234234234234]d x", 1, 2, 3, 4) // ERROR "Printf format has invalid argument index \[2234234234234\]"
	Printf("%d %[3]d %-10d %[2]d x", 1, 2, 3)             // ERROR "Printf format %-10d reads arg #4, but call has only 3 args"
	Printf("%d %[3]d %d %[2]d x", 1, 2, 3, 4, 5)          // ERROR "Printf call needs 4 args but has 5 args"
	Printf("%[1][3]d x", 1, 2)                            // ERROR "Printf format %\[1\]\[ has unknown verb \["

	// wrote Println but meant Fprintln
	Printf("%p\n", os.Stdout)   // OK
	Println(os.Stdout, "hello") // ERROR "Println does not take io.Writer but has first arg os.Stdout"

	Printf(someString(), "hello") // OK

}

func someString() string

type someStruct struct{}

// Log is non-variadic user-define Println-like function.
// Calls to this func must be skipped when checking
// for Println-like arguments.
func (ss *someStruct) Log(f func(), s string) {}

// Error is variadic user-define Println-like function.
// Calls to this func mustn't be checked for Println-like arguments,
// since variadic arguments type isn't interface{}.
func (ss *someStruct) Error(args ...func()) {}

// Println is variadic user-defined Println-like function.
// Calls to this func must be checked for Println-like arguments.
func (ss *someStruct) Println(args ...interface{}) {}

// log is variadic user-defined Println-like function.
// Calls to this func must be checked for Println-like arguments.
func (ss *someStruct) log(f func(), args ...interface{}) {}

// A function we use as a function value; it has no other purpose.
func someFunction() {}

// Printf is used by the test so we must declare it.
func Printf(format string, args ...interface{}) {
	panic("don't call - testing only")
}

// Println is used by the test so we must declare it.
func Println(args ...interface{}) {
	panic("don't call - testing only")
}

// Logf is used by the test so we must declare it.
func Logf(format string, args ...interface{}) {
	panic("don't call - testing only")
}

// Log is used by the test so we must declare it.
func Log(args ...interface{}) {
	panic("don't call - testing only")
}

// printf is used by the test so we must declare it.
func printf(format string, args ...interface{}) {
	panic("don't call - testing only")
}

// Errorf is used by the test for a case in which the first parameter
// is not a format string.
func Errorf(i int, format string, args ...interface{}) {
	panic("don't call - testing only")
}

// errorf is used by the test for a case in which the function accepts multiple
// string parameters before variadic arguments
func errorf(level, format string, args ...interface{}) {
	panic("don't call - testing only")
}

// multi is used by the test.
func multi() []interface{} {
	panic("don't call - testing only")
}

type stringer float64

var stringerv stringer

func (*stringer) String() string {
	return "string"
}

func (*stringer) Warn(int, ...interface{}) string {
	return "warn"
}

func (*stringer) Warnf(int, string, ...interface{}) string {
	return "warnf"
}

type embeddedStringer struct {
	foo string
	stringer
	bar int
}

var embeddedStringerv embeddedStringer

type notstringer struct {
	f float64
}

var notstringerv notstringer

type stringerarray [4]float64

func (stringerarray) String() string {
	return "string"
}

var stringerarrayv stringerarray

type notstringerarray [4]float64

var notstringerarrayv notstringerarray

var nonemptyinterface = interface {
	f()
}(nil)

// A data type we can print with "%d".
type percentDStruct struct {
	a int
	b []byte
	c *float64
}

var percentDV percentDStruct

// A data type we cannot print correctly with "%d".
type notPercentDStruct struct {
	a int
	b []byte
	c bool
}

var notPercentDV notPercentDStruct

// A data type we can print with "%s".
type percentSStruct struct {
	a string
	b []byte
	C stringerarray
}

var percentSV percentSStruct

type recursiveStringer int

func (s recursiveStringer) String() string {
	_ = fmt.Sprintf("%d", s)
	_ = fmt.Sprintf("%#v", s)
	_ = fmt.Sprintf("%v", s)  // ERROR "Sprintf format %v with arg s causes recursive String method call"
	_ = fmt.Sprintf("%v", &s) // ERROR "Sprintf format %v with arg &s causes recursive String method call"
	_ = fmt.Sprintf("%T", s)  // ok; does not recursively call String
	return fmt.Sprintln(s)    // ERROR "Sprintln arg s causes recursive call to String method"
}

type recursivePtrStringer int

func (p *recursivePtrStringer) String() string {
	_ = fmt.Sprintf("%v", *p)
	return fmt.Sprintln(p) // ERROR "Sprintln arg p causes recursive call to String method"
}

type Formatter bool

func (*Formatter) Format(fmt.State, rune) {
}

// Formatter with value receiver
type FormatterVal bool

func (FormatterVal) Format(fmt.State, rune) {
}

type RecursiveSlice []RecursiveSlice

var recursiveSliceV = &RecursiveSlice{}

type RecursiveMap map[int]RecursiveMap

var recursiveMapV = make(RecursiveMap)

type RecursiveStruct struct {
	next *RecursiveStruct
}

var recursiveStructV = &RecursiveStruct{}

type RecursiveStruct1 struct {
	next *Recursive2Struct
}

type RecursiveStruct2 struct {
	next *Recursive1Struct
}

var recursiveStruct1V = &RecursiveStruct1{}

// Fix for issue 7149: Missing return type on String method caused fault.
func (int) String() {
	return ""
}

func (s *unknownStruct) Fprintln(w io.Writer, s string) {}

func UnknownStructFprintln() {
	s := unknownStruct{}
	s.Fprintln(os.Stdout, "hello, world!") // OK
}

// Issue 17798: unexported stringer cannot be formatted.
type unexportedStringer struct {
	t stringer
}
type unexportedStringerOtherFields struct {
	s string
	t stringer
	S string
}

// Issue 17798: unexported error cannot be formatted.
type unexportedError struct {
	e error
}
type unexportedErrorOtherFields struct {
	s string
	e error
	S string
}

type errorer struct{}

func (e errorer) Error() string { return "errorer" }

func UnexportedStringerOrError() {
	us := unexportedStringer{}
	fmt.Printf("%s", us)  // ERROR "Printf format %s has arg us of wrong type testdata.unexportedStringer"
	fmt.Printf("%s", &us) // ERROR "Printf format %s has arg &us of wrong type [*]testdata.unexportedStringer"

	usf := unexportedStringerOtherFields{
		s: "foo",
		S: "bar",
	}
	fmt.Printf("%s", usf)  // ERROR "Printf format %s has arg usf of wrong type testdata.unexportedStringerOtherFields"
	fmt.Printf("%s", &usf) // ERROR "Printf format %s has arg &usf of wrong type [*]testdata.unexportedStringerOtherFields"

	ue := unexportedError{
		e: &errorer{},
	}
	fmt.Printf("%s", ue)  // ERROR "Printf format %s has arg ue of wrong type testdata.unexportedError"
	fmt.Printf("%s", &ue) // ERROR "Printf format %s has arg &ue of wrong type [*]testdata.unexportedError"

	uef := unexportedErrorOtherFields{
		s: "foo",
		e: &errorer{},
		S: "bar",
	}
	fmt.Printf("%s", uef)  // ERROR "Printf format %s has arg uef of wrong type testdata.unexportedErrorOtherFields"
	fmt.Printf("%s", &uef) // ERROR "Printf format %s has arg &uef of wrong type [*]testdata.unexportedErrorOtherFields"
}
