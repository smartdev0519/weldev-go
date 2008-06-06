// $G $D/$F.go && $L $F.$A && ./$A.out

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.


package main

func fun(ia,ib,ic int)int;

func
main()
{
	var x int;

	x = fun(10,20,30);
	if x != 60 { panic x; }
}

func
fun(ia,ib,ic int)int
{
	var o int;

	o = ia+ib+ic;
	if o != 60 { panic o; }
	return o;
}
