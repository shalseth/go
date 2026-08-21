// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "funcdata.h"
#include "textflag.h"

// Point the _crosscall2_ptr C function pointer at crosscall2, through a
// local trampoline so that the address of a dynamically exported
// function is never taken.
TEXT ·set_crosscall2(SB),NOSPLIT,$0-0
	MOVD	_crosscall2_ptr(SB), R1
	MOVD	$crosscall2_trampoline<>(SB), R2
	MOVD	R2, (R1)
	RET

TEXT crosscall2_trampoline<>(SB),NOSPLIT,$0-0
	JMP	crosscall2(SB)

// Reached from crosscall2, which cmd/cgo's C code calls and which opens
// a register window first. See gcc_sparc64.S for why the window matters:
// Go walks the stack pointer onto a goroutine stack from here, and a C
// window left non-current across that loses its registers when the
// goroutine stack moves.
//
// func crosscall2_flat(fn, a unsafe.Pointer, n int32, ctxt uintptr)
// Calls cgocallback with three of those arguments; n is unused.
//
// Entered with the C calling convention, so the arguments arrive in the
// out-registers, and every register the C ABI expects to survive has to
// be preserved. A declared frame gets the prologue to spill RFP, OLR
// and LR, which leaves the local and in registers to save by hand: the
// first 24 bytes of the frame are the outgoing arguments for
// cgocallback, so they go above those.
#define SAVED (176 + 24)

TEXT crosscall2_flat(SB),NOSPLIT,$144-0
	NO_LOCAL_POINTERS

	MOVD	R16, (SAVED+0)(BSP)	// %l0
	MOVD	R17, (SAVED+8)(BSP)
	MOVD	R18, (SAVED+16)(BSP)
	MOVD	R19, (SAVED+24)(BSP)
	MOVD	R20, (SAVED+32)(BSP)
	MOVD	R21, (SAVED+40)(BSP)
	MOVD	g, (SAVED+48)(BSP)	// %l6
	MOVD	R23, (SAVED+56)(BSP)	// %l7
	MOVD	R24, (SAVED+64)(BSP)	// %i0
	MOVD	R25, (SAVED+72)(BSP)
	MOVD	R26, (SAVED+80)(BSP)	// %i2, which load_g clobbers
	MOVD	R27, (SAVED+88)(BSP)
	MOVD	R28, (SAVED+96)(BSP)
	MOVD	R29, (SAVED+104)(BSP)

	// fn, a and ctxt arrived in %o0, %o1 and %o3.
	MOVD	R8, (176+0)(BSP)
	MOVD	R9, (176+8)(BSP)
	MOVD	R11, (176+16)(BSP)

	CALL	runtime·load_g(SB)
	CALL	runtime·cgocallback(SB)

	MOVD	(SAVED+0)(BSP), R16
	MOVD	(SAVED+8)(BSP), R17
	MOVD	(SAVED+16)(BSP), R18
	MOVD	(SAVED+24)(BSP), R19
	MOVD	(SAVED+32)(BSP), R20
	MOVD	(SAVED+40)(BSP), R21
	MOVD	(SAVED+48)(BSP), g
	MOVD	(SAVED+56)(BSP), R23
	MOVD	(SAVED+64)(BSP), R24
	MOVD	(SAVED+72)(BSP), R25
	MOVD	(SAVED+80)(BSP), R26
	MOVD	(SAVED+88)(BSP), R27
	MOVD	(SAVED+96)(BSP), R28
	MOVD	(SAVED+104)(BSP), R29

	RET
