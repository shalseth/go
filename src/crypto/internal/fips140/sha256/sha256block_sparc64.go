// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !purego

package sha256

import (
	"crypto/internal/fips140deps/cpu"
	"crypto/internal/impl"
	"unsafe"
)

var useSHA256 = cpu.SPARC64HasCrypto

func init() {
	impl.Register("sha256", "T4", &useSHA256)
}

//go:noescape
func blockSHA256(dig *Digest, p []byte)

func block(dig *Digest, p []byte) {
	if !useSHA256 {
		blockGeneric(dig, p)
		return
	}
	// The instruction reads its block from the floating-point registers,
	// which are loaded with ldd, and ldd traps on an unaligned address.
	// Misaligned input therefore goes through an aligned copy, a block at
	// a time. A [chunk]byte local would not do: byte arrays carry no
	// alignment, so the buffer is declared as words.
	if len(p) >= chunk && uintptr(unsafe.Pointer(&p[0]))&7 != 0 {
		var words [chunk / 8]uint64
		buf := unsafe.Slice((*byte)(unsafe.Pointer(&words[0])), chunk)
		for len(p) >= chunk {
			copy(buf, p[:chunk])
			blockSHA256(dig, buf)
			p = p[chunk:]
		}
		return
	}
	blockSHA256(dig, p)
}
