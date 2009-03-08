// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This package implements the 32-bit cyclic redundancy check, or CRC-32, checksum.
// See http://en.wikipedia.org/wiki/Cyclic_redundancy_check for information.
package crc32

import "os"

// Predefined polynomials.
const (
	// Far and away the most common CRC-32 polynomial.
	// Used by ethernet (IEEE 802.3), v.42, fddi, gzip, zip, png, mpeg-2, ...
	IEEE = 0xedb88320;

	// Castagnoli's polynomial, used in iSCSI.
	// Has better error detection characteristics than IEEE.
	// http://dx.doi.org/10.1109/26.231911
	Castagnoli = 0x82f63b78;

	// Koopman's polynomial.
	// Also has better error detection characteristics than IEEE.
	// http://dx.doi.org/10.1109/DSN.2002.1028931
	Koopman = 0xeb31d82e;
)

// Table is a 256-word table representing the polynomial for efficient processing.
// TODO(rsc): Change to [256]uint32 once 6g can handle it.
type Table []uint32

// MakeTable returns the Table constructed from the specified polynomial.
func MakeTable(poly uint32) Table {
	t := make(Table, 256);
	for i := 0; i < 256; i++ {
		crc := uint32(i);
		for j := 0; j < 8; j++ {
			if crc&1 == 1 {
				crc = (crc>>1) ^ poly;
			} else {
				crc >>= 1;
			}
		}
		t[i] = crc;
	}
	return t;
}

// IEEETable is the table for the IEEE polynomial.
var IEEETable = MakeTable(IEEE);

// Digest represents the partial evaluation of a checksum.
type Digest struct {
	crc uint32;
	tab Table;
}

// NewDigest creates a new Digest for the checksum based on
// the polynomial represented by the Table.
func NewDigest(tab Table) *Digest {
	return &Digest{0, tab};
}

// NewIEEEDigest creates a new Digest for the checksum based on
// the IEEE polynomial.
func NewIEEEDigest() *Digest {
	return NewDigest(IEEETable);
}

// Write updates the Digest with the incremental checksum generated by p.
// It returns the number of bytes written; err is always nil.
func (d *Digest) Write(p []byte) (n int, err *os.Error) {
	crc := d.crc ^ 0xFFFFFFFF;
	tab := d.tab;
	for i := 0; i < len(p); i++ {
		crc = tab[byte(crc) ^ p[i]] ^ (crc >> 8);
	}
	d.crc = crc ^ 0xFFFFFFFF;
	return len(p), nil;
}

// Sum32 returns the CRC-32 checksum of the data written to the Digest.
func (d *Digest) Sum32() uint32 {
	return d.crc
}

// Sum returns the CRC-32 checksum of the data written to the Digest
// in the form of an array of 4 bytes in big-endian order.
func (d *Digest) Sum() []byte {
	p := make([]byte, 4);
	s := d.Sum32();
	p[0] = byte(s>>24);
	p[1] = byte(s>>16);
	p[2] = byte(s>>8);
	p[3] = byte(s);
	return p;
}


