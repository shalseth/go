// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !purego

package sha512

import (
	"crypto/internal/fips140deps/cpu"
	"crypto/internal/impl"
	"unsafe"
)

var useSHA512 = cpu.SPARC64HasSHA512

func init() {
	impl.Register("sha512", "T4", &useSHA512)
}

//go:noescape
func blockSHA512(dig *Digest, p []byte)

func block(dig *Digest, p []byte) {
	if !useSHA512 {
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
			blockSHA512(dig, buf)
			p = p[chunk:]
		}
		return
	}
	blockSHA512(dig, p)
}
