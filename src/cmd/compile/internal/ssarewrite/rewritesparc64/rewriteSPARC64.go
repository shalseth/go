// Code generated from _gen/SPARC64.rules using 'go generate'; DO NOT EDIT.

package rewritesparc64

import "cmd/compile/internal/types"
import "cmd/compile/internal/ssa/block"
import "cmd/compile/internal/ssa/ssaop"
import "cmd/compile/internal/ssa"

func RewriteValue(v *ssa.Value) bool {
	switch v.Op {
	case ssaop.OpAbs:
		v.Op = ssaop.OpSPARC64FABSD
		return true
	case ssaop.OpAdd16:
		v.Op = ssaop.OpSPARC64ADD
		return true
	case ssaop.OpAdd32:
		v.Op = ssaop.OpSPARC64ADD
		return true
	case ssaop.OpAdd32F:
		v.Op = ssaop.OpSPARC64FADDS
		return true
	case ssaop.OpAdd64:
		v.Op = ssaop.OpSPARC64ADD
		return true
	case ssaop.OpAdd64F:
		v.Op = ssaop.OpSPARC64FADDD
		return true
	case ssaop.OpAdd64carry:
		v.Op = ssaop.OpSPARC64ADDCARRY
		return true
	case ssaop.OpAdd8:
		v.Op = ssaop.OpSPARC64ADD
		return true
	case ssaop.OpAddPtr:
		v.Op = ssaop.OpSPARC64ADD
		return true
	case ssaop.OpAddr:
		return rewriteValue_OpAddr(v)
	case ssaop.OpAnd16:
		v.Op = ssaop.OpSPARC64AND
		return true
	case ssaop.OpAnd32:
		v.Op = ssaop.OpSPARC64AND
		return true
	case ssaop.OpAnd64:
		v.Op = ssaop.OpSPARC64AND
		return true
	case ssaop.OpAnd8:
		v.Op = ssaop.OpSPARC64AND
		return true
	case ssaop.OpAndB:
		v.Op = ssaop.OpSPARC64AND
		return true
	case ssaop.OpAtomicAdd32:
		v.Op = ssaop.OpSPARC64LoweredAtomicAdd32
		return true
	case ssaop.OpAtomicAdd64:
		v.Op = ssaop.OpSPARC64LoweredAtomicAdd64
		return true
	case ssaop.OpAtomicAnd32:
		v.Op = ssaop.OpSPARC64LoweredAtomicAnd32
		return true
	case ssaop.OpAtomicAnd8:
		v.Op = ssaop.OpSPARC64LoweredAtomicAnd8
		return true
	case ssaop.OpAtomicCompareAndSwap32:
		v.Op = ssaop.OpSPARC64LoweredAtomicCas32
		return true
	case ssaop.OpAtomicCompareAndSwap64:
		v.Op = ssaop.OpSPARC64LoweredAtomicCas64
		return true
	case ssaop.OpAtomicExchange32:
		v.Op = ssaop.OpSPARC64LoweredAtomicExchange32
		return true
	case ssaop.OpAtomicExchange64:
		v.Op = ssaop.OpSPARC64LoweredAtomicExchange64
		return true
	case ssaop.OpAtomicLoad32:
		v.Op = ssaop.OpSPARC64LoweredAtomicLoad32
		return true
	case ssaop.OpAtomicLoad64:
		v.Op = ssaop.OpSPARC64LoweredAtomicLoad64
		return true
	case ssaop.OpAtomicLoad8:
		v.Op = ssaop.OpSPARC64LoweredAtomicLoad8
		return true
	case ssaop.OpAtomicLoadPtr:
		v.Op = ssaop.OpSPARC64LoweredAtomicLoad64
		return true
	case ssaop.OpAtomicOr32:
		v.Op = ssaop.OpSPARC64LoweredAtomicOr32
		return true
	case ssaop.OpAtomicOr8:
		v.Op = ssaop.OpSPARC64LoweredAtomicOr8
		return true
	case ssaop.OpAtomicStore32:
		v.Op = ssaop.OpSPARC64LoweredAtomicStore32
		return true
	case ssaop.OpAtomicStore64:
		v.Op = ssaop.OpSPARC64LoweredAtomicStore64
		return true
	case ssaop.OpAtomicStore8:
		v.Op = ssaop.OpSPARC64LoweredAtomicStore8
		return true
	case ssaop.OpAtomicStorePtrNoWB:
		v.Op = ssaop.OpSPARC64LoweredAtomicStore64
		return true
	case ssaop.OpAvg64u:
		return rewriteValue_OpAvg64u(v)
	case ssaop.OpClosureCall:
		v.Op = ssaop.OpSPARC64CALLclosure
		return true
	case ssaop.OpCom16:
		return rewriteValue_OpCom16(v)
	case ssaop.OpCom32:
		return rewriteValue_OpCom32(v)
	case ssaop.OpCom64:
		return rewriteValue_OpCom64(v)
	case ssaop.OpCom8:
		return rewriteValue_OpCom8(v)
	case ssaop.OpConst16:
		return rewriteValue_OpConst16(v)
	case ssaop.OpConst32:
		return rewriteValue_OpConst32(v)
	case ssaop.OpConst32F:
		v.Op = ssaop.OpSPARC64FMOVSconst
		return true
	case ssaop.OpConst64:
		return rewriteValue_OpConst64(v)
	case ssaop.OpConst64F:
		v.Op = ssaop.OpSPARC64FMOVDconst
		return true
	case ssaop.OpConst8:
		return rewriteValue_OpConst8(v)
	case ssaop.OpConstBool:
		return rewriteValue_OpConstBool(v)
	case ssaop.OpConstNil:
		return rewriteValue_OpConstNil(v)
	case ssaop.OpCtz16NonZero:
		v.Op = ssaop.OpCtz64
		return true
	case ssaop.OpCtz32:
		return rewriteValue_OpCtz32(v)
	case ssaop.OpCtz32NonZero:
		v.Op = ssaop.OpCtz64
		return true
	case ssaop.OpCtz64:
		return rewriteValue_OpCtz64(v)
	case ssaop.OpCtz64NonZero:
		v.Op = ssaop.OpCtz64
		return true
	case ssaop.OpCtz8NonZero:
		v.Op = ssaop.OpCtz64
		return true
	case ssaop.OpCvt32Fto32:
		return rewriteValue_OpCvt32Fto32(v)
	case ssaop.OpCvt32Fto64:
		v.Op = ssaop.OpSPARC64FSTOX
		return true
	case ssaop.OpCvt32Fto64F:
		v.Op = ssaop.OpSPARC64FSTOD
		return true
	case ssaop.OpCvt32to32F:
		return rewriteValue_OpCvt32to32F(v)
	case ssaop.OpCvt32to64F:
		return rewriteValue_OpCvt32to64F(v)
	case ssaop.OpCvt64Fto32:
		return rewriteValue_OpCvt64Fto32(v)
	case ssaop.OpCvt64Fto32F:
		v.Op = ssaop.OpSPARC64FDTOS
		return true
	case ssaop.OpCvt64Fto64:
		v.Op = ssaop.OpSPARC64FDTOX
		return true
	case ssaop.OpCvt64to32F:
		v.Op = ssaop.OpSPARC64FXTOS
		return true
	case ssaop.OpCvt64to64F:
		v.Op = ssaop.OpSPARC64FXTOD
		return true
	case ssaop.OpCvtBoolToUint8:
		v.Op = ssaop.OpCopy
		return true
	case ssaop.OpDiv16:
		return rewriteValue_OpDiv16(v)
	case ssaop.OpDiv16u:
		return rewriteValue_OpDiv16u(v)
	case ssaop.OpDiv32:
		return rewriteValue_OpDiv32(v)
	case ssaop.OpDiv32F:
		v.Op = ssaop.OpSPARC64FDIVS
		return true
	case ssaop.OpDiv32u:
		return rewriteValue_OpDiv32u(v)
	case ssaop.OpDiv64:
		return rewriteValue_OpDiv64(v)
	case ssaop.OpDiv64F:
		v.Op = ssaop.OpSPARC64FDIVD
		return true
	case ssaop.OpDiv64u:
		return rewriteValue_OpDiv64u(v)
	case ssaop.OpDiv8:
		return rewriteValue_OpDiv8(v)
	case ssaop.OpDiv8u:
		return rewriteValue_OpDiv8u(v)
	case ssaop.OpEq16:
		return rewriteValue_OpEq16(v)
	case ssaop.OpEq32:
		return rewriteValue_OpEq32(v)
	case ssaop.OpEq32F:
		return rewriteValue_OpEq32F(v)
	case ssaop.OpEq64:
		return rewriteValue_OpEq64(v)
	case ssaop.OpEq64F:
		return rewriteValue_OpEq64F(v)
	case ssaop.OpEq8:
		return rewriteValue_OpEq8(v)
	case ssaop.OpEqB:
		return rewriteValue_OpEqB(v)
	case ssaop.OpEqPtr:
		return rewriteValue_OpEqPtr(v)
	case ssaop.OpGetCallerPC:
		v.Op = ssaop.OpSPARC64LoweredGetCallerPC
		return true
	case ssaop.OpGetCallerSP:
		v.Op = ssaop.OpSPARC64LoweredGetCallerSP
		return true
	case ssaop.OpGetClosurePtr:
		v.Op = ssaop.OpSPARC64LoweredGetClosurePtr
		return true
	case ssaop.OpHmul32:
		return rewriteValue_OpHmul32(v)
	case ssaop.OpHmul32u:
		return rewriteValue_OpHmul32u(v)
	case ssaop.OpHmul64:
		return rewriteValue_OpHmul64(v)
	case ssaop.OpHmul64u:
		return rewriteValue_OpHmul64u(v)
	case ssaop.OpInterCall:
		v.Op = ssaop.OpSPARC64CALLinter
		return true
	case ssaop.OpIsInBounds:
		return rewriteValue_OpIsInBounds(v)
	case ssaop.OpIsNonNil:
		return rewriteValue_OpIsNonNil(v)
	case ssaop.OpIsSliceInBounds:
		return rewriteValue_OpIsSliceInBounds(v)
	case ssaop.OpLeq16:
		return rewriteValue_OpLeq16(v)
	case ssaop.OpLeq16U:
		return rewriteValue_OpLeq16U(v)
	case ssaop.OpLeq32:
		return rewriteValue_OpLeq32(v)
	case ssaop.OpLeq32F:
		return rewriteValue_OpLeq32F(v)
	case ssaop.OpLeq32U:
		return rewriteValue_OpLeq32U(v)
	case ssaop.OpLeq64:
		return rewriteValue_OpLeq64(v)
	case ssaop.OpLeq64F:
		return rewriteValue_OpLeq64F(v)
	case ssaop.OpLeq64U:
		return rewriteValue_OpLeq64U(v)
	case ssaop.OpLeq8:
		return rewriteValue_OpLeq8(v)
	case ssaop.OpLeq8U:
		return rewriteValue_OpLeq8U(v)
	case ssaop.OpLess16:
		return rewriteValue_OpLess16(v)
	case ssaop.OpLess16U:
		return rewriteValue_OpLess16U(v)
	case ssaop.OpLess32:
		return rewriteValue_OpLess32(v)
	case ssaop.OpLess32F:
		return rewriteValue_OpLess32F(v)
	case ssaop.OpLess32U:
		return rewriteValue_OpLess32U(v)
	case ssaop.OpLess64:
		return rewriteValue_OpLess64(v)
	case ssaop.OpLess64F:
		return rewriteValue_OpLess64F(v)
	case ssaop.OpLess64U:
		return rewriteValue_OpLess64U(v)
	case ssaop.OpLess8:
		return rewriteValue_OpLess8(v)
	case ssaop.OpLess8U:
		return rewriteValue_OpLess8U(v)
	case ssaop.OpLoad:
		return rewriteValue_OpLoad(v)
	case ssaop.OpLocalAddr:
		return rewriteValue_OpLocalAddr(v)
	case ssaop.OpLsh16x16:
		return rewriteValue_OpLsh16x16(v)
	case ssaop.OpLsh16x32:
		return rewriteValue_OpLsh16x32(v)
	case ssaop.OpLsh16x64:
		return rewriteValue_OpLsh16x64(v)
	case ssaop.OpLsh16x8:
		return rewriteValue_OpLsh16x8(v)
	case ssaop.OpLsh32x16:
		return rewriteValue_OpLsh32x16(v)
	case ssaop.OpLsh32x32:
		return rewriteValue_OpLsh32x32(v)
	case ssaop.OpLsh32x64:
		return rewriteValue_OpLsh32x64(v)
	case ssaop.OpLsh32x8:
		return rewriteValue_OpLsh32x8(v)
	case ssaop.OpLsh64x16:
		return rewriteValue_OpLsh64x16(v)
	case ssaop.OpLsh64x32:
		return rewriteValue_OpLsh64x32(v)
	case ssaop.OpLsh64x64:
		return rewriteValue_OpLsh64x64(v)
	case ssaop.OpLsh64x8:
		return rewriteValue_OpLsh64x8(v)
	case ssaop.OpLsh8x16:
		return rewriteValue_OpLsh8x16(v)
	case ssaop.OpLsh8x32:
		return rewriteValue_OpLsh8x32(v)
	case ssaop.OpLsh8x64:
		return rewriteValue_OpLsh8x64(v)
	case ssaop.OpLsh8x8:
		return rewriteValue_OpLsh8x8(v)
	case ssaop.OpMod16:
		return rewriteValue_OpMod16(v)
	case ssaop.OpMod16u:
		return rewriteValue_OpMod16u(v)
	case ssaop.OpMod32:
		return rewriteValue_OpMod32(v)
	case ssaop.OpMod32u:
		return rewriteValue_OpMod32u(v)
	case ssaop.OpMod64:
		return rewriteValue_OpMod64(v)
	case ssaop.OpMod64u:
		return rewriteValue_OpMod64u(v)
	case ssaop.OpMod8:
		return rewriteValue_OpMod8(v)
	case ssaop.OpMod8u:
		return rewriteValue_OpMod8u(v)
	case ssaop.OpMove:
		return rewriteValue_OpMove(v)
	case ssaop.OpMul16:
		v.Op = ssaop.OpSPARC64MULD
		return true
	case ssaop.OpMul32:
		v.Op = ssaop.OpSPARC64MULD
		return true
	case ssaop.OpMul32F:
		v.Op = ssaop.OpSPARC64FMULS
		return true
	case ssaop.OpMul64:
		v.Op = ssaop.OpSPARC64MULD
		return true
	case ssaop.OpMul64F:
		v.Op = ssaop.OpSPARC64FMULD
		return true
	case ssaop.OpMul64uhilo:
		v.Op = ssaop.OpSPARC64MULDU
		return true
	case ssaop.OpMul8:
		v.Op = ssaop.OpSPARC64MULD
		return true
	case ssaop.OpNeg16:
		return rewriteValue_OpNeg16(v)
	case ssaop.OpNeg32:
		return rewriteValue_OpNeg32(v)
	case ssaop.OpNeg32F:
		v.Op = ssaop.OpSPARC64FNEGS
		return true
	case ssaop.OpNeg64:
		return rewriteValue_OpNeg64(v)
	case ssaop.OpNeg64F:
		v.Op = ssaop.OpSPARC64FNEGD
		return true
	case ssaop.OpNeg8:
		return rewriteValue_OpNeg8(v)
	case ssaop.OpNeq16:
		return rewriteValue_OpNeq16(v)
	case ssaop.OpNeq32:
		return rewriteValue_OpNeq32(v)
	case ssaop.OpNeq32F:
		return rewriteValue_OpNeq32F(v)
	case ssaop.OpNeq64:
		return rewriteValue_OpNeq64(v)
	case ssaop.OpNeq64F:
		return rewriteValue_OpNeq64F(v)
	case ssaop.OpNeq8:
		return rewriteValue_OpNeq8(v)
	case ssaop.OpNeqB:
		return rewriteValue_OpNeqB(v)
	case ssaop.OpNeqPtr:
		return rewriteValue_OpNeqPtr(v)
	case ssaop.OpNilCheck:
		v.Op = ssaop.OpSPARC64LoweredNilCheck
		return true
	case ssaop.OpNot:
		return rewriteValue_OpNot(v)
	case ssaop.OpOffPtr:
		return rewriteValue_OpOffPtr(v)
	case ssaop.OpOr16:
		v.Op = ssaop.OpSPARC64OR
		return true
	case ssaop.OpOr32:
		v.Op = ssaop.OpSPARC64OR
		return true
	case ssaop.OpOr64:
		v.Op = ssaop.OpSPARC64OR
		return true
	case ssaop.OpOr8:
		v.Op = ssaop.OpSPARC64OR
		return true
	case ssaop.OpOrB:
		v.Op = ssaop.OpSPARC64OR
		return true
	case ssaop.OpPanicBounds:
		v.Op = ssaop.OpSPARC64LoweredPanicBoundsRR
		return true
	case ssaop.OpPubBarrier:
		v.Op = ssaop.OpSPARC64LoweredPubBarrier
		return true
	case ssaop.OpRotateLeft16:
		return rewriteValue_OpRotateLeft16(v)
	case ssaop.OpRotateLeft32:
		return rewriteValue_OpRotateLeft32(v)
	case ssaop.OpRotateLeft64:
		return rewriteValue_OpRotateLeft64(v)
	case ssaop.OpRotateLeft8:
		return rewriteValue_OpRotateLeft8(v)
	case ssaop.OpRound32F:
		v.Op = ssaop.OpCopy
		return true
	case ssaop.OpRound64F:
		v.Op = ssaop.OpCopy
		return true
	case ssaop.OpRsh16Ux16:
		return rewriteValue_OpRsh16Ux16(v)
	case ssaop.OpRsh16Ux32:
		return rewriteValue_OpRsh16Ux32(v)
	case ssaop.OpRsh16Ux64:
		return rewriteValue_OpRsh16Ux64(v)
	case ssaop.OpRsh16Ux8:
		return rewriteValue_OpRsh16Ux8(v)
	case ssaop.OpRsh16x16:
		return rewriteValue_OpRsh16x16(v)
	case ssaop.OpRsh16x32:
		return rewriteValue_OpRsh16x32(v)
	case ssaop.OpRsh16x64:
		return rewriteValue_OpRsh16x64(v)
	case ssaop.OpRsh16x8:
		return rewriteValue_OpRsh16x8(v)
	case ssaop.OpRsh32Ux16:
		return rewriteValue_OpRsh32Ux16(v)
	case ssaop.OpRsh32Ux32:
		return rewriteValue_OpRsh32Ux32(v)
	case ssaop.OpRsh32Ux64:
		return rewriteValue_OpRsh32Ux64(v)
	case ssaop.OpRsh32Ux8:
		return rewriteValue_OpRsh32Ux8(v)
	case ssaop.OpRsh32x16:
		return rewriteValue_OpRsh32x16(v)
	case ssaop.OpRsh32x32:
		return rewriteValue_OpRsh32x32(v)
	case ssaop.OpRsh32x64:
		return rewriteValue_OpRsh32x64(v)
	case ssaop.OpRsh32x8:
		return rewriteValue_OpRsh32x8(v)
	case ssaop.OpRsh64Ux16:
		return rewriteValue_OpRsh64Ux16(v)
	case ssaop.OpRsh64Ux32:
		return rewriteValue_OpRsh64Ux32(v)
	case ssaop.OpRsh64Ux64:
		return rewriteValue_OpRsh64Ux64(v)
	case ssaop.OpRsh64Ux8:
		return rewriteValue_OpRsh64Ux8(v)
	case ssaop.OpRsh64x16:
		return rewriteValue_OpRsh64x16(v)
	case ssaop.OpRsh64x32:
		return rewriteValue_OpRsh64x32(v)
	case ssaop.OpRsh64x64:
		return rewriteValue_OpRsh64x64(v)
	case ssaop.OpRsh64x8:
		return rewriteValue_OpRsh64x8(v)
	case ssaop.OpRsh8Ux16:
		return rewriteValue_OpRsh8Ux16(v)
	case ssaop.OpRsh8Ux32:
		return rewriteValue_OpRsh8Ux32(v)
	case ssaop.OpRsh8Ux64:
		return rewriteValue_OpRsh8Ux64(v)
	case ssaop.OpRsh8Ux8:
		return rewriteValue_OpRsh8Ux8(v)
	case ssaop.OpRsh8x16:
		return rewriteValue_OpRsh8x16(v)
	case ssaop.OpRsh8x32:
		return rewriteValue_OpRsh8x32(v)
	case ssaop.OpRsh8x64:
		return rewriteValue_OpRsh8x64(v)
	case ssaop.OpRsh8x8:
		return rewriteValue_OpRsh8x8(v)
	case ssaop.OpSPARC64ADD:
		return rewriteValue_OpSPARC64ADD(v)
	case ssaop.OpSPARC64ADDconst:
		return rewriteValue_OpSPARC64ADDconst(v)
	case ssaop.OpSPARC64AND:
		return rewriteValue_OpSPARC64AND(v)
	case ssaop.OpSPARC64ANDconst:
		return rewriteValue_OpSPARC64ANDconst(v)
	case ssaop.OpSPARC64CMP:
		return rewriteValue_OpSPARC64CMP(v)
	case ssaop.OpSPARC64FMOVDload:
		return rewriteValue_OpSPARC64FMOVDload(v)
	case ssaop.OpSPARC64FMOVDstore:
		return rewriteValue_OpSPARC64FMOVDstore(v)
	case ssaop.OpSPARC64FMOVSload:
		return rewriteValue_OpSPARC64FMOVSload(v)
	case ssaop.OpSPARC64FMOVSstore:
		return rewriteValue_OpSPARC64FMOVSstore(v)
	case ssaop.OpSPARC64LoweredPanicBoundsCR:
		return rewriteValue_OpSPARC64LoweredPanicBoundsCR(v)
	case ssaop.OpSPARC64LoweredPanicBoundsRC:
		return rewriteValue_OpSPARC64LoweredPanicBoundsRC(v)
	case ssaop.OpSPARC64LoweredPanicBoundsRR:
		return rewriteValue_OpSPARC64LoweredPanicBoundsRR(v)
	case ssaop.OpSPARC64MOVB:
		return rewriteValue_OpSPARC64MOVB(v)
	case ssaop.OpSPARC64MOVBload:
		return rewriteValue_OpSPARC64MOVBload(v)
	case ssaop.OpSPARC64MOVBstore:
		return rewriteValue_OpSPARC64MOVBstore(v)
	case ssaop.OpSPARC64MOVD:
		return rewriteValue_OpSPARC64MOVD(v)
	case ssaop.OpSPARC64MOVDload:
		return rewriteValue_OpSPARC64MOVDload(v)
	case ssaop.OpSPARC64MOVDstore:
		return rewriteValue_OpSPARC64MOVDstore(v)
	case ssaop.OpSPARC64MOVH:
		return rewriteValue_OpSPARC64MOVH(v)
	case ssaop.OpSPARC64MOVHload:
		return rewriteValue_OpSPARC64MOVHload(v)
	case ssaop.OpSPARC64MOVHstore:
		return rewriteValue_OpSPARC64MOVHstore(v)
	case ssaop.OpSPARC64MOVUB:
		return rewriteValue_OpSPARC64MOVUB(v)
	case ssaop.OpSPARC64MOVUBload:
		return rewriteValue_OpSPARC64MOVUBload(v)
	case ssaop.OpSPARC64MOVUH:
		return rewriteValue_OpSPARC64MOVUH(v)
	case ssaop.OpSPARC64MOVUHload:
		return rewriteValue_OpSPARC64MOVUHload(v)
	case ssaop.OpSPARC64MOVUW:
		return rewriteValue_OpSPARC64MOVUW(v)
	case ssaop.OpSPARC64MOVUWload:
		return rewriteValue_OpSPARC64MOVUWload(v)
	case ssaop.OpSPARC64MOVW:
		return rewriteValue_OpSPARC64MOVW(v)
	case ssaop.OpSPARC64MOVWload:
		return rewriteValue_OpSPARC64MOVWload(v)
	case ssaop.OpSPARC64MOVWstore:
		return rewriteValue_OpSPARC64MOVWstore(v)
	case ssaop.OpSPARC64OR:
		return rewriteValue_OpSPARC64OR(v)
	case ssaop.OpSPARC64ORconst:
		return rewriteValue_OpSPARC64ORconst(v)
	case ssaop.OpSPARC64SLLD:
		return rewriteValue_OpSPARC64SLLD(v)
	case ssaop.OpSPARC64SRAD:
		return rewriteValue_OpSPARC64SRAD(v)
	case ssaop.OpSPARC64SRLD:
		return rewriteValue_OpSPARC64SRLD(v)
	case ssaop.OpSPARC64SUB:
		return rewriteValue_OpSPARC64SUB(v)
	case ssaop.OpSPARC64SUBconst:
		return rewriteValue_OpSPARC64SUBconst(v)
	case ssaop.OpSPARC64XOR:
		return rewriteValue_OpSPARC64XOR(v)
	case ssaop.OpSPARC64XORconst:
		return rewriteValue_OpSPARC64XORconst(v)
	case ssaop.OpSignExt16to32:
		v.Op = ssaop.OpSPARC64MOVH
		return true
	case ssaop.OpSignExt16to64:
		v.Op = ssaop.OpSPARC64MOVH
		return true
	case ssaop.OpSignExt32to64:
		v.Op = ssaop.OpSPARC64MOVW
		return true
	case ssaop.OpSignExt8to16:
		v.Op = ssaop.OpSPARC64MOVB
		return true
	case ssaop.OpSignExt8to32:
		v.Op = ssaop.OpSPARC64MOVB
		return true
	case ssaop.OpSignExt8to64:
		v.Op = ssaop.OpSPARC64MOVB
		return true
	case ssaop.OpSlicemask:
		return rewriteValue_OpSlicemask(v)
	case ssaop.OpSqrt:
		v.Op = ssaop.OpSPARC64FSQRTD
		return true
	case ssaop.OpSqrt32:
		v.Op = ssaop.OpSPARC64FSQRTS
		return true
	case ssaop.OpStaticCall:
		v.Op = ssaop.OpSPARC64CALLstatic
		return true
	case ssaop.OpStore:
		return rewriteValue_OpStore(v)
	case ssaop.OpSub16:
		v.Op = ssaop.OpSPARC64SUB
		return true
	case ssaop.OpSub32:
		v.Op = ssaop.OpSPARC64SUB
		return true
	case ssaop.OpSub32F:
		v.Op = ssaop.OpSPARC64FSUBS
		return true
	case ssaop.OpSub64:
		v.Op = ssaop.OpSPARC64SUB
		return true
	case ssaop.OpSub64F:
		v.Op = ssaop.OpSPARC64FSUBD
		return true
	case ssaop.OpSub64borrow:
		v.Op = ssaop.OpSPARC64SUBBORROW
		return true
	case ssaop.OpSub8:
		v.Op = ssaop.OpSPARC64SUB
		return true
	case ssaop.OpSubPtr:
		v.Op = ssaop.OpSPARC64SUB
		return true
	case ssaop.OpTailCall:
		v.Op = ssaop.OpSPARC64CALLtail
		return true
	case ssaop.OpTailCallInter:
		v.Op = ssaop.OpSPARC64CALLtail
		return true
	case ssaop.OpTrunc16to8:
		v.Op = ssaop.OpCopy
		return true
	case ssaop.OpTrunc32to16:
		v.Op = ssaop.OpCopy
		return true
	case ssaop.OpTrunc32to8:
		v.Op = ssaop.OpCopy
		return true
	case ssaop.OpTrunc64to16:
		v.Op = ssaop.OpCopy
		return true
	case ssaop.OpTrunc64to32:
		v.Op = ssaop.OpCopy
		return true
	case ssaop.OpTrunc64to8:
		v.Op = ssaop.OpCopy
		return true
	case ssaop.OpWB:
		v.Op = ssaop.OpSPARC64LoweredWB
		return true
	case ssaop.OpXor16:
		v.Op = ssaop.OpSPARC64XOR
		return true
	case ssaop.OpXor32:
		v.Op = ssaop.OpSPARC64XOR
		return true
	case ssaop.OpXor64:
		v.Op = ssaop.OpSPARC64XOR
		return true
	case ssaop.OpXor8:
		v.Op = ssaop.OpSPARC64XOR
		return true
	case ssaop.OpZero:
		return rewriteValue_OpZero(v)
	case ssaop.OpZeroExt16to32:
		v.Op = ssaop.OpSPARC64MOVUH
		return true
	case ssaop.OpZeroExt16to64:
		v.Op = ssaop.OpSPARC64MOVUH
		return true
	case ssaop.OpZeroExt32to64:
		v.Op = ssaop.OpSPARC64MOVUW
		return true
	case ssaop.OpZeroExt8to16:
		v.Op = ssaop.OpSPARC64MOVUB
		return true
	case ssaop.OpZeroExt8to32:
		v.Op = ssaop.OpSPARC64MOVUB
		return true
	case ssaop.OpZeroExt8to64:
		v.Op = ssaop.OpSPARC64MOVUB
		return true
	}
	return false
}
func rewriteValue_OpAddr(v *ssa.Value) bool {
	v_0 := v.Args[0]
	// match: (Addr {sym} base)
	// result: (MOVDaddr {sym} base)
	for {
		sym := ssa.AuxToSym(v.Aux)
		base := v_0
		v.Reset(ssaop.OpSPARC64MOVDaddr)
		v.Aux = ssa.SymToAux(sym)
		v.AddArg(base)
		return true
	}
}
func rewriteValue_OpAvg64u(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Avg64u <t> x y)
	// result: (ADD (SRLDconst <t> [1] (SUB <t> x y)) y)
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64ADD)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64SRLDconst, t)
		v0.AuxInt = ssa.Int64ToAuxInt(1)
		v1 := b.NewValue0(v.Pos, ssaop.OpSPARC64SUB, t)
		v1.AddArg2(x, y)
		v0.AddArg(v1)
		v.AddArg2(v0, y)
		return true
	}
}
func rewriteValue_OpCom16(v *ssa.Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Com16 x)
	// result: (XNOR x (MOVDconst [0]))
	for {
		x := v_0
		v.Reset(ssaop.OpSPARC64XNOR)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64MOVDconst, typ.UInt64)
		v0.AuxInt = ssa.Int64ToAuxInt(0)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpCom32(v *ssa.Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Com32 x)
	// result: (XNOR x (MOVDconst [0]))
	for {
		x := v_0
		v.Reset(ssaop.OpSPARC64XNOR)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64MOVDconst, typ.UInt64)
		v0.AuxInt = ssa.Int64ToAuxInt(0)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpCom64(v *ssa.Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Com64 x)
	// result: (XNOR x (MOVDconst [0]))
	for {
		x := v_0
		v.Reset(ssaop.OpSPARC64XNOR)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64MOVDconst, typ.UInt64)
		v0.AuxInt = ssa.Int64ToAuxInt(0)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpCom8(v *ssa.Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Com8 x)
	// result: (XNOR x (MOVDconst [0]))
	for {
		x := v_0
		v.Reset(ssaop.OpSPARC64XNOR)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64MOVDconst, typ.UInt64)
		v0.AuxInt = ssa.Int64ToAuxInt(0)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpConst16(v *ssa.Value) bool {
	// match: (Const16 [val])
	// result: (MOVDconst [int64(val)])
	for {
		val := ssa.AuxIntToInt16(v.AuxInt)
		v.Reset(ssaop.OpSPARC64MOVDconst)
		v.AuxInt = ssa.Int64ToAuxInt(int64(val))
		return true
	}
}
func rewriteValue_OpConst32(v *ssa.Value) bool {
	// match: (Const32 [val])
	// result: (MOVDconst [int64(val)])
	for {
		val := ssa.AuxIntToInt32(v.AuxInt)
		v.Reset(ssaop.OpSPARC64MOVDconst)
		v.AuxInt = ssa.Int64ToAuxInt(int64(val))
		return true
	}
}
func rewriteValue_OpConst64(v *ssa.Value) bool {
	// match: (Const64 [val])
	// result: (MOVDconst [int64(val)])
	for {
		val := ssa.AuxIntToInt64(v.AuxInt)
		v.Reset(ssaop.OpSPARC64MOVDconst)
		v.AuxInt = ssa.Int64ToAuxInt(int64(val))
		return true
	}
}
func rewriteValue_OpConst8(v *ssa.Value) bool {
	// match: (Const8 [val])
	// result: (MOVDconst [int64(val)])
	for {
		val := ssa.AuxIntToInt8(v.AuxInt)
		v.Reset(ssaop.OpSPARC64MOVDconst)
		v.AuxInt = ssa.Int64ToAuxInt(int64(val))
		return true
	}
}
func rewriteValue_OpConstBool(v *ssa.Value) bool {
	// match: (ConstBool [t])
	// result: (MOVDconst [ssa.B2i(t)])
	for {
		t := ssa.AuxIntToBool(v.AuxInt)
		v.Reset(ssaop.OpSPARC64MOVDconst)
		v.AuxInt = ssa.Int64ToAuxInt(ssa.B2i(t))
		return true
	}
}
func rewriteValue_OpConstNil(v *ssa.Value) bool {
	// match: (ConstNil)
	// result: (MOVDconst [0])
	for {
		v.Reset(ssaop.OpSPARC64MOVDconst)
		v.AuxInt = ssa.Int64ToAuxInt(0)
		return true
	}
}
func rewriteValue_OpCtz32(v *ssa.Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Ctz32 x)
	// result: (Ctz64 (OR <typ.Int64> x (MOVDconst [1<<32])))
	for {
		x := v_0
		v.Reset(ssaop.OpCtz64)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64OR, typ.Int64)
		v1 := b.NewValue0(v.Pos, ssaop.OpSPARC64MOVDconst, typ.UInt64)
		v1.AuxInt = ssa.Int64ToAuxInt(1 << 32)
		v0.AddArg2(x, v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpCtz64(v *ssa.Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Ctz64 x)
	// result: (POPC (ANDN <typ.Int64> (SUBconst <typ.Int64> [1] x) x))
	for {
		x := v_0
		v.Reset(ssaop.OpSPARC64POPC)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64ANDN, typ.Int64)
		v1 := b.NewValue0(v.Pos, ssaop.OpSPARC64SUBconst, typ.Int64)
		v1.AuxInt = ssa.Int64ToAuxInt(1)
		v1.AddArg(x)
		v0.AddArg2(v1, x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpCvt32Fto32(v *ssa.Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Cvt32Fto32 x)
	// result: (MOVW <typ.Int64> (FSTOI <typ.Int32> x))
	for {
		x := v_0
		v.Reset(ssaop.OpSPARC64MOVW)
		v.Type = typ.Int64
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64FSTOI, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpCvt32to32F(v *ssa.Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Cvt32to32F x)
	// result: (FXTOS (SignExt32to64 <typ.Int64> x))
	for {
		x := v_0
		v.Reset(ssaop.OpSPARC64FXTOS)
		v0 := b.NewValue0(v.Pos, ssaop.OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpCvt32to64F(v *ssa.Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Cvt32to64F x)
	// result: (FXTOD (SignExt32to64 <typ.Int64> x))
	for {
		x := v_0
		v.Reset(ssaop.OpSPARC64FXTOD)
		v0 := b.NewValue0(v.Pos, ssaop.OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpCvt64Fto32(v *ssa.Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Cvt64Fto32 x)
	// result: (MOVW <typ.Int64> (FDTOI <typ.Int32> x))
	for {
		x := v_0
		v.Reset(ssaop.OpSPARC64MOVW)
		v.Type = typ.Int64
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64FDTOI, typ.Int32)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpDiv16(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div16 x y)
	// result: (SDIVD (SignExt16to64 <typ.Int64> x) (SignExt16to64 <typ.Int64> y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64SDIVD)
		v0 := b.NewValue0(v.Pos, ssaop.OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, ssaop.OpSignExt16to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValue_OpDiv16u(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div16u x y)
	// result: (UDIVD (ZeroExt16to64 <typ.UInt64> x) (ZeroExt16to64 <typ.UInt64> y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64UDIVD)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, ssaop.OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValue_OpDiv32(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div32 x y)
	// result: (SDIVD (SignExt32to64 <typ.Int64> x) (SignExt32to64 <typ.Int64> y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64SDIVD)
		v0 := b.NewValue0(v.Pos, ssaop.OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, ssaop.OpSignExt32to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValue_OpDiv32u(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div32u x y)
	// result: (UDIVD (ZeroExt32to64 <typ.UInt64> x) (ZeroExt32to64 <typ.UInt64> y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64UDIVD)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, ssaop.OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValue_OpDiv64(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Div64 [false] x y)
	// result: (SDIVD x y)
	for {
		if ssa.AuxIntToBool(v.AuxInt) != false {
			break
		}
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64SDIVD)
		v.AddArg2(x, y)
		return true
	}
	return false
}
func rewriteValue_OpDiv64u(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Div64u x y)
	// result: (UDIVD x y)
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64UDIVD)
		v.AddArg2(x, y)
		return true
	}
}
func rewriteValue_OpDiv8(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div8 x y)
	// result: (SDIVD (SignExt8to64 <typ.Int64> x) (SignExt8to64 <typ.Int64> y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64SDIVD)
		v0 := b.NewValue0(v.Pos, ssaop.OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, ssaop.OpSignExt8to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValue_OpDiv8u(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div8u x y)
	// result: (UDIVD (ZeroExt8to64 <typ.UInt64> x) (ZeroExt8to64 <typ.UInt64> y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64UDIVD)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, ssaop.OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValue_OpEq16(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq16 x y)
	// result: (Equal (CMP (ZeroExt16to64 x) (ZeroExt16to64 y)))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64Equal)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, ssaop.OpZeroExt16to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, ssaop.OpZeroExt16to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpEq32(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq32 x y)
	// result: (Equal (CMP (ZeroExt32to64 x) (ZeroExt32to64 y)))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64Equal)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, ssaop.OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, ssaop.OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpEq32F(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Eq32F x y)
	// result: (FEqual (FCMPS x y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64FEqual)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64FCMPS, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpEq64(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Eq64 x y)
	// result: (Equal (CMP x y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64Equal)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMP, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpEq64F(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Eq64F x y)
	// result: (FEqual (FCMPD x y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64FEqual)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64FCMPD, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpEq8(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq8 x y)
	// result: (Equal (CMP (ZeroExt8to64 x) (ZeroExt8to64 y)))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64Equal)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, ssaop.OpZeroExt8to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, ssaop.OpZeroExt8to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpEqB(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (EqB x y)
	// result: (Equal (CMP (ZeroExt8to64 x) (ZeroExt8to64 y)))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64Equal)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, ssaop.OpZeroExt8to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, ssaop.OpZeroExt8to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpEqPtr(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (EqPtr x y)
	// result: (Equal (CMP x y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64Equal)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMP, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpHmul32(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Hmul32 x y)
	// result: (SRADconst [32] (MULD <typ.Int64> (SignExt32to64 <typ.Int64> x) (SignExt32to64 <typ.Int64> y)))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64SRADconst)
		v.AuxInt = ssa.Int64ToAuxInt(32)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64MULD, typ.Int64)
		v1 := b.NewValue0(v.Pos, ssaop.OpSignExt32to64, typ.Int64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, ssaop.OpSignExt32to64, typ.Int64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpHmul32u(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Hmul32u x y)
	// result: (SRLDconst [32] (MULD <typ.UInt64> (ZeroExt32to64 <typ.UInt64> x) (ZeroExt32to64 <typ.UInt64> y)))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64SRLDconst)
		v.AuxInt = ssa.Int64ToAuxInt(32)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64MULD, typ.UInt64)
		v1 := b.NewValue0(v.Pos, ssaop.OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, ssaop.OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpHmul64(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Hmul64 x y)
	// result: (SUB <typ.Int64> (SUB <typ.Int64> (UMULXHI <typ.Int64> x y) (AND <typ.Int64> y (SRADconst <typ.Int64> [63] x))) (AND <typ.Int64> x (SRADconst <typ.Int64> [63] y)))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64SUB)
		v.Type = typ.Int64
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64SUB, typ.Int64)
		v1 := b.NewValue0(v.Pos, ssaop.OpSPARC64UMULXHI, typ.Int64)
		v1.AddArg2(x, y)
		v2 := b.NewValue0(v.Pos, ssaop.OpSPARC64AND, typ.Int64)
		v3 := b.NewValue0(v.Pos, ssaop.OpSPARC64SRADconst, typ.Int64)
		v3.AuxInt = ssa.Int64ToAuxInt(63)
		v3.AddArg(x)
		v2.AddArg2(y, v3)
		v0.AddArg2(v1, v2)
		v4 := b.NewValue0(v.Pos, ssaop.OpSPARC64AND, typ.Int64)
		v5 := b.NewValue0(v.Pos, ssaop.OpSPARC64SRADconst, typ.Int64)
		v5.AuxInt = ssa.Int64ToAuxInt(63)
		v5.AddArg(y)
		v4.AddArg2(x, v5)
		v.AddArg2(v0, v4)
		return true
	}
}
func rewriteValue_OpHmul64u(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Hmul64u x y)
	// result: (UMULXHI x y)
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64UMULXHI)
		v.AddArg2(x, y)
		return true
	}
}
func rewriteValue_OpIsInBounds(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (IsInBounds idx len)
	// result: (LessThanU (CMP idx len))
	for {
		idx := v_0
		len := v_1
		v.Reset(ssaop.OpSPARC64LessThanU)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMP, types.TypeFlags)
		v0.AddArg2(idx, len)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpIsNonNil(v *ssa.Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (IsNonNil ptr)
	// result: (NotEqual (CMPconst [0] ptr))
	for {
		ptr := v_0
		v.Reset(ssaop.OpSPARC64NotEqual)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMPconst, types.TypeFlags)
		v0.AuxInt = ssa.Int64ToAuxInt(0)
		v0.AddArg(ptr)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpIsSliceInBounds(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (IsSliceInBounds idx len)
	// result: (LessEqualU (CMP idx len))
	for {
		idx := v_0
		len := v_1
		v.Reset(ssaop.OpSPARC64LessEqualU)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMP, types.TypeFlags)
		v0.AddArg2(idx, len)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpLeq16(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq16 x y)
	// result: (LessEqual (CMP (SignExt16to64 x) (SignExt16to64 y)))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64LessEqual)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, ssaop.OpSignExt16to64, typ.Int64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, ssaop.OpSignExt16to64, typ.Int64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpLeq16U(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq16U x y)
	// result: (LessEqualU (CMP (ZeroExt16to64 x) (ZeroExt16to64 y)))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64LessEqualU)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, ssaop.OpZeroExt16to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, ssaop.OpZeroExt16to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpLeq32(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq32 x y)
	// result: (LessEqual (CMP (SignExt32to64 x) (SignExt32to64 y)))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64LessEqual)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, ssaop.OpSignExt32to64, typ.Int64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, ssaop.OpSignExt32to64, typ.Int64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpLeq32F(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Leq32F x y)
	// result: (FLessEqual (FCMPS x y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64FLessEqual)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64FCMPS, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpLeq32U(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq32U x y)
	// result: (LessEqualU (CMP (ZeroExt32to64 x) (ZeroExt32to64 y)))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64LessEqualU)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, ssaop.OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, ssaop.OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpLeq64(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Leq64 x y)
	// result: (LessEqual (CMP x y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64LessEqual)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMP, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpLeq64F(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Leq64F x y)
	// result: (FLessEqual (FCMPD x y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64FLessEqual)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64FCMPD, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpLeq64U(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Leq64U x y)
	// result: (LessEqualU (CMP x y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64LessEqualU)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMP, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpLeq8(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq8 x y)
	// result: (LessEqual (CMP (SignExt8to64 x) (SignExt8to64 y)))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64LessEqual)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, ssaop.OpSignExt8to64, typ.Int64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, ssaop.OpSignExt8to64, typ.Int64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpLeq8U(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq8U x y)
	// result: (LessEqualU (CMP (ZeroExt8to64 x) (ZeroExt8to64 y)))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64LessEqualU)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, ssaop.OpZeroExt8to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, ssaop.OpZeroExt8to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpLess16(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less16 x y)
	// result: (LessThan (CMP (SignExt16to64 x) (SignExt16to64 y)))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64LessThan)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, ssaop.OpSignExt16to64, typ.Int64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, ssaop.OpSignExt16to64, typ.Int64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpLess16U(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less16U x y)
	// result: (LessThanU (CMP (ZeroExt16to64 x) (ZeroExt16to64 y)))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64LessThanU)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, ssaop.OpZeroExt16to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, ssaop.OpZeroExt16to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpLess32(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less32 x y)
	// result: (LessThan (CMP (SignExt32to64 x) (SignExt32to64 y)))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64LessThan)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, ssaop.OpSignExt32to64, typ.Int64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, ssaop.OpSignExt32to64, typ.Int64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpLess32F(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Less32F x y)
	// result: (FLessThan (FCMPS x y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64FLessThan)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64FCMPS, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpLess32U(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less32U x y)
	// result: (LessThanU (CMP (ZeroExt32to64 x) (ZeroExt32to64 y)))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64LessThanU)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, ssaop.OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, ssaop.OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpLess64(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Less64 x y)
	// result: (LessThan (CMP x y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64LessThan)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMP, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpLess64F(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Less64F x y)
	// result: (FLessThan (FCMPD x y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64FLessThan)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64FCMPD, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpLess64U(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Less64U x y)
	// result: (LessThanU (CMP x y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64LessThanU)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMP, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpLess8(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less8 x y)
	// result: (LessThan (CMP (SignExt8to64 x) (SignExt8to64 y)))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64LessThan)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, ssaop.OpSignExt8to64, typ.Int64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, ssaop.OpSignExt8to64, typ.Int64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpLess8U(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less8U x y)
	// result: (LessThanU (CMP (ZeroExt8to64 x) (ZeroExt8to64 y)))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64LessThanU)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, ssaop.OpZeroExt8to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, ssaop.OpZeroExt8to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpLoad(v *ssa.Value) bool {
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
		v.Reset(ssaop.OpSPARC64MOVUBload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: ssa.Is8BitInt(t) && t.IsSigned()
	// result: (MOVBload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(ssa.Is8BitInt(t) && t.IsSigned()) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVBload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: ssa.Is8BitInt(t) && !t.IsSigned()
	// result: (MOVUBload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(ssa.Is8BitInt(t) && !t.IsSigned()) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVUBload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: ssa.Is16BitInt(t) && t.IsSigned()
	// result: (MOVHload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(ssa.Is16BitInt(t) && t.IsSigned()) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVHload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: ssa.Is16BitInt(t) && !t.IsSigned()
	// result: (MOVUHload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(ssa.Is16BitInt(t) && !t.IsSigned()) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVUHload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: ssa.Is32BitInt(t) && t.IsSigned()
	// result: (MOVWload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(ssa.Is32BitInt(t) && t.IsSigned()) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVWload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: ssa.Is32BitInt(t) && !t.IsSigned()
	// result: (MOVUWload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(ssa.Is32BitInt(t) && !t.IsSigned()) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVUWload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: (ssa.Is64BitInt(t) || ssa.IsPtr(t))
	// result: (MOVDload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(ssa.Is64BitInt(t) || ssa.IsPtr(t)) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVDload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: ssa.Is32BitFloat(t)
	// result: (FMOVSload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(ssa.Is32BitFloat(t)) {
			break
		}
		v.Reset(ssaop.OpSPARC64FMOVSload)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: ssa.Is64BitFloat(t)
	// result: (FMOVDload ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(ssa.Is64BitFloat(t)) {
			break
		}
		v.Reset(ssaop.OpSPARC64FMOVDload)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValue_OpLocalAddr(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (LocalAddr <t> {sym} base mem)
	// cond: t.Elem().HasPointers()
	// result: (MOVDaddr {sym} (SPanchored base mem))
	for {
		t := v.Type
		sym := ssa.AuxToSym(v.Aux)
		base := v_0
		mem := v_1
		if !(t.Elem().HasPointers()) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVDaddr)
		v.Aux = ssa.SymToAux(sym)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPanchored, typ.Uintptr)
		v0.AddArg2(base, mem)
		v.AddArg(v0)
		return true
	}
	// match: (LocalAddr <t> {sym} base _)
	// cond: !t.Elem().HasPointers()
	// result: (MOVDaddr {sym} base)
	for {
		t := v.Type
		sym := ssa.AuxToSym(v.Aux)
		base := v_0
		if !(!t.Elem().HasPointers()) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVDaddr)
		v.Aux = ssa.SymToAux(sym)
		v.AddArg(base)
		return true
	}
	return false
}
func rewriteValue_OpLsh16x16(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x16 x y)
	// result: (Lsh16x64 x (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpLsh16x64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpLsh16x32(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x32 x y)
	// result: (Lsh16x64 x (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpLsh16x64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpLsh16x64(v *ssa.Value) bool {
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
		if v_1.Op != ssaop.OpConst64 {
			break
		}
		c := ssa.AuxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.Reset(ssaop.OpSPARC64SLLDconst)
		v.Type = t
		v.AuxInt = ssa.Int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (Lsh16x64 _ (Const64 [c]))
	// cond: uint64(c) >= 64
	// result: (MOVDconst [0])
	for {
		if v_1.Op != ssaop.OpConst64 {
			break
		}
		c := ssa.AuxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 64) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVDconst)
		v.AuxInt = ssa.Int64ToAuxInt(0)
		return true
	}
	// match: (Lsh16x64 <t> x y)
	// result: (AND (NEG <t> (LessThanU <typ.UInt64> (CMPconst <types.TypeFlags> [64] y))) (SLLD <t> x y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64AND)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64NEG, t)
		v1 := b.NewValue0(v.Pos, ssaop.OpSPARC64LessThanU, typ.UInt64)
		v2 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMPconst, types.TypeFlags)
		v2.AuxInt = ssa.Int64ToAuxInt(64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, ssaop.OpSPARC64SLLD, t)
		v3.AddArg2(x, y)
		v.AddArg2(v0, v3)
		return true
	}
}
func rewriteValue_OpLsh16x8(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x8 x y)
	// result: (Lsh16x64 x (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpLsh16x64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpLsh32x16(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x16 x y)
	// result: (Lsh32x64 x (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpLsh32x64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpLsh32x32(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x32 x y)
	// result: (Lsh32x64 x (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpLsh32x64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpLsh32x64(v *ssa.Value) bool {
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
		if v_1.Op != ssaop.OpConst64 {
			break
		}
		c := ssa.AuxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.Reset(ssaop.OpSPARC64SLLDconst)
		v.Type = t
		v.AuxInt = ssa.Int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (Lsh32x64 _ (Const64 [c]))
	// cond: uint64(c) >= 64
	// result: (MOVDconst [0])
	for {
		if v_1.Op != ssaop.OpConst64 {
			break
		}
		c := ssa.AuxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 64) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVDconst)
		v.AuxInt = ssa.Int64ToAuxInt(0)
		return true
	}
	// match: (Lsh32x64 <t> x y)
	// result: (AND (NEG <t> (LessThanU <typ.UInt64> (CMPconst <types.TypeFlags> [64] y))) (SLLD <t> x y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64AND)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64NEG, t)
		v1 := b.NewValue0(v.Pos, ssaop.OpSPARC64LessThanU, typ.UInt64)
		v2 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMPconst, types.TypeFlags)
		v2.AuxInt = ssa.Int64ToAuxInt(64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, ssaop.OpSPARC64SLLD, t)
		v3.AddArg2(x, y)
		v.AddArg2(v0, v3)
		return true
	}
}
func rewriteValue_OpLsh32x8(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x8 x y)
	// result: (Lsh32x64 x (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpLsh32x64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpLsh64x16(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x16 x y)
	// result: (Lsh64x64 x (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpLsh64x64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpLsh64x32(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x32 x y)
	// result: (Lsh64x64 x (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpLsh64x64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpLsh64x64(v *ssa.Value) bool {
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
		if v_1.Op != ssaop.OpConst64 {
			break
		}
		c := ssa.AuxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.Reset(ssaop.OpSPARC64SLLDconst)
		v.Type = t
		v.AuxInt = ssa.Int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (Lsh64x64 _ (Const64 [c]))
	// cond: uint64(c) >= 64
	// result: (MOVDconst [0])
	for {
		if v_1.Op != ssaop.OpConst64 {
			break
		}
		c := ssa.AuxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 64) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVDconst)
		v.AuxInt = ssa.Int64ToAuxInt(0)
		return true
	}
	// match: (Lsh64x64 <t> x y)
	// result: (AND (NEG <t> (LessThanU <typ.UInt64> (CMPconst <types.TypeFlags> [64] y))) (SLLD <t> x y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64AND)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64NEG, t)
		v1 := b.NewValue0(v.Pos, ssaop.OpSPARC64LessThanU, typ.UInt64)
		v2 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMPconst, types.TypeFlags)
		v2.AuxInt = ssa.Int64ToAuxInt(64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, ssaop.OpSPARC64SLLD, t)
		v3.AddArg2(x, y)
		v.AddArg2(v0, v3)
		return true
	}
}
func rewriteValue_OpLsh64x8(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x8 x y)
	// result: (Lsh64x64 x (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpLsh64x64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpLsh8x16(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x16 x y)
	// result: (Lsh8x64 x (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpLsh8x64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpLsh8x32(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x32 x y)
	// result: (Lsh8x64 x (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpLsh8x64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpLsh8x64(v *ssa.Value) bool {
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
		if v_1.Op != ssaop.OpConst64 {
			break
		}
		c := ssa.AuxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.Reset(ssaop.OpSPARC64SLLDconst)
		v.Type = t
		v.AuxInt = ssa.Int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (Lsh8x64 _ (Const64 [c]))
	// cond: uint64(c) >= 64
	// result: (MOVDconst [0])
	for {
		if v_1.Op != ssaop.OpConst64 {
			break
		}
		c := ssa.AuxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 64) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVDconst)
		v.AuxInt = ssa.Int64ToAuxInt(0)
		return true
	}
	// match: (Lsh8x64 <t> x y)
	// result: (AND (NEG <t> (LessThanU <typ.UInt64> (CMPconst <types.TypeFlags> [64] y))) (SLLD <t> x y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64AND)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64NEG, t)
		v1 := b.NewValue0(v.Pos, ssaop.OpSPARC64LessThanU, typ.UInt64)
		v2 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMPconst, types.TypeFlags)
		v2.AuxInt = ssa.Int64ToAuxInt(64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, ssaop.OpSPARC64SLLD, t)
		v3.AddArg2(x, y)
		v.AddArg2(v0, v3)
		return true
	}
}
func rewriteValue_OpLsh8x8(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x8 x y)
	// result: (Lsh8x64 x (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpLsh8x64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpMod16(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod16 x y)
	// result: (Mod64 (SignExt16to64 <typ.Int64> x) (SignExt16to64 <typ.Int64> y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpMod64)
		v0 := b.NewValue0(v.Pos, ssaop.OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, ssaop.OpSignExt16to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValue_OpMod16u(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod16u x y)
	// result: (Mod64u (ZeroExt16to64 <typ.UInt64> x) (ZeroExt16to64 <typ.UInt64> y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpMod64u)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, ssaop.OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValue_OpMod32(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod32 x y)
	// result: (Mod64 (SignExt32to64 <typ.Int64> x) (SignExt32to64 <typ.Int64> y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpMod64)
		v0 := b.NewValue0(v.Pos, ssaop.OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, ssaop.OpSignExt32to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValue_OpMod32u(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod32u x y)
	// result: (Mod64u (ZeroExt32to64 <typ.UInt64> x) (ZeroExt32to64 <typ.UInt64> y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpMod64u)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, ssaop.OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValue_OpMod64(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod64 x y)
	// result: (SUB <typ.Int64> x (MULD <typ.Int64> (SDIVD <typ.Int64> x y) y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64SUB)
		v.Type = typ.Int64
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64MULD, typ.Int64)
		v1 := b.NewValue0(v.Pos, ssaop.OpSPARC64SDIVD, typ.Int64)
		v1.AddArg2(x, y)
		v0.AddArg2(v1, y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpMod64u(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod64u x y)
	// result: (SUB <typ.UInt64> x (MULD <typ.UInt64> (UDIVD <typ.UInt64> x y) y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64SUB)
		v.Type = typ.UInt64
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64MULD, typ.UInt64)
		v1 := b.NewValue0(v.Pos, ssaop.OpSPARC64UDIVD, typ.UInt64)
		v1.AddArg2(x, y)
		v0.AddArg2(v1, y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpMod8(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod8 x y)
	// result: (Mod64 (SignExt8to64 <typ.Int64> x) (SignExt8to64 <typ.Int64> y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpMod64)
		v0 := b.NewValue0(v.Pos, ssaop.OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, ssaop.OpSignExt8to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValue_OpMod8u(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod8u x y)
	// result: (Mod64u (ZeroExt8to64 <typ.UInt64> x) (ZeroExt8to64 <typ.UInt64> y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpMod64u)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, ssaop.OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValue_OpMove(v *ssa.Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Move [0] _ _ mem)
	// result: mem
	for {
		if ssa.AuxIntToInt64(v.AuxInt) != 0 {
			break
		}
		mem := v_2
		v.CopyOf(mem)
		return true
	}
	// match: (Move [1] dst src mem)
	// result: (MOVBstore dst (MOVBload src mem) mem)
	for {
		if ssa.AuxIntToInt64(v.AuxInt) != 1 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.Reset(ssaop.OpSPARC64MOVBstore)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64MOVBload, typ.Int8)
		v0.AddArg2(src, mem)
		v.AddArg3(dst, v0, mem)
		return true
	}
	// match: (Move [2] {t} dst src mem)
	// cond: t.Alignment()%2 == 0
	// result: (MOVHstore dst (MOVHload src mem) mem)
	for {
		if ssa.AuxIntToInt64(v.AuxInt) != 2 {
			break
		}
		t := ssa.AuxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.Alignment()%2 == 0) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVHstore)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64MOVHload, typ.Int16)
		v0.AddArg2(src, mem)
		v.AddArg3(dst, v0, mem)
		return true
	}
	// match: (Move [4] {t} dst src mem)
	// cond: t.Alignment()%4 == 0
	// result: (MOVWstore dst (MOVWload src mem) mem)
	for {
		if ssa.AuxIntToInt64(v.AuxInt) != 4 {
			break
		}
		t := ssa.AuxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.Alignment()%4 == 0) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVWstore)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64MOVWload, typ.Int32)
		v0.AddArg2(src, mem)
		v.AddArg3(dst, v0, mem)
		return true
	}
	// match: (Move [8] {t} dst src mem)
	// cond: t.Alignment()%8 == 0
	// result: (MOVDstore dst (MOVDload src mem) mem)
	for {
		if ssa.AuxIntToInt64(v.AuxInt) != 8 {
			break
		}
		t := ssa.AuxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.Alignment()%8 == 0) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVDstore)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64MOVDload, typ.UInt64)
		v0.AddArg2(src, mem)
		v.AddArg3(dst, v0, mem)
		return true
	}
	// match: (Move [s] {t} dst src mem)
	// cond: t.Alignment()%8 == 0 && ssa.LogLargeCopyValue(v, s)
	// result: (LoweredMove [s<<4|8] dst src mem)
	for {
		s := ssa.AuxIntToInt64(v.AuxInt)
		t := ssa.AuxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.Alignment()%8 == 0 && ssa.LogLargeCopyValue(v, s)) {
			break
		}
		v.Reset(ssaop.OpSPARC64LoweredMove)
		v.AuxInt = ssa.Int64ToAuxInt(s<<4 | 8)
		v.AddArg3(dst, src, mem)
		return true
	}
	// match: (Move [s] {t} dst src mem)
	// cond: t.Alignment()%4 == 0 && ssa.LogLargeCopyValue(v, s)
	// result: (LoweredMove [s<<4|4] dst src mem)
	for {
		s := ssa.AuxIntToInt64(v.AuxInt)
		t := ssa.AuxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.Alignment()%4 == 0 && ssa.LogLargeCopyValue(v, s)) {
			break
		}
		v.Reset(ssaop.OpSPARC64LoweredMove)
		v.AuxInt = ssa.Int64ToAuxInt(s<<4 | 4)
		v.AddArg3(dst, src, mem)
		return true
	}
	// match: (Move [s] {t} dst src mem)
	// cond: t.Alignment()%2 == 0 && ssa.LogLargeCopyValue(v, s)
	// result: (LoweredMove [s<<4|2] dst src mem)
	for {
		s := ssa.AuxIntToInt64(v.AuxInt)
		t := ssa.AuxToType(v.Aux)
		dst := v_0
		src := v_1
		mem := v_2
		if !(t.Alignment()%2 == 0 && ssa.LogLargeCopyValue(v, s)) {
			break
		}
		v.Reset(ssaop.OpSPARC64LoweredMove)
		v.AuxInt = ssa.Int64ToAuxInt(s<<4 | 2)
		v.AddArg3(dst, src, mem)
		return true
	}
	// match: (Move [s] dst src mem)
	// cond: ssa.LogLargeCopyValue(v, s)
	// result: (LoweredMove [s<<4|1] dst src mem)
	for {
		s := ssa.AuxIntToInt64(v.AuxInt)
		dst := v_0
		src := v_1
		mem := v_2
		if !(ssa.LogLargeCopyValue(v, s)) {
			break
		}
		v.Reset(ssaop.OpSPARC64LoweredMove)
		v.AuxInt = ssa.Int64ToAuxInt(s<<4 | 1)
		v.AddArg3(dst, src, mem)
		return true
	}
	return false
}
func rewriteValue_OpNeg16(v *ssa.Value) bool {
	v_0 := v.Args[0]
	// match: (Neg16 x)
	// result: (NEG x)
	for {
		x := v_0
		v.Reset(ssaop.OpSPARC64NEG)
		v.AddArg(x)
		return true
	}
}
func rewriteValue_OpNeg32(v *ssa.Value) bool {
	v_0 := v.Args[0]
	// match: (Neg32 x)
	// result: (NEG x)
	for {
		x := v_0
		v.Reset(ssaop.OpSPARC64NEG)
		v.AddArg(x)
		return true
	}
}
func rewriteValue_OpNeg64(v *ssa.Value) bool {
	v_0 := v.Args[0]
	// match: (Neg64 x)
	// result: (NEG x)
	for {
		x := v_0
		v.Reset(ssaop.OpSPARC64NEG)
		v.AddArg(x)
		return true
	}
}
func rewriteValue_OpNeg8(v *ssa.Value) bool {
	v_0 := v.Args[0]
	// match: (Neg8 x)
	// result: (NEG x)
	for {
		x := v_0
		v.Reset(ssaop.OpSPARC64NEG)
		v.AddArg(x)
		return true
	}
}
func rewriteValue_OpNeq16(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq16 x y)
	// result: (NotEqual (CMP (ZeroExt16to64 x) (ZeroExt16to64 y)))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64NotEqual)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, ssaop.OpZeroExt16to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, ssaop.OpZeroExt16to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpNeq32(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq32 x y)
	// result: (NotEqual (CMP (ZeroExt32to64 x) (ZeroExt32to64 y)))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64NotEqual)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, ssaop.OpZeroExt32to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, ssaop.OpZeroExt32to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpNeq32F(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Neq32F x y)
	// result: (FNotEqual (FCMPS x y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64FNotEqual)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64FCMPS, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpNeq64(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Neq64 x y)
	// result: (NotEqual (CMP x y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64NotEqual)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMP, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpNeq64F(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (Neq64F x y)
	// result: (FNotEqual (FCMPD x y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64FNotEqual)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64FCMPD, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpNeq8(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq8 x y)
	// result: (NotEqual (CMP (ZeroExt8to64 x) (ZeroExt8to64 y)))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64NotEqual)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, ssaop.OpZeroExt8to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, ssaop.OpZeroExt8to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpNeqB(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (NeqB x y)
	// result: (NotEqual (CMP (ZeroExt8to64 x) (ZeroExt8to64 y)))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64NotEqual)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMP, types.TypeFlags)
		v1 := b.NewValue0(v.Pos, ssaop.OpZeroExt8to64, typ.UInt64)
		v1.AddArg(x)
		v2 := b.NewValue0(v.Pos, ssaop.OpZeroExt8to64, typ.UInt64)
		v2.AddArg(y)
		v0.AddArg2(v1, v2)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpNeqPtr(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	// match: (NeqPtr x y)
	// result: (NotEqual (CMP x y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64NotEqual)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMP, types.TypeFlags)
		v0.AddArg2(x, y)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpNot(v *ssa.Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Not x)
	// result: (XORconst <typ.Bool> [1] x)
	for {
		x := v_0
		v.Reset(ssaop.OpSPARC64XORconst)
		v.Type = typ.Bool
		v.AuxInt = ssa.Int64ToAuxInt(1)
		v.AddArg(x)
		return true
	}
}
func rewriteValue_OpOffPtr(v *ssa.Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (OffPtr [off] ptr:(SP))
	// cond: ssa.Is32Bit(off)
	// result: (MOVDaddr [int32(off)] ptr)
	for {
		off := ssa.AuxIntToInt64(v.AuxInt)
		ptr := v_0
		if ptr.Op != ssaop.OpSP || !(ssa.Is32Bit(off)) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVDaddr)
		v.AuxInt = ssa.Int32ToAuxInt(int32(off))
		v.AddArg(ptr)
		return true
	}
	// match: (OffPtr [off] ptr)
	// cond: ssa.Is13Bit(off)
	// result: (ADDconst [off] ptr)
	for {
		off := ssa.AuxIntToInt64(v.AuxInt)
		ptr := v_0
		if !(ssa.Is13Bit(off)) {
			break
		}
		v.Reset(ssaop.OpSPARC64ADDconst)
		v.AuxInt = ssa.Int64ToAuxInt(off)
		v.AddArg(ptr)
		return true
	}
	// match: (OffPtr [off] ptr)
	// result: (ADD (MOVDconst [off]) ptr)
	for {
		off := ssa.AuxIntToInt64(v.AuxInt)
		ptr := v_0
		v.Reset(ssaop.OpSPARC64ADD)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64MOVDconst, typ.UInt64)
		v0.AuxInt = ssa.Int64ToAuxInt(off)
		v.AddArg2(v0, ptr)
		return true
	}
}
func rewriteValue_OpRotateLeft16(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft16 <t> x (MOVDconst [c]))
	// result: (Or16 (Lsh16x64 <t> x (MOVDconst [c&15])) (Rsh16Ux64 <t> x (MOVDconst [-c&15])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != ssaop.OpSPARC64MOVDconst {
			break
		}
		c := ssa.AuxIntToInt64(v_1.AuxInt)
		v.Reset(ssaop.OpOr16)
		v0 := b.NewValue0(v.Pos, ssaop.OpLsh16x64, t)
		v1 := b.NewValue0(v.Pos, ssaop.OpSPARC64MOVDconst, typ.UInt64)
		v1.AuxInt = ssa.Int64ToAuxInt(c & 15)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, ssaop.OpRsh16Ux64, t)
		v3 := b.NewValue0(v.Pos, ssaop.OpSPARC64MOVDconst, typ.UInt64)
		v3.AuxInt = ssa.Int64ToAuxInt(-c & 15)
		v2.AddArg2(x, v3)
		v.AddArg2(v0, v2)
		return true
	}
	return false
}
func rewriteValue_OpRotateLeft32(v *ssa.Value) bool {
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
		v.Reset(ssaop.OpSPARC64OR)
		v.Type = t
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64SLLW, t)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, ssaop.OpSPARC64SRLW, t)
		v2 := b.NewValue0(v.Pos, ssaop.OpSPARC64SUB, typ.UInt64)
		v3 := b.NewValue0(v.Pos, ssaop.OpSPARC64MOVDconst, typ.UInt64)
		v3.AuxInt = ssa.Int64ToAuxInt(32)
		v2.AddArg2(v3, y)
		v1.AddArg2(x, v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValue_OpRotateLeft64(v *ssa.Value) bool {
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
		v.Reset(ssaop.OpSPARC64OR)
		v.Type = t
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64SLLD, t)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, ssaop.OpSPARC64SRLD, t)
		v2 := b.NewValue0(v.Pos, ssaop.OpSPARC64SUB, typ.UInt64)
		v3 := b.NewValue0(v.Pos, ssaop.OpSPARC64MOVDconst, typ.UInt64)
		v3.AuxInt = ssa.Int64ToAuxInt(64)
		v2.AddArg2(v3, y)
		v1.AddArg2(x, v2)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValue_OpRotateLeft8(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft8 <t> x (MOVDconst [c]))
	// result: (Or8 (Lsh8x64 <t> x (MOVDconst [c&7])) (Rsh8Ux64 <t> x (MOVDconst [-c&7])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != ssaop.OpSPARC64MOVDconst {
			break
		}
		c := ssa.AuxIntToInt64(v_1.AuxInt)
		v.Reset(ssaop.OpOr8)
		v0 := b.NewValue0(v.Pos, ssaop.OpLsh8x64, t)
		v1 := b.NewValue0(v.Pos, ssaop.OpSPARC64MOVDconst, typ.UInt64)
		v1.AuxInt = ssa.Int64ToAuxInt(c & 7)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, ssaop.OpRsh8Ux64, t)
		v3 := b.NewValue0(v.Pos, ssaop.OpSPARC64MOVDconst, typ.UInt64)
		v3.AuxInt = ssa.Int64ToAuxInt(-c & 7)
		v2.AddArg2(x, v3)
		v.AddArg2(v0, v2)
		return true
	}
	return false
}
func rewriteValue_OpRsh16Ux16(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux16 x y)
	// result: (Rsh16Ux64 x (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpRsh16Ux64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpRsh16Ux32(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux32 x y)
	// result: (Rsh16Ux64 x (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpRsh16Ux64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpRsh16Ux64(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux64 <t> x (Const64 [c]))
	// cond: uint64(c) < 64
	// result: (SRLDconst <t> [c] (ZeroExt16to64 x))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != ssaop.OpConst64 {
			break
		}
		c := ssa.AuxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.Reset(ssaop.OpSPARC64SRLDconst)
		v.Type = t
		v.AuxInt = ssa.Int64ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16Ux64 _ (Const64 [c]))
	// cond: uint64(c) >= 64
	// result: (MOVDconst [0])
	for {
		if v_1.Op != ssaop.OpConst64 {
			break
		}
		c := ssa.AuxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 64) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVDconst)
		v.AuxInt = ssa.Int64ToAuxInt(0)
		return true
	}
	// match: (Rsh16Ux64 <t> x y)
	// result: (AND (NEG <t> (LessThanU <typ.UInt64> (CMPconst <types.TypeFlags> [64] y))) (SRLD <t> (ZeroExt16to64 x) y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64AND)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64NEG, t)
		v1 := b.NewValue0(v.Pos, ssaop.OpSPARC64LessThanU, typ.UInt64)
		v2 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMPconst, types.TypeFlags)
		v2.AuxInt = ssa.Int64ToAuxInt(64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, ssaop.OpSPARC64SRLD, t)
		v4 := b.NewValue0(v.Pos, ssaop.OpZeroExt16to64, typ.UInt64)
		v4.AddArg(x)
		v3.AddArg2(v4, y)
		v.AddArg2(v0, v3)
		return true
	}
}
func rewriteValue_OpRsh16Ux8(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux8 x y)
	// result: (Rsh16Ux64 x (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpRsh16Ux64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpRsh16x16(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x16 x y)
	// result: (Rsh16x64 x (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpRsh16x64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpRsh16x32(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x32 x y)
	// result: (Rsh16x64 x (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpRsh16x64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpRsh16x64(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x64 <t> x (Const64 [c]))
	// cond: uint64(c) < 64
	// result: (SRADconst <t> [c] (SignExt16to64 x))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != ssaop.OpConst64 {
			break
		}
		c := ssa.AuxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.Reset(ssaop.OpSPARC64SRADconst)
		v.Type = t
		v.AuxInt = ssa.Int64ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, ssaop.OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16x64 <t> x (Const64 [c]))
	// cond: uint64(c) >= 64
	// result: (SRADconst <t> [63] (SignExt16to64 x))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != ssaop.OpConst64 {
			break
		}
		c := ssa.AuxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 64) {
			break
		}
		v.Reset(ssaop.OpSPARC64SRADconst)
		v.Type = t
		v.AuxInt = ssa.Int64ToAuxInt(63)
		v0 := b.NewValue0(v.Pos, ssaop.OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh16x64 <t> x y)
	// result: (SRAD <t> (SignExt16to64 x) (OR <typ.UInt64> (NEG <typ.UInt64> (GreaterThanU <typ.UInt64> (CMPconst <types.TypeFlags> [63] y))) y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64SRAD)
		v.Type = t
		v0 := b.NewValue0(v.Pos, ssaop.OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, ssaop.OpSPARC64OR, typ.UInt64)
		v2 := b.NewValue0(v.Pos, ssaop.OpSPARC64NEG, typ.UInt64)
		v3 := b.NewValue0(v.Pos, ssaop.OpSPARC64GreaterThanU, typ.UInt64)
		v4 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMPconst, types.TypeFlags)
		v4.AuxInt = ssa.Int64ToAuxInt(63)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg2(v2, y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValue_OpRsh16x8(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x8 x y)
	// result: (Rsh16x64 x (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpRsh16x64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpRsh32Ux16(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux16 x y)
	// result: (Rsh32Ux64 x (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpRsh32Ux64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpRsh32Ux32(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux32 x y)
	// result: (Rsh32Ux64 x (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpRsh32Ux64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpRsh32Ux64(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux64 <t> x (Const64 [c]))
	// cond: uint64(c) < 64
	// result: (SRLDconst <t> [c] (ZeroExt32to64 x))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != ssaop.OpConst64 {
			break
		}
		c := ssa.AuxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.Reset(ssaop.OpSPARC64SRLDconst)
		v.Type = t
		v.AuxInt = ssa.Int64ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh32Ux64 _ (Const64 [c]))
	// cond: uint64(c) >= 64
	// result: (MOVDconst [0])
	for {
		if v_1.Op != ssaop.OpConst64 {
			break
		}
		c := ssa.AuxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 64) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVDconst)
		v.AuxInt = ssa.Int64ToAuxInt(0)
		return true
	}
	// match: (Rsh32Ux64 <t> x y)
	// result: (AND (NEG <t> (LessThanU <typ.UInt64> (CMPconst <types.TypeFlags> [64] y))) (SRLD <t> (ZeroExt32to64 x) y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64AND)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64NEG, t)
		v1 := b.NewValue0(v.Pos, ssaop.OpSPARC64LessThanU, typ.UInt64)
		v2 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMPconst, types.TypeFlags)
		v2.AuxInt = ssa.Int64ToAuxInt(64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, ssaop.OpSPARC64SRLD, t)
		v4 := b.NewValue0(v.Pos, ssaop.OpZeroExt32to64, typ.UInt64)
		v4.AddArg(x)
		v3.AddArg2(v4, y)
		v.AddArg2(v0, v3)
		return true
	}
}
func rewriteValue_OpRsh32Ux8(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux8 x y)
	// result: (Rsh32Ux64 x (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpRsh32Ux64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpRsh32x16(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x16 x y)
	// result: (Rsh32x64 x (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpRsh32x64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpRsh32x32(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x32 x y)
	// result: (Rsh32x64 x (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpRsh32x64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpRsh32x64(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x64 <t> x (Const64 [c]))
	// cond: uint64(c) < 64
	// result: (SRADconst <t> [c] (SignExt32to64 x))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != ssaop.OpConst64 {
			break
		}
		c := ssa.AuxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.Reset(ssaop.OpSPARC64SRADconst)
		v.Type = t
		v.AuxInt = ssa.Int64ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, ssaop.OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh32x64 <t> x (Const64 [c]))
	// cond: uint64(c) >= 64
	// result: (SRADconst <t> [63] (SignExt32to64 x))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != ssaop.OpConst64 {
			break
		}
		c := ssa.AuxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 64) {
			break
		}
		v.Reset(ssaop.OpSPARC64SRADconst)
		v.Type = t
		v.AuxInt = ssa.Int64ToAuxInt(63)
		v0 := b.NewValue0(v.Pos, ssaop.OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh32x64 <t> x y)
	// result: (SRAD <t> (SignExt32to64 x) (OR <typ.UInt64> (NEG <typ.UInt64> (GreaterThanU <typ.UInt64> (CMPconst <types.TypeFlags> [63] y))) y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64SRAD)
		v.Type = t
		v0 := b.NewValue0(v.Pos, ssaop.OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, ssaop.OpSPARC64OR, typ.UInt64)
		v2 := b.NewValue0(v.Pos, ssaop.OpSPARC64NEG, typ.UInt64)
		v3 := b.NewValue0(v.Pos, ssaop.OpSPARC64GreaterThanU, typ.UInt64)
		v4 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMPconst, types.TypeFlags)
		v4.AuxInt = ssa.Int64ToAuxInt(63)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg2(v2, y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValue_OpRsh32x8(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x8 x y)
	// result: (Rsh32x64 x (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpRsh32x64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpRsh64Ux16(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux16 x y)
	// result: (Rsh64Ux64 x (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpRsh64Ux64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpRsh64Ux32(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux32 x y)
	// result: (Rsh64Ux64 x (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpRsh64Ux64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpRsh64Ux64(v *ssa.Value) bool {
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
		if v_1.Op != ssaop.OpConst64 {
			break
		}
		c := ssa.AuxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.Reset(ssaop.OpSPARC64SRLDconst)
		v.Type = t
		v.AuxInt = ssa.Int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (Rsh64Ux64 _ (Const64 [c]))
	// cond: uint64(c) >= 64
	// result: (MOVDconst [0])
	for {
		if v_1.Op != ssaop.OpConst64 {
			break
		}
		c := ssa.AuxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 64) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVDconst)
		v.AuxInt = ssa.Int64ToAuxInt(0)
		return true
	}
	// match: (Rsh64Ux64 <t> x y)
	// result: (AND (NEG <t> (LessThanU <typ.UInt64> (CMPconst <types.TypeFlags> [64] y))) (SRLD <t> x y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64AND)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64NEG, t)
		v1 := b.NewValue0(v.Pos, ssaop.OpSPARC64LessThanU, typ.UInt64)
		v2 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMPconst, types.TypeFlags)
		v2.AuxInt = ssa.Int64ToAuxInt(64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, ssaop.OpSPARC64SRLD, t)
		v3.AddArg2(x, y)
		v.AddArg2(v0, v3)
		return true
	}
}
func rewriteValue_OpRsh64Ux8(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux8 x y)
	// result: (Rsh64Ux64 x (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpRsh64Ux64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpRsh64x16(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64x16 x y)
	// result: (Rsh64x64 x (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpRsh64x64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpRsh64x32(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64x32 x y)
	// result: (Rsh64x64 x (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpRsh64x64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpRsh64x64(v *ssa.Value) bool {
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
		if v_1.Op != ssaop.OpConst64 {
			break
		}
		c := ssa.AuxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.Reset(ssaop.OpSPARC64SRADconst)
		v.Type = t
		v.AuxInt = ssa.Int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (Rsh64x64 <t> x (Const64 [c]))
	// cond: uint64(c) >= 64
	// result: (SRADconst <t> [63] x)
	for {
		t := v.Type
		x := v_0
		if v_1.Op != ssaop.OpConst64 {
			break
		}
		c := ssa.AuxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 64) {
			break
		}
		v.Reset(ssaop.OpSPARC64SRADconst)
		v.Type = t
		v.AuxInt = ssa.Int64ToAuxInt(63)
		v.AddArg(x)
		return true
	}
	// match: (Rsh64x64 <t> x y)
	// result: (SRAD <t> x (OR <typ.UInt64> (NEG <typ.UInt64> (GreaterThanU <typ.UInt64> (CMPconst <types.TypeFlags> [63] y))) y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64SRAD)
		v.Type = t
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64OR, typ.UInt64)
		v1 := b.NewValue0(v.Pos, ssaop.OpSPARC64NEG, typ.UInt64)
		v2 := b.NewValue0(v.Pos, ssaop.OpSPARC64GreaterThanU, typ.UInt64)
		v3 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMPconst, types.TypeFlags)
		v3.AuxInt = ssa.Int64ToAuxInt(63)
		v3.AddArg(y)
		v2.AddArg(v3)
		v1.AddArg(v2)
		v0.AddArg2(v1, y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpRsh64x8(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64x8 x y)
	// result: (Rsh64x64 x (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpRsh64x64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpRsh8Ux16(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux16 x y)
	// result: (Rsh8Ux64 x (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpRsh8Ux64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpRsh8Ux32(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux32 x y)
	// result: (Rsh8Ux64 x (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpRsh8Ux64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpRsh8Ux64(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux64 <t> x (Const64 [c]))
	// cond: uint64(c) < 64
	// result: (SRLDconst <t> [c] (ZeroExt8to64 x))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != ssaop.OpConst64 {
			break
		}
		c := ssa.AuxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.Reset(ssaop.OpSPARC64SRLDconst)
		v.Type = t
		v.AuxInt = ssa.Int64ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8Ux64 _ (Const64 [c]))
	// cond: uint64(c) >= 64
	// result: (MOVDconst [0])
	for {
		if v_1.Op != ssaop.OpConst64 {
			break
		}
		c := ssa.AuxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 64) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVDconst)
		v.AuxInt = ssa.Int64ToAuxInt(0)
		return true
	}
	// match: (Rsh8Ux64 <t> x y)
	// result: (AND (NEG <t> (LessThanU <typ.UInt64> (CMPconst <types.TypeFlags> [64] y))) (SRLD <t> (ZeroExt8to64 x) y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64AND)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64NEG, t)
		v1 := b.NewValue0(v.Pos, ssaop.OpSPARC64LessThanU, typ.UInt64)
		v2 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMPconst, types.TypeFlags)
		v2.AuxInt = ssa.Int64ToAuxInt(64)
		v2.AddArg(y)
		v1.AddArg(v2)
		v0.AddArg(v1)
		v3 := b.NewValue0(v.Pos, ssaop.OpSPARC64SRLD, t)
		v4 := b.NewValue0(v.Pos, ssaop.OpZeroExt8to64, typ.UInt64)
		v4.AddArg(x)
		v3.AddArg2(v4, y)
		v.AddArg2(v0, v3)
		return true
	}
}
func rewriteValue_OpRsh8Ux8(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux8 x y)
	// result: (Rsh8Ux64 x (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpRsh8Ux64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpRsh8x16(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x16 x y)
	// result: (Rsh8x64 x (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpRsh8x64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpRsh8x32(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x32 x y)
	// result: (Rsh8x64 x (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpRsh8x64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpRsh8x64(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x64 <t> x (Const64 [c]))
	// cond: uint64(c) < 64
	// result: (SRADconst <t> [c] (SignExt8to64 x))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != ssaop.OpConst64 {
			break
		}
		c := ssa.AuxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.Reset(ssaop.OpSPARC64SRADconst)
		v.Type = t
		v.AuxInt = ssa.Int64ToAuxInt(c)
		v0 := b.NewValue0(v.Pos, ssaop.OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8x64 <t> x (Const64 [c]))
	// cond: uint64(c) >= 64
	// result: (SRADconst <t> [63] (SignExt8to64 x))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != ssaop.OpConst64 {
			break
		}
		c := ssa.AuxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 64) {
			break
		}
		v.Reset(ssaop.OpSPARC64SRADconst)
		v.Type = t
		v.AuxInt = ssa.Int64ToAuxInt(63)
		v0 := b.NewValue0(v.Pos, ssaop.OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (Rsh8x64 <t> x y)
	// result: (SRAD <t> (SignExt8to64 x) (OR <typ.UInt64> (NEG <typ.UInt64> (GreaterThanU <typ.UInt64> (CMPconst <types.TypeFlags> [63] y))) y))
	for {
		t := v.Type
		x := v_0
		y := v_1
		v.Reset(ssaop.OpSPARC64SRAD)
		v.Type = t
		v0 := b.NewValue0(v.Pos, ssaop.OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, ssaop.OpSPARC64OR, typ.UInt64)
		v2 := b.NewValue0(v.Pos, ssaop.OpSPARC64NEG, typ.UInt64)
		v3 := b.NewValue0(v.Pos, ssaop.OpSPARC64GreaterThanU, typ.UInt64)
		v4 := b.NewValue0(v.Pos, ssaop.OpSPARC64CMPconst, types.TypeFlags)
		v4.AuxInt = ssa.Int64ToAuxInt(63)
		v4.AddArg(y)
		v3.AddArg(v4)
		v2.AddArg(v3)
		v1.AddArg2(v2, y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValue_OpRsh8x8(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x8 x y)
	// result: (Rsh8x64 x (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.Reset(ssaop.OpRsh8x64)
		v0 := b.NewValue0(v.Pos, ssaop.OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValue_OpSPARC64ADD(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (ADD (MOVDconst [c]) x)
	// cond: ssa.Is13Bit(c)
	// result: (ADDconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != ssaop.OpSPARC64MOVDconst {
				continue
			}
			c := ssa.AuxIntToInt64(v_0.AuxInt)
			x := v_1
			if !(ssa.Is13Bit(c)) {
				continue
			}
			v.Reset(ssaop.OpSPARC64ADDconst)
			v.AuxInt = ssa.Int64ToAuxInt(c)
			v.AddArg(x)
			return true
		}
		break
	}
	return false
}
func rewriteValue_OpSPARC64ADDconst(v *ssa.Value) bool {
	v_0 := v.Args[0]
	// match: (ADDconst [c] (MOVDaddr [d] {sym} x))
	// cond: ssa.Is32Bit(c+int64(d))
	// result: (MOVDaddr [int32(c+int64(d))] {sym} x)
	for {
		c := ssa.AuxIntToInt64(v.AuxInt)
		if v_0.Op != ssaop.OpSPARC64MOVDaddr {
			break
		}
		d := ssa.AuxIntToInt32(v_0.AuxInt)
		sym := ssa.AuxToSym(v_0.Aux)
		x := v_0.Args[0]
		if !(ssa.Is32Bit(c + int64(d))) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVDaddr)
		v.AuxInt = ssa.Int32ToAuxInt(int32(c + int64(d)))
		v.Aux = ssa.SymToAux(sym)
		v.AddArg(x)
		return true
	}
	// match: (ADDconst [0] x)
	// result: x
	for {
		if ssa.AuxIntToInt64(v.AuxInt) != 0 {
			break
		}
		x := v_0
		v.CopyOf(x)
		return true
	}
	return false
}
func rewriteValue_OpSPARC64AND(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (AND (MOVDconst [c]) x)
	// cond: ssa.Is13Bit(c)
	// result: (ANDconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != ssaop.OpSPARC64MOVDconst {
				continue
			}
			c := ssa.AuxIntToInt64(v_0.AuxInt)
			x := v_1
			if !(ssa.Is13Bit(c)) {
				continue
			}
			v.Reset(ssaop.OpSPARC64ANDconst)
			v.AuxInt = ssa.Int64ToAuxInt(c)
			v.AddArg(x)
			return true
		}
		break
	}
	return false
}
func rewriteValue_OpSPARC64ANDconst(v *ssa.Value) bool {
	v_0 := v.Args[0]
	// match: (ANDconst [-1] x)
	// result: x
	for {
		if ssa.AuxIntToInt64(v.AuxInt) != -1 {
			break
		}
		x := v_0
		v.CopyOf(x)
		return true
	}
	return false
}
func rewriteValue_OpSPARC64CMP(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (CMP x (MOVDconst [c]))
	// cond: ssa.Is13Bit(c)
	// result: (CMPconst [c] x)
	for {
		x := v_0
		if v_1.Op != ssaop.OpSPARC64MOVDconst {
			break
		}
		c := ssa.AuxIntToInt64(v_1.AuxInt)
		if !(ssa.Is13Bit(c)) {
			break
		}
		v.Reset(ssaop.OpSPARC64CMPconst)
		v.AuxInt = ssa.Int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	// match: (CMP x (MOVUW (MOVDconst [c])))
	// cond: ssa.Is13Bit(int64(uint32(c)))
	// result: (CMPconst [int64(uint32(c))] x)
	for {
		x := v_0
		if v_1.Op != ssaop.OpSPARC64MOVUW {
			break
		}
		v_1_0 := v_1.Args[0]
		if v_1_0.Op != ssaop.OpSPARC64MOVDconst {
			break
		}
		c := ssa.AuxIntToInt64(v_1_0.AuxInt)
		if !(ssa.Is13Bit(int64(uint32(c)))) {
			break
		}
		v.Reset(ssaop.OpSPARC64CMPconst)
		v.AuxInt = ssa.Int64ToAuxInt(int64(uint32(c)))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValue_OpSPARC64FMOVDload(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (FMOVDload [off1] {sym1} (MOVDaddr [off2] {sym2} ptr) mem)
	// cond: ssa.CanMergeSym(sym1, sym2) && sym2 != nil && ptr.Op != ssaop.OpSB && ssa.Is32Bit(int64(off1)+int64(off2))
	// result: (FMOVDload [off1+off2] {ssa.MergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := ssa.AuxIntToInt32(v.AuxInt)
		sym1 := ssa.AuxToSym(v.Aux)
		if v_0.Op != ssaop.OpSPARC64MOVDaddr {
			break
		}
		off2 := ssa.AuxIntToInt32(v_0.AuxInt)
		sym2 := ssa.AuxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(ssa.CanMergeSym(sym1, sym2) && sym2 != nil && ptr.Op != ssaop.OpSB && ssa.Is32Bit(int64(off1)+int64(off2))) {
			break
		}
		v.Reset(ssaop.OpSPARC64FMOVDload)
		v.AuxInt = ssa.Int32ToAuxInt(off1 + off2)
		v.Aux = ssa.SymToAux(ssa.MergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (FMOVDload [off1] {sym} (ADDconst [off2] ptr) mem)
	// cond: ssa.Is32Bit(int64(off1)+off2) && ptr.Op != ssaop.OpSB && ptr.Op != ssaop.OpSP
	// result: (FMOVDload [off1+int32(off2)] {sym} ptr mem)
	for {
		off1 := ssa.AuxIntToInt32(v.AuxInt)
		sym := ssa.AuxToSym(v.Aux)
		if v_0.Op != ssaop.OpSPARC64ADDconst {
			break
		}
		off2 := ssa.AuxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(ssa.Is32Bit(int64(off1)+off2) && ptr.Op != ssaop.OpSB && ptr.Op != ssaop.OpSP) {
			break
		}
		v.Reset(ssaop.OpSPARC64FMOVDload)
		v.AuxInt = ssa.Int32ToAuxInt(off1 + int32(off2))
		v.Aux = ssa.SymToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValue_OpSPARC64FMOVDstore(v *ssa.Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (FMOVDstore [off1] {sym1} (MOVDaddr [off2] {sym2} ptr) val mem)
	// cond: ssa.CanMergeSym(sym1, sym2) && sym2 != nil && ptr.Op != ssaop.OpSB && ssa.Is32Bit(int64(off1)+int64(off2))
	// result: (FMOVDstore [off1+off2] {ssa.MergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := ssa.AuxIntToInt32(v.AuxInt)
		sym1 := ssa.AuxToSym(v.Aux)
		if v_0.Op != ssaop.OpSPARC64MOVDaddr {
			break
		}
		off2 := ssa.AuxIntToInt32(v_0.AuxInt)
		sym2 := ssa.AuxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(ssa.CanMergeSym(sym1, sym2) && sym2 != nil && ptr.Op != ssaop.OpSB && ssa.Is32Bit(int64(off1)+int64(off2))) {
			break
		}
		v.Reset(ssaop.OpSPARC64FMOVDstore)
		v.AuxInt = ssa.Int32ToAuxInt(off1 + off2)
		v.Aux = ssa.SymToAux(ssa.MergeSym(sym1, sym2))
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (FMOVDstore [off1] {sym} (ADDconst [off2] ptr) val mem)
	// cond: ssa.Is32Bit(int64(off1)+off2) && ptr.Op != ssaop.OpSB && ptr.Op != ssaop.OpSP
	// result: (FMOVDstore [off1+int32(off2)] {sym} ptr val mem)
	for {
		off1 := ssa.AuxIntToInt32(v.AuxInt)
		sym := ssa.AuxToSym(v.Aux)
		if v_0.Op != ssaop.OpSPARC64ADDconst {
			break
		}
		off2 := ssa.AuxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(ssa.Is32Bit(int64(off1)+off2) && ptr.Op != ssaop.OpSB && ptr.Op != ssaop.OpSP) {
			break
		}
		v.Reset(ssaop.OpSPARC64FMOVDstore)
		v.AuxInt = ssa.Int32ToAuxInt(off1 + int32(off2))
		v.Aux = ssa.SymToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValue_OpSPARC64FMOVSload(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (FMOVSload [off1] {sym1} (MOVDaddr [off2] {sym2} ptr) mem)
	// cond: ssa.CanMergeSym(sym1, sym2) && sym2 != nil && ptr.Op != ssaop.OpSB && ssa.Is32Bit(int64(off1)+int64(off2))
	// result: (FMOVSload [off1+off2] {ssa.MergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := ssa.AuxIntToInt32(v.AuxInt)
		sym1 := ssa.AuxToSym(v.Aux)
		if v_0.Op != ssaop.OpSPARC64MOVDaddr {
			break
		}
		off2 := ssa.AuxIntToInt32(v_0.AuxInt)
		sym2 := ssa.AuxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(ssa.CanMergeSym(sym1, sym2) && sym2 != nil && ptr.Op != ssaop.OpSB && ssa.Is32Bit(int64(off1)+int64(off2))) {
			break
		}
		v.Reset(ssaop.OpSPARC64FMOVSload)
		v.AuxInt = ssa.Int32ToAuxInt(off1 + off2)
		v.Aux = ssa.SymToAux(ssa.MergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (FMOVSload [off1] {sym} (ADDconst [off2] ptr) mem)
	// cond: ssa.Is32Bit(int64(off1)+off2) && ptr.Op != ssaop.OpSB && ptr.Op != ssaop.OpSP
	// result: (FMOVSload [off1+int32(off2)] {sym} ptr mem)
	for {
		off1 := ssa.AuxIntToInt32(v.AuxInt)
		sym := ssa.AuxToSym(v.Aux)
		if v_0.Op != ssaop.OpSPARC64ADDconst {
			break
		}
		off2 := ssa.AuxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(ssa.Is32Bit(int64(off1)+off2) && ptr.Op != ssaop.OpSB && ptr.Op != ssaop.OpSP) {
			break
		}
		v.Reset(ssaop.OpSPARC64FMOVSload)
		v.AuxInt = ssa.Int32ToAuxInt(off1 + int32(off2))
		v.Aux = ssa.SymToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValue_OpSPARC64FMOVSstore(v *ssa.Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (FMOVSstore [off1] {sym1} (MOVDaddr [off2] {sym2} ptr) val mem)
	// cond: ssa.CanMergeSym(sym1, sym2) && sym2 != nil && ptr.Op != ssaop.OpSB && ssa.Is32Bit(int64(off1)+int64(off2))
	// result: (FMOVSstore [off1+off2] {ssa.MergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := ssa.AuxIntToInt32(v.AuxInt)
		sym1 := ssa.AuxToSym(v.Aux)
		if v_0.Op != ssaop.OpSPARC64MOVDaddr {
			break
		}
		off2 := ssa.AuxIntToInt32(v_0.AuxInt)
		sym2 := ssa.AuxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(ssa.CanMergeSym(sym1, sym2) && sym2 != nil && ptr.Op != ssaop.OpSB && ssa.Is32Bit(int64(off1)+int64(off2))) {
			break
		}
		v.Reset(ssaop.OpSPARC64FMOVSstore)
		v.AuxInt = ssa.Int32ToAuxInt(off1 + off2)
		v.Aux = ssa.SymToAux(ssa.MergeSym(sym1, sym2))
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (FMOVSstore [off1] {sym} (ADDconst [off2] ptr) val mem)
	// cond: ssa.Is32Bit(int64(off1)+off2) && ptr.Op != ssaop.OpSB && ptr.Op != ssaop.OpSP
	// result: (FMOVSstore [off1+int32(off2)] {sym} ptr val mem)
	for {
		off1 := ssa.AuxIntToInt32(v.AuxInt)
		sym := ssa.AuxToSym(v.Aux)
		if v_0.Op != ssaop.OpSPARC64ADDconst {
			break
		}
		off2 := ssa.AuxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(ssa.Is32Bit(int64(off1)+off2) && ptr.Op != ssaop.OpSB && ptr.Op != ssaop.OpSP) {
			break
		}
		v.Reset(ssaop.OpSPARC64FMOVSstore)
		v.AuxInt = ssa.Int32ToAuxInt(off1 + int32(off2))
		v.Aux = ssa.SymToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValue_OpSPARC64LoweredPanicBoundsCR(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (LoweredPanicBoundsCR [kind] {p} (MOVDconst [c]) mem)
	// result: (LoweredPanicBoundsCC [kind] {ssa.PanicBoundsCC{Cx:p.C, Cy:c}} mem)
	for {
		kind := ssa.AuxIntToInt64(v.AuxInt)
		p := ssa.AuxToPanicBoundsC(v.Aux)
		if v_0.Op != ssaop.OpSPARC64MOVDconst {
			break
		}
		c := ssa.AuxIntToInt64(v_0.AuxInt)
		mem := v_1
		v.Reset(ssaop.OpSPARC64LoweredPanicBoundsCC)
		v.AuxInt = ssa.Int64ToAuxInt(kind)
		v.Aux = ssa.PanicBoundsCCToAux(ssa.PanicBoundsCC{Cx: p.C, Cy: c})
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue_OpSPARC64LoweredPanicBoundsRC(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (LoweredPanicBoundsRC [kind] {p} (MOVDconst [c]) mem)
	// result: (LoweredPanicBoundsCC [kind] {ssa.PanicBoundsCC{Cx:c, Cy:p.C}} mem)
	for {
		kind := ssa.AuxIntToInt64(v.AuxInt)
		p := ssa.AuxToPanicBoundsC(v.Aux)
		if v_0.Op != ssaop.OpSPARC64MOVDconst {
			break
		}
		c := ssa.AuxIntToInt64(v_0.AuxInt)
		mem := v_1
		v.Reset(ssaop.OpSPARC64LoweredPanicBoundsCC)
		v.AuxInt = ssa.Int64ToAuxInt(kind)
		v.Aux = ssa.PanicBoundsCCToAux(ssa.PanicBoundsCC{Cx: c, Cy: p.C})
		v.AddArg(mem)
		return true
	}
	return false
}
func rewriteValue_OpSPARC64LoweredPanicBoundsRR(v *ssa.Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (LoweredPanicBoundsRR [kind] x (MOVDconst [c]) mem)
	// result: (LoweredPanicBoundsRC [kind] x {ssa.PanicBoundsC{C:c}} mem)
	for {
		kind := ssa.AuxIntToInt64(v.AuxInt)
		x := v_0
		if v_1.Op != ssaop.OpSPARC64MOVDconst {
			break
		}
		c := ssa.AuxIntToInt64(v_1.AuxInt)
		mem := v_2
		v.Reset(ssaop.OpSPARC64LoweredPanicBoundsRC)
		v.AuxInt = ssa.Int64ToAuxInt(kind)
		v.Aux = ssa.PanicBoundsCToAux(ssa.PanicBoundsC{C: c})
		v.AddArg2(x, mem)
		return true
	}
	// match: (LoweredPanicBoundsRR [kind] (MOVDconst [c]) y mem)
	// result: (LoweredPanicBoundsCR [kind] {ssa.PanicBoundsC{C:c}} y mem)
	for {
		kind := ssa.AuxIntToInt64(v.AuxInt)
		if v_0.Op != ssaop.OpSPARC64MOVDconst {
			break
		}
		c := ssa.AuxIntToInt64(v_0.AuxInt)
		y := v_1
		mem := v_2
		v.Reset(ssaop.OpSPARC64LoweredPanicBoundsCR)
		v.AuxInt = ssa.Int64ToAuxInt(kind)
		v.Aux = ssa.PanicBoundsCToAux(ssa.PanicBoundsC{C: c})
		v.AddArg2(y, mem)
		return true
	}
	return false
}
func rewriteValue_OpSPARC64MOVB(v *ssa.Value) bool {
	v_0 := v.Args[0]
	// match: (MOVB (MOVDconst [c]))
	// result: (MOVDconst [int64(int8(c))])
	for {
		if v_0.Op != ssaop.OpSPARC64MOVDconst {
			break
		}
		c := ssa.AuxIntToInt64(v_0.AuxInt)
		v.Reset(ssaop.OpSPARC64MOVDconst)
		v.AuxInt = ssa.Int64ToAuxInt(int64(int8(c)))
		return true
	}
	return false
}
func rewriteValue_OpSPARC64MOVBload(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBload [off1] {sym1} (MOVDaddr [off2] {sym2} ptr) mem)
	// cond: ssa.CanMergeSym(sym1, sym2) && sym2 != nil && ptr.Op != ssaop.OpSB && ssa.Is32Bit(int64(off1)+int64(off2))
	// result: (MOVBload [off1+off2] {ssa.MergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := ssa.AuxIntToInt32(v.AuxInt)
		sym1 := ssa.AuxToSym(v.Aux)
		if v_0.Op != ssaop.OpSPARC64MOVDaddr {
			break
		}
		off2 := ssa.AuxIntToInt32(v_0.AuxInt)
		sym2 := ssa.AuxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(ssa.CanMergeSym(sym1, sym2) && sym2 != nil && ptr.Op != ssaop.OpSB && ssa.Is32Bit(int64(off1)+int64(off2))) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVBload)
		v.AuxInt = ssa.Int32ToAuxInt(off1 + off2)
		v.Aux = ssa.SymToAux(ssa.MergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVBload [off1] {sym} (ADDconst [off2] ptr) mem)
	// cond: ssa.Is32Bit(int64(off1)+off2) && ptr.Op != ssaop.OpSB && ptr.Op != ssaop.OpSP
	// result: (MOVBload [off1+int32(off2)] {sym} ptr mem)
	for {
		off1 := ssa.AuxIntToInt32(v.AuxInt)
		sym := ssa.AuxToSym(v.Aux)
		if v_0.Op != ssaop.OpSPARC64ADDconst {
			break
		}
		off2 := ssa.AuxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(ssa.Is32Bit(int64(off1)+off2) && ptr.Op != ssaop.OpSB && ptr.Op != ssaop.OpSP) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVBload)
		v.AuxInt = ssa.Int32ToAuxInt(off1 + int32(off2))
		v.Aux = ssa.SymToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValue_OpSPARC64MOVBstore(v *ssa.Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVBstore [off1] {sym1} (MOVDaddr [off2] {sym2} ptr) val mem)
	// cond: ssa.CanMergeSym(sym1, sym2) && sym2 != nil && ptr.Op != ssaop.OpSB && ssa.Is32Bit(int64(off1)+int64(off2))
	// result: (MOVBstore [off1+off2] {ssa.MergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := ssa.AuxIntToInt32(v.AuxInt)
		sym1 := ssa.AuxToSym(v.Aux)
		if v_0.Op != ssaop.OpSPARC64MOVDaddr {
			break
		}
		off2 := ssa.AuxIntToInt32(v_0.AuxInt)
		sym2 := ssa.AuxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(ssa.CanMergeSym(sym1, sym2) && sym2 != nil && ptr.Op != ssaop.OpSB && ssa.Is32Bit(int64(off1)+int64(off2))) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVBstore)
		v.AuxInt = ssa.Int32ToAuxInt(off1 + off2)
		v.Aux = ssa.SymToAux(ssa.MergeSym(sym1, sym2))
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVBstore [off1] {sym} (ADDconst [off2] ptr) val mem)
	// cond: ssa.Is32Bit(int64(off1)+off2) && ptr.Op != ssaop.OpSB && ptr.Op != ssaop.OpSP
	// result: (MOVBstore [off1+int32(off2)] {sym} ptr val mem)
	for {
		off1 := ssa.AuxIntToInt32(v.AuxInt)
		sym := ssa.AuxToSym(v.Aux)
		if v_0.Op != ssaop.OpSPARC64ADDconst {
			break
		}
		off2 := ssa.AuxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(ssa.Is32Bit(int64(off1)+off2) && ptr.Op != ssaop.OpSB && ptr.Op != ssaop.OpSP) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVBstore)
		v.AuxInt = ssa.Int32ToAuxInt(off1 + int32(off2))
		v.Aux = ssa.SymToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValue_OpSPARC64MOVD(v *ssa.Value) bool {
	v_0 := v.Args[0]
	// match: (MOVD (MOVDconst [c]))
	// result: (MOVDconst [c])
	for {
		if v_0.Op != ssaop.OpSPARC64MOVDconst {
			break
		}
		c := ssa.AuxIntToInt64(v_0.AuxInt)
		v.Reset(ssaop.OpSPARC64MOVDconst)
		v.AuxInt = ssa.Int64ToAuxInt(c)
		return true
	}
	return false
}
func rewriteValue_OpSPARC64MOVDload(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVDload [off1] {sym1} (MOVDaddr [off2] {sym2} ptr) mem)
	// cond: ssa.CanMergeSym(sym1, sym2) && sym2 != nil && ptr.Op != ssaop.OpSB && ssa.Is32Bit(int64(off1)+int64(off2))
	// result: (MOVDload [off1+off2] {ssa.MergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := ssa.AuxIntToInt32(v.AuxInt)
		sym1 := ssa.AuxToSym(v.Aux)
		if v_0.Op != ssaop.OpSPARC64MOVDaddr {
			break
		}
		off2 := ssa.AuxIntToInt32(v_0.AuxInt)
		sym2 := ssa.AuxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(ssa.CanMergeSym(sym1, sym2) && sym2 != nil && ptr.Op != ssaop.OpSB && ssa.Is32Bit(int64(off1)+int64(off2))) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVDload)
		v.AuxInt = ssa.Int32ToAuxInt(off1 + off2)
		v.Aux = ssa.SymToAux(ssa.MergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVDload [off1] {sym} (ADDconst [off2] ptr) mem)
	// cond: ssa.Is32Bit(int64(off1)+off2) && ptr.Op != ssaop.OpSB && ptr.Op != ssaop.OpSP
	// result: (MOVDload [off1+int32(off2)] {sym} ptr mem)
	for {
		off1 := ssa.AuxIntToInt32(v.AuxInt)
		sym := ssa.AuxToSym(v.Aux)
		if v_0.Op != ssaop.OpSPARC64ADDconst {
			break
		}
		off2 := ssa.AuxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(ssa.Is32Bit(int64(off1)+off2) && ptr.Op != ssaop.OpSB && ptr.Op != ssaop.OpSP) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVDload)
		v.AuxInt = ssa.Int32ToAuxInt(off1 + int32(off2))
		v.Aux = ssa.SymToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValue_OpSPARC64MOVDstore(v *ssa.Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVDstore [off1] {sym1} (MOVDaddr [off2] {sym2} ptr) val mem)
	// cond: ssa.CanMergeSym(sym1, sym2) && sym2 != nil && ptr.Op != ssaop.OpSB && ssa.Is32Bit(int64(off1)+int64(off2))
	// result: (MOVDstore [off1+off2] {ssa.MergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := ssa.AuxIntToInt32(v.AuxInt)
		sym1 := ssa.AuxToSym(v.Aux)
		if v_0.Op != ssaop.OpSPARC64MOVDaddr {
			break
		}
		off2 := ssa.AuxIntToInt32(v_0.AuxInt)
		sym2 := ssa.AuxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(ssa.CanMergeSym(sym1, sym2) && sym2 != nil && ptr.Op != ssaop.OpSB && ssa.Is32Bit(int64(off1)+int64(off2))) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVDstore)
		v.AuxInt = ssa.Int32ToAuxInt(off1 + off2)
		v.Aux = ssa.SymToAux(ssa.MergeSym(sym1, sym2))
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVDstore [off1] {sym} (ADDconst [off2] ptr) val mem)
	// cond: ssa.Is32Bit(int64(off1)+off2) && ptr.Op != ssaop.OpSB && ptr.Op != ssaop.OpSP
	// result: (MOVDstore [off1+int32(off2)] {sym} ptr val mem)
	for {
		off1 := ssa.AuxIntToInt32(v.AuxInt)
		sym := ssa.AuxToSym(v.Aux)
		if v_0.Op != ssaop.OpSPARC64ADDconst {
			break
		}
		off2 := ssa.AuxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(ssa.Is32Bit(int64(off1)+off2) && ptr.Op != ssaop.OpSB && ptr.Op != ssaop.OpSP) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVDstore)
		v.AuxInt = ssa.Int32ToAuxInt(off1 + int32(off2))
		v.Aux = ssa.SymToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValue_OpSPARC64MOVH(v *ssa.Value) bool {
	v_0 := v.Args[0]
	// match: (MOVH (MOVDconst [c]))
	// result: (MOVDconst [int64(int16(c))])
	for {
		if v_0.Op != ssaop.OpSPARC64MOVDconst {
			break
		}
		c := ssa.AuxIntToInt64(v_0.AuxInt)
		v.Reset(ssaop.OpSPARC64MOVDconst)
		v.AuxInt = ssa.Int64ToAuxInt(int64(int16(c)))
		return true
	}
	return false
}
func rewriteValue_OpSPARC64MOVHload(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHload [off1] {sym1} (MOVDaddr [off2] {sym2} ptr) mem)
	// cond: ssa.CanMergeSym(sym1, sym2) && sym2 != nil && ptr.Op != ssaop.OpSB && ssa.Is32Bit(int64(off1)+int64(off2))
	// result: (MOVHload [off1+off2] {ssa.MergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := ssa.AuxIntToInt32(v.AuxInt)
		sym1 := ssa.AuxToSym(v.Aux)
		if v_0.Op != ssaop.OpSPARC64MOVDaddr {
			break
		}
		off2 := ssa.AuxIntToInt32(v_0.AuxInt)
		sym2 := ssa.AuxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(ssa.CanMergeSym(sym1, sym2) && sym2 != nil && ptr.Op != ssaop.OpSB && ssa.Is32Bit(int64(off1)+int64(off2))) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVHload)
		v.AuxInt = ssa.Int32ToAuxInt(off1 + off2)
		v.Aux = ssa.SymToAux(ssa.MergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVHload [off1] {sym} (ADDconst [off2] ptr) mem)
	// cond: ssa.Is32Bit(int64(off1)+off2) && ptr.Op != ssaop.OpSB && ptr.Op != ssaop.OpSP
	// result: (MOVHload [off1+int32(off2)] {sym} ptr mem)
	for {
		off1 := ssa.AuxIntToInt32(v.AuxInt)
		sym := ssa.AuxToSym(v.Aux)
		if v_0.Op != ssaop.OpSPARC64ADDconst {
			break
		}
		off2 := ssa.AuxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(ssa.Is32Bit(int64(off1)+off2) && ptr.Op != ssaop.OpSB && ptr.Op != ssaop.OpSP) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVHload)
		v.AuxInt = ssa.Int32ToAuxInt(off1 + int32(off2))
		v.Aux = ssa.SymToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValue_OpSPARC64MOVHstore(v *ssa.Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVHstore [off1] {sym1} (MOVDaddr [off2] {sym2} ptr) val mem)
	// cond: ssa.CanMergeSym(sym1, sym2) && sym2 != nil && ptr.Op != ssaop.OpSB && ssa.Is32Bit(int64(off1)+int64(off2))
	// result: (MOVHstore [off1+off2] {ssa.MergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := ssa.AuxIntToInt32(v.AuxInt)
		sym1 := ssa.AuxToSym(v.Aux)
		if v_0.Op != ssaop.OpSPARC64MOVDaddr {
			break
		}
		off2 := ssa.AuxIntToInt32(v_0.AuxInt)
		sym2 := ssa.AuxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(ssa.CanMergeSym(sym1, sym2) && sym2 != nil && ptr.Op != ssaop.OpSB && ssa.Is32Bit(int64(off1)+int64(off2))) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVHstore)
		v.AuxInt = ssa.Int32ToAuxInt(off1 + off2)
		v.Aux = ssa.SymToAux(ssa.MergeSym(sym1, sym2))
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVHstore [off1] {sym} (ADDconst [off2] ptr) val mem)
	// cond: ssa.Is32Bit(int64(off1)+off2) && ptr.Op != ssaop.OpSB && ptr.Op != ssaop.OpSP
	// result: (MOVHstore [off1+int32(off2)] {sym} ptr val mem)
	for {
		off1 := ssa.AuxIntToInt32(v.AuxInt)
		sym := ssa.AuxToSym(v.Aux)
		if v_0.Op != ssaop.OpSPARC64ADDconst {
			break
		}
		off2 := ssa.AuxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(ssa.Is32Bit(int64(off1)+off2) && ptr.Op != ssaop.OpSB && ptr.Op != ssaop.OpSP) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVHstore)
		v.AuxInt = ssa.Int32ToAuxInt(off1 + int32(off2))
		v.Aux = ssa.SymToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValue_OpSPARC64MOVUB(v *ssa.Value) bool {
	v_0 := v.Args[0]
	// match: (MOVUB (MOVDconst [c]))
	// result: (MOVDconst [int64(uint8(c))])
	for {
		if v_0.Op != ssaop.OpSPARC64MOVDconst {
			break
		}
		c := ssa.AuxIntToInt64(v_0.AuxInt)
		v.Reset(ssaop.OpSPARC64MOVDconst)
		v.AuxInt = ssa.Int64ToAuxInt(int64(uint8(c)))
		return true
	}
	return false
}
func rewriteValue_OpSPARC64MOVUBload(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVUBload [off1] {sym1} (MOVDaddr [off2] {sym2} ptr) mem)
	// cond: ssa.CanMergeSym(sym1, sym2) && sym2 != nil && ptr.Op != ssaop.OpSB && ssa.Is32Bit(int64(off1)+int64(off2))
	// result: (MOVUBload [off1+off2] {ssa.MergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := ssa.AuxIntToInt32(v.AuxInt)
		sym1 := ssa.AuxToSym(v.Aux)
		if v_0.Op != ssaop.OpSPARC64MOVDaddr {
			break
		}
		off2 := ssa.AuxIntToInt32(v_0.AuxInt)
		sym2 := ssa.AuxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(ssa.CanMergeSym(sym1, sym2) && sym2 != nil && ptr.Op != ssaop.OpSB && ssa.Is32Bit(int64(off1)+int64(off2))) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVUBload)
		v.AuxInt = ssa.Int32ToAuxInt(off1 + off2)
		v.Aux = ssa.SymToAux(ssa.MergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVUBload [off1] {sym} (ADDconst [off2] ptr) mem)
	// cond: ssa.Is32Bit(int64(off1)+off2) && ptr.Op != ssaop.OpSB && ptr.Op != ssaop.OpSP
	// result: (MOVUBload [off1+int32(off2)] {sym} ptr mem)
	for {
		off1 := ssa.AuxIntToInt32(v.AuxInt)
		sym := ssa.AuxToSym(v.Aux)
		if v_0.Op != ssaop.OpSPARC64ADDconst {
			break
		}
		off2 := ssa.AuxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(ssa.Is32Bit(int64(off1)+off2) && ptr.Op != ssaop.OpSB && ptr.Op != ssaop.OpSP) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVUBload)
		v.AuxInt = ssa.Int32ToAuxInt(off1 + int32(off2))
		v.Aux = ssa.SymToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValue_OpSPARC64MOVUH(v *ssa.Value) bool {
	v_0 := v.Args[0]
	// match: (MOVUH (MOVDconst [c]))
	// result: (MOVDconst [int64(uint16(c))])
	for {
		if v_0.Op != ssaop.OpSPARC64MOVDconst {
			break
		}
		c := ssa.AuxIntToInt64(v_0.AuxInt)
		v.Reset(ssaop.OpSPARC64MOVDconst)
		v.AuxInt = ssa.Int64ToAuxInt(int64(uint16(c)))
		return true
	}
	return false
}
func rewriteValue_OpSPARC64MOVUHload(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVUHload [off1] {sym1} (MOVDaddr [off2] {sym2} ptr) mem)
	// cond: ssa.CanMergeSym(sym1, sym2) && sym2 != nil && ptr.Op != ssaop.OpSB && ssa.Is32Bit(int64(off1)+int64(off2))
	// result: (MOVUHload [off1+off2] {ssa.MergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := ssa.AuxIntToInt32(v.AuxInt)
		sym1 := ssa.AuxToSym(v.Aux)
		if v_0.Op != ssaop.OpSPARC64MOVDaddr {
			break
		}
		off2 := ssa.AuxIntToInt32(v_0.AuxInt)
		sym2 := ssa.AuxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(ssa.CanMergeSym(sym1, sym2) && sym2 != nil && ptr.Op != ssaop.OpSB && ssa.Is32Bit(int64(off1)+int64(off2))) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVUHload)
		v.AuxInt = ssa.Int32ToAuxInt(off1 + off2)
		v.Aux = ssa.SymToAux(ssa.MergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVUHload [off1] {sym} (ADDconst [off2] ptr) mem)
	// cond: ssa.Is32Bit(int64(off1)+off2) && ptr.Op != ssaop.OpSB && ptr.Op != ssaop.OpSP
	// result: (MOVUHload [off1+int32(off2)] {sym} ptr mem)
	for {
		off1 := ssa.AuxIntToInt32(v.AuxInt)
		sym := ssa.AuxToSym(v.Aux)
		if v_0.Op != ssaop.OpSPARC64ADDconst {
			break
		}
		off2 := ssa.AuxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(ssa.Is32Bit(int64(off1)+off2) && ptr.Op != ssaop.OpSB && ptr.Op != ssaop.OpSP) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVUHload)
		v.AuxInt = ssa.Int32ToAuxInt(off1 + int32(off2))
		v.Aux = ssa.SymToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValue_OpSPARC64MOVUW(v *ssa.Value) bool {
	v_0 := v.Args[0]
	// match: (MOVUW (MOVDconst [c]))
	// result: (MOVDconst [int64(uint32(c))])
	for {
		if v_0.Op != ssaop.OpSPARC64MOVDconst {
			break
		}
		c := ssa.AuxIntToInt64(v_0.AuxInt)
		v.Reset(ssaop.OpSPARC64MOVDconst)
		v.AuxInt = ssa.Int64ToAuxInt(int64(uint32(c)))
		return true
	}
	return false
}
func rewriteValue_OpSPARC64MOVUWload(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVUWload [off1] {sym1} (MOVDaddr [off2] {sym2} ptr) mem)
	// cond: ssa.CanMergeSym(sym1, sym2) && sym2 != nil && ptr.Op != ssaop.OpSB && ssa.Is32Bit(int64(off1)+int64(off2))
	// result: (MOVUWload [off1+off2] {ssa.MergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := ssa.AuxIntToInt32(v.AuxInt)
		sym1 := ssa.AuxToSym(v.Aux)
		if v_0.Op != ssaop.OpSPARC64MOVDaddr {
			break
		}
		off2 := ssa.AuxIntToInt32(v_0.AuxInt)
		sym2 := ssa.AuxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(ssa.CanMergeSym(sym1, sym2) && sym2 != nil && ptr.Op != ssaop.OpSB && ssa.Is32Bit(int64(off1)+int64(off2))) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVUWload)
		v.AuxInt = ssa.Int32ToAuxInt(off1 + off2)
		v.Aux = ssa.SymToAux(ssa.MergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVUWload [off1] {sym} (ADDconst [off2] ptr) mem)
	// cond: ssa.Is32Bit(int64(off1)+off2) && ptr.Op != ssaop.OpSB && ptr.Op != ssaop.OpSP
	// result: (MOVUWload [off1+int32(off2)] {sym} ptr mem)
	for {
		off1 := ssa.AuxIntToInt32(v.AuxInt)
		sym := ssa.AuxToSym(v.Aux)
		if v_0.Op != ssaop.OpSPARC64ADDconst {
			break
		}
		off2 := ssa.AuxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(ssa.Is32Bit(int64(off1)+off2) && ptr.Op != ssaop.OpSB && ptr.Op != ssaop.OpSP) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVUWload)
		v.AuxInt = ssa.Int32ToAuxInt(off1 + int32(off2))
		v.Aux = ssa.SymToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValue_OpSPARC64MOVW(v *ssa.Value) bool {
	v_0 := v.Args[0]
	// match: (MOVW (MOVDconst [c]))
	// result: (MOVDconst [int64(int32(c))])
	for {
		if v_0.Op != ssaop.OpSPARC64MOVDconst {
			break
		}
		c := ssa.AuxIntToInt64(v_0.AuxInt)
		v.Reset(ssaop.OpSPARC64MOVDconst)
		v.AuxInt = ssa.Int64ToAuxInt(int64(int32(c)))
		return true
	}
	return false
}
func rewriteValue_OpSPARC64MOVWload(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWload [off1] {sym1} (MOVDaddr [off2] {sym2} ptr) mem)
	// cond: ssa.CanMergeSym(sym1, sym2) && sym2 != nil && ptr.Op != ssaop.OpSB && ssa.Is32Bit(int64(off1)+int64(off2))
	// result: (MOVWload [off1+off2] {ssa.MergeSym(sym1,sym2)} ptr mem)
	for {
		off1 := ssa.AuxIntToInt32(v.AuxInt)
		sym1 := ssa.AuxToSym(v.Aux)
		if v_0.Op != ssaop.OpSPARC64MOVDaddr {
			break
		}
		off2 := ssa.AuxIntToInt32(v_0.AuxInt)
		sym2 := ssa.AuxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		mem := v_1
		if !(ssa.CanMergeSym(sym1, sym2) && sym2 != nil && ptr.Op != ssaop.OpSB && ssa.Is32Bit(int64(off1)+int64(off2))) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVWload)
		v.AuxInt = ssa.Int32ToAuxInt(off1 + off2)
		v.Aux = ssa.SymToAux(ssa.MergeSym(sym1, sym2))
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (MOVWload [off1] {sym} (ADDconst [off2] ptr) mem)
	// cond: ssa.Is32Bit(int64(off1)+off2) && ptr.Op != ssaop.OpSB && ptr.Op != ssaop.OpSP
	// result: (MOVWload [off1+int32(off2)] {sym} ptr mem)
	for {
		off1 := ssa.AuxIntToInt32(v.AuxInt)
		sym := ssa.AuxToSym(v.Aux)
		if v_0.Op != ssaop.OpSPARC64ADDconst {
			break
		}
		off2 := ssa.AuxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(ssa.Is32Bit(int64(off1)+off2) && ptr.Op != ssaop.OpSB && ptr.Op != ssaop.OpSP) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVWload)
		v.AuxInt = ssa.Int32ToAuxInt(off1 + int32(off2))
		v.Aux = ssa.SymToAux(sym)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValue_OpSPARC64MOVWstore(v *ssa.Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (MOVWstore [off1] {sym1} (MOVDaddr [off2] {sym2} ptr) val mem)
	// cond: ssa.CanMergeSym(sym1, sym2) && sym2 != nil && ptr.Op != ssaop.OpSB && ssa.Is32Bit(int64(off1)+int64(off2))
	// result: (MOVWstore [off1+off2] {ssa.MergeSym(sym1,sym2)} ptr val mem)
	for {
		off1 := ssa.AuxIntToInt32(v.AuxInt)
		sym1 := ssa.AuxToSym(v.Aux)
		if v_0.Op != ssaop.OpSPARC64MOVDaddr {
			break
		}
		off2 := ssa.AuxIntToInt32(v_0.AuxInt)
		sym2 := ssa.AuxToSym(v_0.Aux)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(ssa.CanMergeSym(sym1, sym2) && sym2 != nil && ptr.Op != ssaop.OpSB && ssa.Is32Bit(int64(off1)+int64(off2))) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVWstore)
		v.AuxInt = ssa.Int32ToAuxInt(off1 + off2)
		v.Aux = ssa.SymToAux(ssa.MergeSym(sym1, sym2))
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (MOVWstore [off1] {sym} (ADDconst [off2] ptr) val mem)
	// cond: ssa.Is32Bit(int64(off1)+off2) && ptr.Op != ssaop.OpSB && ptr.Op != ssaop.OpSP
	// result: (MOVWstore [off1+int32(off2)] {sym} ptr val mem)
	for {
		off1 := ssa.AuxIntToInt32(v.AuxInt)
		sym := ssa.AuxToSym(v.Aux)
		if v_0.Op != ssaop.OpSPARC64ADDconst {
			break
		}
		off2 := ssa.AuxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(ssa.Is32Bit(int64(off1)+off2) && ptr.Op != ssaop.OpSB && ptr.Op != ssaop.OpSP) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVWstore)
		v.AuxInt = ssa.Int32ToAuxInt(off1 + int32(off2))
		v.Aux = ssa.SymToAux(sym)
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValue_OpSPARC64OR(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (OR (MOVDconst [c]) x)
	// cond: ssa.Is13Bit(c)
	// result: (ORconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != ssaop.OpSPARC64MOVDconst {
				continue
			}
			c := ssa.AuxIntToInt64(v_0.AuxInt)
			x := v_1
			if !(ssa.Is13Bit(c)) {
				continue
			}
			v.Reset(ssaop.OpSPARC64ORconst)
			v.AuxInt = ssa.Int64ToAuxInt(c)
			v.AddArg(x)
			return true
		}
		break
	}
	return false
}
func rewriteValue_OpSPARC64ORconst(v *ssa.Value) bool {
	v_0 := v.Args[0]
	// match: (ORconst [0] x)
	// result: x
	for {
		if ssa.AuxIntToInt64(v.AuxInt) != 0 {
			break
		}
		x := v_0
		v.CopyOf(x)
		return true
	}
	return false
}
func rewriteValue_OpSPARC64SLLD(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SLLD x (MOVDconst [c]))
	// result: (SLLDconst [c&63] x)
	for {
		x := v_0
		if v_1.Op != ssaop.OpSPARC64MOVDconst {
			break
		}
		c := ssa.AuxIntToInt64(v_1.AuxInt)
		v.Reset(ssaop.OpSPARC64SLLDconst)
		v.AuxInt = ssa.Int64ToAuxInt(c & 63)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValue_OpSPARC64SRAD(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SRAD x (MOVDconst [c]))
	// result: (SRADconst [c&63] x)
	for {
		x := v_0
		if v_1.Op != ssaop.OpSPARC64MOVDconst {
			break
		}
		c := ssa.AuxIntToInt64(v_1.AuxInt)
		v.Reset(ssaop.OpSPARC64SRADconst)
		v.AuxInt = ssa.Int64ToAuxInt(c & 63)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValue_OpSPARC64SRLD(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SRLD x (MOVDconst [c]))
	// result: (SRLDconst [c&63] x)
	for {
		x := v_0
		if v_1.Op != ssaop.OpSPARC64MOVDconst {
			break
		}
		c := ssa.AuxIntToInt64(v_1.AuxInt)
		v.Reset(ssaop.OpSPARC64SRLDconst)
		v.AuxInt = ssa.Int64ToAuxInt(c & 63)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValue_OpSPARC64SUB(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (SUB x (MOVDconst [c]))
	// cond: ssa.Is13Bit(c)
	// result: (SUBconst [c] x)
	for {
		x := v_0
		if v_1.Op != ssaop.OpSPARC64MOVDconst {
			break
		}
		c := ssa.AuxIntToInt64(v_1.AuxInt)
		if !(ssa.Is13Bit(c)) {
			break
		}
		v.Reset(ssaop.OpSPARC64SUBconst)
		v.AuxInt = ssa.Int64ToAuxInt(c)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValue_OpSPARC64SUBconst(v *ssa.Value) bool {
	v_0 := v.Args[0]
	// match: (SUBconst [0] x)
	// result: x
	for {
		if ssa.AuxIntToInt64(v.AuxInt) != 0 {
			break
		}
		x := v_0
		v.CopyOf(x)
		return true
	}
	return false
}
func rewriteValue_OpSPARC64XOR(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (XOR (MOVDconst [c]) x)
	// cond: ssa.Is13Bit(c)
	// result: (XORconst [c] x)
	for {
		for _i0 := 0; _i0 <= 1; _i0, v_0, v_1 = _i0+1, v_1, v_0 {
			if v_0.Op != ssaop.OpSPARC64MOVDconst {
				continue
			}
			c := ssa.AuxIntToInt64(v_0.AuxInt)
			x := v_1
			if !(ssa.Is13Bit(c)) {
				continue
			}
			v.Reset(ssaop.OpSPARC64XORconst)
			v.AuxInt = ssa.Int64ToAuxInt(c)
			v.AddArg(x)
			return true
		}
		break
	}
	return false
}
func rewriteValue_OpSPARC64XORconst(v *ssa.Value) bool {
	v_0 := v.Args[0]
	// match: (XORconst [0] x)
	// result: x
	for {
		if ssa.AuxIntToInt64(v.AuxInt) != 0 {
			break
		}
		x := v_0
		v.CopyOf(x)
		return true
	}
	return false
}
func rewriteValue_OpSlicemask(v *ssa.Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	// match: (Slicemask <t> x)
	// result: (SRADconst <t> [63] (NEG <t> x))
	for {
		t := v.Type
		x := v_0
		v.Reset(ssaop.OpSPARC64SRADconst)
		v.Type = t
		v.AuxInt = ssa.Int64ToAuxInt(63)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64NEG, t)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValue_OpStore(v *ssa.Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Store {t} ptr val mem)
	// cond: t.Size() == 1
	// result: (MOVBstore ptr val mem)
	for {
		t := ssa.AuxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.Size() == 1) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVBstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.Size() == 2
	// result: (MOVHstore ptr val mem)
	for {
		t := ssa.AuxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.Size() == 2) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVHstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.Size() == 4 && !t.IsFloat()
	// result: (MOVWstore ptr val mem)
	for {
		t := ssa.AuxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.Size() == 4 && !t.IsFloat()) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVWstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.Size() == 8 && !t.IsFloat()
	// result: (MOVDstore ptr val mem)
	for {
		t := ssa.AuxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.Size() == 8 && !t.IsFloat()) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVDstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.Size() == 4 && t.IsFloat()
	// result: (FMOVSstore ptr val mem)
	for {
		t := ssa.AuxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.Size() == 4 && t.IsFloat()) {
			break
		}
		v.Reset(ssaop.OpSPARC64FMOVSstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.Size() == 8 && t.IsFloat()
	// result: (FMOVDstore ptr val mem)
	for {
		t := ssa.AuxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.Size() == 8 && t.IsFloat()) {
			break
		}
		v.Reset(ssaop.OpSPARC64FMOVDstore)
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValue_OpZero(v *ssa.Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Zero [0] _ mem)
	// result: mem
	for {
		if ssa.AuxIntToInt64(v.AuxInt) != 0 {
			break
		}
		mem := v_1
		v.CopyOf(mem)
		return true
	}
	// match: (Zero [1] ptr mem)
	// result: (MOVBstore ptr (MOVDconst [0]) mem)
	for {
		if ssa.AuxIntToInt64(v.AuxInt) != 1 {
			break
		}
		ptr := v_0
		mem := v_1
		v.Reset(ssaop.OpSPARC64MOVBstore)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64MOVDconst, typ.UInt64)
		v0.AuxInt = ssa.Int64ToAuxInt(0)
		v.AddArg3(ptr, v0, mem)
		return true
	}
	// match: (Zero [2] {t} ptr mem)
	// cond: t.Alignment()%2 == 0
	// result: (MOVHstore ptr (MOVDconst [0]) mem)
	for {
		if ssa.AuxIntToInt64(v.AuxInt) != 2 {
			break
		}
		t := ssa.AuxToType(v.Aux)
		ptr := v_0
		mem := v_1
		if !(t.Alignment()%2 == 0) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVHstore)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64MOVDconst, typ.UInt64)
		v0.AuxInt = ssa.Int64ToAuxInt(0)
		v.AddArg3(ptr, v0, mem)
		return true
	}
	// match: (Zero [4] {t} ptr mem)
	// cond: t.Alignment()%4 == 0
	// result: (MOVWstore ptr (MOVDconst [0]) mem)
	for {
		if ssa.AuxIntToInt64(v.AuxInt) != 4 {
			break
		}
		t := ssa.AuxToType(v.Aux)
		ptr := v_0
		mem := v_1
		if !(t.Alignment()%4 == 0) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVWstore)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64MOVDconst, typ.UInt64)
		v0.AuxInt = ssa.Int64ToAuxInt(0)
		v.AddArg3(ptr, v0, mem)
		return true
	}
	// match: (Zero [8] {t} ptr mem)
	// cond: t.Alignment()%8 == 0
	// result: (MOVDstore ptr (MOVDconst [0]) mem)
	for {
		if ssa.AuxIntToInt64(v.AuxInt) != 8 {
			break
		}
		t := ssa.AuxToType(v.Aux)
		ptr := v_0
		mem := v_1
		if !(t.Alignment()%8 == 0) {
			break
		}
		v.Reset(ssaop.OpSPARC64MOVDstore)
		v0 := b.NewValue0(v.Pos, ssaop.OpSPARC64MOVDconst, typ.UInt64)
		v0.AuxInt = ssa.Int64ToAuxInt(0)
		v.AddArg3(ptr, v0, mem)
		return true
	}
	// match: (Zero [s] {t} ptr mem)
	// cond: t.Alignment()%8 == 0
	// result: (LoweredZero [s<<4|8] ptr mem)
	for {
		s := ssa.AuxIntToInt64(v.AuxInt)
		t := ssa.AuxToType(v.Aux)
		ptr := v_0
		mem := v_1
		if !(t.Alignment()%8 == 0) {
			break
		}
		v.Reset(ssaop.OpSPARC64LoweredZero)
		v.AuxInt = ssa.Int64ToAuxInt(s<<4 | 8)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Zero [s] {t} ptr mem)
	// cond: t.Alignment()%4 == 0
	// result: (LoweredZero [s<<4|4] ptr mem)
	for {
		s := ssa.AuxIntToInt64(v.AuxInt)
		t := ssa.AuxToType(v.Aux)
		ptr := v_0
		mem := v_1
		if !(t.Alignment()%4 == 0) {
			break
		}
		v.Reset(ssaop.OpSPARC64LoweredZero)
		v.AuxInt = ssa.Int64ToAuxInt(s<<4 | 4)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Zero [s] {t} ptr mem)
	// cond: t.Alignment()%2 == 0
	// result: (LoweredZero [s<<4|2] ptr mem)
	for {
		s := ssa.AuxIntToInt64(v.AuxInt)
		t := ssa.AuxToType(v.Aux)
		ptr := v_0
		mem := v_1
		if !(t.Alignment()%2 == 0) {
			break
		}
		v.Reset(ssaop.OpSPARC64LoweredZero)
		v.AuxInt = ssa.Int64ToAuxInt(s<<4 | 2)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Zero [s] ptr mem)
	// result: (LoweredZero [s<<4|1] ptr mem)
	for {
		s := ssa.AuxIntToInt64(v.AuxInt)
		ptr := v_0
		mem := v_1
		v.Reset(ssaop.OpSPARC64LoweredZero)
		v.AuxInt = ssa.Int64ToAuxInt(s<<4 | 1)
		v.AddArg2(ptr, mem)
		return true
	}
}
func RewriteBlock(b *ssa.Block) bool {
	switch b.Kind {
	case block.BlockIf:
		// match: (If (Equal cc) yes no)
		// result: (EQ cc yes no)
		for b.Controls[0].Op == ssaop.OpSPARC64Equal {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.ResetWithControl(block.BlockSPARC64EQ, cc)
			return true
		}
		// match: (If (NotEqual cc) yes no)
		// result: (NE cc yes no)
		for b.Controls[0].Op == ssaop.OpSPARC64NotEqual {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.ResetWithControl(block.BlockSPARC64NE, cc)
			return true
		}
		// match: (If (LessThan cc) yes no)
		// result: (LT cc yes no)
		for b.Controls[0].Op == ssaop.OpSPARC64LessThan {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.ResetWithControl(block.BlockSPARC64LT, cc)
			return true
		}
		// match: (If (LessEqual cc) yes no)
		// result: (LE cc yes no)
		for b.Controls[0].Op == ssaop.OpSPARC64LessEqual {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.ResetWithControl(block.BlockSPARC64LE, cc)
			return true
		}
		// match: (If (GreaterThan cc) yes no)
		// result: (GT cc yes no)
		for b.Controls[0].Op == ssaop.OpSPARC64GreaterThan {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.ResetWithControl(block.BlockSPARC64GT, cc)
			return true
		}
		// match: (If (GreaterEqual cc) yes no)
		// result: (GE cc yes no)
		for b.Controls[0].Op == ssaop.OpSPARC64GreaterEqual {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.ResetWithControl(block.BlockSPARC64GE, cc)
			return true
		}
		// match: (If (LessThanU cc) yes no)
		// result: (ULT cc yes no)
		for b.Controls[0].Op == ssaop.OpSPARC64LessThanU {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.ResetWithControl(block.BlockSPARC64ULT, cc)
			return true
		}
		// match: (If (LessEqualU cc) yes no)
		// result: (ULE cc yes no)
		for b.Controls[0].Op == ssaop.OpSPARC64LessEqualU {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.ResetWithControl(block.BlockSPARC64ULE, cc)
			return true
		}
		// match: (If (GreaterThanU cc) yes no)
		// result: (UGT cc yes no)
		for b.Controls[0].Op == ssaop.OpSPARC64GreaterThanU {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.ResetWithControl(block.BlockSPARC64UGT, cc)
			return true
		}
		// match: (If (GreaterEqualU cc) yes no)
		// result: (UGE cc yes no)
		for b.Controls[0].Op == ssaop.OpSPARC64GreaterEqualU {
			v_0 := b.Controls[0]
			cc := v_0.Args[0]
			b.ResetWithControl(block.BlockSPARC64UGE, cc)
			return true
		}
		// match: (If cond yes no)
		// result: (NE (CMPconst [0] cond) yes no)
		for {
			cond := b.Controls[0]
			v0 := b.NewValue0(cond.Pos, ssaop.OpSPARC64CMPconst, types.TypeFlags)
			v0.AuxInt = ssa.Int64ToAuxInt(0)
			v0.AddArg(cond)
			b.ResetWithControl(block.BlockSPARC64NE, v0)
			return true
		}
	}
	return false
}
