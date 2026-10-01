// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build sparc64 && gc

#include "textflag.h"

#define MEMBAR_FULL	MEMBAR	$15	// #LoadLoad|#StoreLoad|#LoadStore|#StoreStore

TEXT ·rewindAndSetgid(SB),NOSPLIT|NOFRAME,$0-8
	// The kernel may spill at SP even when the handler uses the signal stack.
	// Keep that spill above the pattern and inside the caller's allocation.
	MOVD	RSP, R5
	MOVD	spill+0(FP), R3
	ADD	$15, R3
	AND	$-16, R3
	MOVD	R3, BSP

	// Ask signaller to setgid
	MOVD	$·Baton(SB), R3
	MOVW	$1, R4
	MEMBAR_FULL
	MOVW	R4, (R3)
	MEMBAR_FULL

	// Wait for setgid completion
loop:
	MEMBAR_FULL
	MOVW	(R3), R4
	CMP	ZR, R4
	BNED	loop
	MEMBAR_FULL

	// Restore stack
	MOVD	R5, RSP
	RET
