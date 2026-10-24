// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// SPARC64-specific hardware-assisted CRC32 algorithms. See crc32.go for a
// description of the interface that each architecture-specific file
// implements.

package crc32

import (
	"internal/cpu"
	"math/bits"
	"unsafe"
)

// castagnoliUpdate folds n 8-byte words at p into crc with the T4 CRC32C
// instruction. The CRC is not inverted and is held byte-swapped, as the
// instruction keeps it. p is 8-aligned.
//
//go:noescape
func castagnoliUpdate(crc uint32, p *byte, n int) uint32

var hasCRC32C = cpu.SPARC64.HasCrypto && cpu.SPARC64.HasCRC32C

func archAvailableCastagnoli() bool {
	return hasCRC32C
}

func archInitCastagnoli() {
	if !hasCRC32C {
		panic("arch-specific crc32 instruction for Castagnoli not available")
	}
}

func archUpdateCastagnoli(crc uint32, p []byte) uint32 {
	if !hasCRC32C {
		panic("arch-specific crc32 instruction for Castagnoli not available")
	}
	// The instruction takes aligned words; the bytes before the first
	// and after the last go through the table.
	if n := int(-uintptr(unsafe.Pointer(unsafe.SliceData(p))) & 7); n > 0 {
		n = min(n, len(p))
		crc = simpleUpdate(crc, castagnoliTable, p[:n])
		p = p[n:]
	}
	if n := len(p) &^ 7; n > 0 {
		crc = ^bits.ReverseBytes32(castagnoliUpdate(bits.ReverseBytes32(^crc), &p[0], n/8))
		p = p[n:]
	}
	if len(p) > 0 {
		crc = simpleUpdate(crc, castagnoliTable, p)
	}
	return crc
}

func archAvailableIEEE() bool                    { return false }
func archInitIEEE()                              { panic("not available") }
func archUpdateIEEE(crc uint32, p []byte) uint32 { panic("not available") }
