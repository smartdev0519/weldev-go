// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file encapsulates some of the odd characteristics of the RISCV64
// instruction set, to minimize its interaction with the core of the
// assembler.

package arch

import (
	"cmd/internal/obj"
	"cmd/internal/obj/riscv"
)

// IsRISCV64AMO reports whether the op (as defined by a riscv.A*
// constant) is one of the AMO instructions that requires special
// handling.
func IsRISCV64AMO(op obj.As) bool {
	switch op {
	case riscv.ASCW, riscv.ASCD:
		return true
	}
	return false
}
