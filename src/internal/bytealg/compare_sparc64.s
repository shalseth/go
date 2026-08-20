// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// CMPBODY compares the byte strings described by
//	R8  = a base		R9  = a length
//	R10 = b base		R11 = b length
// and stores -1, 0 or +1 at the return slot RETOFF.
//
// SPARC traps on unaligned access, so words are only ever loaded from
// 8-aligned addresses. If both pointers share a misalignment the
// leading bytes are compared singly until they are aligned. If their
// misalignments differ, one side is aligned by hand and the other's
// words are assembled from two aligned loads and a shift; that path
// only runs while at least 16 bytes remain, so the second load cannot
// reach past the end of the shorter operand. Being big-endian, an
// unsigned comparison of two differing words gives the same answer as
// comparing their first differing byte, so no byte swap is needed.
#define CMPBODY(RETOFF) \
	CMP	R10, R8;		\
	BED	samebytes;		\
	MOVD	R9, R12;		\
	CMP	R11, R12;		\
	BLEUD	haven;			\
	MOVD	R11, R12;		\
haven:;					\
	CMP	ZR, R12;		\
	BED	samebytes;		\
	XOR	R8, R10, R13;		\
	AND	$7, R13, R13;		\
	CMP	ZR, R13;		\
	BNED	mixed;			\
	AND	$7, R8, R13;		\
	CMP	ZR, R13;		\
	BED	words;			\
	MOVD	$8, R16;		\
	SUB	R13, R16, R16;		\
	CMP	R12, R16;		\
	BLEUD	prealign;		\
	MOVD	R12, R16;		\
prealign:;				\
	MOVUB	(R8), R1;		\
	MOVUB	(R10), R2;		\
	CMP	R2, R1;			\
	BCSD	less;			\
	BGUD	greater;		\
	ADD	$1, R8;			\
	ADD	$1, R10;		\
	SUB	$1, R12;		\
	SUB	$1, R16;		\
	BRNZ	R16, prealign;		\
	CMP	ZR, R12;		\
	BED	samebytes;		\
words:;					\
	CMP	$8, R12;		\
	BCSD	bytes;			\
	MOVD	(R8), R1;		\
	MOVD	(R10), R2;		\
	CMP	R2, R1;			\
	BNED	worddiff;		\
	ADD	$8, R8;			\
	ADD	$8, R10;		\
	SUB	$8, R12;		\
	JMP	words;			\
worddiff:;				\
	BCSD	less;			\
	JMP	greater;		\
mixed:;					\
	AND	$7, R8, R13;		\
	CMP	ZR, R13;		\
	BED	mixedaligned;		\
	MOVD	$8, R16;		\
	SUB	R13, R16, R16;		\
	CMP	R12, R16;		\
	BLEUD	mixedpre;		\
	MOVD	R12, R16;		\
mixedpre:;				\
	MOVUB	(R8), R1;		\
	MOVUB	(R10), R2;		\
	CMP	R2, R1;			\
	BCSD	less;			\
	BGUD	greater;		\
	ADD	$1, R8;			\
	ADD	$1, R10;		\
	SUB	$1, R12;		\
	SUB	$1, R16;		\
	BRNZ	R16, mixedpre;		\
	CMP	ZR, R12;		\
	BED	samebytes;		\
mixedaligned:;				\
	AND	$7, R10, R17;		\
	SLLD	$3, R17, R18;		\
	MOVD	$64, R19;		\
	SUB	R18, R19, R19;		\
	SUB	R17, R10, R20;		\
	MOVD	(R20), R21;		\
mixedloop:;				\
	CMP	$16, R12;		\
	BCSD	bytes;			\
	MOVD	8(R20), R24;		\
	SLLD	R18, R21, R25;		\
	SRLD	R19, R24, R1;		\
	OR	R1, R25, R25;		\
	MOVD	(R8), R2;		\
	CMP	R25, R2;		\
	BNED	worddiff;		\
	MOVD	R24, R21;		\
	ADD	$8, R8;			\
	ADD	$8, R10;		\
	ADD	$8, R20;		\
	SUB	$8, R12;		\
	JMP	mixedloop;		\
bytes:;					\
	CMP	ZR, R12;		\
	BED	samebytes;		\
byteloop:;				\
	MOVUB	(R8), R1;		\
	MOVUB	(R10), R2;		\
	CMP	R2, R1;			\
	BCSD	less;			\
	BGUD	greater;		\
	ADD	$1, R8;			\
	ADD	$1, R10;		\
	SUB	$1, R12;		\
	CMP	ZR, R12;		\
	BNED	byteloop;		\
samebytes:;				\
	CMP	R11, R9;		\
	BCSD	less;			\
	BGUD	greater;		\
	MOVD	ZR, ret+RETOFF(FP);	\
	RET;				\
less:;					\
	MOVD	$-1, R1;		\
	MOVD	R1, ret+RETOFF(FP);	\
	RET;				\
greater:;				\
	MOVD	$1, R1;			\
	MOVD	R1, ret+RETOFF(FP);	\
	RET

// func Compare(a, b []byte) int
TEXT ·Compare(SB), NOSPLIT|NOFRAME, $0-56
	MOVD	a_base+0(FP), R8
	MOVD	a_len+8(FP), R9
	MOVD	b_base+24(FP), R10
	MOVD	b_len+32(FP), R11
	CMPBODY(48)

// func runtime.cmpstring(a, b string) int
TEXT runtime·cmpstring(SB), NOSPLIT|NOFRAME, $0-40
	MOVD	a_base+0(FP), R8
	MOVD	a_len+8(FP), R9
	MOVD	b_base+16(FP), R10
	MOVD	b_len+24(FP), R11
	CMPBODY(32)
