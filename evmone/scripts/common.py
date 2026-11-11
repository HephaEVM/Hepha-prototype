"""Paths and subprocess helpers for the standalone native backend."""
from pathlib import Path
import hashlib
import json
import os
import subprocess

ROOT = Path(__file__).resolve().parents[1]
GETH = ROOT.parent / "geth"
VENDOR = ROOT / "vendor"


def sha(path):
    with Path(path).open("rb") as source:
        return hashlib.file_digest(source, "sha256").hexdigest()


def save(path, value):
    Path(path).write_text(json.dumps(value, indent=2) + "\n")


def environment():
    env = {k: v for k, v in os.environ.items()
           if not k.startswith(("HEPHA_", "ICDE_", "PAPER_", "EVMONE_"))
           and k not in ("LD_PRELOAD", "DYLD_INSERT_LIBRARIES")}
    return env


def includes():
    upstream = VENDOR / "evmone"
    paths = [ROOT / "src", upstream, upstream / "include", upstream / "lib",
             upstream / "lib/evmone", upstream / "evmc/include",
             VENDOR / "intx/include", VENDOR / "json", VENDOR / "blst/bindings"]
    return ["-I" + str(p) for p in paths]


def run(command, **kwargs):
    words = list(map(str, command))
    print(" ".join(words) if len(words) < 20 else
          f"{words[0]}: {len(words)} arguments; output {words[-1]}", flush=True)
    subprocess.run(list(map(str, command)), check=True, **kwargs)
