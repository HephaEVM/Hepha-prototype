# Hepha

Hepha shortens EVM execution paths using verified semantic regions: **value
passing**, **computation reuse**, and **check elimination**. Unavailable regions
and unmet safety conditions fall back to the original interpreter.

| Directory | Implementation | Experiments |
| --- | --- | --- |
| [`geth/`](geth/README.md) | Go backend, based on geth v1.16.5 | Warm fixed calls; cold fixed calls and stream |
| [`evmone/`](evmone/README.md) | Native C++ backend, based on evmone v0.25.0 | Warm fixed calls; cold fixed calls and stream |

Start with the README for the backend you want to run. Scripts produce measured
JSON/CSV results and verify execution; no historical results or plotting inputs
are included. The two backends share offline semantic analysis, but execute in
their own language without Go callbacks in the native C++ timed path.

## Dataset

The compact Ethereum dataset is distributed separately in
[Release v0.1.0](https://github.com/HephaEVM/Hepha-prototype/releases/tag/v0.1.0).
Download the dataset archive and its SHA-256 checksum from that release. It covers
**1,000 mainnet blocks (23,400,001-23,401,000), 226,031 transactions**, and the
state needed to replay this range. Extract it outside the repository and pass
its directory with `--data`. See [dataset instructions](geth/docs/DATA.md).
