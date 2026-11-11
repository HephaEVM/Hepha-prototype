// arithmetic-runtime is a reproducible, database-free differential smoke test.
// Use optimizer-audit for historical blocks; these timings include cold binding
// and must not be used as warmed steady-state performance measurements.
package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"go-eth1.16.5-evm/core/state"
	"go-eth1.16.5-evm/core/tracing"
	"go-eth1.16.5-evm/core/types"
	"go-eth1.16.5-evm/core/vm"
	"go-eth1.16.5-evm/core/vm/runtime"
	miner "go-eth1.16.5-evm/template_miner"
)

type result struct {
	Mode           string             `json:"mode"`
	Return         string             `json:"return"`
	Error          string             `json:"error,omitempty"`
	GasUsed        uint64             `json:"gasUsed"`
	StateRoot      common.Hash        `json:"stateRoot"`
	ColdCallMicros int64              `json:"coldCallMicros"`
	Stats          vm.ArithmeticStats `json:"stats"`
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
func run() error {
	artifact := flag.String("artifact", "", "compiler artifact (or use -bytecode)")
	codePath := flag.String("bytecode", "", "runtime hex file")
	source := flag.String("source", "Arithmetic.sol", "source key")
	contract := flag.String("contract", "Arithmetic", "contract name")
	method := flag.String("method", "add", "ABI method")
	args := flag.String("args", "9,7", "comma-separated integer arguments")
	inputHex := flag.String("input", "", "raw calldata hex (overrides ABI method)")
	library := flag.String("templates", "", "mined bundle; omit to exercise negative-cache/feedback")
	feedback := flag.String("feedback-out", "", "write unresolved candidate evidence")
	gas := flag.Uint64("gas", 1000000, "gas supplied to the call")
	output := flag.String("out", "", "JSON results file; stdout if empty")
	flag.Parse()
	if *gas == 0 {
		return fmt.Errorf("-gas must be positive")
	}
	var code, input []byte
	decode := func(s string) ([]byte, error) {
		return hex.DecodeString(strings.TrimPrefix(strings.TrimSpace(s), "0x"))
	}
	if (*artifact == "") == (*codePath == "") {
		return fmt.Errorf("supply exactly one of -artifact or -bytecode")
	}
	if *artifact != "" {
		p, err := miner.Load(*artifact, *source, *contract)
		if err != nil {
			return err
		}
		code = p.Code
		if *inputHex == "" {
			input, err = miner.EncodeCall(p, *method, strings.Split(*args, ","))
			if err != nil {
				return err
			}
		}
	} else {
		raw, err := os.ReadFile(*codePath)
		if err != nil {
			return err
		}
		code, err = decode(string(raw))
		if err != nil {
			return err
		}
	}
	if *inputHex != "" {
		var err error
		input, err = decode(*inputHex)
		if err != nil {
			return err
		}
	}
	var libraryData []byte
	if *library != "" {
		var err error
		libraryData, err = os.ReadFile(*library)
		if err != nil {
			return err
		}
	}
	var results []result
	var lastEngine *vm.ArithmeticEngine
	for i, name := range []string{"baseline", "sync"} {
		var cfg *vm.ArithmeticConfig
		if i > 0 {
			var err error
			if len(libraryData) > 0 {
				lastEngine, err = vm.LoadArithmeticEngine(bytes.NewReader(libraryData), vm.ArithmeticEngineOptions{})
			} else {
				lastEngine, err = vm.NewArithmeticEngine(nil, vm.ArithmeticEngineOptions{})
			}
			if err != nil {
				return err
			}
			cfg = &vm.ArithmeticConfig{Engine: lastEngine, Mode: vm.ArithmeticSync}
		}
		s, err := state.New(types.EmptyRootHash, state.NewDatabaseForTesting())
		if err != nil {
			return err
		}
		addr := common.HexToAddress("0x1000")
		s.CreateAccount(addr)
		s.SetCode(addr, code, tracing.CodeChangeUnspecified)
		now := time.Now()
		ret, left, executionErr := runtime.Call(addr, input, &runtime.Config{State: s, GasLimit: *gas, EVMConfig: vm.Config{Arithmetic: cfg}})
		r := result{Mode: name, Return: hex.EncodeToString(ret), GasUsed: *gas - left, StateRoot: s.IntermediateRoot(true), ColdCallMicros: time.Since(now).Microseconds()}
		if executionErr != nil {
			r.Error = executionErr.Error()
		}
		if cfg != nil {
			r.Stats = cfg.Engine.Stats()
		}
		if i > 0 {
			b := results[0]
			if b.Return != r.Return || b.Error != r.Error || b.GasUsed != r.GasUsed || b.StateRoot != r.StateRoot {
				return fmt.Errorf("%s disagrees with baseline: %+v / %+v", name, b, r)
			}
		}
		results = append(results, r)
	}
	if *feedback != "" {
		f, err := os.Create(*feedback)
		if err != nil {
			return err
		}
		writeErr := lastEngine.WriteFeedback(f)
		closeErr := f.Close()
		if writeErr != nil {
			return writeErr
		}
		if closeErr != nil {
			return closeErr
		}
	}
	data, err := json.MarshalIndent(map[string]any{"validation": "return/error/gas/state-root equal in native and synchronous modes", "timing": "cold call includes recognition; not a performance result", "runs": results}, "", "  ")
	if err != nil {
		return err
	}
	if *output != "" {
		return os.WriteFile(*output, append(data, '\n'), 0644)
	}
	fmt.Println(string(data))
	return nil
}
