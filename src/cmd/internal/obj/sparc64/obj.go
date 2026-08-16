// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sparc64

import (
	"cmd/internal/obj"
	"cmd/internal/sys"
	"fmt"
	"log"
	"math"
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
		AFMOVD, AFMOVS:
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

	// Rewrite 64-bit integer constants and float constants
	// to values stored in memory.
	switch p.As {
	case AMOVD:
		if aclass(p.Ctxt, &p.From) == ClassConst {
			literal := fmt.Sprintf("$i64.%016x", p.From.Offset)
			s := ctxt.Lookup(literal)
			s.Size = 8
			p.From.Type = obj.TYPE_MEM
			p.From.Sym = s
			p.From.Name = obj.NAME_EXTERN
			p.From.Offset = 0
		}

	case AFMOVS:
		if p.From.Type == obj.TYPE_FCONST {
			f32 := float32(p.From.Val.(float64))
			i32 := math.Float32bits(f32)
			literal := fmt.Sprintf("$f32.%08x", uint32(i32))
			s := ctxt.Lookup(literal)
			s.Size = 4
			p.From.Type = obj.TYPE_MEM
			p.From.Sym = s
			p.From.Name = obj.NAME_EXTERN
			p.From.Offset = 0
		}

	case AFMOVD:
		if p.From.Type == obj.TYPE_FCONST {
			i64 := math.Float64bits(p.From.Val.(float64))
			literal := fmt.Sprintf("$f64.%016x", uint64(i64))
			s := ctxt.Lookup(literal)
			s.Size = 8
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

		case isUncondJump[p.As]:
			cursym.Func().Text.Mark &^= LEAF
			fallthrough

		case isCondJump[p.As]:
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

	for p := cursym.Func().Text; p != nil; p = p.Link {
		switch p.As {
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
			if frameSize%16 != 0 {
				ctxt.Diag("%v: unaligned frame size %d - must be 0 mod 16", p, frameSize)
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

			// ADD -(frame+128|176), RSP
			p = obj.Appendp(p, newprog)
			p.As = AADD
			p.From.Type = obj.TYPE_CONST
			p.From.Offset = -int64(frameSize + MinStackFrameSize)
			p.To.Type = obj.TYPE_REG
			p.To.Reg = REG_RSP

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

		case obj.ARET:
			if isNOFRAME(cursym.Func().Text) {
				break
			}

			// MOVD RFP, R1
			q1 = p
			p = obj.Appendp(p, newprog)
			p.As = obj.ARET
			q1.As = AMOVD
			q1.From.Type = obj.TYPE_REG
			q1.From.Reg = REG_RFP
			q1.To.Type = obj.TYPE_REG
			q1.To.Reg = REG_R1

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

			// MOVD (112+StackBias)(RFP), RFP
			q1 = obj.Appendp(q1, newprog)
			q1.As = AMOVD
			q1.From.Type = obj.TYPE_MEM
			q1.From.Reg = REG_RFP
			q1.From.Offset = 112 + StackBias
			q1.To.Type = obj.TYPE_REG
			q1.To.Reg = REG_RFP

			// MOVD R1, RSP
			q1 = obj.Appendp(q1, newprog)
			q1.As = AMOVD
			q1.From.Type = obj.TYPE_REG
			q1.From.Reg = REG_R1
			q1.To.Type = obj.TYPE_REG
			q1.To.Reg = REG_RSP
		}
	}

	// Schedule delay-slots. Only RNOPs for now.
	for p := cursym.Func().Text; p != nil; p = p.Link {
		if !isJump[p.As] {
			continue
		}
		if p.Link != nil && p.Link.As == ARNOP {
			continue
		}
		p = obj.Appendp(p, newprog)
		p.As = ARNOP
	}

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
	Arch:       sys.ArchSPARC64,
	Init:       buildop,
	Preprocess: preprocess,
	Assemble:   span,
	Progedit:   progedit,
	UnaryDst:       unaryDst,
	DWARFRegisters: SPARCDWARFRegisters,
}

// buildop is called once per Link to prepare any derived tables. The
// optab is a plain map built at package init, so there is nothing to
// sort or index here yet; the hook exists because LinkArch requires it
// and because the SSA backend will need per-Link setup later.
func buildop(ctxt *obj.Link) {}
