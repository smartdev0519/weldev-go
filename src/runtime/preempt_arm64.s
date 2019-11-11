// Code generated by mkpreempt.go; DO NOT EDIT.

#include "go_asm.h"
#include "textflag.h"

TEXT ·asyncPreempt(SB),NOSPLIT|NOFRAME,$0-0
	MOVD R30, -496(RSP)
	SUB $496, RSP
	#ifdef GOOS_linux
	MOVD R29, -8(RSP)
	SUB $8, RSP, R29
	#endif
	MOVD R0, 8(RSP)
	MOVD R1, 16(RSP)
	MOVD R2, 24(RSP)
	MOVD R3, 32(RSP)
	MOVD R4, 40(RSP)
	MOVD R5, 48(RSP)
	MOVD R6, 56(RSP)
	MOVD R7, 64(RSP)
	MOVD R8, 72(RSP)
	MOVD R9, 80(RSP)
	MOVD R10, 88(RSP)
	MOVD R11, 96(RSP)
	MOVD R12, 104(RSP)
	MOVD R13, 112(RSP)
	MOVD R14, 120(RSP)
	MOVD R15, 128(RSP)
	MOVD R16, 136(RSP)
	MOVD R17, 144(RSP)
	MOVD R19, 152(RSP)
	MOVD R20, 160(RSP)
	MOVD R21, 168(RSP)
	MOVD R22, 176(RSP)
	MOVD R23, 184(RSP)
	MOVD R24, 192(RSP)
	MOVD R25, 200(RSP)
	MOVD R26, 208(RSP)
	MOVD NZCV, R0
	MOVD R0, 216(RSP)
	MOVD FPSR, R0
	MOVD R0, 224(RSP)
	FMOVD F0, 232(RSP)
	FMOVD F1, 240(RSP)
	FMOVD F2, 248(RSP)
	FMOVD F3, 256(RSP)
	FMOVD F4, 264(RSP)
	FMOVD F5, 272(RSP)
	FMOVD F6, 280(RSP)
	FMOVD F7, 288(RSP)
	FMOVD F8, 296(RSP)
	FMOVD F9, 304(RSP)
	FMOVD F10, 312(RSP)
	FMOVD F11, 320(RSP)
	FMOVD F12, 328(RSP)
	FMOVD F13, 336(RSP)
	FMOVD F14, 344(RSP)
	FMOVD F15, 352(RSP)
	FMOVD F16, 360(RSP)
	FMOVD F17, 368(RSP)
	FMOVD F18, 376(RSP)
	FMOVD F19, 384(RSP)
	FMOVD F20, 392(RSP)
	FMOVD F21, 400(RSP)
	FMOVD F22, 408(RSP)
	FMOVD F23, 416(RSP)
	FMOVD F24, 424(RSP)
	FMOVD F25, 432(RSP)
	FMOVD F26, 440(RSP)
	FMOVD F27, 448(RSP)
	FMOVD F28, 456(RSP)
	FMOVD F29, 464(RSP)
	FMOVD F30, 472(RSP)
	FMOVD F31, 480(RSP)
	CALL ·asyncPreempt2(SB)
	FMOVD 480(RSP), F31
	FMOVD 472(RSP), F30
	FMOVD 464(RSP), F29
	FMOVD 456(RSP), F28
	FMOVD 448(RSP), F27
	FMOVD 440(RSP), F26
	FMOVD 432(RSP), F25
	FMOVD 424(RSP), F24
	FMOVD 416(RSP), F23
	FMOVD 408(RSP), F22
	FMOVD 400(RSP), F21
	FMOVD 392(RSP), F20
	FMOVD 384(RSP), F19
	FMOVD 376(RSP), F18
	FMOVD 368(RSP), F17
	FMOVD 360(RSP), F16
	FMOVD 352(RSP), F15
	FMOVD 344(RSP), F14
	FMOVD 336(RSP), F13
	FMOVD 328(RSP), F12
	FMOVD 320(RSP), F11
	FMOVD 312(RSP), F10
	FMOVD 304(RSP), F9
	FMOVD 296(RSP), F8
	FMOVD 288(RSP), F7
	FMOVD 280(RSP), F6
	FMOVD 272(RSP), F5
	FMOVD 264(RSP), F4
	FMOVD 256(RSP), F3
	FMOVD 248(RSP), F2
	FMOVD 240(RSP), F1
	FMOVD 232(RSP), F0
	MOVD 224(RSP), R0
	MOVD R0, FPSR
	MOVD 216(RSP), R0
	MOVD R0, NZCV
	MOVD 208(RSP), R26
	MOVD 200(RSP), R25
	MOVD 192(RSP), R24
	MOVD 184(RSP), R23
	MOVD 176(RSP), R22
	MOVD 168(RSP), R21
	MOVD 160(RSP), R20
	MOVD 152(RSP), R19
	MOVD 144(RSP), R17
	MOVD 136(RSP), R16
	MOVD 128(RSP), R15
	MOVD 120(RSP), R14
	MOVD 112(RSP), R13
	MOVD 104(RSP), R12
	MOVD 96(RSP), R11
	MOVD 88(RSP), R10
	MOVD 80(RSP), R9
	MOVD 72(RSP), R8
	MOVD 64(RSP), R7
	MOVD 56(RSP), R6
	MOVD 48(RSP), R5
	MOVD 40(RSP), R4
	MOVD 32(RSP), R3
	MOVD 24(RSP), R2
	MOVD 16(RSP), R1
	MOVD 8(RSP), R0
	MOVD 496(RSP), R30
	#ifdef GOOS_linux
	MOVD -8(RSP), R29
	#endif
	MOVD (RSP), R27
	ADD $512, RSP
	JMP (R27)
