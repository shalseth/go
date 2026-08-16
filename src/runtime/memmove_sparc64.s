// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// func memmove(to, from unsafe.Pointer, n uintptr)
//
// Copies forward when to < from and backward otherwise, so overlapping
// regions are handled. 8-byte accesses only when both pointers and the
// length are 8-aligned, since SPARC traps on unaligned access. The 2016
// tree's memmove read registers it never initialized and was not
// salvageable.
TEXT runtime·memmove(SB),NOSPLIT|NOFRAME,$0-24
	MOVD	to+0(FP), R8
	MOVD	from+8(FP), R9
	MOVD	n+16(FP), R10
	CMP	ZR, R10
	BED	done
	// Same pointer: nothing to do.
	CMP	R9, R8
	BED	done
	BCSD	forward		// to < from: copy up

	// Backward copy, from the last byte down.
	OR	R8, R9, R11
	OR	R11, R10, R11
	AND	$7, R11, R11
	CMP	ZR, R11
	BNED	back1
	ADD	R8, R10, R8	// to end
	ADD	R9, R10, R9	// from end
back8loop:
	SUB	$8, R9
	SUB	$8, R8
	MOVD	(R9), R11
	MOVD	R11, (R8)
	CMP	ZR, R10
	SUB	$8, R10
	CMP	ZR, R10
	BNED	back8loop
	RET
back1:
	ADD	R8, R10, R8
	ADD	R9, R10, R9
back1loop:
	SUB	$1, R9
	SUB	$1, R8
	MOVUB	(R9), R11
	MOVB	R11, (R8)
	SUB	$1, R10
	CMP	ZR, R10
	BNED	back1loop
	RET

forward:
	OR	R8, R9, R11
	OR	R11, R10, R11
	AND	$7, R11, R11
	CMP	ZR, R11
	BNED	fwd1
	ADD	R9, R10, R10	// R10 = from end
fwd8loop:
	MOVD	(R9), R11
	MOVD	R11, (R8)
	ADD	$8, R9
	ADD	$8, R8
	CMP	R10, R9
	BNED	fwd8loop
	RET
fwd1:
	ADD	R9, R10, R10
fwd1loop:
	MOVUB	(R9), R11
	MOVB	R11, (R8)
	ADD	$1, R9
	ADD	$1, R8
	CMP	R10, R9
	BNED	fwd1loop
done:
	RET
