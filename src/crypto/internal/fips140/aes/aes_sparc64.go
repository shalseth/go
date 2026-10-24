// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !purego

package aes

import (
	"crypto/internal/fips140deps/cpu"
	"crypto/internal/impl"
	"unsafe"
)

//go:noescape
func encryptBlockAsm(nr int, xk *uint32, dst, src *byte)

//go:noescape
func decryptBlockAsm(nr int, xk *uint32, dst, src *byte)

//go:noescape
func expandKeyAsm(nr int, key *byte, enc *uint32)

var supportsAES = cpu.SPARC64HasAES

func init() {
	impl.Register("aes", "T4", &supportsAES)
}

// checkGenericIsExpected is called by the variable-time implementation to make
// sure it is not used when hardware support is available.
func checkGenericIsExpected() {
	if supportsAES {
		panic("crypto/aes: internal error: using generic implementation despite hardware support")
	}
}

type block struct {
	blockExpanded
}

func newBlock(c *Block, key []byte) *Block {
	switch len(key) {
	case aes128KeySize:
		c.rounds = aes128Rounds
	case aes192KeySize:
		c.rounds = aes192Rounds
	case aes256KeySize:
		c.rounds = aes256Rounds
	}
	if !supportsAES {
		expandKeyGeneric(&c.blockExpanded, key)
		return c
	}
	// The key is loaded into floating-point registers, which needs 8-byte
	// alignment, and the decryption rounds use the encryption schedule.
	var k [4]uint64
	kb := unsafe.Slice((*byte)(unsafe.Pointer(&k[0])), len(k)*8)
	copy(kb, key)
	expandKeyAsm(c.rounds, &kb[0], &c.enc[0])
	clear(k[:])
	return c
}

func encryptBlock(c *Block, dst, src []byte) {
	if !supportsAES {
		encryptBlockGeneric(&c.blockExpanded, dst, src)
		return
	}
	if aligned(dst, src) {
		encryptBlockAsm(c.rounds, &c.enc[0], &dst[0], &src[0])
		return
	}
	var b [2]uint64
	p := (*[BlockSize]byte)(unsafe.Pointer(&b[0]))
	copy(p[:], src)
	encryptBlockAsm(c.rounds, &c.enc[0], &p[0], &p[0])
	copy(dst, p[:])
}

func decryptBlock(c *Block, dst, src []byte) {
	if !supportsAES {
		decryptBlockGeneric(&c.blockExpanded, dst, src)
		return
	}
	if aligned(dst, src) {
		decryptBlockAsm(c.rounds, &c.enc[0], &dst[0], &src[0])
		return
	}
	var b [2]uint64
	p := (*[BlockSize]byte)(unsafe.Pointer(&b[0]))
	copy(p[:], src)
	decryptBlockAsm(c.rounds, &c.enc[0], &p[0], &p[0])
	copy(dst, p[:])
}

// aligned reports whether dst and src are both 8-aligned, as the loads and
// stores into floating-point registers need.
func aligned(dst, src []byte) bool {
	return (uintptr(unsafe.Pointer(&dst[0]))|uintptr(unsafe.Pointer(&src[0])))&7 == 0
}

// EncryptionKeySchedule is used from the GCM implementation to access the
// precomputed AES key schedule, to pass to the assembly implementation.
func EncryptionKeySchedule(c *Block) []uint32 {
	return c.enc[:c.roundKeysSize()]
}
