// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !purego

#include "textflag.h"

// func kdsaSign(message, signature, privateKey []byte) bool
TEXT ·kdsaSign(SB), $4096-73
	// The kdsa instruction takes function code,
	// buffer's location, message's location and message len
	// as parameters. Out of those, the function code and buffer's location
	// should be placed in R0 and R1 respectively. The message's location
	// and message length should be placed in an even-odd register pair. (e.g: R2 and R3)

	// The content of parameter block(buffer) looks like the following:
	// Signature R, Signature S and Private Key all take 32 bytes.
	// In the signing case, the signatures(R and S) will be generated by
	// the signing instruction and get placed in the locations shown in the parameter block.
	//    0 +---------------+
	//      |  Signature(R) |
	//   32 +---------------+
	//      |  Signature(S) |
	//   64 +---------------+
	//      |  Private Key  |
	//   96 +---------------+
	//      |   Reserved    |
	//   112+---------------+
	//      |               |
	//      |     ...       |
	//      |               |
	// 4088 +---------------+

	// The following code section setups the buffer from stack:
	// Get the address of the buffer stack variable.
	MOVD $buffer-4096(SP), R1

	// Zero the buffer.
	MOVD R1, R2
	MOVD $(4096/256), R0 // number of 256 byte chunks to clear

clear:
	XC    $256, (R2), (R2)
	MOVD  $256(R2), R2
	BRCTG R0, clear

	MOVD $40, R0                   // EDDSA-25519 sign has a function code of 40
	LMG  message+0(FP), R2, R3     // R2=base R3=len
	LMG  signature+24(FP), R4, R5  // R4=base R5=len
	LMG  privateKey+48(FP), R6, R7 // R6=base R7=len

	// Checks the length of signature and private key
	CMPBNE R5, $64, panic
	CMPBNE R7, $32, panic

	// The instruction uses RFC 8032's private key, which is the first 32 bytes
	// of the private key in this package. So we copy that into the buffer.
	MVC $32, (R6), 64(R1)

loop:
	WORD $0xB93A0002 // The KDSA instruction
	BVS  loop        // The instruction is exectued by hardware and can be interrupted. This does a retry when that happens.
	BNE  error

success:
	// The signatures generated are in big-endian form, so we
	// need to reverse the bytes of Signature(R) and Signature(S) in the buffers to transform
	// them from big-endian to little-endian.

	// Transform Signature(R) from big endian to little endian and copy into the signature
	MVCIN $32, 31(R1), (R4)

	// Transform Signature(S) from big endian to little endian and copy into the signature
	MVCIN $32, 63(R1), 32(R4)

	MOVB $1, ret+72(FP)
	RET

error:
	// return false
	MOVB $0, ret+72(FP)
	RET

panic:
	UNDEF

// func kdsaVerify(message, signature, publicKey []byte) bool
TEXT ·kdsaVerify(SB), $4096-73
	// The kdsa instruction takes function code,
	// buffer's location, message's location and message len
	// as parameters. Out of those, the function code and buffer's location
	// should be placed in R0 and R1 respectively. The message's location
	// and message length should be placed in an even-odd register pair. (e.g: R2 and R3)

	// The parameter block(buffer) is similar to that of signing, except that
	// we use public key for verification, and Signatures(R and S) are provided
	// as input parameters to the parameter block.
	//    0 +---------------+
	//      |  Signature(R) |
	//   32 +---------------+
	//      |  Signature(S) |
	//   64 +---------------+
	//      |  Public Key   |
	//   96 +---------------+
	//      |   Reserved    |
	//   112+---------------+
	//      |               |
	//      |     ...       |
	//      |               |
	// 4088 +---------------+

	// The following code section setups the buffer from stack:
	// Get the address of the buffer stack variable.
	MOVD $buffer-4096(SP), R1

	// Zero the buffer.
	MOVD R1, R2
	MOVD $(4096/256), R0 // number of 256 byte chunks to clear

clear:
	XC    $256, (R2), (R2)
	MOVD  $256(R2), R2
	BRCTG R0, clear

	MOVD $32, R0                  // EDDSA-25519 verify has a function code of 32
	LMG  message+0(FP), R2, R3    // R2=base R3=len
	LMG  signature+24(FP), R4, R5 // R4=base R5=len
	LMG  publicKey+48(FP), R6, R7 // R6=base R7=len

	// Checks the length of public key and signature
	CMPBNE R5, $64, panic
	CMPBNE R7, $32, panic

verify:
	// The instruction needs Signature(R), Signature(S) and public key
	// to be in big-endian form during computation. Therefore,
	// we do the transformation (from little endian to big endian) and copy those into the buffer.

	// Transform Signature(R) from little endian to big endian and copy into the buffer
	MVCIN $32, 31(R4), (R1)

	// Transform Signature(S) from little endian to big endian and copy into the buffer
	MVCIN $32, 63(R4), 32(R1)

	// Transform Public Key from little endian to big endian and copy into the buffer
	MVCIN $32, 31(R6), 64(R1)

verifyLoop:
	WORD $0xB93A0002 // KDSA instruction
	BVS  verifyLoop  // Retry upon hardware interrupt
	BNE  error

success:
	MOVB $1, ret+72(FP)
	RET

error:
	MOVB $0, ret+72(FP)
	RET

panic:
	UNDEF
