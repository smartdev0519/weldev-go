// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file encapsulates some of the odd characteristics of the ARM
// instruction set, to minimize its interaction with the core of the
// assembler.

package arch

import (
	"strings"

	"cmd/internal/obj"
	"cmd/internal/obj/arm"
)

var armLS = map[string]uint8{
	"U":  arm.C_UBIT,
	"S":  arm.C_SBIT,
	"W":  arm.C_WBIT,
	"P":  arm.C_PBIT,
	"PW": arm.C_WBIT | arm.C_PBIT,
	"WP": arm.C_WBIT | arm.C_PBIT,
}

var armSCOND = map[string]uint8{
	"EQ":  arm.C_SCOND_EQ,
	"NE":  arm.C_SCOND_NE,
	"CS":  arm.C_SCOND_HS,
	"HS":  arm.C_SCOND_HS,
	"CC":  arm.C_SCOND_LO,
	"LO":  arm.C_SCOND_LO,
	"MI":  arm.C_SCOND_MI,
	"PL":  arm.C_SCOND_PL,
	"VS":  arm.C_SCOND_VS,
	"VC":  arm.C_SCOND_VC,
	"HI":  arm.C_SCOND_HI,
	"LS":  arm.C_SCOND_LS,
	"GE":  arm.C_SCOND_GE,
	"LT":  arm.C_SCOND_LT,
	"GT":  arm.C_SCOND_GT,
	"LE":  arm.C_SCOND_LE,
	"AL":  arm.C_SCOND_NONE,
	"U":   arm.C_UBIT,
	"S":   arm.C_SBIT,
	"W":   arm.C_WBIT,
	"P":   arm.C_PBIT,
	"PW":  arm.C_WBIT | arm.C_PBIT,
	"WP":  arm.C_WBIT | arm.C_PBIT,
	"F":   arm.C_FBIT,
	"IBW": arm.C_WBIT | arm.C_PBIT | arm.C_UBIT,
	"IAW": arm.C_WBIT | arm.C_UBIT,
	"DBW": arm.C_WBIT | arm.C_PBIT,
	"DAW": arm.C_WBIT,
	"IB":  arm.C_PBIT | arm.C_UBIT,
	"IA":  arm.C_UBIT,
	"DB":  arm.C_PBIT,
	"DA":  0,
}

var armJump = map[string]bool{
	"B":    true,
	"BL":   true,
	"BEQ":  true,
	"BNE":  true,
	"BCS":  true,
	"BHS":  true,
	"BCC":  true,
	"BLO":  true,
	"BMI":  true,
	"BPL":  true,
	"BVS":  true,
	"BVC":  true,
	"BHI":  true,
	"BLS":  true,
	"BGE":  true,
	"BLT":  true,
	"BGT":  true,
	"BLE":  true,
	"CALL": true,
}

func jumpArm(word string) bool {
	return armJump[word]
}

// IsARMCMP reports whether the op (as defined by an arm.A* constant) is
// one of the comparison instructions that require special handling.
func IsARMCMP(op int) bool {
	switch op {
	case arm.ACMN, arm.ACMP, arm.ATEQ, arm.ATST:
		return true
	}
	return false
}

// IsARMSTREX reports whether the op (as defined by an arm.A* constant) is
// one of the STREX-like instructions that require special handling.
func IsARMSTREX(op int) bool {
	switch op {
	case arm.ASTREX, arm.ASTREXD, arm.ASWPW, arm.ASWPBU:
		return true
	}
	return false
}

// IsARMMRC reports whether the op (as defined by an arm.A* constant) is
// MRC or MCR
func IsARMMRC(op int) bool {
	switch op {
	case arm.AMRC /*, arm.AMCR*/ :
		return true
	}
	return false
}

// ARMMRCOffset implements the peculiar encoding of the MRC and MCR instructions.
func ARMMRCOffset(op int, cond string, x0, x1, x2, x3, x4, x5 int64) (offset int64, ok bool) {
	// TODO only MRC is defined.
	op1 := int64(0)
	if op == arm.AMRC {
		op1 = 1
	}
	bits, ok := ParseARMCondition(cond)
	if !ok {
		return
	}
	offset = (0xe << 24) | // opcode
		(op1 << 20) | // MCR/MRC
		((int64(bits) ^ arm.C_SCOND_XOR) << 28) | // scond
		((x0 & 15) << 8) | //coprocessor number
		((x1 & 7) << 21) | // coprocessor operation
		((x2 & 15) << 12) | // ARM register
		((x3 & 15) << 16) | // Crn
		((x4 & 15) << 0) | // Crm
		((x5 & 7) << 5) | // coprocessor information
		(1 << 4) /* must be set */
	return offset, true
}

// IsARMMULA reports whether the op (as defined by an arm.A* constant) is
// MULA, MULAWT or MULAWB, the 4-operand instructions.
func IsARMMULA(op int) bool {
	switch op {
	case arm.AMULA, arm.AMULAWB, arm.AMULAWT:
		return true
	}
	return false
}

var bcode = []int{
	arm.ABEQ,
	arm.ABNE,
	arm.ABCS,
	arm.ABCC,
	arm.ABMI,
	arm.ABPL,
	arm.ABVS,
	arm.ABVC,
	arm.ABHI,
	arm.ABLS,
	arm.ABGE,
	arm.ABLT,
	arm.ABGT,
	arm.ABLE,
	arm.AB,
	obj.ANOP,
}

// ARMConditionCodes handles the special condition code situation for the ARM.
// It returns a boolean to indicate success; failure means cond was unrecognized.
func ARMConditionCodes(prog *obj.Prog, cond string) bool {
	if cond == "" {
		return true
	}
	bits, ok := ParseARMCondition(cond)
	if !ok {
		return false
	}
	/* hack to make B.NE etc. work: turn it into the corresponding conditional */
	if prog.As == arm.AB {
		prog.As = int16(bcode[(bits^arm.C_SCOND_XOR)&0xf])
		bits = (bits &^ 0xf) | arm.C_SCOND_NONE
	}
	prog.Scond = bits
	return true
}

// ParseARMCondition parses the conditions attached to an ARM instruction.
// The input is a single string consisting of period-separated condition
// codes, such as ".P.W". An initial period is ignored.
func ParseARMCondition(cond string) (uint8, bool) {
	if strings.HasPrefix(cond, ".") {
		cond = cond[1:]
	}
	if cond == "" {
		return arm.C_SCOND_NONE, true
	}
	names := strings.Split(cond, ".")
	bits := uint8(0)
	for _, name := range names {
		if b, present := armLS[name]; present {
			bits |= b
			continue
		}
		if b, present := armSCOND[name]; present {
			bits = (bits &^ arm.C_SCOND) | b
			continue
		}
		return 0, false
	}
	return bits, true
}

func armRegisterNumber(name string, n int16) (int16, bool) {
	if n < 0 || 15 < n {
		return 0, false
	}
	switch name {
	case "R":
		return arm.REG_R0 + n, true
	case "F":
		return arm.REG_F0 + n, true
	}
	return 0, false
}
