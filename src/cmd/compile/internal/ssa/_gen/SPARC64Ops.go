// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "strings"

// Notes:
//  - SPARC V9 is big-endian and traps on unaligned access.
//  - Integer types live in the low portion of registers. Upper portions
//    are junk.
//  - Booleans use the low-order byte of a register. 0=false, 1=true.
//  - Go code never executes SAVE/RESTORE, so the register window does
//    not rotate within Go frames. The 32 architectural registers are
//    treated as a flat file; see cmd/internal/obj/sparc64/a.out.go for
//    the R0..R31 to %g/%o/%l/%i mapping.
//  - Floating-point values are held in the "Y" virtual registers, which
//    the assembler rewrites to F (single) or D (double) registers based
//    on the instruction. This sidesteps SPARC's aliased float file,
//    where %d0 overlaps %f0 and %f1.
//
// Suffixes encode the bit width of various instructions.
// D (double word) = 64 bit
// W (word)        = 32 bit
// H (half word)   = 16 bit
// B (byte)        = 8 bit
// U               = unsigned
// S (single)      = 32 bit float
// D on float ops  = 64 bit float

// Registers not used in regalloc are omitted so the mask stays small.
// Reserved, and therefore absent: R0 (%g0, hardwired zero, named ZR),
// R6/R7 (%g6 kernel, %g7 thread pointer), R14 (%o6, the stack pointer,
// named SP), R22 (%l6, holds g), R23 (%l7, assembler temporary),
// R26 (%i2, temporary), R27/R28 (%i3/%i4, runtime temporaries),
// R29 (%i5, closure context), R30 (%i6, frame pointer), R31 (%i7),
// and Y0 (the float temporary).
var regNamesSPARC64 = []string{
	"ZR", // R0, %g0, constant 0

	// %g1..%g5 are volatile scratch on Linux/sparc64.
	"R1",
	"R2",
	"R3",
	"R4",
	"R5",

	// %o0..%o5, outgoing arguments.
	"R8",
	"R9",
	"R10",
	"R11",
	"R12",
	"R13",

	"R15", // %o7, link register

	// %l0..%l5.
	"R16",
	"R17",
	"R18",
	"R19",
	"R20",
	"R21",

	// %i0..%i1.
	"R24",
	"R25",

	// %i5, the closure context pointer. Named but not allocatable.
	"R29",

	"SP", // R14, %o6
	"g",  // R22, %l6

	// Floating point, as virtual Y registers. Y0 is the temporary.
	"Y1",
	"Y2",
	"Y3",
	"Y4",
	"Y5",
	"Y6",
	"Y7",
	"Y8",
	"Y9",
	"Y10",
	"Y11",
	"Y12",
	"Y13",
	"Y14",
	"Y15",

	// If you add registers, update asyncPreempt in runtime.

	// pseudo-registers
	"SB",
}

func init() {
	// Make map from reg names to reg integers.
	if len(regNamesSPARC64) > 64 {
		panic("too many registers")
	}
	num := map[string]int{}
	for i, name := range regNamesSPARC64 {
		num[name] = i
	}
	buildReg := func(s string) regMask {
		m := regMask{}
		for _, r := range strings.Split(s, " ") {
			if n, ok := num[r]; ok {
				m = m.addReg(uint(n))
				continue
			}
			panic("register " + r + " not found")
		}
		return m
	}

	// Common individual register masks.
	var (
		gp   = buildReg("R1 R2 R3 R4 R5 R8 R9 R10 R11 R12 R13 R15 R16 R17 R18 R19 R20 R21 R24 R25")
		gpg      = gp.union(buildReg("g"))
		gpsp     = gp.union(buildReg("SP"))
		gpspg    = gpg.union(buildReg("SP"))
		gpspsbg  = gpspg.union(buildReg("SB"))
		fp       = buildReg("Y1 Y2 Y3 Y4 Y5 Y6 Y7 Y8 Y9 Y10 Y11 Y12 Y13 Y14 Y15")
		rz       = buildReg("ZR")
		callerSave = gp.union(fp).union(buildReg("g")) // runtime.setg may clobber g
	)
	// Common regInfo.
	var (
		gp01     = regInfo{inputs: nil, outputs: []regMask{gp}}
		gp11     = regInfo{inputs: []regMask{gpg}, outputs: []regMask{gp}}
		gp11sp   = regInfo{inputs: []regMask{gpspg}, outputs: []regMask{gp}}
		gp21     = regInfo{inputs: []regMask{gpg, gpg.union(rz)}, outputs: []regMask{gp}}
		gp2flags = regInfo{inputs: []regMask{gpg, gpg.union(rz)}}
		gp1flags = regInfo{inputs: []regMask{gpg}}
		gpload   = regInfo{inputs: []regMask{gpspsbg}, outputs: []regMask{gp}}
		gpstore  = regInfo{inputs: []regMask{gpspsbg, gpg.union(rz)}}
		fp01     = regInfo{inputs: nil, outputs: []regMask{fp}}
		fp11     = regInfo{inputs: []regMask{fp}, outputs: []regMask{fp}}
		fp21     = regInfo{inputs: []regMask{fp, fp}, outputs: []regMask{fp}}
		fp2flags = regInfo{inputs: []regMask{fp, fp}}
		fpgp     = regInfo{inputs: []regMask{fp}, outputs: []regMask{gp}}
		gpfp     = regInfo{inputs: []regMask{gp}, outputs: []regMask{fp}}
		fpload   = regInfo{inputs: []regMask{gpspsbg}, outputs: []regMask{fp}}
		fpstore  = regInfo{inputs: []regMask{gpspsbg, fp}}
		// Flags are not modelled as an allocatable register, so the
		// flag-reading ops declare no inputs.
		readflags = regInfo{inputs: nil, outputs: []regMask{gp}}
	)
	ops := []opData{
		// Integer arithmetic. SPARC's three-operand form is
		// "op rs1, rs2, rd"; the *const variants take the second
		// operand from the 13-bit signed immediate field, and the
		// assembler expands anything larger via the temporary.
		{name: "ADD", argLength: 2, reg: gp21, asm: "ADD", commutative: true},   // arg0 + arg1
		{name: "ADDconst", argLength: 1, reg: gp11sp, asm: "ADD", aux: "Int64"}, // arg0 + auxInt
		{name: "SUB", argLength: 2, reg: gp21, asm: "SUB"},                      // arg0 - arg1
		{name: "SUBconst", argLength: 1, reg: gp11, asm: "SUB", aux: "Int64"},    // arg0 - auxInt
		{name: "MULD", argLength: 2, reg: gp21, asm: "MULD", commutative: true},  // arg0 * arg1
		{name: "SDIVD", argLength: 2, reg: gp21, asm: "SDIVD"},                   // arg0 / arg1, signed
		{name: "UDIVD", argLength: 2, reg: gp21, asm: "UDIVD"},                   // arg0 / arg1, unsigned

		{name: "AND", argLength: 2, reg: gp21, asm: "AND", commutative: true},   // arg0 & arg1
		{name: "ANDconst", argLength: 1, reg: gp11, asm: "AND", aux: "Int64"},   // arg0 & auxInt
		{name: "OR", argLength: 2, reg: gp21, asm: "OR", commutative: true},     // arg0 | arg1
		{name: "ORconst", argLength: 1, reg: gp11, asm: "OR", aux: "Int64"},     // arg0 | auxInt
		{name: "XOR", argLength: 2, reg: gp21, asm: "XOR", commutative: true},   // arg0 ^ arg1
		{name: "XORconst", argLength: 1, reg: gp11, asm: "XOR", aux: "Int64"},   // arg0 ^ auxInt
		{name: "ANDN", argLength: 2, reg: gp21, asm: "ANDN"},                     // arg0 &^ arg1
		{name: "ORN", argLength: 2, reg: gp21, asm: "ORN"},                       // arg0 |^ arg1
		{name: "XNOR", argLength: 2, reg: gp21, asm: "XNOR", commutative: true}, // ^(arg0 ^ arg1)

		// Shifts. SPARC distinguishes 32-bit (W) and 64-bit (D) forms;
		// the W forms operate on the low word.
		{name: "SLLD", argLength: 2, reg: gp21, asm: "SLLD"},                   // arg0 << arg1, 64 bit
		{name: "SLLDconst", argLength: 1, reg: gp11, asm: "SLLD", aux: "Int64"}, // arg0 << auxInt
		{name: "SRLD", argLength: 2, reg: gp21, asm: "SRLD"},                    // arg0 >> arg1, logical
		{name: "SRLDconst", argLength: 1, reg: gp11, asm: "SRLD", aux: "Int64"},
		{name: "SRAD", argLength: 2, reg: gp21, asm: "SRAD"}, // arg0 >> arg1, arithmetic
		{name: "SRADconst", argLength: 1, reg: gp11, asm: "SRAD", aux: "Int64"},
		{name: "SLLW", argLength: 2, reg: gp21, asm: "SLLW"},
		{name: "SRLW", argLength: 2, reg: gp21, asm: "SRLW"},
		{name: "SRAW", argLength: 2, reg: gp21, asm: "SRAW"},

		// Comparisons write the condition codes. CMP is SUBCC to ZR.
		{name: "CMP", argLength: 2, reg: gp2flags, asm: "CMP", typ: "Flags"},                    // arg0 compare arg1
		{name: "CMPconst", argLength: 1, reg: gp1flags, asm: "CMP", aux: "Int64", typ: "Flags"}, // arg0 compare auxInt

		// Constants. MOVDconst may need several instructions for a
		// 64-bit value; the assembler expands it.
		{name: "MOVDconst", argLength: 0, reg: gp01, asm: "MOVD", aux: "Int64", rematerializeable: true, typ: "UInt64"},
		{name: "FMOVSconst", argLength: 0, reg: fp01, asm: "FMOVS", aux: "Float32", rematerializeable: true},
		{name: "FMOVDconst", argLength: 0, reg: fp01, asm: "FMOVD", aux: "Float64", rematerializeable: true},

		// Address of a symbol or of a stack slot.
		{name: "MOVDaddr", argLength: 1, reg: regInfo{inputs: []regMask{buildReg("SP SB")}, outputs: []regMask{gp}}, aux: "SymOff", asm: "MOVD", rematerializeable: true, symEffect: "Addr"},

		// Loads. The U forms zero-extend, the others sign-extend.
		{name: "MOVBload", argLength: 2, reg: gpload, asm: "MOVB", aux: "SymOff", typ: "Int8", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVUBload", argLength: 2, reg: gpload, asm: "MOVUB", aux: "SymOff", typ: "UInt8", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVHload", argLength: 2, reg: gpload, asm: "MOVH", aux: "SymOff", typ: "Int16", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVUHload", argLength: 2, reg: gpload, asm: "MOVUH", aux: "SymOff", typ: "UInt16", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVWload", argLength: 2, reg: gpload, asm: "MOVW", aux: "SymOff", typ: "Int32", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVUWload", argLength: 2, reg: gpload, asm: "MOVUW", aux: "SymOff", typ: "UInt32", faultOnNilArg0: true, symEffect: "Read"},
		{name: "MOVDload", argLength: 2, reg: gpload, asm: "MOVD", aux: "SymOff", typ: "UInt64", faultOnNilArg0: true, symEffect: "Read"},
		{name: "FMOVSload", argLength: 2, reg: fpload, asm: "FMOVS", aux: "SymOff", typ: "Float32", faultOnNilArg0: true, symEffect: "Read"},
		{name: "FMOVDload", argLength: 2, reg: fpload, asm: "FMOVD", aux: "SymOff", typ: "Float64", faultOnNilArg0: true, symEffect: "Read"},

		// Stores.
		{name: "MOVBstore", argLength: 3, reg: gpstore, asm: "MOVB", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVHstore", argLength: 3, reg: gpstore, asm: "MOVH", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVWstore", argLength: 3, reg: gpstore, asm: "MOVW", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "MOVDstore", argLength: 3, reg: gpstore, asm: "MOVD", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "FMOVSstore", argLength: 3, reg: fpstore, asm: "FMOVS", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},
		{name: "FMOVDstore", argLength: 3, reg: fpstore, asm: "FMOVD", aux: "SymOff", typ: "Mem", faultOnNilArg0: true, symEffect: "Write"},

		// Register moves and conversions.
		{name: "MOVD", argLength: 1, reg: gp11, asm: "MOVD"},     // move, 64 bit
		{name: "MOVW", argLength: 1, reg: gp11, asm: "MOVW"},     // sign extend int32 to int64
		{name: "MOVUW", argLength: 1, reg: gp11, asm: "MOVUW"},   // zero extend uint32 to uint64
		{name: "MOVH", argLength: 1, reg: gp11, asm: "MOVH"},     // sign extend int16 to int64
		{name: "MOVUH", argLength: 1, reg: gp11, asm: "MOVUH"},   // zero extend uint16 to uint64
		{name: "MOVB", argLength: 1, reg: gp11, asm: "MOVB"},     // sign extend int8 to int64
		{name: "MOVUB", argLength: 1, reg: gp11, asm: "MOVUB"},   // zero extend uint8 to uint64
		{name: "NEG", argLength: 1, reg: gp11, asm: "NEG"},       // -arg0

		// Floating point arithmetic.
		{name: "FADDS", argLength: 2, reg: fp21, asm: "FADDS", commutative: true},
		{name: "FADDD", argLength: 2, reg: fp21, asm: "FADDD", commutative: true},
		{name: "FSUBS", argLength: 2, reg: fp21, asm: "FSUBS"},
		{name: "FSUBD", argLength: 2, reg: fp21, asm: "FSUBD"},
		{name: "FMULS", argLength: 2, reg: fp21, asm: "FMULS", commutative: true},
		{name: "FMULD", argLength: 2, reg: fp21, asm: "FMULD", commutative: true},
		{name: "FDIVS", argLength: 2, reg: fp21, asm: "FDIVS"},
		{name: "FDIVD", argLength: 2, reg: fp21, asm: "FDIVD"},
		{name: "FNEGS", argLength: 1, reg: fp11, asm: "FNEGS"},
		{name: "FNEGD", argLength: 1, reg: fp11, asm: "FNEGD"},
		{name: "FABSS", argLength: 1, reg: fp11, asm: "FABSS"},
		{name: "FABSD", argLength: 1, reg: fp11, asm: "FABSD"},
		{name: "FSQRTS", argLength: 1, reg: fp11, asm: "FSQRTS"},
		{name: "FSQRTD", argLength: 1, reg: fp11, asm: "FSQRTD"},
		{name: "FCMPS", argLength: 2, reg: fp2flags, asm: "FCMPS", typ: "Flags"},
		{name: "FCMPD", argLength: 2, reg: fp2flags, asm: "FCMPD", typ: "Flags"},

		// Float/integer conversions.
		{name: "FSTOD", argLength: 1, reg: fp11, asm: "FSTOD"}, // float32 -> float64
		{name: "FDTOS", argLength: 1, reg: fp11, asm: "FDTOS"}, // float64 -> float32
		{name: "FSTOX", argLength: 1, reg: fp11, asm: "FSTOX"}, // float32 -> int64
		{name: "FDTOX", argLength: 1, reg: fp11, asm: "FDTOX"}, // float64 -> int64
		{name: "FXTOS", argLength: 1, reg: fp11, asm: "FXTOS"}, // int64 -> float32
		{name: "FXTOD", argLength: 1, reg: fp11, asm: "FXTOD"}, // int64 -> float64
		{name: "FMOVDgp", argLength: 1, reg: fpgp, asm: "FMOVD"}, // move float64 bits to integer register
		{name: "FMOVDfp", argLength: 1, reg: gpfp, asm: "FMOVD"}, // move integer bits to float register

		// TODO(sparc64): atomics. These need a CASD/CASW retry loop and
		// the right membar placement for the TSO model; omitted until
		// that is written, so the compiler fails loudly instead of
		// emitting something subtly wrong.

		// Calls.
		{name: "CALLstatic", argLength: -1, reg: regInfo{clobbers: callerSave}, aux: "CallOff", clobberFlags: true, call: true},
		{name: "CALLtail", argLength: -1, reg: regInfo{clobbers: callerSave}, aux: "CallOff", clobberFlags: true, call: true, tailCall: true},
		{name: "CALLclosure", argLength: -1, reg: regInfo{inputs: []regMask{gpsp, buildReg("R29"), {}}, clobbers: callerSave}, aux: "CallOff", clobberFlags: true, call: true},
		{name: "CALLinter", argLength: -1, reg: regInfo{inputs: []regMask{gp}, clobbers: callerSave}, aux: "CallOff", clobberFlags: true, call: true},

		// Pseudo-ops the compiler requires of every backend.
		{name: "LoweredNilCheck", argLength: 2, reg: regInfo{inputs: []regMask{gpg}}, nilCheck: true, faultOnNilArg0: true},
		{name: "LoweredGetClosurePtr", reg: regInfo{outputs: []regMask{buildReg("R29")}}},
		{name: "LoweredGetCallerSP", argLength: 1, reg: gp01, rematerializeable: true},
		{name: "LoweredGetCallerPC", reg: gp01, rematerializeable: true},
		// TODO(sparc64): LoweredPanicBounds* need the bounds-call ABI.
		{name: "LoweredWB", argLength: 1, reg: regInfo{clobbers: callerSave.minus(gpg), outputs: []regMask{buildReg("R25")}}, clobberFlags: true, aux: "Int64"},


		// Materialise a boolean from the condition codes. SPARC has no
		// set-on-condition instruction, so each of these becomes a
		// "MOVDconst /bin/bash; MOVcc " pair using the conditional-move
		// family; see ../../sparc64/ssa.go.
		{name: "Equal", argLength: 1, reg: readflags},
		{name: "NotEqual", argLength: 1, reg: readflags},
		{name: "LessThan", argLength: 1, reg: readflags},
		{name: "LessEqual", argLength: 1, reg: readflags},
		{name: "GreaterThan", argLength: 1, reg: readflags},
		{name: "GreaterEqual", argLength: 1, reg: readflags},
		{name: "LessThanU", argLength: 1, reg: readflags},
		{name: "LessEqualU", argLength: 1, reg: readflags},
		{name: "GreaterThanU", argLength: 1, reg: readflags},
		{name: "GreaterEqualU", argLength: 1, reg: readflags},

		// Memory barrier.
		{name: "LoweredPubBarrier", argLength: 1, asm: "MEMBAR", hasSideEffects: true},
	}

	// Blocks. SPARC branches test the condition codes set by a previous
	// CMP; the signed and unsigned orderings use different branch
	// mnemonics, so they are separate block kinds.
	blocks := []blockData{
		{name: "EQ", controls: 1},
		{name: "NE", controls: 1},
		{name: "LT", controls: 1},
		{name: "LE", controls: 1},
		{name: "GT", controls: 1},
		{name: "GE", controls: 1},
		{name: "ULT", controls: 1},
		{name: "ULE", controls: 1},
		{name: "UGT", controls: 1},
		{name: "UGE", controls: 1},
		{name: "FEQ", controls: 1},
		{name: "FNE", controls: 1},
		{name: "FLT", controls: 1},
		{name: "FLE", controls: 1},
		{name: "FGT", controls: 1},
		{name: "FGE", controls: 1},
	}

	archs = append(archs, arch{
		name:            "SPARC64",
		pkg:             "cmd/internal/obj/sparc64",
		genfile:         "../../sparc64/ssa.go",
		ops:             ops,
		blocks:          blocks,
		regnames:        regNamesSPARC64,
		gpregmask:       gp,
		fpregmask:       fp,
		framepointerreg: -1, // not used; Go frames do not chain %fp
		linkreg:         int8(num["R15"]),
	})
}
