// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linux

// Syscall numbers taken from the target's own <sys/syscall.h>. SPARC
// keeps much of the original SunOS numbering, so these do not resemble
// the asm-generic table most newer ports use.
const (
	SYS_CLOSE         = 6
	SYS_MPROTECT      = 74
	SYS_FCNTL         = 92
	SYS_PRCTL         = 147
	SYS_EPOLL_CTL     = 194
	SYS_EPOLL_PWAIT   = 309
	SYS_EPOLL_CREATE1 = 319
	SYS_EPOLL_PWAIT2  = 441
	SYS_EVENTFD2      = 318
	SYS_OPENAT        = 284
	SYS_PREAD64       = 67
	SYS_READ          = 3
	SYS_UNAME         = 189

	// SPARC numbers the O_* style flags differently from the
	// asm-generic values (0x800 for nonblock, 0x80000 for cloexec).
	EFD_NONBLOCK  = 0x4000
	EPOLL_CLOEXEC = 0x400000
	EFD_CLOEXEC   = 0x400000
	O_CLOEXEC     = 0x400000

	// Zero on 64-bit targets, as everywhere else: large-file offsets
	// are already the default.
	O_LARGEFILE = 0x0
)

// EpollEvent matches struct epoll_event on linux/sparc64: 16 bytes with
// data at offset 8, since SPARC does not use the packed layout x86-64
// does. Verified with offsetof on the target.
type EpollEvent struct {
	Events uint32
	_pad   uint32
	Data   [8]byte // unaligned uintptr
}
