// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build race

#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"
#include "asm_sparc64.h"

// The following thunks allow calling the gcc-compiled race runtime directly
// from Go code without going all the way through cgo.
// First, it's much faster (up to 50% speedup for real Go programs).
// Second, it eliminates race-related special cases from cgocall and scheduler.
// Third, in long-term it will allow to remove cyclic runtime/race dependency on cmd/go.

// A brief recap of the sparc64 calling convention.
// C takes its arguments in %o0..%o5 and returns in %o0. %g1..%g5 are
// volatile; the windowed registers are preserved by the hardware across a
// call made from a window of one's own, which is what racecall<> below
// arranges. SP carries the SPARC V9 bias of 2047 and stays 16-byte aligned.
//
// This port has no register ABI, so the thunks take their arguments off the
// stack rather than out of registers, and a raw LR is the address of the
// CALL rather than a return address - see "Notes on the ABI" in
// README.sparc64.md.

// Registers in this file, on the way to racecall<>:
//	O0..O3	the C function's arguments
//	O4	the C function to call
//	L4	racecallatomic<>'s only value that outlives a racecall<>

// A frame's anchors, at these offsets from its own sp+bias. A prologue
// writes its caller's pair there too, which is why RFP reaches them.
// See AnchorFP and AnchorLR in cmd/internal/obj/sparc64.
#define ANCHOR_FP 40
#define ANCHOR_LR 136

// func runtime·raceread(addr uintptr)
// Called from instrumented code.
TEXT	runtime·raceread(SB), NOSPLIT|NOFRAME, $0-8
	MOVD	addr+0(FP), O1
	ADD	$8, LR, O2
	// void __tsan_read(ThreadState *thr, void *addr, void *pc);
	MOVD	$__tsan_read(SB), O4
	JMP	racecalladdr<>(SB)

// func runtime·RaceRead(addr uintptr)
TEXT	runtime·RaceRead(SB), NOSPLIT|NOFRAME, $0-8
	// This needs to be a tail call, because raceread reads caller pc.
	JMP	runtime·raceread(SB)

// func runtime·racereadpc(void *addr, void *callpc, void *pc)
TEXT	runtime·racereadpc(SB), NOSPLIT|NOFRAME, $0-24
	MOVD	addr+0(FP), O1
	MOVD	callpc+8(FP), O2
	MOVD	pc+16(FP), O3
	// void __tsan_read_pc(ThreadState *thr, void *addr, void *callpc, void *pc);
	MOVD	$__tsan_read_pc(SB), O4
	JMP	racecalladdr<>(SB)

// func runtime·racewrite(addr uintptr)
// Called from instrumented code.
TEXT	runtime·racewrite(SB), NOSPLIT|NOFRAME, $0-8
	MOVD	addr+0(FP), O1
	ADD	$8, LR, O2
	// void __tsan_write(ThreadState *thr, void *addr, void *pc);
	MOVD	$__tsan_write(SB), O4
	JMP	racecalladdr<>(SB)

// func runtime·RaceWrite(addr uintptr)
TEXT	runtime·RaceWrite(SB), NOSPLIT|NOFRAME, $0-8
	// This needs to be a tail call, because racewrite reads caller pc.
	JMP	runtime·racewrite(SB)

// func runtime·racewritepc(void *addr, void *callpc, void *pc)
TEXT	runtime·racewritepc(SB), NOSPLIT|NOFRAME, $0-24
	MOVD	addr+0(FP), O1
	MOVD	callpc+8(FP), O2
	MOVD	pc+16(FP), O3
	// void __tsan_write_pc(ThreadState *thr, void *addr, void *callpc, void *pc);
	MOVD	$__tsan_write_pc(SB), O4
	JMP	racecalladdr<>(SB)

// func runtime·racereadrange(addr, size uintptr)
// Called from instrumented code.
TEXT	runtime·racereadrange(SB), NOSPLIT|NOFRAME, $0-16
	MOVD	addr+0(FP), O1
	MOVD	size+8(FP), O2
	ADD	$8, LR, O3
	// void __tsan_read_range(ThreadState *thr, void *addr, uintptr size, void *pc);
	MOVD	$__tsan_read_range(SB), O4
	JMP	racecalladdr<>(SB)

// func runtime·RaceReadRange(addr, size uintptr)
TEXT	runtime·RaceReadRange(SB), NOSPLIT|NOFRAME, $0-16
	// This needs to be a tail call, because racereadrange reads caller pc.
	JMP	runtime·racereadrange(SB)

// func runtime·racereadrangepc1(void *addr, uintptr sz, void *pc)
TEXT	runtime·racereadrangepc1(SB), NOSPLIT|NOFRAME, $0-24
	MOVD	addr+0(FP), O1
	MOVD	size+8(FP), O2
	MOVD	pc+16(FP), O3
	ADD	$8, O3, O3	// pc is function start, tsan wants a return address. It
			// symbolizes one by stepping back a whole delay-slot pair
			// on SPARC, so this is 8 where arm64 uses 4.
	// void __tsan_read_range(ThreadState *thr, void *addr, uintptr size, void *pc);
	MOVD	$__tsan_read_range(SB), O4
	JMP	racecalladdr<>(SB)

// func runtime·racewriterange(addr, size uintptr)
// Called from instrumented code.
TEXT	runtime·racewriterange(SB), NOSPLIT|NOFRAME, $0-16
	MOVD	addr+0(FP), O1
	MOVD	size+8(FP), O2
	ADD	$8, LR, O3
	// void __tsan_write_range(ThreadState *thr, void *addr, uintptr size, void *pc);
	MOVD	$__tsan_write_range(SB), O4
	JMP	racecalladdr<>(SB)

// func runtime·RaceWriteRange(addr, size uintptr)
TEXT	runtime·RaceWriteRange(SB), NOSPLIT|NOFRAME, $0-16
	// This needs to be a tail call, because racewriterange reads caller pc.
	JMP	runtime·racewriterange(SB)

// func runtime·racewriterangepc1(void *addr, uintptr sz, void *pc)
TEXT	runtime·racewriterangepc1(SB), NOSPLIT|NOFRAME, $0-24
	MOVD	addr+0(FP), O1
	MOVD	size+8(FP), O2
	MOVD	pc+16(FP), O3
	ADD	$8, O3, O3	// pc is function start, tsan wants a return address. It
			// symbolizes one by stepping back a whole delay-slot pair
			// on SPARC, so this is 8 where arm64 uses 4.
	// void __tsan_write_range(ThreadState *thr, void *addr, uintptr size, void *pc);
	MOVD	$__tsan_write_range(SB), O4
	JMP	racecalladdr<>(SB)

// If addr (O1) is out of range, do nothing.
// Otherwise, setup goroutine context and invoke racecall. Other arguments already set.
TEXT	racecalladdr<>(SB), NOSPLIT|NOFRAME, $0-0
	MOVD	g_racectx(g), O0
	// Check that addr is within [arenastart, arenaend) or within [racedatastart, racedataend).
	// Userspace addresses are sign-extended from 52 bits, so these
	// comparisons are unsigned.
	MOVD	runtime·racearenastart(SB), L0
	CMP	L0, O1
	BCSD	data
	MOVD	runtime·racearenaend(SB), L0
	CMP	L0, O1
	BCSD	call
data:
	MOVD	runtime·racedatastart(SB), L0
	CMP	L0, O1
	BCSD	ret
	MOVD	runtime·racedataend(SB), L0
	CMP	L0, O1
	BCCD	ret
call:
	JMP	racecall<>(SB)
ret:
	RET

// func runtime·racefuncenter(pc uintptr)
// Called from instrumented code.
TEXT	runtime·racefuncenter(SB), NOSPLIT|NOFRAME, $0-8
	MOVD	callpc+0(FP), O1
	JMP	racefuncenter<>(SB)

// Common code for racefuncenter
// O1 = caller's return address
TEXT	racefuncenter<>(SB), NOSPLIT|NOFRAME, $0-0
	MOVD	g_racectx(g), O0	// goroutine racectx
	// void __tsan_func_enter(ThreadState *thr, void *pc);
	MOVD	$__tsan_func_enter(SB), O4
	JMP	racecall<>(SB)

// func runtime·racefuncexit()
// Called from instrumented code.
TEXT	runtime·racefuncexit(SB), NOSPLIT|NOFRAME, $0-0
	MOVD	g_racectx(g), O0	// race context
	// void __tsan_func_exit(ThreadState *thr);
	MOVD	$__tsan_func_exit(SB), O4
	JMP	racecall<>(SB)

// Atomic operations for sync/atomic package.
// The tsan function goes in O4; racecallatomic<> reaches the argument list
// and the caller pc through this frame's caller anchors, so these must be
// framed - a tail call would leave it nothing to read.

// Load
TEXT	sync∕atomic·LoadInt32(SB), NOSPLIT, $0-12
	GO_ARGS
	MOVD	$__tsan_go_atomic32_load(SB), O4
	CALL	racecallatomic<>(SB)
	RET

TEXT	sync∕atomic·LoadInt64(SB), NOSPLIT, $0-16
	GO_ARGS
	MOVD	$__tsan_go_atomic64_load(SB), O4
	CALL	racecallatomic<>(SB)
	RET

TEXT	sync∕atomic·LoadUint32(SB), NOSPLIT, $0-12
	GO_ARGS
	JMP	sync∕atomic·LoadInt32(SB)

TEXT	sync∕atomic·LoadUint64(SB), NOSPLIT, $0-16
	GO_ARGS
	JMP	sync∕atomic·LoadInt64(SB)

TEXT	sync∕atomic·LoadUintptr(SB), NOSPLIT, $0-16
	GO_ARGS
	JMP	sync∕atomic·LoadInt64(SB)

TEXT	sync∕atomic·LoadPointer(SB), NOSPLIT, $0-16
	GO_ARGS
	JMP	sync∕atomic·LoadInt64(SB)

// Store
TEXT	sync∕atomic·StoreInt32(SB), NOSPLIT, $0-12
	GO_ARGS
	MOVD	$__tsan_go_atomic32_store(SB), O4
	CALL	racecallatomic<>(SB)
	RET

TEXT	sync∕atomic·StoreInt64(SB), NOSPLIT, $0-16
	GO_ARGS
	MOVD	$__tsan_go_atomic64_store(SB), O4
	CALL	racecallatomic<>(SB)
	RET

TEXT	sync∕atomic·StoreUint32(SB), NOSPLIT, $0-12
	GO_ARGS
	JMP	sync∕atomic·StoreInt32(SB)

TEXT	sync∕atomic·StoreUint64(SB), NOSPLIT, $0-16
	GO_ARGS
	JMP	sync∕atomic·StoreInt64(SB)

TEXT	sync∕atomic·StoreUintptr(SB), NOSPLIT, $0-16
	GO_ARGS
	JMP	sync∕atomic·StoreInt64(SB)

// Swap
TEXT	sync∕atomic·SwapInt32(SB), NOSPLIT, $0-20
	GO_ARGS
	MOVD	$__tsan_go_atomic32_exchange(SB), O4
	CALL	racecallatomic<>(SB)
	RET

TEXT	sync∕atomic·SwapInt64(SB), NOSPLIT, $0-24
	GO_ARGS
	MOVD	$__tsan_go_atomic64_exchange(SB), O4
	CALL	racecallatomic<>(SB)
	RET

TEXT	sync∕atomic·SwapUint32(SB), NOSPLIT, $0-20
	GO_ARGS
	JMP	sync∕atomic·SwapInt32(SB)

TEXT	sync∕atomic·SwapUint64(SB), NOSPLIT, $0-24
	GO_ARGS
	JMP	sync∕atomic·SwapInt64(SB)

TEXT	sync∕atomic·SwapUintptr(SB), NOSPLIT, $0-24
	GO_ARGS
	JMP	sync∕atomic·SwapInt64(SB)

// Add
TEXT	sync∕atomic·AddInt32(SB), NOSPLIT, $0-20
	GO_ARGS
	MOVD	$__tsan_go_atomic32_fetch_add(SB), O4
	CALL	racecallatomic<>(SB)
	MOVW	add+8(FP), O0	// convert fetch_add to add_fetch
	MOVW	ret+16(FP), O1
	ADD	O0, O1, O0
	MOVW	O0, ret+16(FP)
	RET

TEXT	sync∕atomic·AddInt64(SB), NOSPLIT, $0-24
	GO_ARGS
	MOVD	$__tsan_go_atomic64_fetch_add(SB), O4
	CALL	racecallatomic<>(SB)
	MOVD	add+8(FP), O0	// convert fetch_add to add_fetch
	MOVD	ret+16(FP), O1
	ADD	O0, O1, O0
	MOVD	O0, ret+16(FP)
	RET

TEXT	sync∕atomic·AddUint32(SB), NOSPLIT, $0-20
	GO_ARGS
	JMP	sync∕atomic·AddInt32(SB)

TEXT	sync∕atomic·AddUint64(SB), NOSPLIT, $0-24
	GO_ARGS
	JMP	sync∕atomic·AddInt64(SB)

TEXT	sync∕atomic·AddUintptr(SB), NOSPLIT, $0-24
	GO_ARGS
	JMP	sync∕atomic·AddInt64(SB)

// And
TEXT	sync∕atomic·AndInt32(SB), NOSPLIT, $0-20
	GO_ARGS
	MOVD	$__tsan_go_atomic32_fetch_and(SB), O4
	CALL	racecallatomic<>(SB)
	RET

TEXT	sync∕atomic·AndInt64(SB), NOSPLIT, $0-24
	GO_ARGS
	MOVD	$__tsan_go_atomic64_fetch_and(SB), O4
	CALL	racecallatomic<>(SB)
	RET

TEXT	sync∕atomic·AndUint32(SB), NOSPLIT, $0-20
	GO_ARGS
	JMP	sync∕atomic·AndInt32(SB)

TEXT	sync∕atomic·AndUint64(SB), NOSPLIT, $0-24
	GO_ARGS
	JMP	sync∕atomic·AndInt64(SB)

TEXT	sync∕atomic·AndUintptr(SB), NOSPLIT, $0-24
	GO_ARGS
	JMP	sync∕atomic·AndInt64(SB)

// Or
TEXT	sync∕atomic·OrInt32(SB), NOSPLIT, $0-20
	GO_ARGS
	MOVD	$__tsan_go_atomic32_fetch_or(SB), O4
	CALL	racecallatomic<>(SB)
	RET

TEXT	sync∕atomic·OrInt64(SB), NOSPLIT, $0-24
	GO_ARGS
	MOVD	$__tsan_go_atomic64_fetch_or(SB), O4
	CALL	racecallatomic<>(SB)
	RET

TEXT	sync∕atomic·OrUint32(SB), NOSPLIT, $0-20
	GO_ARGS
	JMP	sync∕atomic·OrInt32(SB)

TEXT	sync∕atomic·OrUint64(SB), NOSPLIT, $0-24
	GO_ARGS
	JMP	sync∕atomic·OrInt64(SB)

TEXT	sync∕atomic·OrUintptr(SB), NOSPLIT, $0-24
	GO_ARGS
	JMP	sync∕atomic·OrInt64(SB)

// CompareAndSwap
TEXT	sync∕atomic·CompareAndSwapInt32(SB), NOSPLIT, $0-17
	GO_ARGS
	MOVD	$__tsan_go_atomic32_compare_exchange(SB), O4
	CALL	racecallatomic<>(SB)
	RET

TEXT	sync∕atomic·CompareAndSwapInt64(SB), NOSPLIT, $0-25
	GO_ARGS
	MOVD	$__tsan_go_atomic64_compare_exchange(SB), O4
	CALL	racecallatomic<>(SB)
	RET

TEXT	sync∕atomic·CompareAndSwapUint32(SB), NOSPLIT, $0-17
	GO_ARGS
	JMP	sync∕atomic·CompareAndSwapInt32(SB)

TEXT	sync∕atomic·CompareAndSwapUint64(SB), NOSPLIT, $0-25
	GO_ARGS
	JMP	sync∕atomic·CompareAndSwapInt64(SB)

TEXT	sync∕atomic·CompareAndSwapUintptr(SB), NOSPLIT, $0-25
	GO_ARGS
	JMP	sync∕atomic·CompareAndSwapInt64(SB)

// Generic atomic operation implementation.
// O4 = address of the tsan function to call.
TEXT	racecallatomic<>(SB), NOSPLIT, $0
	// This frame's prologue parked the caller's anchors here: the
	// sync/atomic thunk's frame, which its arguments sit above, and the
	// return address into the code that called it.
	MOVD	(ANCHOR_FP+STACK_BIAS)(RFP), L0
	ADD	$(FIXED_FRAME+STACK_BIAS), L0, L0	// the argument list
	MOVD	(L0), L1				// 1st arg is the addr

	// Trigger SIGSEGV early.
	MOVB	(L1), L2	// segv here if addr is bad

	// Check that addr is within [arenastart, arenaend) or within [racedatastart, racedataend).
	MOVD	runtime·racearenastart(SB), L2
	CMP	L2, L1
	BCSD	racecallatomic_data
	MOVD	runtime·racearenaend(SB), L2
	CMP	L2, L1
	BCSD	racecallatomic_ok
racecallatomic_data:
	MOVD	runtime·racedatastart(SB), L2
	CMP	L2, L1
	BCSD	racecallatomic_ignore
	MOVD	runtime·racedataend(SB), L2
	CMP	L2, L1
	BCCD	racecallatomic_ignore
racecallatomic_ok:
	// Addr is within the good range, call the atomic function.
	MOVD	g_racectx(g), O0		// goroutine context
	MOVD	(ANCHOR_LR+STACK_BIAS)(RFP), O1	// caller pc
	ADD	$8, O1, O1			// a raw LR is the CALL's address
	MOVD	O4, O2				// pc
	MOVD	L0, O3				// arguments
	CALL	racecall<>(SB)
	RET

racecallatomic_ignore:
	// Addr is outside the good range.
	// Call __tsan_go_ignore_sync_begin to ignore synchronization during the atomic op.
	// An attempt to synchronize on the address would cause crash.
	MOVD	O4, L4	// remember the original function; racecall<> keeps L4
	MOVD	g_racectx(g), O0		// goroutine context
	MOVD	$__tsan_go_ignore_sync_begin(SB), O4
	CALL	racecall<>(SB)
	// Call the atomic function. L0 is gone, but RFP is not: racecall<>
	// runs C in a window of its own, so this window comes back whole.
	MOVD	(ANCHOR_FP+STACK_BIAS)(RFP), O3
	ADD	$(FIXED_FRAME+STACK_BIAS), O3, O3	// arguments
	MOVD	g_racectx(g), O0		// goroutine context
	MOVD	(ANCHOR_LR+STACK_BIAS)(RFP), O1	// caller pc
	ADD	$8, O1, O1
	MOVD	L4, O2				// pc
	MOVD	L4, O4
	CALL	racecall<>(SB)
	// Call __tsan_go_ignore_sync_end.
	MOVD	g_racectx(g), O0		// goroutine context
	MOVD	$__tsan_go_ignore_sync_end(SB), O4
	CALL	racecall<>(SB)
	RET

// func runtime·racecall(void(*f)(...), ...)
// Calls C function f from race runtime and passes up to 4 arguments to it.
// The arguments are never heap-object-preserving pointers, so we pretend there are no arguments.
TEXT	runtime·racecall(SB), NOSPLIT|NOFRAME, $0-0
	MOVD	fn+0(FP), O4
	MOVD	arg0+8(FP), O0
	MOVD	arg1+16(FP), O1
	MOVD	arg2+24(FP), O2
	MOVD	arg3+32(FP), O3
	JMP	racecall<>(SB)

// Calls the C function in O4 on the g0 stack, with the arguments already in
// O0..O3. Clobbers L0..L3.
TEXT	racecall<>(SB), NOSPLIT|NOFRAME, $0-0
	MOVD	RSP, L0			// this window's stack pointer, to restore
	MOVD	LR, L1			// and our return address, which the C
					// call would otherwise overwrite

	// Move onto the g0 stack unless we are on a system stack already:
	// the race runtime needs far more of it than the nosplit budget of
	// the code that reached us, and this window is spilled while C runs.
	MOVD	g_m(g), L2
	MOVD	m_gsignal(L2), L3
	CMP	L3, g
	BED	onsystemstack
	MOVD	m_g0(L2), L3
	CMP	L3, g
	BED	onsystemstack
	MOVD	(g_sched+gobuf_sp)(L3), L2
	MOVD	L2, BSP

onsystemstack:
	MOVD	BSP, L2
	SUB	$208, L2		// the frame to call C on

	// Hand C a window of its own rather than calling in this one; see
	// the window comment in ·asmcgocall.
	SAVE	$-2047, L2, RSP

	// Describe the frame the way this ABI expects, so an unwinder that
	// starts in C walks back into the Go frames.
	MOVD	R30, (ANCHOR_FP)(BSP)	// %i6: the Go window's stack pointer
	MOVD	OLR, (120)(BSP)
	MOVD	OLR, (ANCHOR_LR)(BSP)

	// The C window: I0..I3 are the arguments, I4 the function. This L0
	// is not the one holding the stack pointer; that belongs to the
	// window we just left.
	MOVD	I0, O0
	MOVD	I1, O1
	MOVD	I2, O2
	MOVD	I3, O3
	MOVD	I4, L0
	CALL	(L0)
	MOVD	O0, I0		// C's result, into the Go window's %o0. This also
				// keeps RESTORE out of the call's delay slot, where
				// the assembler would otherwise take it for the
				// `ret; restore` idiom and hand C this window.
	RESTORE	ZR, ZR, ZR

	MOVD	L0, RSP
	MOVD	L1, LR
	RET

// C->Go callback thunk that allows to call runtime·racesymbolize from C code.
// Called from C in C's own register window, with the command code in O0 and
// its context in O1. See racecallback for command codes.
TEXT	runtime·racecallbackthunk(SB), NOSPLIT|NOFRAME, $0
	// Open a window of our own: this ABI keeps g, the frame anchor and
	// the temporaries in %l registers, which are the C caller's locals.
	SAVE	$-208, RSP, RSP

	// Present it as a frame, so a traceback taken inside the callback
	// walks out instead of decoding what SAVE left in these registers.
	MOVD	R30, RFP
	MOVD	R30, (ANCHOR_FP)(BSP)
	MOVD	OLR, (120)(BSP)
	MOVD	OLR, (ANCHOR_LR)(BSP)

	// Recover g from thread-local storage; the window we just opened has
	// none. A race build links runtime/cgo, so load_g is not a no-op.
	CALL	runtime·load_g(SB)

	// Handle command raceGetProcCmd (0) here.
	// First, code below assumes that we are on curg, while raceGetProcCmd
	// can be executed on g0. Second, it is called frequently, so will
	// benefit from this fast path.
	CMP	ZR, I0
	BNED	rest
	MOVD	g_m(g), L0
	MOVD	m_p(L0), L0
	MOVD	p_raceprocctx(L0), L0
	MOVD	L0, (I1)
	RESTORE	ZR, ZR, ZR
	RET

rest:
	// Set g = g0.
	MOVD	g_m(g), L0
	MOVD	m_g0(L0), L1
	CMP	L1, g
	BED	noswitch	// branch if already on g0
	MOVD	L1, g

noswitch:
	MOVD	I0, FIXED_FRAME+0(BSP)	// func arg
	MOVD	I1, FIXED_FRAME+8(BSP)	// func arg
	CALL	runtime·racecallback(SB)

	// All registers are smashed after Go code, reload.
	MOVD	g_m(g), L0
	MOVD	m_curg(L0), g	// g = m->curg
	RESTORE	ZR, ZR, ZR
	RET
