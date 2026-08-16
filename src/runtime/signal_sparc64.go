// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && sparc64

package runtime

import (
	"internal/abi"
	"internal/goarch"
	"unsafe"
)

func dumpregs(c *sigctxt) {
	r := c.regs()
	for i := 0; i < 8; i++ {
		print("g", i, "  ", hex(r.u_regs[i]), "\t")
		print("o", i, "  ", hex(r.u_regs[8+i]), "\n")
	}
	print("pc  ", hex(c.pc()), "\t")
	print("npc ", hex(c.npc()), "\n")
	print("tstate ", hex(r.tstate), "\n")
}

//go:nosplit
//go:nowritebarrierrec
func (c *sigctxt) sigpc() uintptr { return uintptr(c.pc()) }

func (c *sigctxt) sigsp() uintptr { return uintptr(c.sp()) }
func (c *sigctxt) siglr() uintptr { return uintptr(c.lr()) }
func (c *sigctxt) fault() uintptr { return uintptr(c.sigaddr()) }

// preparePanic sets up the stack to look like a call to sigpanic.
func (c *sigctxt) preparePanic(sig uint32, gp *g) {
	// Arrange the link register and PC so the panicking function looks
	// like it called sigpanic directly. The link register is always
	// spilled so a panic in a leaf function is handled too; this
	// smashes the frame, but execution is not returning there.
	sp := c.sp() - goarch.PtrSize
	c.set_sp(sp)
	*(*uint64)(unsafe.Pointer(uintptr(sp))) = c.lr()

	pc := gp.sigpc

	if shouldPushSigpanic(gp, pc, uintptr(c.lr())) {
		// Make it look like the faulting PC called sigpanic.
		c.set_lr(uint64(pc))
	}

	// In case we are panicking from external C code.
	c.set_g1(uint64(uintptr(unsafe.Pointer(gp))))
	c.set_pc(uint64(abi.FuncPCABIInternal(sigpanic)))
}

func (c *sigctxt) pushCall(targetPC, resumePC uintptr) {
	// Spill the link register, which is about to be clobbered to push
	// the call. The function being pushed restores it and resets the
	// stack pointer. gentraceback knows about this extra slot.
	sp := c.sp() - goarch.PtrSize
	c.set_sp(sp)
	*(*uint64)(unsafe.Pointer(uintptr(sp))) = c.lr()
	// Make the signalled function look like it calls targetPC at
	// resumePC. set_pc moves tnpc along with tpc, which SPARC needs.
	c.set_lr(uint64(resumePC))
	c.set_pc(uint64(targetPC))
}
