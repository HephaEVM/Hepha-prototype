#!/usr/bin/env python3
"""Execute native checkpoints with poisoned new stack slots, including effects."""
import json,subprocess,tempfile,unittest
from pathlib import Path
from codegen import Emitter
from effect_codegen import emit_effect

class Checkpoints(unittest.TestCase):
    def test_new_slots_and_committed_memory(self):
        n=Path(__file__).resolve().parents[1];d=n/'vendor';s=d/'evmone'
        nodes=[dict(op='input',depth=32-i,a=0,b=0,value='0x0') for i in range(32)]
        nodes += [dict(op='constant',depth=0,a=0,b=0,value=hex(i)) for i in [9,0,16]]
        before=list(range(32))+[32,31,34]
        action=dict(op='JUMPI',args=[31],gas=17,pc=6,start=0,start_step=0,op_step=3,result=-1,reuse=-1,before=before,checkpoint=list(range(32)))
        x=dict(nodes=nodes,actions=[action],events=[[i,-1] for i in range(35)]+[[-1,0]],inputs=1,peak=3,writes=[[0,32]],delta=1,exit=10,instructions=4)
        a=emit_effect(Emitter(),'guard',x)
        x['actions']=[dict(op='MSTORE',args=[32,33],gas=3,pc=2,start=0,start_step=0,op_step=2,result=-1,reuse=-1,before=list(range(32))+[32,33],checkpoint=list(range(32))),dict(action,start=5,start_step=1,checkpoint=list(range(32))+[32])]
        x['events'][-1:]=[[-1,0],[-1,1]];b=emit_effect(Emitter(),'gas_resume',x)
        source='\n'.join(['#include "native.hpp"','#include <evmone/instructions.hpp>','#include <stdexcept>','namespace evmone::regional {',a,b,'}',r'''
int main(){using namespace evmone;using namespace evmone::regional;
 Plan p;p.original.resize(20);p.plain=std::make_unique<baseline::CodeAnalysis>(baseline::analyze(p.original));p.patched=std::make_unique<baseline::CodeAnalysis>(baseline::analyze(p.original));ExecutionState s;
 alignas(32) uint256 stack[16];for(auto& v:stack)v=0xdead;stack[0]=0;auto* end=stack+1;
 auto r=guard(p,s,stack,end,1000);
 if(!r.original||r.pc!=p.patched->code().data()+6||r.end!=end+3||r.gas!=993||end[0]!=9||end[1]!=0||end[2]!=16)throw std::runtime_error("new slot omitted at conditional continuation");
 for(auto& v:stack)v=0xdead;stack[0]=0;
 r=gas_resume(p,s,stack,end,10);
 if(!r.original||r.pc!=p.patched->code().data()+5||r.end!=end+1||r.gas!=4||end[0]!=9||s.memory.size()!=32||s.memory[31]!=9)throw std::runtime_error("low-gas continuation lost completed effect or new slot");
 r=guard(p,s,stack,end,0);
 if(!r.original||r.pc!=p.patched->code().data()||r.end!=end||r.gas!=0)throw std::runtime_error("entry continuation changed gas or stack");
}
'''])
        with tempfile.TemporaryDirectory(prefix='evmone-checkpoint-',dir=None) as tmp:
            cpp=Path(tmp)/'test.cpp';binary=Path(tmp)/'test';cpp.write_text(source)
            inc=[n/'src',s/'lib',s/'evmc/include',d/'intx/include']
            cmd=['clang++','-O2','-std=c++20','-Wno-unknown-attributes',str(cpp),*['-I'+str(x) for x in inc],str(s/'lib/evmone/baseline_analysis.cpp'),'-o',str(binary)]
            subprocess.run(cmd,check=True,capture_output=True);subprocess.run([str(binary)],check=True)
if __name__=='__main__':unittest.main()
