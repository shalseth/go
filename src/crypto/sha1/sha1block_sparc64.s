// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !purego

#include "textflag.h"

// func blockSHA1(dig *digest, p []byte)
//
// The T4 SHA1 instruction hashes one 512-bit block per issue, with the
// five state words in %f0-%f4 and the block in %f8-%f23. It clears %f5,
// so the fifth word is loaded and stored on its own.
TEXT ·blockSHA1(SB), NOSPLIT|NOFRAME, $0-32
	MOVD	dig+0(FP), R8
	MOVD	p_base+8(FP), R9
	MOVD	p_len+16(FP), R10

	FMOVD	(R8), D0
	FMOVD	8(R8), D2
	FMOVS	16(R8), F4

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
	SHA1
	ADD	$64, R9
	SUB	$64, R10
	JMP	loop

done:
	FMOVD	D0, (R8)
	FMOVD	D2, 8(R8)
	FMOVS	F4, 16(R8)
	RET
