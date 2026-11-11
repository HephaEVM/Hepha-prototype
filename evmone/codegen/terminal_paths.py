"""Avoid repeatedly admitting loop-exit-only branch summaries in a loop body.

A literal backwards JUMP/JUMPI conservatively marks a loop interval. A branch
summary within it that can only terminate the frame cannot replace a normal
iteration. Leave those instructions to Vanilla, retaining other summaries at
the same PC. This is admission policy, not a bytecode/semantic transformation.
"""
def loop_intervals(code):
    pc=0;previous=None;intervals=[];jumpdests=set()
    while pc<len(code):
        op=code[pc];end=pc+1;literal=None
        if 0x60<=op<=0x7f:
            end+=op-0x5f
            if end<=len(code):literal=int.from_bytes(code[pc+1:end],'big')
        elif op==0x5f:literal=0
        if op==0x5b:jumpdests.add(pc)
        if op in (0x56,0x57) and previous is not None:
            if previous<pc and previous in jumpdests:intervals.append((previous,pc))
        previous=literal;pc=end
    return intervals

def exit_only_in_loop(code,entry,exits,intervals=None):
    intervals=loop_intervals(code) if intervals is None else intervals
    return bool(exits) and all(0<=x<len(code) and code[x] in (0x00,0xf3,0xfd,0xfe) for x in exits) and any(lo<=entry<=hi for lo,hi in intervals)

