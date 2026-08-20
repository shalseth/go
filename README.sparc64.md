# Go on linux/sparc64

An out-of-tree port of the Go gc toolchain to 64-bit SPARC (SPARC V9),
developed on an UltraSPARC T4-1 running Gentoo. This branch adds
`GOARCH=sparc64` to the compiler, assembler, linker and runtime; the
`master` branch tracks upstream unchanged.

Status: `go tool dist test -k` runs 442 packages green, including all of
the `test/` language suite. Five tests fail, and all five are the same
missing piece: there is no sparc64 disassembler, so `cmd/objdump` and
`cmd/pprof` cannot disassemble (see "What is missing"). Grafana Alloy —
an OpenTelemetry collector with roughly two and a half thousand packages
— builds, runs, scrapes metrics and shuts down cleanly on the T4.

## Hardware requirements

VIS3 is the baseline: SPARC T3 or later (T3, T4, T5, S7, M5-M8), or
Fujitsu SPARC64 X or later. Earlier machines - UltraSPARC I through IV,
T1, T2, SPARC64 V/VI/VII - fault with SIGILL during runtime startup.
Verified: a static binary from this branch dies immediately with SIGILL
on an UltraSPARC IIIi (Sun Fire V240).

Three things need VIS3, all of them emitted by the compiler; no
hand-written assembly in the port uses a VIS3 instruction.

  - Moves between the integer and float register files: `MOVXTOD`,
    `MOVDTOX`, `MOVSTOUW`, `MOVWTOS`, emitted by `regMoveOp` in
    `cmd/compile/internal/sparc64/ssa.go`. This is the pervasive one.
    The conversions themselves are plain V9 - `FXTOD` and `FDTOX` work
    everywhere - it is moving the bits across the register files that
    needs VIS3, and the register allocator inserts those moves wherever
    an integer value meets a float one.
  - `UMULXHI`, for `Hmul64u`, `Hmul64` and `Mul64uhilo`.
  - `ADDXC`, for `Add64carry` and `Sub64borrow`.

Supporting a pre-VIS3 machine would not cost a VIS3 build anything: the
selection belongs at build time, in a `GOSPARC64=v9|vis3` feature level
alongside `GOAMD64` and `GOPPC64` in `internal/buildcfg`, so a `vis3`
build would emit exactly the code it emits today. The work is lopsided,
though. Gating `UMULXHI` and `ADDXC` is easy - condition the rules on
the level, drop sparc64 from the `math/bits` intrinsic lists, and give
`Hmul64u` a software fallback, since the generic magic-division rewrite
produces it. The register-file moves are the real cost: without VIS3 the
only route between the files is memory, `stx` then `ldd`, which needs a
scratch slot reserved in every frame before the register allocator
inserts the move. Go used to carry that machinery for 386
(`NeedsFpScratch`); it is gone from the tree, so it would have to be
rebuilt inside frame layout - the part of this port that has produced
every one of its hardest bugs. It has not been attempted.

Independently of VIS3, the runtime reads `%stick` for its cycle counter,
which UltraSPARC-III and later added; `%tick` exists everywhere but
counts each strand's own cycles, and the strands do not agree closely
enough for a process-wide timebase (see "Timebase" below).

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
* The vDSO, for `walltime` and `nanotime`. `time.Now` costs about 278ns
  here rather than the 2043ns two `clock_gettime` syscalls took.
* Atomic intrinsics. `CASW`/`CASD` are SPARC V9's only read-modify-write
  primitives - there is no atomic add and no LL/SC - so exchange, add,
  and and or lower to CAS retry loops, and the 8-bit forms use a 32-bit
  CAS on the containing aligned word. Under TSO a plain load is already
  acquire and a plain store release; only a sequentially consistent
  store needs `MEMBAR #StoreLoad`, and a read-modify-write a full
  barrier.

## What is missing

* cgo and external linking. Everything must be built with
  `CGO_ENABLED=0`.
* The race detector, and the `-buildmode` variants beyond `exe`.
* A disassembler. `cmd/internal/disasm` has no sparc64 support and
  there is no `golang.org/x/arch/sparc64asm` to build on, so `go tool
  objdump` and `pprof`'s annotated-assembly view do not work. This is
  the only part of the test suite still failing — `TestDisasm`,
  `TestDisasmCode`, `TestDisasmGnuAsm`, `TestDisasmGoobj` and
  `TestGoobjFileNumber`, across `cmd/objdump` and `cmd/pprof`.
  Collecting and reading profiles is unaffected; only disassembly is.

* Optimised assembly for the routines that ship a generic Go
  fallback: `math/big`'s `addVV`/`subVV`/`mulAddVWW`/`lshVU`/`rshVU`,
  `internal/bytealg`'s `cmpbody`, and `crypto/internal/fips140/bigmod`'s
  `addMulVVW1024`/`1536`/`2048`. These are correct but slow; they are a
  performance gap, not a correctness one.

  A cause sits underneath them: `bits.Add64` and `bits.Sub64` are not
  intrinsified on this branch, so each limb of a carry chain costs the
  pure-Go body - two adds plus five logical ops to recover the carry -
  where `ADDCC` and VIS3's `ADDXC` need four instructions, and a
  hand-written loop that keeps the carry live in `xcc` needs one.
  (`bits.Mul64` *is* intrinsified, via `MULDU` to `MULD` and `UMULXHI`;
  the multiply itself is already two instructions.) Intrinsifying the
  carry pair would speed up the generic fallbacks themselves, and every
  other user of `math/bits`, before a line of assembly is added.

  Measured on an idle T4: `addVV` 6.2ns per 64-bit limb, `mulAddVWW`
  6.1ns, `addMulVVW` 8.9ns; RSA-2048 sign 19.7ms, verify 0.66ms. Note
  also that the generic implementations are *faster* than the `asm`
  entry points, which on this branch are only forwarding wrappers that
  do not inline: 54.7ns against 71.6ns for a one-word `addVV`.

## Timebase

`cputicks` reads `%stick`, not `%tick`. `%tick` counts the strand's own
cycles: passing a token around a ring of OS-thread-locked goroutines on
the T4, a third of the handoffs saw `%tick` go *backwards* across the
happens-before edge, by as much as 6ms. Callers assume a process-wide
timebase — the debug log merges its per-P shards by this value, and the
mutex profile subtracts two reads taken on different threads — so a
per-strand counter yields impossible orderings and durations.

`%stick` is driven from a system-wide reference and measured zero
backwards steps over the same experiment. It runs slower, about 1GHz
against 2.85GHz on this machine, which costs resolution the runtime does
not need; `ticksPerSecond` calibrates against `nanotime` either way.

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
