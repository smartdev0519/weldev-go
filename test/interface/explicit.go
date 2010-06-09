// errchk $G $D/$F.go

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Static error messages about interface conversions.

package main

type T struct {
	a int
}

var t *T

type I interface {
	M()
}

var i I

type I2 interface {
	M()
	N()
}

var i2 I2

type E interface{}

var e E

func main() {
	e = t // ok
	t = e // ERROR "need explicit|need type assertion"

	// neither of these can work,
	// because i has an extra method
	// that t does not, so i cannot contain a t.
	i = t // ERROR "incompatible|missing M method"
	t = i // ERROR "incompatible|need type assertion"

	i = i2 // ok
	i2 = i // ERROR "missing N method"

	i = I(i2)  // ok
	i2 = I2(i) // ERROR "missing N method"

	e = E(t) // ok
	t = T(e) // ERROR "need explicit|need type assertion|incompatible"
}
