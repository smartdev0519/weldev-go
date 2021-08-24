// run -gcflags=-G=3

// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "reflect"

func f[T any](i interface{}) {
	switch x := i.(type) {
	case T:
		println("T", x)
	case int:
		println("int", x)
	case int32, int16:
		println("int32/int16", reflect.ValueOf(x).Int())
	case struct { a, b T }:
		println("struct{T,T}", x.a, x.b)
	default:
		println("other", reflect.ValueOf(x).Int())
	}
}
func main() {
	f[float64](float64(6))
	f[float64](int(7))
	f[float64](int32(8))
	f[float64](struct{a, b float64}{a:1, b:2})
	f[float64](int8(9))
}
