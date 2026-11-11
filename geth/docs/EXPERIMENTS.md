# Hepha: Semantic Region Execution for Geth

Hepha shortens EVM execution paths by replacing verified bytecode regions with
host-side computations. It preserves the region's observable behavior and
original gas accounting, and falls back to the interpreter when a region is
unavailable or its safety conditions do not hold.

This repository contains the **geth implementation** and two reproduction paths:
fully prepared fixed calls, and chronological replay starting with empty online
plan caches. It is a research execution artifact based on **go-ethereum v1.16.5**,
not a complete networked Ethereum client.

## Design

1. **Semantic summaries.** Offline mining tracks stack-value origins and control
   flow, verifies candidate regions, and normalizes their structure into reusable
   summaries. The checked summary libraries are included in `assets/`.
2. **Region execution.** Binding resolves a summary against actual runtime
   bytecode. Value passing avoids intermediate stack traffic; computation reuse
   removes repeated work; check elimination uses proven conditions to avoid
   redundant checks. The Full policy enables all three mechanisms.
3. **Progressive preparation.** Shared executors serve supported regions directly.
   Selective background preparation specializes other useful regions. Ready plans
   are published incrementally; execution uses the ordinary interpreter otherwise.
   Preparation is bounded and coordinated with foreground execution.

The baseline uses the independent native interpreter in
`core/vm/interpreter_native.go`, with the same latency instrumentation.

## Requirements

- Linux or macOS; **Go 1.24.0**, a C toolchain, and `CGO_ENABLED=1`.
- **Python 3.11+**; the experiment scripts use only the standard library.
- Internet access for the first Go dependency download.
- Historical replay additionally needs the external dataset described in
  [docs/DATA.md](docs/DATA.md).

The reference configuration uses four Go execution threads, a **32 GiB Go soft
memory limit**, and `GOGC=300`. The soft limit is not a hard process-memory cap.
Allow additional memory and disk space for state caches and selective compilation.
Dynamic plugins must be built by the same Go toolchain as the executable. Do not
reuse compiled caches across builds or machines.

## Build and verify

From the repository root:

```sh
python3 scripts/build.py --test
python3 scripts/smoke.py
```

This builds `bin/hepha-replay`, `bin/hepha-smoke`, and `bin/template-miner`.
The tests cover template mining, arithmetic boundaries, region execution and
preparation management. The database-free smoke checks compare normal addition
and overflow against the native interpreter, including output, remaining gas,
and state root. They are correctness checks, not steady-state performance results.

Generated executor sources are included; no historical build overlays or private
experiment directories are required. The original Go module path is retained so
that dynamically generated plugins import the same runtime ABI.

## Fully prepared fixed calls

The frozen selection contains **556 historical root calls**, classified using
`workloads/top20.json`. Each input includes the original transaction context and
pre-state captured during replay. Preparation converges before paired timing.

```sh
python3 scripts/fixed_calls.py \
  --data /path/to/mainnet-23400001-23401000 \
  --out results/fixed --rounds 8 --iterations 32

python3 scripts/summarize.py results/fixed --output results/fixed/summary.json
```

Both methods execute the original deployed bytecode and restore identical inputs
between calls. Method order rotates across rounds. The harness rejects output,
gas, state, or log mismatches and detects plan preparation during timed batches.
This experiment repeatedly executes fixed inputs; it is not continuous transaction
stream throughput. Its paired root-call harness reuses an EVM within each batch.

To use another frozen subset, supply `--targets calls.json`. Each record must have
`hash`, `block`, `index`, and `to`; do not select calls by observed speedup.

## Cold-start experiments

Both experiments preload state through native replay **before creating preparation
workers**. Online code-plan caches then start empty. This isolates execution and
preparation from disk-fetch variation; it does not represent cold storage.

### Continuous transaction stream

```sh
python3 scripts/cold_start.py --kind stream --blocks 1000 --repeats 3 \
  --data /path/to/mainnet-23400001-23401000 --out results/cold-stream

python3 scripts/summarize.py results/cold-stream \
  --output results/cold-stream/summary.json
```

Each method runs in a separate process, preserves the original transaction order,
and creates a new EVM for each transaction. Shared plan caches persist across
transactions. Complete replay order alternates between repetitions. Hepha receives
a default **10 ms preparation credit per block**, recorded separately from EVM
execution time. `--window-ms` changes this credit. The stream also enables the
retained direct-region dispatch path used by the reference stream configuration.

### Fixed calls during cold-to-warm preparation

```sh
python3 scripts/cold_start.py --kind fixed --blocks 1000 \
  --data /path/to/mainnet-23400001-23401000 --out results/cold-fixed

python3 scripts/summarize.py results/cold-fixed \
  --output results/cold-fixed/summary.json
```

This replays the chronological stream to train the preparer, saves immutable
ready-plan views, then measures the same fixed inputs against Vanilla at each
checkpoint. The probes cannot train the online preparer. Checkpoints are at
0–100 processed blocks in steps of 10, then in steps of 50 through 1,000, including
the final block. Both methods use fresh EVM instances for these probes.

The reference fixed-probe policy uses **1,000 ms preparation credit per block**.
It differs from the 10 ms stream policy: fixed probes diagnose plan readiness,
while stream timing includes foreground interference. Use the same `--window-ms`
when comparing these preparation policies directly.

For a quick replay check, use `--kind stream --blocks 1`. A short fixed-probe run
still preloads all blocks needed to capture its selected calls. Use a smaller
`--targets` file to avoid that cost. `--dry-run` records commands and settings
without starting replay. Existing output directories are never overwritten.

## Measurements and outputs

Every run records its command, configuration, asset checksums, binary checksum,
completion status, and logs. Principal files are:

| Experiment | Raw measurements | Validation / preparation |
| --- | --- | --- |
| Warm fixed calls | `observations.jsonl` | `warm-validation.json`, `manifest.json` |
| Continuous replay | `transaction-runs.csv`, `block-runs.csv` | `benchmark-manifest.json`, preparation/window records |
| Cold fixed probes | `checkpoint-*/same-call-readiness.jsonl` | `readiness-protocol.json`, checkpoint metadata |

The primary metric, `lat`, measures top-level EVM instruction-loop elapsed time.
Nested execution is included once. It includes dynamic gas calculation and
memory/state access within that loop, but excludes top-level frame setup,
transaction preprocessing, native data preload, and explicit preparation windows.
It is **not end-to-end block latency**. Preparation and total process time are
reported separately; include them when evaluating end-to-end cost.

The summarizer computes `100 × (1 − Hepha latency / Vanilla latency)` from latency
sums, using per-input medians for repeated fixed calls. It preserves each complete
stream repetition and produces disjoint block-window summaries. It does not adjust
measurements, select the fastest repetition, or synthesize missing samples. Zero
baseline time has no defined reduction percentage. Hardware and run configuration
can materially change the measured improvement.

## Repository layout

```text
core/vm/           Binding, summaries, region execution, native fallback
regionexec/        Shared executor ABI and selective plugin preparation
template_miner/    Discovery, normalization, verification, compiler fixtures
assets/            Frozen summary libraries and preparation priors
workloads/         Fixed-call identities, Top20 cohort, dataset metadata
cmd/               Replay, mining and differential-check commands
scripts/           Build, fixed-call, cold-start and summarization entry points
core/, trie*, ...  Local geth dependencies needed by execution and state replay
```

The priority profiles and precompiled catalogs are frozen development artifacts,
not evidence of performance on unseen contracts. Reproducing this configuration
uses those supplied priors. Claiming held-out generalization requires independent
training inputs and an independently selected evaluation workload.

To mine a small compiler fixture without a database:

```sh
bin/template-miner -artifact template_miner/testdata/arithmetic-0824.json \
  -scenarios template_miner/testdata/scenarios.json -out /tmp/templates.json
```

Large chain/state databases, compiled plugins, run outputs, comparison clients,
and figure-generation pipelines are not distributed in this source directory.
After a run has exited, its `compiled/` directory can be deleted to reclaim space;
keep raw observations, manifests and validation records.

## License and provenance

See [NOTICE](NOTICE), [COPYING](COPYING), and
[COPYING.LESSER](COPYING.LESSER). Upstream copyright and license notices remain in
the source. The extracted geth dependencies come from go-ethereum v1.16.5.
`docs/source-provenance.json` records hashes of the verified Full source inputs
before packaging edits; `MANIFEST.sha256` identifies the released source files.
