// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// Minimax polynomial coefficients and other constants
DATA ·asinrodataL15<> + 0(SB)/8, $-1.309611320495605469
DATA ·asinrodataL15<> + 8(SB)/8, $0x3ff921fb54442d18
DATA ·asinrodataL15<> + 16(SB)/8, $0xbff921fb54442d18
DATA ·asinrodataL15<> + 24(SB)/8, $1.309611320495605469
DATA ·asinrodataL15<> + 32(SB)/8, $-0.0
DATA ·asinrodataL15<> + 40(SB)/8, $1.199437040755305217
DATA ·asinrodataL15<> + 48(SB)/8, $0.166666666666651626E+00
DATA ·asinrodataL15<> + 56(SB)/8, $0.750000000042621169E-01
DATA ·asinrodataL15<> + 64(SB)/8, $0.446428567178116477E-01
DATA ·asinrodataL15<> + 72(SB)/8, $0.303819660378071894E-01
DATA ·asinrodataL15<> + 80(SB)/8, $0.223715011892010405E-01
DATA ·asinrodataL15<> + 88(SB)/8, $0.173659424522364952E-01
DATA ·asinrodataL15<> + 96(SB)/8, $0.137810186504372266E-01
DATA ·asinrodataL15<> + 104(SB)/8, $0.134066870961173521E-01
DATA ·asinrodataL15<> + 112(SB)/8, $-.412335502831898721E-02
DATA ·asinrodataL15<> + 120(SB)/8, $0.867383739532082719E-01
DATA ·asinrodataL15<> + 128(SB)/8, $-.328765950607171649E+00
DATA ·asinrodataL15<> + 136(SB)/8, $0.110401073869414626E+01
DATA ·asinrodataL15<> + 144(SB)/8, $-.270694366992537307E+01
DATA ·asinrodataL15<> + 152(SB)/8, $0.500196500770928669E+01
DATA ·asinrodataL15<> + 160(SB)/8, $-.665866959108585165E+01
DATA ·asinrodataL15<> + 168(SB)/8, $-.344895269334086578E+01
DATA ·asinrodataL15<> + 176(SB)/8, $0.927437952918301659E+00
DATA ·asinrodataL15<> + 184(SB)/8, $0.610487478874645653E+01
DATA ·asinrodataL15<> + 192(SB)/8, $0x7ff8000000000000			//+Inf
DATA ·asinrodataL15<> + 200(SB)/8, $-1.0
DATA ·asinrodataL15<> + 208(SB)/8, $1.0
DATA ·asinrodataL15<> + 216(SB)/8, $1.00000000000000000e-20
GLOBL ·asinrodataL15<> + 0(SB), RODATA, $224

// Asin returns the arcsine, in radians, of the argument.
//
// Special cases are:
//      Asin(±0) = ±0=
//      Asin(x) = NaN if x < -1 or x > 1
// The algorithm used is minimax polynomial approximation
// with coefficients determined with a Remez exchange algorithm.

TEXT	·asinAsm(SB), NOSPLIT, $0-16
	FMOVD	x+0(FP), F0
	MOVD	$·asinrodataL15<>+0(SB), R9
	WORD	$0xB3CD0070	//lgdr %r7, %f0
	FMOVD	F0, F8
	SRAD	$32, R7
	WORD	$0xC0193FE6 //iilf  %r1,1072079005
	BYTE	$0xA0
	BYTE	$0x9D
	WORD	$0xB91700C7 //llgtr %r12,%r7
	MOVW	R12, R8
	MOVW	R1, R6
	CMPBGT	R8, R6, L2
	WORD	$0xC0193BFF //iilf  %r1,1006632959
	BYTE	$0xFF
	BYTE	$0xFF
	MOVW	R1, R6
	CMPBGT	R8, R6, L13
L3:
	FMOVD	216(R9), F0
	FMADD	F0, F8, F8
L1:
	FMOVD	F8, ret+8(FP)
	RET
L2:
	WORD	$0xC0193FEF	//iilf	%r1,1072693247
	BYTE	$0xFF
	BYTE	$0xFF
	CMPW	R12, R1
	BLE	L14
L5:
	WORD	$0xED0090D0	//cdb	%f0,.L17-.L15(%r9)
	BYTE	$0x00
	BYTE	$0x19
	BEQ		L9
	WORD	$0xED0090C8	//cdb	%f0,.L18-.L15(%r9)
	BYTE	$0x00
	BYTE	$0x19
	BEQ	L10
	WFCEDBS	V8, V8, V0
	BVS	L1
	FMOVD	192(R9), F8
	BR	L1
L13:
	WFMDB	V0, V0, V10
L4:
	WFMDB	V10, V10, V0
	FMOVD	184(R9), F6
	FMOVD	176(R9), F2
	FMOVD	168(R9), F4
	WFMADB	V0, V2, V6, V2
	FMOVD	160(R9), F6
	WFMADB	V0, V4, V6, V4
	FMOVD	152(R9), F6
	WFMADB	V0, V2, V6, V2
	FMOVD	144(R9), F6
	WFMADB	V0, V4, V6, V4
	FMOVD	136(R9), F6
	WFMADB	V0, V2, V6, V2
	WORD	$0xC0193FE6	//iilf	%r1,1072079005
	BYTE	$0xA0
	BYTE	$0x9D
	FMOVD	128(R9), F6
	WFMADB	V0, V4, V6, V4
	FMOVD	120(R9), F6
	WFMADB	V0, V2, V6, V2
	FMOVD	112(R9), F6
	WFMADB	V0, V4, V6, V4
	FMOVD	104(R9), F6
	WFMADB	V0, V2, V6, V2
	FMOVD	96(R9), F6
	WFMADB	V0, V4, V6, V4
	FMOVD	88(R9), F6
	WFMADB	V0, V2, V6, V2
	FMOVD	80(R9), F6
	WFMADB	V0, V4, V6, V4
	FMOVD	72(R9), F6
	WFMADB	V0, V2, V6, V2
	FMOVD	64(R9), F6
	WFMADB	V0, V4, V6, V4
	FMOVD	56(R9), F6
	WFMADB	V0, V2, V6, V2
	FMOVD	48(R9), F6
	WFMADB	V0, V4, V6, V0
	WFMDB	V8, V10, V4
	FMADD	F2, F10, F0
	FMADD	F0, F4, F8
	CMPW	R12, R1
	BLE	L1
	FMOVD	40(R9), F0
	FMADD	F0, F1, F8
	FMOVD	F8, ret+8(FP)
	RET
L14:
	FMOVD	200(R9), F0
	FMADD	F8, F8, F0
	WORD	$0xB31300A0	//lcdbr	%f10,%f0
	WORD	$0xED009020	//cdb	%f0,.L39-.L15(%r9)
	BYTE	$0x00
	BYTE	$0x19
	FSQRT	F10, F8
L6:
	MOVW	R7, R6
	CMPBLE	R6, $0, L8
	WORD	$0xB3130088	//lcdbr	%f8,%f8
	FMOVD	24(R9), F1
	BR	L4
L10:
	FMOVD	16(R9), F8
	BR	L1
L9:
	FMOVD	8(R9), F8
	FMOVD	F8, ret+8(FP)
	RET
L8:
	FMOVD	0(R9), F1
	BR	L4
