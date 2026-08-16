// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"
#include "funcdata.h"

// linux/sparc64 passes arguments on the stack, so there is no
// abi.RegArgs to spill and these stubs follow the same shape as the
// other stack-ABI ports. Outgoing arguments start at FIXED_FRAME, the
// 176-byte SPARC V9 minimum frame, and BSP applies the stack bias.
#define FIXED_FRAME 176

// makeFuncStub is the code half of the function returned by MakeFunc.
// See the comment on the declaration of makeFuncStub in makefunc.go.
// No arg size here: the runtime pulls the arg map out of the func value.
TEXT ·makeFuncStub(SB),(NOSPLIT|WRAPPER),$48
	NO_LOCAL_POINTERS
	MOVD	CTXT, (FIXED_FRAME+0)(BSP)
	MOVD	$argframe+0(FP), R1
	MOVD	R1, (FIXED_FRAME+8)(BSP)
	MOVB	ZR, (FIXED_FRAME+32)(BSP)
	MOVD	$(FIXED_FRAME+32)(BSP), R1
	MOVD	R1, (FIXED_FRAME+16)(BSP)
	MOVD	ZR, (FIXED_FRAME+24)(BSP)
	CALL	·callReflect(SB)
	RET

// methodValueCall is the code half of the function returned by
// makeMethodValue. See the comment on its declaration in makefunc.go.
// No arg size here; the runtime pulls the arg map out of the func value.
TEXT ·methodValueCall(SB),(NOSPLIT|WRAPPER),$48
	NO_LOCAL_POINTERS
	MOVD	CTXT, (FIXED_FRAME+0)(BSP)
	MOVD	$argframe+0(FP), R1
	MOVD	R1, (FIXED_FRAME+8)(BSP)
	MOVB	ZR, (FIXED_FRAME+32)(BSP)
	MOVD	$(FIXED_FRAME+32)(BSP), R1
	MOVD	R1, (FIXED_FRAME+16)(BSP)
	MOVD	ZR, (FIXED_FRAME+24)(BSP)
	CALL	·callMethod(SB)
	RET
