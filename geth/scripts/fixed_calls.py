#!/usr/bin/env python3
"""Measure paired, fully prepared historical root calls against Vanilla geth."""
import argparse
import json
from pathlib import Path
from common import ROOT, ANCHOR, arguments, command, environment, run, sha256

p = argparse.ArgumentParser(description=__doc__)
arguments(p)
p.add_argument("--targets", type=Path, default=ROOT / "workloads/fixed-calls.json")
p.add_argument("--rounds", type=int, default=8)
p.add_argument("--iterations", type=int, default=32)
p.add_argument("--order-offset", type=int, default=0, help="rotate the paired method order")
a = p.parse_args()
rows = json.loads(a.targets.read_text())
if not rows or min(a.rounds, a.iterations, a.procs) < 1:
    p.error("targets and positive rounds/iterations/procs are required")
blocks = max(r["block"] for r in rows) - ANCHOR + 1
if not 1 <= blocks <= 1000 or min(r["block"] for r in rows) < ANCHOR:
    p.error("targets are outside the bundled dataset range")
out = a.out.resolve()
env = environment(out / "compiled", a.procs, a.memory)
env.update(PAPER_WARM="1", HASH_TARGETS=str(a.targets.resolve()),
           HASH_ROUNDS=str(a.rounds), HASH_ITERATIONS=str(a.iterations),
           HASH_ORDER_OFFSET=str(a.order_offset))
cmd = command(a.data, out, "fixed", blocks) + ["-benchmark-procs", str(a.procs)]
run(cmd, env, out, a.data, a.dry_run, a.timeout_hours,
    experiment="warm-fixed-calls", targetsSHA256=sha256(a.targets),
    originalBytecode=True, rounds=a.rounds, iterations=a.iterations,
    scope="Fixed original root calls with identical pre-state; preparation completed before timing.")
