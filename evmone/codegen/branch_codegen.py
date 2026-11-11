"""Merge identical SSA prefixes before branch selection; no path prediction."""
import re

def _emit_shared(emitter,name,x,complete):
    paths=x['paths']
    if len(paths)<2:return None
    bodies=[];inputs=set();capacity=1
    for i,p in enumerate(paths):
        lines=emitter.path(name+f'_p{i}',p,complete=complete).splitlines()
        cap=re.search(r'writes\[(\d+)\]',lines[2]);capacity=max(capacity,int(cap[1]))
        body=[]
        for line in lines[3:-1]:
            if re.fullmatch(r'const uint256 v\d+=end\[-\d+\];',line):inputs.add(line)
            else:body.append(line)
        bodies.append(body)
    guard=re.compile(r'if\((v\d+)(==|!=)0\)return \{\};')
    splits=0
    def tree(parts):
        nonlocal splits
        if len(parts)==1:return parts[0]
        n=0
        while all(n<len(p) for p in parts) and len({p[n] for p in parts})==1:n+=1
        out=parts[0][:n];parts=[p[n:] for p in parts]
        if all(not p for p in parts):return out
        if any(not p for p in parts):return None
        matches=[guard.fullmatch(p[0]) for p in parts]
        if not all(matches) or len({m[1] for m in matches})!=1:return None
        zero=[p[1:] for p,m in zip(parts,matches) if m[2]=='!=']
        nonzero=[p[1:] for p,m in zip(parts,matches) if m[2]=='==']
        if not zero or not nonzero:return None
        a,b=tree(zero),tree(nonzero)
        if a is None or b is None:return None
        splits+=1
        return out+[f'if({matches[0][1]}==0){{']+a+['}else{']+b+['}']
    body=tree(bodies)
    if body is None or not splits:return None
    head=[emitter.signature(name),f'const auto n=end-bottom;if(n<{max(p["inputs"] for p in paths)}||n>{min(p["max_stack"] for p in paths)}||gas<{min(p["gas"] for p in paths)})return {{}};',f'StagedWrite writes[{capacity}];size_t nw=0;uint64_t hi=0;',*sorted(inputs)]
    return '\n'.join(head+body+['}'])

def emit_shared(emitter,name,x):
    return _emit_shared(emitter,name,x,False) or _emit_shared(emitter,name,x,True)
