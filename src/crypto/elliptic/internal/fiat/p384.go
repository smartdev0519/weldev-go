// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by generate.go. DO NOT EDIT.

package fiat

import (
	"crypto/subtle"
	"errors"
)

// P384Element is an integer modulo 2^384 - 2^128 - 2^96 + 2^32 - 1.
//
// The zero value is a valid zero element.
type P384Element struct {
	// Values are represented internally always in the Montgomery domain, and
	// converted in Bytes and SetBytes.
	x p384MontgomeryDomainFieldElement
}

const p384ElementLen = 48

type p384UntypedFieldElement = [6]uint64

// One sets e = 1, and returns e.
func (e *P384Element) One() *P384Element {
	p384SetOne(&e.x)
	return e
}

// Equal returns 1 if e == t, and zero otherwise.
func (e *P384Element) Equal(t *P384Element) int {
	eBytes := e.Bytes()
	tBytes := t.Bytes()
	return subtle.ConstantTimeCompare(eBytes, tBytes)
}

var p384ZeroEncoding = new(P384Element).Bytes()

// IsZero returns 1 if e == 0, and zero otherwise.
func (e *P384Element) IsZero() int {
	eBytes := e.Bytes()
	return subtle.ConstantTimeCompare(eBytes, p384ZeroEncoding)
}

// Set sets e = t, and returns e.
func (e *P384Element) Set(t *P384Element) *P384Element {
	e.x = t.x
	return e
}

// Bytes returns the 48-byte big-endian encoding of e.
func (e *P384Element) Bytes() []byte {
	// This function is outlined to make the allocations inline in the caller
	// rather than happen on the heap.
	var out [p384ElementLen]byte
	return e.bytes(&out)
}

func (e *P384Element) bytes(out *[p384ElementLen]byte) []byte {
	var tmp p384NonMontgomeryDomainFieldElement
	p384FromMontgomery(&tmp, &e.x)
	p384ToBytes(out, (*p384UntypedFieldElement)(&tmp))
	p384InvertEndianness(out[:])
	return out[:]
}

// p384MinusOneEncoding is the encoding of -1 mod p, so p - 1, the
// highest canonical encoding. It is used by SetBytes to check for non-canonical
// encodings such as p + k, 2p + k, etc.
var p384MinusOneEncoding = new(P384Element).Sub(
	new(P384Element), new(P384Element).One()).Bytes()

// SetBytes sets e = v, where v is a big-endian 48-byte encoding, and returns e.
// If v is not 48 bytes or it encodes a value higher than 2^384 - 2^128 - 2^96 + 2^32 - 1,
// SetBytes returns nil and an error, and e is unchanged.
func (e *P384Element) SetBytes(v []byte) (*P384Element, error) {
	if len(v) != p384ElementLen {
		return nil, errors.New("invalid P384Element encoding")
	}
	for i := range v {
		if v[i] < p384MinusOneEncoding[i] {
			break
		}
		if v[i] > p384MinusOneEncoding[i] {
			return nil, errors.New("invalid P384Element encoding")
		}
	}
	var in [p384ElementLen]byte
	copy(in[:], v)
	p384InvertEndianness(in[:])
	var tmp p384NonMontgomeryDomainFieldElement
	p384FromBytes((*p384UntypedFieldElement)(&tmp), &in)
	p384ToMontgomery(&e.x, &tmp)
	return e, nil
}

// Add sets e = t1 + t2, and returns e.
func (e *P384Element) Add(t1, t2 *P384Element) *P384Element {
	p384Add(&e.x, &t1.x, &t2.x)
	return e
}

// Sub sets e = t1 - t2, and returns e.
func (e *P384Element) Sub(t1, t2 *P384Element) *P384Element {
	p384Sub(&e.x, &t1.x, &t2.x)
	return e
}

// Mul sets e = t1 * t2, and returns e.
func (e *P384Element) Mul(t1, t2 *P384Element) *P384Element {
	p384Mul(&e.x, &t1.x, &t2.x)
	return e
}

// Square sets e = t * t, and returns e.
func (e *P384Element) Square(t *P384Element) *P384Element {
	p384Square(&e.x, &t.x)
	return e
}

// Select sets v to a if cond == 1, and to b if cond == 0.
func (v *P384Element) Select(a, b *P384Element, cond int) *P384Element {
	p384Selectznz((*p384UntypedFieldElement)(&v.x), p384Uint1(cond),
		(*p384UntypedFieldElement)(&b.x), (*p384UntypedFieldElement)(&a.x))
	return v
}

func p384InvertEndianness(v []byte) {
	for i := 0; i < len(v)/2; i++ {
		v[i], v[len(v)-1-i] = v[len(v)-1-i], v[i]
	}
}
