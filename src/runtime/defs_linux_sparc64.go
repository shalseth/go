// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Constants and layouts read from the target's own headers on an
// UltraSPARC T4 running Linux, rather than copied from another port.
// SPARC inherits a good deal of SunOS numbering, so many of these
// differ from the values every other Linux architecture uses. The ones
// worth noticing, because copying a generic value would produce
// failures a long way from their cause:
//
//	_ENOSYS is 90, not 38.
//	_SA_RESTART is 0x2, _SA_ONSTACK 0x1 and _SA_SIGINFO 0x200, where
//	  most ports use 0x10000000, 0x8000000 and 0x4.
//	_SIGBUS is 10, _SIGUSR1 30, _SIGUSR2 31, _SIGSTOP 17, _SIGCONT 19,
//	  _SIGCHLD 20, _SIGURG 16, _SIGSYS 12 and _SIGIO 23.
//	There is no SIGSTKFLT; signal 7 is SIGEMT.
//	The open flags differ throughout: O_CREAT is 0x200, O_TRUNC 0x400,
//	  O_NONBLOCK 0x4000 and O_CLOEXEC 0x400000.

package runtime

import "unsafe"

const (
	_EINTR  = 0x4
	_EAGAIN = 0xb
	_ENOMEM = 0xc
	_ENOSYS = 0x5a

	_PROT_NONE  = 0x0
	_PROT_READ  = 0x1
	_PROT_WRITE = 0x2
	_PROT_EXEC  = 0x4

	_MAP_ANON    = 0x20
	_MAP_PRIVATE = 0x2
	_MAP_FIXED   = 0x10

	_MADV_DONTNEED   = 0x4
	_MADV_FREE       = 0x8
	_MADV_HUGEPAGE   = 0xe
	_MADV_NOHUGEPAGE = 0xf
	_MADV_COLLAPSE   = 0x19

	_SA_RESTART  = 0x2
	_SA_ONSTACK  = 0x1
	_SA_RESTORER = 0x0 // unused on Linux/SPARC
	_SA_SIGINFO  = 0x200

	_SI_KERNEL = 0x80
	_SI_TIMER  = -0x2

	_SIGHUP  = 0x1
	_SIGINT  = 0x2
	_SIGQUIT = 0x3
	_SIGILL  = 0x4
	_SIGTRAP = 0x5
	_SIGABRT = 0x6
	// SPARC has no SIGSTKFLT. Signal 7 is SIGEMT, and the runtime's
	// signal tables are indexed by number, so it takes that slot.
	_SIGSTKFLT = 0x7
	_SIGFPE    = 0x8
	_SIGKILL   = 0x9
	_SIGBUS    = 0xa
	_SIGSEGV   = 0xb
	_SIGSYS    = 0xc
	_SIGPIPE   = 0xd
	_SIGALRM   = 0xe
	_SIGURG    = 0x10
	_SIGSTOP   = 0x11
	_SIGTSTP   = 0x12
	_SIGCONT   = 0x13
	_SIGCHLD   = 0x14
	_SIGTTIN   = 0x15
	_SIGTTOU   = 0x16
	_SIGIO     = 0x17
	_SIGXCPU   = 0x18
	_SIGXFSZ   = 0x19
	_SIGVTALRM = 0x1a
	_SIGPROF   = 0x1b
	_SIGWINCH  = 0x1c
	_SIGPWR    = 0x1d
	_SIGUSR1   = 0x1e
	_SIGUSR2   = 0x1f
	_SIGRTMIN  = 0x22

	_FPE_INTDIV = 0x1
	_FPE_INTOVF = 0x2
	_FPE_FLTDIV = 0x3
	_FPE_FLTOVF = 0x4
	_FPE_FLTUND = 0x5
	_FPE_FLTRES = 0x6
	_FPE_FLTINV = 0x7
	_FPE_FLTSUB = 0x8

	_BUS_ADRALN = 0x1
	_BUS_ADRERR = 0x2
	_BUS_OBJERR = 0x3

	_SEGV_MAPERR = 0x1
	_SEGV_ACCERR = 0x2

	_ITIMER_REAL    = 0x0
	_ITIMER_VIRTUAL = 0x1
	_ITIMER_PROF    = 0x2

	_CLOCK_THREAD_CPUTIME_ID = 0x3

	_SIGEV_THREAD_ID = 0x4

	_O_RDONLY   = 0x0
	_O_WRONLY   = 0x1
	_O_CREAT    = 0x200
	_O_TRUNC    = 0x400
	_O_NONBLOCK = 0x4000
	_O_CLOEXEC  = 0x400000
)

type timespec struct {
	tv_sec  int64
	tv_nsec int64
}

//go:nosplit
func (ts *timespec) setNsec(ns int64) {
	ts.tv_sec = ns / 1e9
	ts.tv_nsec = ns % 1e9
}

type timeval struct {
	tv_sec  int64
	tv_usec int64
}

func (tv *timeval) set_usec(x int32) {
	tv.tv_usec = int64(x)
}

// sigactiont matches the kernel's struct __new_sigaction from
// asm/signal.h. SPARC orders the fields handler, flags, restorer, mask,
// putting sa_restorer *before* sa_mask, unlike the architectures that
// place it last. sa_mask is a single word because _NSIG is 64 and
// _NSIG_BPW is 64, so _NSIG_WORDS is 1.
type sigactiont struct {
	sa_handler  uintptr
	sa_flags    uint64
	sa_restorer uintptr
	sa_mask     uint64
}

type siginfoFields struct {
	si_signo int32
	si_errno int32
	si_code  int32
	// below here is a union; si_addr is the only field we use
	si_addr uint64
}

type siginfo struct {
	siginfoFields

	// Pad struct to the max size in the kernel.
	_ [_si_max_size - unsafe.Sizeof(siginfoFields{})]byte
}

type itimerspec struct {
	it_interval timespec
	it_value    timespec
}

type itimerval struct {
	it_interval timeval
	it_value    timeval
}

type sigeventFields struct {
	value  uintptr
	signo  int32
	notify int32
	// below here is a union; sigev_notify_thread_id is the only field we use
	sigev_notify_thread_id int32
}

type sigevent struct {
	sigeventFields

	// Pad struct to the max size in the kernel.
	_ [_sigev_max_size - unsafe.Sizeof(sigeventFields{})]byte
}

// stackt is 24 bytes on linux/sparc64, with ss_flags at offset 8 and
// ss_size at 16; the compiler inserts the padding after ss_flags.
type stackt struct {
	ss_sp    *byte
	ss_flags int32
	ss_size  uintptr
}
