// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package interp

import (
	"exp/ssa"
	"syscall"
)

func ext۰syscall۰Close(fn *ssa.Function, args []value) value {
	panic("syscall.Close not yet implemented")
}
func ext۰syscall۰Fstat(fn *ssa.Function, args []value) value {
	panic("syscall.Fstat not yet implemented")
}
func ext۰syscall۰Getdents(fn *ssa.Function, args []value) value {
	panic("syscall.Getdents not yet implemented")
}
func ext۰syscall۰Kill(fn *ssa.Function, args []value) value {
	panic("syscall.Kill not yet implemented")
}
func ext۰syscall۰Lstat(fn *ssa.Function, args []value) value {
	panic("syscall.Lstat not yet implemented")
}
func ext۰syscall۰Open(fn *ssa.Function, args []value) value {
	panic("syscall.Open not yet implemented")
}
func ext۰syscall۰ParseDirent(fn *ssa.Function, args []value) value {
	panic("syscall.ParseDirent not yet implemented")
}
func ext۰syscall۰Read(fn *ssa.Function, args []value) value {
	panic("syscall.Read not yet implemented")
}
func ext۰syscall۰Stat(fn *ssa.Function, args []value) value {
	panic("syscall.Stat not yet implemented")
}

func ext۰syscall۰Write(fn *ssa.Function, args []value) value {
	p := args[1].([]value)
	b := make([]byte, 0, len(p))
	for i := range p {
		b = append(b, p[i].(byte))
	}
	n, err := syscall.Write(args[0].(int), b)
	return tuple{n, wrapError(err)}
}
