# Merging upstream into the sparc64 port

Procedure for pulling golang/go master into the `sparc64` branch, validating
the result, and installing the toolchain.

The merge itself is text work, but the rule generator, the build and the test
suite all need a working sparc64 Go, so in practice every step runs on a
sparc64 host. Long jobs should be detached (`setsid`, `screen`) so they survive
a dropped connection.

Rebuilding downstream software against the new toolchain is out of scope here;
for Grafana Alloy see the `alloy-sparc64` repo.

Paths below are this setup's — adjust for another.

| path | what |
|---|---|
| `/root/goport/go` | port checkout, branch `sparc64` |
| `/root/gomerge` | throwaway worktree for the merge |

Rough costs on an UltraSPARC T4-1: `make.bash` ~12 min, `dist test -k` ~40 min,
`emerge go-9999` ~12 min.

**The port is expected to test clean** — 22 phases, zero failures. That is the
standard every merge is held to, so there is no need to capture a fresh
baseline first. Any failure is a real result: investigate it, do not explain it
away.

---

## 1. Merge onto a throwaway branch

Never merge onto `sparc64` directly — nothing is at risk until it is green.

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

The deltas are usually tiny — a single op in a list, one helper function —
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
fix these one build at a time — scan the whole file against what the other
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

Scan the **whole file**, not just rule conditions — `PanicBoundsC` appears on
the result side and a condition-only scan reports nothing.

**d. Helper signatures, not just names.** `logLargeCopy` was split into
`LogLargeCopy(string, XPos, int64)` and
`LogLargeCopyValue(*Value, int64) bool`. A name-based rename picks the wrong
one and gets as far as the type checker. When a rename target exists, check its
signature against the call site in another arch's rules.

**e. New syscall constants.** Upstream adds a syscall to `exec_linux.go` and
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
    ../bin/go tool dist test -k > /root/mergebuild-k.log 2>&1

Use `-k` (keep going). Without it `dist test` stops at the first failing
package and never reaches the twenty later phases — which is how an
internal-link cgo bug stayed hidden for six days, with every run reaching two
phases out of twenty-two.

    grep -cE '^--- FAIL' /root/mergebuild-k.log   # expect 0
    grep -c '^#####' /root/mergebuild-k.log       # expect 22 phases

Zero failures across 22 phases is the pass condition. The six `cmd/objdump` and
one `cmd/pprof` disassembly tests skip through `mustHaveDisasm`, there being no
sparc64 disassembler; everything else must pass.

**Every failure is an issue to investigate**, whether or not it looks related
to the merge. Two things worth knowing while investigating, neither of them an
excuse to move on:

* `time.TestLongAdjustTimers` is load-sensitive — it needs ~12 s against a hard
  60 s deadline and fails under heavy parallel load. Re-run it alone on an idle
  machine and produce that evidence before concluding it was contention.
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

Annotated, not lightweight — the tag carries which upstream commit it caught
up to, which is the thing worth knowing a year later. To go back to one:

    git checkout $TAG && cd src && ./make.bash

Tags are also why this is a merge rather than a rebase: a fast-forward keeps
every earlier tag reachable. Confirm with
`git merge-base --is-ancestor <tag> sparc64` if in doubt.

---

## 6. Install the toolchain

`dev-lang/go-9999` is a live ebuild pulling branch `sparc64` from the fork, so
it picks up whatever was just pushed. **Push before emerging.**

    emerge -v '=dev-lang/go-9999'
    go version        # must report the merged head, not the old one

It builds with `CGO_ENABLED=1` and bootstraps from the currently installed Go.

---

## Order that matters

Generator before `make.bash`, or the errors are noise. Push before emerging,
since the ebuild pulls from the remote.
