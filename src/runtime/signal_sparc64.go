// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && sparc64

package runtime

import (
	"internal/abi"
	"internal/goarch"
	"internal/runtime/atomic"
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
// siglr returns the link register as a return address. A SPARC CALL
// writes the address of the CALL itself into %o7 and the callee returns
// with JMPL %o7+8, past the delay slot, so the raw register is eight
// bytes short of a return address. Every consumer wants a return
// address: the unwinder seeds frame.lr with this and resolves it as
// "minus one is inside the CALL", and traceback.initAt applies the same
// conversion where it takes the link register from gp.sched. Left raw,
// a signal taken in a leaf resolves its caller one instruction before
// the call, which is outside the range of any call the compiler
// inlined there, so inlined frames vanish from the traceback.
//
// MIPS, the other delay-slot architecture Go supports, needs no such
// conversion: JAL stores PC+8 in $31 directly.
func (c *sigctxt) siglr() uintptr {
	lr := uintptr(c.lr())
	if lr == 0 {
		return 0
	}
	return lr + 8
}
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
		// Make it look like the faulting PC called sigpanic. The link
		// register holds a raw %o7 call address, which the traceback
		// converts to a return PC by adding 8, so store pc-8: the
		// reconstructed return PC is then exactly the faulting
		// instruction, matching what the other architectures put in LR,
		// and pcdata lookups at retpc-1 stay inside the faulting
		// function.
		c.set_lr(uint64(pc) - 8)
	}

	// In case we are panicking from external C code.
	c.set_g1(uint64(uintptr(unsafe.Pointer(gp))))
	c.set_pc(uint64(abi.FuncPCABIInternal(sigpanic)))
}

// pushCallRec is a witness record of one asyncPreempt/sigpanic injection.
// The ring is dumped by fatalthrow on sparc64 so every crash arrives
// annotated with the injections that preceded it. Diagnostic aid; plain
// stores only, safe in signal context.
type pushCallRec struct {
	goid   uint64
	pc     uintptr
	npc    uintptr
	sp     uintptr
	ticks  int64
	target uintptr
}

var pushCallLog [64]pushCallRec
var pushCallIdx atomic.Uint32

func dumpPushCallLog() {
	n := pushCallIdx.Load()
	if n == 0 {
		return
	}
	print("injection witness log (", n, " total, newest last):\n")
	lo := uint32(0)
	if n > 16 {
		lo = n - 16
	}
	now := cputicks()
	for i := lo; i < n; i++ {
		r := &pushCallLog[i%uint32(len(pushCallLog))]
		print("  g", r.goid, " pc=", hex(r.pc))
		if fn := findfunc(r.pc); fn.valid() {
			print(" (", funcname(fn), ")")
		}
		print(" npc-pc=", int64(r.npc)-int64(r.pc), " sp=", hex(r.sp),
			" target=", hex(r.target), " ticksago=", now-r.ticks, "\n")
	}
}

func (c *sigctxt) pushCall(targetPC, resumePC uintptr) {
	// A signal that lands on the delay slot of a taken branch has
	// tnpc != tpc+4: the next instruction is the branch target, not
	// the successor. That state cannot be recreated from userspace -
	// asyncPreempt's final JMPL can only produce a sequential
	// (pc, pc+4) pair - so injecting here would resume with the
	// pending branch silently dropped, sending execution into the
	// wrong basic block with the other path's register state. Skip
	// the injection; preemption is best-effort and will be retried
	// at the next safe opportunity.
	if c.npc() != c.pc()+4 {
		return
	}
	// Freshness check on the kernel's window spill. copyWindow copies
	// [sp+0..127] assuming the kernel spilled the interrupted window
	// there at delivery. Verify it: the spilled %i6 slot must equal the
	// interrupted frame's biased frame pointer, which we can compute
	// from the context. A mismatch means the image is stale (the live
	// window never hit memory), and injecting would resume the
	// goroutine with old register state - so skip; preemption retries.
	if gp := getg(); gp != nil && gp.m != nil && gp.m.curg != nil {
		i := pushCallIdx.Add(1) - 1
		r := &pushCallLog[i%uint32(len(pushCallLog))]
		r.goid = gp.m.curg.goid
		r.pc = uintptr(c.pc())
		r.npc = uintptr(c.npc())
		r.sp = uintptr(c.sp())
		r.ticks = cputicks()
		r.target = targetPC
	}
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
