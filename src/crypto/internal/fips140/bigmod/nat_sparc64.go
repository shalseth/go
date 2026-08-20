// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !purego && sparc64

package bigmod

// Implemented in nat_sparc64.s, using VIS3's UMULXHI for the high half
// of the product and ADDXC for the carries.

//go:noescape
func addMulVVW1024(z, x *uint, y uint) (c uint)

//go:noescape
func addMulVVW1536(z, x *uint, y uint) (c uint)

//go:noescape
func addMulVVW2048(z, x *uint, y uint) (c uint)
