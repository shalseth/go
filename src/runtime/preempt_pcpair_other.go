// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !sparc64

package runtime

// sigctxtSequentialPC reports whether the interrupted context resumes
// at a single sequential PC; only SPARC's PC/nPC pairs can say no.
func sigctxtSequentialPC(c *sigctxt) bool {
	return true
}
