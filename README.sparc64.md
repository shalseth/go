# Go on linux/sparc64

An out-of-tree port of the Go gc toolchain to 64-bit SPARC (SPARC V9),
developed on an UltraSPARC T4-1 running Gentoo. This branch adds
`GOARCH=sparc64` to the compiler, assembler, linker and runtime; the
`master` branch tracks upstream unchanged.

The full test suite passes, with cgo. Known to work: Docker (dockerd,
containerd, runc), Forgejo and its runner, and Grafana Alloy.

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
Fujitsu SPARC64 X or later. The compiler emits the VIS3 register-file
moves, `UMULXHI` and `ADDXC`.

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

## What's implemented

* The compiler, assembler and linker, with an SSA backend.
* The runtime: goroutines, the scheduler with cooperative and asynchronous
  preemption, stack growth, and the garbage collector.
* Signals, timers, the vDSO clock, fork/exec, and the `os`, `net` and
  `net/http` packages.
* cgo, with both internal and external linking.
* `-buildmode=pie`, linked externally. The load address is only randomized
  on a kernel whose sparc64 ELF loader does so (`ARCH_HAS_ELF_RANDOMIZE`).
* `-buildmode=c-archive` and `c-shared`.
* The race detector.
* Atomics.
* Disassembly in `go tool objdump` and `pprof`.
* Native builds: the toolchain builds and tests itself on sparc64.

## What is missing

* The `-buildmode` variants `shared` and `plugin`.

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

## ABI

* Go code runs flat, in one register window; `SAVE`/`RESTORE` appear
  only where it calls C or the vDSO, or C calls Go.
* The stack pointer is biased by 2047, and every frame reserves 176
  bytes: the 128-byte window save area and the argument area.
* g is `%l6`, the frame pointer (RFP) `%l5`, the saved return address
  (OLR) `%i7`, the link register `%o7`. TMP and TMP2 are `%i2` and `%l7`;
  `%g7` is the thread pointer.
* A frame's anchors are at `[sp+bias+40]` (frame pointer) and
  `[sp+bias+136]` (return address).
* `CALL` leaves its own address in `%o7`; the return address is `%o7+8`.
* Linux on SPARC keeps the SunOS signal numbers and several SunOS flag
  values.

## Testing

The port is self-hosting, so the suite runs natively:

```sh
cd src && ./make.bash
cd .. && GO_TEST_TIMEOUT_SCALE=2 ./bin/go tool dist test -k
```

The pass condition is zero failures. `GO_TEST_TIMEOUT_SCALE` gives the
slower packages room, since `dist test` runs as many in parallel as there
are CPUs.

The kernel needs `CONFIG_DUMMY` and `CONFIG_CRYPTO_USER_API_HASH` for a
clean run; neither is sparc64-specific.
