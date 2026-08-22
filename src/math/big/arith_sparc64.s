// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !math_big_pure_go

#include "textflag.h"

// SPARC V9's ADDC and SUBC use the 32-bit carry in icc, so they are
// useless for 64-bit limbs. VIS3's ADDXC and ADDXCcc use the 64-bit
// carry in xcc instead, and ADDXCcc both uses and sets it, so a carry
// chain can run one instruction per limb with the carry never leaving
// the condition codes. Nothing between the adds may disturb xcc: ADD,
// SUB, AND, XNOR, the shifts, the loads and stores, and the register
// branches BRZ and BRNZ all leave it alone.
//
// There is no 64-bit subtract-with-borrow. Borrow chains therefore run
// as add chains over the complement, since x - y - b is x + ^y + (1-b)
// and the carry out is one minus the borrow out.

// func addVV(z, x, y []Word) (c Word)
TEXT ·addVV(SB), NOSPLIT|NOFRAME, $0-80
	MOVD	z+0(FP), R8
	MOVD	z_len+8(FP), R11
	MOVD	x+24(FP), R9
	MOVD	y+48(FP), R10
	SRLD	$2, R11, R12		// groups of four
	AND	$3, R11, R11		// limbs left over
	ADDCC	ZR, ZR, ZR		// carry = 0
	BRZ	R12, addVVtail

addVVgroup:
	MOVD	(R9), R1
	MOVD	8(R9), R2
	MOVD	16(R9), R3
	MOVD	24(R9), R4
	MOVD	(R10), R16
	MOVD	8(R10), R17
	MOVD	16(R10), R18
	MOVD	24(R10), R19
	ADDXCCC	R16, R1, R1
	ADDXCCC	R17, R2, R2
	ADDXCCC	R18, R3, R3
	ADDXCCC	R19, R4, R4
	MOVD	R1, (R8)
	MOVD	R2, 8(R8)
	MOVD	R3, 16(R8)
	MOVD	R4, 24(R8)
	ADD	$32, R8
	ADD	$32, R9
	ADD	$32, R10
	SUB	$1, R12
	BRNZ	R12, addVVgroup

addVVtail:
	BRZ	R11, addVVdone

addVVloop:
	MOVD	(R9), R1
	MOVD	(R10), R2
	ADDXCCC	R2, R1, R1
	MOVD	R1, (R8)
	ADD	$8, R8
	ADD	$8, R9
	ADD	$8, R10
	SUB	$1, R11
	BRNZ	R11, addVVloop

addVVdone:
	ADDXC	ZR, ZR, R1
	MOVD	R1, c+72(FP)
	RET

// func subVV(z, x, y []Word) (c Word)
TEXT ·subVV(SB), NOSPLIT|NOFRAME, $0-80
	MOVD	z+0(FP), R8
	MOVD	z_len+8(FP), R11
	MOVD	x+24(FP), R9
	MOVD	y+48(FP), R10
	SRLD	$2, R11, R12
	AND	$3, R11, R11
	SUBCC	$1, ZR, ZR		// 0-1 borrows: carry = 1, no borrow yet
	BRZ	R12, subVVtail

subVVgroup:
	MOVD	(R9), R1
	MOVD	8(R9), R2
	MOVD	16(R9), R3
	MOVD	24(R9), R4
	MOVD	(R10), R16
	MOVD	8(R10), R17
	MOVD	16(R10), R18
	MOVD	24(R10), R19
	XNOR	R16, ZR, R16
	XNOR	R17, ZR, R17
	XNOR	R18, ZR, R18
	XNOR	R19, ZR, R19
	ADDXCCC	R16, R1, R1
	ADDXCCC	R17, R2, R2
	ADDXCCC	R18, R3, R3
	ADDXCCC	R19, R4, R4
	MOVD	R1, (R8)
	MOVD	R2, 8(R8)
	MOVD	R3, 16(R8)
	MOVD	R4, 24(R8)
	ADD	$32, R8
	ADD	$32, R9
	ADD	$32, R10
	SUB	$1, R12
	BRNZ	R12, subVVgroup

subVVtail:
	BRZ	R11, subVVdone

subVVloop:
	MOVD	(R9), R1
	MOVD	(R10), R2
	XNOR	R2, ZR, R2
	ADDXCCC	R2, R1, R1
	MOVD	R1, (R8)
	ADD	$8, R8
	ADD	$8, R9
	ADD	$8, R10
	SUB	$1, R11
	BRNZ	R11, subVVloop

subVVdone:
	ADDXC	ZR, ZR, R1
	XOR	$1, R1, R1		// carry out 1 means no borrow
	MOVD	R1, c+72(FP)
	RET

// func mulAddVWW(z, x []Word, m, a Word) (c Word)
//
// z[i] = x[i]*y + carry, carrying the high word. The four multiplies of
// a group are independent and issue together; the carry chain that
// follows them is serial, and a limb's high word is the next limb's
// carry-in, so no moves are needed between limbs.
TEXT ·mulAddVWW(SB), NOSPLIT|NOFRAME, $0-72
	MOVD	z+0(FP), R8
	MOVD	z_len+8(FP), R11
	MOVD	x+24(FP), R9
	MOVD	m+48(FP), R10
	MOVD	a+56(FP), R12		// carry
	SRLD	$2, R11, R13
	AND	$3, R11, R11
	BRZ	R13, mulAddtail

mulAddgroup:
	MOVD	(R9), R1
	MOVD	8(R9), R2
	MOVD	16(R9), R3
	MOVD	24(R9), R4
	UMULXHI	R10, R1, R16
	UMULXHI	R10, R2, R17
	UMULXHI	R10, R3, R18
	UMULXHI	R10, R4, R19
	MULD	R10, R1, R1
	MULD	R10, R2, R2
	MULD	R10, R3, R3
	MULD	R10, R4, R4
	ADDCC	R12, R1, R1
	ADDXC	ZR, R16, R16
	ADDCC	R16, R2, R2
	ADDXC	ZR, R17, R17
	ADDCC	R17, R3, R3
	ADDXC	ZR, R18, R18
	ADDCC	R18, R4, R4
	ADDXC	ZR, R19, R19
	MOVD	R1, (R8)
	MOVD	R2, 8(R8)
	MOVD	R3, 16(R8)
	MOVD	R4, 24(R8)
	MOVD	R19, R12
	ADD	$32, R8
	ADD	$32, R9
	SUB	$1, R13
	BRNZ	R13, mulAddgroup

mulAddtail:
	BRZ	R11, mulAdddone

mulAddloop:
	MOVD	(R9), R1
	UMULXHI	R10, R1, R16
	MULD	R10, R1, R1
	ADDCC	R12, R1, R1
	ADDXC	ZR, R16, R12
	MOVD	R1, (R8)
	ADD	$8, R8
	ADD	$8, R9
	SUB	$1, R11
	BRNZ	R11, mulAddloop

mulAdddone:
	MOVD	R12, c+64(FP)
	RET

// func addMulVVWW(z, x, y []Word, m, a Word) (c Word)
//
// z[i] = x[i] + y[i]*m + carry.
TEXT ·addMulVVWW(SB), NOSPLIT|NOFRAME, $0-96
	MOVD	z+0(FP), R8
	MOVD	z_len+8(FP), R11
	MOVD	x+24(FP), R9
	MOVD	y+48(FP), R5
	MOVD	m+72(FP), R10
	MOVD	a+80(FP), R12		// carry
	SRLD	$2, R11, R13
	AND	$3, R11, R11
	BRZ	R13, addMultail

addMulgroup:
	MOVD	(R5), R1
	MOVD	8(R5), R2
	MOVD	16(R5), R3
	MOVD	24(R5), R4
	UMULXHI	R10, R1, R16
	UMULXHI	R10, R2, R17
	UMULXHI	R10, R3, R18
	UMULXHI	R10, R4, R19
	MULD	R10, R1, R1
	MULD	R10, R2, R2
	MULD	R10, R3, R3
	MULD	R10, R4, R4
	MOVD	(R9), R20
	MOVD	8(R9), R28
	MOVD	16(R9), R24
	MOVD	24(R9), R25
	ADDCC	R20, R1, R1
	ADDXC	ZR, R16, R16
	ADDCC	R12, R1, R1
	ADDXC	ZR, R16, R16
	ADDCC	R28, R2, R2
	ADDXC	ZR, R17, R17
	ADDCC	R16, R2, R2
	ADDXC	ZR, R17, R17
	ADDCC	R24, R3, R3
	ADDXC	ZR, R18, R18
	ADDCC	R17, R3, R3
	ADDXC	ZR, R18, R18
	ADDCC	R25, R4, R4
	ADDXC	ZR, R19, R19
	ADDCC	R18, R4, R4
	ADDXC	ZR, R19, R19
	MOVD	R1, (R8)
	MOVD	R2, 8(R8)
	MOVD	R3, 16(R8)
	MOVD	R4, 24(R8)
	MOVD	R19, R12
	ADD	$32, R8
	ADD	$32, R9
	ADD	$32, R5
	SUB	$1, R13
	BRNZ	R13, addMulgroup

addMultail:
	BRZ	R11, addMuldone

addMulloop:
	MOVD	(R5), R1
	MOVD	(R9), R20
	UMULXHI	R10, R1, R16
	MULD	R10, R1, R1
	ADDCC	R20, R1, R1
	ADDXC	ZR, R16, R16
	ADDCC	R12, R1, R1
	ADDXC	ZR, R16, R12
	MOVD	R1, (R8)
	ADD	$8, R8
	ADD	$8, R9
	ADD	$8, R5
	SUB	$1, R11
	BRNZ	R11, addMulloop

addMuldone:
	MOVD	R12, c+88(FP)
	RET

// func lshVU(z, x []Word, s uint) (c Word)
//
// Shifts left, descending, so that z and x may be the same slice.
TEXT ·lshVU(SB), NOSPLIT|NOFRAME, $0-64
	MOVD	z+0(FP), R8
	MOVD	z_len+8(FP), R11
	MOVD	x+24(FP), R9
	MOVD	s+48(FP), R10
	MOVD	ZR, R5			// c
	BRZ	R11, lshdone
	BRZ	R10, lshcopy
	AND	$63, R10, R10
	MOVD	$64, R12
	SUB	R10, R12, R12		// 64 - s
	SLLD	$3, R11, R3
	ADD	R3, R9, R2		// &x[n]
	ADD	R3, R8, R1		// &z[n]
	MOVD	-8(R2), R3
	SRLD	R12, R3, R5		// c = x[n-1] >> (64-s)
	SUB	$1, R11
	BRZ	R11, lshfirst

lshloop:
	MOVD	-8(R2), R3		// x[i]
	MOVD	-16(R2), R4		// x[i-1]
	SLLD	R10, R3, R3
	SRLD	R12, R4, R4
	OR	R4, R3, R3
	MOVD	R3, -8(R1)		// z[i]
	SUB	$8, R2
	SUB	$8, R1
	SUB	$1, R11
	BRNZ	R11, lshloop

lshfirst:
	MOVD	(R9), R3
	SLLD	R10, R3, R3
	MOVD	R3, (R8)		// z[0] = x[0] << s

lshdone:
	MOVD	R5, c+56(FP)
	RET

lshcopy:
	MOVD	(R9), R3
	MOVD	R3, (R8)
	ADD	$8, R9
	ADD	$8, R8
	SUB	$1, R11
	BRNZ	R11, lshcopy
	MOVD	ZR, c+56(FP)
	RET

// func rshVU(z, x []Word, s uint) (c Word)
//
// Shifts right, ascending, so that z and x may be the same slice.
TEXT ·rshVU(SB), NOSPLIT|NOFRAME, $0-64
	MOVD	z+0(FP), R8
	MOVD	z_len+8(FP), R11
	MOVD	x+24(FP), R9
	MOVD	s+48(FP), R10
	MOVD	ZR, R5			// c
	BRZ	R11, rshdone
	BRZ	R10, rshcopy
	AND	$63, R10, R10
	MOVD	$64, R12
	SUB	R10, R12, R12		// 64 - s
	MOVD	(R9), R3
	SLLD	R12, R3, R5		// c = x[0] << (64-s)
	MOVD	R9, R2
	MOVD	R8, R1
	SUB	$1, R11
	BRZ	R11, rshlast

rshloop:
	MOVD	(R2), R3		// x[i-1]
	MOVD	8(R2), R4		// x[i]
	SRLD	R10, R3, R3
	SLLD	R12, R4, R4
	OR	R4, R3, R3
	MOVD	R3, (R1)		// z[i-1]
	ADD	$8, R2
	ADD	$8, R1
	SUB	$1, R11
	BRNZ	R11, rshloop

rshlast:
	MOVD	(R2), R3
	SRLD	R10, R3, R3
	MOVD	R3, (R1)		// z[n-1] = x[n-1] >> s

rshdone:
	MOVD	R5, c+56(FP)
	RET

rshcopy:
	MOVD	(R9), R3
	MOVD	R3, (R8)
	ADD	$8, R9
	ADD	$8, R8
	SUB	$1, R11
	BRNZ	R11, rshcopy
	MOVD	ZR, c+56(FP)
	RET
