// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && gc

package issue9400

var Baton int32

func RewindAndSetgid() {
	const pattern = 0x123456789abcdef
	var frame struct {
		pattern [1024]uint64
		// Leave room for a window spill and alignment above the pattern.
		spill [24]uint64
	}
	for i := range frame.pattern {
		frame.pattern[i] = pattern
	}
	rewindAndSetgid(&frame.spill)
	for _, v := range frame.pattern {
		if v != pattern {
			panic("SIGSETXID clobbered the stack")
		}
	}
}

//go:noescape
func rewindAndSetgid(spill *[24]uint64)
