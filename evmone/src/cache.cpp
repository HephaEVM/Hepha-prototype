// Offline recipes are trusted build artifacts from our existing verified miner.
// They bind by complete bytecode, never by address or a short bytecode prefix.
#include "regions.hpp"
#include <nlohmann/json.hpp>
#include <fstream>
#include <unordered_map>
#include <stdexcept>
#include <cstdlib>
namespace {
using namespace evmone;
using namespace evmone::regional;
using json=nlohmann::json;
void require(bool ok){if(!ok)throw std::runtime_error("invalid region recipe");}
bytes decode(std::string s){if(s.starts_with("0x"))s.erase(0,2);require(s.size()%2==0);bytes b;for(size_t i=0;i<s.size();i+=2)b.push_back(std::stoul(s.substr(i,2),nullptr,16));return b;}
uint256 number(const std::string& s){return intx::from_string<uint256>(s.c_str());}
Affine affine(const json& x,int n){int i=x.at("input");require(i>=-1&&i<n);return {i,number(x.at("add"))};}
Clean clean(const json& x){int k=x.at("kind");require(k>=0&&k<=3);return {k,number(x.at("argument"))};}
Region parse(const json& x){
 Region r;r.kind=x.at("kind");r.inputs=x.at("inputs");r.peak=x.at("peak");r.entry=x.at("entry");r.exit=x.at("exit");r.cost=x.at("cost");r.loop_cost=x.at("loop_cost");r.instructions=x.at("instructions");r.loop_instructions=x.at("loop_instructions");
 r.stage=x.value("stage",0);
 if(r.kind==5){r.native=find_native(x.at("native_id").get<std::string>());r.family=x.at("family_id");require(r.native&&r.stage>=1&&r.stage<=3&&r.family>=1&&r.family<=7);return r;}
 require(r.kind>=1&&r.kind<=4&&r.inputs>0&&r.inputs<=32&&r.peak>=r.inputs&&r.peak<=1024&&r.cost<=INT64_MAX&&r.loop_cost<=INT64_MAX);
 for(const auto& v:x.at("bounds")){int e=v.at("end"),s=v.at("start");require(e>=0&&e<r.inputs&&s>=0&&s<r.inputs);r.bounds.push_back({e,s,number(v.at("size"))});}
 for(const auto& v:x.at("reads"))r.reads.push_back({affine(v.at("offset"),r.inputs),v.at("memory"),clean(v.at("clean"))});
 for(const auto& v:x.at("writes")){int i=v.at("input");require(i>=0&&i<r.inputs);r.writes.push_back({affine(v.at("offset"),r.inputs),i,clean(v.at("clean"))});}
 for(const auto& v:x.at("outputs"))r.outputs.push_back(affine(v,r.inputs));
 for(const auto& v:x.at("read_outputs")){int i=v;require(i>=0&&size_t(i)<r.reads.size());r.read_outputs.push_back(i);}
 require(r.reads.size()<=18&&r.writes.size()<=18&&r.outputs.size()<=18&&r.read_outputs.size()<=18);
 require(r.outputs.size()<=size_t(r.peak)&&r.read_outputs.size()<=size_t(r.peak));
 require(r.kind<3||r.inputs==4);return r;
}
const auto& database(){
 static const auto db=[](){std::unordered_map<std::string,std::vector<Region>> out;const char* file=std::getenv("ICDE_REGIONS");if(!file)throw std::runtime_error("ICDE_REGIONS must name an exported recipe file");json root;std::ifstream(file)>>root;require(root.at("version")==1&&root.at("fork")=="Prague");for(const auto& x:root.at("codes")){auto b=decode(x.at("code"));auto& list=out[std::string(reinterpret_cast<const char*>(b.data()),b.size())];for(const auto& r:x.at("regions"))list.push_back(parse(r));}return out;}();return db;
}
struct Cache {int mode;bool diagnostic=false;std::vector<std::unique_ptr<Plan>> plans;};
std::unique_ptr<CodeAnalysis> overlay(Plan& p){
 const auto n=p.original.size(),aligned=(n+33+7)&~size_t{7},words=(n+64)/64,total=aligned+words*8;
 auto data=std::make_unique<uint8_t[]>(total);std::copy(p.original.begin(),p.original.end(),data.get());
 auto* bits=new(data.get()+aligned) BitsetSpan::word_type[words];BitsetSpan map(bits);
 std::vector<bool> boundaries(n);for(size_t i=0;i<n;i++){boundaries[i]=true;auto op=p.original[i];if(op>=0x60&&op<=0x7f)i+=op-0x5f;else if(op==0x5b)map.set(i);}
 p.entries.resize(n);if(p.adaptive)p.rejected.resize(n);
 for(auto it=p.regions.rbegin();it!=p.regions.rend();++it){auto& r=*it;require(r.entry<n&&boundaries[r.entry]&&p.original[r.entry]!=0x0c);if(r.kind==4)require(p.plain->check_jumpdest(r.exit));r.next=p.entries[r.entry];p.entries[r.entry]=&r;data[r.entry]=0x0c;}
 return std::make_unique<CodeAnalysis>(std::move(data),n,map);
}
}
extern "C" {
void* icde_cache_create(int mode){require(mode>=0&&mode<=12);if(mode!=0&&mode!=8&&mode!=10)database();return new Cache{mode};}
void icde_cache_destroy(void* p){delete static_cast<Cache*>(p);}
void* icde_cache_prepare(void* c,const uint8_t* code,size_t n){auto& cache=*static_cast<Cache*>(c);auto p=std::make_unique<Plan>();p->original.assign(code,code+n);p->plain=std::make_unique<CodeAnalysis>(baseline::analyze(p->original));p->observe=cache.mode==3;p->diagnostic=cache.diagnostic;
 p->stream=cache.mode==10||cache.mode==11?2:cache.mode==8||cache.mode==9?1:0;
 p->adaptive=cache.mode==12;
 if(cache.mode&&cache.mode!=8&&cache.mode!=10){auto i=database().find(std::string(reinterpret_cast<const char*>(code),n));if(i!=database().end())for(const auto& r:i->second)if((cache.mode!=1||r.kind==4)&&(cache.mode!=2||r.stage==0)&&(cache.mode!=4||r.stage<=1)&&(cache.mode!=5||r.stage<=2)&&(cache.mode!=7||r.family!=7))p->regions.push_back(r);}
 if(!p->regions.empty())p->patched=overlay(*p);auto raw=p.get();cache.plans.push_back(std::move(p));return raw;
}
// Called only at a block boundary. Reuse the already analyzed original code;
// negative matches retain the same plan and do not allocate another bytecode.
void* icde_cache_upgrade(void* c,void* existing){
 auto& cache=*static_cast<Cache*>(c);auto& p=*static_cast<Plan*>(existing);
 require(cache.mode==6||cache.mode==9||cache.mode==11||cache.mode==12);
 if(p.patched)return existing;
 p.adaptive=cache.mode==12;
 auto i=database().find(std::string(reinterpret_cast<const char*>(p.original.data()),p.original.size()));
 if(i!=database().end())p.regions=i->second;
 if(!p.regions.empty())p.patched=overlay(p);
 return existing;
}
void icde_cache_execute(void*,evmc_vm* v,void* plan,const evmc_host_interface* h,evmc_host_context* c,evmc_revision rev,const evmc_message* m,evmc_result* out){auto& p=*static_cast<Plan*>(plan);auto& vm=*static_cast<VM*>(v);if((p.patched||p.stream)&&rev==EVMC_PRAGUE&&!vm.get_tracer())*out=regional::execute(vm,*h,c,rev,*m,p.patched?*p.patched:*p.plain,p);else *out=baseline::execute(vm,*h,c,rev,*m,*p.plain);}
void icde_regions_diagnostic(void* c,int enabled){auto& cache=*static_cast<Cache*>(c);cache.diagnostic=enabled;for(auto& p:cache.plans){p->diagnostic=enabled;p->stats={};}}
void icde_regions_counters(void* c,uint64_t* out){std::fill_n(out,5,0);for(const auto& p:static_cast<Cache*>(c)->plans){out[0]+=p->stats.attempts;out[1]+=p->stats.commits;out[2]+=p->stats.fallback;out[3]+=p->stats.instructions;out[4]+=p->stats.bytes;}}
void icde_regions_families(void* c,uint64_t* out){std::fill_n(out,8,0);for(const auto& p:static_cast<Cache*>(c)->plans)for(size_t i=0;i<8;i++)out[i]+=p->stats.families[i];}
}
