// errchk $G $D/$F.go

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

type Inst interface {
	Next()	*Inst;
}

type Regexp struct {
	code *[]Inst;
	start	Inst;
}

type Start struct {
	foo	*Inst;
}

func (start *Start) Next() *Inst { return nil }


func AddInst(Inst) *Inst {
	print("ok in addinst\n");
	return nil
}

func main() {
	re := new(Regexp);
	print("call addinst\n");
	var x Inst = AddInst(new(Start));
	print("return from  addinst\n");
}
