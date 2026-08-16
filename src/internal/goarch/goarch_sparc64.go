// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goarch

const (
	_ArchFamily = SPARC64
	// Linux/sparc64 uses an 8K base page on sun4u and sun4v.
	// Confirmed on an UltraSPARC T4 (sun4v): getconf PAGESIZE == 8192.
	_DefaultPhysPageSize = 8192
	// All SPARC V9 instructions are 4 bytes wide.
	_PCQuantum = 4
	// Flat-frame ABI: Go frames do not execute SAVE/RESTORE, so the only
	// system-reserved word is the saved link register (%o7). See
	// docs/sparc64-port.md, "Register windows", for why this differs from
	// the 176-byte SPARC V9 C frame.
	_MinFrameSize = 8
	// SPARC V9 requires 16-byte stack alignment.
	_StackAlign = 16
)
