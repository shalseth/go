// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package arch

import "testing"

// TestSPARC64Registers checks that the register namespace is bound as
// the assembler expects. obj.Rconv prints the role name for registers
// that have one (R0 is ZR, F0 is FTMP), so archSPARC64 must bind the
// numeric spellings explicitly; this test guards that.
func TestSPARC64Registers(t *testing.T) {
	a := Set("sparc64", false)
	if a == nil {
		t.Fatal("Set(sparc64) returned nil")
	}
	want := []string{
		"R0", "R1", "R31", "F0", "F31", "Y0", "Y15",
		"D0", "D2", "D32", "D62", "FCC0", "FCC3",
		"ZR", "TLS", "RSP", "LR", "RFP", "OLR", "TMP", "TMP2",
		"RT1", "RT2", "CTXT", "FTMP", "DTMP", "BSP", "BFP",
		"ICC", "XCC", "CCR", "TICK", "RPC", "g", "SB", "FP", "PC",
	}
	for _, n := range want {
		if _, ok := a.Register[n]; !ok {
			t.Errorf("register %q not bound", n)
		}
	}
	// R22 is g; it must not be reachable by its raw number, so that
	// assembly cannot clobber g by accident.
	if _, ok := a.Register["R22"]; ok {
		t.Error("R22 is bound; it should be reachable only as g")
	}
	for _, n := range []string{"ADD", "MOVD", "FADDD", "CASD", "MEMBAR", "SETHI", "RET"} {
		if _, ok := a.Instructions[n]; !ok {
			t.Errorf("instruction %q not bound", n)
		}
	}
}

// TestSPARC64RegisterNumber checks the R(n)/F(n)/D(n)/Y(n) forms,
// including the interleaved upper half of the double-precision file.
func TestSPARC64RegisterNumber(t *testing.T) {
	for _, tc := range []struct {
		prefix string
		n      int16
		ok     bool
	}{
		{"R", 0, true}, {"R", 31, true}, {"R", 32, false},
		{"F", 31, true}, {"F", 32, false},
		{"D", 0, true}, {"D", 30, true}, {"D", 32, true}, {"D", 62, true},
		{"D", 1, false},  // odd double registers do not exist
		{"D", 64, false}, // out of range
		{"Y", 15, true}, {"Y", 16, false},
		{"Q", 0, false}, // unknown prefix
	} {
		if _, ok := sparc64RegisterNumber(tc.prefix, tc.n); ok != tc.ok {
			t.Errorf("sparc64RegisterNumber(%q, %d) ok=%v, want %v", tc.prefix, tc.n, ok, tc.ok)
		}
	}
}
