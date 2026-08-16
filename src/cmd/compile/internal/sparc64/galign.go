// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sparc64

import (
	"cmd/compile/internal/ssagen"
	"cmd/internal/obj/sparc64"
)

func Init(arch *ssagen.ArchInfo) {
	arch.LinkArch = &sparc64.Linksparc64
	arch.REGSP = sparc64.REGSP
	arch.MAXWIDTH = 1 << 50

	arch.ZeroRange = zerorange
	arch.Ginsnop = ginsnop

	arch.SSAMarkMoves = ssaMarkMoves
	arch.SSAGenValue = ssaGenValue
	arch.SSAGenBlock = ssaGenBlock
}
