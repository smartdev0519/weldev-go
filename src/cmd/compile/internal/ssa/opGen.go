// autogenerated: do not edit!
// generated from gen/*Ops.go
package ssa

import "cmd/internal/obj/x86"

const (
	blockInvalid BlockKind = iota

	BlockAMD64EQ
	BlockAMD64NE
	BlockAMD64LT
	BlockAMD64LE
	BlockAMD64GT
	BlockAMD64GE
	BlockAMD64ULT
	BlockAMD64ULE
	BlockAMD64UGT
	BlockAMD64UGE

	BlockExit
	BlockDead
	BlockPlain
	BlockIf
	BlockCall
)

var blockString = [...]string{
	blockInvalid: "BlockInvalid",

	BlockAMD64EQ:  "EQ",
	BlockAMD64NE:  "NE",
	BlockAMD64LT:  "LT",
	BlockAMD64LE:  "LE",
	BlockAMD64GT:  "GT",
	BlockAMD64GE:  "GE",
	BlockAMD64ULT: "ULT",
	BlockAMD64ULE: "ULE",
	BlockAMD64UGT: "UGT",
	BlockAMD64UGE: "UGE",

	BlockExit:  "Exit",
	BlockDead:  "Dead",
	BlockPlain: "Plain",
	BlockIf:    "If",
	BlockCall:  "Call",
}

func (k BlockKind) String() string { return blockString[k] }

const (
	OpInvalid Op = iota

	OpAMD64MULQ
	OpAMD64MULQconst
	OpAMD64SHLQ
	OpAMD64SHLQconst
	OpAMD64SHRQ
	OpAMD64SHRQconst
	OpAMD64SARQ
	OpAMD64SARQconst
	OpAMD64XORQconst
	OpAMD64CMPQ
	OpAMD64CMPQconst
	OpAMD64CMPL
	OpAMD64CMPW
	OpAMD64CMPB
	OpAMD64TESTQ
	OpAMD64TESTB
	OpAMD64SBBQcarrymask
	OpAMD64SETEQ
	OpAMD64SETNE
	OpAMD64SETL
	OpAMD64SETLE
	OpAMD64SETG
	OpAMD64SETGE
	OpAMD64SETB
	OpAMD64SETBE
	OpAMD64SETA
	OpAMD64SETAE
	OpAMD64CMOVQCC
	OpAMD64MOVBQSX
	OpAMD64MOVBQZX
	OpAMD64MOVWQSX
	OpAMD64MOVWQZX
	OpAMD64MOVLQSX
	OpAMD64MOVLQZX
	OpAMD64MOVQconst
	OpAMD64LEAQ
	OpAMD64LEAQ1
	OpAMD64LEAQ2
	OpAMD64LEAQ4
	OpAMD64LEAQ8
	OpAMD64MOVBload
	OpAMD64MOVBQZXload
	OpAMD64MOVBQSXload
	OpAMD64MOVWload
	OpAMD64MOVLload
	OpAMD64MOVQload
	OpAMD64MOVQloadidx8
	OpAMD64MOVBstore
	OpAMD64MOVWstore
	OpAMD64MOVLstore
	OpAMD64MOVQstore
	OpAMD64MOVQstoreidx8
	OpAMD64MOVXzero
	OpAMD64REPSTOSQ
	OpAMD64MOVQloadglobal
	OpAMD64MOVQstoreglobal
	OpAMD64CALLstatic
	OpAMD64CALLclosure
	OpAMD64REPMOVSB
	OpAMD64ADDQ
	OpAMD64ADDQconst
	OpAMD64ADDL
	OpAMD64ADDW
	OpAMD64ADDB
	OpAMD64SUBQ
	OpAMD64SUBQconst
	OpAMD64SUBL
	OpAMD64SUBW
	OpAMD64SUBB
	OpAMD64NEGQ
	OpAMD64NEGL
	OpAMD64NEGW
	OpAMD64NEGB
	OpAMD64MULL
	OpAMD64MULW
	OpAMD64ANDQ
	OpAMD64ANDQconst
	OpAMD64ANDL
	OpAMD64ANDW
	OpAMD64ANDB
	OpAMD64InvertFlags

	OpAdd8
	OpAdd16
	OpAdd32
	OpAdd64
	OpAdd8U
	OpAdd16U
	OpAdd32U
	OpAdd64U
	OpAddPtr
	OpSub8
	OpSub16
	OpSub32
	OpSub64
	OpSub8U
	OpSub16U
	OpSub32U
	OpSub64U
	OpMul8
	OpMul16
	OpMul32
	OpMul64
	OpMul8U
	OpMul16U
	OpMul32U
	OpMul64U
	OpMulPtr
	OpAnd8
	OpAnd16
	OpAnd32
	OpAnd64
	OpAnd8U
	OpAnd16U
	OpAnd32U
	OpAnd64U
	OpLsh8
	OpLsh16
	OpLsh32
	OpLsh64
	OpRsh8
	OpRsh8U
	OpRsh16
	OpRsh16U
	OpRsh32
	OpRsh32U
	OpRsh64
	OpRsh64U
	OpEq8
	OpEq16
	OpEq32
	OpEq64
	OpEqPtr
	OpEqFat
	OpNeq8
	OpNeq16
	OpNeq32
	OpNeq64
	OpNeqPtr
	OpNeqFat
	OpLess8
	OpLess8U
	OpLess16
	OpLess16U
	OpLess32
	OpLess32U
	OpLess64
	OpLess64U
	OpLeq8
	OpLeq8U
	OpLeq16
	OpLeq16U
	OpLeq32
	OpLeq32U
	OpLeq64
	OpLeq64U
	OpGreater8
	OpGreater8U
	OpGreater16
	OpGreater16U
	OpGreater32
	OpGreater32U
	OpGreater64
	OpGreater64U
	OpGeq8
	OpGeq8U
	OpGeq16
	OpGeq16U
	OpGeq32
	OpGeq32U
	OpGeq64
	OpGeq64U
	OpNot
	OpNeg8
	OpNeg16
	OpNeg32
	OpNeg64
	OpNeg8U
	OpNeg16U
	OpNeg32U
	OpNeg64U
	OpPhi
	OpCopy
	OpConst
	OpArg
	OpAddr
	OpSP
	OpSB
	OpFunc
	OpLoad
	OpStore
	OpMove
	OpZero
	OpClosureCall
	OpStaticCall
	OpConvert
	OpConvNop
	OpIsNonNil
	OpIsInBounds
	OpArrayIndex
	OpPtrIndex
	OpOffPtr
	OpStructSelect
	OpSliceMake
	OpSlicePtr
	OpSliceLen
	OpSliceCap
	OpStringMake
	OpStringPtr
	OpStringLen
	OpStoreReg
	OpLoadReg
	OpFwdRef
)

var opcodeTable = [...]opInfo{
	{name: "OpInvalid"},

	{
		name: "MULQ",
		asm:  x86.AIMULQ,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "MULQconst",
		asm:  x86.AIMULQ,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SHLQ",
		asm:  x86.ASHLQ,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				2,     // .CX
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SHLQconst",
		asm:  x86.ASHLQ,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SHRQ",
		asm:  x86.ASHRQ,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				2,     // .CX
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SHRQconst",
		asm:  x86.ASHRQ,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SARQ",
		asm:  x86.ASARQ,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				2,     // .CX
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SARQconst",
		asm:  x86.ASARQ,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "XORQconst",
		asm:  x86.AXORQ,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "CMPQ",
		asm:  x86.ACMPQ,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				8589934592, // .FLAGS
			},
		},
	},
	{
		name: "CMPQconst",
		asm:  x86.ACMPQ,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				8589934592, // .FLAGS
			},
		},
	},
	{
		name: "CMPL",
		asm:  x86.ACMPL,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				8589934592, // .FLAGS
			},
		},
	},
	{
		name: "CMPW",
		asm:  x86.ACMPW,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				8589934592, // .FLAGS
			},
		},
	},
	{
		name: "CMPB",
		asm:  x86.ACMPB,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				8589934592, // .FLAGS
			},
		},
	},
	{
		name: "TESTQ",
		asm:  x86.ATESTQ,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				8589934592, // .FLAGS
			},
		},
	},
	{
		name: "TESTB",
		asm:  x86.ATESTB,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				8589934592, // .FLAGS
			},
		},
	},
	{
		name: "SBBQcarrymask",
		asm:  x86.ASBBQ,
		reg: regInfo{
			inputs: []regMask{
				8589934592, // .FLAGS
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SETEQ",
		asm:  x86.ASETEQ,
		reg: regInfo{
			inputs: []regMask{
				8589934592, // .FLAGS
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SETNE",
		asm:  x86.ASETNE,
		reg: regInfo{
			inputs: []regMask{
				8589934592, // .FLAGS
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SETL",
		asm:  x86.ASETLT,
		reg: regInfo{
			inputs: []regMask{
				8589934592, // .FLAGS
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SETLE",
		asm:  x86.ASETLE,
		reg: regInfo{
			inputs: []regMask{
				8589934592, // .FLAGS
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SETG",
		asm:  x86.ASETGT,
		reg: regInfo{
			inputs: []regMask{
				8589934592, // .FLAGS
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SETGE",
		asm:  x86.ASETGE,
		reg: regInfo{
			inputs: []regMask{
				8589934592, // .FLAGS
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SETB",
		asm:  x86.ASETCS,
		reg: regInfo{
			inputs: []regMask{
				8589934592, // .FLAGS
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SETBE",
		asm:  x86.ASETLS,
		reg: regInfo{
			inputs: []regMask{
				8589934592, // .FLAGS
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SETA",
		asm:  x86.ASETHI,
		reg: regInfo{
			inputs: []regMask{
				8589934592, // .FLAGS
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SETAE",
		asm:  x86.ASETCC,
		reg: regInfo{
			inputs: []regMask{
				8589934592, // .FLAGS
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "CMOVQCC",
		reg: regInfo{
			inputs: []regMask{
				8589934592, // .FLAGS
				65519,      // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65519,      // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "MOVBQSX",
		asm:  x86.AMOVBQSX,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "MOVBQZX",
		asm:  x86.AMOVBQZX,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "MOVWQSX",
		asm:  x86.AMOVWQSX,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "MOVWQZX",
		asm:  x86.AMOVWQZX,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "MOVLQSX",
		asm:  x86.AMOVLQSX,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "MOVLQZX",
		asm:  x86.AMOVLQZX,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "MOVQconst",
		reg: regInfo{
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "LEAQ",
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "LEAQ1",
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
				65535,      // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "LEAQ2",
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
				65535,      // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "LEAQ4",
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
				65535,      // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "LEAQ8",
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
				65535,      // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "MOVBload",
		asm:  x86.AMOVB,
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
				0,
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "MOVBQZXload",
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
				0,
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "MOVBQSXload",
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
				0,
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "MOVWload",
		asm:  x86.AMOVW,
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
				0,
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "MOVLload",
		asm:  x86.AMOVL,
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
				0,
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "MOVQload",
		asm:  x86.AMOVQ,
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
				0,
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "MOVQloadidx8",
		asm:  x86.AMOVQ,
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
				65535,      // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				0,
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "MOVBstore",
		asm:  x86.AMOVB,
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
				65535,      // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				0,
			},
		},
	},
	{
		name: "MOVWstore",
		asm:  x86.AMOVW,
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
				65535,      // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				0,
			},
		},
	},
	{
		name: "MOVLstore",
		asm:  x86.AMOVL,
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
				65535,      // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				0,
			},
		},
	},
	{
		name: "MOVQstore",
		asm:  x86.AMOVQ,
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
				65535,      // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				0,
			},
		},
	},
	{
		name: "MOVQstoreidx8",
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
				65535,      // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535,      // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				0,
			},
		},
	},
	{
		name: "MOVXzero",
		reg: regInfo{
			inputs: []regMask{
				4295032831, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15 .SB
				0,
			},
		},
	},
	{
		name: "REPSTOSQ",
		reg: regInfo{
			inputs: []regMask{
				128, // .DI
				2,   // .CX
			},
			clobbers: 131, // .AX .CX .DI
		},
	},
	{
		name: "MOVQloadglobal",
		reg:  regInfo{},
	},
	{
		name: "MOVQstoreglobal",
		reg:  regInfo{},
	},
	{
		name: "CALLstatic",
		reg:  regInfo{},
	},
	{
		name: "CALLclosure",
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				4,     // .DX
				0,
			},
		},
	},
	{
		name: "REPMOVSB",
		reg: regInfo{
			inputs: []regMask{
				128, // .DI
				64,  // .SI
				2,   // .CX
			},
			clobbers: 194, // .CX .SI .DI
		},
	},
	{
		name: "ADDQ",
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "ADDQconst",
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "ADDL",
		asm:  x86.AADDL,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "ADDW",
		asm:  x86.AADDW,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "ADDB",
		asm:  x86.AADDB,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SUBQ",
		asm:  x86.ASUBQ,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SUBQconst",
		asm:  x86.ASUBQ,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SUBL",
		asm:  x86.ASUBL,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SUBW",
		asm:  x86.ASUBW,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "SUBB",
		asm:  x86.ASUBB,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "NEGQ",
		asm:  x86.ANEGQ,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "NEGL",
		asm:  x86.ANEGL,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "NEGW",
		asm:  x86.ANEGW,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "NEGB",
		asm:  x86.ANEGB,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "MULL",
		asm:  x86.AIMULL,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "MULW",
		asm:  x86.AIMULW,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "ANDQ",
		asm:  x86.AANDQ,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "ANDQconst",
		asm:  x86.AANDQ,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "ANDL",
		asm:  x86.AANDL,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "ANDW",
		asm:  x86.AANDW,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "ANDB",
		asm:  x86.AANDB,
		reg: regInfo{
			inputs: []regMask{
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
				65535, // .AX .CX .DX .BX .SP .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
			outputs: []regMask{
				65519, // .AX .CX .DX .BX .BP .SI .DI .R8 .R9 .R10 .R11 .R12 .R13 .R14 .R15
			},
		},
	},
	{
		name: "InvertFlags",
		reg:  regInfo{},
	},

	{
		name:    "Add8",
		generic: true,
	},
	{
		name:    "Add16",
		generic: true,
	},
	{
		name:    "Add32",
		generic: true,
	},
	{
		name:    "Add64",
		generic: true,
	},
	{
		name:    "Add8U",
		generic: true,
	},
	{
		name:    "Add16U",
		generic: true,
	},
	{
		name:    "Add32U",
		generic: true,
	},
	{
		name:    "Add64U",
		generic: true,
	},
	{
		name:    "AddPtr",
		generic: true,
	},
	{
		name:    "Sub8",
		generic: true,
	},
	{
		name:    "Sub16",
		generic: true,
	},
	{
		name:    "Sub32",
		generic: true,
	},
	{
		name:    "Sub64",
		generic: true,
	},
	{
		name:    "Sub8U",
		generic: true,
	},
	{
		name:    "Sub16U",
		generic: true,
	},
	{
		name:    "Sub32U",
		generic: true,
	},
	{
		name:    "Sub64U",
		generic: true,
	},
	{
		name:    "Mul8",
		generic: true,
	},
	{
		name:    "Mul16",
		generic: true,
	},
	{
		name:    "Mul32",
		generic: true,
	},
	{
		name:    "Mul64",
		generic: true,
	},
	{
		name:    "Mul8U",
		generic: true,
	},
	{
		name:    "Mul16U",
		generic: true,
	},
	{
		name:    "Mul32U",
		generic: true,
	},
	{
		name:    "Mul64U",
		generic: true,
	},
	{
		name:    "MulPtr",
		generic: true,
	},
	{
		name:    "And8",
		generic: true,
	},
	{
		name:    "And16",
		generic: true,
	},
	{
		name:    "And32",
		generic: true,
	},
	{
		name:    "And64",
		generic: true,
	},
	{
		name:    "And8U",
		generic: true,
	},
	{
		name:    "And16U",
		generic: true,
	},
	{
		name:    "And32U",
		generic: true,
	},
	{
		name:    "And64U",
		generic: true,
	},
	{
		name:    "Lsh8",
		generic: true,
	},
	{
		name:    "Lsh16",
		generic: true,
	},
	{
		name:    "Lsh32",
		generic: true,
	},
	{
		name:    "Lsh64",
		generic: true,
	},
	{
		name:    "Rsh8",
		generic: true,
	},
	{
		name:    "Rsh8U",
		generic: true,
	},
	{
		name:    "Rsh16",
		generic: true,
	},
	{
		name:    "Rsh16U",
		generic: true,
	},
	{
		name:    "Rsh32",
		generic: true,
	},
	{
		name:    "Rsh32U",
		generic: true,
	},
	{
		name:    "Rsh64",
		generic: true,
	},
	{
		name:    "Rsh64U",
		generic: true,
	},
	{
		name:    "Eq8",
		generic: true,
	},
	{
		name:    "Eq16",
		generic: true,
	},
	{
		name:    "Eq32",
		generic: true,
	},
	{
		name:    "Eq64",
		generic: true,
	},
	{
		name:    "EqPtr",
		generic: true,
	},
	{
		name:    "EqFat",
		generic: true,
	},
	{
		name:    "Neq8",
		generic: true,
	},
	{
		name:    "Neq16",
		generic: true,
	},
	{
		name:    "Neq32",
		generic: true,
	},
	{
		name:    "Neq64",
		generic: true,
	},
	{
		name:    "NeqPtr",
		generic: true,
	},
	{
		name:    "NeqFat",
		generic: true,
	},
	{
		name:    "Less8",
		generic: true,
	},
	{
		name:    "Less8U",
		generic: true,
	},
	{
		name:    "Less16",
		generic: true,
	},
	{
		name:    "Less16U",
		generic: true,
	},
	{
		name:    "Less32",
		generic: true,
	},
	{
		name:    "Less32U",
		generic: true,
	},
	{
		name:    "Less64",
		generic: true,
	},
	{
		name:    "Less64U",
		generic: true,
	},
	{
		name:    "Leq8",
		generic: true,
	},
	{
		name:    "Leq8U",
		generic: true,
	},
	{
		name:    "Leq16",
		generic: true,
	},
	{
		name:    "Leq16U",
		generic: true,
	},
	{
		name:    "Leq32",
		generic: true,
	},
	{
		name:    "Leq32U",
		generic: true,
	},
	{
		name:    "Leq64",
		generic: true,
	},
	{
		name:    "Leq64U",
		generic: true,
	},
	{
		name:    "Greater8",
		generic: true,
	},
	{
		name:    "Greater8U",
		generic: true,
	},
	{
		name:    "Greater16",
		generic: true,
	},
	{
		name:    "Greater16U",
		generic: true,
	},
	{
		name:    "Greater32",
		generic: true,
	},
	{
		name:    "Greater32U",
		generic: true,
	},
	{
		name:    "Greater64",
		generic: true,
	},
	{
		name:    "Greater64U",
		generic: true,
	},
	{
		name:    "Geq8",
		generic: true,
	},
	{
		name:    "Geq8U",
		generic: true,
	},
	{
		name:    "Geq16",
		generic: true,
	},
	{
		name:    "Geq16U",
		generic: true,
	},
	{
		name:    "Geq32",
		generic: true,
	},
	{
		name:    "Geq32U",
		generic: true,
	},
	{
		name:    "Geq64",
		generic: true,
	},
	{
		name:    "Geq64U",
		generic: true,
	},
	{
		name:    "Not",
		generic: true,
	},
	{
		name:    "Neg8",
		generic: true,
	},
	{
		name:    "Neg16",
		generic: true,
	},
	{
		name:    "Neg32",
		generic: true,
	},
	{
		name:    "Neg64",
		generic: true,
	},
	{
		name:    "Neg8U",
		generic: true,
	},
	{
		name:    "Neg16U",
		generic: true,
	},
	{
		name:    "Neg32U",
		generic: true,
	},
	{
		name:    "Neg64U",
		generic: true,
	},
	{
		name:    "Phi",
		generic: true,
	},
	{
		name:    "Copy",
		generic: true,
	},
	{
		name:    "Const",
		generic: true,
	},
	{
		name:    "Arg",
		generic: true,
	},
	{
		name:    "Addr",
		generic: true,
	},
	{
		name:    "SP",
		generic: true,
	},
	{
		name:    "SB",
		generic: true,
	},
	{
		name:    "Func",
		generic: true,
	},
	{
		name:    "Load",
		generic: true,
	},
	{
		name:    "Store",
		generic: true,
	},
	{
		name:    "Move",
		generic: true,
	},
	{
		name:    "Zero",
		generic: true,
	},
	{
		name:    "ClosureCall",
		generic: true,
	},
	{
		name:    "StaticCall",
		generic: true,
	},
	{
		name:    "Convert",
		generic: true,
	},
	{
		name:    "ConvNop",
		generic: true,
	},
	{
		name:    "IsNonNil",
		generic: true,
	},
	{
		name:    "IsInBounds",
		generic: true,
	},
	{
		name:    "ArrayIndex",
		generic: true,
	},
	{
		name:    "PtrIndex",
		generic: true,
	},
	{
		name:    "OffPtr",
		generic: true,
	},
	{
		name:    "StructSelect",
		generic: true,
	},
	{
		name:    "SliceMake",
		generic: true,
	},
	{
		name:    "SlicePtr",
		generic: true,
	},
	{
		name:    "SliceLen",
		generic: true,
	},
	{
		name:    "SliceCap",
		generic: true,
	},
	{
		name:    "StringMake",
		generic: true,
	},
	{
		name:    "StringPtr",
		generic: true,
	},
	{
		name:    "StringLen",
		generic: true,
	},
	{
		name:    "StoreReg",
		generic: true,
	},
	{
		name:    "LoadReg",
		generic: true,
	},
	{
		name:    "FwdRef",
		generic: true,
	},
}

func (o Op) Asm() int       { return opcodeTable[o].asm }
func (o Op) String() string { return opcodeTable[o].name }
