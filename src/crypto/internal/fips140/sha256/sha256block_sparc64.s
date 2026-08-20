// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !purego

#include "textflag.h"

// func blockSHA256(dig *Digest, p []byte)
//
// The T4 SHA256 instruction hashes one 512-bit block per issue. It
// takes no operands: the eight state words come from %f0-%f7, the
// block from %f8-%f23, and the updated state is written back over
// %f0-%f7. So the state is loaded once, each block is loaded over the
// same eight double registers, and the state is stored once at the end.
//
// SHA-256 defines its block as big-endian 32-bit words, which is
// exactly how a load lands them here, so unlike the little-endian
// implementations this needs no byte swapping at all.
//
// dig.h is at offset 0 of Digest and the struct is 8-aligned, so the
// state loads are safe. p is 8-aligned by the caller.
TEXT ·blockSHA256(SB), NOSPLIT|NOFRAME, $0-32
	MOVD	dig+0(FP), R8
	MOVD	p_base+8(FP), R9
	MOVD	p_len+16(FP), R10

	FMOVD	(R8), D0
	FMOVD	8(R8), D2
	FMOVD	16(R8), D4
	FMOVD	24(R8), D6

loop:
	CMP	$64, R10
	BCSD	done
	FMOVD	(R9), D8
	FMOVD	8(R9), D10
	FMOVD	16(R9), D12
	FMOVD	24(R9), D14
	FMOVD	32(R9), D16
	FMOVD	40(R9), D18
	FMOVD	48(R9), D20
	FMOVD	56(R9), D22
	SHA256
	ADD	$64, R9
	SUB	$64, R10
	JMP	loop

done:
	FMOVD	D0, (R8)
	FMOVD	D2, 8(R8)
	FMOVD	D4, 16(R8)
	FMOVD	D6, 24(R8)
	RET
