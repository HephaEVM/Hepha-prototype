#!/usr/bin/env python3
"""Build the release binaries using the same Go toolchain as runtime plugins."""
from pathlib import Path
import argparse
import json
import os
import subprocess
from common import ROOT, TAGS, sha256

p = argparse.ArgumentParser(description=__doc__)
p.add_argument("--test", action="store_true", help="also run self-contained tests")
a = p.parse_args()
env = dict(os.environ, CGO_ENABLED="1")
goroot = subprocess.check_output(["go", "env", "GOROOT"], cwd=ROOT, text=True).strip()
env.update(GOROOT=goroot, GOTOOLCHAIN="local",
           PATH=str(Path(goroot) / "bin") + os.pathsep + env.get("PATH", ""))
(ROOT / "bin").mkdir(exist_ok=True)
for name, package in (("hepha-replay", "optimizer-audit"),
                      ("hepha-smoke", "arithmetic-runtime"),
                      ("template-miner", "template-miner")):
    subprocess.run(["go", "build", "-tags", TAGS, "-o", str(ROOT / "bin" / name),
                    "./cmd/" + package], cwd=ROOT, env=env, check=True)
if a.test:
    subprocess.run(["go", "test", "-tags", TAGS, "./..."], cwd=ROOT, env=env, check=True)
info = {"goVersion": subprocess.check_output(["go", "version"], text=True).strip(),
        "tags": TAGS, "CGO_ENABLED": "1",
        "binaries": {p.name: sha256(p) for p in (ROOT / "bin").iterdir() if p.is_file() and p.name != "build.json"}}
(ROOT / "bin/build.json").write_text(json.dumps(info, indent=2) + "\n")
print("Built Hepha. Try: python3 scripts/smoke.py")
