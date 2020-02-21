// Code generated from gen/ARM.rules; DO NOT EDIT.
// generated with: cd gen; go run *.go

package ssa

import "cmd/internal/objabi"
import "cmd/compile/internal/types"

func rewriteValueARM(v *Value) bool {
	switch v.Op {
	case OpARMADC:
		return rewriteValueARM_OpARMADC(v)
	case OpARMADCconst:
		return rewriteValueARM_OpARMADCconst(v)
	case OpARMADCshiftLL:
		return rewriteValueARM_OpARMADCshiftLL(v)
	case OpARMADCshiftLLreg:
		return rewriteValueARM_OpARMADCshiftLLreg(v)
	case OpARMADCshiftRA:
		return rewriteValueARM_OpARMADCshiftRA(v)
	case OpARMADCshiftRAreg:
		return rewriteValueARM_OpARMADCshiftRAreg(v)
	case OpARMADCshiftRL:
		return rewriteValueARM_OpARMADCshiftRL(v)
	case OpARMADCshiftRLreg:
		return rewriteValueARM_OpARMADCshiftRLreg(v)
	case OpARMADD:
		return rewriteValueARM_OpARMADD(v)
	case OpARMADDD:
		return rewriteValueARM_OpARMADDD(v)
	case OpARMADDF:
		return rewriteValueARM_OpARMADDF(v)
	case OpARMADDS:
		return rewriteValueARM_OpARMADDS(v)
	case OpARMADDSshiftLL:
		return rewriteValueARM_OpARMADDSshiftLL(v)
	case OpARMADDSshiftLLreg:
		return rewriteValueARM_OpARMADDSshiftLLreg(v)
	case OpARMADDSshiftRA:
		return rewriteValueARM_OpARMADDSshiftRA(v)
	case OpARMADDSshiftRAreg:
		return rewriteValueARM_OpARMADDSshiftRAreg(v)
	case OpARMADDSshiftRL:
		return rewriteValueARM_OpARMADDSshiftRL(v)
	case OpARMADDSshiftRLreg:
		return rewriteValueARM_OpARMADDSshiftRLreg(v)
	case OpARMADDconst:
		return rewriteValueARM_OpARMADDconst(v)
	case OpARMADDshiftLL:
		return rewriteValueARM_OpARMADDshiftLL(v)
	case OpARMADDshiftLLreg:
		return rewriteValueARM_OpARMADDshiftLLreg(v)
	case OpARMADDshiftRA:
		return rewriteValueARM_OpARMADDshiftRA(v)
	case OpARMADDshiftRAreg:
		return rewriteValueARM_OpARMADDshiftRAreg(v)
	case OpARMADDshiftRL:
		return rewriteValueARM_OpARMADDshiftRL(v)
	case OpARMADDshiftRLreg:
		return rewriteValueARM_OpARMADDshiftRLreg(v)
	case OpARMAND:
		return rewriteValueARM_OpARMAND(v)
	case OpARMANDconst:
		return rewriteValueARM_OpARMANDconst(v)
	case OpARMANDshiftLL:
		return rewriteValueARM_OpARMANDshiftLL(v)
	case OpARMANDshiftLLreg:
		return rewriteValueARM_OpARMANDshiftLLreg(v)
	case OpARMANDshiftRA:
		return rewriteValueARM_OpARMANDshiftRA(v)
	case OpARMANDshiftRAreg:
		return rewriteValueARM_OpARMANDshiftRAreg(v)
	case OpARMANDshiftRL:
		return rewriteValueARM_OpARMANDshiftRL(v)
	case OpARMANDshiftRLreg:
		return rewriteValueARM_OpARMANDshiftRLreg(v)
	case OpARMBFX:
		return rewriteValueARM_OpARMBFX(v)
	case OpARMBFXU:
		return rewriteValueARM_OpARMBFXU(v)
	case OpARMBIC:
		return rewriteValueARM_OpARMBIC(v)
	case OpARMBICconst:
		return rewriteValueARM_OpARMBICconst(v)
	case OpARMBICshiftLL:
		return rewriteValueARM_OpARMBICshiftLL(v)
	case OpARMBICshiftLLreg:
		return rewriteValueARM_OpARMBICshiftLLreg(v)
	case OpARMBICshiftRA:
		return rewriteValueARM_OpARMBICshiftRA(v)
	case OpARMBICshiftRAreg:
		return rewriteValueARM_OpARMBICshiftRAreg(v)
	case OpARMBICshiftRL:
		return rewriteValueARM_OpARMBICshiftRL(v)
	case OpARMBICshiftRLreg:
		return rewriteValueARM_OpARMBICshiftRLreg(v)
	case OpARMCMN:
		return rewriteValueARM_OpARMCMN(v)
	case OpARMCMNconst:
		return rewriteValueARM_OpARMCMNconst(v)
	case OpARMCMNshiftLL:
		return rewriteValueARM_OpARMCMNshiftLL(v)
	case OpARMCMNshiftLLreg:
		return rewriteValueARM_OpARMCMNshiftLLreg(v)
	case OpARMCMNshiftRA:
		return rewriteValueARM_OpARMCMNshiftRA(v)
	case OpARMCMNshiftRAreg:
		return rewriteValueARM_OpARMCMNshiftRAreg(v)
	case OpARMCMNshiftRL:
		return rewriteValueARM_OpARMCMNshiftRL(v)
	case OpARMCMNshiftRLreg:
		return rewriteValueARM_OpARMCMNshiftRLreg(v)
	case OpARMCMOVWHSconst:
		return rewriteValueARM_OpARMCMOVWHSconst(v)
	case OpARMCMOVWLSconst:
		return rewriteValueARM_OpARMCMOVWLSconst(v)
	case OpARMCMP:
		return rewriteValueARM_OpARMCMP(v)
	case OpARMCMPD:
		return rewriteValueARM_OpARMCMPD(v)
	case OpARMCMPF:
		return rewriteValueARM_OpARMCMPF(v)
	case OpARMCMPconst:
		return rewriteValueARM_OpARMCMPconst(v)
	case OpARMCMPshiftLL:
		return rewriteValueARM_OpARMCMPshiftLL(v)
	case OpARMCMPshiftLLreg:
		return rewriteValueARM_OpARMCMPshiftLLreg(v)
	case OpARMCMPshiftRA:
		return rewriteValueARM_OpARMCMPshiftRA(v)
	case OpARMCMPshiftRAreg:
		return rewriteValueARM_OpARMCMPshiftRAreg(v)
	case OpARMCMPshiftRL:
		return rewriteValueARM_OpARMCMPshiftRL(v)
	case OpARMCMPshiftRLreg:
		return rewriteValueARM_OpARMCMPshiftRLreg(v)
	case OpARMEqual:
		return rewriteValueARM_OpARMEqual(v)
	case OpARMGreaterEqual:
		return rewriteValueARM_OpARMGreaterEqual(v)
	case OpARMGreaterEqualU:
		return rewriteValueARM_OpARMGreaterEqualU(v)
	case OpARMGreaterThan:
		return rewriteValueARM_OpARMGreaterThan(v)
	case OpARMGreaterThanU:
		return rewriteValueARM_OpARMGreaterThanU(v)
	case OpARMLessEqual:
		return rewriteValueARM_OpARMLessEqual(v)
	case OpARMLessEqualU:
		return rewriteValueARM_OpARMLessEqualU(v)
	case OpARMLessThan:
		return rewriteValueARM_OpARMLessThan(v)
	case OpARMLessThanU:
		return rewriteValueARM_OpARMLessThanU(v)
	case OpARMMOVBUload:
		return rewriteValueARM_OpARMMOVBUload(v)
	case OpARMMOVBUloadidx:
		return rewriteValueARM_OpARMMOVBUloadidx(v)
	case OpARMMOVBUreg:
		return rewriteValueARM_OpARMMOVBUreg(v)
	case OpARMMOVBload:
		return rewriteValueARM_OpARMMOVBload(v)
	case OpARMMOVBloadidx:
		return rewriteValueARM_OpARMMOVBloadidx(v)
	case OpARMMOVBreg:
		return rewriteValueARM_OpARMMOVBreg(v)
	case OpARMMOVBstore:
		return rewriteValueARM_OpARMMOVBstore(v)
	case OpARMMOVBstoreidx:
		return rewriteValueARM_OpARMMOVBstoreidx(v)
	case OpARMMOVDload:
		return rewriteValueARM_OpARMMOVDload(v)
	case OpARMMOVDstore:
		return rewriteValueARM_OpARMMOVDstore(v)
	case OpARMMOVFload:
		return rewriteValueARM_OpARMMOVFload(v)
	case OpARMMOVFstore:
		return rewriteValueARM_OpARMMOVFstore(v)
	case OpARMMOVHUload:
		return rewriteValueARM_OpARMMOVHUload(v)
	case OpARMMOVHUloadidx:
		return rewriteValueARM_OpARMMOVHUloadidx(v)
	case OpARMMOVHUreg:
		return rewriteValueARM_OpARMMOVHUreg(v)
	case OpARMMOVHload:
		return rewriteValueARM_OpARMMOVHload(v)
	case OpARMMOVHloadidx:
		return rewriteValueARM_OpARMMOVHloadidx(v)
	case OpARMMOVHreg:
		return rewriteValueARM_OpARMMOVHreg(v)
	case OpARMMOVHstore:
		return rewriteValueARM_OpARMMOVHstore(v)
	case OpARMMOVHstoreidx:
		return rewriteValueARM_OpARMMOVHstoreidx(v)
	case OpARMMOVWload:
		return rewriteValueARM_OpARMMOVWload(v)
	case OpARMMOVWloadidx:
		return rewriteValueARM_OpARMMOVWloadidx(v)
	case OpARMMOVWloadshiftLL:
		return rewriteValueARM_OpARMMOVWloadshiftLL(v)
	case OpARMMOVWloadshiftRA:
		return rewriteValueARM_OpARMMOVWloadshiftRA(v)
	case OpARMMOVWloadshiftRL:
		return rewriteValueARM_OpARMMOVWloadshiftRL(v)
	case OpARMMOVWreg:
		return rewriteValueARM_OpARMMOVWreg(v)
	case OpARMMOVWstore:
		return rewriteValueARM_OpARMMOVWstore(v)
	case OpARMMOVWstoreidx:
		return rewriteValueARM_OpARMMOVWstoreidx(v)
	case OpARMMOVWstoreshiftLL:
		return rewriteValueARM_OpARMMOVWstoreshiftLL(v)
	case OpARMMOVWstoreshiftRA:
		return rewriteValueARM_OpARMMOVWstoreshiftRA(v)
	case OpARMMOVWstoreshiftRL:
		return rewriteValueARM_OpARMMOVWstoreshiftRL(v)
	case OpARMMUL:
		return rewriteValueARM_OpARMMUL(v)
	case OpARMMULA:
		return rewriteValueARM_OpARMMULA(v)
	case OpARMMULD:
		return rewriteValueARM_OpARMMULD(v)
	case OpARMMULF:
		return rewriteValueARM_OpARMMULF(v)
	case OpARMMULS:
		return rewriteValueARM_OpARMMULS(v)
	case OpARMMVN:
		return rewriteValueARM_OpARMMVN(v)
	case OpARMMVNshiftLL:
		return rewriteValueARM_OpARMMVNshiftLL(v)
	case OpARMMVNshiftLLreg:
		return rewriteValueARM_OpARMMVNshiftLLreg(v)
	case OpARMMVNshiftRA:
		return rewriteValueARM_OpARMMVNshiftRA(v)
	case OpARMMVNshiftRAreg:
		return rewriteValueARM_OpARMMVNshiftRAreg(v)
	case OpARMMVNshiftRL:
		return rewriteValueARM_OpARMMVNshiftRL(v)
	case OpARMMVNshiftRLreg:
		return rewriteValueARM_OpARMMVNshiftRLreg(v)
	case OpARMNEGD:
		return rewriteValueARM_OpARMNEGD(v)
	case OpARMNEGF:
		return rewriteValueARM_OpARMNEGF(v)
	case OpARMNMULD:
		return rewriteValueARM_OpARMNMULD(v)
	case OpARMNMULF:
		return rewriteValueARM_OpARMNMULF(v)
	case OpARMNotEqual:
		return rewriteValueARM_OpARMNotEqual(v)
	case OpARMOR:
		return rewriteValueARM_OpARMOR(v)
	case OpARMORconst:
		return rewriteValueARM_OpARMORconst(v)
	case OpARMORshiftLL:
		return rewriteValueARM_OpARMORshiftLL(v)
	case OpARMORshiftLLreg:
		return rewriteValueARM_OpARMORshiftLLreg(v)
	case OpARMORshiftRA:
		return rewriteValueARM_OpARMORshiftRA(v)
	case OpARMORshiftRAreg:
		return rewriteValueARM_OpARMORshiftRAreg(v)
	case OpARMORshiftRL:
		return rewriteValueARM_OpARMORshiftRL(v)
	case OpARMORshiftRLreg:
		return rewriteValueARM_OpARMORshiftRLreg(v)
	case OpARMRSB:
		return rewriteValueARM_OpARMRSB(v)
	case OpARMRSBSshiftLL:
		return rewriteValueARM_OpARMRSBSshiftLL(v)
	case OpARMRSBSshiftLLreg:
		return rewriteValueARM_OpARMRSBSshiftLLreg(v)
	case OpARMRSBSshiftRA:
		return rewriteValueARM_OpARMRSBSshiftRA(v)
	case OpARMRSBSshiftRAreg:
		return rewriteValueARM_OpARMRSBSshiftRAreg(v)
	case OpARMRSBSshiftRL:
		return rewriteValueARM_OpARMRSBSshiftRL(v)
	case OpARMRSBSshiftRLreg:
		return rewriteValueARM_OpARMRSBSshiftRLreg(v)
	case OpARMRSBconst:
		return rewriteValueARM_OpARMRSBconst(v)
	case OpARMRSBshiftLL:
		return rewriteValueARM_OpARMRSBshiftLL(v)
	case OpARMRSBshiftLLreg:
		return rewriteValueARM_OpARMRSBshiftLLreg(v)
	case OpARMRSBshiftRA:
		return rewriteValueARM_OpARMRSBshiftRA(v)
	case OpARMRSBshiftRAreg:
		return rewriteValueARM_OpARMRSBshiftRAreg(v)
	case OpARMRSBshiftRL:
		return rewriteValueARM_OpARMRSBshiftRL(v)
	case OpARMRSBshiftRLreg:
		return rewriteValueARM_OpARMRSBshiftRLreg(v)
	case OpARMRSCconst:
		return rewriteValueARM_OpARMRSCconst(v)
	case OpARMRSCshiftLL:
		return rewriteValueARM_OpARMRSCshiftLL(v)
	case OpARMRSCshiftLLreg:
		return rewriteValueARM_OpARMRSCshiftLLreg(v)
	case OpARMRSCshiftRA:
		return rewriteValueARM_OpARMRSCshiftRA(v)
	case OpARMRSCshiftRAreg:
		return rewriteValueARM_OpARMRSCshiftRAreg(v)
	case OpARMRSCshiftRL:
		return rewriteValueARM_OpARMRSCshiftRL(v)
	case OpARMRSCshiftRLreg:
		return rewriteValueARM_OpARMRSCshiftRLreg(v)
	case OpARMSBC:
		return rewriteValueARM_OpARMSBC(v)
	case OpARMSBCconst:
		return rewriteValueARM_OpARMSBCconst(v)
	case OpARMSBCshiftLL:
		return rewriteValueARM_OpARMSBCshiftLL(v)
	case OpARMSBCshiftLLreg:
		return rewriteValueARM_OpARMSBCshiftLLreg(v)
	case OpARMSBCshiftRA:
		return rewriteValueARM_OpARMSBCshiftRA(v)
	case OpARMSBCshiftRAreg:
		return rewriteValueARM_OpARMSBCshiftRAreg(v)
	case OpARMSBCshiftRL:
		return rewriteValueARM_OpARMSBCshiftRL(v)
	case OpARMSBCshiftRLreg:
		return rewriteValueARM_OpARMSBCshiftRLreg(v)
	case OpARMSLL:
		return rewriteValueARM_OpARMSLL(v)
	case OpARMSLLconst:
		return rewriteValueARM_OpARMSLLconst(v)
	case OpARMSRA:
		return rewriteValueARM_OpARMSRA(v)
	case OpARMSRAcond:
		return rewriteValueARM_OpARMSRAcond(v)
	case OpARMSRAconst:
		return rewriteValueARM_OpARMSRAconst(v)
	case OpARMSRL:
		return rewriteValueARM_OpARMSRL(v)
	case OpARMSRLconst:
		return rewriteValueARM_OpARMSRLconst(v)
	case OpARMSUB:
		return rewriteValueARM_OpARMSUB(v)
	case OpARMSUBD:
		return rewriteValueARM_OpARMSUBD(v)
	case OpARMSUBF:
		return rewriteValueARM_OpARMSUBF(v)
	case OpARMSUBS:
		return rewriteValueARM_OpARMSUBS(v)
	case OpARMSUBSshiftLL:
		return rewriteValueARM_OpARMSUBSshiftLL(v)
	case OpARMSUBSshiftLLreg:
		return rewriteValueARM_OpARMSUBSshiftLLreg(v)
	case OpARMSUBSshiftRA:
		return rewriteValueARM_OpARMSUBSshiftRA(v)
	case OpARMSUBSshiftRAreg:
		return rewriteValueARM_OpARMSUBSshiftRAreg(v)
	case OpARMSUBSshiftRL:
		return rewriteValueARM_OpARMSUBSshiftRL(v)
	case OpARMSUBSshiftRLreg:
		return rewriteValueARM_OpARMSUBSshiftRLreg(v)
	case OpARMSUBconst:
		return rewriteValueARM_OpARMSUBconst(v)
	case OpARMSUBshiftLL:
		return rewriteValueARM_OpARMSUBshiftLL(v)
	case OpARMSUBshiftLLreg:
		return rewriteValueARM_OpARMSUBshiftLLreg(v)
	case OpARMSUBshiftRA:
		return rewriteValueARM_OpARMSUBshiftRA(v)
	case OpARMSUBshiftRAreg:
		return rewriteValueARM_OpARMSUBshiftRAreg(v)
	case OpARMSUBshiftRL:
		return rewriteValueARM_OpARMSUBshiftRL(v)
	case OpARMSUBshiftRLreg:
		return rewriteValueARM_OpARMSUBshiftRLreg(v)
	case OpARMTEQ:
		return rewriteValueARM_OpARMTEQ(v)
	case OpARMTEQconst:
		return rewriteValueARM_OpARMTEQconst(v)
	case OpARMTEQshiftLL:
		return rewriteValueARM_OpARMTEQshiftLL(v)
	case OpARMTEQshiftLLreg:
		return rewriteValueARM_OpARMTEQshiftLLreg(v)
	case OpARMTEQshiftRA:
		return rewriteValueARM_OpARMTEQshiftRA(v)
	case OpARMTEQshiftRAreg:
		return rewriteValueARM_OpARMTEQshiftRAreg(v)
	case OpARMTEQshiftRL:
		return rewriteValueARM_OpARMTEQshiftRL(v)
	case OpARMTEQshiftRLreg:
		return rewriteValueARM_OpARMTEQshiftRLreg(v)
	case OpARMTST:
		return rewriteValueARM_OpARMTST(v)
	case OpARMTSTconst:
		return rewriteValueARM_OpARMTSTconst(v)
	case OpARMTSTshiftLL:
		return rewriteValueARM_OpARMTSTshiftLL(v)
	case OpARMTSTshiftLLreg:
		return rewriteValueARM_OpARMTSTshiftLLreg(v)
	case OpARMTSTshiftRA:
		return rewriteValueARM_OpARMTSTshiftRA(v)
	case OpARMTSTshiftRAreg:
		return rewriteValueARM_OpARMTSTshiftRAreg(v)
	case OpARMTSTshiftRL:
		return rewriteValueARM_OpARMTSTshiftRL(v)
	case OpARMTSTshiftRLreg:
		return rewriteValueARM_OpARMTSTshiftRLreg(v)
	case OpARMXOR:
		return rewriteValueARM_OpARMXOR(v)
	case OpARMXORconst:
		return rewriteValueARM_OpARMXORconst(v)
	case OpARMXORshiftLL:
		return rewriteValueARM_OpARMXORshiftLL(v)
	case OpARMXORshiftLLreg:
		return rewriteValueARM_OpARMXORshiftLLreg(v)
	case OpARMXORshiftRA:
		return rewriteValueARM_OpARMXORshiftRA(v)
	case OpARMXORshiftRAreg:
		return rewriteValueARM_OpARMXORshiftRAreg(v)
	case OpARMXORshiftRL:
		return rewriteValueARM_OpARMXORshiftRL(v)
	case OpARMXORshiftRLreg:
		return rewriteValueARM_OpARMXORshiftRLreg(v)
	case OpARMXORshiftRR:
		return rewriteValueARM_OpARMXORshiftRR(v)
	case OpAbs:
		return rewriteValueARM_OpAbs(v)
	case OpAdd16:
		return rewriteValueARM_OpAdd16(v)
	case OpAdd32:
		return rewriteValueARM_OpAdd32(v)
	case OpAdd32F:
		return rewriteValueARM_OpAdd32F(v)
	case OpAdd32carry:
		return rewriteValueARM_OpAdd32carry(v)
	case OpAdd32withcarry:
		return rewriteValueARM_OpAdd32withcarry(v)
	case OpAdd64F:
		return rewriteValueARM_OpAdd64F(v)
	case OpAdd8:
		return rewriteValueARM_OpAdd8(v)
	case OpAddPtr:
		return rewriteValueARM_OpAddPtr(v)
	case OpAddr:
		return rewriteValueARM_OpAddr(v)
	case OpAnd16:
		return rewriteValueARM_OpAnd16(v)
	case OpAnd32:
		return rewriteValueARM_OpAnd32(v)
	case OpAnd8:
		return rewriteValueARM_OpAnd8(v)
	case OpAndB:
		return rewriteValueARM_OpAndB(v)
	case OpAvg32u:
		return rewriteValueARM_OpAvg32u(v)
	case OpBitLen32:
		return rewriteValueARM_OpBitLen32(v)
	case OpBswap32:
		return rewriteValueARM_OpBswap32(v)
	case OpClosureCall:
		return rewriteValueARM_OpClosureCall(v)
	case OpCom16:
		return rewriteValueARM_OpCom16(v)
	case OpCom32:
		return rewriteValueARM_OpCom32(v)
	case OpCom8:
		return rewriteValueARM_OpCom8(v)
	case OpConst16:
		return rewriteValueARM_OpConst16(v)
	case OpConst32:
		return rewriteValueARM_OpConst32(v)
	case OpConst32F:
		return rewriteValueARM_OpConst32F(v)
	case OpConst64F:
		return rewriteValueARM_OpConst64F(v)
	case OpConst8:
		return rewriteValueARM_OpConst8(v)
	case OpConstBool:
		return rewriteValueARM_OpConstBool(v)
	case OpConstNil:
		return rewriteValueARM_OpConstNil(v)
	case OpCtz16:
		return rewriteValueARM_OpCtz16(v)
	case OpCtz16NonZero:
		return rewriteValueARM_OpCtz16NonZero(v)
	case OpCtz32:
		return rewriteValueARM_OpCtz32(v)
	case OpCtz32NonZero:
		return rewriteValueARM_OpCtz32NonZero(v)
	case OpCtz8:
		return rewriteValueARM_OpCtz8(v)
	case OpCtz8NonZero:
		return rewriteValueARM_OpCtz8NonZero(v)
	case OpCvt32Fto32:
		return rewriteValueARM_OpCvt32Fto32(v)
	case OpCvt32Fto32U:
		return rewriteValueARM_OpCvt32Fto32U(v)
	case OpCvt32Fto64F:
		return rewriteValueARM_OpCvt32Fto64F(v)
	case OpCvt32Uto32F:
		return rewriteValueARM_OpCvt32Uto32F(v)
	case OpCvt32Uto64F:
		return rewriteValueARM_OpCvt32Uto64F(v)
	case OpCvt32to32F:
		return rewriteValueARM_OpCvt32to32F(v)
	case OpCvt32to64F:
		return rewriteValueARM_OpCvt32to64F(v)
	case OpCvt64Fto32:
		return rewriteValueARM_OpCvt64Fto32(v)
	case OpCvt64Fto32F:
		return rewriteValueARM_OpCvt64Fto32F(v)
	case OpCvt64Fto32U:
		return rewriteValueARM_OpCvt64Fto32U(v)
	case OpDiv16:
		return rewriteValueARM_OpDiv16(v)
	case OpDiv16u:
		return rewriteValueARM_OpDiv16u(v)
	case OpDiv32:
		return rewriteValueARM_OpDiv32(v)
	case OpDiv32F:
		return rewriteValueARM_OpDiv32F(v)
	case OpDiv32u:
		return rewriteValueARM_OpDiv32u(v)
	case OpDiv64F:
		return rewriteValueARM_OpDiv64F(v)
	case OpDiv8:
		return rewriteValueARM_OpDiv8(v)
	case OpDiv8u:
		return rewriteValueARM_OpDiv8u(v)
	case OpEq16:
		return rewriteValueARM_OpEq16(v)
	case OpEq32:
		return rewriteValueARM_OpEq32(v)
	case OpEq32F:
		return rewriteValueARM_OpEq32F(v)
	case OpEq64F:
		return rewriteValueARM_OpEq64F(v)
	case OpEq8:
		return rewriteValueARM_OpEq8(v)
	case OpEqB:
		return rewriteValueARM_OpEqB(v)
	case OpEqPtr:
		return rewriteValueARM_OpEqPtr(v)
	case OpFMA:
		return rewriteValueARM_OpFMA(v)
	case OpGeq16:
		return rewriteValueARM_OpGeq16(v)
	case OpGeq16U:
		return rewriteValueARM_OpGeq16U(v)
	case OpGeq32:
		return rewriteValueARM_OpGeq32(v)
	case OpGeq32F:
		return rewriteValueARM_OpGeq32F(v)
	case OpGeq32U:
		return rewriteValueARM_OpGeq32U(v)
	case OpGeq64F:
		return rewriteValueARM_OpGeq64F(v)
	case OpGeq8:
		return rewriteValueARM_OpGeq8(v)
	case OpGeq8U:
		return rewriteValueARM_OpGeq8U(v)
	case OpGetCallerPC:
		return rewriteValueARM_OpGetCallerPC(v)
	case OpGetCallerSP:
		return rewriteValueARM_OpGetCallerSP(v)
	case OpGetClosurePtr:
		return rewriteValueARM_OpGetClosurePtr(v)
	case OpGreater16:
		return rewriteValueARM_OpGreater16(v)
	case OpGreater16U:
		return rewriteValueARM_OpGreater16U(v)
	case OpGreater32:
		return rewriteValueARM_OpGreater32(v)
	case OpGreater32F:
		return rewriteValueARM_OpGreater32F(v)
	case OpGreater32U:
		return rewriteValueARM_OpGreater32U(v)
	case OpGreater64F:
		return rewriteValueARM_OpGreater64F(v)
	case OpGreater8:
		return rewriteValueARM_OpGreater8(v)
	case OpGreater8U:
		return rewriteValueARM_OpGreater8U(v)
	case OpHmul32:
		return rewriteValueARM_OpHmul32(v)
	case OpHmul32u:
		return rewriteValueARM_OpHmul32u(v)
	case OpInterCall:
		return rewriteValueARM_OpInterCall(v)
	case OpIsInBounds:
		return rewriteValueARM_OpIsInBounds(v)
	case OpIsNonNil:
		return rewriteValueARM_OpIsNonNil(v)
	case OpIsSliceInBounds:
		return rewriteValueARM_OpIsSliceInBounds(v)
	case OpLeq16:
		return rewriteValueARM_OpLeq16(v)
	case OpLeq16U:
		return rewriteValueARM_OpLeq16U(v)
	case OpLeq32:
		return rewriteValueARM_OpLeq32(v)
	case OpLeq32F:
		return rewriteValueARM_OpLeq32F(v)
	case OpLeq32U:
		return rewriteValueARM_OpLeq32U(v)
	case OpLeq64F:
		return rewriteValueARM_OpLeq64F(v)
	case OpLeq8:
		return rewriteValueARM_OpLeq8(v)
	case OpLeq8U:
		return rewriteValueARM_OpLeq8U(v)
	case OpLess16:
		return rewriteValueARM_OpLess16(v)
	case OpLess16U:
		return rewriteValueARM_OpLess16U(v)
	case OpLess32:
		return rewriteValueARM_OpLess32(v)
	case OpLess32F:
		return rewriteValueARM_OpLess32F(v)
	case OpLess32U:
		return rewriteValueARM_OpLess32U(v)
	case OpLess64F:
		return rewriteValueARM_OpLess64F(v)
	case OpLess8:
		return rewriteValueARM_OpLess8(v)
	case OpLess8U:
		return rewriteValueARM_OpLess8U(v)
	case OpLoad:
		return rewriteValueARM_OpLoad(v)
	case OpLocalAddr:
		return rewriteValueARM_OpLocalAddr(v)
	case OpLsh16x16:
		return rewriteValueARM_OpLsh16x16(v)
	case OpLsh16x32:
		return rewriteValueARM_OpLsh16x32(v)
	case OpLsh16x64:
		return rewriteValueARM_OpLsh16x64(v)
	case OpLsh16x8:
		return rewriteValueARM_OpLsh16x8(v)
	case OpLsh32x16:
		return rewriteValueARM_OpLsh32x16(v)
	case OpLsh32x32:
		return rewriteValueARM_OpLsh32x32(v)
	case OpLsh32x64:
		return rewriteValueARM_OpLsh32x64(v)
	case OpLsh32x8:
		return rewriteValueARM_OpLsh32x8(v)
	case OpLsh8x16:
		return rewriteValueARM_OpLsh8x16(v)
	case OpLsh8x32:
		return rewriteValueARM_OpLsh8x32(v)
	case OpLsh8x64:
		return rewriteValueARM_OpLsh8x64(v)
	case OpLsh8x8:
		return rewriteValueARM_OpLsh8x8(v)
	case OpMod16:
		return rewriteValueARM_OpMod16(v)
	case OpMod16u:
		return rewriteValueARM_OpMod16u(v)
	case OpMod32:
		return rewriteValueARM_OpMod32(v)
	case OpMod32u:
		return rewriteValueARM_OpMod32u(v)
	case OpMod8:
		return rewriteValueARM_OpMod8(v)
	case OpMod8u:
		return rewriteValueARM_OpMod8u(v)
	case OpMove:
		return rewriteValueARM_OpMove(v)
	case OpMul16:
		return rewriteValueARM_OpMul16(v)
	case OpMul32:
		return rewriteValueARM_OpMul32(v)
	case OpMul32F:
		return rewriteValueARM_OpMul32F(v)
	case OpMul32uhilo:
		return rewriteValueARM_OpMul32uhilo(v)
	case OpMul64F:
		return rewriteValueARM_OpMul64F(v)
	case OpMul8:
		return rewriteValueARM_OpMul8(v)
	case OpNeg16:
		return rewriteValueARM_OpNeg16(v)
	case OpNeg32:
		return rewriteValueARM_OpNeg32(v)
	case OpNeg32F:
		return rewriteValueARM_OpNeg32F(v)
	case OpNeg64F:
		return rewriteValueARM_OpNeg64F(v)
	case OpNeg8:
		return rewriteValueARM_OpNeg8(v)
	case OpNeq16:
		return rewriteValueARM_OpNeq16(v)
	case OpNeq32:
		return rewriteValueARM_OpNeq32(v)
	case OpNeq32F:
		return rewriteValueARM_OpNeq32F(v)
	case OpNeq64F:
		return rewriteValueARM_OpNeq64F(v)
	case OpNeq8:
		return rewriteValueARM_OpNeq8(v)
	case OpNeqB:
		return rewriteValueARM_OpNeqB(v)
	case OpNeqPtr:
		return rewriteValueARM_OpNeqPtr(v)
	case OpNilCheck:
		return rewriteValueARM_OpNilCheck(v)
	case OpNot:
		return rewriteValueARM_OpNot(v)
	case OpOffPtr:
		return rewriteValueARM_OpOffPtr(v)
	case OpOr16:
		return rewriteValueARM_OpOr16(v)
	case OpOr32:
		return rewriteValueARM_OpOr32(v)
	case OpOr8:
		return rewriteValueARM_OpOr8(v)
	case OpOrB:
		return rewriteValueARM_OpOrB(v)
	case OpPanicBounds:
		return rewriteValueARM_OpPanicBounds(v)
	case OpPanicExtend:
		return rewriteValueARM_OpPanicExtend(v)
	case OpRotateLeft16:
		return rewriteValueARM_OpRotateLeft16(v)
	case OpRotateLeft32:
		return rewriteValueARM_OpRotateLeft32(v)
	case OpRotateLeft8:
		return rewriteValueARM_OpRotateLeft8(v)
	case OpRound32F:
		return rewriteValueARM_OpRound32F(v)
	case OpRound64F:
		return rewriteValueARM_OpRound64F(v)
	case OpRsh16Ux16:
		return rewriteValueARM_OpRsh16Ux16(v)
	case OpRsh16Ux32:
		return rewriteValueARM_OpRsh16Ux32(v)
	case OpRsh16Ux64:
		return rewriteValueARM_OpRsh16Ux64(v)
	case OpRsh16Ux8:
		return rewriteValueARM_OpRsh16Ux8(v)
	case OpRsh16x16:
		return rewriteValueARM_OpRsh16x16(v)
	case OpRsh16x32:
		return rewriteValueARM_OpRsh16x32(v)
	case OpRsh16x64:
		return rewriteValueARM_OpRsh16x64(v)
	case OpRsh16x8:
		return rewriteValueARM_OpRsh16x8(v)
	case OpRsh32Ux16:
		return rewriteValueARM_OpRsh32Ux16(v)
	case OpRsh32Ux32:
		return rewriteValueARM_OpRsh32Ux32(v)
	case OpRsh32Ux64:
		return rewriteValueARM_OpRsh32Ux64(v)
	case OpRsh32Ux8:
		return rewriteValueARM_OpRsh32Ux8(v)
	case OpRsh32x16:
		return rewriteValueARM_OpRsh32x16(v)
	case OpRsh32x32:
		return rewriteValueARM_OpRsh32x32(v)
	case OpRsh32x64:
		return rewriteValueARM_OpRsh32x64(v)
	case OpRsh32x8:
		return rewriteValueARM_OpRsh32x8(v)
	case OpRsh8Ux16:
		return rewriteValueARM_OpRsh8Ux16(v)
	case OpRsh8Ux32:
		return rewriteValueARM_OpRsh8Ux32(v)
	case OpRsh8Ux64:
		return rewriteValueARM_OpRsh8Ux64(v)
	case OpRsh8Ux8:
		return rewriteValueARM_OpRsh8Ux8(v)
	case OpRsh8x16:
		return rewriteValueARM_OpRsh8x16(v)
	case OpRsh8x32:
		return rewriteValueARM_OpRsh8x32(v)
	case OpRsh8x64:
		return rewriteValueARM_OpRsh8x64(v)
	case OpRsh8x8:
		return rewriteValueARM_OpRsh8x8(v)
	case OpSelect0:
		return rewriteValueARM_OpSelect0(v)
	case OpSelect1:
		return rewriteValueARM_OpSelect1(v)
	case OpSignExt16to32:
		return rewriteValueARM_OpSignExt16to32(v)
	case OpSignExt8to16:
		return rewriteValueARM_OpSignExt8to16(v)
	case OpSignExt8to32:
		return rewriteValueARM_OpSignExt8to32(v)
	case OpSignmask:
		return rewriteValueARM_OpSignmask(v)
	case OpSlicemask:
		return rewriteValueARM_OpSlicemask(v)
	case OpSqrt:
		return rewriteValueARM_OpSqrt(v)
	case OpStaticCall:
		return rewriteValueARM_OpStaticCall(v)
	case OpStore:
		return rewriteValueARM_OpStore(v)
	case OpSub16:
		return rewriteValueARM_OpSub16(v)
	case OpSub32:
		return rewriteValueARM_OpSub32(v)
	case OpSub32F:
		return rewriteValueARM_OpSub32F(v)
	case OpSub32carry:
		return rewriteValueARM_OpSub32carry(v)
	case OpSub32withcarry:
		return rewriteValueARM_OpSub32withcarry(v)
	case OpSub64F:
		return rewriteValueARM_OpSub64F(v)
	case OpSub8:
		return rewriteValueARM_OpSub8(v)
	case OpSubPtr:
		return rewriteValueARM_OpSubPtr(v)
	case OpTrunc16to8:
		return rewriteValueARM_OpTrunc16to8(v)
	case OpTrunc32to16:
		return rewriteValueARM_OpTrunc32to16(v)
	case OpTrunc32to8:
		return rewriteValueARM_OpTrunc32to8(v)
	case OpWB:
		return rewriteValueARM_OpWB(v)
	case OpXor16:
		return rewriteValueARM_OpXor16(v)
	case OpXor32:
		return rewriteValueARM_OpXor32(v)
	case OpXor8:
		return rewriteValueARM_OpXor8(v)
	case OpZero:
		return rewriteValueARM_OpZero(v)
	case OpZeroExt16to32:
		return rewriteValueARM_OpZeroExt16to32(v)
	case OpZeroExt8to16:
		return rewriteValueARM_OpZeroExt8to16(v)
	case OpZeroExt8to32:
		return rewriteValueARM_OpZeroExt8to32(v)
	case OpZeromask:
		return rewriteValueARM_OpZeromask(v)
	}
	return false
}
func rewriteValueARM_OpARMADC(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ADC (MOVWconst [c]) x flags)
	// result: (ADCconst [c] x flags)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpARMMOVWconst {
				continue
			}
			c := v_0.AuxInt
			x := v_1
			flags := v_2
			v.reset(OpARMADCconst)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(flags)
			return true
		}
		break
	}
	// match: (ADC x (SLLconst [c] y) flags)
	// result: (ADCshiftLL x y [c] flags)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSLLconst {
				continue
			}
			c := v_1.AuxInt
			y := v_1.Args[0]
			flags := v_2
			v.reset(OpARMADCshiftLL)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(flags)
			return true
		}
		break
	}
	// match: (ADC x (SRLconst [c] y) flags)
	// result: (ADCshiftRL x y [c] flags)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRLconst {
				continue
			}
			c := v_1.AuxInt
			y := v_1.Args[0]
			flags := v_2
			v.reset(OpARMADCshiftRL)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(flags)
			return true
		}
		break
	}
	// match: (ADC x (SRAconst [c] y) flags)
	// result: (ADCshiftRA x y [c] flags)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRAconst {
				continue
			}
			c := v_1.AuxInt
			y := v_1.Args[0]
			flags := v_2
			v.reset(OpARMADCshiftRA)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(flags)
			return true
		}
		break
	}
	// match: (ADC x (SLL y z) flags)
	// result: (ADCshiftLLreg x y z flags)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSLL {
				continue
			}
			z := v_1.Args[1]
			y := v_1.Args[0]
			flags := v_2
			v.reset(OpARMADCshiftLLreg)
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(z)
			v.AddArg(flags)
			return true
		}
		break
	}
	// match: (ADC x (SRL y z) flags)
	// result: (ADCshiftRLreg x y z flags)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRL {
				continue
			}
			z := v_1.Args[1]
			y := v_1.Args[0]
			flags := v_2
			v.reset(OpARMADCshiftRLreg)
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(z)
			v.AddArg(flags)
			return true
		}
		break
	}
	// match: (ADC x (SRA y z) flags)
	// result: (ADCshiftRAreg x y z flags)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRA {
				continue
			}
			z := v_1.Args[1]
			y := v_1.Args[0]
			flags := v_2
			v.reset(OpARMADCshiftRAreg)
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(z)
			v.AddArg(flags)
			return true
		}
		break
	}
	return false
}
func rewriteValueARM_OpARMADCconst(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ADCconst [c] (ADDconst [d] x) flags)
	// result: (ADCconst [int64(int32(c+d))] x flags)
	for {
		c := v.AuxInt
		if v_0.Op != OpARMADDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		flags := v_1
		v.reset(OpARMADCconst)
		v.AuxInt = int64(int32(c + d))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	// match: (ADCconst [c] (SUBconst [d] x) flags)
	// result: (ADCconst [int64(int32(c-d))] x flags)
	for {
		c := v.AuxInt
		if v_0.Op != OpARMSUBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		flags := v_1
		v.reset(OpARMADCconst)
		v.AuxInt = int64(int32(c - d))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADCshiftLL(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ADCshiftLL (MOVWconst [c]) x [d] flags)
	// result: (ADCconst [c] (SLLconst <x.Type> x [d]) flags)
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		flags := v_2
		v.reset(OpARMADCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}
	// match: (ADCshiftLL x (MOVWconst [c]) [d] flags)
	// result: (ADCconst x [int64(int32(uint32(c)<<uint64(d)))] flags)
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		flags := v_2
		v.reset(OpARMADCconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADCshiftLLreg(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ADCshiftLLreg (MOVWconst [c]) x y flags)
	// result: (ADCconst [c] (SLL <x.Type> x y) flags)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		flags := v_3
		v.reset(OpARMADCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}
	// match: (ADCshiftLLreg x y (MOVWconst [c]) flags)
	// result: (ADCshiftLL x y [c] flags)
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		flags := v_3
		v.reset(OpARMADCshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADCshiftRA(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ADCshiftRA (MOVWconst [c]) x [d] flags)
	// result: (ADCconst [c] (SRAconst <x.Type> x [d]) flags)
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		flags := v_2
		v.reset(OpARMADCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}
	// match: (ADCshiftRA x (MOVWconst [c]) [d] flags)
	// result: (ADCconst x [int64(int32(c)>>uint64(d))] flags)
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		flags := v_2
		v.reset(OpARMADCconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADCshiftRAreg(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ADCshiftRAreg (MOVWconst [c]) x y flags)
	// result: (ADCconst [c] (SRA <x.Type> x y) flags)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		flags := v_3
		v.reset(OpARMADCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRA, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}
	// match: (ADCshiftRAreg x y (MOVWconst [c]) flags)
	// result: (ADCshiftRA x y [c] flags)
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		flags := v_3
		v.reset(OpARMADCshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADCshiftRL(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ADCshiftRL (MOVWconst [c]) x [d] flags)
	// result: (ADCconst [c] (SRLconst <x.Type> x [d]) flags)
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		flags := v_2
		v.reset(OpARMADCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}
	// match: (ADCshiftRL x (MOVWconst [c]) [d] flags)
	// result: (ADCconst x [int64(int32(uint32(c)>>uint64(d)))] flags)
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		flags := v_2
		v.reset(OpARMADCconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADCshiftRLreg(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ADCshiftRLreg (MOVWconst [c]) x y flags)
	// result: (ADCconst [c] (SRL <x.Type> x y) flags)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		flags := v_3
		v.reset(OpARMADCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}
	// match: (ADCshiftRLreg x y (MOVWconst [c]) flags)
	// result: (ADCshiftRL x y [c] flags)
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		flags := v_3
		v.reset(OpARMADCshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADD(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ADD x (MOVWconst [c]))
	// result: (ADDconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMMOVWconst {
				continue
			}
			c := v_1.AuxInt
			v.reset(OpARMADDconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (ADD x (SLLconst [c] y))
	// result: (ADDshiftLL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSLLconst {
				continue
			}
			c := v_1.AuxInt
			y := v_1.Args[0]
			v.reset(OpARMADDshiftLL)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (ADD x (SRLconst [c] y))
	// result: (ADDshiftRL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRLconst {
				continue
			}
			c := v_1.AuxInt
			y := v_1.Args[0]
			v.reset(OpARMADDshiftRL)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (ADD x (SRAconst [c] y))
	// result: (ADDshiftRA x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRAconst {
				continue
			}
			c := v_1.AuxInt
			y := v_1.Args[0]
			v.reset(OpARMADDshiftRA)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (ADD x (SLL y z))
	// result: (ADDshiftLLreg x y z)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSLL {
				continue
			}
			z := v_1.Args[1]
			y := v_1.Args[0]
			v.reset(OpARMADDshiftLLreg)
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(z)
			return true
		}
		break
	}
	// match: (ADD x (SRL y z))
	// result: (ADDshiftRLreg x y z)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRL {
				continue
			}
			z := v_1.Args[1]
			y := v_1.Args[0]
			v.reset(OpARMADDshiftRLreg)
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(z)
			return true
		}
		break
	}
	// match: (ADD x (SRA y z))
	// result: (ADDshiftRAreg x y z)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRA {
				continue
			}
			z := v_1.Args[1]
			y := v_1.Args[0]
			v.reset(OpARMADDshiftRAreg)
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(z)
			return true
		}
		break
	}
	// match: (ADD x (RSBconst [0] y))
	// result: (SUB x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMRSBconst || v_1.AuxInt != 0 {
				continue
			}
			y := v_1.Args[0]
			v.reset(OpARMSUB)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (ADD <t> (RSBconst [c] x) (RSBconst [d] y))
	// result: (RSBconst [c+d] (ADD <t> x y))
	for {
		t := v.Type
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpARMRSBconst {
				continue
			}
			c := v_0.AuxInt
			x := v_0.Args[0]
			if v_1.Op != OpARMRSBconst {
				continue
			}
			d := v_1.AuxInt
			y := v_1.Args[0]
			v.reset(OpARMRSBconst)
			v.AuxInt = c + d
			v0 := b.NewValue0(v.Pos, OpARMADD, t)
			v0.AddArg(x)
			v0.AddArg(y)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (ADD (MUL x y) a)
	// result: (MULA x y a)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpARMMUL {
				continue
			}
			y := v_0.Args[1]
			x := v_0.Args[0]
			a := v_1
			v.reset(OpARMMULA)
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(a)
			return true
		}
		break
	}
	return false
}
func rewriteValueARM_OpARMADDD(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ADDD a (MULD x y))
	// cond: a.Uses == 1 && objabi.GOARM >= 6
	// result: (MULAD a x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			a := v_0
			if v_1.Op != OpARMMULD {
				continue
			}
			y := v_1.Args[1]
			x := v_1.Args[0]
			if !(a.Uses == 1 && objabi.GOARM >= 6) {
				continue
			}
			v.reset(OpARMMULAD)
			v.AddArg(a)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (ADDD a (NMULD x y))
	// cond: a.Uses == 1 && objabi.GOARM >= 6
	// result: (MULSD a x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			a := v_0
			if v_1.Op != OpARMNMULD {
				continue
			}
			y := v_1.Args[1]
			x := v_1.Args[0]
			if !(a.Uses == 1 && objabi.GOARM >= 6) {
				continue
			}
			v.reset(OpARMMULSD)
			v.AddArg(a)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	return false
}
func rewriteValueARM_OpARMADDF(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ADDF a (MULF x y))
	// cond: a.Uses == 1 && objabi.GOARM >= 6
	// result: (MULAF a x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			a := v_0
			if v_1.Op != OpARMMULF {
				continue
			}
			y := v_1.Args[1]
			x := v_1.Args[0]
			if !(a.Uses == 1 && objabi.GOARM >= 6) {
				continue
			}
			v.reset(OpARMMULAF)
			v.AddArg(a)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (ADDF a (NMULF x y))
	// cond: a.Uses == 1 && objabi.GOARM >= 6
	// result: (MULSF a x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			a := v_0
			if v_1.Op != OpARMNMULF {
				continue
			}
			y := v_1.Args[1]
			x := v_1.Args[0]
			if !(a.Uses == 1 && objabi.GOARM >= 6) {
				continue
			}
			v.reset(OpARMMULSF)
			v.AddArg(a)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	return false
}
func rewriteValueARM_OpARMADDS(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ADDS x (MOVWconst [c]))
	// result: (ADDSconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMMOVWconst {
				continue
			}
			c := v_1.AuxInt
			v.reset(OpARMADDSconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (ADDS x (SLLconst [c] y))
	// result: (ADDSshiftLL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSLLconst {
				continue
			}
			c := v_1.AuxInt
			y := v_1.Args[0]
			v.reset(OpARMADDSshiftLL)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (ADDS x (SRLconst [c] y))
	// result: (ADDSshiftRL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRLconst {
				continue
			}
			c := v_1.AuxInt
			y := v_1.Args[0]
			v.reset(OpARMADDSshiftRL)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (ADDS x (SRAconst [c] y))
	// result: (ADDSshiftRA x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRAconst {
				continue
			}
			c := v_1.AuxInt
			y := v_1.Args[0]
			v.reset(OpARMADDSshiftRA)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (ADDS x (SLL y z))
	// result: (ADDSshiftLLreg x y z)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSLL {
				continue
			}
			z := v_1.Args[1]
			y := v_1.Args[0]
			v.reset(OpARMADDSshiftLLreg)
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(z)
			return true
		}
		break
	}
	// match: (ADDS x (SRL y z))
	// result: (ADDSshiftRLreg x y z)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRL {
				continue
			}
			z := v_1.Args[1]
			y := v_1.Args[0]
			v.reset(OpARMADDSshiftRLreg)
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(z)
			return true
		}
		break
	}
	// match: (ADDS x (SRA y z))
	// result: (ADDSshiftRAreg x y z)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRA {
				continue
			}
			z := v_1.Args[1]
			y := v_1.Args[0]
			v.reset(OpARMADDSshiftRAreg)
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(z)
			return true
		}
		break
	}
	return false
}
func rewriteValueARM_OpARMADDSshiftLL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ADDSshiftLL (MOVWconst [c]) x [d])
	// result: (ADDSconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMADDSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ADDSshiftLL x (MOVWconst [c]) [d])
	// result: (ADDSconst x [int64(int32(uint32(c)<<uint64(d)))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMADDSconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADDSshiftLLreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ADDSshiftLLreg (MOVWconst [c]) x y)
	// result: (ADDSconst [c] (SLL <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMADDSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (ADDSshiftLLreg x y (MOVWconst [c]))
	// result: (ADDSshiftLL x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMADDSshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADDSshiftRA(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ADDSshiftRA (MOVWconst [c]) x [d])
	// result: (ADDSconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMADDSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ADDSshiftRA x (MOVWconst [c]) [d])
	// result: (ADDSconst x [int64(int32(c)>>uint64(d))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMADDSconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADDSshiftRAreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ADDSshiftRAreg (MOVWconst [c]) x y)
	// result: (ADDSconst [c] (SRA <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMADDSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRA, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (ADDSshiftRAreg x y (MOVWconst [c]))
	// result: (ADDSshiftRA x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMADDSshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADDSshiftRL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ADDSshiftRL (MOVWconst [c]) x [d])
	// result: (ADDSconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMADDSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ADDSshiftRL x (MOVWconst [c]) [d])
	// result: (ADDSconst x [int64(int32(uint32(c)>>uint64(d)))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMADDSconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADDSshiftRLreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ADDSshiftRLreg (MOVWconst [c]) x y)
	// result: (ADDSconst [c] (SRL <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMADDSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (ADDSshiftRLreg x y (MOVWconst [c]))
	// result: (ADDSshiftRL x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMADDSshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADDconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ADDconst [off1] (MOVWaddr [off2] {sym} ptr))
	// result: (MOVWaddr [off1+off2] {sym} ptr)
	for {
		off1 := v.AuxInt
		if v_0.Op != OpARMMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym := v_0.Aux
		ptr := v_0.Args[0]
		v.reset(OpARMMOVWaddr)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		return true
	}
	// match: (ADDconst [0] x)
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
	// match: (ADDconst [c] x)
	// cond: !isARMImmRot(uint32(c)) && isARMImmRot(uint32(-c))
	// result: (SUBconst [int64(int32(-c))] x)
	for {
		c := v.AuxInt
		x := v_0
		if !(!isARMImmRot(uint32(c)) && isARMImmRot(uint32(-c))) {
			break
		}
		v.reset(OpARMSUBconst)
		v.AuxInt = int64(int32(-c))
		v.AddArg(x)
		return true
	}
	// match: (ADDconst [c] x)
	// cond: objabi.GOARM==7 && !isARMImmRot(uint32(c)) && uint32(c)>0xffff && uint32(-c)<=0xffff
	// result: (SUBconst [int64(int32(-c))] x)
	for {
		c := v.AuxInt
		x := v_0
		if !(objabi.GOARM == 7 && !isARMImmRot(uint32(c)) && uint32(c) > 0xffff && uint32(-c) <= 0xffff) {
			break
		}
		v.reset(OpARMSUBconst)
		v.AuxInt = int64(int32(-c))
		v.AddArg(x)
		return true
	}
	// match: (ADDconst [c] (MOVWconst [d]))
	// result: (MOVWconst [int64(int32(c+d))])
	for {
		c := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = int64(int32(c + d))
		return true
	}
	// match: (ADDconst [c] (ADDconst [d] x))
	// result: (ADDconst [int64(int32(c+d))] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpARMADDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMADDconst)
		v.AuxInt = int64(int32(c + d))
		v.AddArg(x)
		return true
	}
	// match: (ADDconst [c] (SUBconst [d] x))
	// result: (ADDconst [int64(int32(c-d))] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpARMSUBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMADDconst)
		v.AuxInt = int64(int32(c - d))
		v.AddArg(x)
		return true
	}
	// match: (ADDconst [c] (RSBconst [d] x))
	// result: (RSBconst [int64(int32(c+d))] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpARMRSBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMRSBconst)
		v.AuxInt = int64(int32(c + d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADDshiftLL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ADDshiftLL (MOVWconst [c]) x [d])
	// result: (ADDconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMADDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ADDshiftLL x (MOVWconst [c]) [d])
	// result: (ADDconst x [int64(int32(uint32(c)<<uint64(d)))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMADDconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: (ADDshiftLL [c] (SRLconst x [32-c]) x)
	// result: (SRRconst [32-c] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpARMSRLconst || v_0.AuxInt != 32-c {
			break
		}
		x := v_0.Args[0]
		if x != v_1 {
			break
		}
		v.reset(OpARMSRRconst)
		v.AuxInt = 32 - c
		v.AddArg(x)
		return true
	}
	// match: (ADDshiftLL <typ.UInt16> [8] (BFXU <typ.UInt16> [armBFAuxInt(8, 8)] x) x)
	// result: (REV16 x)
	for {
		if v.Type != typ.UInt16 || v.AuxInt != 8 || v_0.Op != OpARMBFXU || v_0.Type != typ.UInt16 || v_0.AuxInt != armBFAuxInt(8, 8) {
			break
		}
		x := v_0.Args[0]
		if x != v_1 {
			break
		}
		v.reset(OpARMREV16)
		v.AddArg(x)
		return true
	}
	// match: (ADDshiftLL <typ.UInt16> [8] (SRLconst <typ.UInt16> [24] (SLLconst [16] x)) x)
	// cond: objabi.GOARM>=6
	// result: (REV16 x)
	for {
		if v.Type != typ.UInt16 || v.AuxInt != 8 || v_0.Op != OpARMSRLconst || v_0.Type != typ.UInt16 || v_0.AuxInt != 24 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpARMSLLconst || v_0_0.AuxInt != 16 {
			break
		}
		x := v_0_0.Args[0]
		if x != v_1 || !(objabi.GOARM >= 6) {
			break
		}
		v.reset(OpARMREV16)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADDshiftLLreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ADDshiftLLreg (MOVWconst [c]) x y)
	// result: (ADDconst [c] (SLL <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMADDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (ADDshiftLLreg x y (MOVWconst [c]))
	// result: (ADDshiftLL x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMADDshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADDshiftRA(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ADDshiftRA (MOVWconst [c]) x [d])
	// result: (ADDconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMADDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ADDshiftRA x (MOVWconst [c]) [d])
	// result: (ADDconst x [int64(int32(c)>>uint64(d))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMADDconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADDshiftRAreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ADDshiftRAreg (MOVWconst [c]) x y)
	// result: (ADDconst [c] (SRA <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMADDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRA, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (ADDshiftRAreg x y (MOVWconst [c]))
	// result: (ADDshiftRA x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMADDshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADDshiftRL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ADDshiftRL (MOVWconst [c]) x [d])
	// result: (ADDconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMADDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ADDshiftRL x (MOVWconst [c]) [d])
	// result: (ADDconst x [int64(int32(uint32(c)>>uint64(d)))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMADDconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: (ADDshiftRL [c] (SLLconst x [32-c]) x)
	// result: (SRRconst [ c] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpARMSLLconst || v_0.AuxInt != 32-c {
			break
		}
		x := v_0.Args[0]
		if x != v_1 {
			break
		}
		v.reset(OpARMSRRconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMADDshiftRLreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ADDshiftRLreg (MOVWconst [c]) x y)
	// result: (ADDconst [c] (SRL <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMADDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (ADDshiftRLreg x y (MOVWconst [c]))
	// result: (ADDshiftRL x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMADDshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMAND(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (AND x (MOVWconst [c]))
	// result: (ANDconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMMOVWconst {
				continue
			}
			c := v_1.AuxInt
			v.reset(OpARMANDconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (AND x (SLLconst [c] y))
	// result: (ANDshiftLL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSLLconst {
				continue
			}
			c := v_1.AuxInt
			y := v_1.Args[0]
			v.reset(OpARMANDshiftLL)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (AND x (SRLconst [c] y))
	// result: (ANDshiftRL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRLconst {
				continue
			}
			c := v_1.AuxInt
			y := v_1.Args[0]
			v.reset(OpARMANDshiftRL)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (AND x (SRAconst [c] y))
	// result: (ANDshiftRA x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRAconst {
				continue
			}
			c := v_1.AuxInt
			y := v_1.Args[0]
			v.reset(OpARMANDshiftRA)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (AND x (SLL y z))
	// result: (ANDshiftLLreg x y z)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSLL {
				continue
			}
			z := v_1.Args[1]
			y := v_1.Args[0]
			v.reset(OpARMANDshiftLLreg)
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(z)
			return true
		}
		break
	}
	// match: (AND x (SRL y z))
	// result: (ANDshiftRLreg x y z)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRL {
				continue
			}
			z := v_1.Args[1]
			y := v_1.Args[0]
			v.reset(OpARMANDshiftRLreg)
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(z)
			return true
		}
		break
	}
	// match: (AND x (SRA y z))
	// result: (ANDshiftRAreg x y z)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRA {
				continue
			}
			z := v_1.Args[1]
			y := v_1.Args[0]
			v.reset(OpARMANDshiftRAreg)
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(z)
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
	// match: (AND x (MVN y))
	// result: (BIC x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMMVN {
				continue
			}
			y := v_1.Args[0]
			v.reset(OpARMBIC)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (AND x (MVNshiftLL y [c]))
	// result: (BICshiftLL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMMVNshiftLL {
				continue
			}
			c := v_1.AuxInt
			y := v_1.Args[0]
			v.reset(OpARMBICshiftLL)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (AND x (MVNshiftRL y [c]))
	// result: (BICshiftRL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMMVNshiftRL {
				continue
			}
			c := v_1.AuxInt
			y := v_1.Args[0]
			v.reset(OpARMBICshiftRL)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (AND x (MVNshiftRA y [c]))
	// result: (BICshiftRA x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMMVNshiftRA {
				continue
			}
			c := v_1.AuxInt
			y := v_1.Args[0]
			v.reset(OpARMBICshiftRA)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	return false
}
func rewriteValueARM_OpARMANDconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ANDconst [0] _)
	// result: (MOVWconst [0])
	for {
		if v.AuxInt != 0 {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (ANDconst [c] x)
	// cond: int32(c)==-1
	// result: x
	for {
		c := v.AuxInt
		x := v_0
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (ANDconst [c] x)
	// cond: !isARMImmRot(uint32(c)) && isARMImmRot(^uint32(c))
	// result: (BICconst [int64(int32(^uint32(c)))] x)
	for {
		c := v.AuxInt
		x := v_0
		if !(!isARMImmRot(uint32(c)) && isARMImmRot(^uint32(c))) {
			break
		}
		v.reset(OpARMBICconst)
		v.AuxInt = int64(int32(^uint32(c)))
		v.AddArg(x)
		return true
	}
	// match: (ANDconst [c] x)
	// cond: objabi.GOARM==7 && !isARMImmRot(uint32(c)) && uint32(c)>0xffff && ^uint32(c)<=0xffff
	// result: (BICconst [int64(int32(^uint32(c)))] x)
	for {
		c := v.AuxInt
		x := v_0
		if !(objabi.GOARM == 7 && !isARMImmRot(uint32(c)) && uint32(c) > 0xffff && ^uint32(c) <= 0xffff) {
			break
		}
		v.reset(OpARMBICconst)
		v.AuxInt = int64(int32(^uint32(c)))
		v.AddArg(x)
		return true
	}
	// match: (ANDconst [c] (MOVWconst [d]))
	// result: (MOVWconst [c&d])
	for {
		c := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = c & d
		return true
	}
	// match: (ANDconst [c] (ANDconst [d] x))
	// result: (ANDconst [c&d] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpARMANDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMANDconst)
		v.AuxInt = c & d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMANDshiftLL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ANDshiftLL (MOVWconst [c]) x [d])
	// result: (ANDconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMANDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ANDshiftLL x (MOVWconst [c]) [d])
	// result: (ANDconst x [int64(int32(uint32(c)<<uint64(d)))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMANDconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: (ANDshiftLL x y:(SLLconst x [c]) [d])
	// cond: c==d
	// result: y
	for {
		d := v.AuxInt
		x := v_0
		y := v_1
		if y.Op != OpARMSLLconst {
			break
		}
		c := y.AuxInt
		if x != y.Args[0] || !(c == d) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMANDshiftLLreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ANDshiftLLreg (MOVWconst [c]) x y)
	// result: (ANDconst [c] (SLL <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMANDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (ANDshiftLLreg x y (MOVWconst [c]))
	// result: (ANDshiftLL x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMANDshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMANDshiftRA(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ANDshiftRA (MOVWconst [c]) x [d])
	// result: (ANDconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMANDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ANDshiftRA x (MOVWconst [c]) [d])
	// result: (ANDconst x [int64(int32(c)>>uint64(d))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMANDconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	// match: (ANDshiftRA x y:(SRAconst x [c]) [d])
	// cond: c==d
	// result: y
	for {
		d := v.AuxInt
		x := v_0
		y := v_1
		if y.Op != OpARMSRAconst {
			break
		}
		c := y.AuxInt
		if x != y.Args[0] || !(c == d) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMANDshiftRAreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ANDshiftRAreg (MOVWconst [c]) x y)
	// result: (ANDconst [c] (SRA <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMANDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRA, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (ANDshiftRAreg x y (MOVWconst [c]))
	// result: (ANDshiftRA x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMANDshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMANDshiftRL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ANDshiftRL (MOVWconst [c]) x [d])
	// result: (ANDconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMANDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ANDshiftRL x (MOVWconst [c]) [d])
	// result: (ANDconst x [int64(int32(uint32(c)>>uint64(d)))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMANDconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: (ANDshiftRL x y:(SRLconst x [c]) [d])
	// cond: c==d
	// result: y
	for {
		d := v.AuxInt
		x := v_0
		y := v_1
		if y.Op != OpARMSRLconst {
			break
		}
		c := y.AuxInt
		if x != y.Args[0] || !(c == d) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMANDshiftRLreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ANDshiftRLreg (MOVWconst [c]) x y)
	// result: (ANDconst [c] (SRL <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMANDconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (ANDshiftRLreg x y (MOVWconst [c]))
	// result: (ANDshiftRL x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMANDshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMBFX(v *Value) bool {
	v_0 := v.Args[0]
	// match: (BFX [c] (MOVWconst [d]))
	// result: (MOVWconst [int64(int32(d)<<(32-uint32(c&0xff)-uint32(c>>8))>>(32-uint32(c>>8)))])
	for {
		c := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = int64(int32(d) << (32 - uint32(c&0xff) - uint32(c>>8)) >> (32 - uint32(c>>8)))
		return true
	}
	return false
}
func rewriteValueARM_OpARMBFXU(v *Value) bool {
	v_0 := v.Args[0]
	// match: (BFXU [c] (MOVWconst [d]))
	// result: (MOVWconst [int64(int32(uint32(d)<<(32-uint32(c&0xff)-uint32(c>>8))>>(32-uint32(c>>8))))])
	for {
		c := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = int64(int32(uint32(d) << (32 - uint32(c&0xff) - uint32(c>>8)) >> (32 - uint32(c>>8))))
		return true
	}
	return false
}
func rewriteValueARM_OpARMBIC(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (BIC x (MOVWconst [c]))
	// result: (BICconst [c] x)
	for {
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMBICconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (BIC x (SLLconst [c] y))
	// result: (BICshiftLL x y [c])
	for {
		x := v_0
		if v_1.Op != OpARMSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMBICshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (BIC x (SRLconst [c] y))
	// result: (BICshiftRL x y [c])
	for {
		x := v_0
		if v_1.Op != OpARMSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMBICshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (BIC x (SRAconst [c] y))
	// result: (BICshiftRA x y [c])
	for {
		x := v_0
		if v_1.Op != OpARMSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMBICshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (BIC x (SLL y z))
	// result: (BICshiftLLreg x y z)
	for {
		x := v_0
		if v_1.Op != OpARMSLL {
			break
		}
		z := v_1.Args[1]
		y := v_1.Args[0]
		v.reset(OpARMBICshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	// match: (BIC x (SRL y z))
	// result: (BICshiftRLreg x y z)
	for {
		x := v_0
		if v_1.Op != OpARMSRL {
			break
		}
		z := v_1.Args[1]
		y := v_1.Args[0]
		v.reset(OpARMBICshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	// match: (BIC x (SRA y z))
	// result: (BICshiftRAreg x y z)
	for {
		x := v_0
		if v_1.Op != OpARMSRA {
			break
		}
		z := v_1.Args[1]
		y := v_1.Args[0]
		v.reset(OpARMBICshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	// match: (BIC x x)
	// result: (MOVWconst [0])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpARMBICconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (BICconst [0] x)
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
	// match: (BICconst [c] _)
	// cond: int32(c)==-1
	// result: (MOVWconst [0])
	for {
		c := v.AuxInt
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (BICconst [c] x)
	// cond: !isARMImmRot(uint32(c)) && isARMImmRot(^uint32(c))
	// result: (ANDconst [int64(int32(^uint32(c)))] x)
	for {
		c := v.AuxInt
		x := v_0
		if !(!isARMImmRot(uint32(c)) && isARMImmRot(^uint32(c))) {
			break
		}
		v.reset(OpARMANDconst)
		v.AuxInt = int64(int32(^uint32(c)))
		v.AddArg(x)
		return true
	}
	// match: (BICconst [c] x)
	// cond: objabi.GOARM==7 && !isARMImmRot(uint32(c)) && uint32(c)>0xffff && ^uint32(c)<=0xffff
	// result: (ANDconst [int64(int32(^uint32(c)))] x)
	for {
		c := v.AuxInt
		x := v_0
		if !(objabi.GOARM == 7 && !isARMImmRot(uint32(c)) && uint32(c) > 0xffff && ^uint32(c) <= 0xffff) {
			break
		}
		v.reset(OpARMANDconst)
		v.AuxInt = int64(int32(^uint32(c)))
		v.AddArg(x)
		return true
	}
	// match: (BICconst [c] (MOVWconst [d]))
	// result: (MOVWconst [d&^c])
	for {
		c := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = d &^ c
		return true
	}
	// match: (BICconst [c] (BICconst [d] x))
	// result: (BICconst [int64(int32(c|d))] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpARMBICconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMBICconst)
		v.AuxInt = int64(int32(c | d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMBICshiftLL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (BICshiftLL x (MOVWconst [c]) [d])
	// result: (BICconst x [int64(int32(uint32(c)<<uint64(d)))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMBICconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: (BICshiftLL x (SLLconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMSLLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] || !(c == d) {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpARMBICshiftLLreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (BICshiftLLreg x y (MOVWconst [c]))
	// result: (BICshiftLL x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMBICshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMBICshiftRA(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (BICshiftRA x (MOVWconst [c]) [d])
	// result: (BICconst x [int64(int32(c)>>uint64(d))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMBICconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	// match: (BICshiftRA x (SRAconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMSRAconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] || !(c == d) {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpARMBICshiftRAreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (BICshiftRAreg x y (MOVWconst [c]))
	// result: (BICshiftRA x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMBICshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMBICshiftRL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (BICshiftRL x (MOVWconst [c]) [d])
	// result: (BICconst x [int64(int32(uint32(c)>>uint64(d)))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMBICconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: (BICshiftRL x (SRLconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMSRLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] || !(c == d) {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpARMBICshiftRLreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (BICshiftRLreg x y (MOVWconst [c]))
	// result: (BICshiftRL x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMBICshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMCMN(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (CMN x (MOVWconst [c]))
	// result: (CMNconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMMOVWconst {
				continue
			}
			c := v_1.AuxInt
			v.reset(OpARMCMNconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (CMN x (SLLconst [c] y))
	// result: (CMNshiftLL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSLLconst {
				continue
			}
			c := v_1.AuxInt
			y := v_1.Args[0]
			v.reset(OpARMCMNshiftLL)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (CMN x (SRLconst [c] y))
	// result: (CMNshiftRL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRLconst {
				continue
			}
			c := v_1.AuxInt
			y := v_1.Args[0]
			v.reset(OpARMCMNshiftRL)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (CMN x (SRAconst [c] y))
	// result: (CMNshiftRA x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRAconst {
				continue
			}
			c := v_1.AuxInt
			y := v_1.Args[0]
			v.reset(OpARMCMNshiftRA)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (CMN x (SLL y z))
	// result: (CMNshiftLLreg x y z)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSLL {
				continue
			}
			z := v_1.Args[1]
			y := v_1.Args[0]
			v.reset(OpARMCMNshiftLLreg)
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(z)
			return true
		}
		break
	}
	// match: (CMN x (SRL y z))
	// result: (CMNshiftRLreg x y z)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRL {
				continue
			}
			z := v_1.Args[1]
			y := v_1.Args[0]
			v.reset(OpARMCMNshiftRLreg)
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(z)
			return true
		}
		break
	}
	// match: (CMN x (SRA y z))
	// result: (CMNshiftRAreg x y z)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRA {
				continue
			}
			z := v_1.Args[1]
			y := v_1.Args[0]
			v.reset(OpARMCMNshiftRAreg)
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(z)
			return true
		}
		break
	}
	// match: (CMN x (RSBconst [0] y))
	// result: (CMP x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMRSBconst || v_1.AuxInt != 0 {
				continue
			}
			y := v_1.Args[0]
			v.reset(OpARMCMP)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	return false
}
func rewriteValueARM_OpARMCMNconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (CMNconst (MOVWconst [x]) [y])
	// cond: int32(x)==int32(-y)
	// result: (FlagEQ)
	for {
		y := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) == int32(-y)) {
			break
		}
		v.reset(OpARMFlagEQ)
		return true
	}
	// match: (CMNconst (MOVWconst [x]) [y])
	// cond: int32(x)<int32(-y) && uint32(x)<uint32(-y)
	// result: (FlagLT_ULT)
	for {
		y := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) < int32(-y) && uint32(x) < uint32(-y)) {
			break
		}
		v.reset(OpARMFlagLT_ULT)
		return true
	}
	// match: (CMNconst (MOVWconst [x]) [y])
	// cond: int32(x)<int32(-y) && uint32(x)>uint32(-y)
	// result: (FlagLT_UGT)
	for {
		y := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) < int32(-y) && uint32(x) > uint32(-y)) {
			break
		}
		v.reset(OpARMFlagLT_UGT)
		return true
	}
	// match: (CMNconst (MOVWconst [x]) [y])
	// cond: int32(x)>int32(-y) && uint32(x)<uint32(-y)
	// result: (FlagGT_ULT)
	for {
		y := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) > int32(-y) && uint32(x) < uint32(-y)) {
			break
		}
		v.reset(OpARMFlagGT_ULT)
		return true
	}
	// match: (CMNconst (MOVWconst [x]) [y])
	// cond: int32(x)>int32(-y) && uint32(x)>uint32(-y)
	// result: (FlagGT_UGT)
	for {
		y := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) > int32(-y) && uint32(x) > uint32(-y)) {
			break
		}
		v.reset(OpARMFlagGT_UGT)
		return true
	}
	return false
}
func rewriteValueARM_OpARMCMNshiftLL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (CMNshiftLL (MOVWconst [c]) x [d])
	// result: (CMNconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMCMNconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (CMNshiftLL x (MOVWconst [c]) [d])
	// result: (CMNconst x [int64(int32(uint32(c)<<uint64(d)))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMCMNconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMCMNshiftLLreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (CMNshiftLLreg (MOVWconst [c]) x y)
	// result: (CMNconst [c] (SLL <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMCMNconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (CMNshiftLLreg x y (MOVWconst [c]))
	// result: (CMNshiftLL x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMCMNshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMCMNshiftRA(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (CMNshiftRA (MOVWconst [c]) x [d])
	// result: (CMNconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMCMNconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (CMNshiftRA x (MOVWconst [c]) [d])
	// result: (CMNconst x [int64(int32(c)>>uint64(d))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMCMNconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMCMNshiftRAreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (CMNshiftRAreg (MOVWconst [c]) x y)
	// result: (CMNconst [c] (SRA <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMCMNconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRA, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (CMNshiftRAreg x y (MOVWconst [c]))
	// result: (CMNshiftRA x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMCMNshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMCMNshiftRL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (CMNshiftRL (MOVWconst [c]) x [d])
	// result: (CMNconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMCMNconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (CMNshiftRL x (MOVWconst [c]) [d])
	// result: (CMNconst x [int64(int32(uint32(c)>>uint64(d)))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMCMNconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMCMNshiftRLreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (CMNshiftRLreg (MOVWconst [c]) x y)
	// result: (CMNconst [c] (SRL <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMCMNconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (CMNshiftRLreg x y (MOVWconst [c]))
	// result: (CMNshiftRL x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMCMNshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMCMOVWHSconst(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (CMOVWHSconst _ (FlagEQ) [c])
	// result: (MOVWconst [c])
	for {
		c := v.AuxInt
		if v_1.Op != OpARMFlagEQ {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = c
		return true
	}
	// match: (CMOVWHSconst x (FlagLT_ULT))
	// result: x
	for {
		x := v_0
		if v_1.Op != OpARMFlagLT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVWHSconst _ (FlagLT_UGT) [c])
	// result: (MOVWconst [c])
	for {
		c := v.AuxInt
		if v_1.Op != OpARMFlagLT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = c
		return true
	}
	// match: (CMOVWHSconst x (FlagGT_ULT))
	// result: x
	for {
		x := v_0
		if v_1.Op != OpARMFlagGT_ULT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVWHSconst _ (FlagGT_UGT) [c])
	// result: (MOVWconst [c])
	for {
		c := v.AuxInt
		if v_1.Op != OpARMFlagGT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = c
		return true
	}
	// match: (CMOVWHSconst x (InvertFlags flags) [c])
	// result: (CMOVWLSconst x flags [c])
	for {
		c := v.AuxInt
		x := v_0
		if v_1.Op != OpARMInvertFlags {
			break
		}
		flags := v_1.Args[0]
		v.reset(OpARMCMOVWLSconst)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMCMOVWLSconst(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (CMOVWLSconst _ (FlagEQ) [c])
	// result: (MOVWconst [c])
	for {
		c := v.AuxInt
		if v_1.Op != OpARMFlagEQ {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = c
		return true
	}
	// match: (CMOVWLSconst _ (FlagLT_ULT) [c])
	// result: (MOVWconst [c])
	for {
		c := v.AuxInt
		if v_1.Op != OpARMFlagLT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = c
		return true
	}
	// match: (CMOVWLSconst x (FlagLT_UGT))
	// result: x
	for {
		x := v_0
		if v_1.Op != OpARMFlagLT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVWLSconst _ (FlagGT_ULT) [c])
	// result: (MOVWconst [c])
	for {
		c := v.AuxInt
		if v_1.Op != OpARMFlagGT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = c
		return true
	}
	// match: (CMOVWLSconst x (FlagGT_UGT))
	// result: x
	for {
		x := v_0
		if v_1.Op != OpARMFlagGT_UGT {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (CMOVWLSconst x (InvertFlags flags) [c])
	// result: (CMOVWHSconst x flags [c])
	for {
		c := v.AuxInt
		x := v_0
		if v_1.Op != OpARMInvertFlags {
			break
		}
		flags := v_1.Args[0]
		v.reset(OpARMCMOVWHSconst)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMCMP(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (CMP x (MOVWconst [c]))
	// result: (CMPconst [c] x)
	for {
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMCMPconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (CMP (MOVWconst [c]) x)
	// result: (InvertFlags (CMPconst [c] x))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMInvertFlags)
		v0 := b.NewValue0(v.Pos, OpARMCMPconst, types.TypeFlags)
		v0.AuxInt = c
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (CMP x (SLLconst [c] y))
	// result: (CMPshiftLL x y [c])
	for {
		x := v_0
		if v_1.Op != OpARMSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMCMPshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (CMP (SLLconst [c] y) x)
	// result: (InvertFlags (CMPshiftLL x y [c]))
	for {
		if v_0.Op != OpARMSLLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v_1
		v.reset(OpARMInvertFlags)
		v0 := b.NewValue0(v.Pos, OpARMCMPshiftLL, types.TypeFlags)
		v0.AuxInt = c
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (CMP x (SRLconst [c] y))
	// result: (CMPshiftRL x y [c])
	for {
		x := v_0
		if v_1.Op != OpARMSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMCMPshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (CMP (SRLconst [c] y) x)
	// result: (InvertFlags (CMPshiftRL x y [c]))
	for {
		if v_0.Op != OpARMSRLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v_1
		v.reset(OpARMInvertFlags)
		v0 := b.NewValue0(v.Pos, OpARMCMPshiftRL, types.TypeFlags)
		v0.AuxInt = c
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (CMP x (SRAconst [c] y))
	// result: (CMPshiftRA x y [c])
	for {
		x := v_0
		if v_1.Op != OpARMSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMCMPshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (CMP (SRAconst [c] y) x)
	// result: (InvertFlags (CMPshiftRA x y [c]))
	for {
		if v_0.Op != OpARMSRAconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v_1
		v.reset(OpARMInvertFlags)
		v0 := b.NewValue0(v.Pos, OpARMCMPshiftRA, types.TypeFlags)
		v0.AuxInt = c
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (CMP x (SLL y z))
	// result: (CMPshiftLLreg x y z)
	for {
		x := v_0
		if v_1.Op != OpARMSLL {
			break
		}
		z := v_1.Args[1]
		y := v_1.Args[0]
		v.reset(OpARMCMPshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	// match: (CMP (SLL y z) x)
	// result: (InvertFlags (CMPshiftLLreg x y z))
	for {
		if v_0.Op != OpARMSLL {
			break
		}
		z := v_0.Args[1]
		y := v_0.Args[0]
		x := v_1
		v.reset(OpARMInvertFlags)
		v0 := b.NewValue0(v.Pos, OpARMCMPshiftLLreg, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}
	// match: (CMP x (SRL y z))
	// result: (CMPshiftRLreg x y z)
	for {
		x := v_0
		if v_1.Op != OpARMSRL {
			break
		}
		z := v_1.Args[1]
		y := v_1.Args[0]
		v.reset(OpARMCMPshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	// match: (CMP (SRL y z) x)
	// result: (InvertFlags (CMPshiftRLreg x y z))
	for {
		if v_0.Op != OpARMSRL {
			break
		}
		z := v_0.Args[1]
		y := v_0.Args[0]
		x := v_1
		v.reset(OpARMInvertFlags)
		v0 := b.NewValue0(v.Pos, OpARMCMPshiftRLreg, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}
	// match: (CMP x (SRA y z))
	// result: (CMPshiftRAreg x y z)
	for {
		x := v_0
		if v_1.Op != OpARMSRA {
			break
		}
		z := v_1.Args[1]
		y := v_1.Args[0]
		v.reset(OpARMCMPshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	// match: (CMP (SRA y z) x)
	// result: (InvertFlags (CMPshiftRAreg x y z))
	for {
		if v_0.Op != OpARMSRA {
			break
		}
		z := v_0.Args[1]
		y := v_0.Args[0]
		x := v_1
		v.reset(OpARMInvertFlags)
		v0 := b.NewValue0(v.Pos, OpARMCMPshiftRAreg, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v0.AddArg(z)
		v.AddArg(v0)
		return true
	}
	// match: (CMP x (RSBconst [0] y))
	// result: (CMN x y)
	for {
		x := v_0
		if v_1.Op != OpARMRSBconst || v_1.AuxInt != 0 {
			break
		}
		y := v_1.Args[0]
		v.reset(OpARMCMN)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMCMPD(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (CMPD x (MOVDconst [0]))
	// result: (CMPD0 x)
	for {
		x := v_0
		if v_1.Op != OpARMMOVDconst || v_1.AuxInt != 0 {
			break
		}
		v.reset(OpARMCMPD0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMCMPF(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (CMPF x (MOVFconst [0]))
	// result: (CMPF0 x)
	for {
		x := v_0
		if v_1.Op != OpARMMOVFconst || v_1.AuxInt != 0 {
			break
		}
		v.reset(OpARMCMPF0)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMCMPconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (CMPconst (MOVWconst [x]) [y])
	// cond: int32(x)==int32(y)
	// result: (FlagEQ)
	for {
		y := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) == int32(y)) {
			break
		}
		v.reset(OpARMFlagEQ)
		return true
	}
	// match: (CMPconst (MOVWconst [x]) [y])
	// cond: int32(x)<int32(y) && uint32(x)<uint32(y)
	// result: (FlagLT_ULT)
	for {
		y := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) < int32(y) && uint32(x) < uint32(y)) {
			break
		}
		v.reset(OpARMFlagLT_ULT)
		return true
	}
	// match: (CMPconst (MOVWconst [x]) [y])
	// cond: int32(x)<int32(y) && uint32(x)>uint32(y)
	// result: (FlagLT_UGT)
	for {
		y := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) < int32(y) && uint32(x) > uint32(y)) {
			break
		}
		v.reset(OpARMFlagLT_UGT)
		return true
	}
	// match: (CMPconst (MOVWconst [x]) [y])
	// cond: int32(x)>int32(y) && uint32(x)<uint32(y)
	// result: (FlagGT_ULT)
	for {
		y := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) > int32(y) && uint32(x) < uint32(y)) {
			break
		}
		v.reset(OpARMFlagGT_ULT)
		return true
	}
	// match: (CMPconst (MOVWconst [x]) [y])
	// cond: int32(x)>int32(y) && uint32(x)>uint32(y)
	// result: (FlagGT_UGT)
	for {
		y := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x) > int32(y) && uint32(x) > uint32(y)) {
			break
		}
		v.reset(OpARMFlagGT_UGT)
		return true
	}
	// match: (CMPconst (MOVBUreg _) [c])
	// cond: 0xff < c
	// result: (FlagLT_ULT)
	for {
		c := v.AuxInt
		if v_0.Op != OpARMMOVBUreg || !(0xff < c) {
			break
		}
		v.reset(OpARMFlagLT_ULT)
		return true
	}
	// match: (CMPconst (MOVHUreg _) [c])
	// cond: 0xffff < c
	// result: (FlagLT_ULT)
	for {
		c := v.AuxInt
		if v_0.Op != OpARMMOVHUreg || !(0xffff < c) {
			break
		}
		v.reset(OpARMFlagLT_ULT)
		return true
	}
	// match: (CMPconst (ANDconst _ [m]) [n])
	// cond: 0 <= int32(m) && int32(m) < int32(n)
	// result: (FlagLT_ULT)
	for {
		n := v.AuxInt
		if v_0.Op != OpARMANDconst {
			break
		}
		m := v_0.AuxInt
		if !(0 <= int32(m) && int32(m) < int32(n)) {
			break
		}
		v.reset(OpARMFlagLT_ULT)
		return true
	}
	// match: (CMPconst (SRLconst _ [c]) [n])
	// cond: 0 <= n && 0 < c && c <= 32 && (1<<uint32(32-c)) <= uint32(n)
	// result: (FlagLT_ULT)
	for {
		n := v.AuxInt
		if v_0.Op != OpARMSRLconst {
			break
		}
		c := v_0.AuxInt
		if !(0 <= n && 0 < c && c <= 32 && (1<<uint32(32-c)) <= uint32(n)) {
			break
		}
		v.reset(OpARMFlagLT_ULT)
		return true
	}
	return false
}
func rewriteValueARM_OpARMCMPshiftLL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (CMPshiftLL (MOVWconst [c]) x [d])
	// result: (InvertFlags (CMPconst [c] (SLLconst <x.Type> x [d])))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMInvertFlags)
		v0 := b.NewValue0(v.Pos, OpARMCMPconst, types.TypeFlags)
		v0.AuxInt = c
		v1 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v1.AuxInt = d
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (CMPshiftLL x (MOVWconst [c]) [d])
	// result: (CMPconst x [int64(int32(uint32(c)<<uint64(d)))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMCMPconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMCMPshiftLLreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (CMPshiftLLreg (MOVWconst [c]) x y)
	// result: (InvertFlags (CMPconst [c] (SLL <x.Type> x y)))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMInvertFlags)
		v0 := b.NewValue0(v.Pos, OpARMCMPconst, types.TypeFlags)
		v0.AuxInt = c
		v1 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (CMPshiftLLreg x y (MOVWconst [c]))
	// result: (CMPshiftLL x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMCMPshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMCMPshiftRA(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (CMPshiftRA (MOVWconst [c]) x [d])
	// result: (InvertFlags (CMPconst [c] (SRAconst <x.Type> x [d])))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMInvertFlags)
		v0 := b.NewValue0(v.Pos, OpARMCMPconst, types.TypeFlags)
		v0.AuxInt = c
		v1 := b.NewValue0(v.Pos, OpARMSRAconst, x.Type)
		v1.AuxInt = d
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (CMPshiftRA x (MOVWconst [c]) [d])
	// result: (CMPconst x [int64(int32(c)>>uint64(d))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMCMPconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMCMPshiftRAreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (CMPshiftRAreg (MOVWconst [c]) x y)
	// result: (InvertFlags (CMPconst [c] (SRA <x.Type> x y)))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMInvertFlags)
		v0 := b.NewValue0(v.Pos, OpARMCMPconst, types.TypeFlags)
		v0.AuxInt = c
		v1 := b.NewValue0(v.Pos, OpARMSRA, x.Type)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (CMPshiftRAreg x y (MOVWconst [c]))
	// result: (CMPshiftRA x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMCMPshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMCMPshiftRL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (CMPshiftRL (MOVWconst [c]) x [d])
	// result: (InvertFlags (CMPconst [c] (SRLconst <x.Type> x [d])))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMInvertFlags)
		v0 := b.NewValue0(v.Pos, OpARMCMPconst, types.TypeFlags)
		v0.AuxInt = c
		v1 := b.NewValue0(v.Pos, OpARMSRLconst, x.Type)
		v1.AuxInt = d
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (CMPshiftRL x (MOVWconst [c]) [d])
	// result: (CMPconst x [int64(int32(uint32(c)>>uint64(d)))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMCMPconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMCMPshiftRLreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (CMPshiftRLreg (MOVWconst [c]) x y)
	// result: (InvertFlags (CMPconst [c] (SRL <x.Type> x y)))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMInvertFlags)
		v0 := b.NewValue0(v.Pos, OpARMCMPconst, types.TypeFlags)
		v0.AuxInt = c
		v1 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (CMPshiftRLreg x y (MOVWconst [c]))
	// result: (CMPshiftRL x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMCMPshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMEqual(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Equal (FlagEQ))
	// result: (MOVWconst [1])
	for {
		if v_0.Op != OpARMFlagEQ {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (Equal (FlagLT_ULT))
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpARMFlagLT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (Equal (FlagLT_UGT))
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpARMFlagLT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (Equal (FlagGT_ULT))
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpARMFlagGT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (Equal (FlagGT_UGT))
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpARMFlagGT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (Equal (InvertFlags x))
	// result: (Equal x)
	for {
		if v_0.Op != OpARMInvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpARMEqual)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMGreaterEqual(v *Value) bool {
	v_0 := v.Args[0]
	// match: (GreaterEqual (FlagEQ))
	// result: (MOVWconst [1])
	for {
		if v_0.Op != OpARMFlagEQ {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (GreaterEqual (FlagLT_ULT))
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpARMFlagLT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (GreaterEqual (FlagLT_UGT))
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpARMFlagLT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (GreaterEqual (FlagGT_ULT))
	// result: (MOVWconst [1])
	for {
		if v_0.Op != OpARMFlagGT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (GreaterEqual (FlagGT_UGT))
	// result: (MOVWconst [1])
	for {
		if v_0.Op != OpARMFlagGT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (GreaterEqual (InvertFlags x))
	// result: (LessEqual x)
	for {
		if v_0.Op != OpARMInvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpARMLessEqual)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMGreaterEqualU(v *Value) bool {
	v_0 := v.Args[0]
	// match: (GreaterEqualU (FlagEQ))
	// result: (MOVWconst [1])
	for {
		if v_0.Op != OpARMFlagEQ {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (GreaterEqualU (FlagLT_ULT))
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpARMFlagLT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (GreaterEqualU (FlagLT_UGT))
	// result: (MOVWconst [1])
	for {
		if v_0.Op != OpARMFlagLT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (GreaterEqualU (FlagGT_ULT))
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpARMFlagGT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (GreaterEqualU (FlagGT_UGT))
	// result: (MOVWconst [1])
	for {
		if v_0.Op != OpARMFlagGT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (GreaterEqualU (InvertFlags x))
	// result: (LessEqualU x)
	for {
		if v_0.Op != OpARMInvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpARMLessEqualU)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMGreaterThan(v *Value) bool {
	v_0 := v.Args[0]
	// match: (GreaterThan (FlagEQ))
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpARMFlagEQ {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (GreaterThan (FlagLT_ULT))
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpARMFlagLT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (GreaterThan (FlagLT_UGT))
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpARMFlagLT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (GreaterThan (FlagGT_ULT))
	// result: (MOVWconst [1])
	for {
		if v_0.Op != OpARMFlagGT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (GreaterThan (FlagGT_UGT))
	// result: (MOVWconst [1])
	for {
		if v_0.Op != OpARMFlagGT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (GreaterThan (InvertFlags x))
	// result: (LessThan x)
	for {
		if v_0.Op != OpARMInvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpARMLessThan)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMGreaterThanU(v *Value) bool {
	v_0 := v.Args[0]
	// match: (GreaterThanU (FlagEQ))
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpARMFlagEQ {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (GreaterThanU (FlagLT_ULT))
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpARMFlagLT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (GreaterThanU (FlagLT_UGT))
	// result: (MOVWconst [1])
	for {
		if v_0.Op != OpARMFlagLT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (GreaterThanU (FlagGT_ULT))
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpARMFlagGT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (GreaterThanU (FlagGT_UGT))
	// result: (MOVWconst [1])
	for {
		if v_0.Op != OpARMFlagGT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (GreaterThanU (InvertFlags x))
	// result: (LessThanU x)
	for {
		if v_0.Op != OpARMInvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpARMLessThanU)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMLessEqual(v *Value) bool {
	v_0 := v.Args[0]
	// match: (LessEqual (FlagEQ))
	// result: (MOVWconst [1])
	for {
		if v_0.Op != OpARMFlagEQ {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (LessEqual (FlagLT_ULT))
	// result: (MOVWconst [1])
	for {
		if v_0.Op != OpARMFlagLT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (LessEqual (FlagLT_UGT))
	// result: (MOVWconst [1])
	for {
		if v_0.Op != OpARMFlagLT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (LessEqual (FlagGT_ULT))
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpARMFlagGT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (LessEqual (FlagGT_UGT))
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpARMFlagGT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (LessEqual (InvertFlags x))
	// result: (GreaterEqual x)
	for {
		if v_0.Op != OpARMInvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpARMGreaterEqual)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMLessEqualU(v *Value) bool {
	v_0 := v.Args[0]
	// match: (LessEqualU (FlagEQ))
	// result: (MOVWconst [1])
	for {
		if v_0.Op != OpARMFlagEQ {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (LessEqualU (FlagLT_ULT))
	// result: (MOVWconst [1])
	for {
		if v_0.Op != OpARMFlagLT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (LessEqualU (FlagLT_UGT))
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpARMFlagLT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (LessEqualU (FlagGT_ULT))
	// result: (MOVWconst [1])
	for {
		if v_0.Op != OpARMFlagGT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (LessEqualU (FlagGT_UGT))
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpARMFlagGT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (LessEqualU (InvertFlags x))
	// result: (GreaterEqualU x)
	for {
		if v_0.Op != OpARMInvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpARMGreaterEqualU)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMLessThan(v *Value) bool {
	v_0 := v.Args[0]
	// match: (LessThan (FlagEQ))
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpARMFlagEQ {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (LessThan (FlagLT_ULT))
	// result: (MOVWconst [1])
	for {
		if v_0.Op != OpARMFlagLT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (LessThan (FlagLT_UGT))
	// result: (MOVWconst [1])
	for {
		if v_0.Op != OpARMFlagLT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (LessThan (FlagGT_ULT))
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpARMFlagGT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (LessThan (FlagGT_UGT))
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpARMFlagGT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (LessThan (InvertFlags x))
	// result: (GreaterThan x)
	for {
		if v_0.Op != OpARMInvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpARMGreaterThan)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMLessThanU(v *Value) bool {
	v_0 := v.Args[0]
	// match: (LessThanU (FlagEQ))
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpARMFlagEQ {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (LessThanU (FlagLT_ULT))
	// result: (MOVWconst [1])
	for {
		if v_0.Op != OpARMFlagLT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (LessThanU (FlagLT_UGT))
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpARMFlagLT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (LessThanU (FlagGT_ULT))
	// result: (MOVWconst [1])
	for {
		if v_0.Op != OpARMFlagGT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (LessThanU (FlagGT_UGT))
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpARMFlagGT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (LessThanU (InvertFlags x))
	// result: (GreaterThanU x)
	for {
		if v_0.Op != OpARMInvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpARMGreaterThanU)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVBUload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBUload [off1] {sym} (ADDconst [off2] ptr) mem)
	// result: (MOVBUload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpARMADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpARMMOVBUload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBUload [off1] {sym} (SUBconst [off2] ptr) mem)
	// result: (MOVBUload [off1-off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpARMSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpARMMOVBUload)
		v.AuxInt = off1 - off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBUload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVBUload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpARMMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpARMMOVBUload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBUload [off] {sym} ptr (MOVBstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: (MOVBUreg x)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpARMMOVBstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpARMMOVBUreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBUload [0] {sym} (ADD ptr idx) mem)
	// cond: sym == nil
	// result: (MOVBUloadidx ptr idx mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		if v_0.Op != OpARMADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_1
		if !(sym == nil) {
			break
		}
		v.reset(OpARMMOVBUloadidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBUload [off] {sym} (SB) _)
	// cond: symIsRO(sym)
	// result: (MOVWconst [int64(read8(sym, off))])
	for {
		off := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpSB || !(symIsRO(sym)) {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = int64(read8(sym, off))
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVBUloadidx(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBUloadidx ptr idx (MOVBstoreidx ptr2 idx x _))
	// cond: isSamePtr(ptr, ptr2)
	// result: (MOVBUreg x)
	for {
		ptr := v_0
		idx := v_1
		if v_2.Op != OpARMMOVBstoreidx {
			break
		}
		_ = v_2.Args[3]
		ptr2 := v_2.Args[0]
		if idx != v_2.Args[1] {
			break
		}
		x := v_2.Args[2]
		if !(isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpARMMOVBUreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBUloadidx ptr (MOVWconst [c]) mem)
	// result: (MOVBUload [c] ptr mem)
	for {
		ptr := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		mem := v_2
		v.reset(OpARMMOVBUload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBUloadidx (MOVWconst [c]) ptr mem)
	// result: (MOVBUload [c] ptr mem)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		ptr := v_1
		mem := v_2
		v.reset(OpARMMOVBUload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVBUreg(v *Value) bool {
	v_0 := v.Args[0]
	// match: (MOVBUreg x:(MOVBUload _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpARMMOVBUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpARMMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBUreg (ANDconst [c] x))
	// result: (ANDconst [c&0xff] x)
	for {
		if v_0.Op != OpARMANDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMANDconst)
		v.AuxInt = c & 0xff
		v.AddArg(x)
		return true
	}
	// match: (MOVBUreg x:(MOVBUreg _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpARMMOVBUreg {
			break
		}
		v.reset(OpARMMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBUreg (MOVWconst [c]))
	// result: (MOVWconst [int64(uint8(c))])
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = int64(uint8(c))
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVBload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBload [off1] {sym} (ADDconst [off2] ptr) mem)
	// result: (MOVBload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpARMADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpARMMOVBload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBload [off1] {sym} (SUBconst [off2] ptr) mem)
	// result: (MOVBload [off1-off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpARMSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpARMMOVBload)
		v.AuxInt = off1 - off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVBload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpARMMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpARMMOVBload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBload [off] {sym} ptr (MOVBstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: (MOVBreg x)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpARMMOVBstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpARMMOVBreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBload [0] {sym} (ADD ptr idx) mem)
	// cond: sym == nil
	// result: (MOVBloadidx ptr idx mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		if v_0.Op != OpARMADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_1
		if !(sym == nil) {
			break
		}
		v.reset(OpARMMOVBloadidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVBloadidx(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBloadidx ptr idx (MOVBstoreidx ptr2 idx x _))
	// cond: isSamePtr(ptr, ptr2)
	// result: (MOVBreg x)
	for {
		ptr := v_0
		idx := v_1
		if v_2.Op != OpARMMOVBstoreidx {
			break
		}
		_ = v_2.Args[3]
		ptr2 := v_2.Args[0]
		if idx != v_2.Args[1] {
			break
		}
		x := v_2.Args[2]
		if !(isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpARMMOVBreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBloadidx ptr (MOVWconst [c]) mem)
	// result: (MOVBload [c] ptr mem)
	for {
		ptr := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		mem := v_2
		v.reset(OpARMMOVBload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBloadidx (MOVWconst [c]) ptr mem)
	// result: (MOVBload [c] ptr mem)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		ptr := v_1
		mem := v_2
		v.reset(OpARMMOVBload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVBreg(v *Value) bool {
	v_0 := v.Args[0]
	// match: (MOVBreg x:(MOVBload _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpARMMOVBload {
			break
		}
		_ = x.Args[1]
		v.reset(OpARMMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg (ANDconst [c] x))
	// cond: c & 0x80 == 0
	// result: (ANDconst [c&0x7f] x)
	for {
		if v_0.Op != OpARMANDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(c&0x80 == 0) {
			break
		}
		v.reset(OpARMANDconst)
		v.AuxInt = c & 0x7f
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg x:(MOVBreg _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpARMMOVBreg {
			break
		}
		v.reset(OpARMMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVBreg (MOVWconst [c]))
	// result: (MOVWconst [int64(int8(c))])
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = int64(int8(c))
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVBstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBstore [off1] {sym} (ADDconst [off2] ptr) val mem)
	// result: (MOVBstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpARMADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		v.reset(OpARMMOVBstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off1] {sym} (SUBconst [off2] ptr) val mem)
	// result: (MOVBstore [off1-off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpARMSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		v.reset(OpARMMOVBstore)
		v.AuxInt = off1 - off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVBstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpARMMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpARMMOVBstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [off] {sym} ptr (MOVBreg x) mem)
	// result: (MOVBstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpARMMOVBreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpARMMOVBstore)
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
		if v_1.Op != OpARMMOVBUreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpARMMOVBstore)
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
		if v_1.Op != OpARMMOVHreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpARMMOVBstore)
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
		if v_1.Op != OpARMMOVHUreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpARMMOVBstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstore [0] {sym} (ADD ptr idx) val mem)
	// cond: sym == nil
	// result: (MOVBstoreidx ptr idx val mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		if v_0.Op != OpARMADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(sym == nil) {
			break
		}
		v.reset(OpARMMOVBstoreidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVBstoreidx(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBstoreidx ptr (MOVWconst [c]) val mem)
	// result: (MOVBstore [c] ptr val mem)
	for {
		ptr := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		val := v_2
		mem := v_3
		v.reset(OpARMMOVBstore)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVBstoreidx (MOVWconst [c]) ptr val mem)
	// result: (MOVBstore [c] ptr val mem)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		ptr := v_1
		val := v_2
		mem := v_3
		v.reset(OpARMMOVBstore)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVDload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVDload [off1] {sym} (ADDconst [off2] ptr) mem)
	// result: (MOVDload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpARMADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpARMMOVDload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDload [off1] {sym} (SUBconst [off2] ptr) mem)
	// result: (MOVDload [off1-off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpARMSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpARMMOVDload)
		v.AuxInt = off1 - off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVDload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpARMMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpARMMOVDload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDload [off] {sym} ptr (MOVDstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: x
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpARMMOVDstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVDstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVDstore [off1] {sym} (ADDconst [off2] ptr) val mem)
	// result: (MOVDstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpARMADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		v.reset(OpARMMOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDstore [off1] {sym} (SUBconst [off2] ptr) val mem)
	// result: (MOVDstore [off1-off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpARMSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		v.reset(OpARMMOVDstore)
		v.AuxInt = off1 - off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVDstore [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVDstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpARMMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpARMMOVDstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVFload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVFload [off1] {sym} (ADDconst [off2] ptr) mem)
	// result: (MOVFload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpARMADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpARMMOVFload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVFload [off1] {sym} (SUBconst [off2] ptr) mem)
	// result: (MOVFload [off1-off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpARMSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpARMMOVFload)
		v.AuxInt = off1 - off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVFload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVFload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpARMMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpARMMOVFload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVFload [off] {sym} ptr (MOVFstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: x
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpARMMOVFstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVFstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVFstore [off1] {sym} (ADDconst [off2] ptr) val mem)
	// result: (MOVFstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpARMADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		v.reset(OpARMMOVFstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVFstore [off1] {sym} (SUBconst [off2] ptr) val mem)
	// result: (MOVFstore [off1-off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpARMSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		v.reset(OpARMMOVFstore)
		v.AuxInt = off1 - off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVFstore [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVFstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpARMMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpARMMOVFstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVHUload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	// match: (MOVHUload [off1] {sym} (ADDconst [off2] ptr) mem)
	// result: (MOVHUload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpARMADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpARMMOVHUload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHUload [off1] {sym} (SUBconst [off2] ptr) mem)
	// result: (MOVHUload [off1-off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpARMSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpARMMOVHUload)
		v.AuxInt = off1 - off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHUload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVHUload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpARMMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpARMMOVHUload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHUload [off] {sym} ptr (MOVHstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: (MOVHUreg x)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpARMMOVHstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpARMMOVHUreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUload [0] {sym} (ADD ptr idx) mem)
	// cond: sym == nil
	// result: (MOVHUloadidx ptr idx mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		if v_0.Op != OpARMADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_1
		if !(sym == nil) {
			break
		}
		v.reset(OpARMMOVHUloadidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHUload [off] {sym} (SB) _)
	// cond: symIsRO(sym)
	// result: (MOVWconst [int64(read16(sym, off, config.BigEndian))])
	for {
		off := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpSB || !(symIsRO(sym)) {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = int64(read16(sym, off, config.BigEndian))
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVHUloadidx(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHUloadidx ptr idx (MOVHstoreidx ptr2 idx x _))
	// cond: isSamePtr(ptr, ptr2)
	// result: (MOVHUreg x)
	for {
		ptr := v_0
		idx := v_1
		if v_2.Op != OpARMMOVHstoreidx {
			break
		}
		_ = v_2.Args[3]
		ptr2 := v_2.Args[0]
		if idx != v_2.Args[1] {
			break
		}
		x := v_2.Args[2]
		if !(isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpARMMOVHUreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUloadidx ptr (MOVWconst [c]) mem)
	// result: (MOVHUload [c] ptr mem)
	for {
		ptr := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		mem := v_2
		v.reset(OpARMMOVHUload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHUloadidx (MOVWconst [c]) ptr mem)
	// result: (MOVHUload [c] ptr mem)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		ptr := v_1
		mem := v_2
		v.reset(OpARMMOVHUload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVHUreg(v *Value) bool {
	v_0 := v.Args[0]
	// match: (MOVHUreg x:(MOVBUload _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpARMMOVBUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpARMMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg x:(MOVHUload _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpARMMOVHUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpARMMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg (ANDconst [c] x))
	// result: (ANDconst [c&0xffff] x)
	for {
		if v_0.Op != OpARMANDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMANDconst)
		v.AuxInt = c & 0xffff
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg x:(MOVBUreg _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpARMMOVBUreg {
			break
		}
		v.reset(OpARMMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg x:(MOVHUreg _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpARMMOVHUreg {
			break
		}
		v.reset(OpARMMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHUreg (MOVWconst [c]))
	// result: (MOVWconst [int64(uint16(c))])
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = int64(uint16(c))
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVHload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHload [off1] {sym} (ADDconst [off2] ptr) mem)
	// result: (MOVHload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpARMADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpARMMOVHload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHload [off1] {sym} (SUBconst [off2] ptr) mem)
	// result: (MOVHload [off1-off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpARMSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpARMMOVHload)
		v.AuxInt = off1 - off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVHload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpARMMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpARMMOVHload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHload [off] {sym} ptr (MOVHstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: (MOVHreg x)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpARMMOVHstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpARMMOVHreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHload [0] {sym} (ADD ptr idx) mem)
	// cond: sym == nil
	// result: (MOVHloadidx ptr idx mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		if v_0.Op != OpARMADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_1
		if !(sym == nil) {
			break
		}
		v.reset(OpARMMOVHloadidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVHloadidx(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHloadidx ptr idx (MOVHstoreidx ptr2 idx x _))
	// cond: isSamePtr(ptr, ptr2)
	// result: (MOVHreg x)
	for {
		ptr := v_0
		idx := v_1
		if v_2.Op != OpARMMOVHstoreidx {
			break
		}
		_ = v_2.Args[3]
		ptr2 := v_2.Args[0]
		if idx != v_2.Args[1] {
			break
		}
		x := v_2.Args[2]
		if !(isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpARMMOVHreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHloadidx ptr (MOVWconst [c]) mem)
	// result: (MOVHload [c] ptr mem)
	for {
		ptr := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		mem := v_2
		v.reset(OpARMMOVHload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHloadidx (MOVWconst [c]) ptr mem)
	// result: (MOVHload [c] ptr mem)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		ptr := v_1
		mem := v_2
		v.reset(OpARMMOVHload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVHreg(v *Value) bool {
	v_0 := v.Args[0]
	// match: (MOVHreg x:(MOVBload _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpARMMOVBload {
			break
		}
		_ = x.Args[1]
		v.reset(OpARMMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVBUload _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpARMMOVBUload {
			break
		}
		_ = x.Args[1]
		v.reset(OpARMMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVHload _ _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpARMMOVHload {
			break
		}
		_ = x.Args[1]
		v.reset(OpARMMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg (ANDconst [c] x))
	// cond: c & 0x8000 == 0
	// result: (ANDconst [c&0x7fff] x)
	for {
		if v_0.Op != OpARMANDconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(c&0x8000 == 0) {
			break
		}
		v.reset(OpARMANDconst)
		v.AuxInt = c & 0x7fff
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVBreg _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpARMMOVBreg {
			break
		}
		v.reset(OpARMMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVBUreg _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpARMMOVBUreg {
			break
		}
		v.reset(OpARMMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg x:(MOVHreg _))
	// result: (MOVWreg x)
	for {
		x := v_0
		if x.Op != OpARMMOVHreg {
			break
		}
		v.reset(OpARMMOVWreg)
		v.AddArg(x)
		return true
	}
	// match: (MOVHreg (MOVWconst [c]))
	// result: (MOVWconst [int64(int16(c))])
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = int64(int16(c))
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVHstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHstore [off1] {sym} (ADDconst [off2] ptr) val mem)
	// result: (MOVHstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpARMADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		v.reset(OpARMMOVHstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off1] {sym} (SUBconst [off2] ptr) val mem)
	// result: (MOVHstore [off1-off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpARMSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		v.reset(OpARMMOVHstore)
		v.AuxInt = off1 - off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVHstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpARMMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpARMMOVHstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [off] {sym} ptr (MOVHreg x) mem)
	// result: (MOVHstore [off] {sym} ptr x mem)
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpARMMOVHreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpARMMOVHstore)
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
		if v_1.Op != OpARMMOVHUreg {
			break
		}
		x := v_1.Args[0]
		mem := v_2
		v.reset(OpARMMOVHstore)
		v.AuxInt = off
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(x)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstore [0] {sym} (ADD ptr idx) val mem)
	// cond: sym == nil
	// result: (MOVHstoreidx ptr idx val mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		if v_0.Op != OpARMADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(sym == nil) {
			break
		}
		v.reset(OpARMMOVHstoreidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVHstoreidx(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHstoreidx ptr (MOVWconst [c]) val mem)
	// result: (MOVHstore [c] ptr val mem)
	for {
		ptr := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		val := v_2
		mem := v_3
		v.reset(OpARMMOVHstore)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVHstoreidx (MOVWconst [c]) ptr val mem)
	// result: (MOVHstore [c] ptr val mem)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		ptr := v_1
		val := v_2
		mem := v_3
		v.reset(OpARMMOVHstore)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVWload(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	// match: (MOVWload [off1] {sym} (ADDconst [off2] ptr) mem)
	// result: (MOVWload [off1+off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpARMADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpARMMOVWload)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWload [off1] {sym} (SUBconst [off2] ptr) mem)
	// result: (MOVWload [off1-off2] {sym} ptr mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpARMSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		mem := v_1
		v.reset(OpARMMOVWload)
		v.AuxInt = off1 - off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWload [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVWload [off1+off2] {mergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpARMMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		mem := v_1
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpARMMOVWload)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWload [off] {sym} ptr (MOVWstore [off2] {sym2} ptr2 x _))
	// cond: sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)
	// result: x
	for {
		off := v.AuxInt
		sym := v.Aux
		ptr := v_0
		if v_1.Op != OpARMMOVWstore {
			break
		}
		off2 := v_1.AuxInt
		sym2 := v_1.Aux
		_ = v_1.Args[2]
		ptr2 := v_1.Args[0]
		x := v_1.Args[1]
		if !(sym == sym2 && off == off2 && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWload [0] {sym} (ADD ptr idx) mem)
	// cond: sym == nil
	// result: (MOVWloadidx ptr idx mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		if v_0.Op != OpARMADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_1
		if !(sym == nil) {
			break
		}
		v.reset(OpARMMOVWloadidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWload [0] {sym} (ADDshiftLL ptr idx [c]) mem)
	// cond: sym == nil
	// result: (MOVWloadshiftLL ptr idx [c] mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		if v_0.Op != OpARMADDshiftLL {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_1
		if !(sym == nil) {
			break
		}
		v.reset(OpARMMOVWloadshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWload [0] {sym} (ADDshiftRL ptr idx [c]) mem)
	// cond: sym == nil
	// result: (MOVWloadshiftRL ptr idx [c] mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		if v_0.Op != OpARMADDshiftRL {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_1
		if !(sym == nil) {
			break
		}
		v.reset(OpARMMOVWloadshiftRL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWload [0] {sym} (ADDshiftRA ptr idx [c]) mem)
	// cond: sym == nil
	// result: (MOVWloadshiftRA ptr idx [c] mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		if v_0.Op != OpARMADDshiftRA {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		mem := v_1
		if !(sym == nil) {
			break
		}
		v.reset(OpARMMOVWloadshiftRA)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWload [off] {sym} (SB) _)
	// cond: symIsRO(sym)
	// result: (MOVWconst [int64(int32(read32(sym, off, config.BigEndian)))])
	for {
		off := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpSB || !(symIsRO(sym)) {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = int64(int32(read32(sym, off, config.BigEndian)))
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVWloadidx(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWloadidx ptr idx (MOVWstoreidx ptr2 idx x _))
	// cond: isSamePtr(ptr, ptr2)
	// result: x
	for {
		ptr := v_0
		idx := v_1
		if v_2.Op != OpARMMOVWstoreidx {
			break
		}
		_ = v_2.Args[3]
		ptr2 := v_2.Args[0]
		if idx != v_2.Args[1] {
			break
		}
		x := v_2.Args[2]
		if !(isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWloadidx ptr (MOVWconst [c]) mem)
	// result: (MOVWload [c] ptr mem)
	for {
		ptr := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		mem := v_2
		v.reset(OpARMMOVWload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWloadidx (MOVWconst [c]) ptr mem)
	// result: (MOVWload [c] ptr mem)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		ptr := v_1
		mem := v_2
		v.reset(OpARMMOVWload)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWloadidx ptr (SLLconst idx [c]) mem)
	// result: (MOVWloadshiftLL ptr idx [c] mem)
	for {
		ptr := v_0
		if v_1.Op != OpARMSLLconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		mem := v_2
		v.reset(OpARMMOVWloadshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWloadidx (SLLconst idx [c]) ptr mem)
	// result: (MOVWloadshiftLL ptr idx [c] mem)
	for {
		if v_0.Op != OpARMSLLconst {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v_1
		mem := v_2
		v.reset(OpARMMOVWloadshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWloadidx ptr (SRLconst idx [c]) mem)
	// result: (MOVWloadshiftRL ptr idx [c] mem)
	for {
		ptr := v_0
		if v_1.Op != OpARMSRLconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		mem := v_2
		v.reset(OpARMMOVWloadshiftRL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWloadidx (SRLconst idx [c]) ptr mem)
	// result: (MOVWloadshiftRL ptr idx [c] mem)
	for {
		if v_0.Op != OpARMSRLconst {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v_1
		mem := v_2
		v.reset(OpARMMOVWloadshiftRL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWloadidx ptr (SRAconst idx [c]) mem)
	// result: (MOVWloadshiftRA ptr idx [c] mem)
	for {
		ptr := v_0
		if v_1.Op != OpARMSRAconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		mem := v_2
		v.reset(OpARMMOVWloadshiftRA)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWloadidx (SRAconst idx [c]) ptr mem)
	// result: (MOVWloadshiftRA ptr idx [c] mem)
	for {
		if v_0.Op != OpARMSRAconst {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v_1
		mem := v_2
		v.reset(OpARMMOVWloadshiftRA)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVWloadshiftLL(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWloadshiftLL ptr idx [c] (MOVWstoreshiftLL ptr2 idx [d] x _))
	// cond: c==d && isSamePtr(ptr, ptr2)
	// result: x
	for {
		c := v.AuxInt
		ptr := v_0
		idx := v_1
		if v_2.Op != OpARMMOVWstoreshiftLL {
			break
		}
		d := v_2.AuxInt
		_ = v_2.Args[3]
		ptr2 := v_2.Args[0]
		if idx != v_2.Args[1] {
			break
		}
		x := v_2.Args[2]
		if !(c == d && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWloadshiftLL ptr (MOVWconst [c]) [d] mem)
	// result: (MOVWload [int64(uint32(c)<<uint64(d))] ptr mem)
	for {
		d := v.AuxInt
		ptr := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		mem := v_2
		v.reset(OpARMMOVWload)
		v.AuxInt = int64(uint32(c) << uint64(d))
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVWloadshiftRA(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWloadshiftRA ptr idx [c] (MOVWstoreshiftRA ptr2 idx [d] x _))
	// cond: c==d && isSamePtr(ptr, ptr2)
	// result: x
	for {
		c := v.AuxInt
		ptr := v_0
		idx := v_1
		if v_2.Op != OpARMMOVWstoreshiftRA {
			break
		}
		d := v_2.AuxInt
		_ = v_2.Args[3]
		ptr2 := v_2.Args[0]
		if idx != v_2.Args[1] {
			break
		}
		x := v_2.Args[2]
		if !(c == d && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWloadshiftRA ptr (MOVWconst [c]) [d] mem)
	// result: (MOVWload [int64(int32(c)>>uint64(d))] ptr mem)
	for {
		d := v.AuxInt
		ptr := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		mem := v_2
		v.reset(OpARMMOVWload)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVWloadshiftRL(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWloadshiftRL ptr idx [c] (MOVWstoreshiftRL ptr2 idx [d] x _))
	// cond: c==d && isSamePtr(ptr, ptr2)
	// result: x
	for {
		c := v.AuxInt
		ptr := v_0
		idx := v_1
		if v_2.Op != OpARMMOVWstoreshiftRL {
			break
		}
		d := v_2.AuxInt
		_ = v_2.Args[3]
		ptr2 := v_2.Args[0]
		if idx != v_2.Args[1] {
			break
		}
		x := v_2.Args[2]
		if !(c == d && isSamePtr(ptr, ptr2)) {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (MOVWloadshiftRL ptr (MOVWconst [c]) [d] mem)
	// result: (MOVWload [int64(uint32(c)>>uint64(d))] ptr mem)
	for {
		d := v.AuxInt
		ptr := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		mem := v_2
		v.reset(OpARMMOVWload)
		v.AuxInt = int64(uint32(c) >> uint64(d))
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVWreg(v *Value) bool {
	v_0 := v.Args[0]
	// match: (MOVWreg x)
	// cond: x.Uses == 1
	// result: (MOVWnop x)
	for {
		x := v_0
		if !(x.Uses == 1) {
			break
		}
		v.reset(OpARMMOVWnop)
		v.AddArg(x)
		return true
	}
	// match: (MOVWreg (MOVWconst [c]))
	// result: (MOVWconst [c])
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = c
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVWstore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWstore [off1] {sym} (ADDconst [off2] ptr) val mem)
	// result: (MOVWstore [off1+off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpARMADDconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		v.reset(OpARMMOVWstore)
		v.AuxInt = off1 + off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off1] {sym} (SUBconst [off2] ptr) val mem)
	// result: (MOVWstore [off1-off2] {sym} ptr val mem)
	for {
		off1 := v.AuxInt
		sym := v.Aux
		if v_0.Op != OpARMSUBconst {
			break
		}
		off2 := v_0.AuxInt
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		v.reset(OpARMMOVWstore)
		v.AuxInt = off1 - off2
		v.Aux = sym
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [off1] {sym1} (MOVWaddr [off2] {sym2} ptr) val mem)
	// cond: canMergeSym(sym1,sym2)
	// result: (MOVWstore [off1+off2] {mergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := v.AuxInt
		sym1 := v.Aux
		if v_0.Op != OpARMMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym2 := v_0.Aux
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(canMergeSym(sym1, sym2)) {
			break
		}
		v.reset(OpARMMOVWstore)
		v.AuxInt = off1 + off2
		v.Aux = mergeSym(sym1, sym2)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [0] {sym} (ADD ptr idx) val mem)
	// cond: sym == nil
	// result: (MOVWstoreidx ptr idx val mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		if v_0.Op != OpARMADD {
			break
		}
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(sym == nil) {
			break
		}
		v.reset(OpARMMOVWstoreidx)
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [0] {sym} (ADDshiftLL ptr idx [c]) val mem)
	// cond: sym == nil
	// result: (MOVWstoreshiftLL ptr idx [c] val mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		if v_0.Op != OpARMADDshiftLL {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(sym == nil) {
			break
		}
		v.reset(OpARMMOVWstoreshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [0] {sym} (ADDshiftRL ptr idx [c]) val mem)
	// cond: sym == nil
	// result: (MOVWstoreshiftRL ptr idx [c] val mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		if v_0.Op != OpARMADDshiftRL {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(sym == nil) {
			break
		}
		v.reset(OpARMMOVWstoreshiftRL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstore [0] {sym} (ADDshiftRA ptr idx [c]) val mem)
	// cond: sym == nil
	// result: (MOVWstoreshiftRA ptr idx [c] val mem)
	for {
		if v.AuxInt != 0 {
			break
		}
		sym := v.Aux
		if v_0.Op != OpARMADDshiftRA {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[1]
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(sym == nil) {
			break
		}
		v.reset(OpARMMOVWstoreshiftRA)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVWstoreidx(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWstoreidx ptr (MOVWconst [c]) val mem)
	// result: (MOVWstore [c] ptr val mem)
	for {
		ptr := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		val := v_2
		mem := v_3
		v.reset(OpARMMOVWstore)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreidx (MOVWconst [c]) ptr val mem)
	// result: (MOVWstore [c] ptr val mem)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		ptr := v_1
		val := v_2
		mem := v_3
		v.reset(OpARMMOVWstore)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreidx ptr (SLLconst idx [c]) val mem)
	// result: (MOVWstoreshiftLL ptr idx [c] val mem)
	for {
		ptr := v_0
		if v_1.Op != OpARMSLLconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		val := v_2
		mem := v_3
		v.reset(OpARMMOVWstoreshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreidx (SLLconst idx [c]) ptr val mem)
	// result: (MOVWstoreshiftLL ptr idx [c] val mem)
	for {
		if v_0.Op != OpARMSLLconst {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v_1
		val := v_2
		mem := v_3
		v.reset(OpARMMOVWstoreshiftLL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreidx ptr (SRLconst idx [c]) val mem)
	// result: (MOVWstoreshiftRL ptr idx [c] val mem)
	for {
		ptr := v_0
		if v_1.Op != OpARMSRLconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		val := v_2
		mem := v_3
		v.reset(OpARMMOVWstoreshiftRL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreidx (SRLconst idx [c]) ptr val mem)
	// result: (MOVWstoreshiftRL ptr idx [c] val mem)
	for {
		if v_0.Op != OpARMSRLconst {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v_1
		val := v_2
		mem := v_3
		v.reset(OpARMMOVWstoreshiftRL)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreidx ptr (SRAconst idx [c]) val mem)
	// result: (MOVWstoreshiftRA ptr idx [c] val mem)
	for {
		ptr := v_0
		if v_1.Op != OpARMSRAconst {
			break
		}
		c := v_1.AuxInt
		idx := v_1.Args[0]
		val := v_2
		mem := v_3
		v.reset(OpARMMOVWstoreshiftRA)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	// match: (MOVWstoreidx (SRAconst idx [c]) ptr val mem)
	// result: (MOVWstoreshiftRA ptr idx [c] val mem)
	for {
		if v_0.Op != OpARMSRAconst {
			break
		}
		c := v_0.AuxInt
		idx := v_0.Args[0]
		ptr := v_1
		val := v_2
		mem := v_3
		v.reset(OpARMMOVWstoreshiftRA)
		v.AuxInt = c
		v.AddArg(ptr)
		v.AddArg(idx)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVWstoreshiftLL(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWstoreshiftLL ptr (MOVWconst [c]) [d] val mem)
	// result: (MOVWstore [int64(uint32(c)<<uint64(d))] ptr val mem)
	for {
		d := v.AuxInt
		ptr := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		val := v_2
		mem := v_3
		v.reset(OpARMMOVWstore)
		v.AuxInt = int64(uint32(c) << uint64(d))
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVWstoreshiftRA(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWstoreshiftRA ptr (MOVWconst [c]) [d] val mem)
	// result: (MOVWstore [int64(int32(c)>>uint64(d))] ptr val mem)
	for {
		d := v.AuxInt
		ptr := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		val := v_2
		mem := v_3
		v.reset(OpARMMOVWstore)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMOVWstoreshiftRL(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWstoreshiftRL ptr (MOVWconst [c]) [d] val mem)
	// result: (MOVWstore [int64(uint32(c)>>uint64(d))] ptr val mem)
	for {
		d := v.AuxInt
		ptr := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		val := v_2
		mem := v_3
		v.reset(OpARMMOVWstore)
		v.AuxInt = int64(uint32(c) >> uint64(d))
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMUL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (MUL x (MOVWconst [c]))
	// cond: int32(c) == -1
	// result: (RSBconst [0] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMMOVWconst {
				continue
			}
			c := v_1.AuxInt
			if !(int32(c) == -1) {
				continue
			}
			v.reset(OpARMRSBconst)
			v.AuxInt = 0
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (MUL _ (MOVWconst [0]))
	// result: (MOVWconst [0])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_1.Op != OpARMMOVWconst || v_1.AuxInt != 0 {
				continue
			}
			v.reset(OpARMMOVWconst)
			v.AuxInt = 0
			return true
		}
		break
	}
	// match: (MUL x (MOVWconst [1]))
	// result: x
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMMOVWconst || v_1.AuxInt != 1 {
				continue
			}
			v.reset(OpCopy)
			v.Type = x.Type
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (MUL x (MOVWconst [c]))
	// cond: isPowerOfTwo(c)
	// result: (SLLconst [log2(c)] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMMOVWconst {
				continue
			}
			c := v_1.AuxInt
			if !(isPowerOfTwo(c)) {
				continue
			}
			v.reset(OpARMSLLconst)
			v.AuxInt = log2(c)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (MUL x (MOVWconst [c]))
	// cond: isPowerOfTwo(c-1) && int32(c) >= 3
	// result: (ADDshiftLL x x [log2(c-1)])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMMOVWconst {
				continue
			}
			c := v_1.AuxInt
			if !(isPowerOfTwo(c-1) && int32(c) >= 3) {
				continue
			}
			v.reset(OpARMADDshiftLL)
			v.AuxInt = log2(c - 1)
			v.AddArg(x)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (MUL x (MOVWconst [c]))
	// cond: isPowerOfTwo(c+1) && int32(c) >= 7
	// result: (RSBshiftLL x x [log2(c+1)])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMMOVWconst {
				continue
			}
			c := v_1.AuxInt
			if !(isPowerOfTwo(c+1) && int32(c) >= 7) {
				continue
			}
			v.reset(OpARMRSBshiftLL)
			v.AuxInt = log2(c + 1)
			v.AddArg(x)
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (MUL x (MOVWconst [c]))
	// cond: c%3 == 0 && isPowerOfTwo(c/3) && is32Bit(c)
	// result: (SLLconst [log2(c/3)] (ADDshiftLL <x.Type> x x [1]))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMMOVWconst {
				continue
			}
			c := v_1.AuxInt
			if !(c%3 == 0 && isPowerOfTwo(c/3) && is32Bit(c)) {
				continue
			}
			v.reset(OpARMSLLconst)
			v.AuxInt = log2(c / 3)
			v0 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
			v0.AuxInt = 1
			v0.AddArg(x)
			v0.AddArg(x)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (MUL x (MOVWconst [c]))
	// cond: c%5 == 0 && isPowerOfTwo(c/5) && is32Bit(c)
	// result: (SLLconst [log2(c/5)] (ADDshiftLL <x.Type> x x [2]))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMMOVWconst {
				continue
			}
			c := v_1.AuxInt
			if !(c%5 == 0 && isPowerOfTwo(c/5) && is32Bit(c)) {
				continue
			}
			v.reset(OpARMSLLconst)
			v.AuxInt = log2(c / 5)
			v0 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
			v0.AuxInt = 2
			v0.AddArg(x)
			v0.AddArg(x)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (MUL x (MOVWconst [c]))
	// cond: c%7 == 0 && isPowerOfTwo(c/7) && is32Bit(c)
	// result: (SLLconst [log2(c/7)] (RSBshiftLL <x.Type> x x [3]))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMMOVWconst {
				continue
			}
			c := v_1.AuxInt
			if !(c%7 == 0 && isPowerOfTwo(c/7) && is32Bit(c)) {
				continue
			}
			v.reset(OpARMSLLconst)
			v.AuxInt = log2(c / 7)
			v0 := b.NewValue0(v.Pos, OpARMRSBshiftLL, x.Type)
			v0.AuxInt = 3
			v0.AddArg(x)
			v0.AddArg(x)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (MUL x (MOVWconst [c]))
	// cond: c%9 == 0 && isPowerOfTwo(c/9) && is32Bit(c)
	// result: (SLLconst [log2(c/9)] (ADDshiftLL <x.Type> x x [3]))
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMMOVWconst {
				continue
			}
			c := v_1.AuxInt
			if !(c%9 == 0 && isPowerOfTwo(c/9) && is32Bit(c)) {
				continue
			}
			v.reset(OpARMSLLconst)
			v.AuxInt = log2(c / 9)
			v0 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
			v0.AuxInt = 3
			v0.AddArg(x)
			v0.AddArg(x)
			v.AddArg(v0)
			return true
		}
		break
	}
	// match: (MUL (MOVWconst [c]) (MOVWconst [d]))
	// result: (MOVWconst [int64(int32(c*d))])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpARMMOVWconst {
				continue
			}
			c := v_0.AuxInt
			if v_1.Op != OpARMMOVWconst {
				continue
			}
			d := v_1.AuxInt
			v.reset(OpARMMOVWconst)
			v.AuxInt = int64(int32(c * d))
			return true
		}
		break
	}
	return false
}
func rewriteValueARM_OpARMMULA(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (MULA x (MOVWconst [c]) a)
	// cond: int32(c) == -1
	// result: (SUB a x)
	for {
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		a := v_2
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpARMSUB)
		v.AddArg(a)
		v.AddArg(x)
		return true
	}
	// match: (MULA _ (MOVWconst [0]) a)
	// result: a
	for {
		if v_1.Op != OpARMMOVWconst || v_1.AuxInt != 0 {
			break
		}
		a := v_2
		v.reset(OpCopy)
		v.Type = a.Type
		v.AddArg(a)
		return true
	}
	// match: (MULA x (MOVWconst [1]) a)
	// result: (ADD x a)
	for {
		x := v_0
		if v_1.Op != OpARMMOVWconst || v_1.AuxInt != 1 {
			break
		}
		a := v_2
		v.reset(OpARMADD)
		v.AddArg(x)
		v.AddArg(a)
		return true
	}
	// match: (MULA x (MOVWconst [c]) a)
	// cond: isPowerOfTwo(c)
	// result: (ADD (SLLconst <x.Type> [log2(c)] x) a)
	for {
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		a := v_2
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpARMADD)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULA x (MOVWconst [c]) a)
	// cond: isPowerOfTwo(c-1) && int32(c) >= 3
	// result: (ADD (ADDshiftLL <x.Type> x x [log2(c-1)]) a)
	for {
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		a := v_2
		if !(isPowerOfTwo(c-1) && int32(c) >= 3) {
			break
		}
		v.reset(OpARMADD)
		v0 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v0.AuxInt = log2(c - 1)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULA x (MOVWconst [c]) a)
	// cond: isPowerOfTwo(c+1) && int32(c) >= 7
	// result: (ADD (RSBshiftLL <x.Type> x x [log2(c+1)]) a)
	for {
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		a := v_2
		if !(isPowerOfTwo(c+1) && int32(c) >= 7) {
			break
		}
		v.reset(OpARMADD)
		v0 := b.NewValue0(v.Pos, OpARMRSBshiftLL, x.Type)
		v0.AuxInt = log2(c + 1)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULA x (MOVWconst [c]) a)
	// cond: c%3 == 0 && isPowerOfTwo(c/3) && is32Bit(c)
	// result: (ADD (SLLconst <x.Type> [log2(c/3)] (ADDshiftLL <x.Type> x x [1])) a)
	for {
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		a := v_2
		if !(c%3 == 0 && isPowerOfTwo(c/3) && is32Bit(c)) {
			break
		}
		v.reset(OpARMADD)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c / 3)
		v1 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v1.AuxInt = 1
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULA x (MOVWconst [c]) a)
	// cond: c%5 == 0 && isPowerOfTwo(c/5) && is32Bit(c)
	// result: (ADD (SLLconst <x.Type> [log2(c/5)] (ADDshiftLL <x.Type> x x [2])) a)
	for {
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		a := v_2
		if !(c%5 == 0 && isPowerOfTwo(c/5) && is32Bit(c)) {
			break
		}
		v.reset(OpARMADD)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c / 5)
		v1 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v1.AuxInt = 2
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULA x (MOVWconst [c]) a)
	// cond: c%7 == 0 && isPowerOfTwo(c/7) && is32Bit(c)
	// result: (ADD (SLLconst <x.Type> [log2(c/7)] (RSBshiftLL <x.Type> x x [3])) a)
	for {
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		a := v_2
		if !(c%7 == 0 && isPowerOfTwo(c/7) && is32Bit(c)) {
			break
		}
		v.reset(OpARMADD)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c / 7)
		v1 := b.NewValue0(v.Pos, OpARMRSBshiftLL, x.Type)
		v1.AuxInt = 3
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULA x (MOVWconst [c]) a)
	// cond: c%9 == 0 && isPowerOfTwo(c/9) && is32Bit(c)
	// result: (ADD (SLLconst <x.Type> [log2(c/9)] (ADDshiftLL <x.Type> x x [3])) a)
	for {
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		a := v_2
		if !(c%9 == 0 && isPowerOfTwo(c/9) && is32Bit(c)) {
			break
		}
		v.reset(OpARMADD)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c / 9)
		v1 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v1.AuxInt = 3
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULA (MOVWconst [c]) x a)
	// cond: int32(c) == -1
	// result: (SUB a x)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		a := v_2
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpARMSUB)
		v.AddArg(a)
		v.AddArg(x)
		return true
	}
	// match: (MULA (MOVWconst [0]) _ a)
	// result: a
	for {
		if v_0.Op != OpARMMOVWconst || v_0.AuxInt != 0 {
			break
		}
		a := v_2
		v.reset(OpCopy)
		v.Type = a.Type
		v.AddArg(a)
		return true
	}
	// match: (MULA (MOVWconst [1]) x a)
	// result: (ADD x a)
	for {
		if v_0.Op != OpARMMOVWconst || v_0.AuxInt != 1 {
			break
		}
		x := v_1
		a := v_2
		v.reset(OpARMADD)
		v.AddArg(x)
		v.AddArg(a)
		return true
	}
	// match: (MULA (MOVWconst [c]) x a)
	// cond: isPowerOfTwo(c)
	// result: (ADD (SLLconst <x.Type> [log2(c)] x) a)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		a := v_2
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpARMADD)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULA (MOVWconst [c]) x a)
	// cond: isPowerOfTwo(c-1) && int32(c) >= 3
	// result: (ADD (ADDshiftLL <x.Type> x x [log2(c-1)]) a)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		a := v_2
		if !(isPowerOfTwo(c-1) && int32(c) >= 3) {
			break
		}
		v.reset(OpARMADD)
		v0 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v0.AuxInt = log2(c - 1)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULA (MOVWconst [c]) x a)
	// cond: isPowerOfTwo(c+1) && int32(c) >= 7
	// result: (ADD (RSBshiftLL <x.Type> x x [log2(c+1)]) a)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		a := v_2
		if !(isPowerOfTwo(c+1) && int32(c) >= 7) {
			break
		}
		v.reset(OpARMADD)
		v0 := b.NewValue0(v.Pos, OpARMRSBshiftLL, x.Type)
		v0.AuxInt = log2(c + 1)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULA (MOVWconst [c]) x a)
	// cond: c%3 == 0 && isPowerOfTwo(c/3) && is32Bit(c)
	// result: (ADD (SLLconst <x.Type> [log2(c/3)] (ADDshiftLL <x.Type> x x [1])) a)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		a := v_2
		if !(c%3 == 0 && isPowerOfTwo(c/3) && is32Bit(c)) {
			break
		}
		v.reset(OpARMADD)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c / 3)
		v1 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v1.AuxInt = 1
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULA (MOVWconst [c]) x a)
	// cond: c%5 == 0 && isPowerOfTwo(c/5) && is32Bit(c)
	// result: (ADD (SLLconst <x.Type> [log2(c/5)] (ADDshiftLL <x.Type> x x [2])) a)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		a := v_2
		if !(c%5 == 0 && isPowerOfTwo(c/5) && is32Bit(c)) {
			break
		}
		v.reset(OpARMADD)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c / 5)
		v1 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v1.AuxInt = 2
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULA (MOVWconst [c]) x a)
	// cond: c%7 == 0 && isPowerOfTwo(c/7) && is32Bit(c)
	// result: (ADD (SLLconst <x.Type> [log2(c/7)] (RSBshiftLL <x.Type> x x [3])) a)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		a := v_2
		if !(c%7 == 0 && isPowerOfTwo(c/7) && is32Bit(c)) {
			break
		}
		v.reset(OpARMADD)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c / 7)
		v1 := b.NewValue0(v.Pos, OpARMRSBshiftLL, x.Type)
		v1.AuxInt = 3
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULA (MOVWconst [c]) x a)
	// cond: c%9 == 0 && isPowerOfTwo(c/9) && is32Bit(c)
	// result: (ADD (SLLconst <x.Type> [log2(c/9)] (ADDshiftLL <x.Type> x x [3])) a)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		a := v_2
		if !(c%9 == 0 && isPowerOfTwo(c/9) && is32Bit(c)) {
			break
		}
		v.reset(OpARMADD)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c / 9)
		v1 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v1.AuxInt = 3
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULA (MOVWconst [c]) (MOVWconst [d]) a)
	// result: (ADDconst [int64(int32(c*d))] a)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpARMMOVWconst {
			break
		}
		d := v_1.AuxInt
		a := v_2
		v.reset(OpARMADDconst)
		v.AuxInt = int64(int32(c * d))
		v.AddArg(a)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMULD(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MULD (NEGD x) y)
	// cond: objabi.GOARM >= 6
	// result: (NMULD x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpARMNEGD {
				continue
			}
			x := v_0.Args[0]
			y := v_1
			if !(objabi.GOARM >= 6) {
				continue
			}
			v.reset(OpARMNMULD)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	return false
}
func rewriteValueARM_OpARMMULF(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MULF (NEGF x) y)
	// cond: objabi.GOARM >= 6
	// result: (NMULF x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpARMNEGF {
				continue
			}
			x := v_0.Args[0]
			y := v_1
			if !(objabi.GOARM >= 6) {
				continue
			}
			v.reset(OpARMNMULF)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	return false
}
func rewriteValueARM_OpARMMULS(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (MULS x (MOVWconst [c]) a)
	// cond: int32(c) == -1
	// result: (ADD a x)
	for {
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		a := v_2
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpARMADD)
		v.AddArg(a)
		v.AddArg(x)
		return true
	}
	// match: (MULS _ (MOVWconst [0]) a)
	// result: a
	for {
		if v_1.Op != OpARMMOVWconst || v_1.AuxInt != 0 {
			break
		}
		a := v_2
		v.reset(OpCopy)
		v.Type = a.Type
		v.AddArg(a)
		return true
	}
	// match: (MULS x (MOVWconst [1]) a)
	// result: (RSB x a)
	for {
		x := v_0
		if v_1.Op != OpARMMOVWconst || v_1.AuxInt != 1 {
			break
		}
		a := v_2
		v.reset(OpARMRSB)
		v.AddArg(x)
		v.AddArg(a)
		return true
	}
	// match: (MULS x (MOVWconst [c]) a)
	// cond: isPowerOfTwo(c)
	// result: (RSB (SLLconst <x.Type> [log2(c)] x) a)
	for {
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		a := v_2
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpARMRSB)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULS x (MOVWconst [c]) a)
	// cond: isPowerOfTwo(c-1) && int32(c) >= 3
	// result: (RSB (ADDshiftLL <x.Type> x x [log2(c-1)]) a)
	for {
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		a := v_2
		if !(isPowerOfTwo(c-1) && int32(c) >= 3) {
			break
		}
		v.reset(OpARMRSB)
		v0 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v0.AuxInt = log2(c - 1)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULS x (MOVWconst [c]) a)
	// cond: isPowerOfTwo(c+1) && int32(c) >= 7
	// result: (RSB (RSBshiftLL <x.Type> x x [log2(c+1)]) a)
	for {
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		a := v_2
		if !(isPowerOfTwo(c+1) && int32(c) >= 7) {
			break
		}
		v.reset(OpARMRSB)
		v0 := b.NewValue0(v.Pos, OpARMRSBshiftLL, x.Type)
		v0.AuxInt = log2(c + 1)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULS x (MOVWconst [c]) a)
	// cond: c%3 == 0 && isPowerOfTwo(c/3) && is32Bit(c)
	// result: (RSB (SLLconst <x.Type> [log2(c/3)] (ADDshiftLL <x.Type> x x [1])) a)
	for {
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		a := v_2
		if !(c%3 == 0 && isPowerOfTwo(c/3) && is32Bit(c)) {
			break
		}
		v.reset(OpARMRSB)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c / 3)
		v1 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v1.AuxInt = 1
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULS x (MOVWconst [c]) a)
	// cond: c%5 == 0 && isPowerOfTwo(c/5) && is32Bit(c)
	// result: (RSB (SLLconst <x.Type> [log2(c/5)] (ADDshiftLL <x.Type> x x [2])) a)
	for {
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		a := v_2
		if !(c%5 == 0 && isPowerOfTwo(c/5) && is32Bit(c)) {
			break
		}
		v.reset(OpARMRSB)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c / 5)
		v1 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v1.AuxInt = 2
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULS x (MOVWconst [c]) a)
	// cond: c%7 == 0 && isPowerOfTwo(c/7) && is32Bit(c)
	// result: (RSB (SLLconst <x.Type> [log2(c/7)] (RSBshiftLL <x.Type> x x [3])) a)
	for {
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		a := v_2
		if !(c%7 == 0 && isPowerOfTwo(c/7) && is32Bit(c)) {
			break
		}
		v.reset(OpARMRSB)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c / 7)
		v1 := b.NewValue0(v.Pos, OpARMRSBshiftLL, x.Type)
		v1.AuxInt = 3
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULS x (MOVWconst [c]) a)
	// cond: c%9 == 0 && isPowerOfTwo(c/9) && is32Bit(c)
	// result: (RSB (SLLconst <x.Type> [log2(c/9)] (ADDshiftLL <x.Type> x x [3])) a)
	for {
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		a := v_2
		if !(c%9 == 0 && isPowerOfTwo(c/9) && is32Bit(c)) {
			break
		}
		v.reset(OpARMRSB)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c / 9)
		v1 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v1.AuxInt = 3
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULS (MOVWconst [c]) x a)
	// cond: int32(c) == -1
	// result: (ADD a x)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		a := v_2
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpARMADD)
		v.AddArg(a)
		v.AddArg(x)
		return true
	}
	// match: (MULS (MOVWconst [0]) _ a)
	// result: a
	for {
		if v_0.Op != OpARMMOVWconst || v_0.AuxInt != 0 {
			break
		}
		a := v_2
		v.reset(OpCopy)
		v.Type = a.Type
		v.AddArg(a)
		return true
	}
	// match: (MULS (MOVWconst [1]) x a)
	// result: (RSB x a)
	for {
		if v_0.Op != OpARMMOVWconst || v_0.AuxInt != 1 {
			break
		}
		x := v_1
		a := v_2
		v.reset(OpARMRSB)
		v.AddArg(x)
		v.AddArg(a)
		return true
	}
	// match: (MULS (MOVWconst [c]) x a)
	// cond: isPowerOfTwo(c)
	// result: (RSB (SLLconst <x.Type> [log2(c)] x) a)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		a := v_2
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpARMRSB)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULS (MOVWconst [c]) x a)
	// cond: isPowerOfTwo(c-1) && int32(c) >= 3
	// result: (RSB (ADDshiftLL <x.Type> x x [log2(c-1)]) a)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		a := v_2
		if !(isPowerOfTwo(c-1) && int32(c) >= 3) {
			break
		}
		v.reset(OpARMRSB)
		v0 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v0.AuxInt = log2(c - 1)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULS (MOVWconst [c]) x a)
	// cond: isPowerOfTwo(c+1) && int32(c) >= 7
	// result: (RSB (RSBshiftLL <x.Type> x x [log2(c+1)]) a)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		a := v_2
		if !(isPowerOfTwo(c+1) && int32(c) >= 7) {
			break
		}
		v.reset(OpARMRSB)
		v0 := b.NewValue0(v.Pos, OpARMRSBshiftLL, x.Type)
		v0.AuxInt = log2(c + 1)
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULS (MOVWconst [c]) x a)
	// cond: c%3 == 0 && isPowerOfTwo(c/3) && is32Bit(c)
	// result: (RSB (SLLconst <x.Type> [log2(c/3)] (ADDshiftLL <x.Type> x x [1])) a)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		a := v_2
		if !(c%3 == 0 && isPowerOfTwo(c/3) && is32Bit(c)) {
			break
		}
		v.reset(OpARMRSB)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c / 3)
		v1 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v1.AuxInt = 1
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULS (MOVWconst [c]) x a)
	// cond: c%5 == 0 && isPowerOfTwo(c/5) && is32Bit(c)
	// result: (RSB (SLLconst <x.Type> [log2(c/5)] (ADDshiftLL <x.Type> x x [2])) a)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		a := v_2
		if !(c%5 == 0 && isPowerOfTwo(c/5) && is32Bit(c)) {
			break
		}
		v.reset(OpARMRSB)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c / 5)
		v1 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v1.AuxInt = 2
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULS (MOVWconst [c]) x a)
	// cond: c%7 == 0 && isPowerOfTwo(c/7) && is32Bit(c)
	// result: (RSB (SLLconst <x.Type> [log2(c/7)] (RSBshiftLL <x.Type> x x [3])) a)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		a := v_2
		if !(c%7 == 0 && isPowerOfTwo(c/7) && is32Bit(c)) {
			break
		}
		v.reset(OpARMRSB)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c / 7)
		v1 := b.NewValue0(v.Pos, OpARMRSBshiftLL, x.Type)
		v1.AuxInt = 3
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULS (MOVWconst [c]) x a)
	// cond: c%9 == 0 && isPowerOfTwo(c/9) && is32Bit(c)
	// result: (RSB (SLLconst <x.Type> [log2(c/9)] (ADDshiftLL <x.Type> x x [3])) a)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		a := v_2
		if !(c%9 == 0 && isPowerOfTwo(c/9) && is32Bit(c)) {
			break
		}
		v.reset(OpARMRSB)
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = log2(c / 9)
		v1 := b.NewValue0(v.Pos, OpARMADDshiftLL, x.Type)
		v1.AuxInt = 3
		v1.AddArg(x)
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(a)
		return true
	}
	// match: (MULS (MOVWconst [c]) (MOVWconst [d]) a)
	// result: (SUBconst [int64(int32(c*d))] a)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		if v_1.Op != OpARMMOVWconst {
			break
		}
		d := v_1.AuxInt
		a := v_2
		v.reset(OpARMSUBconst)
		v.AuxInt = int64(int32(c * d))
		v.AddArg(a)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMVN(v *Value) bool {
	v_0 := v.Args[0]
	// match: (MVN (MOVWconst [c]))
	// result: (MOVWconst [^c])
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = ^c
		return true
	}
	// match: (MVN (SLLconst [c] x))
	// result: (MVNshiftLL x [c])
	for {
		if v_0.Op != OpARMSLLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMMVNshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (MVN (SRLconst [c] x))
	// result: (MVNshiftRL x [c])
	for {
		if v_0.Op != OpARMSRLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMMVNshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (MVN (SRAconst [c] x))
	// result: (MVNshiftRA x [c])
	for {
		if v_0.Op != OpARMSRAconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMMVNshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (MVN (SLL x y))
	// result: (MVNshiftLLreg x y)
	for {
		if v_0.Op != OpARMSLL {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpARMMVNshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (MVN (SRL x y))
	// result: (MVNshiftRLreg x y)
	for {
		if v_0.Op != OpARMSRL {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpARMMVNshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (MVN (SRA x y))
	// result: (MVNshiftRAreg x y)
	for {
		if v_0.Op != OpARMSRA {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		v.reset(OpARMMVNshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMVNshiftLL(v *Value) bool {
	v_0 := v.Args[0]
	// match: (MVNshiftLL (MOVWconst [c]) [d])
	// result: (MOVWconst [^int64(uint32(c)<<uint64(d))])
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = ^int64(uint32(c) << uint64(d))
		return true
	}
	return false
}
func rewriteValueARM_OpARMMVNshiftLLreg(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MVNshiftLLreg x (MOVWconst [c]))
	// result: (MVNshiftLL x [c])
	for {
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMMVNshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMVNshiftRA(v *Value) bool {
	v_0 := v.Args[0]
	// match: (MVNshiftRA (MOVWconst [c]) [d])
	// result: (MOVWconst [^int64(int32(c)>>uint64(d))])
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = ^int64(int32(c) >> uint64(d))
		return true
	}
	return false
}
func rewriteValueARM_OpARMMVNshiftRAreg(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MVNshiftRAreg x (MOVWconst [c]))
	// result: (MVNshiftRA x [c])
	for {
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMMVNshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMMVNshiftRL(v *Value) bool {
	v_0 := v.Args[0]
	// match: (MVNshiftRL (MOVWconst [c]) [d])
	// result: (MOVWconst [^int64(uint32(c)>>uint64(d))])
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = ^int64(uint32(c) >> uint64(d))
		return true
	}
	return false
}
func rewriteValueARM_OpARMMVNshiftRLreg(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MVNshiftRLreg x (MOVWconst [c]))
	// result: (MVNshiftRL x [c])
	for {
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMMVNshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMNEGD(v *Value) bool {
	v_0 := v.Args[0]
	// match: (NEGD (MULD x y))
	// cond: objabi.GOARM >= 6
	// result: (NMULD x y)
	for {
		if v_0.Op != OpARMMULD {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		if !(objabi.GOARM >= 6) {
			break
		}
		v.reset(OpARMNMULD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMNEGF(v *Value) bool {
	v_0 := v.Args[0]
	// match: (NEGF (MULF x y))
	// cond: objabi.GOARM >= 6
	// result: (NMULF x y)
	for {
		if v_0.Op != OpARMMULF {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		if !(objabi.GOARM >= 6) {
			break
		}
		v.reset(OpARMNMULF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMNMULD(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (NMULD (NEGD x) y)
	// result: (MULD x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpARMNEGD {
				continue
			}
			x := v_0.Args[0]
			y := v_1
			v.reset(OpARMMULD)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	return false
}
func rewriteValueARM_OpARMNMULF(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (NMULF (NEGF x) y)
	// result: (MULF x y)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpARMNEGF {
				continue
			}
			x := v_0.Args[0]
			y := v_1
			v.reset(OpARMMULF)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	return false
}
func rewriteValueARM_OpARMNotEqual(v *Value) bool {
	v_0 := v.Args[0]
	// match: (NotEqual (FlagEQ))
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpARMFlagEQ {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (NotEqual (FlagLT_ULT))
	// result: (MOVWconst [1])
	for {
		if v_0.Op != OpARMFlagLT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (NotEqual (FlagLT_UGT))
	// result: (MOVWconst [1])
	for {
		if v_0.Op != OpARMFlagLT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (NotEqual (FlagGT_ULT))
	// result: (MOVWconst [1])
	for {
		if v_0.Op != OpARMFlagGT_ULT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (NotEqual (FlagGT_UGT))
	// result: (MOVWconst [1])
	for {
		if v_0.Op != OpARMFlagGT_UGT {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 1
		return true
	}
	// match: (NotEqual (InvertFlags x))
	// result: (NotEqual x)
	for {
		if v_0.Op != OpARMInvertFlags {
			break
		}
		x := v_0.Args[0]
		v.reset(OpARMNotEqual)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMOR(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (OR x (MOVWconst [c]))
	// result: (ORconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMMOVWconst {
				continue
			}
			c := v_1.AuxInt
			v.reset(OpARMORconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (OR x (SLLconst [c] y))
	// result: (ORshiftLL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSLLconst {
				continue
			}
			c := v_1.AuxInt
			y := v_1.Args[0]
			v.reset(OpARMORshiftLL)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (OR x (SRLconst [c] y))
	// result: (ORshiftRL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRLconst {
				continue
			}
			c := v_1.AuxInt
			y := v_1.Args[0]
			v.reset(OpARMORshiftRL)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (OR x (SRAconst [c] y))
	// result: (ORshiftRA x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRAconst {
				continue
			}
			c := v_1.AuxInt
			y := v_1.Args[0]
			v.reset(OpARMORshiftRA)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (OR x (SLL y z))
	// result: (ORshiftLLreg x y z)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSLL {
				continue
			}
			z := v_1.Args[1]
			y := v_1.Args[0]
			v.reset(OpARMORshiftLLreg)
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(z)
			return true
		}
		break
	}
	// match: (OR x (SRL y z))
	// result: (ORshiftRLreg x y z)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRL {
				continue
			}
			z := v_1.Args[1]
			y := v_1.Args[0]
			v.reset(OpARMORshiftRLreg)
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(z)
			return true
		}
		break
	}
	// match: (OR x (SRA y z))
	// result: (ORshiftRAreg x y z)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRA {
				continue
			}
			z := v_1.Args[1]
			y := v_1.Args[0]
			v.reset(OpARMORshiftRAreg)
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(z)
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
func rewriteValueARM_OpARMORconst(v *Value) bool {
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
	// match: (ORconst [c] _)
	// cond: int32(c)==-1
	// result: (MOVWconst [-1])
	for {
		c := v.AuxInt
		if !(int32(c) == -1) {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = -1
		return true
	}
	// match: (ORconst [c] (MOVWconst [d]))
	// result: (MOVWconst [c|d])
	for {
		c := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = c | d
		return true
	}
	// match: (ORconst [c] (ORconst [d] x))
	// result: (ORconst [c|d] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpARMORconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMORconst)
		v.AuxInt = c | d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMORshiftLL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ORshiftLL (MOVWconst [c]) x [d])
	// result: (ORconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ORshiftLL x (MOVWconst [c]) [d])
	// result: (ORconst x [int64(int32(uint32(c)<<uint64(d)))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMORconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: ( ORshiftLL [c] (SRLconst x [32-c]) x)
	// result: (SRRconst [32-c] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpARMSRLconst || v_0.AuxInt != 32-c {
			break
		}
		x := v_0.Args[0]
		if x != v_1 {
			break
		}
		v.reset(OpARMSRRconst)
		v.AuxInt = 32 - c
		v.AddArg(x)
		return true
	}
	// match: (ORshiftLL <typ.UInt16> [8] (BFXU <typ.UInt16> [armBFAuxInt(8, 8)] x) x)
	// result: (REV16 x)
	for {
		if v.Type != typ.UInt16 || v.AuxInt != 8 || v_0.Op != OpARMBFXU || v_0.Type != typ.UInt16 || v_0.AuxInt != armBFAuxInt(8, 8) {
			break
		}
		x := v_0.Args[0]
		if x != v_1 {
			break
		}
		v.reset(OpARMREV16)
		v.AddArg(x)
		return true
	}
	// match: (ORshiftLL <typ.UInt16> [8] (SRLconst <typ.UInt16> [24] (SLLconst [16] x)) x)
	// cond: objabi.GOARM>=6
	// result: (REV16 x)
	for {
		if v.Type != typ.UInt16 || v.AuxInt != 8 || v_0.Op != OpARMSRLconst || v_0.Type != typ.UInt16 || v_0.AuxInt != 24 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpARMSLLconst || v_0_0.AuxInt != 16 {
			break
		}
		x := v_0_0.Args[0]
		if x != v_1 || !(objabi.GOARM >= 6) {
			break
		}
		v.reset(OpARMREV16)
		v.AddArg(x)
		return true
	}
	// match: (ORshiftLL x y:(SLLconst x [c]) [d])
	// cond: c==d
	// result: y
	for {
		d := v.AuxInt
		x := v_0
		y := v_1
		if y.Op != OpARMSLLconst {
			break
		}
		c := y.AuxInt
		if x != y.Args[0] || !(c == d) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMORshiftLLreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ORshiftLLreg (MOVWconst [c]) x y)
	// result: (ORconst [c] (SLL <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (ORshiftLLreg x y (MOVWconst [c]))
	// result: (ORshiftLL x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMORshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMORshiftRA(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ORshiftRA (MOVWconst [c]) x [d])
	// result: (ORconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ORshiftRA x (MOVWconst [c]) [d])
	// result: (ORconst x [int64(int32(c)>>uint64(d))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMORconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	// match: (ORshiftRA x y:(SRAconst x [c]) [d])
	// cond: c==d
	// result: y
	for {
		d := v.AuxInt
		x := v_0
		y := v_1
		if y.Op != OpARMSRAconst {
			break
		}
		c := y.AuxInt
		if x != y.Args[0] || !(c == d) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMORshiftRAreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ORshiftRAreg (MOVWconst [c]) x y)
	// result: (ORconst [c] (SRA <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRA, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (ORshiftRAreg x y (MOVWconst [c]))
	// result: (ORshiftRA x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMORshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMORshiftRL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ORshiftRL (MOVWconst [c]) x [d])
	// result: (ORconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (ORshiftRL x (MOVWconst [c]) [d])
	// result: (ORconst x [int64(int32(uint32(c)>>uint64(d)))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMORconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: ( ORshiftRL [c] (SLLconst x [32-c]) x)
	// result: (SRRconst [ c] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpARMSLLconst || v_0.AuxInt != 32-c {
			break
		}
		x := v_0.Args[0]
		if x != v_1 {
			break
		}
		v.reset(OpARMSRRconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (ORshiftRL x y:(SRLconst x [c]) [d])
	// cond: c==d
	// result: y
	for {
		d := v.AuxInt
		x := v_0
		y := v_1
		if y.Op != OpARMSRLconst {
			break
		}
		c := y.AuxInt
		if x != y.Args[0] || !(c == d) {
			break
		}
		v.reset(OpCopy)
		v.Type = y.Type
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMORshiftRLreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (ORshiftRLreg (MOVWconst [c]) x y)
	// result: (ORconst [c] (SRL <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (ORshiftRLreg x y (MOVWconst [c]))
	// result: (ORshiftRL x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMORshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSB(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (RSB (MOVWconst [c]) x)
	// result: (SUBconst [c] x)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMSUBconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (RSB x (MOVWconst [c]))
	// result: (RSBconst [c] x)
	for {
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMRSBconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (RSB x (SLLconst [c] y))
	// result: (RSBshiftLL x y [c])
	for {
		x := v_0
		if v_1.Op != OpARMSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMRSBshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (RSB (SLLconst [c] y) x)
	// result: (SUBshiftLL x y [c])
	for {
		if v_0.Op != OpARMSLLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v_1
		v.reset(OpARMSUBshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (RSB x (SRLconst [c] y))
	// result: (RSBshiftRL x y [c])
	for {
		x := v_0
		if v_1.Op != OpARMSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMRSBshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (RSB (SRLconst [c] y) x)
	// result: (SUBshiftRL x y [c])
	for {
		if v_0.Op != OpARMSRLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v_1
		v.reset(OpARMSUBshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (RSB x (SRAconst [c] y))
	// result: (RSBshiftRA x y [c])
	for {
		x := v_0
		if v_1.Op != OpARMSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMRSBshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (RSB (SRAconst [c] y) x)
	// result: (SUBshiftRA x y [c])
	for {
		if v_0.Op != OpARMSRAconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v_1
		v.reset(OpARMSUBshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (RSB x (SLL y z))
	// result: (RSBshiftLLreg x y z)
	for {
		x := v_0
		if v_1.Op != OpARMSLL {
			break
		}
		z := v_1.Args[1]
		y := v_1.Args[0]
		v.reset(OpARMRSBshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	// match: (RSB (SLL y z) x)
	// result: (SUBshiftLLreg x y z)
	for {
		if v_0.Op != OpARMSLL {
			break
		}
		z := v_0.Args[1]
		y := v_0.Args[0]
		x := v_1
		v.reset(OpARMSUBshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	// match: (RSB x (SRL y z))
	// result: (RSBshiftRLreg x y z)
	for {
		x := v_0
		if v_1.Op != OpARMSRL {
			break
		}
		z := v_1.Args[1]
		y := v_1.Args[0]
		v.reset(OpARMRSBshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	// match: (RSB (SRL y z) x)
	// result: (SUBshiftRLreg x y z)
	for {
		if v_0.Op != OpARMSRL {
			break
		}
		z := v_0.Args[1]
		y := v_0.Args[0]
		x := v_1
		v.reset(OpARMSUBshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	// match: (RSB x (SRA y z))
	// result: (RSBshiftRAreg x y z)
	for {
		x := v_0
		if v_1.Op != OpARMSRA {
			break
		}
		z := v_1.Args[1]
		y := v_1.Args[0]
		v.reset(OpARMRSBshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	// match: (RSB (SRA y z) x)
	// result: (SUBshiftRAreg x y z)
	for {
		if v_0.Op != OpARMSRA {
			break
		}
		z := v_0.Args[1]
		y := v_0.Args[0]
		x := v_1
		v.reset(OpARMSUBshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	// match: (RSB x x)
	// result: (MOVWconst [0])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (RSB (MUL x y) a)
	// cond: objabi.GOARM == 7
	// result: (MULS x y a)
	for {
		if v_0.Op != OpARMMUL {
			break
		}
		y := v_0.Args[1]
		x := v_0.Args[0]
		a := v_1
		if !(objabi.GOARM == 7) {
			break
		}
		v.reset(OpARMMULS)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(a)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSBSshiftLL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (RSBSshiftLL (MOVWconst [c]) x [d])
	// result: (SUBSconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMSUBSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (RSBSshiftLL x (MOVWconst [c]) [d])
	// result: (RSBSconst x [int64(int32(uint32(c)<<uint64(d)))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMRSBSconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSBSshiftLLreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (RSBSshiftLLreg (MOVWconst [c]) x y)
	// result: (SUBSconst [c] (SLL <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMSUBSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (RSBSshiftLLreg x y (MOVWconst [c]))
	// result: (RSBSshiftLL x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMRSBSshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSBSshiftRA(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (RSBSshiftRA (MOVWconst [c]) x [d])
	// result: (SUBSconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMSUBSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (RSBSshiftRA x (MOVWconst [c]) [d])
	// result: (RSBSconst x [int64(int32(c)>>uint64(d))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMRSBSconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSBSshiftRAreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (RSBSshiftRAreg (MOVWconst [c]) x y)
	// result: (SUBSconst [c] (SRA <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMSUBSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRA, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (RSBSshiftRAreg x y (MOVWconst [c]))
	// result: (RSBSshiftRA x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMRSBSshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSBSshiftRL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (RSBSshiftRL (MOVWconst [c]) x [d])
	// result: (SUBSconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMSUBSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (RSBSshiftRL x (MOVWconst [c]) [d])
	// result: (RSBSconst x [int64(int32(uint32(c)>>uint64(d)))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMRSBSconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSBSshiftRLreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (RSBSshiftRLreg (MOVWconst [c]) x y)
	// result: (SUBSconst [c] (SRL <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMSUBSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (RSBSshiftRLreg x y (MOVWconst [c]))
	// result: (RSBSshiftRL x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMRSBSshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSBconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (RSBconst [c] (MOVWconst [d]))
	// result: (MOVWconst [int64(int32(c-d))])
	for {
		c := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = int64(int32(c - d))
		return true
	}
	// match: (RSBconst [c] (RSBconst [d] x))
	// result: (ADDconst [int64(int32(c-d))] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpARMRSBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMADDconst)
		v.AuxInt = int64(int32(c - d))
		v.AddArg(x)
		return true
	}
	// match: (RSBconst [c] (ADDconst [d] x))
	// result: (RSBconst [int64(int32(c-d))] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpARMADDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMRSBconst)
		v.AuxInt = int64(int32(c - d))
		v.AddArg(x)
		return true
	}
	// match: (RSBconst [c] (SUBconst [d] x))
	// result: (RSBconst [int64(int32(c+d))] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpARMSUBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMRSBconst)
		v.AuxInt = int64(int32(c + d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSBshiftLL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (RSBshiftLL (MOVWconst [c]) x [d])
	// result: (SUBconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMSUBconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (RSBshiftLL x (MOVWconst [c]) [d])
	// result: (RSBconst x [int64(int32(uint32(c)<<uint64(d)))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMRSBconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: (RSBshiftLL x (SLLconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMSLLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] || !(c == d) {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSBshiftLLreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (RSBshiftLLreg (MOVWconst [c]) x y)
	// result: (SUBconst [c] (SLL <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMSUBconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (RSBshiftLLreg x y (MOVWconst [c]))
	// result: (RSBshiftLL x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMRSBshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSBshiftRA(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (RSBshiftRA (MOVWconst [c]) x [d])
	// result: (SUBconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMSUBconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (RSBshiftRA x (MOVWconst [c]) [d])
	// result: (RSBconst x [int64(int32(c)>>uint64(d))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMRSBconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	// match: (RSBshiftRA x (SRAconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMSRAconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] || !(c == d) {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSBshiftRAreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (RSBshiftRAreg (MOVWconst [c]) x y)
	// result: (SUBconst [c] (SRA <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMSUBconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRA, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (RSBshiftRAreg x y (MOVWconst [c]))
	// result: (RSBshiftRA x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMRSBshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSBshiftRL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (RSBshiftRL (MOVWconst [c]) x [d])
	// result: (SUBconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMSUBconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (RSBshiftRL x (MOVWconst [c]) [d])
	// result: (RSBconst x [int64(int32(uint32(c)>>uint64(d)))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMRSBconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: (RSBshiftRL x (SRLconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMSRLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] || !(c == d) {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSBshiftRLreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (RSBshiftRLreg (MOVWconst [c]) x y)
	// result: (SUBconst [c] (SRL <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMSUBconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (RSBshiftRLreg x y (MOVWconst [c]))
	// result: (RSBshiftRL x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMRSBshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSCconst(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (RSCconst [c] (ADDconst [d] x) flags)
	// result: (RSCconst [int64(int32(c-d))] x flags)
	for {
		c := v.AuxInt
		if v_0.Op != OpARMADDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		flags := v_1
		v.reset(OpARMRSCconst)
		v.AuxInt = int64(int32(c - d))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	// match: (RSCconst [c] (SUBconst [d] x) flags)
	// result: (RSCconst [int64(int32(c+d))] x flags)
	for {
		c := v.AuxInt
		if v_0.Op != OpARMSUBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		flags := v_1
		v.reset(OpARMRSCconst)
		v.AuxInt = int64(int32(c + d))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSCshiftLL(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (RSCshiftLL (MOVWconst [c]) x [d] flags)
	// result: (SBCconst [c] (SLLconst <x.Type> x [d]) flags)
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		flags := v_2
		v.reset(OpARMSBCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}
	// match: (RSCshiftLL x (MOVWconst [c]) [d] flags)
	// result: (RSCconst x [int64(int32(uint32(c)<<uint64(d)))] flags)
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		flags := v_2
		v.reset(OpARMRSCconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSCshiftLLreg(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (RSCshiftLLreg (MOVWconst [c]) x y flags)
	// result: (SBCconst [c] (SLL <x.Type> x y) flags)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		flags := v_3
		v.reset(OpARMSBCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}
	// match: (RSCshiftLLreg x y (MOVWconst [c]) flags)
	// result: (RSCshiftLL x y [c] flags)
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		flags := v_3
		v.reset(OpARMRSCshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSCshiftRA(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (RSCshiftRA (MOVWconst [c]) x [d] flags)
	// result: (SBCconst [c] (SRAconst <x.Type> x [d]) flags)
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		flags := v_2
		v.reset(OpARMSBCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}
	// match: (RSCshiftRA x (MOVWconst [c]) [d] flags)
	// result: (RSCconst x [int64(int32(c)>>uint64(d))] flags)
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		flags := v_2
		v.reset(OpARMRSCconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSCshiftRAreg(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (RSCshiftRAreg (MOVWconst [c]) x y flags)
	// result: (SBCconst [c] (SRA <x.Type> x y) flags)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		flags := v_3
		v.reset(OpARMSBCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRA, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}
	// match: (RSCshiftRAreg x y (MOVWconst [c]) flags)
	// result: (RSCshiftRA x y [c] flags)
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		flags := v_3
		v.reset(OpARMRSCshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSCshiftRL(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (RSCshiftRL (MOVWconst [c]) x [d] flags)
	// result: (SBCconst [c] (SRLconst <x.Type> x [d]) flags)
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		flags := v_2
		v.reset(OpARMSBCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}
	// match: (RSCshiftRL x (MOVWconst [c]) [d] flags)
	// result: (RSCconst x [int64(int32(uint32(c)>>uint64(d)))] flags)
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		flags := v_2
		v.reset(OpARMRSCconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMRSCshiftRLreg(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (RSCshiftRLreg (MOVWconst [c]) x y flags)
	// result: (SBCconst [c] (SRL <x.Type> x y) flags)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		flags := v_3
		v.reset(OpARMSBCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}
	// match: (RSCshiftRLreg x y (MOVWconst [c]) flags)
	// result: (RSCshiftRL x y [c] flags)
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		flags := v_3
		v.reset(OpARMRSCshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSBC(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SBC (MOVWconst [c]) x flags)
	// result: (RSCconst [c] x flags)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		flags := v_2
		v.reset(OpARMRSCconst)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	// match: (SBC x (MOVWconst [c]) flags)
	// result: (SBCconst [c] x flags)
	for {
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		flags := v_2
		v.reset(OpARMSBCconst)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	// match: (SBC x (SLLconst [c] y) flags)
	// result: (SBCshiftLL x y [c] flags)
	for {
		x := v_0
		if v_1.Op != OpARMSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		flags := v_2
		v.reset(OpARMSBCshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	// match: (SBC (SLLconst [c] y) x flags)
	// result: (RSCshiftLL x y [c] flags)
	for {
		if v_0.Op != OpARMSLLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v_1
		flags := v_2
		v.reset(OpARMRSCshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	// match: (SBC x (SRLconst [c] y) flags)
	// result: (SBCshiftRL x y [c] flags)
	for {
		x := v_0
		if v_1.Op != OpARMSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		flags := v_2
		v.reset(OpARMSBCshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	// match: (SBC (SRLconst [c] y) x flags)
	// result: (RSCshiftRL x y [c] flags)
	for {
		if v_0.Op != OpARMSRLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v_1
		flags := v_2
		v.reset(OpARMRSCshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	// match: (SBC x (SRAconst [c] y) flags)
	// result: (SBCshiftRA x y [c] flags)
	for {
		x := v_0
		if v_1.Op != OpARMSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		flags := v_2
		v.reset(OpARMSBCshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	// match: (SBC (SRAconst [c] y) x flags)
	// result: (RSCshiftRA x y [c] flags)
	for {
		if v_0.Op != OpARMSRAconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v_1
		flags := v_2
		v.reset(OpARMRSCshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	// match: (SBC x (SLL y z) flags)
	// result: (SBCshiftLLreg x y z flags)
	for {
		x := v_0
		if v_1.Op != OpARMSLL {
			break
		}
		z := v_1.Args[1]
		y := v_1.Args[0]
		flags := v_2
		v.reset(OpARMSBCshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		v.AddArg(flags)
		return true
	}
	// match: (SBC (SLL y z) x flags)
	// result: (RSCshiftLLreg x y z flags)
	for {
		if v_0.Op != OpARMSLL {
			break
		}
		z := v_0.Args[1]
		y := v_0.Args[0]
		x := v_1
		flags := v_2
		v.reset(OpARMRSCshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		v.AddArg(flags)
		return true
	}
	// match: (SBC x (SRL y z) flags)
	// result: (SBCshiftRLreg x y z flags)
	for {
		x := v_0
		if v_1.Op != OpARMSRL {
			break
		}
		z := v_1.Args[1]
		y := v_1.Args[0]
		flags := v_2
		v.reset(OpARMSBCshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		v.AddArg(flags)
		return true
	}
	// match: (SBC (SRL y z) x flags)
	// result: (RSCshiftRLreg x y z flags)
	for {
		if v_0.Op != OpARMSRL {
			break
		}
		z := v_0.Args[1]
		y := v_0.Args[0]
		x := v_1
		flags := v_2
		v.reset(OpARMRSCshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		v.AddArg(flags)
		return true
	}
	// match: (SBC x (SRA y z) flags)
	// result: (SBCshiftRAreg x y z flags)
	for {
		x := v_0
		if v_1.Op != OpARMSRA {
			break
		}
		z := v_1.Args[1]
		y := v_1.Args[0]
		flags := v_2
		v.reset(OpARMSBCshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		v.AddArg(flags)
		return true
	}
	// match: (SBC (SRA y z) x flags)
	// result: (RSCshiftRAreg x y z flags)
	for {
		if v_0.Op != OpARMSRA {
			break
		}
		z := v_0.Args[1]
		y := v_0.Args[0]
		x := v_1
		flags := v_2
		v.reset(OpARMRSCshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSBCconst(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SBCconst [c] (ADDconst [d] x) flags)
	// result: (SBCconst [int64(int32(c-d))] x flags)
	for {
		c := v.AuxInt
		if v_0.Op != OpARMADDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		flags := v_1
		v.reset(OpARMSBCconst)
		v.AuxInt = int64(int32(c - d))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	// match: (SBCconst [c] (SUBconst [d] x) flags)
	// result: (SBCconst [int64(int32(c+d))] x flags)
	for {
		c := v.AuxInt
		if v_0.Op != OpARMSUBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		flags := v_1
		v.reset(OpARMSBCconst)
		v.AuxInt = int64(int32(c + d))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSBCshiftLL(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (SBCshiftLL (MOVWconst [c]) x [d] flags)
	// result: (RSCconst [c] (SLLconst <x.Type> x [d]) flags)
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		flags := v_2
		v.reset(OpARMRSCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}
	// match: (SBCshiftLL x (MOVWconst [c]) [d] flags)
	// result: (SBCconst x [int64(int32(uint32(c)<<uint64(d)))] flags)
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		flags := v_2
		v.reset(OpARMSBCconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSBCshiftLLreg(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (SBCshiftLLreg (MOVWconst [c]) x y flags)
	// result: (RSCconst [c] (SLL <x.Type> x y) flags)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		flags := v_3
		v.reset(OpARMRSCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}
	// match: (SBCshiftLLreg x y (MOVWconst [c]) flags)
	// result: (SBCshiftLL x y [c] flags)
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		flags := v_3
		v.reset(OpARMSBCshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSBCshiftRA(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (SBCshiftRA (MOVWconst [c]) x [d] flags)
	// result: (RSCconst [c] (SRAconst <x.Type> x [d]) flags)
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		flags := v_2
		v.reset(OpARMRSCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}
	// match: (SBCshiftRA x (MOVWconst [c]) [d] flags)
	// result: (SBCconst x [int64(int32(c)>>uint64(d))] flags)
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		flags := v_2
		v.reset(OpARMSBCconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSBCshiftRAreg(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (SBCshiftRAreg (MOVWconst [c]) x y flags)
	// result: (RSCconst [c] (SRA <x.Type> x y) flags)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		flags := v_3
		v.reset(OpARMRSCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRA, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}
	// match: (SBCshiftRAreg x y (MOVWconst [c]) flags)
	// result: (SBCshiftRA x y [c] flags)
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		flags := v_3
		v.reset(OpARMSBCshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSBCshiftRL(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (SBCshiftRL (MOVWconst [c]) x [d] flags)
	// result: (RSCconst [c] (SRLconst <x.Type> x [d]) flags)
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		flags := v_2
		v.reset(OpARMRSCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}
	// match: (SBCshiftRL x (MOVWconst [c]) [d] flags)
	// result: (SBCconst x [int64(int32(uint32(c)>>uint64(d)))] flags)
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		flags := v_2
		v.reset(OpARMSBCconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSBCshiftRLreg(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (SBCshiftRLreg (MOVWconst [c]) x y flags)
	// result: (RSCconst [c] (SRL <x.Type> x y) flags)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		flags := v_3
		v.reset(OpARMRSCconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v.AddArg(flags)
		return true
	}
	// match: (SBCshiftRLreg x y (MOVWconst [c]) flags)
	// result: (SBCshiftRL x y [c] flags)
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		flags := v_3
		v.reset(OpARMSBCshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(flags)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSLL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SLL x (MOVWconst [c]))
	// result: (SLLconst x [c&31])
	for {
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMSLLconst)
		v.AuxInt = c & 31
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSLLconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SLLconst [c] (MOVWconst [d]))
	// result: (MOVWconst [int64(int32(uint32(d)<<uint64(c)))])
	for {
		c := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = int64(int32(uint32(d) << uint64(c)))
		return true
	}
	return false
}
func rewriteValueARM_OpARMSRA(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SRA x (MOVWconst [c]))
	// result: (SRAconst x [c&31])
	for {
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMSRAconst)
		v.AuxInt = c & 31
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSRAcond(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SRAcond x _ (FlagEQ))
	// result: (SRAconst x [31])
	for {
		x := v_0
		if v_2.Op != OpARMFlagEQ {
			break
		}
		v.reset(OpARMSRAconst)
		v.AuxInt = 31
		v.AddArg(x)
		return true
	}
	// match: (SRAcond x y (FlagLT_ULT))
	// result: (SRA x y)
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMFlagLT_ULT {
			break
		}
		v.reset(OpARMSRA)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SRAcond x _ (FlagLT_UGT))
	// result: (SRAconst x [31])
	for {
		x := v_0
		if v_2.Op != OpARMFlagLT_UGT {
			break
		}
		v.reset(OpARMSRAconst)
		v.AuxInt = 31
		v.AddArg(x)
		return true
	}
	// match: (SRAcond x y (FlagGT_ULT))
	// result: (SRA x y)
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMFlagGT_ULT {
			break
		}
		v.reset(OpARMSRA)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SRAcond x _ (FlagGT_UGT))
	// result: (SRAconst x [31])
	for {
		x := v_0
		if v_2.Op != OpARMFlagGT_UGT {
			break
		}
		v.reset(OpARMSRAconst)
		v.AuxInt = 31
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSRAconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SRAconst [c] (MOVWconst [d]))
	// result: (MOVWconst [int64(int32(d)>>uint64(c))])
	for {
		c := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = int64(int32(d) >> uint64(c))
		return true
	}
	// match: (SRAconst (SLLconst x [c]) [d])
	// cond: objabi.GOARM==7 && uint64(d)>=uint64(c) && uint64(d)<=31
	// result: (BFX [(d-c)|(32-d)<<8] x)
	for {
		d := v.AuxInt
		if v_0.Op != OpARMSLLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(objabi.GOARM == 7 && uint64(d) >= uint64(c) && uint64(d) <= 31) {
			break
		}
		v.reset(OpARMBFX)
		v.AuxInt = (d - c) | (32-d)<<8
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSRL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SRL x (MOVWconst [c]))
	// result: (SRLconst x [c&31])
	for {
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMSRLconst)
		v.AuxInt = c & 31
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSRLconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SRLconst [c] (MOVWconst [d]))
	// result: (MOVWconst [int64(int32(uint32(d)>>uint64(c)))])
	for {
		c := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = int64(int32(uint32(d) >> uint64(c)))
		return true
	}
	// match: (SRLconst (SLLconst x [c]) [d])
	// cond: objabi.GOARM==7 && uint64(d)>=uint64(c) && uint64(d)<=31
	// result: (BFXU [(d-c)|(32-d)<<8] x)
	for {
		d := v.AuxInt
		if v_0.Op != OpARMSLLconst {
			break
		}
		c := v_0.AuxInt
		x := v_0.Args[0]
		if !(objabi.GOARM == 7 && uint64(d) >= uint64(c) && uint64(d) <= 31) {
			break
		}
		v.reset(OpARMBFXU)
		v.AuxInt = (d - c) | (32-d)<<8
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSUB(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SUB (MOVWconst [c]) x)
	// result: (RSBconst [c] x)
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMRSBconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (SUB x (MOVWconst [c]))
	// result: (SUBconst [c] x)
	for {
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMSUBconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (SUB x (SLLconst [c] y))
	// result: (SUBshiftLL x y [c])
	for {
		x := v_0
		if v_1.Op != OpARMSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMSUBshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SUB (SLLconst [c] y) x)
	// result: (RSBshiftLL x y [c])
	for {
		if v_0.Op != OpARMSLLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v_1
		v.reset(OpARMRSBshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SUB x (SRLconst [c] y))
	// result: (SUBshiftRL x y [c])
	for {
		x := v_0
		if v_1.Op != OpARMSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMSUBshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SUB (SRLconst [c] y) x)
	// result: (RSBshiftRL x y [c])
	for {
		if v_0.Op != OpARMSRLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v_1
		v.reset(OpARMRSBshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SUB x (SRAconst [c] y))
	// result: (SUBshiftRA x y [c])
	for {
		x := v_0
		if v_1.Op != OpARMSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMSUBshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SUB (SRAconst [c] y) x)
	// result: (RSBshiftRA x y [c])
	for {
		if v_0.Op != OpARMSRAconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v_1
		v.reset(OpARMRSBshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SUB x (SLL y z))
	// result: (SUBshiftLLreg x y z)
	for {
		x := v_0
		if v_1.Op != OpARMSLL {
			break
		}
		z := v_1.Args[1]
		y := v_1.Args[0]
		v.reset(OpARMSUBshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	// match: (SUB (SLL y z) x)
	// result: (RSBshiftLLreg x y z)
	for {
		if v_0.Op != OpARMSLL {
			break
		}
		z := v_0.Args[1]
		y := v_0.Args[0]
		x := v_1
		v.reset(OpARMRSBshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	// match: (SUB x (SRL y z))
	// result: (SUBshiftRLreg x y z)
	for {
		x := v_0
		if v_1.Op != OpARMSRL {
			break
		}
		z := v_1.Args[1]
		y := v_1.Args[0]
		v.reset(OpARMSUBshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	// match: (SUB (SRL y z) x)
	// result: (RSBshiftRLreg x y z)
	for {
		if v_0.Op != OpARMSRL {
			break
		}
		z := v_0.Args[1]
		y := v_0.Args[0]
		x := v_1
		v.reset(OpARMRSBshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	// match: (SUB x (SRA y z))
	// result: (SUBshiftRAreg x y z)
	for {
		x := v_0
		if v_1.Op != OpARMSRA {
			break
		}
		z := v_1.Args[1]
		y := v_1.Args[0]
		v.reset(OpARMSUBshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	// match: (SUB (SRA y z) x)
	// result: (RSBshiftRAreg x y z)
	for {
		if v_0.Op != OpARMSRA {
			break
		}
		z := v_0.Args[1]
		y := v_0.Args[0]
		x := v_1
		v.reset(OpARMRSBshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	// match: (SUB x x)
	// result: (MOVWconst [0])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (SUB a (MUL x y))
	// cond: objabi.GOARM == 7
	// result: (MULS x y a)
	for {
		a := v_0
		if v_1.Op != OpARMMUL {
			break
		}
		y := v_1.Args[1]
		x := v_1.Args[0]
		if !(objabi.GOARM == 7) {
			break
		}
		v.reset(OpARMMULS)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(a)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSUBD(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SUBD a (MULD x y))
	// cond: a.Uses == 1 && objabi.GOARM >= 6
	// result: (MULSD a x y)
	for {
		a := v_0
		if v_1.Op != OpARMMULD {
			break
		}
		y := v_1.Args[1]
		x := v_1.Args[0]
		if !(a.Uses == 1 && objabi.GOARM >= 6) {
			break
		}
		v.reset(OpARMMULSD)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SUBD a (NMULD x y))
	// cond: a.Uses == 1 && objabi.GOARM >= 6
	// result: (MULAD a x y)
	for {
		a := v_0
		if v_1.Op != OpARMNMULD {
			break
		}
		y := v_1.Args[1]
		x := v_1.Args[0]
		if !(a.Uses == 1 && objabi.GOARM >= 6) {
			break
		}
		v.reset(OpARMMULAD)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSUBF(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SUBF a (MULF x y))
	// cond: a.Uses == 1 && objabi.GOARM >= 6
	// result: (MULSF a x y)
	for {
		a := v_0
		if v_1.Op != OpARMMULF {
			break
		}
		y := v_1.Args[1]
		x := v_1.Args[0]
		if !(a.Uses == 1 && objabi.GOARM >= 6) {
			break
		}
		v.reset(OpARMMULSF)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SUBF a (NMULF x y))
	// cond: a.Uses == 1 && objabi.GOARM >= 6
	// result: (MULAF a x y)
	for {
		a := v_0
		if v_1.Op != OpARMNMULF {
			break
		}
		y := v_1.Args[1]
		x := v_1.Args[0]
		if !(a.Uses == 1 && objabi.GOARM >= 6) {
			break
		}
		v.reset(OpARMMULAF)
		v.AddArg(a)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSUBS(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SUBS x (MOVWconst [c]))
	// result: (SUBSconst [c] x)
	for {
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMSUBSconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (SUBS x (SLLconst [c] y))
	// result: (SUBSshiftLL x y [c])
	for {
		x := v_0
		if v_1.Op != OpARMSLLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMSUBSshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SUBS (SLLconst [c] y) x)
	// result: (RSBSshiftLL x y [c])
	for {
		if v_0.Op != OpARMSLLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v_1
		v.reset(OpARMRSBSshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SUBS x (SRLconst [c] y))
	// result: (SUBSshiftRL x y [c])
	for {
		x := v_0
		if v_1.Op != OpARMSRLconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMSUBSshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SUBS (SRLconst [c] y) x)
	// result: (RSBSshiftRL x y [c])
	for {
		if v_0.Op != OpARMSRLconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v_1
		v.reset(OpARMRSBSshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SUBS x (SRAconst [c] y))
	// result: (SUBSshiftRA x y [c])
	for {
		x := v_0
		if v_1.Op != OpARMSRAconst {
			break
		}
		c := v_1.AuxInt
		y := v_1.Args[0]
		v.reset(OpARMSUBSshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SUBS (SRAconst [c] y) x)
	// result: (RSBSshiftRA x y [c])
	for {
		if v_0.Op != OpARMSRAconst {
			break
		}
		c := v_0.AuxInt
		y := v_0.Args[0]
		x := v_1
		v.reset(OpARMRSBSshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	// match: (SUBS x (SLL y z))
	// result: (SUBSshiftLLreg x y z)
	for {
		x := v_0
		if v_1.Op != OpARMSLL {
			break
		}
		z := v_1.Args[1]
		y := v_1.Args[0]
		v.reset(OpARMSUBSshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	// match: (SUBS (SLL y z) x)
	// result: (RSBSshiftLLreg x y z)
	for {
		if v_0.Op != OpARMSLL {
			break
		}
		z := v_0.Args[1]
		y := v_0.Args[0]
		x := v_1
		v.reset(OpARMRSBSshiftLLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	// match: (SUBS x (SRL y z))
	// result: (SUBSshiftRLreg x y z)
	for {
		x := v_0
		if v_1.Op != OpARMSRL {
			break
		}
		z := v_1.Args[1]
		y := v_1.Args[0]
		v.reset(OpARMSUBSshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	// match: (SUBS (SRL y z) x)
	// result: (RSBSshiftRLreg x y z)
	for {
		if v_0.Op != OpARMSRL {
			break
		}
		z := v_0.Args[1]
		y := v_0.Args[0]
		x := v_1
		v.reset(OpARMRSBSshiftRLreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	// match: (SUBS x (SRA y z))
	// result: (SUBSshiftRAreg x y z)
	for {
		x := v_0
		if v_1.Op != OpARMSRA {
			break
		}
		z := v_1.Args[1]
		y := v_1.Args[0]
		v.reset(OpARMSUBSshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	// match: (SUBS (SRA y z) x)
	// result: (RSBSshiftRAreg x y z)
	for {
		if v_0.Op != OpARMSRA {
			break
		}
		z := v_0.Args[1]
		y := v_0.Args[0]
		x := v_1
		v.reset(OpARMRSBSshiftRAreg)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(z)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSUBSshiftLL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (SUBSshiftLL (MOVWconst [c]) x [d])
	// result: (RSBSconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMRSBSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (SUBSshiftLL x (MOVWconst [c]) [d])
	// result: (SUBSconst x [int64(int32(uint32(c)<<uint64(d)))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMSUBSconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSUBSshiftLLreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (SUBSshiftLLreg (MOVWconst [c]) x y)
	// result: (RSBSconst [c] (SLL <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMRSBSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (SUBSshiftLLreg x y (MOVWconst [c]))
	// result: (SUBSshiftLL x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMSUBSshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSUBSshiftRA(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (SUBSshiftRA (MOVWconst [c]) x [d])
	// result: (RSBSconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMRSBSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (SUBSshiftRA x (MOVWconst [c]) [d])
	// result: (SUBSconst x [int64(int32(c)>>uint64(d))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMSUBSconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSUBSshiftRAreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (SUBSshiftRAreg (MOVWconst [c]) x y)
	// result: (RSBSconst [c] (SRA <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMRSBSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRA, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (SUBSshiftRAreg x y (MOVWconst [c]))
	// result: (SUBSshiftRA x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMSUBSshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSUBSshiftRL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (SUBSshiftRL (MOVWconst [c]) x [d])
	// result: (RSBSconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMRSBSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (SUBSshiftRL x (MOVWconst [c]) [d])
	// result: (SUBSconst x [int64(int32(uint32(c)>>uint64(d)))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMSUBSconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSUBSshiftRLreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (SUBSshiftRLreg (MOVWconst [c]) x y)
	// result: (RSBSconst [c] (SRL <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMRSBSconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (SUBSshiftRLreg x y (MOVWconst [c]))
	// result: (SUBSshiftRL x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMSUBSshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSUBconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SUBconst [off1] (MOVWaddr [off2] {sym} ptr))
	// result: (MOVWaddr [off2-off1] {sym} ptr)
	for {
		off1 := v.AuxInt
		if v_0.Op != OpARMMOVWaddr {
			break
		}
		off2 := v_0.AuxInt
		sym := v_0.Aux
		ptr := v_0.Args[0]
		v.reset(OpARMMOVWaddr)
		v.AuxInt = off2 - off1
		v.Aux = sym
		v.AddArg(ptr)
		return true
	}
	// match: (SUBconst [0] x)
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
	// match: (SUBconst [c] x)
	// cond: !isARMImmRot(uint32(c)) && isARMImmRot(uint32(-c))
	// result: (ADDconst [int64(int32(-c))] x)
	for {
		c := v.AuxInt
		x := v_0
		if !(!isARMImmRot(uint32(c)) && isARMImmRot(uint32(-c))) {
			break
		}
		v.reset(OpARMADDconst)
		v.AuxInt = int64(int32(-c))
		v.AddArg(x)
		return true
	}
	// match: (SUBconst [c] x)
	// cond: objabi.GOARM==7 && !isARMImmRot(uint32(c)) && uint32(c)>0xffff && uint32(-c)<=0xffff
	// result: (ANDconst [int64(int32(-c))] x)
	for {
		c := v.AuxInt
		x := v_0
		if !(objabi.GOARM == 7 && !isARMImmRot(uint32(c)) && uint32(c) > 0xffff && uint32(-c) <= 0xffff) {
			break
		}
		v.reset(OpARMANDconst)
		v.AuxInt = int64(int32(-c))
		v.AddArg(x)
		return true
	}
	// match: (SUBconst [c] (MOVWconst [d]))
	// result: (MOVWconst [int64(int32(d-c))])
	for {
		c := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = int64(int32(d - c))
		return true
	}
	// match: (SUBconst [c] (SUBconst [d] x))
	// result: (ADDconst [int64(int32(-c-d))] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpARMSUBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMADDconst)
		v.AuxInt = int64(int32(-c - d))
		v.AddArg(x)
		return true
	}
	// match: (SUBconst [c] (ADDconst [d] x))
	// result: (ADDconst [int64(int32(-c+d))] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpARMADDconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMADDconst)
		v.AuxInt = int64(int32(-c + d))
		v.AddArg(x)
		return true
	}
	// match: (SUBconst [c] (RSBconst [d] x))
	// result: (RSBconst [int64(int32(-c+d))] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpARMRSBconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMRSBconst)
		v.AuxInt = int64(int32(-c + d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSUBshiftLL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (SUBshiftLL (MOVWconst [c]) x [d])
	// result: (RSBconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMRSBconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (SUBshiftLL x (MOVWconst [c]) [d])
	// result: (SUBconst x [int64(int32(uint32(c)<<uint64(d)))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMSUBconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: (SUBshiftLL x (SLLconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMSLLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] || !(c == d) {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpARMSUBshiftLLreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (SUBshiftLLreg (MOVWconst [c]) x y)
	// result: (RSBconst [c] (SLL <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMRSBconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (SUBshiftLLreg x y (MOVWconst [c]))
	// result: (SUBshiftLL x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMSUBshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSUBshiftRA(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (SUBshiftRA (MOVWconst [c]) x [d])
	// result: (RSBconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMRSBconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (SUBshiftRA x (MOVWconst [c]) [d])
	// result: (SUBconst x [int64(int32(c)>>uint64(d))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMSUBconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	// match: (SUBshiftRA x (SRAconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMSRAconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] || !(c == d) {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpARMSUBshiftRAreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (SUBshiftRAreg (MOVWconst [c]) x y)
	// result: (RSBconst [c] (SRA <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMRSBconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRA, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (SUBshiftRAreg x y (MOVWconst [c]))
	// result: (SUBshiftRA x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMSUBshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMSUBshiftRL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (SUBshiftRL (MOVWconst [c]) x [d])
	// result: (RSBconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMRSBconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (SUBshiftRL x (MOVWconst [c]) [d])
	// result: (SUBconst x [int64(int32(uint32(c)>>uint64(d)))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMSUBconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: (SUBshiftRL x (SRLconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMSRLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] || !(c == d) {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpARMSUBshiftRLreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (SUBshiftRLreg (MOVWconst [c]) x y)
	// result: (RSBconst [c] (SRL <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMRSBconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (SUBshiftRLreg x y (MOVWconst [c]))
	// result: (SUBshiftRL x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMSUBshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMTEQ(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (TEQ x (MOVWconst [c]))
	// result: (TEQconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMMOVWconst {
				continue
			}
			c := v_1.AuxInt
			v.reset(OpARMTEQconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (TEQ x (SLLconst [c] y))
	// result: (TEQshiftLL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSLLconst {
				continue
			}
			c := v_1.AuxInt
			y := v_1.Args[0]
			v.reset(OpARMTEQshiftLL)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (TEQ x (SRLconst [c] y))
	// result: (TEQshiftRL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRLconst {
				continue
			}
			c := v_1.AuxInt
			y := v_1.Args[0]
			v.reset(OpARMTEQshiftRL)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (TEQ x (SRAconst [c] y))
	// result: (TEQshiftRA x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRAconst {
				continue
			}
			c := v_1.AuxInt
			y := v_1.Args[0]
			v.reset(OpARMTEQshiftRA)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (TEQ x (SLL y z))
	// result: (TEQshiftLLreg x y z)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSLL {
				continue
			}
			z := v_1.Args[1]
			y := v_1.Args[0]
			v.reset(OpARMTEQshiftLLreg)
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(z)
			return true
		}
		break
	}
	// match: (TEQ x (SRL y z))
	// result: (TEQshiftRLreg x y z)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRL {
				continue
			}
			z := v_1.Args[1]
			y := v_1.Args[0]
			v.reset(OpARMTEQshiftRLreg)
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(z)
			return true
		}
		break
	}
	// match: (TEQ x (SRA y z))
	// result: (TEQshiftRAreg x y z)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRA {
				continue
			}
			z := v_1.Args[1]
			y := v_1.Args[0]
			v.reset(OpARMTEQshiftRAreg)
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(z)
			return true
		}
		break
	}
	return false
}
func rewriteValueARM_OpARMTEQconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (TEQconst (MOVWconst [x]) [y])
	// cond: int32(x^y)==0
	// result: (FlagEQ)
	for {
		y := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x^y) == 0) {
			break
		}
		v.reset(OpARMFlagEQ)
		return true
	}
	// match: (TEQconst (MOVWconst [x]) [y])
	// cond: int32(x^y)<0
	// result: (FlagLT_UGT)
	for {
		y := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x^y) < 0) {
			break
		}
		v.reset(OpARMFlagLT_UGT)
		return true
	}
	// match: (TEQconst (MOVWconst [x]) [y])
	// cond: int32(x^y)>0
	// result: (FlagGT_UGT)
	for {
		y := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x^y) > 0) {
			break
		}
		v.reset(OpARMFlagGT_UGT)
		return true
	}
	return false
}
func rewriteValueARM_OpARMTEQshiftLL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (TEQshiftLL (MOVWconst [c]) x [d])
	// result: (TEQconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMTEQconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (TEQshiftLL x (MOVWconst [c]) [d])
	// result: (TEQconst x [int64(int32(uint32(c)<<uint64(d)))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMTEQconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMTEQshiftLLreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (TEQshiftLLreg (MOVWconst [c]) x y)
	// result: (TEQconst [c] (SLL <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMTEQconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (TEQshiftLLreg x y (MOVWconst [c]))
	// result: (TEQshiftLL x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMTEQshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMTEQshiftRA(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (TEQshiftRA (MOVWconst [c]) x [d])
	// result: (TEQconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMTEQconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (TEQshiftRA x (MOVWconst [c]) [d])
	// result: (TEQconst x [int64(int32(c)>>uint64(d))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMTEQconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMTEQshiftRAreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (TEQshiftRAreg (MOVWconst [c]) x y)
	// result: (TEQconst [c] (SRA <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMTEQconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRA, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (TEQshiftRAreg x y (MOVWconst [c]))
	// result: (TEQshiftRA x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMTEQshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMTEQshiftRL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (TEQshiftRL (MOVWconst [c]) x [d])
	// result: (TEQconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMTEQconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (TEQshiftRL x (MOVWconst [c]) [d])
	// result: (TEQconst x [int64(int32(uint32(c)>>uint64(d)))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMTEQconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMTEQshiftRLreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (TEQshiftRLreg (MOVWconst [c]) x y)
	// result: (TEQconst [c] (SRL <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMTEQconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (TEQshiftRLreg x y (MOVWconst [c]))
	// result: (TEQshiftRL x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMTEQshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMTST(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (TST x (MOVWconst [c]))
	// result: (TSTconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMMOVWconst {
				continue
			}
			c := v_1.AuxInt
			v.reset(OpARMTSTconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (TST x (SLLconst [c] y))
	// result: (TSTshiftLL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSLLconst {
				continue
			}
			c := v_1.AuxInt
			y := v_1.Args[0]
			v.reset(OpARMTSTshiftLL)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (TST x (SRLconst [c] y))
	// result: (TSTshiftRL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRLconst {
				continue
			}
			c := v_1.AuxInt
			y := v_1.Args[0]
			v.reset(OpARMTSTshiftRL)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (TST x (SRAconst [c] y))
	// result: (TSTshiftRA x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRAconst {
				continue
			}
			c := v_1.AuxInt
			y := v_1.Args[0]
			v.reset(OpARMTSTshiftRA)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (TST x (SLL y z))
	// result: (TSTshiftLLreg x y z)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSLL {
				continue
			}
			z := v_1.Args[1]
			y := v_1.Args[0]
			v.reset(OpARMTSTshiftLLreg)
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(z)
			return true
		}
		break
	}
	// match: (TST x (SRL y z))
	// result: (TSTshiftRLreg x y z)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRL {
				continue
			}
			z := v_1.Args[1]
			y := v_1.Args[0]
			v.reset(OpARMTSTshiftRLreg)
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(z)
			return true
		}
		break
	}
	// match: (TST x (SRA y z))
	// result: (TSTshiftRAreg x y z)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRA {
				continue
			}
			z := v_1.Args[1]
			y := v_1.Args[0]
			v.reset(OpARMTSTshiftRAreg)
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(z)
			return true
		}
		break
	}
	return false
}
func rewriteValueARM_OpARMTSTconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (TSTconst (MOVWconst [x]) [y])
	// cond: int32(x&y)==0
	// result: (FlagEQ)
	for {
		y := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x&y) == 0) {
			break
		}
		v.reset(OpARMFlagEQ)
		return true
	}
	// match: (TSTconst (MOVWconst [x]) [y])
	// cond: int32(x&y)<0
	// result: (FlagLT_UGT)
	for {
		y := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x&y) < 0) {
			break
		}
		v.reset(OpARMFlagLT_UGT)
		return true
	}
	// match: (TSTconst (MOVWconst [x]) [y])
	// cond: int32(x&y)>0
	// result: (FlagGT_UGT)
	for {
		y := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		x := v_0.AuxInt
		if !(int32(x&y) > 0) {
			break
		}
		v.reset(OpARMFlagGT_UGT)
		return true
	}
	return false
}
func rewriteValueARM_OpARMTSTshiftLL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (TSTshiftLL (MOVWconst [c]) x [d])
	// result: (TSTconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMTSTconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (TSTshiftLL x (MOVWconst [c]) [d])
	// result: (TSTconst x [int64(int32(uint32(c)<<uint64(d)))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMTSTconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMTSTshiftLLreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (TSTshiftLLreg (MOVWconst [c]) x y)
	// result: (TSTconst [c] (SLL <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMTSTconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (TSTshiftLLreg x y (MOVWconst [c]))
	// result: (TSTshiftLL x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMTSTshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMTSTshiftRA(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (TSTshiftRA (MOVWconst [c]) x [d])
	// result: (TSTconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMTSTconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (TSTshiftRA x (MOVWconst [c]) [d])
	// result: (TSTconst x [int64(int32(c)>>uint64(d))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMTSTconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMTSTshiftRAreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (TSTshiftRAreg (MOVWconst [c]) x y)
	// result: (TSTconst [c] (SRA <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMTSTconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRA, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (TSTshiftRAreg x y (MOVWconst [c]))
	// result: (TSTshiftRA x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMTSTshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMTSTshiftRL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (TSTshiftRL (MOVWconst [c]) x [d])
	// result: (TSTconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMTSTconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (TSTshiftRL x (MOVWconst [c]) [d])
	// result: (TSTconst x [int64(int32(uint32(c)>>uint64(d)))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMTSTconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMTSTshiftRLreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (TSTshiftRLreg (MOVWconst [c]) x y)
	// result: (TSTconst [c] (SRL <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMTSTconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (TSTshiftRLreg x y (MOVWconst [c]))
	// result: (TSTshiftRL x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMTSTshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMXOR(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (XOR x (MOVWconst [c]))
	// result: (XORconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMMOVWconst {
				continue
			}
			c := v_1.AuxInt
			v.reset(OpARMXORconst)
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		break
	}
	// match: (XOR x (SLLconst [c] y))
	// result: (XORshiftLL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSLLconst {
				continue
			}
			c := v_1.AuxInt
			y := v_1.Args[0]
			v.reset(OpARMXORshiftLL)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (XOR x (SRLconst [c] y))
	// result: (XORshiftRL x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRLconst {
				continue
			}
			c := v_1.AuxInt
			y := v_1.Args[0]
			v.reset(OpARMXORshiftRL)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (XOR x (SRAconst [c] y))
	// result: (XORshiftRA x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRAconst {
				continue
			}
			c := v_1.AuxInt
			y := v_1.Args[0]
			v.reset(OpARMXORshiftRA)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (XOR x (SRRconst [c] y))
	// result: (XORshiftRR x y [c])
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRRconst {
				continue
			}
			c := v_1.AuxInt
			y := v_1.Args[0]
			v.reset(OpARMXORshiftRR)
			v.AuxInt = c
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		break
	}
	// match: (XOR x (SLL y z))
	// result: (XORshiftLLreg x y z)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSLL {
				continue
			}
			z := v_1.Args[1]
			y := v_1.Args[0]
			v.reset(OpARMXORshiftLLreg)
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(z)
			return true
		}
		break
	}
	// match: (XOR x (SRL y z))
	// result: (XORshiftRLreg x y z)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRL {
				continue
			}
			z := v_1.Args[1]
			y := v_1.Args[0]
			v.reset(OpARMXORshiftRLreg)
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(z)
			return true
		}
		break
	}
	// match: (XOR x (SRA y z))
	// result: (XORshiftRAreg x y z)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			x := v_0
			if v_1.Op != OpARMSRA {
				continue
			}
			z := v_1.Args[1]
			y := v_1.Args[0]
			v.reset(OpARMXORshiftRAreg)
			v.AddArg(x)
			v.AddArg(y)
			v.AddArg(z)
			return true
		}
		break
	}
	// match: (XOR x x)
	// result: (MOVWconst [0])
	for {
		x := v_0
		if x != v_1 {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpARMXORconst(v *Value) bool {
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
	// match: (XORconst [c] (MOVWconst [d]))
	// result: (MOVWconst [c^d])
	for {
		c := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		d := v_0.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = c ^ d
		return true
	}
	// match: (XORconst [c] (XORconst [d] x))
	// result: (XORconst [c^d] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpARMXORconst {
			break
		}
		d := v_0.AuxInt
		x := v_0.Args[0]
		v.reset(OpARMXORconst)
		v.AuxInt = c ^ d
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpARMXORshiftLL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (XORshiftLL (MOVWconst [c]) x [d])
	// result: (XORconst [c] (SLLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMXORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (XORshiftLL x (MOVWconst [c]) [d])
	// result: (XORconst x [int64(int32(uint32(c)<<uint64(d)))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMXORconst)
		v.AuxInt = int64(int32(uint32(c) << uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: (XORshiftLL [c] (SRLconst x [32-c]) x)
	// result: (SRRconst [32-c] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpARMSRLconst || v_0.AuxInt != 32-c {
			break
		}
		x := v_0.Args[0]
		if x != v_1 {
			break
		}
		v.reset(OpARMSRRconst)
		v.AuxInt = 32 - c
		v.AddArg(x)
		return true
	}
	// match: (XORshiftLL <typ.UInt16> [8] (BFXU <typ.UInt16> [armBFAuxInt(8, 8)] x) x)
	// result: (REV16 x)
	for {
		if v.Type != typ.UInt16 || v.AuxInt != 8 || v_0.Op != OpARMBFXU || v_0.Type != typ.UInt16 || v_0.AuxInt != armBFAuxInt(8, 8) {
			break
		}
		x := v_0.Args[0]
		if x != v_1 {
			break
		}
		v.reset(OpARMREV16)
		v.AddArg(x)
		return true
	}
	// match: (XORshiftLL <typ.UInt16> [8] (SRLconst <typ.UInt16> [24] (SLLconst [16] x)) x)
	// cond: objabi.GOARM>=6
	// result: (REV16 x)
	for {
		if v.Type != typ.UInt16 || v.AuxInt != 8 || v_0.Op != OpARMSRLconst || v_0.Type != typ.UInt16 || v_0.AuxInt != 24 {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpARMSLLconst || v_0_0.AuxInt != 16 {
			break
		}
		x := v_0_0.Args[0]
		if x != v_1 || !(objabi.GOARM >= 6) {
			break
		}
		v.reset(OpARMREV16)
		v.AddArg(x)
		return true
	}
	// match: (XORshiftLL x (SLLconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMSLLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] || !(c == d) {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpARMXORshiftLLreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (XORshiftLLreg (MOVWconst [c]) x y)
	// result: (XORconst [c] (SLL <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMXORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (XORshiftLLreg x y (MOVWconst [c]))
	// result: (XORshiftLL x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMXORshiftLL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMXORshiftRA(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (XORshiftRA (MOVWconst [c]) x [d])
	// result: (XORconst [c] (SRAconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMXORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRAconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (XORshiftRA x (MOVWconst [c]) [d])
	// result: (XORconst x [int64(int32(c)>>uint64(d))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMXORconst)
		v.AuxInt = int64(int32(c) >> uint64(d))
		v.AddArg(x)
		return true
	}
	// match: (XORshiftRA x (SRAconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMSRAconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] || !(c == d) {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpARMXORshiftRAreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (XORshiftRAreg (MOVWconst [c]) x y)
	// result: (XORconst [c] (SRA <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMXORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRA, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (XORshiftRAreg x y (MOVWconst [c]))
	// result: (XORshiftRA x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMXORshiftRA)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMXORshiftRL(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (XORshiftRL (MOVWconst [c]) x [d])
	// result: (XORconst [c] (SRLconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMXORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRLconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (XORshiftRL x (MOVWconst [c]) [d])
	// result: (XORconst x [int64(int32(uint32(c)>>uint64(d)))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMXORconst)
		v.AuxInt = int64(int32(uint32(c) >> uint64(d)))
		v.AddArg(x)
		return true
	}
	// match: (XORshiftRL [c] (SLLconst x [32-c]) x)
	// result: (SRRconst [ c] x)
	for {
		c := v.AuxInt
		if v_0.Op != OpARMSLLconst || v_0.AuxInt != 32-c {
			break
		}
		x := v_0.Args[0]
		if x != v_1 {
			break
		}
		v.reset(OpARMSRRconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (XORshiftRL x (SRLconst x [c]) [d])
	// cond: c==d
	// result: (MOVWconst [0])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMSRLconst {
			break
		}
		c := v_1.AuxInt
		if x != v_1.Args[0] || !(c == d) {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpARMXORshiftRLreg(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (XORshiftRLreg (MOVWconst [c]) x y)
	// result: (XORconst [c] (SRL <x.Type> x y))
	for {
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		y := v_2
		v.reset(OpARMXORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
	// match: (XORshiftRLreg x y (MOVWconst [c]))
	// result: (XORshiftRL x y [c])
	for {
		x := v_0
		y := v_1
		if v_2.Op != OpARMMOVWconst {
			break
		}
		c := v_2.AuxInt
		v.reset(OpARMXORshiftRL)
		v.AuxInt = c
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
	return false
}
func rewriteValueARM_OpARMXORshiftRR(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (XORshiftRR (MOVWconst [c]) x [d])
	// result: (XORconst [c] (SRRconst <x.Type> x [d]))
	for {
		d := v.AuxInt
		if v_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0.AuxInt
		x := v_1
		v.reset(OpARMXORconst)
		v.AuxInt = c
		v0 := b.NewValue0(v.Pos, OpARMSRRconst, x.Type)
		v0.AuxInt = d
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (XORshiftRR x (MOVWconst [c]) [d])
	// result: (XORconst x [int64(int32(uint32(c)>>uint64(d)|uint32(c)<<uint64(32-d)))])
	for {
		d := v.AuxInt
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMXORconst)
		v.AuxInt = int64(int32(uint32(c)>>uint64(d) | uint32(c)<<uint64(32-d)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpAbs(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Abs x)
	// result: (ABSD x)
	for {
		x := v_0
		v.reset(OpARMABSD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpAdd16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Add16 x y)
	// result: (ADD x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpAdd32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Add32 x y)
	// result: (ADD x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpAdd32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Add32F x y)
	// result: (ADDF x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMADDF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpAdd32carry(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Add32carry x y)
	// result: (ADDS x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMADDS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpAdd32withcarry(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Add32withcarry x y c)
	// result: (ADC x y c)
	for {
		x := v_0
		y := v_1
		c := v_2
		v.reset(OpARMADC)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(c)
		return true
	}
}
func rewriteValueARM_OpAdd64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Add64F x y)
	// result: (ADDD x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMADDD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpAdd8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Add8 x y)
	// result: (ADD x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpAddPtr(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (AddPtr x y)
	// result: (ADD x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMADD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpAddr(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Addr {sym} base)
	// result: (MOVWaddr {sym} base)
	for {
		sym := v.Aux
		base := v_0
		v.reset(OpARMMOVWaddr)
		v.Aux = sym
		v.AddArg(base)
		return true
	}
}
func rewriteValueARM_OpAnd16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (And16 x y)
	// result: (AND x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMAND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpAnd32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (And32 x y)
	// result: (AND x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMAND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpAnd8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (And8 x y)
	// result: (AND x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMAND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpAndB(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (AndB x y)
	// result: (AND x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMAND)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpAvg32u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Avg32u <t> x y)
	// result: (ADD (SRLconst <t> (SUB <t> x y) [1]) y)
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpARMADD)
		v0 := b.NewValue0(v.Pos, OpARMSRLconst, t)
		v0.AuxInt = 1
		v1 := b.NewValue0(v.Pos, OpARMSUB, t)
		v1.AddArg(x)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpBitLen32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (BitLen32 <t> x)
	// result: (RSBconst [32] (CLZ <t> x))
	for {
		t := v.Type
		x := v_0
		v.reset(OpARMRSBconst)
		v.AuxInt = 32
		v0 := b.NewValue0(v.Pos, OpARMCLZ, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpBswap32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (Bswap32 <t> x)
	// cond: objabi.GOARM==5
	// result: (XOR <t> (SRLconst <t> (BICconst <t> (XOR <t> x (SRRconst <t> [16] x)) [0xff0000]) [8]) (SRRconst <t> x [8]))
	for {
		t := v.Type
		x := v_0
		if !(objabi.GOARM == 5) {
			break
		}
		v.reset(OpARMXOR)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpARMSRLconst, t)
		v0.AuxInt = 8
		v1 := b.NewValue0(v.Pos, OpARMBICconst, t)
		v1.AuxInt = 0xff0000
		v2 := b.NewValue0(v.Pos, OpARMXOR, t)
		v2.AddArg(x)
		v3 := b.NewValue0(v.Pos, OpARMSRRconst, t)
		v3.AuxInt = 16
		v3.AddArg(x)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		v4 := b.NewValue0(v.Pos, OpARMSRRconst, t)
		v4.AuxInt = 8
		v4.AddArg(x)
		v.AddArg(v4)
		return true
	}
	// match: (Bswap32 x)
	// cond: objabi.GOARM>=6
	// result: (REV x)
	for {
		x := v_0
		if !(objabi.GOARM >= 6) {
			break
		}
		v.reset(OpARMREV)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpClosureCall(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ClosureCall [argwid] entry closure mem)
	// result: (CALLclosure [argwid] entry closure mem)
	for {
		argwid := v.AuxInt
		entry := v_0
		closure := v_1
		mem := v_2
		v.reset(OpARMCALLclosure)
		v.AuxInt = argwid
		v.AddArg(entry)
		v.AddArg(closure)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueARM_OpCom16(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Com16 x)
	// result: (MVN x)
	for {
		x := v_0
		v.reset(OpARMMVN)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpCom32(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Com32 x)
	// result: (MVN x)
	for {
		x := v_0
		v.reset(OpARMMVN)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpCom8(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Com8 x)
	// result: (MVN x)
	for {
		x := v_0
		v.reset(OpARMMVN)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpConst16(v *Value) bool {
	// match: (Const16 [val])
	// result: (MOVWconst [val])
	for {
		val := v.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueARM_OpConst32(v *Value) bool {
	// match: (Const32 [val])
	// result: (MOVWconst [val])
	for {
		val := v.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueARM_OpConst32F(v *Value) bool {
	// match: (Const32F [val])
	// result: (MOVFconst [val])
	for {
		val := v.AuxInt
		v.reset(OpARMMOVFconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueARM_OpConst64F(v *Value) bool {
	// match: (Const64F [val])
	// result: (MOVDconst [val])
	for {
		val := v.AuxInt
		v.reset(OpARMMOVDconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueARM_OpConst8(v *Value) bool {
	// match: (Const8 [val])
	// result: (MOVWconst [val])
	for {
		val := v.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = val
		return true
	}
}
func rewriteValueARM_OpConstBool(v *Value) bool {
	// match: (ConstBool [b])
	// result: (MOVWconst [b])
	for {
		b := v.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = b
		return true
	}
}
func rewriteValueARM_OpConstNil(v *Value) bool {
	// match: (ConstNil)
	// result: (MOVWconst [0])
	for {
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
}
func rewriteValueARM_OpCtz16(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Ctz16 <t> x)
	// cond: objabi.GOARM<=6
	// result: (RSBconst [32] (CLZ <t> (SUBconst <typ.UInt32> (AND <typ.UInt32> (ORconst <typ.UInt32> [0x10000] x) (RSBconst <typ.UInt32> [0] (ORconst <typ.UInt32> [0x10000] x))) [1])))
	for {
		t := v.Type
		x := v_0
		if !(objabi.GOARM <= 6) {
			break
		}
		v.reset(OpARMRSBconst)
		v.AuxInt = 32
		v0 := b.NewValue0(v.Pos, OpARMCLZ, t)
		v1 := b.NewValue0(v.Pos, OpARMSUBconst, typ.UInt32)
		v1.AuxInt = 1
		v2 := b.NewValue0(v.Pos, OpARMAND, typ.UInt32)
		v3 := b.NewValue0(v.Pos, OpARMORconst, typ.UInt32)
		v3.AuxInt = 0x10000
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpARMRSBconst, typ.UInt32)
		v4.AuxInt = 0
		v5 := b.NewValue0(v.Pos, OpARMORconst, typ.UInt32)
		v5.AuxInt = 0x10000
		v5.AddArg(x)
		v4.AddArg(v5)
		v2.AddArg(v4)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Ctz16 <t> x)
	// cond: objabi.GOARM==7
	// result: (CLZ <t> (RBIT <typ.UInt32> (ORconst <typ.UInt32> [0x10000] x)))
	for {
		t := v.Type
		x := v_0
		if !(objabi.GOARM == 7) {
			break
		}
		v.reset(OpARMCLZ)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpARMRBIT, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpARMORconst, typ.UInt32)
		v1.AuxInt = 0x10000
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueARM_OpCtz16NonZero(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Ctz16NonZero x)
	// result: (Ctz32 x)
	for {
		x := v_0
		v.reset(OpCtz32)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpCtz32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (Ctz32 <t> x)
	// cond: objabi.GOARM<=6
	// result: (RSBconst [32] (CLZ <t> (SUBconst <t> (AND <t> x (RSBconst <t> [0] x)) [1])))
	for {
		t := v.Type
		x := v_0
		if !(objabi.GOARM <= 6) {
			break
		}
		v.reset(OpARMRSBconst)
		v.AuxInt = 32
		v0 := b.NewValue0(v.Pos, OpARMCLZ, t)
		v1 := b.NewValue0(v.Pos, OpARMSUBconst, t)
		v1.AuxInt = 1
		v2 := b.NewValue0(v.Pos, OpARMAND, t)
		v2.AddArg(x)
		v3 := b.NewValue0(v.Pos, OpARMRSBconst, t)
		v3.AuxInt = 0
		v3.AddArg(x)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Ctz32 <t> x)
	// cond: objabi.GOARM==7
	// result: (CLZ <t> (RBIT <t> x))
	for {
		t := v.Type
		x := v_0
		if !(objabi.GOARM == 7) {
			break
		}
		v.reset(OpARMCLZ)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpARMRBIT, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueARM_OpCtz32NonZero(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Ctz32NonZero x)
	// result: (Ctz32 x)
	for {
		x := v_0
		v.reset(OpCtz32)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpCtz8(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Ctz8 <t> x)
	// cond: objabi.GOARM<=6
	// result: (RSBconst [32] (CLZ <t> (SUBconst <typ.UInt32> (AND <typ.UInt32> (ORconst <typ.UInt32> [0x100] x) (RSBconst <typ.UInt32> [0] (ORconst <typ.UInt32> [0x100] x))) [1])))
	for {
		t := v.Type
		x := v_0
		if !(objabi.GOARM <= 6) {
			break
		}
		v.reset(OpARMRSBconst)
		v.AuxInt = 32
		v0 := b.NewValue0(v.Pos, OpARMCLZ, t)
		v1 := b.NewValue0(v.Pos, OpARMSUBconst, typ.UInt32)
		v1.AuxInt = 1
		v2 := b.NewValue0(v.Pos, OpARMAND, typ.UInt32)
		v3 := b.NewValue0(v.Pos, OpARMORconst, typ.UInt32)
		v3.AuxInt = 0x100
		v3.AddArg(x)
		v2.AddArg(v3)
		v4 := b.NewValue0(v.Pos, OpARMRSBconst, typ.UInt32)
		v4.AuxInt = 0
		v5 := b.NewValue0(v.Pos, OpARMORconst, typ.UInt32)
		v5.AuxInt = 0x100
		v5.AddArg(x)
		v4.AddArg(v5)
		v2.AddArg(v4)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	// match: (Ctz8 <t> x)
	// cond: objabi.GOARM==7
	// result: (CLZ <t> (RBIT <typ.UInt32> (ORconst <typ.UInt32> [0x100] x)))
	for {
		t := v.Type
		x := v_0
		if !(objabi.GOARM == 7) {
			break
		}
		v.reset(OpARMCLZ)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpARMRBIT, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpARMORconst, typ.UInt32)
		v1.AuxInt = 0x100
		v1.AddArg(x)
		v0.AddArg(v1)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueARM_OpCtz8NonZero(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Ctz8NonZero x)
	// result: (Ctz32 x)
	for {
		x := v_0
		v.reset(OpCtz32)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpCvt32Fto32(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Cvt32Fto32 x)
	// result: (MOVFW x)
	for {
		x := v_0
		v.reset(OpARMMOVFW)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpCvt32Fto32U(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Cvt32Fto32U x)
	// result: (MOVFWU x)
	for {
		x := v_0
		v.reset(OpARMMOVFWU)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpCvt32Fto64F(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Cvt32Fto64F x)
	// result: (MOVFD x)
	for {
		x := v_0
		v.reset(OpARMMOVFD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpCvt32Uto32F(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Cvt32Uto32F x)
	// result: (MOVWUF x)
	for {
		x := v_0
		v.reset(OpARMMOVWUF)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpCvt32Uto64F(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Cvt32Uto64F x)
	// result: (MOVWUD x)
	for {
		x := v_0
		v.reset(OpARMMOVWUD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpCvt32to32F(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Cvt32to32F x)
	// result: (MOVWF x)
	for {
		x := v_0
		v.reset(OpARMMOVWF)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpCvt32to64F(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Cvt32to64F x)
	// result: (MOVWD x)
	for {
		x := v_0
		v.reset(OpARMMOVWD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpCvt64Fto32(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Cvt64Fto32 x)
	// result: (MOVDW x)
	for {
		x := v_0
		v.reset(OpARMMOVDW)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpCvt64Fto32F(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Cvt64Fto32F x)
	// result: (MOVDF x)
	for {
		x := v_0
		v.reset(OpARMMOVDF)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpCvt64Fto32U(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Cvt64Fto32U x)
	// result: (MOVDWU x)
	for {
		x := v_0
		v.reset(OpARMMOVDWU)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpDiv16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div16 x y)
	// result: (Div32 (SignExt16to32 x) (SignExt16to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpDiv32)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM_OpDiv16u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div16u x y)
	// result: (Div32u (ZeroExt16to32 x) (ZeroExt16to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpDiv32u)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM_OpDiv32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div32 x y)
	// result: (SUB (XOR <typ.UInt32> (Select0 <typ.UInt32> (CALLudiv (SUB <typ.UInt32> (XOR x <typ.UInt32> (Signmask x)) (Signmask x)) (SUB <typ.UInt32> (XOR y <typ.UInt32> (Signmask y)) (Signmask y)))) (Signmask (XOR <typ.UInt32> x y))) (Signmask (XOR <typ.UInt32> x y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMSUB)
		v0 := b.NewValue0(v.Pos, OpARMXOR, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpSelect0, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpARMCALLudiv, types.NewTuple(typ.UInt32, typ.UInt32))
		v3 := b.NewValue0(v.Pos, OpARMSUB, typ.UInt32)
		v4 := b.NewValue0(v.Pos, OpARMXOR, typ.UInt32)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpSignmask, typ.Int32)
		v5.AddArg(x)
		v4.AddArg(v5)
		v3.AddArg(v4)
		v6 := b.NewValue0(v.Pos, OpSignmask, typ.Int32)
		v6.AddArg(x)
		v3.AddArg(v6)
		v2.AddArg(v3)
		v7 := b.NewValue0(v.Pos, OpARMSUB, typ.UInt32)
		v8 := b.NewValue0(v.Pos, OpARMXOR, typ.UInt32)
		v8.AddArg(y)
		v9 := b.NewValue0(v.Pos, OpSignmask, typ.Int32)
		v9.AddArg(y)
		v8.AddArg(v9)
		v7.AddArg(v8)
		v10 := b.NewValue0(v.Pos, OpSignmask, typ.Int32)
		v10.AddArg(y)
		v7.AddArg(v10)
		v2.AddArg(v7)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v11 := b.NewValue0(v.Pos, OpSignmask, typ.Int32)
		v12 := b.NewValue0(v.Pos, OpARMXOR, typ.UInt32)
		v12.AddArg(x)
		v12.AddArg(y)
		v11.AddArg(v12)
		v0.AddArg(v11)
		v.AddArg(v0)
		v13 := b.NewValue0(v.Pos, OpSignmask, typ.Int32)
		v14 := b.NewValue0(v.Pos, OpARMXOR, typ.UInt32)
		v14.AddArg(x)
		v14.AddArg(y)
		v13.AddArg(v14)
		v.AddArg(v13)
		return true
	}
}
func rewriteValueARM_OpDiv32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Div32F x y)
	// result: (DIVF x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMDIVF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpDiv32u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div32u x y)
	// result: (Select0 <typ.UInt32> (CALLudiv x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect0)
		v.Type = typ.UInt32
		v0 := b.NewValue0(v.Pos, OpARMCALLudiv, types.NewTuple(typ.UInt32, typ.UInt32))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpDiv64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Div64F x y)
	// result: (DIVD x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMDIVD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpDiv8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div8 x y)
	// result: (Div32 (SignExt8to32 x) (SignExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpDiv32)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM_OpDiv8u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div8u x y)
	// result: (Div32u (ZeroExt8to32 x) (ZeroExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpDiv32u)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM_OpEq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq16 x y)
	// result: (Equal (CMP (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpEq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Eq32 x y)
	// result: (Equal (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMP, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpEq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Eq32F x y)
	// result: (Equal (CMPF x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMPF, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpEq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Eq64F x y)
	// result: (Equal (CMPD x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMPD, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpEq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq8 x y)
	// result: (Equal (CMP (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpEqB(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (EqB x y)
	// result: (XORconst [1] (XOR <typ.Bool> x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMXORconst)
		v.AuxInt = 1
		v0 := b.NewValue0(v.Pos, OpARMXOR, typ.Bool)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpEqPtr(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (EqPtr x y)
	// result: (Equal (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMP, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpFMA(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (FMA x y z)
	// result: (FMULAD z x y)
	for {
		x := v_0
		y := v_1
		z := v_2
		v.reset(OpARMFMULAD)
		v.AddArg(z)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpGeq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Geq16 x y)
	// result: (GreaterEqual (CMP (SignExt16to32 x) (SignExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMGreaterEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpGeq16U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Geq16U x y)
	// result: (GreaterEqualU (CMP (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMGreaterEqualU)
		v0 := b.NewValue0(v.Pos, OpARMCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpGeq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Geq32 x y)
	// result: (GreaterEqual (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMGreaterEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMP, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpGeq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Geq32F x y)
	// result: (GreaterEqual (CMPF x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMGreaterEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMPF, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpGeq32U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Geq32U x y)
	// result: (GreaterEqualU (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMGreaterEqualU)
		v0 := b.NewValue0(v.Pos, OpARMCMP, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpGeq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Geq64F x y)
	// result: (GreaterEqual (CMPD x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMGreaterEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMPD, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpGeq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Geq8 x y)
	// result: (GreaterEqual (CMP (SignExt8to32 x) (SignExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMGreaterEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpGeq8U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Geq8U x y)
	// result: (GreaterEqualU (CMP (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMGreaterEqualU)
		v0 := b.NewValue0(v.Pos, OpARMCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpGetCallerPC(v *Value) bool {
	// match: (GetCallerPC)
	// result: (LoweredGetCallerPC)
	for {
		v.reset(OpARMLoweredGetCallerPC)
		return true
	}
}
func rewriteValueARM_OpGetCallerSP(v *Value) bool {
	// match: (GetCallerSP)
	// result: (LoweredGetCallerSP)
	for {
		v.reset(OpARMLoweredGetCallerSP)
		return true
	}
}
func rewriteValueARM_OpGetClosurePtr(v *Value) bool {
	// match: (GetClosurePtr)
	// result: (LoweredGetClosurePtr)
	for {
		v.reset(OpARMLoweredGetClosurePtr)
		return true
	}
}
func rewriteValueARM_OpGreater16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Greater16 x y)
	// result: (GreaterThan (CMP (SignExt16to32 x) (SignExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMGreaterThan)
		v0 := b.NewValue0(v.Pos, OpARMCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpGreater16U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Greater16U x y)
	// result: (GreaterThanU (CMP (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMGreaterThanU)
		v0 := b.NewValue0(v.Pos, OpARMCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpGreater32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Greater32 x y)
	// result: (GreaterThan (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMGreaterThan)
		v0 := b.NewValue0(v.Pos, OpARMCMP, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpGreater32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Greater32F x y)
	// result: (GreaterThan (CMPF x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMGreaterThan)
		v0 := b.NewValue0(v.Pos, OpARMCMPF, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpGreater32U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Greater32U x y)
	// result: (GreaterThanU (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMGreaterThanU)
		v0 := b.NewValue0(v.Pos, OpARMCMP, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpGreater64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Greater64F x y)
	// result: (GreaterThan (CMPD x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMGreaterThan)
		v0 := b.NewValue0(v.Pos, OpARMCMPD, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpGreater8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Greater8 x y)
	// result: (GreaterThan (CMP (SignExt8to32 x) (SignExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMGreaterThan)
		v0 := b.NewValue0(v.Pos, OpARMCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpGreater8U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Greater8U x y)
	// result: (GreaterThanU (CMP (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMGreaterThanU)
		v0 := b.NewValue0(v.Pos, OpARMCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpHmul32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Hmul32 x y)
	// result: (HMUL x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMHMUL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpHmul32u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Hmul32u x y)
	// result: (HMULU x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMHMULU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpInterCall(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (InterCall [argwid] entry mem)
	// result: (CALLinter [argwid] entry mem)
	for {
		argwid := v.AuxInt
		entry := v_0
		mem := v_1
		v.reset(OpARMCALLinter)
		v.AuxInt = argwid
		v.AddArg(entry)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueARM_OpIsInBounds(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (IsInBounds idx len)
	// result: (LessThanU (CMP idx len))
	for {
		idx := v_0
		len := v_1
		v.reset(OpARMLessThanU)
		v0 := b.NewValue0(v.Pos, OpARMCMP, types.TypeFlags)
		v0.AddArg(idx)
		v0.AddArg(len)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpIsNonNil(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (IsNonNil ptr)
	// result: (NotEqual (CMPconst [0] ptr))
	for {
		ptr := v_0
		v.reset(OpARMNotEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMPconst, types.TypeFlags)
		v0.AuxInt = 0
		v0.AddArg(ptr)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpIsSliceInBounds(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (IsSliceInBounds idx len)
	// result: (LessEqualU (CMP idx len))
	for {
		idx := v_0
		len := v_1
		v.reset(OpARMLessEqualU)
		v0 := b.NewValue0(v.Pos, OpARMCMP, types.TypeFlags)
		v0.AddArg(idx)
		v0.AddArg(len)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpLeq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq16 x y)
	// result: (LessEqual (CMP (SignExt16to32 x) (SignExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMLessEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpLeq16U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq16U x y)
	// result: (LessEqualU (CMP (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMLessEqualU)
		v0 := b.NewValue0(v.Pos, OpARMCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpLeq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Leq32 x y)
	// result: (LessEqual (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMLessEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMP, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpLeq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Leq32F x y)
	// result: (GreaterEqual (CMPF y x))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMGreaterEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMPF, types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpLeq32U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Leq32U x y)
	// result: (LessEqualU (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMLessEqualU)
		v0 := b.NewValue0(v.Pos, OpARMCMP, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpLeq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Leq64F x y)
	// result: (GreaterEqual (CMPD y x))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMGreaterEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMPD, types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpLeq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq8 x y)
	// result: (LessEqual (CMP (SignExt8to32 x) (SignExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMLessEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpLeq8U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq8U x y)
	// result: (LessEqualU (CMP (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMLessEqualU)
		v0 := b.NewValue0(v.Pos, OpARMCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpLess16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less16 x y)
	// result: (LessThan (CMP (SignExt16to32 x) (SignExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMLessThan)
		v0 := b.NewValue0(v.Pos, OpARMCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpLess16U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less16U x y)
	// result: (LessThanU (CMP (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMLessThanU)
		v0 := b.NewValue0(v.Pos, OpARMCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpLess32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Less32 x y)
	// result: (LessThan (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMLessThan)
		v0 := b.NewValue0(v.Pos, OpARMCMP, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpLess32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Less32F x y)
	// result: (GreaterThan (CMPF y x))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMGreaterThan)
		v0 := b.NewValue0(v.Pos, OpARMCMPF, types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpLess32U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Less32U x y)
	// result: (LessThanU (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMLessThanU)
		v0 := b.NewValue0(v.Pos, OpARMCMP, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpLess64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Less64F x y)
	// result: (GreaterThan (CMPD y x))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMGreaterThan)
		v0 := b.NewValue0(v.Pos, OpARMCMPD, types.TypeFlags)
		v0.AddArg(y)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpLess8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less8 x y)
	// result: (LessThan (CMP (SignExt8to32 x) (SignExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMLessThan)
		v0 := b.NewValue0(v.Pos, OpARMCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpLess8U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less8U x y)
	// result: (LessThanU (CMP (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMLessThanU)
		v0 := b.NewValue0(v.Pos, OpARMCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpLoad(v *Value) bool {
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
		v.reset(OpARMMOVBUload)
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
		v.reset(OpARMMOVBload)
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
		v.reset(OpARMMOVBUload)
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
		v.reset(OpARMMOVHload)
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
		v.reset(OpARMMOVHUload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is32BitInt(t) || isPtr(t))
	// result: (MOVWload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is32BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(OpARMMOVWload)
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
		v.reset(OpARMMOVFload)
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
		v.reset(OpARMMOVDload)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpLocalAddr(v *Value) bool {
	v_0 := v.Args[0]
	// match: (LocalAddr {sym} base _)
	// result: (MOVWaddr {sym} base)
	for {
		sym := v.Aux
		base := v_0
		v.reset(OpARMMOVWaddr)
		v.Aux = sym
		v.AddArg(base)
		return true
	}
}
func rewriteValueARM_OpLsh16x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x16 x y)
	// result: (CMOVWHSconst (SLL <x.Type> x (ZeroExt16to32 y)) (CMPconst [256] (ZeroExt16to32 y)) [0])
	for {
		x := v_0
		y := v_1
		v.reset(OpARMCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpARMCMPconst, types.TypeFlags)
		v2.AuxInt = 256
		v3 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueARM_OpLsh16x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Lsh16x32 x y)
	// result: (CMOVWHSconst (SLL <x.Type> x y) (CMPconst [256] y) [0])
	for {
		x := v_0
		y := v_1
		v.reset(OpARMCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARMCMPconst, types.TypeFlags)
		v1.AuxInt = 256
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM_OpLsh16x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Lsh16x64 x (Const64 [c]))
	// cond: uint64(c) < 16
	// result: (SLLconst x [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 16) {
			break
		}
		v.reset(OpARMSLLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh16x64 _ (Const64 [c]))
	// cond: uint64(c) >= 16
	// result: (Const16 [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 16) {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpLsh16x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x8 x y)
	// result: (SLL x (ZeroExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMSLL)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpLsh32x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x16 x y)
	// result: (CMOVWHSconst (SLL <x.Type> x (ZeroExt16to32 y)) (CMPconst [256] (ZeroExt16to32 y)) [0])
	for {
		x := v_0
		y := v_1
		v.reset(OpARMCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpARMCMPconst, types.TypeFlags)
		v2.AuxInt = 256
		v3 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueARM_OpLsh32x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Lsh32x32 x y)
	// result: (CMOVWHSconst (SLL <x.Type> x y) (CMPconst [256] y) [0])
	for {
		x := v_0
		y := v_1
		v.reset(OpARMCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARMCMPconst, types.TypeFlags)
		v1.AuxInt = 256
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM_OpLsh32x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Lsh32x64 x (Const64 [c]))
	// cond: uint64(c) < 32
	// result: (SLLconst x [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 32) {
			break
		}
		v.reset(OpARMSLLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh32x64 _ (Const64 [c]))
	// cond: uint64(c) >= 32
	// result: (Const32 [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 32) {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpLsh32x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x8 x y)
	// result: (SLL x (ZeroExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMSLL)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpLsh8x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x16 x y)
	// result: (CMOVWHSconst (SLL <x.Type> x (ZeroExt16to32 y)) (CMPconst [256] (ZeroExt16to32 y)) [0])
	for {
		x := v_0
		y := v_1
		v.reset(OpARMCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpARMCMPconst, types.TypeFlags)
		v2.AuxInt = 256
		v3 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueARM_OpLsh8x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Lsh8x32 x y)
	// result: (CMOVWHSconst (SLL <x.Type> x y) (CMPconst [256] y) [0])
	for {
		x := v_0
		y := v_1
		v.reset(OpARMCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpARMSLL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARMCMPconst, types.TypeFlags)
		v1.AuxInt = 256
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM_OpLsh8x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Lsh8x64 x (Const64 [c]))
	// cond: uint64(c) < 8
	// result: (SLLconst x [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 8) {
			break
		}
		v.reset(OpARMSLLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Lsh8x64 _ (Const64 [c]))
	// cond: uint64(c) >= 8
	// result: (Const8 [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 8) {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpLsh8x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x8 x y)
	// result: (SLL x (ZeroExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMSLL)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpMod16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod16 x y)
	// result: (Mod32 (SignExt16to32 x) (SignExt16to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMod32)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM_OpMod16u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod16u x y)
	// result: (Mod32u (ZeroExt16to32 x) (ZeroExt16to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMod32u)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM_OpMod32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod32 x y)
	// result: (SUB (XOR <typ.UInt32> (Select1 <typ.UInt32> (CALLudiv (SUB <typ.UInt32> (XOR <typ.UInt32> x (Signmask x)) (Signmask x)) (SUB <typ.UInt32> (XOR <typ.UInt32> y (Signmask y)) (Signmask y)))) (Signmask x)) (Signmask x))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMSUB)
		v0 := b.NewValue0(v.Pos, OpARMXOR, typ.UInt32)
		v1 := b.NewValue0(v.Pos, OpSelect1, typ.UInt32)
		v2 := b.NewValue0(v.Pos, OpARMCALLudiv, types.NewTuple(typ.UInt32, typ.UInt32))
		v3 := b.NewValue0(v.Pos, OpARMSUB, typ.UInt32)
		v4 := b.NewValue0(v.Pos, OpARMXOR, typ.UInt32)
		v4.AddArg(x)
		v5 := b.NewValue0(v.Pos, OpSignmask, typ.Int32)
		v5.AddArg(x)
		v4.AddArg(v5)
		v3.AddArg(v4)
		v6 := b.NewValue0(v.Pos, OpSignmask, typ.Int32)
		v6.AddArg(x)
		v3.AddArg(v6)
		v2.AddArg(v3)
		v7 := b.NewValue0(v.Pos, OpARMSUB, typ.UInt32)
		v8 := b.NewValue0(v.Pos, OpARMXOR, typ.UInt32)
		v8.AddArg(y)
		v9 := b.NewValue0(v.Pos, OpSignmask, typ.Int32)
		v9.AddArg(y)
		v8.AddArg(v9)
		v7.AddArg(v8)
		v10 := b.NewValue0(v.Pos, OpSignmask, typ.Int32)
		v10.AddArg(y)
		v7.AddArg(v10)
		v2.AddArg(v7)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v11 := b.NewValue0(v.Pos, OpSignmask, typ.Int32)
		v11.AddArg(x)
		v0.AddArg(v11)
		v.AddArg(v0)
		v12 := b.NewValue0(v.Pos, OpSignmask, typ.Int32)
		v12.AddArg(x)
		v.AddArg(v12)
		return true
	}
}
func rewriteValueARM_OpMod32u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod32u x y)
	// result: (Select1 <typ.UInt32> (CALLudiv x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSelect1)
		v.Type = typ.UInt32
		v0 := b.NewValue0(v.Pos, OpARMCALLudiv, types.NewTuple(typ.UInt32, typ.UInt32))
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpMod8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod8 x y)
	// result: (Mod32 (SignExt8to32 x) (SignExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMod32)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM_OpMod8u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod8u x y)
	// result: (Mod32u (ZeroExt8to32 x) (ZeroExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMod32u)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM_OpMove(v *Value) bool {
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
	// result: (MOVBstore dst (MOVBUload src mem) mem)
	for {
		if v.AuxInt != 1 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpARMMOVBstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpARMMOVBUload, typ.UInt8)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [2] {t} dst src mem)
	// cond: t.(*types.Type).Alignment()%2 == 0
	// result: (MOVHstore dst (MOVHUload src mem) mem)
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
		v.reset(OpARMMOVHstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpARMMOVHUload, typ.UInt16)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [2] dst src mem)
	// result: (MOVBstore [1] dst (MOVBUload [1] src mem) (MOVBstore dst (MOVBUload src mem) mem))
	for {
		if v.AuxInt != 2 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpARMMOVBstore)
		v.AuxInt = 1
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpARMMOVBUload, typ.UInt8)
		v0.AuxInt = 1
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARMMOVBstore, types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpARMMOVBUload, typ.UInt8)
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
		v.reset(OpARMMOVWstore)
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpARMMOVWload, typ.UInt32)
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Move [4] {t} dst src mem)
	// cond: t.(*types.Type).Alignment()%2 == 0
	// result: (MOVHstore [2] dst (MOVHUload [2] src mem) (MOVHstore dst (MOVHUload src mem) mem))
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
		v.reset(OpARMMOVHstore)
		v.AuxInt = 2
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpARMMOVHUload, typ.UInt16)
		v0.AuxInt = 2
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARMMOVHstore, types.TypeMem)
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpARMMOVHUload, typ.UInt16)
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Move [4] dst src mem)
	// result: (MOVBstore [3] dst (MOVBUload [3] src mem) (MOVBstore [2] dst (MOVBUload [2] src mem) (MOVBstore [1] dst (MOVBUload [1] src mem) (MOVBstore dst (MOVBUload src mem) mem))))
	for {
		if v.AuxInt != 4 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpARMMOVBstore)
		v.AuxInt = 3
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpARMMOVBUload, typ.UInt8)
		v0.AuxInt = 3
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARMMOVBstore, types.TypeMem)
		v1.AuxInt = 2
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpARMMOVBUload, typ.UInt8)
		v2.AuxInt = 2
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARMMOVBstore, types.TypeMem)
		v3.AuxInt = 1
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpARMMOVBUload, typ.UInt8)
		v4.AuxInt = 1
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpARMMOVBstore, types.TypeMem)
		v5.AddArg(dst)
		v6 := b.NewValue0(v.Pos, OpARMMOVBUload, typ.UInt8)
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
	// result: (MOVBstore [2] dst (MOVBUload [2] src mem) (MOVBstore [1] dst (MOVBUload [1] src mem) (MOVBstore dst (MOVBUload src mem) mem)))
	for {
		if v.AuxInt != 3 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpARMMOVBstore)
		v.AuxInt = 2
		v.AddArg(dst)
		v0 := b.NewValue0(v.Pos, OpARMMOVBUload, typ.UInt8)
		v0.AuxInt = 2
		v0.AddArg(src)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARMMOVBstore, types.TypeMem)
		v1.AuxInt = 1
		v1.AddArg(dst)
		v2 := b.NewValue0(v.Pos, OpARMMOVBUload, typ.UInt8)
		v2.AuxInt = 1
		v2.AddArg(src)
		v2.AddArg(mem)
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARMMOVBstore, types.TypeMem)
		v3.AddArg(dst)
		v4 := b.NewValue0(v.Pos, OpARMMOVBUload, typ.UInt8)
		v4.AddArg(src)
		v4.AddArg(mem)
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Move [s] {t} dst src mem)
	// cond: s%4 == 0 && s > 4 && s <= 512 && t.(*types.Type).Alignment()%4 == 0 && !config.noDuffDevice
	// result: (DUFFCOPY [8 * (128 - s/4)] dst src mem)
	for {
		s := v.AuxInt
		t := v.Aux
		dst := v_0
		src := v_1
		mem := v_2
		if !(s%4 == 0 && s > 4 && s <= 512 && t.(*types.Type).Alignment()%4 == 0 && !config.noDuffDevice) {
			break
		}
		v.reset(OpARMDUFFCOPY)
		v.AuxInt = 8 * (128 - s/4)
		v.AddArg(dst)
		v.AddArg(src)
		v.AddArg(mem)
		return true
	}
	// match: (Move [s] {t} dst src mem)
	// cond: (s > 512 || config.noDuffDevice) || t.(*types.Type).Alignment()%4 != 0
	// result: (LoweredMove [t.(*types.Type).Alignment()] dst src (ADDconst <src.Type> src [s-moveSize(t.(*types.Type).Alignment(), config)]) mem)
	for {
		s := v.AuxInt
		t := v.Aux
		dst := v_0
		src := v_1
		mem := v_2
		if !((s > 512 || config.noDuffDevice) || t.(*types.Type).Alignment()%4 != 0) {
			break
		}
		v.reset(OpARMLoweredMove)
		v.AuxInt = t.(*types.Type).Alignment()
		v.AddArg(dst)
		v.AddArg(src)
		v0 := b.NewValue0(v.Pos, OpARMADDconst, src.Type)
		v0.AuxInt = s - moveSize(t.(*types.Type).Alignment(), config)
		v0.AddArg(src)
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpMul16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Mul16 x y)
	// result: (MUL x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMMUL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpMul32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Mul32 x y)
	// result: (MUL x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMMUL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpMul32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Mul32F x y)
	// result: (MULF x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMMULF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpMul32uhilo(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Mul32uhilo x y)
	// result: (MULLU x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMMULLU)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpMul64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Mul64F x y)
	// result: (MULD x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMMULD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpMul8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Mul8 x y)
	// result: (MUL x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMMUL)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpNeg16(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Neg16 x)
	// result: (RSBconst [0] x)
	for {
		x := v_0
		v.reset(OpARMRSBconst)
		v.AuxInt = 0
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpNeg32(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Neg32 x)
	// result: (RSBconst [0] x)
	for {
		x := v_0
		v.reset(OpARMRSBconst)
		v.AuxInt = 0
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpNeg32F(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Neg32F x)
	// result: (NEGF x)
	for {
		x := v_0
		v.reset(OpARMNEGF)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpNeg64F(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Neg64F x)
	// result: (NEGD x)
	for {
		x := v_0
		v.reset(OpARMNEGD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpNeg8(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Neg8 x)
	// result: (RSBconst [0] x)
	for {
		x := v_0
		v.reset(OpARMRSBconst)
		v.AuxInt = 0
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpNeq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq16 x y)
	// result: (NotEqual (CMP (ZeroExt16to32 x) (ZeroExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMNotEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpNeq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Neq32 x y)
	// result: (NotEqual (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMNotEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMP, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpNeq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Neq32F x y)
	// result: (NotEqual (CMPF x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMNotEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMPF, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpNeq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Neq64F x y)
	// result: (NotEqual (CMPD x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMNotEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMPD, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpNeq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq8 x y)
	// result: (NotEqual (CMP (ZeroExt8to32 x) (ZeroExt8to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMNotEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpNeqB(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (NeqB x y)
	// result: (XOR x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMXOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpNeqPtr(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (NeqPtr x y)
	// result: (NotEqual (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMNotEqual)
		v0 := b.NewValue0(v.Pos, OpARMCMP, types.TypeFlags)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpNilCheck(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (NilCheck ptr mem)
	// result: (LoweredNilCheck ptr mem)
	for {
		ptr := v_0
		mem := v_1
		v.reset(OpARMLoweredNilCheck)
		v.AddArg(ptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueARM_OpNot(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Not x)
	// result: (XORconst [1] x)
	for {
		x := v_0
		v.reset(OpARMXORconst)
		v.AuxInt = 1
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpOffPtr(v *Value) bool {
	v_0 := v.Args[0]
	// match: (OffPtr [off] ptr:(SP))
	// result: (MOVWaddr [off] ptr)
	for {
		off := v.AuxInt
		ptr := v_0
		if ptr.Op != OpSP {
			break
		}
		v.reset(OpARMMOVWaddr)
		v.AuxInt = off
		v.AddArg(ptr)
		return true
	}
	// match: (OffPtr [off] ptr)
	// result: (ADDconst [off] ptr)
	for {
		off := v.AuxInt
		ptr := v_0
		v.reset(OpARMADDconst)
		v.AuxInt = off
		v.AddArg(ptr)
		return true
	}
}
func rewriteValueARM_OpOr16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Or16 x y)
	// result: (OR x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpOr32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Or32 x y)
	// result: (OR x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpOr8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Or8 x y)
	// result: (OR x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpOrB(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (OrB x y)
	// result: (OR x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpPanicBounds(v *Value) bool {
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
		v.reset(OpARMLoweredPanicBoundsA)
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
		v.reset(OpARMLoweredPanicBoundsB)
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
		v.reset(OpARMLoweredPanicBoundsC)
		v.AuxInt = kind
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpPanicExtend(v *Value) bool {
	v_3 := v.Args[3]
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (PanicExtend [kind] hi lo y mem)
	// cond: boundsABI(kind) == 0
	// result: (LoweredPanicExtendA [kind] hi lo y mem)
	for {
		kind := v.AuxInt
		hi := v_0
		lo := v_1
		y := v_2
		mem := v_3
		if !(boundsABI(kind) == 0) {
			break
		}
		v.reset(OpARMLoweredPanicExtendA)
		v.AuxInt = kind
		v.AddArg(hi)
		v.AddArg(lo)
		v.AddArg(y)
		v.AddArg(mem)
		return true
	}
	// match: (PanicExtend [kind] hi lo y mem)
	// cond: boundsABI(kind) == 1
	// result: (LoweredPanicExtendB [kind] hi lo y mem)
	for {
		kind := v.AuxInt
		hi := v_0
		lo := v_1
		y := v_2
		mem := v_3
		if !(boundsABI(kind) == 1) {
			break
		}
		v.reset(OpARMLoweredPanicExtendB)
		v.AuxInt = kind
		v.AddArg(hi)
		v.AddArg(lo)
		v.AddArg(y)
		v.AddArg(mem)
		return true
	}
	// match: (PanicExtend [kind] hi lo y mem)
	// cond: boundsABI(kind) == 2
	// result: (LoweredPanicExtendC [kind] hi lo y mem)
	for {
		kind := v.AuxInt
		hi := v_0
		lo := v_1
		y := v_2
		mem := v_3
		if !(boundsABI(kind) == 2) {
			break
		}
		v.reset(OpARMLoweredPanicExtendC)
		v.AuxInt = kind
		v.AddArg(hi)
		v.AddArg(lo)
		v.AddArg(y)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpRotateLeft16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft16 <t> x (MOVWconst [c]))
	// result: (Or16 (Lsh16x32 <t> x (MOVWconst [c&15])) (Rsh16Ux32 <t> x (MOVWconst [-c&15])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpOr16)
		v0 := b.NewValue0(v.Pos, OpLsh16x32, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
		v1.AuxInt = c & 15
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpRsh16Ux32, t)
		v2.AddArg(x)
		v3 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
		v3.AuxInt = -c & 15
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
	return false
}
func rewriteValueARM_OpRotateLeft32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (RotateLeft32 x (MOVWconst [c]))
	// result: (SRRconst [-c&31] x)
	for {
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpARMSRRconst)
		v.AuxInt = -c & 31
		v.AddArg(x)
		return true
	}
	// match: (RotateLeft32 x y)
	// result: (SRR x (RSBconst [0] <y.Type> y))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMSRR)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpARMRSBconst, y.Type)
		v0.AuxInt = 0
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpRotateLeft8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft8 <t> x (MOVWconst [c]))
	// result: (Or8 (Lsh8x32 <t> x (MOVWconst [c&7])) (Rsh8Ux32 <t> x (MOVWconst [-c&7])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpARMMOVWconst {
			break
		}
		c := v_1.AuxInt
		v.reset(OpOr8)
		v0 := b.NewValue0(v.Pos, OpLsh8x32, t)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
		v1.AuxInt = c & 7
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpRsh8Ux32, t)
		v2.AddArg(x)
		v3 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
		v3.AuxInt = -c & 7
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
	return false
}
func rewriteValueARM_OpRound32F(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Round32F x)
	// result: x
	for {
		x := v_0
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpRound64F(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Round64F x)
	// result: x
	for {
		x := v_0
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpRsh16Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux16 x y)
	// result: (CMOVWHSconst (SRL <x.Type> (ZeroExt16to32 x) (ZeroExt16to32 y)) (CMPconst [256] (ZeroExt16to32 y)) [0])
	for {
		x := v_0
		y := v_1
		v.reset(OpARMCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpARMCMPconst, types.TypeFlags)
		v3.AuxInt = 256
		v4 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueARM_OpRsh16Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux32 x y)
	// result: (CMOVWHSconst (SRL <x.Type> (ZeroExt16to32 x) y) (CMPconst [256] y) [0])
	for {
		x := v_0
		y := v_1
		v.reset(OpARMCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpARMCMPconst, types.TypeFlags)
		v2.AuxInt = 256
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueARM_OpRsh16Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux64 x (Const64 [c]))
	// cond: uint64(c) < 16
	// result: (SRLconst (SLLconst <typ.UInt32> x [16]) [c+16])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 16) {
			break
		}
		v.reset(OpARMSRLconst)
		v.AuxInt = c + 16
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, typ.UInt32)
		v0.AuxInt = 16
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16Ux64 _ (Const64 [c]))
	// cond: uint64(c) >= 16
	// result: (Const16 [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 16) {
			break
		}
		v.reset(OpConst16)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpRsh16Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux8 x y)
	// result: (SRL (ZeroExt16to32 x) (ZeroExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMSRL)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM_OpRsh16x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x16 x y)
	// result: (SRAcond (SignExt16to32 x) (ZeroExt16to32 y) (CMPconst [256] (ZeroExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMSRAcond)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpARMCMPconst, types.TypeFlags)
		v2.AuxInt = 256
		v3 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueARM_OpRsh16x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x32 x y)
	// result: (SRAcond (SignExt16to32 x) y (CMPconst [256] y))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMSRAcond)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpARMCMPconst, types.TypeFlags)
		v1.AuxInt = 256
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM_OpRsh16x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x64 x (Const64 [c]))
	// cond: uint64(c) < 16
	// result: (SRAconst (SLLconst <typ.UInt32> x [16]) [c+16])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 16) {
			break
		}
		v.reset(OpARMSRAconst)
		v.AuxInt = c + 16
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, typ.UInt32)
		v0.AuxInt = 16
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16x64 x (Const64 [c]))
	// cond: uint64(c) >= 16
	// result: (SRAconst (SLLconst <typ.UInt32> x [16]) [31])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 16) {
			break
		}
		v.reset(OpARMSRAconst)
		v.AuxInt = 31
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, typ.UInt32)
		v0.AuxInt = 16
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueARM_OpRsh16x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x8 x y)
	// result: (SRA (SignExt16to32 x) (ZeroExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMSRA)
		v0 := b.NewValue0(v.Pos, OpSignExt16to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM_OpRsh32Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux16 x y)
	// result: (CMOVWHSconst (SRL <x.Type> x (ZeroExt16to32 y)) (CMPconst [256] (ZeroExt16to32 y)) [0])
	for {
		x := v_0
		y := v_1
		v.reset(OpARMCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v0.AddArg(v1)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpARMCMPconst, types.TypeFlags)
		v2.AuxInt = 256
		v3 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueARM_OpRsh32Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh32Ux32 x y)
	// result: (CMOVWHSconst (SRL <x.Type> x y) (CMPconst [256] y) [0])
	for {
		x := v_0
		y := v_1
		v.reset(OpARMCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v0.AddArg(x)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARMCMPconst, types.TypeFlags)
		v1.AuxInt = 256
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM_OpRsh32Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Rsh32Ux64 x (Const64 [c]))
	// cond: uint64(c) < 32
	// result: (SRLconst x [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 32) {
			break
		}
		v.reset(OpARMSRLconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh32Ux64 _ (Const64 [c]))
	// cond: uint64(c) >= 32
	// result: (Const32 [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 32) {
			break
		}
		v.reset(OpConst32)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpRsh32Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux8 x y)
	// result: (SRL x (ZeroExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMSRL)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpRsh32x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x16 x y)
	// result: (SRAcond x (ZeroExt16to32 y) (CMPconst [256] (ZeroExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMSRAcond)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARMCMPconst, types.TypeFlags)
		v1.AuxInt = 256
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v1.AddArg(v2)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM_OpRsh32x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Rsh32x32 x y)
	// result: (SRAcond x y (CMPconst [256] y))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMSRAcond)
		v.AddArg(x)
		v.AddArg(y)
		v0 := b.NewValue0(v.Pos, OpARMCMPconst, types.TypeFlags)
		v0.AuxInt = 256
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpRsh32x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Rsh32x64 x (Const64 [c]))
	// cond: uint64(c) < 32
	// result: (SRAconst x [c])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 32) {
			break
		}
		v.reset(OpARMSRAconst)
		v.AuxInt = c
		v.AddArg(x)
		return true
	}
	// match: (Rsh32x64 x (Const64 [c]))
	// cond: uint64(c) >= 32
	// result: (SRAconst x [31])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 32) {
			break
		}
		v.reset(OpARMSRAconst)
		v.AuxInt = 31
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueARM_OpRsh32x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x8 x y)
	// result: (SRA x (ZeroExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMSRA)
		v.AddArg(x)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpRsh8Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux16 x y)
	// result: (CMOVWHSconst (SRL <x.Type> (ZeroExt8to32 x) (ZeroExt16to32 y)) (CMPconst [256] (ZeroExt16to32 y)) [0])
	for {
		x := v_0
		y := v_1
		v.reset(OpARMCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v2.AddArg(y)
		v0.AddArg(v2)
		v.AddArg(v0)
		v3 := b.NewValue0(v.Pos, OpARMCMPconst, types.TypeFlags)
		v3.AuxInt = 256
		v4 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v4.AddArg(y)
		v3.AddArg(v4)
		v.AddArg(v3)
		return true
	}
}
func rewriteValueARM_OpRsh8Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux32 x y)
	// result: (CMOVWHSconst (SRL <x.Type> (ZeroExt8to32 x) y) (CMPconst [256] y) [0])
	for {
		x := v_0
		y := v_1
		v.reset(OpARMCMOVWHSconst)
		v.AuxInt = 0
		v0 := b.NewValue0(v.Pos, OpARMSRL, x.Type)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(x)
		v0.AddArg(v1)
		v0.AddArg(y)
		v.AddArg(v0)
		v2 := b.NewValue0(v.Pos, OpARMCMPconst, types.TypeFlags)
		v2.AuxInt = 256
		v2.AddArg(y)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueARM_OpRsh8Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux64 x (Const64 [c]))
	// cond: uint64(c) < 8
	// result: (SRLconst (SLLconst <typ.UInt32> x [24]) [c+24])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 8) {
			break
		}
		v.reset(OpARMSRLconst)
		v.AuxInt = c + 24
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, typ.UInt32)
		v0.AuxInt = 24
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8Ux64 _ (Const64 [c]))
	// cond: uint64(c) >= 8
	// result: (Const8 [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 8) {
			break
		}
		v.reset(OpConst8)
		v.AuxInt = 0
		return true
	}
	return false
}
func rewriteValueARM_OpRsh8Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux8 x y)
	// result: (SRL (ZeroExt8to32 x) (ZeroExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMSRL)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM_OpRsh8x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x16 x y)
	// result: (SRAcond (SignExt8to32 x) (ZeroExt16to32 y) (CMPconst [256] (ZeroExt16to32 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMSRAcond)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		v2 := b.NewValue0(v.Pos, OpARMCMPconst, types.TypeFlags)
		v2.AuxInt = 256
		v3 := b.NewValue0(v.Pos, OpZeroExt16to32, typ.UInt32)
		v3.AddArg(y)
		v2.AddArg(v3)
		v.AddArg(v2)
		return true
	}
}
func rewriteValueARM_OpRsh8x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x32 x y)
	// result: (SRAcond (SignExt8to32 x) y (CMPconst [256] y))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMSRAcond)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v.AddArg(y)
		v1 := b.NewValue0(v.Pos, OpARMCMPconst, types.TypeFlags)
		v1.AuxInt = 256
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM_OpRsh8x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x64 x (Const64 [c]))
	// cond: uint64(c) < 8
	// result: (SRAconst (SLLconst <typ.UInt32> x [24]) [c+24])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) < 8) {
			break
		}
		v.reset(OpARMSRAconst)
		v.AuxInt = c + 24
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, typ.UInt32)
		v0.AuxInt = 24
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8x64 x (Const64 [c]))
	// cond: uint64(c) >= 8
	// result: (SRAconst (SLLconst <typ.UInt32> x [24]) [31])
	for {
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := v_1.AuxInt
		if !(uint64(c) >= 8) {
			break
		}
		v.reset(OpARMSRAconst)
		v.AuxInt = 31
		v0 := b.NewValue0(v.Pos, OpARMSLLconst, typ.UInt32)
		v0.AuxInt = 24
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueARM_OpRsh8x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x8 x y)
	// result: (SRA (SignExt8to32 x) (ZeroExt8to32 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpARMSRA)
		v0 := b.NewValue0(v.Pos, OpSignExt8to32, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to32, typ.UInt32)
		v1.AddArg(y)
		v.AddArg(v1)
		return true
	}
}
func rewriteValueARM_OpSelect0(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Select0 (CALLudiv x (MOVWconst [1])))
	// result: x
	for {
		if v_0.Op != OpARMCALLudiv {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpARMMOVWconst || v_0_1.AuxInt != 1 {
			break
		}
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
	// match: (Select0 (CALLudiv x (MOVWconst [c])))
	// cond: isPowerOfTwo(c)
	// result: (SRLconst [log2(c)] x)
	for {
		if v_0.Op != OpARMCALLudiv {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpARMMOVWconst {
			break
		}
		c := v_0_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpARMSRLconst)
		v.AuxInt = log2(c)
		v.AddArg(x)
		return true
	}
	// match: (Select0 (CALLudiv (MOVWconst [c]) (MOVWconst [d])))
	// result: (MOVWconst [int64(int32(uint32(c)/uint32(d)))])
	for {
		if v_0.Op != OpARMCALLudiv {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpARMMOVWconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = int64(int32(uint32(c) / uint32(d)))
		return true
	}
	return false
}
func rewriteValueARM_OpSelect1(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Select1 (CALLudiv _ (MOVWconst [1])))
	// result: (MOVWconst [0])
	for {
		if v_0.Op != OpARMCALLudiv {
			break
		}
		_ = v_0.Args[1]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpARMMOVWconst || v_0_1.AuxInt != 1 {
			break
		}
		v.reset(OpARMMOVWconst)
		v.AuxInt = 0
		return true
	}
	// match: (Select1 (CALLudiv x (MOVWconst [c])))
	// cond: isPowerOfTwo(c)
	// result: (ANDconst [c-1] x)
	for {
		if v_0.Op != OpARMCALLudiv {
			break
		}
		_ = v_0.Args[1]
		x := v_0.Args[0]
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpARMMOVWconst {
			break
		}
		c := v_0_1.AuxInt
		if !(isPowerOfTwo(c)) {
			break
		}
		v.reset(OpARMANDconst)
		v.AuxInt = c - 1
		v.AddArg(x)
		return true
	}
	// match: (Select1 (CALLudiv (MOVWconst [c]) (MOVWconst [d])))
	// result: (MOVWconst [int64(int32(uint32(c)%uint32(d)))])
	for {
		if v_0.Op != OpARMCALLudiv {
			break
		}
		_ = v_0.Args[1]
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpARMMOVWconst {
			break
		}
		c := v_0_0.AuxInt
		v_0_1 := v_0.Args[1]
		if v_0_1.Op != OpARMMOVWconst {
			break
		}
		d := v_0_1.AuxInt
		v.reset(OpARMMOVWconst)
		v.AuxInt = int64(int32(uint32(c) % uint32(d)))
		return true
	}
	return false
}
func rewriteValueARM_OpSignExt16to32(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SignExt16to32 x)
	// result: (MOVHreg x)
	for {
		x := v_0
		v.reset(OpARMMOVHreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpSignExt8to16(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SignExt8to16 x)
	// result: (MOVBreg x)
	for {
		x := v_0
		v.reset(OpARMMOVBreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpSignExt8to32(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SignExt8to32 x)
	// result: (MOVBreg x)
	for {
		x := v_0
		v.reset(OpARMMOVBreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpSignmask(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Signmask x)
	// result: (SRAconst x [31])
	for {
		x := v_0
		v.reset(OpARMSRAconst)
		v.AuxInt = 31
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpSlicemask(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (Slicemask <t> x)
	// result: (SRAconst (RSBconst <t> [0] x) [31])
	for {
		t := v.Type
		x := v_0
		v.reset(OpARMSRAconst)
		v.AuxInt = 31
		v0 := b.NewValue0(v.Pos, OpARMRSBconst, t)
		v0.AuxInt = 0
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueARM_OpSqrt(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Sqrt x)
	// result: (SQRTD x)
	for {
		x := v_0
		v.reset(OpARMSQRTD)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpStaticCall(v *Value) bool {
	v_0 := v.Args[0]
	// match: (StaticCall [argwid] {target} mem)
	// result: (CALLstatic [argwid] {target} mem)
	for {
		argwid := v.AuxInt
		target := v.Aux
		mem := v_0
		v.reset(OpARMCALLstatic)
		v.AuxInt = argwid
		v.Aux = target
		v.AddArg(mem)
		return true
	}
}
func rewriteValueARM_OpStore(v *Value) bool {
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
		v.reset(OpARMMOVBstore)
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
		v.reset(OpARMMOVHstore)
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
		v.reset(OpARMMOVWstore)
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
		v.reset(OpARMMOVFstore)
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
		v.reset(OpARMMOVDstore)
		v.AddArg(ptr)
		v.AddArg(val)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpSub16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Sub16 x y)
	// result: (SUB x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpSub32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Sub32 x y)
	// result: (SUB x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpSub32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Sub32F x y)
	// result: (SUBF x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMSUBF)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpSub32carry(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Sub32carry x y)
	// result: (SUBS x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMSUBS)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpSub32withcarry(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Sub32withcarry x y c)
	// result: (SBC x y c)
	for {
		x := v_0
		y := v_1
		c := v_2
		v.reset(OpARMSBC)
		v.AddArg(x)
		v.AddArg(y)
		v.AddArg(c)
		return true
	}
}
func rewriteValueARM_OpSub64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Sub64F x y)
	// result: (SUBD x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMSUBD)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpSub8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Sub8 x y)
	// result: (SUB x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpSubPtr(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SubPtr x y)
	// result: (SUB x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMSUB)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpTrunc16to8(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Trunc16to8 x)
	// result: x
	for {
		x := v_0
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpTrunc32to16(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Trunc32to16 x)
	// result: x
	for {
		x := v_0
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpTrunc32to8(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Trunc32to8 x)
	// result: x
	for {
		x := v_0
		v.reset(OpCopy)
		v.Type = x.Type
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpWB(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (WB {fn} destptr srcptr mem)
	// result: (LoweredWB {fn} destptr srcptr mem)
	for {
		fn := v.Aux
		destptr := v_0
		srcptr := v_1
		mem := v_2
		v.reset(OpARMLoweredWB)
		v.Aux = fn
		v.AddArg(destptr)
		v.AddArg(srcptr)
		v.AddArg(mem)
		return true
	}
}
func rewriteValueARM_OpXor16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Xor16 x y)
	// result: (XOR x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMXOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpXor32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Xor32 x y)
	// result: (XOR x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMXOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpXor8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Xor8 x y)
	// result: (XOR x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpARMXOR)
		v.AddArg(x)
		v.AddArg(y)
		return true
	}
}
func rewriteValueARM_OpZero(v *Value) bool {
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
	// result: (MOVBstore ptr (MOVWconst [0]) mem)
	for {
		if v.AuxInt != 1 {
			break
		}
		ptr := v_0
		mem := v_1
		v.reset(OpARMMOVBstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [2] {t} ptr mem)
	// cond: t.(*types.Type).Alignment()%2 == 0
	// result: (MOVHstore ptr (MOVWconst [0]) mem)
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
		v.reset(OpARMMOVHstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [2] ptr mem)
	// result: (MOVBstore [1] ptr (MOVWconst [0]) (MOVBstore [0] ptr (MOVWconst [0]) mem))
	for {
		if v.AuxInt != 2 {
			break
		}
		ptr := v_0
		mem := v_1
		v.reset(OpARMMOVBstore)
		v.AuxInt = 1
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARMMOVBstore, types.TypeMem)
		v1.AuxInt = 0
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [4] {t} ptr mem)
	// cond: t.(*types.Type).Alignment()%4 == 0
	// result: (MOVWstore ptr (MOVWconst [0]) mem)
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
		v.reset(OpARMMOVWstore)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [4] {t} ptr mem)
	// cond: t.(*types.Type).Alignment()%2 == 0
	// result: (MOVHstore [2] ptr (MOVWconst [0]) (MOVHstore [0] ptr (MOVWconst [0]) mem))
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
		v.reset(OpARMMOVHstore)
		v.AuxInt = 2
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARMMOVHstore, types.TypeMem)
		v1.AuxInt = 0
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [4] ptr mem)
	// result: (MOVBstore [3] ptr (MOVWconst [0]) (MOVBstore [2] ptr (MOVWconst [0]) (MOVBstore [1] ptr (MOVWconst [0]) (MOVBstore [0] ptr (MOVWconst [0]) mem))))
	for {
		if v.AuxInt != 4 {
			break
		}
		ptr := v_0
		mem := v_1
		v.reset(OpARMMOVBstore)
		v.AuxInt = 3
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARMMOVBstore, types.TypeMem)
		v1.AuxInt = 2
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARMMOVBstore, types.TypeMem)
		v3.AuxInt = 1
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v5 := b.NewValue0(v.Pos, OpARMMOVBstore, types.TypeMem)
		v5.AuxInt = 0
		v5.AddArg(ptr)
		v6 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
		v6.AuxInt = 0
		v5.AddArg(v6)
		v5.AddArg(mem)
		v3.AddArg(v5)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [3] ptr mem)
	// result: (MOVBstore [2] ptr (MOVWconst [0]) (MOVBstore [1] ptr (MOVWconst [0]) (MOVBstore [0] ptr (MOVWconst [0]) mem)))
	for {
		if v.AuxInt != 3 {
			break
		}
		ptr := v_0
		mem := v_1
		v.reset(OpARMMOVBstore)
		v.AuxInt = 2
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARMMOVBstore, types.TypeMem)
		v1.AuxInt = 1
		v1.AddArg(ptr)
		v2 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
		v2.AuxInt = 0
		v1.AddArg(v2)
		v3 := b.NewValue0(v.Pos, OpARMMOVBstore, types.TypeMem)
		v3.AuxInt = 0
		v3.AddArg(ptr)
		v4 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
		v4.AuxInt = 0
		v3.AddArg(v4)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Zero [s] {t} ptr mem)
	// cond: s%4 == 0 && s > 4 && s <= 512 && t.(*types.Type).Alignment()%4 == 0 && !config.noDuffDevice
	// result: (DUFFZERO [4 * (128 - s/4)] ptr (MOVWconst [0]) mem)
	for {
		s := v.AuxInt
		t := v.Aux
		ptr := v_0
		mem := v_1
		if !(s%4 == 0 && s > 4 && s <= 512 && t.(*types.Type).Alignment()%4 == 0 && !config.noDuffDevice) {
			break
		}
		v.reset(OpARMDUFFZERO)
		v.AuxInt = 4 * (128 - s/4)
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
		v0.AuxInt = 0
		v.AddArg(v0)
		v.AddArg(mem)
		return true
	}
	// match: (Zero [s] {t} ptr mem)
	// cond: (s > 512 || config.noDuffDevice) || t.(*types.Type).Alignment()%4 != 0
	// result: (LoweredZero [t.(*types.Type).Alignment()] ptr (ADDconst <ptr.Type> ptr [s-moveSize(t.(*types.Type).Alignment(), config)]) (MOVWconst [0]) mem)
	for {
		s := v.AuxInt
		t := v.Aux
		ptr := v_0
		mem := v_1
		if !((s > 512 || config.noDuffDevice) || t.(*types.Type).Alignment()%4 != 0) {
			break
		}
		v.reset(OpARMLoweredZero)
		v.AuxInt = t.(*types.Type).Alignment()
		v.AddArg(ptr)
		v0 := b.NewValue0(v.Pos, OpARMADDconst, ptr.Type)
		v0.AuxInt = s - moveSize(t.(*types.Type).Alignment(), config)
		v0.AddArg(ptr)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpARMMOVWconst, typ.UInt32)
		v1.AuxInt = 0
		v.AddArg(v1)
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueARM_OpZeroExt16to32(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ZeroExt16to32 x)
	// result: (MOVHUreg x)
	for {
		x := v_0
		v.reset(OpARMMOVHUreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpZeroExt8to16(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ZeroExt8to16 x)
	// result: (MOVBUreg x)
	for {
		x := v_0
		v.reset(OpARMMOVBUreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpZeroExt8to32(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ZeroExt8to32 x)
	// result: (MOVBUreg x)
	for {
		x := v_0
		v.reset(OpARMMOVBUreg)
		v.AddArg(x)
		return true
	}
}
func rewriteValueARM_OpZeromask(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Zeromask x)
	// result: (SRAconst (RSBshiftRL <typ.Int32> x x [1]) [31])
	for {
		x := v_0
		v.reset(OpARMSRAconst)
		v.AuxInt = 31
		v0 := b.NewValue0(v.Pos, OpARMRSBshiftRL, typ.Int32)
		v0.AuxInt = 1
		v0.AddArg(x)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteBlockARM(b *Block) bool {
	switch b.Kind {
	case BlockARMEQ:
		// match: (EQ (FlagEQ) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpARMFlagEQ {
			b.Reset(BlockFirst)
			return true
		}
		// match: (EQ (FlagLT_ULT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpARMFlagLT_ULT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (EQ (FlagLT_UGT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpARMFlagLT_UGT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (EQ (FlagGT_ULT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpARMFlagGT_ULT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (EQ (FlagGT_UGT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpARMFlagGT_UGT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (EQ (InvertFlags cmp) yes no)
		// result: (EQ cmp yes no)
		for b.Controls[0].Op == OpARMInvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockARMEQ)
			b.AddControl(cmp)
			return true
		}
		// match: (EQ (CMPconst [0] l:(SUB x y)) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMP x y) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUB {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMEQ)
			v0 := b.NewValue0(v_0.Pos, OpARMCMP, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(MULS x y a)) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMP a (MUL <x.Type> x y)) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMMULS {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMEQ)
			v0 := b.NewValue0(v_0.Pos, OpARMCMP, types.TypeFlags)
			v0.AddArg(a)
			v1 := b.NewValue0(v_0.Pos, OpARMMUL, x.Type)
			v1.AddArg(x)
			v1.AddArg(y)
			v0.AddArg(v1)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(SUBconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMPconst [c] x) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMEQ)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(SUBshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMPshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMEQ)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(SUBshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMPshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMEQ)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(SUBshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMPshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMEQ)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(SUBshiftLLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMPshiftLLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftLLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMEQ)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftLLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(SUBshiftRLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMPshiftRLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftRLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMEQ)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftRLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(SUBshiftRAreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMPshiftRAreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftRAreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMEQ)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftRAreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(ADD x y)) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMN x y) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADD {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				b.Reset(BlockARMEQ)
				v0 := b.NewValue0(v_0.Pos, OpARMCMN, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (EQ (CMPconst [0] l:(MULA x y a)) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMN a (MUL <x.Type> x y)) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMMULA {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMEQ)
			v0 := b.NewValue0(v_0.Pos, OpARMCMN, types.TypeFlags)
			v0.AddArg(a)
			v1 := b.NewValue0(v_0.Pos, OpARMMUL, x.Type)
			v1.AddArg(x)
			v1.AddArg(y)
			v0.AddArg(v1)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(ADDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMNconst [c] x) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMEQ)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(ADDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMNshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMEQ)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(ADDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMNshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMEQ)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(ADDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMNshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMEQ)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(ADDshiftLLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMNshiftLLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftLLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMEQ)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftLLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(ADDshiftRLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMNshiftRLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftRLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMEQ)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftRLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(ADDshiftRAreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (EQ (CMNshiftRAreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftRAreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMEQ)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftRAreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(AND x y)) yes no)
		// cond: l.Uses==1
		// result: (EQ (TST x y) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMAND {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				b.Reset(BlockARMEQ)
				v0 := b.NewValue0(v_0.Pos, OpARMTST, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (EQ (CMPconst [0] l:(ANDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (EQ (TSTconst [c] x) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMEQ)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(ANDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (TSTshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMEQ)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(ANDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (TSTshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMEQ)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(ANDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (TSTshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMEQ)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(ANDshiftLLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (EQ (TSTshiftLLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftLLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMEQ)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftLLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(ANDshiftRLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (EQ (TSTshiftRLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftRLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMEQ)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftRLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(ANDshiftRAreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (EQ (TSTshiftRAreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftRAreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMEQ)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftRAreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(XOR x y)) yes no)
		// cond: l.Uses==1
		// result: (EQ (TEQ x y) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXOR {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				b.Reset(BlockARMEQ)
				v0 := b.NewValue0(v_0.Pos, OpARMTEQ, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (EQ (CMPconst [0] l:(XORconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (EQ (TEQconst [c] x) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMEQ)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(XORshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (TEQshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMEQ)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(XORshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (TEQshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMEQ)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(XORshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (EQ (TEQshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMEQ)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(XORshiftLLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (EQ (TEQshiftLLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftLLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMEQ)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftLLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(XORshiftRLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (EQ (TEQshiftRLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftRLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMEQ)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftRLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (EQ (CMPconst [0] l:(XORshiftRAreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (EQ (TEQshiftRAreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftRAreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMEQ)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftRAreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
	case BlockARMGE:
		// match: (GE (FlagEQ) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpARMFlagEQ {
			b.Reset(BlockFirst)
			return true
		}
		// match: (GE (FlagLT_ULT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpARMFlagLT_ULT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (GE (FlagLT_UGT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpARMFlagLT_UGT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (GE (FlagGT_ULT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpARMFlagGT_ULT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (GE (FlagGT_UGT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpARMFlagGT_UGT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (GE (InvertFlags cmp) yes no)
		// result: (LE cmp yes no)
		for b.Controls[0].Op == OpARMInvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockARMLE)
			b.AddControl(cmp)
			return true
		}
		// match: (GE (CMPconst [0] l:(SUB x y)) yes no)
		// cond: l.Uses==1
		// result: (GE (CMP x y) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUB {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMP, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(MULS x y a)) yes no)
		// cond: l.Uses==1
		// result: (GE (CMP a (MUL <x.Type> x y)) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMMULS {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMP, types.TypeFlags)
			v0.AddArg(a)
			v1 := b.NewValue0(v_0.Pos, OpARMMUL, x.Type)
			v1.AddArg(x)
			v1.AddArg(y)
			v0.AddArg(v1)
			b.AddControl(v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(SUBconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (GE (CMPconst [c] x) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(SUBshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GE (CMPshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(SUBshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GE (CMPshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(SUBshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GE (CMPshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(SUBshiftLLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (GE (CMPshiftLLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftLLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftLLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(SUBshiftRLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (GE (CMPshiftRLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftRLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftRLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(SUBshiftRAreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (GE (CMPshiftRAreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftRAreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftRAreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(ADD x y)) yes no)
		// cond: l.Uses==1
		// result: (GE (CMN x y) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADD {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				b.Reset(BlockARMGE)
				v0 := b.NewValue0(v_0.Pos, OpARMCMN, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (GE (CMPconst [0] l:(MULA x y a)) yes no)
		// cond: l.Uses==1
		// result: (GE (CMN a (MUL <x.Type> x y)) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMMULA {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMN, types.TypeFlags)
			v0.AddArg(a)
			v1 := b.NewValue0(v_0.Pos, OpARMMUL, x.Type)
			v1.AddArg(x)
			v1.AddArg(y)
			v0.AddArg(v1)
			b.AddControl(v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(ADDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (GE (CMNconst [c] x) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(ADDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GE (CMNshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(ADDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GE (CMNshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(ADDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GE (CMNshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(ADDshiftLLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (GE (CMNshiftLLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftLLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftLLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(ADDshiftRLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (GE (CMNshiftRLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftRLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftRLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(ADDshiftRAreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (GE (CMNshiftRAreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftRAreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftRAreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(AND x y)) yes no)
		// cond: l.Uses==1
		// result: (GE (TST x y) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMAND {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				b.Reset(BlockARMGE)
				v0 := b.NewValue0(v_0.Pos, OpARMTST, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (GE (CMPconst [0] l:(ANDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (GE (TSTconst [c] x) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGE)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(ANDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GE (TSTshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGE)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(ANDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GE (TSTshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGE)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(ANDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GE (TSTshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGE)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(ANDshiftLLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (GE (TSTshiftLLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftLLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGE)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftLLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(ANDshiftRLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (GE (TSTshiftRLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftRLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGE)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftRLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(ANDshiftRAreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (GE (TSTshiftRAreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftRAreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGE)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftRAreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(XOR x y)) yes no)
		// cond: l.Uses==1
		// result: (GE (TEQ x y) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXOR {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				b.Reset(BlockARMGE)
				v0 := b.NewValue0(v_0.Pos, OpARMTEQ, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (GE (CMPconst [0] l:(XORconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (GE (TEQconst [c] x) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGE)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(XORshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GE (TEQshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGE)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(XORshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GE (TEQshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGE)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(XORshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GE (TEQshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGE)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(XORshiftLLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (GE (TEQshiftLLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftLLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGE)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftLLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(XORshiftRLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (GE (TEQshiftRLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftRLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGE)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftRLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (GE (CMPconst [0] l:(XORshiftRAreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (GE (TEQshiftRAreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftRAreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGE)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftRAreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
	case BlockARMGT:
		// match: (GT (FlagEQ) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpARMFlagEQ {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (GT (FlagLT_ULT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpARMFlagLT_ULT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (GT (FlagLT_UGT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpARMFlagLT_UGT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (GT (FlagGT_ULT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpARMFlagGT_ULT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (GT (FlagGT_UGT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpARMFlagGT_UGT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (GT (InvertFlags cmp) yes no)
		// result: (LT cmp yes no)
		for b.Controls[0].Op == OpARMInvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockARMLT)
			b.AddControl(cmp)
			return true
		}
		// match: (GT (CMPconst [0] l:(SUB x y)) yes no)
		// cond: l.Uses==1
		// result: (GT (CMP x y) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUB {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGT)
			v0 := b.NewValue0(v_0.Pos, OpARMCMP, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(MULS x y a)) yes no)
		// cond: l.Uses==1
		// result: (GT (CMP a (MUL <x.Type> x y)) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMMULS {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGT)
			v0 := b.NewValue0(v_0.Pos, OpARMCMP, types.TypeFlags)
			v0.AddArg(a)
			v1 := b.NewValue0(v_0.Pos, OpARMMUL, x.Type)
			v1.AddArg(x)
			v1.AddArg(y)
			v0.AddArg(v1)
			b.AddControl(v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(SUBconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (GT (CMPconst [c] x) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGT)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(SUBshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GT (CMPshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGT)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(SUBshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GT (CMPshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGT)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(SUBshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GT (CMPshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGT)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(SUBshiftLLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (GT (CMPshiftLLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftLLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGT)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftLLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(SUBshiftRLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (GT (CMPshiftRLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftRLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGT)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftRLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(SUBshiftRAreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (GT (CMPshiftRAreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftRAreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGT)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftRAreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(ADD x y)) yes no)
		// cond: l.Uses==1
		// result: (GT (CMN x y) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADD {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				b.Reset(BlockARMGT)
				v0 := b.NewValue0(v_0.Pos, OpARMCMN, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (GT (CMPconst [0] l:(ADDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (GT (CMNconst [c] x) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGT)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(ADDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GT (CMNshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGT)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(ADDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GT (CMNshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGT)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(ADDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GT (CMNshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGT)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(ADDshiftLLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (GT (CMNshiftLLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftLLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGT)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftLLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(ADDshiftRLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (GT (CMNshiftRLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftRLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGT)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftRLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(ADDshiftRAreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (GT (CMNshiftRAreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftRAreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGT)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftRAreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(AND x y)) yes no)
		// cond: l.Uses==1
		// result: (GT (TST x y) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMAND {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				b.Reset(BlockARMGT)
				v0 := b.NewValue0(v_0.Pos, OpARMTST, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (GT (CMPconst [0] l:(MULA x y a)) yes no)
		// cond: l.Uses==1
		// result: (GT (CMN a (MUL <x.Type> x y)) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMMULA {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGT)
			v0 := b.NewValue0(v_0.Pos, OpARMCMN, types.TypeFlags)
			v0.AddArg(a)
			v1 := b.NewValue0(v_0.Pos, OpARMMUL, x.Type)
			v1.AddArg(x)
			v1.AddArg(y)
			v0.AddArg(v1)
			b.AddControl(v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(ANDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (GT (TSTconst [c] x) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGT)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(ANDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GT (TSTshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGT)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(ANDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GT (TSTshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGT)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(ANDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GT (TSTshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGT)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(ANDshiftLLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (GT (TSTshiftLLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftLLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGT)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftLLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(ANDshiftRLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (GT (TSTshiftRLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftRLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGT)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftRLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(ANDshiftRAreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (GT (TSTshiftRAreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftRAreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGT)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftRAreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(XOR x y)) yes no)
		// cond: l.Uses==1
		// result: (GT (TEQ x y) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXOR {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				b.Reset(BlockARMGT)
				v0 := b.NewValue0(v_0.Pos, OpARMTEQ, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (GT (CMPconst [0] l:(XORconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (GT (TEQconst [c] x) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGT)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(XORshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GT (TEQshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGT)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(XORshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GT (TEQshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGT)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(XORshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (GT (TEQshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGT)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(XORshiftLLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (GT (TEQshiftLLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftLLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGT)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftLLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(XORshiftRLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (GT (TEQshiftRLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftRLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGT)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftRLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (GT (CMPconst [0] l:(XORshiftRAreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (GT (TEQshiftRAreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftRAreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMGT)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftRAreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
	case BlockIf:
		// match: (If (Equal cc) yes no)
		// result: (EQ cc yes no)
		for b.Controls[0].Op == OpARMEqual {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.Reset(BlockARMEQ)
			b.AddControl(cc)
			return true
		}
		// match: (If (NotEqual cc) yes no)
		// result: (NE cc yes no)
		for b.Controls[0].Op == OpARMNotEqual {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.Reset(BlockARMNE)
			b.AddControl(cc)
			return true
		}
		// match: (If (LessThan cc) yes no)
		// result: (LT cc yes no)
		for b.Controls[0].Op == OpARMLessThan {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.Reset(BlockARMLT)
			b.AddControl(cc)
			return true
		}
		// match: (If (LessThanU cc) yes no)
		// result: (ULT cc yes no)
		for b.Controls[0].Op == OpARMLessThanU {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.Reset(BlockARMULT)
			b.AddControl(cc)
			return true
		}
		// match: (If (LessEqual cc) yes no)
		// result: (LE cc yes no)
		for b.Controls[0].Op == OpARMLessEqual {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.Reset(BlockARMLE)
			b.AddControl(cc)
			return true
		}
		// match: (If (LessEqualU cc) yes no)
		// result: (ULE cc yes no)
		for b.Controls[0].Op == OpARMLessEqualU {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.Reset(BlockARMULE)
			b.AddControl(cc)
			return true
		}
		// match: (If (GreaterThan cc) yes no)
		// result: (GT cc yes no)
		for b.Controls[0].Op == OpARMGreaterThan {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.Reset(BlockARMGT)
			b.AddControl(cc)
			return true
		}
		// match: (If (GreaterThanU cc) yes no)
		// result: (UGT cc yes no)
		for b.Controls[0].Op == OpARMGreaterThanU {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.Reset(BlockARMUGT)
			b.AddControl(cc)
			return true
		}
		// match: (If (GreaterEqual cc) yes no)
		// result: (GE cc yes no)
		for b.Controls[0].Op == OpARMGreaterEqual {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.Reset(BlockARMGE)
			b.AddControl(cc)
			return true
		}
		// match: (If (GreaterEqualU cc) yes no)
		// result: (UGE cc yes no)
		for b.Controls[0].Op == OpARMGreaterEqualU {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.Reset(BlockARMUGE)
			b.AddControl(cc)
			return true
		}
		// match: (If cond yes no)
		// result: (NE (CMPconst [0] cond) yes no)
		for {
			cond := b.Controls[0]
			b.Reset(BlockARMNE)
			v0 := b.NewValue0(cond.Pos, OpARMCMPconst, types.TypeFlags)
			v0.AuxInt = 0
			v0.AddArg(cond)
			b.AddControl(v0)
			return true
		}
	case BlockARMLE:
		// match: (LE (FlagEQ) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpARMFlagEQ {
			b.Reset(BlockFirst)
			return true
		}
		// match: (LE (FlagLT_ULT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpARMFlagLT_ULT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (LE (FlagLT_UGT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpARMFlagLT_UGT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (LE (FlagGT_ULT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpARMFlagGT_ULT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (LE (FlagGT_UGT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpARMFlagGT_UGT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (LE (InvertFlags cmp) yes no)
		// result: (GE cmp yes no)
		for b.Controls[0].Op == OpARMInvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockARMGE)
			b.AddControl(cmp)
			return true
		}
		// match: (LE (CMPconst [0] l:(SUB x y)) yes no)
		// cond: l.Uses==1
		// result: (LE (CMP x y) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUB {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMP, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(MULS x y a)) yes no)
		// cond: l.Uses==1
		// result: (LE (CMP a (MUL <x.Type> x y)) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMMULS {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMP, types.TypeFlags)
			v0.AddArg(a)
			v1 := b.NewValue0(v_0.Pos, OpARMMUL, x.Type)
			v1.AddArg(x)
			v1.AddArg(y)
			v0.AddArg(v1)
			b.AddControl(v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(SUBconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (LE (CMPconst [c] x) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(SUBshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LE (CMPshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(SUBshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LE (CMPshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(SUBshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LE (CMPshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(SUBshiftLLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (LE (CMPshiftLLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftLLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftLLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(SUBshiftRLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (LE (CMPshiftRLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftRLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftRLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(SUBshiftRAreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (LE (CMPshiftRAreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftRAreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftRAreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(ADD x y)) yes no)
		// cond: l.Uses==1
		// result: (LE (CMN x y) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADD {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				b.Reset(BlockARMLE)
				v0 := b.NewValue0(v_0.Pos, OpARMCMN, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (LE (CMPconst [0] l:(MULA x y a)) yes no)
		// cond: l.Uses==1
		// result: (LE (CMN a (MUL <x.Type> x y)) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMMULA {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMN, types.TypeFlags)
			v0.AddArg(a)
			v1 := b.NewValue0(v_0.Pos, OpARMMUL, x.Type)
			v1.AddArg(x)
			v1.AddArg(y)
			v0.AddArg(v1)
			b.AddControl(v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(ADDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (LE (CMNconst [c] x) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(ADDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LE (CMNshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(ADDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LE (CMNshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(ADDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LE (CMNshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(ADDshiftLLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (LE (CMNshiftLLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftLLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftLLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(ADDshiftRLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (LE (CMNshiftRLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftRLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftRLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(ADDshiftRAreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (LE (CMNshiftRAreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftRAreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftRAreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(AND x y)) yes no)
		// cond: l.Uses==1
		// result: (LE (TST x y) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMAND {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				b.Reset(BlockARMLE)
				v0 := b.NewValue0(v_0.Pos, OpARMTST, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (LE (CMPconst [0] l:(ANDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (LE (TSTconst [c] x) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLE)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(ANDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LE (TSTshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLE)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(ANDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LE (TSTshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLE)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(ANDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LE (TSTshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLE)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(ANDshiftLLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (LE (TSTshiftLLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftLLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLE)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftLLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(ANDshiftRLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (LE (TSTshiftRLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftRLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLE)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftRLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(ANDshiftRAreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (LE (TSTshiftRAreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftRAreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLE)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftRAreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(XOR x y)) yes no)
		// cond: l.Uses==1
		// result: (LE (TEQ x y) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXOR {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				b.Reset(BlockARMLE)
				v0 := b.NewValue0(v_0.Pos, OpARMTEQ, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (LE (CMPconst [0] l:(XORconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (LE (TEQconst [c] x) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLE)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(XORshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LE (TEQshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLE)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(XORshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LE (TEQshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLE)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(XORshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LE (TEQshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLE)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(XORshiftLLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (LE (TEQshiftLLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftLLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLE)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftLLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(XORshiftRLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (LE (TEQshiftRLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftRLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLE)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftRLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (LE (CMPconst [0] l:(XORshiftRAreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (LE (TEQshiftRAreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftRAreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLE)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftRAreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
	case BlockARMLT:
		// match: (LT (FlagEQ) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpARMFlagEQ {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (LT (FlagLT_ULT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpARMFlagLT_ULT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (LT (FlagLT_UGT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpARMFlagLT_UGT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (LT (FlagGT_ULT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpARMFlagGT_ULT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (LT (FlagGT_UGT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpARMFlagGT_UGT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (LT (InvertFlags cmp) yes no)
		// result: (GT cmp yes no)
		for b.Controls[0].Op == OpARMInvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockARMGT)
			b.AddControl(cmp)
			return true
		}
		// match: (LT (CMPconst [0] l:(SUB x y)) yes no)
		// cond: l.Uses==1
		// result: (LT (CMP x y) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUB {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLT)
			v0 := b.NewValue0(v_0.Pos, OpARMCMP, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(MULS x y a)) yes no)
		// cond: l.Uses==1
		// result: (LT (CMP a (MUL <x.Type> x y)) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMMULS {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLT)
			v0 := b.NewValue0(v_0.Pos, OpARMCMP, types.TypeFlags)
			v0.AddArg(a)
			v1 := b.NewValue0(v_0.Pos, OpARMMUL, x.Type)
			v1.AddArg(x)
			v1.AddArg(y)
			v0.AddArg(v1)
			b.AddControl(v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(SUBconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (LT (CMPconst [c] x) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLT)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(SUBshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LT (CMPshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLT)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(SUBshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LT (CMPshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLT)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(SUBshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LT (CMPshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLT)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(SUBshiftLLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (LT (CMPshiftLLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftLLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLT)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftLLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(SUBshiftRLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (LT (CMPshiftRLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftRLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLT)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftRLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(SUBshiftRAreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (LT (CMPshiftRAreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftRAreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLT)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftRAreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(ADD x y)) yes no)
		// cond: l.Uses==1
		// result: (LT (CMN x y) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADD {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				b.Reset(BlockARMLT)
				v0 := b.NewValue0(v_0.Pos, OpARMCMN, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (LT (CMPconst [0] l:(MULA x y a)) yes no)
		// cond: l.Uses==1
		// result: (LT (CMN a (MUL <x.Type> x y)) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMMULA {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLT)
			v0 := b.NewValue0(v_0.Pos, OpARMCMN, types.TypeFlags)
			v0.AddArg(a)
			v1 := b.NewValue0(v_0.Pos, OpARMMUL, x.Type)
			v1.AddArg(x)
			v1.AddArg(y)
			v0.AddArg(v1)
			b.AddControl(v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(ADDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (LT (CMNconst [c] x) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLT)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(ADDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LT (CMNshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLT)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(ADDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LT (CMNshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLT)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(ADDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LT (CMNshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLT)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(ADDshiftLLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (LT (CMNshiftLLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftLLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLT)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftLLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(ADDshiftRLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (LT (CMNshiftRLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftRLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLT)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftRLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(ADDshiftRAreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (LT (CMNshiftRAreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftRAreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLT)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftRAreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(AND x y)) yes no)
		// cond: l.Uses==1
		// result: (LT (TST x y) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMAND {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				b.Reset(BlockARMLT)
				v0 := b.NewValue0(v_0.Pos, OpARMTST, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (LT (CMPconst [0] l:(ANDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (LT (TSTconst [c] x) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLT)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(ANDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LT (TSTshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLT)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(ANDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LT (TSTshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLT)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(ANDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LT (TSTshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLT)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(ANDshiftLLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (LT (TSTshiftLLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftLLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLT)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftLLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(ANDshiftRLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (LT (TSTshiftRLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftRLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLT)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftRLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(ANDshiftRAreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (LT (TSTshiftRAreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftRAreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLT)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftRAreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(XOR x y)) yes no)
		// cond: l.Uses==1
		// result: (LT (TEQ x y) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXOR {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				b.Reset(BlockARMLT)
				v0 := b.NewValue0(v_0.Pos, OpARMTEQ, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (LT (CMPconst [0] l:(XORconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (LT (TEQconst [c] x) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLT)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(XORshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LT (TEQshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLT)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(XORshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LT (TEQshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLT)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(XORshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (LT (TEQshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLT)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(XORshiftLLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (LT (TEQshiftLLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftLLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLT)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftLLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(XORshiftRLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (LT (TEQshiftRLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftRLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLT)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftRLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (LT (CMPconst [0] l:(XORshiftRAreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (LT (TEQshiftRAreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftRAreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMLT)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftRAreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
	case BlockARMNE:
		// match: (NE (CMPconst [0] (Equal cc)) yes no)
		// result: (EQ cc yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpARMEqual {
				break
			}
			cc := v_0_0.Args[0]
			b.Reset(BlockARMEQ)
			b.AddControl(cc)
			return true
		}
		// match: (NE (CMPconst [0] (NotEqual cc)) yes no)
		// result: (NE cc yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpARMNotEqual {
				break
			}
			cc := v_0_0.Args[0]
			b.Reset(BlockARMNE)
			b.AddControl(cc)
			return true
		}
		// match: (NE (CMPconst [0] (LessThan cc)) yes no)
		// result: (LT cc yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpARMLessThan {
				break
			}
			cc := v_0_0.Args[0]
			b.Reset(BlockARMLT)
			b.AddControl(cc)
			return true
		}
		// match: (NE (CMPconst [0] (LessThanU cc)) yes no)
		// result: (ULT cc yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpARMLessThanU {
				break
			}
			cc := v_0_0.Args[0]
			b.Reset(BlockARMULT)
			b.AddControl(cc)
			return true
		}
		// match: (NE (CMPconst [0] (LessEqual cc)) yes no)
		// result: (LE cc yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpARMLessEqual {
				break
			}
			cc := v_0_0.Args[0]
			b.Reset(BlockARMLE)
			b.AddControl(cc)
			return true
		}
		// match: (NE (CMPconst [0] (LessEqualU cc)) yes no)
		// result: (ULE cc yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpARMLessEqualU {
				break
			}
			cc := v_0_0.Args[0]
			b.Reset(BlockARMULE)
			b.AddControl(cc)
			return true
		}
		// match: (NE (CMPconst [0] (GreaterThan cc)) yes no)
		// result: (GT cc yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpARMGreaterThan {
				break
			}
			cc := v_0_0.Args[0]
			b.Reset(BlockARMGT)
			b.AddControl(cc)
			return true
		}
		// match: (NE (CMPconst [0] (GreaterThanU cc)) yes no)
		// result: (UGT cc yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpARMGreaterThanU {
				break
			}
			cc := v_0_0.Args[0]
			b.Reset(BlockARMUGT)
			b.AddControl(cc)
			return true
		}
		// match: (NE (CMPconst [0] (GreaterEqual cc)) yes no)
		// result: (GE cc yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpARMGreaterEqual {
				break
			}
			cc := v_0_0.Args[0]
			b.Reset(BlockARMGE)
			b.AddControl(cc)
			return true
		}
		// match: (NE (CMPconst [0] (GreaterEqualU cc)) yes no)
		// result: (UGE cc yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			v_0_0 := v_0.Args[0]
			if v_0_0.Op != OpARMGreaterEqualU {
				break
			}
			cc := v_0_0.Args[0]
			b.Reset(BlockARMUGE)
			b.AddControl(cc)
			return true
		}
		// match: (NE (FlagEQ) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpARMFlagEQ {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (NE (FlagLT_ULT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpARMFlagLT_ULT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (NE (FlagLT_UGT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpARMFlagLT_UGT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (NE (FlagGT_ULT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpARMFlagGT_ULT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (NE (FlagGT_UGT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpARMFlagGT_UGT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (NE (InvertFlags cmp) yes no)
		// result: (NE cmp yes no)
		for b.Controls[0].Op == OpARMInvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockARMNE)
			b.AddControl(cmp)
			return true
		}
		// match: (NE (CMPconst [0] l:(SUB x y)) yes no)
		// cond: l.Uses==1
		// result: (NE (CMP x y) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUB {
				break
			}
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMNE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMP, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(MULS x y a)) yes no)
		// cond: l.Uses==1
		// result: (NE (CMP a (MUL <x.Type> x y)) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMMULS {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMNE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMP, types.TypeFlags)
			v0.AddArg(a)
			v1 := b.NewValue0(v_0.Pos, OpARMMUL, x.Type)
			v1.AddArg(x)
			v1.AddArg(y)
			v0.AddArg(v1)
			b.AddControl(v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(SUBconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (NE (CMPconst [c] x) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMNE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(SUBshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (CMPshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMNE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(SUBshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (CMPshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMNE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(SUBshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (CMPshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMNE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(SUBshiftLLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (NE (CMPshiftLLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftLLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMNE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftLLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(SUBshiftRLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (NE (CMPshiftRLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftRLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMNE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftRLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(SUBshiftRAreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (NE (CMPshiftRAreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMSUBshiftRAreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMNE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMPshiftRAreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(ADD x y)) yes no)
		// cond: l.Uses==1
		// result: (NE (CMN x y) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADD {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				b.Reset(BlockARMNE)
				v0 := b.NewValue0(v_0.Pos, OpARMCMN, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (NE (CMPconst [0] l:(MULA x y a)) yes no)
		// cond: l.Uses==1
		// result: (NE (CMN a (MUL <x.Type> x y)) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMMULA {
				break
			}
			a := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMNE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMN, types.TypeFlags)
			v0.AddArg(a)
			v1 := b.NewValue0(v_0.Pos, OpARMMUL, x.Type)
			v1.AddArg(x)
			v1.AddArg(y)
			v0.AddArg(v1)
			b.AddControl(v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(ADDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (NE (CMNconst [c] x) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMNE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(ADDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (CMNshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMNE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(ADDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (CMNshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMNE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(ADDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (CMNshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMNE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(ADDshiftLLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (NE (CMNshiftLLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftLLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMNE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftLLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(ADDshiftRLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (NE (CMNshiftRLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftRLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMNE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftRLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(ADDshiftRAreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (NE (CMNshiftRAreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMADDshiftRAreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMNE)
			v0 := b.NewValue0(v_0.Pos, OpARMCMNshiftRAreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(AND x y)) yes no)
		// cond: l.Uses==1
		// result: (NE (TST x y) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMAND {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				b.Reset(BlockARMNE)
				v0 := b.NewValue0(v_0.Pos, OpARMTST, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (NE (CMPconst [0] l:(ANDconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (NE (TSTconst [c] x) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMNE)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(ANDshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (TSTshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMNE)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(ANDshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (TSTshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMNE)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(ANDshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (TSTshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMNE)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(ANDshiftLLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (NE (TSTshiftLLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftLLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMNE)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftLLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(ANDshiftRLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (NE (TSTshiftRLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftRLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMNE)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftRLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(ANDshiftRAreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (NE (TSTshiftRAreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMANDshiftRAreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMNE)
			v0 := b.NewValue0(v_0.Pos, OpARMTSTshiftRAreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(XOR x y)) yes no)
		// cond: l.Uses==1
		// result: (NE (TEQ x y) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXOR {
				break
			}
			_ = l.Args[1]
			l_0 := l.Args[0]
			l_1 := l.Args[1]
			for _i0 := 0; _i0 <= 1; _i0, l_0, l_1 = _i0+1, l_1, l_0 {
				x := l_0
				y := l_1
				if !(l.Uses == 1) {
					continue
				}
				b.Reset(BlockARMNE)
				v0 := b.NewValue0(v_0.Pos, OpARMTEQ, types.TypeFlags)
				v0.AddArg(x)
				v0.AddArg(y)
				b.AddControl(v0)
				return true
			}
			break
		}
		// match: (NE (CMPconst [0] l:(XORconst [c] x)) yes no)
		// cond: l.Uses==1
		// result: (NE (TEQconst [c] x) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORconst {
				break
			}
			c := l.AuxInt
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMNE)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQconst, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			b.AddControl(v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(XORshiftLL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (TEQshiftLL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftLL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMNE)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftLL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(XORshiftRL x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (TEQshiftRL x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftRL {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMNE)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftRL, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(XORshiftRA x y [c])) yes no)
		// cond: l.Uses==1
		// result: (NE (TEQshiftRA x y [c]) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftRA {
				break
			}
			c := l.AuxInt
			y := l.Args[1]
			x := l.Args[0]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMNE)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftRA, types.TypeFlags)
			v0.AuxInt = c
			v0.AddArg(x)
			v0.AddArg(y)
			b.AddControl(v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(XORshiftLLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (NE (TEQshiftLLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftLLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMNE)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftLLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(XORshiftRLreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (NE (TEQshiftRLreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftRLreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMNE)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftRLreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
		// match: (NE (CMPconst [0] l:(XORshiftRAreg x y z)) yes no)
		// cond: l.Uses==1
		// result: (NE (TEQshiftRAreg x y z) yes no)
		for b.Controls[0].Op == OpARMCMPconst {
			v_0 := b.Controls[0]
			if v_0.AuxInt != 0 {
				break
			}
			l := v_0.Args[0]
			if l.Op != OpARMXORshiftRAreg {
				break
			}
			z := l.Args[2]
			x := l.Args[0]
			y := l.Args[1]
			if !(l.Uses == 1) {
				break
			}
			b.Reset(BlockARMNE)
			v0 := b.NewValue0(v_0.Pos, OpARMTEQshiftRAreg, types.TypeFlags)
			v0.AddArg(x)
			v0.AddArg(y)
			v0.AddArg(z)
			b.AddControl(v0)
			return true
		}
	case BlockARMUGE:
		// match: (UGE (FlagEQ) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpARMFlagEQ {
			b.Reset(BlockFirst)
			return true
		}
		// match: (UGE (FlagLT_ULT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpARMFlagLT_ULT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (UGE (FlagLT_UGT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpARMFlagLT_UGT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (UGE (FlagGT_ULT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpARMFlagGT_ULT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (UGE (FlagGT_UGT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpARMFlagGT_UGT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (UGE (InvertFlags cmp) yes no)
		// result: (ULE cmp yes no)
		for b.Controls[0].Op == OpARMInvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockARMULE)
			b.AddControl(cmp)
			return true
		}
	case BlockARMUGT:
		// match: (UGT (FlagEQ) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpARMFlagEQ {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (UGT (FlagLT_ULT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpARMFlagLT_ULT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (UGT (FlagLT_UGT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpARMFlagLT_UGT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (UGT (FlagGT_ULT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpARMFlagGT_ULT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (UGT (FlagGT_UGT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpARMFlagGT_UGT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (UGT (InvertFlags cmp) yes no)
		// result: (ULT cmp yes no)
		for b.Controls[0].Op == OpARMInvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockARMULT)
			b.AddControl(cmp)
			return true
		}
	case BlockARMULE:
		// match: (ULE (FlagEQ) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpARMFlagEQ {
			b.Reset(BlockFirst)
			return true
		}
		// match: (ULE (FlagLT_ULT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpARMFlagLT_ULT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (ULE (FlagLT_UGT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpARMFlagLT_UGT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (ULE (FlagGT_ULT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpARMFlagGT_ULT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (ULE (FlagGT_UGT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpARMFlagGT_UGT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (ULE (InvertFlags cmp) yes no)
		// result: (UGE cmp yes no)
		for b.Controls[0].Op == OpARMInvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockARMUGE)
			b.AddControl(cmp)
			return true
		}
	case BlockARMULT:
		// match: (ULT (FlagEQ) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpARMFlagEQ {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (ULT (FlagLT_ULT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpARMFlagLT_ULT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (ULT (FlagLT_UGT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpARMFlagLT_UGT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (ULT (FlagGT_ULT) yes no)
		// result: (First yes no)
		for b.Controls[0].Op == OpARMFlagGT_ULT {
			b.Reset(BlockFirst)
			return true
		}
		// match: (ULT (FlagGT_UGT) yes no)
		// result: (First no yes)
		for b.Controls[0].Op == OpARMFlagGT_UGT {
			b.Reset(BlockFirst)
			b.swapSuccessors()
			return true
		}
		// match: (ULT (InvertFlags cmp) yes no)
		// result: (UGT cmp yes no)
		for b.Controls[0].Op == OpARMInvertFlags {
			v_0 := b.Controls[0]
			cmp := v_0.Args[0]
			b.Reset(BlockARMUGT)
			b.AddControl(cmp)
			return true
		}
	}
	return false
}
