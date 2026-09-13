// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sparc64asm

import (
	"fmt"
	"strings"
)

// goRegName maps a SPARC integer register to the name the Go assembler
// uses for it. The port gives several hardware registers a role and a
// name of their own; everything else is R<n>.
var goRegName = map[Reg]string{
	G0: "ZR", G7: "TLS", O6: "RSP", O7: "LR",
	L5: "RFP", L6: "g", L7: "TMP2",
	I2: "TMP", I3: "RT1", I4: "RT2", I5: "CTXT", I7: "OLR",
}

// goOpName maps SPARC opcodes to the Go assembler mnemonic where they
// differ. Loads and stores in particular are all spelled MOV in Go.
var goOpName = map[Op]string{
	LDX: "MOVD", STX: "MOVD",
	LDUW: "MOVWU", LDSW: "MOVW", STW: "MOVW",
	LDUH: "MOVHU", LDSH: "MOVH", STH: "MOVH",
	LDUB: "MOVBU", LDSB: "MOVB", STB: "MOVB",
	LDF: "FMOVS", STF: "FMOVS", LDDF: "FMOVD", STDF: "FMOVD",
	JMPL: "JMPL", CALL: "CALL", NOP: "RNOP",
}

// GoSyntax returns the Go assembler syntax for the instruction.
// The pc is the program counter of the instruction, and symname, if
// non-nil, returns the name and base of the symbol containing an
// address, so branch and call targets can be shown by name.
func GoSyntax(inst Inst, pc uint64, symname func(uint64) (string, uint64)) string {
	if symname == nil {
		symname = func(uint64) (string, uint64) { return "", 0 }
	}
	// Go spells the return and unconditional-branch idioms RET and JMP,
	// as the assembler does, rather than showing their encodings.
	if op, args, ok := goSynthetic(inst, pc, symname); ok {
		if len(args) == 0 {
			return op
		}
		return op + " " + strings.Join(args, ", ")
	}

	op := goOpName[inst.Op]
	if op == "" {
		op = strings.ToUpper(inst.Op.String())
	}

	var args []string
	for _, a := range inst.Args {
		if a == nil {
			break
		}
		args = append(args, goArg(a, pc, symname))
	}

	// Go writes the destination last and otherwise reverses SPARC's
	// operand order: "add rs1, rs2, rd" prints as "ADD rs2, rs1, rd".
	if len(args) == 3 && !isMemOp(inst.Op) {
		args[0], args[1] = args[1], args[0]
	}
	if len(args) == 0 {
		return op
	}
	return op + " " + strings.Join(args, ", ")
}

func isMemOp(o Op) bool {
	switch o {
	case LDX, LDUW, LDSW, LDUH, LDSH, LDUB, LDSB, LDD, LDF, LDDF,
		STX, STW, STH, STB, STD, STF, STDF:
		return true
	}
	return false
}

func goArg(a Arg, pc uint64, symname func(uint64) (string, uint64)) string {
	switch v := a.(type) {
	case Reg:
		return goReg(v)
	case Imm:
		return fmt.Sprintf("$%d", int64(v))
	case PCRel:
		target := uint64(int64(pc) + int64(v))
		if name, base := symname(target); name != "" {
			if target == base {
				return fmt.Sprintf("%s(SB)", name)
			}
			return fmt.Sprintf("%s+%d(SB)", name, target-base)
		}
		return fmt.Sprintf("%#x", target)
	case Mem:
		return goMem(v)
	}
	return a.String()
}

func goReg(r Reg) string {
	if n, ok := goRegName[r]; ok {
		return n
	}
	switch {
	case r <= I7:
		return fmt.Sprintf("R%d", int(r))
	case r >= F0 && r <= F31:
		return fmt.Sprintf("F%d", int(r-F0))
	case r >= D0 && r <= D62:
		return fmt.Sprintf("D%d", int(r-D0))
	}
	switch r {
	case ICC:
		return "ICC"
	case XCC:
		return "XCC"
	case CCR:
		return "CCR"
	}
	return r.String()
}

func goMem(m Mem) string {
	if m.Ind != G0 {
		return fmt.Sprintf("(%s)(%s)", goReg(m.Base), goReg(m.Ind))
	}
	if m.Off != 0 {
		return fmt.Sprintf("%d(%s)", int64(m.Off), goReg(m.Base))
	}
	return fmt.Sprintf("(%s)", goReg(m.Base))
}

// goSynthetic recognises the encodings the Go assembler writes under a
// different name.
func goSynthetic(inst Inst, pc uint64, symname func(uint64) (string, uint64)) (string, []string, bool) {
	a := inst.Args
	switch inst.Op {
	case JMPL:
		// jmpl LR+8, ZR is the return sequence.
		rs1, ok1 := a[0].(Reg)
		imm, oki := a[1].(Imm)
		rd, okd := a[2].(Reg)
		if ok1 && oki && okd && rs1 == O7 && imm == 8 && rd == G0 {
			return "RET", nil, true
		}
		// jmpl rs1+0, ZR is an indirect jump.
		if ok1 && oki && okd && imm == 0 && rd == G0 {
			return "JMP", []string{"(" + goReg(rs1) + ")"}, true
		}
	case BA:
		// An always-taken branch is a plain jump.
		for _, x := range a {
			if p, ok := x.(PCRel); ok {
				return "JMP", []string{goArg(p, pc, symname)}, true
			}
		}
	}
	return "", nil, false
}
