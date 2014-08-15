// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime_test

import (
	"bytes"
	"runtime"
	"testing"
)

// TestGCInfo tests that various objects in heap, data and bss receive correct GC pointer type info.
func TestGCInfo(t *testing.T) {
	verifyGCInfo(t, "bss ScalarPtr", &bssScalarPtr, nonStackInfo(infoScalarPtr))
	verifyGCInfo(t, "bss PtrScalar", &bssPtrScalar, nonStackInfo(infoPtrScalar))
	verifyGCInfo(t, "bss Complex", &bssComplex, nonStackInfo(infoComplex()))
	verifyGCInfo(t, "bss string", &bssString, nonStackInfo(infoString))
	verifyGCInfo(t, "bss eface", &bssEface, nonStackInfo(infoEface))

	verifyGCInfo(t, "data ScalarPtr", &dataScalarPtr, nonStackInfo(infoScalarPtr))
	verifyGCInfo(t, "data PtrScalar", &dataPtrScalar, nonStackInfo(infoPtrScalar))
	verifyGCInfo(t, "data Complex", &dataComplex, nonStackInfo(infoComplex()))
	verifyGCInfo(t, "data string", &dataString, nonStackInfo(infoString))
	verifyGCInfo(t, "data eface", &dataEface, nonStackInfo(infoEface))

	verifyGCInfo(t, "stack ScalarPtr", new(ScalarPtr), infoScalarPtr)
	verifyGCInfo(t, "stack PtrScalar", new(PtrScalar), infoPtrScalar)
	verifyGCInfo(t, "stack Complex", new(Complex), infoComplex())
	verifyGCInfo(t, "stack string", new(string), infoString)
	verifyGCInfo(t, "stack eface", new(interface{}), infoEface)

	for i := 0; i < 3; i++ {
		verifyGCInfo(t, "heap ScalarPtr", escape(new(ScalarPtr)), nonStackInfo(infoScalarPtr))
		verifyGCInfo(t, "heap PtrScalar", escape(new(PtrScalar)), nonStackInfo(infoPtrScalar))
		verifyGCInfo(t, "heap Complex", escape(new(Complex)), nonStackInfo(infoComplex()))
		verifyGCInfo(t, "heap string", escape(new(string)), nonStackInfo(infoString))
		verifyGCInfo(t, "heap eface", escape(new(interface{})), nonStackInfo(infoEface))
	}

}

func verifyGCInfo(t *testing.T, name string, p interface{}, mask0 []byte) {
	mask := runtime.GCMask(p)
	if len(mask) > len(mask0) {
		mask0 = append(mask0, BitsDead)
		mask = mask[:len(mask0)]
	}
	if bytes.Compare(mask, mask0) != 0 {
		t.Errorf("bad GC program for %v:\nwant %+v\ngot  %+v", name, mask0, mask)
		return
	}
}

func nonStackInfo(mask []byte) []byte {
	// BitsDead is replaced with BitsScalar everywhere except stacks.
	mask1 := make([]byte, len(mask))
	mw := false
	for i, v := range mask {
		if !mw && v == BitsDead {
			v = BitsScalar
		}
		mw = !mw && v == BitsMultiWord
		mask1[i] = v
	}
	return mask1
}

var gcinfoSink interface{}

func escape(p interface{}) interface{} {
	gcinfoSink = p
	return p
}

const (
	BitsDead = iota
	BitsScalar
	BitsPointer
	BitsMultiWord
)

const (
	BitsString = iota
	BitsSlice
	BitsIface
	BitsEface
)

type ScalarPtr struct {
	q int
	w *int
	e int
	r *int
	t int
	y *int
}

var infoScalarPtr = []byte{BitsScalar, BitsPointer, BitsScalar, BitsPointer, BitsScalar, BitsPointer}

type PtrScalar struct {
	q *int
	w int
	e *int
	r int
	t *int
	y int
}

var infoPtrScalar = []byte{BitsPointer, BitsScalar, BitsPointer, BitsScalar, BitsPointer, BitsScalar}

type Complex struct {
	q *int
	w byte
	e [17]byte
	r []byte
	t int
	y uint16
	u uint64
	i string
}

func infoComplex() []byte {
	switch runtime.GOARCH {
	case "386", "arm":
		return []byte{
			BitsPointer, BitsScalar, BitsScalar, BitsScalar,
			BitsScalar, BitsScalar, BitsMultiWord, BitsSlice,
			BitsDead, BitsScalar, BitsScalar, BitsScalar,
			BitsScalar, BitsMultiWord, BitsString,
		}
	case "amd64":
		return []byte{
			BitsPointer, BitsScalar, BitsScalar, BitsScalar,
			BitsMultiWord, BitsSlice, BitsDead, BitsScalar,
			BitsScalar, BitsScalar, BitsMultiWord, BitsString,
		}
	case "amd64p32":
		return []byte{
			BitsPointer, BitsScalar, BitsScalar, BitsScalar,
			BitsScalar, BitsScalar, BitsMultiWord, BitsSlice,
			BitsDead, BitsScalar, BitsScalar, BitsDead,
			BitsScalar, BitsScalar, BitsMultiWord, BitsString,
		}
	default:
		panic("unknown arch")
	}
}

var (
	// BSS
	bssScalarPtr ScalarPtr
	bssPtrScalar PtrScalar
	bssComplex   Complex
	bssString    string
	bssEface     interface{}

	// DATA
	dataScalarPtr             = ScalarPtr{q: 1}
	dataPtrScalar             = PtrScalar{w: 1}
	dataComplex               = Complex{w: 1}
	dataString                = "foo"
	dataEface     interface{} = 42

	infoString = []byte{BitsMultiWord, BitsString}
	infoEface  = []byte{BitsMultiWord, BitsEface}
)
