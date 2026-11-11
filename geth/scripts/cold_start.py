#!/usr/bin/env python3
"""Replay original blocks with empty plans, optionally probing frozen fixed calls."""
import argparse
import json
from pathlib import Path
from common import ROOT, ANCHOR, arguments, command, environment, run, sha256

p = argparse.ArgumentParser(description=__doc__)
arguments(p)
p.add_argument("--kind", choices=("stream", "fixed"), default="stream")
p.add_argument("--blocks", type=int, default=1000)
p.add_argument("--repeats", type=int, default=1, help="independent complete replays")
p.add_argument("--window-ms", type=int, help="per-block preparation credit: stream 10, fixed 1000")
p.add_argument("--targets", type=Path, default=ROOT / "workloads/fixed-calls.json")
p.add_argument("--preload-blocks", type=int, help="native data preload; at least the measured range")
a = p.parse_args()
if not 1 <= a.blocks <= 1000 or min(a.repeats, a.procs) < 1:
    p.error("blocks must be in [1,1000]; repeats and procs must be positive")
window = a.window_ms if a.window_ms is not None else (10 if a.kind == "stream" else 1000)
if not 1 <= window <= 10000:
    p.error("window-ms must be in [1,10000]")
targets = json.loads(a.targets.read_text()) if a.kind == "fixed" else []
preload = a.preload_blocks or max([a.blocks] + [r["block"] - ANCHOR + 1 for r in targets])
if not a.blocks <= preload <= 1000:
    p.error("preload-blocks must cover execution and fit within the dataset")
if targets and (min(r["block"] for r in targets) < ANCHOR or
                max(r["block"] for r in targets) >= ANCHOR + preload):
    p.error("fixed targets must be inside the native preload range")
policy = json.loads((ROOT / "config/preparation.json").read_text())
for repeat in range(a.repeats):
    methods = ["vanilla", "hepha"] if a.kind == "stream" else ["hepha"]
    if repeat % 2:
        methods.reverse()
    for method in methods:
        out = a.out.resolve() / f"repeat-{repeat}" / method
        env = environment(out / "compiled", a.procs, a.memory)
        active = method == "hepha"
        on = "1" if active else "0"
        env.update(
            PAPER_STREAM="1", PAPER_NO_PROBES="1" if a.kind == "stream" else "0",
            BENCHMARK_FRESH_TX_EVM="1", COLD_FIXED_REUSE="0",
            BENCHMARK_PRELOAD_GC="1", BENCHMARK_PRELOAD_END=str(ANCHOR + preload - 1),
            BENCHMARK_MEASURE_FROM=str(ANCHOR), REGION_POLICY="new",
            PREP_EXPERIMENT_POLICY="priority", COLD_EXECUTION_COUNTERS="off",
            COLD_PUBLISH_DEMAND="on", COLD_WORK_POLICY="all", COLD_COOPERATIVE_PREPARATION="0",
            COLD_FOCUSED_PREPARATION="0", COLD_FAIR_STAGE="0",
            COLD_HOTSET=on, COLD_HOTSET_GUARD="1",
            COLD_HOTSET_FILE=str(ROOT / "assets/historical-top100.json"),
            COLD_VALUE_GATE="1", COLD_VALUE_POLICY=on,
            COLD_VALUE_PROFILE=str(ROOT / "assets/benefit-profile.json"),
            COLD_VALUE_PREPARATION="background", COLD_VALUE_BUDGET_US="2000",
            COLD_VALUE_COMPILER="selective" if active else "none",
            COLD_HOTCALL_POLICY=on, COLD_SELECTIVE_DELIVERY=on, COLD_TRADE=on,
            COLD_PREPARATION_DUTY="25", COLD_ADMISSION_THRESHOLD="1",
            COLD_PREPARATION_MIN_HITS="1", COLD_ENTRY_MIN_HITS="1",
            COLD_REGION_PRIORITY=on, COLD_REGION_PROFILE=str(ROOT / "assets/region-profile.json"),
            COLD_WINDOW_COMPILER="1", FINAL_BACKGROUND="0", COLD_WINDOW_WARM="0",
            EXECUTION_COST_STUDY="0", EXECUTION_AVAILABLE_FIRST="0", EXECUTION_PROFIT_GATE="0",
            ADMISSION_PRIMARY="0", ADMISSION_BOOTSTRAP="0",
        )
        if active:
            env.update({"COLD_TRADE_" + k: str(v) for k, v in policy.items()})
            env["COLD_TRADE_WINDOW_US"] = str(window * 1000)
        if a.kind == "stream" and active:
            env["BENCHMARK_STREAM_REGIONS"] = "1"
        if a.kind == "fixed":
            env.update(READINESS_TARGETS=str(a.targets.resolve()), PAPER_FIXED_DENSE="1")
        cmd = command(a.data, out, "cold", a.blocks) + [
            "-benchmark-modes", "sync" if active else "baseline",
            "-benchmark-repeats", "1", "-benchmark-procs", str(a.procs),
            "-benchmark-loop-timing=true"]
        run(cmd, env, out, a.data, a.dry_run, a.timeout_hours,
            experiment="cold-" + a.kind, method=method, repeat=repeat, blocks=a.blocks,
            preloadBlocks=preload, windowMS=window if active else 0,
            targetsSHA256=sha256(a.targets) if targets else None,
            scope="Chronological replay with empty online plans. Native preload warms state only. "
                  "Fixed probes use frozen ready-plan snapshots and cannot train the preparer.")
