// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && !sparc64

package syscall

// direntBufAlign is the alignment getdents64 requires of the buffer it
// is given. Only strict-alignment targets need more than one.
const direntBufAlign = 1
