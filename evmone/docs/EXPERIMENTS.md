# Experiment boundaries

## Fixed calls

The default identities are the same 556 calls as `geth/workloads/fixed-calls.json`.
`capture.py` reconstructs their historical pre-state through canonical Geth
replay, captures every requested state value, and records an untimed executed-PC
census. `export.py` re-proves region semantics using the current sibling Geth
implementation. Observed PCs limit preparation work; they do not prove safety.
The full policy enables value passing, computation reuse and check elimination.

`build.py` generates direct C++ functions, then compiles Vanilla evmone and the
region backend with the same compiler and optimization flags. The original
interpreter implementation remains the control. Inputs are restored between
measurements. Eight rounds alternate method order, with 32 repetitions per batch
by default. Timed lookup/analysis misses fail the warm experiment. Return values,
remaining gas, refunds, accounts, storage and logs are validated.

The result is root execution time, including nested calls once. Preparation,
witness decoding and state restoration are outside timing. The summarizer takes
each call's median across rounds, then compares latency sums; it does not average
individual percentage speedups or select the fastest round.

## Cold binding stream

The compact database supplies an untimed reference capture of original mainnet
transactions in canonical order. Every block is validated against its recorded
state root, receipt root and gas. The native driver subsequently runs each root
against that transaction's own pre-state witness and validates observable effects.
It does not carry a native trie forward or measure disk I/O.

Every replay starts a new process with empty analyzed-code and bound-plan caches.
Both methods reuse their VM across transactions. First-use analysis is included
in root time. Hepha queues observed code and binds installed executors in the
next between-block window, up to 10 ms per boundary. No future transaction is
executed to train a plan. Uncovered code follows Vanilla. Preparation time is
recorded separately. Three complete, original-order replays alternate method
order; all completed repetitions are retained.

The portable source release contains the core Full backend and real progressive
binding. Historical platform-specific instruction-layout/LD_PRELOAD experiments
are not bundled. This release does not promise identical historical percentages;
compiler, hardware, input corpus and preparation boundary affect results.

## Outputs and checks

- Fixed calls: `observations.jsonl`, `summary.json`, `protocol.json` and completion.
- Stream: per-method, per-repeat `observations.jsonl`, `complete.json`, and logs.
- `complete.json` records root latency and between-block binding separately.
- Zero-instruction transactions remain in stream output with zero root latency.
- No simulation, manual percentage adjustment or plot rendering is performed.

Run the focused correctness tests with:

```sh
PYTHONPATH=codegen python3 -m unittest discover -s tests -v
```

The tests check exact stack restoration after effects and low gas, including
poisoned new stack slots, and conservative admission of loop-exit summaries.
