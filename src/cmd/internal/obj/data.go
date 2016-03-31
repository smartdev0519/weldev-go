// Derived from Inferno utils/6l/obj.c and utils/6l/span.c
// http://code.google.com/p/inferno-os/source/browse/utils/6l/obj.c
// http://code.google.com/p/inferno-os/source/browse/utils/6l/span.c
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
//	Portions Copyright © 2009 The Go Authors.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package obj

import (
	"log"
	"math"
)

// Grow increases the length of s.P to lsiz.
func (s *LSym) Grow(lsiz int64) {
	siz := int(lsiz)
	if int64(siz) != lsiz {
		log.Fatalf("LSym.Grow size %d too long", lsiz)
	}
	if len(s.P) >= siz {
		return
	}
	// TODO(dfc) append cap-len at once, rather than
	// one byte at a time.
	for cap(s.P) < siz {
		s.P = append(s.P[:cap(s.P)], 0)
	}
	s.P = s.P[:siz]
}

// GrowCap increases the capacity of s.P to c.
func (s *LSym) GrowCap(c int64) {
	if int64(cap(s.P)) >= c {
		return
	}
	if s.P == nil {
		s.P = make([]byte, 0, c)
		return
	}
	b := make([]byte, len(s.P), c)
	copy(b, s.P)
	s.P = b
}

// prepwrite prepares to write data of size siz into s at offset off.
func (s *LSym) prepwrite(ctxt *Link, off int64, siz int) {
	if off < 0 || siz < 0 || off >= 1<<30 {
		log.Fatalf("prepwrite: bad off=%d siz=%d", off, siz)
	}
	if s.Type == SBSS || s.Type == STLSBSS {
		ctxt.Diag("cannot supply data for BSS var")
	}
	l := off + int64(siz)
	s.Grow(l)
	if l > s.Size {
		s.Size = l
	}
}

// WriteFloat32 writes f into s at offset off.
func (s *LSym) WriteFloat32(ctxt *Link, off int64, f float32) {
	s.prepwrite(ctxt, off, 4)
	ctxt.Arch.ByteOrder.PutUint32(s.P[off:], math.Float32bits(f))
}

// WriteFloat64 writes f into s at offset off.
func (s *LSym) WriteFloat64(ctxt *Link, off int64, f float64) {
	s.prepwrite(ctxt, off, 8)
	ctxt.Arch.ByteOrder.PutUint64(s.P[off:], math.Float64bits(f))
}

// WriteInt writes an integer i of size siz into s at offset off.
func (s *LSym) WriteInt(ctxt *Link, off int64, siz int, i int64) {
	s.prepwrite(ctxt, off, siz)
	switch siz {
	default:
		ctxt.Diag("WriteInt: bad integer size: %d", siz)
	case 1:
		s.P[off] = byte(i)
	case 2:
		ctxt.Arch.ByteOrder.PutUint16(s.P[off:], uint16(i))
	case 4:
		ctxt.Arch.ByteOrder.PutUint32(s.P[off:], uint32(i))
	case 8:
		ctxt.Arch.ByteOrder.PutUint64(s.P[off:], uint64(i))
	}
}

// WriteAddr writes an address of size siz into s at offset off.
// rsym and roff specify the relocation for the address.
func (s *LSym) WriteAddr(ctxt *Link, off int64, siz int, rsym *LSym, roff int64) {
	if siz != ctxt.Arch.PtrSize {
		ctxt.Diag("WriteAddr: bad address size %d in %s", siz, s.Name)
	}
	s.prepwrite(ctxt, off, siz)
	r := Addrel(s)
	r.Off = int32(off)
	if int64(r.Off) != off {
		ctxt.Diag("WriteAddr: off overflow %d in %s", off, s.Name)
	}
	r.Siz = uint8(siz)
	r.Sym = rsym
	r.Type = R_ADDR
	r.Add = roff
}

// WriteOff writes a 4 byte offset to rsym+roff into s at offset off.
// After linking the 4 bytes stored at s+off will be
// rsym+roff-(start of section that s is in).
func (s *LSym) WriteOff(ctxt *Link, off int64, rsym *LSym, roff int64) {
	s.prepwrite(ctxt, off, 4)
	r := Addrel(s)
	r.Off = int32(off)
	if int64(r.Off) != off {
		ctxt.Diag("WriteOff: off overflow %d in %s", off, s.Name)
	}
	r.Siz = 4
	r.Sym = rsym
	r.Type = R_ADDROFF
	r.Add = roff
}

// WriteString writes a string of size siz into s at offset off.
func (s *LSym) WriteString(ctxt *Link, off int64, siz int, str string) {
	if siz < len(str) {
		ctxt.Diag("WriteString: bad string size: %d < %d", siz, len(str))
	}
	s.prepwrite(ctxt, off, siz)
	copy(s.P[off:off+int64(siz)], str)
}

// WriteBytes writes a slice of bytes into s at offset off.
func (s *LSym) WriteBytes(ctxt *Link, off int64, b []byte) int64 {
	s.prepwrite(ctxt, off, len(b))
	copy(s.P[off:], b)
	return off + int64(len(b))
}

func Addrel(s *LSym) *Reloc {
	s.R = append(s.R, Reloc{})
	return &s.R[len(s.R)-1]
}

func Setuintxx(ctxt *Link, s *LSym, off int64, v uint64, wid int64) int64 {
	if s.Type == 0 {
		s.Type = SDATA
	}
	if s.Size < off+wid {
		s.Size = off + wid
		s.Grow(s.Size)
	}

	switch wid {
	case 1:
		s.P[off] = uint8(v)
	case 2:
		ctxt.Arch.ByteOrder.PutUint16(s.P[off:], uint16(v))
	case 4:
		ctxt.Arch.ByteOrder.PutUint32(s.P[off:], uint32(v))
	case 8:
		ctxt.Arch.ByteOrder.PutUint64(s.P[off:], uint64(v))
	}

	return off + wid
}
