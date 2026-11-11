#!/usr/bin/env python3
"""Paired Vanilla/Hepha measurements of prepared historical calls."""
import argparse
import json
from pathlib import Path
import statistics
import subprocess
import time
from common import ROOT, GETH, environment, save, sha

p = argparse.ArgumentParser(description=__doc__)
p.add_argument("--fixtures", type=Path, required=True)
p.add_argument("--targets", type=Path, default=GETH / "workloads/fixed-calls.json")
p.add_argument("--build", type=Path, default=ROOT / "build")
p.add_argument("--out", type=Path, required=True)
p.add_argument("--rounds", type=int, default=8)
p.add_argument("--iterations", type=int, default=32)
a = p.parse_args()
if min(a.rounds, a.iterations) < 1: p.error("rounds and iterations must be positive")
files = [a.fixtures.resolve() / (x["hash"] + ".json") for x in json.loads(a.targets.read_text())]
if not files or not all(p.is_file() for p in files): p.error("missing target fixtures")
out = a.out.resolve(); out.mkdir(parents=True, exist_ok=False)
binary = a.build.resolve() / "fixed-calls"; plans = a.build.resolve() / "generated/plans.json"
env = environment(); env["ICDE_REGIONS"] = str(plans)
save(out / "protocol.json", dict(rounds=a.rounds, iterations=a.iterations,
     binarySHA256=sha(binary), plansSHA256=sha(plans), targetsSHA256=sha(a.targets),
     fixtures={p.name: sha(p) for p in files},
     scope="Native C++ state-witness host; paired root execution with offline preparation and restored input state."))
start = time.monotonic(); completed = 0
with (out / "observations.jsonl").open("w") as dest, (out / "run.log").open("w") as log:
    for path in files:
        subprocess.run([str(binary), str(path), str(a.rounds), str(a.iterations), "baseline,stage3"],
                       env=env, stdout=dest, stderr=log, check=True)
        completed += 1
save(out / "complete.json", dict(calls=completed, seconds=time.monotonic()-start))
top = {k.lower() for k in json.loads((GETH / "workloads/top20.json").read_text())}
targets = {x["hash"]: x for x in json.loads(a.targets.read_text())}
samples = {}
for text in (out / "observations.jsonl").read_text().splitlines():
    r = json.loads(text); samples.setdefault((r["hash"], r["mode"]), []).append(r["frame_ns"])
groups = {n: [0., 0.] for n in ("Top20", "Non-Top20", "All")}
for h, t in targets.items():
    times = [statistics.median(samples[(h, m)]) for m in ("baseline", "stage3")]
    for group in ("Top20" if t["to"].lower() in top else "Non-Top20", "All"):
        for i in range(2): groups[group][i] += times[i]
save(out / "summary.json", {k: dict(vanillaNS=v[0], hephaNS=v[1],
     reductionPct=100*(1-v[1]/v[0]) if v[0] else None) for k, v in groups.items()})
print(out / "summary.json")
