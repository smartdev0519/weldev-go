// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// makeFuncStub is jumped to by the code generated by MakeFunc.
// See the comment on the declaration of makeFuncStub in makefunc.go
// for more details.
// No argsize here, gc generates argsize info at call site.
TEXT ·makeFuncStub(SB),(NOSPLIT|WRAPPER),$8
	MOVW	R7, 4(R13)
	MOVW	$argframe+0(FP), R1
	MOVW	R1, 8(R13)
	BL	·callReflect(SB)
	RET

// methodValueCall is the code half of the function returned by makeMethodValue.
// See the comment on the declaration of methodValueCall in makefunc.go
// for more details.
// No argsize here, gc generates argsize info at call site.
TEXT ·methodValueCall(SB),(NOSPLIT|WRAPPER),$8
	MOVW	R7, 4(R13)
	MOVW	$argframe+0(FP), R1
	MOVW	R1, 8(R13)
	BL	·callMethod(SB)
	RET

// Stubs to give reflect package access to runtime services
// TODO: should probably be done another way.
TEXT ·makemap(SB),NOSPLIT,$-4-0
	B	runtime·reflect_makemap(SB)
TEXT ·mapaccess(SB),NOSPLIT,$-4-0
	B	runtime·reflect_mapaccess(SB)
TEXT ·mapassign(SB),NOSPLIT,$-4-0
	B	runtime·reflect_mapassign(SB)
TEXT ·mapdelete(SB),NOSPLIT,$-4-0
	B	runtime·reflect_mapdelete(SB)
TEXT ·mapiterinit(SB),NOSPLIT,$-4-0
	B	runtime·reflect_mapiterinit(SB)
TEXT ·mapiterkey(SB),NOSPLIT,$-4-0
	B	runtime·reflect_mapiterkey(SB)
TEXT ·mapiternext(SB),NOSPLIT,$-4-0
	B	runtime·reflect_mapiternext(SB)
TEXT ·maplen(SB),NOSPLIT,$-4-0
	B	runtime·reflect_maplen(SB)
TEXT ·ismapkey(SB),NOSPLIT,$-4-0
	B	runtime·reflect_ismapkey(SB)
TEXT ·ifaceE2I(SB),NOSPLIT,$0-0
	B	runtime·reflect_ifaceE2I(SB)
TEXT ·unsafe_New(SB),NOSPLIT,$0-0
	B	runtime·newobject(SB)
TEXT ·unsafe_NewArray(SB),NOSPLIT,$0-0
	B	runtime·newarray(SB)
