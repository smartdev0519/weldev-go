// $G $F.go && $L $F.$A && ./$A.out

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.


package main

func assertequal(is, shouldbe int, msg string) {
	if is != shouldbe {
		print "assertion fail", msg, "\n";
		panic 1;
	}
}

func f1() {
}

func f2(a int) {
}

func f3(a, b int) int {
	return a+b;
}

func f4(a, b int, c float) int {
	return (a+b)/2 + int(c);
}

func f5(a int) int {
	return 5;
}

func f6(a int) (r int) {
	return 6;
}

func f7(a int) (x int, y float) {
	return 7, 7.0;
}


func f8(a int) (x int, y float) {
	return 8, 8.0;
}

type T struct {
	x, y int;
}

func (t *T) m10(a int, b float) int {
	return (t.x+a) * (t.y+int(b));
}


//BUG func f9(a int) (i int, f float) {  // multiple returns not ready yet
// BUG funny return value
//	var i float = 9;
//	var f float = 9.0;
//	return i, f;
// return
// }


func main() {
	f1();
	f2(1);
	r3 := f3(1, 2);
	assertequal(r3, 3, "3");
	r4 := f4(0, 2, 3.0);
	assertequal(r4, 4, "4");
	r5 := f5(1);
	assertequal(r5, 5, "5");
	r6 := f6(1);
	assertequal(r6, 6, "6");
	var r7 int;
	var s7 float;
	//BUG r7, s7 = f7(1);
	//BUG assertequal(r7, 7, "r7");
	//BUG assertequal(int(s7), 7, "s7");
	var r8 int;
	var s8 float;
	//BUG r8, s8 = f8(1);
	//BUG assertequal(r8, 8, "r8");
	//BUG assertequal(int(s8), 8, "s8");
	var r9 int;
	var s9 float;
	//BUG r9, s9 = f9(1);
	//BUG assertequal(r9, 9, "r9");
	//BUG assertequal(int(s9), 9, "s9");
	var t *T = new(T);
	t.x = 1;
	t.y = 2;
	r10 := t.m10(1, 3.0);
	assertequal(r10, 10, "10");
}
