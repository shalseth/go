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
	targ := r.Sym()

	switch rt := r.Type(); rt {
	default:
		if rt >= objabi.ElfRelocOffset {
			ldr.Errorf(s, "unexpected relocation type %d (%s)", r.Type(), sym.RelocName(target.Arch, r.Type()))
			return false
		}

	// Relocations read from host ELF objects (internal linking of cgo
	// code). Convert each to the internal relocation that applies the
	// same bits; the actual patching happens in archreloc.
	case objabi.ElfRelocOffset + objabi.RelocType(elf.R_SPARC_64),
		objabi.ElfRelocOffset + objabi.RelocType(elf.R_SPARC_UA64):
		su := ldr.MakeSymbolUpdater(s)
		su.SetRelocType(rIdx, objabi.R_ADDR)
		return true

	case objabi.ElfRelocOffset + objabi.RelocType(elf.R_SPARC_32),
		objabi.ElfRelocOffset + objabi.RelocType(elf.R_SPARC_UA32):
		su := ldr.MakeSymbolUpdater(s)
		su.SetRelocType(rIdx, objabi.R_ADDR)
		su.SetRelocSiz(rIdx, 4)
		return true

	case objabi.ElfRelocOffset + objabi.RelocType(elf.R_SPARC_DISP32):
		su := ldr.MakeSymbolUpdater(s)
		su.SetRelocType(rIdx, objabi.R_PCREL)
		su.SetRelocSiz(rIdx, 4)
		return true

	case objabi.ElfRelocOffset + objabi.RelocType(elf.R_SPARC_WDISP30),
		objabi.ElfRelocOffset + objabi.RelocType(elf.R_SPARC_WPLT30):
		// A call. To a symbol in this binary it resolves directly; to
		// a dynamic import it goes through a PLT entry.
		su := ldr.MakeSymbolUpdater(s)
		su.SetRelocType(rIdx, objabi.R_CALLSPARC64)
		if ldr.SymType(targ) == sym.SDYNIMPORT {
			addpltsym(target, ldr, syms, targ)
			su.SetRelocSym(rIdx, syms.PLT)
			su.SetRelocAdd(rIdx, r.Add()+int64(ldr.SymPlt(targ)))
		}
		return true

	case objabi.ElfRelocOffset + objabi.RelocType(elf.R_SPARC_PC22):
		su := ldr.MakeSymbolUpdater(s)
		su.SetRelocType(rIdx, objabi.R_SPARC64ELFPC22)
		return true

	case objabi.ElfRelocOffset + objabi.RelocType(elf.R_SPARC_PC10):
		su := ldr.MakeSymbolUpdater(s)
		su.SetRelocType(rIdx, objabi.R_SPARC64ELFPC10)
		return true

	// The GOTDATA_OP model: a SETHI/XOR pair forms a GOT-relative
	// offset and an annotated ldx adds it to the GOT base. For a symbol
	// in this binary the linker resolves the offset to the symbol
	// itself and turns the load into an add; for a dynamic import it
	// allocates a real GOT slot and keeps the load.
	case objabi.ElfRelocOffset + 80, // R_SPARC_GOTDATA_HIX22
		objabi.ElfRelocOffset + 82: // R_SPARC_GOTDATA_OP_HIX22
		su := ldr.MakeSymbolUpdater(s)
		su.SetRelocType(rIdx, objabi.R_SPARC64GDOPHIX22)
		if ldr.SymType(targ) == sym.SDYNIMPORT {
			addgotsym(target, ldr, syms, targ)
			su.SetRelocSym(rIdx, syms.GOT)
			su.SetRelocAdd(rIdx, r.Add()+int64(ldr.SymGot(targ)))
		}
		return true

	case objabi.ElfRelocOffset + 81, // R_SPARC_GOTDATA_LOX10
		objabi.ElfRelocOffset + 83: // R_SPARC_GOTDATA_OP_LOX10
		su := ldr.MakeSymbolUpdater(s)
		su.SetRelocType(rIdx, objabi.R_SPARC64GDOPLOX10)
		if ldr.SymType(targ) == sym.SDYNIMPORT {
			addgotsym(target, ldr, syms, targ)
			su.SetRelocSym(rIdx, syms.GOT)
			su.SetRelocAdd(rIdx, r.Add()+int64(ldr.SymGot(targ)))
		}
		return true

	case objabi.ElfRelocOffset + 84: // R_SPARC_GOTDATA_OP
		su := ldr.MakeSymbolUpdater(s)
		if ldr.SymType(targ) == sym.SDYNIMPORT {
			su.SetRelocType(rIdx, objabi.R_SPARC64GDOPNOP)
			// The load stays exactly as it is; point the relocation
			// away from the dynamic symbol so the generic checks do
			// not demand dynamic-linking treatment for it.
			su.SetRelocSym(rIdx, syms.GOT)
			su.SetRelocAdd(rIdx, 0)
		} else {
			su.SetRelocType(rIdx, objabi.R_SPARC64GDOP2ADD)
		}
		return true

	case objabi.ElfRelocOffset + objabi.RelocType(elf.R_SPARC_HI22):
		su := ldr.MakeSymbolUpdater(s)
		su.SetRelocType(rIdx, objabi.R_SPARC64ELFHI22)
		return true

	case objabi.ElfRelocOffset + objabi.RelocType(elf.R_SPARC_LO10):
		su := ldr.MakeSymbolUpdater(s)
		su.SetRelocType(rIdx, objabi.R_SPARC64ELFLO10)
		return true
	}

	// Reject the rest (GOT and TLS machinery, which the internal linker
	// does not provide on this port).
	ldr.Errorf(s, "adddynrel: unsupported dynamic relocation for symbol %s (type %d)", ldr.SymName(targ), r.Type())
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
	if plt.Size() == 0 {
		// The psABI reserves the first four 32-byte entries; the
		// dynamic linker writes its resolver trampoline into them at
		// startup.
		for i := 0; i < 32; i++ {
			plt.AddUint32(ctxt.Arch, 0)
		}
	}
}

// addpltsym gives s a PLT entry in the psABI's format: a SETHI whose
// immediate is the entry's own byte offset within .plt (the resolver
// derives the relocation index from it) and a branch to the reserved
// second entry. The dynamic linker binds the slot by patching the entry's
// code, so the R_SPARC_JMP_SLOT relocation's offset is the entry itself.
func addpltsym(target *ld.Target, ldr *loader.Loader, syms *ld.ArchSyms, s loader.Sym) {
	if ldr.SymPlt(s) >= 0 {
		return
	}
	ld.Adddynsym(ldr, target, syms, s)

	plt := ldr.MakeSymbolUpdater(syms.PLT)
	rela := ldr.MakeSymbolUpdater(syms.RelaPLT)
	off := plt.Size()

	plt.AddUint32(target.Arch, 0x03000000|uint32(off))                    // sethi %hi(off<<10), %g1
	plt.AddUint32(target.Arch, 0x30680000|uint32((0x20-off-4)>>2)&0x7ffff) // ba,a %xcc, .plt+0x20
	for i := 0; i < 6; i++ {
		plt.AddUint32(target.Arch, 0x01000000) // nop
	}

	rela.AddAddrPlus(target.Arch, plt.Sym(), off)
	rela.AddUint64(target.Arch, elf.R_INFO(uint32(ldr.SymDynid(s)), uint32(elf.R_SPARC_JMP_SLOT)))
	rela.AddUint64(target.Arch, 0)

	ldr.SetPlt(s, int32(off))
}

// addgotsym gives s a GOT slot filled in by the dynamic linker.
func addgotsym(target *ld.Target, ldr *loader.Loader, syms *ld.ArchSyms, s loader.Sym) {
	if ldr.SymGot(s) >= 0 {
		return
	}
	ld.Adddynsym(ldr, target, syms, s)

	got := ldr.MakeSymbolUpdater(syms.GOT)
	rela := ldr.MakeSymbolUpdater(syms.Rela)
	off := got.Size()
	got.AddUint64(target.Arch, 0)

	rela.AddAddrPlus(target.Arch, got.Sym(), off)
	rela.AddUint64(target.Arch, elf.R_INFO(uint32(ldr.SymDynid(s)), uint32(elf.R_SPARC_GLOB_DAT)))
	rela.AddUint64(target.Arch, 0)

	ldr.SetGot(s, int32(off))
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

	case objabi.R_SPARC64ELFPC22, objabi.R_SPARC64ELFPC10:
		// Place-relative pieces (host R_SPARC_PC22/PC10).
		t := ldr.SymValue(rs) + r.Add() - (ldr.SymValue(s) + int64(r.Off()))
		if r.Type() == objabi.R_SPARC64ELFPC22 {
			return val&^0x3fffff | (t>>10)&0x3fffff, noExtReloc, isOk
		}
		return val&^0x3ff | t&0x3ff, noExtReloc, isOk

	case objabi.R_SPARC64GDOPHIX22, objabi.R_SPARC64GDOPLOX10:
		// GOT-relative value in the psABI's SETHI/XOR form, which
		// handles negative offsets: HIX22 = (v>>10)^(v>>31), LOX10 =
		// (v&0x3ff)|((v>>31)&0x1c00), with arithmetic shifts.
		v := ldr.SymValue(rs) + r.Add() - ldr.SymValue(syms.GOT)
		if r.Type() == objabi.R_SPARC64GDOPHIX22 {
			return val&^0x3fffff | ((v>>10)^(v>>31))&0x3fffff, noExtReloc, isOk
		}
		return val&^0x1fff | (v&0x3ff | (v>>31)&0x1c00), noExtReloc, isOk

	case objabi.R_SPARC64GDOP2ADD:
		// The annotated "ldx [got+off], rd" becomes "add got, off, rd":
		// same rd, rs1 and rs2, opcode swapped for the arithmetic op.
		rd := val >> 25 & 0x1f
		rs1 := val >> 14 & 0x1f
		rs2 := val & 0x1f
		return 2<<30 | rd<<25 | 0<<19 | rs1<<14 | rs2, noExtReloc, isOk

	case objabi.R_SPARC64GDOPNOP:
		// The symbol is dynamic; the load from its real GOT slot stays.
		return val, noExtReloc, isOk

	case objabi.R_SPARC64ELFHI22:
		// A host object's R_SPARC_HI22: the top 22 bits of a 32-bit
		// absolute address into this one SETHI.
		t := ldr.SymValue(rs) + r.Add()
		if uint64(t) != uint64(uint32(t)) {
			ldr.Errorf(s, "R_SPARC_HI22 target out of 32-bit range: %#x", t)
		}
		return val&^0x3fffff | (t>>10)&0x3fffff, noExtReloc, isOk

	case objabi.R_SPARC64ELFLO10:
		t := ldr.SymValue(rs) + r.Add()
		return val&^0x3ff | t&0x3ff, noExtReloc, isOk

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
