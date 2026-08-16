// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && !sparc64

package linux

// The close-on-exec flags equal O_CLOEXEC, which is 0x80000 everywhere
// except SPARC.
const (
	EPOLL_CLOEXEC = 0x80000
	EFD_CLOEXEC   = 0x80000
	O_CLOEXEC     = 0x80000
)
