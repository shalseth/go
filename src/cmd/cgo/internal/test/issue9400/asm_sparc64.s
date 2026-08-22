// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build sparc64 && gc

#include "textflag.h"

#define STACK_BIAS	2047

#define MEMBAR_FULL	MEMBAR	$15	// #LoadLoad|#StoreLoad|#LoadStore|#StoreStore

// The rewind has to clear the caller's pattern by more than the pattern's
// own length. A SPARC frame reserves 176 bytes at %sp+STACK_BIAS for the
// register window, and the hardware and the kernel write there on their own
// account - a window spill lands exactly there, whichever stack the signal
// handler itself runs on. Rewinding by just the pattern's length leaves that
// reserved area sitting inside the pattern, so the test reads a legitimate
// 128-byte window image as a failure. Step past the bias and the reserved
// area as well, and what remains below the new stack pointer is still the
// pattern - which is what the test is actually watching.
#define REWIND	(1024*8 + STACK_BIAS + 1 + 176)

TEXT ·RewindAndSetgid(SB),NOSPLIT|NOFRAME,$0-0
	// Rewind stack pointer so anything that happens on the stack
	// will clobber the test pattern created by the caller.
	//
	// Move it through a register: the assembler tracks a constant added
	// to the stack pointer as a frame adjustment, and one this large
	// overflows that.
	MOVD	BSP, R3
	ADD	$REWIND, R3
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
	MOVD	BSP, R3
	SUB	$REWIND, R3
	MOVD	R3, BSP
	RET
