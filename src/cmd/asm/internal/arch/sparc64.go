// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file encapsulates some of the odd characteristics of the
// SPARC V9 instruction set, to minimize its interaction with the
// core of the assembler.

package arch

import (
	"cmd/internal/obj"
	"cmd/internal/obj/sparc64"
)

// jumpSPARC64 reports whether word is a SPARC V9 branch or call.
// SPARC has a lot of these: the integer branches exist in a plain
// form and in ICC-implied (W) and XCC-implied (D) aliases, and the
// register-compare branches (BR*) and float branches (FB*) are
// separate families again.
func jumpSPARC64(word string) bool {
	switch word {
	// RET and the register-indirect JMPL are not listed: the parser
	// routes IsJump opcodes to a branch-target form, and those two are
	// handled by the generic instruction path instead.
	case "CALL", "JMP",
		// BPcc, condition given explicitly
		"BN", "BNE", "BE", "BG", "BLE", "BGE", "BL", "BGU", "BLEU",
		"BCC", "BCS", "BPOS", "BNEG", "BVC", "BVS",
		// BPcc with ICC implied
		"BNW", "BNEW", "BEW", "BGW", "BLEW", "BGEW", "BLW", "BGUW",
		"BLEUW", "BCCW", "BCSW", "BPOSW", "BNEGW", "BVCW", "BVSW",
		// BPcc with XCC implied
		"BND", "BNED", "BED", "BGD", "BLED", "BGED", "BLD", "BGUD",
		"BLEUD", "BCCD", "BCSD", "BPOSD", "BNEGD", "BVCD", "BVSD",
		// BPr, branch on register contents
		"BRZ", "BRLEZ", "BRLZ", "BRNZ", "BRGZ", "BRGEZ",
		// FBfcc
		"FBA", "FBN", "FBU", "FBG", "FBUG", "FBL", "FBUL", "FBLG",
		"FBNE", "FBE", "FBUE", "FBGE", "FBUGE", "FBLE", "FBULE", "FBO":
		return true
	}
	return false
}

// sparc64RegisterNumber converts R(10) into sparc64.REG_R10 and so on.
func sparc64RegisterNumber(name string, n int16) (int16, bool) {
	switch name {
	case "R":
		if 0 <= n && n <= 31 {
			return sparc64.REG_R0 + n, true
		}
	case "G":
		if 0 <= n && n <= 7 {
			return sparc64.REG_R0 + n, true
		}
	case "O":
		if 0 <= n && n <= 7 {
			return sparc64.REG_R8 + n, true
		}
	case "L":
		if 0 <= n && n <= 7 {
			return sparc64.REG_R16 + n, true
		}
	case "I":
		if 0 <= n && n <= 7 {
			return sparc64.REG_R24 + n, true
		}
	case "F":
		if 0 <= n && n <= 31 {
			return sparc64.REG_F0 + n, true
		}
	case "D":
		// Double-precision registers are D0, D2, ... D62. The upper
		// half (D32..D62) is interleaved with the lower half in the
		// register space; see a.out.go.
		if 0 <= n && n <= 62 && n%2 == 0 {
			if n < 32 {
				return sparc64.REG_D0 + n, true
			}
			return sparc64.REG_D32 + (n - 32), true
		}
	case "Y":
		if 0 <= n && n <= 15 {
			return sparc64.REG_Y0 + n, true
		}
	}
	return 0, false
}

// IsSPARC64CMP reports whether op is a comparison that writes only the
// condition codes. These take their second operand in Prog.Reg rather
// than Prog.To, because they have no destination register.
func IsSPARC64CMP(op obj.As) bool {
	switch op {
	case sparc64.ACMP, sparc64.AFCMPS, sparc64.AFCMPD:
		return true
	}
	return false
}
