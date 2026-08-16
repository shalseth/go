// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build amd64 || arm64 || loong64 || mips64 || mips64le || ppc64 || ppc64le || riscv64 || s390x || sparc64 || wasm

package runtime

import (
	"internal/goarch"
	"internal/goos"
	"unsafe"
)

const (
	// addrBits is the number of bits needed to represent a virtual address.
	//
	// See heapAddrBits for a table of address space sizes on
	// various architectures. 48 bits is enough for all
	// arch/os combos except s390x, aix, and riscv64.
	//
	// On AMD64, virtual addresses are 48-bit (or 57-bit) sign-extended.
	// Other archs are 48-bit zero-extended.
	//
	// We use one extra bit to placate systems which simulate amd64 binaries on
	// an arm64 host. Allocated arm64 addresses could be as high as 1<<48-1,
	// which would be invalid if we assumed 48-bit sign-extended addresses.
	// See issue 69255.
	// (Note that this does not help the other way around, simluating arm64
	// on amd64, but we don't have that problem at the moment.)
	//
	// On s390x, virtual addresses are 64-bit. There's not much we
	// can do about this, so we just hope that the kernel doesn't
	// get to really high addresses and panic if it does.
	defaultAddrBits = 48 + 1

	// On AIX, 64-bit addresses are split into 36-bit segment number and 28-bit
	// offset in segment.  Segment numbers in the range 0x0A0000000-0x0AFFFFFFF(LSA)
	// are available for mmap.
	// We assume all tagged addresses are from memory allocated with mmap.
	// We use one bit to distinguish between the two ranges.
	aixAddrBits = 57

	// Later versions of FreeBSD enable amd64's la57 by default.
	freebsdAmd64AddrBits = 57

	// riscv64 SV57 mode gives 56 bits of userspace VA.
	// tagged pointer code supports it,
	// but broader support for SV57 mode is incomplete,
	// and there may be other issues (see #54104).
	riscv64AddrBits = 56

	// sparc64 (sun4v) has a 52-bit sign-extended VA space: userspace
	// addresses live in [0, 2^51) and [2^64-2^51, 2^64), and the
	// kernel serves unhinted mmaps from the high half.
	sparc64AddrBits = 52

	addrBits = goos.IsAix*aixAddrBits + goarch.IsRiscv64*riscv64AddrBits + goarch.IsSparc64*sparc64AddrBits + goos.IsFreebsd*goarch.IsAmd64*freebsdAmd64AddrBits + (1-goos.IsAix)*(1-goarch.IsRiscv64)*(1-goarch.IsSparc64)*(1-goos.IsFreebsd*goarch.IsAmd64)*defaultAddrBits

	// In addition to the 16 bits (or other, depending on arch/os) taken from the top,
	// we can take 9 from the bottom, because we require pointers to be well-aligned
	// (see tagptr.go:tagAlignBits). That gives us a total of 25 bits for the tag.
	tagBits = 64 - addrBits + tagAlignBits
)

// taggedPointerPack created a taggedPointer from a pointer and a tag.
// Tag bits that don't fit in the result are discarded.
func taggedPointerPack(ptr unsafe.Pointer, tag uintptr) taggedPointer {
	t := taggedPointer(uint64(uintptr(ptr))<<(tagBits-tagAlignBits) | uint64(tag&(1<<tagBits-1)))
	if t.pointer() != ptr || t.tag() != tag {
		print("runtime: taggedPointerPack invalid packing: ptr=", ptr, " tag=", hex(tag), " packed=", hex(t), " -> ptr=", t.pointer(), " tag=", hex(t.tag()), "\n")
		throw("taggedPointerPack")
	}
	return t
}

// Pointer returns the pointer from a taggedPointer.
func (tp taggedPointer) pointer() unsafe.Pointer {
	if GOARCH == "amd64" || GOARCH == "sparc64" {
		// amd64 and sparc64 addresses are sign-extended: the upper half
		// of the VA space sits at the top of the address range, so sign
		// extend when unpacking.
		return unsafe.Pointer(uintptr(int64(tp) >> tagBits << tagAlignBits))
	}
	if GOOS == "aix" {
		return unsafe.Pointer(uintptr((tp >> tagBits << tagAlignBits) | 0xa<<56))
	}
	return unsafe.Pointer(uintptr(tp >> tagBits << tagAlignBits))
}

// Tag returns the tag from a taggedPointer.
func (tp taggedPointer) tag() uintptr {
	return uintptr(tp & (1<<tagBits - 1))
}
