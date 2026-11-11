# Hepha for evmone

Native C++ execution of verified semantic regions, based on **evmone v0.25.0**.
The offline exporter shares Geth's analyzer; timed execution uses the native
C++ state host without Go callbacks. Original gas, outputs and state effects
are checked against reference witnesses.

**Requirements:** Python 3.11+, Go 1.24.0, Clang with C++20, and Perl. Linux and
macOS are supported. Pinned upstream sources are included in `vendor/`.

Run from this directory, using a new output directory for each experiment:

```sh
# Capture the frozen historical calls, then export and compile the Full regions.
python3 scripts/capture.py --kind fixed --data /path/to/mainnet-23400001-23401000 --out work/fixed
python3 scripts/export.py --fixtures work/fixed/fixtures --out work/analysis
python3 scripts/build.py --analysis work/analysis/analysis.json

# Paired, fully prepared Vanilla and Hepha calls.
python3 scripts/fixed_calls.py --fixtures work/fixed/fixtures --out results/fixed

# Original-order stream with empty online binding caches.
python3 scripts/capture.py --kind stream --blocks 1000 \
  --data /path/to/mainnet-23400001-23401000 --out work/stream
python3 scripts/cold_start.py --witnesses work/stream/stream.jsonl.gz --out results/stream
```

For a quick check, pass a one-call `--targets` file to capture and fixed-call
measurement, and use `--blocks 1` for stream capture. The scripts produce fresh
measurements; no paper figures or previous timing results are included.

Cold binding publishes plans for previously observed code between blocks. It is
not an online C++ compilation benchmark. Both arms reuse a VM. State witnesses
are restored outside root-execution timing. See [measurement details](docs/EXPERIMENTS.md).

Core implementation: `src/`; C++ generation: `codegen/`; offline Go adapters:
`analysis/`. The regular `geth/` build is unchanged by these optional adapters.
Dependency licenses are retained under `vendor/`; see [NOTICE](NOTICE).
