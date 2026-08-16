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
	// Go frames do not execute SAVE/RESTORE, but they must still reserve
	// the SPARC V9 frame: a register-window overflow trap, a signal, or
	// an explicit FLUSHW spills the *current* window to %sp+StackBias
	// regardless of who set up the frame. That makes the 128-byte window
	// save area architecturally mandatory, not a calling convention.
	// The remaining 48 bytes are the outgoing-argument area, kept so Go
	// frames stay walkable by C tooling and the cgo boundary needs no
	// special case. Matches sparc64.MinStackFrameSize.
	_MinFrameSize = 176
	// SPARC V9 requires 16-byte stack alignment.
	_StackAlign = 16
)
