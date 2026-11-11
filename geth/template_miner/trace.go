package templateminer

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/core/vm/runtime"
	"github.com/ethereum/go-ethereum/params"
)

type ValueNode struct {
	ID           int    `json:"id"`
	ProducerStep int    `json:"producerStep"`
	Operation    string `json:"operation"`
	Inputs       []int  `json:"inputs,omitempty"`
	Value        string `json:"value,omitempty"`
}
type ValueRef struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}
type Step struct {
	Memory string         `json:"entryMemory,omitempty"`
	PC     uint64         `json:"pc"`
	Op     string         `json:"op"`
	Stack  []ValueRef     `json:"stackBefore"`
	Inputs []int          `json:"inputs,omitempty"`
	Output int            `json:"output,omitempty"`
	Source SourceLocation `json:"source"`
}
type Trace struct {
	Steps          []Step      `json:"steps"`
	Values         []ValueNode `json:"values"`
	Return         string      `json:"return"`
	ExecutionError string      `json:"executionError,omitempty"`
	TrackingError  string      `json:"trackingError,omitempty"`
	GasUsed        uint64      `json:"gasUsed"`
}

// Replay uses upstream Geth, so mining is unaffected by the local skipper or
// synthetic baseline work. Each produced value gets a fresh ID; DUP/SWAP keep
// IDs. Equal concrete values are never treated as equal provenance.
func Replay(p *Program, input []byte) *Trace {
	t := &Trace{}
	table, _ := vm.LookupInstructionSet(params.Rules{IsCancun: true})
	ids := []int{}
	memory := map[uint64]int{}
	newValue := func(op string, inputs []int) int {
		id := len(t.Values) + 1
		t.Values = append(t.Values, ValueNode{ID: id, ProducerStep: len(t.Steps), Operation: op, Inputs: append([]int(nil), inputs...)})
		return id
	}
	hooks := &tracing.Hooks{}
	hooks.OnOpcode = func(pc uint64, op byte, gas, cost uint64, scope tracing.OpContext, rData []byte, depth int, err error) {
		if depth != 1 || !bytes.Equal(scope.ContractCode(), p.Code) || t.TrackingError != "" {
			return
		}
		actual := scope.StackData()
		if err != nil {
			return
		}
		if len(actual) != len(ids) {
			t.TrackingError = fmt.Sprintf("stack provenance lost at PC %d: predicted %d, actual %d", pc, len(ids), len(actual))
			return
		}
		ins, ok := p.ByPC[pc]
		if !ok {
			t.TrackingError = "PC not on instruction boundary"
			return
		}
		step := Step{PC: pc, Op: vm.OpCode(op).String(), Source: ins.Source}
		if p.Artifact != nil && legacyCompiler(p.Artifact.CompilerVersion) && vm.OpCode(op) == vm.JUMPDEST && len(actual) >= 4 && len(scope.MemoryData()) <= 1<<20 {
			step.Memory = hex.EncodeToString(scope.MemoryData())
		}
		for i, v := range actual {
			value := v.Hex()
			step.Stack = append(step.Stack, ValueRef{ids[i], value})
			if ids[i] > 0 {
				t.Values[ids[i]-1].Value = value
			}
		}
		n := len(ids)
		opcode := vm.OpCode(op)
		switch {
		case opcode >= vm.DUP1 && opcode <= vm.DUP16:
			d := int(opcode-vm.DUP1) + 1
			step.Inputs = []int{ids[n-d]}
			ids = append(ids, ids[n-d])
		case opcode >= vm.SWAP1 && opcode <= vm.SWAP16:
			d := int(opcode-vm.SWAP1) + 1
			ids[n-1], ids[n-1-d] = ids[n-1-d], ids[n-1]
		default:
			pops, maxStack := table[op].Stack()
			pushes := pops + int(params.StackLimit) - maxStack
			if pops > n || pushes < 0 || pushes > 1 {
				t.TrackingError = fmt.Sprintf("unsupported stack effect at %d", pc)
				return
			}
			for i := 0; i < pops; i++ {
				step.Inputs = append(step.Inputs, ids[n-1-i])
			}
			if opcode == vm.MSTORE {
				offset := actual[n-1].Uint64()
				if !actual[n-1].IsUint64() {
					memory = map[uint64]int{}
				} else {
					for k := range memory {
						if k >= offset && k-offset < 32 || offset >= k && offset-k < 32 {
							delete(memory, k)
						}
					}
					memory[offset] = ids[n-2]
				}
			}
			if opcode == vm.MSTORE8 || opcode == vm.CALLDATACOPY || opcode == vm.CODECOPY || opcode == vm.RETURNDATACOPY || opcode == vm.MCOPY || opcode == vm.CALL || opcode == vm.DELEGATECALL || opcode == vm.STATICCALL || opcode == vm.CALLCODE {
				memory = map[uint64]int{}
			}
			ids = ids[:n-pops]
			if pushes == 1 {
				inputs := step.Inputs
				if opcode == vm.MLOAD && actual[n-1].IsUint64() {
					if source, ok := memory[actual[n-1].Uint64()]; ok {
						inputs = append(append([]int(nil), inputs...), source)
					}
				}
				id := newValue(opcode.String(), inputs)
				step.Output = id
				ids = append(ids, id)
			}
		}
		t.Steps = append(t.Steps, step)
	}
	hooks.OnExit = func(depth int, output []byte, gasUsed uint64, err error, reverted bool) {
		if depth == 0 {
			t.GasUsed = gasUsed
		}
	}
	ret, _, err := runtime.Execute(p.Code, input, &runtime.Config{GasLimit: 10_000_000, EVMConfig: vm.Config{Tracer: hooks}})
	t.Return = hex.EncodeToString(ret)
	if err != nil {
		t.ExecutionError = err.Error()
	}
	return t
}

func EncodeCall(p *Program, method string, arguments []string) ([]byte, error) {
	parsed, err := abi.JSON(strings.NewReader(string(p.Contract.ABI)))
	if err != nil {
		return nil, err
	}
	m, ok := parsed.Methods[method]
	if !ok {
		return nil, fmt.Errorf("ABI method %q not found", method)
	}
	if len(m.Inputs) != len(arguments) {
		return nil, fmt.Errorf("%s requires %d arguments", method, len(m.Inputs))
	}
	vals := make([]any, len(arguments))
	for i, arg := range arguments {
		typ := m.Inputs[i].Type
		switch typ.T {
		case abi.UintTy, abi.IntTy:
			n, ok := new(big.Int).SetString(arg, 0)
			if !ok {
				return nil, fmt.Errorf("invalid integer %q", arg)
			}
			domain := integerType{bits: typ.Size, signed: typ.T == abi.IntTy}
			lo, hi := domain.bounds()
			if n.Cmp(lo) < 0 || n.Cmp(hi) > 0 {
				return nil, fmt.Errorf("integer outside %s range", typ.String())
			}
			v := reflect.New(typ.GetType()).Elem()
			if v.Kind() == reflect.Ptr {
				vals[i] = n
			} else {
				if typ.T == abi.IntTy {
					v.SetInt(n.Int64())
				} else {
					v.SetUint(n.Uint64())
				}
				vals[i] = v.Interface()
			}
		case abi.StringTy:
			vals[i] = arg
		case abi.AddressTy:
			if !common.IsHexAddress(arg) {
				return nil, fmt.Errorf("invalid address %q", arg)
			}
			vals[i] = common.HexToAddress(arg)
		default:
			return nil, fmt.Errorf("unsupported CLI input type %s", typ.String())
		}
	}
	return parsed.Pack(method, vals...)
}

func (t *Trace) Depends(value, ancestor int) bool {
	seen := map[int]bool{}
	var visit func(int) bool
	visit = func(id int) bool {
		if id == ancestor {
			return true
		}
		if id <= 0 || id > len(t.Values) || seen[id] {
			return false
		}
		seen[id] = true
		for _, parent := range t.Values[id-1].Inputs {
			if visit(parent) {
				return true
			}
		}
		return false
	}
	return visit(value)
}
