// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cpu

// CacheLinePadSize is used to prevent false sharing. The UltraSPARC T4
// reports no line size through getconf, and the 2016 Solaris port used
// 32; 64 is used here because over-padding is harmless while
// under-padding reintroduces the false sharing this exists to prevent.
const CacheLinePadSize = 64

func doinit() {}
