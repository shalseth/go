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

func doinit() {
	options = []option{
		{Name: "crypto", Feature: &SPARC64.HasCrypto},
	}
	SPARC64.HasCrypto = HWCap&hwcap_SPARC_CRYPTO != 0
}
