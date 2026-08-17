// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// stackBias is the constant SPARC V9 subtracts from %sp and %fp in the
// 64-bit ABI: the real stack top is %sp+stackBias. It must agree with
// sparc64.StackBias in cmd/internal/obj.
const stackBias = 2047

func osArchInit() {}

// SPARC keeps the SunOS values for the sigprocmask "how" argument
// instead of the ones the rest of Linux uses (0, 1, 2). Getting these
// wrong turns the mask-everything call that brackets fork into an
// unblock-everything call, and a signal arriving in that window kills
// the process with "signal received during fork".
const (
	_SIG_BLOCK   = 1
	_SIG_UNBLOCK = 2
	_SIG_SETMASK = 4

	// SPARC's SS_DISABLE is 2 as elsewhere, but sigaltstack takes the
	// SunOS flag values, so keep it beside its friends.
	_SS_DISABLE = 2
	_NSIG       = 65
)

type sigset [2]uint32

var sigset_all = sigset{^uint32(0), ^uint32(0)}

//go:nosplit
//go:nowritebarrierrec
func sigaddset(mask *sigset, i int) {
	(*mask)[(i-1)/32] |= 1 << ((uint32(i) - 1) & 31)
}

func sigdelset(mask *sigset, i int) {
	(*mask)[(i-1)/32] &^= 1 << ((uint32(i) - 1) & 31)
}

//go:nosplit
func sigfillset(mask *uint64) {
	*mask = ^uint64(0)
}
