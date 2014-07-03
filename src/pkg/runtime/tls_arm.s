// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "zasm_GOOS_GOARCH.h"
#include "funcdata.h"
#include "../../cmd/ld/textflag.h"

// We have to resort to TLS variable to save g(R10).
// One reason is that external code might trigger
// SIGSEGV, and our runtime.sigtramp don't even know we
// are in external code, and will continue to use R10,
// this might as well result in another SIGSEGV.
// Note: both functions will clobber R0 and R11 and
// can be called from 5c ABI code.

// On android, runtime.tlsg is a normal variable.
// TLS offset is computed in x_cgo_inittls.

// save_g saves the g register into pthread-provided
// thread-local memory, so that we can call externally compiled
// ARM code that will overwrite those registers.
// NOTE: runtime.gogo assumes that R1 is preserved by this function.
//       runtime.mcall assumes this function only clobbers R0 and R11.
TEXT runtime·save_g(SB),NOSPLIT,$0
	MRC		15, 0, R0, C13, C0, 3 // fetch TLS base pointer
	// $runtime.tlsg(SB) is a special linker symbol.
	// It is the offset from the TLS base pointer to our
	// thread-local storage for g.
#ifdef GOOS_android
	MOVW	runtime·tlsg(SB), R11
#else
	MOVW	$runtime·tlsg(SB), R11
#endif
	ADD	R11, R0
	MOVW	g, 0(R0)
	RET

// load_g loads the g register from pthread-provided
// thread-local memory, for use after calling externally compiled
// ARM code that overwrote those registers.
TEXT runtime·load_g(SB),NOSPLIT,$0
	MRC		15, 0, R0, C13, C0, 3 // fetch TLS base pointer
	// $runtime.tlsg(SB) is a special linker symbol.
	// It is the offset from the TLS base pointer to our
	// thread-local storage for g.
#ifdef GOOS_android
	MOVW	runtime·tlsg(SB), R11
#else
	MOVW	$runtime·tlsg(SB), R11
#endif
	ADD	R11, R0
	MOVW	0(R0), g
	RET
