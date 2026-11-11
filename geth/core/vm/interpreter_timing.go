package vm

// InterpreterLoopTiming observes the existing Run instruction-loop timer.
// Nanoseconds covers normal and early exits of top-level bytecode runs, including
// initcode. Child execution is already inside its parent's elapsed time. It
// excludes the top-level frame setup/cleanup and transaction pre/post processing;
// opcode dynamic gas, memory/state operations, child setup/cleanup and nested
// precompiles remain included. Top-level precompiles and empty code have no loop.
// Async completion outside the loop is not included; use synchronous comparisons.
type InterpreterLoopTiming struct {
	Nanoseconds, LegacyCallMicros  int64
	Runs, EarlyExits, CreationRuns uint64
}
