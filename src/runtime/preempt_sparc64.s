// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_asm.h"
#include "textflag.h"

// asyncPreempt saves the interrupted register state and calls
// asyncPreempt2, which parks the goroutine. It is entered by having the
// signal handler rewrite the PC (pushCall), so every register the
// compiler can allocate must be preserved: the interrupted code has no
// idea this ran.
//
// pushCall pushed a MinFrameSize area (with the interrupted %o7 at its
// base) and set LR = resumePC-8, so this behaves exactly like a framed
// function: the hand-written prologue below mirrors the compiled one,
// asyncPreempt2's prologue publishes our return address for the
// unwinder, and the final jump through OLR+8 resumes the interrupted
// code exactly. TMP is clobbered on exit — instructions keeping the
// assembler temporary live are marked non-preemptible.
//
// Frame: the 176-byte SPARC minimum, the 20 allocatable integer
// registers (R1-R5, R8-R13, R16-R20, R24, R25, R29), the reserved
// temporaries TMP2/RT1/RT2 (live across instruction sequences that are
// not individually marked unsafe, like the Zero/Move loops), and the
// 15 allocatable float registers (Y1-Y15 = D2..D30).
TEXT ·asyncPreempt(SB),NOSPLIT|NOFRAME,$0-0
	// Framed prologue by hand.
	//
	// Order matters more here than in a compiled prologue. The kernel
	// spills the live %i6/%i7 to [%sp+bias+112/120] on every trap, and
	// the slots at the entry %sp are the injected frame's copied
	// window - the exit path below reads the interrupted OLR back out
	// of [entry sp+120]. LR at this point holds pushCall's synthetic
	// value (resumePC-8), so OLR must not be overwritten until %sp has
	// moved to this frame, whose own +112/+120 slots are scratch.
	// With the old order (MOVD LR, OLR before the SUB) a signal
	// landing in that one-instruction window let the kernel spill the
	// synthetic LR over the real return address; the goroutine then
	// resumed into the middle of itself and executed the wrong basic
	// blocks with the other path's register meanings. Compiled
	// prologues have the same window but are healed by their birth
	// stores; this frame's record is the copied window, so the window
	// must simply never exist.
	MOVD	RFP, (40)(BSP)
	MOVD	OLR, (120)(BSP)
	SUB	$512, BSP
	ADD	$512, RSP, RFP
	MOVD	LR, OLR

	MOVD	R1, (176+0)(BSP)
	MOVD	R2, (176+8)(BSP)
	MOVD	R3, (176+16)(BSP)
	MOVD	R4, (176+24)(BSP)
	MOVD	R5, (176+32)(BSP)
	MOVD	R8, (176+40)(BSP)
	MOVD	R9, (176+48)(BSP)
	MOVD	R10, (176+56)(BSP)
	MOVD	R11, (176+64)(BSP)
	MOVD	R12, (176+72)(BSP)
	MOVD	R13, (176+80)(BSP)
	MOVD	R16, (176+88)(BSP)
	MOVD	R17, (176+96)(BSP)
	MOVD	R18, (176+104)(BSP)
	MOVD	R19, (176+112)(BSP)
	MOVD	R20, (176+120)(BSP)
	MOVD	R24, (176+136)(BSP)
	MOVD	R25, (176+144)(BSP)
	MOVD	R29, (176+152)(BSP)
	MOVD	TMP2, (176+288)(BSP)
	MOVD	RT1, (176+296)(BSP)
	MOVD	RT2, (176+304)(BSP)
	// The interrupted code may be preempted between a compare and its
	// branch: the integer and float condition codes are live state.
	// Nothing above sets flags (plain SUB/ADD/MOVD), so CCR is still
	// the interrupted value here. R1 is already saved, so it can be
	// scratch.
	RD	CCR, R1
	MOVD	R1, (176+312)(BSP)
	STXFSR	(176+320)(BSP)
	FMOVD	D2, (176+160)(BSP)
	FMOVD	D4, (176+168)(BSP)
	FMOVD	D6, (176+176)(BSP)
	FMOVD	D8, (176+184)(BSP)
	FMOVD	D10, (176+192)(BSP)
	FMOVD	D12, (176+200)(BSP)
	FMOVD	D14, (176+208)(BSP)
	FMOVD	D16, (176+216)(BSP)
	FMOVD	D18, (176+224)(BSP)
	FMOVD	D20, (176+232)(BSP)
	FMOVD	D22, (176+240)(BSP)
	FMOVD	D24, (176+248)(BSP)
	FMOVD	D26, (176+256)(BSP)
	FMOVD	D28, (176+264)(BSP)
	FMOVD	D30, (176+272)(BSP)
	CALL	·asyncPreempt2(SB)
	FMOVD	(176+272)(BSP), D30
	FMOVD	(176+264)(BSP), D28
	FMOVD	(176+256)(BSP), D26
	FMOVD	(176+248)(BSP), D24
	FMOVD	(176+240)(BSP), D22
	FMOVD	(176+232)(BSP), D20
	FMOVD	(176+224)(BSP), D18
	FMOVD	(176+216)(BSP), D16
	FMOVD	(176+208)(BSP), D14
	FMOVD	(176+200)(BSP), D12
	FMOVD	(176+192)(BSP), D10
	FMOVD	(176+184)(BSP), D8
	FMOVD	(176+176)(BSP), D6
	FMOVD	(176+168)(BSP), D4
	FMOVD	(176+160)(BSP), D2
	MOVD	(176+152)(BSP), R29
	MOVD	(176+144)(BSP), R25
	MOVD	(176+136)(BSP), R24
	MOVD	(176+120)(BSP), R20
	MOVD	(176+112)(BSP), R19
	MOVD	(176+104)(BSP), R18
	MOVD	(176+96)(BSP), R17
	MOVD	(176+88)(BSP), R16
	MOVD	(176+80)(BSP), R13
	MOVD	(176+72)(BSP), R12
	MOVD	(176+64)(BSP), R11
	MOVD	(176+56)(BSP), R10
	MOVD	(176+48)(BSP), R9
	MOVD	(176+40)(BSP), R8
	MOVD	(176+32)(BSP), R5
	MOVD	(176+24)(BSP), R4
	MOVD	(176+16)(BSP), R3
	MOVD	(176+8)(BSP), R2
	// Restore the condition codes (no instruction below sets flags),
	// using R1 as scratch before its own final reload.
	LDXFSR	(176+320)(BSP)
	MOVD	(176+312)(BSP), R1
	WR	R1, CCR
	MOVD	(176+0)(BSP), R1
	MOVD	(176+288)(BSP), TMP2
	MOVD	(176+296)(BSP), RT1
	MOVD	(176+304)(BSP), RT2

	// Framed epilogue by hand, plus popping the pushCall area and
	// restoring the interrupted %o7 from its base. The jump goes
	// through TMP because LR must carry the interrupted value.
	//
	// The anchors are read RFP-relative BEFORE RSP moves: the kernel
	// spills %i6 to [sp+bias+112] on a context switch, so RSP must
	// never point at a slot pair that disagrees with the live
	// %i6/%i7 (see the epilogue in cmd/internal/obj/sparc64).
	ADD	$8, OLR, TMP		// resumePC
	MOVD	(128+2047)(RFP), LR	// interrupted LR (spilled by pushCall at +128)
	MOVD	(120+2047)(RFP), OLR	// interrupted OLR
	MOVD	(40+2047)(RFP), RFP	// interrupted RFP (RFP dies last)
	ADD	$(512+176), RSP		// pop our frame and the pushCall area
	JMPL	TMP, ZR
