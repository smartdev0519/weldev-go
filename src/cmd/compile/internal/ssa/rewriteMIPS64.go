// Code generated from gen/MIPS64.rules; DO NOT EDIT.
// generated with: cd gen; go run *.go

package ssa

import "cmd/compile/internal/types"

func rewriteValueMIPS64(v *Value) bool {
	switch v.Op {
	case OpAdd16:
		v.Op = OpMIPS64ADDV
		return true
	case OpAdd32:
		v.Op = OpMIPS64ADDV
		return true
	case OpAdd32F:
		v.Op = OpMIPS64ADDF
		return true
	case OpAdd64:
		v.Op = OpMIPS64ADDV
		return true
	case OpAdd64F:
		v.Op = OpMIPS64ADDD
		return true
	case OpAdd8:
		v.Op = OpMIPS64ADDV
		return true
	case OpAddPtr:
		v.Op = OpMIPS64ADDV
		return true
	case OpAddr:
		v.Op = OpMIPS64MOVVaddr
		return true
	case OpAnd16:
		v.Op = OpMIPS64AND
		return true
	case OpAnd32:
		v.Op = OpMIPS64AND
		return true
	case OpAnd64:
		v.Op = OpMIPS64AND
		return true
	case OpAnd8:
		v.Op = OpMIPS64AND
		return true
	case OpAndB:
		v.Op = OpMIPS64AND
		return true
	case OpAtomicAdd32:
		v.Op = OpMIPS64LoweredAtomicAdd32
		return true
	case OpAtomicAdd64:
		v.Op = OpMIPS64LoweredAtomicAdd64
		return true
	case OpAtomicCompareAndSwap32:
		v.Op = OpMIPS64LoweredAtomicCas32
		return true
	case OpAtomicCompareAndSwap64:
		v.Op = OpMIPS64LoweredAtomicCas64
		return true
	case OpAtomicExchange32:
		v.Op = OpMIPS64LoweredAtomicExchange32
		return true
	case OpAtomicExchange64:
		v.Op = OpMIPS64LoweredAtomicExchange64
		return true
	case OpAtomicLoad32:
		v.Op = OpMIPS64LoweredAtomicLoad32
		return true
	case OpAtomicLoad64:
		v.Op = OpMIPS64LoweredAtomicLoad64
		return true
	case OpAtomicLoad8:
		v.Op = OpMIPS64LoweredAtomicLoad8
		return true
	case OpAtomicLoadPtr:
		v.Op = OpMIPS64LoweredAtomicLoad64
		return true
	case OpAtomicStore32:
		v.Op = OpMIPS64LoweredAtomicStore32
		return true
	case OpAtomicStore64:
		v.Op = OpMIPS64LoweredAtomicStore64
		return true
	case OpAtomicStore8:
		v.Op = OpMIPS64LoweredAtomicStore8
		return true
	case OpAtomicStorePtrNoWB:
		v.Op = OpMIPS64LoweredAtomicStore64
		return true
	case OpAvg64u:
		return rewriteValueMIPS64_OpAvg64u(v)
	case OpClosureCall:
		v.Op = OpMIPS64CALLclosure
		return true
	case OpCom16:
		return rewriteValueMIPS64_OpCom16(v)
	case OpCom32:
		return rewriteValueMIPS64_OpCom32(v)
	case OpCom64:
		return rewriteValueMIPS64_OpCom64(v)
	case OpCom8:
		return rewriteValueMIPS64_OpCom8(v)
	case OpConst16:
		v.Op = OpMIPS64MOVVconst
		return true
	case OpConst32:
		v.Op = OpMIPS64MOVVconst
		return true
	case OpConst32F:
		v.Op = OpMIPS64MOVFconst
		return true
	case OpConst64:
		v.Op = OpMIPS64MOVVconst
		return true
	case OpConst64F:
		v.Op = OpMIPS64MOVDconst
		return true
	case OpConst8:
		v.Op = OpMIPS64MOVVconst
		return true
	case OpConstBool:
		v.Op = OpMIPS64MOVVconst
		return true
	case OpConstNil:
		return rewriteValueMIPS64_OpConstNil(v)
	case OpCvt32Fto32:
		v.Op = OpMIPS64TRUNCFW
		return true
	case OpCvt32Fto64:
		v.Op = OpMIPS64TRUNCFV
		return true
	case OpCvt32Fto64F:
		v.Op = OpMIPS64MOVFD
		return true
	case OpCvt32to32F:
		v.Op = OpMIPS64MOVWF
		return true
	case OpCvt32to64F:
		v.Op = OpMIPS64MOVWD
		return true
	case OpCvt64Fto32:
		v.Op = OpMIPS64TRUNCDW
		return true
	case OpCvt64Fto32F:
		v.Op = OpMIPS64MOVDF
		return true
	case OpCvt64Fto64:
		v.Op = OpMIPS64TRUNCDV
		return true
	case OpCvt64to32F:
		v.Op = OpMIPS64MOVVF
		return true
	case OpCvt64to64F:
		v.Op = OpMIPS64MOVVD
		return true
	case OpDiv16:
		return rewriteValueMIPS64_OpDiv16(v)
	case OpDiv16u:
		return rewriteValueMIPS64_OpDiv16u(v)
	case OpDiv32:
		return rewriteValueMIPS64_OpDiv32(v)
	case OpDiv32F:
		v.Op = OpMIPS64DIVF
		return true
	case OpDiv32u:
		return rewriteValueMIPS64_OpDiv32u(v)
	case OpDiv64:
		return rewriteValueMIPS64_OpDiv64(v)
	case OpDiv64F:
		v.Op = OpMIPS64DIVD
		return true
	case OpDiv64u:
		return rewriteValueMIPS64_OpDiv64u(v)
	case OpDiv8:
		return rewriteValueMIPS64_OpDiv8(v)
	case OpDiv8u:
		return rewriteValueMIPS64_OpDiv8u(v)
	case OpEq16:
		return rewriteValueMIPS64_OpEq16(v)
	case OpEq32:
		return rewriteValueMIPS64_OpEq32(v)
	case OpEq32F:
		return rewriteValueMIPS64_OpEq32F(v)
	case OpEq64:
		return rewriteValueMIPS64_OpEq64(v)
	case OpEq64F:
		return rewriteValueMIPS64_OpEq64F(v)
	case OpEq8:
		return rewriteValueMIPS64_OpEq8(v)
	case OpEqB:
		return rewriteValueMIPS64_OpEqB(v)
	case OpEqPtr:
		return rewriteValueMIPS64_OpEqPtr(v)
	case OpGeq32F:
		return rewriteValueMIPS64_OpGeq32F(v)
	case OpGeq64F:
		return rewriteValueMIPS64_OpGeq64F(v)
	case OpGetCallerPC:
		v.Op = OpMIPS64LoweredGetCallerPC
		return true
	case OpGetCallerSP:
		v.Op = OpMIPS64LoweredGetCallerSP
		return true
	case OpGetClosurePtr:
		v.Op = OpMIPS64LoweredGetClosurePtr
		return true
	case OpGreater32F:
		return rewriteValueMIPS64_OpGreater32F(v)
	case OpGreater64F:
		return rewriteValueMIPS64_OpGreater64F(v)
	case OpHmul32:
		return rewriteValueMIPS64_OpHmul32(v)
	case OpHmul32u:
		return rewriteValueMIPS64_OpHmul32u(v)
	case OpHmul64:
		return rewriteValueMIPS64_OpHmul64(v)
	case OpHmul64u:
		return rewriteValueMIPS64_OpHmul64u(v)
	case OpInterCall:
		v.Op = OpMIPS64CALLinter
		return true
	case OpIsInBounds:
		return rewriteValueMIPS64_OpIsInBounds(v)
	case OpIsNonNil:
		return rewriteValueMIPS64_OpIsNonNil(v)
	case OpIsSliceInBounds:
		return rewriteValueMIPS64_OpIsSliceInBounds(v)
	case OpLeq16:
		return rewriteValueMIPS64_OpLeq16(v)
	case OpLeq16U:
		return rewriteValueMIPS64_OpLeq16U(v)
	case OpLeq32:
		return rewriteValueMIPS64_OpLeq32(v)
	case OpLeq32F:
		return rewriteValueMIPS64_OpLeq32F(v)
	case OpLeq32U:
		return rewriteValueMIPS64_OpLeq32U(v)
	case OpLeq64:
		return rewriteValueMIPS64_OpLeq64(v)
	case OpLeq64F:
		return rewriteValueMIPS64_OpLeq64F(v)
	case OpLeq64U:
		return rewriteValueMIPS64_OpLeq64U(v)
	case OpLeq8:
		return rewriteValueMIPS64_OpLeq8(v)
	case OpLeq8U:
		return rewriteValueMIPS64_OpLeq8U(v)
	case OpLess16:
		return rewriteValueMIPS64_OpLess16(v)
	case OpLess16U:
		return rewriteValueMIPS64_OpLess16U(v)
	case OpLess32:
		return rewriteValueMIPS64_OpLess32(v)
	case OpLess32F:
		return rewriteValueMIPS64_OpLess32F(v)
	case OpLess32U:
		return rewriteValueMIPS64_OpLess32U(v)
	case OpLess64:
		return rewriteValueMIPS64_OpLess64(v)
	case OpLess64F:
		return rewriteValueMIPS64_OpLess64F(v)
	case OpLess64U:
		return rewriteValueMIPS64_OpLess64U(v)
	case OpLess8:
		return rewriteValueMIPS64_OpLess8(v)
	case OpLess8U:
		return rewriteValueMIPS64_OpLess8U(v)
	case OpLoad:
		return rewriteValueMIPS64_OpLoad(v)
	case OpLocalAddr:
		return rewriteValueMIPS64_OpLocalAddr(v)
	case OpLsh16x16:
		return rewriteValueMIPS64_OpLsh16x16(v)
	case OpLsh16x32:
		return rewriteValueMIPS64_OpLsh16x32(v)
	case OpLsh16x64:
		return rewriteValueMIPS64_OpLsh16x64(v)
	case OpLsh16x8:
		return rewriteValueMIPS64_OpLsh16x8(v)
	case OpLsh32x16:
		return rewriteValueMIPS64_OpLsh32x16(v)
	case OpLsh32x32:
		return rewriteValueMIPS64_OpLsh32x32(v)
	case OpLsh32x64:
		return rewriteValueMIPS64_OpLsh32x64(v)
	case OpLsh32x8:
		return rewriteValueMIPS64_OpLsh32x8(v)
	case OpLsh64x16:
		return rewriteValueMIPS64_OpLsh64x16(v)
	case OpLsh64x32:
		return rewriteValueMIPS64_OpLsh64x32(v)
	case OpLsh64x64:
		return rewriteValueMIPS64_OpLsh64x64(v)
	case OpLsh64x8:
		return rewriteValueMIPS64_OpLsh64x8(v)
	case OpLsh8x16:
		return rewriteValueMIPS64_OpLsh8x16(v)
	case OpLsh8x32:
		return rewriteValueMIPS64_OpLsh8x32(v)
	case OpLsh8x64:
		return rewriteValueMIPS64_OpLsh8x64(v)
	case OpLsh8x8:
		return rewriteValueMIPS64_OpLsh8x8(v)
	case OpMIPS64ADDV:
		return rewriteValueMIPS64_OpMIPS64ADDV(v)
	case OpMIPS64ADDVconst:
		return rewriteValueMIPS64_OpMIPS64ADDVconst(v)
	case OpMIPS64AND:
		return rewriteValueMIPS64_OpMIPS64AND(v)
	case OpMIPS64ANDconst:
		return rewriteValueMIPS64_OpMIPS64ANDconst(v)
	case OpMIPS64LoweredAtomicAdd32:
		return rewriteValueMIPS64_OpMIPS64LoweredAtomicAdd32(v)
	case OpMIPS64LoweredAtomicAdd64:
		return rewriteValueMIPS64_OpMIPS64LoweredAtomicAdd64(v)
	case OpMIPS64LoweredAtomicStore32:
		return rewriteValueMIPS64_OpMIPS64LoweredAtomicStore32(v)
	case OpMIPS64LoweredAtomicStore64:
		return rewriteValueMIPS64_OpMIPS64LoweredAtomicStore64(v)
	case OpMIPS64MOVBUload:
		return rewriteValueMIPS64_OpMIPS64MOVBUload(v)
	case OpMIPS64MOVBUreg:
		return rewriteValueMIPS64_OpMIPS64MOVBUreg(v)
	case OpMIPS64MOVBload:
		return rewriteValueMIPS64_OpMIPS64MOVBload(v)
	case OpMIPS64MOVBreg:
		return rewriteValueMIPS64_OpMIPS64MOVBreg(v)
	case OpMIPS64MOVBstore:
		return rewriteValueMIPS64_OpMIPS64MOVBstore(v)
	case OpMIPS64MOVBstorezero:
		return rewriteValueMIPS64_OpMIPS64MOVBstorezero(v)
	case OpMIPS64MOVDload:
		return rewriteValueMIPS64_OpMIPS64MOVDload(v)
	case OpMIPS64MOVDstore:
		return rewriteValueMIPS64_OpMIPS64MOVDstore(v)
	case OpMIPS64MOVFload:
		return rewriteValueMIPS64_OpMIPS64MOVFload(v)
	case OpMIPS64MOVFstore:
		return rewriteValueMIPS64_OpMIPS64MOVFstore(v)
	case OpMIPS64MOVHUload:
		return rewriteValueMIPS64_OpMIPS64MOVHUload(v)
	case OpMIPS64MOVHUreg:
		return rewriteValueMIPS64_OpMIPS64MOVHUreg(v)
	case OpMIPS64MOVHload:
		return rewriteValueMIPS64_OpMIPS64MOVHload(v)
	case OpMIPS64MOVHreg:
		return rewriteValueMIPS64_OpMIPS64MOVHreg(v)
	case OpMIPS64MOVHstore:
		return rewriteValueMIPS64_OpMIPS64MOVHstore(v)
	case OpMIPS64MOVHstorezero:
		return rewriteValueMIPS64_OpMIPS64MOVHstorezero(v)
	case OpMIPS64MOVVload:
		return rewriteValueMIPS64_OpMIPS64MOVVload(v)
	case OpMIPS64MOVVreg:
		return rewriteValueMIPS64_OpMIPS64MOVVreg(v)
	case OpMIPS64MOVVstore:
		return rewriteValueMIPS64_OpMIPS64MOVVstore(v)
	case OpMIPS64MOVVstorezero:
		return rewriteValueMIPS64_OpMIPS64MOVVstorezero(v)
	case OpMIPS64MOVWUload:
		return rewriteValueMIPS64_OpMIPS64MOVWUload(v)
	case OpMIPS64MOVWUreg:
		return rewriteValueMIPS64_OpMIPS64MOVWUreg(v)
	case OpMIPS64MOVWload:
		return rewriteValueMIPS64_OpMIPS64MOVWload(v)
	case OpMIPS64MOVWreg:
		return rewriteValueMIPS64_OpMIPS64MOVWreg(v)
	case OpMIPS64MOVWstore:
		return rewriteValueMIPS64_OpMIPS64MOVWstore(v)
	case OpMIPS64MOVWstorezero:
		return rewriteValueMIPS64_OpMIPS64MOVWstorezero(v)
	case OpMIPS64NEGV:
		return rewriteValueMIPS64_OpMIPS64NEGV(v)
	case OpMIPS64NOR:
		return rewriteValueMIPS64_OpMIPS64NOR(v)
	case OpMIPS64NORconst:
		return rewriteValueMIPS64_OpMIPS64NORconst(v)
	case OpMIPS64OR:
		return rewriteValueMIPS64_OpMIPS64OR(v)
	case OpMIPS64ORconst:
		return rewriteValueMIPS64_OpMIPS64ORconst(v)
	case OpMIPS64SGT:
		return rewriteValueMIPS64_OpMIPS64SGT(v)
	case OpMIPS64SGTU:
		return rewriteValueMIPS64_OpMIPS64SGTU(v)
	case OpMIPS64SGTUconst:
		return rewriteValueMIPS64_OpMIPS64SGTUconst(v)
	case OpMIPS64SGTconst:
		return rewriteValueMIPS64_OpMIPS64SGTconst(v)
	case OpMIPS64SLLV:
		return rewriteValueMIPS64_OpMIPS64SLLV(v)
	case OpMIPS64SLLVconst:
		return rewriteValueMIPS64_OpMIPS64SLLVconst(v)
	case OpMIPS64SRAV:
		return rewriteValueMIPS64_OpMIPS64SRAV(v)
	case OpMIPS64SRAVconst:
		return rewriteValueMIPS64_OpMIPS64SRAVconst(v)
	case OpMIPS64SRLV:
		return rewriteValueMIPS64_OpMIPS64SRLV(v)
	case OpMIPS64SRLVconst:
		return rewriteValueMIPS64_OpMIPS64SRLVconst(v)
	case OpMIPS64SUBV:
		return rewriteValueMIPS64_OpMIPS64SUBV(v)
	case OpMIPS64SUBVconst:
		return rewriteValueMIPS64_OpMIPS64SUBVconst(v)
	case OpMIPS64XOR:
		return rewriteValueMIPS64_OpMIPS64XOR(v)
	case OpMIPS64XORconst:
		return rewriteValueMIPS64_OpMIPS64XORconst(v)
	case OpMod16:
		return rewriteValueMIPS64_OpMod16(v)
	case OpMod16u:
		return rewriteValueMIPS64_OpMod16u(v)
	case OpMod32:
		return rewriteValueMIPS64_OpMod32(v)
	case OpMod32u:
		return rewriteValueMIPS64_OpMod32u(v)
	case OpMod64:
		return rewriteValueMIPS64_OpMod64(v)
	case OpMod64u:
		return rewriteValueMIPS64_OpMod64u(v)
	case OpMod8:
		return rewriteValueMIPS64_OpMod8(v)
	case OpMod8u:
		return rewriteValueMIPS64_OpMod8u(v)
	case OpMove:
		return rewriteValueMIPS64_OpMove(v)
	case OpMul16:
		return rewriteValueMIPS64_OpMul16(v)
	case OpMul32:
		return rewriteValueMIPS64_OpMul32(v)
	case OpMul32F:
		v.Op = OpMIPS64MULF
		return true
	case OpMul64:
		return rewriteValueMIPS64_OpMul64(v)
	case OpMul64F:
		v.Op = OpMIPS64MULD
		return true
	case OpMul64uhilo:
		v.Op = OpMIPS64MULVU
		return true
	case OpMul8:
		return rewriteValueMIPS64_OpMul8(v)
	case OpNeg16:
		v.Op = OpMIPS64NEGV
		return true
	case OpNeg32:
		v.Op = OpMIPS64NEGV
		return true
	case OpNeg32F:
		v.Op = OpMIPS64NEGF
		return true
	case OpNeg64:
		v.Op = OpMIPS64NEGV
		return true
	case OpNeg64F:
		v.Op = OpMIPS64NEGD
		return true
	case OpNeg8:
		v.Op = OpMIPS64NEGV
		return true
	case OpNeq16:
		return rewriteValueMIPS64_OpNeq16(v)
	case OpNeq32:
		return rewriteValueMIPS64_OpNeq32(v)
	case OpNeq32F:
		return rewriteValueMIPS64_OpNeq32F(v)
	case OpNeq64:
		return rewriteValueMIPS64_OpNeq64(v)
	case OpNeq64F:
		return rewriteValueMIPS64_OpNeq64F(v)
	case OpNeq8:
		return rewriteValueMIPS64_OpNeq8(v)
	case OpNeqB:
		v.Op = OpMIPS64XOR
		return true
	case OpNeqPtr:
		return rewriteValueMIPS64_OpNeqPtr(v)
	case OpNilCheck:
		v.Op = OpMIPS64LoweredNilCheck
		return true
	case OpNot:
		return rewriteValueMIPS64_OpNot(v)
	case OpOffPtr:
		return rewriteValueMIPS64_OpOffPtr(v)
	case OpOr16:
		v.Op = OpMIPS64OR
		return true
	case OpOr32:
		v.Op = OpMIPS64OR
		return true
	case OpOr64:
		v.Op = OpMIPS64OR
		return true
	case OpOr8:
		v.Op = OpMIPS64OR
		return true
	case OpOrB:
		v.Op = OpMIPS64OR
		return true
	case OpPanicBounds:
		return rewriteValueMIPS64_OpPanicBounds(v)
	case OpRotateLeft16:
		return rewriteValueMIPS64_OpRotateLeft16(v)
	case OpRotateLeft32:
		return rewriteValueMIPS64_OpRotateLeft32(v)
	case OpRotateLeft64:
		return rewriteValueMIPS64_OpRotateLeft64(v)
	case OpRotateLeft8:
		return rewriteValueMIPS64_OpRotateLeft8(v)
	case OpRound32F:
		v.Op = OpCopy
		return true
	case OpRound64F:
		v.Op = OpCopy
		return true
	case OpRsh16Ux16:
		return rewriteValueMIPS64_OpRsh16Ux16(v)
	case OpRsh16Ux32:
		return rewriteValueMIPS64_OpRsh16Ux32(v)
	case OpRsh16Ux64:
		return rewriteValueMIPS64_OpRsh16Ux64(v)
	case OpRsh16Ux8:
		return rewriteValueMIPS64_OpRsh16Ux8(v)
	case OpRsh16x16:
		return rewriteValueMIPS64_OpRsh16x16(v)
	case OpRsh16x32:
		return rewriteValueMIPS64_OpRsh16x32(v)
	case OpRsh16x64:
		return rewriteValueMIPS64_OpRsh16x64(v)
	case OpRsh16x8:
		return rewriteValueMIPS64_OpRsh16x8(v)
	case OpRsh32Ux16:
		return rewriteValueMIPS64_OpRsh32Ux16(v)
	case OpRsh32Ux32:
		return rewriteValueMIPS64_OpRsh32Ux32(v)
	case OpRsh32Ux64:
		return rewriteValueMIPS64_OpRsh32Ux64(v)
	case OpRsh32Ux8:
		return rewriteValueMIPS64_OpRsh32Ux8(v)
	case OpRsh32x16:
		return rewriteValueMIPS64_OpRsh32x16(v)
	case OpRsh32x32:
		return rewriteValueMIPS64_OpRsh32x32(v)
	case OpRsh32x64:
		return rewriteValueMIPS64_OpRsh32x64(v)
	case OpRsh32x8:
		return rewriteValueMIPS64_OpRsh32x8(v)
	case OpRsh64Ux16:
		return rewriteValueMIPS64_OpRsh64Ux16(v)
	case OpRsh64Ux32:
		return rewriteValueMIPS64_OpRsh64Ux32(v)
	case OpRsh64Ux64:
		return rewriteValueMIPS64_OpRsh64Ux64(v)
	case OpRsh64Ux8:
		return rewriteValueMIPS64_OpRsh64Ux8(v)
	case OpRsh64x16:
		return rewriteValueMIPS64_OpRsh64x16(v)
	case OpRsh64x32:
		return rewriteValueMIPS64_OpRsh64x32(v)
	case OpRsh64x64:
		return rewriteValueMIPS64_OpRsh64x64(v)
	case OpRsh64x8:
		return rewriteValueMIPS64_OpRsh64x8(v)
	case OpRsh8Ux16:
		return rewriteValueMIPS64_OpRsh8Ux16(v)
	case OpRsh8Ux32:
		return rewriteValueMIPS64_OpRsh8Ux32(v)
	case OpRsh8Ux64:
		return rewriteValueMIPS64_OpRsh8Ux64(v)
	case OpRsh8Ux8:
		return rewriteValueMIPS64_OpRsh8Ux8(v)
	case OpRsh8x16:
		return rewriteValueMIPS64_OpRsh8x16(v)
	case OpRsh8x32:
		return rewriteValueMIPS64_OpRsh8x32(v)
	case OpRsh8x64:
		return rewriteValueMIPS64_OpRsh8x64(v)
	case OpRsh8x8:
		return rewriteValueMIPS64_OpRsh8x8(v)
	case OpSelect0:
		return rewriteValueMIPS64_OpSelect0(v)
	case OpSelect1:
		return rewriteValueMIPS64_OpSelect1(v)
	case OpSignExt16to32:
		v.Op = OpMIPS64MOVHreg
		return true
	case OpSignExt16to64:
		v.Op = OpMIPS64MOVHreg
		return true
	case OpSignExt32to64:
		v.Op = OpMIPS64MOVWreg
		return true
	case OpSignExt8to16:
		v.Op = OpMIPS64MOVBreg
		return true
	case OpSignExt8to32:
		v.Op = OpMIPS64MOVBreg
		return true
	case OpSignExt8to64:
		v.Op = OpMIPS64MOVBreg
		return true
	case OpSlicemask:
		return rewriteValueMIPS64_OpSlicemask(v)
	case OpSqrt:
		v.Op = OpMIPS64SQRTD
		return true
	case OpStaticCall:
		v.Op = OpMIPS64CALLstatic
		return true
	case OpStore:
		return rewriteValueMIPS64_OpStore(v)
	case OpSub16:
		v.Op = OpMIPS64SUBV
		return true
	case OpSub32:
		v.Op = OpMIPS64SUBV
		return true
	case OpSub32F:
		v.Op = OpMIPS64SUBF
		return true
	case OpSub64:
		v.Op = OpMIPS64SUBV
		return true
	case OpSub64F:
		v.Op = OpMIPS64SUBD
		return true
	case OpSub8:
		v.Op = OpMIPS64SUBV
		return true
	case OpSubPtr:
		v.Op = OpMIPS64SUBV
		return true
	case OpTrunc16to8:
		v.Op = OpCopy
		return true
	case OpTrunc32to16:
		v.Op = OpCopy
		return true
	case OpTrunc32to8:
		v.Op = OpCopy
		return true
	case OpTrunc64to16:
		v.Op = OpCopy
		return true
	case OpTrunc64to32:
		v.Op = OpCopy
		return true
	case OpTrunc64to8:
		v.Op = OpCopy
		return true
	case OpWB:
		v.Op = OpMIPS64LoweredWB
		return true
	case OpXor16:
		v.Op = OpMIPS64XOR
		return true
	case OpXor32:
		v.Op = OpMIPS64XOR
		return true
	case OpXor64:
		v.Op = OpMIPS64XOR
		return true
	case OpXor8:
		v.Op = OpMIPS64XOR
		return true
	case OpZero:
		return rewriteValueMIPS64_OpZero(v)
	case OpZeroExt16to32:
		v.Op = OpMIPS64MOVHUreg
		return true
	case OpZeroExt16to64:
		v.Op = OpMIPS64MOVHUreg
		return true
	case OpZeroExt32to64:
		v.Op = OpMIPS64MOVWUreg
		return true
	case OpZeroExt8to16:
		v.Op = OpMIPS64MOVBUreg
		return true
	case OpZeroExt8to32:
		v.Op = OpMIPS64MOVBUreg
		return true
	case OpZeroExt8to64:
		v.Op = OpMIPS64MOVBUreg
		return true
	}
	return false
}
func rewriteValueMIPS64_OpAvg64u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Avg64u <t> x y)
	// result: (ADDV (SRLVconst <t> (SUBV <t> x y) [1]) y)
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64ADDV)
		v0 := b.NewValue0(v.Pos, OpMIPS64SRLVconst, t)
		v0.AuxInt = 1
		v1 := b.NewValue0(v.Pos, OpMIPS64SUBV, t)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
}
func rewriteValueMIPS64_OpCom16(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Com16 x)
	// result: (NOR (MOVVconst [0]) x)
	for {
		x := v_0
		v.reset(OpMIPS64NOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpCom32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Com32 x)
	// result: (NOR (MOVVconst [0]) x)
	for {
		x := v_0
		v.reset(OpMIPS64NOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpCom64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Com64 x)
	// result: (NOR (MOVVconst [0]) x)
	for {
		x := v_0
		v.reset(OpMIPS64NOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpCom8(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Com8 x)
	// result: (NOR (MOVVconst [0]) x)
	for {
		x := v_0
		v.reset(OpMIPS64NOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpConstNil(v *Value) bool {
	// match: (ConstNil)
	// result: (MOVVconst [0])
	for {
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}
}
func rewriteValueMIPS64_OpDiv16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div16 x y)
	// result: (Select1 (DIVV (SignExt16to64 x) (SignExt16to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVV, types.NewTuple(typ.Int64, typ.Int64))
		v1 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpDiv16u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div16u x y)
	// result: (Select1 (DIVVU (ZeroExt16to64 x) (ZeroExt16to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVVU, types.NewTuple(typ.UInt64, typ.UInt64))
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpDiv32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div32 x y)
	// result: (Select1 (DIVV (SignExt32to64 x) (SignExt32to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVV, types.NewTuple(typ.Int64, typ.Int64))
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpDiv32u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div32u x y)
	// result: (Select1 (DIVVU (ZeroExt32to64 x) (ZeroExt32to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVVU, types.NewTuple(typ.UInt64, typ.UInt64))
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpDiv64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div64 x y)
	// result: (Select1 (DIVV x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVV, types.NewTuple(typ.Int64, typ.Int64))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpDiv64u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div64u x y)
	// result: (Select1 (DIVVU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVVU, types.NewTuple(typ.UInt64, typ.UInt64))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpDiv8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div8 x y)
	// result: (Select1 (DIVV (SignExt8to64 x) (SignExt8to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVV, types.NewTuple(typ.Int64, typ.Int64))
		v1 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpDiv8u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div8u x y)
	// result: (Select1 (DIVVU (ZeroExt8to64 x) (ZeroExt8to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVVU, types.NewTuple(typ.UInt64, typ.UInt64))
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpEq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq16 x y)
	// result: (SGTU (MOVVconst [1]) (XOR (ZeroExt16to64 x) (ZeroExt16to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64XOR, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpEq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq32 x y)
	// result: (SGTU (MOVVconst [1]) (XOR (ZeroExt32to64 x) (ZeroExt32to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64XOR, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpEq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Eq32F x y)
	// result: (FPFlagTrue (CMPEQF x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64FPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPEQF, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpEq64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq64 x y)
	// result: (SGTU (MOVVconst [1]) (XOR x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64XOR, typ.UInt64)
		v1.AddArg(x)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpEq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Eq64F x y)
	// result: (FPFlagTrue (CMPEQD x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64FPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPEQD, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpEq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq8 x y)
	// result: (SGTU (MOVVconst [1]) (XOR (ZeroExt8to64 x) (ZeroExt8to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64XOR, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpEqB(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (EqB x y)
	// result: (XOR (MOVVconst [1]) (XOR <typ.Bool> x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64XOR, typ.Bool)
		v1.AddArg(x)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpEqPtr(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (EqPtr x y)
	// result: (SGTU (MOVVconst [1]) (XOR x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64XOR, typ.UInt64)
		v1.AddArg(x)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpGeq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Geq32F x y)
	// result: (FPFlagTrue (CMPGEF x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64FPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPGEF, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpGeq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Geq64F x y)
	// result: (FPFlagTrue (CMPGED x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64FPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPGED, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpGreater32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Greater32F x y)
	// result: (FPFlagTrue (CMPGTF x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64FPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPGTF, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpGreater64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Greater64F x y)
	// result: (FPFlagTrue (CMPGTD x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64FPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPGTD, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpHmul32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Hmul32 x y)
	// result: (SRAVconst (Select1 <typ.Int64> (MULV (SignExt32to64 x) (SignExt32to64 y))) [32])
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64SRAVconst)
		v.AuxInt = 32
		v0 := b.NewValue0(v.Pos, OpSelect1, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpMIPS64MULV, types.NewTuple(typ.Int64, typ.Int64))
		v2 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpHmul32u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Hmul32u x y)
	// result: (SRLVconst (Select1 <typ.UInt64> (MULVU (ZeroExt32to64 x) (ZeroExt32to64 y))) [32])
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64SRLVconst)
		v.AuxInt = 32
		v0 := b.NewValue0(v.Pos, OpSelect1, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpMIPS64MULVU, types.NewTuple(typ.UInt64, typ.UInt64))
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpHmul64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Hmul64 x y)
	// result: (Select0 (MULV x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPS64MULV, types.NewTuple(typ.Int64, typ.Int64))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpHmul64u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Hmul64u x y)
	// result: (Select0 (MULVU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPS64MULVU, types.NewTuple(typ.UInt64, typ.UInt64))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpIsInBounds(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (IsInBounds idx len)
	// result: (SGTU len idx)
	for {
		idx := v_0
		len := v_1
		v.reset(OpMIPS64SGTU)
		v.AddArg(len)
		v.AddArg(idx)
		return true
	}
}
func rewriteValueMIPS64_OpIsNonNil(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (IsNonNil ptr)
	// result: (SGTU ptr (MOVVconst [0]))
	for {
		ptr := v_0
		v.reset(OpMIPS64SGTU)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpIsSliceInBounds(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (IsSliceInBounds idx len)
	// result: (XOR (MOVVconst [1]) (SGTU idx len))
	for {
		idx := v_0
		len := v_1
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v1.AddArg(idx)
		v1.AddArg(len)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLeq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq16 x y)
	// result: (XOR (MOVVconst [1]) (SGT (SignExt16to64 x) (SignExt16to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLeq16U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq16U x y)
	// result: (XOR (MOVVconst [1]) (SGTU (ZeroExt16to64 x) (ZeroExt16to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLeq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq32 x y)
	// result: (XOR (MOVVconst [1]) (SGT (SignExt32to64 x) (SignExt32to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLeq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Leq32F x y)
	// result: (FPFlagTrue (CMPGEF y x))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64FPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPGEF, types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpLeq32U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq32U x y)
	// result: (XOR (MOVVconst [1]) (SGTU (ZeroExt32to64 x) (ZeroExt32to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLeq64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq64 x y)
	// result: (XOR (MOVVconst [1]) (SGT x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGT, typ.Bool)
		v1.AddArg(x)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLeq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Leq64F x y)
	// result: (FPFlagTrue (CMPGED y x))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64FPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPGED, types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpLeq64U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq64U x y)
	// result: (XOR (MOVVconst [1]) (SGTU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v1.AddArg(x)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLeq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq8 x y)
	// result: (XOR (MOVVconst [1]) (SGT (SignExt8to64 x) (SignExt8to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGT, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLeq8U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq8U x y)
	// result: (XOR (MOVVconst [1]) (SGTU (ZeroExt8to64 x) (ZeroExt8to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64XOR)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 1
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v2.AddArg(x)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLess16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less16 x y)
	// result: (SGT (SignExt16to64 y) (SignExt16to64 x))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64SGT)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLess16U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less16U x y)
	// result: (SGTU (ZeroExt16to64 y) (ZeroExt16to64 x))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLess32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less32 x y)
	// result: (SGT (SignExt32to64 y) (SignExt32to64 x))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64SGT)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLess32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Less32F x y)
	// result: (FPFlagTrue (CMPGTF y x))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64FPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPGTF, types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpLess32U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less32U x y)
	// result: (SGTU (ZeroExt32to64 y) (ZeroExt32to64 x))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLess64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Less64 x y)
	// result: (SGT y x)
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64SGT)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpLess64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Less64F x y)
	// result: (FPFlagTrue (CMPGTD y x))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64FPFlagTrue)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPGTD, types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpLess64U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Less64U x y)
	// result: (SGTU y x)
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64SGTU)
		v.AddArg(y)
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpLess8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less8 x y)
	// result: (SGT (SignExt8to64 y) (SignExt8to64 x))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64SGT)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLess8U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less8U x y)
	// result: (SGTU (ZeroExt8to64 y) (ZeroExt8to64 x))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(x)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpLoad(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Load <t> ptr mem)
	// cond: t.IsBoolean()
	// result: (MOVBUload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(t.IsBoolean()) {
			break
		}
		v.reset(OpMIPS64MOVBUload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is8BitInt(t) && isSigned(t))
	// result: (MOVBload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is8BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpMIPS64MOVBload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is8BitInt(t) && !isSigned(t))
	// result: (MOVBUload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is8BitInt(t) && !isSigned(t)) {
			break
		}
		v.reset(OpMIPS64MOVBUload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is16BitInt(t) && isSigned(t))
	// result: (MOVHload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is16BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpMIPS64MOVHload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is16BitInt(t) && !isSigned(t))
	// result: (MOVHUload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is16BitInt(t) && !isSigned(t)) {
			break
		}
		v.reset(OpMIPS64MOVHUload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is32BitInt(t) && isSigned(t))
	// result: (MOVWload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is32BitInt(t) && isSigned(t)) {
			break
		}
		v.reset(OpMIPS64MOVWload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is32BitInt(t) && !isSigned(t))
	// result: (MOVWUload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is32BitInt(t) && !isSigned(t)) {
			break
		}
		v.reset(OpMIPS64MOVWUload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is64BitInt(t) || isPtr(t))
	// result: (MOVVload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is64BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(OpMIPS64MOVVload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is32BitFloat(t)
	// result: (MOVFload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is32BitFloat(t)) {
			break
		}
		v.reset(OpMIPS64MOVFload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is64BitFloat(t)
	// result: (MOVDload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is64BitFloat(t)) {
			break
		}
		v.reset(OpMIPS64MOVDload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpLocalAddr(v *Value) bool {
	v_0 := v.Args[0]
	// match: (LocalAddr {sym} base _)
	// result: (MOVVaddr {sym} base)
	for {
		sym := v.Aux
		base := v_0
		v.reset(OpMIPS64MOVVaddr)
		v.Aux = sym
		v.AddArg(base)
		return true
	}
}
func rewriteValueMIPS64_OpLsh16x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x16 <t> x y)
	// result: (AND (NEGV <t> (SGTU (MOVVconst <typ.UInt64> [64]) (ZeroExt16to64 y))) (SLLV <t> x (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpLsh16x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x32 <t> x y)
	// result: (AND (NEGV <t> (SGTU (MOVVconst <typ.UInt64> [64]) (ZeroExt32to64 y))) (SLLV <t> x (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpLsh16x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x64 <t> x y)
	// result: (AND (NEGV <t> (SGTU (MOVVconst <typ.UInt64> [64]) y)) (SLLV <t> x y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v3.AddArg(x)
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS64_OpLsh16x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x8 <t> x y)
	// result: (AND (NEGV <t> (SGTU (MOVVconst <typ.UInt64> [64]) (ZeroExt8to64 y))) (SLLV <t> x (ZeroExt8to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpLsh32x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x16 <t> x y)
	// result: (AND (NEGV <t> (SGTU (MOVVconst <typ.UInt64> [64]) (ZeroExt16to64 y))) (SLLV <t> x (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpLsh32x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x32 <t> x y)
	// result: (AND (NEGV <t> (SGTU (MOVVconst <typ.UInt64> [64]) (ZeroExt32to64 y))) (SLLV <t> x (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpLsh32x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x64 <t> x y)
	// result: (AND (NEGV <t> (SGTU (MOVVconst <typ.UInt64> [64]) y)) (SLLV <t> x y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v3.AddArg(x)
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS64_OpLsh32x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x8 <t> x y)
	// result: (AND (NEGV <t> (SGTU (MOVVconst <typ.UInt64> [64]) (ZeroExt8to64 y))) (SLLV <t> x (ZeroExt8to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpLsh64x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x16 <t> x y)
	// result: (AND (NEGV <t> (SGTU (MOVVconst <typ.UInt64> [64]) (ZeroExt16to64 y))) (SLLV <t> x (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpLsh64x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x32 <t> x y)
	// result: (AND (NEGV <t> (SGTU (MOVVconst <typ.UInt64> [64]) (ZeroExt32to64 y))) (SLLV <t> x (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpLsh64x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x64 <t> x y)
	// result: (AND (NEGV <t> (SGTU (MOVVconst <typ.UInt64> [64]) y)) (SLLV <t> x y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v3.AddArg(x)
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS64_OpLsh64x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x8 <t> x y)
	// result: (AND (NEGV <t> (SGTU (MOVVconst <typ.UInt64> [64]) (ZeroExt8to64 y))) (SLLV <t> x (ZeroExt8to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpLsh8x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x16 <t> x y)
	// result: (AND (NEGV <t> (SGTU (MOVVconst <typ.UInt64> [64]) (ZeroExt16to64 y))) (SLLV <t> x (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpLsh8x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x32 <t> x y)
	// result: (AND (NEGV <t> (SGTU (MOVVconst <typ.UInt64> [64]) (ZeroExt32to64 y))) (SLLV <t> x (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpLsh8x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x64 <t> x y)
	// result: (AND (NEGV <t> (SGTU (MOVVconst <typ.UInt64> [64]) y)) (SLLV <t> x y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v3.AddArg(x)
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS64_OpLsh8x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x8 <t> x y)
	// result: (AND (NEGV <t> (SGTU (MOVVconst <typ.UInt64> [64]) (ZeroExt8to64 y))) (SLLV <t> x (ZeroExt8to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SLLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpMIPS64ADDV(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ADDV x (MOVVconst [c]))
	// cond: is32Bit(c)
	// result: (ADDVconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpMIPS64MOVVconst {
				continue
			}
			c := v_1.AuxInt
			if !(is32Bit(c)) {
				continue
			}
			v.reset(OpMIPS64ADDVconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (ADDV x (NEGV y))
	// result: (SUBV x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpMIPS64NEGV {
				continue
			}
			y := v_1.Args[0]
			v.reset(OpMIPS64SUBV)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64ADDVconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ADDVconst [off1] (MOVVaddr [off2] {sym} ptr))
	// result: (MOVVaddr [off1+off2] {sym} ptr)
	for {
		off1 := v.AuxInt
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym := v_0.Aux
		ptr := v_0.Args[0]
		v.reset(OpMIPS64MOVVaddr)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		return true
	}
	// match: (ADDVconst [0] x)
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v_0
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ADDVconst [c] (MOVVconst [d]))
	// result: (MOVVconst [c+d])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = c + d
		return true
	}
	// match: (ADDVconst [c] (ADDVconst [d] x))
	// cond: is32Bit(c+d)
	// result: (ADDVconst [c+d] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(c + d)) {
			break
		}
		v.reset(OpMIPS64ADDVconst)
		v.AuxInt = c + d
		v.AddArg(x)
		return true
	}
	// match: (ADDVconst [c] (SUBVconst [d] x))
	// cond: is32Bit(c-d)
	// result: (ADDVconst [c-d] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64SUBVconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(c - d)) {
			break
		}
		v.reset(OpMIPS64ADDVconst)
		v.AuxInt = c - d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64AND(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (AND x (MOVVconst [c]))
	// cond: is32Bit(c)
	// result: (ANDconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpMIPS64MOVVconst {
				continue
			}
			c := v_1.AuxInt
			if !(is32Bit(c)) {
				continue
			}
			v.reset(OpMIPS64ANDconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (AND x x)
	// result: x
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64ANDconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ANDconst [0] _)
	// result: (MOVVconst [0])
	for {
		if v.AuxInt != 0 {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}
	// match: (ANDconst [-1] x)
	// result: x
	for {
		if v.AuxInt != -1 {
			break
		}
		x := v_0
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ANDconst [c] (MOVVconst [d]))
	// result: (MOVVconst [c&d])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = c & d
		return true
	}
	// match: (ANDconst [c] (ANDconst [d] x))
	// result: (ANDconst [c&d] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64ANDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpMIPS64ANDconst)
		v.AuxInt = c & d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64LoweredAtomicAdd32(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (LoweredAtomicAdd32 ptr (MOVVconst [c]) mem)
	// cond: is32Bit(c)
	// result: (LoweredAtomicAddconst32 [c] ptr mem)
	for {
		ptr := v_0
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		mem := v_2
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpMIPS64LoweredAtomicAddconst32)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64LoweredAtomicAdd64(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (LoweredAtomicAdd64 ptr (MOVVconst [c]) mem)
	// cond: is32Bit(c)
	// result: (LoweredAtomicAddconst64 [c] ptr mem)
	for {
		ptr := v_0
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		mem := v_2
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpMIPS64LoweredAtomicAddconst64)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64LoweredAtomicStore32(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (LoweredAtomicStore32 ptr (MOVVconst [0]) mem)
	// result: (LoweredAtomicStorezero32 ptr mem)
	for {
		ptr := v_0
		if v_1.Op != OpMIPS64MOVVconst || v_1.AuxInt != 0 {
			break
		}
		mem := v_2
		v.reset(OpMIPS64LoweredAtomicStorezero32)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64LoweredAtomicStore64(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (LoweredAtomicStore64 ptr (MOVVconst [0]) mem)
	// result: (LoweredAtomicStorezero64 ptr mem)
	for {
		ptr := v_0
		if v_1.Op != OpMIPS64MOVVconst || v_1.AuxInt != 0 {
			break
		}
		mem := v_2
		v.reset(OpMIPS64LoweredAtomicStorezero64)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVBUload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBUload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVBUload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVBUload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBUload [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVBUload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVBUload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVBUreg(v *Value) bool {
	v_0 := v.Args[0]
	// match: (MOVBUreg x:(MOVBUload _ _))
	// result: (MOVVreg x)
	for {
		x := v_0
		if x.Op != OpMIPS64MOVBUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBUreg x:(MOVBUreg _))
	// result: (MOVVreg x)
	for {
		x := v_0
		if x.Op != OpMIPS64MOVBUreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBUreg (MOVVconst [c]))
	// result: (MOVVconst [int64(uint8(c))])
	for {
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = int64(uint8(c))
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVBload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVBload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVBload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBload [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVBload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVBload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVBreg(v *Value) bool {
	v_0 := v.Args[0]
	// match: (MOVBreg x:(MOVBload _ _))
	// result: (MOVVreg x)
	for {
		x := v_0
		if x.Op != OpMIPS64MOVBload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg x:(MOVBreg _))
	// result: (MOVVreg x)
	for {
		x := v_0
		if x.Op != OpMIPS64MOVBreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg (MOVVconst [c]))
	// result: (MOVVconst [int64(int8(c))])
	for {
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = int64(int8(c))
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVBstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBstore [off1] {sym} (ADDVconst [off2] ptr) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVBstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVBstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVBstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVBstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVVconst [0]) mem)
	// result: (MOVBstorezero [off] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPS64MOVVconst || v_1.AuxInt != 0 {
			break
		}
		mem := v_2
		v.reset(OpMIPS64MOVBstorezero)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVBreg x) mem)
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPS64MOVBreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpMIPS64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVBUreg x) mem)
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPS64MOVBUreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpMIPS64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVHreg x) mem)
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPS64MOVHreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpMIPS64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVHUreg x) mem)
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPS64MOVHUreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpMIPS64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVWreg x) mem)
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPS64MOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpMIPS64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVWUreg x) mem)
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPS64MOVWUreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpMIPS64MOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVBstorezero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBstorezero [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVBstorezero [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVBstorezero)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstorezero [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVBstorezero [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVBstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVDload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVDload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVDload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVDload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDload [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVDload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVDload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVDstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVDstore [off1] {sym} (ADDVconst [off2] ptr) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVDstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDstore [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVDstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVFload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVFload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVFload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVFload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVFload [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVFload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVFload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVFstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVFstore [off1] {sym} (ADDVconst [off2] ptr) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVFstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVFstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVFstore [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVFstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVFstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVHUload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHUload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVHUload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVHUload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHUload [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVHUload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVHUload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVHUreg(v *Value) bool {
	v_0 := v.Args[0]
	// match: (MOVHUreg x:(MOVBUload _ _))
	// result: (MOVVreg x)
	for {
		x := v_0
		if x.Op != OpMIPS64MOVBUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg x:(MOVHUload _ _))
	// result: (MOVVreg x)
	for {
		x := v_0
		if x.Op != OpMIPS64MOVHUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg x:(MOVBUreg _))
	// result: (MOVVreg x)
	for {
		x := v_0
		if x.Op != OpMIPS64MOVBUreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg x:(MOVHUreg _))
	// result: (MOVVreg x)
	for {
		x := v_0
		if x.Op != OpMIPS64MOVHUreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg (MOVVconst [c]))
	// result: (MOVVconst [int64(uint16(c))])
	for {
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = int64(uint16(c))
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVHload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVHload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVHload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHload [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVHload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVHload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVHreg(v *Value) bool {
	v_0 := v.Args[0]
	// match: (MOVHreg x:(MOVBload _ _))
	// result: (MOVVreg x)
	for {
		x := v_0
		if x.Op != OpMIPS64MOVBload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVBUload _ _))
	// result: (MOVVreg x)
	for {
		x := v_0
		if x.Op != OpMIPS64MOVBUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVHload _ _))
	// result: (MOVVreg x)
	for {
		x := v_0
		if x.Op != OpMIPS64MOVHload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVBreg _))
	// result: (MOVVreg x)
	for {
		x := v_0
		if x.Op != OpMIPS64MOVBreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVBUreg _))
	// result: (MOVVreg x)
	for {
		x := v_0
		if x.Op != OpMIPS64MOVBUreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVHreg _))
	// result: (MOVVreg x)
	for {
		x := v_0
		if x.Op != OpMIPS64MOVHreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg (MOVVconst [c]))
	// result: (MOVVconst [int64(int16(c))])
	for {
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = int64(int16(c))
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVHstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHstore [off1] {sym} (ADDVconst [off2] ptr) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVHstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVHstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVHstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVHstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVVconst [0]) mem)
	// result: (MOVHstorezero [off] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPS64MOVVconst || v_1.AuxInt != 0 {
			break
		}
		mem := v_2
		v.reset(OpMIPS64MOVHstorezero)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVHreg x) mem)
	// result: (MOVHstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPS64MOVHreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpMIPS64MOVHstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVHUreg x) mem)
	// result: (MOVHstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPS64MOVHUreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpMIPS64MOVHstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVWreg x) mem)
	// result: (MOVHstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPS64MOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpMIPS64MOVHstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVWUreg x) mem)
	// result: (MOVHstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPS64MOVWUreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpMIPS64MOVHstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVHstorezero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHstorezero [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVHstorezero [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVHstorezero)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstorezero [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVHstorezero [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVHstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVVload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVVload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVVload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVVload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVVload [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVVload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVVload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVVreg(v *Value) bool {
	v_0 := v.Args[0]
	// match: (MOVVreg x)
	// cond: x.Uses == 1
	// result: (MOVVnop x)
	for {
		x := v_0
		if !(x.Uses == 1) {
			break
		}
		v.reset(OpMIPS64MOVVnop)
		v.AddArg(x)
		return true
	}
	// match: (MOVVreg (MOVVconst [c]))
	// result: (MOVVconst [c])
	for {
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = c
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVVstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVVstore [off1] {sym} (ADDVconst [off2] ptr) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVVstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVVstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVVstore [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVVstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVVstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVVstore [off] {sym} ptr (MOVVconst [0]) mem)
	// result: (MOVVstorezero [off] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPS64MOVVconst || v_1.AuxInt != 0 {
			break
		}
		mem := v_2
		v.reset(OpMIPS64MOVVstorezero)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVVstorezero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVVstorezero [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVVstorezero [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVVstorezero)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVVstorezero [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVVstorezero [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVVstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVWUload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWUload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVWUload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVWUload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWUload [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVWUload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVWUload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVWUreg(v *Value) bool {
	v_0 := v.Args[0]
	// match: (MOVWUreg x:(MOVBUload _ _))
	// result: (MOVVreg x)
	for {
		x := v_0
		if x.Op != OpMIPS64MOVBUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWUreg x:(MOVHUload _ _))
	// result: (MOVVreg x)
	for {
		x := v_0
		if x.Op != OpMIPS64MOVHUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWUreg x:(MOVWUload _ _))
	// result: (MOVVreg x)
	for {
		x := v_0
		if x.Op != OpMIPS64MOVWUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWUreg x:(MOVBUreg _))
	// result: (MOVVreg x)
	for {
		x := v_0
		if x.Op != OpMIPS64MOVBUreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWUreg x:(MOVHUreg _))
	// result: (MOVVreg x)
	for {
		x := v_0
		if x.Op != OpMIPS64MOVHUreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWUreg x:(MOVWUreg _))
	// result: (MOVVreg x)
	for {
		x := v_0
		if x.Op != OpMIPS64MOVWUreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWUreg (MOVVconst [c]))
	// result: (MOVVconst [int64(uint32(c))])
	for {
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = int64(uint32(c))
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVWload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWload [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVWload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVWload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWload [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVWload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVWload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVWreg(v *Value) bool {
	v_0 := v.Args[0]
	// match: (MOVWreg x:(MOVBload _ _))
	// result: (MOVVreg x)
	for {
		x := v_0
		if x.Op != OpMIPS64MOVBload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg x:(MOVBUload _ _))
	// result: (MOVVreg x)
	for {
		x := v_0
		if x.Op != OpMIPS64MOVBUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg x:(MOVHload _ _))
	// result: (MOVVreg x)
	for {
		x := v_0
		if x.Op != OpMIPS64MOVHload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg x:(MOVHUload _ _))
	// result: (MOVVreg x)
	for {
		x := v_0
		if x.Op != OpMIPS64MOVHUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg x:(MOVWload _ _))
	// result: (MOVVreg x)
	for {
		x := v_0
		if x.Op != OpMIPS64MOVWload {
			break
		}
		_ = x.Args[1]
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg x:(MOVBreg _))
	// result: (MOVVreg x)
	for {
		x := v_0
		if x.Op != OpMIPS64MOVBreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg x:(MOVBUreg _))
	// result: (MOVVreg x)
	for {
		x := v_0
		if x.Op != OpMIPS64MOVBUreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg x:(MOVHreg _))
	// result: (MOVVreg x)
	for {
		x := v_0
		if x.Op != OpMIPS64MOVHreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg x:(MOVWreg _))
	// result: (MOVVreg x)
	for {
		x := v_0
		if x.Op != OpMIPS64MOVWreg {
			break
		}
		v.reset(OpMIPS64MOVVreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg (MOVVconst [c]))
	// result: (MOVVconst [int64(int32(c))])
	for {
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = int64(int32(c))
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVWstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWstore [off1] {sym} (ADDVconst [off2] ptr) val mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVWstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVWstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVWstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVWstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off] {sym} ptr (MOVVconst [0]) mem)
	// result: (MOVWstorezero [off] {sym} ptr mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPS64MOVVconst || v_1.AuxInt != 0 {
			break
		}
		mem := v_2
		v.reset(OpMIPS64MOVWstorezero)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off] {sym} ptr (MOVWreg x) mem)
	// result: (MOVWstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPS64MOVWreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpMIPS64MOVWstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off] {sym} ptr (MOVWUreg x) mem)
	// result: (MOVWstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpMIPS64MOVWUreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpMIPS64MOVWstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64MOVWstorezero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWstorezero [off1] {sym} (ADDVconst [off2] ptr) mem)
	// cond: is32Bit(off1+off2)
	// result: (MOVWstorezero [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		if !(is32Bit(off1 + off2)) {
			break
		}
		v.reset(OpMIPS64MOVWstorezero)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstorezero [off1] {sym1} (MOVVaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2) && is32Bit(off1+off2)
	// result: (MOVWstorezero [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpMIPS64MOVVaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2) && is32Bit(off1+off2)) {
			break
		}
		v.reset(OpMIPS64MOVWstorezero)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64NEGV(v *Value) bool {
	v_0 := v.Args[0]
	// match: (NEGV (MOVVconst [c]))
	// result: (MOVVconst [-c])
	for {
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = -c
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64NOR(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (NOR x (MOVVconst [c]))
	// cond: is32Bit(c)
	// result: (NORconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpMIPS64MOVVconst {
				continue
			}
			c := v_1.AuxInt
			if !(is32Bit(c)) {
				continue
			}
			v.reset(OpMIPS64NORconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64NORconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (NORconst [c] (MOVVconst [d]))
	// result: (MOVVconst [^(c|d)])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = ^(c | d)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64OR(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (OR x (MOVVconst [c]))
	// cond: is32Bit(c)
	// result: (ORconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpMIPS64MOVVconst {
				continue
			}
			c := v_1.AuxInt
			if !(is32Bit(c)) {
				continue
			}
			v.reset(OpMIPS64ORconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (OR x x)
	// result: x
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64ORconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ORconst [0] x)
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v_0
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ORconst [-1] _)
	// result: (MOVVconst [-1])
	for {
		if v.AuxInt != -1 {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = -1
		return true
	}
	// match: (ORconst [c] (MOVVconst [d]))
	// result: (MOVVconst [c|d])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = c | d
		return true
	}
	// match: (ORconst [c] (ORconst [d] x))
	// cond: is32Bit(c|d)
	// result: (ORconst [c|d] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64ORconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(c | d)) {
			break
		}
		v.reset(OpMIPS64ORconst)
		v.AuxInt = c | d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SGT(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SGT (MOVVconst [c]) x)
	// cond: is32Bit(c)
	// result: (SGTconst [c] x)
	for {
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpMIPS64SGTconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SGTU(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SGTU (MOVVconst [c]) x)
	// cond: is32Bit(c)
	// result: (SGTUconst [c] x)
	for {
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpMIPS64SGTUconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SGTUconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SGTUconst [c] (MOVVconst [d]))
	// cond: uint64(c)>uint64(d)
	// result: (MOVVconst [1])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		if !(uint64(c) > uint64(d)) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTUconst [c] (MOVVconst [d]))
	// cond: uint64(c)<=uint64(d)
	// result: (MOVVconst [0])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		if !(uint64(c) <= uint64(d)) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}
	// match: (SGTUconst [c] (MOVBUreg _))
	// cond: 0xff < uint64(c)
	// result: (MOVVconst [1])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64MOVBUreg || !(0xff < uint64(c)) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTUconst [c] (MOVHUreg _))
	// cond: 0xffff < uint64(c)
	// result: (MOVVconst [1])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64MOVHUreg || !(0xffff < uint64(c)) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTUconst [c] (ANDconst [m] _))
	// cond: uint64(m) < uint64(c)
	// result: (MOVVconst [1])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64ANDconst {
			break
		}
		m := v_0.AuxInt
		if !(uint64(m) < uint64(c)) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTUconst [c] (SRLVconst _ [d]))
	// cond: 0 < d && d <= 63 && 0xffffffffffffffff>>uint64(d) < uint64(c)
	// result: (MOVVconst [1])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64SRLVconst {
			break
		}
		d := v_0.AuxInt
		if !(0 < d && d <= 63 && 0xffffffffffffffff>>uint64(d) < uint64(c)) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SGTconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SGTconst [c] (MOVVconst [d]))
	// cond: c>d
	// result: (MOVVconst [1])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		if !(c > d) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTconst [c] (MOVVconst [d]))
	// cond: c<=d
	// result: (MOVVconst [0])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		if !(c <= d) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}
	// match: (SGTconst [c] (MOVBreg _))
	// cond: 0x7f < c
	// result: (MOVVconst [1])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64MOVBreg || !(0x7f < c) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTconst [c] (MOVBreg _))
	// cond: c <= -0x80
	// result: (MOVVconst [0])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64MOVBreg || !(c <= -0x80) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}
	// match: (SGTconst [c] (MOVBUreg _))
	// cond: 0xff < c
	// result: (MOVVconst [1])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64MOVBUreg || !(0xff < c) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTconst [c] (MOVBUreg _))
	// cond: c < 0
	// result: (MOVVconst [0])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64MOVBUreg || !(c < 0) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}
	// match: (SGTconst [c] (MOVHreg _))
	// cond: 0x7fff < c
	// result: (MOVVconst [1])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64MOVHreg || !(0x7fff < c) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTconst [c] (MOVHreg _))
	// cond: c <= -0x8000
	// result: (MOVVconst [0])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64MOVHreg || !(c <= -0x8000) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}
	// match: (SGTconst [c] (MOVHUreg _))
	// cond: 0xffff < c
	// result: (MOVVconst [1])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64MOVHUreg || !(0xffff < c) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTconst [c] (MOVHUreg _))
	// cond: c < 0
	// result: (MOVVconst [0])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64MOVHUreg || !(c < 0) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}
	// match: (SGTconst [c] (MOVWUreg _))
	// cond: c < 0
	// result: (MOVVconst [0])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64MOVWUreg || !(c < 0) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}
	// match: (SGTconst [c] (ANDconst [m] _))
	// cond: 0 <= m && m < c
	// result: (MOVVconst [1])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64ANDconst {
			break
		}
		m := v_0.AuxInt
		if !(0 <= m && m < c) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}
	// match: (SGTconst [c] (SRLVconst _ [d]))
	// cond: 0 <= c && 0 < d && d <= 63 && 0xffffffffffffffff>>uint64(d) < uint64(c)
	// result: (MOVVconst [1])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64SRLVconst {
			break
		}
		d := v_0.AuxInt
		if !(0 <= c && 0 < d && d <= 63 && 0xffffffffffffffff>>uint64(d) < uint64(c)) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 1
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SLLV(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SLLV _ (MOVVconst [c]))
	// cond: uint64(c)>=64
	// result: (MOVVconst [0])
	for {
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}
	// match: (SLLV x (MOVVconst [c]))
	// result: (SLLVconst x [c])
	for {
		x := v_0
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMIPS64SLLVconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SLLVconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SLLVconst [c] (MOVVconst [d]))
	// result: (MOVVconst [d<<uint64(c)])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = d << uint64(c)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SRAV(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SRAV x (MOVVconst [c]))
	// cond: uint64(c)>=64
	// result: (SRAVconst x [63])
	for {
		x := v_0
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpMIPS64SRAVconst)
		v.AuxInt = 63
		v.AddArg(x)
		return true
	}
	// match: (SRAV x (MOVVconst [c]))
	// result: (SRAVconst x [c])
	for {
		x := v_0
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMIPS64SRAVconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SRAVconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SRAVconst [c] (MOVVconst [d]))
	// result: (MOVVconst [d>>uint64(c)])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = d >> uint64(c)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SRLV(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SRLV _ (MOVVconst [c]))
	// cond: uint64(c)>=64
	// result: (MOVVconst [0])
	for {
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}
	// match: (SRLV x (MOVVconst [c]))
	// result: (SRLVconst x [c])
	for {
		x := v_0
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpMIPS64SRLVconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SRLVconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SRLVconst [c] (MOVVconst [d]))
	// result: (MOVVconst [int64(uint64(d)>>uint64(c))])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = int64(uint64(d) >> uint64(c))
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SUBV(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SUBV x (MOVVconst [c]))
	// cond: is32Bit(c)
	// result: (SUBVconst [c] x)
	for {
		x := v_0
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		if !(is32Bit(c)) {
			break
		}
		v.reset(OpMIPS64SUBVconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (SUBV x x)
	// result: (MOVVconst [0])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}
	// match: (SUBV (MOVVconst [0]) x)
	// result: (NEGV x)
	for {
		if v_0.Op != OpMIPS64MOVVconst || v_0.AuxInt != 0 {
			break
		}
		x := v_1
		v.reset(OpMIPS64NEGV)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64SUBVconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SUBVconst [0] x)
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v_0
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (SUBVconst [c] (MOVVconst [d]))
	// result: (MOVVconst [d-c])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = d - c
		return true
	}
	// match: (SUBVconst [c] (SUBVconst [d] x))
	// cond: is32Bit(-c-d)
	// result: (ADDVconst [-c-d] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64SUBVconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(-c - d)) {
			break
		}
		v.reset(OpMIPS64ADDVconst)
		v.AuxInt = -c - d
		v.AddArg(x)
		return true
	}
	// match: (SUBVconst [c] (ADDVconst [d] x))
	// cond: is32Bit(-c+d)
	// result: (ADDVconst [-c+d] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64ADDVconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(-c + d)) {
			break
		}
		v.reset(OpMIPS64ADDVconst)
		v.AuxInt = -c + d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64XOR(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (XOR x (MOVVconst [c]))
	// cond: is32Bit(c)
	// result: (XORconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpMIPS64MOVVconst {
				continue
			}
			c := v_1.AuxInt
			if !(is32Bit(c)) {
				continue
			}
			v.reset(OpMIPS64XORconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (XOR x x)
	// result: (MOVVconst [0])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMIPS64XORconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (XORconst [0] x)
	// result: x
	for {
		if v.AuxInt != 0 {
			break
		}
		x := v_0
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (XORconst [-1] x)
	// result: (NORconst [0] x)
	for {
		if v.AuxInt != -1 {
			break
		}
		x := v_0
		v.reset(OpMIPS64NORconst)
		v.AuxInt = 0
		v.AddArg(x)
		return true
	}
	// match: (XORconst [c] (MOVVconst [d]))
	// result: (MOVVconst [c^d])
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = c ^ d
		return true
	}
	// match: (XORconst [c] (XORconst [d] x))
	// cond: is32Bit(c^d)
	// result: (XORconst [c^d] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpMIPS64XORconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		if !(is32Bit(c ^ d)) {
			break
		}
		v.reset(OpMIPS64XORconst)
		v.AuxInt = c ^ d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMod16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod16 x y)
	// result: (Select0 (DIVV (SignExt16to64 x) (SignExt16to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVV, types.NewTuple(typ.Int64, typ.Int64))
		v1 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpMod16u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod16u x y)
	// result: (Select0 (DIVVU (ZeroExt16to64 x) (ZeroExt16to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVVU, types.NewTuple(typ.UInt64, typ.UInt64))
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpMod32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod32 x y)
	// result: (Select0 (DIVV (SignExt32to64 x) (SignExt32to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVV, types.NewTuple(typ.Int64, typ.Int64))
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpMod32u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod32u x y)
	// result: (Select0 (DIVVU (ZeroExt32to64 x) (ZeroExt32to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVVU, types.NewTuple(typ.UInt64, typ.UInt64))
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpMod64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod64 x y)
	// result: (Select0 (DIVV x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVV, types.NewTuple(typ.Int64, typ.Int64))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpMod64u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod64u x y)
	// result: (Select0 (DIVVU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVVU, types.NewTuple(typ.UInt64, typ.UInt64))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpMod8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod8 x y)
	// result: (Select0 (DIVV (SignExt8to64 x) (SignExt8to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVV, types.NewTuple(typ.Int64, typ.Int64))
		v1 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpMod8u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod8u x y)
	// result: (Select0 (DIVVU (ZeroExt8to64 x) (ZeroExt8to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect0)
		v0 := b.NewValue0(v.Pos, OpMIPS64DIVVU, types.NewTuple(typ.UInt64, typ.UInt64))
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpMove(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: (Move [0] _ _ mem)
	// result: mem
	for {
		if v.AuxInt != 0 {
			break
		}
		mem := v_2
		v.reset(OpCopy)
		v.Type = mem.Type
		v.AddArg(mem)
		return true
	}
	// match: (Move [1] dst src mem)
	// result: (MOVBstore dst (MOVBload src mem) mem)
	for {
		if v.AuxInt != 1 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpMIPS64MOVBstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVBload, typ.Int8)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [2] {t} dst src mem)
	// cond: t.(*types.Type).Alignment()%2 == 0
	// result: (MOVHstore dst (MOVHload src mem) mem)
	for {
		if v.AuxInt != 2 {
			break
		}
		t := v.Aux
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.(*types.Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpMIPS64MOVHstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVHload, typ.Int16)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [2] dst src mem)
	// result: (MOVBstore [1] dst (MOVBload [1] src mem) (MOVBstore dst (MOVBload src mem) mem))
	for {
		if v.AuxInt != 2 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpMIPS64MOVBstore)
		v.AuxInt = 1
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVBload, typ.Int8)
		v0.AuxInt = 1
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVBload, typ.Int8)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Move [4] {t} dst src mem)
	// cond: t.(*types.Type).Alignment()%4 == 0
	// result: (MOVWstore dst (MOVWload src mem) mem)
	for {
		if v.AuxInt != 4 {
			break
		}
		t := v.Aux
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.(*types.Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpMIPS64MOVWstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVWload, typ.Int32)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [4] {t} dst src mem)
	// cond: t.(*types.Type).Alignment()%2 == 0
	// result: (MOVHstore [2] dst (MOVHload [2] src mem) (MOVHstore dst (MOVHload src mem) mem))
	for {
		if v.AuxInt != 4 {
			break
		}
		t := v.Aux
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.(*types.Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpMIPS64MOVHstore)
		v.AuxInt = 2
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVHload, typ.Int16)
		v0.AuxInt = 2
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVHload, typ.Int16)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Move [4] dst src mem)
	// result: (MOVBstore [3] dst (MOVBload [3] src mem) (MOVBstore [2] dst (MOVBload [2] src mem) (MOVBstore [1] dst (MOVBload [1] src mem) (MOVBstore dst (MOVBload src mem) mem))))
	for {
		if v.AuxInt != 4 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpMIPS64MOVBstore)
		v.AuxInt = 3
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVBload, typ.Int8)
		v0.AuxInt = 3
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, types.TypeMem)
		v1.AuxInt = 2
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVBload, typ.Int8)
		v2.AuxInt = 2
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, types.TypeMem)
		v3.AuxInt = 1
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVBload, typ.Int8)
		v4.AuxInt = 1
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, types.TypeMem)
		v5.AddArg(dst)
		v6 := b.NewValue0(v.Pos, OpMIPS64MOVBload, typ.Int8)
		v6.AddArg(src)
		v6.AddArg(mem)
		v5.AddArg(v6)
		v5.AddArg(mem)
		v3.AddArg(v5)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Move [8] {t} dst src mem)
	// cond: t.(*types.Type).Alignment()%8 == 0
	// result: (MOVVstore dst (MOVVload src mem) mem)
	for {
		if v.AuxInt != 8 {
			break
		}
		t := v.Aux
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.(*types.Type).Alignment()%8 == 0) {
			break
		}
		v.reset(OpMIPS64MOVVstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVload, typ.UInt64)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [8] {t} dst src mem)
	// cond: t.(*types.Type).Alignment()%4 == 0
	// result: (MOVWstore [4] dst (MOVWload [4] src mem) (MOVWstore dst (MOVWload src mem) mem))
	for {
		if v.AuxInt != 8 {
			break
		}
		t := v.Aux
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.(*types.Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpMIPS64MOVWstore)
		v.AuxInt = 4
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVWload, typ.Int32)
		v0.AuxInt = 4
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVWstore, types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVWload, typ.Int32)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Move [8] {t} dst src mem)
	// cond: t.(*types.Type).Alignment()%2 == 0
	// result: (MOVHstore [6] dst (MOVHload [6] src mem) (MOVHstore [4] dst (MOVHload [4] src mem) (MOVHstore [2] dst (MOVHload [2] src mem) (MOVHstore dst (MOVHload src mem) mem))))
	for {
		if v.AuxInt != 8 {
			break
		}
		t := v.Aux
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.(*types.Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpMIPS64MOVHstore)
		v.AuxInt = 6
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVHload, typ.Int16)
		v0.AuxInt = 6
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, types.TypeMem)
		v1.AuxInt = 4
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVHload, typ.Int16)
		v2.AuxInt = 4
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, types.TypeMem)
		v3.AuxInt = 2
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVHload, typ.Int16)
		v4.AuxInt = 2
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, types.TypeMem)
		v5.AddArg(dst)
		v6 := b.NewValue0(v.Pos, OpMIPS64MOVHload, typ.Int16)
		v6.AddArg(src)
		v6.AddArg(mem)
		v5.AddArg(v6)
		v5.AddArg(mem)
		v3.AddArg(v5)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Move [3] dst src mem)
	// result: (MOVBstore [2] dst (MOVBload [2] src mem) (MOVBstore [1] dst (MOVBload [1] src mem) (MOVBstore dst (MOVBload src mem) mem)))
	for {
		if v.AuxInt != 3 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpMIPS64MOVBstore)
		v.AuxInt = 2
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVBload, typ.Int8)
		v0.AuxInt = 2
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, types.TypeMem)
		v1.AuxInt = 1
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVBload, typ.Int8)
		v2.AuxInt = 1
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, types.TypeMem)
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVBload, typ.Int8)
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Move [6] {t} dst src mem)
	// cond: t.(*types.Type).Alignment()%2 == 0
	// result: (MOVHstore [4] dst (MOVHload [4] src mem) (MOVHstore [2] dst (MOVHload [2] src mem) (MOVHstore dst (MOVHload src mem) mem)))
	for {
		if v.AuxInt != 6 {
			break
		}
		t := v.Aux
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.(*types.Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpMIPS64MOVHstore)
		v.AuxInt = 4
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVHload, typ.Int16)
		v0.AuxInt = 4
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, types.TypeMem)
		v1.AuxInt = 2
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVHload, typ.Int16)
		v2.AuxInt = 2
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, types.TypeMem)
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVHload, typ.Int16)
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Move [12] {t} dst src mem)
	// cond: t.(*types.Type).Alignment()%4 == 0
	// result: (MOVWstore [8] dst (MOVWload [8] src mem) (MOVWstore [4] dst (MOVWload [4] src mem) (MOVWstore dst (MOVWload src mem) mem)))
	for {
		if v.AuxInt != 12 {
			break
		}
		t := v.Aux
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.(*types.Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpMIPS64MOVWstore)
		v.AuxInt = 8
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVWload, typ.Int32)
		v0.AuxInt = 8
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVWstore, types.TypeMem)
		v1.AuxInt = 4
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVWload, typ.Int32)
		v2.AuxInt = 4
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVWstore, types.TypeMem)
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVWload, typ.Int32)
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Move [16] {t} dst src mem)
	// cond: t.(*types.Type).Alignment()%8 == 0
	// result: (MOVVstore [8] dst (MOVVload [8] src mem) (MOVVstore dst (MOVVload src mem) mem))
	for {
		if v.AuxInt != 16 {
			break
		}
		t := v.Aux
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.(*types.Type).Alignment()%8 == 0) {
			break
		}
		v.reset(OpMIPS64MOVVstore)
		v.AuxInt = 8
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVload, typ.UInt64)
		v0.AuxInt = 8
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVVstore, types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVload, typ.UInt64)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Move [24] {t} dst src mem)
	// cond: t.(*types.Type).Alignment()%8 == 0
	// result: (MOVVstore [16] dst (MOVVload [16] src mem) (MOVVstore [8] dst (MOVVload [8] src mem) (MOVVstore dst (MOVVload src mem) mem)))
	for {
		if v.AuxInt != 24 {
			break
		}
		t := v.Aux
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.(*types.Type).Alignment()%8 == 0) {
			break
		}
		v.reset(OpMIPS64MOVVstore)
		v.AuxInt = 16
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVload, typ.UInt64)
		v0.AuxInt = 16
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVVstore, types.TypeMem)
		v1.AuxInt = 8
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVload, typ.UInt64)
		v2.AuxInt = 8
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVVstore, types.TypeMem)
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVVload, typ.UInt64)
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Move [s] {t} dst src mem)
	// cond: s%8 == 0 && s >= 24 && s <= 8*128 && t.(*types.Type).Alignment()%8 == 0 && !config.noDuffDevice
	// result: (DUFFCOPY [16 * (128 - s/8)] dst src mem)
	for {
		s := v.AuxInt
		t := v.Aux
		dst := v_0
		src := v_1
		mem := v_2
		if !(s%8 == 0 && s >= 24 && s <= 8*128 && t.(*types.Type).Alignment()%8 == 0 && !config.noDuffDevice) {
			break
		}
		v.reset(OpMIPS64DUFFCOPY)
		v.AuxInt = 16 * (128 - s/8)
		v.AddArg(dst)
		v.AddArg(src)
		v.AddArg(mem)
		return true
	}
	// match: (Move [s] {t} dst src mem)
	// cond: s > 24 || t.(*types.Type).Alignment()%8 != 0
	// result: (LoweredMove [t.(*types.Type).Alignment()] dst src (ADDVconst <src.Type> src [s-moveSize(t.(*types.Type).Alignment(), config)]) mem)
	for {
		s := v.AuxInt
		t := v.Aux
		dst := v_0
		src := v_1
		mem := v_2
		if !(s > 24 || t.(*types.Type).Alignment()%8 != 0) {
			break
		}
		v.reset(OpMIPS64LoweredMove)
		v.AuxInt = t.(*types.Type).Alignment()
		v.AddArg(dst)
		v.AddArg(src)
		v0 := b.NewValue0(v.Pos, OpMIPS64ADDVconst, src.Type)
		v0.AuxInt = s - moveSize(t.(*types.Type).Alignment(), config)
		v0.AddArg(src)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpMul16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mul16 x y)
	// result: (Select1 (MULVU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64MULVU, types.NewTuple(typ.UInt64, typ.UInt64))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpMul32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mul32 x y)
	// result: (Select1 (MULVU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64MULVU, types.NewTuple(typ.UInt64, typ.UInt64))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpMul64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mul64 x y)
	// result: (Select1 (MULVU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64MULVU, types.NewTuple(typ.UInt64, typ.UInt64))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpMul8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mul8 x y)
	// result: (Select1 (MULVU x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect1)
		v0 := b.NewValue0(v.Pos, OpMIPS64MULVU, types.NewTuple(typ.UInt64, typ.UInt64))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpNeq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq16 x y)
	// result: (SGTU (XOR (ZeroExt16to32 x) (ZeroExt16to64 y)) (MOVVconst [0]))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpMIPS64XOR, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v3.AuxInt = 0
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS64_OpNeq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq32 x y)
	// result: (SGTU (XOR (ZeroExt32to64 x) (ZeroExt32to64 y)) (MOVVconst [0]))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpMIPS64XOR, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v3.AuxInt = 0
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS64_OpNeq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Neq32F x y)
	// result: (FPFlagFalse (CMPEQF x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64FPFlagFalse)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPEQF, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpNeq64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq64 x y)
	// result: (SGTU (XOR x y) (MOVVconst [0]))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpMIPS64XOR, typ.UInt64)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpNeq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Neq64F x y)
	// result: (FPFlagFalse (CMPEQD x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64FPFlagFalse)
		v0 := b.NewValue0(v.Pos, OpMIPS64CMPEQD, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpNeq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq8 x y)
	// result: (SGTU (XOR (ZeroExt8to64 x) (ZeroExt8to64 y)) (MOVVconst [0]))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpMIPS64XOR, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v3.AuxInt = 0
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS64_OpNeqPtr(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (NeqPtr x y)
	// result: (SGTU (XOR x y) (MOVVconst [0]))
	for {
		x := v_0
		y := v_1
		v.reset(OpMIPS64SGTU)
		v0 := b.NewValue0(v.Pos, OpMIPS64XOR, typ.UInt64)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v1.AuxInt = 0
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpNot(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Not x)
	// result: (XORconst [1] x)
	for {
		x := v_0
		v.reset(OpMIPS64XORconst)
		v.AuxInt = 1
		v.AddArg(x)
		return true
	}
}
func rewriteValueMIPS64_OpOffPtr(v *Value) bool {
	v_0 := v.Args[0]
	// match: (OffPtr [off] ptr:(SP))
	// result: (MOVVaddr [off] ptr)
	for {
		off := v.AuxInt
		ptr := v_0
		if ptr.Op != OpSP {
			break
		}
		v.reset(OpMIPS64MOVVaddr)
		v.AuxInt = off
		v.AddArg(ptr)
		return true
	}
	// match: (OffPtr [off] ptr)
	// result: (ADDVconst [off] ptr)
	for {
		off := v.AuxInt
		ptr := v_0
		v.reset(OpMIPS64ADDVconst)
		v.AuxInt = off
		v.AddArg(ptr)
		return true
	}
}
func rewriteValueMIPS64_OpPanicBounds(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (PanicBounds [kind] x y mem)
	// cond: boundsABI(kind) == 0
	// result: (LoweredPanicBoundsA [kind] x y mem)
	for {
		kind := v.AuxInt
		x := v_0
		y := v_1
		mem := v_2
		if !(boundsABI(kind) == 0) {
			break
		}
		v.reset(OpMIPS64LoweredPanicBoundsA)
		v.AuxInt = kind
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(mem)
		return true
	}
	// match: (PanicBounds [kind] x y mem)
	// cond: boundsABI(kind) == 1
	// result: (LoweredPanicBoundsB [kind] x y mem)
	for {
		kind := v.AuxInt
		x := v_0
		y := v_1
		mem := v_2
		if !(boundsABI(kind) == 1) {
			break
		}
		v.reset(OpMIPS64LoweredPanicBoundsB)
		v.AuxInt = kind
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(mem)
		return true
	}
	// match: (PanicBounds [kind] x y mem)
	// cond: boundsABI(kind) == 2
	// result: (LoweredPanicBoundsC [kind] x y mem)
	for {
		kind := v.AuxInt
		x := v_0
		y := v_1
		mem := v_2
		if !(boundsABI(kind) == 2) {
			break
		}
		v.reset(OpMIPS64LoweredPanicBoundsC)
		v.AuxInt = kind
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpRotateLeft16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft16 <t> x (MOVVconst [c]))
	// result: (Or16 (Lsh16x64 <t> x (MOVVconst [c&15])) (Rsh16Ux64 <t> x (MOVVconst [-c&15])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpOr16)
		v0 := b.NewValue0(v.Pos, OpLsh16x64, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v1.AuxInt = c & 15
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpRsh16Ux64, t)
		v2.AddArg(x)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v3.AuxInt = -c & 15
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpRotateLeft32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft32 <t> x (MOVVconst [c]))
	// result: (Or32 (Lsh32x64 <t> x (MOVVconst [c&31])) (Rsh32Ux64 <t> x (MOVVconst [-c&31])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpOr32)
		v0 := b.NewValue0(v.Pos, OpLsh32x64, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v1.AuxInt = c & 31
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpRsh32Ux64, t)
		v2.AddArg(x)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v3.AuxInt = -c & 31
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpRotateLeft64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft64 <t> x (MOVVconst [c]))
	// result: (Or64 (Lsh64x64 <t> x (MOVVconst [c&63])) (Rsh64Ux64 <t> x (MOVVconst [-c&63])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpOr64)
		v0 := b.NewValue0(v.Pos, OpLsh64x64, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v1.AuxInt = c & 63
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpRsh64Ux64, t)
		v2.AddArg(x)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v3.AuxInt = -c & 63
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpRotateLeft8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft8 <t> x (MOVVconst [c]))
	// result: (Or8 (Lsh8x64 <t> x (MOVVconst [c&7])) (Rsh8Ux64 <t> x (MOVVconst [-c&7])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpOr8)
		v0 := b.NewValue0(v.Pos, OpLsh8x64, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v1.AuxInt = c & 7
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpRsh8Ux64, t)
		v2.AddArg(x)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v3.AuxInt = -c & 7
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpRsh16Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux16 <t> x y)
	// result: (AND (NEGV <t> (SGTU (MOVVconst <typ.UInt64> [64]) (ZeroExt16to64 y))) (SRLV <t> (ZeroExt16to64 x) (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v6.AddArg(y)
		v4.AddArg(v6)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh16Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux32 <t> x y)
	// result: (AND (NEGV <t> (SGTU (MOVVconst <typ.UInt64> [64]) (ZeroExt32to64 y))) (SRLV <t> (ZeroExt16to64 x) (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v6.AddArg(y)
		v4.AddArg(v6)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh16Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux64 <t> x y)
	// result: (AND (NEGV <t> (SGTU (MOVVconst <typ.UInt64> [64]) y)) (SRLV <t> (ZeroExt16to64 x) y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(x)
		v3.AddArg(v4)
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS64_OpRsh16Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux8 <t> x y)
	// result: (AND (NEGV <t> (SGTU (MOVVconst <typ.UInt64> [64]) (ZeroExt8to64 y))) (SRLV <t> (ZeroExt16to64 x) (ZeroExt8to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v6.AddArg(y)
		v4.AddArg(v6)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh16x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x16 <t> x y)
	// result: (SRAV (SignExt16to64 x) (OR <t> (NEGV <t> (SGTU (ZeroExt16to64 y) (MOVVconst <typ.UInt64> [63]))) (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v5.AuxInt = 63
		v3.AddArg(v5)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v6 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v6.AddArg(y)
		v1.AddArg(v6)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpRsh16x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x32 <t> x y)
	// result: (SRAV (SignExt16to64 x) (OR <t> (NEGV <t> (SGTU (ZeroExt32to64 y) (MOVVconst <typ.UInt64> [63]))) (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v5.AuxInt = 63
		v3.AddArg(v5)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v6 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v6.AddArg(y)
		v1.AddArg(v6)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpRsh16x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x64 <t> x y)
	// result: (SRAV (SignExt16to64 x) (OR <t> (NEGV <t> (SGTU y (MOVVconst <typ.UInt64> [63]))) y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v3.AddArg(y)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v4.AuxInt = 63
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpRsh16x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x8 <t> x y)
	// result: (SRAV (SignExt16to64 x) (OR <t> (NEGV <t> (SGTU (ZeroExt8to64 y) (MOVVconst <typ.UInt64> [63]))) (ZeroExt8to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v5.AuxInt = 63
		v3.AddArg(v5)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v6 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v6.AddArg(y)
		v1.AddArg(v6)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpRsh32Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux16 <t> x y)
	// result: (AND (NEGV <t> (SGTU (MOVVconst <typ.UInt64> [64]) (ZeroExt16to64 y))) (SRLV <t> (ZeroExt32to64 x) (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v6.AddArg(y)
		v4.AddArg(v6)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh32Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux32 <t> x y)
	// result: (AND (NEGV <t> (SGTU (MOVVconst <typ.UInt64> [64]) (ZeroExt32to64 y))) (SRLV <t> (ZeroExt32to64 x) (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v6.AddArg(y)
		v4.AddArg(v6)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh32Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux64 <t> x y)
	// result: (AND (NEGV <t> (SGTU (MOVVconst <typ.UInt64> [64]) y)) (SRLV <t> (ZeroExt32to64 x) y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(x)
		v3.AddArg(v4)
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS64_OpRsh32Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux8 <t> x y)
	// result: (AND (NEGV <t> (SGTU (MOVVconst <typ.UInt64> [64]) (ZeroExt8to64 y))) (SRLV <t> (ZeroExt32to64 x) (ZeroExt8to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v6.AddArg(y)
		v4.AddArg(v6)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh32x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x16 <t> x y)
	// result: (SRAV (SignExt32to64 x) (OR <t> (NEGV <t> (SGTU (ZeroExt16to64 y) (MOVVconst <typ.UInt64> [63]))) (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v5.AuxInt = 63
		v3.AddArg(v5)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v6 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v6.AddArg(y)
		v1.AddArg(v6)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpRsh32x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x32 <t> x y)
	// result: (SRAV (SignExt32to64 x) (OR <t> (NEGV <t> (SGTU (ZeroExt32to64 y) (MOVVconst <typ.UInt64> [63]))) (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v5.AuxInt = 63
		v3.AddArg(v5)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v6 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v6.AddArg(y)
		v1.AddArg(v6)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpRsh32x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x64 <t> x y)
	// result: (SRAV (SignExt32to64 x) (OR <t> (NEGV <t> (SGTU y (MOVVconst <typ.UInt64> [63]))) y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v3.AddArg(y)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v4.AuxInt = 63
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpRsh32x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x8 <t> x y)
	// result: (SRAV (SignExt32to64 x) (OR <t> (NEGV <t> (SGTU (ZeroExt8to64 y) (MOVVconst <typ.UInt64> [63]))) (ZeroExt8to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v5.AuxInt = 63
		v3.AddArg(v5)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v6 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v6.AddArg(y)
		v1.AddArg(v6)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpRsh64Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux16 <t> x y)
	// result: (AND (NEGV <t> (SGTU (MOVVconst <typ.UInt64> [64]) (ZeroExt16to64 y))) (SRLV <t> x (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh64Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux32 <t> x y)
	// result: (AND (NEGV <t> (SGTU (MOVVconst <typ.UInt64> [64]) (ZeroExt32to64 y))) (SRLV <t> x (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh64Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux64 <t> x y)
	// result: (AND (NEGV <t> (SGTU (MOVVconst <typ.UInt64> [64]) y)) (SRLV <t> x y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v3.AddArg(x)
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS64_OpRsh64Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux8 <t> x y)
	// result: (AND (NEGV <t> (SGTU (MOVVconst <typ.UInt64> [64]) (ZeroExt8to64 y))) (SRLV <t> x (ZeroExt8to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v5.AddArg(y)
		v4.AddArg(v5)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh64x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64x16 <t> x y)
	// result: (SRAV x (OR <t> (NEGV <t> (SGTU (ZeroExt16to64 y) (MOVVconst <typ.UInt64> [63]))) (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64SRAV)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v4.AuxInt = 63
		v2.AddArg(v4)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v5 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v5.AddArg(y)
		v0.AddArg(v5)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpRsh64x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64x32 <t> x y)
	// result: (SRAV x (OR <t> (NEGV <t> (SGTU (ZeroExt32to64 y) (MOVVconst <typ.UInt64> [63]))) (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64SRAV)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v4.AuxInt = 63
		v2.AddArg(v4)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v5 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v5.AddArg(y)
		v0.AddArg(v5)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpRsh64x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64x64 <t> x y)
	// result: (SRAV x (OR <t> (NEGV <t> (SGTU y (MOVVconst <typ.UInt64> [63]))) y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64SRAV)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2.AddArg(y)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v3.AuxInt = 63
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpRsh64x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64x8 <t> x y)
	// result: (SRAV x (OR <t> (NEGV <t> (SGTU (ZeroExt8to64 y) (MOVVconst <typ.UInt64> [63]))) (ZeroExt8to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64SRAV)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v4.AuxInt = 63
		v2.AddArg(v4)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v5.AddArg(y)
		v0.AddArg(v5)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpRsh8Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux16 <t> x y)
	// result: (AND (NEGV <t> (SGTU (MOVVconst <typ.UInt64> [64]) (ZeroExt16to64 y))) (SRLV <t> (ZeroExt8to64 x) (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v6.AddArg(y)
		v4.AddArg(v6)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh8Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux32 <t> x y)
	// result: (AND (NEGV <t> (SGTU (MOVVconst <typ.UInt64> [64]) (ZeroExt32to64 y))) (SRLV <t> (ZeroExt8to64 x) (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v6.AddArg(y)
		v4.AddArg(v6)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh8Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux64 <t> x y)
	// result: (AND (NEGV <t> (SGTU (MOVVconst <typ.UInt64> [64]) y)) (SRLV <t> (ZeroExt8to64 x) y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(x)
		v3.AddArg(v4)
		v3.AddArg(y)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueMIPS64_OpRsh8Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux8 <t> x y)
	// result: (AND (NEGV <t> (SGTU (MOVVconst <typ.UInt64> [64]) (ZeroExt8to64 y))) (SRLV <t> (ZeroExt8to64 x) (ZeroExt8to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64AND)
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v1 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 64
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v3.AddArg(y)
		v1.AddArg(v3)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpMIPS64SRLV, t)
		v5 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v5.AddArg(x)
		v4.AddArg(v5)
		v6 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v6.AddArg(y)
		v4.AddArg(v6)
		v.AddArg(v4)
		return true
	}
}
func rewriteValueMIPS64_OpRsh8x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x16 <t> x y)
	// result: (SRAV (SignExt8to64 x) (OR <t> (NEGV <t> (SGTU (ZeroExt16to64 y) (MOVVconst <typ.UInt64> [63]))) (ZeroExt16to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v5.AuxInt = 63
		v3.AddArg(v5)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v6 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v6.AddArg(y)
		v1.AddArg(v6)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpRsh8x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x32 <t> x y)
	// result: (SRAV (SignExt8to64 x) (OR <t> (NEGV <t> (SGTU (ZeroExt32to64 y) (MOVVconst <typ.UInt64> [63]))) (ZeroExt32to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v5.AuxInt = 63
		v3.AddArg(v5)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v6 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v6.AddArg(y)
		v1.AddArg(v6)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpRsh8x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x64 <t> x y)
	// result: (SRAV (SignExt8to64 x) (OR <t> (NEGV <t> (SGTU y (MOVVconst <typ.UInt64> [63]))) y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v3.AddArg(y)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v4.AuxInt = 63
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpRsh8x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x8 <t> x y)
	// result: (SRAV (SignExt8to64 x) (OR <t> (NEGV <t> (SGTU (ZeroExt8to64 y) (MOVVconst <typ.UInt64> [63]))) (ZeroExt8to64 y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpMIPS64SRAV)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64OR, t)
		v2 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v3 := b.NewValue0(v.Pos, OpMIPS64SGTU, typ.Bool)
		v4 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v4.AddArg(y)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v5.AuxInt = 63
		v3.AddArg(v5)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v6 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v6.AddArg(y)
		v1.AddArg(v6)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueMIPS64_OpSelect0(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Select0 (DIVVU _ (MOVVconst [1])))
	// result: (MOVVconst [0])
	for {
		if v_0.Op != OpMIPS64DIVVU {
			break
		}
		_ = v_0.Args[1]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst || v_0_1.AuxInt != 1 {
			break
		}
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = 0
		return true
	}
	// match: (Select0 (DIVVU x (MOVVconst [c])))
	// cond: isPowerOfTwo(c)
	// result: (ANDconst [c-1] x)
	for {
		if v_0.Op != OpMIPS64DIVVU {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpMIPS64ANDconst)
		v.AuxInt = c - 1
		v.AddArg(x)
		return true
	}
	// match: (Select0 (DIVV (MOVVconst [c]) (MOVVconst [d])))
	// result: (MOVVconst [c%d])
	for {
		if v_0.Op != OpMIPS64DIVV {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = c % d
		return true
	}
	// match: (Select0 (DIVVU (MOVVconst [c]) (MOVVconst [d])))
	// result: (MOVVconst [int64(uint64(c)%uint64(d))])
	for {
		if v_0.Op != OpMIPS64DIVVU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = int64(uint64(c) % uint64(d))
		return true
	}
	return false
}
func rewriteValueMIPS64_OpSelect1(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Select1 (MULVU x (MOVVconst [-1])))
	// result: (NEGV x)
	for {
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			x := v_0_0
			if v_0_1.Op != OpMIPS64MOVVconst || v_0_1.AuxInt != -1 {
				continue
			}
			v.reset(OpMIPS64NEGV)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (Select1 (MULVU _ (MOVVconst [0])))
	// result: (MOVVconst [0])
	for {
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_1.Op != OpMIPS64MOVVconst || v_0_1.AuxInt != 0 {
				continue
			}
			v.reset(OpMIPS64MOVVconst)
			v.AuxInt = 0
			return true
		}
		break
	}
	// match: (Select1 (MULVU x (MOVVconst [1])))
	// result: x
	for {
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			x := v_0_0
			if v_0_1.Op != OpMIPS64MOVVconst || v_0_1.AuxInt != 1 {
				continue
			}
			v.reset(OpCopy)
			v.Type = x.Type
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (Select1 (MULVU x (MOVVconst [c])))
	// cond: isPowerOfTwo(c)
	// result: (SLLVconst [log2(c)] x)
	for {
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			x := v_0_0
			if v_0_1.Op != OpMIPS64MOVVconst {
				continue
			}
			c := v_0_1.AuxInt
			if !(isPowerOfTwo(c)) {
				continue
			}
			v.reset(OpMIPS64SLLVconst)
			v.AuxInt = log2(c)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (Select1 (DIVVU x (MOVVconst [1])))
	// result: x
	for {
		if v_0.Op != OpMIPS64DIVVU {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst || v_0_1.AuxInt != 1 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (Select1 (DIVVU x (MOVVconst [c])))
	// cond: isPowerOfTwo(c)
	// result: (SRLVconst [log2(c)] x)
	for {
		if v_0.Op != OpMIPS64DIVVU {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpMIPS64SRLVconst)
		v.AuxInt = log2(c)
		v.AddArg(x)
		return true
	}
	// match: (Select1 (MULVU (MOVVconst [c]) (MOVVconst [d])))
	// result: (MOVVconst [c*d])
	for {
		if v_0.Op != OpMIPS64MULVU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		for _i0 := 0; _i0 <= 1; _i0, v_0_0, v_0_1 = _i0+1, v_0_1, v_0_0 {
			if v_0_0.Op != OpMIPS64MOVVconst {
				continue
			}
			c := v_0_0.AuxInt
			if v_0_1.Op != OpMIPS64MOVVconst {
				continue
			}
			d := v_0_1.AuxInt
			v.reset(OpMIPS64MOVVconst)
			v.AuxInt = c * d
			return true
		}
		break
	}
	// match: (Select1 (DIVV (MOVVconst [c]) (MOVVconst [d])))
	// result: (MOVVconst [c/d])
	for {
		if v_0.Op != OpMIPS64DIVV {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = c / d
		return true
	}
	// match: (Select1 (DIVVU (MOVVconst [c]) (MOVVconst [d])))
	// result: (MOVVconst [int64(uint64(c)/uint64(d))])
	for {
		if v_0.Op != OpMIPS64DIVVU {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpMIPS64MOVVconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpMIPS64MOVVconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpMIPS64MOVVconst)
		v.AuxInt = int64(uint64(c) / uint64(d))
		return true
	}
	return false
}
func rewriteValueMIPS64_OpSlicemask(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (Slicemask <t> x)
	// result: (SRAVconst (NEGV <t> x) [63])
	for {
		t := v.Type
		x := v_0
		v.reset(OpMIPS64SRAVconst)
		v.AuxInt = 63
		v0 := b.NewValue0(v.Pos, OpMIPS64NEGV, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueMIPS64_OpStore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 1
	// result: (MOVBstore ptr val mem)
	for {
		t := v.Aux
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 1) {
			break
		}
		v.reset(OpMIPS64MOVBstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 2
	// result: (MOVHstore ptr val mem)
	for {
		t := v.Aux
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 2) {
			break
		}
		v.reset(OpMIPS64MOVHstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 4 && !is32BitFloat(val.Type)
	// result: (MOVWstore ptr val mem)
	for {
		t := v.Aux
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 4 && !is32BitFloat(val.Type)) {
			break
		}
		v.reset(OpMIPS64MOVWstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 8 && !is64BitFloat(val.Type)
	// result: (MOVVstore ptr val mem)
	for {
		t := v.Aux
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 8 && !is64BitFloat(val.Type)) {
			break
		}
		v.reset(OpMIPS64MOVVstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 4 && is32BitFloat(val.Type)
	// result: (MOVFstore ptr val mem)
	for {
		t := v.Aux
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 4 && is32BitFloat(val.Type)) {
			break
		}
		v.reset(OpMIPS64MOVFstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.(*types.Type).Size() == 8 && is64BitFloat(val.Type)
	// result: (MOVDstore ptr val mem)
	for {
		t := v.Aux
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.(*types.Type).Size() == 8 && is64BitFloat(val.Type)) {
			break
		}
		v.reset(OpMIPS64MOVDstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueMIPS64_OpZero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: (Zero [0] _ mem)
	// result: mem
	for {
		if v.AuxInt != 0 {
			break
		}
		mem := v_1
		v.reset(OpCopy)
		v.Type = mem.Type
		v.AddArg(mem)
		return true
	}
	// match: (Zero [1] ptr mem)
	// result: (MOVBstore ptr (MOVVconst [0]) mem)
	for {
		if v.AuxInt != 1 {
			break
		}
		ptr := v_0
		mem := v_1
		v.reset(OpMIPS64MOVBstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [2] {t} ptr mem)
	// cond: t.(*types.Type).Alignment()%2 == 0
	// result: (MOVHstore ptr (MOVVconst [0]) mem)
	for {
		if v.AuxInt != 2 {
			break
		}
		t := v.Aux
		ptr := v_0
		mem := v_1
		if !(t.(*types.Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpMIPS64MOVHstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [2] ptr mem)
	// result: (MOVBstore [1] ptr (MOVVconst [0]) (MOVBstore [0] ptr (MOVVconst [0]) mem))
	for {
		if v.AuxInt != 2 {
			break
		}
		ptr := v_0
		mem := v_1
		v.reset(OpMIPS64MOVBstore)
		v.AuxInt = 1
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, types.TypeMem)
		v1.AuxInt = 0
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [4] {t} ptr mem)
	// cond: t.(*types.Type).Alignment()%4 == 0
	// result: (MOVWstore ptr (MOVVconst [0]) mem)
	for {
		if v.AuxInt != 4 {
			break
		}
		t := v.Aux
		ptr := v_0
		mem := v_1
		if !(t.(*types.Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpMIPS64MOVWstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [4] {t} ptr mem)
	// cond: t.(*types.Type).Alignment()%2 == 0
	// result: (MOVHstore [2] ptr (MOVVconst [0]) (MOVHstore [0] ptr (MOVVconst [0]) mem))
	for {
		if v.AuxInt != 4 {
			break
		}
		t := v.Aux
		ptr := v_0
		mem := v_1
		if !(t.(*types.Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpMIPS64MOVHstore)
		v.AuxInt = 2
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, types.TypeMem)
		v1.AuxInt = 0
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [4] ptr mem)
	// result: (MOVBstore [3] ptr (MOVVconst [0]) (MOVBstore [2] ptr (MOVVconst [0]) (MOVBstore [1] ptr (MOVVconst [0]) (MOVBstore [0] ptr (MOVVconst [0]) mem))))
	for {
		if v.AuxInt != 4 {
			break
		}
		ptr := v_0
		mem := v_1
		v.reset(OpMIPS64MOVBstore)
		v.AuxInt = 3
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, types.TypeMem)
		v1.AuxInt = 2
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, types.TypeMem)
		v3.AuxInt = 1
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, types.TypeMem)
		v5.AuxInt = 0
		v5.AddArg(ptr)
		v6 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v6.AuxInt = 0
		v5.AddArg(v6)
		v5.AddArg(mem)
		v3.AddArg(v5)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [8] {t} ptr mem)
	// cond: t.(*types.Type).Alignment()%8 == 0
	// result: (MOVVstore ptr (MOVVconst [0]) mem)
	for {
		if v.AuxInt != 8 {
			break
		}
		t := v.Aux
		ptr := v_0
		mem := v_1
		if !(t.(*types.Type).Alignment()%8 == 0) {
			break
		}
		v.reset(OpMIPS64MOVVstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [8] {t} ptr mem)
	// cond: t.(*types.Type).Alignment()%4 == 0
	// result: (MOVWstore [4] ptr (MOVVconst [0]) (MOVWstore [0] ptr (MOVVconst [0]) mem))
	for {
		if v.AuxInt != 8 {
			break
		}
		t := v.Aux
		ptr := v_0
		mem := v_1
		if !(t.(*types.Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpMIPS64MOVWstore)
		v.AuxInt = 4
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVWstore, types.TypeMem)
		v1.AuxInt = 0
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [8] {t} ptr mem)
	// cond: t.(*types.Type).Alignment()%2 == 0
	// result: (MOVHstore [6] ptr (MOVVconst [0]) (MOVHstore [4] ptr (MOVVconst [0]) (MOVHstore [2] ptr (MOVVconst [0]) (MOVHstore [0] ptr (MOVVconst [0]) mem))))
	for {
		if v.AuxInt != 8 {
			break
		}
		t := v.Aux
		ptr := v_0
		mem := v_1
		if !(t.(*types.Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpMIPS64MOVHstore)
		v.AuxInt = 6
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, types.TypeMem)
		v1.AuxInt = 4
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, types.TypeMem)
		v3.AuxInt = 2
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, types.TypeMem)
		v5.AuxInt = 0
		v5.AddArg(ptr)
		v6 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v6.AuxInt = 0
		v5.AddArg(v6)
		v5.AddArg(mem)
		v3.AddArg(v5)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [3] ptr mem)
	// result: (MOVBstore [2] ptr (MOVVconst [0]) (MOVBstore [1] ptr (MOVVconst [0]) (MOVBstore [0] ptr (MOVVconst [0]) mem)))
	for {
		if v.AuxInt != 3 {
			break
		}
		ptr := v_0
		mem := v_1
		v.reset(OpMIPS64MOVBstore)
		v.AuxInt = 2
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, types.TypeMem)
		v1.AuxInt = 1
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVBstore, types.TypeMem)
		v3.AuxInt = 0
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [6] {t} ptr mem)
	// cond: t.(*types.Type).Alignment()%2 == 0
	// result: (MOVHstore [4] ptr (MOVVconst [0]) (MOVHstore [2] ptr (MOVVconst [0]) (MOVHstore [0] ptr (MOVVconst [0]) mem)))
	for {
		if v.AuxInt != 6 {
			break
		}
		t := v.Aux
		ptr := v_0
		mem := v_1
		if !(t.(*types.Type).Alignment()%2 == 0) {
			break
		}
		v.reset(OpMIPS64MOVHstore)
		v.AuxInt = 4
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, types.TypeMem)
		v1.AuxInt = 2
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVHstore, types.TypeMem)
		v3.AuxInt = 0
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [12] {t} ptr mem)
	// cond: t.(*types.Type).Alignment()%4 == 0
	// result: (MOVWstore [8] ptr (MOVVconst [0]) (MOVWstore [4] ptr (MOVVconst [0]) (MOVWstore [0] ptr (MOVVconst [0]) mem)))
	for {
		if v.AuxInt != 12 {
			break
		}
		t := v.Aux
		ptr := v_0
		mem := v_1
		if !(t.(*types.Type).Alignment()%4 == 0) {
			break
		}
		v.reset(OpMIPS64MOVWstore)
		v.AuxInt = 8
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVWstore, types.TypeMem)
		v1.AuxInt = 4
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVWstore, types.TypeMem)
		v3.AuxInt = 0
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [16] {t} ptr mem)
	// cond: t.(*types.Type).Alignment()%8 == 0
	// result: (MOVVstore [8] ptr (MOVVconst [0]) (MOVVstore [0] ptr (MOVVconst [0]) mem))
	for {
		if v.AuxInt != 16 {
			break
		}
		t := v.Aux
		ptr := v_0
		mem := v_1
		if !(t.(*types.Type).Alignment()%8 == 0) {
			break
		}
		v.reset(OpMIPS64MOVVstore)
		v.AuxInt = 8
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVVstore, types.TypeMem)
		v1.AuxInt = 0
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [24] {t} ptr mem)
	// cond: t.(*types.Type).Alignment()%8 == 0
	// result: (MOVVstore [16] ptr (MOVVconst [0]) (MOVVstore [8] ptr (MOVVconst [0]) (MOVVstore [0] ptr (MOVVconst [0]) mem)))
	for {
		if v.AuxInt != 24 {
			break
		}
		t := v.Aux
		ptr := v_0
		mem := v_1
		if !(t.(*types.Type).Alignment()%8 == 0) {
			break
		}
		v.reset(OpMIPS64MOVVstore)
		v.AuxInt = 16
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpMIPS64MOVVstore, types.TypeMem)
		v1.AuxInt = 8
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpMIPS64MOVVstore, types.TypeMem)
		v3.AuxInt = 0
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpMIPS64MOVVconst, typ.UInt64)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [s] {t} ptr mem)
	// cond: s%8 == 0 && s > 24 && s <= 8*128 && t.(*types.Type).Alignment()%8 == 0 && !config.noDuffDevice
	// result: (DUFFZERO [8 * (128 - s/8)] ptr mem)
	for {
		s := v.AuxInt
		t := v.Aux
		ptr := v_0
		mem := v_1
		if !(s%8 == 0 && s > 24 && s <= 8*128 && t.(*types.Type).Alignment()%8 == 0 && !config.noDuffDevice) {
			break
		}
		v.reset(OpMIPS64DUFFZERO)
		v.AuxInt = 8 * (128 - s/8)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [s] {t} ptr mem)
	// cond: (s > 8*128 || config.noDuffDevice) || t.(*types.Type).Alignment()%8 != 0
	// result: (LoweredZero [t.(*types.Type).Alignment()] ptr (ADDVconst <ptr.Type> ptr [s-moveSize(t.(*types.Type).Alignment(), config)]) mem)
	for {
		s := v.AuxInt
		t := v.Aux
		ptr := v_0
		mem := v_1
		if !((s > 8*128 || config.noDuffDevice) || t.(*types.Type).Alignment()%8 != 0) {
			break
		}
		v.reset(OpMIPS64LoweredZero)
		v.AuxInt = t.(*types.Type).Alignment()
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpMIPS64ADDVconst, ptr.Type)
		v0.AuxInt = s - moveSize(t.(*types.Type).Alignment(), config)
		v0.AddArg(ptr)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteBlockMIPS64(b *Block) bool {
	switch b.Kind {
	case BlockMIPS64EQ:
		// match: (EQ (FPFlagTrue cmp) yes no)
		// result: (FPF cmp yes no)
		for b.Controls[0].Op == OpMIPS64FPFlagTrue {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockMIPS64FPF)
			b.AddControl(cmp)
			return true
		}
		// match: (EQ (FPFlagFalse cmp) yes no)
		// result: (FPT cmp yes no)
		for b.Controls[0].Op == OpMIPS64FPFlagFalse {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockMIPS64FPT)
			b.AddControl(cmp)
			return true
		}
		// match: (EQ (XORconst [1] cmp:(SGT _ _)) yes no)
		// result: (NE cmp yes no)
		for b.Controls[0].Op == OpMIPS64XORconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpMIPS64SGT {
				break
			}
			_ = cmp.Args[1]
			b.Reset(BlockMIPS64NE)
			b.AddControl(cmp)
			return true
		}
		// match: (EQ (XORconst [1] cmp:(SGTU _ _)) yes no)
		// result: (NE cmp yes no)
		for b.Controls[0].Op == OpMIPS64XORconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpMIPS64SGTU {
				break
			}
			_ = cmp.Args[1]
			b.Reset(BlockMIPS64NE)
			b.AddControl(cmp)
			return true
		}
		// match: (EQ (XORconst [1] cmp:(SGTconst _)) yes no)
		// result: (NE cmp yes no)
		for b.Controls[0].Op == OpMIPS64XORconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpMIPS64SGTconst {
				break
			}
			b.Reset(BlockMIPS64NE)
			b.AddControl(cmp)
			return true
		}
		// match: (EQ (XORconst [1] cmp:(SGTUconst _)) yes no)
		// result: (NE cmp yes no)
		for b.Controls[0].Op == OpMIPS64XORconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpMIPS64SGTUconst {
				break
			}
			b.Reset(BlockMIPS64NE)
			b.AddControl(cmp)
			return true
		}
		// match: (EQ (SGTUconst [1] x) yes no)
		// result: (NE x yes no)
		for b.Controls[0].Op == OpMIPS64SGTUconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 1 {
				break
			}
			x := v_0.Args[0]
			b.Reset(BlockMIPS64NE)
			b.AddControl(x)
			return true
		}
		// match: (EQ (SGTU x (MOVVconst [0])) yes no)
		// result: (EQ x yes no)
		for b.Controls[0].Op == OpMIPS64SGTU {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpMIPS64MOVVconst || v_0_1.AuxInt != 0 {
				break
			}
			b.Reset(BlockMIPS64EQ)
			b.AddControl(x)
			return true
		}
		// match: (EQ (SGTconst [0] x) yes no)
		// result: (GEZ x yes no)
		for b.Controls[0].Op == OpMIPS64SGTconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			x := v_0.Args[0]
			b.Reset(BlockMIPS64GEZ)
			b.AddControl(x)
			return true
		}
		// match: (EQ (SGT x (MOVVconst [0])) yes no)
		// result: (LEZ x yes no)
		for b.Controls[0].Op == OpMIPS64SGT {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpMIPS64MOVVconst || v_0_1.AuxInt != 0 {
				break
			}
			b.Reset(BlockMIPS64LEZ)
			b.AddControl(x)
			return true
		}
		// match: (EQ (MOVVconst [0]) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpMIPS64MOVVconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (EQ (MOVVconst [c]) yes no)
		// cond: c != 0
		// result: (First no yes)
		for b.Controls[0].Op == OpMIPS64MOVVconst {
			v_0 := b.Controls[0]
			c := v_0.AuxInt
			if !(c != 0) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
	case BlockMIPS64GEZ:
		// match: (GEZ (MOVVconst [c]) yes no)
		// cond: c >= 0
		// result: (First yes no)
		for b.Controls[0].Op == OpMIPS64MOVVconst {
			v_0 := b.Controls[0]
			c := v_0.AuxInt
			if !(c >= 0) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (GEZ (MOVVconst [c]) yes no)
		// cond: c < 0
		// result: (First no yes)
		for b.Controls[0].Op == OpMIPS64MOVVconst {
			v_0 := b.Controls[0]
			c := v_0.AuxInt
			if !(c < 0) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
	case BlockMIPS64GTZ:
		// match: (GTZ (MOVVconst [c]) yes no)
		// cond: c > 0
		// result: (First yes no)
		for b.Controls[0].Op == OpMIPS64MOVVconst {
			v_0 := b.Controls[0]
			c := v_0.AuxInt
			if !(c > 0) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (GTZ (MOVVconst [c]) yes no)
		// cond: c <= 0
		// result: (First no yes)
		for b.Controls[0].Op == OpMIPS64MOVVconst {
			v_0 := b.Controls[0]
			c := v_0.AuxInt
			if !(c <= 0) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
	case BlockIf:
		// match: (If cond yes no)
		// result: (NE cond yes no)
		for {
			cond := b.Controls[0]
			b.Reset(BlockMIPS64NE)
			b.AddControl(cond)
			return true
		}
	case BlockMIPS64LEZ:
		// match: (LEZ (MOVVconst [c]) yes no)
		// cond: c <= 0
		// result: (First yes no)
		for b.Controls[0].Op == OpMIPS64MOVVconst {
			v_0 := b.Controls[0]
			c := v_0.AuxInt
			if !(c <= 0) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (LEZ (MOVVconst [c]) yes no)
		// cond: c > 0
		// result: (First no yes)
		for b.Controls[0].Op == OpMIPS64MOVVconst {
			v_0 := b.Controls[0]
			c := v_0.AuxInt
			if !(c > 0) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
	case BlockMIPS64LTZ:
		// match: (LTZ (MOVVconst [c]) yes no)
		// cond: c < 0
		// result: (First yes no)
		for b.Controls[0].Op == OpMIPS64MOVVconst {
			v_0 := b.Controls[0]
			c := v_0.AuxInt
			if !(c < 0) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
		// match: (LTZ (MOVVconst [c]) yes no)
		// cond: c >= 0
		// result: (First no yes)
		for b.Controls[0].Op == OpMIPS64MOVVconst {
			v_0 := b.Controls[0]
			c := v_0.AuxInt
			if !(c >= 0) {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
	case BlockMIPS64NE:
		// match: (NE (FPFlagTrue cmp) yes no)
		// result: (FPT cmp yes no)
		for b.Controls[0].Op == OpMIPS64FPFlagTrue {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockMIPS64FPT)
			b.AddControl(cmp)
			return true
		}
		// match: (NE (FPFlagFalse cmp) yes no)
		// result: (FPF cmp yes no)
		for b.Controls[0].Op == OpMIPS64FPFlagFalse {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockMIPS64FPF)
			b.AddControl(cmp)
			return true
		}
		// match: (NE (XORconst [1] cmp:(SGT _ _)) yes no)
		// result: (EQ cmp yes no)
		for b.Controls[0].Op == OpMIPS64XORconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpMIPS64SGT {
				break
			}
			_ = cmp.Args[1]
			b.Reset(BlockMIPS64EQ)
			b.AddControl(cmp)
			return true
		}
		// match: (NE (XORconst [1] cmp:(SGTU _ _)) yes no)
		// result: (EQ cmp yes no)
		for b.Controls[0].Op == OpMIPS64XORconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpMIPS64SGTU {
				break
			}
			_ = cmp.Args[1]
			b.Reset(BlockMIPS64EQ)
			b.AddControl(cmp)
			return true
		}
		// match: (NE (XORconst [1] cmp:(SGTconst _)) yes no)
		// result: (EQ cmp yes no)
		for b.Controls[0].Op == OpMIPS64XORconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpMIPS64SGTconst {
				break
			}
			b.Reset(BlockMIPS64EQ)
			b.AddControl(cmp)
			return true
		}
		// match: (NE (XORconst [1] cmp:(SGTUconst _)) yes no)
		// result: (EQ cmp yes no)
		for b.Controls[0].Op == OpMIPS64XORconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 1 {
				break
			}
			cmp := v_0.Args[0]
			if cmp.Op != OpMIPS64SGTUconst {
				break
			}
			b.Reset(BlockMIPS64EQ)
			b.AddControl(cmp)
			return true
		}
		// match: (NE (SGTUconst [1] x) yes no)
		// result: (EQ x yes no)
		for b.Controls[0].Op == OpMIPS64SGTUconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 1 {
				break
			}
			x := v_0.Args[0]
			b.Reset(BlockMIPS64EQ)
			b.AddControl(x)
			return true
		}
		// match: (NE (SGTU x (MOVVconst [0])) yes no)
		// result: (NE x yes no)
		for b.Controls[0].Op == OpMIPS64SGTU {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpMIPS64MOVVconst || v_0_1.AuxInt != 0 {
				break
			}
			b.Reset(BlockMIPS64NE)
			b.AddControl(x)
			return true
		}
		// match: (NE (SGTconst [0] x) yes no)
		// result: (LTZ x yes no)
		for b.Controls[0].Op == OpMIPS64SGTconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			x := v_0.Args[0]
			b.Reset(BlockMIPS64LTZ)
			b.AddControl(x)
			return true
		}
		// match: (NE (SGT x (MOVVconst [0])) yes no)
		// result: (GTZ x yes no)
		for b.Controls[0].Op == OpMIPS64SGT {
			v_0 := b.Controls[0]
			_ = v_0.Args[1]
			x := v_0.Args[0]
			v_0_1 := v_0.Args[1]
			if v_0_1.Op != OpMIPS64MOVVconst || v_0_1.AuxInt != 0 {
				break
			}
			b.Reset(BlockMIPS64GTZ)
			b.AddControl(x)
			return true
		}
		// match: (NE (MOVVconst [0]) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpMIPS64MOVVconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (NE (MOVVconst [c]) yes no)
		// cond: c != 0
		// result: (First yes no)
		for b.Controls[0].Op == OpMIPS64MOVVconst {
			v_0 := b.Controls[0]
			c := v_0.AuxInt
			if !(c != 0) {
				break
			}
			b.Reset(BlockFirst)
			return true
		}
	}
	return false
}
