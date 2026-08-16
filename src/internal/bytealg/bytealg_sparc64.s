// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// linux/sparc64 uses the generic Go implementations in this package, so
// there is no assembly to provide. The file still has to exist: without
// at least one assembly file for the target, cmd/go compiles the package
// with -complete, and MakeNoZero, which the runtime supplies via
// linkname, is rejected as a function with no body.
