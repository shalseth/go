// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !purego

package gcm

import (
	"crypto/internal/fips140/aes"
	"crypto/internal/fips140/subtle"
	"crypto/internal/fips140deps/cpu"
	"crypto/internal/impl"
	"unsafe"
)

// The following functions are defined in gcm_sparc64.s. Their pointers
// are 8-aligned: the data is loaded with ldx and ldd.

//go:noescape
func gcmGhashBlocks(h, y *[2]uint64, p *byte, n int)

//go:noescape
func gcmCtrBlocks(nr int, xk *uint32, dst, src *byte, n int, ctr *[2]uint64)

// XMULX is VIS3, which the port already requires.
var supportsAESGCM = cpu.SPARC64HasAES

func init() {
	impl.Register("gcm", "T4", &supportsAESGCM)
}

// checkGenericIsExpected is called by the variable-time implementation to make
// sure it is not used when hardware support is available.
func checkGenericIsExpected() {
	if supportsAESGCM {
		panic("gcm: internal error: using generic implementation despite hardware support")
	}
}

type gcmPlatformData struct {
	h [2]uint64 // the hash key, as big-endian halves
}

func initGCM(g *GCM) {
	if !supportsAESGCM {
		return
	}
	aes.EncryptBlockInternal(&g.cipher, words(&g.h), words(&g.h))
}

// chunk is the size of the aligned buffer that misaligned data goes
// through.
const chunk = 256

func ghashUpdate(g *GCM, y *[2]uint64, data []byte) {
	if n := len(data) &^ (gcmBlockSize - 1); n > 0 {
		if aligned(data) {
			gcmGhashBlocks(&g.h, y, &data[0], n/gcmBlockSize)
		} else {
			var buf [chunk / 8]uint64
			b := unsafe.Slice((*byte)(unsafe.Pointer(&buf[0])), chunk)
			for p := data[:n]; len(p) > 0; {
				m := copy(b, p)
				gcmGhashBlocks(&g.h, y, &b[0], m/gcmBlockSize)
				p = p[m:]
			}
		}
		data = data[n:]
	}
	if len(data) > 0 {
		var last [2]uint64
		copy(words(&last), data)
		gcmGhashBlocks(&g.h, y, &words(&last)[0], 1)
	}
}

func counterCrypt(g *GCM, out, src []byte, ctr *[2]uint64) {
	ks := aes.EncryptionKeySchedule(&g.cipher)
	nr := len(ks)/4 - 1
	if n := len(src) &^ (gcmBlockSize - 1); n > 0 {
		if aligned(out) && aligned(src) {
			gcmCtrBlocks(nr, &ks[0], &out[0], &src[0], n/gcmBlockSize, ctr)
		} else {
			var buf [chunk / 8]uint64
			b := unsafe.Slice((*byte)(unsafe.Pointer(&buf[0])), chunk)
			for o, p := out[:n], src[:n]; len(p) > 0; {
				m := copy(b, p)
				gcmCtrBlocks(nr, &ks[0], &b[0], &b[0], m/gcmBlockSize, ctr)
				copy(o, b[:m])
				o, p = o[m:], p[m:]
			}
		}
		out, src = out[n:], src[n:]
	}
	if len(src) > 0 {
		var last [2]uint64
		b := words(&last)
		copy(b, src)
		gcmCtrBlocks(nr, &ks[0], &b[0], &b[0], 1, ctr)
		copy(out, b)
	}
}

// start derives the initial counter from the nonce, returns the tag mask
// E(counter), and leaves ctr at the first block for the data.
func start(g *GCM, nonce []byte, ctr *[2]uint64) (tagMask [2]uint64) {
	if len(nonce) == gcmStandardNonceSize {
		copy(words(ctr), nonce)
		ctr[1] = ctr[1]&^0xffffffff | 1
	} else {
		ghashUpdate(g, ctr, nonce)
		lens := [2]uint64{0, uint64(len(nonce)) * 8}
		gcmGhashBlocks(&g.h, ctr, &words(&lens)[0], 1)
	}
	aes.EncryptBlockInternal(&g.cipher, words(&tagMask), words(ctr))
	ctr[1] = ctr[1]&^0xffffffff | uint64(uint32(ctr[1])+1)
	return tagMask
}

func finish(g *GCM, y, tagMask *[2]uint64, dataLen, textLen int) {
	lens := [2]uint64{uint64(dataLen) * 8, uint64(textLen) * 8}
	gcmGhashBlocks(&g.h, y, &words(&lens)[0], 1)
	y[0] ^= tagMask[0]
	y[1] ^= tagMask[1]
}

func seal(out []byte, g *GCM, nonce, plaintext, data []byte) {
	if !supportsAESGCM {
		sealGeneric(out, g, nonce, plaintext, data)
		return
	}
	var ctr, y [2]uint64
	tagMask := start(g, nonce, &ctr)
	ghashUpdate(g, &y, data)
	counterCrypt(g, out[:len(plaintext)], plaintext, &ctr)
	ghashUpdate(g, &y, out[:len(plaintext)])
	finish(g, &y, &tagMask, len(data), len(plaintext))
	copy(out[len(plaintext):], words(&y))
}

func open(out []byte, g *GCM, nonce, ciphertext, data []byte) error {
	if !supportsAESGCM {
		return openGeneric(out, g, nonce, ciphertext, data)
	}
	tag := ciphertext[len(ciphertext)-g.tagSize:]
	ciphertext = ciphertext[:len(ciphertext)-g.tagSize]

	var ctr, y [2]uint64
	tagMask := start(g, nonce, &ctr)
	ghashUpdate(g, &y, data)
	ghashUpdate(g, &y, ciphertext)
	finish(g, &y, &tagMask, len(data), len(ciphertext))
	if subtle.ConstantTimeCompare(words(&y)[:g.tagSize], tag) != 1 {
		return errOpen
	}
	counterCrypt(g, out, ciphertext, &ctr)
	return nil
}

// words returns the bytes of w, which on this big-endian machine are the
// two halves in order.
func words(w *[2]uint64) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(&w[0])), 16)
}

func aligned(b []byte) bool {
	return uintptr(unsafe.Pointer(unsafe.SliceData(b)))&7 == 0
}
