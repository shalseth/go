// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// func castagnoliUpdate(crc uint32, p *byte, n int) uint32
TEXT ·castagnoliUpdate(SB), NOSPLIT|NOFRAME, $0-28
	MOVUW	crc+0(FP), R8
	MOVD	p+8(FP), R9
	MOVD	n+16(FP), R10
	MOVXTOD	R8, D2

loop:
	BRZ	R10, done
	FMOVD	(R9), D4
	CRC32C	D4, D2, D2
	ADD	$8, R9
	SUB	$1, R10
	JMP	loop

done:
	MOVDTOX	D2, R8
	MOVW	R8, ret+24(FP)
	RET
