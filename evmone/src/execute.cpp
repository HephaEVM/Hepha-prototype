#include "native.hpp"
#include <cstring>
#include <algorithm>
namespace evmone::regional {
uint256 Clean::apply(uint256 v)const {
 if(kind==1)return v&arg;
 if(kind==2 && arg<31){const auto bit=static_cast<unsigned>(arg[0]*8+7);const auto mask=(uint256{1}<<bit)-1;return ((v>>bit)&1)?v|~mask:v&mask;}
 if(kind==3)return v!=0?uint256{1}:uint256{};
 return v;
}
static bool try_region(const Region& r,const Plan& p,ExecutionState& s,uint256* bottom,uint256*& end,int64_t& gas,const uint8_t*& pc) noexcept {
 const auto height=end-bottom;if(p.diagnostic)++p.stats.attempts;
 if(p.observe||height<r.inputs||height-r.inputs+r.peak>1024||gas<0||uint64_t(gas)<r.cost)return false;
 auto* in=end-r.inputs;uint64_t dest=r.exit;const auto* code=p.patched->code().data();
 if(r.kind!=4){if(!u64(in[0])||!p.patched->check_jumpdest(in[0][0]))return false;dest=in[0][0];}
 uint64_t size=0,total=0,bytes_copied=0,count=r.instructions;
 if(r.kind==4){
  if(!(in[3]<in[0])){gas-=r.cost;pc=code+dest;}
  else {
   for(int i=0;i<4;i++)if(!u64(in[i]))return false;
   const auto offset=in[3][0],limit=in[0][0],words=(limit-offset)/32+((limit-offset)%32!=0);
   if(words>(1<<18))return false;const auto length=words*32;
   if(offset>max_memory||length>max_memory-offset||in[1][0]>max_memory-offset-length||in[2][0]>max_memory-offset-length)return false;
   const auto src=in[2][0]+offset,dst=in[1][0]+offset;
   if(src!=dst&&src<dst+length&&dst<src+length)return false;
   if(r.loop_cost && words>(uint64_t(gas)-r.cost)/r.loop_cost)return false;
   if(!memory_cost(s,std::max(src+length,dst+length),r.cost+words*r.loop_cost,gas,size,total))return false;
   grow(s,size);if(src!=dst)std::memcpy(&s.memory[dst],&s.memory[src],length);
   in[3]=offset+length;gas-=total;pc=code+dest;bytes_copied=length;count+=words*r.loop_instructions;
  }
 } else if(r.kind==3){
  if(r.inputs!=4||!u64(in[1])||in[1][0]>4096*32||!u64(in[2])||in[2][0]>max_memory-in[1][0]-32)return false;
  const auto len=in[1][0],dst=in[2][0],words=(len+31)/32;uint64_t src=0,hi=dst+len+32;
  if(words){if(!u64(in[3])||in[3][0]>max_memory-words*32)return false;src=in[3][0];hi=std::max(hi,src+words*32);}
  if(r.loop_cost && words>(uint64_t(gas)-r.cost)/r.loop_cost)return false;
  if(!memory_cost(s,hi,r.cost+words*r.loop_cost,gas,size,total))return false;
  grow(s,size);for(uint64_t i=0;i<words;i++){std::array<uint8_t,32> v;std::memcpy(v.data(),&s.memory[src+i*32],32);std::memcpy(&s.memory[dst+i*32],v.data(),32);}
  std::memset(&s.memory[dst+len],0,32);end=in;gas-=total;pc=code+dest;bytes_copied=words*32;count+=words*r.loop_instructions;
 } else {
  std::array<uint256,18> values{};std::array<uint64_t,18> offsets{};uint64_t hi=0;
  if(r.kind==1){
   for(const auto& b:r.bounds)if(signed_less(in[b.end]-in[b.start],b.size))return false;
   for(size_t i=0;i<r.reads.size();i++){const auto& rd=r.reads[i];const auto off=rd.offset.eval(in);
    if(rd.memory){if(!u64(off)||off[0]>max_memory-32)return false;hi=std::max(hi,off[0]+32);values[i]=load(s.memory.size()?&s.memory[0]:nullptr,s.memory.size(),off);}
    else values[i]=load(s.msg->input_data,s.msg->input_size,off);
    if(rd.clean.apply(values[i])!=values[i])return false;
   }
  } else if(r.kind==2){
   for(size_t i=0;i<r.writes.size();i++){const auto off=r.writes[i].offset.eval(in);if(!u64(off)||off[0]>max_memory-32)return false;offsets[i]=off[0];hi=std::max(hi,off[0]+32);}
   for(size_t i=0;i<r.outputs.size();i++)values[i]=r.outputs[i].eval(in);
  } else return false;
  if(!memory_cost(s,hi,r.cost,gas,size,total))return false;
  grow(s,size);
  if(r.kind==2)for(size_t i=0;i<r.writes.size();i++){const auto& w=r.writes[i];const auto v=w.clean.apply(in[w.input]);intx::be::store(std::span<uint8_t,32>{&s.memory[offsets[i]],32},v);}
  if(r.kind==1){for(size_t i=0;i<r.read_outputs.size();i++)in[i]=values[r.read_outputs[i]];end=in+r.read_outputs.size();}
  else {std::copy_n(values.data(),r.outputs.size(),in);end=in+r.outputs.size();}
  gas-=total;pc=code+dest;
 }
 if(p.diagnostic){++p.stats.commits;++p.stats.families[0];p.stats.instructions+=count;p.stats.bytes+=bytes_copied;}return true;
}
RegionResult run_region(const Region& r,const Plan& p,ExecutionState& s,uint256* bottom,uint256* end,int64_t gas) noexcept {
 if(r.kind==5){
  if(p.diagnostic)++p.stats.attempts;
  if(!p.observe){auto result=r.native(p,s,bottom,end,gas);if(s.status!=EVMC_SUCCESS)return result;if(result.pc){if(p.diagnostic){++p.stats.commits;++p.stats.families[r.family];p.stats.instructions+=result.instructions;}return result;}}
 }else{
 const uint8_t* pc=nullptr;
 if(try_region(r,p,s,bottom,end,gas,pc))return {pc,end,gas};
 }
 if(r.next&&!p.observe)return run_region(*r.next,p,s,bottom,end,gas);
 return {nullptr,end,gas};
}
}
