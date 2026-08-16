// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// Memory ordering notes.
//
// SPARC V9 under Linux runs in TSO, which already orders Load->Load,
// Load->Store and Store->Store. Only Store->Load may be reordered, so:
//
//   - plain loads give acquire semantics for free, hence Load and
//     LoadAcq are the same instruction with no barrier;
//   - plain stores give release semantics for free, hence StoreRel and
//     StorepNoWB need no barrier;
//   - Store is specified as sequentially consistent, so it takes a
//     trailing "membar #StoreLoad" — the same reason amd64, also TSO,
//     implements Store with an XCHG rather than a plain MOV;
//   - the read-modify-write operations take a full barrier after the
//     CAS so they are sequentially consistent.
//
// CASA/CASXA semantics: "casx [rs1], rs2, rd" compares the memory word
// at rs1 against rs2 and, if they are equal, swaps it with rd. rd
// always receives the original memory value, so a CAS succeeded exactly
// when rd comes back equal to rs2.
//
// The 8-bit read-modify-write operations are done with a 32-bit CAS on
// the containing aligned word, because SPARC has no byte-width CAS.
// SPARC is big-endian, so the byte at address p sits at bit position
// (3 - (p & 3)) * 8 within that word.

#define MEMBAR_FULL	MEMBAR	$15	// #LoadLoad|#StoreLoad|#LoadStore|#StoreStore
#define MEMBAR_SL	MEMBAR	$2	// #StoreLoad

// func Cas(ptr *uint32, old, new uint32) bool
TEXT ·Cas(SB), NOSPLIT, $0-17
	MOVD	ptr+0(FP), R8
	MOVUW	old+8(FP), R9
	MOVUW	new+12(FP), R10
	CASW	(R8), R9, R10
	MEMBAR_FULL
	SUBCC	R10, R9, ZR
	MOVD	ZR, R11
	MOVE	XCC, $1, R11
	MOVB	R11, ret+16(FP)
	RET

// func Cas64(ptr *uint64, old, new uint64) bool
TEXT ·Cas64(SB), NOSPLIT, $0-25
	MOVD	ptr+0(FP), R8
	MOVD	old+8(FP), R9
	MOVD	new+16(FP), R10
	CASD	(R8), R9, R10
	MEMBAR_FULL
	SUBCC	R10, R9, ZR
	MOVD	ZR, R11
	MOVE	XCC, $1, R11
	MOVB	R11, ret+24(FP)
	RET

// func Load(ptr *uint32) uint32
TEXT ·Load(SB), NOSPLIT, $0-12
	MOVD	ptr+0(FP), R8
	MOVUW	(R8), R9
	MOVW	R9, ret+8(FP)
	RET

// func Load8(ptr *uint8) uint8
TEXT ·Load8(SB), NOSPLIT, $0-9
	MOVD	ptr+0(FP), R8
	MOVUB	(R8), R9
	MOVB	R9, ret+8(FP)
	RET

// func Load64(ptr *uint64) uint64
TEXT ·Load64(SB), NOSPLIT, $0-16
	MOVD	ptr+0(FP), R8
	MOVD	(R8), R9
	MOVD	R9, ret+8(FP)
	RET

// func Loadp(ptr unsafe.Pointer) unsafe.Pointer
TEXT ·Loadp(SB), NOSPLIT, $0-16
	MOVD	ptr+0(FP), R8
	MOVD	(R8), R9
	MOVD	R9, ret+8(FP)
	RET

// func Store(ptr *uint32, val uint32)
TEXT ·Store(SB), NOSPLIT, $0-12
	MOVD	ptr+0(FP), R8
	MOVUW	val+8(FP), R9
	MOVW	R9, (R8)
	MEMBAR_SL
	RET

// func Store8(ptr *uint8, val uint8)
TEXT ·Store8(SB), NOSPLIT, $0-9
	MOVD	ptr+0(FP), R8
	MOVUB	val+8(FP), R9
	MOVB	R9, (R8)
	MEMBAR_SL
	RET

// func Store64(ptr *uint64, val uint64)
TEXT ·Store64(SB), NOSPLIT, $0-16
	MOVD	ptr+0(FP), R8
	MOVD	val+8(FP), R9
	MOVD	R9, (R8)
	MEMBAR_SL
	RET

// func StoreRel(ptr *uint32, val uint32)
TEXT ·StoreRel(SB), NOSPLIT, $0-12
	MOVD	ptr+0(FP), R8
	MOVUW	val+8(FP), R9
	MOVW	R9, (R8)
	RET

// func StoreRel64(ptr *uint64, val uint64)
TEXT ·StoreRel64(SB), NOSPLIT, $0-16
	MOVD	ptr+0(FP), R8
	MOVD	val+8(FP), R9
	MOVD	R9, (R8)
	RET

// func StorepNoWB(ptr unsafe.Pointer, val unsafe.Pointer)
TEXT ·StorepNoWB(SB), NOSPLIT, $0-16
	MOVD	ptr+0(FP), R8
	MOVD	val+8(FP), R9
	MOVD	R9, (R8)
	RET

// func Xchg(ptr *uint32, new uint32) uint32
TEXT ·Xchg(SB), NOSPLIT, $0-20
	MOVD	ptr+0(FP), R8
	MOVUW	new+8(FP), R9
xchg_again:
	MOVUW	(R8), R10
	MOVD	R9, R11
	CASW	(R8), R10, R11
	SUBCC	R11, R10, ZR
	BNED	xchg_again
	MEMBAR_FULL
	MOVW	R10, ret+16(FP)
	RET

// func Xchg64(ptr *uint64, new uint64) uint64
TEXT ·Xchg64(SB), NOSPLIT, $0-24
	MOVD	ptr+0(FP), R8
	MOVD	new+8(FP), R9
xchg64_again:
	MOVD	(R8), R10
	MOVD	R9, R11
	CASD	(R8), R10, R11
	SUBCC	R11, R10, ZR
	BNED	xchg64_again
	MEMBAR_FULL
	MOVD	R10, ret+16(FP)
	RET

// func Xadd(ptr *uint32, delta int32) uint32
TEXT ·Xadd(SB), NOSPLIT, $0-20
	MOVD	ptr+0(FP), R8
	MOVW	delta+8(FP), R9
xadd_again:
	MOVUW	(R8), R10
	ADD	R9, R10, R11
	MOVD	R11, R12
	CASW	(R8), R10, R12
	SUBCC	R12, R10, ZR
	BNED	xadd_again
	MEMBAR_FULL
	MOVW	R11, ret+16(FP)
	RET

// func Xadd64(ptr *uint64, delta int64) uint64
TEXT ·Xadd64(SB), NOSPLIT, $0-24
	MOVD	ptr+0(FP), R8
	MOVD	delta+8(FP), R9
xadd64_again:
	MOVD	(R8), R10
	ADD	R9, R10, R11
	MOVD	R11, R12
	CASD	(R8), R10, R12
	SUBCC	R12, R10, ZR
	BNED	xadd64_again
	MEMBAR_FULL
	MOVD	R11, ret+16(FP)
	RET

// func And32(ptr *uint32, val uint32) uint32
TEXT ·And32(SB), NOSPLIT, $0-20
	MOVD	ptr+0(FP), R8
	MOVUW	val+8(FP), R9
and32_again:
	MOVUW	(R8), R10
	AND	R9, R10, R11
	MOVD	R11, R12
	CASW	(R8), R10, R12
	SUBCC	R12, R10, ZR
	BNED	and32_again
	MEMBAR_FULL
	MOVW	R11, ret+16(FP)
	RET

// func Or32(ptr *uint32, val uint32) uint32
TEXT ·Or32(SB), NOSPLIT, $0-20
	MOVD	ptr+0(FP), R8
	MOVUW	val+8(FP), R9
or32_again:
	MOVUW	(R8), R10
	OR	R9, R10, R11
	MOVD	R11, R12
	CASW	(R8), R10, R12
	SUBCC	R12, R10, ZR
	BNED	or32_again
	MEMBAR_FULL
	MOVW	R11, ret+16(FP)
	RET

// func And64(ptr *uint64, val uint64) uint64
TEXT ·And64(SB), NOSPLIT, $0-24
	MOVD	ptr+0(FP), R8
	MOVD	val+8(FP), R9
and64_again:
	MOVD	(R8), R10
	AND	R9, R10, R11
	MOVD	R11, R12
	CASD	(R8), R10, R12
	SUBCC	R12, R10, ZR
	BNED	and64_again
	MEMBAR_FULL
	MOVD	R11, ret+16(FP)
	RET

// func Or64(ptr *uint64, val uint64) uint64
TEXT ·Or64(SB), NOSPLIT, $0-24
	MOVD	ptr+0(FP), R8
	MOVD	val+8(FP), R9
or64_again:
	MOVD	(R8), R10
	OR	R9, R10, R11
	MOVD	R11, R12
	CASD	(R8), R10, R12
	SUBCC	R12, R10, ZR
	BNED	or64_again
	MEMBAR_FULL
	MOVD	R11, ret+16(FP)
	RET

// byteSetup computes, for the byte pointer in R8:
//	R8  = the containing aligned word address
//	R1  = the shift that moves a byte into position
//	R2  = 0xff << shift, the byte's mask within the word
// It clobbers R3.
#define byteSetup \
	AND	$3, R8, R1;	\
	MOVD	$3, R3;		\
	SUB	R1, R3, R1;	\
	SLLD	$3, R1, R1;	\
	AND	$-4, R8, R8;	\
	MOVD	$255, R2;	\
	SLLD	R1, R2, R2

// func And8(ptr *uint8, val uint8)
TEXT ·And8(SB), NOSPLIT, $0-9
	MOVD	ptr+0(FP), R8
	MOVUB	val+8(FP), R9
	byteSetup
	SLLD	R1, R9, R9		// val in position
	XOR	$-1, R2, R3		// ~mask: the bytes to leave alone
	OR	R9, R3, R9		// keep = ~mask | (val<<shift)
and8_again:
	MOVUW	(R8), R10
	AND	R9, R10, R11
	MOVD	R11, R12
	CASW	(R8), R10, R12
	SUBCC	R12, R10, ZR
	BNED	and8_again
	MEMBAR_FULL
	RET

// func Or8(ptr *uint8, val uint8)
TEXT ·Or8(SB), NOSPLIT, $0-9
	MOVD	ptr+0(FP), R8
	MOVUB	val+8(FP), R9
	byteSetup
	SLLD	R1, R9, R9		// val in position; other bytes are 0
or8_again:
	MOVUW	(R8), R10
	OR	R9, R10, R11
	MOVD	R11, R12
	CASW	(R8), R10, R12
	SUBCC	R12, R10, ZR
	BNED	or8_again
	MEMBAR_FULL
	RET

// func Xchg8(ptr *uint8, new uint8) uint8
TEXT ·Xchg8(SB), NOSPLIT, $0-17
	MOVD	ptr+0(FP), R8
	MOVUB	new+8(FP), R9
	byteSetup
	SLLD	R1, R9, R9		// new byte in position
xchg8_again:
	MOVUW	(R8), R10
	ANDN	R10, R2, R11		// clear the target byte
	OR	R9, R11, R11		// insert the new byte
	MOVD	R11, R12
	CASW	(R8), R10, R12
	SUBCC	R12, R10, ZR
	BNED	xchg8_again
	MEMBAR_FULL
	AND	R2, R10, R10		// isolate the old byte
	SRLD	R1, R10, R10		// and shift it down
	MOVB	R10, ret+16(FP)
	RET

// Aliases. These differ only in the Go signature.
TEXT ·Casint32(SB), NOSPLIT, $0-17
	JMP	·Cas(SB)

TEXT ·Casint64(SB), NOSPLIT, $0-25
	JMP	·Cas64(SB)

TEXT ·Casuintptr(SB), NOSPLIT, $0-25
	JMP	·Cas64(SB)

TEXT ·Casp1(SB), NOSPLIT, $0-25
	JMP	·Cas64(SB)

TEXT ·CasRel(SB), NOSPLIT, $0-17
	JMP	·Cas(SB)

TEXT ·Loaduintptr(SB), NOSPLIT, $0-16
	JMP	·Load64(SB)

TEXT ·Loaduint(SB), NOSPLIT, $0-16
	JMP	·Load64(SB)

TEXT ·Loadint32(SB), NOSPLIT, $0-12
	JMP	·Load(SB)

TEXT ·Loadint64(SB), NOSPLIT, $0-16
	JMP	·Load64(SB)

TEXT ·LoadAcq(SB), NOSPLIT, $0-12
	JMP	·Load(SB)

TEXT ·LoadAcq64(SB), NOSPLIT, $0-16
	JMP	·Load64(SB)

TEXT ·LoadAcquintptr(SB), NOSPLIT, $0-16
	JMP	·Load64(SB)

TEXT ·Storeint32(SB), NOSPLIT, $0-12
	JMP	·Store(SB)

TEXT ·Storeint64(SB), NOSPLIT, $0-16
	JMP	·Store64(SB)

TEXT ·Storeuintptr(SB), NOSPLIT, $0-16
	JMP	·Store64(SB)

TEXT ·StoreReluintptr(SB), NOSPLIT, $0-16
	JMP	·StoreRel64(SB)

TEXT ·Xaddint32(SB), NOSPLIT, $0-20
	JMP	·Xadd(SB)

TEXT ·Xaddint64(SB), NOSPLIT, $0-24
	JMP	·Xadd64(SB)

TEXT ·Xadduintptr(SB), NOSPLIT, $0-24
	JMP	·Xadd64(SB)

TEXT ·Xchgint32(SB), NOSPLIT, $0-20
	JMP	·Xchg(SB)

TEXT ·Xchgint64(SB), NOSPLIT, $0-24
	JMP	·Xchg64(SB)

TEXT ·Xchguintptr(SB), NOSPLIT, $0-24
	JMP	·Xchg64(SB)

TEXT ·And(SB), NOSPLIT, $0-12
	JMP	·And32(SB)

TEXT ·Or(SB), NOSPLIT, $0-12
	JMP	·Or32(SB)

TEXT ·Anduintptr(SB), NOSPLIT, $0-24
	JMP	·And64(SB)

TEXT ·Oruintptr(SB), NOSPLIT, $0-24
	JMP	·Or64(SB)
