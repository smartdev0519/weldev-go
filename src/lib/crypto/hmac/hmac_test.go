// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hmac

// TODO(rsc): better test

import (
	"hash";
	"crypto/hmac";
	"io";
	"fmt";
	"testing";
)

type hmacTest struct {
	hash func([]byte) hash.Hash;
	key []byte;
	in []byte;
	out string;
}

// Tests from US FIPS 198
// http://csrc.nist.gov/publications/fips/fips198/fips-198a.pdf
var hmacTests = []hmacTest {
	hmacTest {
		NewSHA1,
		[]byte {
			0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
			0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
			0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
			0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
			0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27,
			0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
			0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37,
			0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
		},
		io.StringBytes("Sample #1"),
		"4f4ca3d5d68ba7cc0a1208c9c61e9c5da0403c0a",
	},
	hmacTest {
		NewSHA1,
		[]byte {
			0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37,
			0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
			0x40, 0x41, 0x42, 0x43,
		},
		io.StringBytes("Sample #2"),
		"0922d3405faa3d194f82a45830737d5cc6c75d24",
	},
	hmacTest {
		NewSHA1,
		[]byte {
			0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57,
			0x58, 0x59, 0x5a, 0x5b, 0x5c, 0x5d, 0x5e, 0x5f,
			0x60, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67,
			0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6e, 0x6f,
			0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77,
			0x78, 0x79, 0x7a, 0x7b, 0x7c, 0x7d, 0x7e, 0x7f,
			0x80, 0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87,
			0x88, 0x89, 0x8a, 0x8b, 0x8c, 0x8d, 0x8e, 0x8f,
			0x90, 0x91, 0x92, 0x93, 0x94, 0x95, 0x96, 0x97,
			0x98, 0x99, 0x9a, 0x9b, 0x9c, 0x9d, 0x9e, 0x9f,
			0xa0, 0xa1, 0xa2, 0xa3, 0xa4, 0xa5, 0xa6, 0xa7,
			0xa8, 0xa9, 0xaa, 0xab, 0xac, 0xad, 0xae, 0xaf,
			0xb0, 0xb1, 0xb2, 0xb3,
		},
		io.StringBytes("Sample #3"),
		"bcf41eab8bb2d802f3d05caf7cb092ecf8d1a3aa",
	},

	// Test from Plan 9.
	hmacTest {
		NewMD5,
		io.StringBytes("Jefe"),
		io.StringBytes("what do ya want for nothing?"),
		"750c783e6ab0b503eaa86e310a5db738",
	}
}

func TestHMAC(t *testing.T) {
	for i, tt := range hmacTests {
		h := tt.hash(tt.key);
		for j := 0; j < 2; j++ {
			n, err := h.Write(tt.in);
			if n != len(tt.in) || err != nil {
				t.Errorf("test %d.%d: Write(%d) = %d, %v", i, j, len(tt.in), n, err);
				continue;
			}
			sum := fmt.Sprintf("%x", h.Sum());
			if sum != tt.out {
				t.Errorf("test %d.%d: have %s want %s\n", i, j, sum, tt.out);
			}

			// Second iteration: make sure reset works.
			h.Reset();
		}
	}
}
