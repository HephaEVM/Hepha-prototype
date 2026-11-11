#!/usr/bin/env python3
"""Export validated native witnesses without changing the regular Geth build."""
import argparse
import json
import os
from pathlib import Path
import subprocess
import sys
from common import ROOT, GETH, run, save

p = argparse.ArgumentParser(description=__doc__)
p.add_argument("--kind", choices=("fixed", "stream"), default="fixed")
p.add_argument("--data", type=Path, required=True)
p.add_argument("--out", type=Path, required=True)
p.add_argument("--targets", type=Path, default=GETH / "workloads/fixed-calls.json")
p.add_argument("--blocks", type=int, default=1000)
a = p.parse_args()
if not 1 <= a.blocks <= 1000: p.error("blocks must be in [1,1000]")
out = a.out.resolve(); out.mkdir(parents=True, exist_ok=False)
build = out / "capture-build"; build.mkdir()
mapping = {str(GETH / "cmd/optimizer-audit" / f.name): str(f)
           for f in (ROOT / "analysis/capture").glob("*.go")}


def patched(relative, old, new):
    path = GETH / relative
    text = Path(mapping.get(str(path), path)).read_text()
    if text.count(old) != 1:
        raise RuntimeError("capture overlay no longer matches " + relative)
    target = build / path.name
    target.write_text(text.replace(old, new))
    mapping[str(path)] = str(target)


# Fixed export ends before Hepha warm preparation. The original measurement
# function stays typechecked but is unreachable in this dedicated binary.
patched("cmd/optimizer-audit/hash_ablation.go", "\tmodifiedExecuted := false",
        '\tif os.Getenv("EVMONE_NATIVE_EXPORT") != "" { return exportNativeHostFixture(block, headers, tx, fixture) }\n\tmodifiedExecuted := false')
patched("cmd/optimizer-audit/paper_warm.go",
        "func runPaperWarmCases(out string, enc *json.Encoder, rounds int) error {",
        'func runPaperWarmCases(out string, enc *json.Encoder, rounds int) error {\n if os.Getenv("EVMONE_NATIVE_EXPORT") != "" { return nil }')
patched("core/vm/interpreter.go", "type Config struct {",
        "type Config struct {\n ExternalVM interface { RunFrame(*EVM, *Contract, []byte, bool) ([]byte, error, time.Duration) }")
patched("core/vm/interpreter.go", 'import (', 'import (\n "time"')
patched("core/vm/interpreter_dispatch.go", "return evm.runSelected(contract, input, readOnly)",
        "if b := evm.Config.ExternalVM; b != nil { return b.RunFrame(evm, contract, input, readOnly) }; return evm.runSelected(contract, input, readOnly)")
patched("cmd/optimizer-audit/benchmark.go", "warmConfigs[i] = makeConfig(armEngine, m)",
        'warmConfigs[i] = makeConfig(armEngine, m)\n if path := os.Getenv("EVMONE_STREAM_CAPTURE"); path != "" { c, er := newNativeStreamCapture(path); if er != nil { return er }; defer c.close(); warmConfigs[i].ExternalVM = c }')
patched("cmd/optimizer-audit/benchmark.go", "Before: func(index int, tx *types.Transaction) {",
        "Before: func(index int, tx *types.Transaction) {\n if c, ok := configs[i].ExternalVM.(*nativeStreamCapture); ok { c.tx=tx; c.index=index; c.block=number; c.recorded=false }")
patched("cmd/optimizer-audit/benchmark.go", "After: func(index int, tx *types.Transaction, receipt *types.Receipt, elapsed time.Duration) {",
        'After: func(index int, tx *types.Transaction, receipt *types.Receipt, elapsed time.Duration) {\n if c, ok := configs[i].ExternalVM.(*nativeStreamCapture); ok { if !c.recorded { c.write(map[string]any{"frame":false}) }; c.tx=nil }')
patched("cmd/optimizer-audit/benchmark.go", "untimedCapture := false", "untimedCapture := true")
save(build / "overlay.json", {"Replace": mapping})
sys.path.insert(0, str(GETH / "scripts"))
# Load the Geth environment helper under a distinct module name.
import importlib.util
spec = importlib.util.spec_from_file_location("geth_common", GETH / "scripts/common.py")
config = importlib.util.module_from_spec(spec); spec.loader.exec_module(config)
env = config.environment(out / "compiled")
env["GOCACHE"] = os.environ.get("GOCACHE", str(build / "go-cache"))
binary = build / "capture"
run(["go", "build", "-p", "4", "-overlay", build / "overlay.json", "-tags",
     "hepha_export,paper_evaluation,precompiled_regions,region_families", "-o", binary,
     "./cmd/optimizer-audit"], cwd=GETH, env=env)
if a.kind == "fixed":
    rows = json.loads(a.targets.read_text())
    blocks = max(r["block"] for r in rows) - config.ANCHOR + 1
    if not 1 <= blocks <= 1000: p.error("targets are outside the dataset")
    env.update(EVMONE_NATIVE_EXPORT=str(out / "fixtures"), PAPER_WARM="1",
               HASH_TARGETS=str(a.targets.resolve()), HASH_ROUNDS="1", HASH_ITERATIONS="1")
    command = config.command(a.data, out / "replay", "fixed", blocks)
else:
    env.update(EVMONE_STREAM_CAPTURE=str(out / "stream.jsonl.gz"), PAPER_STREAM="1",
               PAPER_NO_PROBES="1", BENCHMARK_FRESH_TX_EVM="1",
               BENCHMARK_PRELOAD_END=str(config.ANCHOR + a.blocks - 1))
    command = config.command(a.data, out / "replay", "cold", a.blocks)
    command += ["-benchmark-modes", "baseline", "-benchmark-repeats", "1"]
command[0] = str(binary)
save(out / "protocol.json", dict(kind=a.kind, command=command,
     originalBytecode=True, scope="Untimed Vanilla reference capture; block root/receipt/gas validation."))
with (out / "capture.log").open("w") as log:
    subprocess.run(command, cwd=GETH, env=env, stdout=log, stderr=subprocess.STDOUT, check=True)
save(out / "complete.json", dict(complete=True, kind=a.kind))
print("Captured:", out)
