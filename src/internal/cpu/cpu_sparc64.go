// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cpu

// CacheLinePadSize is used to prevent false sharing. The UltraSPARC T4
// reports no line size through getconf, and the 2016 Solaris port used
// 32; 64 is used here because over-padding is harmless while
// under-padding reintroduces the false sharing this exists to prevent.
const CacheLinePadSize = 64

// HWCap is set by the runtime from the AT_HWCAP entry of the auxiliary
// vector; see runtime/os_linux_sparc64.go.
var HWCap uint

// From the kernel's arch/sparc/include/uapi/asm/elf.h. A T4 reports
// 0x0747fbdf, which is every capability up to and including this one.
const hwcap_SPARC_CRYPTO = 0x04000000

// Crypto Function Register bits.
const (
	cfrAES    = 1 << 0
	cfrSHA1   = 1 << 5
	cfrSHA256 = 1 << 6
	cfrSHA512 = 1 << 7
	cfrCRC32C = 1 << 11
)

// readCFR reads the Crypto Function Register, %asr26. It is an illegal
// instruction without the crypto unit.
func readCFR() uint64

func doinit() {
	options = []option{
		{Name: "crypto", Feature: &SPARC64.HasCrypto},
		{Name: "aes", Feature: &SPARC64.HasAES},
		{Name: "sha1", Feature: &SPARC64.HasSHA1},
		{Name: "sha256", Feature: &SPARC64.HasSHA256},
		{Name: "sha512", Feature: &SPARC64.HasSHA512},
		{Name: "crc32c", Feature: &SPARC64.HasCRC32C},
	}
	SPARC64.HasCrypto = HWCap&hwcap_SPARC_CRYPTO != 0
	if SPARC64.HasCrypto {
		cfr := readCFR()
		SPARC64.HasAES = cfr&cfrAES != 0
		SPARC64.HasSHA1 = cfr&cfrSHA1 != 0
		SPARC64.HasSHA256 = cfr&cfrSHA256 != 0
		SPARC64.HasSHA512 = cfr&cfrSHA512 != 0
		SPARC64.HasCRC32C = cfr&cfrCRC32C != 0
	}
}
