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

// _rt0_sparc64_linux_lib is the entry point for -buildmode=c-archive and
// c-shared: the linker runs it as a global constructor, with the C ABI
// and argc and argv in %o0 and %o1. It opens a register window of its
// own, so every register C expects back survives the Go code below, and
// flushes the caller's windows to their stacks before Go runs in this
// one. libInit starts the runtime on a thread of its own and returns.
TEXT _rt0_sparc64_linux_lib(SB),NOSPLIT|NOFRAME,$0
	SAVE	$-192, RSP, RSP
	FLUSHW
	MOVD	R30, (176+0)(BSP)	// %i6 and %i7, the way back to C
	MOVD	R31, (176+8)(BSP)
	MOVD	ZR, g
	MOVD	I0, _rt0_sparc64_lib_argc<>(SB)
	MOVD	I1, _rt0_sparc64_lib_argv<>(SB)
	CALL	runtime·libInit(SB)
	MOVD	(176+0)(BSP), R30
	MOVD	(176+8)(BSP), R31
	RESTORE	ZR, ZR, ZR
	RET

// rt0_lib_go starts the runtime on the thread libInit created. It is
// entered from C and does not return, like main above.
TEXT runtime·rt0_lib_go<ABIInternal>(SB),NOSPLIT|NOFRAME,$0
	MOVD	_rt0_sparc64_lib_argc<>(SB), R9
	MOVD	_rt0_sparc64_lib_argv<>(SB), R8
	JMP	runtime·rt0_go(SB)

DATA _rt0_sparc64_lib_argc<>(SB)/8, $0
GLOBL _rt0_sparc64_lib_argc<>(SB), NOPTR, $8
DATA _rt0_sparc64_lib_argv<>(SB)/8, $0
GLOBL _rt0_sparc64_lib_argv<>(SB), NOPTR, $8
