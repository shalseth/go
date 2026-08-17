// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	"unsafe"

	"internal/abi"
	"internal/runtime/sys"
)

// gostartcall adjusts a gobuf as if it had executed a call to fn with
// context ctxt and then immediately done a Gosave.
func gostartcall(buf *gobuf, fn, ctxt unsafe.Pointer) {
	if buf.lr != 0 {
		throw("invalid use of gostartcall")
	}
	buf.lr = abi.FuncPCABI0(goexit) + sys.PCQuantum
	// The goroutine's first prologue stores OLR at [sp+120] as its
	// caller's return-address anchor; make that the goexit frame.
	buf.olr = buf.lr
	buf.pc = uintptr(fn)
	buf.ctxt = ctxt
}
