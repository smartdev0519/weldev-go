package main

import (
	p0 "./bug0"
	p1 "./bug1"

	"reflect"
	"strings"
)

var v0 p0.T
var v1 p1.T

type I0 interface {
	M(p0.T)
}

type I1 interface {
	M(p1.T)
}

type t0 int

func (t0) M(p0.T) {}

type t1 float

func (t1) M(p1.T) {}

var i0 I0 = t0(0) // ok
var i1 I1 = t1(0) // ok

var p0i p0.I = t0(0) // ok
var p1i p1.I = t1(0) // ok

func main() {
	// check that reflect paths are correct,
	// meaning that reflect data for v0, v1 didn't get confused.

	// path is full (rooted) path name.  check suffix for gc, prefix for gccgo
	if s := reflect.Typeof(v0).PkgPath(); !strings.HasSuffix(s, "/bug0") && !strings.HasPrefix(s, "bug0") {
		panicln("bad v0 path", len(s), s)
	}
	if s := reflect.Typeof(v1).PkgPath(); !strings.HasSuffix(s, "/bug1") && !strings.HasPrefix(s, "bug1") {
		panicln("bad v1 path", s)
	}

	// check that dynamic interface check doesn't get confused
	var i interface{} = t0(0)
	if _, ok := i.(I1); ok {
		panicln("used t0 as i1")
	}
	if _, ok := i.(p1.I); ok {
		panicln("used t0 as p1.I")
	}

	i = t1(1)
	if _, ok := i.(I0); ok {
		panicln("used t1 as i0")
	}
	if _, ok := i.(p0.I); ok {
		panicln("used t1 as p0.I")
	}

	// check that type switch works.
	// the worry is that if p0.T and p1.T have the same hash,
	// the binary search will handle one of them incorrectly.
	for j := 0; j < 3; j++ {
		switch j {
		case 0:
			i = p0.T{}
		case 1:
			i = p1.T{}
		case 2:
			i = 3.14
		}
		switch k := i.(type) {
		case p0.T:
			if j != 0 {
				panicln("type switch p0.T")
			}
		case p1.T:
			if j != 1 {
				panicln("type switch p1.T")
			}
		default:
			if j != 2 {
				panicln("type switch default", j)
			}
		}
	}
}
