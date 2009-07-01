// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gob

import (
	"io";
	"math";
	"os";
	"unsafe";
)

// The global execution state of an instance of the decoder.
type DecState struct {
	r	io.Reader;
	err	os.Error;
	base	uintptr;
	buf [1]byte;	// buffer used by the decoder; here to avoid allocation.
}

// DecodeUint reads an encoded unsigned integer from state.r.
// Sets state.err.  If state.err is already non-nil, it does nothing.
func DecodeUint(state *DecState) (x uint64) {
	if state.err != nil {
		return
	}
	for shift := uint(0);; shift += 7 {
		var n int;
		n, state.err = state.r.Read(&state.buf);
		if n != 1 {
			return 0
		}
		b := uint64(state.buf[0]);
		x |= b << shift;
		if b&0x80 != 0 {
			x &^= 0x80 << shift;
			break
		}
	}
	return x;
}

// DecodeInt reads an encoded signed integer from state.r.
// Sets state.err.  If state.err is already non-nil, it does nothing.
func DecodeInt(state *DecState) int64 {
	x := DecodeUint(state);
	if state.err != nil {
		return 0
	}
	if x & 1 != 0 {
		return ^int64(x>>1)
	}
	return int64(x >> 1)
}

// The 'instructions' of the decoding machine
type decInstr struct {
	op	func(i *decInstr, state *DecState);
	field		int;	// field number
	indir	int;	// how many pointer indirections to reach the value in the struct
	offset	uintptr;	// offset in the structure of the field to encode
}

// Since the encoder writes no zeros, if we arrive at a decoder we have
// a value to extract and store.  The field number has already been read
// (it's how we knew to call this decoder).
// Each decoder is responsible for handling any indirections associated
// with the data structure.  If any pointer so reached is nil, allocation must
// be done.

// Walk the pointer hierarchy, allocating if we find a nil.  Stop one before the end.
func decIndirect(p unsafe.Pointer, indir int) unsafe.Pointer {
	for ; indir > 1; indir-- {
		if *(*unsafe.Pointer)(p) == nil {
			// Allocation required
			*(*unsafe.Pointer)(p) = unsafe.Pointer(new(unsafe.Pointer));
		}
		p = *(*unsafe.Pointer)(p);
	}
	return p
}

func decBool(i *decInstr, state *DecState) {
	p := unsafe.Pointer(state.base+i.offset);
	if i.indir > 0 {
		if i.indir > 1 {
			p = decIndirect(p, i.indir);
		}
		if *(*unsafe.Pointer)(p) == nil {
			*(*unsafe.Pointer)(p) = unsafe.Pointer(new(int));
			p = *(*unsafe.Pointer)(p);
		}
	}
	v := int(DecodeInt(state));
	if state.err == nil {
		*(*bool)(p) = v != 0;
	}
}

func decInt(i *decInstr, state *DecState) {
	p := unsafe.Pointer(state.base+i.offset);
	if i.indir > 0 {
		if i.indir > 1 {
			p = decIndirect(p, i.indir);
		}
		if *(*unsafe.Pointer)(p) == nil {
			*(*unsafe.Pointer)(p) = unsafe.Pointer(new(int));
			p = *(*unsafe.Pointer)(p);
		}
	}
	v := int(DecodeInt(state));
	if state.err == nil {
		*(*int)(p) = v;
	}
}

func decUint(i *decInstr, state *DecState) {
	p := unsafe.Pointer(state.base+i.offset);
	if i.indir > 0 {
		if i.indir > 1 {
			p = decIndirect(p, i.indir);
		}
		if *(*unsafe.Pointer)(p) == nil {
			*(*unsafe.Pointer)(p) = unsafe.Pointer(new(int));
			p = *(*unsafe.Pointer)(p);
		}
	}
	v := uint(DecodeUint(state));
	if state.err == nil {
		*(*uint)(p) = v;
	}
}

func decInt8(i *decInstr, state *DecState) {
	p := unsafe.Pointer(state.base+i.offset);
	if i.indir > 0 {
		if i.indir > 1 {
			p = decIndirect(p, i.indir);
		}
		if *(*unsafe.Pointer)(p) == nil {
			*(*unsafe.Pointer)(p) = unsafe.Pointer(new(int));
			p = *(*unsafe.Pointer)(p);
		}
	}
	v := int8(DecodeInt(state));
	if state.err == nil {
		*(*int8)(p) = v;
	}
}

func decUint8(i *decInstr, state *DecState) {
	p := unsafe.Pointer(state.base+i.offset);
	if i.indir > 0 {
		if i.indir > 1 {
			p = decIndirect(p, i.indir);
		}
		if *(*unsafe.Pointer)(p) == nil {
			*(*unsafe.Pointer)(p) = unsafe.Pointer(new(int));
			p = *(*unsafe.Pointer)(p);
		}
	}
	v := uint8(DecodeUint(state));
	if state.err == nil {
		*(*uint8)(p) = v;
	}
}

func decInt16(i *decInstr, state *DecState) {
	p := unsafe.Pointer(state.base+i.offset);
	if i.indir > 0 {
		if i.indir > 1 {
			p = decIndirect(p, i.indir);
		}
		if *(*unsafe.Pointer)(p) == nil {
			*(*unsafe.Pointer)(p) = unsafe.Pointer(new(int));
			p = *(*unsafe.Pointer)(p);
		}
	}
	v := int16(DecodeInt(state));
	if state.err == nil {
		*(*int16)(p) = v;
	}
}

func decUint16(i *decInstr, state *DecState) {
	p := unsafe.Pointer(state.base+i.offset);
	if i.indir > 0 {
		if i.indir > 1 {
			p = decIndirect(p, i.indir);
		}
		if *(*unsafe.Pointer)(p) == nil {
			*(*unsafe.Pointer)(p) = unsafe.Pointer(new(int));
			p = *(*unsafe.Pointer)(p);
		}
	}
	v := uint16(DecodeUint(state));
	if state.err == nil {
		*(*uint16)(p) = v;
	}
}

func decInt32(i *decInstr, state *DecState) {
	p := unsafe.Pointer(state.base+i.offset);
	if i.indir > 0 {
		if i.indir > 1 {
			p = decIndirect(p, i.indir);
		}
		if *(*unsafe.Pointer)(p) == nil {
			*(*unsafe.Pointer)(p) = unsafe.Pointer(new(int));
			p = *(*unsafe.Pointer)(p);
		}
	}
	v := int32(DecodeInt(state));
	if state.err == nil {
		*(*int32)(p) = v;
	}
}

func decUint32(i *decInstr, state *DecState) {
	p := unsafe.Pointer(state.base+i.offset);
	if i.indir > 0 {
		if i.indir > 1 {
			p = decIndirect(p, i.indir);
		}
		if *(*unsafe.Pointer)(p) == nil {
			*(*unsafe.Pointer)(p) = unsafe.Pointer(new(int));
			p = *(*unsafe.Pointer)(p);
		}
	}
	v := uint32(DecodeUint(state));
	if state.err == nil {
		*(*uint32)(p) = v;
	}
}

func decInt64(i *decInstr, state *DecState) {
	p := unsafe.Pointer(state.base+i.offset);
	if i.indir > 0 {
		if i.indir > 1 {
			p = decIndirect(p, i.indir);
		}
		if *(*unsafe.Pointer)(p) == nil {
			*(*unsafe.Pointer)(p) = unsafe.Pointer(new(int));
			p = *(*unsafe.Pointer)(p);
		}
	}
	v := int64(DecodeInt(state));
	if state.err == nil {
		*(*int64)(p) = v;
	}
}

func decUint64(i *decInstr, state *DecState) {
	p := unsafe.Pointer(state.base+i.offset);
	if i.indir > 0 {
		if i.indir > 1 {
			p = decIndirect(p, i.indir);
		}
		if *(*unsafe.Pointer)(p) == nil {
			*(*unsafe.Pointer)(p) = unsafe.Pointer(new(int));
			p = *(*unsafe.Pointer)(p);
		}
	}
	v := uint64(DecodeUint(state));
	if state.err == nil {
		*(*uint64)(p) = v;
	}
}

// Floating-point numbers are transmitted as uint64s holding the bits
// of the underlying representation.  They are sent byte-reversed, with
// the exponent end coming out first, so integer floating point numbers
// (for example) transmit more compactly.  This routine does the
// unswizzling.
func floatFromBits(u uint64) float64 {
	var v uint64;
	for i := 0; i < 8; i++ {
		v <<= 8;
		v |= u & 0xFF;
		u >>= 8;
	}
	return math.Float64frombits(v);
}

func decFloat(i *decInstr, state *DecState) {
	p := unsafe.Pointer(state.base+i.offset);
	if i.indir > 0 {
		if i.indir > 1 {
			p = decIndirect(p, i.indir);
		}
		if *(*unsafe.Pointer)(p) == nil {
			*(*unsafe.Pointer)(p) = unsafe.Pointer(new(int));
			p = *(*unsafe.Pointer)(p);
		}
	}
	v := float(floatFromBits(uint64(DecodeUint(state))));
	if state.err == nil {
		*(*float)(p) = v;
	}
}

func decFloat32(i *decInstr, state *DecState) {
	p := unsafe.Pointer(state.base+i.offset);
	if i.indir > 0 {
		if i.indir > 1 {
			p = decIndirect(p, i.indir);
		}
		if *(*unsafe.Pointer)(p) == nil {
			*(*unsafe.Pointer)(p) = unsafe.Pointer(new(int));
			p = *(*unsafe.Pointer)(p);
		}
	}
	v := float32(floatFromBits(uint64(DecodeUint(state))));
	if state.err == nil {
		*(*float32)(p) = v;
	}
}

func decFloat64(i *decInstr, state *DecState) {
	p := unsafe.Pointer(state.base+i.offset);
	if i.indir > 0 {
		if i.indir > 1 {
			p = decIndirect(p, i.indir);
		}
		if *(*unsafe.Pointer)(p) == nil {
			*(*unsafe.Pointer)(p) = unsafe.Pointer(new(int));
			p = *(*unsafe.Pointer)(p);
		}
	}
	v := floatFromBits(uint64(DecodeUint(state)));
	if state.err == nil {
		*(*float64)(p) = v;
	}
}
