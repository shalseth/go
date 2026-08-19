// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build sparc64 && gc

#include "textflag.h"

#define MEMBAR_FULL	MEMBAR	$15	// #LoadLoad|#StoreLoad|#LoadStore|#StoreStore

TEXT ·RewindAndSetgid(SB),NOSPLIT|NOFRAME,$0-0
	// Rewind stack pointer so anything that happens on the stack
	// will clobber the test pattern created by the caller. On SPARC
	// that includes the register window the kernel spills on entry to
	// the signal handler, which is the write this test is watching for.
	ADD	$(1024*8), BSP

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
	SUB	$(1024*8), BSP
	RET
