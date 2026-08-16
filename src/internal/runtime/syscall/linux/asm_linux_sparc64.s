// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// func Syscall6(num, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2, errno uintptr)
//
// The linux/sparc64 convention: the syscall number goes in %g1 (R1),
// arguments in %o0..%o5 (R8..R13), and the kernel is entered with
// "ta 0x6d". The result comes back in %o0, and unlike most ports the
// error indication is not a register value but the carry bit of the
// 32-bit condition codes: carry set means %o0 holds an errno.
TEXT ·Syscall6(SB),NOSPLIT,$0-80
	MOVD	num+0(FP), R1
	MOVD	a1+8(FP), R8
	MOVD	a2+16(FP), R9
	MOVD	a3+24(FP), R10
	MOVD	a4+32(FP), R11
	MOVD	a5+40(FP), R12
	MOVD	a6+48(FP), R13
	TA	$0x6d
	BCSW	err
	MOVD	R8, r1+56(FP)
	MOVD	R9, r2+64(FP)
	MOVD	ZR, errno+72(FP)
	RET
err:
	MOVD	$-1, R2
	MOVD	R2, r1+56(FP)
	MOVD	ZR, r2+64(FP)
	MOVD	R8, errno+72(FP)
	RET
