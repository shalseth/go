// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && sparc64

package syscall

// direntBufAlign is the alignment getdents64 requires of the buffer it
// is given. The kernel writes a 64-bit d_ino at the start of every
// entry, and SPARC traps on an unaligned store, which the kernel reports
// to us as EFAULT.
const direntBufAlign = 8
