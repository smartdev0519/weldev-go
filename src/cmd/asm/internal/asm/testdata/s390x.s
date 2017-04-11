// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "../../../../../runtime/textflag.h"

TEXT main·foo(SB),DUPOK|NOSPLIT,$16-0 // TEXT main.foo(SB), DUPOK|NOSPLIT, $16-0
	MOVD	R1, R2                // b9040021
	MOVW	R3, R4                // b9140043
	MOVH	R5, R6                // b9070065
	MOVB	R7, R8                // b9060087
	MOVWZ	R1, R2                // b9160021
	MOVHZ	R2, R3                // b9850032
	MOVBZ	R4, R5                // b9840054
	MOVDBR	R1, R2                // b90f0021
	MOVWBR	R3, R4                // b91f0043

	MOVDEQ	R0, R1                // b9e28010
	MOVDGE	R2, R3                // b9e2a032
	MOVDGT	R4, R5                // b9e22054
	MOVDLE	R6, R7                // b9e2c076
	MOVDLT	R8, R9                // b9e24098
	MOVDNE	R10, R11              // b9e270ba

	MOVD	(R15), R1             // e310f0000004
	MOVW	(R15), R2             // e320f0000014
	MOVH	(R15), R3             // e330f0000015
	MOVB	(R15), R4             // e340f0000077
	MOVWZ	(R15), R5             // e350f0000016
	MOVHZ	(R15), R6             // e360f0000091
	MOVBZ	(R15), R7             // e370f0000090
	MOVDBR	(R15), R8             // e380f000000f
	MOVWBR	(R15), R9             // e390f000001e

	MOVD	R1, n-8(SP)           // e310f0100024
	MOVW	R2, n-8(SP)           // e320f0100050
	MOVH	R3, n-8(SP)           // e330f0100070
	MOVB	R4, n-8(SP)           // e340f0100072
	MOVWZ	R5, n-8(SP)           // e350f0100050
	MOVHZ	R6, n-8(SP)           // e360f0100070
	MOVBZ	R7, n-8(SP)           // e370f0100072
	MOVDBR	R8, n-8(SP)           // e380f010002f
	MOVWBR	R9, n-8(SP)           // e390f010003e

	MOVD	$-8589934592, R1      // c01efffffffe
	MOVW	$-131072, R2          // c021fffe0000
	MOVH	$-512, R3             // a739fe00
	MOVB	$-1, R4               // a749ffff

	MOVD	$-2147483648, n-8(SP) // c0b180000000e3b0f0100024
	MOVW	$-131072, n-8(SP)     // c0b1fffe0000e3b0f0100050
	MOVH	$-512, n-8(SP)        // e544f010fe00
	MOVB	$-1, n-8(SP)          // 92fff010

	ADD	R1, R2                // b9e81022
	ADD	R1, R2, R3            // b9e81032
	ADD	$8192, R1             // a71b2000
	ADD	$8192, R1, R2         // ec21200000d9
	ADD	$32768, R1            // c21800008000
	ADD	$32768, R1, R2        // b9040021c22800008000
	ADDC	R1, R2                // b9ea1022
	ADDC	$1, R1, R2            // ec21000100db
	ADDC	R1, R2, R3            // b9ea1032
	ADDW	R1, R2                // 1a21
	ADDW	R1, R2, R3            // b9f81032
	ADDW	$8192, R1             // a71a2000
	ADDW	$8192, R1, R2         // ec21200000d8
	SUB	R3, R4                // b9090043
	SUB	R3, R4, R5            // b9e93054
	SUB	$8192, R3             // a73be000
	SUB	$8192, R3, R4         // ec43e00000d9
	SUBC	R1, R2                // b90b0021
	SUBC	$1, R1, R2            // ec21ffff00db
	SUBC	R2, R3, R4            // b9eb2043
	SUBW	R3, R4                // 1b43
	SUBW	R3, R4, R5            // b9f93054
	SUBW	$8192, R1             // c21500002000
	SUBW	$8192, R1, R2         // 1821c22500002000
	MULLW	R6, R7                // b91c0076
	MULLW	R6, R7, R8            // b9040087b91c0086
	MULLW	$8192, R6             // a76c2000
	MULLW	$8192, R6, R7         // 1876a77c2000
	MULLW	$-32769, R8           // c281ffff7fff
	MULLW   $-32769, R8, R9       // 1898c291ffff7fff
	MULLD	$-2147483648, R1      // c21080000000
	MULLD   $-2147483648, R1, R2  // b9040021c22080000000
	MULHD	R9, R8                // b90400b8b98600a9ebb9003f000ab98000b8b90900abebb8003f000ab98000b9b9e9b08a
	MULHD	R7, R2, R1            // b90400b2b98600a7ebb7003f000ab98000b2b90900abebb2003f000ab98000b7b9e9b01a
	MULHDU	R3, R4                // b90400b4b98600a3b904004a
	MULHDU	R5, R6, R7            // b90400b6b98600a5b904007a
	DIVD	R1, R2                // b90400b2b90d00a1b904002b
	DIVD	R1, R2, R3            // b90400b2b90d00a1b904003b
	DIVW	R4, R5                // b90400b5b91d00a4b904005b
	DIVW	R4, R5, R6            // b90400b5b91d00a4b904006b
	DIVDU	R7, R8                // a7a90000b90400b8b98700a7b904008b
	DIVDU	R7, R8, R9            // a7a90000b90400b8b98700a7b904009b
	DIVWU	R1, R2                // a7a90000b90400b2b99700a1b904002b
	DIVWU	R1, R2, R3            // a7a90000b90400b2b99700a1b904003b
	MODD	R1, R2                // b90400b2b90d00a1b904002a
	MODD	R1, R2, R3            // b90400b2b90d00a1b904003a
	MODW	R4, R5                // b90400b5b91d00a4b904005a
	MODW	R4, R5, R6            // b90400b5b91d00a4b904006a
	MODDU	R7, R8                // a7a90000b90400b8b98700a7b904008a
	MODDU	R7, R8, R9            // a7a90000b90400b8b98700a7b904009a
	MODWU	R1, R2                // a7a90000b90400b2b99700a1b904002a
	MODWU	R1, R2, R3            // a7a90000b90400b2b99700a1b904003a
	NEG	R1                    // b9030011
	NEG	R1, R2                // b9030021
	NEGW	R1                    // b9130011
	NEGW	R1, R2                // b9130021
	FLOGR	R2, R2                // b9830022

	AND	R1, R2                // b9800021
	AND	R1, R2, R3            // b9e42031
	AND	$-2, R1               // a517fffe
	AND	$-65536, R1           // c01bffff0000
	AND	$1, R1                // c0a100000001b980001a
	ANDW	R1, R2                // 1421
	ANDW	R1, R2, R3            // b9f42031
	ANDW	$1, R1                // c01b00000001
	ANDW	$131071, R1           // a5160001
	ANDW	$65536, R1            // c01b00010000
	ANDW	$-2, R1               // a517fffe
	OR	R1, R2                // b9810021
	OR	R1, R2, R3            // b9e62031
	OR	$1, R1                // a51b0001
	OR	$131071, R1           // c01d0001ffff
	OR	$65536, R1            // c01d00010000
	OR	$-2, R1               // c0a1fffffffeb981001a
	ORW	R1, R2                // 1621
	ORW	R1, R2, R3            // b9f62031
	ORW	$1, R1                // a51b0001
	ORW	$131071, R1           // c01d0001ffff
	ORW	$65536, R1            // a51a0001
	ORW	$-2, R1               // c01dfffffffe
	XOR	R1, R2                // b9820021
	XOR	R1, R2, R3            // b9e72031
	XOR	$1, R1                // c01700000001
	XOR	$131071, R1           // c0170001ffff
	XOR	$65536, R1            // c01700010000
	XOR	$-2, R1               // c0a1fffffffeb982001a
	XORW	R1, R2                // 1721
	XORW	R1, R2, R3            // b9f72031
	XORW	$1, R1                // c01700000001
	XORW	$131071, R1           // c0170001ffff
	XORW	$65536, R1            // c01700010000
	XORW	$-2, R1               // c017fffffffe

	ADD	-524288(R1), R2       // e32010008008
	ADD	524287(R3), R4        // e3403fff7f08
	ADD	-524289(R1), R2       // c0a1fff7ffffe32a10000008
	ADD	524288(R3), R4        // c0a100080000e34a30000008
	ADD	-524289(R1)(R2*1), R3 // c0a1fff7ffff41aa2000e33a10000008
	ADD	524288(R3)(R4*1), R5  // c0a10008000041aa4000e35a30000008
	ADDC	(R1), R2              // e3201000000a
	ADDW	(R5), R6              // 5a605000
	ADDW	4095(R7), R8          // 5a807fff
	ADDW	-1(R1), R2            // e3201fffff5a
	ADDW	4096(R3), R4          // e3403000015a
	MULLD	(R1)(R2*1), R3        // e3321000000c
	MULLW	(R3)(R4*1), R5        // 71543000
	MULLW	4096(R3), R4          // e34030000151
	SUB	(R1), R2              // e32010000009
	SUBC	(R1), R2              // e3201000000b
	SUBE	(R1), R2              // e32010000089
	SUBW	(R1), R2              // 5b201000
	SUBW	-1(R1), R2            // e3201fffff5b
	AND	(R1), R2              // e32010000080
	ANDW	(R1), R2              // 54201000
	ANDW	-1(R1), R2            // e3201fffff54
	OR	(R1), R2              // e32010000081
	ORW	(R1), R2              // 56201000
	ORW	-1(R1), R2            // e3201fffff56
	XOR	(R1), R2              // e32010000082
	XORW	(R1), R2              // 57201000
	XORW	-1(R1), R2            // e3201fffff57

	LAA	R1, R2, 524287(R3)    // eb213fff7ff8
	LAAG	R4, R5, -524288(R6)   // eb54600080e8
	LAAL	R7, R8, 8192(R9)      // eb87900002fa
	LAALG	R10, R11, -8192(R12)  // ebbac000feea
	LAN	R1, R2, (R3)          // eb21300000f4
	LANG	R4, R5, (R6)          // eb54600000e4
	LAX	R7, R8, (R9)          // eb87900000f7
	LAXG	R10, R11, (R12)       // ebbac00000e7
	LAO	R1, R2, (R3)          // eb21300000f6
	LAOG	R4, R5, (R6)          // eb54600000e6

	XC	$8, (R15), n-8(SP)       // d707f010f000
	NC	$8, (R15), n-8(SP)       // d407f010f000
	OC	$8, (R15), n-8(SP)       // d607f010f000
	MVC	$8, (R15), n-8(SP)       // d207f010f000
	CLC	$8, (R15), n-8(SP)       // d507f000f010
	XC	$256, -8(R15), -8(R15)   // b90400afc2a8fffffff8d7ffa000a000
	MVC	$256, 8192(R1), 8192(R2) // b90400a2c2a800002000b90400b1c2b800002000d2ffa000b000

	CMP	R1, R2                 // b9200012
	CMP	R3, $32767             // a73f7fff
	CMP	R3, $32768             // c23c00008000
	CMP	R3, $-2147483648       // c23c80000000
	CMPU	R4, R5                 // b9210045
	CMPU	R6, $4294967295        // c26effffffff
	CMPW	R7, R8                 // 1978
	CMPW	R9, $-32768            // a79e8000
	CMPW	R9, $-32769            // c29dffff7fff
	CMPW	R9, $-2147483648       // c29d80000000
	CMPWU	R1, R2                 // 1512
	CMPWU	R3, $4294967295        // c23fffffffff

	BNE	0(PC)                  // a7740000
	BEQ	0(PC)                  // a7840000
	BLT	0(PC)                  // a7440000
	BLE	0(PC)                  // a7c40000
	BGT	0(PC)                  // a7240000
	BGE	0(PC)                  // a7a40000
	BLTU	0(PC)                  // a7540000
	BLEU	0(PC)                  // a7d40000

	CMPBNE	R1, R2, 0(PC)          // ec1200007064
	CMPBEQ	R3, R4, 0(PC)          // ec3400008064
	CMPBLT	R5, R6, 0(PC)          // ec5600004064
	CMPBLE	R7, R8, 0(PC)          // ec780000c064
	CMPBGT	R9, R1, 0(PC)          // ec9100002064
	CMPBGE	R2, R3, 0(PC)          // ec230000a064

	CMPBNE	R1, $-127, 0(PC)       // ec170000817c
	CMPBEQ	R3, $0, 0(PC)          // ec380000007c
	CMPBLT	R5, $128, 0(PC)        // ec540000807c
	CMPBLE	R7, $127, 0(PC)        // ec7c00007f7c
	CMPBGT	R9, $0, 0(PC)          // ec920000007c
	CMPBGE	R2, $128, 0(PC)        // ec2a0000807c

	CMPUBNE	R1, R2, 0(PC)          // ec1200007065
	CMPUBEQ	R3, R4, 0(PC)          // ec3400008065
	CMPUBLT	R5, R6, 0(PC)          // ec5600004065
	CMPUBLE	R7, R8, 0(PC)          // ec780000c065
	CMPUBGT	R9, R1, 0(PC)          // ec9100002065
	CMPUBGE	R2, R3, 0(PC)          // ec230000a065

	CMPUBNE	R1, $256, 0(PC)        // ec170000007d
	CMPUBEQ	R3, $0, 0(PC)          // ec380000007d
	CMPUBLT	R5, $256, 0(PC)        // ec540000007d
	CMPUBLE	R7, $0, 0(PC)          // ec7c0000007d
	CMPUBGT	R9, $256, 0(PC)        // ec920000007d
	CMPUBGE	R2, $0, 0(PC)          // ec2a0000007d

	LGDR	F1, R12                // b3cd00c1
	LDGR	R2, F15                // b3c100f2

	CEFBRA	R0, F15                // b39400f0
	CDFBRA	R1, F14                // b39500e1
	CEGBRA	R2, F13                // b3a400d2
	CDGBRA	R3, F12                // b3a500c3

	CELFBR	R0, F15                // b39000f0
	CDLFBR	R1, F14                // b39100e1
	CELGBR	R2, F13                // b3a000d2
	CDLGBR	R3, F12                // b3a100c3

	CFEBRA	F15, R1                // b398501f
	CFDBRA	F14, R2                // b399502e
	CGEBRA	F13, R3                // b3a8503d
	CGDBRA	F12, R4                // b3a9504c

	CLFEBR	F15, R1                // b39c501f
	CLFDBR	F14, R2                // b39d502e
	CLGEBR	F13, R3                // b3ac503d
	CLGDBR	F12, R4                // b3ad504c

	FMOVS	$0, F11                // b37400b0
	FMOVD	$0, F12                // b37500c0
	FMOVS	(R1)(R2*1), F0         // ed0210000064
	FMOVS	n-8(SP), F15           // edf0f0100064
	FMOVD	-9999999(R8)(R9*1), F8 // c0a1ff67698141aa9000ed8a80000065
	FMOVD	F4, F5                 // 2854
	FADDS	F0, F15                // b30a00f0
	FADD	F1, F14                // b31a00e1
	FSUBS	F2, F13                // b30b00d2
	FSUB	F3, F12                // b31b00c3
	FMULS	F4, F11                // b31700b4
	FMUL	F5, F10                // b31c00a5
	FDIVS	F6, F9                 // b30d0096
	FDIV	F7, F8                 // b31d0087
	FABS	F1, F2                 // b3100021
	FSQRTS	F3, F4                 // b3140043
	FSQRT	F5, F15                // b31500f5
	FIEBR	$0, F0, F1             // b3570010
	FIDBR	$7, F2, F3             // b35f7032
	FMADD	F1, F1, F1             // b31e1011
	FMADDS	F1, F2, F3             // b30e3012
	FMSUB	F4, F5, F5             // b31f5045
	FMSUBS	F6, F6, F7             // b30f7066

	VL	(R15), V1               // e710f0000006
	VST	V1, (R15)               // e710f000000e
	VL	(R15), V31              // e7f0f0000806
	VST	V31, (R15)              // e7f0f000080e
	VESLB	$5, V14                 // e7ee00050030
	VESRAG	$0, V15, V16            // e70f0000383a
	VLM	(R15), V8, V23          // e787f0000436
	VSTM	V8, V23, (R15)          // e787f000043e
	VONE	V1                      // e710ffff0044
	VZERO	V16                     // e70000000844
	VGBM	$52428, V31             // e7f0cccc0844
	VREPIB	$255, V4                // e74000ff0045
	VREPIH	$-1, V16                // e700ffff1845
	VREPIF	$-32768, V0             // e70080002045
	VREPIG	$32767, V31             // e7f07fff3845
	VREPG	$1, V4, V16             // e7040001384d
	VREPB	$4, V31, V1             // e71f0004044d
	VFTCIDB	$4095, V1, V2           // e721fff0304a
	WFTCIDB	$3276, V15, V16         // e70fccc8384a
	VPOPCT	V8, V19                 // e73800000850
	VFEEZBS	V1, V2, V31             // e7f120300880
	WFCHDBS	V22, V23, V4            // e746701836eb
	VMNH	V1, V2, V30             // e7e1200018fe
	VO	V2, V1, V0              // e7021000006a
	VERLLVF	V2, V30, V27            // e7be20002c73
	VSCBIB	V0, V23, V24            // e78700000cf5
	VNOT	V16, V1                 // e7101000046b
	VCLZF	V16, V17                // e71000002c53
	VLVGP	R3, R4, V8              // e78340000062
	VGEG	$1, 8(R15)(V30*1), V31  // e7fef0081c12
	VSCEG	$1, V31, 16(R15)(V30*1) // e7fef0101c1a
	VGEF	$0, 2048(R15)(V1*1), V2 // e721f8000013
	VSCEF	$0, V2, 4095(R15)(V1*1) // e721ffff001b
	VLL	R0, (R15), V1           // e710f0000037
	VSTL	R0, V16, (R15)          // e700f000083f
	VGMH	$8, $16, V12            // e7c008101046
	VLEIB	$15, $255, V0           // e70000fff040
	VLEIH	$7, $-32768, V15        // e7f080007041
	VLEIF	$2, $-43, V16           // e700ffd52843
	VLEIG	$1, $32767, V31         // e7f07fff1842
	VSLDB	$3, V1, V16, V18        // e72100030a77
	VERIMB	$2, V31, V1, V2         // e72f10020472
	VSEL	V1, V2, V3, V4          // e7412000308d
	VGFMAH	V21, V31, V24, V0       // e705f10087bc
	VFMADB	V16, V8, V9, V10        // e7a08300948f
	WFMADB	V17, V18, V19, V20      // e74123083f8f
	VFMSDB	V2, V25, V24, V31       // e7f293008b8e
	WFMSDB	V31, V2, V3, V4         // e74f2308348e
	VPERM	V31, V0, V2, V3         // e73f0000248c
	VPDI	$1, V2, V31, V1         // e712f0001284
	VLEG	$1, (R3), V1            // e71030001002
	VLEF	$2, (R0), V31           // e7f000002803
	VLEH	$3, (R12), V16          // e700c0003801
	VLEB	$15, 4095(R9), V15      // e7f09ffff000
	VSTEG	$1, V30, (R1)(R2*1)     // e7e21000180a
	VSTEF	$3, V2, (R9)            // e7209000300b
	VSTEH	$7, V31, (R2)           // e7f020007809
	VSTEB	$15, V29, 4094(R12)     // e7d0cffef808

	RET

TEXT main·init(SB),DUPOK|NOSPLIT,$0 // TEXT main.init(SB), DUPOK|NOSPLIT, $0
	RET

TEXT main·main(SB),DUPOK|NOSPLIT,$0 // TEXT main.main(SB), DUPOK|NOSPLIT, $0
	BL      main·foo(SB)    // CALL main.foo(SB)
	RET
