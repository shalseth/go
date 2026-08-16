// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sparc64

import (
	"cmd/compile/internal/base"
	"cmd/compile/internal/ir"
	"cmd/compile/internal/logopt"
	"cmd/compile/internal/objw"
	"cmd/compile/internal/ssa"
	"cmd/compile/internal/ssa/block"
	"cmd/compile/internal/ssagen"
	"cmd/compile/internal/types"
	"internal/abi"
	"cmd/internal/obj"
	"cmd/internal/obj/sparc64"
	"math"
)

// ssaMarkMoves marks any MOVDconst ops that need to avoid clobbering
// the condition codes. SPARC's plain MOVD does not write the condition
// codes, so nothing needs marking.
func ssaMarkMoves(s *ssagen.State, b *ssa.Block) {}

// isFPReg reports whether r is one of the virtual float registers the
// backend allocates. The assembler's yfix pass later rewrites these to
// the F or D register the chosen instruction needs.
func isFPReg(r int16) bool {
	return sparc64.REG_Y0 <= r && r <= sparc64.REG_Y15
}

// regMoveOp returns the instruction that moves size bytes between two
// registers. A move that crosses between the integer and float files
// cannot be a plain MOVD on SPARC; it needs the VIS3 file-crossing
// moves, which is why this port requires a T3 or later.
func regMoveOp(src, dst int16, size int64) obj.As {
	srcFP, dstFP := isFPReg(src), isFPReg(dst)
	switch {
	case srcFP && dstFP:
		if size == 4 {
			return sparc64.AFMOVS
		}
		return sparc64.AFMOVD
	case srcFP && !dstFP:
		if size == 4 {
			return sparc64.AMOVSTOUW
		}
		return sparc64.AMOVDTOX
	case !srcFP && dstFP:
		if size == 4 {
			return sparc64.AMOVWTOS
		}
		return sparc64.AMOVXTOD
	}
	return sparc64.AMOVD
}

// loadByType returns the load instruction for the given type.
func loadByType(t *types.Type, r int16) obj.As {
	if t.IsFloat() || isFPReg(r) {
		if t.Size() == 4 {
			return sparc64.AFMOVS
		}
		return sparc64.AFMOVD
	}
	switch t.Size() {
	case 1:
		if t.IsSigned() {
			return sparc64.AMOVB
		}
		return sparc64.AMOVUB
	case 2:
		if t.IsSigned() {
			return sparc64.AMOVH
		}
		return sparc64.AMOVUH
	case 4:
		if t.IsSigned() {
			return sparc64.AMOVW
		}
		return sparc64.AMOVUW
	case 8:
		return sparc64.AMOVD
	}
	panic("bad load type")
}

// storeByType returns the store instruction for the given type.
func storeByType(t *types.Type, r int16) obj.As {
	if t.IsFloat() || isFPReg(r) {
		if t.Size() == 4 {
			return sparc64.AFMOVS
		}
		return sparc64.AFMOVD
	}
	switch t.Size() {
	case 1:
		return sparc64.AMOVB
	case 2:
		return sparc64.AMOVH
	case 4:
		return sparc64.AMOVW
	case 8:
		return sparc64.AMOVD
	}
	panic("bad store type")
}

// condMove maps a flag-reading op to the SPARC conditional move that
// implements it. SPARC has no set-on-condition instruction, so a
// boolean is produced by zeroing a register and then conditionally
// moving 1 into it.
func condMove(op ssa.Op) obj.As {
	switch op {
	case ssa.OpSPARC64Equal:
		return sparc64.AMOVE
	case ssa.OpSPARC64NotEqual:
		return sparc64.AMOVNE
	case ssa.OpSPARC64LessThan:
		return sparc64.AMOVL
	case ssa.OpSPARC64LessEqual:
		return sparc64.AMOVLE
	case ssa.OpSPARC64GreaterThan:
		return sparc64.AMOVG
	case ssa.OpSPARC64GreaterEqual:
		return sparc64.AMOVGE
	case ssa.OpSPARC64LessThanU:
		return sparc64.AMOVCS
	case ssa.OpSPARC64LessEqualU:
		return sparc64.AMOVLEU
	case ssa.OpSPARC64GreaterThanU:
		return sparc64.AMOVGU
	case ssa.OpSPARC64GreaterEqualU:
		return sparc64.AMOVCC
	}
	panic("bad conditional move op")
}

// boundsRegs lists the registers a PanicBounds operand may live in, in
// the order the PCData encoding numbers them. SPARC's allocatable
// registers are not a contiguous run, so unlike the other backends the
// index cannot be computed as reg-REG_R1 and this table is the mapping.
//
// The runtime's bounds-panic handler must use the same order when it
// recovers the operands from the signal context.
var boundsRegs = [16]int16{
	sparc64.REG_R1, sparc64.REG_R2, sparc64.REG_R3, sparc64.REG_R4,
	sparc64.REG_R5, sparc64.REG_R8, sparc64.REG_R9, sparc64.REG_R10,
	sparc64.REG_R11, sparc64.REG_R12, sparc64.REG_R13, sparc64.REG_R15,
	sparc64.REG_R16, sparc64.REG_R17, sparc64.REG_R18, sparc64.REG_R19,
}

// boundsRegIndex returns the PCData index for register r.
func boundsRegIndex(r int16) int {
	for i, x := range boundsRegs {
		if x == r {
			return i
		}
	}
	panic("register not in the PanicBounds set")
}

// panicBounds emits the PCData entry and call for a failed bounds check.
func panicBounds(s *ssagen.State, v *ssa.Value) {
	code, signed := ssa.BoundsKind(v.AuxInt).Code()
	xIsReg, yIsReg := false, false
	xVal, yVal := 0, 0

	// loadConst puts c in a register from the bounds set and returns its
	// index, avoiding the index already used by the other operand.
	loadConst := func(c int64, avoid int) int {
		idx := 0
		if idx == avoid {
			idx = 1
		}
		p := s.Prog(sparc64.AMOVD)
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = c
		p.To.Type = obj.TYPE_REG
		p.To.Reg = boundsRegs[idx]
		return idx
	}

	switch v.Op {
	case ssa.OpSPARC64LoweredPanicBoundsRR:
		xIsReg, yIsReg = true, true
		xVal = boundsRegIndex(v.Args[0].Reg())
		yVal = boundsRegIndex(v.Args[1].Reg())
	case ssa.OpSPARC64LoweredPanicBoundsRC:
		xIsReg = true
		xVal = boundsRegIndex(v.Args[0].Reg())
		c := v.Aux.(ssa.PanicBoundsC).C
		if c >= 0 && c <= abi.BoundsMaxConst {
			yVal = int(c)
		} else {
			yIsReg = true
			yVal = loadConst(c, xVal)
		}
	case ssa.OpSPARC64LoweredPanicBoundsCR:
		yIsReg = true
		yVal = boundsRegIndex(v.Args[0].Reg())
		c := v.Aux.(ssa.PanicBoundsC).C
		if c >= 0 && c <= abi.BoundsMaxConst {
			xVal = int(c)
		} else {
			xIsReg = true
			xVal = loadConst(c, yVal)
		}
	case ssa.OpSPARC64LoweredPanicBoundsCC:
		cx := v.Aux.(ssa.PanicBoundsCC).Cx
		if cx >= 0 && cx <= abi.BoundsMaxConst {
			xVal = int(cx)
		} else {
			xIsReg = true
			xVal = loadConst(cx, -1)
		}
		cy := v.Aux.(ssa.PanicBoundsCC).Cy
		if cy >= 0 && cy <= abi.BoundsMaxConst {
			yVal = int(cy)
		} else {
			yIsReg = true
			yVal = loadConst(cy, xVal)
		}
	}

	c := abi.BoundsEncode(code, signed, xIsReg, yIsReg, xVal, yVal)
	p := s.Prog(obj.APCDATA)
	p.From.SetConst(abi.PCDATA_PanicBounds)
	p.To.SetConst(int64(c))
	p = s.Prog(obj.ACALL)
	p.To.Type = obj.TYPE_MEM
	p.To.Name = obj.NAME_EXTERN
	p.To.Sym = ir.Syms.PanicBounds
}

// moveSizeAndOp picks the widest store that respects the given
// alignment. SPARC traps on unaligned access, so this must not
// over-promise.
func moveSizeAndOp(align int64) (int64, obj.As) {
	switch {
	case align%8 == 0:
		return 8, sparc64.AMOVD
	case align%4 == 0:
		return 4, sparc64.AMOVW
	case align%2 == 0:
		return 2, sparc64.AMOVH
	}
	return 1, sparc64.AMOVB
}

// fcondMove maps a float flag-reading op to the SPARC conditional move
// that implements it. These are the FMOVcc opcodes, whose cond field
// uses the float encoding rather than the integer one.
func fcondMove(op ssa.Op) obj.As {
	switch op {
	case ssa.OpSPARC64FEqual:
		return sparc64.AMOVFE
	case ssa.OpSPARC64FNotEqual:
		return sparc64.AMOVFNE
	case ssa.OpSPARC64FLessThan:
		return sparc64.AMOVFL
	case ssa.OpSPARC64FLessEqual:
		return sparc64.AMOVFLE
	case ssa.OpSPARC64FGreaterThan:
		return sparc64.AMOVFG
	case ssa.OpSPARC64FGreaterEqual:
		return sparc64.AMOVFGE
	}
	panic("bad float conditional move op")
}

func ssaGenValue(s *ssagen.State, v *ssa.Value) {
	switch v.Op {
	case ssa.OpCopy:
		if v.Type.IsMemory() {
			return
		}
		x := v.Args[0].Reg()
		y := v.Reg()
		if x == y {
			return
		}
		p := s.Prog(regMoveOp(x, y, v.Type.Size()))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = x
		p.To.Type = obj.TYPE_REG
		p.To.Reg = y

	case ssa.OpLoadReg:
		if v.Type.IsFlags() {
			v.Fatalf("load flags not implemented: %v", v.LongString())
			return
		}
		p := s.Prog(loadByType(v.Type, v.Reg()))
		ssagen.AddrAuto(&p.From, v.Args[0])
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()

	case ssa.OpStoreReg:
		if v.Type.IsFlags() {
			v.Fatalf("store flags not implemented: %v", v.LongString())
			return
		}
		p := s.Prog(storeByType(v.Type, v.Args[0].Reg()))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg()
		ssagen.AddrAuto(&p.To, v)

	case ssa.OpArgIntReg, ssa.OpArgFloatReg:
		ssagen.CheckArgReg(v)

	case ssa.OpPhi:
		ssagen.CheckLoweredPhi(v)

	case ssa.OpInitMem, ssa.OpVarDef, ssa.OpVarLive, ssa.OpKeepAlive,
		ssa.OpSelect0, ssa.OpSelect1, ssa.OpGetG, ssa.OpSP, ssa.OpSB:
		// nothing to do

	case ssa.OpSPARC64ADD, ssa.OpSPARC64SUB, ssa.OpSPARC64MULD,
		ssa.OpSPARC64SDIVD, ssa.OpSPARC64UDIVD, ssa.OpSPARC64UMULXHI,
		ssa.OpSPARC64AND, ssa.OpSPARC64OR, ssa.OpSPARC64XOR,
		ssa.OpSPARC64ANDN, ssa.OpSPARC64ORN, ssa.OpSPARC64XNOR,
		ssa.OpSPARC64SLLD, ssa.OpSPARC64SRLD, ssa.OpSPARC64SRAD,
		ssa.OpSPARC64SLLW, ssa.OpSPARC64SRLW, ssa.OpSPARC64SRAW,
		ssa.OpSPARC64FADDS, ssa.OpSPARC64FADDD,
		ssa.OpSPARC64FSUBS, ssa.OpSPARC64FSUBD,
		ssa.OpSPARC64FMULS, ssa.OpSPARC64FMULD,
		ssa.OpSPARC64FDIVS, ssa.OpSPARC64FDIVD:
		// SPARC three-operand form: op rs1, rs2, rd. Go's assembler
		// takes rs2 in From, rs1 in Reg and rd in To.
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg()
		p.Reg = v.Args[0].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()

	case ssa.OpSPARC64ADDconst, ssa.OpSPARC64SUBconst,
		ssa.OpSPARC64ANDconst, ssa.OpSPARC64ORconst, ssa.OpSPARC64XORconst,
		ssa.OpSPARC64SLLDconst, ssa.OpSPARC64SRLDconst, ssa.OpSPARC64SRADconst:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.Reg = v.Args[0].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()

	case ssa.OpSPARC64NEG,
		ssa.OpSPARC64MOVB, ssa.OpSPARC64MOVUB,
		ssa.OpSPARC64MOVH, ssa.OpSPARC64MOVUH,
		ssa.OpSPARC64MOVW, ssa.OpSPARC64MOVUW, ssa.OpSPARC64MOVD,
		ssa.OpSPARC64FNEGS, ssa.OpSPARC64FNEGD,
		ssa.OpSPARC64FABSS, ssa.OpSPARC64FABSD,
		ssa.OpSPARC64FSQRTS, ssa.OpSPARC64FSQRTD,
		ssa.OpSPARC64FSTOD, ssa.OpSPARC64FDTOS,
		ssa.OpSPARC64FSTOX, ssa.OpSPARC64FDTOX,
		ssa.OpSPARC64FXTOS, ssa.OpSPARC64FXTOD,
		ssa.OpSPARC64FMOVDgp, ssa.OpSPARC64FMOVDfp:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[0].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()

	case ssa.OpSPARC64MULDU:
		// UMULXHI rs1, rs2, hi ; MULD rs1, rs2, lo
		p := s.Prog(sparc64.AUMULXHI)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg()
		p.Reg = v.Args[0].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg0()
		p = s.Prog(sparc64.AMULD)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg()
		p.Reg = v.Args[0].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg1()

	case ssa.OpSPARC64MOVDconst:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()

	case ssa.OpSPARC64CMP, ssa.OpSPARC64FCMPS, ssa.OpSPARC64FCMPD:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg()
		p.Reg = v.Args[0].Reg()

	case ssa.OpSPARC64CMPconst:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = v.AuxInt
		p.Reg = v.Args[0].Reg()

	case ssa.OpSPARC64MOVDaddr:
		p := s.Prog(sparc64.AMOVD)
		p.From.Type = obj.TYPE_ADDR
		p.From.Reg = v.Args[0].Reg()
		ssagen.AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()

	case ssa.OpSPARC64MOVBload, ssa.OpSPARC64MOVUBload,
		ssa.OpSPARC64MOVHload, ssa.OpSPARC64MOVUHload,
		ssa.OpSPARC64MOVWload, ssa.OpSPARC64MOVUWload,
		ssa.OpSPARC64MOVDload,
		ssa.OpSPARC64FMOVSload, ssa.OpSPARC64FMOVDload:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg()
		ssagen.AddAux(&p.From, v)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()

	case ssa.OpSPARC64MOVBstore, ssa.OpSPARC64MOVHstore,
		ssa.OpSPARC64MOVWstore, ssa.OpSPARC64MOVDstore,
		ssa.OpSPARC64FMOVSstore, ssa.OpSPARC64FMOVDstore:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg()
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg()
		ssagen.AddAux(&p.To, v)

	case ssa.OpSPARC64FEqual, ssa.OpSPARC64FNotEqual,
		ssa.OpSPARC64FLessThan, ssa.OpSPARC64FLessEqual,
		ssa.OpSPARC64FGreaterThan, ssa.OpSPARC64FGreaterEqual:
		// MOVD $0, rd ; MOVF<cc> FCC0, $1, rd
		r := v.Reg()
		p := s.Prog(sparc64.AMOVD)
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = 0
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
		p = s.Prog(fcondMove(v.Op))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = sparc64.REG_FCC0
		p.AddRestSourceConst(1)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r

	case ssa.OpSPARC64Equal, ssa.OpSPARC64NotEqual,
		ssa.OpSPARC64LessThan, ssa.OpSPARC64LessEqual,
		ssa.OpSPARC64GreaterThan, ssa.OpSPARC64GreaterEqual,
		ssa.OpSPARC64LessThanU, ssa.OpSPARC64LessEqualU,
		ssa.OpSPARC64GreaterThanU, ssa.OpSPARC64GreaterEqualU:
		// MOVD $0, rd ; MOV<cc> XCC, $1, rd
		r := v.Reg()
		p := s.Prog(sparc64.AMOVD)
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = 0
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r
		p = s.Prog(condMove(v.Op))
		p.From.Type = obj.TYPE_REG
		p.From.Reg = sparc64.REG_XCC
		p.AddRestSourceConst(1)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = r

	case ssa.OpSPARC64FMOVSconst, ssa.OpSPARC64FMOVDconst:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_FCONST
		p.From.Val = math.Float64frombits(uint64(v.AuxInt))
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()

	case ssa.OpSPARC64LoweredZero:
		// SUB	$sz, R1
		// MOVD	ZR, sz(R1)
		// ADD	$sz, R1
		// BNED	R1, Rarg1, -2(PC)
		// arg1 is the address of the last element to zero.
		sz, mov := moveSizeAndOp(v.AuxInt)
		p := s.Prog(sparc64.ASUB)
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = sz
		p.Reg = sparc64.REG_R1
		p.To.Type = obj.TYPE_REG
		p.To.Reg = sparc64.REG_R1
		p2 := s.Prog(mov)
		p2.From.Type = obj.TYPE_REG
		p2.From.Reg = sparc64.REG_ZR
		p2.To.Type = obj.TYPE_MEM
		p2.To.Reg = sparc64.REG_R1
		p2.To.Offset = sz
		p3 := s.Prog(sparc64.AADD)
		p3.From.Type = obj.TYPE_CONST
		p3.From.Offset = sz
		p3.Reg = sparc64.REG_R1
		p3.To.Type = obj.TYPE_REG
		p3.To.Reg = sparc64.REG_R1
		p4 := s.Prog(sparc64.ACMP)
		p4.From.Type = obj.TYPE_REG
		p4.From.Reg = v.Args[1].Reg()
		p4.Reg = sparc64.REG_R1
		p5 := s.Prog(sparc64.ABNED)
		p5.To.Type = obj.TYPE_BRANCH
		p5.To.SetTarget(p2)

	case ssa.OpSPARC64LoweredMove:
		// SUB	$sz, R1
		// MOVD	sz(R1), TMP
		// MOVD	TMP, (R2)
		// ADD	$sz, R1
		// ADD	$sz, R2
		// BNED	R1, Rarg2, -4(PC)
		sz, mov := moveSizeAndOp(v.AuxInt)
		p := s.Prog(sparc64.ASUB)
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = sz
		p.Reg = sparc64.REG_R1
		p.To.Type = obj.TYPE_REG
		p.To.Reg = sparc64.REG_R1
		p2 := s.Prog(mov)
		p2.From.Type = obj.TYPE_MEM
		p2.From.Reg = sparc64.REG_R1
		p2.From.Offset = sz
		p2.To.Type = obj.TYPE_REG
		p2.To.Reg = sparc64.REGTMP
		p3 := s.Prog(mov)
		p3.From.Type = obj.TYPE_REG
		p3.From.Reg = sparc64.REGTMP
		p3.To.Type = obj.TYPE_MEM
		p3.To.Reg = sparc64.REG_R2
		p4 := s.Prog(sparc64.AADD)
		p4.From.Type = obj.TYPE_CONST
		p4.From.Offset = sz
		p4.Reg = sparc64.REG_R1
		p4.To.Type = obj.TYPE_REG
		p4.To.Reg = sparc64.REG_R1
		p5 := s.Prog(sparc64.AADD)
		p5.From.Type = obj.TYPE_CONST
		p5.From.Offset = sz
		p5.Reg = sparc64.REG_R2
		p5.To.Type = obj.TYPE_REG
		p5.To.Reg = sparc64.REG_R2
		p6 := s.Prog(sparc64.ACMP)
		p6.From.Type = obj.TYPE_REG
		p6.From.Reg = v.Args[2].Reg()
		p6.Reg = sparc64.REG_R1
		p7 := s.Prog(sparc64.ABNED)
		p7.To.Type = obj.TYPE_BRANCH
		p7.To.SetTarget(p2)

	case ssa.OpSPARC64LoweredPanicBoundsRR, ssa.OpSPARC64LoweredPanicBoundsRC,
		ssa.OpSPARC64LoweredPanicBoundsCR, ssa.OpSPARC64LoweredPanicBoundsCC:
		panicBounds(s, v)

	case ssa.OpSPARC64LoweredWB:
		p := s.Prog(obj.ACALL)
		p.To.Type = obj.TYPE_MEM
		p.To.Name = obj.NAME_EXTERN
		// AuxInt encodes how many buffer entries we need.
		p.To.Sym = ir.Syms.GCWriteBarrier[v.AuxInt-1]

	case ssa.OpSPARC64CALLstatic, ssa.OpSPARC64CALLclosure, ssa.OpSPARC64CALLinter:
		s.Call(v)

	case ssa.OpSPARC64CALLtail:
		s.TailCall(v)

	case ssa.OpSPARC64LoweredGetClosurePtr:
		ssagen.CheckLoweredGetClosurePtr(v)

	case ssa.OpSPARC64LoweredGetCallerSP:
		p := s.Prog(sparc64.AMOVD)
		p.From.Type = obj.TYPE_ADDR
		p.From.Offset = -base.Ctxt.Arch.FixedFrameSize
		p.From.Name = obj.NAME_PARAM
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()

	case ssa.OpSPARC64LoweredGetCallerPC:
		p := s.Prog(obj.AGETCALLERPC)
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg()

	case ssa.OpSPARC64LoweredNilCheck:
		// Issue a load through the pointer; a nil pointer faults.
		p := s.Prog(sparc64.AMOVUB)
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = sparc64.REGTMP
		if logopt.Enabled() {
			logopt.LogOpt(v.Pos, "nilcheck", "genssa", v.Block.Func.Name)
		}
		if base.Debug.Nil != 0 && v.Pos.Line() > 1 {
			base.WarnfAt(v.Pos, "generated nil check")
		}

	case ssa.OpSPARC64LoweredPubBarrier:
		p := s.Prog(v.Op.Asm())
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = 0xf // #Sync

	case ssa.OpClobber, ssa.OpClobberReg:
		// nothing to do

	default:
		v.Fatalf("genValue not implemented: %s", v.LongString())
	}
}

// blockJump maps each block kind to its branch and the inverse branch,
// so the generator can fall through to whichever successor follows.
// The D suffix selects the XCC (64-bit) condition codes.
var blockJump = map[block.BlockKind]struct {
	asm, invasm obj.As
}{
	block.BlockSPARC64EQ:  {sparc64.ABED, sparc64.ABNED},
	block.BlockSPARC64NE:  {sparc64.ABNED, sparc64.ABED},
	block.BlockSPARC64LT:  {sparc64.ABLD, sparc64.ABGED},
	block.BlockSPARC64GE:  {sparc64.ABGED, sparc64.ABLD},
	block.BlockSPARC64LE:  {sparc64.ABLED, sparc64.ABGD},
	block.BlockSPARC64GT:  {sparc64.ABGD, sparc64.ABLED},
	block.BlockSPARC64ULT: {sparc64.ABCSD, sparc64.ABCCD},
	block.BlockSPARC64UGE: {sparc64.ABCCD, sparc64.ABCSD},
	block.BlockSPARC64ULE: {sparc64.ABLEUD, sparc64.ABGUD},
	block.BlockSPARC64UGT: {sparc64.ABGUD, sparc64.ABLEUD},
	block.BlockSPARC64FEQ: {sparc64.AFBE, sparc64.AFBNE},
	block.BlockSPARC64FNE: {sparc64.AFBNE, sparc64.AFBE},
	block.BlockSPARC64FLT: {sparc64.AFBL, sparc64.AFBGE},
	block.BlockSPARC64FGE: {sparc64.AFBGE, sparc64.AFBL},
	block.BlockSPARC64FLE: {sparc64.AFBLE, sparc64.AFBG},
	block.BlockSPARC64FGT: {sparc64.AFBG, sparc64.AFBLE},
}

func ssaGenBlock(s *ssagen.State, b, next *ssa.Block) {
	switch b.Kind {
	case block.BlockPlain, block.BlockDefer:
		if b.Succs[0].Block() != next {
			p := s.Prog(obj.AJMP)
			p.To.Type = obj.TYPE_BRANCH
			s.Branches = append(s.Branches, ssagen.Branch{P: p, B: b.Succs[0].Block()})
		}

	case block.BlockExit, block.BlockRetJmp:
		// nothing to do

	case block.BlockRet:
		s.Prog(obj.ARET)

	default:
		jmp, ok := blockJump[b.Kind]
		if !ok {
			b.Fatalf("branch not implemented: %s", b.LongString())
		}
		// The condition codes were set by the controlling CMP or FCMP.
		switch next {
		case b.Succs[0].Block():
			s.Br(jmp.invasm, b.Succs[1].Block())
		case b.Succs[1].Block():
			s.Br(jmp.asm, b.Succs[0].Block())
		default:
			if b.Likely != ssa.BranchUnlikely {
				s.Br(jmp.asm, b.Succs[0].Block())
				s.Br(obj.AJMP, b.Succs[1].Block())
			} else {
				s.Br(jmp.invasm, b.Succs[1].Block())
				s.Br(obj.AJMP, b.Succs[0].Block())
			}
		}
	}
}

// zerorange zeroes the stack range [off, off+cnt) at function entry.
func zerorange(pp *objw.Progs, p *obj.Prog, off, cnt int64, _ *uint32) *obj.Prog {
	if cnt == 0 {
		return p
	}
	for i := int64(0); i < cnt; i += 8 {
		p = pp.Append(p, sparc64.AMOVD, obj.TYPE_REG, sparc64.REG_ZR, 0,
			obj.TYPE_MEM, sparc64.REGSP, sparc64.StackBias+off+i)
	}
	return p
}

// ginsnop emits a no-op, used to pad call sites.
func ginsnop(pp *objw.Progs) *obj.Prog {
	return pp.Prog(sparc64.ARNOP)
}
