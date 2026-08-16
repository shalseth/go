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
