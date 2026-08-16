// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	"internal/goarch"
	"unsafe"
)

// Linux/sparc64 does not hand the signal handler a ucontext_t. The
// kernel passes the address of the struct sigcontext embedded in the
// signal frame, and glibc's ucontext_t does not describe it: reading
// mc_gregs[MC_PC] through glibc's headers yields zero.
//
// The layout below was established by faulting on a known address under
// a SA_SIGINFO handler on an UltraSPARC T4 and dumping the words the
// handler received:
//
//	  0..127   siginfo, copied into the frame
//	128..255   u_regs[0..15], that is %g0..%g7 then %o0..%o7
//	     256   tstate
//	     264   tpc     -- the faulting PC
//	     272   tnpc
//	     280   y, then fprs
//
// %o6 read back as an odd value because it carries the SPARC V9 stack
// bias; sp() removes it so the runtime sees an ordinary stack address.

type sigcontext struct {
	sigc_info [128]byte
	u_regs    [16]uint64 // %g0..%g7, %o0..%o7
	tstate    uint64
	tpc       uint64
	tnpc      uint64
	y         uint32
	fprs      uint32
}

// Indices into u_regs.
const (
	_UREG_G0 = 0
	_UREG_G1 = 1
	_UREG_G7 = 7
	_UREG_O0 = 8
	_UREG_O6 = 14 // stack pointer, biased
	_UREG_O7 = 15 // link register
)

type sigctxt struct {
	info *siginfo
	ctxt unsafe.Pointer
}

//go:nosplit
//go:nowritebarrierrec
func (c *sigctxt) regs() *sigcontext { return (*sigcontext)(c.ctxt) }

func (c *sigctxt) g1() uint64 { return c.regs().u_regs[_UREG_G1] }
func (c *sigctxt) g7() uint64 { return c.regs().u_regs[_UREG_G7] }
func (c *sigctxt) o0() uint64 { return c.regs().u_regs[_UREG_O0] }
func (c *sigctxt) lr() uint64 { return c.regs().u_regs[_UREG_O7] }

// sp returns the unbiased stack pointer. The register holds %sp as the
// hardware sees it, which is the real stack address minus the bias.
func (c *sigctxt) sp() uint64 { return c.regs().u_regs[_UREG_O6] + stackBias }

func (c *sigctxt) pc() uint64  { return c.regs().tpc }
func (c *sigctxt) npc() uint64 { return c.regs().tnpc }

func (c *sigctxt) sigcode() uint64 { return uint64(c.info.si_code) }
func (c *sigctxt) sigaddr() uint64 { return c.info.si_addr }

// set_pc also sets tnpc. SPARC executes with a program counter and a
// next program counter, so moving the PC without moving nPC would
// resume at an unrelated instruction after the first one retires.
func (c *sigctxt) set_pc(x uint64) {
	c.regs().tpc = x
	c.regs().tnpc = x + 4
}

func (c *sigctxt) set_lr(x uint64) { c.regs().u_regs[_UREG_O7] = x }
func (c *sigctxt) set_sp(x uint64) { c.regs().u_regs[_UREG_O6] = x - stackBias }
func (c *sigctxt) set_g1(x uint64) { c.regs().u_regs[_UREG_G1] = x }

func (c *sigctxt) set_sigcode(x uint64) { c.info.si_code = int32(x) }

func (c *sigctxt) set_sigaddr(x uint64) {
	*(*uintptr)(add(unsafe.Pointer(c.info), 2*goarch.PtrSize)) = uintptr(x)
}
