# Hepha for Geth

Verified EVM regions execute directly in Go, with original gas and interpreter
fallback. Based on go-ethereum v1.16.5; this is an execution research artifact.

**Requirements:** Go 1.24.0, Python 3.11+, and a C compiler (Linux or macOS).
The reference runs use four Go threads and a 32 GiB Go soft memory limit.

```sh
python3 scripts/build.py --test
python3 scripts/smoke.py

# Fully prepared historical calls.
python3 scripts/fixed_calls.py --data /path/to/mainnet-23400001-23401000 --out results/fixed
python3 scripts/summarize.py results/fixed --output results/fixed/summary.json

# Chronological replay from empty plan caches.
python3 scripts/cold_start.py --kind stream --blocks 1000 --repeats 3 \
  --data /path/to/mainnet-23400001-23401000 --out results/stream
python3 scripts/summarize.py results/stream --output results/stream/summary.json

# Probe the same calls at cold-to-warm checkpoints.
python3 scripts/cold_start.py --kind fixed --blocks 1000 \
  --data /path/to/mainnet-23400001-23401000 --out results/cold-fixed
```

Use a new output directory per experiment. For a quick check, replay one block.
Scripts write fresh measurements and validation records; they do not render plots.
Both methods use original deployed bytecode. EVM `lat` excludes data preload and
explicit preparation windows; preparation time is recorded separately.

- [Dataset and checksums](docs/DATA.md)
- [Measurement details](docs/EXPERIMENTS.md)
- [Validation](docs/VALIDATION.md)
- Core implementation: `core/vm/`; offline miner: `template_miner/`.

Upstream license notices are retained in `COPYING`, `COPYING.LESSER`, and `NOTICE`.
