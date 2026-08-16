// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// SPARC system calls trap with ta 0x6d; the number goes in %g1 (R1),
// arguments in %o0-%o5 (R8-R13). An error is signaled by the carry
// flag, with the (positive) errno in %o0.

// func rawVforkSyscall(trap, a1, a2, a3 uintptr) (r1 uintptr, err Errno)
TEXT ·rawVforkSyscall(SB),NOSPLIT|NOFRAME,$0-48
	MOVD	trap+0(FP), R1
	MOVD	a1+8(FP), R8
	MOVD	a2+16(FP), R9
	MOVD	a3+24(FP), R10
	TA	$0x6d
	BCSW	err
	// Fork-like syscalls return the pid in %o0 in both processes;
	// %o1 is 0 in the parent and 1 in the child. Convert to the
	// Linux convention of returning 0 in the child.
	CMP	R9, ZR
	BED	parent
	MOVD	ZR, R8
parent:
	MOVD	R8, r1+32(FP)
	MOVD	ZR, err+40(FP)
	RET
err:
	MOVD	$-1, R2
	MOVD	R2, r1+32(FP)
	MOVD	R8, err+40(FP)
	RET

// func rawSyscallNoError(trap, a1, a2, a3 uintptr) (r1, r2 uintptr)
TEXT ·rawSyscallNoError(SB),NOSPLIT|NOFRAME,$0-48
	MOVD	trap+0(FP), R1
	MOVD	a1+8(FP), R8
	MOVD	a2+16(FP), R9
	MOVD	a3+24(FP), R10
	TA	$0x6d
	MOVD	R8, r1+32(FP)
	MOVD	R9, r2+40(FP)
	RET
