// $G $D/$F.go && $L $F.$A && ./$A.out

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Random malloc test.

package main

import (
	"flag";
	"malloc";
	"rand";
	"unsafe";
)

var chatty = flag.Bool("v", false, "chatty");

var footprint uint64;
var allocated uint64;
func bigger() {
	if f := malloc.GetStats().Sys; footprint < f {
		footprint = f;
		if *chatty {
			println("Footprint", footprint, " for ", allocated);
		}
		if footprint > 1e9 {
			panicln("too big");
		}
	}
}

// Prime the data structures by allocating one of
// each block in order.  After this, there should be
// little reason to ask for more memory from the OS.
func prime() {
	for i := 0; i < 16; i++ {
		b := malloc.Alloc(1<<uint(i));
		malloc.Free(b);
	}
	for i := uint64(0); i < 256; i++ {
		b := malloc.Alloc(i<<12);
		malloc.Free(b);
	}
}

func memset(b *byte, c byte, n uint64) {
	np := uintptr(n);
	for i := uintptr(0); i < np; i++ {
		*(b.(unsafe.Pointer).(uintptr)+i).(unsafe.Pointer).(*byte) = c;
	}
}

func main() {
	flag.Parse();
//	prime();
	var blocks [1] struct { base *byte; siz uint64; };
	for i := 0; i < 1<<12; i++ {
		if i%(1<<10) == 0 && *chatty {
			println(i);
		}
		b := rand.Int() % len(blocks);
		if blocks[b].base != nil {
		//	println("Free", blocks[b].siz, blocks[b].base);
			malloc.Free(blocks[b].base);
			blocks[b].base = nil;
			allocated -= blocks[b].siz;
			continue
		}
		siz := uint64(rand.Int() >> (11 + rand.Uint32() % 20));
		base := malloc.Alloc(siz);
	//	ptr := uint64(syscall.BytePtr(base))+uint64(siz/2);
	//	obj, size, ref, ok := allocator.find(ptr);
	//	if obj != base || *ref != 0 || !ok {
	//		panicln("find", siz, obj, ref, ok);
	//	}
		blocks[b].base = base;
		blocks[b].siz = siz;
		allocated += siz;
	//	println("Alloc", siz, base);
		memset(base, 0xbb, siz);
		bigger();
	}
}
