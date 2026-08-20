// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// At process entry the kernel places the argument block directly after
// the 16-word register window save area: argc at %sp+BIAS+128 and argv
// right after. (glibc reads these at +176, but only because its _start
// first drops the stack pointer by another 48 bytes.) Verified with a
// debugger at the entry breakpoint on the T4.
TEXT _rt0_sparc64_linux(SB),NOSPLIT|NOFRAME,$0
	MOVD	128(BSP), R9	// argc
	MOVD	$136(BSP), R8	// argv
	JMP	runtime·rt0_go(SB)

// main is the entry point when the program is linked externally: the C
// startup code calls it following the C ABI, with argc in %o0 and argv
// in %o1. rt0_go wants them the other way round, argc in R9 and argv
// in R8, so the two are swapped on the way through.
TEXT main(SB),NOSPLIT|NOFRAME,$0
	MOVD	R8, TMP		// argc
	MOVD	R9, R8		// argv
	MOVD	TMP, R9		// argc
	JMP	runtime·rt0_go(SB)
