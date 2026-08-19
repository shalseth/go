// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sparc64

const (
	maxAlign  = 32 // max data alignment
	minAlign  = 1  // min data alignment
	funcAlign = 8
)

// Used by ../internal/ld/dwarf.go.
//
// The SPARC DWARF register numbering follows the architectural register
// file: 0-7 are %g0-%g7, 8-15 are %o0-%o7, 16-23 are %l0-%l7 and 24-31
// are %i0-%i7. The stack pointer is %o6 and the link register %o7.
const (
	dwarfRegSP = 14
	dwarfRegLR = 15
	// The SPARC V9 ABI stores %sp biased by 2047, so the canonical
	// frame address is %sp+2047, as gcc also emits.
	dwarfCFABias = 2047
	// A frame's return address is not spilled at the bottom of the frame.
	// Every callee's prologue publishes its caller's %i7 into the caller's
	// register save area at [sp+120], which is also where the kernel's
	// window spill mirrors it.
	dwarfRAOffset = 120
)
