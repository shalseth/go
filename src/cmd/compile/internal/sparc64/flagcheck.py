#!/usr/bin/env python3
"""Report sparc64 ops that write the condition codes without saying so.

flagalloc keys off ClobberFlags to know where a Flags value dies. An op whose
expansion emits a CC-writing instruction but does not declare clobberFlags
leaves the allocator believing a comparison result survives across it, and the
consumer then branches on codes the op destroyed. That miscompiles silently and
shows up far away - as heap corruption, or a nil dereference inside the GC.

This has now happened twice: the atomics in 2026-08 and ADDCARRY/SUBBORROW in
2026-09. Run it after adding or changing any multi-instruction op.

    python3 src/cmd/compile/internal/sparc64/flagcheck.py

Exits non-zero if anything is unaccounted for.
"""
import os
import re
import sys

HERE = os.path.dirname(os.path.abspath(__file__))
ROOT = os.path.normpath(os.path.join(HERE, '..', '..', '..', '..', '..'))
SSA = os.path.join(HERE, 'ssa.go')
OPS = os.path.join(ROOT, 'src/cmd/compile/internal/ssa/_gen/SPARC64Ops.go')

# Arithmetic and logical forms ending in CC write the integer condition codes,
# as does CMP. MOV<cond> also ends in CC (MOVCC = move if carry clear) but
# *reads* them, so anything starting with MOV is excluded. FCMP writes fcc.
CC_WRITER = re.compile(r'sparc64\.A(?!MOV)(\w+CC|CMP\w*|FCMP\w*)\b')


def declared_ops(text):
    """Ops declaring clobberFlags, handling both one-line and multi-line forms."""
    out = set()
    for m in re.finditer(r'name:\s*"(\w+)"', text):
        # scan to the end of this op's literal, whichever form it takes
        tail = text[m.end():m.end() + 600]
        end = tail.find('name:')
        if end != -1:
            tail = tail[:end]
        if 'clobberFlags' in tail:
            out.add(m.group(1))
    return out


def main():
    ssa = open(SSA).read()
    declared = declared_ops(open(OPS).read())

    cur = []
    problems = []
    for i, line in enumerate(ssa.split('\n'), 1):
        m = re.match(r'\s*case ((?:ssaop\.OpSPARC64\w+[,\s]*)+):', line)
        if m:
            cur = re.findall(r'ssaop\.OpSPARC64(\w+)', m.group(1))
        hit = CC_WRITER.search(line)
        if hit and cur:
            missing = [n for n in cur if n not in declared]
            if missing:
                problems.append((i, line.strip()[:50], missing))

    if not problems:
        print('ok: every sparc64 op that writes condition codes declares clobberFlags')
        return 0
    print('sparc64 ops writing condition codes without clobberFlags:\n')
    for line_no, text, names in problems:
        print('  ssa.go:%d  %s' % (line_no, text))
        print('      ops: %s\n' % ', '.join(names))
    return 1


if __name__ == '__main__':
    sys.exit(main())
