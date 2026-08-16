// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// func memclrNoHeapPointers(ptr unsafe.Pointer, n uintptr)
//
// Straightforward and correct rather than fast: an 8-byte loop when
// both the pointer and the length are 8-aligned (SPARC traps on
// unaligned access), a byte loop otherwise. The 2016 tree's memclr read
// registers it never initialized and was not salvageable.
TEXT runtime·memclrNoHeapPointers(SB),NOSPLIT|NOFRAME,$0-16
	MOVD	ptr+0(FP), R8
	MOVD	n+8(FP), R9
	CMP	ZR, R9
	BED	done
	ADD	R8, R9, R9	// R9 = end
	OR	R8, R9, R10
	AND	$7, R10, R10
	CMP	ZR, R10
	BNED	byteloop
wordloop:
	MOVD	ZR, (R8)
	ADD	$8, R8
	CMP	R9, R8
	BNED	wordloop
	RET
byteloop:
	MOVB	ZR, (R8)
	ADD	$1, R8
	CMP	R9, R8
	BNED	byteloop
done:
	RET
