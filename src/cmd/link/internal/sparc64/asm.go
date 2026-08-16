// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sparc64

import (
	"cmd/internal/objabi"
	"cmd/internal/sys"
	"cmd/link/internal/ld"
	"cmd/link/internal/loader"
	"cmd/link/internal/sym"
	"log"
)

// External linking, dynamic linking and cgo are not supported on
// linux/sparc64 yet. The hooks below fail loudly rather than emitting
// silently wrong output; internal linking of a static binary is the
// only supported mode.

func gentext(ctxt *ld.Link, ldr *loader.Loader) {}

func adddynrel(target *ld.Target, ldr *loader.Loader, syms *ld.ArchSyms, s loader.Sym, r loader.Reloc, rIdx int) bool {
	ldr.Errorf(s, "adddynrel: dynamic linking is not supported on linux/sparc64")
	return false
}

func elfreloc1(ctxt *ld.Link, out *ld.OutBuf, ldr *loader.Loader, s loader.Sym, r loader.ExtReloc, ri int, sectoff int64) bool {
	// External linking would need the R_SPARC_* encodings here.
	return false
}

func elfsetupplt(ctxt *ld.Link, ldr *loader.Loader, plt, gotplt *loader.SymbolBuilder, dynamic loader.Sym) {
}

func machoreloc1(*sys.Arch, *ld.OutBuf, *loader.Loader, loader.Sym, loader.ExtReloc, int64) bool {
	return false
}

func archreloc(target *ld.Target, ldr *loader.Loader, syms *ld.ArchSyms, r loader.Reloc, s loader.Sym, val int64) (o int64, nExtReloc int, ok bool) {
	if target.IsExternal() {
		// Nothing is wired up for external linking yet; report the
		// relocation as unhandled so the linker complains rather than
		// writing a wrong value.
		return val, 0, false
	}

	const isOk = true
	const noExtReloc = 0
	rs := r.Sym()

	switch r.Type() {
	case objabi.R_CALLSPARC64:
		// SPARC V9 CALL carries a 30-bit word displacement, so the
		// reachable range is +-2GB rather than the +-8MB a branch gets.
		t := ldr.SymValue(rs) + r.Add() - (ldr.SymValue(s) + int64(r.Off()))
		if t > 1<<31-4 || t < -1<<31 {
			ldr.Errorf(s, "program too large, call relocation distance = %d", t)
		}
		if t&3 != 0 {
			ldr.Errorf(s, "unaligned call target, distance = %d", t)
		}
		return val | (t>>2)&0x3fffffff, noExtReloc, isOk

	case objabi.R_ADDRSPARC64HI, objabi.R_ADDRSPARC64LO:
		// These relocate a SETHI/OR pair, so the value spans two
		// instruction words: the high half of val is the SETHI and the
		// low half the instruction that follows it. HI takes bits
		// 63..32 of the target address, LO bits 31..0; within each
		// half, SETHI takes the top 22 bits and the OR the low 10.
		t := ldr.SymValue(rs) + r.Add()
		half := uint32(t)
		if r.Type() == objabi.R_ADDRSPARC64HI {
			half = uint32(uint64(t) >> 32)
		}
		o0 := uint32(val>>32) | half>>10
		o1 := uint32(val) | half&0x3ff
		return int64(o0)<<32 | int64(o1), noExtReloc, isOk

	case objabi.R_SPARC64_TLS_LE:
		// Local exec: the offset is relative to the thread pointer in
		// %g7, which the runtime sets up.
		t := ldr.SymValue(rs) + r.Add()
		if t < -4096 || t >= 4096 {
			ldr.Errorf(s, "TLS offset out of range %d", t)
		}
		return val | t&0x1fff, noExtReloc, isOk
	}

	return val, 0, false
}

func archrelocvariant(*ld.Target, *loader.Loader, loader.Reloc, sym.RelocVariant, loader.Sym, int64, []byte) int64 {
	log.Fatalf("unexpected relocation variant on sparc64")
	return -1
}

func extreloc(target *ld.Target, ldr *loader.Loader, r loader.Reloc, s loader.Sym) (loader.ExtReloc, bool) {
	return loader.ExtReloc{}, false
}
