// Code generated from _gen/SPARC64.rules using 'go generate'; DO NOT EDIT.

package ssa

import "cmd/compile/internal/types"
import "cmd/compile/internal/ssa/block"

func rewriteValueSPARC64(v *Value) bool {
	switch v.Op {
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
	case OpCvt32Fto64:
		v.Op = OpSPARC64FSTOX
		return true
	case OpCvt32Fto64F:
		v.Op = OpSPARC64FSTOD
		return true
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
	case OpDiv32F:
		v.Op = OpSPARC64FDIVS
		return true
	case OpDiv64:
		return rewriteValueSPARC64_OpDiv64(v)
	case OpDiv64F:
		v.Op = OpSPARC64FDIVD
		return true
	case OpDiv64u:
		return rewriteValueSPARC64_OpDiv64u(v)
	case OpEq64:
		return rewriteValueSPARC64_OpEq64(v)
	case OpGetCallerPC:
		v.Op = OpSPARC64LoweredGetCallerPC
		return true
	case OpGetCallerSP:
		v.Op = OpSPARC64LoweredGetCallerSP
		return true
	case OpGetClosurePtr:
		v.Op = OpSPARC64LoweredGetClosurePtr
		return true
	case OpInterCall:
		v.Op = OpSPARC64CALLinter
		return true
	case OpIsInBounds:
		return rewriteValueSPARC64_OpIsInBounds(v)
	case OpIsNonNil:
		return rewriteValueSPARC64_OpIsNonNil(v)
	case OpIsSliceInBounds:
		return rewriteValueSPARC64_OpIsSliceInBounds(v)
	case OpLeq64:
		return rewriteValueSPARC64_OpLeq64(v)
	case OpLess64:
		return rewriteValueSPARC64_OpLess64(v)
	case OpLoad:
		return rewriteValueSPARC64_OpLoad(v)
	case OpLocalAddr:
		return rewriteValueSPARC64_OpLocalAddr(v)
	case OpLsh64x64:
		return rewriteValueSPARC64_OpLsh64x64(v)
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
	case OpNeq64:
		return rewriteValueSPARC64_OpNeq64(v)
	case OpNilCheck:
		v.Op = OpSPARC64LoweredNilCheck
		return true
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
	case OpPubBarrier:
		v.Op = OpSPARC64LoweredPubBarrier
		return true
	case OpRsh64Ux64:
		return rewriteValueSPARC64_OpRsh64Ux64(v)
	case OpRsh64x64:
		return rewriteValueSPARC64_OpRsh64x64(v)
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
func rewriteValueSPARC64_OpLsh64x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Lsh64x64 x y)
	// result: (SLLD x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64SLLD)
		v.AddArg2(x, y)
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
func rewriteValueSPARC64_OpRsh64Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Rsh64Ux64 x y)
	// result: (SRLD x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64SRLD)
		v.AddArg2(x, y)
		return true
	}
}
func rewriteValueSPARC64_OpRsh64x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Rsh64x64 x y)
	// result: (SRAD x y)
	for {
		x := v_0
		y := v_1
		v.reset(OpSPARC64SRAD)
		v.AddArg2(x, y)
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
