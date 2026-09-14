// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

TEXT errors(SB),$0
	// A RESTORE after a CALL would land in the call's delay slot and
	// rotate the window away before the callee runs. "ret; restore" is
	// the only place the restore belongs in a slot.
	CALL	foo(SB)			// ERROR "RESTORE immediately after CALL"
	RESTORE	ZR, ZR, ZR
	RET
