// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// System calls and other system stubs for Linux/sparc64.
//
// The convention: the syscall number goes in %g1 (R1), arguments in
// %o0..%o5 (R8..R13), and the kernel is entered with "ta 0x6d". The
// result comes back in %o0. Unlike most ports the error indication is
// not a register value but the carry bit of the 32-bit condition codes:
// carry set means %o0 holds an errno, so the runtime's convention of
// returning a negative errno is produced by negating it.
//
// Syscall numbers were read from the target's own <sys/syscall.h>.
// SPARC keeps much of the original SunOS numbering, so they bear no
// relation to the asm-generic table.

#include "go_asm.h"
#include "go_tls.h"
#include "textflag.h"

#define SYS_exit		1
#define SYS_read		3
#define SYS_write		4
#define SYS_close		6
#define SYS_brk			17
#define SYS_getpid		20
#define SYS_sigaltstack		28
#define SYS_kill		37
#define SYS_mmap		71
#define SYS_munmap		73
#define SYS_madvise		75
#define SYS_mincore		78
#define SYS_setitimer		83
#define SYS_rt_sigreturn	101
#define SYS_rt_sigaction	102
#define SYS_rt_sigprocmask	103
#define SYS_futex		142
#define SYS_gettid		143
#define SYS_exit_group		188
#define SYS_tgkill		211
#define SYS_clone		217
#define SYS_sched_yield		245
#define SYS_nanosleep		249
#define SYS_clock_gettime	257
#define SYS_sched_getaffinity	260
#define SYS_timer_settime	262
#define SYS_timer_delete	265
#define SYS_timer_create	266
#define SYS_openat		284
#define SYS_pipe2		321
#define SYS_getrandom		347

// SYS puts a syscall number in %g1 and traps.
#define SYS(n) \
	MOVD	$n, R1;	\
	TA	$0x6d

// func exit(code int32)
TEXT runtime·exit(SB),NOSPLIT|NOFRAME,$0-4
	MOVW	code+0(FP), R8
	SYS(SYS_exit_group)
	RET

// func exitThread(wait *atomic.Uint32)
TEXT runtime·exitThread(SB),NOSPLIT|NOFRAME,$0-8
	MOVD	wait+0(FP), R9
	// The caller can no longer touch the stack once this store is
	// visible, so it is done last before the syscall.
	MOVW	ZR, (R9)
	MOVD	$0, R8
	SYS(SYS_exit)
	// Not reached: the thread is gone. Spin defensively rather than
	// fall through into whatever follows.
hang:
	MOVD	$0, R8
	JMP	hang

// func open(name *byte, mode, perm int32) int32
TEXT runtime·open(SB),NOSPLIT|NOFRAME,$0-20
	MOVD	$-100, R8	// AT_FDCWD
	MOVD	name+0(FP), R9
	MOVW	mode+8(FP), R10
	MOVW	perm+12(FP), R11
	SYS(SYS_openat)
	BCSW	openerr
	MOVW	R8, ret+16(FP)
	RET
openerr:
	MOVD	$-1, R8
	MOVW	R8, ret+16(FP)
	RET

// func closefd(fd int32) int32
TEXT runtime·closefd(SB),NOSPLIT|NOFRAME,$0-12
	MOVW	fd+0(FP), R8
	SYS(SYS_close)
	BCSW	closeerr
	MOVW	R8, ret+8(FP)
	RET
closeerr:
	MOVD	$-1, R8
	MOVW	R8, ret+8(FP)
	RET

// func write1(fd uintptr, p unsafe.Pointer, n int32) int32
TEXT runtime·write1(SB),NOSPLIT|NOFRAME,$0-28
	MOVD	fd+0(FP), R8
	MOVD	p+8(FP), R9
	MOVW	n+16(FP), R10
	SYS(SYS_write)
	BCSW	writeerr
	MOVW	R8, ret+24(FP)
	RET
writeerr:
	NEG	R8, R8		// the runtime wants a negative errno
	MOVW	R8, ret+24(FP)
	RET

// func read(fd int32, p unsafe.Pointer, n int32) int32
TEXT runtime·read(SB),NOSPLIT|NOFRAME,$0-28
	MOVW	fd+0(FP), R8
	MOVD	p+8(FP), R9
	MOVW	n+16(FP), R10
	SYS(SYS_read)
	BCSW	readerr
	MOVW	R8, ret+24(FP)
	RET
readerr:
	NEG	R8, R8
	MOVW	R8, ret+24(FP)
	RET

// func pipe2(flags int32) (r, w int32, errno int32)
TEXT runtime·pipe2(SB),NOSPLIT|NOFRAME,$0-20
	MOVD	$r+8(FP), R8
	MOVW	flags+0(FP), R9
	SYS(SYS_pipe2)
	BCSW	pipeerr
	MOVW	ZR, errno+16(FP)
	RET
pipeerr:
	NEG	R8, R8
	MOVW	R8, errno+16(FP)
	RET

// func usleep(usec uint32)
TEXT runtime·usleep(SB),NOSPLIT,$32-4
	MOVUW	usec+0(FP), R9
	MOVD	$1000, R10
	MULD	R10, R9, R9		// nanoseconds
	MOVD	$1000000000, R10
	UDIVD	R10, R9, R11		// seconds
	MOVD	R11, R12
	MULD	R10, R12, R12
	SUB	R12, R9, R9		// remaining nanoseconds
	// The timespec must live above the fixed 176-byte frame area:
	// on a context switch the kernel refills the current register
	// window from [sp+bias+0..127], so any data kept there comes back
	// as l/i register contents - corrupting the CALLER's registers
	// with tv_sec/tv_nsec.
	MOVD	R11, (176+8)(BSP)
	MOVD	R9, (176+16)(BSP)
	MOVD	$(176+8)(BSP), R8
	MOVD	$0, R9
	SYS(SYS_nanosleep)
	RET

// func gettid() uint32
TEXT runtime·gettid(SB),NOSPLIT|NOFRAME,$0-4
	SYS(SYS_gettid)
	MOVW	R8, ret+0(FP)
	RET

// func raise(sig uint32)
TEXT runtime·raise(SB),NOSPLIT|NOFRAME,$0
	SYS(SYS_gettid)
	MOVD	R8, R9			// tid
	MOVD	R8, R10
	SYS(SYS_getpid)
	MOVD	R8, R11			// pid
	MOVD	R11, R8
	MOVD	R10, R9
	MOVW	sig+0(FP), R10
	SYS(SYS_tgkill)
	RET

// func raiseproc(sig uint32)
TEXT runtime·raiseproc(SB),NOSPLIT|NOFRAME,$0
	SYS(SYS_getpid)
	MOVD	R8, R9
	MOVD	R9, R8
	MOVW	sig+0(FP), R9
	SYS(SYS_kill)
	RET

// func getpid() int
TEXT runtime·getpid(SB),NOSPLIT|NOFRAME,$0-8
	SYS(SYS_getpid)
	MOVD	R8, ret+0(FP)
	RET

// func tgkill(tgid, tid, sig int)
TEXT runtime·tgkill(SB),NOSPLIT|NOFRAME,$0-24
	MOVD	tgid+0(FP), R8
	MOVD	tid+8(FP), R9
	MOVD	sig+16(FP), R10
	SYS(SYS_tgkill)
	RET

// func setitimer(mode int32, new, old *itimerval)
TEXT runtime·setitimer(SB),NOSPLIT|NOFRAME,$0-24
	MOVW	mode+0(FP), R8
	MOVD	new+8(FP), R9
	MOVD	old+16(FP), R10
	SYS(SYS_setitimer)
	RET

// func timer_create(clockid int32, sevp *sigevent, timerid *int32) int32
TEXT runtime·timer_create(SB),NOSPLIT|NOFRAME,$0-28
	MOVW	clockid+0(FP), R8
	MOVD	sevp+8(FP), R9
	MOVD	timerid+16(FP), R10
	SYS(SYS_timer_create)
	BCSW	tcerr
	MOVW	R8, ret+24(FP)
	RET
tcerr:
	NEG	R8, R8
	MOVW	R8, ret+24(FP)
	RET

// func timer_settime(timerid int32, flags int32, new, old *itimerspec) int32
TEXT runtime·timer_settime(SB),NOSPLIT|NOFRAME,$0-28
	MOVW	timerid+0(FP), R8
	MOVW	flags+4(FP), R9
	MOVD	new+8(FP), R10
	MOVD	old+16(FP), R11
	SYS(SYS_timer_settime)
	BCSW	tserr
	MOVW	R8, ret+24(FP)
	RET
tserr:
	NEG	R8, R8
	MOVW	R8, ret+24(FP)
	RET

// func timer_delete(timerid int32) int32
TEXT runtime·timer_delete(SB),NOSPLIT|NOFRAME,$0-12
	MOVW	timerid+0(FP), R8
	SYS(SYS_timer_delete)
	BCSW	tderr
	MOVW	R8, ret+8(FP)
	RET
tderr:
	NEG	R8, R8
	MOVW	R8, ret+8(FP)
	RET

// func mincore(addr unsafe.Pointer, n uintptr, dst *byte) int32
TEXT runtime·mincore(SB),NOSPLIT|NOFRAME,$0-28
	MOVD	addr+0(FP), R8
	MOVD	n+8(FP), R9
	MOVD	dst+16(FP), R10
	SYS(SYS_mincore)
	BCSW	mcerr
	MOVW	R8, ret+24(FP)
	RET
mcerr:
	NEG	R8, R8
	MOVW	R8, ret+24(FP)
	RET

#define CLOCK_REALTIME	0
#define CLOCK_MONOTONIC	1

// Calling the vDSO from the flat-frame ABI needs three things that a
// plain syscall does not.
//
// It is ordinary C code, so it may execute SAVE and open a frame of its
// own; run it on g0's stack, which has room, rather than a goroutine's.
// %l and %i registers survive it either way - a callee that SAVEs gets a
// fresh window, and a leaf callee may touch only %o and %g - so state is
// carried across the call in %l0-%l3, after spilling what was there.
//
// While it runs, the current window is the vDSO's, so %l6 is not our g.
// A signal taken there would find garbage, so leave g at the base of the
// signal stack for sigFetchG, which knows to look when the faulting PC
// is inside the vDSO.
//
// Finally m.vdsoPC and m.vdsoSP let a SIGPROF traceback step over code
// that has no Go frame information.

// func walltime() (sec int64, nsec int32)
TEXT runtime·walltime(SB),NOSPLIT,$64-12
	MOVD	R16, (176+16)(BSP)
	MOVD	R17, (176+24)(BSP)
	MOVD	R18, (176+32)(BSP)
	MOVD	R19, (176+40)(BSP)

	MOVD	BSP, R16		// our stack pointer, unbiased
	MOVD	g_m(g), R17		// m

	MOVD	m_vdsoPC(R17), R3
	MOVD	m_vdsoSP(R17), R4
	MOVD	R3, (176+0)(BSP)
	MOVD	R4, (176+8)(BSP)
	ADD	$8, LR, R3		// our return address
	MOVD	R3, m_vdsoPC(R17)
	MOVD	$sec+0(FP), R4
	// The caller's frame pointer, not the address of its argument
	// area: outgoing arguments sit MinStackFrameSize into the
	// caller's frame, and sigprof unwinds from this as a frame SP.
	SUB	$176, R4
	MOVD	R4, m_vdsoSP(R17)

	MOVD	runtime·vdsoClockgettimeSym(SB), R2
	CMP	ZR, R2
	BED	walltime_fallback

	MOVD	BSP, R1
	MOVD	m_curg(R17), R5
	CMP	g, R5
	BNED	walltime_noswitch
	MOVD	m_g0(R17), R5
	MOVD	(g_sched+gobuf_sp)(R5), R1
walltime_noswitch:
	SUB	$224, R1
	AND	$~15, R1
	MOVD	R1, BSP

	MOVD	ZR, R18
	MOVD	m_gsignal(R17), R19
	CMP	ZR, R19
	BED	walltime_nosaveg
	CMP	g, R19
	BED	walltime_nosaveg
	MOVD	(g_stack+stack_lo)(R19), R18
	MOVD	g, (R18)
walltime_nosaveg:
	// The vDSO is C code and may clobber any register. R16 (our SP),
	// R17 (m) and R18 (the gsignal slot) are all live across the call,
	// so park them on the stack we are calling on and reload them on
	// return: the frame they were saved in at entry is not addressable
	// until BSP is back, and BSP comes from R16.
	MOVD	R16, (192+0)(BSP)
	MOVD	R17, (192+8)(BSP)
	MOVD	R18, (192+16)(BSP)
	MOVD	$CLOCK_REALTIME, R8
	MOVD	$176(BSP), R9		// timespec, clear of the window save area
	CALL	(R2)

	MOVD	(192+8)(BSP), R17
	MOVD	(192+16)(BSP), R18
	CMP	ZR, R18
	BED	walltime_noclearg
	MOVD	ZR, (R18)
walltime_noclearg:
	MOVD	176(BSP), R9
	MOVD	(176+8)(BSP), R10
	MOVD	(192+0)(BSP), R16
	MOVD	R16, BSP
	JMP	walltime_finish

walltime_fallback:
	MOVD	$CLOCK_REALTIME, R8
	MOVD	$(176+48)(BSP), R9
	SYS(SYS_clock_gettime)
	MOVD	(176+48)(BSP), R9
	MOVD	(176+56)(BSP), R10

walltime_finish:
	MOVD	(176+8)(BSP), R3
	MOVD	R3, m_vdsoSP(R17)
	MOVD	(176+0)(BSP), R3
	MOVD	R3, m_vdsoPC(R17)

	MOVD	R9, sec+0(FP)
	MOVW	R10, nsec+8(FP)

	MOVD	(176+16)(BSP), R16
	MOVD	(176+24)(BSP), R17
	MOVD	(176+32)(BSP), R18
	MOVD	(176+40)(BSP), R19
	RET

// func nanotime1() int64
TEXT runtime·nanotime1(SB),NOSPLIT,$64-8
	MOVD	R16, (176+16)(BSP)
	MOVD	R17, (176+24)(BSP)
	MOVD	R18, (176+32)(BSP)
	MOVD	R19, (176+40)(BSP)

	MOVD	BSP, R16
	MOVD	g_m(g), R17

	MOVD	m_vdsoPC(R17), R3
	MOVD	m_vdsoSP(R17), R4
	MOVD	R3, (176+0)(BSP)
	MOVD	R4, (176+8)(BSP)
	ADD	$8, LR, R3
	MOVD	R3, m_vdsoPC(R17)
	MOVD	$ret+0(FP), R4
	// The caller's frame pointer, not the address of its argument
	// area: outgoing arguments sit MinStackFrameSize into the
	// caller's frame, and sigprof unwinds from this as a frame SP.
	SUB	$176, R4
	MOVD	R4, m_vdsoSP(R17)

	MOVD	runtime·vdsoClockgettimeSym(SB), R2
	CMP	ZR, R2
	BED	nanotime_fallback

	MOVD	BSP, R1
	MOVD	m_curg(R17), R5
	CMP	g, R5
	BNED	nanotime_noswitch
	MOVD	m_g0(R17), R5
	MOVD	(g_sched+gobuf_sp)(R5), R1
nanotime_noswitch:
	SUB	$224, R1
	AND	$~15, R1
	MOVD	R1, BSP

	MOVD	ZR, R18
	MOVD	m_gsignal(R17), R19
	CMP	ZR, R19
	BED	nanotime_nosaveg
	CMP	g, R19
	BED	nanotime_nosaveg
	MOVD	(g_stack+stack_lo)(R19), R18
	MOVD	g, (R18)
nanotime_nosaveg:
	// The vDSO is C code and may clobber any register. R16 (our SP),
	// R17 (m) and R18 (the gsignal slot) are all live across the call,
	// so park them on the stack we are calling on and reload them on
	// return: the frame they were saved in at entry is not addressable
	// until BSP is back, and BSP comes from R16.
	MOVD	R16, (192+0)(BSP)
	MOVD	R17, (192+8)(BSP)
	MOVD	R18, (192+16)(BSP)
	MOVD	$CLOCK_MONOTONIC, R8
	MOVD	$176(BSP), R9
	CALL	(R2)

	MOVD	(192+8)(BSP), R17
	MOVD	(192+16)(BSP), R18
	CMP	ZR, R18
	BED	nanotime_noclearg
	MOVD	ZR, (R18)
nanotime_noclearg:
	MOVD	176(BSP), R9
	MOVD	(176+8)(BSP), R10
	MOVD	(192+0)(BSP), R16
	MOVD	R16, BSP
	JMP	nanotime_finish

nanotime_fallback:
	MOVD	$CLOCK_MONOTONIC, R8
	MOVD	$(176+48)(BSP), R9
	SYS(SYS_clock_gettime)
	MOVD	(176+48)(BSP), R9
	MOVD	(176+56)(BSP), R10

nanotime_finish:
	MOVD	(176+8)(BSP), R3
	MOVD	R3, m_vdsoSP(R17)
	MOVD	(176+0)(BSP), R3
	MOVD	R3, m_vdsoPC(R17)

	MOVD	$1000000000, R11
	MULD	R11, R9, R9
	ADD	R10, R9, R9
	MOVD	R9, ret+0(FP)

	MOVD	(176+16)(BSP), R16
	MOVD	(176+24)(BSP), R17
	MOVD	(176+32)(BSP), R18
	MOVD	(176+40)(BSP), R19
	RET

// func rtsigprocmask(how int32, new, old *sigset, size int32)
TEXT runtime·rtsigprocmask(SB),NOSPLIT|NOFRAME,$0-28
	MOVW	how+0(FP), R8
	MOVD	new+8(FP), R9
	MOVD	old+16(FP), R10
	MOVW	size+24(FP), R11
	SYS(SYS_rt_sigprocmask)
	BCSW	badsig
	RET
badsig:
	// A failed sigprocmask leaves signal handling in an unknown state;
	// crash immediately with a store to address zero.
	MOVD	R8, (ZR)
	RET

// Signal return trampoline. The sparc64 kernel takes the restorer as
// the fourth argument to rt_sigaction (the sa_restorer struct field is
// ignored) and sets it as the return address of the signal handler.
// A SPARC return jumps to the return address plus 8 (skipping the call
// and its delay slot), so the handler comes back at stub+8: the first
// two instructions are never executed and exist only as padding.
// glibc instead passes the address of its stub minus 8.
TEXT runtime·sigreturn_stub(SB),NOSPLIT|NOFRAME,$0
	RNOP
	RNOP
	SYS(SYS_rt_sigreturn)

// func rt_sigaction(sig uintptr, new, old *sigactiont, size uintptr) int32
TEXT runtime·rt_sigaction(SB),NOSPLIT|NOFRAME,$0-36
	MOVD	sig+0(FP), R8
	MOVD	new+8(FP), R9
	MOVD	old+16(FP), R10
	MOVD	$runtime·sigreturn_stub(SB), R11
	MOVD	size+24(FP), R12
	SYS(SYS_rt_sigaction)
	BCSW	rterr
	MOVW	R8, ret+32(FP)
	RET
rterr:
	NEG	R8, R8
	MOVW	R8, ret+32(FP)
	RET

// func sigfwd(fn uintptr, sig uint32, info *siginfo, ctx unsafe.Pointer)
TEXT runtime·sigfwd(SB),NOSPLIT,$0-32
	MOVW	sig+8(FP), R8
	MOVD	info+16(FP), R9
	MOVD	ctx+24(FP), R10
	MOVD	fn+0(FP), R11
	CALL	(R11)
	RET

// func sigtramp(signo, ureg, ctxt unsafe.Pointer)
//
// The kernel enters here with the signal number in %o0, a siginfo
// pointer in %o1 and, on this architecture, the address of the struct
// sigcontext embedded in the signal frame in %o2. See
// signal_linux_sparc64.go for that layout.
// The handler must run in the SAME register window as the interrupted
// code. A SAVE here would create a second user window, and the kernel
// gets the spill/refill bookkeeping of that second window wrong across
// syscalls made from the handler: a futex(FUTEX_WAKE) issued by the
// handler returned with %i6/%i7 refilled from the window of the
// interrupted futex(FUTEX_WAIT) frame, silently replacing the frame
// anchors of the running Go function. Everything after that returns to
// the wrong address.
//
// Instead keep one window for the whole process, exactly as ordinary Go
// code does, and preserve the interrupted function's %l0-%l7/%i0-%i7 by
// hand. g lives in %l6 and is saved with them, so it needs no shuffling
// through a global register either. The %g and %o registers are part of
// the signal context and rt_sigreturn restores them, so they are free.
TEXT runtime·sigtramp(SB),NOSPLIT|NOFRAME|TOPFRAME,$0
	// Run the handler in the INTERRUPTED register window.
	//
	// The obvious implementation - SAVE a fresh window for the handler -
	// is wrong on Linux/sparc64. A SAVE here leaves two live user
	// windows, and the kernel's window bookkeeping does not survive a
	// syscall made from inside the handler: the window comes back
	// holding the registers of the frame the signal interrupted (a
	// thread parked in futex(FUTEX_WAIT) hands its notesleep frame
	// anchors to the handler). The flat-frame ABI keeps its frame
	// anchors in %i6/%i7, so that silently redirects the handler's
	// returns.
	//
	// Go code never executes SAVE, so one window serves the whole
	// program; keep it that way here and preserve the interrupted
	// window by hand instead. %g registers need no saving: the kernel
	// restores them from the signal context.
	ADD	$-352, RSP

	// Save the interrupted window: %l0-%l7, %i0-%i7, and the return
	// address the kernel left in %o7 (CALL below clobbers it).
	MOVD	R16, (208+0)(BSP)
	MOVD	R17, (208+8)(BSP)
	MOVD	R18, (208+16)(BSP)
	MOVD	R19, (208+24)(BSP)
	MOVD	R20, (208+32)(BSP)
	MOVD	R21, (208+40)(BSP)
	MOVD	g, (208+48)(BSP)
	MOVD	R23, (208+56)(BSP)
	MOVD	R24, (208+64)(BSP)
	MOVD	R25, (208+72)(BSP)
	MOVD	R26, (208+80)(BSP)
	MOVD	R27, (208+88)(BSP)
	MOVD	R28, (208+96)(BSP)
	MOVD	R29, (208+104)(BSP)
	MOVD	R30, (208+112)(BSP)
	MOVD	R31, (208+120)(BSP)
	MOVD	LR, (208+128)(BSP)

	// sigtrampgo(sig uint32, info, ctx unsafe.Pointer); the kernel
	// passed them in %o0-%o2, which this window still holds.
	MOVW	R8, (176+0)(BSP)
	MOVD	R9, (176+8)(BSP)
	MOVD	R10, (176+16)(BSP)
	MOVD	$runtime·sigtrampgo(SB), R11
	CALL	(R11)

	// Restore the interrupted window.
	MOVD	(208+0)(BSP), R16
	MOVD	(208+8)(BSP), R17
	MOVD	(208+16)(BSP), R18
	MOVD	(208+24)(BSP), R19
	MOVD	(208+32)(BSP), R20
	MOVD	(208+40)(BSP), R21
	MOVD	(208+48)(BSP), g
	MOVD	(208+56)(BSP), R23
	MOVD	(208+64)(BSP), R24
	MOVD	(208+72)(BSP), R25
	MOVD	(208+80)(BSP), R26
	MOVD	(208+88)(BSP), R27
	MOVD	(208+96)(BSP), R28
	MOVD	(208+104)(BSP), R29
	MOVD	(208+112)(BSP), R30
	MOVD	(208+120)(BSP), R31
	// The restorer address goes to a %g register: every %l and %i is
	// live again, and the kernel reloads %g from the signal context.
	MOVD	(208+128)(BSP), R5

	ADD	$352, RSP
	// Return to the kernel's restorer stub, which lands at +8 as usual.
	JMPL	$8(R5), ZR
	RNOP

TEXT runtime·cgoSigtramp(SB),NOSPLIT|NOFRAME,$0
	JMP	runtime·sigtramp(SB)

// func mmap(addr unsafe.Pointer, n uintptr, prot, flags, fd int32, off uint32) (p unsafe.Pointer, err int)
TEXT runtime·mmap(SB),NOSPLIT|NOFRAME,$0
	MOVD	addr+0(FP), R8
	MOVD	n+8(FP), R9
	MOVW	prot+16(FP), R10
	MOVW	flags+20(FP), R11
	MOVW	fd+24(FP), R12
	MOVW	off+28(FP), R13
	SYS(SYS_mmap)
	BCSW	mmaperr
	MOVD	R8, p+32(FP)
	MOVD	ZR, err+40(FP)
	RET
mmaperr:
	MOVD	ZR, p+32(FP)
	MOVD	R8, err+40(FP)
	RET

// Unused alias kept for symmetry with cgo-capable ports.
TEXT runtime·sysMunmap(SB),NOSPLIT|NOFRAME,$0
	MOVD	addr+0(FP), R8
	MOVD	n+8(FP), R9
	SYS(SYS_munmap)
	BCSW	munmapfail
	RET
munmapfail:
	MOVD	R8, (ZR)	// crash
	RET

// func munmap(addr unsafe.Pointer, n uintptr)
TEXT runtime·munmap(SB),NOSPLIT|NOFRAME,$0
	MOVD	addr+0(FP), R8
	MOVD	n+8(FP), R9
	SYS(SYS_munmap)
	BCSW	munmapfault
	RET
munmapfault:
	MOVD	R8, (ZR)	// crash
	RET

// func madvise(addr unsafe.Pointer, n uintptr, flags int32) int32
TEXT runtime·madvise(SB),NOSPLIT|NOFRAME,$0
	MOVD	addr+0(FP), R8
	MOVD	n+8(FP), R9
	MOVW	flags+16(FP), R10
	SYS(SYS_madvise)
	BCSW	madverr
	MOVW	R8, ret+24(FP)
	RET
madverr:
	NEG	R8, R8
	MOVW	R8, ret+24(FP)
	RET

// func futex(addr unsafe.Pointer, op int32, val uint32, ts, addr2 unsafe.Pointer, val3 uint32) int32
TEXT runtime·futex(SB),NOSPLIT|NOFRAME,$0
	MOVD	addr+0(FP), R8
	MOVW	op+8(FP), R9
	MOVUW	val+12(FP), R10
	MOVD	ts+16(FP), R11
	MOVD	addr2+24(FP), R12
	MOVUW	val3+32(FP), R13
	SYS(SYS_futex)
	BCSW	futexerr
	MOVW	R8, ret+40(FP)
	RET
futexerr:
	NEG	R8, R8
	MOVW	R8, ret+40(FP)
	RET

// func clone(flags int32, stk, mp, gp, fn unsafe.Pointer) int32
TEXT runtime·clone(SB),NOSPLIT|NOFRAME,$0
	MOVW	flags+0(FP), R8
	MOVD	stk+8(FP), R9

	// Stash mp, gp and fn at the top of the child's stack, then hand
	// the kernel a stack pointer with the SPARC V9 bias applied and a
	// full 192-byte frame reserved below the stash — the same shape
	// glibc's clone uses. The child recovers the stash relative to its
	// own %sp.
	MOVD	mp+16(FP), R16
	MOVD	gp+24(FP), R17
	MOVD	fn+32(FP), R18
	SUB	$32, R9			// stash base
	MOVD	R16, 0(R9)
	MOVD	R17, 8(R9)
	MOVD	R18, 16(R9)
	SUB	$192, R9		// child frame
	SUB	$2047, R9		// stack bias

	MOVD	$0, R10			// parent tid
	MOVD	$0, R11			// child tid
	MOVD	$0, R12			// tls
	SYS(SYS_clone)
	BCSW	cloneerr

	// SPARC keeps the SunOS convention: both parent and child return
	// here, with %o1 zero in the parent and nonzero in the child.
	CMP	ZR, R9
	BNED	child
	MOVW	R8, ret+40(FP)
	RET
cloneerr:
	NEG	R8, R8
	MOVW	R8, ret+40(FP)
	RET

child:
	// On the child stack now; the stash sits 192 bytes above our
	// (unbiased) stack pointer.
	MOVD	(192+0)(BSP), R16	// mp
	MOVD	(192+8)(BSP), R17	// gp
	MOVD	(192+16)(BSP), R18	// fn

	CMP	ZR, R16
	BED	nog
	CMP	ZR, R17
	BED	nog

	// Store the new thread's id in m.procid.
	SYS(SYS_gettid)
	MOVD	R8, m_procid(R16)

	// Set up g and its m.
	MOVD	R16, g_m(R17)
	MOVD	R17, g
	CALL	runtime·save_g(SB)

nog:
	// Call fn. It must not return.
	CALL	(R18)
childexit:
	MOVD	$0, R8
	SYS(SYS_exit)
	JMP	childexit

// func sigaltstack(new, old *stackt)
TEXT runtime·sigaltstack(SB),NOSPLIT|NOFRAME,$0
	MOVD	new+0(FP), R8
	MOVD	old+8(FP), R9
	SYS(SYS_sigaltstack)
	BCSW	altstackfail
	RET
altstackfail:
	MOVD	R8, (ZR)	// crash
	RET

// func osyield()
TEXT runtime·osyield(SB),NOSPLIT|NOFRAME,$0
	SYS(SYS_sched_yield)
	RET

// func sched_getaffinity(pid, len uintptr, buf *uintptr) int32
TEXT runtime·sched_getaffinity(SB),NOSPLIT|NOFRAME,$0
	MOVD	pid+0(FP), R8
	MOVD	len+8(FP), R9
	MOVD	buf+16(FP), R10
	SYS(SYS_sched_getaffinity)
	BCSW	affierr
	MOVW	R8, ret+24(FP)
	RET
affierr:
	NEG	R8, R8
	MOVW	R8, ret+24(FP)
	RET

// func sbrk0() uintptr
TEXT runtime·sbrk0(SB),NOSPLIT|NOFRAME,$0-8
	MOVD	$0, R8
	SYS(SYS_brk)
	MOVD	R8, ret+0(FP)
	RET

