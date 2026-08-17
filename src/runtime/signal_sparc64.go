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

// _MinFrameSizeWords is goarch.MinFrameSize in pointer-size words.
const _MinFrameSizeWords = 176 / goarch.PtrSize

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

// copyWindow copies the 128-byte register window save area from the
// old stack pointer to the new one when an injected call moves the
// stack pointer. The sparc64 kernel refills the current window's
// %l0-%l7/%i0-%i7 from [%sp+bias] on the way back to userspace, so
// after sigreturn every local and in register - including g, which
// lives in %l6 - is loaded from the NEW sp. If the window words are
// not moved along with sp, the resumed code runs with sixteen garbage
// registers.
//
//go:nosplit
func copyWindow(newsp, oldsp uint64) {
	dst := (*[16]uint64)(unsafe.Pointer(uintptr(newsp)))
	src := (*[16]uint64)(unsafe.Pointer(uintptr(oldsp)))
	*dst = *src
}

// preparePanic sets up the stack to look like a call to sigpanic.
func (c *sigctxt) preparePanic(sig uint32, gp *g) {
	// Arrange the link register and PC so the panicking function looks
	// like it called sigpanic directly. Push a MinFrameSize area with
	// the old link register spilled at its base: the generic traceback
	// and stack scanner expect exactly this shape for injected calls.
	oldsp := c.sp()
	sp := oldsp - goarch.PtrSize*(_MinFrameSizeWords)
	copyWindow(sp, oldsp)
	c.set_sp(sp)
	// sp+128: above the register window save area, see pushCall.
	*(*uint64)(unsafe.Pointer(uintptr(sp) + 128)) = c.lr()

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
	// Push a MinFrameSize area with the clobbered link register
	// spilled at its base; asyncPreempt pops both on the way out, and
	// the traceback machinery knows this shape for injected calls.
	oldsp := c.sp()
	sp := oldsp - goarch.PtrSize*(_MinFrameSizeWords)
	copyWindow(sp, oldsp)
	c.set_sp(sp)
	// The link register is spilled at sp+128, above the register
	// window save area: a nested signal spills the interrupted
	// window's registers to [sp+0..127], which would destroy a value
	// kept at sp+0.
	*(*uint64)(unsafe.Pointer(uintptr(sp) + 128)) = c.lr()
	// Make the signalled function look like it calls targetPC at
	// resumePC. LR gets resumePC-8: asyncPreempt behaves like a normal
	// framed function, and a SPARC return jumps to the return address
	// plus 8. set_pc moves tnpc along with tpc, which SPARC needs.
	c.set_lr(uint64(resumePC) - 8)
	c.set_pc(uint64(targetPC))
}
