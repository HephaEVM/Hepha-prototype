# Historical replay dataset

The reproduction scripts expect an external directory with this layout:

```text
mainnet-23400001-23401000/
  chaindata/    Pebble chain database
  state/        LevelDB state subset at block 23,400,000
```

The reference range is Ethereum mainnet blocks **23,400,001–23,401,000**,
containing **226,031 transactions**. Its parent identifiers are:

- Block hash: `0x91f74cc48392039b9db99a47fd0dbee28cae2ac92ad685b0a80135f41ac0ae05`
- State root: `0xab955667a17523870dde8dc14fa6bf4456252e480c8a61d0c2a701a6bf6caca5`

The compact state subset contains the trie nodes, accounts, storage and code read
by this replay range. It is not a complete mainnet snapshot and cannot be assumed
to support unrelated transactions or later blocks. The chain and state stores
are opened read-only; replay writes go to a separate in-memory overlay.

The database is **not included in the Git source tree**. Download the
[dataset archive](https://github.com/HephaEVM/Hepha-prototype/releases/download/v0.1.0/hepha-mainnet-23400001-23401000.tar.gz)
and its
[SHA-256 checksum](https://github.com/HephaEVM/Hepha-prototype/releases/download/v0.1.0/hepha-mainnet-23400001-23401000.tar.gz.sha256)
from Release v0.1.0 (approximately 740 MiB compressed). From a directory outside
the repository:

```sh
base=https://github.com/HephaEVM/Hepha-prototype/releases/download/v0.1.0
file=hepha-mainnet-23400001-23401000.tar.gz
curl -fL --retry 3 -O "$base/$file"
curl -fL --retry 3 -O "$base/$file.sha256"
shasum -a 256 -c "$file.sha256"
tar -xzf "$file"
```

A directory containing only block bodies, or an ordinary current-state Geth
database, is not a substitute. Database-free smoke tests remain runnable without
this dataset.

`workloads/dataset.json` describes the expected range and database formats.
`workloads/dataset.sha256` contains the reference database-file checksums. Check
them from the dataset directory:

```sh
cd /path/to/mainnet-23400001-23401000
shasum -a 256 -c /path/to/Hepha/geth/workloads/dataset.sha256
```

Replay also validates canonical block continuity, state roots, receipt roots, and
gas against the recorded block headers. Missing state or mismatched roots must be
treated as an incomplete reproduction, not omitted from timing aggregates.

The fixed-call manifest stores transaction identities only. The replay captures
their actual pre-state, calldata and call context from the dataset; it does not
fabricate contract inputs or insert additional overflow checks into deployed code.
