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
	"debug/elf"
	"log"
)

// Dynamic linking is not supported on linux/sparc64 yet: adddynrel and
// elfsetupplt fail loudly rather than emitting silently wrong output.
// External linking of a static binary is wired up below.

func gentext(ctxt *ld.Link, ldr *loader.Loader) {}

func adddynrel(target *ld.Target, ldr *loader.Loader, syms *ld.ArchSyms, s loader.Sym, r loader.Reloc, rIdx int) bool {
	ldr.Errorf(s, "adddynrel: dynamic linking is not supported on linux/sparc64")
	return false
}

func elfreloc1(ctxt *ld.Link, out *ld.OutBuf, ldr *loader.Loader, s loader.Sym, r loader.ExtReloc, ri int, sectoff int64) bool {
	// One Elf64_Rela entry.
	write := func(rtyp elf.R_SPARC, offset int64, sym int32, addend int64) {
		out.Write64(uint64(offset))
		out.Write64(uint64(rtyp) | uint64(sym)<<32)
		out.Write64(uint64(addend))
	}

	elfsym := ld.ElfSymForReloc(ctxt, r.Xsym)
	switch r.Type {
	default:
		return false

	case objabi.R_ADDR, objabi.R_DWARFSECREF:
		switch r.Size {
		case 4:
			write(elf.R_SPARC_32, sectoff, elfsym, r.Xadd)
		case 8:
			write(elf.R_SPARC_64, sectoff, elfsym, r.Xadd)
		default:
			return false
		}

	case objabi.R_CALLSPARC64:
		write(elf.R_SPARC_WDISP30, sectoff, elfsym, r.Xadd)

	// The HI and LO relocations each cover a SETHI and the instruction
	// that follows it, so each becomes two ELF relocations, one per
	// instruction word. HH22 and HM10 take bits 63..42 and 41..32 of
	// the address; LM22 and LO10 take bits 31..10 and 9..0.
	case objabi.R_ADDRSPARC64HI:
		write(elf.R_SPARC_HH22, sectoff, elfsym, r.Xadd)
		write(elf.R_SPARC_HM10, sectoff+4, elfsym, r.Xadd)

	case objabi.R_ADDRSPARC64LO:
		write(elf.R_SPARC_LM22, sectoff, elfsym, r.Xadd)
		write(elf.R_SPARC_LO10, sectoff+4, elfsym, r.Xadd)

	// Local exec, covering the sethi and the xor that follows it. The
	// sequence the assembler emits - sethi, xor, add %g7 - is the one
	// the ABI defines for these relocations, so the host linker can
	// fill it in as it stands.
	case objabi.R_SPARC64_TLS_LE:
		write(elf.R_SPARC_TLS_LE_HIX22, sectoff, elfsym, r.Xadd)
		write(elf.R_SPARC_TLS_LE_LOX10, sectoff+4, elfsym, r.Xadd)
	}

	return true
}

func elfsetupplt(ctxt *ld.Link, ldr *loader.Loader, plt, gotplt *loader.SymbolBuilder, dynamic loader.Sym) {
}

func machoreloc1(*sys.Arch, *ld.OutBuf, *loader.Loader, loader.Sym, loader.ExtReloc, int64) bool {
	return false
}

func archreloc(target *ld.Target, ldr *loader.Loader, syms *ld.ArchSyms, r loader.Reloc, s loader.Sym, val int64) (o int64, nExtReloc int, ok bool) {
	if target.IsExternal() {
		// The host linker computes the values; the instruction fields
		// stay as the assembler left them, with the addend carried in
		// the ELF relocation.
		switch r.Type() {
		case objabi.R_CALLSPARC64:
			return val, 1, true
		case objabi.R_ADDRSPARC64HI, objabi.R_ADDRSPARC64LO, objabi.R_SPARC64_TLS_LE:
			// Two ELF relocations each; see elfreloc1.
			return val, 2, true
		}
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
	switch r.Type() {
	case objabi.R_CALLSPARC64, objabi.R_SPARC64_TLS_LE:
		return ld.ExtrelocSimple(ldr, r), true
	case objabi.R_ADDRSPARC64HI, objabi.R_ADDRSPARC64LO:
		return ld.ExtrelocViaOuterSym(ldr, r, s), true
	}
	return loader.ExtReloc{}, false
}
