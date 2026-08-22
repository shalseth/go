// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// IBODY scans for a byte, eight at a time.
//	R8 = pointer	R9 = length	R18 = byte to find
// and stores the index, or -1, at RETOFF.
//
// The word test is the usual one: xor the word with the byte broadcast
// across all eight lanes, so a match becomes a zero byte, then
// (v - 0x01..01) &^ v & 0x80..80 says whether any lane is zero. That
// test is exact as a yes/no answer but not per lane - the subtraction
// borrows across lanes, so a non-matching lane can pick up a marker -
// which does not matter here, because the lane is found by rescanning
// rather than read out of the mask. CBODY below, which does read the
// mask, uses the more careful form.
// Which lane matched is resolved by rescanning those eight bytes one at
// a time - that costs a few instructions once per call, not per word,
// which is why this needs no LZD.
//
// SPARC traps on unaligned access, so the leading bytes are scanned
// singly until the pointer is 8-aligned.
#define IBODY(RETOFF) \
	MOVD	R8, R16;		\ // base, for the index
	CMP	ZR, R9;			\
	BED	ibnone;			\
	/* broadcast the byte into all eight lanes */	\
	MOVD	R18, R10;		\
	SLLD	$8, R10, R11;		\
	OR	R11, R10, R10;		\
	SLLD	$16, R10, R11;		\
	OR	R11, R10, R10;		\
	SLLD	$32, R10, R11;		\
	OR	R11, R10, R10;		\
	/* 0x0101010101010101 and 0x8080808080808080 */	\
	MOVD	$1, R17;		\
	SLLD	$8, R17, R11;		\
	OR	R11, R17, R17;		\
	SLLD	$16, R17, R11;		\
	OR	R11, R17, R17;		\
	SLLD	$32, R17, R11;		\
	OR	R11, R17, R17;		\
	SLLD	$7, R17, R19;		\
ibalign:;				\
	AND	$7, R8, R11;		\
	CMP	ZR, R11;		\
	BED	ibwords;		\
	MOVUB	(R8), R12;		\
	CMP	R12, R18;		\
	BED	ibhit;			\
	ADD	$1, R8;			\
	SUB	$1, R9;			\
	CMP	ZR, R9;			\
	BNED	ibalign;		\
	JMP	ibnone;			\
ibwords:;				\
	CMP	$8, R9;			\
	BCSD	ibtail;			\
	MOVD	(R8), R20;		\
	XOR	R10, R20, R20;		\
	SUB	R17, R20, R28;		\
	ANDN	R20, R28, R28;		\
	AND	R19, R28, R28;		\
	CMP	ZR, R28;		\
	BNED	ibfound;		\
	ADD	$8, R8;			\
	SUB	$8, R9;			\
	JMP	ibwords;		\
ibfound:;				\
	MOVD	$8, R9;			\ // the match is inside these eight
ibtail:;				\
	CMP	ZR, R9;			\
	BED	ibnone;			\
	MOVUB	(R8), R12;		\
	CMP	R12, R18;		\
	BED	ibhit;			\
	ADD	$1, R8;			\
	SUB	$1, R9;			\
	JMP	ibtail;			\
ibhit:;					\
	SUB	R16, R8, R8;		\
	MOVD	R8, ret+RETOFF(FP);	\
	RET;				\
ibnone:;				\
	MOVD	$-1, R8;		\
	MOVD	R8, ret+RETOFF(FP);	\
	RET

// func IndexByte(b []byte, c byte) int
TEXT ·IndexByte(SB), NOSPLIT|NOFRAME, $0-40
	MOVD	b_base+0(FP), R8
	MOVD	b_len+8(FP), R9
	MOVUB	c+24(FP), R18
	IBODY(32)

// func IndexByteString(s string, c byte) int
TEXT ·IndexByteString(SB), NOSPLIT|NOFRAME, $0-32
	MOVD	s_base+0(FP), R8
	MOVD	s_len+8(FP), R9
	MOVUB	c+16(FP), R18
	IBODY(24)

