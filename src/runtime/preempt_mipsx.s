// Code generated by mkpreempt.go; DO NOT EDIT.

// +build mips mipsle

#include "go_asm.h"
#include "textflag.h"

// Note: asyncPreempt doesn't use the internal ABI, but we must be able to inject calls to it from the signal handler, so Go code has to see the PC of this function literally.
TEXT ·asyncPreempt<ABIInternal>(SB),NOSPLIT|NOFRAME,$0-0
	MOVW R31, -244(R29)
	SUB $244, R29
	MOVW R1, 4(R29)
	MOVW R2, 8(R29)
	MOVW R3, 12(R29)
	MOVW R4, 16(R29)
	MOVW R5, 20(R29)
	MOVW R6, 24(R29)
	MOVW R7, 28(R29)
	MOVW R8, 32(R29)
	MOVW R9, 36(R29)
	MOVW R10, 40(R29)
	MOVW R11, 44(R29)
	MOVW R12, 48(R29)
	MOVW R13, 52(R29)
	MOVW R14, 56(R29)
	MOVW R15, 60(R29)
	MOVW R16, 64(R29)
	MOVW R17, 68(R29)
	MOVW R18, 72(R29)
	MOVW R19, 76(R29)
	MOVW R20, 80(R29)
	MOVW R21, 84(R29)
	MOVW R22, 88(R29)
	MOVW R24, 92(R29)
	MOVW R25, 96(R29)
	MOVW R28, 100(R29)
	MOVW HI, R1
	MOVW R1, 104(R29)
	MOVW LO, R1
	MOVW R1, 108(R29)
	#ifndef GOMIPS_softfloat
	MOVW FCR31, R1
	MOVW R1, 112(R29)
	MOVF F0, 116(R29)
	MOVF F1, 120(R29)
	MOVF F2, 124(R29)
	MOVF F3, 128(R29)
	MOVF F4, 132(R29)
	MOVF F5, 136(R29)
	MOVF F6, 140(R29)
	MOVF F7, 144(R29)
	MOVF F8, 148(R29)
	MOVF F9, 152(R29)
	MOVF F10, 156(R29)
	MOVF F11, 160(R29)
	MOVF F12, 164(R29)
	MOVF F13, 168(R29)
	MOVF F14, 172(R29)
	MOVF F15, 176(R29)
	MOVF F16, 180(R29)
	MOVF F17, 184(R29)
	MOVF F18, 188(R29)
	MOVF F19, 192(R29)
	MOVF F20, 196(R29)
	MOVF F21, 200(R29)
	MOVF F22, 204(R29)
	MOVF F23, 208(R29)
	MOVF F24, 212(R29)
	MOVF F25, 216(R29)
	MOVF F26, 220(R29)
	MOVF F27, 224(R29)
	MOVF F28, 228(R29)
	MOVF F29, 232(R29)
	MOVF F30, 236(R29)
	MOVF F31, 240(R29)
	#endif
	CALL ·asyncPreempt2(SB)
	#ifndef GOMIPS_softfloat
	MOVF 240(R29), F31
	MOVF 236(R29), F30
	MOVF 232(R29), F29
	MOVF 228(R29), F28
	MOVF 224(R29), F27
	MOVF 220(R29), F26
	MOVF 216(R29), F25
	MOVF 212(R29), F24
	MOVF 208(R29), F23
	MOVF 204(R29), F22
	MOVF 200(R29), F21
	MOVF 196(R29), F20
	MOVF 192(R29), F19
	MOVF 188(R29), F18
	MOVF 184(R29), F17
	MOVF 180(R29), F16
	MOVF 176(R29), F15
	MOVF 172(R29), F14
	MOVF 168(R29), F13
	MOVF 164(R29), F12
	MOVF 160(R29), F11
	MOVF 156(R29), F10
	MOVF 152(R29), F9
	MOVF 148(R29), F8
	MOVF 144(R29), F7
	MOVF 140(R29), F6
	MOVF 136(R29), F5
	MOVF 132(R29), F4
	MOVF 128(R29), F3
	MOVF 124(R29), F2
	MOVF 120(R29), F1
	MOVF 116(R29), F0
	MOVW 112(R29), R1
	MOVW R1, FCR31
	#endif
	MOVW 108(R29), R1
	MOVW R1, LO
	MOVW 104(R29), R1
	MOVW R1, HI
	MOVW 100(R29), R28
	MOVW 96(R29), R25
	MOVW 92(R29), R24
	MOVW 88(R29), R22
	MOVW 84(R29), R21
	MOVW 80(R29), R20
	MOVW 76(R29), R19
	MOVW 72(R29), R18
	MOVW 68(R29), R17
	MOVW 64(R29), R16
	MOVW 60(R29), R15
	MOVW 56(R29), R14
	MOVW 52(R29), R13
	MOVW 48(R29), R12
	MOVW 44(R29), R11
	MOVW 40(R29), R10
	MOVW 36(R29), R9
	MOVW 32(R29), R8
	MOVW 28(R29), R7
	MOVW 24(R29), R6
	MOVW 20(R29), R5
	MOVW 16(R29), R4
	MOVW 12(R29), R3
	MOVW 8(R29), R2
	MOVW 4(R29), R1
	MOVW 244(R29), R31
	MOVW (R29), R23
	ADD $248, R29
	JMP (R23)
