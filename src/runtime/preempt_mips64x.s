// Code generated by mkpreempt.go; DO NOT EDIT.

// +build mips64 mips64le

#include "go_asm.h"
#include "textflag.h"

TEXT ·asyncPreempt(SB),NOSPLIT|NOFRAME,$0-0
	MOVV R31, -488(R29)
	SUBV $488, R29
	MOVV R1, 8(R29)
	MOVV R2, 16(R29)
	MOVV R3, 24(R29)
	MOVV R4, 32(R29)
	MOVV R5, 40(R29)
	MOVV R6, 48(R29)
	MOVV R7, 56(R29)
	MOVV R8, 64(R29)
	MOVV R9, 72(R29)
	MOVV R10, 80(R29)
	MOVV R11, 88(R29)
	MOVV R12, 96(R29)
	MOVV R13, 104(R29)
	MOVV R14, 112(R29)
	MOVV R15, 120(R29)
	MOVV R16, 128(R29)
	MOVV R17, 136(R29)
	MOVV R18, 144(R29)
	MOVV R19, 152(R29)
	MOVV R20, 160(R29)
	MOVV R21, 168(R29)
	MOVV R22, 176(R29)
	MOVV R24, 184(R29)
	MOVV R25, 192(R29)
	MOVV RSB, 200(R29)
	MOVV HI, R1
	MOVV R1, 208(R29)
	MOVV LO, R1
	MOVV R1, 216(R29)
	MOVV FCR31, R1
	MOVV R1, 224(R29)
	MOVD F0, 232(R29)
	MOVD F1, 240(R29)
	MOVD F2, 248(R29)
	MOVD F3, 256(R29)
	MOVD F4, 264(R29)
	MOVD F5, 272(R29)
	MOVD F6, 280(R29)
	MOVD F7, 288(R29)
	MOVD F8, 296(R29)
	MOVD F9, 304(R29)
	MOVD F10, 312(R29)
	MOVD F11, 320(R29)
	MOVD F12, 328(R29)
	MOVD F13, 336(R29)
	MOVD F14, 344(R29)
	MOVD F15, 352(R29)
	MOVD F16, 360(R29)
	MOVD F17, 368(R29)
	MOVD F18, 376(R29)
	MOVD F19, 384(R29)
	MOVD F20, 392(R29)
	MOVD F21, 400(R29)
	MOVD F22, 408(R29)
	MOVD F23, 416(R29)
	MOVD F24, 424(R29)
	MOVD F25, 432(R29)
	MOVD F26, 440(R29)
	MOVD F27, 448(R29)
	MOVD F28, 456(R29)
	MOVD F29, 464(R29)
	MOVD F30, 472(R29)
	MOVD F31, 480(R29)
	CALL ·asyncPreempt2(SB)
	MOVD 480(R29), F31
	MOVD 472(R29), F30
	MOVD 464(R29), F29
	MOVD 456(R29), F28
	MOVD 448(R29), F27
	MOVD 440(R29), F26
	MOVD 432(R29), F25
	MOVD 424(R29), F24
	MOVD 416(R29), F23
	MOVD 408(R29), F22
	MOVD 400(R29), F21
	MOVD 392(R29), F20
	MOVD 384(R29), F19
	MOVD 376(R29), F18
	MOVD 368(R29), F17
	MOVD 360(R29), F16
	MOVD 352(R29), F15
	MOVD 344(R29), F14
	MOVD 336(R29), F13
	MOVD 328(R29), F12
	MOVD 320(R29), F11
	MOVD 312(R29), F10
	MOVD 304(R29), F9
	MOVD 296(R29), F8
	MOVD 288(R29), F7
	MOVD 280(R29), F6
	MOVD 272(R29), F5
	MOVD 264(R29), F4
	MOVD 256(R29), F3
	MOVD 248(R29), F2
	MOVD 240(R29), F1
	MOVD 232(R29), F0
	MOVV 224(R29), R1
	MOVV R1, FCR31
	MOVV 216(R29), R1
	MOVV R1, LO
	MOVV 208(R29), R1
	MOVV R1, HI
	MOVV 200(R29), RSB
	MOVV 192(R29), R25
	MOVV 184(R29), R24
	MOVV 176(R29), R22
	MOVV 168(R29), R21
	MOVV 160(R29), R20
	MOVV 152(R29), R19
	MOVV 144(R29), R18
	MOVV 136(R29), R17
	MOVV 128(R29), R16
	MOVV 120(R29), R15
	MOVV 112(R29), R14
	MOVV 104(R29), R13
	MOVV 96(R29), R12
	MOVV 88(R29), R11
	MOVV 80(R29), R10
	MOVV 72(R29), R9
	MOVV 64(R29), R8
	MOVV 56(R29), R7
	MOVV 48(R29), R6
	MOVV 40(R29), R5
	MOVV 32(R29), R4
	MOVV 24(R29), R3
	MOVV 16(R29), R2
	MOVV 8(R29), R1
	MOVV 488(R29), R31
	MOVV (R29), R23
	ADDV $496, R29
	JMP (R23)
