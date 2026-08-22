// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

.file "gcc_sparc64.S"

/*
 * void crosscall1(void (*fn)(void), void (*setg_gcc)(void *g), void *g)
 *
 * Calling into the gc toolchain, where every register is caller-save.
 *
 * The obvious way to do this on SPARC is a save/restore pair, letting
 * the register window preserve the caller's locals. That is not used
 * here: this port's Go code runs flat, one window for the whole
 * program, and it treats %i6 and %i7 as frame anchors it may overwrite.
 * Mixing the two conventions at this boundary is hard to reason about,
 * so the frame is flat here too and every register the C ABI expects to
 * survive a call is saved by hand, as the ports without register
 * windows do.
 *
 * %sp is biased by 2047; the frame keeps the ABI's 176-byte minimum -
 * a 128-byte window save area the kernel may spill into, and 48 bytes
 * of argument slots - and adds room for the saved registers.
 */
#define BIAS 2047
#define SAVED (BIAS + 176)

.globl crosscall1
.type crosscall1, #function
crosscall1:
	add	%sp, -(176 + 176 + 192), %sp

	/* Same reason as crosscall2 below: Go is about to use %i6, which is
	   the stack pointer of the window beneath this one. */
	flushw

	stx	%l0, [%sp + SAVED + 0]
	stx	%l1, [%sp + SAVED + 8]
	stx	%l2, [%sp + SAVED + 16]
	stx	%l3, [%sp + SAVED + 24]
	stx	%l4, [%sp + SAVED + 32]
	stx	%l5, [%sp + SAVED + 40]
	stx	%l6, [%sp + SAVED + 48]
	stx	%l7, [%sp + SAVED + 56]
	stx	%i0, [%sp + SAVED + 64]
	stx	%i1, [%sp + SAVED + 72]
	stx	%i2, [%sp + SAVED + 80]
	stx	%i3, [%sp + SAVED + 88]
	stx	%i4, [%sp + SAVED + 96]
	stx	%i5, [%sp + SAVED + 104]
	stx	%i6, [%sp + SAVED + 112]
	stx	%i7, [%sp + SAVED + 120]

	/*
	 * The first Go function's epilogue restores its caller's frame
	 * anchor from this frame's AnchorFP slot - the one belonging to
	 * whichever register RFP is. Publish it there as well.
	 */
	stx	%i6, [%sp + BIAS + 40]
	stx	%o7, [%sp + SAVED + 128]

	/* fn and setg_gcc outlive calls into Go, so they go to memory. */
	stx	%o0, [%sp + SAVED + 136]
	stx	%o1, [%sp + SAVED + 144]

	mov	%o2, %o0			/* g */
	ldx	[%sp + SAVED + 144], %g1
	call	%g1				/* setg_gcc(g) */
	 nop

	ldx	[%sp + SAVED + 136], %g1
	/* Slot 40 above is %l5's spill slot: keep the live register equal
	   to the published anchor so a trap here cannot clobber it. %l5 was
	   saved above and is restored below. */
	mov	%i6, %l5

	/* Park %i6 on the scratch region carved at the top of this frame:
	   the kernel dereferences the live %i6 (clone reads a frame through
	   it, window spills write through it), and Go is about to run for
	   an arbitrarily long time with this window's registers. The
	   scratch's own %i6 image points back at itself. %i6 was saved
	   above and is restored below. */
	add	%sp, (176 + 176), %i6
	stx	%i6, [%i6 + BIAS + 112]

	call	%g1				/* fn() */
	 nop

	ldx	[%sp + SAVED + 0], %l0
	ldx	[%sp + SAVED + 8], %l1
	ldx	[%sp + SAVED + 16], %l2
	ldx	[%sp + SAVED + 24], %l3
	ldx	[%sp + SAVED + 32], %l4
	ldx	[%sp + SAVED + 40], %l5
	ldx	[%sp + SAVED + 48], %l6
	ldx	[%sp + SAVED + 56], %l7
	ldx	[%sp + SAVED + 64], %i0
	ldx	[%sp + SAVED + 72], %i1
	ldx	[%sp + SAVED + 80], %i2
	ldx	[%sp + SAVED + 88], %i3
	ldx	[%sp + SAVED + 96], %i4
	ldx	[%sp + SAVED + 104], %i5
	ldx	[%sp + SAVED + 112], %i6
	ldx	[%sp + SAVED + 120], %i7
	ldx	[%sp + SAVED + 128], %o7

	add	%sp, (176 + 176 + 192), %sp
	retl
	 nop

.size crosscall1, .-crosscall1

#ifdef __ELF__
/*
 * void crosscall2(void (*fn)(void), void *a, int n, void *ctxt)
 *
 * Opens a register window before handing control to the Go side, which
 * is entered flat as crosscall2_flat.
 *
 * This is not bookkeeping for the C ABI's benefit - crosscall2_flat
 * saves by hand everything C expects to survive. It is about which
 * window's %sp the kernel uses when it spills the register file. Go
 * code runs flat and walks the stack pointer off the C stack and onto a
 * goroutine stack, and a goroutine stack moves when a callback grows
 * it. A window that is not current while that happens gets refilled
 * from an address that no longer holds it, and the C frame underneath
 * silently loses registers it was holding across the call.
 *
 * Opening a window here keeps every C window's %sp on the C stack,
 * where nothing moves, and leaves the window Go walks the stack in
 * current for as long as it does so.
 */
.globl crosscall2
.type crosscall2, #function
crosscall2:
	save	%sp, -(176 + 192), %sp

	/*
	 * Push the C windows out to the frames they belong to before Go
	 * code touches %i6.
	 *
	 * %i6 in one window is the same physical register as %o6 - the
	 * stack pointer - of the window beneath it. Go runs flat and uses
	 * %i6 as its frame pointer, so every Go prologue entered from here
	 * repoints the stack pointer of the C window underneath at the Go
	 * stack. Spill that window afterwards and it lands on a goroutine
	 * stack, on top of live Go frames.
	 *
	 * After this the C windows are in memory, at their own frames, so
	 * nothing needs spilling while Go runs; the fills on the way out
	 * read those frames, because the epilogues restore %i6 first.
	 */
	flushw

	/* Same as crosscall1: advertise the anchor where Go looks, and
	   keep live %l5 equal to its published spill slot. */
	stx	%i6, [%sp + BIAS + 40]
	stx	%i7, [%sp + BIAS + 120]
	mov	%i6, %l5

	/* %i6 must NOT be parked here, unlike crosscall1: this window came
	   from a save, so %i6 is physically the caller's live %o6 - its
	   stack pointer - and a restore does not undo a write to it. It
	   already names a valid C frame, which is all the parking is for. */

	mov	%i0, %o0
	mov	%i1, %o1
	mov	%i2, %o2
	mov	%i3, %o3
	call	crosscall2_flat
	 nop

	return	%i7 + 8
	 nop
.size crosscall2, .-crosscall2

.section .note.GNU-stack,"",%progbits
#endif
