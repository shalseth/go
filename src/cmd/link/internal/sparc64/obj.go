// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sparc64

import (
	"cmd/internal/objabi"
	"cmd/internal/sys"
	"cmd/link/internal/ld"
)

func Init() (*sys.Arch, ld.Arch) {
	arch := sys.ArchSPARC64

	theArch := ld.Arch{
		Funcalign:        funcAlign,
		Maxalign:         maxAlign,
		Minalign:         minAlign,
		Dwarfregsp:       dwarfRegSP,
		Dwarfreglr:       dwarfRegLR,
		Dwarfcfabias:     dwarfCFABias,
		Dwarfrainreg:     dwarfRAInReg,
		Dwarfraspill:     dwarfRASpill,
		Adddynrel:        adddynrel,
		Archinit:         archinit,
		Archreloc:        archreloc,
		Archrelocvariant: archrelocvariant,
		Extreloc:         extreloc,
		Gentext:          gentext,
		Machoreloc1:      machoreloc1,

		ELF: ld.ELFArch{
			Linuxdynld: "/lib64/ld-linux.so.2",
			// Only linux/sparc64 is supported; the other ELF systems
			// that once ran on SPARC are not Go targets.
			LinuxdynldMusl: "XXX",
			Freebsddynld:   "XXX",
			Openbsddynld:   "XXX",
			Netbsddynld:    "XXX",
			Dragonflydynld: "XXX",
			Solarisdynld:   "XXX",

			Reloc1:    elfreloc1,
			RelocSize: 24, // Elf64_Rela
			SetupPLT:  elfsetupplt,
		},
	}

	return arch, theArch
}

func archinit(ctxt *ld.Link) {
	switch ctxt.HeadType {
	default:
		ld.Exitf("unknown -H option: %v", ctxt.HeadType)

	case objabi.Hlinux: /* sparc64 elf */
		ld.Elfinit(ctxt)
		ld.HEADR = ld.ELFRESERVE
		if *ld.FlagRound == -1 {
			// SPARC64 uses an 8K base page; round segments to it.
			*ld.FlagRound = 0x2000
		}
		if *ld.FlagTextAddr == -1 {
			*ld.FlagTextAddr = ld.Rnd(0x100000, *ld.FlagRound) + int64(ld.HEADR)
		}
	}
}
