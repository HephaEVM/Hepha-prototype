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
using json=nlohmann::json;
using namespace evmc;
using namespace evmone::state;
extern "C" void* icde_cache_create(int);
extern "C" void icde_cache_destroy(void*);
extern "C" void* icde_cache_prepare(void*,const uint8_t*,size_t);
extern "C" void icde_cache_execute(void*,evmc_vm*,void*,const evmc_host_interface*,evmc_host_context*,evmc_revision,const evmc_message*,evmc_result*);
extern "C" void* icde_cache_upgrade(void*,void*);
extern "C" void icde_regions_diagnostic(void*,int);
extern "C" void icde_regions_counters(void*,uint64_t*);
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
 void* region_cache=nullptr;std::vector<void*> pending;size_t next_pending=0;uint64_t binding_ns=0;
 explicit CachedVM(int mode):evmc_vm{EVMC_ABI_VERSION},cache(icde_cache_create(mode==10||mode==11?10:0)){
  if(mode==6||mode==11||mode==12)region_cache=icde_cache_create(mode);
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
 ~CachedVM(){if(region_cache)icde_cache_destroy(region_cache);icde_cache_destroy(cache);}
 size_t window(){if(!region_cache)return 0;auto t=std::chrono::steady_clock::now();size_t n=0;
   while(next_pending<pending.size()&&std::chrono::steady_clock::now()-t<std::chrono::milliseconds(10)){
     require(icde_cache_upgrade(region_cache,pending[next_pending++])!=nullptr,"upgrade failed");++n;}
   binding_ns+=std::chrono::duration_cast<std::chrono::nanoseconds>(std::chrono::steady_clock::now()-t).count();return n;}

 void* prepare(const uint8_t* code,size_t n,bool stable){std::string key(reinterpret_cast<const char*>(code),n);auto i=content.find(key);void* p;
  if(i==content.end()){if(strict){std::cerr<<"timed analysis miss\n";std::exit(6);}++misses;p=icde_cache_prepare(cache,code,n);content.emplace(std::move(key),p);if(region_cache)pending.push_back(p);}else p=i->second;
  if(stable)pointers[code]={n,p};return p;
 }
};
static void seed(State& state,const View& view,CachedVM& vm){vm.pointers.clear();for(const auto& [a,x]:view.accounts){auto& acc=state.get(a);acc.access_status=x["warm"]?EVMC_ACCESS_WARM:EVMC_ACCESS_COLD;
 if(x["exists"].get<bool>()){auto code=state.get_code(a);if(!code.empty()){auto i=vm.content.find(std::string(reinterpret_cast<const char*>(code.data()),code.size()));if(i!=vm.content.end())vm.pointers[code.data()]={code.size(),i->second};}}
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
  for(auto k=x["storage"].begin();k!=x["storage"].end();++k)require(st.get_storage(a,word<bytes32>(k.key())).current==word<bytes32>(k.value()),prefix+"storage "+k.key()+" actual="+evmc::hex(st.get_storage(a,word<bytes32>(k.key())).current)+" expected="+k.value().get<std::string>());
  for(auto k=x["transient"].begin();k!=x["transient"].end();++k)require(acc.transient_storage[word<bytes32>(k.key())]==word<bytes32>(k.value()),prefix+"transient");
 }
 auto logs=host.take_logs();require(logs.size()==f["logs"].size(),"log count");for(size_t i=0;i<logs.size();i++){auto& want=f["logs"][i];require(logs[i].addr==word<address>(want["address"]),"log address");require(logs[i].data==decode(want["data"]),"log data");require(logs[i].topics.size()==want["topics"].size(),"log topics");for(size_t j=0;j<logs[i].topics.size();j++)require(logs[i].topics[j]==word<bytes32>(want["topics"][j]),"log topic");}
 // No unaccounted mutations may disappear merely because the export was partial.
 const auto diff=st.build_diff(EVMC_PRAGUE);
 for(const auto& a:diff.deleted_accounts){const auto key="0x"+evmc::hex(a);require(f["post"].contains(key),"unexpected deletion "+key);const auto& want=f["post"][key];require(!want["exists"].get<bool>()||want["destructed"].get<bool>()||(number(want["balance"])==0&&want["nonce"].get<uint64_t>()==0&&want["code"].get<std::string>().empty()),"unexpected nonempty deletion");}
 for(const auto& change:diff.modified_accounts){const auto key="0x"+evmc::hex(change.addr);require(f["post"].contains(key),"unexpected modified account "+key);for(const auto& [k,v]:change.modified_storage){const auto sk="0x"+evmc::hex(k);require(f["post"][key]["storage"].contains(sk),"unexpected modified storage "+sk);require(word<bytes32>(f["post"][key]["storage"][sk])==v,"diff storage");}}
}
// Fixtures are decoded before a block; validation and JSON encoding follow it.
// This avoids interleaving the artificial witness codec with every timed root.
struct StreamFrame {
 json f;View view;State state;BlockInfo block;Transaction tx;
 bytes input,code,expected;evmc_message message{};evmc::Result result;
 std::unique_ptr<evmone::state::Host> host;
 std::vector<std::pair<const uint8_t*,size_t>> code_views;
 uint64_t checkpoint;int64_t ns=0;
 StreamFrame(json data,CachedVM& cache,evmc::VM& vm):f(std::move(data)),view(f),state(view){
  seed(state,view,cache);
  const auto& b=f["block"];
  block.number=b["number"];block.timestamp=b["time"];block.gas_limit=b["gas_limit"];block.coinbase=word<address>(b["coinbase"]);block.prev_randao=word<bytes32>(b["randao"]);block.base_fee=static_cast<uint64_t>(number(b["base_fee"]));block.blob_base_fee=number(b["blob_base_fee"]);block.chain_id=1;
  tx.sender=word<address>(b["origin"]);tx.max_gas_price=number(b["gas_price"]);tx.max_priority_gas_price=tx.max_gas_price;
  for(const auto& x:b["blob_hashes"])tx.blob_hashes.push_back(word<bytes32>(x));
  input=decode(f["input"]);code=decode(f["root_code"]);expected=decode(f["output"]);
  message.sender=word<address>(f["from"]);message.recipient=word<address>(f["to"]);message.code_address=message.recipient;message.value=word<uint256be>(f["value"]);message.gas=f["gas"];message.input_data=input.data();message.input_size=input.size();message.flags=f["read_only"].get<bool>()?EVMC_STATIC:0;
  if(f["creation"].get<bool>()){message.kind=EVMC_CREATE;state.get(message.recipient).just_created=true;}
  for(const auto& [a,x]:view.accounts)if(x["exists"].get<bool>()){auto c=state.get_code(a);if(!c.empty())code_views.emplace_back(c.data(),c.size());}
  code_views.emplace_back(code.data(),code.size());
  host=std::make_unique<evmone::state::Host>(EVMC_PRAGUE,vm,state,block,view,tx);checkpoint=state.checkpoint();
 }
 void execute(CachedVM& cache,evmc::VM& vm){
  cache.pointers.clear();
  // Register only plans already observed, never analyze future code here.
  for(auto [ptr,n]:code_views){auto it=cache.content.find(std::string(reinterpret_cast<const char*>(ptr),n));if(it!=cache.content.end())cache.pointers[ptr]={n,it->second};}
  cache.root_ns=0;
  result=vm.execute(*host,EVMC_PRAGUE,message,code.data(),code.size());ns=cache.root_ns;
 }
 void verify(){
  const auto prefix=f["hash"].get<std::string>()+" ";
  const int status=result.status_code==EVMC_SUCCESS?0:result.status_code==EVMC_REVERT?2:1;
  require(status==f["status"].get<int>(),prefix+"status");require(result.gas_left==f["gas_left"].get<int64_t>(),prefix+"gas");require(result.gas_refund==f["refund"].get<int64_t>(),prefix+"refund");
  require(result.output_size==expected.size()&&(!result.output_size||!std::memcmp(result.output_data,expected.data(),expected.size())),prefix+"output");
  if(result.status_code!=EVMC_SUCCESS){state.rollback(checkpoint);host->take_logs().clear();}
  try{check_state(state,*host,f);}catch(const std::exception& e){throw std::runtime_error(prefix+e.what());}
 }
};
int main(int argc,char**argv){try{
 require(argc==2,"usage: native-stream-batch MODE < witnesses.jsonl");int mode=std::stoi(argv[1]);
 auto start=std::chrono::steady_clock::now();auto* cache=new CachedVM(mode);evmc::VM wrapped{cache};
 auto startup_ns=std::chrono::duration_cast<std::chrono::nanoseconds>(std::chrono::steady_clock::now()-start).count();
 std::unordered_map<std::string,std::string> codes;std::vector<json> rows;std::string line;
 uint64_t previous=0,transactions=0,frames=0;int64_t total=0;
 auto flush=[&](){
  if(rows.empty())return;
  cache->window();
  std::vector<std::unique_ptr<StreamFrame>> prepared;prepared.reserve(rows.size());
  for(const auto& f:rows)prepared.push_back(f["frame"].get<bool>()?std::make_unique<StreamFrame>(f,*cache,wrapped):nullptr);
  for(auto& f:prepared)if(f)f->execute(*cache,wrapped);
  for(size_t i=0;i<rows.size();++i){auto& f=prepared[i];if(f){f->verify();++frames;}const auto& r=rows[i];const int64_t ns=f?f->ns:0;total+=ns;++transactions;
   std::cout<<json{{"hash",r["hash"]},{"block",r["block_number"]},{"index",r["tx_index"]},{"to",r["outer_to"]},{"ns",ns},{"mode",mode},{"diagnostic",false},{"warmup",false}}.dump()<<'\n';
  }
  std::cout.flush();if(previous%20==0)std::cerr<<"transactions="<<transactions<<" block="<<previous<<" lat_s="<<double(total)/1e9<<std::endl;
  rows.clear();
 };
 while(std::getline(std::cin,line)){
  auto f=json::parse(line);auto height=f["block_number"].get<uint64_t>();
  require(!previous||height==previous||height==previous+1,"non-chronological witness");
  if(height!=previous){flush();previous=height;}
  if(f.contains("new_codes"))for(auto it=f["new_codes"].begin();it!=f["new_codes"].end();++it){require(!codes.contains(it.key()),"duplicate code definition");codes.emplace(it.key(),it.value().get<std::string>());}
  if(f["frame"].get<bool>()){
   if(f.contains("code_ref")){f["root_code"]=codes.at(f["code_ref"].get<std::string>());for(const char* p:{"accounts","post"})for(auto& x:f[p])x["code"]=codes.at(x["code_ref"].get<std::string>());}
   for(const char* p:{"accounts","post"}){json norm=json::object();for(auto it=f[p].begin();it!=f[p].end();++it)norm["0x"+evmc::hex(word<address>(it.key()))]=it.value();f[p]=norm;}
  }
  f.erase("new_codes");rows.push_back(std::move(f));
 }
 flush();
 std::cerr<<json{{"complete",true},{"transactions",transactions},{"frames",frames},{"lat_ns",total},{"binding_ns",cache->binding_ns},{"bindings",cache->next_pending},{"queued",cache->pending.size()},{"startup_ns",startup_ns},{"analyses",cache->misses},{"fixture_batching","one block; no JSON decoding or validation between timed roots"}}.dump()<<std::endl;
 return 0;
 }catch(const std::exception& e){std::cerr<<"native block stream failed: "<<e.what()<<std::endl;return 1;}}
