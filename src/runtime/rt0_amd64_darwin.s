// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.


TEXT	_rt0_amd64_darwin(SB),7,$-8

// copy arguments forward on an even stack


	MOVQ	0(SP), AX		// argc
	LEAQ	8(SP), BX		// argv
	ANDQ	$~7, SP
	SUBQ	$32, SP
	MOVQ	AX, 16(SP)
	MOVQ	BX, 24(SP)

// allocate the per-user block

	LEAQ	peruser<>(SB), R15	// dedicated u. register
	MOVQ	SP, AX
	SUBQ	$4096, AX
	MOVQ	AX, 0(R15)

	CALL	check(SB)

// process the arguments

	MOVL	16(SP), AX
	MOVL	AX, 0(SP)
	MOVQ	24(SP), AX
	MOVQ	AX, 8(SP)
	CALL	args(SB)

	CALL	main·main(SB)

	MOVQ	$0, AX
	MOVQ	AX, 0(SP)		// exit status
	CALL	sys·exit(SB)

	CALL	notok(SB)

	ADDQ	$32, SP
	RET

TEXT	_morestack(SB), 7, $0
	MOVQ	SP, AX
	SUBQ	$1024, AX
	MOVQ	AX, 0(R15)
	RET

TEXT	FLUSH(SB),7,$-8
	RET

TEXT	sys·exit(SB),1,$-8
	MOVL	8(SP), DI		// arg 1 exit status
	MOVL	$(0x2000000+1), AX
	SYSCALL
	JCC	2(PC)
	CALL	notok(SB)
	RET

TEXT	sys·write(SB),1,$-8
	MOVL	8(SP), DI		// arg 1 fid
	MOVQ	16(SP), SI		// arg 2 buf
	MOVL	24(SP), DX		// arg 3 count
	MOVL	$(0x2000000+4), AX	// syscall entry
	SYSCALL
	JCC	2(PC)
	CALL	notok(SB)
	RET

TEXT	open(SB),1,$-8
	MOVQ	8(SP), DI
	MOVL	16(SP), SI
	MOVQ	$0, R10
	MOVL	$(0x2000000+5), AX	// syscall entry
	SYSCALL
	RET

TEXT	close(SB),1,$-8
	MOVL	8(SP), DI
	MOVL	$(0x2000000+6), AX	// syscall entry
	SYSCALL
	RET

TEXT	fstat(SB),1,$-8
	MOVL	8(SP), DI
	MOVQ	16(SP), SI
	MOVL	$(0x2000000+339), AX	// syscall entry; really fstat64
	SYSCALL
	RET

TEXT	read(SB),1,$-8
	MOVL	8(SP), DI
	MOVQ	16(SP), SI
	MOVL	24(SP), DX
	MOVL	$(0x2000000+3), AX	// syscall entry
	SYSCALL
	RET

TEXT	sys·sigaction(SB),1,$-8
	MOVL	8(SP), DI		// arg 1 sig
	MOVQ	16(SP), SI		// arg 2 act
	MOVQ	24(SP), DX		// arg 3 oact
	MOVQ	24(SP), CX		// arg 3 oact
	MOVQ	24(SP), R10		// arg 3 oact
	MOVL	$(0x2000000+46), AX	// syscall entry
	SYSCALL
	JCC	2(PC)
	CALL	notok(SB)
	RET

TEXT sigtramp(SB),1,$24
	MOVL	DX,0(SP)
	MOVQ	CX,8(SP)
	MOVQ	R8,16(SP)
	CALL	sighandler(SB)
	RET

TEXT	sys·breakpoint(SB),1,$-8
	BYTE	$0xcc
	RET

TEXT	sys·mmap(SB),1,$-8
	MOVQ	8(SP), DI		// arg 1 addr
	MOVL	16(SP), SI		// arg 2 len
	MOVL	20(SP), DX		// arg 3 prot
	MOVL	24(SP), R10		// arg 4 flags
	MOVL	28(SP), R8		// arg 5 fid
	MOVL	32(SP), R9		// arg 6 offset
	MOVL	$(0x2000000+197), AX	// syscall entry
	SYSCALL
	JCC	2(PC)
	CALL	notok(SB)
	RET

TEXT	notok(SB),1,$-8
	MOVL	$0xf1, BP
	MOVQ	BP, (BP)
	RET

TEXT	sys·memclr(SB),1,$-8
	MOVQ	8(SP), DI		// arg 1 addr
	MOVL	16(SP), CX		// arg 2 count
	ADDL	$7, CX
	SHRL	$3, CX
	MOVQ	$0, AX
	CLD
	REP
	STOSQ
	RET

TEXT	sys·getcallerpc+0(SB),1,$0
	MOVQ	x+0(FP),AX
	MOVQ	-8(AX),AX
	RET

GLOBL	peruser<>(SB),$64
