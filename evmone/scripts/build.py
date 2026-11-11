#!/usr/bin/env python3
"""Build both controls and the Full backend from vendored source."""
import argparse
from concurrent.futures import ThreadPoolExecutor
import hashlib
import json
import os
from pathlib import Path
import platform
import subprocess
import sys
from common import ROOT, VENDOR, includes, run, save, sha

p = argparse.ArgumentParser(description=__doc__)
p.add_argument("--analysis", type=Path, required=True, help="trusted exported region IR")
p.add_argument("--out", type=Path, default=ROOT / "build")
p.add_argument("--jobs", type=int, default=4)
a = p.parse_args()
if a.jobs < 1:
    p.error("jobs must be positive")
out = a.out.resolve(); out.mkdir(parents=True, exist_ok=True)
sys.path.insert(0, str(ROOT / "codegen"))
from codegen import generate
generated = generate(a.analysis, out / "generated")
upstream = VENDOR / "evmone"
blst = VENDOR / "blst"
with (out / "blst.log").open("w") as log:
    subprocess.run(["sh", "build.sh"], cwd=blst, stdout=log,
                   stderr=subprocess.STDOUT, check=True)
core = ["advanced_analysis", "advanced_execution", "advanced_instructions",
        "baseline_analysis", "baseline_execution", "baseline_instruction_table",
        "create_address", "delegation", "instructions_calls", "instructions_storage",
        "tracing", "vm"]
sources = [upstream / "lib/evmone" / (n + ".cpp") for n in core]
sources += sorted((upstream / "lib/evmone_precompiles").glob("*.cpp"))
sources += [upstream / "lib/evmone_precompiles/pairing/bn254/pairing.cpp",
            upstream / "lib/evmone_precompiles/keccak.c"]
host = ["host", "state", "precompiles", "block", "transaction", "authorization",
        "bloom_filter", "system_contracts", "rlp_decode", "bal", "requests", "ethash_difficulty"]
sources += [upstream / "test/state" / (n + ".cpp") for n in host]
sources += [ROOT / "src" / n for n in ("regional_execution.cpp", "execute.cpp", "cache.cpp")]
sources += generated
headers = sorted(p for base in (ROOT / "src", VENDOR) for p in base.rglob("*")
                 if p.suffix in (".h", ".hpp") and p.is_file())
header_hash = hashlib.sha256(b"".join(p.read_bytes() for p in headers)).hexdigest()
compiler = subprocess.check_output(["clang++", "--version"], text=True)


def compile(source):
    c = source.suffix == ".c"
    cmd = [os.environ.get("CC", "clang") if c else os.environ.get("CXX", "clang++"),
           "-O3", "-DNDEBUG", "-fPIC", "-ffunction-sections", "-fdata-sections"]
    if not c:
        cmd += ["-std=c++20", '-DPROJECT_VERSION="0.25.0"']
    cmd += includes()
    key = hashlib.sha256(source.read_bytes() + header_hash.encode() +
                         compiler.encode() + json.dumps(cmd).encode()).hexdigest()
    obj = out / "objects" / (key + ".o"); obj.parent.mkdir(exist_ok=True)
    if not obj.exists():
        with obj.with_suffix(".log").open("w") as log:
            result = subprocess.run(cmd + ["-c", str(source), "-o", str(obj)],
                                    stdout=log, stderr=subprocess.STDOUT)
        if result.returncode:
            obj.unlink(missing_ok=True)
            raise RuntimeError(obj.with_suffix(".log").read_text())
    return obj


with ThreadPoolExecutor(max_workers=a.jobs) as pool:
    objects = list(pool.map(compile, sources))
    mains = list(pool.map(compile, [ROOT / "src/fixed_calls.cpp", ROOT / "src/cold_stream.cpp"]))
system = platform.system()
for name, main in zip(("fixed-calls", "cold-stream"), mains):
    link = ["clang++", "-O3", "-Wl,-dead_strip" if system == "Darwin" else "-Wl,--gc-sections",
            str(main), *map(str, objects), str(blst / "libblst.a"), "-o", str(out / name)]
    run(link)
save(out / "build.json", dict(compiler=compiler, platform=platform.platform(),
     analysisSHA256=sha(a.analysis), plansSHA256=sha(out / "generated/plans.json"),
     sourceSHA256={str(p.relative_to(ROOT)) if p.is_relative_to(ROOT) else str(p): sha(p)
                   for p in sources + headers}, executables={n: sha(out / n) for n in ("fixed-calls", "cold-stream")}))
print("Built", out, flush=True)
