// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import "unsafe"

// adjust Gobuf as if it executed a call to fn with context ctxt
// and then did an immediate Gosave.
func gostartcall(buf *gobuf, fn, ctxt unsafe.Pointer) {
	if buf.lr != 0 {
		throw("invalid use of gostartcall")
	}
	buf.lr = buf.pc
	buf.pc = uintptr(fn)
	buf.ctxt = ctxt
}

// Called to rewind context saved during morestack back to beginning of function.
// To help us, the linker emits a jmp back to the beginning right after the
// call to morestack. We just have to decode and apply that jump.
func rewindmorestack(buf *gobuf) {
	var inst uint32
	if buf.pc&3 == 0 && buf.pc != 0 {
		inst = *(*uint32)(unsafe.Pointer(buf.pc))
		if inst>>24 == 0x9a || inst>>24 == 0xea {
			buf.pc += uintptr(int32(inst<<8)>>6) + 8
			return
		}
	}

	print("runtime: pc=", hex(buf.pc), " ", hex(inst), "\n")
	throw("runtime: misuse of rewindmorestack")
}

// for testing
func usplit(x uint32) (q, r uint32)
