// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sparc64asm

import (
	"encoding/binary"
	"errors"
)

// ErrTruncated is returned when the instruction is cut short.
var ErrTruncated = errors.New("truncated instruction")

// ErrUnrecognized is returned when the encoding is not a known instruction.
var ErrUnrecognized = errors.New("unrecognized instruction")

// Decode decodes the 4-byte big-endian SPARC V9 instruction at the start
// of src. Every SPARC instruction is exactly four bytes.
func Decode(src []byte) (Inst, error) {
	if len(src) < 4 {
		return Inst{}, ErrTruncated
	}
	x := binary.BigEndian.Uint32(src)
	inst := Inst{Enc: x}

	switch x >> 30 {
	case 0:
		return decodeFormat2(inst, x)
	case 1:
		// CALL: 30-bit word displacement.
		inst.Op = CALL
		inst.Args[0] = PCRel(int64(int32(x<<2)) )
		return inst, nil
	default:
		return decodeFormat3(inst, x)
	}
}

// sign extends the low n bits of v.
func signExtend(v uint32, n uint) int64 {
	shift := 32 - n
	return int64(int32(v<<shift) >> shift)
}

var branchIcc = [16]Op{BN, BE, BLE, BL, BLEU, BCS, BNEG, BVS,
	BA, BNE, BG, BGE, BGU, BCC, BPOS, BVC}

var branchFcc = [16]Op{FBN, FBNE, FBLG, FBUL, FBL, FBUG, FBG, FBU,
	FBA, FBE, FBUE, FBGE, FBUGE, FBLE, FBULE, FBO}

var branchReg = [8]Op{0, BRZ, BRLEZ, BRLZ, 0, BRNZ, BRGZ, BRGEZ}

func decodeFormat2(inst Inst, x uint32) (Inst, error) {
	op2 := (x >> 22) & 7
	rd := Reg((x >> 25) & 0x1f)
	cond := (x >> 25) & 0xf

	switch op2 {
	case 0:
		inst.Op = ILLTRAP
		inst.Args[0] = Imm(x & 0x3fffff)
	case 1: // BPcc, with %icc/%xcc selector and prediction
		inst.Op = branchIcc[cond]
		cc := Reg(ICC)
		if (x>>21)&1 == 1 {
			cc = XCC
		}
		inst.Args[0] = cc
		inst.Args[1] = PCRel(signExtend(x&0x7ffff, 19) * 4)
	case 2: // Bicc
		inst.Op = branchIcc[cond]
		inst.Args[0] = PCRel(signExtend(x&0x3fffff, 22) * 4)
	case 3: // BPr, branch on register
		inst.Op = branchReg[(x>>25)&7]
		if inst.Op == 0 {
			return inst, ErrUnrecognized
		}
		d16 := ((x >> 20) & 3 << 14) | (x & 0x3fff)
		inst.Args[0] = Reg((x >> 14) & 0x1f)
		inst.Args[1] = PCRel(signExtend(d16, 16) * 4)
	case 4: // SETHI, and NOP as its degenerate form
		imm := x & 0x3fffff
		if rd == G0 && imm == 0 {
			inst.Op = NOP
			return inst, nil
		}
		inst.Op = SETHI
		inst.Args[0] = Imm(int64(imm))
		inst.Args[1] = rd
	case 5: // FBPfcc
		inst.Op = branchFcc[cond]
		inst.Args[0] = Reg(FCC0 + Reg((x>>20)&3))
		inst.Args[1] = PCRel(signExtend(x&0x7ffff, 19) * 4)
	case 6: // FBfcc
		inst.Op = branchFcc[cond]
		inst.Args[0] = PCRel(signExtend(x&0x3fffff, 22) * 4)
	default:
		return inst, ErrUnrecognized
	}
	return inst, nil
}

// arithOp3 maps op=2 op3 values to opcodes. Zero means unhandled.
var arithOp3 = [64]Op{
	0x00: ADD, 0x01: AND, 0x02: OR, 0x03: XOR,
	0x04: SUB, 0x05: ANDN, 0x06: ORN, 0x07: XNOR,
	0x08: ADDC, 0x09: MULX, 0x0a: UMUL, 0x0b: SMUL,
	0x0c: SUBC, 0x0d: UDIVX, 0x0e: UDIV, 0x0f: SDIV,
	0x10: ADDcc, 0x11: ANDcc, 0x12: ORcc, 0x13: XORcc,
	0x14: SUBcc, 0x15: ANDNcc, 0x16: ORNcc, 0x17: XNORcc,
	0x18: ADDCcc, 0x1a: UMULcc, 0x1b: SMULcc, 0x1c: SUBCcc,
	0x1e: UDIVcc, 0x1f: SDIVcc,
	0x2d: SDIVX, 0x2e: POPC,
	0x38: JMPL, 0x39: RETURN, 0x3c: SAVE, 0x3d: RESTORE,
}

// loadStoreOp3 maps op=3 op3 values to opcodes.
var loadStoreOp3 = [64]Op{
	0x00: LDUW, 0x01: LDUB, 0x02: LDUH, 0x03: LDD,
	0x04: STW, 0x05: STB, 0x06: STH, 0x07: STD,
	0x08: LDSW, 0x09: LDSB, 0x0a: LDSH, 0x0b: LDX,
	0x0e: STX,
	0x10: LDUWA, 0x11: LDUBA, 0x12: LDUHA, 0x13: LDDA,
	0x14: STWA, 0x15: STBA, 0x16: STHA, 0x17: STDA,
	0x18: LDSWA, 0x19: LDSBA, 0x1a: LDSHA, 0x1b: LDXA,
	0x1e: STXA,
	0x20: LDF, 0x23: LDDF, 0x24: STF, 0x27: STDF,
	0x2d: PREFETCH,
	0x3c: CASA, 0x3e: CASXA,
}

var movcc = [16]Op{MOVN, MOVE, MOVLE, MOVL, MOVLEU, MOVCS, MOVNEG, MOVVS,
	MOVA, MOVNE, MOVG, MOVGE, MOVGU, MOVCC, MOVPOS, MOVVC}

// fpop1 maps the opf field of FPop1 (op3=0x34) to opcodes.
var fpop1 = map[uint32]Op{
	0x001: FMOVs, 0x002: FMOVd, 0x005: FNEGs, 0x006: FNEGd,
	0x009: FABSs, 0x00a: FABSd, 0x029: FSQRTs, 0x02a: FSQRTd,
	0x041: FADDs, 0x042: FADDd, 0x045: FSUBs, 0x046: FSUBd,
	0x049: FMULs, 0x04a: FMULd, 0x04d: FDIVs, 0x04e: FDIVd,
	0x0c4: FITOS, 0x0c8: FITOD, 0x0c6: FDTOS, 0x0c9: FSTOD,
	0x0d1: FSTOI, 0x0d2: FDTOI, 0x081: FSTOX, 0x082: FDTOX,
	0x084: FXTOS, 0x088: FXTOD,
}

// fpop2 maps the opf field of FPop2 (op3=0x35) to compare opcodes.
var fpop2 = map[uint32]Op{
	0x051: FCMPs, 0x052: FCMPd, 0x055: FCMPEs, 0x056: FCMPEd,
}

// impdep1 maps the VIS3 opf values the port actually emits.
var impdep1 = map[uint32]Op{
	0x11: ADDXC, 0x13: ADDXCcc, 0x16: UMULXHI, 0x142: SHA256,
}

// isDoubleOp reports whether the opcode's register operands are
// double-precision, so the 5-bit field needs the V9 bit-5 fixup.
func isDoubleOp(o Op) bool {
	switch o {
	case FMOVd, FNEGd, FABSd, FSQRTd, FADDd, FSUBd, FMULd, FDIVd,
		FCMPd, FCMPEd, FDTOS, FSTOD, FDTOI, FDTOX, FXTOD, FITOD:
		return true
	}
	return false
}

// fpReg decodes a 5-bit float register field. For double-precision the
// V9 encoding folds bit 5 of the register number into bit 0.
func fpReg(f uint32, double bool) Reg {
	if double {
		n := (f & 0x1e) | ((f & 1) << 5)
		return RegD(uint8(n))
	}
	return RegF(uint8(f))
}

// fmovcc is the floating-point condition encoding for FMOVcc.
var fmovcc = [16]Op{MOVN, MOVNE, MOVLG, MOVUL, MOVL, MOVUG, MOVG, MOVU,
	MOVA, MOVE, MOVUE, MOVGE, MOVUGE, MOVLE, MOVULE, MOVO}

var movr = [8]Op{0, MOVRZ, MOVRLEZ, MOVRLZ, 0, MOVRNZ, MOVRGZ, MOVRGEZ}

func decodeFormat3(inst Inst, x uint32) (Inst, error) {
	op := x >> 30
	rd := Reg((x >> 25) & 0x1f)
	op3 := (x >> 19) & 0x3f
	rs1 := Reg((x >> 14) & 0x1f)
	imm := (x>>13)&1 == 1
	rs2 := Reg(x & 0x1f)
	simm13 := Imm(signExtend(x&0x1fff, 13))

	// second operand: register or immediate
	var second Arg = rs2
	if imm {
		second = simm13
	}

	if op == 3 {
		o := loadStoreOp3[op3]
		if o == 0 {
			return inst, ErrUnrecognized
		}
		inst.Op = o
		m := Mem{Base: rs1}
		if imm {
			m.Off = simm13
		} else {
			m.Ind = rs2
		}
		if op3&0x10 != 0 && op3 < 0x20 { // the *a alternate-space forms
			m.HasA = true
			m.ASI = uint8((x >> 5) & 0xff)
			m.Ind = rs2
			m.Off = 0
		}
		if o == CASA || o == CASXA {
			// cas [rs1], rs2, rd - the address has no displacement.
			inst.Args[0] = Mem{Base: rs1}
			inst.Args[1] = rs2
			inst.Args[2] = rd
			return inst, nil
		}
		reg := rd
		switch o {
		case LDF, STF:
			reg = RegF(uint8(rd))
		case LDDF, STDF:
			reg = fpReg(uint32(rd), true)
		}
		if isStore(o) {
			inst.Args[0] = reg
			inst.Args[1] = m
		} else {
			inst.Args[0] = m
			inst.Args[1] = reg
		}
		return inst, nil
	}

	// op == 2
	switch {
	case op3 == 0x25 || op3 == 0x26 || op3 == 0x27: // shifts
		xbit := (x>>12)&1 == 1
		switch op3 {
		case 0x25:
			inst.Op = SLL
			if xbit {
				inst.Op = SLLX
			}
		case 0x26:
			inst.Op = SRL
			if xbit {
				inst.Op = SRLX
			}
		default:
			inst.Op = SRA
			if xbit {
				inst.Op = SRAX
			}
		}
		if imm {
			if xbit {
				second = Imm(x & 0x3f)
			} else {
				second = Imm(x & 0x1f)
			}
		}
		inst.Args[0] = rs1
		inst.Args[1] = second
		inst.Args[2] = rd
		return inst, nil

	case op3 == 0x2c: // MOVcc / FMOVcc
		cond := (x >> 14) & 0xf
		cc2 := (x >> 18) & 1
		var cc Reg
		if cc2 == 1 {
			// Integer condition codes.
			inst.Op = movcc[cond]
			cc = ICC
			if (x>>12)&1 == 1 {
				cc = XCC
			}
		} else {
			// Floating-point condition codes use their own encoding,
			// in which 1 is "not equal" and 9 is "equal" - the reverse
			// of the integer table.
			inst.Op = fmovcc[cond]
			cc = FCC0 + Reg((x>>11)&3)
		}
		if imm {
			second = Imm(signExtend(x&0x7ff, 11))
		}
		inst.Args[0] = cc
		inst.Args[1] = second
		inst.Args[2] = rd
		return inst, nil

	case op3 == 0x2f: // MOVr
		inst.Op = movr[(x>>10)&7]
		if inst.Op == 0 {
			return inst, ErrUnrecognized
		}
		if imm {
			second = Imm(signExtend(x&0x3ff, 10))
		}
		inst.Args[0] = rs1
		inst.Args[1] = second
		inst.Args[2] = rd
		return inst, nil

	case op3 == 0x28: // RDasr family, plus MEMBAR/STBAR at rs1==15
		if rs1 == 15 && rd == 0 {
			if imm {
				inst.Op = MEMBAR
				inst.Args[0] = Imm(x & 0x7f)
			} else {
				inst.Op = STBAR
			}
			return inst, nil
		}
		switch rs1 {
		case 0:
			inst.Op = RDY
		case 2:
			inst.Op = RDCCR
		case 3:
			inst.Op = RDASI
		case 5:
			inst.Op = RDPC
		case 6:
			inst.Op = RDFPRS
		default:
			return inst, ErrUnrecognized
		}
		inst.Args[0] = rd
		return inst, nil

	case op3 == 0x30: // WRasr family
		switch rd {
		case 0:
			inst.Op = WRY
		case 2:
			inst.Op = WRCCR
		case 3:
			inst.Op = WRASI
		case 6:
			inst.Op = WRFPRS
		default:
			return inst, ErrUnrecognized
		}
		inst.Args[0] = rs1
		inst.Args[1] = second
		return inst, nil

	case op3 == 0x2b: // FLUSHW
		inst.Op = FLUSHW
		return inst, nil

	case op3 == 0x34: // FPop1
		opf := (x >> 5) & 0x1ff
		o, ok := fpop1[opf]
		if !ok {
			return inst, ErrUnrecognized
		}
		inst.Op = o
		d := isDoubleOp(o)
		// Conversions read and write different widths; decode each side
		// by what the opcode says rather than assuming a single width.
		srcDouble, dstDouble := d, d
		switch o {
		case FITOD, FSTOD, FXTOD:
			srcDouble = o == FXTOD
		case FDTOS, FDTOI, FDTOX:
			dstDouble = false
			srcDouble = true
		case FITOS, FSTOI, FSTOX:
			srcDouble, dstDouble = false, false
		case FXTOS:
			srcDouble, dstDouble = true, false
		}
		if o == FITOD || o == FITOS {
			srcDouble = false
		}
		inst.Args[0] = fpReg(x&0x1f, srcDouble)
		inst.Args[1] = fpReg(uint32(rd), dstDouble)
		if o == FADDs || o == FADDd || o == FSUBs || o == FSUBd ||
			o == FMULs || o == FMULd || o == FDIVs || o == FDIVd {
			inst.Args[0] = fpReg(uint32(rs1), d)
			inst.Args[1] = fpReg(x&0x1f, d)
			inst.Args[2] = fpReg(uint32(rd), d)
		}
		return inst, nil

	case op3 == 0x35: // FPop2, the compares
		opf := (x >> 5) & 0x1ff
		o, ok := fpop2[opf]
		if !ok {
			return inst, ErrUnrecognized
		}
		inst.Op = o
		d := isDoubleOp(o)
		inst.Args[0] = Reg(FCC0 + Reg((x>>25)&3))
		inst.Args[1] = fpReg(uint32(rs1), d)
		inst.Args[2] = fpReg(x&0x1f, d)
		return inst, nil

	case op3 == 0x36: // IMPDEP1: the VIS3 ops
		opf := (x >> 5) & 0x1ff
		o, ok := impdep1[opf]
		if !ok {
			return inst, ErrUnrecognized
		}
		inst.Op = o
		if o == SHA256 {
			return inst, nil
		}
		inst.Args[0] = rs1
		inst.Args[1] = rs2
		inst.Args[2] = rd
		return inst, nil

	case op3 == 0x3a: // Tcc
		inst.Op = Tcc
		inst.Args[0] = rs1
		inst.Args[1] = second
		return inst, nil
	}

	o := arithOp3[op3]
	if o == 0 {
		return inst, ErrUnrecognized
	}
	inst.Op = o
	switch o {
	case POPC:
		inst.Args[0] = second
		inst.Args[1] = rd
		inst.Args[2] = nil
		return inst, nil
	case RETURN:
		inst.Args[0] = rs1
		inst.Args[1] = second
	case SAVE, RESTORE, JMPL:
		inst.Args[0] = rs1
		inst.Args[1] = second
		inst.Args[2] = rd
	default:
		inst.Args[0] = rs1
		inst.Args[1] = second
		inst.Args[2] = rd
	}
	return inst, nil
}

func isStore(o Op) bool {
	switch o {
	case STB, STH, STW, STX, STD, STF, STDF, STQF, STFSR, STXFSR,
		STBA, STHA, STWA, STXA, STDA:
		return true
	}
	return false
}
