// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// errchk $G $F.go

package main

type T int
func (t T) V()
func (t *T) P()

type V interface { V() }
type P interface { P(); V() }

type S struct { T; }
type SP struct { *T; }

func main() {
	var t T;
	var v V;
	var p P;
	var s S;
	var sp SP;

	v = t;
	p = t;	// ERROR "is not"
	v = &t;
	p = &t;

	v = s;
	p = s;	// ERROR "is not"
	v = &s;
	p = &s;

	v = sp;
	p = sp;	// no error!
	v = &sp;
	p = &sp;
}

