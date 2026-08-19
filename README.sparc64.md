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
* Atomic intrinsics. The atomics are correct - hand-written assembly in
  `internal/runtime/atomic` - but the compiler does not inline them, so
  every atomic operation is an out-of-line call. That also blocks
  inlining of the `sync` fast paths, which is why `test/inline_sync.go`
  excludes this architecture alongside 386, arm and wasm. Fixing it
  means adding the SSA ops: a CASD/CASW retry loop with the right
  membar placement for the TSO model.
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

### Return addresses

The single fact that has produced the most bugs on this branch: a SPARC
`CALL` writes **the address of the CALL instruction itself** into `%o7`,
and the callee returns with `JMPL %o7+8`, stepping over both the call
and its delay slot. A raw `%o7` is therefore *not* a return address; it
is eight bytes short of one. Three rules follow, and every one of them
has been violated at least once:

1. **Converting.** Any raw `%o7` read out of a register, a signal
   context or a frame anchor must have 8 added before the rest of the
   runtime sees it. The runtime resolves a return address by the
   universal rule "subtract one and you are inside the CALL", so an
   unconverted value resolves one instruction too early. That is usually
   still the right *line* — the instructions before a call belong to the
   same statement — which is what makes the mistake so quiet. What it
   loses is anything finer: inlined frames vanish, because the earlier
   instruction lies outside the range of the call that was inlined
   there.

2. **Storing.** Anything that arranges to be *returned into* — an
   injected `sigpanic` or `asyncPreempt` call — must store `target-8`,
   so that the conversion above reconstructs the intended address.

3. **Trusting.** The link register is authoritative only for a frame
   that has not saved it: a leaf, or a frame still inside its prologue.
   Every other frame has stored its own return address at `[sp+120]`,
   and that slot wins. Any call a function makes overwrites `%o7` with
   an address pointing back into that same function, so a signal
   delivered to a framed function generally finds a stale
   return-to-itself value in the register. Trusting it duplicates the
   frame.

The conversions are deliberately centralised in the unwinder and the
signal context rather than spread across their callers, so profiling,
tracing, heap dumps and `testing.T.Helper` all inherit them. Add new
ones there, not at the call site.

There is a related consequence for generated code. Because the return
address points past the delay slot, "return address minus one" lands in
the delay slot, never in the call. The slot must therefore carry the
call's own position, which is why the assembler always appends its own
`RNOP` rather than adopting whatever instruction follows a jump — an
existing `RNOP` may be an inline mark, whose position the inline tree
records as the parent frame's call site.

MIPS, the only other delay-slot architecture Go supports, needs none of
this: `JAL` writes `PC+8` into `$31` directly, so its link register is
already a return address.

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
