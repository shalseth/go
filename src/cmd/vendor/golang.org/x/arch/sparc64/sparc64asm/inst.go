// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sparc64asm implements decoding of SPARC V9 machine code.
package sparc64asm

import (
	"fmt"
	"strings"
)

// An Op is a SPARC V9 opcode.
type Op uint16

// An Inst is a single instruction.
type Inst struct {
	Op   Op     // opcode mnemonic
	Enc  uint32 // raw encoding
	Args Args   // instruction arguments, in SPARC manual order
}

func (i Inst) String() string {
	var buf strings.Builder
	buf.WriteString(i.Op.String())
	for j, arg := range i.Args {
		if arg == nil {
			break
		}
		if j == 0 {
			buf.WriteString(" ")
		} else {
			buf.WriteString(", ")
		}
		buf.WriteString(arg.String())
	}
	return buf.String()
}

// An Args holds the arguments of an instruction. Unused slots are nil.
// SPARC never needs more than four.
type Args [4]Arg

// An Arg is a single instruction argument.
type Arg interface {
	String() string
	isArg()
}

// A Reg is an integer or floating-point register.
// Integer registers are %g0-%g7, %o0-%o7, %l0-%l7, %i0-%i7.
type Reg uint16

// Integer registers occupy 0..31 in encoding order:
// %g0-%g7, %o0-%o7, %l0-%l7, %i0-%i7.
const (
	G0 Reg = iota
	G1
	G2
	G3
	G4
	G5
	G6
	G7
	O0
	O1
	O2
	O3
	O4
	O5
	O6 // %sp
	O7
	L0
	L1
	L2
	L3
	L4
	L5
	L6
	L7
	I0
	I1
	I2
	I3
	I4
	I5
	I6 // %fp
	I7
)

// Single-precision float registers %f0-%f31 occupy 32..63,
// double-precision %f0,%f2,...,%f62 occupy 64..126 (even only).
const (
	F0 Reg = 32
	F31    = F0 + 31
	D0  Reg = 64
	D62    = D0 + 62
)

// Special registers.
const (
	ICC Reg = 200 + iota
	XCC
	FCC0
	FCC1
	FCC2
	FCC3
	Y
	ASI
	FPRS
	CCR
)

// RegF returns the single-precision register %f<n>.
func RegF(n uint8) Reg { return F0 + Reg(n) }

// RegD returns the double-precision register %f<n>, n even.
func RegD(n uint8) Reg { return D0 + Reg(n) }

func (Reg) isArg() {}

func (r Reg) String() string {
	switch {
	case r <= G7:
		return fmt.Sprintf("%%g%d", int(r-G0))
	case r <= O7:
		if r == O6 {
			return "%sp"
		}
		return fmt.Sprintf("%%o%d", int(r-O0))
	case r <= L7:
		return fmt.Sprintf("%%l%d", int(r-L0))
	case r <= I7:
		if r == I6 {
			return "%fp"
		}
		return fmt.Sprintf("%%i%d", int(r-I0))
	case r >= F0 && r <= F31:
		return fmt.Sprintf("%%f%d", int(r-F0))
	case r >= D0 && r <= D62:
		return fmt.Sprintf("%%f%d", int(r-D0))
	}
	switch r {
	case ICC:
		return "%icc"
	case XCC:
		return "%xcc"
	case FCC0:
		return "%fcc0"
	case FCC1:
		return "%fcc1"
	case FCC2:
		return "%fcc2"
	case FCC3:
		return "%fcc3"
	case Y:
		return "%y"
	case ASI:
		return "%asi"
	case FPRS:
		return "%fprs"
	case CCR:
		return "%ccr"
	}
	return fmt.Sprintf("Reg(%d)", uint16(r))
}

// An Imm is an immediate (signed, already sign-extended where the
// encoding calls for it).
type Imm int64

func (Imm) isArg() {}

func (i Imm) String() string { return fmt.Sprintf("%d", int64(i)) }

// A Mem is a memory operand: [rs1 + rs2] or [rs1 + simm13].
type Mem struct {
	Base Reg
	Ind  Reg // zero value G0 means "no index register"
	Off  Imm
	ASI  uint8 // alternate address space, for the *a forms
	HasA bool
}

func (Mem) isArg() {}

func (m Mem) String() string {
	var inner string
	if m.Ind != G0 {
		inner = fmt.Sprintf("%s + %s", m.Base, m.Ind)
	} else if m.Off != 0 {
		if m.Off < 0 {
			inner = fmt.Sprintf("%s + %d", m.Base, int64(m.Off))
		} else if m.Off < 10 {
			inner = fmt.Sprintf("%s + %d", m.Base, int64(m.Off))
		} else {
			inner = fmt.Sprintf("%s + 0x%x", m.Base, int64(m.Off))
		}
	} else {
		inner = m.Base.String()
	}
	s := fmt.Sprintf("[ %s ]", inner)
	if m.HasA {
		s += fmt.Sprintf(" %d", m.ASI)
	}
	return s
}

// A PCRel is a PC-relative branch or call target, in bytes.
type PCRel int64

func (PCRel) isArg() {}

func (p PCRel) String() string { return fmt.Sprintf("%+d", int64(p)) }
