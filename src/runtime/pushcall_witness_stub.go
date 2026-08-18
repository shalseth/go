// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !(linux && sparc64)

package runtime

// dumpPushCallLog is a diagnostic for the linux/sparc64 port; see
// signal_sparc64.go. No-op elsewhere.
func dumpPushCallLog() {}
