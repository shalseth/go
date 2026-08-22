# Go on linux/sparc64

An out-of-tree port of the Go gc toolchain to 64-bit SPARC (SPARC V9),
developed on an UltraSPARC T4-1 running Gentoo. This branch adds
`GOARCH=sparc64` to the compiler, assembler, linker and runtime; the
`master` branch tracks upstream unchanged.

Status: with **cgo enabled**, `go test std cmd` runs 380 packages green,
including the full runtime suite, every cgo, callback, traceback and
profiling test, and all of `cmd/go`'s script tests. The only failures are
the missing disassembler (`cmd/objdump`, `cmd/pprof`; see "What is
missing"), two `net` tests that need a kernel with `CONFIG_DUMMY`, and the
`moddeps` provenance check, which an out-of-tree port that patches the
vendored `golang.org/x/sys` cannot satisfy until that support is upstream.
Both internal and external linking of cgo programs work, including a psABI
PLT and GOT for dynamic imports under internal linking. Grafana Alloy — an
OpenTelemetry collector with roughly two and a half thousand packages —
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
Fujitsu SPARC64 X or later. Earlier machines - UltraSPARC I through IV,
T1, T2, SPARC64 V/VI/VII - fault with SIGILL during runtime startup; a
binary from this branch dies immediately on an UltraSPARC IIIi.

Three things need it, all emitted by the compiler: the register-file
moves `MOVXTOD`, `MOVDTOX`, `MOVSTOUW` and `MOVWTOS`, which the register
allocator inserts wherever an integer value meets a float one;
`UMULXHI`, for the high half of a 64x64 multiply; and `ADDXC`, for
`Add64carry` and `Sub64borrow`. Going below VIS3 would mean a build-time
feature level in the style of `GOAMD64`, software fallbacks for the
multiply and the carries, and - the hard part - reimplementing the
register-file moves as memory round trips, which needs a scratch slot
reserved in every frame. It has not been attempted.

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
native `go build` doubles as a stress test of the port. Some runtime
bugs surface only under that kind of load.

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
* cgo, in both link modes. Go→C calls, C→Go callbacks, signals and
  profiling during cgo execution, and thread exit through pthread TSD
  destructors all work. The port survives the register-window hazards
  this required: the frame pointer lives in `%l5` rather than `%i6`
  (which is the hardware's stack pointer for the adjacent window and is
  walked by the kernel in `clone`), the return-address chain slot sits
  outside the `%i7` image the hardware spills over during signal
  handling, and every Go↔C transition opens or flushes register windows
  explicitly. See the commit history for the full analysis.
* Internal linking of cgo binaries: the linker consumes host ELF
  objects, applies the psABI GOTDATA code models, and emits the
  patched-code PLT with `R_SPARC_JMP_SLOT` and a GOT with
  `R_SPARC_GLOB_DAT` for dynamic imports; `DT_PLTGOT` names the PLT as
  the SPARC dynamic linker expects.
* `TrailingZeros` via the hardware `POPC` instruction, keeping the
  allocator's hot path inlined.
* Atomic intrinsics. `CASW`/`CASD` are SPARC V9's only read-modify-write
  primitives - there is no atomic add and no LL/SC - so exchange, add,
  and and or lower to CAS retry loops, and the 8-bit forms use a 32-bit
  CAS on the containing aligned word. Under TSO a plain load is already
  acquire and a plain store release; only a sequentially consistent
  store needs `MEMBAR #StoreLoad`, and a read-modify-write a full
  barrier.

## What is missing

* The race detector, and the `-buildmode` variants beyond `exe`:
  `pie`, `c-archive` and `c-shared` need a position-independent code
  model in the compiler, which SPARC makes expensive (no PC-relative
  addressing; PIC costs a dedicated GOT register and GOT-relative
  sequences for every global).
* A disassembler. `cmd/internal/disasm` has no sparc64 support and
  there is no `golang.org/x/arch/sparc64asm` to build on, so `go tool
  objdump` and `pprof`'s annotated-assembly view do not work. This is
  the only part of the test suite still failing — `TestDisasm`,
  `TestDisasmCode`, `TestDisasmGnuAsm`, `TestDisasmGoobj` and
  `TestGoobjFileNumber`, across `cmd/objdump` and `cmd/pprof`.
  Collecting and reading profiles is unaffected; only disassembly is.

## Performance

### Carry chains

VIS3's `ADDXC` copies the 64-bit carry out of `xcc` into a register, so
`bits.Add64` is four instructions - `ADDCC`, `ADDXC`, `ADDCC`, `ADDXC` -
rather than the two adds plus five logical operations the pure-Go body
needs to recover the carry bit. `bits.Sub64` is the same shape with
`SUBCC`, and `bits.Mul64` is `MULD` with `UMULXHI`.

`ADDXCcc` both uses and sets the carry, so a chain in assembly runs one
instruction per limb with the carry never leaving the condition codes.
Nothing between the adds may disturb `xcc`, which rules out the usual
loop-counter compare; `BRZ` and `BRNZ` branch on a register and leave
the codes alone, so the loops use those. There is no 64-bit
subtract-with-borrow - VIS3 added `ADDXC` and `ADDXCcc` but no subtract
counterpart - so borrow chains run as add chains over the complement:
`x - y - b` is `x + ^y + (1-b)`, and the carry out is one minus the
borrow out, at one `XNOR` per limb.

Two files use this:

  - `math/big/arith_sparc64.s` holds `addVV`, `subVV`, `mulAddVWW`,
    `addMulVVWW`, `lshVU` and `rshVU`, four limbs per iteration. The
    shifts run descending and ascending respectively so that `z` and `x`
    may alias.
  - `crypto/internal/fips140/bigmod/nat_sparc64.s` holds
    `addMulVVW1024`, `1536` and `2048` as three entry points that set a
    group count and tail-jump into a shared core. `UMULXHI` and `MULD`
    give the two halves of each product; the four multiplies of a group
    are independent and issue together, overlapping the serial carry
    chain that follows them.

`math/big`'s `internal/asmgen`, which generates `arith_$GOARCH.s` for
the other architectures, has no sparc64 definition: its model wants a
single subtract-with-borrow instruction. `arith_sparc64.s` is
hand-written, and the generator's test only checks the architectures it
knows, so the file is neither flagged nor kept in step with the
generator.

### String and byte scanning

SPARC traps on unaligned access, so words are only ever loaded from
8-aligned addresses. Being big-endian pays off twice over: an unsigned
comparison of two differing words gives the same answer as comparing
their first differing byte, and the first match in a word is its most
significant marked lane. Neither needs the byte swap the little-endian
ports perform, which matters because SPARC has no byte-swap
instruction.

  - `internal/bytealg/compare_sparc64.s` holds `Compare` and
    `runtime.cmpstring`, sharing one body. Operands that share a
    misalignment are byte-compared up to alignment; operands whose
    addresses differ mod 8 have one side aligned by hand while the
    other's words are assembled from two aligned loads and a shift. That
    path runs only while 16 bytes remain, so the second load cannot
    reach past the end of the shorter operand.
  - `internal/bytealg/indexbyte_sparc64.s` and `count_sparc64.s` hold
    `IndexByte`, `IndexByteString`, `Count` and `CountString`. Xor a
    word with the byte broadcast across all eight lanes and a match
    becomes a zero byte, which `(v - 0x01..01) &^ v & 0x80..80` detects.
    That test is exact as a yes/no answer but not per lane, since the
    subtraction borrows between lanes, so `IndexByte` finds the lane by
    rescanning those eight bytes - a few instructions once per call, and
    no need for `LZD`. `Count` reads the mask itself, so it builds an
    exact one the other way: masking off each lane's high bit and adding
    `0x7f` cannot carry out of the lane. Summing the lanes needs no
    `POPC` either, since shifting the markers down and multiplying by
    `0x01..01` accumulates them into the top byte.

`bytealg.Index` is not implemented, so `bytes.Index` and `strings.Index`
use their own loop over `IndexByte` and `Equal`. A real one earns its
keep only with a vector prefilter, which this port cannot emit.

### SHA-256

`crypto/internal/fips140/sha256/sha256block_sparc64.s` uses the T4's
`sha256` instruction, which hashes a whole 512-bit block per issue. It
takes no operands: the eight state words come from `%f0-%f7`, the block
from `%f8-%f23`, and the result is written back over the state. So the
assembly loads the state once, loads each block over the same eight
double registers, and stores the state at the end. SHA-256 defines its
block as big-endian words, which is how a load already lands them, so
this too byte-swaps nothing.

The instruction reads its block with `ldd`, which traps on an unaligned
address, so misaligned input is copied a block at a time through an
aligned buffer. That buffer is declared as words: a `[64]byte` local
carries no alignment guarantee.

The path is gated on `SPARC64.HasCrypto`, which `internal/cpu` derives
from `AT_HWCAP`, and `GODEBUG=cpu.crypto=off` selects the generic
implementation instead.

### Measurements

Measured on an idle T4-1.

Carry-chain arithmetic gains from both the intrinsics and the assembly.
The first column is what the generic Go code compiles to without the
`Add64` and `Sub64` intrinsics, as on a compiler that lacks them; the
second is the generic path in this toolchain, selected with `purego` and
`math_big_pure_go`; the third is the assembly.

| | no intrinsics | generic | assembly | |
|---|---|---|---|---|
| `addVV`, per limb | 6.20ns | 3.97ns | 1.58ns | 3.9x |
| `subVV`, per limb | 6.20ns | 3.98ns | 1.69ns | 3.7x |
| `mulAddVWW`, per limb | 6.10ns | 4.18ns | 1.88ns | 3.2x |
| `addMulVVWW`, per limb | 8.90ns | 6.32ns | 2.84ns | 3.1x |
| bigmod `MontgomeryMul` | 24.5us | 19.3us | 7.11us | 3.4x |
| bigmod `ExpBig` | 54.8ms | 39.3ms | 19.8ms | 2.8x |
| RSA-2048 sign | 19.66ms | 15.66ms | 6.78ms | 2.9x |
| RSA-2048 verify | 0.656ms | 0.526ms | 0.211ms | 3.1x |

Scanning and hashing depend only on the assembly, and on the hardware
instruction for SHA-256, which `GODEBUG=cpu.crypto=off` turns off.

| | generic | assembly | |
|---|---|---|---|
| `bytes.Compare`, 64B | 331ns | 32.0ns | 10.4x |
| `bytes.Compare`, 4KB | 19.6us | 1.29us | 15.2x |
| `bytes.IndexByte`, 4KB | 14.0us | 1.52us | 9.2x |
| `bytes.Count`, 4KB | 6.24us | 1.69us | 3.7x |
| `sha256`, 1KB | 17.9 MB/s | 566 MB/s | 32x |
| `sha256`, 8KB | 19.2 MB/s | 846 MB/s | 44x |

`bytes.Compare` reaches 3.1 GB/s on 4KB buffers against 197 MB/s for the
byte-at-a-time generic version, and 2.3 GB/s when the two operands'
addresses differ mod 8. `bytes.IndexByte` reaches 2.8 GB/s. SHA-256 runs
at 3.2 cycles per byte against roughly 150 for the Go code, and the
block loop alone measures 898 MB/s.

Substring search gains less than `IndexByte` alone suggests, because
`bytes.Index` scans with `IndexByte` and confirms with `Equal`: over a
5KB block of HTTP headers, a needle that is absent runs 3.9x faster than
the generic version, `Content-Length` 2.0x, and `\r\n\r\n` - whose first
byte occurs 160 times in that block - 1.5x, because the confirmations
dominate. The `bytes` package's own `BenchmarkIndex` shows no difference
at all: it searches an all-zero buffer for a needle whose leading bytes
are also zero, so the first-byte test never fails and `IndexByte` is
never reached.

ECDSA has no assembly here. P-256 and friends use `nistec`'s own field
arithmetic rather than `bigmod`, and that code is generic Go.

### Hardware left on the table

The T4 reports a good deal more than SHA-256:

```
aes, des, kasumi, camellia, md5, sha1, sha256, sha512,
mpmul, montmul, montsqr, crc32c, popc
```

`montmul` and `montsqr` are the Montgomery multiply and square that
`bigmod` does in software; `aes` is the equivalent of AES-NI; `crc32c`
is what `hash/crc32` wants for the Castagnoli polynomial; `md5`, `sha1`
and `sha512` are single instructions like `sha256`. Going by what
SHA-256 measures above, and by what OpenSSL gets from the same
instructions, AES-GCM is worth 10-20x, CRC32C 10x, and Montgomery
multiply another 3x on top of the assembly above.

Each one needs its opcode in the assembler and the assembly around it;
the encoding and the `AT_HWCAP` gating are already in place. OpenSSL's
`aest4-sparcv9.pl` and `sparct4-mont.pl` are working references for the
sequences.

### Assembly checking

`cmd/vet`'s `asmdecl` pass carries a sparc64 entry, so `go vet` checks
every `FP` reference in the port's assembly against the Go declaration
it belongs to.

## Timebase

`cputicks` reads `%stick`, not `%tick`. `%tick` counts the strand's own
cycles: passing a token around a ring of OS-thread-locked goroutines on
the T4, a third of the handoffs saw `%tick` go *backwards* across the
happens-before edge, by as much as 6ms. Callers assume a process-wide
timebase — the debug log merges its per-P shards by this value, and the
mutex profile subtracts two reads taken on different threads — so a
per-strand counter yields impossible orderings and durations.

Wall and monotonic clock reads are separate from `cputicks`: they go
through the vDSO. The runtime looks up `__vdso_clock_gettime` at startup
(`vdso_linux_sparc64.go`) and calls it from `nanotime` and `walltime`,
falling back to a `ta 0x6d` syscall if the lookup fails. On the T4 that
is worth a great deal, because a trap here is expensive: one monotonic
read costs 142ns through the vDSO against 1056ns through the kernel, so
the vDSO saves about 0.9us on every clock read, and `time.Now`, which
reads both clocks, costs 260ns. It is also why `sigFetchG` needs its
recovery path - a signal can land inside vDSO code, where the g register
does not hold a valid g.

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
always agree with the ones the unwinder reads. The agreement is briefly untrue in three
places, and each is either ordered so that the invariant holds or marked
non-preemptible: the prologue between pushing the frame and publishing
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
