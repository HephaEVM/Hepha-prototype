#pragma once
#include "regions.hpp"
#include <evmone_precompiles/keccak.hpp>
#include <algorithm>
#include <cstring>

namespace evmone::regional {
constexpr uint64_t max_memory = 0x1fffffffe0;
inline bool u64(const uint256& v) { return (v[1] | v[2] | v[3]) == 0; }
inline bool signed_less(const uint256& a, const uint256& b) {
    const bool x = a[3] >> 63, y = b[3] >> 63;
    return x != y ? x : a < b;
}
inline uint256 sext(const uint256& index, uint256 value) {
    if (index >= 31) return value;
    const auto bit = unsigned(index[0] * 8 + 7);
    const auto mask = (uint256{1} << bit) - 1;
    return ((value >> bit) & 1) ? value | ~mask : value & mask;
}
inline uint256 sar(const uint256& shift, const uint256& value) {
    const auto sign = (value[3] >> 63) ? ~uint256{} : uint256{};
    return (value >> shift) | (sign << (shift < 256 ? 256 - shift[0] : 0));
}
inline bool memory_cost(ExecutionState& s, uint64_t end, uint64_t fixed,
    int64_t gas, uint64_t& size, uint64_t& total) {
    if (gas < 0 || fixed > uint64_t(gas) || end > max_memory) return false;
    size = s.memory.size(); total = fixed; end = (end + 31) & ~uint64_t{31};
    if (end > size) {
        const auto ow = size / 32, nw = end / 32;
        const auto cost = 3*nw + nw*nw/512 - (3*ow + ow*ow/512);
        if (cost > uint64_t(gas) - total) return false;
        total += cost; size = end;
    }
    return true;
}
inline void grow(ExecutionState& s, uint64_t size) {
    if (size > s.memory.size()) s.memory.grow(size);
}
inline uint256 load(const uint8_t* data, size_t size, const uint256& offset) {
    std::array<uint8_t,32> b{};
    if (u64(offset) && offset[0] < size)
        std::memcpy(b.data(), data+offset[0], std::min<size_t>(size-offset[0],32));
    return intx::be::load<uint256>(std::span<const uint8_t,32>{b});
}
inline void store(ExecutionState& s, uint64_t offset, const uint256& value) {
    intx::be::store(std::span<uint8_t,32>{&s.memory[offset],32},value);
}
inline bool jump(const Plan& p, const uint256& v) {
    return u64(v) && p.plain->check_jumpdest(v[0]);
}
struct StagedWrite { uint64_t offset; uint256 value; };
inline uint256 staged_load(ExecutionState& s, uint64_t offset,
    const StagedWrite* writes, size_t count) {
    std::array<uint8_t,32> word{};
    if (offset < s.memory.size())
        std::memcpy(word.data(), &s.memory[offset], std::min<size_t>(s.memory.size()-offset,32));
    for (size_t i=0;i<count;i++) {
        const auto& w=writes[i];
        const auto first=std::max(offset,w.offset), last=std::min(offset+32,w.offset+32);
        if (first<last) {
            uint8_t bytes[32]; intx::be::store(bytes,w.value);
            std::memcpy(word.data()+first-offset,bytes+first-w.offset,last-first);
        }
    }
    return intx::be::load<uint256>(std::span<const uint8_t,32>{word});
}
struct ArithmeticResult { uint256 value; uint64_t cost, instructions; bool ok; };
struct MulDivResult { uint256 value; bool wide, rounded, ok; };
inline MulDivResult muldiv(const uint256& a,const uint256& b,const uint256& d,bool ceil) {
    if (d==0) return {};
    const auto product=intx::umul(a,b);
    const uint256 low{product[0],product[1],product[2],product[3]};
    const uint256 high{product[4],product[5],product[6],product[7]};
    if (high>=d) return {};
    const bool wide=high!=0;
    uint256 value;
    if (wide) {
        const auto q=intx::udivrem(product,d).quot;
        value={q[0],q[1],q[2],q[3]};
    } else value=low/d;
    const bool rounded=ceil && low-value*d!=0;
    if (rounded) { value += 1; if (value==0) return {}; }
    return {value,wide,rounded,true};
}
}
