// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// SPARC system calls trap with ta 0x6d; the number goes in %g1 (R1),
// arguments in %o0-%o5 (R8-R13). An error is signaled by the carry
// flag, with the (positive) errno in %o0.

#define SYS_clone	217

// func rawVforkSyscall(trap, a1, a2, a3 uintptr) (r1 uintptr, err Errno)
TEXT ·rawVforkSyscall(SB),NOSPLIT|NOFRAME,$0-48
	MOVD	trap+0(FP), R1
	MOVD	a1+8(FP), R8
	MOVD	a2+16(FP), R9
	MOVD	a3+24(FP), R10
	TA	$0x6d
	BCSW	err
	// sys_clone has a sparc wrapper that keeps the SunOS convention:
	// the pid comes back in %o0 to *both* processes and %o1 tells them
	// apart, zero in the parent and nonzero in the child. Convert that
	// to the Linux convention of returning 0 in the child.
	//
	// clone3 has no such wrapper - it is generic kernel code - so it
	// already follows the Linux convention, and %o1 is left holding
	// the size argument the caller passed in, which is never zero.
	// Applying the sys_clone rule there tells the *parent* it is the
	// child, and the parent goes on to execve itself.
	//
	// Re-read the syscall number from the frame rather than keeping it
	// in a register: this runs with CLONE_VM, so the child shares this
	// stack, but both processes read the same still-live argument slot.
	MOVD	trap+0(FP), R2
	MOVD	$SYS_clone, R3
	CMP	R2, R3
	BNED	done
	CMP	R9, ZR
	BED	done
	MOVD	ZR, R8
done:
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
