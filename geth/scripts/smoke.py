#!/usr/bin/env python3
"""Database-free differential checks on normal and failing arithmetic inputs."""
import os
import subprocess
from common import ROOT

for args in ("9,7", f"{2**256 - 1},1"):
    subprocess.run([str(ROOT / "bin/hepha-smoke"), "-artifact",
                    str(ROOT / "template_miner/testdata/arithmetic-0824.json"),
                    "-templates", str(ROOT / "assets/arithmetic.json"),
                    "-method", "add", "-args", args], cwd=ROOT, check=True)
print("Smoke checks completed; see the output for return value, gas, state root and fallback counts.")
