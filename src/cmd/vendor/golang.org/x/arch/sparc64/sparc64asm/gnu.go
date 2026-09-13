// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sparc64asm

import (
	"fmt"
	"strings"
)

// GNUSyntax returns the GNU assembler syntax for the instruction, as
// printed by objdump. GNU prints a number of synthetic instructions in
// place of their real encodings - mov, clr, cmp, retl and so on - and
// this reproduces them, since matching objdump is the whole point.
func GNUSyntax(inst Inst) string {
	if op, args, ok := synthetic(inst); ok {
		return join(op, args)
	}
	var args []string
	for _, a := range inst.Args {
		if a == nil {
			break
		}
		args = append(args, gnuArg(a))
	}
	return join(inst.Op.String(), args)
}

func join(op string, args []string) string {
	if len(args) == 0 {
		return op
	}
	return op + " " + strings.Join(args, ", ")
}

func gnuArg(a Arg) string {
	switch v := a.(type) {
	case Imm:
		return gnuImm(int64(v))
	case PCRel:
		return fmt.Sprintf("%d", int64(v))
	}
	return a.String()
}

// gnuImm formats an immediate the way objdump does: small magnitudes in
// decimal, anything larger in hex.
func gnuImm(v int64) string {
	// objdump prints negative immediates and small magnitudes in
	// decimal, larger positives in hex.
	if v < 0 || v < 10 {
		return fmt.Sprintf("%d", v)
	}
	return fmt.Sprintf("0x%x", v)
}

// synthetic recognises the encodings objdump prints under another name.
func synthetic(inst Inst) (string, []string, bool) {
	a := inst.Args
	switch inst.Op {
	case OR:
		// The zero register and a zero immediate are interchangeable
		// here, so treat both as "no value".
		zero := func(x Arg) bool {
			if r, ok := x.(Reg); ok {
				return r == G0
			}
			if i, ok := x.(Imm); ok {
				return i == 0
			}
			return false
		}
		immZero := func(x Arg) bool {
			i, ok := x.(Imm)
			return ok && i == 0
		}
		if zero(a[0]) && immZero(a[1]) {
			return "clr", []string{gnuArg(a[2])}, true
		}
		if zero(a[0]) {
			return "mov", []string{gnuArg(a[1]), gnuArg(a[2])}, true
		}
		if zero(a[1]) {
			return "mov", []string{gnuArg(a[0]), gnuArg(a[2])}, true
		}
	case SUBcc:
		// subcc rs1, X, %g0  ->  cmp rs1, X
		if rd, ok := a[2].(Reg); ok && rd == G0 {
			return "cmp", []string{gnuArg(a[0]), gnuArg(a[1])}, true
		}
	case SUB:
		// sub %g0, X, rd -> neg; sub rd, 1, rd -> dec
		if rs1, ok := a[0].(Reg); ok && rs1 == G0 {
			if src, ok1 := a[1].(Reg); ok1 {
				if dst, ok2 := a[2].(Reg); ok2 && src == dst {
					return "neg", []string{gnuArg(a[2])}, true
				}
			}
			return "neg", []string{gnuArg(a[1]), gnuArg(a[2])}, true
		}
		if imm, ok := a[1].(Imm); ok && imm == 1 {
			if rs1, ok1 := a[0].(Reg); ok1 {
				if rd, ok2 := a[2].(Reg); ok2 && rs1 == rd {
					return "dec", []string{gnuArg(a[2])}, true
				}
			}
		}
	case ADD:
		// add rd, 1, rd -> inc
		if imm, ok := a[1].(Imm); ok && imm == 1 {
			if rs1, ok1 := a[0].(Reg); ok1 {
				if rd, ok2 := a[2].(Reg); ok2 && rs1 == rd {
					return "inc", []string{gnuArg(a[2])}, true
				}
			}
		}
	case JMPL:
		// jmpl %o7+8, %g0 -> retl;  jmpl X, %o7 -> call
		rd, okd := a[2].(Reg)
		rs1, ok1 := a[0].(Reg)
		imm, oki := a[1].(Imm)
		if okd && ok1 && oki && rd == G0 && rs1 == O7 && imm == 8 {
			return "retl", nil, true
		}
		if okd && oki && rd == G0 && imm == 0 {
			return "jmp", []string{gnuArg(a[0])}, true
		}
		if okd && rd == O7 {
			if oki && imm == 0 {
				return "call", []string{gnuArg(a[0])}, true
			}
			return "call", []string{gnuArg(a[0]), gnuArg(a[1])}, true
		}
	case STX, STW, STH, STB:
		// storing %g0 is printed as clrx/clr/clrh/clrb
		if rd, ok := a[0].(Reg); ok && rd == G0 {
			name := map[Op]string{STX: "clrx", STW: "clr", STH: "clrh", STB: "clrb"}[inst.Op]
			return name, []string{gnuArg(a[1])}, true
		}
		if inst.Op == STW {
			return "st", []string{gnuArg(a[0]), gnuArg(a[1])}, true
		}
	case SETHI:
		return "sethi", []string{
			gnuHi(uint64(a[0].(Imm)) << 10), gnuArg(a[1])}, true
	case LDUW:
		return "ld", []string{gnuArg(a[0]), gnuArg(a[1])}, true
	case MEMBAR:
		return "membar", []string{membarMask(uint32(a[0].(Imm)))}, true
	case LDDF:
		return "ldd", []string{gnuArg(a[0]), gnuArg(a[1])}, true
	case STDF:
		return "std", []string{gnuArg(a[0]), gnuArg(a[1])}, true
	case LDF:
		return "ld", []string{gnuArg(a[0]), gnuArg(a[1])}, true
	case STF:
		return "st", []string{gnuArg(a[0]), gnuArg(a[1])}, true
	case FCMPs, FCMPd, FCMPEs, FCMPEd:
		// objdump omits %fcc0, naming it only when it is not the default.
		if cc, ok := a[0].(Reg); ok && cc == FCC0 {
			return inst.Op.String(), []string{gnuArg(a[1]), gnuArg(a[2])}, true
		}
	case CASA:
		return "cas", []string{gnuArg(a[0]), gnuArg(a[1]), gnuArg(a[2])}, true
	case CASXA:
		return "casx", []string{gnuArg(a[0]), gnuArg(a[1]), gnuArg(a[2])}, true
	case RESTORE, SAVE:
		// objdump prints a bare restore/save when every field is %g0.
		allZero := true
		for _, x := range a[:3] {
			r, ok := x.(Reg)
			if !ok || r != G0 {
				allZero = false
			}
		}
		if allZero {
			return inst.Op.String(), nil, true
		}
	case Tcc:
		// objdump prints only the trap number when rs1 is %g0.
		if rs1, ok := a[0].(Reg); ok && rs1 == G0 {
			return "ta", []string{gnuArg(a[1])}, true
		}
		return "ta", []string{gnuArg(a[0]), gnuArg(a[1])}, true
	}
	return "", nil, false
}

// membarMask renders MEMBAR's mmask/cmask the way objdump does.
func membarMask(m uint32) string {
	names := []struct {
		bit  uint32
		name string
	}{
		{0x01, "#LoadLoad"}, {0x02, "#StoreLoad"},
		{0x04, "#LoadStore"}, {0x08, "#StoreStore"},
		{0x10, "#Lookaside"}, {0x20, "#MemIssue"}, {0x40, "#Sync"},
	}
	// objdump lists them most-significant first.
	var out []string
	for i := len(names) - 1; i >= 0; i-- {
		if m&names[i].bit != 0 {
			out = append(out, names[i].name)
		}
	}
	if len(out) == 0 {
		return "0"
	}
	return strings.Join(out, "|")
}

// gnuHi renders sethi's operand; objdump prints a bare 0 for zero.
func gnuHi(v uint64) string {
	if v == 0 {
		return "%hi(0)"
	}
	return fmt.Sprintf("%%hi(%#x)", v)
}
