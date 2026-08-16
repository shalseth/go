// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sparc64

import "cmd/internal/obj"

//go:generate go run ../stringer.go -i $GOFILE -o anames.go -p sparc64

// SPARC V9 has 32 architecturally visible integer registers at any
// moment, drawn from a register window: 8 global, 8 out, 8 local and
// 8 in. Go numbers them flat as R0..R31, which maps onto the usual
// SPARC assembler names as:
//
//	R0..R7    %g0..%g7   globals (%g0 is hardwired zero)
//	R8..R15   %o0..%o7   out    (%o6 is the stack pointer, %o7 the link register)
//	R16..R23  %l0..%l7   local
//	R24..R31  %i0..%i7   in     (%i6 is the frame pointer, %i7 the caller's return address)
//
// Go code does not execute SAVE/RESTORE, so the window never rotates
// within Go frames; see obj.go for how the frame is laid out and
// docs/sparc64-port.md for why.
const (
	// integer
	REG_R0 = obj.RBaseSPARC64 + iota
	REG_R1
	REG_R2
	REG_R3
	REG_R4
	REG_R5
	REG_R6
	REG_R7
	REG_R8
	REG_R9
	REG_R10
	REG_R11
	REG_R12
	REG_R13
	REG_R14
	REG_R15
	REG_R16
	REG_R17
	REG_R18
	REG_R19
	REG_R20
	REG_R21
	REG_R22
	REG_R23
	REG_R24
	REG_R25
	REG_R26
	REG_R27
	REG_R28
	REG_R29
	REG_R30
	REG_R31

	// single-precision floating point
	REG_F0
	REG_F1
	REG_F2
	REG_F3
	REG_F4
	REG_F5
	REG_F6
	REG_F7
	REG_F8
	REG_F9
	REG_F10
	REG_F11
	REG_F12
	REG_F13
	REG_F14
	REG_F15
	REG_F16
	REG_F17
	REG_F18
	REG_F19
	REG_F20
	REG_F21
	REG_F22
	REG_F23
	REG_F24
	REG_F25
	REG_F26
	REG_F27
	REG_F28
	REG_F29
	REG_F30
	REG_F31

	// double-precision floating point; the first half is aliased to
	// single-precision registers, that is: Dn is aliased to Fn, Fn+1,
	// where n ≤ 30.
	REG_D0
	REG_D32
	REG_D2
	REG_D34
	REG_D4
	REG_D36
	REG_D6
	REG_D38
	REG_D8
	REG_D40
	REG_D10
	REG_D42
	REG_D12
	REG_D44
	REG_D14
	REG_D46
	REG_D16
	REG_D48
	REG_D18
	REG_D50
	REG_D20
	REG_D52
	REG_D22
	REG_D54
	REG_D24
	REG_D56
	REG_D26
	REG_D58
	REG_D28
	REG_D60
	REG_D30
	REG_D62

	// common single/double-precision virtualized registers.
	// Yn is aliased to F2n, F2n+1, D2n.
	REG_Y0
	REG_Y1
	REG_Y2
	REG_Y3
	REG_Y4
	REG_Y5
	REG_Y6
	REG_Y7
	REG_Y8
	REG_Y9
	REG_Y10
	REG_Y11
	REG_Y12
	REG_Y13
	REG_Y14
	REG_Y15
)

const (
	// floating-point condition-code registers
	REG_FCC0 = REG_R0 + 256 + iota
	REG_FCC1
	REG_FCC2
	REG_FCC3
)

const (
	// integer condition-code flags
	REG_ICC = REG_R0 + 384     // 32-bit condition codes
	REG_XCC = REG_R0 + 384 + 2 // 64-bit condition codes
)

const (
	REG_SPECIAL = REG_R0 + 512

	REG_CCR  = REG_SPECIAL + 2
	REG_TICK = REG_SPECIAL + 4
	REG_RPC  = REG_SPECIAL + 5

	// Biased views of the stack and frame pointers. On SPARC V9 the
	// 64-bit ABI stores %sp and %fp biased by StackBias; these encode
	// "the biased form of" the register in an Addr.
	REG_BSP = REG_RSP + 256
	REG_BFP = REG_RFP + 256

	REG_LAST = REG_R0 + 1024
)

// Register assignments.
const (
	REG_ZR   = REG_R0  // %g0, hardwired zero
	REG_TLS  = REG_R7  // %g7, thread pointer (set by the kernel)
	REG_RSP  = REG_R14 // %o6, hardware stack pointer (biased)
	REG_LR   = REG_R15 // %o7, link register written by CALL
	REG_G    = REG_R22 // %l6, current goroutine
	REG_TMP2 = REG_R23 // %l7
	REG_TMP  = REG_R26 // %i2
	REG_RT1  = REG_R27 // %i3
	REG_RT2  = REG_R28 // %i4
	REG_CTXT = REG_R29 // %i5, closure context pointer
	REG_RFP  = REG_R30 // %i6, frame pointer (biased)
	REG_OLR  = REG_R31 // %i7
	REG_FTMP = REG_F0
	REG_DTMP = REG_D0
	REG_YTMP = REG_Y0
)

// Conventional names shared with the other backends. cmd/compile's SSA
// generator emits references to REGSP, REGG and REGZERO by those exact
// names, and the runtime and linker use REGCTXT, REGTMP and REGLINK.
const (
	REGZERO = REG_ZR
	REGSP   = REG_RSP
	REGG    = REG_G
	REGCTXT = REG_CTXT
	REGTMP  = REG_TMP
	REGLINK = REG_LR
)

// SPARCDWARFRegisters maps Go's register numbers onto the DWARF
// numbering, which follows the architectural register file: 0-7 are
// %g0-%g7, 8-15 %o0-%o7, 16-23 %l0-%l7 and 24-31 %i0-%i7. The 32
// single-precision float registers follow at 32.
var SPARCDWARFRegisters = map[int16]int16{}

func init() {
	// f maps [from, to] to DWARF numbers starting at base.
	f := func(from, to, base int16) {
		for r := int16(from); r <= to; r++ {
			SPARCDWARFRegisters[r] = (r - from) + base
		}
	}
	f(REG_R0, REG_R31, 0)
	f(REG_F0, REG_F31, 32)
}

const (
	REG_MIN = REG_R0
	REG_MAX = REG_R25
)

const (
	// StackAlign is the alignment SPARC V9 requires of the stack
	// pointer. The 2016 Solaris port used 8 here despite the ABI
	// requiring 16; 16 is used instead, which is strictly more
	// conservative.
	StackAlign = 16

	// StackBias is the constant SPARC V9 subtracts from %sp and %fp
	// in the 64-bit ABI: the real stack top is %sp+StackBias. The
	// kernel assumes this when delivering signals and C code assumes
	// it at the cgo boundary.
	StackBias = 0x7ff

	// WindowSaveAreaSize is the 16-extended-word region every frame
	// must reserve at %sp+StackBias for the register window. This is
	// not optional: a window-overflow trap, a signal, or an explicit
	// FLUSHW spills the current window there, whether or not Go code
	// ever executed SAVE.
	WindowSaveAreaSize = 16 * 8

	// ArgumentsSaveAreaSize is the 6-extended-word outgoing-argument
	// region the SPARC V9 C ABI reserves. Go does not use it, but it
	// is kept so Go frames remain walkable by C tooling and so the
	// cgo boundary needs no special case.
	ArgumentsSaveAreaSize = 6 * 8

	MinStackFrameSize = WindowSaveAreaSize + ArgumentsSaveAreaSize // 176
)

const (
	BIG = 1<<12 - 1 // magnitude of smallest negative immediate
)

// Prog.Mark
const (
	FOLL = 1 << iota
	LABEL
	LEAF
)

// Operand classes. The assembler classifies each Addr before
// selecting an encoding.
const (
	ClassUnknown = iota

	ClassReg    // R1..R31
	ClassFReg   // F0..F31
	ClassDReg   // D0..D62
	ClassCond   // ICC, XCC
	ClassFCond  // FCC0..FCC3
	ClassSpcReg // TICK, CCR, etc

	ClassZero     // $0 or ZR
	ClassConst5   // unsigned 5-bit constant
	ClassConst6   // unsigned 6-bit constant
	ClassConst10  // signed 10-bit constant
	ClassConst11  // signed 11-bit constant
	ClassConst13  // signed 13-bit constant
	ClassConst31_ // signed 32-bit constant, negative
	ClassConst31  // signed 32-bit constant, positive or zero
	ClassConst32  // 32-bit constant
	ClassConst    // 64-bit constant
	ClassFConst   // floating-point constant

	ClassRegReg     // $(Rn+Rm) or $(Rn)(Rm*1)
	ClassRegConst13 // $n(R), n is 13-bit signed
	ClassRegConst   // $n(R), n large

	ClassIndirRegReg // (Rn+Rm) or (Rn)(Rm*1)
	ClassIndir0      // (R)
	ClassIndir13     // n(R), n is 13-bit signed
	ClassIndir       // n(R), n large

	ClassBranch // n(PC) branch target, n is 21-bit signed, mod 4

	ClassAddr    // $sym(SB)
	ClassMem     // sym(SB)
	ClassTLSAddr // $tlssym(SB)
	ClassTLSMem  // tlssym(SB)

	ClassTextSize
	ClassNone

	ClassBias = 64 // BFP or BSP present in Addr, bitwise OR with classes above
)

var cnames = []string{
	ClassUnknown:     "ClassUnknown",
	ClassReg:         "ClassReg",
	ClassFReg:        "ClassFReg",
	ClassDReg:        "ClassDReg",
	ClassCond:        "ClassCond",
	ClassFCond:       "ClassFCond",
	ClassSpcReg:      "ClassSpcReg",
	ClassZero:        "ClassZero",
	ClassConst5:      "ClassConst5",
	ClassConst6:      "ClassConst6",
	ClassConst10:     "ClassConst10",
	ClassConst11:     "ClassConst11",
	ClassConst13:     "ClassConst13",
	ClassConst31_:    "ClassConst31-",
	ClassConst31:     "ClassConst31+",
	ClassConst32:     "ClassConst32",
	ClassConst:       "ClassConst",
	ClassFConst:      "ClassFConst",
	ClassRegReg:      "ClassRegReg",
	ClassRegConst13:  "ClassRegConst13",
	ClassRegConst:    "ClassRegConst",
	ClassIndirRegReg: "ClassIndirRegReg",
	ClassIndir0:      "ClassIndir0",
	ClassIndir13:     "ClassIndir13",
	ClassIndir:       "ClassIndir",
	ClassBranch:      "ClassBranch",
	ClassAddr:        "ClassAddr",
	ClassMem:         "ClassMem",
	ClassTLSAddr:     "ClassTLSAddr",
	ClassTLSMem:      "ClassTLSMem",
	ClassTextSize:    "ClassTextSize",
	ClassNone:        "ClassNone",
	ClassBias:        "ClassBias",
}

const (
	AADD = obj.ABaseSPARC64 + obj.A_ARCHSPECIFIC + iota
	AADDCC
	AADDC
	AADDCCC
	AAND
	AANDCC
	AANDN
	AANDNCC

	// These are the two-operand SPARCv9 32-, and 64-bit, branch
	// on integer condition codes with prediction (BPcc), not the
	// single-operand SPARCv8 32-bit branch on integer condition
	// codes (Bicc).
	ABN
	ABNE
	ABE
	ABG
	ABLE
	ABGE
	ABL
	ABGU
	ABLEU
	ABCC
	ABCS
	ABPOS
	ABNEG
	ABVC
	ABVS

	ABRZ
	ABRLEZ
	ABRLZ
	ABRNZ
	ABRGZ
	ABRGEZ
	ACASW
	ACASD
	AFABSS
	AFABSD
	AFADDS
	AFADDD
	AFBA
	AFBN
	AFBU
	AFBG
	AFBUG
	AFBL
	AFBUL
	AFBLG
	AFBNE
	AFBE
	AFBUE
	AFBGE
	AFBUGE
	AFBLE
	AFBULE
	AFBO
	AFCMPS
	AFCMPD
	AFDIVS
	AFDIVD
	AFITOS
	AFITOD
	AFLUSH
	AFMOVS // the SPARC64 instruction, and alias for loads and stores
	AFMOVD // the SPARC64 instruction, and alias for loads and stores
	AFMULS
	AFMULD
	AFSMULD
	AFNEGS
	AFNEGD
	AFSQRTS
	AFSQRTD
	AFSTOX
	AFDTOX
	AFSTOI
	AFDTOI
	AFSTOD
	AFDTOS
	AFSUBS
	AFSUBD
	AFXTOS
	AFXTOD
	AJMPL
	ALDSB
	ALDSH
	ALDSW
	ALDUB
	ALDD
	ALDDF
	ALDSF
	ALDUH
	ALDUW
	AMEMBAR
	AMOVA
	AMOVCC
	AMOVCS
	AMOVE
	AMOVG
	AMOVGE
	AMOVGU
	AMOVL
	AMOVLE
	AMOVLEU
	AMOVN
	AMOVNE
	AMOVNEG
	AMOVPOS
	AMOVRGEZ
	AMOVRGZ
	AMOVRLEZ
	AMOVRLZ
	AMOVRNZ
	AMOVRZ
	AMOVVC
	AMOVVS

	// Move Integer Register on Floating-point Condition (FMOVcc). These
	// are separate opcodes rather than the integer MOVcc with cc2
	// cleared, because the cond field is a different encoding: integer
	// cond 0001 is E, but float cond 0001 is NE.
	AMOVFE
	AMOVFNE
	AMOVFL
	AMOVFLE
	AMOVFG
	AMOVFGE
	AMULD
	AOR
	AORCC
	AORN
	AORNCC
	ARD
	ARESTORE // not used under normal circumstances
	ASAVE    // not used under normal circumstances
	ASDIVD
	AUMULXHI // VIS3: high 64 bits of an unsigned 64x64 multiply
	ASETHI
	AUDIVD
	ASLLW
	ASRLW
	ASRAW
	ASLLD
	ASRLD
	ASRAD
	ASTB
	ASTH
	ASTW
	ASTD
	ASTSF
	ASTDF
	ASUB
	ASUBCC
	ASUBC
	ASUBCCC
	ATA
	AXOR
	AXORCC
	AXNOR
	AXNORCC

	// Pseudo-instructions, aliases to SPARC64 instructions and
	// synthetic instructions.
	ACMP // SUBCC R1, R2, ZR
	ANEG
	AMOVUB
	AMOVB
	AMOVUH
	AMOVH
	AMOVUW
	AMOVW
	AMOVD // also the SPARC64 synthetic instruction
	ARNOP // SETHI $0, ZR

	// These are aliases to two-operand SPARCv9 32-, and 64-bit,
	// branch on integer condition codes with prediction (BPcc),
	// with ICC implied.
	ABNW
	ABNEW
	ABEW
	ABGW
	ABLEW
	ABGEW
	ABLW
	ABGUW
	ABLEUW
	ABCCW
	ABCSW
	ABPOSW
	ABNEGW
	ABVCW
	ABVSW

	// These are aliases to two-operand SPARCv9 32-, and 64-bit,
	// branch on integer condition codes with prediction (BPcc),
	// with XCC implied.
	ABND
	ABNED
	ABED
	ABGD
	ABLED
	ABGED
	ABLD
	ABGUD
	ABLEUD
	ABCCD
	ABCSD
	ABPOSD
	ABNEGD
	ABVCD
	ABVSD

	AWORD
	ADWORD

	ALAST
)
