// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

// Read from the target's own headers. The newer syscalls share the
// common numbering, but getrandom and copy_file_range sit where SPARC's
// own table put them.
const (
	getrandomTrap       uintptr = 347
	copyFileRangeTrap   uintptr = 357
	pidfdSendSignalTrap uintptr = 424
	pidfdOpenTrap       uintptr = 434
	openat2Trap         uintptr = 437
)
