// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_asm.h"
#include "textflag.h"

// func memequal(a, b unsafe.Pointer, size uintptr) bool
//
// A byte loop with an 8-byte fast path when both pointers are aligned;
// SPARC traps on unaligned access.
TEXT runtime·memequal(SB),NOSPLIT|NOFRAME,$0-25
	MOVD	a+0(FP), R8
	MOVD	b+8(FP), R9
	MOVD	size+16(FP), R10
	CMP	R9, R8
	BED	eq
	CMP	ZR, R10
	BED	eq
	// Word loop only when both pointers are 8-aligned and at least a
	// word remains; the tail falls through to the byte loop.
	OR	R8, R9, R11
	AND	$7, R11, R11
	CMP	ZR, R11
	BNED	byteentry
wordloop:
	CMP	$8, R10
	BCSD	byteentry
	MOVD	(R8), R11
	MOVD	(R9), R12
	SUBCC	R12, R11, ZR
	BNED	noteq
	ADD	$8, R8
	ADD	$8, R9
	SUB	$8, R10
	JMP	wordloop
byteentry:
	CMP	ZR, R10
	BED	eq
byteloop:
	MOVUB	(R8), R11
	MOVUB	(R9), R12
	SUBCC	R12, R11, ZR
	BNED	noteq
	ADD	$1, R8
	ADD	$1, R9
	SUB	$1, R10
	CMP	ZR, R10
	BNED	byteloop
eq:
	MOVD	$1, R8
	MOVB	R8, ret+24(FP)
	RET
noteq:
	MOVB	ZR, ret+24(FP)
	RET

// func memequal_varlen(a, b unsafe.Pointer) bool
TEXT runtime·memequal_varlen(SB),NOSPLIT,$32-17
	MOVD	a+0(FP), R8
	MOVD	b+8(FP), R9
	CMP	R9, R8
	BED	eq
	// The compiler stores the size at offset 8 in the closure.
	MOVD	8(CTXT), R10
	MOVD	R8, (176+0)(BSP)
	MOVD	R9, (176+8)(BSP)
	MOVD	R10, (176+16)(BSP)
	CALL	runtime·memequal(SB)
	MOVUB	(176+24)(BSP), R8
	MOVB	R8, ret+16(FP)
	RET
eq:
	MOVD	$1, R8
	MOVB	R8, ret+16(FP)
	RET
