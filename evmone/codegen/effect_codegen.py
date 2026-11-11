"""Ordered effect lowering. State primitives remain upstream evmone operations."""
from codegen import expr,word

def emit_effect(emitter,name,x):
    nodes=x['nodes'];actions=x['actions'];used={i for _,i in x['writes']}
    def live_stack(stack):
        return [(i,j) for i,j in enumerate(stack) if i>=32-x['inputs'] and not (i<32 and i==j)]
    for a in actions:
        used.update(a['args'])
        if a['reuse']>=0:used.add(a['reuse'])
        for stack in ['checkpoint','before']:used.update(j for _,j in live_stack(a[stack]))
    for i in reversed(range(len(nodes))):
        if i in used and nodes[i]['op'] not in ['constant','input','effect','CALLER']:
            used.add(nodes[i]['a'])
            if nodes[i]['op'] not in ['NOT','ISZERO']:used.add(nodes[i]['b'])
    lines=[emitter.signature(name),f'const auto n=end-bottom;if(n<{x["inputs"]}||n+{x["peak"]}>1024)return {{}};', 'alignas(32) uint256 args[6];']
    def resume(stack,pc,cost,steps):
        r=[]
        for i,j in live_stack(stack):r.append(f'end[{i-32}]=v{j};')
        r.append(f'return {{p.patched->code().data()+{pc},end+{len(stack)-32},gas-{cost},{steps},true}};')
        return ''.join(r)
    for node,action in x['events']:
        if node>=0:
            if node not in used:continue
            v=nodes[node];op=v['op']
            if op=='input':
                assert v['depth']<=x['inputs'],(name,v,x['inputs'])
                value=f'end[-{v["depth"]}]'
            elif op=='constant':value=word(v['value'])
            elif op=='CALLER':value='intx::be::load<uint256>(s.msg->sender)'
            else:value=expr(op,[f'v{v["a"]}',f'v{v["b"]}'],v['value'])
            lines.append(f'const uint256 v{node}={value};');continue
        a=actions[action];op=a['op'];prefix=a['gas'];ids=a['args']
        lines.append(f'if(gas<{prefix}){{'+resume(a['checkpoint'],a['start'],0,a['start_step'])+'}')
        if op=='JUMPI':
            lines.append(f'if(v{ids[0]}==0){{'+resume(a['before'],a['pc'],prefix-10,a['op_step'])+'}')
            lines.append(f'gas-={prefix};');continue
        if op=='helper':
            fn=emitter.helper(a['helper']);h=a['helper_gas']
            lines.append(f'const auto q{action}={fn}(v{ids[0]},v{ids[1]});if(!q{action}.ok||q{action}.cost!={h}){{'+resume(a['before'],a['pc'],prefix-8-h,a['op_step'])+'}')
            lines.append(f'const uint256 v{a["result"]}=q{action}.value;gas-={prefix};');continue
        lines.append(f'gas-={prefix};')
        if op=='KECCAK256' and a['reuse']>=0:
            lines.append(f'if(!check_memory(gas,s.memory,v{ids[1]},v{ids[0]})||(gas-=12)<0){{s.status=EVMC_OUT_OF_GAS;return {{nullptr,end,gas}};}}')
            continue
        for j,i in enumerate(ids):lines.append(f'args[{j}]=v{i};')
        fn=f'log<{int(op[3:])}>' if op.startswith('LOG') else op.lower()
        lines.append(f'const auto r{action}=instr::core::{fn}(args+{len(ids)},gas,s);gas=r{action}.gas_left;if(r{action}.status!=EVMC_SUCCESS){{s.status=r{action}.status;return {{nullptr,end,gas}};}}')
        if a['result']>=0:lines.append(f'const uint256 v{a["result"]}=args[0];')
    for offset,i in x['writes']:lines.append(f'end[{offset}]=v{i};')
    lines += [f'return {{p.patched->code().data()+{x["exit"]},end+{x["delta"]},gas,{x["instructions"]}}};','}']
    return '\n'.join(lines)
