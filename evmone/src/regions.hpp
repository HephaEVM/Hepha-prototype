#pragma once
#include <evmone/baseline.hpp>
#include <evmone/execution_state.hpp>
#include <evmone/vm.hpp>
#include <array>
#include <vector>
#include <string_view>
namespace evmone::regional {
using baseline::CodeAnalysis;
struct Plan;
struct RegionResult {const uint8_t* pc;uint256* end;int64_t gas;uint64_t instructions=0;bool original=false;};
using Native = RegionResult (*)(const Plan&,ExecutionState&,uint256*,uint256*,int64_t) noexcept;
Native find_native(std::string_view);
struct Affine {int input;uint256 add;uint256 eval(const uint256* s)const {return input<0?add:s[input]+add;}};
struct Clean {int kind;uint256 arg;uint256 apply(uint256 v)const;};
struct Read {Affine offset;bool memory;Clean clean;};
struct Write {Affine offset;int input;Clean clean;};
struct Bound {int end,start;uint256 size;};
struct Region {
 int stage=0,family=0;Native native=nullptr;const Region* next=nullptr;
 int kind=0,inputs=0,peak=0;uint64_t entry=0,exit=0,cost=0,loop_cost=0,instructions=0,loop_instructions=0;
 std::vector<Read> reads;std::vector<Write> writes;std::vector<Bound> bounds;std::vector<Affine> outputs;std::vector<int> read_outputs;
};
struct Counters {uint64_t attempts=0,commits=0,fallback=0,instructions=0,bytes=0;std::array<uint64_t,8> families{};};
struct Plan {
 bytes original;std::unique_ptr<CodeAnalysis> plain,patched;
 std::vector<Region> regions;std::vector<const Region*> entries;
 bool observe=false,diagnostic=false;mutable Counters stats;
 uint8_t stream=0;
 bool adaptive=false;mutable std::vector<uint8_t> rejected;
};
// Interpreter registers must never escape by reference into the slow helper.
RegionResult run_region(const Region&,const Plan&,ExecutionState&,uint256* bottom,uint256* end,int64_t gas) noexcept;
evmc_result execute(VM&,const evmc_host_interface&,evmc_host_context*,evmc_revision,const evmc_message&,const CodeAnalysis&,const Plan&) noexcept;
}
