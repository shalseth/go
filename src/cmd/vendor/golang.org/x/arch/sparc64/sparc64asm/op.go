// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sparc64asm

import "fmt"

const (
	_ Op = iota
	ADD
	ADDcc
	ADDC
	ADDCcc
	AND
	ANDcc
	ANDN
	ANDNcc
	OR
	ORcc
	ORN
	ORNcc
	XOR
	XORcc
	XNOR
	XNORcc
	SUB
	SUBcc
	SUBC
	SUBCcc
	MULX
	UMUL
	UMULcc
	SMUL
	SMULcc
	UDIVX
	SDIVX
	UDIV
	UDIVcc
	SDIV
	SDIVcc
	SLL
	SRL
	SRA
	SLLX
	SRLX
	SRAX
	JMPL
	RETURN
	SAVE
	RESTORE
	FLUSHW
	SETHI
	NOP
	CALL
	BA
	BN
	BNE
	BE
	BG
	BLE
	BGE
	BL
	BGU
	BLEU
	BCC
	BCS
	BPOS
	BNEG
	BVC
	BVS
	BRZ
	BRLEZ
	BRLZ
	BRNZ
	BRGZ
	BRGEZ
	MOVA
	MOVN
	MOVNE
	MOVE
	MOVG
	MOVLE
	MOVGE
	MOVL
	MOVGU
	MOVLEU
	MOVCC
	MOVCS
	MOVPOS
	MOVNEG
	MOVVC
	MOVVS
	MOVLG
	MOVUL
	MOVUG
	MOVUE
	MOVUGE
	MOVULE
	MOVO
	MOVU
	MOVRZ
	MOVRLEZ
	MOVRLZ
	MOVRNZ
	MOVRGZ
	MOVRGEZ
	LDSB
	LDSH
	LDSW
	LDUB
	LDUH
	LDUW
	LDX
	LDD
	LDF
	LDDF
	LDQF
	LDFSR
	LDXFSR
	STB
	STH
	STW
	STX
	STD
	STF
	STDF
	STQF
	STFSR
	STXFSR
	LDSBA
	LDSHA
	LDSWA
	LDUBA
	LDUHA
	LDUWA
	LDXA
	LDDA
	STBA
	STHA
	STWA
	STXA
	STDA
	CASA
	CASXA
	PREFETCH
	PREFETCHA
	MEMBAR
	STBAR
	RDY
	WRY
	RDCCR
	WRCCR
	RDASI
	WRASI
	RDPC
	RDFPRS
	WRFPRS
	POPC
	UMULXHI
	ADDXC
	ADDXCcc
	FADDs
	FADDd
	FSUBs
	FSUBd
	FMULs
	FMULd
	FDIVs
	FDIVd
	FSQRTs
	FSQRTd
	FMOVs
	FMOVd
	FNEGs
	FNEGd
	FABSs
	FABSd
	FCMPs
	FCMPd
	FCMPEs
	FCMPEd
	FSTOI
	FDTOI
	FITOS
	FITOD
	FSTOD
	FDTOS
	FSTOX
	FDTOX
	FXTOS
	FXTOD
	FBA
	FBN
	FBU
	FBG
	FBUG
	FBL
	FBUL
	FBLG
	FBNE
	FBE
	FBUE
	FBGE
	FBUGE
	FBLE
	FBULE
	FBO
	ILLTRAP
	Tcc
	SHA256
)

var opstr = [...]string{
	ADD: "add",
	ADDcc: "addcc",
	ADDC: "addc",
	ADDCcc: "addccc",
	AND: "and",
	ANDcc: "andcc",
	ANDN: "andn",
	ANDNcc: "andncc",
	OR: "or",
	ORcc: "orcc",
	ORN: "orn",
	ORNcc: "orncc",
	XOR: "xor",
	XORcc: "xorcc",
	XNOR: "xnor",
	XNORcc: "xnorcc",
	SUB: "sub",
	SUBcc: "subcc",
	SUBC: "subc",
	SUBCcc: "subccc",
	MULX: "mulx",
	UMUL: "umul",
	UMULcc: "umulcc",
	SMUL: "smul",
	SMULcc: "smulcc",
	UDIVX: "udivx",
	SDIVX: "sdivx",
	UDIV: "udiv",
	UDIVcc: "udivcc",
	SDIV: "sdiv",
	SDIVcc: "sdivcc",
	SLL: "sll",
	SRL: "srl",
	SRA: "sra",
	SLLX: "sllx",
	SRLX: "srlx",
	SRAX: "srax",
	JMPL: "jmpl",
	RETURN: "return",
	SAVE: "save",
	RESTORE: "restore",
	FLUSHW: "flushw",
	SETHI: "sethi",
	NOP: "nop",
	CALL: "call",
	BA: "ba",
	BN: "bn",
	BNE: "bne",
	BE: "be",
	BG: "bg",
	BLE: "ble",
	BGE: "bge",
	BL: "bl",
	BGU: "bgu",
	BLEU: "bleu",
	BCC: "bcc",
	BCS: "bcs",
	BPOS: "bpos",
	BNEG: "bneg",
	BVC: "bvc",
	BVS: "bvs",
	BRZ: "brz",
	BRLEZ: "brlez",
	BRLZ: "brlz",
	BRNZ: "brnz",
	BRGZ: "brgz",
	BRGEZ: "brgez",
	MOVA: "mova",
	MOVN: "movn",
	MOVNE: "movne",
	MOVE: "move",
	MOVG: "movg",
	MOVLE: "movle",
	MOVGE: "movge",
	MOVL: "movl",
	MOVGU: "movgu",
	MOVLEU: "movleu",
	MOVCC: "movcc",
	MOVCS: "movcs",
	MOVPOS: "movpos",
	MOVNEG: "movneg",
	MOVVC: "movvc",
	MOVVS: "movvs",
	MOVLG: "movlg",
	MOVUL: "movul",
	MOVUG: "movug",
	MOVUE: "movue",
	MOVUGE: "movuge",
	MOVULE: "movule",
	MOVO: "movo",
	MOVU: "movu",
	MOVRZ: "movrz",
	MOVRLEZ: "movrlez",
	MOVRLZ: "movrlz",
	MOVRNZ: "movrnz",
	MOVRGZ: "movrgz",
	MOVRGEZ: "movrgez",
	LDSB: "ldsb",
	LDSH: "ldsh",
	LDSW: "ldsw",
	LDUB: "ldub",
	LDUH: "lduh",
	LDUW: "lduw",
	LDX: "ldx",
	LDD: "ldd",
	LDF: "ldf",
	LDDF: "lddf",
	LDQF: "ldqf",
	LDFSR: "ldfsr",
	LDXFSR: "ldxfsr",
	STB: "stb",
	STH: "sth",
	STW: "stw",
	STX: "stx",
	STD: "std",
	STF: "stf",
	STDF: "stdf",
	STQF: "stqf",
	STFSR: "stfsr",
	STXFSR: "stxfsr",
	LDSBA: "ldsba",
	LDSHA: "ldsha",
	LDSWA: "ldswa",
	LDUBA: "lduba",
	LDUHA: "lduha",
	LDUWA: "lduwa",
	LDXA: "ldxa",
	LDDA: "ldda",
	STBA: "stba",
	STHA: "stha",
	STWA: "stwa",
	STXA: "stxa",
	STDA: "stda",
	CASA: "casa",
	CASXA: "casxa",
	PREFETCH: "prefetch",
	PREFETCHA: "prefetcha",
	MEMBAR: "membar",
	STBAR: "stbar",
	RDY: "rdy",
	WRY: "wry",
	RDCCR: "rdccr",
	WRCCR: "wrccr",
	RDASI: "rdasi",
	WRASI: "wrasi",
	RDPC: "rdpc",
	RDFPRS: "rdfprs",
	WRFPRS: "wrfprs",
	POPC: "popc",
	UMULXHI: "umulxhi",
	ADDXC: "addxc",
	ADDXCcc: "addxccc",
	FADDs: "fadds",
	FADDd: "faddd",
	FSUBs: "fsubs",
	FSUBd: "fsubd",
	FMULs: "fmuls",
	FMULd: "fmuld",
	FDIVs: "fdivs",
	FDIVd: "fdivd",
	FSQRTs: "fsqrts",
	FSQRTd: "fsqrtd",
	FMOVs: "fmovs",
	FMOVd: "fmovd",
	FNEGs: "fnegs",
	FNEGd: "fnegd",
	FABSs: "fabss",
	FABSd: "fabsd",
	FCMPs: "fcmps",
	FCMPd: "fcmpd",
	FCMPEs: "fcmpes",
	FCMPEd: "fcmped",
	FSTOI: "fstoi",
	FDTOI: "fdtoi",
	FITOS: "fitos",
	FITOD: "fitod",
	FSTOD: "fstod",
	FDTOS: "fdtos",
	FSTOX: "fstox",
	FDTOX: "fdtox",
	FXTOS: "fxtos",
	FXTOD: "fxtod",
	FBA: "fba",
	FBN: "fbn",
	FBU: "fbu",
	FBG: "fbg",
	FBUG: "fbug",
	FBL: "fbl",
	FBUL: "fbul",
	FBLG: "fblg",
	FBNE: "fbne",
	FBE: "fbe",
	FBUE: "fbue",
	FBGE: "fbge",
	FBUGE: "fbuge",
	FBLE: "fble",
	FBULE: "fbule",
	FBO: "fbo",
	ILLTRAP: "illtrap",
	Tcc: "tcc",
	SHA256: "sha256",
}

func (op Op) String() string {
	if int(op) < len(opstr) && opstr[op] != "" {
		return opstr[op]
	}
	return fmt.Sprintf("Op(%d)", int(op))
}
