// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http_test

import (
	"io"
	"log"
	"net/http"
	"net/http/internal/http2"
	"os"
	"strconv"
	"testing"
	"time"
)

var quietLog = log.New(io.Discard, "", 0)

func TestMain(m *testing.M) {
	*http.MaxWriteWaitBeforeConnReuse = 60 * time.Minute
	// The HTTP/2 server gives a client 2 seconds to send its SETTINGS
	// frame. Scale that like the other deadlines when the machine is
	// slow or loaded.
	http2.SetFirstSettingsTimeout(2 * time.Second * time.Duration(timeoutScale()))
	os.Exit(m.Run())
}

// waitCondition waits for fn to return true,
// checking immediately and then at exponentially increasing intervals.
func waitCondition(t testing.TB, delay time.Duration, fn func(time.Duration) bool) {
	t.Helper()
	start := time.Now()
	var since time.Duration
	for !fn(since) {
		time.Sleep(delay)
		delay = 2*delay - (delay / 2) // 1.5x, rounded up
		since = time.Since(start)
	}
}

// timeoutScale returns GO_TEST_TIMEOUT_SCALE, or 1.
func timeoutScale() int {
	if s := os.Getenv("GO_TEST_TIMEOUT_SCALE"); s != "" {
		if n, err := strconv.Atoi(s); err == nil && n > 0 {
			return n
		}
	}
	return 1
}
