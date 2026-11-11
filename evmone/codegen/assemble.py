#!/usr/bin/env python3
"""Combine immutable previous analysis with newly proved stage3 programs."""
from pathlib import Path
import json,collections,argparse,copy
from codegen import Emitter
from branch_codegen import emit_shared
p=argparse.ArgumentParser();p.add_argument('base');p.add_argument('effects');p.add_argument('output');a=p.parse_args()
r=json.loads(Path(a.base).read_text());by={x['code']:x for x in r['codes']}
for row in json.loads(Path(a.effects).read_text())['codes']:
 if row['code'] not in by:r['codes'].append(row);continue
 by[row['code']]['regions'].extend(row['regions'] or [])
merged=0
for row in r['codes']:
 new=[]
 for item in row['regions']:
  x=item.get('native',{})
  if item.get('stage')!=2 or len(x.get('paths',[]))<2:continue
  if emit_shared(Emitter(),'probe',x) is None:continue
  item=copy.deepcopy(item);item['stage']=3;item['native']['shared_branches']=True;new.append(item);merged+=1
 row['regions'].extend(new);row['regions'].sort(key=lambda x:(x['entry'],-x.get('stage',0),x.get('native',{}).get('family')!='effects'))
r['counts']=dict(collections.Counter(str(x['kind']) for row in r['codes'] for x in row['regions']))
Path(a.output).write_text(json.dumps(r,separators=(',',':'))+'\n');print('shared branch regions',merged)
