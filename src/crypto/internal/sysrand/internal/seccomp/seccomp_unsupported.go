// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !linux || !cgo

package seccomp

import "errors"

// ErrUnsupported reports that this system cannot install a seccomp filter.
var ErrUnsupported = errors.New("seccomp filters are not supported on this system")

func DisableGetrandom() error {
	return ErrUnsupported
}
