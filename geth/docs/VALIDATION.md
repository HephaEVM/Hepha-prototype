# Release validation

The extracted source was checked locally on macOS/arm64 with Go 1.24.0.
These checks validate packaging and execution; they do not constitute a new
full-range performance evaluation.

| Check | Result |
| --- | --- |
| Build all three public commands with the release tags | Passed |
| `go test` across all included packages with the release tags | Passed |
| Database-free normal-addition and overflow differential checks | Passed |
| Mine the bundled arithmetic compiler fixture | Passed: three template families |
| Native and Hepha replay of block 23,400,001 | Passed: 285 transactions; canonical state, receipts and gas validated |
| Warm fixed-call reproduction entry | Passed on one original root call, two paired rounds |
| Cold fixed-call reproduction entry | Passed on the same call with frozen checkpoints at 0 and 1 processed blocks |
| Raw-output summarizers and full-range command planning | Passed |
| Source scan for Chinese text, personal absolute paths, private keys and database files | Passed |

The warm check exercised host-executor preparation using the repository-local
sources and an empty compilation directory. The cold checks created fresh EVM
instances per transaction and recorded preparation separately. The supplied
database remained external to the source package.

The complete 1,000-block performance experiments were not rerun as part of this
packaging pass. Run the documented experiments on the target machine and retain
their raw outputs, validation records and configuration before reporting results.
