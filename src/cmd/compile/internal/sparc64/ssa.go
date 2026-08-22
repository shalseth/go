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
	"cmd/internal/obj"
	"cmd/internal/obj/sparc64"
	"internal/abi"
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

// SPARC's memory model is TSO: loads are ordered against loads and
// stores, and stores against stores, without any barrier. Only
// store-then-load can be reordered, so a sequentially consistent store
// needs #StoreLoad and a read-modify-write needs a full barrier before
// anything after it may be observed.
const (
	membarStoreLoad = 2  // #StoreLoad
	membarFull      = 15 // #LoadLoad|#StoreLoad|#LoadStore|#StoreStore
)

func membar(s *ssagen.State, mask int64) {
	p := s.Prog(sparc64.AMEMBAR)
	p.From.Type = obj.TYPE_CONST
	p.From.Offset = mask
}

// casLoop emits the tail of a compare-and-swap retry loop: compare the
// value CAS returned against the value we expected to find, and branch
// back to top if they differ. CASW and CASD leave the old memory
// contents in their destination register, so equality means the swap
// happened.
func casLoop(s *ssagen.State, got, want int16, top *obj.Prog) {
	p := s.Prog(sparc64.ASUBCC)
	p.From.Type = obj.TYPE_REG
	p.From.Reg = got
	p.Reg = want
	p.To.Type = obj.TYPE_REG
	p.To.Reg = sparc64.REG_ZR
	b := s.Prog(sparc64.ABNED)
	b.To.Type = obj.TYPE_BRANCH
	b.To.SetTarget(top)
	membar(s, membarFull)
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
	sparc64.REG_R11, sparc64.REG_R12, sparc64.REG_R13, sparc64.REG_R16,
	sparc64.REG_R17, sparc64.REG_R18, sparc64.REG_R19, sparc64.REG_R20,
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

// zeroMoveAux unpacks a LoweredZero/LoweredMove AuxInt into the byte
// count, the access width, and the move instruction for that width.
// The lowering rules pack these as size<<4 | width, with the width
// chosen from the type's alignment; SPARC traps on unaligned access,
// so the width must not over-promise.
func zeroMoveAux(v *ssa.Value) (sz, chunk int64, mov obj.As) {
	sz = v.AuxInt >> 4
	chunk = v.AuxInt & 0xf
	switch chunk {
	case 8:
		mov = sparc64.AMOVD
	case 4:
		mov = sparc64.AMOVW
	case 2:
		mov = sparc64.AMOVH
	case 1:
		mov = sparc64.AMOVB
	default:
		v.Fatalf("bad LoweredZero/LoweredMove width %d", chunk)
	}
	if sz <= 0 || sz%chunk != 0 {
		// The loop would run past the end. Go type sizes are multiples
		// of their alignment, so this cannot happen for rule-generated
		// ops; fail loudly if it somehow does.
		v.Fatalf("bad LoweredZero/LoweredMove size %d for width %d", sz, chunk)
	}
	return sz, chunk, mov
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

	case ssa.OpSPARC64NEG, ssa.OpSPARC64POPC,
		ssa.OpSPARC64MOVB, ssa.OpSPARC64MOVUB,
		ssa.OpSPARC64MOVH, ssa.OpSPARC64MOVUH,
		ssa.OpSPARC64MOVW, ssa.OpSPARC64MOVUW, ssa.OpSPARC64MOVD,
		ssa.OpSPARC64FNEGS, ssa.OpSPARC64FNEGD,
		ssa.OpSPARC64FABSS, ssa.OpSPARC64FABSD,
		ssa.OpSPARC64FSQRTS, ssa.OpSPARC64FSQRTD,
		ssa.OpSPARC64FSTOD, ssa.OpSPARC64FDTOS,
		ssa.OpSPARC64FSTOX, ssa.OpSPARC64FDTOX,
		ssa.OpSPARC64FSTOI, ssa.OpSPARC64FDTOI,
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

	case ssa.OpSPARC64ADDCARRY, ssa.OpSPARC64SUBBORROW:
		// (sum, carry) = arg0 +- arg1 +- arg2, with arg2 and carry in {0, 1}:
		//
		//	ADDCC/SUBCC arg1, arg0, sum   // xcc.c = c1
		//	ADDXC       ZR, ZR, carry     // carry = c1
		//	ADDCC/SUBCC arg2, sum, sum    // xcc.c = c2
		//	ADDXC       carry, ZR, carry  // carry = c1 + c2
		//
		// At most one of c1 and c2 can be set, so the last add cannot
		// itself carry. resultNotInArgs keeps sum clear of arg2, which
		// is still live for the third instruction.
		op := sparc64.AADDCC
		if v.Op == ssa.OpSPARC64SUBBORROW {
			op = sparc64.ASUBCC
		}
		sum, carry := v.Reg0(), v.Reg1()
		p := s.Prog(op)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg()
		p.Reg = v.Args[0].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = sum
		p = s.Prog(sparc64.AADDXC)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = sparc64.REG_ZR
		p.Reg = sparc64.REG_ZR
		p.To.Type = obj.TYPE_REG
		p.To.Reg = carry
		p = s.Prog(op)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[2].Reg()
		p.Reg = sum
		p.To.Type = obj.TYPE_REG
		p.To.Reg = sum
		p = s.Prog(sparc64.AADDXC)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = sparc64.REG_ZR
		p.Reg = carry
		p.To.Type = obj.TYPE_REG
		p.To.Reg = carry

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
		if p.From.Reg == sparc64.REG_RSP && p.From.Sym == nil {
			// Raw stack-pointer arithmetic, as produced by the
			// OffPtr(SP) lowering for outgoing call arguments. %sp
			// carries the SPARC V9 bias, so the real address is
			// %sp + StackBias + offset. Named slots do not take this
			// path: the assembler's autoedit pass adds the bias when
			// it resolves NAME_AUTO and NAME_PARAM.
			p.From.Offset += sparc64.StackBias
		}
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

	case ssa.OpSPARC64LoweredAtomicLoad8, ssa.OpSPARC64LoweredAtomicLoad32, ssa.OpSPARC64LoweredAtomicLoad64:
		// A plain load. Under TSO it is already an acquire.
		ld := sparc64.AMOVD
		switch v.Op {
		case ssa.OpSPARC64LoweredAtomicLoad8:
			ld = sparc64.AMOVUB
		case ssa.OpSPARC64LoweredAtomicLoad32:
			ld = sparc64.AMOVUW
		}
		p := s.Prog(ld)
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = v.Args[0].Reg()
		p.To.Type = obj.TYPE_REG
		p.To.Reg = v.Reg0()

	case ssa.OpSPARC64LoweredAtomicStore8, ssa.OpSPARC64LoweredAtomicStore32, ssa.OpSPARC64LoweredAtomicStore64:
		// A plain store is already a release; #StoreLoad is what keeps
		// a later load from passing it.
		st := sparc64.AMOVD
		switch v.Op {
		case ssa.OpSPARC64LoweredAtomicStore8:
			st = sparc64.AMOVB
		case ssa.OpSPARC64LoweredAtomicStore32:
			st = sparc64.AMOVW
		}
		p := s.Prog(st)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = v.Args[1].Reg()
		p.To.Type = obj.TYPE_MEM
		p.To.Reg = v.Args[0].Reg()
		membar(s, membarStoreLoad)

	case ssa.OpSPARC64LoweredAtomicExchange32, ssa.OpSPARC64LoweredAtomicExchange64:
		// again:
		//	MOV(UW|D)	(ptr), out
		//	MOVD		val, TMP
		//	CAS(W|D)	(ptr), out, TMP
		//	SUBCC		TMP, out, ZR
		//	BNED		again
		//	MEMBAR		$15
		ld, cas := sparc64.AMOVD, sparc64.ACASD
		if v.Op == ssa.OpSPARC64LoweredAtomicExchange32 {
			ld, cas = sparc64.AMOVUW, sparc64.ACASW
		}
		ptr, val, out := v.Args[0].Reg(), v.Args[1].Reg(), v.Reg0()
		top := s.Prog(ld)
		top.From.Type = obj.TYPE_MEM
		top.From.Reg = ptr
		top.To.Type = obj.TYPE_REG
		top.To.Reg = out
		p1 := s.Prog(sparc64.AMOVD)
		p1.From.Type = obj.TYPE_REG
		p1.From.Reg = val
		p1.To.Type = obj.TYPE_REG
		p1.To.Reg = sparc64.REG_TMP
		p2 := s.Prog(cas)
		p2.From.Type = obj.TYPE_MEM
		p2.From.Reg = ptr
		p2.Reg = out
		p2.To.Type = obj.TYPE_REG
		p2.To.Reg = sparc64.REG_TMP
		casLoop(s, sparc64.REG_TMP, out, top)

	case ssa.OpSPARC64LoweredAtomicAdd32, ssa.OpSPARC64LoweredAtomicAdd64:
		// again:
		//	MOV(UW|D)	(ptr), TMP
		//	ADD		delta, TMP, out
		//	MOVD		out, TMP2
		//	CAS(W|D)	(ptr), TMP, TMP2
		//	SUBCC		TMP2, TMP, ZR
		//	BNED		again
		//	MEMBAR		$15
		// The result is the new value, which is what Xadd returns.
		ld, cas := sparc64.AMOVD, sparc64.ACASD
		if v.Op == ssa.OpSPARC64LoweredAtomicAdd32 {
			ld, cas = sparc64.AMOVUW, sparc64.ACASW
		}
		ptr, delta, out := v.Args[0].Reg(), v.Args[1].Reg(), v.Reg0()
		top := s.Prog(ld)
		top.From.Type = obj.TYPE_MEM
		top.From.Reg = ptr
		top.To.Type = obj.TYPE_REG
		top.To.Reg = sparc64.REG_TMP
		p1 := s.Prog(sparc64.AADD)
		p1.From.Type = obj.TYPE_REG
		p1.From.Reg = delta
		p1.Reg = sparc64.REG_TMP
		p1.To.Type = obj.TYPE_REG
		p1.To.Reg = out
		p2 := s.Prog(sparc64.AMOVD)
		p2.From.Type = obj.TYPE_REG
		p2.From.Reg = out
		p2.To.Type = obj.TYPE_REG
		p2.To.Reg = sparc64.REG_TMP2
		p3 := s.Prog(cas)
		p3.From.Type = obj.TYPE_MEM
		p3.From.Reg = ptr
		p3.Reg = sparc64.REG_TMP
		p3.To.Type = obj.TYPE_REG
		p3.To.Reg = sparc64.REG_TMP2
		casLoop(s, sparc64.REG_TMP2, sparc64.REG_TMP, top)

	case ssa.OpSPARC64LoweredAtomicCas32, ssa.OpSPARC64LoweredAtomicCas64:
		//	MOVD		new, TMP	// CAS overwrites its destination
		//	CAS(W|D)	(ptr), old, TMP
		//	MEMBAR		$15
		//	SUBCC		TMP, old, ZR
		//	MOVD		ZR, out
		//	MOVE		XCC, $1, out
		// Read the 32-bit condition codes for the 32-bit form. CASW
		// zero-extends the word it returns, but the old value it is
		// compared against is an arbitrary SSA value that may be held
		// sign-extended, so a 64-bit compare of the two disagrees
		// whenever bit 31 is set. SUBCC sets both code sets, and icc
		// looks only at the low word.
		cas, cc := sparc64.ACASD, int16(sparc64.REG_XCC)
		if v.Op == ssa.OpSPARC64LoweredAtomicCas32 {
			cas, cc = sparc64.ACASW, sparc64.REG_ICC
		}
		ptr, old, new, out := v.Args[0].Reg(), v.Args[1].Reg(), v.Args[2].Reg(), v.Reg0()
		p := s.Prog(sparc64.AMOVD)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = new
		p.To.Type = obj.TYPE_REG
		p.To.Reg = sparc64.REG_TMP
		p1 := s.Prog(cas)
		p1.From.Type = obj.TYPE_MEM
		p1.From.Reg = ptr
		p1.Reg = old
		p1.To.Type = obj.TYPE_REG
		p1.To.Reg = sparc64.REG_TMP
		membar(s, membarFull)
		p2 := s.Prog(sparc64.ASUBCC)
		p2.From.Type = obj.TYPE_REG
		p2.From.Reg = sparc64.REG_TMP
		p2.Reg = old
		p2.To.Type = obj.TYPE_REG
		p2.To.Reg = sparc64.REG_ZR
		p3 := s.Prog(sparc64.AMOVD)
		p3.From.Type = obj.TYPE_REG
		p3.From.Reg = sparc64.REG_ZR
		p3.To.Type = obj.TYPE_REG
		p3.To.Reg = out
		p4 := s.Prog(sparc64.AMOVE)
		p4.From.Type = obj.TYPE_REG
		p4.From.Reg = cc
		p4.AddRestSourceConst(1)
		p4.To.Type = obj.TYPE_REG
		p4.To.Reg = out

	case ssa.OpSPARC64LoweredAtomicAnd32, ssa.OpSPARC64LoweredAtomicOr32:
		// again:
		//	MOVUW	(ptr), TMP
		//	AND/OR	val, TMP, TMP2
		//	CASW	(ptr), TMP, TMP2
		//	SUBCC	TMP2, TMP, ZR
		//	BNED	again
		//	MEMBAR	$15
		logical := sparc64.AAND
		if v.Op == ssa.OpSPARC64LoweredAtomicOr32 {
			logical = sparc64.AOR
		}
		ptr, val := v.Args[0].Reg(), v.Args[1].Reg()
		top := s.Prog(sparc64.AMOVUW)
		top.From.Type = obj.TYPE_MEM
		top.From.Reg = ptr
		top.To.Type = obj.TYPE_REG
		top.To.Reg = sparc64.REG_TMP
		p1 := s.Prog(logical)
		p1.From.Type = obj.TYPE_REG
		p1.From.Reg = val
		p1.Reg = sparc64.REG_TMP
		p1.To.Type = obj.TYPE_REG
		p1.To.Reg = sparc64.REG_TMP2
		p2 := s.Prog(sparc64.ACASW)
		p2.From.Type = obj.TYPE_MEM
		p2.From.Reg = ptr
		p2.Reg = sparc64.REG_TMP
		p2.To.Type = obj.TYPE_REG
		p2.To.Reg = sparc64.REG_TMP2
		casLoop(s, sparc64.REG_TMP2, sparc64.REG_TMP, top)

	case ssa.OpSPARC64LoweredAtomicAnd8, ssa.OpSPARC64LoweredAtomicOr8:
		// SPARC has no byte-width CAS, so operate on the containing
		// aligned word. Big-endian, so the byte at p sits at bit
		// (3 - (p & 3)) * 8.
		//
		//	AND	$3, ptr, RT1		// RT1 = ptr & 3
		//	MOVD	$3, RT2
		//	SUB	RT1, RT2, RT1
		//	SLLD	$3, RT1, RT1		// RT1 = shift
		//	AND	$-4, ptr, TMP		// TMP = aligned pointer
		//	AND	$255, val, TMP2\n		//	SLLD	RT1, TMP2, TMP2		// TMP2 = (val&0xff) << shift
		// and only, so the other bytes survive the AND:
		//	MOVD	$255, RT2
		//	SLLD	RT1, RT2, RT2
		//	XOR	$-1, RT2, RT2		// RT2 = ~mask
		//	OR	RT2, TMP2, TMP2
		// again:
		//	MOVUW	(TMP), RT2
		//	AND/OR	TMP2, RT2, RT1
		//	CASW	(TMP), RT2, RT1
		//	SUBCC	RT1, RT2, ZR
		//	BNED	again
		//	MEMBAR	$15
		isAnd := v.Op == ssa.OpSPARC64LoweredAtomicAnd8
		ptr, val := v.Args[0].Reg(), v.Args[1].Reg()

		p := s.Prog(sparc64.AAND)
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = 3
		p.Reg = ptr
		p.To.Type = obj.TYPE_REG
		p.To.Reg = sparc64.REG_RT1
		p = s.Prog(sparc64.AMOVD)
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = 3
		p.To.Type = obj.TYPE_REG
		p.To.Reg = sparc64.REG_RT2
		p = s.Prog(sparc64.ASUB)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = sparc64.REG_RT1
		p.Reg = sparc64.REG_RT2
		p.To.Type = obj.TYPE_REG
		p.To.Reg = sparc64.REG_RT1
		p = s.Prog(sparc64.ASLLD)
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = 3
		p.Reg = sparc64.REG_RT1
		p.To.Type = obj.TYPE_REG
		p.To.Reg = sparc64.REG_RT1
		p = s.Prog(sparc64.AAND)
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = -4
		p.Reg = ptr
		p.To.Type = obj.TYPE_REG
		p.To.Reg = sparc64.REG_TMP
		// Mask the value to its byte before shifting it into place: it
		// is an arbitrary SSA value and may be held sign-extended, and
		// stray high bits would land on the neighbouring bytes.
		p = s.Prog(sparc64.AAND)
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = 255
		p.Reg = val
		p.To.Type = obj.TYPE_REG
		p.To.Reg = sparc64.REG_TMP2
		p = s.Prog(sparc64.ASLLD)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = sparc64.REG_RT1
		p.Reg = sparc64.REG_TMP2
		p.To.Type = obj.TYPE_REG
		p.To.Reg = sparc64.REG_TMP2
		if isAnd {
			p = s.Prog(sparc64.AMOVD)
			p.From.Type = obj.TYPE_CONST
			p.From.Offset = 255
			p.To.Type = obj.TYPE_REG
			p.To.Reg = sparc64.REG_RT2
			p = s.Prog(sparc64.ASLLD)
			p.From.Type = obj.TYPE_REG
			p.From.Reg = sparc64.REG_RT1
			p.Reg = sparc64.REG_RT2
			p.To.Type = obj.TYPE_REG
			p.To.Reg = sparc64.REG_RT2
			p = s.Prog(sparc64.AXOR)
			p.From.Type = obj.TYPE_CONST
			p.From.Offset = -1
			p.Reg = sparc64.REG_RT2
			p.To.Type = obj.TYPE_REG
			p.To.Reg = sparc64.REG_RT2
			p = s.Prog(sparc64.AOR)
			p.From.Type = obj.TYPE_REG
			p.From.Reg = sparc64.REG_RT2
			p.Reg = sparc64.REG_TMP2
			p.To.Type = obj.TYPE_REG
			p.To.Reg = sparc64.REG_TMP2
		}
		logical := sparc64.AAND
		if !isAnd {
			logical = sparc64.AOR
		}
		top := s.Prog(sparc64.AMOVUW)
		top.From.Type = obj.TYPE_MEM
		top.From.Reg = sparc64.REG_TMP
		top.To.Type = obj.TYPE_REG
		top.To.Reg = sparc64.REG_RT2
		p = s.Prog(logical)
		p.From.Type = obj.TYPE_REG
		p.From.Reg = sparc64.REG_TMP2
		p.Reg = sparc64.REG_RT2
		p.To.Type = obj.TYPE_REG
		p.To.Reg = sparc64.REG_RT1
		p = s.Prog(sparc64.ACASW)
		p.From.Type = obj.TYPE_MEM
		p.From.Reg = sparc64.REG_TMP
		p.Reg = sparc64.REG_RT2
		p.To.Type = obj.TYPE_REG
		p.To.Reg = sparc64.REG_RT1
		casLoop(s, sparc64.REG_RT1, sparc64.REG_RT2, top)

	case ssa.OpSPARC64LoweredZero:
		// MOVD	$size, TMP2
		// ADD	R1, TMP2, TMP2	// TMP2 = end pointer
		// loop:
		// MOV(w)	ZR, (R1)
		// ADD	$width, R1
		// CMP	TMP2, R1
		// BNED	loop
		//
		// The end pointer lives in the reserved TMP2 rather than in an
		// SSA value; see the op definition for why.
		sz, chunk, mov := zeroMoveAux(v)
		p := s.Prog(sparc64.AMOVD)
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = sz
		p.To.Type = obj.TYPE_REG
		p.To.Reg = sparc64.REG_TMP2
		p1 := s.Prog(sparc64.AADD)
		p1.From.Type = obj.TYPE_REG
		p1.From.Reg = sparc64.REG_R1
		p1.Reg = sparc64.REG_TMP2
		p1.To.Type = obj.TYPE_REG
		p1.To.Reg = sparc64.REG_TMP2
		p2 := s.Prog(mov)
		p2.From.Type = obj.TYPE_REG
		p2.From.Reg = sparc64.REG_ZR
		p2.To.Type = obj.TYPE_MEM
		p2.To.Reg = sparc64.REG_R1
		p3 := s.Prog(sparc64.AADD)
		p3.From.Type = obj.TYPE_CONST
		p3.From.Offset = chunk
		p3.Reg = sparc64.REG_R1
		p3.To.Type = obj.TYPE_REG
		p3.To.Reg = sparc64.REG_R1
		p4 := s.Prog(sparc64.ACMP)
		p4.From.Type = obj.TYPE_REG
		p4.From.Reg = sparc64.REG_TMP2
		p4.Reg = sparc64.REG_R1
		p5 := s.Prog(sparc64.ABNED)
		p5.To.Type = obj.TYPE_BRANCH
		p5.To.SetTarget(p2)

	case ssa.OpSPARC64LoweredMove:
		// MOVD	$size, TMP2
		// ADD	R1, TMP2, TMP2	// TMP2 = end of src
		// loop:
		// MOV(w)	(R1), TMP
		// MOV(w)	TMP, (R2)
		// ADD	$width, R1
		// ADD	$width, R2
		// CMP	TMP2, R1
		// BNED	loop
		sz, chunk, mov := zeroMoveAux(v)
		p := s.Prog(sparc64.AMOVD)
		p.From.Type = obj.TYPE_CONST
		p.From.Offset = sz
		p.To.Type = obj.TYPE_REG
		p.To.Reg = sparc64.REG_TMP2
		p1 := s.Prog(sparc64.AADD)
		p1.From.Type = obj.TYPE_REG
		p1.From.Reg = sparc64.REG_R1
		p1.Reg = sparc64.REG_TMP2
		p1.To.Type = obj.TYPE_REG
		p1.To.Reg = sparc64.REG_TMP2
		p2 := s.Prog(mov)
		p2.From.Type = obj.TYPE_MEM
		p2.From.Reg = sparc64.REG_R1
		p2.To.Type = obj.TYPE_REG
		p2.To.Reg = sparc64.REGTMP
		p3 := s.Prog(mov)
		p3.From.Type = obj.TYPE_REG
		p3.From.Reg = sparc64.REGTMP
		p3.To.Type = obj.TYPE_MEM
		p3.To.Reg = sparc64.REG_R2
		p4 := s.Prog(sparc64.AADD)
		p4.From.Type = obj.TYPE_CONST
		p4.From.Offset = chunk
		p4.Reg = sparc64.REG_R1
		p4.To.Type = obj.TYPE_REG
		p4.To.Reg = sparc64.REG_R1
		p5 := s.Prog(sparc64.AADD)
		p5.From.Type = obj.TYPE_CONST
		p5.From.Offset = chunk
		p5.Reg = sparc64.REG_R2
		p5.To.Type = obj.TYPE_REG
		p5.To.Reg = sparc64.REG_R2
		p6 := s.Prog(sparc64.ACMP)
		p6.From.Type = obj.TYPE_REG
		p6.From.Reg = sparc64.REG_TMP2
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
// off is relative to the base of the locals, which sit above the
// fixed 176-byte frame area (the register window save area plus the
// outgoing arguments), so the emitted offset must include both the
// stack bias and the fixed frame size.
func zerorange(pp *objw.Progs, p *obj.Prog, off, cnt int64, _ *uint32) *obj.Prog {
	if cnt == 0 {
		return p
	}
	for i := int64(0); i < cnt; i += 8 {
		p = pp.Append(p, sparc64.AMOVD, obj.TYPE_REG, sparc64.REG_ZR, 0,
			obj.TYPE_MEM, sparc64.REGSP, sparc64.StackBias+sparc64.MinStackFrameSize+off+i)
	}
	return p
}

// ginsnop emits a no-op, used to pad call sites.
func ginsnop(pp *objw.Progs) *obj.Prog {
	return pp.Prog(sparc64.ARNOP)
}
