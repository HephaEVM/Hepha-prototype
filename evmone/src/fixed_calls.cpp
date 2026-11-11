// Native C++ control using the pinned upstream state host. No Go callbacks.
#include <test/state/host.hpp>
#include <evmone/evmone.h>
#include <nlohmann/json.hpp>
#include <chrono>
#include <fstream>
#include <iostream>
#include <stdexcept>
#include <unordered_map>
#include <cstring>
#include <sstream>
using json=nlohmann::json;
using namespace evmc;
using namespace evmone::state;
extern "C" void* icde_cache_create(int);
extern "C" void icde_cache_destroy(void*);
extern "C" void* icde_cache_prepare(void*,const uint8_t*,size_t);
extern "C" void icde_cache_execute(void*,evmc_vm*,void*,const evmc_host_interface*,evmc_host_context*,evmc_revision,const evmc_message*,evmc_result*);
extern "C" void icde_regions_diagnostic(void*,int);
extern "C" void icde_regions_counters(void*,uint64_t*);
extern "C" void icde_regions_families(void*,uint64_t*);
static bytes decode(std::string s) {if(s.starts_with("0x"))s.erase(0,2);if(s.size()%2)s="0"+s;bytes b;for(size_t i=0;i<s.size();i+=2)b.push_back(static_cast<uint8_t>(std::stoul(s.substr(i,2),nullptr,16)));return b;}
template<typename T> static T word(const std::string& s){T x{};auto b=decode(s);if(b.size()>sizeof(x))throw std::runtime_error("oversized word");std::memcpy(reinterpret_cast<uint8_t*>(&x)+sizeof(x)-b.size(),b.data(),b.size());return x;}
static intx::uint256 number(const std::string& s){return intx::be::load<intx::uint256>(word<bytes32>(s));}
static std::string hexbytes(bytes_view b){return evmc::hex(b);}
static void require(bool yes,const std::string& message){if(!yes)throw std::runtime_error(message);}
struct View:StateView,BlockHashes {
 const json& f;std::unordered_map<address,json> accounts;std::unordered_map<int64_t,bytes32> hashes;
 explicit View(const json& data):f(data){for(auto it=f["accounts"].begin();it!=f["accounts"].end();++it)accounts.emplace(word<address>(it.key()),it.value());for(auto it=f["block_hashes"].begin();it!=f["block_hashes"].end();++it)hashes[std::stoll(it.key())]=word<bytes32>(it.value());}
 const json& checked(const address& a)const noexcept {auto i=accounts.find(a);if(i==accounts.end()){std::cerr<<"unexported account "<<evmc::hex(a)<<'\n';std::exit(3);}return i->second;}
 std::optional<StateView::Account> get_account(const address& a)const noexcept override {const auto& x=checked(a);if(!x["exists"].get<bool>())return {};return StateView::Account{x["nonce"].get<uint64_t>(),number(x["balance"]),word<bytes32>(x["code_hash"]),x["has_storage"]};}
 bytes get_account_code(const address& a)const noexcept override {return decode(checked(a)["code"]);}
 bytes32 get_storage(const address& a,const bytes32& k)const noexcept override {const auto& x=checked(a)["storage"];std::string key="0x"+evmc::hex(k);if(!x.contains(key)){std::cerr<<"unexported storage "<<evmc::hex(a)<<" "<<key<<'\n';std::exit(4);}return word<bytes32>(x[key]["current"]);}
 bytes32 get_block_hash(int64_t n)const noexcept override {auto i=hashes.find(n);if(i==hashes.end()){std::cerr<<"unexported blockhash "<<n<<'\n';std::exit(5);}return i->second;}
};
struct CachedVM:evmc_vm {
 evmc::VM engine{evmc_create_evmone()};void* cache;std::unordered_map<std::string,void*> content;std::unordered_map<const uint8_t*,std::pair<size_t,void*>> pointers;
 bool strict=false;uint64_t misses=0;int64_t root_ns=0;
 explicit CachedVM(int adv):evmc_vm{EVMC_ABI_VERSION},cache(icde_cache_create(adv)){
  name="evmone-native-cached";version="0.25.0";destroy=[](evmc_vm* v){delete static_cast<CachedVM*>(v);};
  set_option=nullptr;
  execute=[](evmc_vm* v,const evmc_host_interface* h,evmc_host_context* c,evmc_revision rev,const evmc_message* m,const uint8_t* code,size_t size){
   auto& self=*static_cast<CachedVM*>(v);auto start=std::chrono::steady_clock::now();
   void* plan=nullptr;auto it=self.pointers.find(code);
   if(it!=self.pointers.end()&&it->second.first==size)plan=it->second.second;
   else plan=self.prepare(code,size,false);
   evmc_result result{};icde_cache_execute(self.cache,self.engine.get_raw_pointer(),plan,h,c,rev,m,&result);
   if(m->depth==0)self.root_ns+=std::chrono::duration_cast<std::chrono::nanoseconds>(std::chrono::steady_clock::now()-start).count();
   return result;
  };
 }
 ~CachedVM(){icde_cache_destroy(cache);}
 void* prepare(const uint8_t* code,size_t n,bool stable){std::string key(reinterpret_cast<const char*>(code),n);auto i=content.find(key);void* p;
  if(i==content.end()){if(strict){std::cerr<<"timed analysis miss\n";std::exit(6);}++misses;p=icde_cache_prepare(cache,code,n);content.emplace(std::move(key),p);}else p=i->second;
  if(stable)pointers[code]={n,p};return p;
 }
};
static void seed(State& state,const View& view,CachedVM& vm){vm.pointers.clear();for(const auto& [a,x]:view.accounts){auto& acc=state.get(a);acc.access_status=x["warm"]?EVMC_ACCESS_WARM:EVMC_ACCESS_COLD;
 if(x["exists"].get<bool>()){auto code=state.get_code(a);if(!code.empty())vm.prepare(code.data(),code.size(),true);}
 if(x.contains("warm_slots"))for(auto it=x["warm_slots"].begin();it!=x["warm_slots"].end();++it)state.access_storage(a,word<bytes32>(it.key())).access_status=it.value().get<bool>()?EVMC_ACCESS_WARM:EVMC_ACCESS_COLD;
 for(auto it=x["storage"].begin();it!=x["storage"].end();++it){auto& slot=state.access_storage(a,word<bytes32>(it.key()));slot.current=word<bytes32>(it.value()["current"]);slot.original=word<bytes32>(it.value()["original"]);slot.access_status=it.value()["warm"]?EVMC_ACCESS_WARM:EVMC_ACCESS_COLD;slot.loaded=true;}
 for(auto it=x["transient"].begin();it!=x["transient"].end();++it)acc.transient_storage[word<bytes32>(it.key())]=word<bytes32>(it.value());
}}
static void check_state(State& st,evmone::state::Host& host,const json& f){
 for(auto it=f["post"].begin();it!=f["post"].end();++it){auto a=word<address>(it.key());auto& x=it.value();auto& acc=st.get(a);auto prefix=it.key()+" ";
  // Upstream may keep a newly touched empty placeholder until EIP-161 cleanup.
  // Geth skips creating it for a zero-value call. Permit only an explicitly
  // erasable empty account when Geth says absent; all value fields still match.
  const bool absent_equivalent=acc.nonexistent||(acc.erase_if_empty&&acc.is_empty());
  require(x["exists"].get<bool>()?!acc.nonexistent:absent_equivalent,prefix+"existence");require(acc.balance==number(x["balance"]),prefix+"balance");require(acc.nonce==x["nonce"].get<uint64_t>(),prefix+"nonce");require(st.get_code(a)==decode(x["code"]),prefix+"code");require(acc.destructed==x["destructed"].get<bool>(),prefix+"destructed");
  for(auto k=x["storage"].begin();k!=x["storage"].end();++k)require(st.get_storage(a,word<bytes32>(k.key())).current==word<bytes32>(k.value()),prefix+"storage "+k.key());
  for(auto k=x["transient"].begin();k!=x["transient"].end();++k)require(acc.transient_storage[word<bytes32>(k.key())]==word<bytes32>(k.value()),prefix+"transient");
 }
 auto logs=host.take_logs();require(logs.size()==f["logs"].size(),"log count");for(size_t i=0;i<logs.size();i++){auto& want=f["logs"][i];require(logs[i].addr==word<address>(want["address"]),"log address");require(logs[i].data==decode(want["data"]),"log data");require(logs[i].topics.size()==want["topics"].size(),"log topics");for(size_t j=0;j<logs[i].topics.size();j++)require(logs[i].topics[j]==word<bytes32>(want["topics"][j]),"log topic");}
 // No unaccounted mutations may disappear merely because the export was partial.
 const auto diff=st.build_diff(EVMC_PRAGUE);
 for(const auto& a:diff.deleted_accounts){const auto key="0x"+evmc::hex(a);require(f["post"].contains(key),"unexpected deletion "+key);const auto& want=f["post"][key];require(!want["exists"].get<bool>()||want["destructed"].get<bool>()||(number(want["balance"])==0&&want["nonce"].get<uint64_t>()==0&&want["code"].get<std::string>().empty()),"unexpected nonempty deletion");}
 for(const auto& change:diff.modified_accounts){const auto key="0x"+evmc::hex(change.addr);require(f["post"].contains(key),"unexpected modified account "+key);for(const auto& [k,v]:change.modified_storage){const auto sk="0x"+evmc::hex(k);require(f["post"][key]["storage"].contains(sk),"unexpected modified storage "+sk);require(word<bytes32>(f["post"][key]["storage"][sk])==v,"diff storage");}}
}
int main(int argc,char** argv){try{
 if(argc<2)throw std::runtime_error("usage: native-host FIXTURE [rounds] [iterations]");json f;std::ifstream(argv[1])>>f;for(const char* part:{"accounts","post"}){json norm=json::object();for(auto it=f[part].begin();it!=f[part].end();++it)norm["0x"+evmc::hex(word<address>(it.key()))]=it.value();f[part]=norm;}View view(f);
 const int rounds=argc>2?std::stoi(argv[2]):8,iterations=argc>3?std::stoi(argv[3]):64;
 BlockInfo block;auto& b=f["block"];block.number=b["number"];block.timestamp=b["time"];block.gas_limit=b["gas_limit"];block.coinbase=word<address>(b["coinbase"]);block.prev_randao=word<bytes32>(b["randao"]);block.base_fee=static_cast<uint64_t>(number(b["base_fee"]));block.blob_base_fee=number(b["blob_base_fee"]);block.chain_id=1;
 Transaction tx;tx.sender=word<address>(b["origin"]);tx.max_gas_price=number(b["gas_price"]);tx.max_priority_gas_price=tx.max_gas_price;for(const auto& x:b["blob_hashes"])tx.blob_hashes.push_back(word<bytes32>(x));
 auto input=decode(f["input"]),expected=decode(f["output"]);evmc_message msg{};msg.sender=word<address>(f["from"]);msg.recipient=word<address>(f["to"]);msg.code_address=msg.recipient;msg.value=word<uint256be>(f["value"]);msg.gas=f["gas"];msg.input_data=input.data();msg.input_size=input.size();
 const auto rootcode=view.get_account_code(msg.recipient);if(rootcode.size()==23&&rootcode[0]==0xef&&rootcode[1]==1&&rootcode[2]==0){std::memcpy(&msg.code_address,rootcode.data()+3,20);msg.flags|=EVMC_DELEGATED;}
 evmc::VM vms[7];CachedVM* caches[7];int modes[]={0,2,4,5,7,6,3};for(int i=0;i<7;i++){caches[i]=new CachedVM(modes[i]);vms[i]=evmc::VM{caches[i]};} const char* names[]={"baseline","copy-abi","stage1","stage2","branch-fixed","stage3","bind-only"}; uint64_t counters[7][5]{},families[7][8]{};
 auto run=[&](int mode,bool check){auto& cache=*caches[mode];State st(view);seed(st,view,cache);evmone::state::Host host(EVMC_PRAGUE,vms[mode],st,block,view,tx);cache.root_ns=0;auto result=host.call(msg);
  const int status=result.status_code==EVMC_SUCCESS?0:result.status_code==EVMC_REVERT?2:1;
  require(status==f["status"].get<int>(),"status "+std::to_string(result.status_code));require(result.gas_left==f["gas_left"].get<int64_t>(),"gas "+std::to_string(result.gas_left)+" expected "+std::to_string(f["gas_left"].get<int64_t>()));require(result.gas_refund==f["refund"].get<int64_t>(),"refund");require(result.output_size==expected.size()&&(!result.output_size||!std::memcmp(result.output_data,expected.data(),expected.size())),"output");if(check)check_state(st,host,f);return cache.root_ns;};
 for(int m : {0,5}){run(m,true);icde_regions_diagnostic(caches[m]->cache,1);run(m,true);icde_regions_counters(caches[m]->cache,counters[m]);icde_regions_families(caches[m]->cache,families[m]);icde_regions_diagnostic(caches[m]->cache,0);caches[m]->strict=true;}
 std::vector<int> selected;std::istringstream selection(argc>4?argv[4]:"baseline,copy-abi,stage1,stage2,branch-fixed,stage3,bind-only");std::string item;while(std::getline(selection,item,',')){bool found=false;for(int m=0;m<7;m++)if(item==names[m]){selected.push_back(m);found=true;}require(found,"unknown measurement mode");}require(!selected.empty(),"empty modes");
 for(int round=0;round<rounds;round++)for(int order=0;order<int(selected.size());order++){int m=selected[(round+order)%selected.size()];run(m,false);int64_t total=0;for(int i=0;i<iterations;i++)total+=run(m,i==0||i+1==iterations);json row={{"hash",f["hash"]},{"mode",names[m]},{"round",round},{"iterations",iterations},{"frame_ns",double(total)/iterations},{"analyses",caches[m]->misses},{"attempts",counters[m][0]},{"commits",counters[m][1]},{"fallback",counters[m][2]},{"instructions",counters[m][3]},{"bytes",counters[m][4]},{"families",families[m]},{"scope","root execute; native upstream state host; no Go"}};std::cout<<row.dump()<<'\n';}
 return 0;
}catch(const std::exception& e){std::cerr<<"native fixture failed: "<<e.what()<<'\n';return 1;}}
