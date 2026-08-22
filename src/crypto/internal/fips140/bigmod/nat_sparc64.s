// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !purego

#include "textflag.h"

// func addMulVVW1024(z, x *uint, y uint) (c uint)
TEXT ·addMulVVW1024(SB), NOSPLIT|NOFRAME, $0-32
	MOVD	$4, R11
	JMP	amvw<>(SB)

// func addMulVVW1536(z, x *uint, y uint) (c uint)
TEXT ·addMulVVW1536(SB), NOSPLIT|NOFRAME, $0-32
	MOVD	$6, R11
	JMP	amvw<>(SB)

// func addMulVVW2048(z, x *uint, y uint) (c uint)
TEXT ·addMulVVW2048(SB), NOSPLIT|NOFRAME, $0-32
	MOVD	$8, R11
	JMP	amvw<>(SB)

// amvw computes z[i] += x[i] * y over R11 groups of four limbs,
// returning the carry out. It is entered by a tail jump, so LR still
// holds the original caller's return address and the arguments sit at
// the same offsets.
//
// The four multiplies of a group are independent, so they issue
// together and overlap the serial carry chain that follows. Within the
// chain a limb's high word is the next limb's carry-in, so no register
// moves are needed between limbs.
TEXT amvw<>(SB), NOSPLIT|NOFRAME, $0-32
	MOVD	z+0(FP), R8
	MOVD	x+8(FP), R9
	MOVD	y+16(FP), R10
	MOVD	ZR, R12			// carry

loop:
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

	MOVD	(R8), R20
	MOVD	8(R8), R28
	MOVD	16(R8), R24
	MOVD	24(R8), R25

	// lo += z[i]; hi += carry; lo += carry-in; hi += carry. The high
	// word cannot overflow: it is at most 2**64-2, and the two carries
	// into it cannot both be set.
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
	SUBCC	$1, R11, R11
	BNED	loop

	MOVD	R12, c+24(FP)
	RET
