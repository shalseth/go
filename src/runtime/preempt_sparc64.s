// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_asm.h"
#include "textflag.h"

// asyncPreempt saves the caller's register state and calls
// asyncPreempt2, which parks the goroutine. It is entered by having the
// signal handler rewrite the PC, so every live register must be
// preserved: the interrupted code has no idea this ran.
//
// The frame is the 176-byte SPARC minimum plus room for %g1..%g5,
// %o0..%o5 and %l0..%l7, which are the registers the compiler
// allocates. %g0 is hardwired zero, %o6 and %i6 are the stack and frame
// pointers, %o7 is saved by the prologue, and %g7 is the thread
// pointer, so none of those need saving here.
TEXT ·asyncPreempt(SB),NOSPLIT|NOFRAME,$0-0
	SUB	$304, BSP
	MOVD	R1, (176+0)(BSP)
	MOVD	R2, (176+8)(BSP)
	MOVD	R3, (176+16)(BSP)
	MOVD	R4, (176+24)(BSP)
	MOVD	R5, (176+32)(BSP)
	MOVD	R8, (176+40)(BSP)
	MOVD	R9, (176+48)(BSP)
	MOVD	R10, (176+56)(BSP)
	MOVD	R11, (176+64)(BSP)
	MOVD	R12, (176+72)(BSP)
	MOVD	R13, (176+80)(BSP)
	MOVD	R16, (176+88)(BSP)
	MOVD	R17, (176+96)(BSP)
	MOVD	R18, (176+104)(BSP)
	MOVD	R19, (176+112)(BSP)
	MOVD	R20, (176+120)(BSP)
	MOVD	R21, (176+128)(BSP)
	CALL	·asyncPreempt2(SB)
	MOVD	(176+128)(BSP), R21
	MOVD	(176+120)(BSP), R20
	MOVD	(176+112)(BSP), R19
	MOVD	(176+104)(BSP), R18
	MOVD	(176+96)(BSP), R17
	MOVD	(176+88)(BSP), R16
	MOVD	(176+80)(BSP), R13
	MOVD	(176+72)(BSP), R12
	MOVD	(176+64)(BSP), R11
	MOVD	(176+56)(BSP), R10
	MOVD	(176+48)(BSP), R9
	MOVD	(176+40)(BSP), R8
	MOVD	(176+32)(BSP), R5
	MOVD	(176+24)(BSP), R4
	MOVD	(176+16)(BSP), R3
	MOVD	(176+8)(BSP), R2
	MOVD	(176+0)(BSP), R1
	ADD	$304, BSP
	// The signal handler pushed the interrupted PC as our return
	// address, so returning resumes the preempted code.
	RET
