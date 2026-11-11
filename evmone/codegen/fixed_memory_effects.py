"""Prove constant-address memory operations once on entry to an effect region.

EVM frame memory never shrinks. If its existing size covers all substituted
constant accesses, they require no expansion and cannot fail a bounds check.
All gas charges remain in place; unsupported accesses retain upstream helpers.
"""
import re

SIGNATURE = '(const Plan&,ExecutionState&,uint256*,uint256*,int64_t) noexcept;\n'
CONSTANT = re.compile(r'^const uint256 (v\d+)=uint256\{(0x[0-9a-f]+)ULL,0x0ULL,0x0ULL,0x0ULL\};$')
ARGUMENT = re.compile(r'^args\[(\d+)\]=(v\d+);$')
CALL = re.compile(r'^const auto r(\d+)=instr::core::(mload|mstore|keccak256)\(args\+(\d+),gas,s\);gas=r\1.gas_left;if\(r\1.status!=EVMC_SUCCESS\)(?: \[\[unlikely\]\] )?\{s.status=r\1.status;return \{nullptr,end,gas\};\}$')


def transform(name, body):
    # The supported effect emitter has no CALL/CREATE/RETURN or memory shrink.
    operations = re.findall(r'instr::core::([\w<>]+)\(', body)
    if any(op not in ('mload','mstore','keccak256','sload','sstore','log<0>','log<1>','log<2>','log<3>','log<4>') for op in operations):
        return body
    constants, arguments, changed, required = {}, {}, [], 0
    count = 0
    for line in body.splitlines():
        if m := CONSTANT.fullmatch(line): constants[m[1]] = int(m[2],16)
        if m := ARGUMENT.fullmatch(line): arguments[int(m[1])] = m[2]
        if m := CALL.fullmatch(line):
            op, n = m[2], int(m[3]); offset = constants.get(arguments.get(n-1))
            length = constants.get(arguments.get(0)) if op == 'keccak256' else 32
            if offset is not None and length is not None and 0 <= offset <= 4096 and 0 <= length <= 256:
                required = max(required,offset+length if length else 0); count += 1
                if op == 'mstore':
                    line = f'store(s,{offset},args[0]);'
                elif op == 'mload':
                    line = f'args[0]=intx::be::load<uint256>(std::span<const uint8_t,32>{{&s.memory[{offset}],32}});'
                else:
                    cost = 6*((length+31)//32); data = f'&s.memory[{offset}]' if length else 'nullptr'
                    line = f'if((gas-={cost})<0){{s.status=EVMC_OUT_OF_GAS;return {{nullptr,end,gas}};}}\nargs[0]=intx::be::load<uint256>(ethash::keccak256({data},{length}));'
            arguments.clear()
        elif 'instr::core::' in line:
            arguments.clear()
        changed.append(line)
    if count < 3:return body
    # Static literals are bounded; the guard adds no memory access before the
    # original stack precondition and the fallback runs the exact old function.
    assert changed[1].startswith('const auto n=end-bottom;')
    changed.insert(2, f'if(s.memory.size()<{required}) [[unlikely]] return ORIGINAL_EXECUTOR(p,s,bottom,end,gas);')
    return 'RegionResult ORIGINAL_EXECUTOR'+SIGNATURE+'\n'+'\n'.join(changed)+'\n'
