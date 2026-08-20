// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && sparc64

package runtime

import "internal/cpu"

// stackBias is the constant SPARC V9 subtracts from %sp and %fp in the
// 64-bit ABI: the real stack top is %sp+stackBias. It must agree with
// sparc64.StackBias in cmd/internal/obj.
const stackBias = 2047

func osArchInit() {}

func archauxv(tag, val uintptr) {
	switch tag {
	case _AT_HWCAP:
		cpu.HWCap = uint(val)
	}
}

const (
	_SS_DISABLE = 2
	_NSIG       = 65

	// SPARC keeps the SunOS values for the sigprocmask "how"
	// argument instead of the ones the rest of Linux uses (0, 1, 2).
	// Getting these wrong turns the mask-everything call that
	// brackets fork into an unblock-everything call, and a signal
	// arriving in that window kills the process with "signal
	// received during fork".
	_SIG_BLOCK   = 1
	_SIG_UNBLOCK = 2
	_SIG_SETMASK = 4
)

// The kernel sigset_t is a single 64-bit word, as on the other big-endian
// 64-bit ports: the [2]uint32 form in os_linux_generic.go would put signals
// 1-32 in the high half on a big-endian machine and silently mask the wrong
// ones. This duplicates os_linux_be64.go rather than sharing it, because the
// _SIG_* values above differ from the ones ppc64 and s390x use.

type sigset uint64

var sigset_all = sigset(^uint64(0))

//go:nosplit
//go:nowritebarrierrec
func sigaddset(mask *sigset, i int) {
	if i > 64 {
		throw("unexpected signal greater than 64")
	}
	*mask |= 1 << (uint(i) - 1)
}

func sigdelset(mask *sigset, i int) {
	if i > 64 {
		throw("unexpected signal greater than 64")
	}
	*mask &^= 1 << (uint(i) - 1)
}

//go:nosplit
func sigfillset(mask *uint64) {
	*mask = ^uint64(0)
}
