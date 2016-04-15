// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package aes

import (
	"crypto/cipher"
)

// defined in asm_amd64.s
func hasAsm() bool
func encryptBlockAsm(nr int, xk *uint32, dst, src *byte)
func decryptBlockAsm(nr int, xk *uint32, dst, src *byte)
func expandKeyAsm(nr int, key *byte, enc *uint32, dec *uint32)

type aesCipherAsm struct {
	aesCipher
}

var useAsm = hasAsm()

func newCipher(key []byte) (cipher.Block, error) {
	if !useAsm {
		return newCipherGeneric(key)
	}
	n := len(key) + 28
	c := aesCipherAsm{aesCipher{make([]uint32, n), make([]uint32, n)}}
	rounds := 10
	switch len(key) {
	case 128 / 8:
		rounds = 10
	case 192 / 8:
		rounds = 12
	case 256 / 8:
		rounds = 14
	}
	expandKeyAsm(rounds, &key[0], &c.enc[0], &c.dec[0])
	if hasGCMAsm() {
		return &aesCipherGCM{c}, nil
	}

	return &c, nil
}

func (c *aesCipherAsm) BlockSize() int { return BlockSize }

func (c *aesCipherAsm) Encrypt(dst, src []byte) {
	if len(src) < BlockSize {
		panic("crypto/aes: input not full block")
	}
	if len(dst) < BlockSize {
		panic("crypto/aes: output not full block")
	}
	encryptBlockAsm(len(c.enc)/4-1, &c.enc[0], &dst[0], &src[0])
}

func (c *aesCipherAsm) Decrypt(dst, src []byte) {
	if len(src) < BlockSize {
		panic("crypto/aes: input not full block")
	}
	if len(dst) < BlockSize {
		panic("crypto/aes: output not full block")
	}
	decryptBlockAsm(len(c.dec)/4-1, &c.dec[0], &dst[0], &src[0])
}
