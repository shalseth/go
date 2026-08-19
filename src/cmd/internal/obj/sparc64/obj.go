// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sparc64

import (
	"cmd/internal/obj"
	"cmd/internal/sys"
	"internal/abi"
	"log"
)

var isUncondJump = map[obj.As]bool{
	obj.ACALL:     true,
	obj.ADUFFZERO: true,
	obj.ADUFFCOPY: true,
	obj.AJMP:      true,
	obj.ARET:      true,
	AFBA:          true,
	AJMPL:         true,
}

var isCondJump = map[obj.As]bool{
	ABN:    true,
	ABNE:   true,
	ABE:    true,
	ABG:    true,
	ABLE:   true,
	ABGE:   true,
	ABL:    true,
	ABGU:   true,
	ABLEU:  true,
	ABCC:   true,
	ABCS:   true,
	ABPOS:  true,
	ABNEG:  true,
	ABVC:   true,
	ABVS:   true,
	ABNW:   true,
	ABNEW:  true,
	ABEW:   true,
	ABGW:   true,
	ABLEW:  true,
	ABGEW:  true,
	ABLW:   true,
	ABGUW:  true,
	ABLEUW: true,
	ABCCW:  true,
	ABCSW:  true,
	ABPOSW: true,
	ABNEGW: true,
	ABVCW:  true,
	ABVSW:  true,
	ABND:   true,
	ABNED:  true,
	ABED:   true,
	ABGD:   true,
	ABLED:  true,
	ABGED:  true,
	ABLD:   true,
	ABGUD:  true,
	ABLEUD: true,
	ABCCD:  true,
	ABCSD:  true,
	ABPOSD: true,
	ABNEGD: true,
	ABVCD:  true,
	ABVSD:  true,
	ABRZ:   true,
	ABRLEZ: true,
	ABRLZ:  true,
	ABRNZ:  true,
	ABRGZ:  true,
	ABRGEZ: true,
	AFBN:   true,
	AFBU:   true,
	AFBG:   true,
	AFBUG:  true,
	AFBL:   true,
	AFBUL:  true,
	AFBLG:  true,
	AFBNE:  true,
	AFBE:   true,
	AFBUE:  true,
	AFBGE:  true,
	AFBUGE: true,
	AFBLE:  true,
	AFBULE: true,
	AFBO:   true,
}

var isJump = make(map[obj.As]bool)

func init() {
	for k := range isUncondJump {
		isJump[k] = true
	}
	for k := range isCondJump {
		isJump[k] = true
	}
}

// autoeditprog returns a new obj.Prog, with off(SP), off(FP), $off(SP),
// and $off(FP) replaced with new(RFP).
func autoeditprog(ctxt *obj.Link, cursym *obj.LSym, p *obj.Prog) *obj.Prog {
	r := new(obj.Prog)
	*r = *p
	// *r = *p copies the RestArgs slice header, so the extra operands
	// would still alias p's backing array. Clone it before editing.
	if len(p.RestArgs) > 0 {
		r.RestArgs = append([]obj.AddrPos(nil), p.RestArgs...)
	}
	r.From = *autoeditaddr(ctxt, cursym, &r.From)
	if f3 := r.GetFrom3(); f3 != nil {
		*f3 = *autoeditaddr(ctxt, cursym, f3)
	}
	r.To = *autoeditaddr(ctxt, cursym, &r.To)
	return r
}

// autoeditaddr returns a new obj.Addr, with off(SP), off(FP), $off(SP),
// and $off(FP) replaced with new(RFP).
func autoeditaddr(ctxt *obj.Link, cursym *obj.LSym, a *obj.Addr) *obj.Addr {
	if a == nil {
		return nil
	}
	if a.Type != obj.TYPE_MEM && a.Type != obj.TYPE_ADDR {
		return a
	}
	r := new(obj.Addr)
	*r = *a
	if r.Name == obj.NAME_PARAM {
		r.Reg = REG_RFP
		// cursym may be nil when called from aclass, which only needs
		// the operand class. RFP and RSP classify identically, so the
		// NOFRAME distinction does not affect that caller's result.
		if cursym != nil && cursym.Func().Text.From.Sym.NoFrame() {
			// NOFRAME functions live in caller's frame.
			r.Reg = REG_RSP
		}
		r.Offset += MinStackFrameSize + StackBias
		r.Name = obj.NAME_NONE
		return r
	}
	if r.Name == obj.NAME_AUTO {
		r.Reg = REG_RFP
		r.Offset += StackBias
		r.Name = obj.NAME_NONE
	}
	return r
}

// yfix rewrites references to Y registers (issued by compiler)
// to F and D registers.
func yfix(p *obj.Prog) {
	if REG_Y0 <= p.From.Reg && p.From.Reg <= REG_Y15 {
		if isInstDouble[p.As] || isSrcDouble[p.As] {
			p.From.Reg = REG_D0 + (p.From.Reg-REG_Y0)*2
		} else if isInstFloat[p.As] || isSrcFloat[p.As] {
			p.From.Reg = REG_F0 + (p.From.Reg-REG_Y0)*2
		}
	}
	if REG_Y0 <= p.Reg && p.Reg <= REG_Y15 {
		if isInstDouble[p.As] {
			p.Reg = REG_D0 + (p.Reg-REG_Y0)*2
		} else {
			p.Reg = REG_F0 + (p.Reg-REG_Y0)*2
		}
	}
	if p.GetFrom3() != nil && REG_Y0 <= p.GetFrom3().Reg && p.GetFrom3().Reg <= REG_Y15 {
		if isInstDouble[p.As] {
			p.GetFrom3().Reg = REG_D0 + (p.GetFrom3().Reg-REG_Y0)*2
		} else {
			p.GetFrom3().Reg = REG_F0 + (p.GetFrom3().Reg-REG_Y0)*2
		}
	}
	if REG_Y0 <= p.To.Reg && p.To.Reg <= REG_Y15 {
		if isInstDouble[p.As] || isDstDouble[p.As] {
			p.To.Reg = REG_D0 + (p.To.Reg-REG_Y0)*2
		} else if isInstFloat[p.As] || isDstFloat[p.As] {
			p.To.Reg = REG_F0 + (p.To.Reg-REG_Y0)*2
		}
	}
}

// biasfix rewrites referencing to BSP and BFP to RSP and RFP and
// adding the stack bias.
func biasfix(p *obj.Prog) {
	// Only match 2-operand instructions.
	if p.GetFrom3() != nil || p.Reg != 0 {
		return
	}
	switch p.As {
	case AMOVD:
		switch aclass(p.Ctxt, &p.From) {
		case ClassReg, ClassZero:
			switch {
			// MOVD	R, BSP	-> ADD	-$STACK_BIAS, R, RSP
			case aclass(p.Ctxt, &p.To) == ClassReg|ClassBias:
				p.As = AADD
				p.Reg = p.From.Reg
				if p.From.Type == obj.TYPE_CONST {
					p.Reg = REG_ZR
				}
				p.From.Reg = 0
				p.From.Offset = -StackBias
				p.From.Type = obj.TYPE_CONST
				p.From.Class = aclass(p.Ctxt, &p.From)
				p.To.Reg -= 256 // must match a.out.go:/REG_BSP
				p.To.Class = aclass(p.Ctxt, &p.To)
			}

		case ClassReg | ClassBias:
			// MOVD	BSP, R	-> ADD	$STACK_BIAS, RSP, R
			if aclass(p.Ctxt, &p.To) == ClassReg {
				p.Reg = p.From.Reg - 256 // must match a.out.go:/REG_BSP
				p.As = AADD
				p.From.Reg = 0
				p.From.Offset = StackBias
				p.From.Type = obj.TYPE_CONST
				p.From.Class = aclass(p.Ctxt, &p.From)
			}

		// MOVD	$off(BSP), R	-> MOVD	$(off+STACK_BIAS)(RSP), R
		case ClassRegConst13 | ClassBias, ClassRegConst | ClassBias:
			p.From.Reg -= 256 // must match a.out.go:/REG_BSP
			p.From.Offset += StackBias
			p.From.Class = aclass(p.Ctxt, &p.From)
		}

	case AADD, ASUB:
		// ADD	$const, BSP	-> ADD	$const, RSP
		if isAddrCompatible(p.Ctxt, &p.From, ClassConst) && aclass(p.Ctxt, &p.To) == ClassReg|ClassBias {
			p.To.Reg -= 256 // must match a.out.go:/REG_BSP
			p.To.Class = aclass(p.Ctxt, &p.To)
		}
	}
	switch p.As {
	case AMOVD, AMOVW, AMOVUW, AMOVH, AMOVUH, AMOVB, AMOVUB,
		AFMOVD, AFMOVS, ASTXFSR, ALDXFSR:
		switch aclass(p.Ctxt, &p.From) {
		case ClassZero, ClassReg, ClassFReg, ClassDReg:
			switch {
			// MOVD	R, off(BSP)	-> MOVD	R, (off+STACK_BIAS)(RSP)
			case aclass(p.Ctxt, &p.To)&ClassBias != 0 && isAddrCompatible(p.Ctxt, &p.To, ClassIndir):
				p.To.Offset += StackBias
				p.To.Reg -= 256 // must match a.out.go:/REG_BSP
				p.To.Class = aclass(p.Ctxt, &p.To)
			}

		// MOVD	off(BSP), R	-> MOVD	(off+STACK_BIAS)(RSP), R
		case ClassIndir0 | ClassBias, ClassIndir13 | ClassBias, ClassIndir | ClassBias:
			p.From.Reg -= 256 // must match a.out.go:/REG_BSP
			p.From.Offset += StackBias
			p.From.Class = aclass(p.Ctxt, &p.From)
		}
	}
}

func progedit(ctxt *obj.Link, p *obj.Prog, newprog obj.ProgAlloc) {
	// Rewrite constant moves to memory to go through an intermediary
	// register
	switch p.As {
	case AMOVD:
		if (p.From.Type == obj.TYPE_CONST || p.From.Type == obj.TYPE_ADDR) && (p.To.Type == obj.TYPE_MEM) {
			q := obj.Appendp(p, newprog)
			q.As = p.As
			q.To = p.To
			q.From.Type = obj.TYPE_REG
			q.From.Reg = REG_TMP
			q.From.Offset = 0

			p.To = obj.Addr{}
			p.To.Type = obj.TYPE_REG
			p.To.Reg = REG_TMP
			p.To.Offset = 0
		}

	case AFMOVS:
		if (p.From.Type == obj.TYPE_FCONST || p.From.Type == obj.TYPE_ADDR) && (p.To.Type == obj.TYPE_MEM) {
			q := obj.Appendp(p, newprog)
			q.As = p.As
			q.To = p.To
			q.From.Type = obj.TYPE_REG
			q.From.Reg = REG_FTMP
			q.From.Offset = 0

			p.To = obj.Addr{}
			p.To.Type = obj.TYPE_REG
			p.To.Reg = REG_FTMP
			p.To.Offset = 0
		}

	case AFMOVD:
		if (p.From.Type == obj.TYPE_FCONST || p.From.Type == obj.TYPE_ADDR) && (p.To.Type == obj.TYPE_MEM) {
			q := obj.Appendp(p, newprog)
			q.As = p.As
			q.To = p.To
			q.From.Type = obj.TYPE_REG
			q.From.Reg = REG_DTMP
			q.From.Offset = 0

			p.To = obj.Addr{}
			p.To.Type = obj.TYPE_REG
			p.To.Reg = REG_DTMP
			p.To.Offset = 0
		}
	}

	// Rewrite 64-bit integer constants and float constants to loads
	// from values stored in memory. The ctxt helpers both intern the
	// symbol and give it its data; a bare Lookup leaves an undefined
	// symbol behind for the linker to trip over.
	switch p.As {
	case AMOVD:
		if aclass(p.Ctxt, &p.From) == ClassConst {
			s := ctxt.Int64Sym(p.From.Offset)
			p.From.Type = obj.TYPE_MEM
			p.From.Sym = s
			p.From.Name = obj.NAME_EXTERN
			p.From.Offset = 0
		}

	case AFMOVS:
		if p.From.Type == obj.TYPE_FCONST {
			s := ctxt.Float32Sym(float32(p.From.Val.(float64)))
			p.From.Type = obj.TYPE_MEM
			p.From.Sym = s
			p.From.Name = obj.NAME_EXTERN
			p.From.Offset = 0
		}

	case AFMOVD:
		if p.From.Type == obj.TYPE_FCONST {
			s := ctxt.Float64Sym(p.From.Val.(float64))
			p.From.Type = obj.TYPE_MEM
			p.From.Sym = s
			p.From.Name = obj.NAME_EXTERN
			p.From.Offset = 0
		}
	}

	// TODO(aram): remove this when compiler can use F and
	// D registers directly.
	yfix(p)

	biasfix(p)
}

// isNOFRAME reports whether the TEXT prog p is marked NOFRAME. In 2016
// the textflags lived in a third operand; they are now attributes on the
// TEXT symbol itself.
func isNOFRAME(p *obj.Prog) bool {
	return p.From.Sym != nil && p.From.Sym.NoFrame()
}

// TODO(aram):
func preprocess(ctxt *obj.Link, cursym *obj.LSym, newprog obj.ProgAlloc) {
	cursym.Func().Text.Pc = 0
	cursym.Func().Args = cursym.Func().Text.To.Val.(int32)
	cursym.Func().Locals = int32(cursym.Func().Text.To.Offset)

	// Find leaf subroutines,
	// Strip NOPs.
	var q *obj.Prog
	var q1 *obj.Prog
	for p := cursym.Func().Text; p != nil; p = p.Link {
		switch {
		case p.As == obj.ATEXT:
			p.Mark |= LEAF

		case p.As == obj.ARET:
			break

		case p.As == obj.ANOP:
			q1 = p.Link
			q.Link = q1 /* q is non-nop */
			q1.Mark |= p.Mark
			continue

		case p.As == obj.ACALL || p.As == obj.ADUFFZERO || p.As == obj.ADUFFCOPY || p.As == AJMPL:
			// Only instructions that write the link register make a
			// function a non-leaf. Plain jumps — including the tail
			// call a compiled "RET target(SB)" becomes — do not.
			cursym.Func().Text.Mark &^= LEAF
			fallthrough

		case isUncondJump[p.As] || isCondJump[p.As]:
			q1 = p.To.Target()

			if q1 != nil {
				for q1.As == obj.ANOP {
					q1 = q1.Link
					p.To.SetTarget(q1)
				}
			}

			break
		}

		q = p
	}

	// A leaf with no locals needs no frame at all. Treat it as NOFRAME:
	// FP-relative arguments then resolve against RSP, and SP stays
	// unchanged — which the tail-call wrappers the compiler generates
	// (a bare "RET target(SB)") depend on. Pushing a frame there would
	// displace every argument the wrapper forwards.
	if cursym.Func().Locals == 0 && cursym.Func().Text.Mark&LEAF != 0 {
		cursym.Func().Text.From.Sym.Set(obj.AttrNoFrame, true)
	}

	for p := cursym.Func().Text; p != nil; p = p.Link {
		switch p.As {
		case obj.AGETCALLERPC:
			// The prologue moves the incoming return address from LR
			// (%o7, where CALL leaves it) into OLR (%i7), so a leaf
			// still has it in LR while everyone else reads OLR.
			//
			// %o7 holds the address of the CALL instruction itself.
			// Every other architecture's GetCallerPC returns the real
			// return address, and the runtime treats it as a resumable
			// PC (recovery jumps to _panic.startPC to resume a pending
			// Goexit), so add 8 to skip the call and its delay slot.
			p.As = AADD
			p.From.Type = obj.TYPE_CONST
			p.From.Offset = 8
			if cursym.Leaf() {
				p.Reg = REG_LR
			} else {
				p.Reg = REG_OLR
			}

		case obj.ATEXT:
			if cursym.Func().Text.Mark&LEAF != 0 {
				cursym.Set(obj.AttrLeaf, true)
			}
		}
	}

	for p := cursym.Func().Text; p != nil; p = p.Link {
		switch p.As {
		case obj.ATEXT:
			frameSize := cursym.Func().Locals
			if frameSize < 0 {
				ctxt.Diag("%v: negative frame size %d", p, frameSize)
			}
			// SPARC requires the stack pointer to stay 16-byte
			// aligned. The compiler already rounds its frames, but
			// hand-written assembly may declare any size, so pad here
			// rather than rejecting it - arm64, which has the same
			// requirement, does the same.
			if pad := frameSize % 16; pad != 0 {
				frameSize += 16 - pad
				cursym.Func().Locals = frameSize
				p.To.Offset = int64(frameSize)
			}
			if frameSize != 0 && isNOFRAME(p) {
				ctxt.Diag("%v: non-zero framesize for NOFRAME function", p)
			}

			if isNOFRAME(p) {
				// Without this NOP, DTrace changes the execution of the binary,
				// This should never happen, but this NOP seems to fix it.
				// Keep this NOP in here until we understand the DTrace behavior.
				p = obj.Appendp(p, newprog)
				p.As = ARNOP
				break
			}
			// Without this NOP, DTrace changes the execution of the binary,
			// This should never happen, but this NOP seems to fix it.
			// Keep this NOP in here until we understand the DTrace behavior.
			p = obj.Appendp(p, newprog)
			p.As = ARNOP

			if ctxt.Flag_maymorestack != "" && !cursym.Func().Text.From.Sym.NoSplit() {
				// Call the maymorestack hook once per call, above
				// the point morestack resumes at, so growing the
				// stack does not call it again.
				//
				// Publish the caller's anchors first: this is a real
				// call and must be unwindable. The frame is not open
				// yet, so open one of our own for it, and keep LR -
				// this function's return address, which the call
				// would otherwise clobber - along with the context
				// register.
				p = obj.Appendp(p, newprog)
				p.As = AMOVD
				p.From.Type = obj.TYPE_REG
				p.From.Reg = REG_RFP
				p.To.Type = obj.TYPE_MEM
				p.To.Reg = REG_RSP
				p.To.Offset = int64(112 + StackBias)

				p = obj.Appendp(p, newprog)
				p.As = AMOVD
				p.From.Type = obj.TYPE_REG
				p.From.Reg = REG_R31
				p.To.Type = obj.TYPE_MEM
				p.To.Reg = REG_RSP
				p.To.Offset = int64(120 + StackBias)

				p = cursym.Func().SpillRegisterArgs(p, newprog)

				const hookFrame = MinStackFrameSize + 16

				p = obj.Appendp(p, newprog)
				p.As = AADD
				p.From.Type = obj.TYPE_CONST
				p.From.Offset = -hookFrame
				p.To.Type = obj.TYPE_REG
				p.To.Reg = REG_RSP
				p.Spadj = hookFrame

				p = obj.Appendp(p, newprog)
				p.As = AMOVD
				p.From.Type = obj.TYPE_REG
				p.From.Reg = REG_LR
				p.To.Type = obj.TYPE_MEM
				p.To.Reg = REG_RSP
				p.To.Offset = MinStackFrameSize + StackBias

				p = obj.Appendp(p, newprog)
				p.As = AMOVD
				p.From.Type = obj.TYPE_REG
				p.From.Reg = REG_CTXT
				p.To.Type = obj.TYPE_MEM
				p.To.Reg = REG_RSP
				p.To.Offset = MinStackFrameSize + 8 + StackBias

				p = obj.Appendp(p, newprog)
				p.As = obj.ACALL
				p.To.Type = obj.TYPE_MEM
				p.To.Name = obj.NAME_EXTERN
				p.To.Sym = ctxt.LookupABI(ctxt.Flag_maymorestack, cursym.ABI())

				p = obj.Appendp(p, newprog)
				p.As = AMOVD
				p.From.Type = obj.TYPE_MEM
				p.From.Reg = REG_RSP
				p.From.Offset = MinStackFrameSize + 8 + StackBias
				p.To.Type = obj.TYPE_REG
				p.To.Reg = REG_CTXT

				p = obj.Appendp(p, newprog)
				p.As = AMOVD
				p.From.Type = obj.TYPE_MEM
				p.From.Reg = REG_RSP
				p.From.Offset = MinStackFrameSize + StackBias
				p.To.Type = obj.TYPE_REG
				p.To.Reg = REG_LR

				p = obj.Appendp(p, newprog)
				p.As = AADD
				p.From.Type = obj.TYPE_CONST
				p.From.Offset = hookFrame
				p.To.Type = obj.TYPE_REG
				p.To.Reg = REG_RSP
				p.Spadj = -hookFrame

				p = cursym.Func().UnspillRegisterArgs(p, newprog)
			}

			// MOVD RFP, (112+bias)(RSP)
			p = obj.Appendp(p, newprog)
			storeAnchors := p
			p.As = AMOVD
			p.From.Type = obj.TYPE_REG
			p.From.Reg = REG_RFP
			p.To.Type = obj.TYPE_MEM
			p.To.Reg = REG_RSP
			p.To.Offset = int64(112 + StackBias)

			// MOVD R31, (120+bias)(RSP)
			//
			// The caller's anchors go out before the stack split check
			// below: this slot doubles as the *caller's* return
			// address for the stack unwinder ([caller sp + 120]), and
			// a goroutine parked in the split check must already be
			// unwindable through its caller.
			p = obj.Appendp(p, newprog)
			p.As = AMOVD
			p.From.Type = obj.TYPE_REG
			p.From.Reg = REG_R31
			p.To.Type = obj.TYPE_MEM
			p.To.Reg = REG_RSP
			p.To.Offset = int64(120 + StackBias)

			if !cursym.Func().Text.From.Sym.NoSplit() {
				p = stacksplit(ctxt, cursym, newprog, p, int64(frameSize)+MinStackFrameSize, storeAnchors)
			}

			// The frame push and the anchor handoff below must not be
			// preempted asynchronously. Between them the live %i7 does
			// not describe the frame that %sp now names, and the
			// kernel spills the register window to [%sp+bias+112/120]
			// on signal delivery - exactly the slots the unwinder
			// reads as this frame's anchors. A preemption landing in
			// the middle publishes the wrong return address and the
			// traceback (and therefore the GC's stack scan) skips a
			// frame.
			p = ctxt.StartUnsafePoint(p, newprog)

			// ADD -(frame+128|176), RSP
			p = obj.Appendp(p, newprog)
			p.As = AADD
			p.From.Type = obj.TYPE_CONST
			p.From.Offset = -int64(frameSize + MinStackFrameSize)
			p.To.Type = obj.TYPE_REG
			p.To.Reg = REG_RSP
			p.Spadj = frameSize + int32(MinStackFrameSize)

			// SUB -(frame+128|176), RSP, RFP
			p = obj.Appendp(p, newprog)
			p.As = ASUB
			p.From.Type = obj.TYPE_CONST
			p.From.Offset = -int64(frameSize + MinStackFrameSize)
			p.Reg = REG_RSP
			p.To.Type = obj.TYPE_REG
			p.To.Reg = REG_RFP

			// MOVD LR, R31
			p = obj.Appendp(p, newprog)
			p.As = AMOVD
			p.From.Type = obj.TYPE_REG
			p.From.Reg = REG_LR
			p.To.Type = obj.TYPE_REG
			p.To.Reg = REG_R31

			// Publish this frame's own anchors at [own sp+112/120],
			// the same slots and values the kernel spills there on a
			// trap. Without this a frame is described only by its
			// callees: a framed function whose calls so far have all
			// been to leaf functions has an unwritten return-address
			// slot, and any unwind that crosses it - the GC stack
			// scan, save(), a panic's defer walk - reads whatever a
			// previous occupant of that stack memory left behind and
			// decodes the rest of the stack from fossils. Storing the
			// pair at frame birth closes that hole for good.

			// MOVD RFP, (112+bias)(RSP)
			p = obj.Appendp(p, newprog)
			p.As = AMOVD
			p.From.Type = obj.TYPE_REG
			p.From.Reg = REG_RFP
			p.To.Type = obj.TYPE_MEM
			p.To.Reg = REG_RSP
			p.To.Offset = int64(112 + StackBias)

			// MOVD R31, (120+bias)(RSP)
			p = obj.Appendp(p, newprog)
			p.As = AMOVD
			p.From.Type = obj.TYPE_REG
			p.From.Reg = REG_R31
			p.To.Type = obj.TYPE_MEM
			p.To.Reg = REG_RSP
			p.To.Offset = int64(120 + StackBias)

			p = ctxt.EndUnsafePoint(p, newprog, -1)

		case obj.ARET:
			retSym := p.To.Sym
			retReg := int16(0)
			if p.To.Type == obj.TYPE_REG && p.To.Sym == nil {
				// RET (Rn): an indirect tail call, used by
				// compiler-generated method wrappers.
				retReg = p.To.Reg
			}
			if isNOFRAME(cursym.Func().Text) {
				if retSym != nil {
					// RET target(SB): a tail call out of a
					// frameless function is a plain jump.
					p.As = obj.AJMP
				} else if retReg != 0 {
					p.As = AJMPL
					p.From.Type = obj.TYPE_REG
					p.From.Reg = retReg
					p.To.Type = obj.TYPE_REG
					p.To.Reg = REG_ZR
				}
				break
			}

			// The epilogue must not touch any scratch register:
			// assembly functions like gcWriteBarrier promise to
			// preserve every general-purpose register, and this
			// epilogue runs after their own restores. Everything
			// derives from RFP and the frame slots.

			// Same hazard as the prologue: once the anchors below have
			// been reloaded from this frame's slots, the live %i6/%i7
			// belong to the caller while %sp still names this frame,
			// so a window spill would publish the caller's return
			// address as this frame's. Keep the whole restore
			// sequence non-preemptible.
			// This RET prog becomes the PCDATA that opens the region;
			// the real return is appended at the end.
			q1 = p
			q1.As = obj.APCDATA
			q1.From = obj.Addr{}
			q1.From.SetConst(abi.PCDATA_UnsafePoint)
			q1.To = obj.Addr{}
			q1.To.SetConst(abi.UnsafePointUnsafe)

			// MOVD R31, LR
			q1 = obj.Appendp(q1, newprog)
			q1.As = AMOVD
			q1.From.Type = obj.TYPE_REG
			q1.From.Reg = REG_R31
			q1.To.Type = obj.TYPE_REG
			q1.To.Reg = REG_LR

			// MOVD (120+StackBias)(RFP), R31
			q1 = obj.Appendp(q1, newprog)
			q1.As = AMOVD
			q1.From.Type = obj.TYPE_MEM
			q1.From.Reg = REG_RFP
			q1.From.Offset = 120 + StackBias
			q1.To.Type = obj.TYPE_REG
			q1.To.Reg = REG_R31

			// MOVD (112+StackBias)(RFP), RFP - restore the caller's
			// frame anchor BEFORE raising SP. The kernel spills %i6
			// to [sp+bias+112] on every involuntary context switch:
			// if SP were raised first, there would be a window where
			// [sp+112] is the caller's anchor slot while %i6 still
			// holds the dying frame's fp, and a context switch there
			// would overwrite the anchor with fp itself (the classic
			// "RFP == RSP after return" corruption). With this order
			// the transient spill clobbers only the dying frame's
			// own slot, and once SP is raised, %i6 and [sp+112]
			// already agree.
			q1 = obj.Appendp(q1, newprog)
			q1.As = AMOVD
			q1.From.Type = obj.TYPE_MEM
			q1.From.Reg = REG_RFP
			q1.From.Offset = 112 + StackBias
			q1.To.Type = obj.TYPE_REG
			q1.To.Reg = REG_RFP

			// ADD $framesize, RSP (pop the frame by constant - RFP
			// no longer holds the entry sp)
			q1 = obj.Appendp(q1, newprog)
			q1.As = AADD
			q1.From.Type = obj.TYPE_CONST
			q1.From.Offset = int64(cursym.Func().Locals) + MinStackFrameSize
			q1.To.Type = obj.TYPE_REG
			q1.To.Reg = REG_RSP
			q1.Spadj = -(cursym.Func().Locals + int32(MinStackFrameSize))

			q1 = ctxt.EndUnsafePoint(q1, newprog, -1)

			// The return itself.
			p = obj.Appendp(q1, newprog)
			p.As = obj.ARET

			// The epilogue's SP restore carries Spadj=-frame; the
			// return itself compensates so instructions after it
			// (other body code with the frame live) keep the right
			// pcsp value.
			p.Spadj = cursym.Func().Locals + int32(MinStackFrameSize)

			if retSym != nil {
				// RET target(SB): after the epilogue, jump to the
				// target instead of returning. Appendp gave the
				// final instruction an empty To, so the original
				// target must be put back explicitly.
				p.As = obj.AJMP
				p.To.Type = obj.TYPE_MEM
				p.To.Name = obj.NAME_EXTERN
				p.To.Sym = retSym
			} else if retReg != 0 {
				// RET (Rn): after the epilogue, jump through the
				// register. The epilogue only touches R1, RFP,
				// OLR and RSP, so the target register survives.
				p.As = AJMPL
				p.From.Type = obj.TYPE_REG
				p.From.Reg = retReg
				p.To.Type = obj.TYPE_REG
				p.To.Reg = REG_ZR
			}
		}
	}

	// Track stack-pointer adjustments in hand-written assembly so the
	// pcsp tables cover asm frames too. The compiler-generated
	// prologue and epilogue set Spadj explicitly above.
	for p := cursym.Func().Text; p != nil; p = p.Link {
		// BSP is the same physical register as RSP; it only tells the
		// operand classifier that memory references through it carry the
		// stack bias. Assembly that adjusts the stack pointer may spell it
		// either way, and asyncPreempt spells its prologue with BSP and its
		// epilogue with RSP. Missing the BSP form left that frame out of the
		// pcsp table entirely, so the unwinder placed the interrupted frame
		// 512 bytes too low and died with "traceback did not unwind
		// completely".
		to, reg := p.To.Reg, p.Reg
		if to == REG_BSP {
			to = REG_RSP
		}
		if reg == REG_BSP {
			reg = REG_RSP
		}
		if p.Spadj != 0 || p.From.Type != obj.TYPE_CONST || p.To.Type != obj.TYPE_REG || to != REG_RSP {
			continue
		}
		if reg != 0 && reg != REG_RSP {
			continue
		}
		switch p.As {
		case AADD:
			p.Spadj = int32(-p.From.Offset)
		case ASUB:
			p.Spadj = int32(p.From.Offset)
		}
	}

	// Mark functions that write the stack pointer in a way the spdelta
	// table cannot describe -- the stack switches in hand-written assembly,
	// which load a whole new SP rather than adjust it by a constant
	// (systemstack, mcall, gogo). Two things depend on this flag, and both
	// were silently wrong without it:
	//
	// The unwinder refuses to walk past such a frame, because past a stack
	// switch it may not even be on the stack it thinks it is. Missing the
	// flag, it walked on and read a stack address as a return PC, which is
	// the "traceback did not unwind completely" this port kept dying on.
	//
	// More seriously, isAsyncSafePoint refuses to preempt inside one. Missing
	// the flag, the runtime async-preempted stack-switching assembly midway
	// through the switch; proc.go has an explicit assertion against exactly
	// that. That is a plausible source of the rare, load-dependent memory
	// corruption seen when building large programs natively.
	//
	// Both spellings of the stack pointer have to be recognised: BSP is the
	// same register as RSP and only marks operands as carrying the bias.
	for p := cursym.Func().Text; p != nil; p = p.Link {
		if p.To.Type != obj.TYPE_REG || p.Spadj != 0 {
			continue
		}
		if p.To.Reg != REG_RSP && p.To.Reg != REG_BSP {
			continue
		}
		f := cursym.Func()
		if f.FuncFlag&abi.FuncFlagSPWrite == 0 {
			f.FuncFlag |= abi.FuncFlagSPWrite
			if ctxt.Debugvlog {
				ctxt.Logf("auto-SPWRITE: %s %v\n", cursym.Name, p)
			}
		}
	}

	// Schedule delay-slots. Only RNOPs for now. A RESTORE directly
	// after a jump is hand-written asm deliberately placing the window
	// restore in the delay slot (the `ret; restore` idiom) — leave it.
	//
	// The slot must belong to the jump, and Appendp gives the RNOP the
	// jump's position. A call's return address points past the delay
	// slot, so the runtime's universal "a return address minus one is
	// inside the CALL" rule — unwinder.symPC, and through it
	// CallersFrames, testing.T.Helper attribution and every profile and
	// trace stack — resolves to the slot rather than to the call. A slot
	// carrying a foreign position reports the call as coming from
	// whichever line that instruction belongs to.
	//
	// Never adopt an RNOP that is already there instead of appending
	// one. Ginsnop emits an RNOP for ssa.OpInlMark, and the inline tree
	// records that instruction's PC as the position of the parent frame
	// at an inlined call site. Such a mark frequently follows a call,
	// and reusing it as the delay slot would force one instruction to
	// carry two different positions: the caller would be reported at the
	// inlined call site, or the inlined frame at the caller's previous
	// call, depending on which position won.
	for p := cursym.Func().Text; p != nil; p = p.Link {
		if !isJump[p.As] {
			continue
		}
		if p.Link != nil && p.Link.As == ARESTORE {
			continue
		}
		p = obj.Appendp(p, newprog)
		p.As = ARNOP
	}

	// A call to a function that never returns (gopanic, panicwrap) may
	// be the last code in a function. Its return address - the call
	// plus 8, past the delay slot - would then be the entry of the NEXT
	// function, and the traceback machinery would resolve the frame to
	// the wrong symbol with a zero pcsp. Append an UNDEF so the return
	// address stays inside this function.
	isLinkingCall := func(p *obj.Prog) bool {
		return p != nil && (p.As == obj.ACALL || p.As == obj.ADUFFZERO || p.As == obj.ADUFFCOPY ||
			(p.As == AJMPL && p.To.Reg != REG_ZR))
	}
	var prev *obj.Prog
	for p := cursym.Func().Text; p != nil; prev, p = p, p.Link {
		if p.Link != nil {
			continue
		}
		if isLinkingCall(p) || (p.As == ARNOP && isLinkingCall(prev)) {
			q := obj.Appendp(p, newprog)
			q.As = obj.AUNDEF
		}
	}

	// Mark unsafe points for asynchronous preemption. The asyncPreempt
	// return path clobbers TMP, so any instruction explicitly keeping
	// a value there must not be preempted; multi-instruction
	// expansions that go through TMP internally have no side effects
	// before their final instruction and can simply restart.
	isUnsafePoint := func(p *obj.Prog) bool {
		if p.From.Reg == REG_TMP || p.Reg == REG_TMP || p.To.Reg == REG_TMP ||
			p.From.Index == REG_TMP || p.To.Index == REG_TMP {
			return true
		}
		// Multi-instruction expansions that go through TMP implicitly
		// (ClobberTMP) and control-flow expansions must not be
		// preempted mid-sequence either; treating them as unsafe
		// rather than restartable sidesteps mid-sequence resume
		// entirely.
		//
		// oplook keys the table on the operand classes, and this runs
		// before preprocess caches them on the real progs, so classify
		// the copy here: without this every lookup fails and nothing
		// is ever marked unsafe.
		q := autoeditprog(ctxt, cursym, p)
		q.From.Class = aclass(ctxt, &q.From)
		if f3 := q.GetFrom3(); f3 != nil {
			f3.Class = aclass(ctxt, f3)
		}
		q.To.Class = aclass(ctxt, &q.To)
		o, err := oplook(q)
		return err == nil && (o.OpInfo&ClobberTMP != 0 || o.size > 4)
	}
	isRestartable := func(p *obj.Prog) bool {
		return false
	}
	obj.MarkUnsafePoints(ctxt, cursym.Func().Text, newprog, isUnsafePoint, isRestartable)

	// For future use by oplook and friends.
	for p := cursym.Func().Text; p != nil; p = p.Link {
		p.From.Class = aclass(ctxt, &p.From)
		if p.GetFrom3() != nil {
			p.GetFrom3().Class = aclass(ctxt, p.GetFrom3())
		}
		p.To.Class = aclass(ctxt, &p.To)
	}
}

func relinv(a obj.As) obj.As {
	switch a {
	case obj.AJMP:
		return ABN
	case ABN:
		return obj.AJMP
	case ABE:
		return ABNE
	case ABNE:
		return ABE
	case ABG:
		return ABLE
	case ABLE:
		return ABG
	case ABGE:
		return ABL
	case ABL:
		return ABGE
	case ABGU:
		return ABLEU
	case ABLEU:
		return ABGU
	case ABCC:
		return ABCS
	case ABCS:
		return ABCC
	case ABPOS:
		return ABNEG
	case ABNEG:
		return ABPOS
	case ABVC:
		return ABVS
	case ABVS:
		return ABVC
	case ABNW:
		return obj.AJMP
	case ABEW:
		return ABNEW
	case ABNEW:
		return ABEW
	case ABGW:
		return ABLEW
	case ABLEW:
		return ABGW
	case ABGEW:
		return ABLW
	case ABLW:
		return ABGEW
	case ABGUW:
		return ABLEUW
	case ABLEUW:
		return ABGUW
	case ABCCW:
		return ABCSW
	case ABCSW:
		return ABCCW
	case ABPOSW:
		return ABNEGW
	case ABNEGW:
		return ABPOSW
	case ABVCW:
		return ABVSW
	case ABVSW:
		return ABVCW
	case ABND:
		return obj.AJMP
	case ABED:
		return ABNED
	case ABNED:
		return ABED
	case ABGD:
		return ABLED
	case ABLED:
		return ABGD
	case ABGED:
		return ABLD
	case ABLD:
		return ABGED
	case ABGUD:
		return ABLEUD
	case ABLEUD:
		return ABGUD
	case ABCCD:
		return ABCSD
	case ABCSD:
		return ABCCD
	case ABPOSD:
		return ABNEGD
	case ABNEGD:
		return ABPOSD
	case ABVCD:
		return ABVSD
	case ABVSD:
		return ABVCD
	case AFBN:
		return AFBA
	case AFBA:
		return AFBN
	case AFBE:
		return AFBLG
	case AFBLG:
		return AFBE
	case AFBG:
		return AFBLE
	case AFBLE:
		return AFBG
	case AFBGE:
		return AFBL
	case AFBL:
		return AFBGE
	}

	log.Fatalf("unknown relation: %s", a.String())
	return 0
}

var unaryDst = map[obj.As]bool{
	obj.ACALL: true,
	obj.AJMP:  true,
	AWORD:     true,
	ADWORD:    true,
	ABNW:      true,
	ABNEW:     true,
	ABEW:      true,
	ABGW:      true,
	ABLEW:     true,
	ABGEW:     true,
	ABLW:      true,
	ABGUW:     true,
	ABLEUW:    true,
	ABCCW:     true,
	ABCSW:     true,
	ABPOSW:    true,
	ABNEGW:    true,
	ABVCW:     true,
	ABVSW:     true,
	ABND:      true,
	ABNED:     true,
	ABED:      true,
	ABGD:      true,
	ABLED:     true,
	ABGED:     true,
	ABLD:      true,
	ABGUD:     true,
	ABLEUD:    true,
	ABCCD:     true,
	ABCSD:     true,
	ABPOSD:    true,
	ABNEGD:    true,
	ABVCD:     true,
	ABVSD:     true,
}

var Linksparc64 = obj.LinkArch{
	Arch:           sys.ArchSPARC64,
	Init:           buildop,
	Preprocess:     preprocess,
	Assemble:       span,
	Progedit:       progedit,
	UnaryDst:       unaryDst,
	DWARFRegisters: SPARCDWARFRegisters,
}

// buildop is called once per Link to prepare any derived tables. The
// optab is a plain map built at package init, so there is nothing to
// sort or index here yet; the hook exists because LinkArch requires it
// and because the SSA backend will need per-Link setup later.
func buildop(ctxt *obj.Link) {}

// stacksplit inserts the stack-bound check and morestack call at p,
// which sits between the TEXT prog and the frame-establishing part of
// the prologue. totalframe is the full frame the prologue is about to
// push (locals plus the fixed frame), so it is always larger than
// abi.StackSmall and the small-stack fast path of other ports does not
// apply. Returns the last prog of the inserted sequence.
func stacksplit(ctxt *obj.Link, cursym *obj.LSym, newprog obj.ProgAlloc, p *obj.Prog, totalframe int64, resume *obj.Prog) *obj.Prog {
	// MOVD g_stackguard(g), RT1
	p = obj.Appendp(p, newprog)
	// (the guard load; the split resumes at the anchor stores instead)
	p.As = AMOVD
	p.From.Type = obj.TYPE_MEM
	p.From.Reg = REG_G
	p.From.Offset = 2 * int64(ctxt.Arch.PtrSize) // G.stackguard0
	if cursym.CFunc() {
		p.From.Offset = 3 * int64(ctxt.Arch.PtrSize) // G.stackguard1
	}
	p.To.Type = obj.TYPE_REG
	p.To.Reg = REG_RT1

	// Mark the stack bound check and morestack call async nonpreemptible.
	// If we get preempted here, when resumed the preemption request is
	// cleared, but we'll still call morestack, which will double the stack
	// unnecessarily. See issue #35470.
	p = ctxt.StartUnsafePoint(p, newprog)

	// ADD $StackBias, RSP, RT2 (the unbiased stack pointer; the BSP
	// pseudo-register is only rewritten for parsed assembly, not for
	// progs appended here)
	p = obj.Appendp(p, newprog)
	p.As = AADD
	p.From.Type = obj.TYPE_CONST
	p.From.Offset = StackBias
	p.Reg = REG_RSP
	p.To.Type = obj.TYPE_REG
	p.To.Reg = REG_RT2

	// MOVD $(totalframe-StackSmall), R3
	// The offset may not fit in a 13-bit immediate, so always go
	// through a register; R3 is dead until the morestack handoff.
	offset := totalframe - abi.StackSmall
	p = obj.Appendp(p, newprog)
	p.As = AMOVD
	p.From.Type = obj.TYPE_CONST
	p.From.Offset = offset
	p.To.Type = obj.TYPE_REG
	p.To.Reg = REG_R3

	var wrapCheck *obj.Prog
	if totalframe > abi.StackBig {
		// The runtime guarantees SP > StackBig on entry, but with a
		// frame this large SP-offset may still underflow. Grow the
		// stack when SP < offset rather than comparing garbage.
		//	CMP R3, RT2
		//	BCSD morestack
		p = obj.Appendp(p, newprog)
		p.As = ACMP
		p.From.Type = obj.TYPE_REG
		p.From.Reg = REG_R3
		p.Reg = REG_RT2
		p.To.Type = obj.TYPE_NONE

		p = obj.Appendp(p, newprog)
		wrapCheck = p
		p.As = ABCSD
		p.To.Type = obj.TYPE_BRANCH
	}

	// SUB R3, RT2, RT2 (SP - (totalframe - StackSmall))
	p = obj.Appendp(p, newprog)
	p.As = ASUB
	p.From.Type = obj.TYPE_REG
	p.From.Reg = REG_R3
	p.Reg = REG_RT2
	p.To.Type = obj.TYPE_REG
	p.To.Reg = REG_RT2

	// CMP RT1, RT2 (flags = adjusted SP - stackguard)
	p = obj.Appendp(p, newprog)
	p.As = ACMP
	p.From.Type = obj.TYPE_REG
	p.From.Reg = REG_RT1
	p.Reg = REG_RT2
	p.To.Type = obj.TYPE_NONE

	// BCCD ok (unsigned >=: enough stack)
	p = obj.Appendp(p, newprog)
	enough := p
	p.As = ABCCD
	p.To.Type = obj.TYPE_BRANCH

	// MOVD LR, R3: the morestack contract wants the prologue's LR
	// (this function's raw return address) in R3.
	p = obj.Appendp(p, newprog)
	p.As = AMOVD
	p.From.Type = obj.TYPE_REG
	p.From.Reg = REG_LR
	p.To.Type = obj.TYPE_REG
	p.To.Reg = REG_R3
	if wrapCheck != nil {
		wrapCheck.To.SetTarget(p)
	}

	p = ctxt.EmitEntryStackMap(cursym, p, newprog)

	// CALL runtime.morestack(SB)
	p = obj.Appendp(p, newprog)
	p.As = obj.ACALL
	p.To.Type = obj.TYPE_MEM
	p.To.Name = obj.NAME_EXTERN
	if cursym.CFunc() {
		p.To.Sym = ctxt.Lookup("runtime.morestackc")
	} else if !cursym.Func().Text.From.Sym.NeedCtxt() {
		p.To.Sym = ctxt.Lookup("runtime.morestack_noctxt")
	} else {
		p.To.Sym = ctxt.Lookup("runtime.morestack")
	}

	p = ctxt.EndUnsafePoint(p, newprog, -1)

	// JMP back to the guard load: after the stack has grown the check
	// runs again (and passes).
	p = obj.Appendp(p, newprog)
	p.As = obj.AJMP
	p.To.Type = obj.TYPE_BRANCH
	// Resume at the anchor stores, not at the guard load: morestack
	// resumes here with the frame's RFP/OLR reloaded from the goroutine
	// context, and the slots this frame wrote before growing hold
	// pre-copy values that stack copying cannot fix up (the frame has
	// no size yet, so it owns no anchor slot to adjust).
	p.To.SetTarget(resume)

	// Branch-over target. A real nop: this assembler has no zero-width
	// placeholder handling.
	p = obj.Appendp(p, newprog)
	p.As = ARNOP
	enough.To.SetTarget(p)

	return p
}
