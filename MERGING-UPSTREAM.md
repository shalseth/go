# Merging upstream into the sparc64 port

Procedure for pulling golang/go master into the `sparc64` branch, validating
the result, and installing the toolchain.

The merge itself is text work, but the rule generator, the build and the test
suite all need a working sparc64 Go, so in practice every step runs on a
sparc64 host. Long jobs should be detached (`setsid`, `screen`) so they survive
a dropped connection.

Rebuilding downstream software against the new toolchain is out of scope here;
for Grafana Alloy see the `alloy-sparc64` repo.

Paths below are this setup's; adjust for another.

| path | what |
|---|---|
| `/root/goport/go` | port checkout, branch `sparc64` |
| `/root/gomerge` | throwaway worktree for the merge |

Rough costs on an UltraSPARC T4-1: `make.bash` ~12 min, `dist test -k` ~40 min,
`emerge go-9999` ~12 min.

**The port is expected to test clean**: 22 phases, zero failures. That is the
standard every merge is held to, so there is no need to capture a fresh
baseline first. Any failure is a real result: investigate it, do not explain it
away.

---

## 1. Merge onto a throwaway branch

Never merge onto `sparc64` directly. Nothing is at risk until it is green.

    cd /root/goport/go
    git remote add upstream https://github.com/golang/go.git   # first time only
    git fetch upstream master
    git rev-list --count sparc64..upstream/master              # how far behind

    git worktree add /root/gomerge -b merge-upstream-$(date +%Y%m%d) sparc64
    cd /root/gomerge && git merge upstream/master

A worktree keeps the working checkout untouched and buildable while the merge
is still in doubt.

**Merge, never rebase.** Rebasing replays 150+ port commits through the drift,
conflicting on generated files repeatedly, and rewrites hashes so release tags
point into dead history. The fast-forward property is what keeps old release
tags reachable.

---

## 2. Resolve the conflicts

For each conflicted file, find out what *this port* actually changed before
deciding anything:

    BASE=$(git merge-base sparc64 upstream/master)
    git diff $BASE sparc64 -- <file>

The deltas are usually tiny: a single op in a list, one helper function,
while the conflict spans hundreds of lines because upstream moved the file's
contents elsewhere. In that shape the resolution is *take upstream wholesale,
then re-apply the small delta in its new home*:

    git checkout --theirs <file> && git add <file>

Watch for a delta upstream has since made redundant: taking upstream then
leaves an unused variable behind and the build fails on "declared and not
used". Delete the leftovers as part of the same resolution.

---

## 3. The adaptations no conflict marker shows

The part that costs a day if skipped. A package *move* appears as
delete-plus-add, so comparing changed-file lists never flags it. None of the
following produced a conflict in the 2026-08 merge.

**a. Stale generated files.** When the generator's output path moves, the old
file is left behind and still compiles. After running the generator, look for
duplicates:

    git status --porcelain | grep '^??'
    ls src/cmd/compile/internal/ssa/rewriteSPARC64.go   # should NOT exist

**b. Package qualifier drift.** Ops moved `ssa` → `ssaop`. The generator's own
check catches the arch file (`OpSPARC64ADD has no code generation in
../../sparc64/ssa.go`), but only after a full run. A regex on `ssa\.Op([A-Z])`
misses the bare `ssa.Op` *type*; match `\bssa\.Op\b` as well.

**c. Unqualified helpers in `SPARC64.rules`.** Arch rules files call
`ssa.Is32Bit`, `ssa.B2i`, `ssa.IsPtr` rather than the unexported names. Do not
fix these one build at a time. Scan the whole file against what the other
arches do:

```python
# run from src/cmd/compile/internal/ssa/_gen
import re, glob, collections
qualified = collections.defaultdict(set)
for f in (f for f in glob.glob("*.rules") if f != "SPARC64.rules"):
    for m in re.finditer(r"\b(ssa|ssaop)\.([A-Za-z_]\w*)", open(f).read()):
        qualified[m.group(2)].add(m.group(1))
src = "".join(l.split("//")[0] for l in open("SPARC64.rules"))
for name in sorted(set(re.findall(r"(?<![.\w])([A-Za-z_]\w*)", src))):
    if name in qualified:
        print("%-24s -> %s.%s" % (name, sorted(qualified[name])[0], name))
```

Scan the **whole file**, not just rule conditions: `PanicBoundsC` appears on
the result side and a condition-only scan reports nothing.

**d. Helper signatures, not just names.** `logLargeCopy` was split into
`LogLargeCopy(string, XPos, int64)` and
`LogLargeCopyValue(*Value, int64) bool`. A name-based rename picks the wrong
one and gets as far as the type checker. When a rename target exists, check its
signature against the call site in another arch's rules.

**e. Ops that write the condition codes.** `flagalloc` uses an op's
`clobberFlags` to know where a `Flags` value dies. An op whose expansion emits
a CC-writing instruction without declaring it leaves the allocator believing a
comparison result survives across it, and the consumer then branches on codes
the op destroyed. Nothing catches this: it builds, the tests that exercise the
op pass, and the damage surfaces somewhere else entirely, as heap corruption,
or a nil dereference inside the GC sweeper.

This has happened twice, both times in newly added multi-instruction ops: the
atomics in 2026-08 and `ADDCARRY`/`SUBBORROW` in 2026-09, the latter reaching a
released toolchain. Run the check after any op change:

    python3 misc/sparc64/flagcheck.py

It exits non-zero and names the ops. Note `MOV<cond>` reads the codes rather
than writing them, so `MOVCC` is not a hit; `CMP`, `FCMP` and anything ending
in `CC` are.

**f. New syscall constants.** Upstream adds a syscall to `exec_linux.go` and
the generated `zsysnum_linux_sparc64.go` has never heard of it. Take the number
from the target's own headers, never from another architecture:

    grep landlock /usr/include/asm/unistd_64.h

Then regenerate and iterate:

    cd src/cmd/compile/internal/ssa && go run -C=_gen .   # must exit 0
    cd /root/gomerge/src && ./make.bash

`make.bash` fails fast (~2 min) at the bootstrap type-check, so iterating on it
is cheap. Once it passes it goes on to build the standard library with the
*new* compiler, which is where codegen problems surface.

---

## 4. Test

    cd /root/gomerge/src
    GO_TEST_TIMEOUT_SCALE=2 ../bin/go tool dist test -k > /root/mergebuild-k.log 2>&1

Use `-k` (keep going). Without it `dist test` stops at the first failing
package and never reaches the twenty later phases, which is how an
internal-link cgo bug stayed hidden for six days, with every run reaching two
phases out of twenty-two.

    grep -cE '^--- FAIL' /root/mergebuild-k.log     # expect 0
    grep -c '^#####' /root/mergebuild-k.log         # expect 22 phases
    grep -c 'test timed out' /root/mergebuild-k.log # expect 0
    grep -c 'signal:' /root/mergebuild-k.log        # expect 0

Zero failures across 22 phases is the pass condition, with nothing skipped for
the port: the `cmd/objdump` and `cmd/pprof` disassembly tests used to skip
through `mustHaveDisasm` and have run since the sparc64 disassembler landed
(2026-09-14).

**Every failure is an issue to investigate**, whether or not it looks related
to the merge. Two failure modes produce no `--- FAIL` line at all and so are
invisible to that grep alone:

* A crash of the *compiler itself* - `signal: segmentation fault` against a
  package name, with no Go traceback.
* A package-level timeout - `panic: test timed out after 10m0s`, which names
  the hung test under `running tests:` and reports as a bare `FAIL <pkg>`.

Both make `dist test` exit non-zero while `--- FAIL` counts zero, so a run
scored on `--- FAIL` alone reads as clean when it is not. Check the exit
status as well as the greps.

Two things worth knowing while investigating, neither of them an excuse
to move on:

* `time.TestLongAdjustTimers` is load-sensitive: about 12 s of work against a
  hard 60 s deadline. Re-run it alone on an idle machine before blaming
  contention. `runtime.TestEINTR` is the same category.

  The load is the suite's own. `cmd/dist` sets `maxbg` to `runtime.NumCPU()`
  whenever the worklist holds the `GOMAXPROCS=2 runtime` phase, which it always
  does, so a 64-thread T4 runs 64 test binaries at once: measured at 108 test
  processes and 259 runnable against 32 CPUs. Affinity does not help. `NumCPU`
  reads `sched_getaffinity`, so `taskset` scales `maxbg` down with it and the
  ratio is unchanged.

* **Two kinds of timeout. Do not confuse them.**

  **Harness deadline.** `panic: test timed out after 10m0s`, reported as a bare
  `FAIL <pkg>` with no `--- FAIL`. Scaled by `GO_TEST_TIMEOUT_SCALE`, which
  `cmd/dist/test.go` parses as an integer and applies in `scaledTimeout`. No
  source change, and it survives every merge. Preferred over re-adding the
  per-arch timeout table upstream deleted, which would conflict on every merge.

  **A test's own budget.** `TestLongAdjustTimers`' `AfterFunc(60*Second, ...)`
  giving `timer expired`. The variable does nothing here. Only the test source
  helps, gated on GOARCH: `src/time/tick_test.go` widens 60 s to 5 min and the
  read timeout 5 s to 30 s, sparc64 only.

  Step 4 sets `GO_TEST_TIMEOUT_SCALE=2`. Measured 2026-09-28:
  `runtime` alone on an idle box is `ok 449.137s` against a 600 s budget, so a
  1.34x slowdown busts it, and the suite's 45x parallelism supplies far more.
  `../test` (`cmd/internal/testdir`) sits at 471 s. Scale 2 gives both 1200 s,
  about 2.6x their idle cost. Raise it only if that proves tight: the budget is
  also the hang detector, so keep it as low as works.

  The knob is for a known load characteristic, not for an unexplained failure.
  On a timeout, read what `running tests:` names and how long it had run. The
  2026-09-28 case named `TestSUID (1s)`: nothing hung, the package was slow. A
  real hang names a test whose duration is close to the whole budget.
* `runtime.TestSegv/SegvInCgo` printing `runtime: g N: unexpected return pc`
  was a real backend bug, fixed 2026-09-08: a jump's `Spadj` landed on its
  branch delay slot, so the epilogue described the slot as still owning the
  frame the `ADD $framesize, RSP` before it had released. If it returns, the
  `pctospadj` table is the place to look, not the test:

      go build -gcflags="-d pctab=pctospadj" runtime

  The slot after a `RET` must not carry the frame size. Note the sibling
  `unknown pc` message is a *different* case that the test skips by design
  (go.dev/issue/50979), so grep for the exact wording before concluding
  anything.

* Some failures depend on kernel configuration rather than on the port. If one
  looks like that, confirm it against the running kernel's config before
  changing any Go code.

---

## 5. Land it

    cd /root/gomerge && git add -A && git commit
    cd /root/goport/go && git merge --ff-only merge-upstream-YYYYMMDD
    git merge-base --is-ancestor origin/sparc64 sparc64   # confirm fast-forward
    git push origin sparc64
    git worktree remove /root/gomerge

Then tag the merge, so any past toolchain can be checked out and rebuilt:

    TAG=sparc64-$(date +%Y%m%d)
    git tag -a $TAG -m "merged upstream through $(git rev-parse --short upstream/master)"
    git push origin $TAG

Annotated, not lightweight: the tag carries which upstream commit it caught
up to, which is the thing worth knowing a year later. To go back to one:

    git checkout $TAG && cd src && ./make.bash

Tags are also why this is a merge rather than a rebase: a fast-forward keeps
every earlier tag reachable. Confirm with
`git merge-base --is-ancestor <tag> sparc64` if in doubt.

---

## 6. Release

Build the tarball from `/root/goport/go`, never from `/usr/lib/go`: the
ebuild's `src_install` strips every `testdata` directory, so a GOROOT taken
from there is missing the test corpus.

    TAG=sparc64-$(date +%Y%m%d)
    V=go-1.28.0_pre$(date +%Y%m%d)-linux-sparc64.tar.xz
    cd src && ./make.bash            # rebuild AFTER the last commit
    ../bin/go version                # must report the tagged head
    tar -C /root/goport --exclude=.git -I 'xz -T0' -cf /var/cache/distfiles/$V go

`go version` embeds the git HEAD at build time, so a tarball built before the
final commit reports the previous one while carrying the new code. Rebuilding
is cheap, about 6 minutes warm. Verify the artifact rather than the tree:

    rm -rf /tmp/relverify && mkdir -p /tmp/relverify
    tar -C /tmp/relverify -xf /var/cache/distfiles/$V
    /tmp/relverify/go/bin/go version                          # its own bin/go
    find /tmp/relverify/go/src -type d -name testdata | wc -l  # must be non-zero

Then publish. Notes stay short: the port's scope belongs in
`README.sparc64.md` and the changelog in the history, so one line per fix
plus a compare link is enough.

    gh release create $TAG --repo shalseth/go --title "$TAG" \
        --notes-file notes.md /var/cache/distfiles/$V

`gh release upload --clobber` replaces an asset that turns out wrong.

---

## 7. Install the toolchain

`dev-lang/go-9999` is a live ebuild pulling branch `sparc64` from the fork, so
it picks up whatever was just pushed. **Push before emerging.**

    emerge -v '=dev-lang/go-9999'
    go version        # must report the merged head, not the old one

It builds with `CGO_ENABLED=1` and bootstraps from the currently installed Go.

---

## 8. Rebuild Alloy

The end-to-end check, and the reason step 7 is not optional: Alloy is roughly
2500 packages and exercises far more of the toolchain than the suite's own
tests do.

    /root/alloybuild/alloy-sparc64/build.sh

Leave `GOROOT_SPARC64` unset. `build.sh` defaults it to `$(go env GOROOT)`,
which after step 7 is the toolchain just installed, and it prints which one it
picked. Confirm the binary really came from it rather than from a cached
build:

    go version -m <out> | head -1

About 20 minutes. A `josharian/native: unrecognized arch sparc64` line at
startup is expected and is not ours to fix: upstream deleted that code path in
2023 but has never tagged a release, so everyone still pins v1.1.0.

---

## Order that matters

Generator before `make.bash`, or the errors are noise. Push before emerging,
since the ebuild pulls from the remote. Rebuild before packaging, or the
tarball reports the wrong commit.

Emerge last, and always. A release can be tagged, published and verified
while the machine still runs a toolchain weeks old, because nothing in the
release path touches `/usr/lib/go`. The gap is silent until someone types
`go version`. Steps 7 and 8 are what close it, and they also mean the next
merge bootstraps from current code rather than from whatever was installed
months ago.
