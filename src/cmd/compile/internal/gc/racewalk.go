// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"cmd/compile/internal/base"
	"cmd/compile/internal/ir"
	"cmd/compile/internal/types"
	"cmd/internal/src"
	"cmd/internal/sys"
)

func instrument(fn *ir.Func) {
	if fn.Pragma&ir.Norace != 0 || (fn.Sym().Linksym() != nil && fn.Sym().Linksym().ABIWrapper()) {
		return
	}

	if !base.Flag.Race || !base.Compiling(base.NoRacePkgs) {
		fn.SetInstrumentBody(true)
	}

	if base.Flag.Race {
		lno := base.Pos
		base.Pos = src.NoXPos

		if thearch.LinkArch.Arch.Family != sys.AMD64 {
			fn.Enter.Prepend(mkcall("racefuncenterfp", nil, nil))
			fn.Exit.Append(mkcall("racefuncexit", nil, nil))
		} else {

			// nodpc is the PC of the caller as extracted by
			// getcallerpc. We use -widthptr(FP) for x86.
			// This only works for amd64. This will not
			// work on arm or others that might support
			// race in the future.
			nodpc := ir.RegFP.CloneName()
			nodpc.SetType(types.Types[types.TUINTPTR])
			nodpc.SetFrameOffset(int64(-types.PtrSize))
			fn.Dcl = append(fn.Dcl, nodpc)
			fn.Enter.Prepend(mkcall("racefuncenter", nil, nil, nodpc))
			fn.Exit.Append(mkcall("racefuncexit", nil, nil))
		}
		base.Pos = lno
	}
}
