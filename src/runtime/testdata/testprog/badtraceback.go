// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"runtime"
	"runtime/debug"
	"unsafe"
)

func init() {
	register("BadTraceback", BadTraceback)
}

func BadTraceback() {
	// Disable GC to prevent traceback at unexpected time.
	debug.SetGCPercent(-1)
	// Out of an abundance of caution, also make sure that there are
	// no GCs actively in progress.
	runtime.GC()

	// Run badLR1 on its own stack to minimize the stack size and
	// exercise the stack bounds logic in the hex dump.
	go badLR1()
	select {}
}

//go:noinline
func badLR1() {
	// We need two frames on LR machines because we'll smash this
	// frame's saved LR.
	badLR2(0)
}

//go:noinline
func badLR2(arg int) {
	// Smash the return PC or saved LR.
	lrOff := unsafe.Sizeof(uintptr(0))
	if runtime.GOARCH == "ppc64" || runtime.GOARCH == "ppc64le" {
		lrOff = 32 // FIXED_FRAME or sys.MinFrameSize
	}
	if runtime.GOARCH == "sparc64" {
		// Arguments start at the 176-byte minimum frame and the return
		// address chain slot sits at sp+136 (AnchorLR), clear of the
		// window save area the hardware spills %i7 into.
		lrOff = 176 - 136
	}
	lrPtr := (*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&arg)) - lrOff))
	badPC := uintptr(0xbad)
	if runtime.GOARCH == "sparc64" {
		// The anchor holds a raw %o7, the address of the call, which the
		// unwinder turns into a return address by adding 8. Store the
		// value that converts to 0xbad.
		badPC -= 8
	}
	*lrPtr = badPC

	runtime.KeepAlive(lrPtr) // prevent dead store elimination

	// Print a backtrace. This should include diagnostics for the
	// bad return PC and a hex dump.
	panic("backtrace")
}
