// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !purego

#include "textflag.h"

// func blockSHA512(dig *Digest, p []byte)
//
// The T4 SHA512 instruction hashes one 1024-bit block per issue, with the
// state in %d0-%d14 and the block in %d16-%d46. The state is loaded
// once and stored once; SHA-512's big-endian words need no byte swap.
TEXT ·blockSHA512(SB), NOSPLIT|NOFRAME, $0-32
	MOVD	dig+0(FP), R8
	MOVD	p_base+8(FP), R9
	MOVD	p_len+16(FP), R10

	FMOVD	(R8), D0
	FMOVD	8(R8), D2
	FMOVD	16(R8), D4
	FMOVD	24(R8), D6
	FMOVD	32(R8), D8
	FMOVD	40(R8), D10
	FMOVD	48(R8), D12
	FMOVD	56(R8), D14

loop:
	CMP	$128, R10
	BCSD	done
	FMOVD	(R9), D16
	FMOVD	8(R9), D18
	FMOVD	16(R9), D20
	FMOVD	24(R9), D22
	FMOVD	32(R9), D24
	FMOVD	40(R9), D26
	FMOVD	48(R9), D28
	FMOVD	56(R9), D30
	FMOVD	64(R9), D32
	FMOVD	72(R9), D34
	FMOVD	80(R9), D36
	FMOVD	88(R9), D38
	FMOVD	96(R9), D40
	FMOVD	104(R9), D42
	FMOVD	112(R9), D44
	FMOVD	120(R9), D46
	SHA512
	ADD	$128, R9
	SUB	$128, R10
	JMP	loop

done:
	FMOVD	D0, (R8)
	FMOVD	D2, 8(R8)
	FMOVD	D4, 16(R8)
	FMOVD	D6, 24(R8)
	FMOVD	D8, 32(R8)
	FMOVD	D10, 40(R8)
	FMOVD	D12, 48(R8)
	FMOVD	D14, 56(R8)
	RET
