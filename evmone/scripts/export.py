#!/usr/bin/env python3
"""Build the offline exporter against the sibling Geth semantic analyzer."""
import argparse
import json
import os
from pathlib import Path
import subprocess
import sys
from common import ROOT, GETH, run, save, environment

p = argparse.ArgumentParser(description=__doc__)
p.add_argument("--fixtures", type=Path, required=True)
p.add_argument("--out", type=Path, required=True)
a = p.parse_args()
out = a.out.resolve(); out.mkdir(parents=True, exist_ok=False)
files = sorted(a.fixtures.glob("*.json"))
if not files:
    p.error("fixtures directory contains no JSON witnesses")
sites = {}
for path in files:
    record = json.loads(path.read_text())
    if "executed_pcs" not in record:
        p.error(f"{path.name}: missing untimed instruction census")
    for code, pcs in record["executed_pcs"].items():
        sites.setdefault(code, set()).update(pcs)
save(out / "sites.json", {k: sorted(v) for k, v in sites.items()})
overlay = {str(GETH / "core/vm" / f.name): str(f)
           for f in (ROOT / "analysis/vm").glob("*.go")}
save(out / "overlay.json", {"Replace": overlay})
env = environment(); env.update(GOMAXPROCS="4", GOMEMLIMIT="32GiB", GOGC="300")
env["GOROOT"] = subprocess.check_output(["go", "env", "GOROOT"], cwd=GETH, text=True).strip()
env["PATH"] = str(Path(env["GOROOT"]) / "bin") + os.pathsep + env["PATH"]
env["GOTOOLCHAIN"] = "local"
run(["go", "build", "-p", "4", "-overlay", out / "overlay.json", "-tags",
     "hepha_export,paper_evaluation,precompiled_regions,region_families", "-o", out / "export",
     ROOT / "analysis/export/main.go"], cwd=GETH, env=env)
for kind in ("full", "effects"):
    run([out / "export", "-fixtures", a.fixtures.resolve(), "-sites", out / "sites.json",
         "-out", out / (kind + ".json"), "-" + kind], cwd=GETH, env=env)
run([sys.executable, ROOT / "codegen/assemble.py", out / "full.json", out / "effects.json",
     out / "analysis.json"], env=env)
print("Region IR:", out / "analysis.json")
