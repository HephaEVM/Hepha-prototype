#!/usr/bin/env python3
"""Measure chronological witness replay from empty code and binding caches."""
import argparse
import gzip
import json
from pathlib import Path
import subprocess
import time
from common import ROOT, GETH, environment, save, sha

p = argparse.ArgumentParser(description=__doc__)
p.add_argument("--witnesses", type=Path, required=True)
p.add_argument("--build", type=Path, default=ROOT / "build")
p.add_argument("--out", type=Path, required=True)
p.add_argument("--repeats", type=int, default=3)
a = p.parse_args()
if a.repeats < 1: p.error("repeats must be positive")
out = a.out.resolve(); out.mkdir(parents=True, exist_ok=False)
binary = a.build.resolve() / "cold-stream"; plans = a.build.resolve() / "generated/plans.json"
env = environment(); env["ICDE_REGIONS"] = str(plans)
save(out / "protocol.json", dict(repeats=a.repeats, binarySHA256=sha(binary),
     plansSHA256=sha(plans), witnessesSHA256=sha(a.witnesses),
     scope="Native C++ witness host; original order; empty online caches; reused VM in both arms; precompiled executors, observed-code binding between blocks with 10 ms credit. No online C++ compilation. Binding is separate from root latency."))
for repeat in range(a.repeats):
    methods = [("vanilla", 0), ("hepha", 12)]
    if repeat % 2: methods.reverse()
    for name, mode in methods:
        directory = out / f"repeat-{repeat}" / name; directory.mkdir(parents=True)
        start = time.monotonic()
        with (directory / "observations.jsonl").open("w") as dest, (directory / "run.log").open("w") as log:
            child = subprocess.Popen([str(binary), str(mode)], stdin=subprocess.PIPE,
                                     stdout=dest, stderr=log, env=env)
            try:
                opener = gzip.open if a.witnesses.suffix == ".gz" else open
                with opener(a.witnesses, "rb") as source:
                    for block in iter(lambda: source.read(1 << 20), b""):
                        child.stdin.write(block)
                child.stdin.close()
                code = child.wait()
            except BaseException:
                child.kill(); child.wait(); raise
        if code: raise SystemExit(f"Replay failed: {directory / 'run.log'}")
        record = json.loads((directory / "run.log").read_text().splitlines()[-1])
        if not record.get("complete"): raise SystemExit("Replay incomplete")
        record["wallSeconds"] = time.monotonic() - start
        save(directory / "complete.json", record)
        print(name, repeat, record["transactions"], "transactions", flush=True)

top = {k.lower() for k in json.loads((GETH / "workloads/top20.json").read_text())}
summaries = []
for repeat in range(a.repeats):
    folder = out / f"repeat-{repeat}"
    totals = {n: [0, 0, 0] for n in ("Top20", "Non-Top20", "All")}
    with (folder / "vanilla/observations.jsonl").open() as va, (folder / "hepha/observations.jsonl").open() as he:
        for left, right in zip(va, he, strict=True):
            v, h = json.loads(left), json.loads(right)
            if any(v[k] != h[k] for k in ("hash", "block", "index", "to")):
                raise RuntimeError("replay input identity/order mismatch")
            for group in ("Top20" if v["to"].lower() in top else "Non-Top20", "All"):
                totals[group][0] += v["ns"]; totals[group][1] += h["ns"]; totals[group][2] += 1
    summaries.append(dict(repeat=repeat, groups={k: dict(vanillaNS=v[0], hephaNS=v[1],
        transactions=v[2], reductionPct=100*(1-v[1]/v[0]) if v[0] else None)
        for k, v in totals.items()}))
save(out / "summary.json", summaries)
