// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"
#include "asm_sparc64.h"

DATA dbgbuf(SB)/8, $"\n\n"
GLOBL dbgbuf(SB), $8

TEXT runtime·rt0_go(SB),NOSPLIT|TOPFRAME,$0
	// BSP = stack; R9 = argc; R8 = argv

	// Push whatever windows the loader and libc startup left resident
	// out to their own frames, while %i6 still names the frame it
	// belongs to. From here on Go runs flat and rewrites %i6 as a frame
	// anchor, so any window still resident above this one would be
	// spilled through a pointer into a goroutine stack. crosscall1 and
	// crosscall2 do the same where a thread enters Go from C; this is
	// the initial thread's turn.
	FLUSHW

	SUB	$(FIXED_FRAME+16), BSP
	MOVD	$(FIXED_FRAME+0)(BSP), RT1
	MOVW	R9, (RT1) // argc
	MOVD	R8, FIXED_FRAME+8(BSP) // argv

	// create istack out of the given (operating system) stack.
	// _cgo_init may update stackguard.
	MOVD	$runtime·g0(SB), g
	MOVD BSP, RT1
	MOVD	$(-64*1024)(RT1), RT2
	MOVD	RT2, g_stackguard0(g)
	MOVD	RT2, g_stackguard1(g)
	MOVD	RT2, (g_stack+stack_lo)(g)
	MOVD	RT1, (g_stack+stack_hi)(g)

	// if there is a _cgo_init, call it using the gcc ABI.
	MOVD	_cgo_init(SB), R12
	CMP	ZR, R12
	BED	nocgo

	MOVD	TLS, O3			// arg 3: TLS base pointer
	MOVD	$runtime·tls_g(SB), O2 	// arg 2: &tls_g
	MOVD	$setg_gcc<>(SB), O1	// arg 1: setg
	MOVD	g, O0			// arg 0: G
	// C functions expect FIXED_FRAME bytes of space on caller stack frame.
	MOVD	BSP, L1
	SUB	$FIXED_FRAME, BSP
	CALL	(R12)
	MOVD	L1, BSP
	
	MOVD	_cgo_init(SB), R12
	CMP	ZR, R12
	BED	nocgo

nocgo:
	// update stackguard after _cgo_init
	MOVD	(g_stack+stack_lo)(g), R3
	ADD	$const_stackGuard, R3
	MOVD	R3, g_stackguard0(g)
	MOVD	R3, g_stackguard1(g)

	// set the per-goroutine and per-mach "registers"
	MOVD	$runtime·m0(SB), R3

	// save m->g0 = g0
	MOVD	g, m_g0(R3)
	// save m0 to g0->m
	MOVD	R3, g_m(g)

	CALL	runtime·check(SB)

	// argc and argv were stored at FIXED_FRAME+0 and +8 above, which
	// is exactly where runtime.args expects its arguments.
	CALL	runtime·args(SB)
	CALL	runtime·osinit(SB)
	CALL	runtime·schedinit(SB)

	// create a new goroutine to start program
	MOVD	$runtime·mainPC(SB), RT1		// entry
	SUB	$(32+FIXED_FRAME), BSP
	MOVD	RT1, FIXED_FRAME+0(BSP)
	MOVD	ZR, FIXED_FRAME+8(BSP)
	MOVD	ZR, FIXED_FRAME+16(BSP)
	MOVD	ZR, FIXED_FRAME+24(BSP)
	CALL	runtime·newproc(SB)
	ADD	$(32+FIXED_FRAME), BSP

	// start this M
	CALL	runtime·mstart(SB)

	MOVD	ZR, (ZR)	// boom
	UNDEF

DATA	runtime·mainPC+0(SB)/8,$runtime·main<ABIInternal>(SB)
GLOBL	runtime·mainPC(SB),RODATA,$8

// mstart is the ABI0 entry point for a new m; the work is in mstart0.
// mstart must keep a frame: its RET is reached, not dead code. When cgo is
// in use every m runs on a pthread created by _cgo_sys_thread_start, so g0's
// stack is the OS thread's and mexit takes its osStack path - it returns,
// unwinding mstart0 back into here. Frameless, the CALL above would have
// overwritten LR with its own address and the RET would jump to itself,
// spinning the thread forever instead of letting it exit; the pthread's TSD
// destructors would then never run. With a frame, the prologue parks the
// return address in the frame's OLR slot and the epilogue restores it, so
// the RET lands back in crosscall1 and the thread unwinds into C and exits.
TEXT runtime·mstart(SB),NOSPLIT|TOPFRAME,$0
	CALL	runtime·mstart0(SB)
	RET	// reached on a cgo thread's exit; returns into crosscall1

TEXT runtime·breakpoint(SB),NOSPLIT|NOFRAME,$0-0
	TA	$0x81
	RET

TEXT runtime·asminit(SB),NOSPLIT|NOFRAME,$0-0
	RET

TEXT runtime·gogo(SB), NOSPLIT|NOFRAME, $0-8
	MOVD	buf+0(FP), R5
	MOVD	gobuf_g(R5), g
	CALL	runtime·save_g(SB)

	MOVD	0(g), R4	// make sure g is not nil
	// The frame anchor registers RFP (%i6) and OLR (%i7) are part of a
	// goroutine's context in the flat-frame ABI: a framed function's
	// epilogue unwinds through them, and the kernel's register-window
	// spill keeps [sp+bias+112/120] mirroring them. gobuf.bp holds the
	// parked frame's RFP; gobuf.lr holds its OLR (the frame's own
	// return address). LR gets the same value: for a fresh goroutine
	// (gostartcall) the entry prologue captures its return address
	// from LR, while for a parked frame LR is dead anyway.
	//
	// Order matters here in a way no PCDATA can express: the kernel
	// spills the live %i6/%i7 to [%sp+bias+112/120] on every trap,
	// with no Go code involved. If %sp already pointed at the incoming
	// goroutine's frame while %i6/%i7 still held scheduler values, any
	// interrupt in that window would plant foreign pointers in the
	// frame's anchor slots, to be read by some unwind minutes later.
	// So: install the anchors while %sp still points at the dead
	// scheduler frame (spills there are harmless), and switch %sp
	// last, when a spill writes exactly the right values.
	MOVD	gobuf_lr(R5), LR
	MOVD	gobuf_olr(R5), OLR
	MOVD	gobuf_bp(R5), RFP
	MOVD	gobuf_ctxt(R5), CTXT
	MOVD	gobuf_pc(R5), R8
	MOVD	gobuf_sp(R5), R1
	MOVD	ZR, gobuf_sp(R5)
	MOVD	ZR, gobuf_lr(R5)
	MOVD	ZR, gobuf_olr(R5)
	MOVD	ZR, gobuf_bp(R5)
	MOVD	ZR, gobuf_ctxt(R5)
	CMP	ZR, ZR // set condition codes for == test, needed by stack split
	MOVD	R1, BSP
	JMPL	R8, ZR

// void mcall(fn func(*g))
// Switch to m->g0's stack, call fn(g).
// Fn must never return. It should gogo(&g->sched)
// to keep running g.
TEXT runtime·mcall(SB), NOSPLIT|NOFRAME, $0-8
	// Save caller state in g->sched.
	// gobuf_pc must be an exact resume address for gogo. LR holds the
	// address of the CALL instruction itself (SPARC %o7), so the
	// instruction to resume at is LR+8 (skipping the delay slot).
	MOVD	BSP, TMP
	MOVD	TMP, (g_sched+gobuf_sp)(g)
	ADD	$8, LR, TMP
	MOVD	TMP, (g_sched+gobuf_pc)(g)
	// Save the caller's frame anchors so gogo can rebuild them:
	// mcall is NOFRAME, so RFP/OLR still belong to the calling frame.
	MOVD	OLR, (g_sched+gobuf_lr)(g)
	MOVD	OLR, (g_sched+gobuf_olr)(g)
	MOVD	RFP, (g_sched+gobuf_bp)(g)
	MOVD	g, (g_sched+gobuf_g)(g)

	// Switch to m->g0 & its stack, call fn.
	MOVD	g, R3
	MOVD	g_m(g), R8
	MOVD	m_g0(R8), g
	CALL	runtime·save_g(SB)
	CMP	g, R3
	BNED	ok
	JMP	runtime·badmcall(SB)
ok:
	MOVD	fn+0(FP), CTXT			// context
	MOVD	0(CTXT), R4			// code pointer
	MOVD	(g_sched+gobuf_sp)(g), TMP
	MOVD	TMP, BSP	// sp = m->g0->sched.sp
	SUB	$16, BSP
	MOVD	R3, (176+0)(BSP)
	MOVD	$0, (176+8)(BSP)
	CALL	(R4)
	JMP	runtime·badmcall2(SB)

// systemstack_switch is a dummy routine that systemstack leaves at the bottom
// of the G stack. We need to distinguish the routine that
// lives at the bottom of the G stack from the one that lives
// at the top of the system stack because the one at the top of
// the system stack terminates the stack walk (see topofstack()).
TEXT runtime·systemstack_switch(SB), NOSPLIT, $0-0
	UNDEF
	CALL	(LR)	// make sure this function is not leaf
	RET

// func systemstack(fn func())
TEXT runtime·systemstack(SB), NOSPLIT, $0-8
	MOVD	fn+0(FP), R3	// R3 = fn
	MOVD	R3, CTXT		// context
	MOVD	g_m(g), R4	// R4 = m

	MOVD	m_gsignal(R4), R5	// R5 = gsignal
	CMP	g, R5
	BED	noswitch

	MOVD	m_g0(R4), R5	// R5 = g0
	CMP	g, R5
	BED	noswitch

	MOVD	m_curg(R4), R8
	CMP	g, R8
	BED	switch

	// Bad: g is not gsignal, not g0, not curg. What is it?
	// Hide call from linker nosplit analysis.
	MOVD	$runtime·badsystemstack(SB), R3
	CALL	(R3)

switch:
	// save our state in g->sched. Pretend to
	// be systemstack_switch if the G stack is scanned.
	// The fake pc points past the prologue's frame push (offset 16, the
	// instruction after the generated ADD that carries the Spadj), so
	// the unwinder sees a 176-byte frame at the saved sp and lands
	// exactly on the caller's frame; the saved lr is our return
	// address into that caller.
	MOVD	$runtime·systemstack_switch(SB), R8
	ADD	$16, R8
	MOVD	R8, (g_sched+gobuf_pc)(g)
	MOVD	BSP, TMP
	MOVD	TMP, (g_sched+gobuf_sp)(g)
	MOVD	OLR, (g_sched+gobuf_lr)(g)
	MOVD	OLR, (g_sched+gobuf_olr)(g)
	// RFP is part of the goroutine's context: if this goroutine is
	// later resumed through gogo (it can be preempted or parked while
	// the system stack runs), gogo reloads RFP from gobuf.bp. Leaving
	// a stale value here hands the resumed code a frame anchor into a
	// stack that copystack has since moved.
	MOVD	RFP, (g_sched+gobuf_bp)(g)
	MOVD	g, (g_sched+gobuf_g)(g)

	// switch to g0
	MOVD	R5, g
	CALL	runtime·save_g(SB)
	MOVD	(g_sched+gobuf_sp)(g), R3
	// make it look like mstart called systemstack on g0, to stop traceback
	SUB	$16, R3
	AND	$~15, R3
	MOVD	$runtime·mstart(SB), R4
	MOVD	R4, 0(R3)
	MOVD	R3, BSP

	// call target function
	MOVD	0(CTXT), R3	// code pointer
	CALL	(R3)

	// switch back to g
	MOVD	g_m(g), R3
	MOVD	m_curg(R3), g
	CALL	runtime·save_g(SB)
	// Restore the frame anchor BEFORE the stack pointer: the kernel
	// spills the live %i6/%i7 to [%sp+bias+112/120] on every trap, so
	// the moment %sp names the goroutine frame, %i6/%i7 must already
	// hold that frame's anchors or an interrupt plants g0 values in
	// the frame's anchor slots.
	MOVD	(g_sched+gobuf_bp)(g), RFP
	MOVD	(g_sched+gobuf_sp)(g), TMP
	MOVD	TMP, BSP
	MOVD	$0, (g_sched+gobuf_sp)(g)
	MOVD	$0, (g_sched+gobuf_bp)(g)
	RET

noswitch:
	// already on m stack, just call directly.
	// Tail-call so that systemstack does not remain as an intermediate
	// frame: an unwinder reaching it jumps to the goroutine stack, and
	// nested systemstack calls - which all land here after the first -
	// would each lose the frames above them. The epilogue leaves CTXT
	// and R3 alone, so fn still gets its context.
	MOVD	0(CTXT), R3	// code pointer
	RET	R3


// func switchToCrashStack0(fn func())
TEXT runtime·switchToCrashStack0(SB), NOSPLIT, $0-8
	MOVD	fn+0(FP), R3
	MOVD	R3, CTXT	// context register
	MOVD	g_m(g), R4	// R4 = curm

	// set g to gcrash
	MOVD	$runtime·gcrash(SB), g	// g = &gcrash
	CALL	runtime·save_g(SB)
	MOVD	R4, g_m(g)	// g.m = curm
	MOVD	g, m_g0(R4)	// curm.g0 = g

	// Switch to the crash stack. A SPARC frame's register-window and
	// argument save area sits *above* %sp, so the new stack pointer
	// has to leave a whole minimum frame below stack.hi or the first
	// window spill runs off the end of the stack; clone reserves the
	// same 192 bytes when it hands a child its stack. Assigning to BSP
	// applies the stack bias.
	MOVD	(g_stack+stack_hi)(g), R5
	SUB	$192, R5	// MinStackFrameSize, rounded up to 16-byte alignment
	AND	$~15, R5
	MOVD	R5, BSP

	// Root the frame chain here. fn's prologue publishes our RFP and
	// OLR into this frame's anchor slots, and the kernel spills the
	// same two registers there on any trap, so zeroing them is what
	// stops a traceback at this frame instead of sending it back into
	// the stack we just abandoned.
	MOVD	ZR, RFP
	MOVD	ZR, OLR

	// call target function
	MOVD	0(CTXT), R3	// code pointer
	CALL	(R3)

	// should never return
	CALL	runtime·abort(SB)
	UNDEF
/*
 * support for morestack
 */

// Called during function prolog when more stack is needed.
// Caller has already loaded:
// R3 prolog's LR
//
// The traceback routines see morestack on a g0 as being
// the top of a stack (for example, morestack calling newstack
// calling the scheduler calling newm calling gc), so we must
// record an argument size. For that purpose, it has no arguments.
TEXT runtime·morestack(SB),NOSPLIT|NOFRAME,$0-0
	// Called from f.
	// Set g->sched to context in f.
	// gogo resumes at gobuf_pc exactly; LR is the address of the CALL
	// morestack in f's split block, so resume at LR+8 (the jump back
	// to the start of f). gobuf_lr stays the raw %o7 value of f's
	// caller: it is restored into LR, and f returns through it with
	// the usual +8.
	//
	// This has to come before the g0 and gsignal checks below:
	// badmorestackg0 tracebacks from g->sched, so leaving a stale
	// one there makes it print an unrelated stack.
	MOVD	CTXT, (g_sched+gobuf_ctxt)(g)
	MOVD	BSP, TMP
	MOVD	TMP, (g_sched+gobuf_sp)(g)
	ADD	$8, LR, TMP
	MOVD	TMP, (g_sched+gobuf_pc)(g)
	MOVD	R3, (g_sched+gobuf_lr)(g)
	// f's prologue has not run, so RFP and OLR still belong to f's
	// caller. OLR must be preserved exactly: while SP is still at f's
	// entry, the kernel window spill mirrors %i7 into [sp+bias+120],
	// the very slot holding the caller's return address anchor.
	MOVD	RFP, (g_sched+gobuf_bp)(g)
	MOVD	OLR, TMP
	MOVD	TMP, (g_sched+gobuf_olr)(g)

	// Cannot grow scheduler stack (m->g0).
	MOVD	g_m(g), R8
	MOVD	m_g0(R8), R4
	CMP	g, R4
	BNED	3(PC)
	CALL	runtime·badmorestackg0(SB)
	JMP	runtime·abort(SB)

	// Cannot grow signal stack (m->gsignal).
	MOVD	m_gsignal(R8), R4
	CMP	g, R4
	BNED	3(PC)
	CALL	runtime·badmorestackgsignal(SB)
	JMP	runtime·abort(SB)

	// Called from f.
	// Set m->morebuf to f's callers.
	MOVD	R3, (m_morebuf+gobuf_pc)(R8)	// f's caller's PC
	MOVD	BSP, TMP
	MOVD	TMP, (m_morebuf+gobuf_sp)(R8)	// f's caller's BSP
	MOVD	g, (m_morebuf+gobuf_g)(R8)

	// Push any register window the kernel is holding out to the stack
	// frame it belongs to, while that stack is still mapped. A buffered
	// window is written back through the stack pointer it was buffered
	// with, and newstack is about to copy this goroutine's stack and
	// free the old one: a window still owed to the old stack is lost,
	// and it takes g and whatever else the window held with it.
	FLUSHW

	// Call newstack on m->g0's stack.
	MOVD	m_g0(R8), g
	CALL	runtime·save_g(SB)
	MOVD	(g_sched+gobuf_sp)(g), TMP
	MOVD	TMP, BSP
	CALL	runtime·newstack(SB)

	// Not reached, but make sure the return PC from the call to newstack
	// is still in this function, and not the beginning of the next.
	UNDEF

TEXT runtime·morestack_noctxt(SB),NOSPLIT|NOFRAME,$0-0
	MOVD	ZR, CTXT
	JMP	runtime·morestack(SB)

// reflectcall has no variable-sized frames, so a small number of
// constant-sized-frame functions encode a few bits of size in the PC.
#define DISPATCH(NAME,MAXSIZE)		\
	MOVD	$MAXSIZE, TMP;		\
	CMP	TMP, RT1;		\
	BGD	3(PC);			\
	MOVD	$NAME(SB), RT1;		\
	JMPL	RT1, ZR
// Note: can't just "JMP NAME(SB)" - bad inlining results.

TEXT ·reflectcall(SB), NOSPLIT|NOFRAME, $0-48
	MOVUW	frameSize+32(FP), RT1
	DISPATCH(runtime·call16, 16)
	DISPATCH(runtime·call32, 32)
	DISPATCH(runtime·call64, 64)
	DISPATCH(runtime·call128, 128)
	DISPATCH(runtime·call256, 256)
	DISPATCH(runtime·call512, 512)
	DISPATCH(runtime·call1024, 1024)
	DISPATCH(runtime·call2048, 2048)
	DISPATCH(runtime·call4096, 4096)
	DISPATCH(runtime·call8192, 8192)
	DISPATCH(runtime·call16384, 16384)
	DISPATCH(runtime·call32768, 32768)
	DISPATCH(runtime·call65536, 65536)
	DISPATCH(runtime·call131072, 131072)
	DISPATCH(runtime·call262144, 262144)
	DISPATCH(runtime·call524288, 524288)
	DISPATCH(runtime·call1048576, 1048576)
	DISPATCH(runtime·call2097152, 2097152)
	DISPATCH(runtime·call4194304, 4194304)
	DISPATCH(runtime·call8388608, 8388608)
	DISPATCH(runtime·call16777216, 16777216)
	DISPATCH(runtime·call33554432, 33554432)
	DISPATCH(runtime·call67108864, 67108864)
	DISPATCH(runtime·call134217728, 134217728)
	DISPATCH(runtime·call268435456, 268435456)
	DISPATCH(runtime·call536870912, 536870912)
	DISPATCH(runtime·call1073741824, 1073741824)
	MOVD	$runtime·badreflectcall(SB), R1
	JMPL	R1, ZR

// There is no register ABI on sparc64, so regArgs is ignored and the
// spill/unspill helpers of other ports are not needed.
#define CALLFN(NAME,MAXSIZE)			\
TEXT NAME(SB), WRAPPER, $MAXSIZE-48;		\
	NO_LOCAL_POINTERS;			\
	/* copy arguments to the callee's stack space */ \
	MOVD	stackArgs+16(FP), R3;		\
	MOVUW	stackArgsSize+24(FP), R4;	\
	MOVD	BSP, R5;			\
	ADD	$FIXED_FRAME, R5;		\
	ADD	R5, R4, R4;			\
	CMP	R4, R5;				\
	BED	6(PC);				\
	MOVUB	(R3), R9;			\
	ADD	$1, R3;				\
	MOVB	R9, (R5);			\
	ADD	$1, R5;				\
	JMP	-6(PC);				\
	/* call function */			\
	MOVD	f+8(FP), CTXT;			\
	MOVD	(CTXT), R1;			\
	PCDATA	$PCDATA_StackMapIndex, $0;	\
	CALL	(R1);				\
	/* copy return values back */		\
	MOVD	stackArgsType+0(FP), R8;	\
	MOVD	stackArgs+16(FP), R3;		\
	MOVUW	stackArgsSize+24(FP), R4;	\
	MOVUW	stackRetOffset+28(FP), R9;	\
	MOVD	BSP, R5;			\
	ADD	$FIXED_FRAME, R5;		\
	ADD	R9, R5, R5;			\
	ADD	R9, R3, R3;			\
	SUB	R9, R4, R4;			\
	CALL	callRet<>(SB);			\
	RET

// callRet copies return values back at the end of call*. This is a
// separate function so it can allocate stack space for the arguments
// to reflectcallmove. It does not follow the Go ABI; it expects its
// arguments in registers: R8 = argtype, R3 = dst, R5 = src, R4 = size.
TEXT callRet<>(SB), NOSPLIT, $48-0
	MOVD	R8, (FIXED_FRAME+0)(BSP)
	MOVD	R3, (FIXED_FRAME+8)(BSP)
	MOVD	R5, (FIXED_FRAME+16)(BSP)
	MOVD	R4, (FIXED_FRAME+24)(BSP)
	MOVD	ZR, (FIXED_FRAME+32)(BSP)
	CALL	runtime·reflectcallmove(SB)
	RET

CALLFN(·call16, 16)
CALLFN(·call32, 32)
CALLFN(·call64, 64)
CALLFN(·call128, 128)
CALLFN(·call256, 256)
CALLFN(·call512, 512)
CALLFN(·call1024, 1024)
CALLFN(·call2048, 2048)
CALLFN(·call4096, 4096)
CALLFN(·call8192, 8192)
CALLFN(·call16384, 16384)
CALLFN(·call32768, 32768)
CALLFN(·call65536, 65536)
CALLFN(·call131072, 131072)
CALLFN(·call262144, 262144)
CALLFN(·call524288, 524288)
CALLFN(·call1048576, 1048576)
CALLFN(·call2097152, 2097152)
CALLFN(·call4194304, 4194304)
CALLFN(·call8388608, 8388608)
CALLFN(·call16777216, 16777216)
CALLFN(·call33554432, 33554432)
CALLFN(·call67108864, 67108864)
CALLFN(·call134217728, 134217728)
CALLFN(·call268435456, 268435456)
CALLFN(·call536870912, 536870912)
CALLFN(·call1073741824, 1073741824)


// AES hashing not implemented for SPARC64.
TEXT runtime·procyieldAsm(SB),NOSPLIT,$0-0
	RD	CCR, R2
	RET

// void jmpdefer(fv, sp);
// called from deferreturn.
// 1. grab stored LR for caller
// 2. sub 4 bytes to get back to BL deferreturn
// 3. BR to fn
// gosave<> parks the goroutine that is about to leave for C. Its only
// caller is asmcgocall, and the parked PC it records is asmcgocall's own
// return address, which asmcgocall keeps in R21 - not a PC inside
// asmcgocall itself.
//
// The distinction matters because asmcgocall writes SP, and the unwinder
// refuses to walk through a function that does. A callback that grows the
// goroutine's stack has to walk exactly here: copystack unwinds the whole
// stack, crosses the C boundary through the anchors cgocallback plants,
// and would arrive inside asmcgocall. Parking one frame out lands it in
// cgocall instead, which is an ordinary framed function. Since asmcgocall
// carries no frame of its own, nothing is lost by skipping it: the stack
// pointer recorded here is already the frame cgocall is standing on.
TEXT gosave<>(SB),NOSPLIT|NOFRAME,$0
	ADD	$8, R21, TMP2
	MOVD	TMP2, (g_sched+gobuf_pc)(g)
	MOVD	BSP, TMP
	MOVD	TMP, (g_sched+gobuf_sp)(g)
	MOVD	$0, (g_sched+gobuf_lr)(g)
	MOVD	$0, (g_sched+gobuf_ctxt)(g)
	RET

// func asmcgocall(fn, arg unsafe.Pointer) int32
// Call fn(arg) on the scheduler stack,
// aligned appropriately for the gcc ABI.
// See cgocall.go for more details.
TEXT ·asmcgocall(SB),NOSPLIT|NOFRAME,$0-20
	MOVD	LR, R21			// gosave<> below clobbers LR
	MOVD	fn+0(FP), R19
	MOVD	arg+8(FP), R20
	MOVD	g, R24			// the g we came in on

	// How deep our stack pointer and frame anchor sit below the top of
	// the stack we came in on. A callback can grow this goroutine's
	// stack, which moves it, so both are rebuilt from the top of the
	// stack as it is on the way back rather than simply kept: this ABI
	// addresses locals through the frame pointer, and a caller resumed
	// with a frame pointer into a stack that has moved reads and writes
	// its own locals in freed memory.
	//
	// All of this is kept in this window's own registers. They are the
	// one place the C call cannot reach: the hardware preserves a
	// window across a call made from it, and spills it to its own
	// frame. A stack slot here is not safe - anything that ends up with
	// a stack pointer near this frame writes through it.
	MOVD	(g_stack+stack_hi)(g), R16
	MOVD	R16, R17
	MOVD	BSP, R18
	SUB	R18, R16		// depth to the stack pointer
	SUB	RFP, R17		// depth to the frame anchor

	// Move onto m->g0's stack, unless we are on it already (creating an
	// OS thread comes in that way). This window's stack pointer has to
	// end up on a stack that cannot move: a callback can grow the
	// goroutine stack, and a window spilled to a stack that is then
	// freed gets refilled from memory that no longer holds it.
	MOVD	g_m(g), R9
	MOVD	m_g0(R9), R28		// g0, kept where the O-registers below do not alias it
	CMP	R28, g
	BED	oncurrentstack
	CALL	gosave<>(SB)
	MOVD	(g_sched+gobuf_sp)(R28), R18
	MOVD	R18, BSP
	MOVD	R28, g
	CALL	runtime·save_g(SB)

oncurrentstack:
	MOVD	BSP, R13
	SUB	$208, R13		// the frame to call C on

	// Hand control to C in a register window of its own.
	//
	// %i6 is this ABI's frame anchor, but it is also the hardware's
	// stack pointer for the window above - they are the same registers -
	// and a window is spilled and refilled through it. Called in this
	// window, C fills the register file underneath a window whose %i6
	// names a goroutine frame while its own stack pointer has been
	// walked onto the g0 stack; the window above then gets spilled into
	// live goroutine memory, over the anchor slots a frame keeps at
	// [sp+bias+112/120], and the frame it lands on returns through
	// whatever the spill left there.
	//
	// Opening a window gives C one whose %i6 is this window's stack
	// pointer - a g0 frame, on a stack that never moves - and leaves
	// this window's stack pointer alone for as long as C runs.
	MOVD	R20, O0			// arg, C's first argument
	MOVD	R19, O1			// fn
	MOVD	R28, O3			// g0
	SAVE	$-2047, R13, RSP

	// The C window: I0 = arg, I1 = fn, I3 = g0.
	MOVD	I3, g

	// Describe this frame the way every other frame in this ABI is
	// described, so an unwinder that starts here walks back into the Go
	// frames rather than reading whatever the slots held, and so a
	// refill of this window finds the right %i6. Only the slots are
	// written: %i7 is the window below's link register, and a C callee
	// that returns with "ret" would jump through it.
	MOVD	RFP, (112)(BSP)
	MOVD	OLR, (120)(BSP)

	MOVD	I0, O0
	MOVD	I0, FIXED_FRAME(BSP)	// the C ABI's argument save slot
	MOVD	I1, R16
	CALL	(R16)
	MOVD	O0, I0			// errno, into the Go window's %o0
	RESTORE	ZR, ZR, ZR

	// Back in the Go window, still on the g0 stack and still carrying
	// g0. Put the goroutine back and rebuild both anchors from the top
	// of its stack as it is now - a callback may have moved it.
	MOVD	R24, g
	CALL	runtime·save_g(SB)
	MOVD	(g_stack+stack_hi)(g), R18
	MOVD	R18, R19
	SUB	R17, R19
	MOVD	R19, RFP
	SUB	R16, R18
	MOVD	R18, BSP

	MOVD	R21, LR
	MOVW	R8, ret+16(FP)
	RET

// cgocallback is where C code calls into Go: crosscall2 arrives here
// with the callback's function, its argument frame and a context.
//
// The work is to find or borrow an m, switch from the C stack we are
// standing on to the goroutine stack that m owns, run cgocallbackg
// there, and put everything back. gobuf.sp holds an unbiased address
// here, which is why it moves through BSP rather than RSP.
//
// Unlike the other ports this does not plant the parked PC below the
// goroutine's stack pointer to make a traceback continue seamlessly
// into the Go frames beneath the C ones. A traceback taken inside a
// callback therefore stops at the boundary instead of spanning it.
//
// func cgocallback(fn, frame unsafe.Pointer, ctxt uintptr)
TEXT ·cgocallback(SB),NOSPLIT,$32-24
	NO_LOCAL_POINTERS

	// A nil fn means: do not call anything, just drop the m. frame
	// carries the g to restore. This is how a dying thread releases
	// what it borrowed.
	MOVD	fn+0(FP), R1
	CMP	ZR, R1
	BNED	loadg
	MOVD	frame+8(FP), g
	JMP	dropm

loadg:
	// Recover g from thread-local storage. A nil g means Go never
	// created this thread, so borrow an m for the duration.
	CALL	runtime·load_g(SB)
	CMP	ZR, g
	BED	needm

	MOVD	g_m(g), R2
	MOVD	R2, savedm-8(SP)
	JMP	havem

needm:
	MOVD	g, savedm-8(SP)		// g is zero, and so is m
	MOVD	$runtime·needAndBindM(SB), R3
	CALL	(R3)			// indirect, to hide from the nosplit check

	// Give m->g0->sched.sp a usable value, so that a panic inside the
	// callback has somewhere to unwind to. It points one frame below
	// this one for the reason spelled out at havem.
	MOVD	g_m(g), R2
	MOVD	m_g0(R2), R3
	MOVD	BSP, R4
	SUB	$352, R4
	MOVD	R4, (g_sched+gobuf_sp)(R3)
	MOVD	RFP, (g_sched+gobuf_bp)(R3)

havem:
	// Save m->g0->sched.sp and point it at the C stack we are on, so
	// that the switch back has something to return to.
	//
	// It points a frame's worth below this one rather than at it. A
	// stack pointer handed to the runtime is a caller's stack pointer,
	// and in this ABI a callee writes into its caller's frame: the
	// caller's frame pointer and return address go to 112 and 120,
	// and outgoing arguments start at 176. Anything that runs on the
	// g0 stack while the callback is in flight - newstack when the
	// goroutine grows its stack, systemstack, the scheduler - would
	// otherwise write through this frame's own reserved slots and
	// locals, losing the address this function has to return to.
	MOVD	m_g0(R2), R3
	MOVD	(g_sched+gobuf_sp)(R3), R4
	MOVD	R4, savedsp-16(SP)
	MOVD	BSP, R11
	SUB	$352, R11

	// unwindm restores m->g0->sched.sp from the m.g0 stack if the
	// callback panics, reading it at MinFrameSize above the value
	// installed here - the contract described at the top of
	// cgocall.go. Leave it there as well as in this frame: the frame
	// copy serves the ordinary return below, and the panic unwinds
	// through the copy on the stack.
	MOVD	R4, (176)(R11)

	MOVD	R11, (g_sched+gobuf_sp)(R3)

	// Move to m->curg and its stack.
	MOVD	m_curg(R2), g
	CALL	runtime·save_g(SB)
	MOVD	(g_sched+gobuf_sp)(g), R4	// parked sp, unbiased
	MOVD	(g_sched+gobuf_pc)(g), R5	// where the goroutine resumes
	MOVD	(g_sched+gobuf_bp)(g), R8	// its frame anchor

	// Gather the arguments before the stack moves out from under them.
	MOVD	fn+0(FP), R16
	MOVD	frame+8(FP), R17
	MOVD	ctxt+16(FP), R18

	// The anchors this frame will advertise: the caller's stack pointer
	// in raw form, and a return address into the parked code. R6 and R7
	// are %g6 and %g7 - the kernel's thread-info register and the
	// thread pointer - so scratch comes from elsewhere. gobuf.pc
	// is an exact resume address, while an OLR holds the address of the
	// CALL, which every reader turns back into a resume address by
	// adding 8 - so it goes in eight bytes lower.
	SUB	$2047, R4, R9
	SUB	$8, R5

	// Open a frame on the goroutine stack: the ABI minimum, which
	// covers the window save area and the reserved slots, plus the
	// three arguments for cgocallbackg and a slot holding this
	// frame's own stack pointer for the return trip.
	SUB	$208, R4
	MOVD	BSP, R20
	MOVD	R4, BSP
	MOVD	R20, (176+24)(BSP)

	MOVD	R9, RFP
	MOVD	R5, OLR
	MOVD	RFP, (112)(BSP)
	MOVD	OLR, (120)(BSP)

	MOVD	R16, (176+0)(BSP)
	MOVD	R17, (176+8)(BSP)
	MOVD	R18, (176+16)(BSP)
	MOVD	$runtime·cgocallbackg(SB), R19
	CALL	(R19)			// indirect: we are on another stack now

	// Restore the goroutine's parked stack pointer. cgocallbackg
	// leaves BSP where it found it, so undoing the frame is enough.
	MOVD	BSP, R4
	ADD	$208, R4
	MOVD	R4, (g_sched+gobuf_sp)(g)

	// Back to m->g0 and this frame's stack pointer, which comes from
	// the slot parked above rather than from m->g0->sched.sp:
	// systemstack repoints that at itself whenever it runs, and a
	// callback that grows the goroutine's stack is certain to run it.
	MOVD	g_m(g), R2
	MOVD	m_g0(R2), g
	CALL	runtime·save_g(SB)
	MOVD	(176+24)(BSP), R4
	MOVD	R4, BSP

	// The anchors planted above were RFP and OLR, which are not spare
	// registers: they are this frame's own frame pointer and return
	// address, and cgocallbackg dutifully restored them on the way
	// out. Both have to come back before anything downstream uses
	// them - the local slots below and the epilogue are addressed
	// through RFP, and the epilogue returns through OLR.
	//
	// RFP is the raw frame pointer, one frame above the stack pointer
	// just restored: the ABI minimum plus this frame's declared
	// locals. OLR is where the prologue spilled it.
	SUB	$2047, R4, R9
	ADD	$(176+32), R9
	MOVD	R9, RFP
	MOVD	(120)(BSP), OLR

	MOVD	savedsp-16(SP), R4
	MOVD	R4, (g_sched+gobuf_sp)(g)

	// If there was an m on entry, it was ours to keep; otherwise the
	// borrowed one goes back, unless a pthread key holds it for the
	// next call from this thread.
	MOVD	savedm-8(SP), R8
	CMP	ZR, R8
	BNED	droppedm

	MOVD	_cgo_pthread_key_created(SB), R8
	CMP	ZR, R8
	BED	dropm
	MOVD	(R8), R8
	CMP	ZR, R8
	BNED	droppedm

dropm:
	MOVD	$runtime·dropm(SB), R3
	CALL	(R3)

droppedm:
	RET

// Called from the C code cgo generates, so it obeys the C ABI: the
// result goes in %o0, and every register C expects to survive a call
// has to survive.
//
// That is more than it looks. No register window is opened here - this
// port's Go code runs flat - so the local and in registers still belong
// to the C caller while Go code runs on top of them, and the Go code
// below is free to use them. gcc keeps values in %l0-%l7 and %i0-%i5
// across a call precisely because the window is meant to preserve them,
// so all fourteen are saved by hand, as in crosscall1 and crosscall2.
// The prologue takes care of RFP, OLR and LR, and %g1-%g5 are volatile
// in the C ABI.
//
// The stakes are concrete: asmcgocall parks the goroutine's g in %l3
// across its call into C, so clobbering that register loses g for the
// rest of the program.
#define TOPSAVED (176 + 24)

TEXT _cgo_topofstack(SB),NOSPLIT,$144
	NO_LOCAL_POINTERS

	MOVD	R16, (TOPSAVED+0)(BSP)	// %l0
	MOVD	R17, (TOPSAVED+8)(BSP)
	MOVD	R18, (TOPSAVED+16)(BSP)
	MOVD	R19, (TOPSAVED+24)(BSP)
	MOVD	R20, (TOPSAVED+32)(BSP)
	MOVD	R21, (TOPSAVED+40)(BSP)
	MOVD	g, (TOPSAVED+48)(BSP)	// %l6
	MOVD	R23, (TOPSAVED+56)(BSP)	// %l7
	MOVD	R24, (TOPSAVED+64)(BSP)	// %i0
	MOVD	R25, (TOPSAVED+72)(BSP)
	MOVD	R26, (TOPSAVED+80)(BSP)	// %i2, which load_g clobbers
	MOVD	R27, (TOPSAVED+88)(BSP)
	MOVD	R28, (TOPSAVED+96)(BSP)
	MOVD	R29, (TOPSAVED+104)(BSP)

	CALL	runtime·load_g(SB)
	MOVD	g_m(g), R1
	MOVD	m_curg(R1), R1
	MOVD	(g_stack+stack_hi)(R1), R1

	MOVD	(TOPSAVED+0)(BSP), R16
	MOVD	(TOPSAVED+8)(BSP), R17
	MOVD	(TOPSAVED+16)(BSP), R18
	MOVD	(TOPSAVED+24)(BSP), R19
	MOVD	(TOPSAVED+32)(BSP), R20
	MOVD	(TOPSAVED+40)(BSP), R21
	MOVD	(TOPSAVED+48)(BSP), g
	MOVD	(TOPSAVED+56)(BSP), R23
	MOVD	(TOPSAVED+64)(BSP), R24
	MOVD	(TOPSAVED+72)(BSP), R25
	MOVD	(TOPSAVED+80)(BSP), R26
	MOVD	(TOPSAVED+88)(BSP), R27
	MOVD	(TOPSAVED+96)(BSP), R28
	MOVD	(TOPSAVED+104)(BSP), R29

	MOVD	R1, R8			// %o0: the C return register
	RET

// void setg(G*); set g. for use by needm.
TEXT runtime·setg(SB), NOSPLIT, $0-8
	MOVD	gg+0(FP), g
	// This only happens if iscgo, so jump straight to save_g
	CALL	runtime·save_g(SB)
	RET

// void setg_gcc(G*); set g called from gcc
TEXT setg_gcc<>(SB),NOSPLIT,$16
	MOVD	R8, g
	MOVD	TMP, savedTMP-8(SP)
	CALL	runtime·save_g(SB)
	MOVD	savedTMP-8(SP), TMP
	RET

// The signal handler recognises an abort by the faulting PC (isAbortPC),
// so the fault must happen at an instruction inside this function. A jump
// to a bad address faults with PC=0 instead, which findfunc cannot resolve:
// the abort is then downgraded to a recoverable nil-pointer panic and every
// traceback below it reads "unknown pc 0x0".
TEXT runtime·abort(SB),NOSPLIT|NOFRAME,$0-0
	MOVD	ZR, (ZR)	// boom
	UNDEF

// func cputicks() int64
//
// Read %stick, not %tick. %tick counts the strand's own cycles and the
// strands do not agree: on an UltraSPARC T4, passing a token around a
// ring of thread-locked goroutines, a third of the handoffs saw %tick
// go *backwards* across the happens-before edge, by as much as 6ms.
// Callers assume a process-wide timebase - the debug log merges its
// shards by this value, and the mutex profile subtracts two reads taken
// on different threads - so a per-strand counter produces garbage
// orderings and negative-turned-huge durations.
//
// %stick is driven from a system-wide reference and measured zero
// backwards steps over the same experiment. It ticks slower (about 1GHz
// against 2.85GHz here), which costs resolution the runtime does not
// need: ticksPerSecond calibrates against nanotime either way.
TEXT runtime·cputicks(SB),NOSPLIT,$0-0
	RD	STICK, R1
	MOVD	R1, ret+0(FP)
	RET

TEXT runtime·goexit(SB),NOSPLIT|NOFRAME|TOPFRAME,$0-0
	RNOP	// +0
	RNOP	// +4: goexit+PCQuantum, the value planted in LR
	RNOP	// +8
	CALL	runtime·goexit1(SB)	// +12: where the RET lands; does not return

// TODO(aram):
TEXT runtime·addmoduledata(SB),NOSPLIT,$0-0
	MOVD	runtime·lastmoduledatap(SB), R1
	MOVD	R8, moduledata_next(R1)
	MOVD	R8, runtime·lastmoduledatap(SB)
	RET

TEXT gcWriteBarrier<>(SB),NOSPLIT,$176
	// Save the registers clobbered by the fast path.
	MOVD	R1, (176+0)(BSP)
	MOVD	R2, (176+8)(BSP)
	MOVD	R3, (176+16)(BSP)
retry:
	MOVD	g_m(g), R1
	MOVD	m_p(R1), R1
	MOVD	(p_wbBuf+wbBuf_next)(R1), R2
	MOVD	(p_wbBuf+wbBuf_end)(R1), R3
	// Increment wbBuf.next position.
	ADD	R25, R2, R2
	// Is the buffer full? Flush if next > end.
	CMP	R3, R2
	BGUD	flush
	// Commit to the larger buffer.
	MOVD	R2, (p_wbBuf+wbBuf_next)(R1)
	// Make the return value: the original next position.
	SUB	R25, R2, R25
	// Restore registers.
	MOVD	(176+0)(BSP), R1
	MOVD	(176+8)(BSP), R2
	MOVD	(176+16)(BSP), R3
	RET

flush:
	// Save all allocatable integer registers: these could be
	// clobbered by wbBufFlush and were not saved by the caller.
	// R1, R2, R3 are already saved.
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
	MOVD	R21, (176+128)(BSP)
	MOVD	R24, (176+136)(BSP)
	MOVD	R25, (176+144)(BSP)
	MOVD	CTXT, (176+152)(BSP)
	// R15 is the link register, saved by the prologue and declared
	// clobbered by LoweredWB. TMP and TMP2 are assembler temporaries.
	// g is not clobbered by wbBufFlush.

	CALL	runtime·wbBufFlush(SB)

	MOVD	(176+24)(BSP), R4
	MOVD	(176+32)(BSP), R5
	MOVD	(176+40)(BSP), R8
	MOVD	(176+48)(BSP), R9
	MOVD	(176+56)(BSP), R10
	MOVD	(176+64)(BSP), R11
	MOVD	(176+72)(BSP), R12
	MOVD	(176+80)(BSP), R13
	MOVD	(176+88)(BSP), R16
	MOVD	(176+96)(BSP), R17
	MOVD	(176+104)(BSP), R18
	MOVD	(176+112)(BSP), R19
	MOVD	(176+120)(BSP), R20
	MOVD	(176+128)(BSP), R21
	MOVD	(176+136)(BSP), R24
	MOVD	(176+144)(BSP), R25
	MOVD	(176+152)(BSP), CTXT
	JMP	retry

TEXT runtime·gcWriteBarrier1(SB),NOSPLIT|NOFRAME,$0
	MOVD	$8, R25
	JMP	gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier2(SB),NOSPLIT|NOFRAME,$0
	MOVD	$16, R25
	JMP	gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier3(SB),NOSPLIT|NOFRAME,$0
	MOVD	$24, R25
	JMP	gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier4(SB),NOSPLIT|NOFRAME,$0
	MOVD	$32, R25
	JMP	gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier5(SB),NOSPLIT|NOFRAME,$0
	MOVD	$40, R25
	JMP	gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier6(SB),NOSPLIT|NOFRAME,$0
	MOVD	$48, R25
	JMP	gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier7(SB),NOSPLIT|NOFRAME,$0
	MOVD	$56, R25
	JMP	gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier8(SB),NOSPLIT|NOFRAME,$0
	MOVD	$64, R25
	JMP	gcWriteBarrier<>(SB)

// panicBounds is called by compiled code at a failed bounds check, with
// the PCData entry describing which registers or constants hold the
// operands. It saves the 16 candidate registers — in exactly the order
// of the compiler's boundsRegs table (R1-R5, R8-R13, R16-R20; R15 is
// the link register and excluded) — and hands runtime.panicBounds64 the
// call PC and a pointer to the save area. The registers may hold dead
// pointers; panicBounds64 only reads the ones the PCData names.
TEXT runtime·panicBounds(SB),NOSPLIT,$144-0
	NO_LOCAL_POINTERS
	MOVD	R1, (176+16)(BSP)
	MOVD	R2, (176+24)(BSP)
	MOVD	R3, (176+32)(BSP)
	MOVD	R4, (176+40)(BSP)
	MOVD	R5, (176+48)(BSP)
	MOVD	R8, (176+56)(BSP)
	MOVD	R9, (176+64)(BSP)
	MOVD	R10, (176+72)(BSP)
	MOVD	R11, (176+80)(BSP)
	MOVD	R12, (176+88)(BSP)
	MOVD	R13, (176+96)(BSP)
	MOVD	R16, (176+104)(BSP)
	MOVD	R17, (176+112)(BSP)
	MOVD	R18, (176+120)(BSP)
	MOVD	R19, (176+128)(BSP)
	MOVD	R20, (176+136)(BSP)

	// The prologue moved the incoming return address into OLR, but it
	// is a raw %o7: the address of the CALL to panicBounds itself.
	// panicBounds64 wants a return address - it reads the bounds
	// metadata with pcdatavalue at pc-1 and expects that to land inside
	// the call - so step over the call and its delay slot. R1 is already
	// saved above, so it is free as scratch.
	ADD	$8, OLR, R1
	MOVD	R1, (176+0)(BSP)
	MOVD	$(176+16)(BSP), R1	// pointer to the save area
	MOVD	R1, (176+8)(BSP)
	CALL	runtime·panicBounds64(SB)
	RET	// not reached

// publicationBarrier must order prior stores before later stores.
// Linux/sparc64 runs in TSO, where stores are already ordered, so like
// amd64 this is a plain return.
TEXT runtime·publicationBarrier(SB),NOSPLIT|NOFRAME,$0-0
	RET
