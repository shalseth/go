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
	MOVD	R11, 8(BSP)
	MOVD	R9, 16(BSP)
	MOVD	$8(BSP), R8
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

// func walltime() (sec int64, nsec int32)
TEXT runtime·walltime(SB),NOSPLIT,$32-12
	MOVD	$0, R8			// CLOCK_REALTIME
	MOVD	$8(BSP), R9
	SYS(SYS_clock_gettime)
	MOVD	8(BSP), R9
	MOVD	16(BSP), R10
	MOVD	R9, sec+0(FP)
	MOVW	R10, nsec+8(FP)
	RET

// func nanotime1() int64
TEXT runtime·nanotime1(SB),NOSPLIT,$32-8
	MOVD	$1, R8			// CLOCK_MONOTONIC
	MOVD	$8(BSP), R9
	SYS(SYS_clock_gettime)
	MOVD	8(BSP), R9
	MOVD	16(BSP), R10
	MOVD	$1000000000, R11
	MULD	R11, R9, R9
	ADD	R10, R9, R9
	MOVD	R9, ret+0(FP)
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
TEXT runtime·sigtramp(SB),NOSPLIT|NOFRAME|TOPFRAME,$0
	// Go code never executes SAVE, so the interrupted function's %l and
	// %i registers (including g in %l6) are live in the current window,
	// and rt_sigreturn restores only the globals, outs, and PC/nPC.
	// Shift to a fresh window for the handler: the interrupted window
	// was flushed to its own stack by the kernel at delivery, and the
	// RESTORE on the way out refills it untouched.
	//
	// g must move across the window shift by hand: copy it through a
	// global register.
	MOVD	g, R5
	SAVE	$-208, RSP, RSP		// 176 mandatory + args for sigtrampgo
	MOVD	R5, g

	// The arguments arrived in %o0-%o2 and are now %i0-%i2 (R24-R26).
	MOVW	R24, (176+0)(BSP)
	MOVD	R25, (176+8)(BSP)
	MOVD	R26, (176+16)(BSP)
	MOVD	$runtime·sigtrampgo(SB), R11
	CALL	(R11)

	// ret; restore: return to the kernel's restorer stub (in %i7, +8
	// as usual) while switching back to the interrupted window.
	JMPL	$8(OLR), ZR
	RESTORE	ZR, ZR, ZR

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
