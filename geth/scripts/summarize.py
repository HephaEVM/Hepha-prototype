#!/usr/bin/env python3
"""Summarize measured latency only; retain every complete paired repetition."""
import argparse
from collections import defaultdict
import csv
import json
from pathlib import Path
import statistics
from common import ROOT, ANCHOR


def paired(records, group_for):
    """Median per input/method, then ratio of cohort latency sums."""
    by_input = defaultdict(lambda: defaultdict(list))
    for identity, method, ns in records:
        if ns < 0:
            raise ValueError("negative latency")
        by_input[identity][method].append(ns)
    sums = {k: [0., 0., 0] for k in ("top20", "non-top20", "all")}
    for identity, times in by_input.items():
        if set(times) != {"vanilla", "hepha"}:
            raise ValueError(f"missing paired measurement: {identity}")
        vanilla, hepha = (statistics.median(times[m]) for m in ("vanilla", "hepha"))
        for group in (group_for(identity), "all"):
            sums[group][0] += vanilla
            sums[group][1] += hepha
            sums[group][2] += 1
    return {g: {"inputs": n, "vanillaTotalNS": v, "hephaTotalNS": h,
                "latencyReductionPct": 100 * (1 - h / v) if v else None}
            for g, (v, h, n) in sums.items()}


p = argparse.ArgumentParser(description=__doc__)
p.add_argument("result", type=Path)
p.add_argument("--targets", type=Path, default=ROOT / "workloads/fixed-calls.json")
p.add_argument("--top20", type=Path, default=ROOT / "workloads/top20.json")
p.add_argument("--output", type=Path, help="optional summary JSON path")
a = p.parse_args()
targets = {r["hash"].lower(): r for r in json.loads(a.targets.read_text())}
top = set(x.lower() for x in json.loads(a.top20.read_text()))


def group(identity):
    r = targets[identity]
    return "top20" if r["to"].lower() in top else "non-top20"


obs = a.result / "observations.jsonl"
csvs = sorted(a.result.glob("repeat-*/vanilla/transaction-runs.csv"))
if obs.exists():
    status = a.result / "exit.json"
    validation = a.result / "warm-validation.json"
    if (not status.exists() or json.loads(status.read_text())["exitCode"] or
            not validation.exists() or json.loads(validation.read_text())["status"] != "complete"):
        raise SystemExit("Warm run is incomplete or unvalidated; raw data is retained for diagnosis.")
    records = []
    for line in obs.open():
        r = json.loads(line)
        if r.get("phase") == "paper-timed":
            method = "vanilla" if r["mode"] == "vanilla-native" else "hepha"
            records.append((r["hash"].lower(), method, r["lat_ns"]))
    if not records:
        raise SystemExit("No timed records; confirm that the run completed.")
    result = {"scope": "warm fixed calls", "groups": paired(records, group)}
elif csvs:
    repetitions = []
    for native in csvs:
        other = native.parent.parent / "hepha/transaction-runs.csv"
        if not other.exists():
            raise SystemExit(f"Missing paired replay: {other}")
        records, destinations, blocks = [], {}, {}
        for method, file in (("vanilla", native), ("hepha", other)):
            exit_file = file.parent / "exit.json"
            if not exit_file.exists() or json.loads(exit_file.read_text())["exitCode"]:
                raise SystemExit(f"Incomplete replay: {file.parent}")
            for r in csv.DictReader(file.open()):
                h = r["tx_hash"].lower()
                records.append((h, method, int(r["loop_ns"])))
                destinations[h], blocks[h] = r["to"].lower(), int(r["block"]) - ANCHOR + 1
        classify = lambda h: "top20" if destinations[h] in top else "non-top20"
        points, previous = [], 0
        last = max(blocks.values())
        ends = sorted(set([n for n in range(10, 101, 10) if n <= last] +
                          [n for n in range(150, 1001, 50) if n <= last] + [last]))
        for end in ends:
            subset = [r for r in records if previous < blocks[r[0]] <= end]
            points.append({"fromBlock": previous + 1, "toBlock": end,
                           "groups": paired(subset, classify)})
            previous = end
        repetitions.append({"repeat": native.parent.parent.name,
                            "groups": paired(records, classify), "intervals": points})
    result = {"scope": "chronological stream; intervals are disjoint, not cumulative",
              "repetitions": repetitions}
else:
    files = sorted(a.result.glob("**/same-call-readiness.jsonl"))
    if not files:
        raise SystemExit("No supported measurement files found.")
    points = []
    for file in files:
        parent = file.parent
        while parent != parent.parent and not (parent / "exit.json").exists():
            parent = parent.parent
        status = parent / "exit.json"
        if not status.exists() or json.loads(status.read_text())["exitCode"]:
            raise SystemExit(f"Incomplete cold replay for {file}")
        records, checkpoint = [], None
        for line in file.open():
            r = json.loads(line)
            method = "vanilla" if r["condition"] == "native" else "hepha"
            if method == "hepha":
                checkpoint = r["checkpointBlock"]
            records.append((r["hash"].lower(), method, r["latNS"]))
        points.append({"file": str(file), "processedBlocks": checkpoint,
                       "groups": paired(records, group)})
    result = {"scope": "fixed calls on immutable cold-start plan checkpoints",
              "points": points}
text = json.dumps(result, indent=2) + "\n"
if a.output:
    a.output.write_text(text)
else:
    print(text, end="")
