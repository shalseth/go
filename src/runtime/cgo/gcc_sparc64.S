// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

.file "gcc_sparc64.S"

/*
 * void crosscall1(void (*fn)(void), void (*setg_gcc)(void *g), void *g)
 *
 * Calling into the gc toolchain, where every register is caller-save.
 * On SPARC the register window does most of that work: save gives this
 * frame its own local and in registers for the Go code to trash, and
 * restore brings the caller's back.
 *
 * Two registers need explicit care. Window in-registers are physically
 * the caller's out-registers, so %i6 and %i7 - the caller's stack
 * pointer and this frame's return address - would be lost when the Go
 * port overwrites them as its own frame anchors. Both go to memory for
 * the duration. fn goes with them, since nothing may be assumed about
 * what survives a call into Go.
 *
 * The frame is the ABI minimum of 176 bytes, a 128-byte window save
 * area and 48 bytes of argument slots, plus 32 for those locals. %sp is
 * biased by 2047.
 */
#define BIAS 2047
#define LOCALS (BIAS + 176)

.globl crosscall1
.type crosscall1, #function
crosscall1:
	save	%sp, -(176 + 32), %sp

	stx	%i7, [%sp + LOCALS + 0]		/* our return address */
	stx	%i6, [%sp + LOCALS + 8]		/* the caller's stack pointer */
	stx	%i0, [%sp + LOCALS + 16]	/* fn */

	mov	%i2, %o0			/* g */
	call	%i1				/* setg_gcc(g) */
	 nop

	ldx	[%sp + LOCALS + 16], %g1
	call	%g1				/* fn() */
	 nop

	ldx	[%sp + LOCALS + 0], %i7
	ldx	[%sp + LOCALS + 8], %i6
	ret
	 restore

.size crosscall1, .-crosscall1

#ifdef __ELF__
.section .note.GNU-stack,"",%progbits
#endif
