"""Lower trusted, re-proved Go IR to native C++; no runtime DAG interpreter."""
import hashlib,json
from pathlib import Path

def word(x):
    v=int(x,0) if isinstance(x,str) else int(x)
    return 'uint256{'+','.join(f'0x{(v>>(64*i))&((1<<64)-1):x}ULL' for i in range(4))+'}'

def expr(op,args,value='0x0'):
    a=args[0] if args else '';b=args[1] if len(args)>1 else ''
    binary={'ADD':'+','SUB':'-','MUL':'*','AND':'&','OR':'|','XOR':'^','EQ':'==','LT':'<','GT':'>'}
    if op in binary:return f'({a}{binary[op]}{b})'
    if op=='PUSH32':return word(value)
    if op in ['DIV','MOD','SDIV','SMOD']:
        fn='sdivrem' if op.startswith('S') else 'udivrem';field='rem' if op.endswith('MOD') else 'quot'
        return f'({b}!=0?intx::{fn}({a},{b}).{field}:uint256{{}})'
    if op in ['ADDMOD','MULMOD']:return f'({args[2]}!=0?intx::{op.lower()}({a},{b},{args[2]}):uint256{{}})'
    if op=='SLT':return f'signed_less({a},{b})'
    if op=='SGT':return f'signed_less({b},{a})'
    if op=='ISZERO':return f'({a}==0)'
    if op=='NOT':return f'(~{a})'
    if op=='SHL':return f'({b}<<{a})'
    if op=='SHR':return f'({b}>>{a})'
    if op=='SAR':return f'sar({a},{b})'
    if op=='SIGNEXTEND':return f'sext({a},{b})'
    if op=='BYTE':return f'({a}<32?({b}>>((31-{a}[0])*8))&255:uint256{{}})'
    if op=='EXP':return f'intx::exp({a},{b})'
    if op=='CALLDATASIZE':return 'uint256{s.msg->input_size}'
    if op=='CALLDATALOAD':return f'load(s.msg->input_data,s.msg->input_size,{a})'
    raise ValueError('missing C++ operation '+op)

def digest(x):return hashlib.sha256(json.dumps(x,sort_keys=True,separators=(',',':')).encode()).hexdigest()

class Emitter:
    def __init__(self):self.helpers={}
    def helper(self,h):
        name='check_'+digest(h)
        if name in self.helpers:return name
        lines=[f'static inline ArithmeticResult {name}(uint256 v0,uint256 v1) noexcept {{']
        if h['bits']<256:
            for v in ['v0','v1']:
                test=f'sext(uint256{{{h["bits"]//8-1}}},{v})!={v}' if h['signed'] else f'({v}>>{h["bits"]})!=0'
                lines.append(f'if({test})return {{}};')
        def value(i,seen):
            if i in seen:return
            n=h['nodes'][i];args=[] if n['op']=='PUSH32' else n['args'][:1 if n['op'] in ['ISZERO','NOT'] else 2]
            for a in args:value(a,seen)
            lines.append(f'const uint256 v{i} = uint256{{{expr(n["op"],[f"v{a}" for a in args],n["value"])}}};');seen.add(i)
        def decision(i,seen):
            v,z,n,path=h['decisions'][i]
            if path>=0:
                if h['costs'][path]<0:lines.append('return {};');return
                result=h['results'][path];value(result,seen)
                lines.append(f'return {{v{result},{h["costs"][path]},{h["steps"][path]},true}};');return
            value(v,seen);lines.append(f'if(v{v}==0){{');decision(z,set(seen));lines.append('}else{');decision(n,set(seen));lines.append('}')
        decision(h['root'],{0,1});lines.append('}');self.helpers[name]='\n'.join(lines);return name
    @staticmethod
    def signature(name):return f'RegionResult {name}(const Plan& p,ExecutionState& s,uint256* bottom,uint256* end,int64_t gas) noexcept {{'
    def arithmetic(self,name,x):
        h=self.helper(x['helper']);n=x['inputs'];low=min(c for c in x['helper']['costs'] if c>=0)
        source=lambda a:word(a['constant']) if a['index']<0 else f'in[{a["index"]}]'
        lines=[self.signature(name),f'const auto n=end-bottom;if(n<{n}||n-{n}+{x["peak"]}>1024||gas<{low})return {{}};',f'auto* in=end-{n};',f'const auto k={source(x["k"])};if(!jump(p,k))return {{}};',f'const auto q={h}({source(x["a"])},{source(x["b"])});if(!q.ok||uint64_t(gas)<q.cost)return {{}};']
        for i,v in enumerate(x['prefix']):lines.append(f'const auto o{i}={source(v)};')
        for i in range(len(x['prefix'])):lines.append(f'in[{i}]=o{i};')
        count=len(x['prefix']);lines += [f'in[{count}]=q.value;return {{p.patched->code().data()+k[0],in+{count+1},gas-int64_t(q.cost),q.instructions}};','}'];return '\n'.join(lines)
    def muldiv(self,name,x):
        arrays=''.join(f'constexpr uint64_t {k}[]={{'+','.join(map(str,x[k]))+'};' for k in ['costs','peaks','steps'])
        index='(q.wide?2:0)+(q.rounded?1:0)' if x['rounding'] else 'q.wide?1:0'
        return '\n'.join([self.signature(name),f'const auto n=end-bottom;if(n<4||gas<{min(x["costs"])})return {{}};auto* in=end-4;', 'const auto k=in[0];if(!jump(p,k))return {};',f'const auto q=muldiv(in[{x["a"]}],in[{x["b"]}],in[{x["d"]}],{str(x["rounding"]).lower()});if(!q.ok)return {{}};',arrays,f'const auto i={index};if(n-4+peaks[i]>1024||uint64_t(gas)<costs[i])return {{}};', 'in[0]=q.value;return {p.patched->code().data()+k[0],in+1,gas-int64_t(costs[i]),steps[i]};','}'])
    def hash(self,name,x):
        r=x['region'];n=r['Inputs'];lines=[self.signature(name),f'const auto n=end-bottom;if(n<{n}||n-{n}+{r["Peak"]}>1024)return {{}};auto* in=end-{n};',f'uint64_t size,total;if(!memory_cost(s,{r["MemoryEnd"]},{x["gas"]},gas,size,total))return {{}};']
        def source(v):
            k=v['Kind'];value=[f'in[{v["Slot"]}]',word(v['Word']),'intx::be::load<uint256>(s.msg->sender)','intx::be::load<uint256>(s.msg->recipient)','result'][k]
            return value if k==1 else f'({value}&{word(v["Mask"])})'
        for i,w in enumerate(r['Writes']):lines.append(f'const auto w{i}={source(w["Value"])};')
        lines.append('grow(s,size);')
        for i,w in enumerate(r['Writes']):lines.append(f'store(s,{w["Offset"]},w{i});')
        lines.append(f'const auto result=intx::be::load<uint256>(ethash::keccak256(&s.memory[{r["Offset"]}],{r["Size"]}));')
        for i,v in enumerate(r['Outputs']):lines.append(f'const auto o{i}={source(v)};')
        for i in range(len(r['Outputs'])):lines.append(f'in[{i}]=o{i};')
        lines += [f'return {{p.patched->code().data()+{r["Exit"]},in+{len(r["Outputs"])},gas-int64_t(total),{len(r["PCs"])}}};','}'];return '\n'.join(lines)
    def path(self,name,r,complete=False):
        r=dict(r,outputs=r['outputs'] or []);nodes=r['nodes'];used=set(r['outputs']);forward={int(k):v for k,v in (r['forward'] or {}).items()};dead={int(k) for k,v in (r['dead'] or {}).items() if v}
        if complete:used.update(range(32,32+len(nodes)))
        for i in reversed(range(len(nodes))):
            n=nodes[i]
            if 'helper' in n or n['op'] in ['JUMPI','MLOAD','MSTORE']:used.add(32+i)
            if i in forward:used.add(forward[i])
            if 32+i in used:used.update(n['args'] or [])
        count=sum(n['op']=='MSTORE' and 'helper' not in n for n in nodes)
        lines=['static inline '+self.signature(name),f'const auto n=end-bottom;if(n<{r["inputs"]}||n>{r["max_stack"]}||gas<{r["gas"]})return {{}};',f'StagedWrite writes[{max(1,count)}];size_t nw=0;uint64_t hi=0;']
        for i in sorted(used & set(range(32))):lines.append(f'const uint256 v{i}=end[-{32-i}];')
        for i,node in enumerate(nodes):
            v=f'v{32+i}';args=[f'v{a}' for a in (node['args'] or [])];op=node['op']
            if 32+i not in used:continue
            if 'helper' in node:
                fn=self.helper(node['helper']);lines += [f'const auto q{i}={fn}({args[0]},{args[1]});if(!q{i}.ok||q{i}.cost!={node["cost"]})return {{}};',f'const uint256 {v}=q{i}.value;'];continue
            if op=='JUMPI':lines.append(f'if({args[0]}{"==" if node["nonzero"] else "!="}0)return {{}};');continue
            if op in ['MLOAD','MSTORE']:
                lines += [f'if(!u64({args[0]})||{args[0]}[0]>max_memory-32)return {{}};',f'const auto o{i}={args[0]}[0];hi=std::max<uint64_t>(hi,o{i}+32);']
                if op=='MLOAD':lines.append(f'const uint256 {v}='+ (f'v{forward[i]}' if i in forward else f'staged_load(s,o{i},writes,nw)')+';')
                else:lines.append(f'writes[nw++]={{o{i},{args[1]}}};')
                continue
            lines.append(f'const uint256 {v} = uint256{{{expr(op,args,node["value"])}}};')
        lines += [f'uint64_t size,total;if(!memory_cost(s,hi,{r["gas"]},gas,size,total))return {{}};','grow(s,size);']
        for i in range(count):
            if i not in dead:lines.append(f'store(s,writes[{i}].offset,writes[{i}].value);')
        for i,v in enumerate(r['outputs']):lines.append(f'end[{i-r["drop"]}]=v{v};')
        lines += [f'return {{p.patched->code().data()+{r["exit"]},end+{len(r["outputs"])-r["drop"]},gas-int64_t(total),{r["instructions"]}}};','}'];return '\n'.join(lines)
    def kernel(self,name,x):
        family=x['family']
        if x.get('shared_branches'):
            from branch_codegen import emit_shared
            result=emit_shared(self,name,x)
            if result is None:raise ValueError('unmergeable branch artifact')
            return result
        if family=='effects':
            from effect_codegen import emit_effect
            return emit_effect(self,name,x)
        if family in ['arithmetic','muldiv','hash']:return getattr(self,family)(name,x)
        parts=[self.path(name+f'_p{i}',p) for i,p in enumerate(x['paths'])]
        body=[self.signature(name)]
        for i in range(len(parts)):body.append(f'if(auto result={name}_p{i}(p,s,bottom,end,gas);result.pc)return result;')
        return '\n'.join(parts+body+['return {};','}'])

def generate(source,directory):
    from terminal_paths import loop_intervals, exit_only_in_loop
    root=json.loads(Path(source).read_text());jobs={};families=['arithmetic','muldiv','hash','region','branch','pipeline','effects']
    excluded=[]
    for code in root['codes']:
        raw=bytes.fromhex(code['code'].removeprefix('0x'));intervals=loop_intervals(raw);keep=[]
        for r in code['regions']:
            x=r.get('native',{});exits=[p['exit'] for p in x.get('paths',[])]
            if x.get('family')=='branch' and exit_only_in_loop(raw,r['entry'],exits,intervals):
                excluded.append(dict(native_id=digest(x),entry=r['entry'],exits=exits,loop_intervals=intervals))
            else:keep.append(r)
        code['regions']=keep
        for r in code['regions']:
            if r['kind']!=5:continue
            x=r.pop('native');key=digest(x);r['native_id']=key;r['family_id']=families.index(x['family'])+1;jobs[key]=x
    directory.mkdir(exist_ok=True);(directory/'plans.json').write_text(json.dumps(root,separators=(',',':'))+'\n')
    (directory/'terminal-path-admission.json').write_text(json.dumps(excluded,indent=2)+'\n')
    # Keep translation-unit membership stable across additions. Rebuilding a
    # small new family must not recompile every existing native region.
    import re
    previous={}
    for f in sorted(directory.glob('native_*.cpp')):
        unit=int(f.stem.split('_')[1])
        for key in re.findall(r'^RegionResult native_([0-9a-f]{64})\(',f.read_text(),re.M):previous[key]=unit
    groups={}
    for key in sorted(jobs):
        if key in previous:groups.setdefault(previous[key],[]).append(key)
    fresh=[key for key in sorted(jobs) if key not in previous];start=max(groups,default=-1)+1
    for at in range(0,len(fresh),64):groups[start+at//64]=fresh[at:at+64]
    keys=sorted(jobs);files=[];registry=['#include "regions.hpp"','namespace evmone::regional {']
    for unit,selected in sorted(groups.items()):
        emitter=Emitter();bodies=[emitter.kernel('native_'+k,jobs[k]) for k in selected]
        content='\n'.join(['#include "native.hpp"','#include <evmone/instructions.hpp>','namespace evmone::regional {',*emitter.helpers.values(),*bodies,'}'])
        path=directory/f'native_{unit:04}.cpp';path.write_text(content);files.append(path)
        for k in selected:registry.append(Emitter.signature('native_'+k).removesuffix(' {')+';')
    registry += ['Native find_native(std::string_view key) {','static const std::pair<std::string_view,Native> entries[]={']
    registry += [f'{{"{k}",native_{k}}},' for k in keys]
    registry += ['{\"~\",nullptr},','};','size_t lo=0,hi=std::size(entries);while(lo<hi){auto mid=(lo+hi)/2;if(entries[mid].first<key)lo=mid+1;else hi=mid;}return lo<std::size(entries)&&entries[lo].first==key?entries[lo].second:nullptr;','}','}']
    path=directory/'registry.cpp';path.write_text('\n'.join(registry)+'\n');files.append(path)
    print('generated',len(keys),'native regions;',len(files),'translation units',flush=True);return files
