// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !sparc64

package runtime

import "unsafe"

// unwindTrail is the sparc64 frame trail; see traceback_trail_sparc64.go.
// Empty here, so an unwinder costs exactly what it did before the port.
type unwindTrail struct{}

// Compile-time proof that it is free: an array of this length only exists
// if the length is zero.
var _ [0]byte = [unsafe.Sizeof(unwindTrail{})]byte{}

func (u *unwinder) record()              {}
func (u *unwinder) dumpTrail(why string) {}
