// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// CBODY counts occurrences of a byte, eight at a time.
//	R8 = pointer	R9 = length	R18 = byte to count
// and stores the total at RETOFF.
//
// The marker word here must be exact per lane, so it is built the other
// way: masking off each lane's high bit and adding 0x7f cannot carry out
// of the lane, so lanes stay independent. The count is then summed
// without a POPC: shifting the markers down to 0x01 per lane and multiplying by
// 0x0101010101010101 accumulates every lane into the top byte, which
// cannot overflow because there are only eight of them.
#define CBODY(RETOFF) \
	MOVD	ZR, R24;		\ // running count
	CMP	ZR, R9;			\
	BED	cbdone;			\
	MOVD	R18, R10;		\
	SLLD	$8, R10, R11;		\
	OR	R11, R10, R10;		\
	SLLD	$16, R10, R11;		\
	OR	R11, R10, R10;		\
	SLLD	$32, R10, R11;		\
	OR	R11, R10, R10;		\
	MOVD	$1, R17;		\
	SLLD	$8, R17, R11;		\
	OR	R11, R17, R17;		\
	SLLD	$16, R17, R11;		\
	OR	R11, R17, R17;		\
	SLLD	$32, R17, R11;		\
	OR	R11, R17, R17;		\
	SLLD	$7, R17, R19;		\
	SUB	R17, R19, R25;		\ // 0x7f7f7f7f7f7f7f7f
cbalign:;				\
	AND	$7, R8, R11;		\
	CMP	ZR, R11;		\
	BED	cbwords;		\
	MOVUB	(R8), R12;		\
	CMP	R12, R18;		\
	BNED	cbalignnext;		\
	ADD	$1, R24;		\
cbalignnext:;				\
	ADD	$1, R8;			\
	SUB	$1, R9;			\
	CMP	ZR, R9;			\
	BNED	cbalign;		\
	JMP	cbdone;			\
cbwords:;				\
	CMP	$8, R9;			\
	BCSD	cbtail;			\
	MOVD	(R8), R20;		\
	XOR	R10, R20, R20;		\
	AND	R25, R20, R21;		\
	ADD	R25, R21, R21;		\
	OR	R20, R21, R21;		\
	ANDN	R21, R19, R21;		\
	SRLD	$7, R21, R21;		\
	MULD	R17, R21, R21;		\
	SRLD	$56, R21, R21;		\
	ADD	R21, R24, R24;		\
	ADD	$8, R8;			\
	SUB	$8, R9;			\
	JMP	cbwords;		\
cbtail:;				\
	CMP	ZR, R9;			\
	BED	cbdone;			\
	MOVUB	(R8), R12;		\
	CMP	R12, R18;		\
	BNED	cbtailnext;		\
	ADD	$1, R24;		\
cbtailnext:;				\
	ADD	$1, R8;			\
	SUB	$1, R9;			\
	JMP	cbtail;			\
cbdone:;				\
	MOVD	R24, ret+RETOFF(FP);	\
	RET

// func Count(b []byte, c byte) int
TEXT ·Count(SB), NOSPLIT|NOFRAME, $0-40
	MOVD	b_base+0(FP), R8
	MOVD	b_len+8(FP), R9
	MOVUB	c+24(FP), R18
	CBODY(32)

// func CountString(s string, c byte) int
TEXT ·CountString(SB), NOSPLIT|NOFRAME, $0-32
	MOVD	s_base+0(FP), R8
	MOVD	s_len+8(FP), R9
	MOVUB	c+16(FP), R18
	CBODY(24)
