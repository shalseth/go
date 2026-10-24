# Go on linux/sparc64

An out-of-tree port of the Go gc toolchain to 64-bit SPARC (SPARC V9),
developed on an UltraSPARC T4-1 running Gentoo. This branch adds
`GOARCH=sparc64` to the compiler, assembler, linker and runtime; the
`master` branch tracks upstream unchanged.

Status: with **cgo enabled**, the entire test suite is green.
Both internal and external linking of cgo programs work, including a psABI
PLT and GOT for dynamic imports under internal linking. Grafana Alloy, an
OpenTelemetry collector of roughly two and a half thousand packages,
builds, runs, scrapes metrics and shuts down cleanly on the T4.

## Provenance

This port does not start from nothing. The assembler backend
(`cmd/internal/obj/sparc64`) and the core runtime assembly
(`runtime/asm_sparc64.s`, `runtime/asm_sparc64.h`, `runtime/tls_sparc64.s`)
were revived from the 2016 linux/solaris-sparc64 tree,
`minux/go.sparc64 @ 9b8610d`, and rebased onto the current
`cmd/internal/obj` API. The instruction encodings are unchanged from that
work.

That code comes from the original gc SPARC64 port effort, and this port
would have been considerably harder to begin without it. Particular
thanks to **Aram Hăvărneanu**, whose SPARC64 work it descends from, and
to **Shenghou Ma**, in whose tree it was found.

Everything else here is new: the compiler backend and SSA lowering rules,
the linker support, the flat-frame ABI and the register-window discipline
that goes with it, the signal, vDSO and syscall layers, and the cgo
boundary.

## Hardware requirements

VIS3 is the baseline: SPARC T3 or later (T3, T4, T5, S7, M5-M8), or
Fujitsu SPARC64 X or later.

Three things need it, all emitted by the compiler: the register-file
moves `MOVXTOD`, `MOVDTOX`, `MOVSTOUW` and `MOVWTOS`, which the register
allocator inserts wherever an integer value meets a float one;
`UMULXHI`, for the high half of a 64x64 multiply; and `ADDXC`, for
`Add64carry` and `Sub64borrow`.

The runtime reads `%stick` for its cycle counter (see "Timebase" below).

## Building

Cross-compile from any supported host (amd64 here). Bootstrapping the
tree needs Go 1.26 or later:

```sh
cd src && GOROOT_BOOTSTRAP=/path/to/go1.26 ./make.bash
GOOS=linux GOARCH=sparc64 ../bin/go build ./yourprogram
```

The toolchain also builds and runs natively on sparc64. Cross-build a
bootstrap tree, copy it to the target, and build the source there:

```sh
cd src && GOOS=linux GOARCH=sparc64 GOROOT_BOOTSTRAP=/path/to/go1.26 ./bootstrap.bash
scp ../../go-linux-sparc64-bootstrap.tbz target:
ssh target 'tar xjf go-linux-sparc64-bootstrap.tbz'
ssh target 'git clone https://github.com/shalseth/go.git'
ssh target 'GOROOT_BOOTSTRAP=$HOME/go-linux-sparc64-bootstrap go/src/make.bash'
```

The gc compiler is one of the heaviest Go workloads there is, so a
native `go build` doubles as a stress test of the port. Some runtime
bugs surface only under that kind of load.

## What's implemented

* The compiler, assembler and linker, with an SSA backend.
* The runtime: goroutines, the scheduler with cooperative and asynchronous
  preemption, stack growth, and the garbage collector.
* Signals, timers, the vDSO clock, fork/exec, and the `os`, `net` and
  `net/http` packages.
* cgo, with both internal and external linking.
* `-buildmode=pie`, linked externally. The load address is only randomized
  on a kernel whose sparc64 ELF loader does so (`ARCH_HAS_ELF_RANDOMIZE`).
* The race detector.
* Atomics.
* Disassembly in `go tool objdump` and `pprof`.
* Native builds: the toolchain builds and tests itself on sparc64.

## What is missing

* The `-buildmode` variants `c-archive`, `c-shared`, `shared` and
  `plugin`.

## Performance

Hardware instructions in use:

* VIS3 `ADDXC` and `UMULXHI` for 64-bit carry chains and multiplies, in
  `math/big` and in the FIPS `bigmod` package: about 3x for RSA-2048
  against code without them.
* The T4 `sha1`, `sha256` and `sha512` instructions: 20-45x.
* The T4 `aes` instructions and VIS3 `XMULX` for AES-GCM: 60-75x.
* The T4 `crc32c` instruction for CRC-32C: about 14x.
* `POPC` for counting trailing zeros.
* `FSQRT` for `math.Sqrt`.

`bytes.Compare`, `bytes.Equal`, `bytes.IndexByte` and `bytes.Count` have
word-at-a-time assembly. `Compare`, `IndexByte` and `Count` measure 4-15x
faster than the generic code.

Not used yet, although the T4 and later have them: the `des`,
`camellia`, `kasumi`, `md5`, `mpmul`, `montmul` and `montsqr`
instructions. There is no dedicated assembly for substring search or for
the elliptic-curve field arithmetic.

## Timebase

`cputicks` reads `%stick`, not `%tick`. `%tick` counts the strand's own
cycles: passing a token around a ring of OS-thread-locked goroutines on
the T4, a third of the handoffs saw `%tick` go *backwards* across the
happens-before edge, by as much as 6ms. Callers assume a process-wide
timebase: the debug log merges its per-P shards by this value, and the
mutex profile subtracts two reads taken on different threads. A
per-strand counter yields impossible orderings and durations.

Wall and monotonic clock reads are separate from `cputicks`: they go
through the vDSO. The runtime looks up `__vdso_clock_gettime` at startup
(`vdso_linux_sparc64.go`) and calls it from `nanotime` and `walltime`,
falling back to a `ta 0x6d` syscall if the lookup fails. That matters on
the T4, where a trap is expensive. It is also why `sigFetchG` needs its
recovery path: a signal can land inside vDSO code, where the g register
does not hold a valid g.

`%stick` is driven from a system-wide reference and measured zero
backwards steps over the same experiment. It runs slower, about 1GHz
against 2.85GHz on this machine, which costs resolution the runtime does
not need; `ticksPerSecond` calibrates against `nanotime` either way.

## Notes on the ABI

Go does not use SPARC register windows: `SAVE`/`RESTORE` appear nowhere
in generated code, and one window serves the whole program. Frames are
flat, and two registers act as per-frame anchors:

* `%l5` (RFP): the frame's entry stack pointer, biased.
* `%i7` (OLR): the frame's own return address.

RFP cannot be `%i6`. That register is the hardware's stack pointer for
the window above, so a frame pointer every prologue rewrites would
decide where the hardware spills a window unrelated to this frame, onto
the 128 bytes backing whatever registers were interrupted there. With
RFP in `%l5`, `%i6` keeps the caller stack pointer the C side left in
it, and such a spill lands harmlessly on the C stack.

A frame keeps its anchors at `[sp+bias+40]`, the slot the hardware
spills `%l5` to, and `[sp+bias+136]`. The first is idempotent under a
hardware spill: spilling this window writes the value the slot already
holds, so a trap cannot damage it. The second cannot be the `%i7` image
at +120, because during signal handling the live `%i7` holds the
interrupted frame's return address, so a trap taken mid-handler spills
that stale value over the slot. +120 is still written, since the
hardware expects it there, but nothing reads it back; `pushCall` owns
+128.

The agreement is briefly untrue in three places, each either ordered so
the invariant holds or marked non-preemptible: the prologue between pushing the frame and publishing
the return address, the epilogue between reloading the caller's anchors
and raising the stack pointer, and any signal handler that opens a
second register window. Code that runs in those windows and expects to
unwind is the classic way to break this port.

### Return addresses

The most consequential difference from the other architectures: a SPARC
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
   still the right *line*, since the instructions before a call belong
   to the same statement, which is what makes the mistake so quiet. What
   it loses is anything finer: inlined frames vanish, because the earlier
   instruction lies outside the range of the call that was inlined
   there.

2. **Storing.** Anything that arranges to be *returned into*, such as
   an injected `sigpanic` or `asyncPreempt` call, must store `target-8`,
   so that the conversion above reconstructs the intended address.

3. **Trusting.** The link register is authoritative only for a frame
   that has not saved it: a leaf, or a frame still inside its prologue.
   Every other frame has stored its own return address at `[sp+136]`,
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
call's own position, which is why the assembler appends its own `RNOP`
rather than adopting whatever instruction follows a jump: an existing
`RNOP` may be an inline mark, whose position the inline tree records as
the parent frame's call site.

The one exemption is `ret; restore`, the SPARC idiom for leaving a
register window: there the restore is meant to run as part of the
return, so it keeps the slot. A `RESTORE` after a `CALL` is refused
outright, because control comes back from a call and a restore in its
slot would rotate the window away before the callee ever runs, taking
`%o0`, where the callee's result arrives, with it. Put the result
somewhere first (`MOVD O0, I0`, as `·asmcgocall` does).

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

The port is self-hosting, so the suite runs natively:

```sh
cd src && ./make.bash
cd .. && GO_TEST_TIMEOUT_SCALE=2 ./bin/go tool dist test -k
```

`-k` keeps going past a failing package. Without it the run stops at the
first failure and never reaches the twenty later phases. The pass condition
is zero `--- FAIL` across 22 phases.

`GO_TEST_TIMEOUT_SCALE` scales the per-package harness deadline, and is set
because `dist test` sizes its background parallelism to `runtime.NumCPU()`:
the suite competes with itself, and `runtime`, which needs about 450 s alone,
can otherwise exceed its 600 s budget. No concurrency limit is set.
`MERGING-UPSTREAM.md` has the detail, including why a test's own internal
budget is a different problem that this variable does not touch.

On an UltraSPARC T4-1, 64 threads at 2.85 GHz, that is about 12 minutes for
`make.bash` and 29 for the suite.

Test binaries can also be cross-compiled, for a target with no toolchain on
it:

```sh
GOOS=linux GOARCH=sparc64 go test -c -o sort.test sort
scp sort.test target: && ssh target ./sort.test -test.short
```

Tests that read files from their own source directory (`os`, `net`,
`io/fs`, `text/template`) need that directory copied alongside the
binary.

Two kernel options are needed for a clean sweep, neither of them
sparc64-specific:

* `CONFIG_DUMMY`: `net` uses a dummy interface for two tests.
* `CONFIG_CRYPTO_USER_API_HASH`: `golang.org/x/sys/unix`'s
  `TestSockaddrALG` binds an `AF_ALG` socket.
