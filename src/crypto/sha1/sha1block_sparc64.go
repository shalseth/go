// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !purego

package sha1

import (
	"crypto/internal/impl"
	"internal/cpu"
	"unsafe"
)

var useSHA1 = cpu.SPARC64.HasCrypto && cpu.SPARC64.HasSHA1

func init() {
	impl.Register("sha1", "T4", &useSHA1)
}

//go:noescape
func blockSHA1(dig *digest, p []byte)

func block(dig *digest, p []byte) {
	if !useSHA1 {
		blockGeneric(dig, p)
		return
	}
	// The block is loaded with ldd, which traps on an unaligned address,
	// so misaligned input goes through an aligned copy, a block at a time.
	if len(p) >= chunk && uintptr(unsafe.Pointer(&p[0]))&7 != 0 {
		var words [chunk / 8]uint64
		buf := unsafe.Slice((*byte)(unsafe.Pointer(&words[0])), chunk)
		for len(p) >= chunk {
			copy(buf, p[:chunk])
			blockSHA1(dig, buf)
			p = p[chunk:]
		}
		return
	}
	blockSHA1(dig, p)
}
