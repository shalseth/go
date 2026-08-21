// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgo

// On this port crosscall2 is a stub in gcc_sparc64.S that opens a
// register window and enters the Go side as crosscall2_flat, so that
// symbol has to be visible to the C half of the boundary too. See
// gcc_sparc64.S for why the window is there.

//go:cgo_export_static crosscall2_flat
//go:cgo_export_dynamic crosscall2_flat
