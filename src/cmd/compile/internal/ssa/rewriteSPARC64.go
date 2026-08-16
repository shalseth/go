// Code generated from _gen/SPARC64.rules using 'go generate'; DO NOT EDIT.

package ssa

import "cmd/compile/internal/types"
import "cmd/compile/internal/ssa/block"

func rewriteValueSPARC64(v *Value) bool {
	switch v.Op {
	case OpAbs:
		v.Op = OpSPARC64FABSD
		return true
	case OpAdd16:
		v.Op = OpSPARC64ADD
		return true
	case OpAdd32:
		v.Op = OpSPARC64ADD
		return true
	case OpAdd32F:
		v.Op = OpSPARC64FADDS
		return true
	case OpAdd64:
		v.Op = OpSPARC64ADD
		return true
	case OpAdd64F:
		v.Op = OpSPARC64FADDD
		return true
	case OpAdd8:
		v.Op = OpSPARC64ADD
		return true
	case OpAddPtr:
		v.Op = OpSPARC64ADD
		return true
	case OpAddr:
		return rewriteValueSPARC64_OpAddr(v)
	case OpAnd16:
		v.Op = OpSPARC64AND
		return true
	case OpAnd32:
		v.Op = OpSPARC64AND
		return true
	case OpAnd64:
		v.Op = OpSPARC64AND
		return true
	case OpAnd8:
		v.Op = OpSPARC64AND
		return true
	case OpAndB:
		v.Op = OpSPARC64AND
		return true
	case OpAvg64u:
		return rewriteValueSPARC64_OpAvg64u(v)
	case OpClosureCall:
		v.Op = OpSPARC64CALLclosure
		return true
	case OpCom16:
		return rewriteValueSPARC64_OpCom16(v)
	case OpCom32:
		return rewriteValueSPARC64_OpCom32(v)
	case OpCom64:
		return rewriteValueSPARC64_OpCom64(v)
	case OpCom8:
		return rewriteValueSPARC64_OpCom8(v)
	case OpConst16:
		return rewriteValueSPARC64_OpConst16(v)
	case OpConst32:
		return rewriteValueSPARC64_OpConst32(v)
	case OpConst32F:
		v.Op = OpSPARC64FMOVSconst
		return true
	case OpConst64:
		return rewriteValueSPARC64_OpConst64(v)
	case OpConst64F:
		v.Op = OpSPARC64FMOVDconst
		return true
	case OpConst8:
		return rewriteValueSPARC64_OpConst8(v)
	case OpConstBool:
		return rewriteValueSPARC64_OpConstBool(v)
	case OpConstNil:
		return rewriteValueSPARC64_OpConstNil(v)
	case OpCvt32Fto32:
		return rewriteValueSPARC64_OpCvt32Fto32(v)
	case OpCvt32Fto64:
		v.Op = OpSPARC64FSTOX
		return true
	case OpCvt32Fto64F:
		v.Op = OpSPARC64FSTOD
		return true
	case OpCvt32to32F:
		return rewriteValueSPARC64_OpCvt32to32F(v)
	case OpCvt32to64F:
		return rewriteValueSPARC64_OpCvt32to64F(v)
	case OpCvt64Fto32:
		return rewriteValueSPARC64_OpCvt64Fto32(v)
	case OpCvt64Fto32F:
		v.Op = OpSPARC64FDTOS
		return true
	case OpCvt64Fto64:
		v.Op = OpSPARC64FDTOX
		return true
	case OpCvt64to32F:
		v.Op = OpSPARC64FXTOS
		return true
	case OpCvt64to64F:
		v.Op = OpSPARC64FXTOD
		return true
	case OpCvtBoolToUint8:
		v.Op = OpCopy
		return true
	case OpDiv16:
		return rewriteValueSPARC64_OpDiv16(v)
	case OpDiv16u:
		return rewriteValueSPARC64_OpDiv16u(v)
	case OpDiv32:
		return rewriteValueSPARC64_OpDiv32(v)
	case OpDiv32F:
		v.Op = OpSPARC64FDIVS
		return true
	case OpDiv32u:
		return rewriteValueSPARC64_OpDiv32u(v)
	case OpDiv64:
		return rewriteValueSPARC64_OpDiv64(v)
	case OpDiv64F:
		v.Op = OpSPARC64FDIVD
		return true
	case OpDiv64u:
		return rewriteValueSPARC64_OpDiv64u(v)
	case OpDiv8:
		return rewriteValueSPARC64_OpDiv8(v)
	case OpDiv8u:
		return rewriteValueSPARC64_OpDiv8u(v)
	case OpEq16:
		return rewriteValueSPARC64_OpEq16(v)
	case OpEq32:
		return rewriteValueSPARC64_OpEq32(v)
	case OpEq32F:
		return rewriteValueSPARC64_OpEq32F(v)
	case OpEq64:
		return rewriteValueSPARC64_OpEq64(v)
	case OpEq64F:
		return rewriteValueSPARC64_OpEq64F(v)
	case OpEq8:
		return rewriteValueSPARC64_OpEq8(v)
	case OpEqB:
		return rewriteValueSPARC64_OpEqB(v)
	case OpEqPtr:
		return rewriteValueSPARC64_OpEqPtr(v)
	case OpGetCallerPC:
		v.Op = OpSPARC64LoweredGetCallerPC
		return true
	case OpGetCallerSP:
		v.Op = OpSPARC64LoweredGetCallerSP
		return true
	case OpGetClosurePtr:
		v.Op = OpSPARC64LoweredGetClosurePtr
		return true
	case OpHmul32:
		return rewriteValueSPARC64_OpHmul32(v)
	case OpHmul32u:
		return rewriteValueSPARC64_OpHmul32u(v)
	case OpHmul64:
		return rewriteValueSPARC64_OpHmul64(v)
	case OpHmul64u:
		return rewriteValueSPARC64_OpHmul64u(v)
	case OpInterCall:
		v.Op = OpSPARC64CALLinter
		return true
	case OpIsInBounds:
		return rewriteValueSPARC64_OpIsInBounds(v)
	case OpIsNonNil:
		return rewriteValueSPARC64_OpIsNonNil(v)
	case OpIsSliceInBounds:
		return rewriteValueSPARC64_OpIsSliceInBounds(v)
	case OpLeq16:
		return rewriteValueSPARC64_OpLeq16(v)
	case OpLeq16U:
		return rewriteValueSPARC64_OpLeq16U(v)
	case OpLeq32:
		return rewriteValueSPARC64_OpLeq32(v)
	case OpLeq32F:
		return rewriteValueSPARC64_OpLeq32F(v)
	case OpLeq32U:
		return rewriteValueSPARC64_OpLeq32U(v)
	case OpLeq64:
		return rewriteValueSPARC64_OpLeq64(v)
	case OpLeq64F:
		return rewriteValueSPARC64_OpLeq64F(v)
	case OpLeq64U:
		return rewriteValueSPARC64_OpLeq64U(v)
	case OpLeq8:
		return rewriteValueSPARC64_OpLeq8(v)
	case OpLeq8U:
		return rewriteValueSPARC64_OpLeq8U(v)
	case OpLess16:
		return rewriteValueSPARC64_OpLess16(v)
	case OpLess16U:
		return rewriteValueSPARC64_OpLess16U(v)
	case OpLess32:
		return rewriteValueSPARC64_OpLess32(v)
	case OpLess32F:
		return rewriteValueSPARC64_OpLess32F(v)
	case OpLess32U:
		return rewriteValueSPARC64_OpLess32U(v)
	case OpLess64:
		return rewriteValueSPARC64_OpLess64(v)
	case OpLess64F:
		return rewriteValueSPARC64_OpLess64F(v)
	case OpLess64U:
		return rewriteValueSPARC64_OpLess64U(v)
	case OpLess8:
		return rewriteValueSPARC64_OpLess8(v)
	case OpLess8U:
		return rewriteValueSPARC64_OpLess8U(v)
	case OpLoad:
		return rewriteValueSPARC64_OpLoad(v)
	case OpLocalAddr:
		return rewriteValueSPARC64_OpLocalAddr(v)
	case OpLsh16x16:
		return rewriteValueSPARC64_OpLsh16x16(v)
	case OpLsh16x32:
		return rewriteValueSPARC64_OpLsh16x32(v)
	case OpLsh16x64:
		return rewriteValueSPARC64_OpLsh16x64(v)
	case OpLsh16x8:
		return rewriteValueSPARC64_OpLsh16x8(v)
	case OpLsh32x16:
		return rewriteValueSPARC64_OpLsh32x16(v)
	case OpLsh32x32:
		return rewriteValueSPARC64_OpLsh32x32(v)
	case OpLsh32x64:
		return rewriteValueSPARC64_OpLsh32x64(v)
	case OpLsh32x8:
		return rewriteValueSPARC64_OpLsh32x8(v)
	case OpLsh64x16:
		return rewriteValueSPARC64_OpLsh64x16(v)
	case OpLsh64x32:
		return rewriteValueSPARC64_OpLsh64x32(v)
	case OpLsh64x64:
		return rewriteValueSPARC64_OpLsh64x64(v)
	case OpLsh64x8:
		return rewriteValueSPARC64_OpLsh64x8(v)
	case OpLsh8x16:
		return rewriteValueSPARC64_OpLsh8x16(v)
	case OpLsh8x32:
		return rewriteValueSPARC64_OpLsh8x32(v)
	case OpLsh8x64:
		return rewriteValueSPARC64_OpLsh8x64(v)
	case OpLsh8x8:
		return rewriteValueSPARC64_OpLsh8x8(v)
	case OpMod16:
		return rewriteValueSPARC64_OpMod16(v)
	case OpMod16u:
		return rewriteValueSPARC64_OpMod16u(v)
	case OpMod32:
		return rewriteValueSPARC64_OpMod32(v)
	case OpMod32u:
		return rewriteValueSPARC64_OpMod32u(v)
	case OpMod64:
		return rewriteValueSPARC64_OpMod64(v)
	case OpMod64u:
		return rewriteValueSPARC64_OpMod64u(v)
	case OpMod8:
		return rewriteValueSPARC64_OpMod8(v)
	case OpMod8u:
		return rewriteValueSPARC64_OpMod8u(v)
	case OpMove:
		return rewriteValueSPARC64_OpMove(v)
	case OpMul16:
		v.Op = OpSPARC64MULD
		return true
	case OpMul32:
		v.Op = OpSPARC64MULD
		return true
	case OpMul32F:
		v.Op = OpSPARC64FMULS
		return true
	case OpMul64:
		v.Op = OpSPARC64MULD
		return true
	case OpMul64F:
		v.Op = OpSPARC64FMULD
		return true
	case OpMul64uhilo:
		v.Op = OpSPARC64MULDU
		return true
	case OpMul8:
		v.Op = OpSPARC64MULD
		return true
	case OpNeg16:
		return rewriteValueSPARC64_OpNeg16(v)
	case OpNeg32:
		return rewriteValueSPARC64_OpNeg32(v)
	case OpNeg32F:
		v.Op = OpSPARC64FNEGS
		return true
	case OpNeg64:
		return rewriteValueSPARC64_OpNeg64(v)
	case OpNeg64F:
		v.Op = OpSPARC64FNEGD
		return true
	case OpNeg8:
		return rewriteValueSPARC64_OpNeg8(v)
	case OpNeq16:
		return rewriteValueSPARC64_OpNeq16(v)
	case OpNeq32:
		return rewriteValueSPARC64_OpNeq32(v)
	case OpNeq32F:
		return rewriteValueSPARC64_OpNeq32F(v)
	case OpNeq64:
		return rewriteValueSPARC64_OpNeq64(v)
	case OpNeq64F:
		return rewriteValueSPARC64_OpNeq64F(v)
	case OpNeq8:
		return rewriteValueSPARC64_OpNeq8(v)
	case OpNeqB:
		return rewriteValueSPARC64_OpNeqB(v)
	case OpNeqPtr:
		return rewriteValueSPARC64_OpNeqPtr(v)
	case OpNilCheck:
		v.Op = OpSPARC64LoweredNilCheck
		return true
	case OpNot:
		return rewriteValueSPARC64_OpNot(v)
	case OpOffPtr:
		return rewriteValueSPARC64_OpOffPtr(v)
	case OpOr16:
		v.Op = OpSPARC64OR
		return true
	case OpOr32:
		v.Op = OpSPARC64OR
		return true
	case OpOr64:
		v.Op = OpSPARC64OR
		return true
	case OpOr8:
		v.Op = OpSPARC64OR
		return true
	case OpOrB:
		v.Op = OpSPARC64OR
		return true
	case OpPanicBounds:
		v.Op = OpSPARC64LoweredPanicBoundsRR
		return true
	case OpPubBarrier:
		v.Op = OpSPARC64LoweredPubBarrier
		return true
	case OpRotateLeft32:
		return rewriteValueSPARC64_OpRotateLeft32(v)
	case OpRotateLeft64:
		return rewriteValueSPARC64_OpRotateLeft64(v)
	case OpRound32F:
		v.Op = OpCopy
		return true
	case OpRound64F:
		v.Op = OpCopy
		return true
	case OpRsh16Ux16:
		return rewriteValueSPARC64_OpRsh16Ux16(v)
	case OpRsh16Ux32:
		return rewriteValueSPARC64_OpRsh16Ux32(v)
	case OpRsh16Ux64:
		return rewriteValueSPARC64_OpRsh16Ux64(v)
	case OpRsh16Ux8:
		return rewriteValueSPARC64_OpRsh16Ux8(v)
	case OpRsh16x16:
		return rewriteValueSPARC64_OpRsh16x16(v)
	case OpRsh16x32:
		return rewriteValueSPARC64_OpRsh16x32(v)
	case OpRsh16x64:
		return rewriteValueSPARC64_OpRsh16x64(v)
	case OpRsh16x8:
		return rewriteValueSPARC64_OpRsh16x8(v)
	case OpRsh32Ux16:
		return rewriteValueSPARC64_OpRsh32Ux16(v)
	case OpRsh32Ux32:
		return rewriteValueSPARC64_OpRsh32Ux32(v)
	case OpRsh32Ux64:
		return rewriteValueSPARC64_OpRsh32Ux64(v)
	case OpRsh32Ux8:
		return rewriteValueSPARC64_OpRsh32Ux8(v)
	case OpRsh32x16:
		return rewriteValueSPARC64_OpRsh32x16(v)
	case OpRsh32x32:
		return rewriteValueSPARC64_OpRsh32x32(v)
	case OpRsh32x64:
		return rewriteValueSPARC64_OpRsh32x64(v)
	case OpRsh32x8:
		return rewriteValueSPARC64_OpRsh32x8(v)
	case OpRsh64Ux16:
		return rewriteValueSPARC64_OpRsh64Ux16(v)
	case OpRsh64Ux32:
		return rewriteValueSPARC64_OpRsh64Ux32(v)
	case OpRsh64Ux64:
		return rewriteValueSPARC64_OpRsh64Ux64(v)
	case OpRsh64Ux8:
		return rewriteValueSPARC64_OpRsh64Ux8(v)
	case OpRsh64x16:
		return rewriteValueSPARC64_OpRsh64x16(v)
	case OpRsh64x32:
		return rewriteValueSPARC64_OpRsh64x32(v)
	case OpRsh64x64:
		return rewriteValueSPARC64_OpRsh64x64(v)
	case OpRsh64x8:
		return rewriteValueSPARC64_OpRsh64x8(v)
	case OpRsh8Ux16:
		return rewriteValueSPARC64_OpRsh8Ux16(v)
	case OpRsh8Ux32:
		return rewriteValueSPARC64_OpRsh8Ux32(v)
	case OpRsh8Ux64:
		return rewriteValueSPARC64_OpRsh8Ux64(v)
	case OpRsh8Ux8:
		return rewriteValueSPARC64_OpRsh8Ux8(v)
	case OpRsh8x16:
		return rewriteValueSPARC64_OpRsh8x16(v)
	case OpRsh8x32:
		return rewriteValueSPARC64_OpRsh8x32(v)
	case OpRsh8x64:
		return rewriteValueSPARC64_OpRsh8x64(v)
	case OpRsh8x8:
		return rewriteValueSPARC64_OpRsh8x8(v)
	case OpSPARC64ADD:
		return rewriteValueSPARC64_OpSPARC64ADD(v)
	case OpSPARC64ADDconst:
		return rewriteValueSPARC64_OpSPARC64ADDconst(v)
	case OpSPARC64AND:
		return rewriteValueSPARC64_OpSPARC64AND(v)
	case OpSPARC64ANDconst:
		return rewriteValueSPARC64_OpSPARC64ANDconst(v)
	case OpSPARC64CMP:
		return rewriteValueSPARC64_OpSPARC64CMP(v)
	case OpSPARC64LoweredPanicBoundsCR:
		return rewriteValueSPARC64_OpSPARC64LoweredPanicBoundsCR(v)
	case OpSPARC64LoweredPanicBoundsRC:
		return rewriteValueSPARC64_OpSPARC64LoweredPanicBoundsRC(v)
	case OpSPARC64LoweredPanicBoundsRR:
		return rewriteValueSPARC64_OpSPARC64LoweredPanicBoundsRR(v)
	case OpSPARC64MOVB:
		return rewriteValueSPARC64_OpSPARC64MOVB(v)
	case OpSPARC64MOVD:
		return rewriteValueSPARC64_OpSPARC64MOVD(v)
	case OpSPARC64MOVH:
		return rewriteValueSPARC64_OpSPARC64MOVH(v)
	case OpSPARC64MOVUB:
		return rewriteValueSPARC64_OpSPARC64MOVUB(v)
	case OpSPARC64MOVUH:
		return rewriteValueSPARC64_OpSPARC64MOVUH(v)
	case OpSPARC64MOVUW:
		return rewriteValueSPARC64_OpSPARC64MOVUW(v)
	case OpSPARC64MOVW:
		return rewriteValueSPARC64_OpSPARC64MOVW(v)
	case OpSPARC64OR:
		return rewriteValueSPARC64_OpSPARC64OR(v)
	case OpSPARC64ORconst:
		return rewriteValueSPARC64_OpSPARC64ORconst(v)
	case OpSPARC64SLLD:
		return rewriteValueSPARC64_OpSPARC64SLLD(v)
	case OpSPARC64SRAD:
		return rewriteValueSPARC64_OpSPARC64SRAD(v)
	case OpSPARC64SRLD:
		return rewriteValueSPARC64_OpSPARC64SRLD(v)
	case OpSPARC64SUB:
		return rewriteValueSPARC64_OpSPARC64SUB(v)
	case OpSPARC64SUBconst:
		return rewriteValueSPARC64_OpSPARC64SUBconst(v)
	case OpSPARC64XOR:
		return rewriteValueSPARC64_OpSPARC64XOR(v)
	case OpSPARC64XORconst:
		return rewriteValueSPARC64_OpSPARC64XORconst(v)
	case OpSignExt16to32:
		v.Op = OpSPARC64MOVH
		return true
	case OpSignExt16to64:
		v.Op = OpSPARC64MOVH
		return true
	case OpSignExt32to64:
		v.Op = OpSPARC64MOVW
		return true
	case OpSignExt8to16:
		v.Op = OpSPARC64MOVB
		return true
	case OpSignExt8to32:
		v.Op = OpSPARC64MOVB
		return true
	case OpSignExt8to64:
		v.Op = OpSPARC64MOVB
		return true
	case OpSlicemask:
		return rewriteValueSPARC64_OpSlicemask(v)
	case OpSqrt:
		v.Op = OpSPARC64FSQRTD
		return true
	case OpSqrt32:
		v.Op = OpSPARC64FSQRTS
		return true
	case OpStaticCall:
		v.Op = OpSPARC64CALLstatic
		return true
	case OpStore:
		return rewriteValueSPARC64_OpStore(v)
	case OpSub16:
		v.Op = OpSPARC64SUB
		return true
	case OpSub32:
		v.Op = OpSPARC64SUB
		return true
	case OpSub32F:
		v.Op = OpSPARC64FSUBS
		return true
	case OpSub64:
		v.Op = OpSPARC64SUB
		return true
	case OpSub64F:
		v.Op = OpSPARC64FSUBD
		return true
	case OpSub8:
		v.Op = OpSPARC64SUB
		return true
	case OpSubPtr:
		v.Op = OpSPARC64SUB
		return true
	case OpTailCall:
		v.Op = OpSPARC64CALLtail
		return true
	case OpTailCallInter:
		v.Op = OpSPARC64CALLtail
		return true
	case OpTrunc16to8:
		v.Op = OpCopy
		return true
	case OpTrunc32to16:
		v.Op = OpCopy
		return true
	case OpTrunc32to8:
		v.Op = OpCopy
		return true
	case OpTrunc64to16:
		v.Op = OpCopy
		return true
	case OpTrunc64to32:
		v.Op = OpCopy
		return true
	case OpTrunc64to8:
		v.Op = OpCopy
		return true
	case OpWB:
		v.Op = OpSPARC64LoweredWB
		return true
	case OpXor16:
		v.Op = OpSPARC64XOR
		return true
	case OpXor32:
		v.Op = OpSPARC64XOR
		return true
	case OpXor64:
		v.Op = OpSPARC64XOR
		return true
	case OpXor8:
		v.Op = OpSPARC64XOR
		return true
	case OpZero:
		return rewriteValueSPARC64_OpZero(v)
	case OpZeroExt16to32:
		v.Op = OpSPARC64MOVUH
		return true
	case OpZeroExt16to64:
		v.Op = OpSPARC64MOVUH
		return true
	case OpZeroExt32to64:
		v.Op = OpSPARC64MOVUW
		return true
	case OpZeroExt8to16:
		v.Op = OpSPARC64MOVUB
		return true
	case OpZeroExt8to32:
		v.Op = OpSPARC64MOVUB
		return true
	case OpZeroExt8to64:
		v.Op = OpSPARC64MOVUB
		return true
	}
	return false
}
func rewriteValueSPARC64_OpAddr(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Addr {sym} base)
	// result: (MOVDaddr {sym} base)
	for {
		sym := auxToSym(v.Aux)
		base := v_0
		v.reset(OpSPARC64MOVDaddr)
		v.Aux = symToAux(sym)
		v.AddArg(base)
		return true
	}
}
func rewriteValueSPARC64_OpAvg64u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Avg64u <t> x y)
	// result: (ADD (SRLDconst <t> [1] (SUB <t> x y)) y)
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSPARC64ADD)
		v0 := b.NewValue0(v.Pos, OpSPARC64SRLDconst, t)
		v0.AuxInt = int64ToAuxInt(1)
		v1 := b.NewValue0(v.Pos, OpSPARC64SUB, t)
		v1.AddArg2(x, y)
		v0.AddArg(v1)
		v.AddArg2(v0, y)
		return true
	}
}
func rewriteValueSPARC64_OpCom16(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Com16 x)
	// result: (XNOR x (MOVDconst [0]))
	for {
		x := v_0
		v.reset(OpSPARC64XNOR)
		v0 := b.NewValue0(v.Pos, OpSPARC64MOVDconst, typ.UInt64)
		v0.AuxInt = int64ToAuxInt(0)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpCom32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Com32 x)
	// result: (XNOR x (MOVDconst [0]))
	for {
		x := v_0
		v.reset(OpSPARC64XNOR)
		v0 := b.NewValue0(v.Pos, OpSPARC64MOVDconst, typ.UInt64)
		v0.AuxInt = int64ToAuxInt(0)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpCom64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Com64 x)
	// result: (XNOR x (MOVDconst [0]))
	for {
		x := v_0
		v.reset(OpSPARC64XNOR)
		v0 := b.NewValue0(v.Pos, OpSPARC64MOVDconst, typ.UInt64)
		v0.AuxInt = int64ToAuxInt(0)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpCom8(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Com8 x)
	// result: (XNOR x (MOVDconst [0]))
	for {
		x := v_0
		v.reset(OpSPARC64XNOR)
		v0 := b.NewValue0(v.Pos, OpSPARC64MOVDconst, typ.UInt64)
		v0.AuxInt = int64ToAuxInt(0)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpConst16(v *Value) bool {
	// match: (Const16 [val])
	// result: (MOVDconst [int64(val)])
	for {
		val := auxIntToInt16(v.AuxInt)
		v.reset(OpSPARC64MOVDconst)
		v.AuxInt = int64ToAuxInt(int64(val))
		return true
	}
}
func rewriteValueSPARC64_OpConst32(v *Value) bool {
	// match: (Const32 [val])
	// result: (MOVDconst [int64(val)])
	for {
		val := auxIntToInt32(v.AuxInt)
		v.reset(OpSPARC64MOVDconst)
		v.AuxInt = int64ToAuxInt(int64(val))
		return true
	}
}
func rewriteValueSPARC64_OpConst64(v *Value) bool {
	// match: (Const64 [val])
	// result: (MOVDconst [int64(val)])
	for {
		val := auxIntToInt64(v.AuxInt)
		v.reset(OpSPARC64MOVDconst)
		v.AuxInt = int64ToAuxInt(int64(val))
		return true
	}
}
func rewriteValueSPARC64_OpConst8(v *Value) bool {
	// match: (Const8 [val])
	// result: (MOVDconst [int64(val)])
	for {
		val := auxIntToInt8(v.AuxInt)
		v.reset(OpSPARC64MOVDconst)
		v.AuxInt = int64ToAuxInt(int64(val))
		return true
	}
}
func rewriteValueSPARC64_OpConstBool(v *Value) bool {
	// match: (ConstBool [t])
	// result: (MOVDconst [b2i(t)])
	for {
		t := auxIntToBool(v.AuxInt)
		v.reset(OpSPARC64MOVDconst)
		v.AuxInt = int64ToAuxInt(b2i(t))
		return true
	}
}
func rewriteValueSPARC64_OpConstNil(v *Value) bool {
	// match: (ConstNil)
	// result: (MOVDconst [0])
	for {
		v.reset(OpSPARC64MOVDconst)
		v.AuxInt = int64ToAuxInt(0)
		return true
	}
}
func rewriteValueSPARC64_OpCvt32Fto32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Cvt32Fto32 x)
	// result: (MOVW <typ.Int64> (FSTOX <typ.Int64> x))
	for {
		x := v_0
		v.reset(OpSPARC64MOVW)
		v.Type = typ.Int64
		v0 := b.NewValue0(v.Pos, OpSPARC64FSTOX, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpCvt32to32F(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Cvt32to32F x)
	// result: (FXTOS (SignExt32to64 <typ.Int64> x))
	for {
		x := v_0
		v.reset(OpSPARC64FXTOS)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpCvt32to64F(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Cvt32to64F x)
	// result: (FXTOD (SignExt32to64 <typ.Int64> x))
	for {
		x := v_0
		v.reset(OpSPARC64FXTOD)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpCvt64Fto32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Cvt64Fto32 x)
	// result: (MOVW <typ.Int64> (FDTOX <typ.Int64> x))
	for {
		x := v_0
		v.reset(OpSPARC64MOVW)
		v.Type = typ.Int64
		v0 := b.NewValue0(v.Pos, OpSPARC64FDTOX, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpDiv16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div16 x y)
	// result: (SDIVD (SignExt16to64 <typ.Int64> x) (SignExt16to64 <typ.Int64> y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64SDIVD)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSPARC64_OpDiv16u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div16u x y)
	// result: (UDIVD (ZeroExt16to64 <typ.UInt64> x) (ZeroExt16to64 <typ.UInt64> y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64UDIVD)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSPARC64_OpDiv32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div32 x y)
	// result: (SDIVD (SignExt32to64 <typ.Int64> x) (SignExt32to64 <typ.Int64> y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64SDIVD)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSPARC64_OpDiv32u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div32u x y)
	// result: (UDIVD (ZeroExt32to64 <typ.UInt64> x) (ZeroExt32to64 <typ.UInt64> y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64UDIVD)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSPARC64_OpDiv64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Div64 [false] x y)
	// result: (SDIVD x y)
	for {
		if auxIntToBool(v.AuxInt) != false {
			break
		}
		x := v_0
		y := v_1
		v.reset(OpSPARC64SDIVD)
		v.AddArg2(x, y)
		return true
	}
	return false
}
func rewriteValueSPARC64_OpDiv64u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Div64u x y)
	// result: (UDIVD x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64UDIVD)
		v.AddArg2(x, y)
		return true
	}
}
func rewriteValueSPARC64_OpDiv8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div8 x y)
	// result: (SDIVD (SignExt8to64 <typ.Int64> x) (SignExt8to64 <typ.Int64> y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64SDIVD)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSPARC64_OpDiv8u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div8u x y)
	// result: (UDIVD (ZeroExt8to64 <typ.UInt64> x) (ZeroExt8to64 <typ.UInt64> y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64UDIVD)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSPARC64_OpEq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq16 x y)
	// result: (Equal (CMP (ZeroExt16to64 x) (ZeroExt16to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64Equal)
		v0 := b.NewValue0(v.Pos, OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpEq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq32 x y)
	// result: (Equal (CMP (ZeroExt32to64 x) (ZeroExt32to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64Equal)
		v0 := b.NewValue0(v.Pos, OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpEq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Eq32F x y)
	// result: (FEqual (FCMPS x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64FEqual)
		v0 := b.NewValue0(v.Pos, OpSPARC64FCMPS, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpEq64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Eq64 x y)
	// result: (Equal (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64Equal)
		v0 := b.NewValue0(v.Pos, OpSPARC64CMP, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpEq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Eq64F x y)
	// result: (FEqual (FCMPD x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64FEqual)
		v0 := b.NewValue0(v.Pos, OpSPARC64FCMPD, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpEq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq8 x y)
	// result: (Equal (CMP (ZeroExt8to64 x) (ZeroExt8to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64Equal)
		v0 := b.NewValue0(v.Pos, OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpEqB(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (EqB x y)
	// result: (Equal (CMP (ZeroExt8to64 x) (ZeroExt8to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64Equal)
		v0 := b.NewValue0(v.Pos, OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpEqPtr(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (EqPtr x y)
	// result: (Equal (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64Equal)
		v0 := b.NewValue0(v.Pos, OpSPARC64CMP, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpHmul32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Hmul32 x y)
	// result: (SRADconst [32] (MULD <typ.Int64> (SignExt32to64 <typ.Int64> x) (SignExt32to64 <typ.Int64> y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64SRADconst)
		v.AuxInt = int64ToAuxInt(32)
		v0 := b.NewValue0(v.Pos, OpSPARC64MULD, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpHmul32u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Hmul32u x y)
	// result: (SRLDconst [32] (MULD <typ.UInt64> (ZeroExt32to64 <typ.UInt64> x) (ZeroExt32to64 <typ.UInt64> y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64SRLDconst)
		v.AuxInt = int64ToAuxInt(32)
		v0 := b.NewValue0(v.Pos, OpSPARC64MULD, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpHmul64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Hmul64 x y)
	// result: (SUB <typ.Int64> (SUB <typ.Int64> (UMULXHI <typ.Int64> x y) (AND <typ.Int64> y (SRADconst <typ.Int64> [63] x))) (AND <typ.Int64> x (SRADconst <typ.Int64> [63] y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64SUB)
		v.Type = typ.Int64
		v0 := b.NewValue0(v.Pos, OpSPARC64SUB, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpSPARC64UMULXHI, typ.Int64)
		v1.AddArg2(x, y)
		v2 := b.NewValue0(v.Pos, OpSPARC64AND, typ.Int64)
		v3 := b.NewValue0(v.Pos, OpSPARC64SRADconst, typ.Int64)
		v3.AuxInt = int64ToAuxInt(63)
		v3.AddArg(x)
		v2.AddArg2(y, v3)
		v0.AddArg2(v1, v2)
		v4 := b.NewValue0(v.Pos, OpSPARC64AND, typ.Int64)
		v5 := b.NewValue0(v.Pos, OpSPARC64SRADconst, typ.Int64)
		v5.AuxInt = int64ToAuxInt(63)
		v5.AddArg(y)
		v4.AddArg2(x, v5)
		v.AddArg2(v0, v4)
		return true
	}
}
func rewriteValueSPARC64_OpHmul64u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Hmul64u x y)
	// result: (UMULXHI x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64UMULXHI)
		v.AddArg2(x, y)
		return true
	}
}
func rewriteValueSPARC64_OpIsInBounds(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (IsInBounds idx len)
	// result: (LessThanU (CMP idx len))
	for {
		idx := v_0
		len := v_1
		v.reset(OpSPARC64LessThanU)
		v0 := b.NewValue0(v.Pos, OpSPARC64CMP, types.TypeFlags)
		v0.AddArg2(idx, len)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpIsNonNil(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (IsNonNil ptr)
	// result: (NotEqual (CMPconst [0] ptr))
	for {
		ptr := v_0
		v.reset(OpSPARC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpSPARC64CMPconst, types.TypeFlags)
		v0.AuxInt = int64ToAuxInt(0)
		v0.AddArg(ptr)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpIsSliceInBounds(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (IsSliceInBounds idx len)
	// result: (LessEqualU (CMP idx len))
	for {
		idx := v_0
		len := v_1
		v.reset(OpSPARC64LessEqualU)
		v0 := b.NewValue0(v.Pos, OpSPARC64CMP, types.TypeFlags)
		v0.AddArg2(idx, len)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpLeq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq16 x y)
	// result: (LessEqual (CMP (SignExt16to64 x) (SignExt16to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64LessEqual)
		v0 := b.NewValue0(v.Pos, OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpLeq16U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq16U x y)
	// result: (LessEqualU (CMP (ZeroExt16to64 x) (ZeroExt16to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64LessEqualU)
		v0 := b.NewValue0(v.Pos, OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpLeq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq32 x y)
	// result: (LessEqual (CMP (SignExt32to64 x) (SignExt32to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64LessEqual)
		v0 := b.NewValue0(v.Pos, OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpLeq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Leq32F x y)
	// result: (FLessEqual (FCMPS x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64FLessEqual)
		v0 := b.NewValue0(v.Pos, OpSPARC64FCMPS, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpLeq32U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq32U x y)
	// result: (LessEqualU (CMP (ZeroExt32to64 x) (ZeroExt32to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64LessEqualU)
		v0 := b.NewValue0(v.Pos, OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpLeq64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Leq64 x y)
	// result: (LessEqual (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64LessEqual)
		v0 := b.NewValue0(v.Pos, OpSPARC64CMP, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpLeq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Leq64F x y)
	// result: (FLessEqual (FCMPD x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64FLessEqual)
		v0 := b.NewValue0(v.Pos, OpSPARC64FCMPD, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpLeq64U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Leq64U x y)
	// result: (LessEqualU (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64LessEqualU)
		v0 := b.NewValue0(v.Pos, OpSPARC64CMP, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpLeq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq8 x y)
	// result: (LessEqual (CMP (SignExt8to64 x) (SignExt8to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64LessEqual)
		v0 := b.NewValue0(v.Pos, OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpLeq8U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq8U x y)
	// result: (LessEqualU (CMP (ZeroExt8to64 x) (ZeroExt8to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64LessEqualU)
		v0 := b.NewValue0(v.Pos, OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpLess16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less16 x y)
	// result: (LessThan (CMP (SignExt16to64 x) (SignExt16to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64LessThan)
		v0 := b.NewValue0(v.Pos, OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpLess16U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less16U x y)
	// result: (LessThanU (CMP (ZeroExt16to64 x) (ZeroExt16to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64LessThanU)
		v0 := b.NewValue0(v.Pos, OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpLess32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less32 x y)
	// result: (LessThan (CMP (SignExt32to64 x) (SignExt32to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64LessThan)
		v0 := b.NewValue0(v.Pos, OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpLess32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Less32F x y)
	// result: (FLessThan (FCMPS x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64FLessThan)
		v0 := b.NewValue0(v.Pos, OpSPARC64FCMPS, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpLess32U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less32U x y)
	// result: (LessThanU (CMP (ZeroExt32to64 x) (ZeroExt32to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64LessThanU)
		v0 := b.NewValue0(v.Pos, OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpLess64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Less64 x y)
	// result: (LessThan (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64LessThan)
		v0 := b.NewValue0(v.Pos, OpSPARC64CMP, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpLess64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Less64F x y)
	// result: (FLessThan (FCMPD x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64FLessThan)
		v0 := b.NewValue0(v.Pos, OpSPARC64FCMPD, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpLess64U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Less64U x y)
	// result: (LessThanU (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64LessThanU)
		v0 := b.NewValue0(v.Pos, OpSPARC64CMP, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpLess8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less8 x y)
	// result: (LessThan (CMP (SignExt8to64 x) (SignExt8to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64LessThan)
		v0 := b.NewValue0(v.Pos, OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpLess8U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less8U x y)
	// result: (LessThanU (CMP (ZeroExt8to64 x) (ZeroExt8to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64LessThanU)
		v0 := b.NewValue0(v.Pos, OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpLoad(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Load <t> ptr mem)
	// cond: t.IsBoolean()
	// result: (MOVUBload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(t.IsBoolean()) {
			break
		}
		v.reset(OpSPARC64MOVUBload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is8BitInt(t) && t.IsSigned()
	// result: (MOVBload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is8BitInt(t) && t.IsSigned()) {
			break
		}
		v.reset(OpSPARC64MOVBload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is8BitInt(t) && !t.IsSigned()
	// result: (MOVUBload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is8BitInt(t) && !t.IsSigned()) {
			break
		}
		v.reset(OpSPARC64MOVUBload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is16BitInt(t) && t.IsSigned()
	// result: (MOVHload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is16BitInt(t) && t.IsSigned()) {
			break
		}
		v.reset(OpSPARC64MOVHload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is16BitInt(t) && !t.IsSigned()
	// result: (MOVUHload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is16BitInt(t) && !t.IsSigned()) {
			break
		}
		v.reset(OpSPARC64MOVUHload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is32BitInt(t) && t.IsSigned()
	// result: (MOVWload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is32BitInt(t) && t.IsSigned()) {
			break
		}
		v.reset(OpSPARC64MOVWload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is32BitInt(t) && !t.IsSigned()
	// result: (MOVUWload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is32BitInt(t) && !t.IsSigned()) {
			break
		}
		v.reset(OpSPARC64MOVUWload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (is64BitInt(t) || isPtr(t))
	// result: (MOVDload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is64BitInt(t) || isPtr(t)) {
			break
		}
		v.reset(OpSPARC64MOVDload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is32BitFloat(t)
	// result: (FMOVSload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is32BitFloat(t)) {
			break
		}
		v.reset(OpSPARC64FMOVSload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is64BitFloat(t)
	// result: (FMOVDload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is64BitFloat(t)) {
			break
		}
		v.reset(OpSPARC64FMOVDload)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueSPARC64_OpLocalAddr(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (LocalAddr <t> {sym} base mem)
	// cond: t.Elem().HasPointers()
	// result: (MOVDaddr {sym} (SPanchored base mem))
	for {
		t := v.Type
		sym := auxToSym(v.Aux)
		base := v_0
		mem := v_1
		if !(t.Elem().HasPointers()) {
			break
		}
		v.reset(OpSPARC64MOVDaddr)
		v.Aux = symToAux(sym)
		v0 := b.NewValue0(v.Pos, OpSPanchored, typ.Uintptr)
		v0.AddArg2(base, mem)
		v.AddArg(v0)
		return true
	}
	// match: (LocalAddr <t> {sym} base _)
	// cond: !t.Elem().HasPointers()
	// result: (MOVDaddr {sym} base)
	for {
		t := v.Type
		sym := auxToSym(v.Aux)
		base := v_0
		if !(!t.Elem().HasPointers()) {
			break
		}
		v.reset(OpSPARC64MOVDaddr)
		v.Aux = symToAux(sym)
		v.AddArg(base)
		return true
	}
	return false
}
func rewriteValueSPARC64_OpLsh16x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x16 x y)
	// result: (Lsh16x64 x (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpLsh16x64)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpLsh16x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x32 x y)
	// result: (Lsh16x64 x (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpLsh16x64)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpLsh16x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x64 <t> x (Const64 [c]))
	// cond: uint64(c) < 64
	// result: (SLLDconst <t> [c] x)
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpSPARC64SLLDconst)
		v.Type = t
		v.AuxInt = int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (Lsh16x64 _ (Const64 [c]))
	// cond: uint64(c) >= 64
	// result: (MOVDconst [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpSPARC64MOVDconst)
		v.AuxInt = int64ToAuxInt(0)
		return true
	}
	// match: (Lsh16x64 <t> x y)
	// result: (AND (NEG <t> (LessThanU <typ.UInt64> (CMPconst <types.TypeFlags> [64] y))) (SLLD <t> x y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSPARC64AND)
		v0 := b.NewValue0(v.Pos, OpSPARC64NEG, t)
		v1 := b.NewValue0(v.Pos, OpSPARC64LessThanU, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpSPARC64CMPconst, types.TypeFlags)
		v2.AuxInt = int64ToAuxInt(64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, OpSPARC64SLLD, t)
		v3.AddArg2(x, y)
		v.AddArg2(v0, v3)
		return true
	}
}
func rewriteValueSPARC64_OpLsh16x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x8 x y)
	// result: (Lsh16x64 x (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpLsh16x64)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpLsh32x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x16 x y)
	// result: (Lsh32x64 x (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpLsh32x64)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpLsh32x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x32 x y)
	// result: (Lsh32x64 x (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpLsh32x64)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpLsh32x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x64 <t> x (Const64 [c]))
	// cond: uint64(c) < 64
	// result: (SLLDconst <t> [c] x)
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpSPARC64SLLDconst)
		v.Type = t
		v.AuxInt = int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (Lsh32x64 _ (Const64 [c]))
	// cond: uint64(c) >= 64
	// result: (MOVDconst [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpSPARC64MOVDconst)
		v.AuxInt = int64ToAuxInt(0)
		return true
	}
	// match: (Lsh32x64 <t> x y)
	// result: (AND (NEG <t> (LessThanU <typ.UInt64> (CMPconst <types.TypeFlags> [64] y))) (SLLD <t> x y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSPARC64AND)
		v0 := b.NewValue0(v.Pos, OpSPARC64NEG, t)
		v1 := b.NewValue0(v.Pos, OpSPARC64LessThanU, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpSPARC64CMPconst, types.TypeFlags)
		v2.AuxInt = int64ToAuxInt(64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, OpSPARC64SLLD, t)
		v3.AddArg2(x, y)
		v.AddArg2(v0, v3)
		return true
	}
}
func rewriteValueSPARC64_OpLsh32x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x8 x y)
	// result: (Lsh32x64 x (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpLsh32x64)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpLsh64x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x16 x y)
	// result: (Lsh64x64 x (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpLsh64x64)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpLsh64x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x32 x y)
	// result: (Lsh64x64 x (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpLsh64x64)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpLsh64x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x64 <t> x (Const64 [c]))
	// cond: uint64(c) < 64
	// result: (SLLDconst <t> [c] x)
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpSPARC64SLLDconst)
		v.Type = t
		v.AuxInt = int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (Lsh64x64 _ (Const64 [c]))
	// cond: uint64(c) >= 64
	// result: (MOVDconst [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpSPARC64MOVDconst)
		v.AuxInt = int64ToAuxInt(0)
		return true
	}
	// match: (Lsh64x64 <t> x y)
	// result: (AND (NEG <t> (LessThanU <typ.UInt64> (CMPconst <types.TypeFlags> [64] y))) (SLLD <t> x y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSPARC64AND)
		v0 := b.NewValue0(v.Pos, OpSPARC64NEG, t)
		v1 := b.NewValue0(v.Pos, OpSPARC64LessThanU, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpSPARC64CMPconst, types.TypeFlags)
		v2.AuxInt = int64ToAuxInt(64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, OpSPARC64SLLD, t)
		v3.AddArg2(x, y)
		v.AddArg2(v0, v3)
		return true
	}
}
func rewriteValueSPARC64_OpLsh64x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x8 x y)
	// result: (Lsh64x64 x (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpLsh64x64)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpLsh8x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x16 x y)
	// result: (Lsh8x64 x (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpLsh8x64)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpLsh8x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x32 x y)
	// result: (Lsh8x64 x (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpLsh8x64)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpLsh8x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x64 <t> x (Const64 [c]))
	// cond: uint64(c) < 64
	// result: (SLLDconst <t> [c] x)
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpSPARC64SLLDconst)
		v.Type = t
		v.AuxInt = int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (Lsh8x64 _ (Const64 [c]))
	// cond: uint64(c) >= 64
	// result: (MOVDconst [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpSPARC64MOVDconst)
		v.AuxInt = int64ToAuxInt(0)
		return true
	}
	// match: (Lsh8x64 <t> x y)
	// result: (AND (NEG <t> (LessThanU <typ.UInt64> (CMPconst <types.TypeFlags> [64] y))) (SLLD <t> x y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSPARC64AND)
		v0 := b.NewValue0(v.Pos, OpSPARC64NEG, t)
		v1 := b.NewValue0(v.Pos, OpSPARC64LessThanU, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpSPARC64CMPconst, types.TypeFlags)
		v2.AuxInt = int64ToAuxInt(64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, OpSPARC64SLLD, t)
		v3.AddArg2(x, y)
		v.AddArg2(v0, v3)
		return true
	}
}
func rewriteValueSPARC64_OpLsh8x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x8 x y)
	// result: (Lsh8x64 x (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpLsh8x64)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpMod16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod16 x y)
	// result: (Mod64 (SignExt16to64 <typ.Int64> x) (SignExt16to64 <typ.Int64> y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMod64)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSPARC64_OpMod16u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod16u x y)
	// result: (Mod64u (ZeroExt16to64 <typ.UInt64> x) (ZeroExt16to64 <typ.UInt64> y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMod64u)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSPARC64_OpMod32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod32 x y)
	// result: (Mod64 (SignExt32to64 <typ.Int64> x) (SignExt32to64 <typ.Int64> y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMod64)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSPARC64_OpMod32u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod32u x y)
	// result: (Mod64u (ZeroExt32to64 <typ.UInt64> x) (ZeroExt32to64 <typ.UInt64> y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMod64u)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSPARC64_OpMod64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod64 x y)
	// result: (SUB <typ.Int64> x (MULD <typ.Int64> (SDIVD <typ.Int64> x y) y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64SUB)
		v.Type = typ.Int64
		v0 := b.NewValue0(v.Pos, OpSPARC64MULD, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpSPARC64SDIVD, typ.Int64)
		v1.AddArg2(x, y)
		v0.AddArg2(v1, y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpMod64u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod64u x y)
	// result: (SUB <typ.UInt64> x (MULD <typ.UInt64> (UDIVD <typ.UInt64> x y) y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64SUB)
		v.Type = typ.UInt64
		v0 := b.NewValue0(v.Pos, OpSPARC64MULD, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpSPARC64UDIVD, typ.UInt64)
		v1.AddArg2(x, y)
		v0.AddArg2(v1, y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpMod8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod8 x y)
	// result: (Mod64 (SignExt8to64 <typ.Int64> x) (SignExt8to64 <typ.Int64> y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMod64)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSPARC64_OpMod8u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod8u x y)
	// result: (Mod64u (ZeroExt8to64 <typ.UInt64> x) (ZeroExt8to64 <typ.UInt64> y))
	for {
		x := v_0
		y := v_1
		v.reset(OpMod64u)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSPARC64_OpMove(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: (Move [0] _ _ mem)
	// result: mem
	for {
		if auxIntToInt64(v.AuxInt) != 0 {
			break
		}
		mem := v_2
		v.copyOf(mem)
		return true
	}
	// match: (Move [1] dst src mem)
	// result: (MOVBstore dst (MOVBload src mem) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 1 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpSPARC64MOVBstore)
		v0 := b.NewValue0(v.Pos, OpSPARC64MOVBload, typ.Int8)
		v0.AddArg2(src, mem)
		v.AddArg3(dst, v0, mem)
		return true
	}
	// match: (Move [2] {t} dst src mem)
	// cond: t.Alignment()%2 == 0
	// result: (MOVHstore dst (MOVHload src mem) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 2 {
			break
		}
		t := auxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.Alignment()%2 == 0) {
			break
		}
		v.reset(OpSPARC64MOVHstore)
		v0 := b.NewValue0(v.Pos, OpSPARC64MOVHload, typ.Int16)
		v0.AddArg2(src, mem)
		v.AddArg3(dst, v0, mem)
		return true
	}
	// match: (Move [4] {t} dst src mem)
	// cond: t.Alignment()%4 == 0
	// result: (MOVWstore dst (MOVWload src mem) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 4 {
			break
		}
		t := auxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.Alignment()%4 == 0) {
			break
		}
		v.reset(OpSPARC64MOVWstore)
		v0 := b.NewValue0(v.Pos, OpSPARC64MOVWload, typ.Int32)
		v0.AddArg2(src, mem)
		v.AddArg3(dst, v0, mem)
		return true
	}
	// match: (Move [8] {t} dst src mem)
	// cond: t.Alignment()%8 == 0
	// result: (MOVDstore dst (MOVDload src mem) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 8 {
			break
		}
		t := auxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.Alignment()%8 == 0) {
			break
		}
		v.reset(OpSPARC64MOVDstore)
		v0 := b.NewValue0(v.Pos, OpSPARC64MOVDload, typ.UInt64)
		v0.AddArg2(src, mem)
		v.AddArg3(dst, v0, mem)
		return true
	}
	// match: (Move [s] {t} dst src mem)
	// result: (LoweredMove [t.Alignment()] dst src (ADDconst <src.Type> [s-moveSize(t.Alignment(), config)] src) mem)
	for {
		s := auxIntToInt64(v.AuxInt)
		t := auxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpSPARC64LoweredMove)
		v.AuxInt = int64ToAuxInt(t.Alignment())
		v0 := b.NewValue0(v.Pos, OpSPARC64ADDconst, src.Type)
		v0.AuxInt = int64ToAuxInt(s - moveSize(t.Alignment(), config))
		v0.AddArg(src)
		v.AddArg4(dst, src, v0, mem)
		return true
	}
}
func rewriteValueSPARC64_OpNeg16(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Neg16 x)
	// result: (NEG x)
	for {
		x := v_0
		v.reset(OpSPARC64NEG)
		v.AddArg(x)
		return true
	}
}
func rewriteValueSPARC64_OpNeg32(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Neg32 x)
	// result: (NEG x)
	for {
		x := v_0
		v.reset(OpSPARC64NEG)
		v.AddArg(x)
		return true
	}
}
func rewriteValueSPARC64_OpNeg64(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Neg64 x)
	// result: (NEG x)
	for {
		x := v_0
		v.reset(OpSPARC64NEG)
		v.AddArg(x)
		return true
	}
}
func rewriteValueSPARC64_OpNeg8(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Neg8 x)
	// result: (NEG x)
	for {
		x := v_0
		v.reset(OpSPARC64NEG)
		v.AddArg(x)
		return true
	}
}
func rewriteValueSPARC64_OpNeq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq16 x y)
	// result: (NotEqual (CMP (ZeroExt16to64 x) (ZeroExt16to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpNeq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq32 x y)
	// result: (NotEqual (CMP (ZeroExt32to64 x) (ZeroExt32to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpNeq32F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Neq32F x y)
	// result: (FNotEqual (FCMPS x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64FNotEqual)
		v0 := b.NewValue0(v.Pos, OpSPARC64FCMPS, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpNeq64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Neq64 x y)
	// result: (NotEqual (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpSPARC64CMP, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpNeq64F(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Neq64F x y)
	// result: (FNotEqual (FCMPD x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64FNotEqual)
		v0 := b.NewValue0(v.Pos, OpSPARC64FCMPD, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpNeq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq8 x y)
	// result: (NotEqual (CMP (ZeroExt8to64 x) (ZeroExt8to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpNeqB(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (NeqB x y)
	// result: (NotEqual (CMP (ZeroExt8to64 x) (ZeroExt8to64 y)))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpNeqPtr(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (NeqPtr x y)
	// result: (NotEqual (CMP x y))
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64NotEqual)
		v0 := b.NewValue0(v.Pos, OpSPARC64CMP, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpNot(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Not x)
	// result: (XORconst <typ.Bool> [1] x)
	for {
		x := v_0
		v.reset(OpSPARC64XORconst)
		v.Type = typ.Bool
		v.AuxInt = int64ToAuxInt(1)
		v.AddArg(x)
		return true
	}
}
func rewriteValueSPARC64_OpOffPtr(v *Value) bool {
	v_0 := v.Args[0]
	// match: (OffPtr [off] ptr)
	// result: (ADDconst [off] ptr)
	for {
		off := auxIntToInt64(v.AuxInt)
		ptr := v_0
		v.reset(OpSPARC64ADDconst)
		v.AuxInt = int64ToAuxInt(off)
		v.AddArg(ptr)
		return true
	}
}
func rewriteValueSPARC64_OpRotateLeft32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft32 <t> x y)
	// result: (OR <t> (SLLW <t> x y) (SRLW <t> x (SUB <typ.UInt64> (MOVDconst <typ.UInt64> [32]) y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSPARC64OR)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpSPARC64SLLW, t)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpSPARC64SRLW, t)
		v2 := b.NewValue0(v.Pos, OpSPARC64SUB, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpSPARC64MOVDconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(32)
		v2.AddArg2(v3, y)
		v1.AddArg2(x, v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSPARC64_OpRotateLeft64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft64 <t> x y)
	// result: (OR <t> (SLLD <t> x y) (SRLD <t> x (SUB <typ.UInt64> (MOVDconst <typ.UInt64> [64]) y)))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSPARC64OR)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpSPARC64SLLD, t)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpSPARC64SRLD, t)
		v2 := b.NewValue0(v.Pos, OpSPARC64SUB, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpSPARC64MOVDconst, typ.UInt64)
		v3.AuxInt = int64ToAuxInt(64)
		v2.AddArg2(v3, y)
		v1.AddArg2(x, v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueSPARC64_OpRsh16Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux16 x y)
	// result: (Rsh16Ux64 x (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRsh16Ux64)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpRsh16Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux32 x y)
	// result: (Rsh16Ux64 x (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRsh16Ux64)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpRsh16Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux64 <t> x (Const64 [c]))
	// cond: uint64(c) < 64
	// result: (SRLDconst <t> [c] x)
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpSPARC64SRLDconst)
		v.Type = t
		v.AuxInt = int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (Rsh16Ux64 _ (Const64 [c]))
	// cond: uint64(c) >= 64
	// result: (MOVDconst [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpSPARC64MOVDconst)
		v.AuxInt = int64ToAuxInt(0)
		return true
	}
	// match: (Rsh16Ux64 <t> x y)
	// result: (AND (NEG <t> (LessThanU <typ.UInt64> (CMPconst <types.TypeFlags> [64] y))) (SRLD <t> x y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSPARC64AND)
		v0 := b.NewValue0(v.Pos, OpSPARC64NEG, t)
		v1 := b.NewValue0(v.Pos, OpSPARC64LessThanU, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpSPARC64CMPconst, types.TypeFlags)
		v2.AuxInt = int64ToAuxInt(64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, OpSPARC64SRLD, t)
		v3.AddArg2(x, y)
		v.AddArg2(v0, v3)
		return true
	}
}
func rewriteValueSPARC64_OpRsh16Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux8 x y)
	// result: (Rsh16Ux64 x (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRsh16Ux64)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpRsh16x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x16 x y)
	// result: (Rsh16x64 x (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRsh16x64)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpRsh16x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x32 x y)
	// result: (Rsh16x64 x (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRsh16x64)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpRsh16x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x64 <t> x (Const64 [c]))
	// cond: uint64(c) < 64
	// result: (SRADconst <t> [c] x)
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpSPARC64SRADconst)
		v.Type = t
		v.AuxInt = int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (Rsh16x64 <t> x (Const64 [c]))
	// cond: uint64(c) >= 64
	// result: (SRADconst <t> [63] x)
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpSPARC64SRADconst)
		v.Type = t
		v.AuxInt = int64ToAuxInt(63)
		v.AddArg(x)
		return true
	}
	// match: (Rsh16x64 <t> x y)
	// result: (SRAD <t> x (OR <typ.UInt64> (NEG <typ.UInt64> (GreaterThanU <typ.UInt64> (CMPconst <types.TypeFlags> [63] y))) y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSPARC64SRAD)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpSPARC64OR, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpSPARC64NEG, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpSPARC64GreaterThanU, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpSPARC64CMPconst, types.TypeFlags)
		v3.AuxInt = int64ToAuxInt(63)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg2(v1, y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpRsh16x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x8 x y)
	// result: (Rsh16x64 x (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRsh16x64)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpRsh32Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux16 x y)
	// result: (Rsh32Ux64 x (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRsh32Ux64)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpRsh32Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux32 x y)
	// result: (Rsh32Ux64 x (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRsh32Ux64)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpRsh32Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux64 <t> x (Const64 [c]))
	// cond: uint64(c) < 64
	// result: (SRLDconst <t> [c] x)
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpSPARC64SRLDconst)
		v.Type = t
		v.AuxInt = int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (Rsh32Ux64 _ (Const64 [c]))
	// cond: uint64(c) >= 64
	// result: (MOVDconst [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpSPARC64MOVDconst)
		v.AuxInt = int64ToAuxInt(0)
		return true
	}
	// match: (Rsh32Ux64 <t> x y)
	// result: (AND (NEG <t> (LessThanU <typ.UInt64> (CMPconst <types.TypeFlags> [64] y))) (SRLD <t> x y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSPARC64AND)
		v0 := b.NewValue0(v.Pos, OpSPARC64NEG, t)
		v1 := b.NewValue0(v.Pos, OpSPARC64LessThanU, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpSPARC64CMPconst, types.TypeFlags)
		v2.AuxInt = int64ToAuxInt(64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, OpSPARC64SRLD, t)
		v3.AddArg2(x, y)
		v.AddArg2(v0, v3)
		return true
	}
}
func rewriteValueSPARC64_OpRsh32Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux8 x y)
	// result: (Rsh32Ux64 x (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRsh32Ux64)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpRsh32x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x16 x y)
	// result: (Rsh32x64 x (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRsh32x64)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpRsh32x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x32 x y)
	// result: (Rsh32x64 x (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRsh32x64)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpRsh32x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x64 <t> x (Const64 [c]))
	// cond: uint64(c) < 64
	// result: (SRADconst <t> [c] x)
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpSPARC64SRADconst)
		v.Type = t
		v.AuxInt = int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (Rsh32x64 <t> x (Const64 [c]))
	// cond: uint64(c) >= 64
	// result: (SRADconst <t> [63] x)
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpSPARC64SRADconst)
		v.Type = t
		v.AuxInt = int64ToAuxInt(63)
		v.AddArg(x)
		return true
	}
	// match: (Rsh32x64 <t> x y)
	// result: (SRAD <t> x (OR <typ.UInt64> (NEG <typ.UInt64> (GreaterThanU <typ.UInt64> (CMPconst <types.TypeFlags> [63] y))) y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSPARC64SRAD)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpSPARC64OR, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpSPARC64NEG, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpSPARC64GreaterThanU, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpSPARC64CMPconst, types.TypeFlags)
		v3.AuxInt = int64ToAuxInt(63)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg2(v1, y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpRsh32x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x8 x y)
	// result: (Rsh32x64 x (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRsh32x64)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpRsh64Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux16 x y)
	// result: (Rsh64Ux64 x (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRsh64Ux64)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpRsh64Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux32 x y)
	// result: (Rsh64Ux64 x (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRsh64Ux64)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpRsh64Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux64 <t> x (Const64 [c]))
	// cond: uint64(c) < 64
	// result: (SRLDconst <t> [c] x)
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpSPARC64SRLDconst)
		v.Type = t
		v.AuxInt = int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (Rsh64Ux64 _ (Const64 [c]))
	// cond: uint64(c) >= 64
	// result: (MOVDconst [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpSPARC64MOVDconst)
		v.AuxInt = int64ToAuxInt(0)
		return true
	}
	// match: (Rsh64Ux64 <t> x y)
	// result: (AND (NEG <t> (LessThanU <typ.UInt64> (CMPconst <types.TypeFlags> [64] y))) (SRLD <t> x y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSPARC64AND)
		v0 := b.NewValue0(v.Pos, OpSPARC64NEG, t)
		v1 := b.NewValue0(v.Pos, OpSPARC64LessThanU, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpSPARC64CMPconst, types.TypeFlags)
		v2.AuxInt = int64ToAuxInt(64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, OpSPARC64SRLD, t)
		v3.AddArg2(x, y)
		v.AddArg2(v0, v3)
		return true
	}
}
func rewriteValueSPARC64_OpRsh64Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux8 x y)
	// result: (Rsh64Ux64 x (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRsh64Ux64)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpRsh64x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64x16 x y)
	// result: (Rsh64x64 x (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRsh64x64)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpRsh64x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64x32 x y)
	// result: (Rsh64x64 x (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRsh64x64)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpRsh64x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64x64 <t> x (Const64 [c]))
	// cond: uint64(c) < 64
	// result: (SRADconst <t> [c] x)
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpSPARC64SRADconst)
		v.Type = t
		v.AuxInt = int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (Rsh64x64 <t> x (Const64 [c]))
	// cond: uint64(c) >= 64
	// result: (SRADconst <t> [63] x)
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpSPARC64SRADconst)
		v.Type = t
		v.AuxInt = int64ToAuxInt(63)
		v.AddArg(x)
		return true
	}
	// match: (Rsh64x64 <t> x y)
	// result: (SRAD <t> x (OR <typ.UInt64> (NEG <typ.UInt64> (GreaterThanU <typ.UInt64> (CMPconst <types.TypeFlags> [63] y))) y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSPARC64SRAD)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpSPARC64OR, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpSPARC64NEG, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpSPARC64GreaterThanU, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpSPARC64CMPconst, types.TypeFlags)
		v3.AuxInt = int64ToAuxInt(63)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg2(v1, y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpRsh64x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64x8 x y)
	// result: (Rsh64x64 x (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRsh64x64)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpRsh8Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux16 x y)
	// result: (Rsh8Ux64 x (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRsh8Ux64)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpRsh8Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux32 x y)
	// result: (Rsh8Ux64 x (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRsh8Ux64)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpRsh8Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux64 <t> x (Const64 [c]))
	// cond: uint64(c) < 64
	// result: (SRLDconst <t> [c] x)
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpSPARC64SRLDconst)
		v.Type = t
		v.AuxInt = int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (Rsh8Ux64 _ (Const64 [c]))
	// cond: uint64(c) >= 64
	// result: (MOVDconst [0])
	for {
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpSPARC64MOVDconst)
		v.AuxInt = int64ToAuxInt(0)
		return true
	}
	// match: (Rsh8Ux64 <t> x y)
	// result: (AND (NEG <t> (LessThanU <typ.UInt64> (CMPconst <types.TypeFlags> [64] y))) (SRLD <t> x y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSPARC64AND)
		v0 := b.NewValue0(v.Pos, OpSPARC64NEG, t)
		v1 := b.NewValue0(v.Pos, OpSPARC64LessThanU, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpSPARC64CMPconst, types.TypeFlags)
		v2.AuxInt = int64ToAuxInt(64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, OpSPARC64SRLD, t)
		v3.AddArg2(x, y)
		v.AddArg2(v0, v3)
		return true
	}
}
func rewriteValueSPARC64_OpRsh8Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux8 x y)
	// result: (Rsh8Ux64 x (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRsh8Ux64)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpRsh8x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x16 x y)
	// result: (Rsh8x64 x (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRsh8x64)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpRsh8x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x32 x y)
	// result: (Rsh8x64 x (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRsh8x64)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpRsh8x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x64 <t> x (Const64 [c]))
	// cond: uint64(c) < 64
	// result: (SRADconst <t> [c] x)
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpSPARC64SRADconst)
		v.Type = t
		v.AuxInt = int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (Rsh8x64 <t> x (Const64 [c]))
	// cond: uint64(c) >= 64
	// result: (SRADconst <t> [63] x)
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpConst64 {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpSPARC64SRADconst)
		v.Type = t
		v.AuxInt = int64ToAuxInt(63)
		v.AddArg(x)
		return true
	}
	// match: (Rsh8x64 <t> x y)
	// result: (SRAD <t> x (OR <typ.UInt64> (NEG <typ.UInt64> (GreaterThanU <typ.UInt64> (CMPconst <types.TypeFlags> [63] y))) y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.reset(OpSPARC64SRAD)
		v.Type = t
		v0 := b.NewValue0(v.Pos, OpSPARC64OR, typ.UInt64)
		v1 := b.NewValue0(v.Pos, OpSPARC64NEG, typ.UInt64)
		v2 := b.NewValue0(v.Pos, OpSPARC64GreaterThanU, typ.UInt64)
		v3 := b.NewValue0(v.Pos, OpSPARC64CMPconst, types.TypeFlags)
		v3.AuxInt = int64ToAuxInt(63)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg2(v1, y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpRsh8x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x8 x y)
	// result: (Rsh8x64 x (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpRsh8x64)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueSPARC64_OpSPARC64ADD(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ADD (MOVDconst [c]) x)
	// cond: is13Bit(c)
	// result: (ADDconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpSPARC64MOVDconst {
				continue
			}
			c := auxIntToInt64(v_0.AuxInt)
			x := v_1
			if !(is13Bit(c)) {
				continue
			}
			v.reset(OpSPARC64ADDconst)
			v.AuxInt = int64ToAuxInt(c)
			v.AddArg(x)
			return true
		}
		break
	}
	return false
}
func rewriteValueSPARC64_OpSPARC64ADDconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ADDconst [0] x)
	// result: x
	for {
		if auxIntToInt64(v.AuxInt) != 0 {
			break
		}
		x := v_0
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteValueSPARC64_OpSPARC64AND(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (AND (MOVDconst [c]) x)
	// cond: is13Bit(c)
	// result: (ANDconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpSPARC64MOVDconst {
				continue
			}
			c := auxIntToInt64(v_0.AuxInt)
			x := v_1
			if !(is13Bit(c)) {
				continue
			}
			v.reset(OpSPARC64ANDconst)
			v.AuxInt = int64ToAuxInt(c)
			v.AddArg(x)
			return true
		}
		break
	}
	return false
}
func rewriteValueSPARC64_OpSPARC64ANDconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ANDconst [-1] x)
	// result: x
	for {
		if auxIntToInt64(v.AuxInt) != -1 {
			break
		}
		x := v_0
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteValueSPARC64_OpSPARC64CMP(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (CMP x (MOVDconst [c]))
	// cond: is13Bit(c)
	// result: (CMPconst [c] x)
	for {
		x := v_0
		if v_1.Op != OpSPARC64MOVDconst {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(is13Bit(c)) {
			break
		}
		v.reset(OpSPARC64CMPconst)
		v.AuxInt = int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueSPARC64_OpSPARC64LoweredPanicBoundsCR(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (LoweredPanicBoundsCR [kind] {p} (MOVDconst [c]) mem)
	// result: (LoweredPanicBoundsCC [kind] {PanicBoundsCC{Cx:p.C, Cy:c}} mem)
	for {
		kind := auxIntToInt64(v.AuxInt)
		p := auxToPanicBoundsC(v.Aux)
		if v_0.Op != OpSPARC64MOVDconst {
			break
		}
		c := auxIntToInt64(v_0.AuxInt)
		mem := v_1
		v.reset(OpSPARC64LoweredPanicBoundsCC)
		v.AuxInt = int64ToAuxInt(kind)
		v.Aux = panicBoundsCCToAux(PanicBoundsCC{Cx: p.C, Cy: c})
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueSPARC64_OpSPARC64LoweredPanicBoundsRC(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (LoweredPanicBoundsRC [kind] {p} (MOVDconst [c]) mem)
	// result: (LoweredPanicBoundsCC [kind] {PanicBoundsCC{Cx:c, Cy:p.C}} mem)
	for {
		kind := auxIntToInt64(v.AuxInt)
		p := auxToPanicBoundsC(v.Aux)
		if v_0.Op != OpSPARC64MOVDconst {
			break
		}
		c := auxIntToInt64(v_0.AuxInt)
		mem := v_1
		v.reset(OpSPARC64LoweredPanicBoundsCC)
		v.AuxInt = int64ToAuxInt(kind)
		v.Aux = panicBoundsCCToAux(PanicBoundsCC{Cx: c, Cy: p.C})
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValueSPARC64_OpSPARC64LoweredPanicBoundsRR(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (LoweredPanicBoundsRR [kind] x (MOVDconst [c]) mem)
	// result: (LoweredPanicBoundsRC [kind] x {PanicBoundsC{C:c}} mem)
	for {
		kind := auxIntToInt64(v.AuxInt)
		x := v_0
		if v_1.Op != OpSPARC64MOVDconst {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		mem := v_2
		v.reset(OpSPARC64LoweredPanicBoundsRC)
		v.AuxInt = int64ToAuxInt(kind)
		v.Aux = panicBoundsCToAux(PanicBoundsC{C: c})
		v.AddArg2(x, mem)
		return true
	}
	// match: (LoweredPanicBoundsRR [kind] (MOVDconst [c]) y mem)
	// result: (LoweredPanicBoundsCR [kind] {PanicBoundsC{C:c}} y mem)
	for {
		kind := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpSPARC64MOVDconst {
			break
		}
		c := auxIntToInt64(v_0.AuxInt)
		y := v_1
		mem := v_2
		v.reset(OpSPARC64LoweredPanicBoundsCR)
		v.AuxInt = int64ToAuxInt(kind)
		v.Aux = panicBoundsCToAux(PanicBoundsC{C: c})
		v.AddArg2(y, mem)
		return true
	}
	return false
}
func rewriteValueSPARC64_OpSPARC64MOVB(v *Value) bool {
	v_0 := v.Args[0]
	// match: (MOVB (MOVDconst [c]))
	// result: (MOVDconst [int64(int8(c))])
	for {
		if v_0.Op != OpSPARC64MOVDconst {
			break
		}
		c := auxIntToInt64(v_0.AuxInt)
		v.reset(OpSPARC64MOVDconst)
		v.AuxInt = int64ToAuxInt(int64(int8(c)))
		return true
	}
	return false
}
func rewriteValueSPARC64_OpSPARC64MOVD(v *Value) bool {
	v_0 := v.Args[0]
	// match: (MOVD (MOVDconst [c]))
	// result: (MOVDconst [c])
	for {
		if v_0.Op != OpSPARC64MOVDconst {
			break
		}
		c := auxIntToInt64(v_0.AuxInt)
		v.reset(OpSPARC64MOVDconst)
		v.AuxInt = int64ToAuxInt(c)
		return true
	}
	return false
}
func rewriteValueSPARC64_OpSPARC64MOVH(v *Value) bool {
	v_0 := v.Args[0]
	// match: (MOVH (MOVDconst [c]))
	// result: (MOVDconst [int64(int16(c))])
	for {
		if v_0.Op != OpSPARC64MOVDconst {
			break
		}
		c := auxIntToInt64(v_0.AuxInt)
		v.reset(OpSPARC64MOVDconst)
		v.AuxInt = int64ToAuxInt(int64(int16(c)))
		return true
	}
	return false
}
func rewriteValueSPARC64_OpSPARC64MOVUB(v *Value) bool {
	v_0 := v.Args[0]
	// match: (MOVUB (MOVDconst [c]))
	// result: (MOVDconst [int64(uint8(c))])
	for {
		if v_0.Op != OpSPARC64MOVDconst {
			break
		}
		c := auxIntToInt64(v_0.AuxInt)
		v.reset(OpSPARC64MOVDconst)
		v.AuxInt = int64ToAuxInt(int64(uint8(c)))
		return true
	}
	return false
}
func rewriteValueSPARC64_OpSPARC64MOVUH(v *Value) bool {
	v_0 := v.Args[0]
	// match: (MOVUH (MOVDconst [c]))
	// result: (MOVDconst [int64(uint16(c))])
	for {
		if v_0.Op != OpSPARC64MOVDconst {
			break
		}
		c := auxIntToInt64(v_0.AuxInt)
		v.reset(OpSPARC64MOVDconst)
		v.AuxInt = int64ToAuxInt(int64(uint16(c)))
		return true
	}
	return false
}
func rewriteValueSPARC64_OpSPARC64MOVUW(v *Value) bool {
	v_0 := v.Args[0]
	// match: (MOVUW (MOVDconst [c]))
	// result: (MOVDconst [int64(uint32(c))])
	for {
		if v_0.Op != OpSPARC64MOVDconst {
			break
		}
		c := auxIntToInt64(v_0.AuxInt)
		v.reset(OpSPARC64MOVDconst)
		v.AuxInt = int64ToAuxInt(int64(uint32(c)))
		return true
	}
	return false
}
func rewriteValueSPARC64_OpSPARC64MOVW(v *Value) bool {
	v_0 := v.Args[0]
	// match: (MOVW (MOVDconst [c]))
	// result: (MOVDconst [int64(int32(c))])
	for {
		if v_0.Op != OpSPARC64MOVDconst {
			break
		}
		c := auxIntToInt64(v_0.AuxInt)
		v.reset(OpSPARC64MOVDconst)
		v.AuxInt = int64ToAuxInt(int64(int32(c)))
		return true
	}
	return false
}
func rewriteValueSPARC64_OpSPARC64OR(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (OR (MOVDconst [c]) x)
	// cond: is13Bit(c)
	// result: (ORconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpSPARC64MOVDconst {
				continue
			}
			c := auxIntToInt64(v_0.AuxInt)
			x := v_1
			if !(is13Bit(c)) {
				continue
			}
			v.reset(OpSPARC64ORconst)
			v.AuxInt = int64ToAuxInt(c)
			v.AddArg(x)
			return true
		}
		break
	}
	return false
}
func rewriteValueSPARC64_OpSPARC64ORconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (ORconst [0] x)
	// result: x
	for {
		if auxIntToInt64(v.AuxInt) != 0 {
			break
		}
		x := v_0
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteValueSPARC64_OpSPARC64SLLD(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SLLD x (MOVDconst [c]))
	// result: (SLLDconst [c&63] x)
	for {
		x := v_0
		if v_1.Op != OpSPARC64MOVDconst {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		v.reset(OpSPARC64SLLDconst)
		v.AuxInt = int64ToAuxInt(c & 63)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueSPARC64_OpSPARC64SRAD(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SRAD x (MOVDconst [c]))
	// result: (SRADconst [c&63] x)
	for {
		x := v_0
		if v_1.Op != OpSPARC64MOVDconst {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		v.reset(OpSPARC64SRADconst)
		v.AuxInt = int64ToAuxInt(c & 63)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueSPARC64_OpSPARC64SRLD(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SRLD x (MOVDconst [c]))
	// result: (SRLDconst [c&63] x)
	for {
		x := v_0
		if v_1.Op != OpSPARC64MOVDconst {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		v.reset(OpSPARC64SRLDconst)
		v.AuxInt = int64ToAuxInt(c & 63)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueSPARC64_OpSPARC64SUB(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SUB x (MOVDconst [c]))
	// cond: is13Bit(c)
	// result: (SUBconst [c] x)
	for {
		x := v_0
		if v_1.Op != OpSPARC64MOVDconst {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(is13Bit(c)) {
			break
		}
		v.reset(OpSPARC64SUBconst)
		v.AuxInt = int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueSPARC64_OpSPARC64SUBconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (SUBconst [0] x)
	// result: x
	for {
		if auxIntToInt64(v.AuxInt) != 0 {
			break
		}
		x := v_0
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteValueSPARC64_OpSPARC64XOR(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (XOR (MOVDconst [c]) x)
	// cond: is13Bit(c)
	// result: (XORconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != OpSPARC64MOVDconst {
				continue
			}
			c := auxIntToInt64(v_0.AuxInt)
			x := v_1
			if !(is13Bit(c)) {
				continue
			}
			v.reset(OpSPARC64XORconst)
			v.AuxInt = int64ToAuxInt(c)
			v.AddArg(x)
			return true
		}
		break
	}
	return false
}
func rewriteValueSPARC64_OpSPARC64XORconst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (XORconst [0] x)
	// result: x
	for {
		if auxIntToInt64(v.AuxInt) != 0 {
			break
		}
		x := v_0
		v.copyOf(x)
		return true
	}
	return false
}
func rewriteValueSPARC64_OpSlicemask(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (Slicemask <t> x)
	// result: (SRADconst <t> [63] (NEG <t> x))
	for {
		t := v.Type
		x := v_0
		v.reset(OpSPARC64SRADconst)
		v.Type = t
		v.AuxInt = int64ToAuxInt(63)
		v0 := b.NewValue0(v.Pos, OpSPARC64NEG, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueSPARC64_OpStore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Store {t} ptr val mem)
	// cond: t.Size() == 1
	// result: (MOVBstore ptr val mem)
	for {
		t := auxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.Size() == 1) {
			break
		}
		v.reset(OpSPARC64MOVBstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.Size() == 2
	// result: (MOVHstore ptr val mem)
	for {
		t := auxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.Size() == 2) {
			break
		}
		v.reset(OpSPARC64MOVHstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.Size() == 4 && !t.IsFloat()
	// result: (MOVWstore ptr val mem)
	for {
		t := auxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.Size() == 4 && !t.IsFloat()) {
			break
		}
		v.reset(OpSPARC64MOVWstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.Size() == 8 && !t.IsFloat()
	// result: (MOVDstore ptr val mem)
	for {
		t := auxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.Size() == 8 && !t.IsFloat()) {
			break
		}
		v.reset(OpSPARC64MOVDstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.Size() == 4 && t.IsFloat()
	// result: (FMOVSstore ptr val mem)
	for {
		t := auxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.Size() == 4 && t.IsFloat()) {
			break
		}
		v.reset(OpSPARC64FMOVSstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.Size() == 8 && t.IsFloat()
	// result: (FMOVDstore ptr val mem)
	for {
		t := auxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.Size() == 8 && t.IsFloat()) {
			break
		}
		v.reset(OpSPARC64FMOVDstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValueSPARC64_OpZero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	typ := &b.Func.Config.Types
	// match: (Zero [0] _ mem)
	// result: mem
	for {
		if auxIntToInt64(v.AuxInt) != 0 {
			break
		}
		mem := v_1
		v.copyOf(mem)
		return true
	}
	// match: (Zero [1] ptr mem)
	// result: (MOVBstore ptr (MOVDconst [0]) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 1 {
			break
		}
		ptr := v_0
		mem := v_1
		v.reset(OpSPARC64MOVBstore)
		v0 := b.NewValue0(v.Pos, OpSPARC64MOVDconst, typ.UInt64)
		v0.AuxInt = int64ToAuxInt(0)
		v.AddArg3(ptr, v0, mem)
		return true
	}
	// match: (Zero [2] {t} ptr mem)
	// cond: t.Alignment()%2 == 0
	// result: (MOVHstore ptr (MOVDconst [0]) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 2 {
			break
		}
		t := auxToType(v.Aux)
		ptr := v_0
		mem := v_1
		if !(t.Alignment()%2 == 0) {
			break
		}
		v.reset(OpSPARC64MOVHstore)
		v0 := b.NewValue0(v.Pos, OpSPARC64MOVDconst, typ.UInt64)
		v0.AuxInt = int64ToAuxInt(0)
		v.AddArg3(ptr, v0, mem)
		return true
	}
	// match: (Zero [4] {t} ptr mem)
	// cond: t.Alignment()%4 == 0
	// result: (MOVWstore ptr (MOVDconst [0]) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 4 {
			break
		}
		t := auxToType(v.Aux)
		ptr := v_0
		mem := v_1
		if !(t.Alignment()%4 == 0) {
			break
		}
		v.reset(OpSPARC64MOVWstore)
		v0 := b.NewValue0(v.Pos, OpSPARC64MOVDconst, typ.UInt64)
		v0.AuxInt = int64ToAuxInt(0)
		v.AddArg3(ptr, v0, mem)
		return true
	}
	// match: (Zero [8] {t} ptr mem)
	// cond: t.Alignment()%8 == 0
	// result: (MOVDstore ptr (MOVDconst [0]) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 8 {
			break
		}
		t := auxToType(v.Aux)
		ptr := v_0
		mem := v_1
		if !(t.Alignment()%8 == 0) {
			break
		}
		v.reset(OpSPARC64MOVDstore)
		v0 := b.NewValue0(v.Pos, OpSPARC64MOVDconst, typ.UInt64)
		v0.AuxInt = int64ToAuxInt(0)
		v.AddArg3(ptr, v0, mem)
		return true
	}
	// match: (Zero [s] {t} ptr mem)
	// result: (LoweredZero [t.Alignment()] ptr (ADDconst <ptr.Type> [s-moveSize(t.Alignment(), config)] ptr) mem)
	for {
		s := auxIntToInt64(v.AuxInt)
		t := auxToType(v.Aux)
		ptr := v_0
		mem := v_1
		v.reset(OpSPARC64LoweredZero)
		v.AuxInt = int64ToAuxInt(t.Alignment())
		v0 := b.NewValue0(v.Pos, OpSPARC64ADDconst, ptr.Type)
		v0.AuxInt = int64ToAuxInt(s - moveSize(t.Alignment(), config))
		v0.AddArg(ptr)
		v.AddArg3(ptr, v0, mem)
		return true
	}
}
func rewriteBlockSPARC64(b *Block) bool {
	switch b.Kind {
	case block.BlockIf:
		// match: (If (Equal cc) yes no)
		// result: (EQ cc yes no)
		for b.Controls[0].Op == OpSPARC64Equal {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.resetWithControl(block.BlockSPARC64EQ, cc)
			return true
		}
		// match: (If (NotEqual cc) yes no)
		// result: (NE cc yes no)
		for b.Controls[0].Op == OpSPARC64NotEqual {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.resetWithControl(block.BlockSPARC64NE, cc)
			return true
		}
		// match: (If (LessThan cc) yes no)
		// result: (LT cc yes no)
		for b.Controls[0].Op == OpSPARC64LessThan {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.resetWithControl(block.BlockSPARC64LT, cc)
			return true
		}
		// match: (If (LessEqual cc) yes no)
		// result: (LE cc yes no)
		for b.Controls[0].Op == OpSPARC64LessEqual {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.resetWithControl(block.BlockSPARC64LE, cc)
			return true
		}
		// match: (If (GreaterThan cc) yes no)
		// result: (GT cc yes no)
		for b.Controls[0].Op == OpSPARC64GreaterThan {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.resetWithControl(block.BlockSPARC64GT, cc)
			return true
		}
		// match: (If (GreaterEqual cc) yes no)
		// result: (GE cc yes no)
		for b.Controls[0].Op == OpSPARC64GreaterEqual {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.resetWithControl(block.BlockSPARC64GE, cc)
			return true
		}
		// match: (If (LessThanU cc) yes no)
		// result: (ULT cc yes no)
		for b.Controls[0].Op == OpSPARC64LessThanU {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.resetWithControl(block.BlockSPARC64ULT, cc)
			return true
		}
		// match: (If (LessEqualU cc) yes no)
		// result: (ULE cc yes no)
		for b.Controls[0].Op == OpSPARC64LessEqualU {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.resetWithControl(block.BlockSPARC64ULE, cc)
			return true
		}
		// match: (If (GreaterThanU cc) yes no)
		// result: (UGT cc yes no)
		for b.Controls[0].Op == OpSPARC64GreaterThanU {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.resetWithControl(block.BlockSPARC64UGT, cc)
			return true
		}
		// match: (If (GreaterEqualU cc) yes no)
		// result: (UGE cc yes no)
		for b.Controls[0].Op == OpSPARC64GreaterEqualU {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.resetWithControl(block.BlockSPARC64UGE, cc)
			return true
		}
		// match: (If cond yes no)
		// result: (NE (CMPconst [0] cond) yes no)
		for {
			cond := b.Controls[0]
			v0 := b.NewValue0(cond.Pos, OpSPARC64CMPconst, types.TypeFlags)
			v0.AuxInt = int64ToAuxInt(0)
			v0.AddArg(cond)
			b.resetWithControl(block.BlockSPARC64NE, v0)
			return true
		}
	}
	return false
}
