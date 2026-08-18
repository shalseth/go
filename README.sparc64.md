# Go on linux/sparc64

An out-of-tree port of the Go gc toolchain to 64-bit SPARC (SPARC V9),
developed on an UltraSPARC T4-1 running Gentoo. This branch adds
`GOARCH=sparc64` to the compiler, assembler, linker and runtime; the
`master` branch tracks upstream unchanged.

Status: the standard library test suite passes, and Grafana Alloy — an
OpenTelemetry collector with roughly two and a half thousand packages —
builds, runs, scrapes metrics and shuts down cleanly on the T4.

## Building

Cross-compile from any supported host (amd64 here):

```sh
cd src && GOROOT_BOOTSTRAP=/path/to/go1.24 ./make.bash
GOOS=linux GOARCH=sparc64 ../bin/go build ./yourprogram
```

The toolchain also builds and runs natively on sparc64. Cross-build a
bootstrap tree once, copy it to the target, and build from there:

```sh
cd src && GOOS=linux GOARCH=sparc64 ./bootstrap.bash
scp ../../go-linux-sparc64-bootstrap.tbz target:
ssh target 'tar xjf go-linux-sparc64-bootstrap.tbz'
ssh target 'GOROOT_BOOTSTRAP=$HOME/go-linux-sparc64-bootstrap go/src/make.bash'
```

The gc compiler is one of the heaviest Go workloads there is, so a
native `go build` doubles as a stress test of the port; several of the
subtlest runtime bugs on this branch were only ever reproducible that
way.

## What works

* The compiler, assembler and linker, including the SSA backend, the
  big-endian and 64-bit-address paths, and the delay-slot scheduler.
* Goroutines, channels, timers, `select`, defer/panic/recover, `Goexit`.
* Stack growth and copying, cooperative *and* asynchronous preemption.
* The garbage collector, including the Green Tea span marker
  (`GOEXPERIMENT=greenteagc`, on by default), verified with
  `GODEBUG=gccheckmark=1`.
* Signals: `os/signal`, `signal.NotifyContext`, SIGPIPE semantics,
  fork/exec, `net`, `net/http`, `os`, `syscall`.

## What is missing

* cgo and external linking. Everything must be built with
  `CGO_ENABLED=0`.
* VDSO support: `time.Now` and friends go through real syscalls.
* The race detector, and the `-buildmode` variants beyond `exe`.

## Notes on the ABI

Go does not use SPARC register windows: `SAVE`/`RESTORE` appear nowhere
in generated code, and one window serves the whole program. Frames are
flat, and two "in" registers act as per-frame anchors:

* `%i6` (RFP) — the frame's entry stack pointer, biased.
* `%i7` (OLR) — the frame's own return address.

A function's prologue stores its caller's pair at `[sp+112]` and
`[sp+120]` — the same offsets the hardware uses for `%i6`/`%i7` in a
window save area — so the values the kernel spills there on a trap
always agree with the ones the unwinder reads. Several of the subtler
bugs fixed on this branch come from moments where that agreement is
briefly untrue: the prologue between pushing the frame and publishing
the return address, the epilogue between reloading the caller's anchors
and raising the stack pointer, and any signal handler that opens a
second register window. Those regions are now either ordered so the
invariant holds or marked non-preemptible.

Other things worth knowing before touching this code:

* The stack pointer carries the SPARC V9 bias of 2047; real addresses
  are `%sp + 2047`, and every frame reserves 176 bytes (a 128-byte
  window save area plus the ABI's argument area).
* SPARC keeps SunOS signal numbers (`SIGSTOP` 17, `SIGCHLD` 20,
  `SIGUSR1` 30) and SunOS `sigprocmask` values (`SIG_BLOCK` 1,
  `SIG_UNBLOCK` 2, `SIG_SETMASK` 4).
* The kernel `sigset_t` is a single 64-bit word, so the generic
  `[2]uint32` form puts signals 1-32 in the wrong half on a big-endian
  machine.
* `O_CLOEXEC`, `EPOLL_CLOEXEC` and `EFD_CLOEXEC` are `0x400000`;
  `EFD_NONBLOCK` is `0x4000`.
* Userspace addresses are 52-bit and sign-extended, so unhinted `mmap`
  returns high-half addresses such as `0xfff8...`.
* `rt_sigaction` takes the restorer as its fourth argument, and the
  handler returns to `restorer+8`.

## Testing

Standard library test binaries are cross-compiled and run on the
target:

```sh
GOOS=linux GOARCH=sparc64 go test -c -o sort.test sort
scp sort.test target: && ssh target ./sort.test -test.short
```

Tests that read files from their own source directory (`os`, `net`,
`io/fs`, `text/template`) need that directory copied alongside the
binary.
