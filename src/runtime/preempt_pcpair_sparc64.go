// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build sparc64

package runtime

// sigctxtSequentialPC reports whether the interrupted context resumes
// at a single sequential PC. A signal delivered in a branch delay slot
// has npc != pc+4 (npc is the branch target), and rewriting the PC
// pair to inject a call would lose that target, so such contexts must
// not be preempted.
func sigctxtSequentialPC(c *sigctxt) bool {
	return c.npc() == c.pc()+4
}
