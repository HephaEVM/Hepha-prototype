"""Shared, dependency-free experiment configuration and process management."""
from pathlib import Path
import hashlib
import json
import os
import platform
import shutil
import signal
import subprocess
import time

ROOT = Path(__file__).resolve().parents[1]
ANCHOR = 23400001
TAGS = "paper_evaluation,precompiled_regions,region_families"


def sha256(path):
    return hashlib.file_digest(Path(path).open("rb"), "sha256").hexdigest()


def environment(cache, procs=4, memory="32GiB"):
    env = dict(os.environ)
    prefixes = ("COLD_", "ADMISSION_", "ORACLE_", "BENCHMARK_", "EXECUTION_",
                "FINAL_", "REGION_", "PAPER_", "HASH_", "DEMAND_", "PREP_")
    for key in list(env):
        if key.startswith(prefixes) or key in ("READINESS_TARGETS", "CHECKED_COMPILED",
                                              "TOP20_SCOPE_FILE", "COST_ARMS"):
            del env[key]
    env.update(GOMAXPROCS=str(procs), GOMEMLIMIT=memory, GOGC="300",
               DEMAND_ROOT=str(ROOT), DEMAND_DIR=str(cache), DEMAND_MAX_SLOTS="262144",
               HISTORICAL_UNMODIFIED="1", PAPER_ABLATION="full", PAPER_CE_EXTEND="1",
               CATALOG_DISABLE="0", COST_STUDY="0", ORACLE_ARTIFACTS="")
    # Resolve automatic toolchain selection once. Runtime plugin compilers use
    # GOTOOLCHAIN=local and must not fall back to an older Go binary on PATH.
    goroot = subprocess.check_output(["go", "env", "GOROOT"], cwd=ROOT, text=True).strip()
    env.update(GOROOT=goroot, GOTOOLCHAIN="local",
               PATH=str(Path(goroot) / "bin") + os.pathsep + env.get("PATH", ""))
    for name, file in {
        "ABI_SCALAR_TEMPLATES": "abi-scalar.json",
        "ABI_EXPANDED_TEMPLATES": "abi-expanded.json",
        "HOT_REGION_CORPUS": "hot-region-corpus.json",
        "HOT_REGION_SELECTED": "hot-region-selected.json",
    }.items():
        env[name] = str(ROOT / "assets" / file)
    return env


def arguments(parser):
    parser.add_argument("--data", type=Path, required=True,
                        help="dataset containing chaindata/ and state/")
    parser.add_argument("--out", type=Path, required=True, help="new output directory")
    parser.add_argument("--procs", type=int, default=4)
    parser.add_argument("--memory", default="32GiB", help="Go soft memory limit")
    parser.add_argument("--dry-run", action="store_true")
    parser.add_argument("--timeout-hours", type=float, default=12)


def command(data, out, mode, blocks):
    return [str(ROOT / "bin" / "hepha-replay"), "-mode", mode,
            "-chain", str(data.resolve() / "chaindata"),
            "-snapshot", str(data.resolve() / "state"),
            "-start", str(ANCHOR), "-end", str(ANCHOR + blocks - 1),
            "-out", str(out), "-arithmetic-templates", str(ROOT / "assets/arithmetic.json")]


def run(cmd, env, out, data, dry_run=False, timeout_hours=12, **protocol):
    out = Path(out).resolve()
    out.mkdir(parents=True, exist_ok=False)
    cache = Path(env["DEMAND_DIR"])
    cache.mkdir(parents=True, exist_ok=False)
    retained = {k: v for k, v in env.items()
                if k.startswith(("COLD_", "BENCHMARK_", "PAPER_", "HASH_", "DEMAND_",
                                 "REGION_", "FINAL_", "PREP_", "EXECUTION_"))
                or k in ("GOMAXPROCS", "GOMEMLIMIT", "GOGC", "READINESS_TARGETS",
                         "HISTORICAL_UNMODIFIED")}
    plan = dict(protocol, command=cmd, environment=retained,
                host=platform.platform(), timing="EVM lat; preparation recorded separately",
                assets={str(p.relative_to(ROOT)): sha256(p) for p in (ROOT / "assets").glob("*.json")})
    if Path(cmd[0]).is_file():
        plan["binarySHA256"] = sha256(cmd[0])
    (out / "plan.json").write_text(json.dumps(plan, indent=2) + "\n")
    print(json.dumps({"output": str(out), "command": cmd}), flush=True)
    if dry_run:
        return
    if not Path(cmd[0]).is_file():
        raise SystemExit("Build first: python3 scripts/build.py")
    if not all((data / p).is_dir() for p in ("chaindata", "state")):
        raise SystemExit("--data must contain chaindata/ and state/; see docs/DATA.md")
    start = time.monotonic()
    with (out / "run.log").open("w") as log:
        child = subprocess.Popen(cmd, cwd=ROOT, env=env, stdout=log,
                                 stderr=subprocess.STDOUT, start_new_session=True)
        reason = None
        try:
            while child.poll() is None:
                time.sleep(2)
                if time.monotonic() - start > timeout_hours * 3600:
                    reason = "time limit reached"
                    break
                if shutil.disk_usage(cache).free < 4 * 1024**3:
                    reason = "less than 4 GiB free disk space"
                    break
        finally:
            if child.poll() is None:
                os.killpg(child.pid, signal.SIGTERM)
                try:
                    child.wait(timeout=20)
                except subprocess.TimeoutExpired:
                    os.killpg(child.pid, signal.SIGKILL)
                    child.wait()
        record = dict(exitCode=child.returncode, seconds=time.monotonic() - start,
                      stoppedReason=reason)
        (out / "exit.json").write_text(json.dumps(record, indent=2) + "\n")
    if child.returncode or reason:
        raise SystemExit(f"Run incomplete: {out / 'run.log'} ({reason or child.returncode})")
    print(f"Completed in {record['seconds']:.1f}s: {out}", flush=True)
