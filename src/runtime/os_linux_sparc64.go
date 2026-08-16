// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// stackBias is the constant SPARC V9 subtracts from %sp and %fp in the
// 64-bit ABI: the real stack top is %sp+stackBias. It must agree with
// sparc64.StackBias in cmd/internal/obj.
const stackBias = 2047

func osArchInit() {}
