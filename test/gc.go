// $G $F.go && $L $F.$A && ./$A.out

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "malloc"

func mk2() {
	b := new([10000]byte);
//	println(b, "stored at", &b);
}

func mk1() {
	mk2();
}

func main() {
	for i := 0; i < 10; i++ {
		mk1();
		malloc.GC();
	}
}
