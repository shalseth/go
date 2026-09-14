// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build sparc64

package runtime

// unwindTrail records the frames an unwind has resolved. The flat-frame ABI
// gives a mis-sized frame no other symptom than a walk that ends somewhere
// unexpected, several frames later, so the trail is what names the frame that
// was actually wrong.
type unwindTrail struct {
	pc    [8]uintptr
	sp    [8]uintptr
	fp    [8]uintptr
	delta [8]int32
	n     int
}

// record appends the current frame to the trail.
func (u *unwinder) record() {
	t := &u.trail
	i := t.n
	if i >= len(t.pc) {
		// keep the first frames and the most recent one
		i = len(t.pc) - 1
	}
	t.pc[i] = u.frame.pc
	t.sp[i] = u.frame.sp
	t.fp[i] = u.frame.fp
	if u.frame.fn.valid() {
		t.delta[i] = funcspdelta(u.frame.fn, u.frame.pc)
	} else {
		t.delta[i] = -1
	}
	t.n++
}

// dumpTrail prints the recorded frame chain. Each line is the frame as the
// unwinder resolved it; a frame whose fp does not equal the next frame's sp,
// or whose spdelta disagrees with fp-sp, is the one that mis-sized itself.
func (u *unwinder) dumpTrail(why string) {
	t := &u.trail
	print("unwind trail (", why, "), ", t.n, " frames, newest last:\n")
	n := t.n
	if n > len(t.pc) {
		n = len(t.pc)
	}
	for i := 0; i < n; i++ {
		print("  [", i, "] ")
		if fn := findfunc(t.pc[i]); fn.valid() {
			print(funcname(fn))
		} else {
			print("<invalid>")
		}
		print(" pc=", hex(t.pc[i]),
			" sp=", hex(t.sp[i]),
			" fp=", hex(t.fp[i]),
			" fp-sp=", t.fp[i]-t.sp[i],
			" spdelta=", t.delta[i], "\n")
	}
}
