package templateminer

import (
	"bytes"
	"compress/gzip"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
)

// Convert actual EVM fixture executions to the same format as the chain tracer.
// The miner receives no AST, source map, Yul, or pre-existing template library.
func auditFixture(t *testing.T, programs []*Program, methods []string) (string, []AuditRecord) {
	t.Helper()
	dir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(dir, "bytecode"), 0755); err != nil {
		t.Fatal(err)
	}
	var records []AuditRecord
	for i, p := range programs {
		input, err := EncodeCall(p, methods[i], []string{"7", "9"})
		if err != nil {
			t.Fatal(err)
		}
		tr := Replay(p, input)
		hash := crypto.Keccak256Hash(p.Code).Hex()
		if err = os.WriteFile(filepath.Join(dir, "bytecode", hash+".hex"), []byte(hex.EncodeToString(p.Code)), 0644); err != nil {
			t.Fatal(err)
		}
		id := uint64(i + 1)
		addr := "0x" + strings.Repeat("0", 39) + string(rune('1'+i))
		records = append(records, AuditRecord{Kind: "frame", CallType: "CALL", CallID: id, Block: 1, Tx: hash, CodeHash: hash, CodeAddress: addr})
		for j, s := range tr.Steps {
			rec := AuditRecord{Kind: "opcode", CallID: id, Step: j, PC: s.PC, Op: s.Op}
			for _, v := range s.Stack {
				rec.Stack = append(rec.Stack, v.Value)
			}
			records = append(records, rec)
		}
		records = append(records, AuditRecord{Kind: "exit", CallID: id, Output: "0x" + tr.Return, Reverted: tr.ExecutionError != "", Error: tr.ExecutionError})
	}
	data := []byte(`[{"number":1,"transactions":2,"gasUsed":1,"expectedGasUsed":1,"stateRoot":"test-state","expectedStateRoot":"test-state","receiptRoot":"test-receipt","expectedReceiptRoot":"test-receipt"}]`)
	if err := os.WriteFile(filepath.Join(dir, "block-validation.json"), data, 0644); err != nil {
		t.Fatal(err)
	}
	writeAuditFixture(t, dir, records)
	return dir, records
}
func writeAuditFixture(t *testing.T, dir string, records []AuditRecord) {
	t.Helper()
	f, err := os.Create(filepath.Join(dir, "traces.jsonl.gz"))
	if err != nil {
		t.Fatal(err)
	}
	z := gzip.NewWriter(f)
	enc := json.NewEncoder(z)
	for _, rec := range records {
		if err = enc.Encode(rec); err != nil {
			t.Fatal(err)
		}
	}
	if err = z.Close(); err != nil {
		t.Fatal(err)
	}
	if err = f.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestMainnetMiningWithoutCompilerArtifacts(t *testing.T) {
	p := fixture(t, "arithmetic-0824.json")
	q := fixture(t, "arithmetic-0819.json")
	dir, _ := auditFixture(t, []*Program{p, q}, []string{"add", "add"})
	r, err := MineAudit(dir, "")
	if err != nil {
		t.Fatal(err)
	}
	if len(r.Templates) != 1 || len(r.Contracts) != 2 || len(r.Families) != 1 || len(r.Families[0].CodeHashes) != 2 || len(r.Families[0].CompilerVersions) != 2 {
		t.Fatalf("cross-version/contract normalization failed: %+v", r.Families)
	}
	for _, c := range r.Contracts {
		if c.OptimizerEnabled != nil || c.OptimizerMode != "unknown" {
			t.Fatal("missing settings became optimizer=false")
		}
		if len(c.Instances) != 1 || c.Instances[0].Executions != 1 {
			t.Fatal("missing execution evidence")
		}
		w := c.Instances[0].Witnesses[0]
		if w.Result != "10" {
			t.Fatal("wrong concrete sum")
		}
		origins := 0
		for _, j := range w.Jumps {
			if j.ProducerPC != nil {
				origins++
			}
		}
		if origins < 4 || !w.Jumps[len(w.Jumps)-1].FromEntry {
			t.Fatal("cleanup and final return provenance lost")
		}
	}
	// Existing template matching remains compatible with a mainnet-mined family.
	m, err := MatchBytecode(q.Code, r.Templates)
	if err != nil || len(m.Bindings) != 1 {
		t.Fatalf("mined family not reusable: %v", err)
	}
}

func TestMainnetRejectsUncheckedAndIncompleteOrDamagedTrace(t *testing.T) {
	p := fixture(t, "arithmetic-0824.json")
	// Use the unchecked method name declared by the Solidity fixture.
	input, err := EncodeCall(p, "uncheckedAdd", []string{"7", "9"})
	if err != nil {
		t.Fatal(err)
	}
	_ = input
	dir, records := auditFixture(t, []*Program{p}, []string{"uncheckedAdd"})
	r, err := MineAudit(dir, "")
	if err != nil || len(r.Templates) != 0 {
		t.Fatalf("unchecked accepted: %v", err)
	}
	writeAuditFixture(t, dir, records[:len(records)-1])
	if _, err = MineAudit(dir, ""); err == nil {
		t.Fatal("unfinished frame accepted")
	}
	dir, records = auditFixture(t, []*Program{p}, []string{"add"})
	// Mutate the observed result without changing bytecode or the branch.
	for i := range records {
		if records[i].Kind == "opcode" && records[i].PC == 630 {
			records[i].Stack[len(records[i].Stack)-1] = "0x11"
			break
		}
	}
	writeAuditFixture(t, dir, records)
	r, err = MineAudit(dir, "")
	if err != nil {
		t.Fatal(err)
	}
	if len(r.Templates) != 0 {
		t.Fatal("wrong concrete result accepted")
	}
	for i := range records {
		if records[i].Kind == "opcode" {
			records[i].Op = "INVALID"
			break
		}
	}
	writeAuditFixture(t, dir, records)
	if _, err = MineAudit(dir, ""); err == nil {
		t.Fatal("opcode/bytecode mismatch accepted")
	}
}

func TestHistoricalCompilerEvidenceAndCreationExclusion(t *testing.T) {
	p := fixture(t, "arithmetic-0824.json")
	dir, records := auditFixture(t, []*Program{p}, []string{"add"})
	address := records[0].CodeAddress
	evidence := map[string]any{"sourceURL": "https://sourcify.dev/example", "response": map[string]any{"address": address, "chainId": "1", "runtimeMatch": "exact_match", "runtimeBytecode": map[string]any{"onchainBytecode": "0x" + hex.EncodeToString(p.Code)}, "compilation": map[string]any{"compilerVersion": "0.8.24+commit.e11b9ed9", "compilerSettings": map[string]any{"optimizer": map[string]any{"enabled": false, "runs": 200}}}}}
	write := func() {
		data, _ := json.Marshal(evidence)
		if err := os.WriteFile(filepath.Join(dir, "evidence.json"), data, 0644); err != nil {
			t.Fatal(err)
		}
	}
	write()
	l := ChainLabel{Address: address, EvidenceFile: "evidence.json"}
	verified := verifyChainLabel(l, dir, p.Code, "0.8.24")
	if verified.OptimizerEnabled == nil || *verified.OptimizerEnabled || verified.OptimizerMode != "false" || verified.EvidenceSHA256 == "" {
		t.Fatal("exact runtime evidence not accepted")
	}
	bad := append([]byte(nil), p.Code...)
	bad[0] ^= 1
	if verifyChainLabel(l, dir, bad, "0.8.24").OptimizerMode != "unknown" {
		t.Fatal("different historical runtime labelled")
	}
	settings := evidence["response"].(map[string]any)["compilation"].(map[string]any)["compilerSettings"].(map[string]any)
	settings["optimizer"] = map[string]any{"details": map[string]any{"yul": true}}
	write()
	if got := verifyChainLabel(l, dir, p.Code, "0.8.24"); got.OptimizerMode != "custom_details" || got.OptimizerEnabled != nil {
		t.Fatal("custom optimizer forced to boolean")
	}
	records[0].CallType = "CREATE"
	writeAuditFixture(t, dir, records)
	r, err := MineAudit(dir, "")
	if err != nil || len(r.Contracts) != 0 || r.ExcludedFrames["creation code, not deployed runtime"] != 1 {
		t.Fatalf("init code accepted: %v", err)
	}
}

func TestMainnetOverflowWitnessAndOldCompilerExclusion(t *testing.T) {
	p := fixture(t, "arithmetic-0824.json")
	dir, _ := auditFixture(t, []*Program{p}, []string{"sub"}) // 7 - 9 must panic.
	r, err := MineAudit(dir, "")
	if err != nil {
		t.Fatal(err)
	}
	if len(r.Templates) != 1 || r.Templates[0].Operation != "SUB" || r.Contracts[0].Instances[0].Panics != 1 {
		t.Fatal("real panic witness not retained")
	}
	// Change only the compiler trailer: identical arithmetic under a pre-0.8
	// label remains out of scope in the explicitly requested 0.8-only mode.
	p.Code = append([]byte(nil), p.Code...)
	i := bytes.LastIndex(p.Code, []byte{0x64, 's', 'o', 'l', 'c', 0x43})
	if i < 0 {
		t.Fatal("missing compiler metadata")
	}
	p.Code[i+7] = 7
	p.Code[i+8] = 6
	dir, _ = auditFixture(t, []*Program{p}, []string{"add"})
	r, err = MineAuditWithModel(dir, "", "solidity08")
	if err != nil || len(r.Templates) != 0 {
		t.Fatalf("pre-0.8 accepted: %v", err)
	}
}

func TestMainnetRejectsWrongPredicateAndSeparatesInterleavedFrames(t *testing.T) {
	p := fixture(t, "arithmetic-0824.json")
	q := fixture(t, "arithmetic-0819.json")
	dir, records := auditFixture(t, []*Program{p, q}, []string{"add", "add"})
	boundary := 0
	for i, r := range records {
		if r.Kind == "frame" && r.CallID == 2 {
			boundary = i
			break
		}
	}
	interleaved := append([]AuditRecord(nil), records[:20]...)
	interleaved = append(interleaved, records[boundary:]...)
	interleaved = append(interleaved, records[20:boundary]...)
	writeAuditFixture(t, dir, interleaved)
	r, err := MineAudit(dir, "")
	if err != nil || len(r.Templates) != 1 || r.Families[0].ObservedExecutions != 2 {
		t.Fatalf("interleaved frames mixed: %v", err)
	}
	var add Helper
	for _, h := range p.Helpers {
		if h.Operation == "ADD" {
			add = h
		}
	}
	a, err := Analyze(p, add)
	if err != nil {
		t.Fatal(err)
	}
	for _, pc := range a.ReachablePCs {
		if p.ByPC[pc].Op == vm.GT {
			p.Code = append([]byte(nil), p.Code...)
			p.Code[pc] = byte(vm.LT)
			ins := p.ByPC[pc]
			ins.Op = vm.LT
			p.ByPC[pc] = ins
			break
		}
	}
	dir, _ = auditFixture(t, []*Program{p}, []string{"add"})
	r, err = MineAudit(dir, "")
	if err != nil || len(r.Templates) != 0 {
		t.Fatalf("wrong overflow predicate accepted: %v", err)
	}
}
