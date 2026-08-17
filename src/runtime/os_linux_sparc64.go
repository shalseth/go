// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// stackBias is the constant SPARC V9 subtracts from %sp and %fp in the
// 64-bit ABI: the real stack top is %sp+stackBias. It must agree with
// sparc64.StackBias in cmd/internal/obj.
const stackBias = 2047

func osArchInit() {}

const (
	// SPARC keeps the SunOS values for the sigprocmask "how"
	// argument instead of the ones the rest of Linux uses (0, 1, 2).
	// Getting these wrong turns the mask-everything call that
	// brackets fork into an unblock-everything call, and a signal
	// arriving in that window kills the process with "signal
	// received during fork".
	//
	// The sigset type itself comes from os_linux_be64.go: the kernel
	// sigset_t is a single 64-bit word, so the [2]uint32 form the
	// generic file uses would put signals 1-32 in the high half on a
	// big-endian machine and silently mask the wrong signals.
	_SIG_BLOCK   = 1
	_SIG_UNBLOCK = 2
	_SIG_SETMASK = 4
)
