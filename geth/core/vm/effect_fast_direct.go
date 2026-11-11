// Code generated from effect summaries. DO NOT EDIT.
package vm

func chainFastShape0(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	v[39].And(&v[26], &v[38])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[40] = p.nodes[40].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[41] = p.nodes[41].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[44] = p.nodes[44].c
	v[45] = p.nodes[45].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[47] = p.nodes[47].c
	v[48] = p.nodes[48].c
	v[49].And(&v[27], &v[38])
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[52] = p.nodes[52].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["02541769305f6e00521686802a09fcd97f4d32c00f8b9edc1da798fada867639"] = chainFastShape0
}

func chainFastShape1(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = chainPure(OpCode(21), v[30], v[33])
	v[35] = p.nodes[35].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[41] = chainPure(OpCode(21), v[29], v[33])
	v[42] = p.nodes[42].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	v[45] = p.nodes[45].c
	v[46].And(&v[26], &v[45])
	v[47] = p.nodes[47].c
	v[48] = p.nodes[48].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[50] = p.nodes[50].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[51] = p.nodes[51].c
	v[52].Add(&v[49], &v[51])
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[54].Sub(&v[52], &v[53])
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[55] = p.nodes[55].c
	v[56] = p.nodes[56].c
	v[57].And(&v[25], &v[45])
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[60] = p.nodes[60].c
	v[61] = p.nodes[61].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[63] = p.nodes[63].c
	v[64] = p.nodes[64].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[67] = p.nodes[67].c
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["02d2bfc93fee4c69840cea97dc266add002d463a80fe0263ad0f0f2cb6655c93"] = chainFastShape1
}

func chainFastShape2(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37].And(&v[28], &v[35])
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[40] = p.nodes[40].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[43] = p.nodes[43].c
	v[44].And(&v[42], &v[43])
	v[45] = chainPure(OpCode(17), v[30], v[44])
	v[46] = chainPure(OpCode(21), v[45], v[36])
	v[47] = p.nodes[47].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[48] = p.nodes[48].c
	v[49] = p.nodes[49].c
	v[50] = p.nodes[50].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[52].And(&v[43], &v[51])
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[54] = p.nodes[54].c
	v[55] = chainPure(OpCode(17), v[53], v[43])
	v[56] = chainPure(OpCode(21), v[55], v[36])
	v[57] = p.nodes[57].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[58] = p.nodes[58].c
	v[59] = p.nodes[59].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[61] = p.nodes[61].c
	v[62] = chainPure(OpCode(28), v[61], v[60])
	v[63] = p.nodes[63].c
	v[64] = chainPure(OpCode(20), v[62], v[63])
	v[65] = chainPure(OpCode(21), v[64], v[36])
	v[66] = p.nodes[66].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[67] = p.nodes[67].c
	v[68] = p.nodes[68].c
	v[69].And(&v[29], &v[35])
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[72].And(&v[43], &v[71])
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	v[74] = chainPure(OpCode(17), v[73], v[43])
	v[75] = chainPure(OpCode(21), v[74], v[36])
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	v[77] = chainPure(OpCode(28), v[61], v[76])
	v[78] = chainPure(OpCode(20), v[63], v[77])
	v[79] = chainPure(OpCode(21), v[78], v[36])
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	v[80] = p.nodes[80].c
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	v[82].Add(&v[39], &v[81])
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	v[84].Sub(&v[82], &v[83])
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["03334448e3f2ba5d6bc707e7e5abf56dd2af7c8ae84eb868ae22e20548446ce3"] = chainFastShape2
}

func chainFastShape3(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	v[39].And(&v[34], &v[38])
	v[40].And(&v[38], &v[39])
	v[41].And(&v[27], &v[38])
	v[42] = chainPure(OpCode(20), v[40], v[41])
	v[43] = chainPure(OpCode(21), v[42], v[33])
	v[44] = chainPure(OpCode(21), v[43], v[33])
	v[45] = p.nodes[45].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[46] = p.nodes[46].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[47] = p.nodes[47].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[48] = p.nodes[48].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[49] = p.nodes[49].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[50] = p.nodes[50].c
	v[51] = p.nodes[51].c
	v[52] = p.nodes[52].c
	v[53] = p.nodes[53].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[54] = chainPure(OpCode(21), v[31], v[33])
	v[55] = p.nodes[55].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[56].And(&v[28], &v[38])
	v[57] = p.nodes[57].c
	v[58] = p.nodes[58].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[60] = p.nodes[60].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[61] = p.nodes[61].c
	v[62].Add(&v[59], &v[61])
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[64].Sub(&v[62], &v[63])
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[65] = p.nodes[65].c
	v[66] = p.nodes[66].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[69] = p.nodes[69].c
	v[70] = p.nodes[70].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[72] = p.nodes[72].c
	v[73] = p.nodes[73].c
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[76] = p.nodes[76].c
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["0334ef1980ce87741fc3394410e686e421ed740a5050a3138159668a5a4616b3"] = chainFastShape3
}

func chainFastShape4(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	v[41] = p.nodes[41].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[46] = p.nodes[46].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[47] = p.nodes[47].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[50] = p.nodes[50].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[54] = p.nodes[54].c
	v[55].Sub(&v[52], &v[53])
	v[56].Add(&v[46], &v[55])
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[57] = p.nodes[57].c
	v[58] = chainPure(OpCode(17), v[24], v[37])
	v[59] = p.nodes[59].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[60] = p.nodes[60].c
	v[61].And(&v[23], &v[36])
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	v[69].Sub(&v[67], &v[68])
	v[70].Add(&v[46], &v[69])
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	v[71] = p.nodes[71].c
	v[72] = p.nodes[72].c
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	v[74] = p.nodes[74].c
	v[75] = p.nodes[75].c
	v[76] = p.nodes[76].c
	v[77] = chainPure(OpCode(17), v[73], v[76])
	v[78] = chainPure(OpCode(21), v[77], v[37])
	v[79] = p.nodes[79].c
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	v[80] = p.nodes[80].c
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	v[82] = p.nodes[82].c
	v[83].And(&v[81], &v[82])
	v[84].And(&v[73], &v[76])
	v[85].Or(&v[83], &v[84])
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["03f08dfaf309f2647eaa08f35e509d38d82bb8770d900c18745c189604f065f5"] = chainFastShape4
}

func chainFastShape5(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = chainPure(OpCode(21), v[30], v[33])
	v[35] = p.nodes[35].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[36] = p.nodes[36].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[37] = p.nodes[37].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[38] = chainPure(OpCode(17), v[28], v[33])
	v[39] = chainPure(OpCode(21), v[38], v[33])
	v[40] = p.nodes[40].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[41] = p.nodes[41].c
	v[42] = p.nodes[42].c
	v[43] = p.nodes[43].c
	v[44].And(&v[25], &v[43])
	v[45].And(&v[43], &v[44])
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[46] = p.nodes[46].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[47] = p.nodes[47].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[50] = p.nodes[50].c
	v[51] = p.nodes[51].c
	v[52] = p.nodes[52].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[54].Add(&v[47], &v[53])
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[55] = p.nodes[55].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[56].Add(&v[46], &v[53])
	v[57] = p.nodes[57].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[58] = p.nodes[58].c
	v[59] = chainPure(OpCode(17), v[27], v[49])
	v[60] = chainPure(OpCode(21), v[59], v[33])
	v[61] = p.nodes[61].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[62] = p.nodes[62].c
	v[63] = p.nodes[63].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[65] = p.nodes[65].c
	v[66] = p.nodes[66].c
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[68].Add(&v[47], &v[67])
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[69].Add(&v[46], &v[67])
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[70] = chainPure(OpCode(17), v[28], v[27])
	v[71] = chainPure(OpCode(21), v[70], v[33])
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[73].And(&v[26], &v[43])
	v[74].And(&v[43], &v[73])
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	v[77] = p.nodes[77].c
	v[78] = p.nodes[78].c
	v[79] = p.nodes[79].c
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	v[81] = chainPure(OpCode(16), v[80], v[76])
	v[82] = chainPure(OpCode(21), v[81], v[33])
	v[83] = p.nodes[83].c
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	v[84] = p.nodes[84].c
	v[85] = p.nodes[85].c
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	v[87].Add(&v[47], &v[86])
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	v[88].Add(&v[46], &v[86])
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	v[91] = p.nodes[91].c
	v[92] = p.nodes[92].c
	v[93].Add(&v[46], &v[90])
	v[94] = p.nodes[94].c
	v[95].Add(&v[33], &v[90])
	v[96] = p.nodes[96].c
	v[97] = p.nodes[97].c
	v[98] = p.nodes[98].c
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	v[100].Sub(&v[93], &v[99])
	code, err = p.actionFast(44, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 44 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["04dcb53e6fd44053cb8dfd57aa8350f0148bd54f521d9c4642b96834d1defae1"] = chainFastShape5
}

func chainFastShape6(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = chainPure(OpCode(17), v[29], v[33])
	v[35] = chainPure(OpCode(21), v[34], v[33])
	v[36] = p.nodes[36].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	v[40].And(&v[26], &v[39])
	v[41].And(&v[39], &v[40])
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[42] = p.nodes[42].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[43] = p.nodes[43].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[46] = p.nodes[46].c
	v[47] = p.nodes[47].c
	v[48] = p.nodes[48].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[50].Add(&v[43], &v[49])
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[51] = p.nodes[51].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[52].Add(&v[42], &v[49])
	v[53] = p.nodes[53].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[54] = p.nodes[54].c
	v[55] = chainPure(OpCode(17), v[28], v[45])
	v[56] = chainPure(OpCode(21), v[55], v[33])
	v[57] = p.nodes[57].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[58] = p.nodes[58].c
	v[59] = p.nodes[59].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[61] = p.nodes[61].c
	v[62] = p.nodes[62].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[64].Add(&v[43], &v[63])
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[65].Add(&v[42], &v[63])
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[66] = chainPure(OpCode(17), v[29], v[28])
	v[67] = chainPure(OpCode(21), v[66], v[33])
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[69].And(&v[27], &v[39])
	v[70].And(&v[39], &v[69])
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	v[73] = p.nodes[73].c
	v[74] = p.nodes[74].c
	v[75] = p.nodes[75].c
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[77] = chainPure(OpCode(16), v[76], v[72])
	v[78] = chainPure(OpCode(21), v[77], v[33])
	v[79] = p.nodes[79].c
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	v[80] = p.nodes[80].c
	v[81] = p.nodes[81].c
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	v[83].Add(&v[43], &v[82])
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	v[84].Add(&v[42], &v[82])
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	v[87] = p.nodes[87].c
	v[88] = p.nodes[88].c
	v[89].Add(&v[42], &v[86])
	v[90] = p.nodes[90].c
	v[91].Add(&v[33], &v[86])
	v[92] = p.nodes[92].c
	v[93] = p.nodes[93].c
	v[94] = p.nodes[94].c
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	v[96].Sub(&v[89], &v[95])
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["06e01740366ef9909ab71a1118f789984a46e7e8694561f60c428eebf982e448"] = chainFastShape6
}

func chainFastShape7(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35].SetBytes(effectiveCaller(s))
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38].And(&v[35], &v[37])
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[41] = p.nodes[41].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[44] = p.nodes[44].c
	v[45].And(&v[43], &v[44])
	v[46] = chainPure(OpCode(17), v[29], v[33])
	v[47] = p.nodes[47].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[48] = chainPure(OpCode(16), v[45], v[29])
	v[49] = chainPure(OpCode(21), v[48], v[33])
	v[50] = p.nodes[50].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[51] = p.nodes[51].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[53] = p.nodes[53].c
	v[54] = p.nodes[54].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[56] = p.nodes[56].c
	v[57] = p.nodes[57].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[59] = p.nodes[59].c
	v[60] = chainPure(OpCode(17), v[58], v[44])
	v[61] = chainPure(OpCode(21), v[60], v[33])
	v[62] = p.nodes[62].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[63] = p.nodes[63].c
	v[64] = p.nodes[64].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[66] = p.nodes[66].c
	v[67] = chainPure(OpCode(28), v[66], v[65])
	v[68] = p.nodes[68].c
	v[69] = chainPure(OpCode(20), v[67], v[68])
	v[70] = chainPure(OpCode(21), v[69], v[33])
	v[71] = p.nodes[71].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[74] = p.nodes[74].c
	v[75].Sub(&v[72], &v[73])
	v[76].Add(&v[40], &v[75])
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	v[79] = p.nodes[79].c
	v[80].Sub(&v[77], &v[78])
	v[81].Add(&v[40], &v[80])
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["0999524f944b05b8df58f2fe488f833c56470711a50a9b4d662a65353b84e230"] = chainFastShape7
}

func chainFastShape8(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	v[41] = chainPure(OpCode(20), v[27], v[40])
	v[42] = chainPure(OpCode(21), v[41], v[40])
	v[43] = p.nodes[43].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[44].Mul(&v[27], &v[37])
	v[45] = chainPure(OpCode(21), v[27], v[40])
	v[46] = chainPure(OpCode(21), v[45], v[40])
	v[47] = p.nodes[47].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[48].Div(&v[44], &v[27])
	v[49] = chainPure(OpCode(20), v[37], v[48])
	v[50] = chainPure(OpCode(21), v[49], v[40])
	v[51] = chainPure(OpCode(21), v[50], v[40])
	v[52] = p.nodes[52].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[53] = p.nodes[53].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[55] = p.nodes[55].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[57] = chainPure(OpCode(17), v[54], v[56])
	v[58] = chainPure(OpCode(21), v[57], v[40])
	v[59] = p.nodes[59].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[60] = p.nodes[60].c
	v[61] = p.nodes[61].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[63] = p.nodes[63].c
	v[64] = p.nodes[64].c
	v[65].SetBytes(effectiveCaller(s))
	v[66] = p.nodes[66].c
	v[67].And(&v[65], &v[66])
	v[68].And(&v[66], &v[67])
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[69] = p.nodes[69].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[70] = p.nodes[70].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[74] = p.nodes[74].c
	v[75].And(&v[26], &v[66])
	v[76].And(&v[66], &v[75])
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[79] = p.nodes[79].c
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[81] = chainPure(OpCode(17), v[54], v[40])
	v[82] = chainPure(OpCode(21), v[81], v[40])
	v[83] = p.nodes[83].c
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	v[84] = p.nodes[84].c
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	v[86].Add(&v[69], &v[85])
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	v[88].Sub(&v[86], &v[87])
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["0ae58ef581548615176694418b02aba5efdc150e5ea7fe645c68a33bd04f9237"] = chainFastShape8
}

func chainFastShape9(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = chainPure(OpCode(21), v[30], v[33])
	v[35] = p.nodes[35].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[36] = p.nodes[36].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[37] = chainPure(OpCode(17), v[28], v[33])
	v[38] = chainPure(OpCode(21), v[37], v[33])
	v[39] = p.nodes[39].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[40] = p.nodes[40].c
	v[41] = p.nodes[41].c
	v[42] = p.nodes[42].c
	v[43].And(&v[25], &v[42])
	v[44].And(&v[42], &v[43])
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[45] = p.nodes[45].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[46] = p.nodes[46].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[49] = p.nodes[49].c
	v[50] = p.nodes[50].c
	v[51] = p.nodes[51].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[53].Add(&v[46], &v[52])
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[54] = p.nodes[54].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[55].Add(&v[45], &v[52])
	v[56] = p.nodes[56].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[57] = p.nodes[57].c
	v[58] = chainPure(OpCode(17), v[27], v[48])
	v[59] = chainPure(OpCode(21), v[58], v[33])
	v[60] = p.nodes[60].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[61] = p.nodes[61].c
	v[62] = p.nodes[62].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[64] = p.nodes[64].c
	v[65] = p.nodes[65].c
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[67].Add(&v[46], &v[66])
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[68].Add(&v[45], &v[66])
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[69] = chainPure(OpCode(17), v[28], v[27])
	v[70] = chainPure(OpCode(21), v[69], v[33])
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[72].And(&v[26], &v[42])
	v[73].And(&v[42], &v[72])
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	v[76] = p.nodes[76].c
	v[77] = p.nodes[77].c
	v[78] = p.nodes[78].c
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	v[80] = chainPure(OpCode(16), v[79], v[75])
	v[81] = chainPure(OpCode(21), v[80], v[33])
	v[82] = p.nodes[82].c
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	v[83] = p.nodes[83].c
	v[84] = p.nodes[84].c
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	v[86].Add(&v[46], &v[85])
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	v[87].Add(&v[45], &v[85])
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	v[90] = p.nodes[90].c
	v[91] = p.nodes[91].c
	v[92].Add(&v[45], &v[89])
	v[93] = p.nodes[93].c
	v[94].Add(&v[33], &v[89])
	v[95] = p.nodes[95].c
	v[96] = p.nodes[96].c
	v[97] = p.nodes[97].c
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	v[99].Sub(&v[92], &v[98])
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["0b596564a09d23269b2269f881029bf6f558669ce0a1d893dc09b552aec55df6"] = chainFastShape9
}

func chainFastShape10(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38].And(&v[28], &v[37])
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[41] = p.nodes[41].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[44] = p.nodes[44].c
	v[45] = p.nodes[45].c
	v[46] = p.nodes[46].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[48] = p.nodes[48].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[52] = p.nodes[52].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[53] = p.nodes[53].c
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[57] = p.nodes[57].c
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	v[59] = p.nodes[59].c
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	v[62] = p.nodes[62].c
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	v[63] = p.nodes[63].c
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	v[66] = p.nodes[66].c
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	code, err = p.actionFast(44, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 44 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["0cd55585a3778bcb1e153d01a11690cce4a5c6de4c9da1e8dbb0b6c72bd2cb00"] = chainFastShape10
}

func chainFastShape11(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33].SetBytes(effectiveCaller(s))
	v[34] = p.nodes[34].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[37] = p.nodes[37].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[40] = p.nodes[40].c
	v[41].And(&v[39], &v[40])
	v[42] = chainPure(OpCode(21), v[41], v[34])
	v[43] = p.nodes[43].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[44] = p.nodes[44].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[47] = p.nodes[47].c
	v[48] = p.nodes[48].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[50] = p.nodes[50].c
	v[51] = p.nodes[51].c
	v[52] = p.nodes[52].c
	v[53].And(&v[28], &v[52])
	v[54] = p.nodes[54].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[55] = p.nodes[55].c
	v[56] = p.nodes[56].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[58] = p.nodes[58].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[60] = p.nodes[60].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[63] = p.nodes[63].c
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	v[66] = p.nodes[66].c
	v[67] = p.nodes[67].c
	v[68] = p.nodes[68].c
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	v[69].Add(&v[36], &v[65])
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	v[71].Sub(&v[69], &v[70])
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["0ce87deca04a04fb962fa9adbdf81bc82787b2c44d49aa4125468c3b34bc13b6"] = chainFastShape11
}

func chainFastShape12(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36].And(&v[31], &v[34])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[39] = p.nodes[39].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[42] = p.nodes[42].c
	v[43] = chainPure(OpCode(28), v[42], v[41])
	v[44] = p.nodes[44].c
	v[45] = chainPure(OpCode(20), v[43], v[44])
	v[46] = chainPure(OpCode(21), v[45], v[35])
	v[47] = p.nodes[47].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[48].And(&v[25], &v[34])
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[49] = p.nodes[49].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[51].SetBytes(effectiveCaller(s))
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[54] = chainPure(OpCode(17), v[27], v[53])
	v[55] = chainPure(OpCode(21), v[54], v[35])
	v[56] = p.nodes[56].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[57] = p.nodes[57].c
	v[58] = p.nodes[58].c
	v[59] = p.nodes[59].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[60].And(&v[26], &v[34])
	v[61] = p.nodes[61].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[62] = p.nodes[62].c
	v[63] = p.nodes[63].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[66] = p.nodes[66].c
	v[67].And(&v[65], &v[66])
	v[68] = chainPure(OpCode(17), v[27], v[67])
	v[69] = chainPure(OpCode(21), v[68], v[35])
	v[70] = p.nodes[70].c
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[71] = p.nodes[71].c
	v[72] = p.nodes[72].c
	v[73] = p.nodes[73].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[75].And(&v[66], &v[74])
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[77] = p.nodes[77].c
	v[78] = chainPure(OpCode(17), v[76], v[66])
	v[79] = chainPure(OpCode(21), v[78], v[35])
	v[80] = p.nodes[80].c
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	v[81] = p.nodes[81].c
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	v[83] = chainPure(OpCode(28), v[42], v[82])
	v[84] = chainPure(OpCode(20), v[44], v[83])
	v[85] = chainPure(OpCode(21), v[84], v[35])
	v[86] = p.nodes[86].c
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	v[87] = p.nodes[87].c
	v[88] = p.nodes[88].c
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	v[91].And(&v[66], &v[90])
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	v[93] = chainPure(OpCode(17), v[92], v[66])
	v[94] = chainPure(OpCode(21), v[93], v[35])
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	code, err = p.actionFast(44, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 44 != 0, err
	}
	v[96] = chainPure(OpCode(28), v[42], v[95])
	v[97] = chainPure(OpCode(20), v[44], v[96])
	v[98] = chainPure(OpCode(21), v[97], v[35])
	code, err = p.actionFast(45, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 45 != 0, err
	}
	code, err = p.actionFast(46, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 46 != 0, err
	}
	code, err = p.actionFast(47, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 47 != 0, err
	}
	code, err = p.actionFast(48, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 48 != 0, err
	}
	code, err = p.actionFast(49, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 49 != 0, err
	}
	v[99] = p.nodes[99].c
	code, err = p.actionFast(50, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 50 != 0, err
	}
	code, err = p.actionFast(51, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 51 != 0, err
	}
	v[101].Add(&v[38], &v[100])
	code, err = p.actionFast(52, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 52 != 0, err
	}
	v[103].Sub(&v[101], &v[102])
	code, err = p.actionFast(53, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 53 != 0, err
	}
	code, err = p.actionFast(54, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 54 != 0, err
	}
	code, err = p.actionFast(55, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 55 != 0, err
	}
	code, err = p.actionFast(56, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 56 != 0, err
	}
	code, err = p.actionFast(57, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 57 != 0, err
	}
	code, err = p.actionFast(58, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 58 != 0, err
	}
	code, err = p.actionFast(59, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 59 != 0, err
	}
	code, err = p.actionFast(60, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 60 != 0, err
	}
	v[105] = p.nodes[105].c
	v[106] = p.nodes[106].c
	code, err = p.actionFast(61, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 61 != 0, err
	}
	code, err = p.actionFast(62, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 62 != 0, err
	}
	code, err = p.actionFast(63, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 63 != 0, err
	}
	code, err = p.actionFast(64, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 64 != 0, err
	}
	code, err = p.actionFast(65, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 65 != 0, err
	}
	code, err = p.actionFast(66, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 66 != 0, err
	}
	code, err = p.actionFast(67, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 67 != 0, err
	}
	code, err = p.actionFast(68, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 68 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["0d1b410ddde539d19b744ea4ed266404183578c2bb498161f0e88deceb83503a"] = chainFastShape12
}

func chainFastShape13(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = chainPure(OpCode(17), v[27], v[33])
	v[35] = p.nodes[35].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[36].SetBytes(effectiveCaller(s))
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[39] = p.nodes[39].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[42] = chainPure(OpCode(17), v[27], v[41])
	v[43] = chainPure(OpCode(21), v[42], v[33])
	v[44] = p.nodes[44].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[45] = p.nodes[45].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[47] = p.nodes[47].c
	v[48] = p.nodes[48].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[50] = p.nodes[50].c
	v[51] = p.nodes[51].c
	v[52] = p.nodes[52].c
	v[53] = p.nodes[53].c
	v[54] = p.nodes[54].c
	v[55].And(&v[26], &v[54])
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[56] = p.nodes[56].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[59] = p.nodes[59].c
	v[60].And(&v[58], &v[59])
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[62] = p.nodes[62].c
	v[63] = chainPure(OpCode(17), v[61], v[59])
	v[64] = chainPure(OpCode(21), v[63], v[33])
	v[65] = p.nodes[65].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[66] = p.nodes[66].c
	v[67] = p.nodes[67].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[69] = p.nodes[69].c
	v[70] = chainPure(OpCode(28), v[69], v[68])
	v[71] = p.nodes[71].c
	v[72] = chainPure(OpCode(20), v[70], v[71])
	v[73] = chainPure(OpCode(21), v[72], v[33])
	v[74] = p.nodes[74].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[75] = p.nodes[75].c
	v[76] = p.nodes[76].c
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	v[80] = p.nodes[80].c
	v[81].Sub(&v[78], &v[79])
	v[82].Add(&v[38], &v[81])
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	v[85] = p.nodes[85].c
	v[86].Sub(&v[83], &v[84])
	v[87].Add(&v[38], &v[86])
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["0f51e6642967fff3c3ec290e9edde74fd7fbed40d2a7d9b1010b9c2a5504fc4f"] = chainFastShape13
}

func chainFastShape14(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33].SetBytes(effectiveCaller(s))
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38].And(&v[33], &v[37])
	v[39].And(&v[37], &v[38])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[40] = p.nodes[40].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[41] = p.nodes[41].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[44] = chainPure(OpCode(16), v[43], v[30])
	v[45] = chainPure(OpCode(21), v[44], v[35])
	v[46] = chainPure(OpCode(21), v[45], v[35])
	v[47] = chainPure(OpCode(21), v[46], v[35])
	v[48] = p.nodes[48].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[49] = chainPure(OpCode(20), v[38], v[38])
	v[50] = chainPure(OpCode(21), v[49], v[35])
	v[51] = chainPure(OpCode(21), v[50], v[35])
	v[52] = p.nodes[52].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[53] = p.nodes[53].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[54] = p.nodes[54].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[56] = p.nodes[56].c
	v[57] = p.nodes[57].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[59] = p.nodes[59].c
	v[60].And(&v[29], &v[37])
	v[61].And(&v[37], &v[60])
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[64] = p.nodes[64].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[66] = p.nodes[66].c
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[68].Add(&v[40], &v[67])
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	v[70].Sub(&v[68], &v[69])
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["11e9be8c8e31f3296f2a2c6b7c7068dbb49cc1a503028a2a74e884ac10d78d63"] = chainFastShape14
}

func chainFastShape15(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = chainPure(OpCode(21), v[30], v[33])
	v[35] = p.nodes[35].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	v[40].And(&v[26], &v[39])
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[41] = p.nodes[41].c
	v[42] = p.nodes[42].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[43] = p.nodes[43].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[46] = p.nodes[46].c
	v[47] = p.nodes[47].c
	v[48].Div(&v[45], &v[47])
	v[49] = p.nodes[49].c
	v[50].And(&v[48], &v[49])
	v[51] = chainPure(OpCode(21), v[50], v[33])
	v[52] = p.nodes[52].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[53] = p.nodes[53].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[54] = p.nodes[54].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[56] = p.nodes[56].c
	v[57] = p.nodes[57].c
	v[58] = p.nodes[58].c
	v[59].Div(&v[55], &v[58])
	v[60] = p.nodes[60].c
	v[61].And(&v[59], &v[60])
	v[62] = p.nodes[62].c
	v[63] = p.nodes[63].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[65] = p.nodes[65].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[68] = p.nodes[68].c
	v[69] = p.nodes[69].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[71] = chainPure(OpCode(17), v[70], v[64])
	v[72] = chainPure(OpCode(21), v[71], v[33])
	v[73] = p.nodes[73].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[74] = p.nodes[74].c
	v[75] = p.nodes[75].c
	v[76] = p.nodes[76].c
	v[77] = p.nodes[77].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[78] = chainPure(OpCode(21), v[29], v[33])
	v[79] = p.nodes[79].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[80] = p.nodes[80].c
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[82] = p.nodes[82].c
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[83].Add(&v[42], &v[81])
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[85].Sub(&v[83], &v[84])
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[86] = p.nodes[86].c
	v[87] = p.nodes[87].c
	v[88].And(&v[25], &v[39])
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	v[91] = p.nodes[91].c
	v[92] = p.nodes[92].c
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	v[94] = p.nodes[94].c
	v[95] = p.nodes[95].c
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["138f748911e180c853e748bb55c44f9c3ce8171bf9bd74a3ed0528efa316d236"] = chainFastShape15
}

func chainFastShape16(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = chainPure(OpCode(21), v[30], v[33])
	v[35] = p.nodes[35].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[36] = chainPure(OpCode(17), v[28], v[33])
	v[37] = chainPure(OpCode(21), v[36], v[33])
	v[38] = p.nodes[38].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	v[41] = p.nodes[41].c
	v[42].And(&v[25], &v[41])
	v[43].And(&v[41], &v[42])
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[44] = p.nodes[44].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[45] = p.nodes[45].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[48] = p.nodes[48].c
	v[49] = p.nodes[49].c
	v[50] = p.nodes[50].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[52].Add(&v[45], &v[51])
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[53] = p.nodes[53].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[54].Add(&v[44], &v[51])
	v[55] = p.nodes[55].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[56] = p.nodes[56].c
	v[57] = chainPure(OpCode(17), v[27], v[47])
	v[58] = chainPure(OpCode(21), v[57], v[33])
	v[59] = p.nodes[59].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[60] = p.nodes[60].c
	v[61] = p.nodes[61].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[63] = p.nodes[63].c
	v[64] = p.nodes[64].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[66].Add(&v[45], &v[65])
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[67].Add(&v[44], &v[65])
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[68] = chainPure(OpCode(17), v[28], v[27])
	v[69] = chainPure(OpCode(21), v[68], v[33])
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[71].And(&v[26], &v[41])
	v[72].And(&v[41], &v[71])
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[75] = p.nodes[75].c
	v[76] = p.nodes[76].c
	v[77] = p.nodes[77].c
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	v[79] = chainPure(OpCode(16), v[78], v[74])
	v[80] = chainPure(OpCode(21), v[79], v[33])
	v[81] = p.nodes[81].c
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	v[82] = p.nodes[82].c
	v[83] = p.nodes[83].c
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	v[85].Add(&v[45], &v[84])
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	v[86].Add(&v[44], &v[84])
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	v[89] = p.nodes[89].c
	v[90] = p.nodes[90].c
	v[91].Add(&v[44], &v[88])
	v[92] = p.nodes[92].c
	v[93].Add(&v[33], &v[88])
	v[94] = p.nodes[94].c
	v[95] = p.nodes[95].c
	v[96] = p.nodes[96].c
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	v[98].Sub(&v[91], &v[97])
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["13e3be939c25fd203f2ee5ca79c61f32f7acd45d64c3e9700e6b46ae18167c56"] = chainFastShape16
}

func chainFastShape17(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34].SetBytes(effectiveCaller(s))
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38].And(&v[34], &v[36])
	v[39] = p.nodes[39].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[40].And(&v[26], &v[36])
	v[41] = p.nodes[41].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[42] = p.nodes[42].c
	v[43] = p.nodes[43].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[44] = p.nodes[44].c
	v[45] = p.nodes[45].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[46] = p.nodes[46].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[49] = p.nodes[49].c
	v[50].And(&v[48], &v[49])
	v[51] = chainPure(OpCode(17), v[27], v[50])
	v[52] = chainPure(OpCode(21), v[51], v[37])
	v[53] = p.nodes[53].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[54] = p.nodes[54].c
	v[55] = p.nodes[55].c
	v[56] = p.nodes[56].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[58].And(&v[49], &v[57])
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[60] = p.nodes[60].c
	v[61] = chainPure(OpCode(17), v[59], v[49])
	v[62] = chainPure(OpCode(21), v[61], v[37])
	v[63] = p.nodes[63].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[64] = p.nodes[64].c
	v[65] = p.nodes[65].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[67] = p.nodes[67].c
	v[68] = chainPure(OpCode(28), v[67], v[66])
	v[69] = p.nodes[69].c
	v[70] = chainPure(OpCode(20), v[68], v[69])
	v[71] = chainPure(OpCode(21), v[70], v[37])
	v[72] = p.nodes[72].c
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[73] = p.nodes[73].c
	v[74] = p.nodes[74].c
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[77].And(&v[49], &v[76])
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	v[79] = chainPure(OpCode(17), v[78], v[49])
	v[80] = chainPure(OpCode(21), v[79], v[37])
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	v[82] = chainPure(OpCode(28), v[67], v[81])
	v[83] = chainPure(OpCode(20), v[69], v[82])
	v[84] = chainPure(OpCode(21), v[83], v[37])
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	v[85] = p.nodes[85].c
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	v[87].Add(&v[45], &v[86])
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	v[89].Sub(&v[87], &v[88])
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["142adcbc88b712c69a25caa5f7c32bde09b627e61478e74d7cb6c26f23f28060"] = chainFastShape17
}

func chainFastShape18(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = chainPure(OpCode(21), v[30], v[33])
	v[35] = p.nodes[35].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[36] = p.nodes[36].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	v[41] = p.nodes[41].c
	v[42].Div(&v[37], &v[41])
	v[43] = p.nodes[43].c
	v[44].And(&v[42], &v[43])
	v[45] = p.nodes[45].c
	v[46] = p.nodes[46].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[48] = p.nodes[48].c
	v[49] = p.nodes[49].c
	v[50] = p.nodes[50].c
	v[51] = p.nodes[51].c
	v[52].And(&v[26], &v[51])
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[53] = p.nodes[53].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[54] = p.nodes[54].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[57] = p.nodes[57].c
	v[58] = p.nodes[58].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[60] = chainPure(OpCode(17), v[59], v[47])
	v[61] = chainPure(OpCode(21), v[60], v[33])
	v[62] = p.nodes[62].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[63] = p.nodes[63].c
	v[64] = p.nodes[64].c
	v[65] = p.nodes[65].c
	v[66] = p.nodes[66].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[67] = chainPure(OpCode(21), v[29], v[33])
	v[68] = p.nodes[68].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[69] = p.nodes[69].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[71] = p.nodes[71].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[72].Add(&v[53], &v[70])
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[74].Sub(&v[72], &v[73])
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[75] = p.nodes[75].c
	v[76] = p.nodes[76].c
	v[77].And(&v[25], &v[51])
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[80] = p.nodes[80].c
	v[81] = p.nodes[81].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[83] = p.nodes[83].c
	v[84] = p.nodes[84].c
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["14868eab1e935ae55db69a24be7d386b92bc329517b0f701cb22530d5a52069c"] = chainFastShape18
}

func chainFastShape19(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	v[41].And(&v[25], &v[40])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[42] = p.nodes[42].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[43] = p.nodes[43].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[46] = p.nodes[46].c
	v[47] = p.nodes[47].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[49] = p.nodes[49].c
	v[50] = p.nodes[50].c
	v[51].And(&v[26], &v[40])
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[54] = p.nodes[54].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["17a34aa87efb9150c83485089b225e829b09773488209bc31f6394e258e54bac"] = chainFastShape19
}

func chainFastShape20(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36].And(&v[28], &v[35])
	v[37].And(&v[35], &v[36])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[38] = p.nodes[38].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[39] = p.nodes[39].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[42] = p.nodes[42].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[44] = p.nodes[44].c
	v[45].And(&v[29], &v[35])
	v[46].And(&v[35], &v[45])
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[49] = p.nodes[49].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[51] = p.nodes[51].c
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[53].Add(&v[38], &v[52])
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[55].Sub(&v[53], &v[54])
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["1a6c368cb517713aeee4825ebc4824114e92cbe59c5a5ec07940c5047bdf2269"] = chainFastShape20
}

func chainFastShape21(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[40] = p.nodes[40].c
	v[41] = p.nodes[41].c
	v[42] = p.nodes[42].c
	v[43] = p.nodes[43].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[44] = p.nodes[44].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[45] = p.nodes[45].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[48] = p.nodes[48].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[52] = p.nodes[52].c
	v[53].Sub(&v[50], &v[51])
	v[54].Add(&v[44], &v[53])
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[55] = p.nodes[55].c
	v[56] = chainPure(OpCode(17), v[23], v[33])
	v[57] = p.nodes[57].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[58] = p.nodes[58].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[61].And(&v[22], &v[43])
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	v[67].Sub(&v[65], &v[66])
	v[68].Add(&v[44], &v[67])
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["1baf27255f3f68692a948cc48aedf34aa8c9677265dc44d3ab8abf4afe76b042"] = chainFastShape21
}

func chainFastShape22(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	v[41] = p.nodes[41].c
	v[42] = p.nodes[42].c
	v[43] = p.nodes[43].c
	v[44].And(&v[25], &v[42])
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[45] = p.nodes[45].c
	v[46] = p.nodes[46].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[47] = p.nodes[47].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[50] = p.nodes[50].c
	v[51].And(&v[49], &v[50])
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[53] = p.nodes[53].c
	v[54] = chainPure(OpCode(17), v[52], v[50])
	v[55] = chainPure(OpCode(21), v[54], v[43])
	v[56] = p.nodes[56].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[57] = p.nodes[57].c
	v[58] = p.nodes[58].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[60] = p.nodes[60].c
	v[61] = chainPure(OpCode(28), v[60], v[59])
	v[62] = p.nodes[62].c
	v[63] = chainPure(OpCode(20), v[61], v[62])
	v[64] = chainPure(OpCode(21), v[63], v[43])
	v[65] = p.nodes[65].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[66] = p.nodes[66].c
	v[67] = p.nodes[67].c
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[69].SetBytes(effectiveCaller(s))
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[70] = p.nodes[70].c
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[74] = p.nodes[74].c
	v[75].Sub(&v[72], &v[73])
	v[76].Add(&v[46], &v[75])
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	v[79] = p.nodes[79].c
	v[80].Sub(&v[77], &v[78])
	v[81].Add(&v[46], &v[80])
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["20cef1fa5fc3bd09aa7bd5586981cafc48c35f047e79302923b4475013e09ffb"] = chainFastShape22
}

func chainFastShape23(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35].Add(&v[31], &v[33])
	v[36] = p.nodes[36].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[38].Sub(&v[35], &v[37])
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	v[41] = p.nodes[41].c
	v[42] = p.nodes[42].c
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	v[45].And(&v[22], &v[44])
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[48] = p.nodes[48].c
	v[49] = p.nodes[49].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[51] = p.nodes[51].c
	v[52] = p.nodes[52].c
	v[53].And(&v[23], &v[44])
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[56] = p.nodes[56].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["2209fcc85e494f0497810113352b8e82d9571bdd16e16a2de970c37d745cde29"] = chainFastShape23
}

func chainFastShape24(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38].And(&v[28], &v[37])
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[41] = p.nodes[41].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[43] = p.nodes[43].c
	v[44].And(&v[29], &v[43])
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[47] = p.nodes[47].c
	v[48].Add(&v[46], &v[47])
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[50] = p.nodes[50].c
	v[51].And(&v[49], &v[50])
	v[52] = chainPure(OpCode(21), v[51], v[35])
	v[53] = chainPure(OpCode(21), v[52], v[35])
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[54].Add(&v[33], &v[45])
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[56].Add(&v[40], &v[46])
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[57] = p.nodes[57].c
	v[58].Add(&v[45], &v[57])
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[60].Add(&v[41], &v[46])
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[61] = p.nodes[61].c
	v[62].Add(&v[45], &v[61])
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[64] = p.nodes[64].c
	v[65].Add(&v[46], &v[64])
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[66] = p.nodes[66].c
	v[67].Add(&v[45], &v[66])
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[69] = p.nodes[69].c
	v[70].Add(&v[46], &v[69])
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[71] = p.nodes[71].c
	v[72].Add(&v[45], &v[71])
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[74].Add(&v[34], &v[46])
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[75] = p.nodes[75].c
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[76] = chainPure(OpCode(17), v[30], v[35])
	v[77] = chainPure(OpCode(21), v[76], v[35])
	v[78] = p.nodes[78].c
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[79] = p.nodes[79].c
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	v[82].Add(&v[40], &v[81])
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	v[83].Add(&v[41], &v[81])
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	v[84].Add(&v[64], &v[81])
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	v[85] = p.nodes[85].c
	v[86].Add(&v[69], &v[81])
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	v[88].Sub(&v[86], &v[87])
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["230400bb4a89fcd4ceb24fdebfbf4f4f6bd977411508cb9d2a7116ac883af831"] = chainFastShape24
}

func chainFastShape25(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33].SetBytes(effectiveCaller(s))
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38].And(&v[30], &v[37])
	v[39].And(&v[37], &v[38])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[40] = p.nodes[40].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[41] = p.nodes[41].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[44] = chainPure(OpCode(16), v[43], v[31])
	v[45] = chainPure(OpCode(21), v[44], v[35])
	v[46] = p.nodes[46].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[47].And(&v[33], &v[37])
	v[48] = chainPure(OpCode(20), v[38], v[47])
	v[49] = chainPure(OpCode(21), v[48], v[35])
	v[50] = chainPure(OpCode(21), v[49], v[35])
	v[51] = p.nodes[51].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[52] = p.nodes[52].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[53] = p.nodes[53].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[55] = p.nodes[55].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[57] = p.nodes[57].c
	v[58].And(&v[37], &v[47])
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[61] = p.nodes[61].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[63] = p.nodes[63].c
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[65].Add(&v[40], &v[64])
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	v[67].Sub(&v[65], &v[66])
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["23d1d3f8254728562c03bccf2683cfe24de588878f05c19813c5d4172f33d8b3"] = chainFastShape25
}

func chainFastShape26(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37].And(&v[27], &v[35])
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[40] = p.nodes[40].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[43] = p.nodes[43].c
	v[44] = chainPure(OpCode(28), v[43], v[42])
	v[45] = p.nodes[45].c
	v[46] = chainPure(OpCode(20), v[44], v[45])
	v[47] = chainPure(OpCode(21), v[46], v[36])
	v[48] = p.nodes[48].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[49] = p.nodes[49].c
	v[50].SetBytes(effectiveCaller(s))
	v[51] = p.nodes[51].c
	v[52].And(&v[35], &v[50])
	v[53] = p.nodes[53].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[54] = p.nodes[54].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[55] = p.nodes[55].c
	v[56] = p.nodes[56].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[59] = p.nodes[59].c
	v[60].And(&v[58], &v[59])
	v[61] = chainPure(OpCode(17), v[28], v[60])
	v[62] = chainPure(OpCode(21), v[61], v[36])
	v[63] = p.nodes[63].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[64] = p.nodes[64].c
	v[65] = p.nodes[65].c
	v[66] = p.nodes[66].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[68].And(&v[59], &v[67])
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[70] = p.nodes[70].c
	v[71] = chainPure(OpCode(17), v[69], v[59])
	v[72] = chainPure(OpCode(21), v[71], v[36])
	v[73] = p.nodes[73].c
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[74] = p.nodes[74].c
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[76] = chainPure(OpCode(28), v[43], v[75])
	v[77] = chainPure(OpCode(20), v[45], v[76])
	v[78] = chainPure(OpCode(21), v[77], v[36])
	v[79] = p.nodes[79].c
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	v[80] = p.nodes[80].c
	v[81] = p.nodes[81].c
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	v[83].And(&v[59], &v[82])
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	v[85] = chainPure(OpCode(17), v[84], v[59])
	v[86] = chainPure(OpCode(21), v[85], v[36])
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	v[88] = chainPure(OpCode(28), v[43], v[87])
	v[89] = chainPure(OpCode(20), v[45], v[88])
	v[90] = chainPure(OpCode(21), v[89], v[36])
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	v[91] = p.nodes[91].c
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	code, err = p.actionFast(44, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 44 != 0, err
	}
	v[93].Add(&v[39], &v[92])
	code, err = p.actionFast(45, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 45 != 0, err
	}
	v[95].Sub(&v[93], &v[94])
	code, err = p.actionFast(46, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 46 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["278b7570b9b7ab52b9ffeff74b573da063d5f719ee7a8ec313444d655345dc10"] = chainFastShape26
}

func chainFastShape27(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36].And(&v[29], &v[34])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[39] = p.nodes[39].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[42] = p.nodes[42].c
	v[43].And(&v[41], &v[42])
	v[44] = chainPure(OpCode(17), v[31], v[43])
	v[45] = chainPure(OpCode(21), v[44], v[35])
	v[46] = p.nodes[46].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[47] = p.nodes[47].c
	v[48] = p.nodes[48].c
	v[49] = p.nodes[49].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[51].And(&v[42], &v[50])
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[53] = p.nodes[53].c
	v[54] = chainPure(OpCode(17), v[52], v[42])
	v[55] = chainPure(OpCode(21), v[54], v[35])
	v[56] = p.nodes[56].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[57] = p.nodes[57].c
	v[58] = p.nodes[58].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[60] = p.nodes[60].c
	v[61] = chainPure(OpCode(28), v[60], v[59])
	v[62] = p.nodes[62].c
	v[63] = chainPure(OpCode(20), v[61], v[62])
	v[64] = chainPure(OpCode(21), v[63], v[35])
	v[65] = p.nodes[65].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[66] = p.nodes[66].c
	v[67] = p.nodes[67].c
	v[68].And(&v[30], &v[34])
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[71].And(&v[42], &v[70])
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[73] = chainPure(OpCode(17), v[72], v[42])
	v[74] = chainPure(OpCode(21), v[73], v[35])
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	v[76] = chainPure(OpCode(28), v[60], v[75])
	v[77] = chainPure(OpCode(20), v[62], v[76])
	v[78] = chainPure(OpCode(21), v[77], v[35])
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	v[79] = p.nodes[79].c
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	v[81].Add(&v[38], &v[80])
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	v[83].Sub(&v[81], &v[82])
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["2948be0b415eabb8cc5fc7a14c5d16521ec19518bb58903a7ff967554a6997f7"] = chainFastShape27
}

func chainFastShape28(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33].SetBytes(effectiveCaller(s))
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38].And(&v[33], &v[36])
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[41] = p.nodes[41].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[44] = p.nodes[44].c
	v[45] = chainPure(OpCode(28), v[44], v[43])
	v[46] = p.nodes[46].c
	v[47] = chainPure(OpCode(20), v[45], v[46])
	v[48] = chainPure(OpCode(21), v[47], v[37])
	v[49] = p.nodes[49].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[50] = p.nodes[50].c
	v[51].And(&v[28], &v[36])
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[54] = chainPure(OpCode(28), v[44], v[53])
	v[55] = chainPure(OpCode(20), v[46], v[54])
	v[56] = chainPure(OpCode(21), v[55], v[37])
	v[57] = p.nodes[57].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[58] = p.nodes[58].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[59] = chainPure(OpCode(17), v[29], v[37])
	v[60] = p.nodes[60].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[61] = p.nodes[61].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[64] = chainPure(OpCode(17), v[29], v[63])
	v[65] = chainPure(OpCode(21), v[64], v[37])
	v[66] = p.nodes[66].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[67] = p.nodes[67].c
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[69] = p.nodes[69].c
	v[70] = p.nodes[70].c
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[72] = p.nodes[72].c
	v[73] = p.nodes[73].c
	v[74] = p.nodes[74].c
	v[75] = p.nodes[75].c
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[77] = p.nodes[77].c
	v[78].And(&v[76], &v[77])
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	v[80] = p.nodes[80].c
	v[81] = chainPure(OpCode(17), v[79], v[77])
	v[82] = chainPure(OpCode(21), v[81], v[37])
	v[83] = p.nodes[83].c
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[84] = p.nodes[84].c
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	v[86] = chainPure(OpCode(28), v[44], v[85])
	v[87] = chainPure(OpCode(20), v[46], v[86])
	v[88] = chainPure(OpCode(21), v[87], v[37])
	v[89] = p.nodes[89].c
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	v[90] = p.nodes[90].c
	v[91] = p.nodes[91].c
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	v[95] = p.nodes[95].c
	v[96].Sub(&v[93], &v[94])
	v[97].Add(&v[40], &v[96])
	code, err = p.actionFast(44, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 44 != 0, err
	}
	code, err = p.actionFast(45, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 45 != 0, err
	}
	code, err = p.actionFast(46, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 46 != 0, err
	}
	code, err = p.actionFast(47, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 47 != 0, err
	}
	v[100] = p.nodes[100].c
	v[101].Sub(&v[98], &v[99])
	v[102].Add(&v[40], &v[101])
	code, err = p.actionFast(48, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 48 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["2dbcaa39b088de97b24fb1f0c03b592228e2efed92badf93cfeeead0d92c1aa9"] = chainFastShape28
}

func chainFastShape29(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33].SetBytes(effectiveCaller(s))
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38].And(&v[33], &v[36])
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[41] = p.nodes[41].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[44] = p.nodes[44].c
	v[45] = chainPure(OpCode(28), v[44], v[43])
	v[46] = p.nodes[46].c
	v[47] = chainPure(OpCode(20), v[45], v[46])
	v[48] = chainPure(OpCode(21), v[47], v[37])
	v[49] = p.nodes[49].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[50] = p.nodes[50].c
	v[51] = p.nodes[51].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[53] = p.nodes[53].c
	v[54].And(&v[52], &v[53])
	v[55] = chainPure(OpCode(17), v[30], v[37])
	v[56] = p.nodes[56].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[57] = chainPure(OpCode(16), v[54], v[30])
	v[58] = chainPure(OpCode(21), v[57], v[37])
	v[59] = p.nodes[59].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[60] = p.nodes[60].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[62] = p.nodes[62].c
	v[63] = p.nodes[63].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[65] = p.nodes[65].c
	v[66] = p.nodes[66].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[68] = p.nodes[68].c
	v[69] = chainPure(OpCode(17), v[67], v[53])
	v[70] = chainPure(OpCode(21), v[69], v[37])
	v[71] = p.nodes[71].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[72] = p.nodes[72].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[74] = chainPure(OpCode(28), v[44], v[73])
	v[75] = chainPure(OpCode(20), v[46], v[74])
	v[76] = chainPure(OpCode(21), v[75], v[37])
	v[77] = p.nodes[77].c
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	v[80] = p.nodes[80].c
	v[81].Sub(&v[78], &v[79])
	v[82].Add(&v[40], &v[81])
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	v[85] = p.nodes[85].c
	v[86].Sub(&v[83], &v[84])
	v[87].Add(&v[40], &v[86])
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["2ded5a85eabc49169eae49702d1edb8b71a698bd16ba46b2463ad50b434eb589"] = chainFastShape29
}

func chainFastShape30(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37].And(&v[26], &v[35])
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[40] = p.nodes[40].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[43] = p.nodes[43].c
	v[44] = chainPure(OpCode(28), v[43], v[42])
	v[45] = p.nodes[45].c
	v[46] = chainPure(OpCode(20), v[44], v[45])
	v[47] = chainPure(OpCode(21), v[46], v[36])
	v[48] = p.nodes[48].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[49].And(&v[25], &v[35])
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[50] = p.nodes[50].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[52].SetBytes(effectiveCaller(s))
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[55] = chainPure(OpCode(17), v[27], v[54])
	v[56] = chainPure(OpCode(21), v[55], v[36])
	v[57] = p.nodes[57].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[58] = p.nodes[58].c
	v[59] = p.nodes[59].c
	v[60] = p.nodes[60].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[61] = p.nodes[61].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[62] = p.nodes[62].c
	v[63] = p.nodes[63].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[66] = p.nodes[66].c
	v[67].And(&v[65], &v[66])
	v[68] = chainPure(OpCode(17), v[27], v[67])
	v[69] = chainPure(OpCode(21), v[68], v[36])
	v[70] = p.nodes[70].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[71] = p.nodes[71].c
	v[72] = p.nodes[72].c
	v[73] = p.nodes[73].c
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[75].And(&v[66], &v[74])
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	v[77] = p.nodes[77].c
	v[78] = chainPure(OpCode(17), v[76], v[66])
	v[79] = chainPure(OpCode(21), v[78], v[36])
	v[80] = p.nodes[80].c
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[81] = p.nodes[81].c
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	v[83] = chainPure(OpCode(28), v[43], v[82])
	v[84] = chainPure(OpCode(20), v[45], v[83])
	v[85] = chainPure(OpCode(21), v[84], v[36])
	v[86] = p.nodes[86].c
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	v[87] = p.nodes[87].c
	v[88] = p.nodes[88].c
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	v[90].And(&v[66], &v[89])
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	v[92] = chainPure(OpCode(17), v[91], v[66])
	v[93] = chainPure(OpCode(21), v[92], v[36])
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	code, err = p.actionFast(44, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 44 != 0, err
	}
	code, err = p.actionFast(45, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 45 != 0, err
	}
	v[95] = chainPure(OpCode(28), v[43], v[94])
	v[96] = chainPure(OpCode(20), v[45], v[95])
	v[97] = chainPure(OpCode(21), v[96], v[36])
	code, err = p.actionFast(46, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 46 != 0, err
	}
	code, err = p.actionFast(47, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 47 != 0, err
	}
	code, err = p.actionFast(48, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 48 != 0, err
	}
	code, err = p.actionFast(49, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 49 != 0, err
	}
	code, err = p.actionFast(50, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 50 != 0, err
	}
	v[98] = p.nodes[98].c
	code, err = p.actionFast(51, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 51 != 0, err
	}
	code, err = p.actionFast(52, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 52 != 0, err
	}
	v[100].Add(&v[39], &v[99])
	code, err = p.actionFast(53, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 53 != 0, err
	}
	v[102].Sub(&v[100], &v[101])
	code, err = p.actionFast(54, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 54 != 0, err
	}
	code, err = p.actionFast(55, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 55 != 0, err
	}
	code, err = p.actionFast(56, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 56 != 0, err
	}
	code, err = p.actionFast(57, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 57 != 0, err
	}
	code, err = p.actionFast(58, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 58 != 0, err
	}
	code, err = p.actionFast(59, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 59 != 0, err
	}
	code, err = p.actionFast(60, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 60 != 0, err
	}
	code, err = p.actionFast(61, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 61 != 0, err
	}
	v[104] = p.nodes[104].c
	v[105] = p.nodes[105].c
	code, err = p.actionFast(62, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 62 != 0, err
	}
	code, err = p.actionFast(63, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 63 != 0, err
	}
	code, err = p.actionFast(64, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 64 != 0, err
	}
	code, err = p.actionFast(65, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 65 != 0, err
	}
	code, err = p.actionFast(66, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 66 != 0, err
	}
	code, err = p.actionFast(67, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 67 != 0, err
	}
	code, err = p.actionFast(68, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 68 != 0, err
	}
	code, err = p.actionFast(69, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 69 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["2f03404a4b86833fe23eb4029c9ce82f7eb080c2638b1591144ba1ae94095bfe"] = chainFastShape30
}

func chainFastShape31(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	v[39].And(&v[28], &v[37])
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[40] = p.nodes[40].c
	v[41] = p.nodes[41].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[42] = p.nodes[42].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[45] = p.nodes[45].c
	v[46].And(&v[44], &v[45])
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[48] = p.nodes[48].c
	v[49] = chainPure(OpCode(17), v[47], v[45])
	v[50] = chainPure(OpCode(21), v[49], v[38])
	v[51] = p.nodes[51].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[52] = p.nodes[52].c
	v[53] = p.nodes[53].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[55] = p.nodes[55].c
	v[56] = chainPure(OpCode(28), v[55], v[54])
	v[57] = p.nodes[57].c
	v[58] = chainPure(OpCode(20), v[56], v[57])
	v[59] = chainPure(OpCode(21), v[58], v[38])
	v[60] = p.nodes[60].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[61] = p.nodes[61].c
	v[62] = p.nodes[62].c
	v[63].And(&v[29], &v[37])
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[66].And(&v[45], &v[65])
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[68] = chainPure(OpCode(17), v[67], v[45])
	v[69] = chainPure(OpCode(21), v[68], v[38])
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	v[71] = chainPure(OpCode(28), v[55], v[70])
	v[72] = chainPure(OpCode(20), v[57], v[71])
	v[73] = chainPure(OpCode(21), v[72], v[38])
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	v[74] = p.nodes[74].c
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	v[76].Add(&v[41], &v[75])
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	v[78].Sub(&v[76], &v[77])
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["30fe27f0d5e9c10dc67b1ba7841db4fba7604916bf75ad33401cbb96da8841a1"] = chainFastShape31
}

func chainFastShape32(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33].SetBytes(effectiveCaller(s))
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37].And(&v[33], &v[35])
	v[38] = p.nodes[38].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[39].And(&v[27], &v[35])
	v[40] = p.nodes[40].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[41] = p.nodes[41].c
	v[42] = p.nodes[42].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[45] = p.nodes[45].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[48] = p.nodes[48].c
	v[49].And(&v[47], &v[48])
	v[50] = chainPure(OpCode(17), v[28], v[49])
	v[51] = chainPure(OpCode(21), v[50], v[36])
	v[52] = p.nodes[52].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[53] = p.nodes[53].c
	v[54] = p.nodes[54].c
	v[55] = p.nodes[55].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[57].And(&v[48], &v[56])
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[59] = p.nodes[59].c
	v[60] = chainPure(OpCode(17), v[58], v[48])
	v[61] = chainPure(OpCode(21), v[60], v[36])
	v[62] = p.nodes[62].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[63] = p.nodes[63].c
	v[64] = p.nodes[64].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[66] = p.nodes[66].c
	v[67] = chainPure(OpCode(28), v[66], v[65])
	v[68] = p.nodes[68].c
	v[69] = chainPure(OpCode(20), v[67], v[68])
	v[70] = chainPure(OpCode(21), v[69], v[36])
	v[71] = p.nodes[71].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[72] = p.nodes[72].c
	v[73] = p.nodes[73].c
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	v[76].And(&v[48], &v[75])
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[78] = chainPure(OpCode(17), v[77], v[48])
	v[79] = chainPure(OpCode(21), v[78], v[36])
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	v[81] = chainPure(OpCode(28), v[66], v[80])
	v[82] = chainPure(OpCode(20), v[68], v[81])
	v[83] = chainPure(OpCode(21), v[82], v[36])
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	v[84] = p.nodes[84].c
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	v[86].Add(&v[44], &v[85])
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	v[88].Sub(&v[86], &v[87])
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["332413647442373380504859f448957b325709d0db2232b37d27223610de9b4a"] = chainFastShape32
}

func chainFastShape33(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38].And(&v[29], &v[37])
	v[39] = p.nodes[39].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[40] = chainPure(OpCode(17), v[30], v[35])
	v[41] = p.nodes[41].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[42] = p.nodes[42].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[44].And(&v[37], &v[43])
	v[45].And(&v[37], &v[44])
	v[46].And(&v[28], &v[37])
	v[47] = chainPure(OpCode(20), v[45], v[46])
	v[48] = chainPure(OpCode(21), v[47], v[35])
	v[49] = chainPure(OpCode(21), v[48], v[35])
	v[50] = p.nodes[50].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[51] = p.nodes[51].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[52] = p.nodes[52].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[53] = p.nodes[53].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[54] = p.nodes[54].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[55] = p.nodes[55].c
	v[56] = p.nodes[56].c
	v[57] = p.nodes[57].c
	v[58] = p.nodes[58].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[59] = p.nodes[59].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[60] = p.nodes[60].c
	v[61] = p.nodes[61].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[63] = p.nodes[63].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[64] = p.nodes[64].c
	v[65].Add(&v[62], &v[64])
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[67].Sub(&v[65], &v[66])
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[68] = p.nodes[68].c
	v[69] = p.nodes[69].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[72] = p.nodes[72].c
	v[73] = p.nodes[73].c
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[75] = p.nodes[75].c
	v[76] = p.nodes[76].c
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	v[79] = p.nodes[79].c
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["3351ded32d8a7f95a2a68def1deb845b7a00ddf88aa6ac60a440fad62e4bac73"] = chainFastShape33
}

func chainFastShape34(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35].And(&v[26], &v[33])
	v[36] = p.nodes[36].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[37] = chainPure(OpCode(17), v[27], v[34])
	v[38] = p.nodes[38].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[39].SetBytes(effectiveCaller(s))
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[40] = p.nodes[40].c
	v[41] = p.nodes[41].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[42] = p.nodes[42].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[45] = chainPure(OpCode(17), v[27], v[44])
	v[46] = chainPure(OpCode(21), v[45], v[34])
	v[47] = p.nodes[47].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[48] = p.nodes[48].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[50] = p.nodes[50].c
	v[51] = p.nodes[51].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[53] = p.nodes[53].c
	v[54] = p.nodes[54].c
	v[55] = p.nodes[55].c
	v[56] = p.nodes[56].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[57] = p.nodes[57].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[60] = p.nodes[60].c
	v[61].And(&v[59], &v[60])
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[63] = p.nodes[63].c
	v[64] = chainPure(OpCode(17), v[62], v[60])
	v[65] = chainPure(OpCode(21), v[64], v[34])
	v[66] = p.nodes[66].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[67] = p.nodes[67].c
	v[68] = p.nodes[68].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[70] = p.nodes[70].c
	v[71] = chainPure(OpCode(28), v[70], v[69])
	v[72] = p.nodes[72].c
	v[73] = chainPure(OpCode(20), v[71], v[72])
	v[74] = chainPure(OpCode(21), v[73], v[34])
	v[75] = p.nodes[75].c
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	v[76] = p.nodes[76].c
	v[77] = p.nodes[77].c
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	v[81] = p.nodes[81].c
	v[82].Sub(&v[79], &v[80])
	v[83].Add(&v[41], &v[82])
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	v[86] = p.nodes[86].c
	v[87].Sub(&v[84], &v[85])
	v[88].Add(&v[41], &v[87])
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["34cba75e9878f86ec7b0de8421a38ab904fedefe92db64d6e879009391a91301"] = chainFastShape34
}

func chainFastShape35(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	v[40].And(&v[34], &v[39])
	v[41] = p.nodes[41].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[42].And(&v[29], &v[39])
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[45] = p.nodes[45].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[47] = p.nodes[47].c
	v[48].And(&v[30], &v[47])
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[51] = p.nodes[51].c
	v[52].And(&v[50], &v[51])
	v[53] = chainPure(OpCode(21), v[52], v[37])
	v[54] = chainPure(OpCode(21), v[53], v[37])
	v[55] = chainPure(OpCode(20), v[35], v[54])
	v[56] = p.nodes[56].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[58] = p.nodes[58].c
	v[59].And(&v[57], &v[58])
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[60].Add(&v[35], &v[49])
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[61] = p.nodes[61].c
	v[62].Add(&v[49], &v[61])
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[63] = p.nodes[63].c
	v[64].Add(&v[49], &v[63])
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[65] = p.nodes[65].c
	v[66].Add(&v[49], &v[65])
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[67] = p.nodes[67].c
	v[68].Add(&v[49], &v[67])
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	v[70] = p.nodes[70].c
	v[71].Add(&v[44], &v[69])
	v[72] = p.nodes[72].c
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[74].Sub(&v[71], &v[73])
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["39026ad16a5112c83a14100695222a9a9c22257a00ce6414f0cd449deb7d0d57"] = chainFastShape35
}

func chainFastShape36(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	v[40].SetBytes(effectiveCaller(s))
	v[41] = p.nodes[41].c
	v[42].And(&v[40], &v[41])
	v[43].And(&v[41], &v[42])
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[44] = p.nodes[44].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[45] = p.nodes[45].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[49] = p.nodes[49].c
	v[50].And(&v[26], &v[41])
	v[51].And(&v[41], &v[50])
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[54] = p.nodes[54].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[56] = chainPure(OpCode(17), v[28], v[36])
	v[57] = chainPure(OpCode(21), v[56], v[36])
	v[58] = p.nodes[58].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[59] = p.nodes[59].c
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[61].Add(&v[44], &v[60])
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[63].Sub(&v[61], &v[62])
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["399aded9af9c8864b5b0e54872560c405d11658e5aa9edc37796699a35e3b79c"] = chainFastShape36
}

func chainFastShape37(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[35] = p.nodes[35].c
	v[36] = chainPure(OpCode(17), v[34], v[29])
	v[37] = chainPure(OpCode(21), v[36], v[35])
	v[38] = p.nodes[38].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	v[41] = p.nodes[41].c
	v[42] = p.nodes[42].c
	v[43].And(&v[27], &v[42])
	v[44] = chainPure(OpCode(21), v[43], v[35])
	v[45] = p.nodes[45].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[46] = chainPure(OpCode(21), v[28], v[35])
	v[47] = p.nodes[47].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[48] = p.nodes[48].c
	v[49] = p.nodes[49].c
	v[50] = p.nodes[50].c
	v[51] = p.nodes[51].c
	v[52].And(&v[25], &v[51])
	v[53] = p.nodes[53].c
	v[54] = p.nodes[54].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[56] = p.nodes[56].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[57] = p.nodes[57].c
	v[58].Add(&v[55], &v[57])
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[60].Sub(&v[58], &v[59])
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[61] = p.nodes[61].c
	v[62] = p.nodes[62].c
	v[63].And(&v[24], &v[51])
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[66] = p.nodes[66].c
	v[67] = p.nodes[67].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[69] = p.nodes[69].c
	v[70] = p.nodes[70].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["41b51ed5d890ca7adc6148192804a1b55ce7cb975b76130c113ce6506b7cf397"] = chainFastShape37
}

func chainFastShape38(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33].SetBytes(effectiveCaller(s))
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38].And(&v[33], &v[37])
	v[39].And(&v[37], &v[38])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[40] = p.nodes[40].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[41] = p.nodes[41].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[44] = chainPure(OpCode(16), v[43], v[31])
	v[45] = chainPure(OpCode(21), v[44], v[35])
	v[46] = p.nodes[46].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[47] = chainPure(OpCode(20), v[38], v[38])
	v[48] = chainPure(OpCode(21), v[47], v[35])
	v[49] = chainPure(OpCode(21), v[48], v[35])
	v[50] = p.nodes[50].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[51] = p.nodes[51].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[52] = p.nodes[52].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[54] = p.nodes[54].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[56] = p.nodes[56].c
	v[57].And(&v[30], &v[37])
	v[58].And(&v[37], &v[57])
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[61] = p.nodes[61].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[63] = p.nodes[63].c
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[65].Add(&v[40], &v[64])
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	v[67].Sub(&v[65], &v[66])
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["438d1d9db5fc76a59414d01ae116db01d692ba651ee57ab939a55cb6af0d481e"] = chainFastShape38
}

func chainFastShape39(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36].And(&v[24], &v[35])
	v[37].And(&v[35], &v[36])
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[38] = p.nodes[38].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[39] = p.nodes[39].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[41].SetBytes(effectiveCaller(s))
	v[42].And(&v[35], &v[41])
	v[43].And(&v[35], &v[42])
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[46] = p.nodes[46].c
	v[47] = p.nodes[47].c
	v[48] = p.nodes[48].c
	v[49] = p.nodes[49].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[51] = p.nodes[51].c
	v[52] = p.nodes[52].c
	v[53] = chainPure(OpCode(20), v[26], v[34])
	v[54] = chainPure(OpCode(21), v[53], v[34])
	v[55] = p.nodes[55].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[56].Mul(&v[26], &v[50])
	v[57] = chainPure(OpCode(21), v[26], v[34])
	v[58] = chainPure(OpCode(21), v[57], v[34])
	v[59] = p.nodes[59].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[60].Div(&v[56], &v[26])
	v[61] = chainPure(OpCode(20), v[50], v[60])
	v[62] = chainPure(OpCode(21), v[61], v[34])
	v[63] = chainPure(OpCode(21), v[62], v[34])
	v[64] = p.nodes[64].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[65] = p.nodes[65].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[67] = p.nodes[67].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[69] = chainPure(OpCode(17), v[66], v[68])
	v[70] = chainPure(OpCode(21), v[69], v[34])
	v[71] = p.nodes[71].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[72] = p.nodes[72].c
	v[73] = chainPure(OpCode(16), v[45], v[72])
	v[74] = chainPure(OpCode(21), v[73], v[34])
	v[75] = p.nodes[75].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[76] = p.nodes[76].c
	v[77] = p.nodes[77].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[79] = p.nodes[79].c
	v[80] = p.nodes[80].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	v[84] = p.nodes[84].c
	v[85].And(&v[25], &v[35])
	v[86].And(&v[35], &v[85])
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	v[89] = p.nodes[89].c
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	v[91] = chainPure(OpCode(17), v[66], v[34])
	v[92] = chainPure(OpCode(21), v[91], v[34])
	v[93] = p.nodes[93].c
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	v[94] = p.nodes[94].c
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	v[96].Add(&v[38], &v[95])
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	v[98].Sub(&v[96], &v[97])
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["452514790ef91140bd39bb440188b0eb425e3358df7d2e74d6a06e41687aab98"] = chainFastShape39
}

func chainFastShape40(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37].And(&v[24], &v[35])
	v[38] = p.nodes[38].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[39].And(&v[25], &v[35])
	v[40] = p.nodes[40].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[41] = p.nodes[41].c
	v[42] = p.nodes[42].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[45] = p.nodes[45].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[48] = p.nodes[48].c
	v[49].And(&v[47], &v[48])
	v[50] = chainPure(OpCode(17), v[26], v[49])
	v[51] = chainPure(OpCode(21), v[50], v[36])
	v[52] = p.nodes[52].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[53] = p.nodes[53].c
	v[54] = p.nodes[54].c
	v[55] = p.nodes[55].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[57].And(&v[48], &v[56])
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[59] = p.nodes[59].c
	v[60] = chainPure(OpCode(17), v[58], v[48])
	v[61] = chainPure(OpCode(21), v[60], v[36])
	v[62] = p.nodes[62].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[63] = p.nodes[63].c
	v[64] = p.nodes[64].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[66] = p.nodes[66].c
	v[67] = chainPure(OpCode(28), v[66], v[65])
	v[68] = p.nodes[68].c
	v[69] = chainPure(OpCode(20), v[67], v[68])
	v[70] = chainPure(OpCode(21), v[69], v[36])
	v[71] = p.nodes[71].c
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[72] = p.nodes[72].c
	v[73] = p.nodes[73].c
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[76].And(&v[48], &v[75])
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	v[78] = chainPure(OpCode(17), v[77], v[48])
	v[79] = chainPure(OpCode(21), v[78], v[36])
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	v[81] = chainPure(OpCode(28), v[66], v[80])
	v[82] = chainPure(OpCode(20), v[68], v[81])
	v[83] = chainPure(OpCode(21), v[82], v[36])
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	v[84] = p.nodes[84].c
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	v[86].Add(&v[44], &v[85])
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	v[88].Sub(&v[86], &v[87])
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	v[89] = p.nodes[89].c
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	code, err = p.actionFast(44, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 44 != 0, err
	}
	v[91].SetBytes(effectiveCaller(s))
	code, err = p.actionFast(45, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 45 != 0, err
	}
	code, err = p.actionFast(46, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 46 != 0, err
	}
	code, err = p.actionFast(47, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 47 != 0, err
	}
	code, err = p.actionFast(48, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 48 != 0, err
	}
	v[94] = p.nodes[94].c
	v[95] = p.nodes[95].c
	code, err = p.actionFast(49, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 49 != 0, err
	}
	code, err = p.actionFast(50, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 50 != 0, err
	}
	code, err = p.actionFast(51, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 51 != 0, err
	}
	code, err = p.actionFast(52, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 52 != 0, err
	}
	code, err = p.actionFast(53, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 53 != 0, err
	}
	code, err = p.actionFast(54, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 54 != 0, err
	}
	code, err = p.actionFast(55, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 55 != 0, err
	}
	code, err = p.actionFast(56, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 56 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["462c76723f2af1ec9d0000906d68150aaac4032e139644f152b5a4d1f8eea6e2"] = chainFastShape40
}

func chainFastShape41(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36].And(&v[28], &v[35])
	v[37].And(&v[35], &v[36])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[38] = p.nodes[38].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[39] = p.nodes[39].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[42] = p.nodes[42].c
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[46].Add(&v[39], &v[45])
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[47] = p.nodes[47].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[48].Add(&v[38], &v[45])
	v[49] = p.nodes[49].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[50] = p.nodes[50].c
	v[51] = chainPure(OpCode(17), v[30], v[41])
	v[52] = chainPure(OpCode(21), v[51], v[34])
	v[53] = p.nodes[53].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[54] = p.nodes[54].c
	v[55] = p.nodes[55].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[57] = p.nodes[57].c
	v[58] = p.nodes[58].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[60].Add(&v[39], &v[59])
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[61].Add(&v[38], &v[59])
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[62] = chainPure(OpCode(17), v[31], v[30])
	v[63] = chainPure(OpCode(21), v[62], v[34])
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[65].And(&v[29], &v[35])
	v[66].And(&v[35], &v[65])
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[69] = p.nodes[69].c
	v[70] = p.nodes[70].c
	v[71] = p.nodes[71].c
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[73] = chainPure(OpCode(16), v[72], v[68])
	v[74] = chainPure(OpCode(21), v[73], v[34])
	v[75] = p.nodes[75].c
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	v[76] = p.nodes[76].c
	v[77] = p.nodes[77].c
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	v[79].Add(&v[39], &v[78])
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	v[80].Add(&v[38], &v[78])
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	v[83] = p.nodes[83].c
	v[84] = p.nodes[84].c
	v[85].Add(&v[38], &v[82])
	v[86] = p.nodes[86].c
	v[87].Add(&v[34], &v[82])
	v[88] = p.nodes[88].c
	v[89] = p.nodes[89].c
	v[90] = p.nodes[90].c
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	v[92].Sub(&v[85], &v[91])
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["4a038cd4d78d47f63effedb73f97f1eec2d07428c28213c6c737bf8e5fb87e4e"] = chainFastShape41
}

func chainFastShape42(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36].And(&v[31], &v[34])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[39] = p.nodes[39].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[42] = p.nodes[42].c
	v[43] = chainPure(OpCode(28), v[42], v[41])
	v[44] = p.nodes[44].c
	v[45] = chainPure(OpCode(20), v[43], v[44])
	v[46] = chainPure(OpCode(21), v[45], v[35])
	v[47] = p.nodes[47].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[48] = p.nodes[48].c
	v[49].And(&v[27], &v[34])
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[52] = chainPure(OpCode(28), v[42], v[51])
	v[53] = chainPure(OpCode(20), v[44], v[52])
	v[54] = chainPure(OpCode(21), v[53], v[35])
	v[55] = p.nodes[55].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[56] = p.nodes[56].c
	v[57].And(&v[28], &v[34])
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[60] = chainPure(OpCode(28), v[42], v[59])
	v[61] = chainPure(OpCode(20), v[44], v[60])
	v[62] = chainPure(OpCode(21), v[61], v[35])
	v[63] = p.nodes[63].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[64] = p.nodes[64].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[66].SetBytes(effectiveCaller(s))
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[69] = chainPure(OpCode(17), v[29], v[68])
	v[70] = chainPure(OpCode(21), v[69], v[35])
	v[71] = p.nodes[71].c
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[72] = p.nodes[72].c
	v[73] = p.nodes[73].c
	v[74] = p.nodes[74].c
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[75] = p.nodes[75].c
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[76] = p.nodes[76].c
	v[77] = p.nodes[77].c
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	v[79] = p.nodes[79].c
	v[80].And(&v[78], &v[79])
	v[81] = chainPure(OpCode(17), v[29], v[80])
	v[82] = chainPure(OpCode(21), v[81], v[35])
	v[83] = p.nodes[83].c
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	v[84] = p.nodes[84].c
	v[85] = p.nodes[85].c
	v[86] = p.nodes[86].c
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	v[88].And(&v[79], &v[87])
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	v[90] = p.nodes[90].c
	v[91] = chainPure(OpCode(17), v[89], v[79])
	v[92] = chainPure(OpCode(21), v[91], v[35])
	v[93] = p.nodes[93].c
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	v[94] = p.nodes[94].c
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	v[96] = chainPure(OpCode(28), v[42], v[95])
	v[97] = chainPure(OpCode(20), v[44], v[96])
	v[98] = chainPure(OpCode(21), v[97], v[35])
	v[99] = p.nodes[99].c
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	code, err = p.actionFast(44, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 44 != 0, err
	}
	v[100] = p.nodes[100].c
	v[101] = p.nodes[101].c
	code, err = p.actionFast(45, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 45 != 0, err
	}
	code, err = p.actionFast(46, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 46 != 0, err
	}
	code, err = p.actionFast(47, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 47 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["4b7d58ae45579c23e37e9059147eed06b79bcf81843eca0d07cfb09224c488b7"] = chainFastShape42
}

func chainFastShape43(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36].Add(&v[31], &v[35])
	v[37] = p.nodes[37].c
	v[38].Add(&v[31], &v[34])
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	v[41] = p.nodes[41].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[42] = p.nodes[42].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[44].Sub(&v[36], &v[43])
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[45] = p.nodes[45].c
	v[46] = p.nodes[46].c
	v[47] = p.nodes[47].c
	v[48].And(&v[23], &v[47])
	v[49].And(&v[47], &v[48])
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[52] = p.nodes[52].c
	v[53] = p.nodes[53].c
	v[54] = p.nodes[54].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[56].Add(&v[42], &v[55])
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[57] = p.nodes[57].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[58].Add(&v[35], &v[55])
	v[59] = p.nodes[59].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[60] = p.nodes[60].c
	v[61] = chainPure(OpCode(17), v[25], v[51])
	v[62] = chainPure(OpCode(21), v[61], v[34])
	v[63] = p.nodes[63].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[64] = p.nodes[64].c
	v[65] = p.nodes[65].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[67] = p.nodes[67].c
	v[68] = p.nodes[68].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[70].Add(&v[42], &v[69])
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[71].Add(&v[35], &v[69])
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[72] = chainPure(OpCode(17), v[26], v[25])
	v[73] = chainPure(OpCode(21), v[72], v[34])
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[75].And(&v[24], &v[47])
	v[76].And(&v[47], &v[75])
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[79] = p.nodes[79].c
	v[80] = p.nodes[80].c
	v[81] = p.nodes[81].c
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	v[83] = chainPure(OpCode(16), v[82], v[78])
	v[84] = chainPure(OpCode(21), v[83], v[34])
	v[85] = p.nodes[85].c
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	v[86] = p.nodes[86].c
	v[87] = p.nodes[87].c
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	v[89].Add(&v[42], &v[88])
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	v[90].Add(&v[35], &v[88])
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	v[93] = p.nodes[93].c
	v[94].Add(&v[35], &v[92])
	v[95].Add(&v[34], &v[92])
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	v[97].Sub(&v[94], &v[96])
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["4c058a2b42d9d325eb5c6684e1aa469f9d84f747f108d5638c6c19889e87bc27"] = chainFastShape43
}

func chainFastShape44(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33].SetBytes(effectiveCaller(s))
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37].And(&v[33], &v[35])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[40] = p.nodes[40].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[43] = p.nodes[43].c
	v[44].And(&v[42], &v[43])
	v[45] = chainPure(OpCode(17), v[29], v[36])
	v[46] = p.nodes[46].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[47] = chainPure(OpCode(16), v[44], v[29])
	v[48] = chainPure(OpCode(21), v[47], v[36])
	v[49] = p.nodes[49].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[50] = p.nodes[50].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[52] = p.nodes[52].c
	v[53] = p.nodes[53].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[55] = p.nodes[55].c
	v[56] = p.nodes[56].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[58] = p.nodes[58].c
	v[59] = chainPure(OpCode(17), v[57], v[43])
	v[60] = chainPure(OpCode(21), v[59], v[36])
	v[61] = p.nodes[61].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[62] = p.nodes[62].c
	v[63] = p.nodes[63].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[65] = p.nodes[65].c
	v[66] = chainPure(OpCode(28), v[65], v[64])
	v[67] = p.nodes[67].c
	v[68] = chainPure(OpCode(20), v[66], v[67])
	v[69] = chainPure(OpCode(21), v[68], v[36])
	v[70] = p.nodes[70].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[73] = p.nodes[73].c
	v[74].Sub(&v[71], &v[72])
	v[75].Add(&v[39], &v[74])
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[78] = p.nodes[78].c
	v[79].Sub(&v[76], &v[77])
	v[80].Add(&v[39], &v[79])
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["4fbab321deef9978a3b9bf27ac9df3195b14f6ed5b1e280424b85e5b021083de"] = chainFastShape44
}

func chainFastShape45(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[37] = chainPure(OpCode(16), v[36], v[35])
	v[38] = chainPure(OpCode(21), v[37], v[34])
	v[39] = p.nodes[39].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[40] = p.nodes[40].c
	v[41].Add(&v[28], &v[40])
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[43] = chainPure(OpCode(16), v[42], v[35])
	v[44] = chainPure(OpCode(21), v[43], v[34])
	v[45] = p.nodes[45].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[47].Add(&v[40], &v[41])
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[49] = p.nodes[49].c
	v[50] = chainPure(OpCode(20), v[46], v[49])
	v[51] = chainPure(OpCode(21), v[50], v[34])
	v[52] = p.nodes[52].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[53] = p.nodes[53].c
	v[54] = p.nodes[54].c
	v[55] = p.nodes[55].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[56] = p.nodes[56].c
	v[57] = p.nodes[57].c
	v[58] = p.nodes[58].c
	v[59] = p.nodes[59].c
	v[60].And(&v[25], &v[59])
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[61] = p.nodes[61].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[64] = p.nodes[64].c
	v[65] = p.nodes[65].c
	v[66] = p.nodes[66].c
	v[67] = p.nodes[67].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[69].Add(&v[35], &v[68])
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[70] = p.nodes[70].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[71].Add(&v[40], &v[68])
	v[72] = p.nodes[72].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[73] = p.nodes[73].c
	v[74] = chainPure(OpCode(17), v[27], v[63])
	v[75] = chainPure(OpCode(21), v[74], v[34])
	v[76] = p.nodes[76].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[77].Sub(&v[63], &v[27])
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[78] = p.nodes[78].c
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[82] = p.nodes[82].c
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	v[84].Add(&v[35], &v[83])
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	v[85].Add(&v[40], &v[83])
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	v[86] = chainPure(OpCode(17), v[27], v[81])
	v[87] = chainPure(OpCode(21), v[86], v[34])
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	v[88].Sub(&v[81], &v[27])
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	v[89] = p.nodes[89].c
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	v[92] = p.nodes[92].c
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	v[94].Add(&v[35], &v[93])
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	code, err = p.actionFast(44, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 44 != 0, err
	}
	v[95].Add(&v[40], &v[93])
	code, err = p.actionFast(45, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 45 != 0, err
	}
	v[96] = chainPure(OpCode(17), v[27], v[91])
	v[97] = chainPure(OpCode(21), v[96], v[34])
	code, err = p.actionFast(46, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 46 != 0, err
	}
	v[98].Sub(&v[91], &v[27])
	code, err = p.actionFast(47, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 47 != 0, err
	}
	code, err = p.actionFast(48, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 48 != 0, err
	}
	code, err = p.actionFast(49, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 49 != 0, err
	}
	code, err = p.actionFast(50, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 50 != 0, err
	}
	v[99] = chainPure(OpCode(21), v[98], v[34])
	v[100] = chainPure(OpCode(21), v[99], v[34])
	v[101] = p.nodes[101].c
	code, err = p.actionFast(51, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 51 != 0, err
	}
	v[102] = p.nodes[102].c
	code, err = p.actionFast(52, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 52 != 0, err
	}
	code, err = p.actionFast(53, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 53 != 0, err
	}
	code, err = p.actionFast(54, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 54 != 0, err
	}
	code, err = p.actionFast(55, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 55 != 0, err
	}
	code, err = p.actionFast(56, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 56 != 0, err
	}
	code, err = p.actionFast(57, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 57 != 0, err
	}
	code, err = p.actionFast(58, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 58 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["5200699f40fd45b12b0325a3ae3064303cd1e2a59bf30aa0d5a6a325a5a95a61"] = chainFastShape45
}

func chainFastShape46(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33].SetBytes(effectiveCaller(s))
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	v[41].And(&v[30], &v[40])
	v[42] = p.nodes[42].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[43] = p.nodes[43].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[45] = p.nodes[45].c
	v[46] = p.nodes[46].c
	v[47] = p.nodes[47].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[49] = p.nodes[49].c
	v[50] = p.nodes[50].c
	v[51] = p.nodes[51].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[52] = p.nodes[52].c
	v[53] = p.nodes[53].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[54] = p.nodes[54].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[57] = p.nodes[57].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[59] = p.nodes[59].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[63] = p.nodes[63].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[64] = p.nodes[64].c
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	v[68] = p.nodes[68].c
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	v[70] = p.nodes[70].c
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	v[73] = p.nodes[73].c
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	v[74] = p.nodes[74].c
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	v[77] = p.nodes[77].c
	code, err = p.actionFast(44, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 44 != 0, err
	}
	code, err = p.actionFast(45, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 45 != 0, err
	}
	code, err = p.actionFast(46, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 46 != 0, err
	}
	code, err = p.actionFast(47, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 47 != 0, err
	}
	code, err = p.actionFast(48, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 48 != 0, err
	}
	v[79] = p.nodes[79].c
	code, err = p.actionFast(49, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 49 != 0, err
	}
	v[81].Add(&v[53], &v[80])
	code, err = p.actionFast(50, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 50 != 0, err
	}
	code, err = p.actionFast(51, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 51 != 0, err
	}
	code, err = p.actionFast(52, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 52 != 0, err
	}
	v[83].Add(&v[53], &v[82])
	code, err = p.actionFast(53, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 53 != 0, err
	}
	code, err = p.actionFast(54, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 54 != 0, err
	}
	v[84] = p.nodes[84].c
	v[85] = p.nodes[85].c
	v[86] = p.nodes[86].c
	v[87] = p.nodes[87].c
	v[88] = p.nodes[88].c
	v[89] = p.nodes[89].c
	v[90] = p.nodes[90].c
	code, err = p.actionFast(55, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 55 != 0, err
	}
	v[91] = p.nodes[91].c
	code, err = p.actionFast(56, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 56 != 0, err
	}
	code, err = p.actionFast(57, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 57 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["539bfc40fa3989bb39a19463b41695a360e9b86f1dcf569ce23a0a55d88e121e"] = chainFastShape46
}

func chainFastShape47(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = chainPure(OpCode(21), v[30], v[33])
	v[35] = p.nodes[35].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[36] = chainPure(OpCode(17), v[29], v[33])
	v[37] = chainPure(OpCode(21), v[36], v[33])
	v[38] = p.nodes[38].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	v[41] = p.nodes[41].c
	v[42].And(&v[26], &v[41])
	v[43].And(&v[41], &v[42])
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[44] = p.nodes[44].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[45] = p.nodes[45].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[48] = p.nodes[48].c
	v[49] = p.nodes[49].c
	v[50] = p.nodes[50].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[52].Add(&v[45], &v[51])
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[53] = p.nodes[53].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[54].Add(&v[44], &v[51])
	v[55] = p.nodes[55].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[56] = p.nodes[56].c
	v[57] = chainPure(OpCode(17), v[28], v[47])
	v[58] = chainPure(OpCode(21), v[57], v[33])
	v[59] = p.nodes[59].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[60] = p.nodes[60].c
	v[61] = p.nodes[61].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[63] = p.nodes[63].c
	v[64] = p.nodes[64].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[66].Add(&v[45], &v[65])
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[67].Add(&v[44], &v[65])
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[68] = chainPure(OpCode(17), v[29], v[28])
	v[69] = chainPure(OpCode(21), v[68], v[33])
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[71].And(&v[27], &v[41])
	v[72].And(&v[41], &v[71])
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[75] = p.nodes[75].c
	v[76] = p.nodes[76].c
	v[77] = p.nodes[77].c
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	v[79] = chainPure(OpCode(16), v[78], v[74])
	v[80] = chainPure(OpCode(21), v[79], v[33])
	v[81] = p.nodes[81].c
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	v[82] = p.nodes[82].c
	v[83] = p.nodes[83].c
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	v[85].Add(&v[45], &v[84])
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	v[86].Add(&v[44], &v[84])
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	v[89] = p.nodes[89].c
	v[90] = p.nodes[90].c
	v[91].Add(&v[44], &v[88])
	v[92] = p.nodes[92].c
	v[93].Add(&v[33], &v[88])
	v[94] = p.nodes[94].c
	v[95] = p.nodes[95].c
	v[96] = p.nodes[96].c
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	v[98].Sub(&v[91], &v[97])
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["56e3965fab082d6866d0277b1d3ed644aafa5b1877207a4509ab20af06195825"] = chainFastShape47
}

func chainFastShape48(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38].And(&v[28], &v[37])
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[41] = p.nodes[41].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[45] = p.nodes[45].c
	v[46] = p.nodes[46].c
	v[47] = p.nodes[47].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[49] = p.nodes[49].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[52] = p.nodes[52].c
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[53] = p.nodes[53].c
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[56] = p.nodes[56].c
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["5b8271a96ff61e6dbcc1f65d478071defbd85e4f70bd9c41bd7b63e90508fe52"] = chainFastShape48
}

func chainFastShape49(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38].And(&v[26], &v[37])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[39] = p.nodes[39].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[40] = p.nodes[40].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[46] = chainPure(OpCode(17), v[45], v[30])
	v[47] = chainPure(OpCode(21), v[46], v[35])
	v[48] = p.nodes[48].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[49] = p.nodes[49].c
	v[50] = p.nodes[50].c
	v[51] = p.nodes[51].c
	v[52] = p.nodes[52].c
	v[53].And(&v[28], &v[52])
	v[54] = chainPure(OpCode(21), v[53], v[35])
	v[55] = p.nodes[55].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[56] = chainPure(OpCode(21), v[29], v[35])
	v[57] = p.nodes[57].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[58] = p.nodes[58].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[60] = p.nodes[60].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[61].Add(&v[39], &v[59])
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[63].Sub(&v[61], &v[62])
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[64] = p.nodes[64].c
	v[65] = p.nodes[65].c
	v[66].And(&v[25], &v[37])
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[69] = p.nodes[69].c
	v[70] = p.nodes[70].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[72] = p.nodes[72].c
	v[73] = p.nodes[73].c
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["5f531a320e341cdfd402537853f047530f50e0758774a258f2eb5ccdb60df70f"] = chainFastShape49
}

func chainFastShape50(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37].And(&v[29], &v[36])
	v[38].And(&v[36], &v[37])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[39] = p.nodes[39].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[40] = p.nodes[40].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[43] = chainPure(OpCode(16), v[42], v[31])
	v[44] = chainPure(OpCode(21), v[43], v[34])
	v[45] = p.nodes[45].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[46].SetBytes(effectiveCaller(s))
	v[47].And(&v[36], &v[46])
	v[48] = chainPure(OpCode(20), v[37], v[47])
	v[49] = chainPure(OpCode(21), v[48], v[34])
	v[50] = chainPure(OpCode(21), v[49], v[34])
	v[51] = p.nodes[51].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[52] = p.nodes[52].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[53] = p.nodes[53].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[55] = p.nodes[55].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[57] = p.nodes[57].c
	v[58].And(&v[30], &v[36])
	v[59].And(&v[36], &v[58])
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[62] = p.nodes[62].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[64] = p.nodes[64].c
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[66].Add(&v[39], &v[65])
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	v[68].Sub(&v[66], &v[67])
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["64522be6f9ff9ce300b58ad715e027f13cda9e42d6ce6e0a54c14137a7f4003a"] = chainFastShape50
}

func chainFastShape51(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38].And(&v[28], &v[37])
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[41] = p.nodes[41].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[43] = p.nodes[43].c
	v[44].And(&v[29], &v[43])
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[47] = p.nodes[47].c
	v[48].Add(&v[46], &v[47])
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[50] = p.nodes[50].c
	v[51].And(&v[49], &v[50])
	v[52] = chainPure(OpCode(21), v[51], v[35])
	v[53] = chainPure(OpCode(21), v[52], v[35])
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[54].Add(&v[33], &v[45])
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[56].Add(&v[40], &v[46])
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[57] = p.nodes[57].c
	v[58].Add(&v[45], &v[57])
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[60].Add(&v[41], &v[46])
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[61] = p.nodes[61].c
	v[62].Add(&v[45], &v[61])
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[64] = p.nodes[64].c
	v[65].Add(&v[46], &v[64])
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[66] = p.nodes[66].c
	v[67].Add(&v[45], &v[66])
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[69] = p.nodes[69].c
	v[70].Add(&v[46], &v[69])
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[71] = p.nodes[71].c
	v[72].Add(&v[45], &v[71])
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[74].Add(&v[34], &v[46])
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[75] = p.nodes[75].c
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[77] = chainPure(OpCode(17), v[30], v[76])
	v[78] = chainPure(OpCode(21), v[77], v[35])
	v[79] = p.nodes[79].c
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[80] = chainPure(OpCode(17), v[30], v[35])
	v[81] = chainPure(OpCode(21), v[80], v[35])
	v[82] = p.nodes[82].c
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[83] = p.nodes[83].c
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	v[86].Add(&v[40], &v[85])
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	v[87].Add(&v[41], &v[85])
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	v[88].Add(&v[64], &v[85])
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	v[89] = p.nodes[89].c
	v[90].Add(&v[69], &v[85])
	v[91] = p.nodes[91].c
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	v[93].Sub(&v[90], &v[92])
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["66dd6271899f67f9aa3406b7df613caecb575561022c60961f81d00b339a364f"] = chainFastShape51
}

func chainFastShape52(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37].Div(&v[31], &v[36])
	v[38].And(&v[30], &v[37])
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[42] = p.nodes[42].c
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	v[45] = p.nodes[45].c
	v[46].And(&v[26], &v[45])
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[47] = p.nodes[47].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[48] = p.nodes[48].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[51] = p.nodes[51].c
	v[52] = p.nodes[52].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[54] = chainPure(OpCode(17), v[53], v[41])
	v[55] = chainPure(OpCode(21), v[54], v[35])
	v[56] = p.nodes[56].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[57] = p.nodes[57].c
	v[58] = p.nodes[58].c
	v[59] = p.nodes[59].c
	v[60] = p.nodes[60].c
	v[61].And(&v[28], &v[60])
	v[62] = chainPure(OpCode(21), v[61], v[35])
	v[63] = p.nodes[63].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[64] = chainPure(OpCode(21), v[29], v[35])
	v[65] = p.nodes[65].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[66] = p.nodes[66].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[68] = p.nodes[68].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[69].Add(&v[47], &v[67])
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[71].Sub(&v[69], &v[70])
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[72] = p.nodes[72].c
	v[73] = p.nodes[73].c
	v[74].And(&v[25], &v[45])
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[77] = p.nodes[77].c
	v[78] = p.nodes[78].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[80] = p.nodes[80].c
	v[81] = p.nodes[81].c
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["677b911688490aef3cab46571cc695164466dc74e2f359adc4fa6acd32b904ed"] = chainFastShape52
}

func chainFastShape53(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33].SetBytes(effectiveCaller(s))
	v[34] = p.nodes[34].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[37] = p.nodes[37].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[40] = p.nodes[40].c
	v[41] = p.nodes[41].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	v[45] = p.nodes[45].c
	v[46].And(&v[28], &v[45])
	v[47] = p.nodes[47].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[48] = p.nodes[48].c
	v[49] = p.nodes[49].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[51] = p.nodes[51].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[53] = p.nodes[53].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[56] = p.nodes[56].c
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[59] = p.nodes[59].c
	v[60] = p.nodes[60].c
	v[61] = p.nodes[61].c
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[62].Add(&v[36], &v[58])
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	v[64].Sub(&v[62], &v[63])
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["69758e82226bc7ea5ab75ff4a24b83925727b5ca2eca3a3ed8f6a32436839b71"] = chainFastShape53
}

func chainFastShape54(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = chainPure(OpCode(21), v[30], v[33])
	v[35] = p.nodes[35].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[36] = p.nodes[36].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[37] = p.nodes[37].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[38] = p.nodes[38].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	v[41] = p.nodes[41].c
	v[42] = p.nodes[42].c
	v[43] = p.nodes[43].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[44] = chainPure(OpCode(21), v[29], v[33])
	v[45] = p.nodes[45].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[46] = p.nodes[46].c
	v[47] = p.nodes[47].c
	v[48] = p.nodes[48].c
	v[49].And(&v[26], &v[48])
	v[50] = p.nodes[50].c
	v[51] = p.nodes[51].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[53] = p.nodes[53].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[54] = p.nodes[54].c
	v[55].Add(&v[52], &v[54])
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[57].Sub(&v[55], &v[56])
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[58] = p.nodes[58].c
	v[59] = p.nodes[59].c
	v[60].And(&v[25], &v[48])
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[63] = p.nodes[63].c
	v[64] = p.nodes[64].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[66] = p.nodes[66].c
	v[67] = p.nodes[67].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[70] = p.nodes[70].c
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["6bc1b082628fc6a2f3f3696c2768261ab68d402d5f67435b5b739a4ae97b7101"] = chainFastShape54
}

func chainFastShape55(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33].SetBytes(effectiveCaller(s))
	v[34] = p.nodes[34].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[37] = p.nodes[37].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[40] = p.nodes[40].c
	v[41] = p.nodes[41].c
	v[42].And(&v[39], &v[41])
	v[43] = p.nodes[43].c
	v[44].Or(&v[42], &v[43])
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[45] = p.nodes[45].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[48] = p.nodes[48].c
	v[49] = p.nodes[49].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[51] = p.nodes[51].c
	v[52] = p.nodes[52].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[53] = p.nodes[53].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[55] = p.nodes[55].c
	v[56] = p.nodes[56].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[58] = chainPure(OpCode(20), v[47], v[57])
	v[59] = p.nodes[59].c
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[61].And(&v[41], &v[60])
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["6f080b72f4dc45ed13c4e51c512925bbd20311fd4902c6a9d48c3bab5e5f4e86"] = chainFastShape55
}

func chainFastShape56(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37].And(&v[30], &v[35])
	v[38] = chainPure(OpCode(21), v[37], v[36])
	v[39] = p.nodes[39].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[40] = chainPure(OpCode(21), v[31], v[36])
	v[41] = p.nodes[41].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[42] = p.nodes[42].c
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	v[45] = p.nodes[45].c
	v[46].And(&v[28], &v[45])
	v[47] = p.nodes[47].c
	v[48] = p.nodes[48].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[50] = p.nodes[50].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[51] = p.nodes[51].c
	v[52].Add(&v[49], &v[51])
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[54].Sub(&v[52], &v[53])
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[55] = p.nodes[55].c
	v[56] = p.nodes[56].c
	v[57].And(&v[27], &v[45])
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[60] = p.nodes[60].c
	v[61] = p.nodes[61].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[63] = p.nodes[63].c
	v[64] = p.nodes[64].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[67] = p.nodes[67].c
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["6ff465ec21c3b6950831173c9d95d8ff71b1e1baeb60327461afebdcf58a4241"] = chainFastShape56
}

func chainFastShape57(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = chainPure(OpCode(17), v[30], v[33])
	v[36] = chainPure(OpCode(21), v[35], v[34])
	v[37] = p.nodes[37].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[38] = p.nodes[38].c
	v[39] = chainPure(OpCode(16), v[30], v[38])
	v[40] = chainPure(OpCode(21), v[39], v[34])
	v[41] = p.nodes[41].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[42].SetBytes(effectiveCaller(s))
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[45] = p.nodes[45].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[48] = p.nodes[48].c
	v[49].And(&v[47], &v[48])
	v[50] = chainPure(OpCode(21), v[49], v[34])
	v[51] = p.nodes[51].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[53] = p.nodes[53].c
	v[54].And(&v[52], &v[53])
	v[55] = p.nodes[55].c
	v[56].Or(&v[54], &v[55])
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[57] = p.nodes[57].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[60] = p.nodes[60].c
	v[61] = p.nodes[61].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[63] = p.nodes[63].c
	v[64] = p.nodes[64].c
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[65] = p.nodes[65].c
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[67] = p.nodes[67].c
	v[68] = p.nodes[68].c
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[70] = chainPure(OpCode(20), v[59], v[69])
	v[71] = p.nodes[71].c
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	v[73].And(&v[53], &v[72])
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["733a38e328b94f2f344f4feaff458db350aa6e38242131f64c2e787b95b56dfe"] = chainFastShape57
}

func chainFastShape58(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33].SetBytes(effectiveCaller(s))
	v[34] = p.nodes[34].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[37] = p.nodes[37].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[40] = p.nodes[40].c
	v[41].And(&v[39], &v[40])
	v[42] = chainPure(OpCode(21), v[41], v[34])
	v[43] = p.nodes[43].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[45] = p.nodes[45].c
	v[46].And(&v[44], &v[45])
	v[47] = p.nodes[47].c
	v[48].Or(&v[46], &v[47])
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[49] = p.nodes[49].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[52] = p.nodes[52].c
	v[53] = p.nodes[53].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[55] = p.nodes[55].c
	v[56] = p.nodes[56].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[57] = p.nodes[57].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[59] = p.nodes[59].c
	v[60] = p.nodes[60].c
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[62] = chainPure(OpCode(20), v[51], v[61])
	v[63] = p.nodes[63].c
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	v[65].And(&v[45], &v[64])
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["73bed98be973c765e83de15c459689d97c3f640d2866b0cfbcd9718452b3ae57"] = chainFastShape58
}

func chainFastShape59(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	v[39].SetBytes(effectiveCaller(s))
	v[40] = p.nodes[40].c
	v[41].And(&v[39], &v[40])
	v[42].And(&v[40], &v[41])
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[43] = p.nodes[43].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[44] = p.nodes[44].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[48] = p.nodes[48].c
	v[49].And(&v[27], &v[40])
	v[50].And(&v[40], &v[49])
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[53] = p.nodes[53].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[55] = chainPure(OpCode(17), v[29], v[35])
	v[56] = chainPure(OpCode(21), v[55], v[35])
	v[57] = p.nodes[57].c
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[58] = p.nodes[58].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[60].Add(&v[43], &v[59])
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[62].Sub(&v[60], &v[61])
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["7583774cfbf262f707c555d841d0ad29271a2def3bf183b3d033b514a917200c"] = chainFastShape59
}

func chainFastShape60(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = chainPure(OpCode(21), v[30], v[33])
	v[35] = p.nodes[35].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[36] = p.nodes[36].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[37] = p.nodes[37].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	v[41] = p.nodes[41].c
	v[42] = p.nodes[42].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[43] = chainPure(OpCode(21), v[29], v[33])
	v[44] = p.nodes[44].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[45] = p.nodes[45].c
	v[46] = p.nodes[46].c
	v[47] = p.nodes[47].c
	v[48].And(&v[26], &v[47])
	v[49] = p.nodes[49].c
	v[50] = p.nodes[50].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[52] = p.nodes[52].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[53] = p.nodes[53].c
	v[54].Add(&v[51], &v[53])
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[56].Sub(&v[54], &v[55])
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[57] = p.nodes[57].c
	v[58] = p.nodes[58].c
	v[59].And(&v[25], &v[47])
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[62] = p.nodes[62].c
	v[63] = p.nodes[63].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[65] = p.nodes[65].c
	v[66] = p.nodes[66].c
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[69] = p.nodes[69].c
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["784aa87f7f3a6d1f4920628b0018001e30456886916a96e1d0dcef268a32e3bb"] = chainFastShape60
}

func chainFastShape61(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38].And(&v[26], &v[37])
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[41] = p.nodes[41].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[44] = p.nodes[44].c
	v[45] = p.nodes[45].c
	v[46].Div(&v[43], &v[45])
	v[47] = p.nodes[47].c
	v[48].And(&v[46], &v[47])
	v[49] = chainPure(OpCode(21), v[48], v[35])
	v[50] = p.nodes[50].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[51] = p.nodes[51].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[52].And(&v[27], &v[37])
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[55].Div(&v[54], &v[45])
	v[56].And(&v[47], &v[55])
	v[57] = chainPure(OpCode(21), v[56], v[35])
	v[58] = p.nodes[58].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[59] = p.nodes[59].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[60] = p.nodes[60].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[62] = p.nodes[62].c
	v[63] = p.nodes[63].c
	v[64] = p.nodes[64].c
	v[65].Div(&v[61], &v[64])
	v[66] = p.nodes[66].c
	v[67].And(&v[65], &v[66])
	v[68] = p.nodes[68].c
	v[69] = p.nodes[69].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[71] = p.nodes[71].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[74] = p.nodes[74].c
	v[75] = p.nodes[75].c
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[77] = chainPure(OpCode(17), v[76], v[70])
	v[78] = chainPure(OpCode(21), v[77], v[35])
	v[79] = p.nodes[79].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[80] = p.nodes[80].c
	v[81] = p.nodes[81].c
	v[82] = p.nodes[82].c
	v[83] = p.nodes[83].c
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[84] = chainPure(OpCode(21), v[30], v[35])
	v[85] = p.nodes[85].c
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[86] = p.nodes[86].c
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[88] = p.nodes[88].c
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[89].Add(&v[40], &v[87])
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	v[91].Sub(&v[89], &v[90])
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[92] = p.nodes[92].c
	v[93] = p.nodes[93].c
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	v[96] = p.nodes[96].c
	v[97] = p.nodes[97].c
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	v[99] = p.nodes[99].c
	v[100] = p.nodes[100].c
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	code, err = p.actionFast(44, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 44 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["794ac3a9b48befadc4efcf84a58aa34734d58611073ba1db087ecb785b5ef13d"] = chainFastShape61
}

func chainFastShape62(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37].And(&v[31], &v[35])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[40] = p.nodes[40].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[43] = p.nodes[43].c
	v[44].And(&v[42], &v[43])
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[46] = p.nodes[46].c
	v[47] = chainPure(OpCode(17), v[45], v[43])
	v[48] = chainPure(OpCode(21), v[47], v[36])
	v[49] = p.nodes[49].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[50] = p.nodes[50].c
	v[51] = p.nodes[51].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[53] = p.nodes[53].c
	v[54] = chainPure(OpCode(28), v[53], v[52])
	v[55] = p.nodes[55].c
	v[56] = chainPure(OpCode(20), v[54], v[55])
	v[57] = chainPure(OpCode(21), v[56], v[36])
	v[58] = p.nodes[58].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["84a2ed0e4af6e00203874f1cf9fdeef8faf9536d36d1d119f9eb15cabb30cf31"] = chainFastShape62
}

func chainFastShape63(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = chainPure(OpCode(21), v[30], v[33])
	v[35] = p.nodes[35].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[36] = p.nodes[36].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	v[41] = p.nodes[41].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[42] = chainPure(OpCode(21), v[29], v[33])
	v[43] = p.nodes[43].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[44] = p.nodes[44].c
	v[45] = p.nodes[45].c
	v[46] = p.nodes[46].c
	v[47].And(&v[26], &v[46])
	v[48] = p.nodes[48].c
	v[49] = p.nodes[49].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[51] = p.nodes[51].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[52] = p.nodes[52].c
	v[53].Add(&v[50], &v[52])
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[55].Sub(&v[53], &v[54])
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[56] = p.nodes[56].c
	v[57] = p.nodes[57].c
	v[58].And(&v[25], &v[46])
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[61] = p.nodes[61].c
	v[62] = p.nodes[62].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[64] = p.nodes[64].c
	v[65] = p.nodes[65].c
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[68] = p.nodes[68].c
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["850d7892e4bb69219c58bad9b97ba7338845f34741bf614a2e0d34953df875e5"] = chainFastShape63
}

func chainFastShape64(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	v[41].And(&v[28], &v[40])
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[42] = p.nodes[42].c
	v[43] = p.nodes[43].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[44] = p.nodes[44].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[46] = p.nodes[46].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[50] = p.nodes[50].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[51] = p.nodes[51].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[55] = p.nodes[55].c
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[57] = p.nodes[57].c
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	v[60] = p.nodes[60].c
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	v[61] = p.nodes[61].c
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	v[64] = p.nodes[64].c
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["8513039208c9a65169806f4b7e3366dc503260ce622e285186dd8175db3ffe04"] = chainFastShape64
}

func chainFastShape65(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = chainPure(OpCode(17), v[30], v[33])
	v[36] = chainPure(OpCode(21), v[35], v[34])
	v[37] = p.nodes[37].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	v[41] = p.nodes[41].c
	v[42].And(&v[28], &v[41])
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[45] = p.nodes[45].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[48] = p.nodes[48].c
	v[49] = p.nodes[49].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[52] = p.nodes[52].c
	v[53] = p.nodes[53].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[55].And(&v[29], &v[41])
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[58] = p.nodes[58].c
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[61] = p.nodes[61].c
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["87324b70b66ddbbabf5b420a03e8a8a64de673a70ce11fad71968ed55f3518e9"] = chainFastShape65
}

func chainFastShape66(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	v[40] = chainPure(OpCode(20), v[27], v[39])
	v[41] = chainPure(OpCode(21), v[40], v[39])
	v[42] = p.nodes[42].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[43].Mul(&v[27], &v[36])
	v[44] = chainPure(OpCode(21), v[27], v[39])
	v[45] = chainPure(OpCode(21), v[44], v[39])
	v[46] = p.nodes[46].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[47].Div(&v[43], &v[27])
	v[48] = chainPure(OpCode(20), v[36], v[47])
	v[49] = chainPure(OpCode(21), v[48], v[39])
	v[50] = chainPure(OpCode(21), v[49], v[39])
	v[51] = p.nodes[51].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[52] = p.nodes[52].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[54] = p.nodes[54].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[56] = chainPure(OpCode(17), v[53], v[55])
	v[57] = chainPure(OpCode(21), v[56], v[39])
	v[58] = p.nodes[58].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[59] = p.nodes[59].c
	v[60] = chainPure(OpCode(16), v[28], v[59])
	v[61] = chainPure(OpCode(21), v[60], v[39])
	v[62] = p.nodes[62].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[63] = p.nodes[63].c
	v[64] = p.nodes[64].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[66] = p.nodes[66].c
	v[67] = p.nodes[67].c
	v[68] = p.nodes[68].c
	v[69].And(&v[25], &v[68])
	v[70].And(&v[68], &v[69])
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[71] = p.nodes[71].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[72] = p.nodes[72].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[76] = p.nodes[76].c
	v[77].And(&v[26], &v[68])
	v[78].And(&v[68], &v[77])
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[81] = p.nodes[81].c
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[83] = chainPure(OpCode(17), v[53], v[39])
	v[84] = chainPure(OpCode(21), v[83], v[39])
	v[85] = p.nodes[85].c
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	v[86] = p.nodes[86].c
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	v[88].Add(&v[71], &v[87])
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	v[90].Sub(&v[88], &v[89])
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["87ac809792640cfb04118f661a82697fa292dd8eacd3aac772fc71a7c2128ec3"] = chainFastShape66
}

func chainFastShape67(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	v[40].And(&v[26], &v[39])
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[41] = p.nodes[41].c
	v[42] = p.nodes[42].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[43] = p.nodes[43].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[46] = p.nodes[46].c
	v[47] = p.nodes[47].c
	v[48] = p.nodes[48].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[50] = p.nodes[50].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[54] = p.nodes[54].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[55] = p.nodes[55].c
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[59] = p.nodes[59].c
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	v[61] = p.nodes[61].c
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	v[64] = p.nodes[64].c
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	v[65] = p.nodes[65].c
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	v[68] = p.nodes[68].c
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	code, err = p.actionFast(44, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 44 != 0, err
	}
	v[70] = p.nodes[70].c
	v[71] = p.nodes[71].c
	v[72] = p.nodes[72].c
	v[73] = p.nodes[73].c
	v[74] = p.nodes[74].c
	v[75] = p.nodes[75].c
	v[76] = p.nodes[76].c
	v[77] = p.nodes[77].c
	v[78].And(&v[31], &v[77])
	code, err = p.actionFast(45, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 45 != 0, err
	}
	v[79] = p.nodes[79].c
	code, err = p.actionFast(46, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 46 != 0, err
	}
	code, err = p.actionFast(47, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 47 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["8b394be3e531284e4957715c8c1bed4afc7e5061412470cbb2bf1360fec83aa9"] = chainFastShape67
}

func chainFastShape68(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	v[40].And(&v[26], &v[39])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[41] = p.nodes[41].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[42] = p.nodes[42].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[45] = p.nodes[45].c
	v[46] = p.nodes[46].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[48] = p.nodes[48].c
	v[49] = p.nodes[49].c
	v[50].And(&v[27], &v[39])
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[53] = p.nodes[53].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["8c006fa8bdc5b67b4df09a3461a9cc9c9dc504efe20733cb63ef85bd32ca12be"] = chainFastShape68
}

func chainFastShape69(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	v[39].And(&v[25], &v[38])
	v[40].And(&v[38], &v[39])
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[41] = p.nodes[41].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[42] = p.nodes[42].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[44].SetBytes(effectiveCaller(s))
	v[45].And(&v[38], &v[44])
	v[46].And(&v[38], &v[45])
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[48] = p.nodes[48].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[50] = p.nodes[50].c
	v[51] = p.nodes[51].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[55] = p.nodes[55].c
	v[56].And(&v[26], &v[38])
	v[57].And(&v[38], &v[56])
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[60] = p.nodes[60].c
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[62] = chainPure(OpCode(17), v[29], v[35])
	v[63] = chainPure(OpCode(21), v[62], v[35])
	v[64] = p.nodes[64].c
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	v[65] = p.nodes[65].c
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	v[67].Add(&v[41], &v[66])
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	v[69].Sub(&v[67], &v[68])
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["8eeee27c8f222dc289db4c746fc87da0fd8a15719b7c83cc82788de9a8ba78d6"] = chainFastShape69
}

func chainFastShape70(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[39].Add(&v[37], &v[38])
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[40] = p.nodes[40].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[41] = p.nodes[41].c
	v[42].Add(&v[38], &v[41])
	v[43] = p.nodes[43].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[44] = p.nodes[44].c
	v[45] = chainPure(OpCode(17), v[30], v[31])
	v[46] = chainPure(OpCode(21), v[45], v[35])
	v[47] = p.nodes[47].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[48].Sub(&v[31], &v[30])
	v[49] = p.nodes[49].c
	v[50] = p.nodes[50].c
	v[51] = p.nodes[51].c
	v[52] = p.nodes[52].c
	v[53].And(&v[28], &v[52])
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[54] = p.nodes[54].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[56] = p.nodes[56].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[60] = p.nodes[60].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[62].Add(&v[37], &v[61])
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[63].Add(&v[41], &v[61])
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[64] = chainPure(OpCode(17), v[30], v[59])
	v[65] = chainPure(OpCode(21), v[64], v[35])
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[66].Sub(&v[59], &v[30])
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[67] = p.nodes[67].c
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	v[70] = p.nodes[70].c
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	v[72].Add(&v[37], &v[71])
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	v[73].Add(&v[41], &v[71])
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	v[74] = chainPure(OpCode(17), v[30], v[69])
	v[75] = chainPure(OpCode(21), v[74], v[35])
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	v[76].Sub(&v[69], &v[30])
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	v[77] = chainPure(OpCode(21), v[76], v[35])
	v[78] = chainPure(OpCode(21), v[77], v[35])
	v[79] = p.nodes[79].c
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	v[80] = p.nodes[80].c
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	code, err = p.actionFast(44, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 44 != 0, err
	}
	code, err = p.actionFast(45, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 45 != 0, err
	}
	code, err = p.actionFast(46, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 46 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["900a44b8e6779f8cfceaf548f8d20179701d1241cb148d825a594bc651b62a26"] = chainFastShape70
}

func chainFastShape71(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = chainPure(OpCode(21), v[29], v[33])
	v[35] = p.nodes[35].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	v[40].And(&v[26], &v[39])
	v[41] = p.nodes[41].c
	v[42] = p.nodes[42].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[44] = p.nodes[44].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[45] = p.nodes[45].c
	v[46].Add(&v[43], &v[45])
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[48].Sub(&v[46], &v[47])
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[49] = p.nodes[49].c
	v[50] = p.nodes[50].c
	v[51].And(&v[25], &v[39])
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[54] = p.nodes[54].c
	v[55] = p.nodes[55].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[57] = p.nodes[57].c
	v[58] = p.nodes[58].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[61] = p.nodes[61].c
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["9103c70a062c07a32fd99fa0f1483d606d32e7ca3dd9ebd4595328a8335c86d8"] = chainFastShape71
}

func chainFastShape72(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38].And(&v[26], &v[37])
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[41] = p.nodes[41].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[43].SetBytes(effectiveCaller(s))
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[46] = chainPure(OpCode(17), v[28], v[45])
	v[47] = chainPure(OpCode(21), v[46], v[35])
	v[48] = p.nodes[48].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[50] = p.nodes[50].c
	v[51] = p.nodes[51].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["9426e7c8fd0919be05d0da1e97783f88fa7dccf6de22acbc84f54598560cff64"] = chainFastShape72
}

func chainFastShape73(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36].And(&v[31], &v[34])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[39] = p.nodes[39].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[42] = p.nodes[42].c
	v[43] = chainPure(OpCode(28), v[42], v[41])
	v[44] = p.nodes[44].c
	v[45] = chainPure(OpCode(20), v[43], v[44])
	v[46] = chainPure(OpCode(21), v[45], v[35])
	v[47] = p.nodes[47].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[48] = p.nodes[48].c
	v[49].SetBytes(effectiveCaller(s))
	v[50] = p.nodes[50].c
	v[51].And(&v[34], &v[49])
	v[52] = p.nodes[52].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[53].And(&v[27], &v[34])
	v[54] = p.nodes[54].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[55] = p.nodes[55].c
	v[56] = p.nodes[56].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[59] = p.nodes[59].c
	v[60].And(&v[58], &v[59])
	v[61] = chainPure(OpCode(17), v[28], v[60])
	v[62] = chainPure(OpCode(21), v[61], v[35])
	v[63] = p.nodes[63].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[64] = p.nodes[64].c
	v[65] = p.nodes[65].c
	v[66] = p.nodes[66].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[68].And(&v[59], &v[67])
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[70] = p.nodes[70].c
	v[71] = chainPure(OpCode(17), v[69], v[59])
	v[72] = chainPure(OpCode(21), v[71], v[35])
	v[73] = p.nodes[73].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[74] = p.nodes[74].c
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[76] = chainPure(OpCode(28), v[42], v[75])
	v[77] = chainPure(OpCode(20), v[44], v[76])
	v[78] = chainPure(OpCode(21), v[77], v[35])
	v[79] = p.nodes[79].c
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[80] = p.nodes[80].c
	v[81] = p.nodes[81].c
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	v[84].And(&v[59], &v[83])
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	v[86] = chainPure(OpCode(17), v[85], v[59])
	v[87] = chainPure(OpCode(21), v[86], v[35])
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	v[89] = chainPure(OpCode(28), v[42], v[88])
	v[90] = chainPure(OpCode(20), v[44], v[89])
	v[91] = chainPure(OpCode(21), v[90], v[35])
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	v[92] = p.nodes[92].c
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	v[94].Add(&v[38], &v[93])
	code, err = p.actionFast(44, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 44 != 0, err
	}
	v[96].Sub(&v[94], &v[95])
	code, err = p.actionFast(45, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 45 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["94289a09facc368f5816c8443461925fdf098dde4a2b27ae93b540aea857d74b"] = chainFastShape73
}

func chainFastShape74(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33].SetBytes(effectiveCaller(s))
	v[34] = p.nodes[34].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[37] = p.nodes[37].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[40] = p.nodes[40].c
	v[41] = p.nodes[41].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[43] = p.nodes[43].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[46] = chainPure(OpCode(17), v[42], v[45])
	v[47] = chainPure(OpCode(21), v[46], v[34])
	v[48] = p.nodes[48].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[49] = p.nodes[49].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[52] = p.nodes[52].c
	v[53].And(&v[51], &v[52])
	v[54] = chainPure(OpCode(21), v[53], v[34])
	v[55] = p.nodes[55].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[57] = p.nodes[57].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[59] = p.nodes[59].c
	v[60] = p.nodes[60].c
	v[61] = p.nodes[61].c
	v[62].And(&v[29], &v[61])
	v[63] = p.nodes[63].c
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	v[64] = p.nodes[64].c
	v[65] = p.nodes[65].c
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[67] = p.nodes[67].c
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	v[69] = p.nodes[69].c
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	v[72] = p.nodes[72].c
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	v[75] = p.nodes[75].c
	v[76] = p.nodes[76].c
	v[77] = p.nodes[77].c
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	v[78].Add(&v[36], &v[74])
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	v[80].Sub(&v[78], &v[79])
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["9698ffeab00e12233c225f3feb030df8cfb4f0b0e9fed184307b43401afa3c97"] = chainFastShape74
}

func chainFastShape75(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = chainPure(OpCode(17), v[28], v[33])
	v[35] = chainPure(OpCode(21), v[34], v[33])
	v[36] = p.nodes[36].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	v[40].And(&v[25], &v[39])
	v[41].And(&v[39], &v[40])
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[42] = p.nodes[42].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[43] = p.nodes[43].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[46] = p.nodes[46].c
	v[47] = p.nodes[47].c
	v[48] = p.nodes[48].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[50].Add(&v[43], &v[49])
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[51] = p.nodes[51].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[52].Add(&v[42], &v[49])
	v[53] = p.nodes[53].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[54] = p.nodes[54].c
	v[55] = chainPure(OpCode(17), v[27], v[45])
	v[56] = chainPure(OpCode(21), v[55], v[33])
	v[57] = p.nodes[57].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[58] = p.nodes[58].c
	v[59] = p.nodes[59].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[61] = p.nodes[61].c
	v[62] = p.nodes[62].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[64].Add(&v[43], &v[63])
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[65].Add(&v[42], &v[63])
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[66] = chainPure(OpCode(17), v[28], v[27])
	v[67] = chainPure(OpCode(21), v[66], v[33])
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[69].And(&v[26], &v[39])
	v[70].And(&v[39], &v[69])
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	v[73] = p.nodes[73].c
	v[74] = p.nodes[74].c
	v[75] = p.nodes[75].c
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[77] = chainPure(OpCode(16), v[76], v[72])
	v[78] = chainPure(OpCode(21), v[77], v[33])
	v[79] = p.nodes[79].c
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	v[80] = p.nodes[80].c
	v[81] = p.nodes[81].c
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	v[83].Add(&v[43], &v[82])
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	v[84].Add(&v[42], &v[82])
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	v[87] = p.nodes[87].c
	v[88] = p.nodes[88].c
	v[89].Add(&v[42], &v[86])
	v[90] = p.nodes[90].c
	v[91].Add(&v[33], &v[86])
	v[92] = p.nodes[92].c
	v[93] = p.nodes[93].c
	v[94] = p.nodes[94].c
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	v[96].Sub(&v[89], &v[95])
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["9a2b912b19ac631072e468dbb8a62a9d3a189e5f6f7d1778047599a1c8cff50a"] = chainFastShape75
}

func chainFastShape76(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	v[40].And(&v[27], &v[39])
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[41] = p.nodes[41].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[42] = p.nodes[42].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[45] = p.nodes[45].c
	v[46] = p.nodes[46].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[48].And(&v[28], &v[39])
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[51] = p.nodes[51].c
	v[52] = p.nodes[52].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[56] = p.nodes[56].c
	v[57].Sub(&v[54], &v[55])
	v[58].Add(&v[41], &v[57])
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["9a3a45ff0f5c2dbed552feea9542b68f2f69dd7f581b473bde07b5abc0ccd173"] = chainFastShape76
}

func chainFastShape77(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36].And(&v[25], &v[34])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[39] = p.nodes[39].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[43] = p.nodes[43].c
	v[44].And(&v[42], &v[43])
	v[45] = p.nodes[45].c
	v[46].Or(&v[44], &v[45])
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[48] = p.nodes[48].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[49] = p.nodes[49].c
	v[50] = p.nodes[50].c
	v[51] = p.nodes[51].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[52].And(&v[26], &v[34])
	v[53] = p.nodes[53].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[54] = p.nodes[54].c
	v[55] = p.nodes[55].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[56] = p.nodes[56].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[59] = p.nodes[59].c
	v[60].And(&v[58], &v[59])
	v[61] = chainPure(OpCode(17), v[27], v[60])
	v[62] = chainPure(OpCode(21), v[61], v[35])
	v[63] = p.nodes[63].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[64] = p.nodes[64].c
	v[65] = p.nodes[65].c
	v[66] = p.nodes[66].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[68].And(&v[59], &v[67])
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[70] = p.nodes[70].c
	v[71] = chainPure(OpCode(17), v[69], v[59])
	v[72] = chainPure(OpCode(21), v[71], v[35])
	v[73] = p.nodes[73].c
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[74] = p.nodes[74].c
	v[75] = p.nodes[75].c
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[77] = p.nodes[77].c
	v[78] = chainPure(OpCode(28), v[77], v[76])
	v[79] = chainPure(OpCode(20), v[45], v[78])
	v[80] = chainPure(OpCode(21), v[79], v[35])
	v[81] = p.nodes[81].c
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	v[82] = p.nodes[82].c
	v[83] = p.nodes[83].c
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	v[86].And(&v[59], &v[85])
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	v[88] = chainPure(OpCode(17), v[87], v[59])
	v[89] = chainPure(OpCode(21), v[88], v[35])
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	v[91] = chainPure(OpCode(28), v[77], v[90])
	v[92] = chainPure(OpCode(20), v[45], v[91])
	v[93] = chainPure(OpCode(21), v[92], v[35])
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	code, err = p.actionFast(44, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 44 != 0, err
	}
	code, err = p.actionFast(45, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 45 != 0, err
	}
	code, err = p.actionFast(46, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 46 != 0, err
	}
	v[94] = p.nodes[94].c
	code, err = p.actionFast(47, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 47 != 0, err
	}
	code, err = p.actionFast(48, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 48 != 0, err
	}
	v[96].Add(&v[38], &v[95])
	code, err = p.actionFast(49, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 49 != 0, err
	}
	v[98].Sub(&v[96], &v[97])
	code, err = p.actionFast(50, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 50 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["9a4a333816ec7415ac295f6c05c807ab3d3acab87f4330e5fda830b0bcf39c18"] = chainFastShape77
}

func chainFastShape78(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	v[40].And(&v[30], &v[39])
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[41] = p.nodes[41].c
	v[42] = p.nodes[42].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[43] = p.nodes[43].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[46] = p.nodes[46].c
	v[47] = p.nodes[47].c
	v[48] = p.nodes[48].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[50] = p.nodes[50].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[54] = p.nodes[54].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[55] = p.nodes[55].c
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[59] = p.nodes[59].c
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	v[61] = p.nodes[61].c
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	v[64] = p.nodes[64].c
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	v[65] = p.nodes[65].c
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	v[68] = p.nodes[68].c
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	code, err = p.actionFast(44, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 44 != 0, err
	}
	v[70] = p.nodes[70].c
	code, err = p.actionFast(45, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 45 != 0, err
	}
	v[72].Add(&v[42], &v[71])
	code, err = p.actionFast(46, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 46 != 0, err
	}
	code, err = p.actionFast(47, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 47 != 0, err
	}
	code, err = p.actionFast(48, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 48 != 0, err
	}
	v[74].Add(&v[42], &v[73])
	code, err = p.actionFast(49, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 49 != 0, err
	}
	code, err = p.actionFast(50, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 50 != 0, err
	}
	v[75] = p.nodes[75].c
	v[76] = p.nodes[76].c
	v[77] = p.nodes[77].c
	v[78] = p.nodes[78].c
	v[79] = p.nodes[79].c
	v[80] = p.nodes[80].c
	v[81] = p.nodes[81].c
	code, err = p.actionFast(51, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 51 != 0, err
	}
	v[82] = p.nodes[82].c
	code, err = p.actionFast(52, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 52 != 0, err
	}
	code, err = p.actionFast(53, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 53 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["9bd09b0513c333bd9198c97ea1e3705584cd3fa52ea6806f294b5b05640bc72a"] = chainFastShape78
}

func chainFastShape79(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[37] = p.nodes[37].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[41] = p.nodes[41].c
	v[42] = p.nodes[42].c
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	v[45].And(&v[29], &v[44])
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[46] = p.nodes[46].c
	v[47] = p.nodes[47].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[48] = p.nodes[48].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[51] = p.nodes[51].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[53] = p.nodes[53].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[57] = p.nodes[57].c
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[58] = p.nodes[58].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	v[62] = p.nodes[62].c
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	v[64] = p.nodes[64].c
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	v[67] = p.nodes[67].c
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	v[68] = p.nodes[68].c
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	v[71] = p.nodes[71].c
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	code, err = p.actionFast(44, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 44 != 0, err
	}
	code, err = p.actionFast(45, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 45 != 0, err
	}
	code, err = p.actionFast(46, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 46 != 0, err
	}
	v[73] = p.nodes[73].c
	code, err = p.actionFast(47, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 47 != 0, err
	}
	v[75].Add(&v[47], &v[74])
	code, err = p.actionFast(48, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 48 != 0, err
	}
	code, err = p.actionFast(49, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 49 != 0, err
	}
	code, err = p.actionFast(50, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 50 != 0, err
	}
	v[77].Add(&v[47], &v[76])
	code, err = p.actionFast(51, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 51 != 0, err
	}
	code, err = p.actionFast(52, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 52 != 0, err
	}
	v[78] = p.nodes[78].c
	v[79] = p.nodes[79].c
	v[80] = p.nodes[80].c
	v[81] = p.nodes[81].c
	v[82] = p.nodes[82].c
	v[83] = p.nodes[83].c
	v[84] = p.nodes[84].c
	code, err = p.actionFast(53, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 53 != 0, err
	}
	v[85] = p.nodes[85].c
	code, err = p.actionFast(54, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 54 != 0, err
	}
	code, err = p.actionFast(55, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 55 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["9cabe4ed682c3074dd1f94146d06af1494857d7ac5e662cada1fd34be88b76c2"] = chainFastShape79
}

func chainFastShape80(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	v[40].And(&v[25], &v[39])
	v[41].And(&v[39], &v[40])
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[42] = p.nodes[42].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[43] = p.nodes[43].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[47] = p.nodes[47].c
	v[48].And(&v[26], &v[39])
	v[49].And(&v[39], &v[48])
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[52] = p.nodes[52].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[54] = chainPure(OpCode(17), v[29], v[35])
	v[55] = chainPure(OpCode(21), v[54], v[35])
	v[56] = p.nodes[56].c
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[57] = p.nodes[57].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[59].Add(&v[42], &v[58])
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[61].Sub(&v[59], &v[60])
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["9da4b298c5d5e0f4fdd8536db24ceb0f9accaac410ad4f7f770807277d58da73"] = chainFastShape80
}

func chainFastShape81(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	v[39].And(&v[28], &v[38])
	v[40] = p.nodes[40].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[41].And(&v[29], &v[38])
	v[42] = p.nodes[42].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[43] = chainPure(OpCode(17), v[30], v[36])
	v[44] = p.nodes[44].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[45] = p.nodes[45].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[47].And(&v[38], &v[46])
	v[48].And(&v[38], &v[47])
	v[49] = chainPure(OpCode(20), v[39], v[48])
	v[50] = chainPure(OpCode(21), v[49], v[36])
	v[51] = chainPure(OpCode(21), v[50], v[36])
	v[52] = p.nodes[52].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[53] = p.nodes[53].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[54] = p.nodes[54].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[55] = p.nodes[55].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[56] = p.nodes[56].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[57] = p.nodes[57].c
	v[58] = p.nodes[58].c
	v[59] = p.nodes[59].c
	v[60] = p.nodes[60].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[61] = p.nodes[61].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[62] = p.nodes[62].c
	v[63] = p.nodes[63].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[65] = p.nodes[65].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[66] = p.nodes[66].c
	v[67].Add(&v[64], &v[66])
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[69].Sub(&v[67], &v[68])
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[70] = p.nodes[70].c
	v[71] = p.nodes[71].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[74] = p.nodes[74].c
	v[75] = p.nodes[75].c
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[77] = p.nodes[77].c
	v[78] = p.nodes[78].c
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	v[81] = p.nodes[81].c
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["a73fd38b328d1b0173acdcf0cee7b13e02ecbac2ba7a58384b7e3a98c7ed9972"] = chainFastShape81
}

func chainFastShape82(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[41] = p.nodes[41].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[42] = p.nodes[42].c
	v[43] = p.nodes[43].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[44] = p.nodes[44].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[47] = p.nodes[47].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[51] = p.nodes[51].c
	v[52].Sub(&v[49], &v[50])
	v[53].Add(&v[43], &v[52])
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[54] = p.nodes[54].c
	v[55] = chainPure(OpCode(17), v[23], v[33])
	v[56] = p.nodes[56].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[57] = p.nodes[57].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[60].And(&v[22], &v[41])
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	v[66].Sub(&v[64], &v[65])
	v[67].Add(&v[43], &v[66])
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["a7b2ff621f5ac3ae0eac9093db4c6bc983634094cc7c7e986e7cb130cf347c21"] = chainFastShape82
}

func chainFastShape83(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37].And(&v[26], &v[35])
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[40] = p.nodes[40].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[43] = p.nodes[43].c
	v[44] = chainPure(OpCode(28), v[43], v[42])
	v[45] = p.nodes[45].c
	v[46] = chainPure(OpCode(20), v[44], v[45])
	v[47] = chainPure(OpCode(21), v[46], v[36])
	v[48] = p.nodes[48].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[49] = p.nodes[49].c
	v[50].And(&v[27], &v[35])
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[53] = chainPure(OpCode(28), v[43], v[52])
	v[54] = chainPure(OpCode(20), v[45], v[53])
	v[55] = chainPure(OpCode(21), v[54], v[36])
	v[56] = p.nodes[56].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[57] = p.nodes[57].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[59].SetBytes(effectiveCaller(s))
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[62] = chainPure(OpCode(17), v[28], v[61])
	v[63] = chainPure(OpCode(21), v[62], v[36])
	v[64] = p.nodes[64].c
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[65] = p.nodes[65].c
	v[66] = p.nodes[66].c
	v[67] = p.nodes[67].c
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[68] = p.nodes[68].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[69] = p.nodes[69].c
	v[70] = p.nodes[70].c
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[72] = p.nodes[72].c
	v[73].And(&v[71], &v[72])
	v[74] = chainPure(OpCode(17), v[28], v[73])
	v[75] = chainPure(OpCode(21), v[74], v[36])
	v[76] = p.nodes[76].c
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	v[77] = p.nodes[77].c
	v[78] = p.nodes[78].c
	v[79] = p.nodes[79].c
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	v[81].And(&v[72], &v[80])
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	v[83] = p.nodes[83].c
	v[84] = chainPure(OpCode(17), v[82], v[72])
	v[85] = chainPure(OpCode(21), v[84], v[36])
	v[86] = p.nodes[86].c
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	v[87] = p.nodes[87].c
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	v[89] = chainPure(OpCode(28), v[43], v[88])
	v[90] = chainPure(OpCode(20), v[45], v[89])
	v[91] = chainPure(OpCode(21), v[90], v[36])
	v[92] = p.nodes[92].c
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	v[93] = p.nodes[93].c
	v[94] = p.nodes[94].c
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	code, err = p.actionFast(44, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 44 != 0, err
	}
	v[96].And(&v[72], &v[95])
	code, err = p.actionFast(45, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 45 != 0, err
	}
	v[98] = chainPure(OpCode(17), v[97], v[72])
	v[99] = chainPure(OpCode(21), v[98], v[36])
	code, err = p.actionFast(46, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 46 != 0, err
	}
	code, err = p.actionFast(47, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 47 != 0, err
	}
	code, err = p.actionFast(48, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 48 != 0, err
	}
	code, err = p.actionFast(49, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 49 != 0, err
	}
	code, err = p.actionFast(50, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 50 != 0, err
	}
	v[101] = chainPure(OpCode(28), v[43], v[100])
	v[102] = chainPure(OpCode(20), v[45], v[101])
	v[103] = chainPure(OpCode(21), v[102], v[36])
	code, err = p.actionFast(51, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 51 != 0, err
	}
	code, err = p.actionFast(52, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 52 != 0, err
	}
	code, err = p.actionFast(53, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 53 != 0, err
	}
	code, err = p.actionFast(54, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 54 != 0, err
	}
	code, err = p.actionFast(55, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 55 != 0, err
	}
	v[104] = p.nodes[104].c
	code, err = p.actionFast(56, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 56 != 0, err
	}
	code, err = p.actionFast(57, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 57 != 0, err
	}
	v[106].Add(&v[39], &v[105])
	code, err = p.actionFast(58, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 58 != 0, err
	}
	v[108].Sub(&v[106], &v[107])
	code, err = p.actionFast(59, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 59 != 0, err
	}
	code, err = p.actionFast(60, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 60 != 0, err
	}
	code, err = p.actionFast(61, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 61 != 0, err
	}
	code, err = p.actionFast(62, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 62 != 0, err
	}
	code, err = p.actionFast(63, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 63 != 0, err
	}
	code, err = p.actionFast(64, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 64 != 0, err
	}
	code, err = p.actionFast(65, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 65 != 0, err
	}
	code, err = p.actionFast(66, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 66 != 0, err
	}
	v[110] = p.nodes[110].c
	v[111] = p.nodes[111].c
	code, err = p.actionFast(67, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 67 != 0, err
	}
	code, err = p.actionFast(68, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 68 != 0, err
	}
	code, err = p.actionFast(69, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 69 != 0, err
	}
	code, err = p.actionFast(70, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 70 != 0, err
	}
	code, err = p.actionFast(71, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 71 != 0, err
	}
	code, err = p.actionFast(72, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 72 != 0, err
	}
	code, err = p.actionFast(73, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 73 != 0, err
	}
	code, err = p.actionFast(74, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 74 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["b49d0b0112b0065fc6c9b7e4443dabd11363b3066af7985f34d10d56b07aca9e"] = chainFastShape83
}

func chainFastShape84(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36].And(&v[31], &v[34])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[39] = p.nodes[39].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[42] = p.nodes[42].c
	v[43] = chainPure(OpCode(28), v[42], v[41])
	v[44] = p.nodes[44].c
	v[45] = chainPure(OpCode(20), v[43], v[44])
	v[46] = chainPure(OpCode(21), v[45], v[35])
	v[47] = p.nodes[47].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[48] = p.nodes[48].c
	v[49].And(&v[28], &v[34])
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[52] = chainPure(OpCode(28), v[42], v[51])
	v[53] = chainPure(OpCode(20), v[44], v[52])
	v[54] = chainPure(OpCode(21), v[53], v[35])
	v[55] = p.nodes[55].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[56] = p.nodes[56].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[57] = chainPure(OpCode(17), v[29], v[35])
	v[58] = p.nodes[58].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[59].SetBytes(effectiveCaller(s))
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[60] = p.nodes[60].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[63] = chainPure(OpCode(17), v[29], v[62])
	v[64] = chainPure(OpCode(21), v[63], v[35])
	v[65] = p.nodes[65].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[66] = p.nodes[66].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[68] = p.nodes[68].c
	v[69] = p.nodes[69].c
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[71] = p.nodes[71].c
	v[72] = p.nodes[72].c
	v[73] = p.nodes[73].c
	v[74] = p.nodes[74].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[76] = p.nodes[76].c
	v[77].And(&v[75], &v[76])
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[79] = p.nodes[79].c
	v[80] = chainPure(OpCode(17), v[78], v[76])
	v[81] = chainPure(OpCode(21), v[80], v[35])
	v[82] = p.nodes[82].c
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	v[83] = p.nodes[83].c
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	v[85] = chainPure(OpCode(28), v[42], v[84])
	v[86] = chainPure(OpCode(20), v[44], v[85])
	v[87] = chainPure(OpCode(21), v[86], v[35])
	v[88] = p.nodes[88].c
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	v[89] = p.nodes[89].c
	v[90] = p.nodes[90].c
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	v[94] = p.nodes[94].c
	v[95].Sub(&v[92], &v[93])
	v[96].Add(&v[38], &v[95])
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	code, err = p.actionFast(44, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 44 != 0, err
	}
	code, err = p.actionFast(45, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 45 != 0, err
	}
	code, err = p.actionFast(46, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 46 != 0, err
	}
	v[99] = p.nodes[99].c
	v[100].Sub(&v[97], &v[98])
	v[101].Add(&v[38], &v[100])
	code, err = p.actionFast(47, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 47 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["b56453a8bf5a2bd0ce7815e306f4bec8f4724482def642304ea7b4483ee52d07"] = chainFastShape84
}

func chainFastShape85(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36].And(&v[28], &v[34])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[39] = p.nodes[39].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[42] = p.nodes[42].c
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[46].And(&v[29], &v[34])
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[49] = p.nodes[49].c
	v[50] = p.nodes[50].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[54] = p.nodes[54].c
	v[55].Sub(&v[52], &v[53])
	v[56].Add(&v[38], &v[55])
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["b5f91f6adf138c70ee551c0803e2ca33caaee5a75663015ac7681ebf70aabea8"] = chainFastShape85
}

func chainFastShape86(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36].And(&v[31], &v[34])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[39] = p.nodes[39].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[42] = p.nodes[42].c
	v[43] = chainPure(OpCode(28), v[42], v[41])
	v[44] = p.nodes[44].c
	v[45] = chainPure(OpCode(20), v[43], v[44])
	v[46] = chainPure(OpCode(21), v[45], v[35])
	v[47] = p.nodes[47].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[48] = p.nodes[48].c
	v[49].And(&v[28], &v[34])
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[52] = chainPure(OpCode(28), v[42], v[51])
	v[53] = chainPure(OpCode(20), v[44], v[52])
	v[54] = chainPure(OpCode(21), v[53], v[35])
	v[55] = p.nodes[55].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[56] = p.nodes[56].c
	v[57].SetBytes(effectiveCaller(s))
	v[58] = p.nodes[58].c
	v[59].And(&v[34], &v[57])
	v[60] = p.nodes[60].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[61] = p.nodes[61].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[62] = p.nodes[62].c
	v[63] = p.nodes[63].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[66] = p.nodes[66].c
	v[67].And(&v[65], &v[66])
	v[68] = chainPure(OpCode(17), v[29], v[67])
	v[69] = chainPure(OpCode(21), v[68], v[35])
	v[70] = p.nodes[70].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[71] = p.nodes[71].c
	v[72] = p.nodes[72].c
	v[73] = p.nodes[73].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[75].And(&v[66], &v[74])
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[77] = p.nodes[77].c
	v[78] = chainPure(OpCode(17), v[76], v[66])
	v[79] = chainPure(OpCode(21), v[78], v[35])
	v[80] = p.nodes[80].c
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[81] = p.nodes[81].c
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[83] = chainPure(OpCode(28), v[42], v[82])
	v[84] = chainPure(OpCode(20), v[44], v[83])
	v[85] = chainPure(OpCode(21), v[84], v[35])
	v[86] = p.nodes[86].c
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	v[87] = p.nodes[87].c
	v[88] = p.nodes[88].c
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	v[90].And(&v[66], &v[89])
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	v[92] = chainPure(OpCode(17), v[91], v[66])
	v[93] = chainPure(OpCode(21), v[92], v[35])
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	v[95] = chainPure(OpCode(28), v[42], v[94])
	v[96] = chainPure(OpCode(20), v[44], v[95])
	v[97] = chainPure(OpCode(21), v[96], v[35])
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	code, err = p.actionFast(44, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 44 != 0, err
	}
	code, err = p.actionFast(45, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 45 != 0, err
	}
	code, err = p.actionFast(46, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 46 != 0, err
	}
	v[98] = p.nodes[98].c
	code, err = p.actionFast(47, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 47 != 0, err
	}
	code, err = p.actionFast(48, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 48 != 0, err
	}
	v[100].Add(&v[38], &v[99])
	code, err = p.actionFast(49, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 49 != 0, err
	}
	v[102].Sub(&v[100], &v[101])
	code, err = p.actionFast(50, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 50 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["b66002954169309bb6f3d070e2b10afa1f35eb55792a64d184fa420cddbc0964"] = chainFastShape86
}

func chainFastShape87(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[35] = p.nodes[35].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	v[41] = p.nodes[41].c
	v[42].And(&v[25], &v[40])
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[45] = p.nodes[45].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[48] = p.nodes[48].c
	v[49].And(&v[47], &v[48])
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[51] = p.nodes[51].c
	v[52] = chainPure(OpCode(17), v[50], v[48])
	v[53] = chainPure(OpCode(21), v[52], v[41])
	v[54] = p.nodes[54].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[55] = p.nodes[55].c
	v[56] = p.nodes[56].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[58] = p.nodes[58].c
	v[59] = chainPure(OpCode(28), v[58], v[57])
	v[60] = p.nodes[60].c
	v[61] = chainPure(OpCode(20), v[59], v[60])
	v[62] = chainPure(OpCode(21), v[61], v[41])
	v[63] = p.nodes[63].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[64] = p.nodes[64].c
	v[65] = p.nodes[65].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[67].SetBytes(effectiveCaller(s))
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[68] = p.nodes[68].c
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[72] = p.nodes[72].c
	v[73].Sub(&v[70], &v[71])
	v[74].Add(&v[44], &v[73])
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	v[77] = p.nodes[77].c
	v[78].Sub(&v[75], &v[76])
	v[79].Add(&v[44], &v[78])
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["b770a3be1223a83fa7d0f3c2c81a755a94d151aa4bae725bfa948795ade3849a"] = chainFastShape87
}

func chainFastShape88(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38].And(&v[28], &v[37])
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[41] = p.nodes[41].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[44] = p.nodes[44].c
	v[45] = p.nodes[45].c
	v[46].And(&v[43], &v[45])
	v[47].Or(&v[33], &v[46])
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[48].SetBytes(effectiveCaller(s))
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[49] = p.nodes[49].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[52].And(&v[44], &v[51])
	v[53] = p.nodes[53].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[54] = p.nodes[54].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[57] = chainPure(OpCode(17), v[29], v[56])
	v[58] = chainPure(OpCode(21), v[57], v[35])
	v[59] = p.nodes[59].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[61] = p.nodes[61].c
	v[62] = p.nodes[62].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	v[65].And(&v[45], &v[64])
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["b7f438de10f091d744faeecdefde946012ae8aa560ba7f0a3c04f06eb9fe5ca3"] = chainFastShape88
}

func chainFastShape89(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38].And(&v[26], &v[36])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[41] = p.nodes[41].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[44] = p.nodes[44].c
	v[45].And(&v[43], &v[44])
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[47] = p.nodes[47].c
	v[48] = chainPure(OpCode(17), v[46], v[44])
	v[49] = chainPure(OpCode(21), v[48], v[37])
	v[50] = p.nodes[50].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[51] = p.nodes[51].c
	v[52] = p.nodes[52].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[54] = p.nodes[54].c
	v[55] = chainPure(OpCode(28), v[54], v[53])
	v[56] = p.nodes[56].c
	v[57] = chainPure(OpCode(20), v[55], v[56])
	v[58] = chainPure(OpCode(21), v[57], v[37])
	v[59] = p.nodes[59].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[60] = p.nodes[60].c
	v[61] = p.nodes[61].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[63].SetBytes(effectiveCaller(s))
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[64] = p.nodes[64].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[68] = p.nodes[68].c
	v[69].Sub(&v[66], &v[67])
	v[70].Add(&v[40], &v[69])
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[73] = p.nodes[73].c
	v[74].Sub(&v[71], &v[72])
	v[75].Add(&v[40], &v[74])
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["b962ff136fa756358d02dfa01db3a63464b5c637b88786c1a0aa50d780f6db0f"] = chainFastShape89
}

func chainFastShape90(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35].And(&v[29], &v[33])
	v[36] = p.nodes[36].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[37].SetBytes(effectiveCaller(s))
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[40] = p.nodes[40].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[42] = p.nodes[42].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[44] = chainPure(OpCode(21), v[30], v[34])
	v[45] = p.nodes[45].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[48] = p.nodes[48].c
	v[49].And(&v[47], &v[48])
	v[50] = p.nodes[50].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[52] = p.nodes[52].c
	v[53].And(&v[51], &v[52])
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["bc4e6eaa2e7158ef4f7e9ac314f2a6b6fe730740361eb3c6fbd0f75123d23c73"] = chainFastShape90
}

func chainFastShape91(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33].SetBytes(effectiveCaller(s))
	v[34] = p.nodes[34].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[37] = p.nodes[37].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[40] = chainPure(OpCode(17), v[27], v[39])
	v[41] = chainPure(OpCode(21), v[40], v[34])
	v[42] = p.nodes[42].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[43] = p.nodes[43].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[45] = p.nodes[45].c
	v[46] = p.nodes[46].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[48] = p.nodes[48].c
	v[49] = p.nodes[49].c
	v[50] = p.nodes[50].c
	v[51] = p.nodes[51].c
	v[52] = p.nodes[52].c
	v[53].And(&v[26], &v[52])
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[54] = p.nodes[54].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[57] = p.nodes[57].c
	v[58].And(&v[56], &v[57])
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[60] = p.nodes[60].c
	v[61] = chainPure(OpCode(17), v[59], v[57])
	v[62] = chainPure(OpCode(21), v[61], v[34])
	v[63] = p.nodes[63].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[64] = p.nodes[64].c
	v[65] = p.nodes[65].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[67] = p.nodes[67].c
	v[68] = chainPure(OpCode(28), v[67], v[66])
	v[69] = p.nodes[69].c
	v[70] = chainPure(OpCode(20), v[68], v[69])
	v[71] = chainPure(OpCode(21), v[70], v[34])
	v[72] = p.nodes[72].c
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[73] = p.nodes[73].c
	v[74] = p.nodes[74].c
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	v[78] = p.nodes[78].c
	v[79].Sub(&v[76], &v[77])
	v[80].Add(&v[36], &v[79])
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	v[83] = p.nodes[83].c
	v[84].Sub(&v[81], &v[82])
	v[85].Add(&v[36], &v[84])
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["be4edef877a943ea638dc1931b308c496037334aea57f3cadb5200feeddc5112"] = chainFastShape91
}

func chainFastShape92(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37].And(&v[27], &v[36])
	v[38].And(&v[36], &v[37])
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[39] = p.nodes[39].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[40] = p.nodes[40].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[42].SetBytes(effectiveCaller(s))
	v[43].And(&v[36], &v[42])
	v[44].And(&v[36], &v[43])
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[47] = p.nodes[47].c
	v[48] = p.nodes[48].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[50] = p.nodes[50].c
	v[51] = p.nodes[51].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[55] = p.nodes[55].c
	v[56].And(&v[28], &v[36])
	v[57].And(&v[36], &v[56])
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	v[60] = p.nodes[60].c
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	v[62] = p.nodes[62].c
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	v[64].Add(&v[39], &v[63])
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	v[66].Sub(&v[64], &v[65])
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["bf6f1355181332843d0425526e9c5c4ef8f23bbbb158a0bb9459a5e1d7980bfb"] = chainFastShape92
}

func chainFastShape93(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = chainPure(OpCode(16), v[27], v[33])
	v[36] = chainPure(OpCode(21), v[35], v[34])
	v[37] = p.nodes[37].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[42] = p.nodes[42].c
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	v[45].And(&v[24], &v[44])
	v[46].And(&v[44], &v[45])
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[47] = p.nodes[47].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[48] = p.nodes[48].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[52] = p.nodes[52].c
	v[53].And(&v[25], &v[44])
	v[54].And(&v[44], &v[53])
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[57] = p.nodes[57].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[59] = chainPure(OpCode(17), v[28], v[34])
	v[60] = chainPure(OpCode(21), v[59], v[34])
	v[61] = p.nodes[61].c
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[62] = p.nodes[62].c
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[64].Add(&v[47], &v[63])
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[66].Sub(&v[64], &v[65])
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["bf92c4032f750ad67c94e9474158e4b246713594ece19d0319e87838b805c86d"] = chainFastShape93
}

func chainFastShape94(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[36] = p.nodes[36].c
	v[37].SetBytes(effectiveCaller(s))
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	v[41] = p.nodes[41].c
	v[42] = p.nodes[42].c
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	v[45].And(&v[27], &v[44])
	v[46] = p.nodes[46].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[47] = p.nodes[47].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[49] = p.nodes[49].c
	v[50] = p.nodes[50].c
	v[51] = p.nodes[51].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[53] = p.nodes[53].c
	v[54] = p.nodes[54].c
	v[55] = p.nodes[55].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[56] = p.nodes[56].c
	v[57] = p.nodes[57].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[58] = p.nodes[58].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[61] = p.nodes[61].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[63] = p.nodes[63].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[67] = p.nodes[67].c
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[68] = p.nodes[68].c
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	v[72] = p.nodes[72].c
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	v[74] = p.nodes[74].c
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	v[77] = p.nodes[77].c
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	v[78] = p.nodes[78].c
	code, err = p.actionFast(44, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 44 != 0, err
	}
	code, err = p.actionFast(45, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 45 != 0, err
	}
	code, err = p.actionFast(46, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 46 != 0, err
	}
	v[81] = p.nodes[81].c
	code, err = p.actionFast(47, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 47 != 0, err
	}
	code, err = p.actionFast(48, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 48 != 0, err
	}
	code, err = p.actionFast(49, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 49 != 0, err
	}
	code, err = p.actionFast(50, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 50 != 0, err
	}
	code, err = p.actionFast(51, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 51 != 0, err
	}
	v[83] = p.nodes[83].c
	code, err = p.actionFast(52, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 52 != 0, err
	}
	v[85].Add(&v[57], &v[84])
	code, err = p.actionFast(53, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 53 != 0, err
	}
	code, err = p.actionFast(54, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 54 != 0, err
	}
	code, err = p.actionFast(55, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 55 != 0, err
	}
	v[87].Add(&v[57], &v[86])
	code, err = p.actionFast(56, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 56 != 0, err
	}
	code, err = p.actionFast(57, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 57 != 0, err
	}
	v[88] = p.nodes[88].c
	v[89] = p.nodes[89].c
	v[90] = p.nodes[90].c
	v[91] = p.nodes[91].c
	v[92] = p.nodes[92].c
	v[93] = p.nodes[93].c
	v[94] = p.nodes[94].c
	code, err = p.actionFast(58, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 58 != 0, err
	}
	v[95] = p.nodes[95].c
	code, err = p.actionFast(59, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 59 != 0, err
	}
	code, err = p.actionFast(60, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 60 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["c23a6530fade87a852dd45d17f23648c1cebe1600b3d0df9c0457a55481c6e62"] = chainFastShape94
}

func chainFastShape95(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35].And(&v[24], &v[33])
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[38] = p.nodes[38].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[40].SetBytes(effectiveCaller(s))
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[43] = chainPure(OpCode(17), v[26], v[42])
	v[44] = chainPure(OpCode(21), v[43], v[34])
	v[45] = p.nodes[45].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[46] = p.nodes[46].c
	v[47] = p.nodes[47].c
	v[48] = p.nodes[48].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[49].And(&v[25], &v[33])
	v[50] = p.nodes[50].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[51] = p.nodes[51].c
	v[52] = p.nodes[52].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[53] = p.nodes[53].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[56] = p.nodes[56].c
	v[57].And(&v[55], &v[56])
	v[58] = chainPure(OpCode(17), v[26], v[57])
	v[59] = chainPure(OpCode(21), v[58], v[34])
	v[60] = p.nodes[60].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[61] = p.nodes[61].c
	v[62] = p.nodes[62].c
	v[63] = p.nodes[63].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[65].And(&v[56], &v[64])
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[67] = p.nodes[67].c
	v[68] = chainPure(OpCode(17), v[66], v[56])
	v[69] = chainPure(OpCode(21), v[68], v[34])
	v[70] = p.nodes[70].c
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[71] = p.nodes[71].c
	v[72] = p.nodes[72].c
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	v[74] = p.nodes[74].c
	v[75] = chainPure(OpCode(28), v[74], v[73])
	v[76] = p.nodes[76].c
	v[77] = chainPure(OpCode(20), v[75], v[76])
	v[78] = chainPure(OpCode(21), v[77], v[34])
	v[79] = p.nodes[79].c
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	v[80] = p.nodes[80].c
	v[81] = p.nodes[81].c
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	v[84].And(&v[56], &v[83])
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	v[86] = chainPure(OpCode(17), v[85], v[56])
	v[87] = chainPure(OpCode(21), v[86], v[34])
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	v[89] = chainPure(OpCode(28), v[74], v[88])
	v[90] = chainPure(OpCode(20), v[76], v[89])
	v[91] = chainPure(OpCode(21), v[90], v[34])
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	code, err = p.actionFast(44, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 44 != 0, err
	}
	code, err = p.actionFast(45, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 45 != 0, err
	}
	v[92] = p.nodes[92].c
	code, err = p.actionFast(46, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 46 != 0, err
	}
	code, err = p.actionFast(47, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 47 != 0, err
	}
	v[94].Add(&v[37], &v[93])
	code, err = p.actionFast(48, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 48 != 0, err
	}
	v[96].Sub(&v[94], &v[95])
	code, err = p.actionFast(49, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 49 != 0, err
	}
	code, err = p.actionFast(50, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 50 != 0, err
	}
	code, err = p.actionFast(51, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 51 != 0, err
	}
	code, err = p.actionFast(52, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 52 != 0, err
	}
	code, err = p.actionFast(53, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 53 != 0, err
	}
	code, err = p.actionFast(54, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 54 != 0, err
	}
	code, err = p.actionFast(55, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 55 != 0, err
	}
	code, err = p.actionFast(56, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 56 != 0, err
	}
	v[98] = p.nodes[98].c
	v[99] = p.nodes[99].c
	code, err = p.actionFast(57, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 57 != 0, err
	}
	code, err = p.actionFast(58, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 58 != 0, err
	}
	code, err = p.actionFast(59, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 59 != 0, err
	}
	code, err = p.actionFast(60, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 60 != 0, err
	}
	code, err = p.actionFast(61, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 61 != 0, err
	}
	code, err = p.actionFast(62, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 62 != 0, err
	}
	code, err = p.actionFast(63, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 63 != 0, err
	}
	code, err = p.actionFast(64, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 64 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["c4e7bbeaa970facd71b30bca740136e82ad79d8aaeaece9434fa5ff18f2182e0"] = chainFastShape95
}

func chainFastShape96(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33].SetBytes(effectiveCaller(s))
	v[34] = p.nodes[34].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[37] = p.nodes[37].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[40] = p.nodes[40].c
	v[41].And(&v[39], &v[40])
	v[42] = chainPure(OpCode(21), v[41], v[34])
	v[43] = p.nodes[43].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[44] = p.nodes[44].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[47] = p.nodes[47].c
	v[48] = p.nodes[48].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[50] = p.nodes[50].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[53] = chainPure(OpCode(17), v[49], v[52])
	v[54] = chainPure(OpCode(21), v[53], v[34])
	v[55] = p.nodes[55].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[56] = p.nodes[56].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[59].And(&v[40], &v[58])
	v[60] = chainPure(OpCode(21), v[59], v[34])
	v[61] = p.nodes[61].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[63] = p.nodes[63].c
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	v[65] = p.nodes[65].c
	v[66] = p.nodes[66].c
	v[67] = p.nodes[67].c
	v[68].And(&v[29], &v[67])
	v[69] = p.nodes[69].c
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	v[70] = p.nodes[70].c
	v[71] = p.nodes[71].c
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	v[73] = p.nodes[73].c
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	v[75] = p.nodes[75].c
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	v[78] = p.nodes[78].c
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	v[81] = p.nodes[81].c
	v[82] = p.nodes[82].c
	v[83] = p.nodes[83].c
	code, err = p.actionFast(44, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 44 != 0, err
	}
	v[84].Add(&v[36], &v[80])
	code, err = p.actionFast(45, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 45 != 0, err
	}
	v[86].Sub(&v[84], &v[85])
	code, err = p.actionFast(46, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 46 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["c76b47a1b51b977c983c06fe790121925e9e3c1307502cfeb6366c8758b2b782"] = chainFastShape96
}

func chainFastShape97(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38].And(&v[28], &v[37])
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[41] = p.nodes[41].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[43] = p.nodes[43].c
	v[44].And(&v[29], &v[43])
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[47] = p.nodes[47].c
	v[48].Add(&v[46], &v[47])
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[50] = p.nodes[50].c
	v[51].And(&v[49], &v[50])
	v[52] = chainPure(OpCode(21), v[51], v[35])
	v[53] = chainPure(OpCode(21), v[52], v[35])
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[54].Add(&v[33], &v[45])
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[56].Add(&v[40], &v[46])
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[57] = p.nodes[57].c
	v[58].Add(&v[45], &v[57])
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[60].Add(&v[41], &v[46])
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[61] = p.nodes[61].c
	v[62].Add(&v[45], &v[61])
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[64] = p.nodes[64].c
	v[65].Add(&v[46], &v[64])
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[66] = p.nodes[66].c
	v[67].Add(&v[45], &v[66])
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[69] = p.nodes[69].c
	v[70].Add(&v[46], &v[69])
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[71] = p.nodes[71].c
	v[72].Add(&v[45], &v[71])
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[74].Add(&v[34], &v[46])
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[75] = p.nodes[75].c
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[76] = chainPure(OpCode(17), v[30], v[35])
	v[77] = chainPure(OpCode(21), v[76], v[35])
	v[78] = p.nodes[78].c
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[79] = p.nodes[79].c
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	v[82].Add(&v[40], &v[81])
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	v[83].Add(&v[41], &v[81])
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	v[84].Add(&v[64], &v[81])
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	v[85] = p.nodes[85].c
	v[86].Add(&v[69], &v[81])
	v[87] = p.nodes[87].c
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	v[89].Sub(&v[86], &v[88])
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["c93e854ee174ade10176e8bb8b0f111e83503b6cd43511ca7c6680afd64b8281"] = chainFastShape97
}

func chainFastShape98(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	v[40].And(&v[27], &v[39])
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[41] = p.nodes[41].c
	v[42] = p.nodes[42].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[43] = p.nodes[43].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[45].SetBytes(effectiveCaller(s))
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[47] = p.nodes[47].c
	v[48] = p.nodes[48].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[51] = p.nodes[51].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[53].And(&v[28], &v[39])
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[56] = p.nodes[56].c
	v[57] = p.nodes[57].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	v[61] = p.nodes[61].c
	v[62].Sub(&v[59], &v[60])
	v[63].Add(&v[42], &v[62])
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["c956b2ee20d44d4c941f4138919b8139e1b064a29489fd46b1033981523102a6"] = chainFastShape98
}

func chainFastShape99(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36].And(&v[31], &v[34])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[39] = p.nodes[39].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[42] = p.nodes[42].c
	v[43] = chainPure(OpCode(28), v[42], v[41])
	v[44] = p.nodes[44].c
	v[45] = chainPure(OpCode(20), v[43], v[44])
	v[46] = chainPure(OpCode(21), v[45], v[35])
	v[47] = p.nodes[47].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[48] = p.nodes[48].c
	v[49].And(&v[27], &v[34])
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[52] = chainPure(OpCode(28), v[42], v[51])
	v[53] = chainPure(OpCode(20), v[44], v[52])
	v[54] = chainPure(OpCode(21), v[53], v[35])
	v[55] = p.nodes[55].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[56].And(&v[26], &v[34])
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[57] = p.nodes[57].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[59].SetBytes(effectiveCaller(s))
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[62] = chainPure(OpCode(17), v[28], v[61])
	v[63] = chainPure(OpCode(21), v[62], v[35])
	v[64] = p.nodes[64].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[65] = p.nodes[65].c
	v[66] = p.nodes[66].c
	v[67] = p.nodes[67].c
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[68] = p.nodes[68].c
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[69] = p.nodes[69].c
	v[70] = p.nodes[70].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[73] = p.nodes[73].c
	v[74].And(&v[72], &v[73])
	v[75] = chainPure(OpCode(17), v[28], v[74])
	v[76] = chainPure(OpCode(21), v[75], v[35])
	v[77] = p.nodes[77].c
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[78] = p.nodes[78].c
	v[79] = p.nodes[79].c
	v[80] = p.nodes[80].c
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	v[82].And(&v[73], &v[81])
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	v[84] = p.nodes[84].c
	v[85] = chainPure(OpCode(17), v[83], v[73])
	v[86] = chainPure(OpCode(21), v[85], v[35])
	v[87] = p.nodes[87].c
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	v[88] = p.nodes[88].c
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	v[90] = chainPure(OpCode(28), v[42], v[89])
	v[91] = chainPure(OpCode(20), v[44], v[90])
	v[92] = chainPure(OpCode(21), v[91], v[35])
	v[93] = p.nodes[93].c
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	v[94] = p.nodes[94].c
	v[95] = p.nodes[95].c
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	v[97].And(&v[73], &v[96])
	code, err = p.actionFast(44, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 44 != 0, err
	}
	v[99] = chainPure(OpCode(17), v[98], v[73])
	v[100] = chainPure(OpCode(21), v[99], v[35])
	code, err = p.actionFast(45, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 45 != 0, err
	}
	code, err = p.actionFast(46, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 46 != 0, err
	}
	code, err = p.actionFast(47, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 47 != 0, err
	}
	code, err = p.actionFast(48, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 48 != 0, err
	}
	code, err = p.actionFast(49, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 49 != 0, err
	}
	v[102] = chainPure(OpCode(28), v[42], v[101])
	v[103] = chainPure(OpCode(20), v[44], v[102])
	v[104] = chainPure(OpCode(21), v[103], v[35])
	code, err = p.actionFast(50, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 50 != 0, err
	}
	code, err = p.actionFast(51, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 51 != 0, err
	}
	code, err = p.actionFast(52, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 52 != 0, err
	}
	code, err = p.actionFast(53, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 53 != 0, err
	}
	code, err = p.actionFast(54, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 54 != 0, err
	}
	v[105] = p.nodes[105].c
	code, err = p.actionFast(55, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 55 != 0, err
	}
	code, err = p.actionFast(56, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 56 != 0, err
	}
	v[107].Add(&v[38], &v[106])
	code, err = p.actionFast(57, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 57 != 0, err
	}
	v[109].Sub(&v[107], &v[108])
	code, err = p.actionFast(58, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 58 != 0, err
	}
	code, err = p.actionFast(59, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 59 != 0, err
	}
	code, err = p.actionFast(60, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 60 != 0, err
	}
	code, err = p.actionFast(61, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 61 != 0, err
	}
	code, err = p.actionFast(62, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 62 != 0, err
	}
	code, err = p.actionFast(63, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 63 != 0, err
	}
	code, err = p.actionFast(64, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 64 != 0, err
	}
	code, err = p.actionFast(65, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 65 != 0, err
	}
	v[111] = p.nodes[111].c
	v[112] = p.nodes[112].c
	code, err = p.actionFast(66, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 66 != 0, err
	}
	code, err = p.actionFast(67, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 67 != 0, err
	}
	code, err = p.actionFast(68, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 68 != 0, err
	}
	code, err = p.actionFast(69, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 69 != 0, err
	}
	code, err = p.actionFast(70, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 70 != 0, err
	}
	code, err = p.actionFast(71, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 71 != 0, err
	}
	code, err = p.actionFast(72, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 72 != 0, err
	}
	code, err = p.actionFast(73, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 73 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["cb833e32eef660df1a4013189b191a3b90164f752be01d66d4e87bdb35e1e25f"] = chainFastShape99
}

func chainFastShape100(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	v[39].And(&v[28], &v[38])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[40] = p.nodes[40].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[41] = p.nodes[41].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[44] = p.nodes[44].c
	v[45] = p.nodes[45].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[47].And(&v[29], &v[38])
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[50] = p.nodes[50].c
	v[51] = p.nodes[51].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[55] = p.nodes[55].c
	v[56].Sub(&v[53], &v[54])
	v[57].Add(&v[40], &v[56])
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["cd0118c22a0112c252ac89f3c1041492be578a5cb720447b559445fe1fddcf9d"] = chainFastShape100
}

func chainFastShape101(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[40] = p.nodes[40].c
	v[41] = p.nodes[41].c
	v[42] = p.nodes[42].c
	v[43] = p.nodes[43].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[44] = p.nodes[44].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[45] = p.nodes[45].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[48] = p.nodes[48].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[52] = p.nodes[52].c
	v[53].Sub(&v[50], &v[51])
	v[54].Add(&v[44], &v[53])
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[55] = p.nodes[55].c
	v[56] = chainPure(OpCode(17), v[22], v[33])
	v[57] = p.nodes[57].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[58] = p.nodes[58].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[61].And(&v[21], &v[43])
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	v[67].Sub(&v[65], &v[66])
	v[68].Add(&v[44], &v[67])
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["cd0a7ed7b4460ca69a8524b0caeb1227ef8fe1693fc394f2595b5d0c0a038577"] = chainFastShape101
}

func chainFastShape102(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38].And(&v[30], &v[36])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[41] = p.nodes[41].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[44] = p.nodes[44].c
	v[45].And(&v[43], &v[44])
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[47] = p.nodes[47].c
	v[48] = chainPure(OpCode(17), v[46], v[44])
	v[49] = chainPure(OpCode(21), v[48], v[37])
	v[50] = p.nodes[50].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[51] = p.nodes[51].c
	v[52] = p.nodes[52].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[54] = p.nodes[54].c
	v[55] = chainPure(OpCode(28), v[54], v[53])
	v[56] = p.nodes[56].c
	v[57] = chainPure(OpCode(20), v[55], v[56])
	v[58] = chainPure(OpCode(21), v[57], v[37])
	v[59] = p.nodes[59].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[60].And(&v[29], &v[36])
	v[61] = p.nodes[61].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[63].Add(&v[40], &v[62])
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[65].Sub(&v[63], &v[64])
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["d0f3e87c56300ef21cf616ec04f2e30af07d5a257413cbf550118b1e638a3f70"] = chainFastShape102
}

func chainFastShape103(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	v[40] = chainPure(OpCode(20), v[28], v[39])
	v[41] = chainPure(OpCode(21), v[40], v[39])
	v[42] = p.nodes[42].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[43].Mul(&v[28], &v[36])
	v[44] = chainPure(OpCode(21), v[28], v[39])
	v[45] = chainPure(OpCode(21), v[44], v[39])
	v[46] = p.nodes[46].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[47].Div(&v[43], &v[28])
	v[48] = chainPure(OpCode(20), v[36], v[47])
	v[49] = chainPure(OpCode(21), v[48], v[39])
	v[50] = chainPure(OpCode(21), v[49], v[39])
	v[51] = p.nodes[51].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[52] = p.nodes[52].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[54] = p.nodes[54].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[56] = chainPure(OpCode(17), v[53], v[55])
	v[57] = chainPure(OpCode(21), v[56], v[39])
	v[58] = p.nodes[58].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[59] = p.nodes[59].c
	v[60] = p.nodes[60].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[62] = p.nodes[62].c
	v[63] = p.nodes[63].c
	v[64].SetBytes(effectiveCaller(s))
	v[65] = p.nodes[65].c
	v[66].And(&v[64], &v[65])
	v[67].And(&v[65], &v[66])
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[68] = p.nodes[68].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[69] = p.nodes[69].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[73] = p.nodes[73].c
	v[74].And(&v[27], &v[65])
	v[75].And(&v[65], &v[74])
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[78] = p.nodes[78].c
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	v[80] = chainPure(OpCode(17), v[53], v[39])
	v[81] = chainPure(OpCode(21), v[80], v[39])
	v[82] = p.nodes[82].c
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[83] = p.nodes[83].c
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	v[85].Add(&v[68], &v[84])
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	v[87].Sub(&v[85], &v[86])
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["d18415e2ac2ec8a0a024972e1b6a58ba2eef3bcc6b23fe7523f986d8625910ad"] = chainFastShape103
}

func chainFastShape104(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	v[40].And(&v[25], &v[39])
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[41] = p.nodes[41].c
	v[42] = p.nodes[42].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[43] = p.nodes[43].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[46] = p.nodes[46].c
	v[47] = p.nodes[47].c
	v[48] = p.nodes[48].c
	v[49] = p.nodes[49].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[51].Add(&v[43], &v[50])
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[52] = p.nodes[52].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[53].Add(&v[42], &v[50])
	v[54] = p.nodes[54].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[55] = p.nodes[55].c
	v[56] = chainPure(OpCode(17), v[27], v[45])
	v[57] = chainPure(OpCode(21), v[56], v[37])
	v[58] = p.nodes[58].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[59].Sub(&v[45], &v[27])
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[60] = p.nodes[60].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[64] = p.nodes[64].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[66].Add(&v[43], &v[65])
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[67].Add(&v[42], &v[65])
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[68] = chainPure(OpCode(17), v[27], v[63])
	v[69] = chainPure(OpCode(21), v[68], v[37])
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[70].Sub(&v[63], &v[27])
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	v[71] = p.nodes[71].c
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	v[74] = p.nodes[74].c
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	v[76].Add(&v[43], &v[75])
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	v[77].Add(&v[42], &v[75])
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	v[78] = chainPure(OpCode(17), v[27], v[73])
	v[79] = chainPure(OpCode(21), v[78], v[37])
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	v[80].Sub(&v[73], &v[27])
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	v[81] = chainPure(OpCode(21), v[80], v[37])
	v[82] = chainPure(OpCode(21), v[81], v[37])
	v[83] = p.nodes[83].c
	code, err = p.actionFast(44, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 44 != 0, err
	}
	v[84] = p.nodes[84].c
	code, err = p.actionFast(45, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 45 != 0, err
	}
	code, err = p.actionFast(46, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 46 != 0, err
	}
	code, err = p.actionFast(47, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 47 != 0, err
	}
	code, err = p.actionFast(48, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 48 != 0, err
	}
	code, err = p.actionFast(49, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 49 != 0, err
	}
	code, err = p.actionFast(50, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 50 != 0, err
	}
	code, err = p.actionFast(51, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 51 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["d2287db8f5e8d8f3fee39d17682ea31d0fa033c6504490822e847424669ddfea"] = chainFastShape104
}

func chainFastShape105(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[39].Mul(&v[28], &v[29])
	v[40] = p.nodes[40].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[41].Div(&v[39], &v[28])
	v[42] = chainPure(OpCode(20), v[29], v[41])
	v[43] = p.nodes[43].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[44] = p.nodes[44].c
	v[45] = p.nodes[45].c
	v[46] = chainPure(OpCode(17), v[39], v[45])
	v[47] = chainPure(OpCode(21), v[46], v[37])
	v[48] = p.nodes[48].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[49] = chainPure(OpCode(21), v[39], v[37])
	v[50] = p.nodes[50].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[52] = p.nodes[52].c
	v[53] = p.nodes[53].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[55] = p.nodes[55].c
	v[56] = p.nodes[56].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[58] = p.nodes[58].c
	v[59] = p.nodes[59].c
	v[60] = p.nodes[60].c
	v[61] = p.nodes[61].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[62] = p.nodes[62].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[63] = p.nodes[63].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[66] = p.nodes[66].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[70] = p.nodes[70].c
	v[71].Sub(&v[68], &v[69])
	v[72].Add(&v[62], &v[71])
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[73] = p.nodes[73].c
	v[74] = chainPure(OpCode(17), v[51], v[37])
	v[75] = p.nodes[75].c
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[76] = p.nodes[76].c
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	v[79].And(&v[22], &v[61])
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	v[85].Sub(&v[83], &v[84])
	v[86].Add(&v[62], &v[85])
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["d30ff375124732cf357d45a21029d3795133ca96750bfd5d891f3a971412a096"] = chainFastShape105
}

func chainFastShape106(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37].And(&v[27], &v[35])
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[40] = p.nodes[40].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[43] = p.nodes[43].c
	v[44] = chainPure(OpCode(28), v[43], v[42])
	v[45] = p.nodes[45].c
	v[46] = chainPure(OpCode(20), v[44], v[45])
	v[47] = chainPure(OpCode(21), v[46], v[36])
	v[48] = p.nodes[48].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[49] = p.nodes[49].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[50] = chainPure(OpCode(17), v[28], v[36])
	v[51] = p.nodes[51].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[52].SetBytes(effectiveCaller(s))
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[53] = p.nodes[53].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[56] = chainPure(OpCode(17), v[28], v[55])
	v[57] = chainPure(OpCode(21), v[56], v[36])
	v[58] = p.nodes[58].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[59] = p.nodes[59].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[61] = p.nodes[61].c
	v[62] = p.nodes[62].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[64] = p.nodes[64].c
	v[65] = p.nodes[65].c
	v[66] = p.nodes[66].c
	v[67] = p.nodes[67].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[69] = p.nodes[69].c
	v[70].And(&v[68], &v[69])
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[72] = p.nodes[72].c
	v[73] = chainPure(OpCode(17), v[71], v[69])
	v[74] = chainPure(OpCode(21), v[73], v[36])
	v[75] = p.nodes[75].c
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[76] = p.nodes[76].c
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	v[78] = chainPure(OpCode(28), v[43], v[77])
	v[79] = chainPure(OpCode(20), v[45], v[78])
	v[80] = chainPure(OpCode(21), v[79], v[36])
	v[81] = p.nodes[81].c
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	v[82] = p.nodes[82].c
	v[83] = p.nodes[83].c
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	v[87] = p.nodes[87].c
	v[88].Sub(&v[85], &v[86])
	v[89].Add(&v[39], &v[88])
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	v[92] = p.nodes[92].c
	v[93].Sub(&v[90], &v[91])
	v[94].Add(&v[39], &v[93])
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["d423e0578eb448f065cde90593c790f1bb83bd6fe90916b1ba0f8f406c3fbd60"] = chainFastShape106
}

func chainFastShape107(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36].And(&v[25], &v[35])
	v[37].And(&v[35], &v[36])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[38] = p.nodes[38].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[39] = p.nodes[39].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[42] = p.nodes[42].c
	v[43] = p.nodes[43].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[45] = p.nodes[45].c
	v[46].And(&v[26], &v[35])
	v[47].And(&v[35], &v[46])
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[50] = p.nodes[50].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[52] = chainPure(OpCode(17), v[29], v[34])
	v[53] = chainPure(OpCode(21), v[52], v[34])
	v[54] = p.nodes[54].c
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[55] = p.nodes[55].c
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[57].Add(&v[38], &v[56])
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[59].Sub(&v[57], &v[58])
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["d474f5ce9130d843476803a2369cc5919e7553f0ef0854d26ecaf30495b54a9b"] = chainFastShape107
}

func chainFastShape108(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36].And(&v[28], &v[35])
	v[37].And(&v[35], &v[36])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[38] = p.nodes[38].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[39] = p.nodes[39].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[42] = p.nodes[42].c
	v[43] = p.nodes[43].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[45] = p.nodes[45].c
	v[46].And(&v[29], &v[35])
	v[47].And(&v[35], &v[46])
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[50] = p.nodes[50].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[52] = p.nodes[52].c
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[54].Add(&v[38], &v[53])
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[56].Sub(&v[54], &v[55])
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["d8aab68f46939c54aa43ce82d34d15cbfb34d484ab57b6af176fa9c4b083fc6b"] = chainFastShape108
}

func chainFastShape109(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38].And(&v[28], &v[37])
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[41] = p.nodes[41].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[44] = p.nodes[44].c
	v[45] = p.nodes[45].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[48] = p.nodes[48].c
	v[49] = p.nodes[49].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[51].And(&v[29], &v[37])
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[54] = p.nodes[54].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[57] = p.nodes[57].c
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["d8ce30ae2858edd4725682cea21e4055d93cc09174e6676804d2f9e5532d0c77"] = chainFastShape109
}

func chainFastShape110(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[35].SetBytes(effectiveCaller(s))
	v[36] = p.nodes[36].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[39] = p.nodes[39].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[42] = chainPure(OpCode(17), v[34], v[41])
	v[43] = chainPure(OpCode(21), v[42], v[36])
	v[44] = p.nodes[44].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[45] = p.nodes[45].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[48] = p.nodes[48].c
	v[49].And(&v[47], &v[48])
	v[50] = chainPure(OpCode(21), v[49], v[36])
	v[51] = p.nodes[51].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[52] = p.nodes[52].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[55] = p.nodes[55].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[57] = p.nodes[57].c
	v[58] = p.nodes[58].c
	v[59] = p.nodes[59].c
	v[60].And(&v[28], &v[59])
	v[61] = p.nodes[61].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[62] = p.nodes[62].c
	v[63] = p.nodes[63].c
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[65] = p.nodes[65].c
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[67] = p.nodes[67].c
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	v[70] = p.nodes[70].c
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	v[73] = p.nodes[73].c
	v[74] = p.nodes[74].c
	v[75] = p.nodes[75].c
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	v[76].Add(&v[38], &v[72])
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	v[78].Sub(&v[76], &v[77])
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["da5454f454eca3f69ca4a785cf0ffd8f8b51a40d907d8da563a2620ba19974cc"] = chainFastShape110
}

func chainFastShape111(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33].SetBytes(effectiveCaller(s))
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38].And(&v[33], &v[37])
	v[39].And(&v[37], &v[38])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[40] = p.nodes[40].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[41] = p.nodes[41].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[44] = chainPure(OpCode(16), v[43], v[30])
	v[45] = chainPure(OpCode(21), v[44], v[35])
	v[46] = p.nodes[46].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[47] = chainPure(OpCode(20), v[38], v[38])
	v[48] = chainPure(OpCode(21), v[47], v[35])
	v[49] = chainPure(OpCode(21), v[48], v[35])
	v[50] = p.nodes[50].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[51] = p.nodes[51].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[52] = p.nodes[52].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[54] = p.nodes[54].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[56] = p.nodes[56].c
	v[57].And(&v[29], &v[37])
	v[58].And(&v[37], &v[57])
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[61] = p.nodes[61].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[63] = p.nodes[63].c
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[65].Add(&v[40], &v[64])
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	v[67].Sub(&v[65], &v[66])
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["ddf71fbbb78a9f63894627607310a4af4aa5a64e2f04a1ef84122e1950f4e90c"] = chainFastShape111
}

func chainFastShape112(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33].SetBytes(effectiveCaller(s))
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36].And(&v[33], &v[34])
	v[37].And(&v[27], &v[34])
	v[38] = chainPure(OpCode(20), v[36], v[37])
	v[39] = chainPure(OpCode(21), v[38], v[35])
	v[40] = chainPure(OpCode(21), v[39], v[35])
	v[41] = p.nodes[41].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[42] = p.nodes[42].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	v[45].And(&v[34], &v[37])
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[46] = p.nodes[46].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[47] = p.nodes[47].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[50] = p.nodes[50].c
	v[51] = p.nodes[51].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[53] = p.nodes[53].c
	v[54].And(&v[28], &v[34])
	v[55].And(&v[34], &v[54])
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[58] = p.nodes[58].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[60] = p.nodes[60].c
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[62].Add(&v[46], &v[61])
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[64].Sub(&v[62], &v[63])
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["ddfa0a36932ac42bf3529b093a70b87e7f60b711d4f10a2dce366d675fbb8c1d"] = chainFastShape112
}

func chainFastShape113(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = chainPure(OpCode(21), v[30], v[33])
	v[35] = p.nodes[35].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[36] = p.nodes[36].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[37] = p.nodes[37].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[38] = p.nodes[38].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[39] = chainPure(OpCode(17), v[28], v[33])
	v[40] = chainPure(OpCode(21), v[39], v[33])
	v[41] = p.nodes[41].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[42] = p.nodes[42].c
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	v[45].And(&v[25], &v[44])
	v[46].And(&v[44], &v[45])
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[47] = p.nodes[47].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[48] = p.nodes[48].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[51] = p.nodes[51].c
	v[52] = p.nodes[52].c
	v[53] = p.nodes[53].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[55].Add(&v[48], &v[54])
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[56] = p.nodes[56].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[57].Add(&v[47], &v[54])
	v[58] = p.nodes[58].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[59] = p.nodes[59].c
	v[60] = chainPure(OpCode(17), v[27], v[50])
	v[61] = chainPure(OpCode(21), v[60], v[33])
	v[62] = p.nodes[62].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[63] = p.nodes[63].c
	v[64] = p.nodes[64].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[66] = p.nodes[66].c
	v[67] = p.nodes[67].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[69].Add(&v[48], &v[68])
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[70].Add(&v[47], &v[68])
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[71] = chainPure(OpCode(17), v[28], v[27])
	v[72] = chainPure(OpCode(21), v[71], v[33])
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	v[74].And(&v[26], &v[44])
	v[75].And(&v[44], &v[74])
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	v[78] = p.nodes[78].c
	v[79] = p.nodes[79].c
	v[80] = p.nodes[80].c
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	v[82] = chainPure(OpCode(16), v[81], v[77])
	v[83] = chainPure(OpCode(21), v[82], v[33])
	v[84] = p.nodes[84].c
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	v[85] = p.nodes[85].c
	v[86] = p.nodes[86].c
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	v[88].Add(&v[48], &v[87])
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	v[89].Add(&v[47], &v[87])
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	v[92] = p.nodes[92].c
	v[93] = p.nodes[93].c
	v[94].Add(&v[47], &v[91])
	v[95] = p.nodes[95].c
	v[96].Add(&v[33], &v[91])
	v[97] = p.nodes[97].c
	v[98] = p.nodes[98].c
	v[99] = p.nodes[99].c
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	code, err = p.actionFast(44, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 44 != 0, err
	}
	v[101].Sub(&v[94], &v[100])
	code, err = p.actionFast(45, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 45 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["e05befa340b2d81ee0645a30815fcde9c4ecfc7ab8a4af12dd93b10036cb3945"] = chainFastShape113
}

func chainFastShape114(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36].And(&v[31], &v[34])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[39] = p.nodes[39].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[42] = p.nodes[42].c
	v[43] = chainPure(OpCode(28), v[42], v[41])
	v[44] = p.nodes[44].c
	v[45] = chainPure(OpCode(20), v[43], v[44])
	v[46] = chainPure(OpCode(21), v[45], v[35])
	v[47] = p.nodes[47].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[48] = p.nodes[48].c
	v[49].SetBytes(effectiveCaller(s))
	v[50] = p.nodes[50].c
	v[51].And(&v[34], &v[49])
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[54] = p.nodes[54].c
	v[55].And(&v[53], &v[54])
	v[56] = chainPure(OpCode(17), v[30], v[35])
	v[57] = p.nodes[57].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[58] = chainPure(OpCode(16), v[55], v[30])
	v[59] = chainPure(OpCode(21), v[58], v[35])
	v[60] = p.nodes[60].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[61] = p.nodes[61].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[63] = p.nodes[63].c
	v[64] = p.nodes[64].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[66] = p.nodes[66].c
	v[67] = p.nodes[67].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[69] = p.nodes[69].c
	v[70] = chainPure(OpCode(17), v[68], v[54])
	v[71] = chainPure(OpCode(21), v[70], v[35])
	v[72] = p.nodes[72].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[73] = p.nodes[73].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[75] = chainPure(OpCode(28), v[42], v[74])
	v[76] = chainPure(OpCode(20), v[44], v[75])
	v[77] = chainPure(OpCode(21), v[76], v[35])
	v[78] = p.nodes[78].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	v[81] = p.nodes[81].c
	v[82].Sub(&v[79], &v[80])
	v[83].Add(&v[38], &v[82])
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	v[86] = p.nodes[86].c
	v[87].Sub(&v[84], &v[85])
	v[88].Add(&v[38], &v[87])
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["e069a73b73c0fec25ff7f7d1175632ce997e599a4457c7c1cd280fe12a57136b"] = chainFastShape114
}

func chainFastShape115(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	v[40].Div(&v[34], &v[39])
	v[41] = p.nodes[41].c
	v[42].And(&v[40], &v[41])
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[46] = p.nodes[46].c
	v[47] = p.nodes[47].c
	v[48] = p.nodes[48].c
	v[49] = p.nodes[49].c
	v[50].And(&v[27], &v[49])
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[51] = p.nodes[51].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[52] = p.nodes[52].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[55] = p.nodes[55].c
	v[56] = p.nodes[56].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[58] = chainPure(OpCode(17), v[57], v[45])
	v[59] = chainPure(OpCode(21), v[58], v[35])
	v[60] = p.nodes[60].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[61] = p.nodes[61].c
	v[62] = p.nodes[62].c
	v[63] = p.nodes[63].c
	v[64] = p.nodes[64].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[65] = chainPure(OpCode(21), v[30], v[35])
	v[66] = p.nodes[66].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[67] = p.nodes[67].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[69] = p.nodes[69].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[70].Add(&v[51], &v[68])
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[72].Sub(&v[70], &v[71])
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[73] = p.nodes[73].c
	v[74] = p.nodes[74].c
	v[75].And(&v[26], &v[49])
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[78] = p.nodes[78].c
	v[79] = p.nodes[79].c
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[81] = p.nodes[81].c
	v[82] = p.nodes[82].c
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["e12a09d757f904d02709ddd82fa57cbb6a289e192aef3ccc5d7a134c7e164391"] = chainFastShape115
}

func chainFastShape116(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36].And(&v[25], &v[34])
	v[37] = p.nodes[37].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[38].And(&v[26], &v[34])
	v[39] = p.nodes[39].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[40] = p.nodes[40].c
	v[41] = p.nodes[41].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[42] = p.nodes[42].c
	v[43] = p.nodes[43].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[44] = p.nodes[44].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[47] = p.nodes[47].c
	v[48].And(&v[46], &v[47])
	v[49] = chainPure(OpCode(17), v[27], v[48])
	v[50] = chainPure(OpCode(21), v[49], v[35])
	v[51] = p.nodes[51].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[52] = p.nodes[52].c
	v[53] = p.nodes[53].c
	v[54] = p.nodes[54].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[56].And(&v[47], &v[55])
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[58] = p.nodes[58].c
	v[59] = chainPure(OpCode(17), v[57], v[47])
	v[60] = chainPure(OpCode(21), v[59], v[35])
	v[61] = p.nodes[61].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[62] = p.nodes[62].c
	v[63] = p.nodes[63].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[65] = p.nodes[65].c
	v[66] = chainPure(OpCode(28), v[65], v[64])
	v[67] = p.nodes[67].c
	v[68] = chainPure(OpCode(20), v[66], v[67])
	v[69] = chainPure(OpCode(21), v[68], v[35])
	v[70] = p.nodes[70].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[71] = p.nodes[71].c
	v[72] = p.nodes[72].c
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	v[75].And(&v[47], &v[74])
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[77] = chainPure(OpCode(17), v[76], v[47])
	v[78] = chainPure(OpCode(21), v[77], v[35])
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	v[80] = chainPure(OpCode(28), v[65], v[79])
	v[81] = chainPure(OpCode(20), v[67], v[80])
	v[82] = chainPure(OpCode(21), v[81], v[35])
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	v[83] = p.nodes[83].c
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	v[85].Add(&v[43], &v[84])
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	v[87].Sub(&v[85], &v[86])
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	v[88] = p.nodes[88].c
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	v[90].SetBytes(effectiveCaller(s))
	code, err = p.actionFast(44, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 44 != 0, err
	}
	code, err = p.actionFast(45, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 45 != 0, err
	}
	code, err = p.actionFast(46, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 46 != 0, err
	}
	code, err = p.actionFast(47, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 47 != 0, err
	}
	v[93] = p.nodes[93].c
	v[94] = p.nodes[94].c
	code, err = p.actionFast(48, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 48 != 0, err
	}
	code, err = p.actionFast(49, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 49 != 0, err
	}
	code, err = p.actionFast(50, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 50 != 0, err
	}
	code, err = p.actionFast(51, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 51 != 0, err
	}
	code, err = p.actionFast(52, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 52 != 0, err
	}
	code, err = p.actionFast(53, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 53 != 0, err
	}
	code, err = p.actionFast(54, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 54 != 0, err
	}
	code, err = p.actionFast(55, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 55 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["e1d060dbabb400a5a930ad1958a5961ca6c15f05285e935e4675561f011bd6f0"] = chainFastShape116
}

func chainFastShape117(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38].And(&v[29], &v[36])
	v[39] = chainPure(OpCode(21), v[38], v[37])
	v[40] = p.nodes[40].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[41] = chainPure(OpCode(21), v[30], v[37])
	v[42] = p.nodes[42].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	v[45] = p.nodes[45].c
	v[46] = p.nodes[46].c
	v[47].And(&v[27], &v[46])
	v[48] = p.nodes[48].c
	v[49] = p.nodes[49].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[51] = p.nodes[51].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[52] = p.nodes[52].c
	v[53].Add(&v[50], &v[52])
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[55].Sub(&v[53], &v[54])
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[56] = p.nodes[56].c
	v[57] = p.nodes[57].c
	v[58].And(&v[26], &v[46])
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[61] = p.nodes[61].c
	v[62] = p.nodes[62].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[64] = p.nodes[64].c
	v[65] = p.nodes[65].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[68] = p.nodes[68].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["e1d6af049d20901bb8b611c3f160bc0251046b7eff4641cf3c7a77fa960172ae"] = chainFastShape117
}

func chainFastShape118(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = chainPure(OpCode(17), v[30], v[33])
	v[35] = p.nodes[35].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[36] = p.nodes[36].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	v[41] = p.nodes[41].c
	v[42].And(&v[37], &v[41])
	v[43].And(&v[41], &v[42])
	v[44].And(&v[28], &v[41])
	v[45] = chainPure(OpCode(20), v[43], v[44])
	v[46] = chainPure(OpCode(21), v[45], v[33])
	v[47] = chainPure(OpCode(21), v[46], v[33])
	v[48] = p.nodes[48].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[49] = p.nodes[49].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[50] = p.nodes[50].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[51] = p.nodes[51].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[52] = p.nodes[52].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[53] = p.nodes[53].c
	v[54] = p.nodes[54].c
	v[55] = p.nodes[55].c
	v[56] = p.nodes[56].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[57] = p.nodes[57].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[58].And(&v[29], &v[41])
	v[59] = p.nodes[59].c
	v[60] = p.nodes[60].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[62] = p.nodes[62].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[63] = p.nodes[63].c
	v[64].Add(&v[61], &v[63])
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[66].Sub(&v[64], &v[65])
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[67] = p.nodes[67].c
	v[68] = p.nodes[68].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[71] = p.nodes[71].c
	v[72] = p.nodes[72].c
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[74] = p.nodes[74].c
	v[75] = p.nodes[75].c
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[78] = p.nodes[78].c
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["e3729438f1b53bcf76b1ede8e4acbc8e7fc3de566181b7b0ba8d552916690c88"] = chainFastShape118
}

func chainFastShape119(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35].SetBytes(effectiveCaller(s))
	v[36] = p.nodes[36].c
	v[37].And(&v[35], &v[36])
	v[38].And(&v[36], &v[37])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[39] = p.nodes[39].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[40] = p.nodes[40].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[46] = p.nodes[46].c
	v[47].And(&v[27], &v[36])
	v[48].And(&v[36], &v[47])
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[51] = p.nodes[51].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[53] = chainPure(OpCode(17), v[29], v[34])
	v[54] = chainPure(OpCode(21), v[53], v[34])
	v[55] = p.nodes[55].c
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[56] = p.nodes[56].c
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[58].Add(&v[39], &v[57])
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[60].Sub(&v[58], &v[59])
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["e4de5faf548f52a579f141afa8dc5b358b83e4dff108e386bd0e2c6d433b699e"] = chainFastShape119
}

func chainFastShape120(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	v[40].And(&v[35], &v[39])
	v[41].And(&v[39], &v[40])
	v[42].And(&v[28], &v[39])
	v[43] = chainPure(OpCode(20), v[41], v[42])
	v[44] = chainPure(OpCode(21), v[43], v[33])
	v[45] = chainPure(OpCode(21), v[44], v[33])
	v[46] = p.nodes[46].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[47] = p.nodes[47].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[48] = p.nodes[48].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[49] = p.nodes[49].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[50] = p.nodes[50].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[51] = p.nodes[51].c
	v[52] = p.nodes[52].c
	v[53] = p.nodes[53].c
	v[54] = p.nodes[54].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[55] = p.nodes[55].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[56].And(&v[29], &v[39])
	v[57] = p.nodes[57].c
	v[58] = p.nodes[58].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[60] = p.nodes[60].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[61] = p.nodes[61].c
	v[62].Add(&v[59], &v[61])
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[64].Sub(&v[62], &v[63])
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[65] = p.nodes[65].c
	v[66] = p.nodes[66].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[69] = p.nodes[69].c
	v[70] = p.nodes[70].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[72] = p.nodes[72].c
	v[73] = p.nodes[73].c
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	v[76] = p.nodes[76].c
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["e70b80cb49e307ba403498d3539abbe8b9886922f9370a8c35ba65d5632e18c2"] = chainFastShape120
}

func chainFastShape121(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35].SetBytes(effectiveCaller(s))
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	v[41] = p.nodes[41].c
	v[42].And(&v[29], &v[41])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[43] = p.nodes[43].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[44] = p.nodes[44].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[46].And(&v[35], &v[41])
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[49] = chainPure(OpCode(16), v[48], v[30])
	v[50] = chainPure(OpCode(21), v[49], v[39])
	v[51] = p.nodes[51].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[52] = p.nodes[52].c
	v[53] = p.nodes[53].c
	v[54] = p.nodes[54].c
	v[55] = p.nodes[55].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[57] = p.nodes[57].c
	v[58] = p.nodes[58].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[59] = p.nodes[59].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[61] = p.nodes[61].c
	v[62] = p.nodes[62].c
	v[63] = p.nodes[63].c
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[64].Add(&v[43], &v[60])
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[66].Sub(&v[64], &v[65])
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[67] = p.nodes[67].c
	v[68] = p.nodes[68].c
	v[69] = p.nodes[69].c
	v[70] = p.nodes[70].c
	v[71] = p.nodes[71].c
	v[72] = p.nodes[72].c
	v[73] = p.nodes[73].c
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[74] = p.nodes[74].c
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[77] = chainPure(OpCode(16), v[76], v[30])
	v[78] = chainPure(OpCode(21), v[77], v[39])
	v[79] = p.nodes[79].c
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	v[80] = p.nodes[80].c
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	v[82] = p.nodes[82].c
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	v[84] = p.nodes[84].c
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	v[87] = p.nodes[87].c
	v[88] = p.nodes[88].c
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	v[89].Add(&v[43], &v[86])
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	v[91].Sub(&v[89], &v[90])
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	v[92] = p.nodes[92].c
	v[93] = p.nodes[93].c
	v[94] = p.nodes[94].c
	v[95] = p.nodes[95].c
	v[96] = p.nodes[96].c
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	v[97] = p.nodes[97].c
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	v[100].And(&v[41], &v[99])
	v[101] = p.nodes[101].c
	code, err = p.actionFast(44, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 44 != 0, err
	}
	code, err = p.actionFast(45, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 45 != 0, err
	}
	code, err = p.actionFast(46, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 46 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["e74144587a421edb9ba89b88730801cefe43ba83674ce8a4ceb778fda43bc14a"] = chainFastShape121
}

func chainFastShape122(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36].And(&v[31], &v[34])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[39] = p.nodes[39].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[42] = p.nodes[42].c
	v[43] = chainPure(OpCode(28), v[42], v[41])
	v[44] = p.nodes[44].c
	v[45] = chainPure(OpCode(20), v[43], v[44])
	v[46] = chainPure(OpCode(21), v[45], v[35])
	v[47] = p.nodes[47].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[48].And(&v[27], &v[34])
	v[49] = p.nodes[49].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[50] = chainPure(OpCode(17), v[28], v[35])
	v[51] = p.nodes[51].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[52].SetBytes(effectiveCaller(s))
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[53] = p.nodes[53].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[56] = chainPure(OpCode(17), v[28], v[55])
	v[57] = chainPure(OpCode(21), v[56], v[35])
	v[58] = p.nodes[58].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[59] = p.nodes[59].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[61] = p.nodes[61].c
	v[62] = p.nodes[62].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[64] = p.nodes[64].c
	v[65] = p.nodes[65].c
	v[66] = p.nodes[66].c
	v[67] = p.nodes[67].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[70] = p.nodes[70].c
	v[71].And(&v[69], &v[70])
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[73] = p.nodes[73].c
	v[74] = chainPure(OpCode(17), v[72], v[70])
	v[75] = chainPure(OpCode(21), v[74], v[35])
	v[76] = p.nodes[76].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[77] = p.nodes[77].c
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[79] = chainPure(OpCode(28), v[42], v[78])
	v[80] = chainPure(OpCode(20), v[44], v[79])
	v[81] = chainPure(OpCode(21), v[80], v[35])
	v[82] = p.nodes[82].c
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	v[83] = p.nodes[83].c
	v[84] = p.nodes[84].c
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	v[88] = p.nodes[88].c
	v[89].Sub(&v[86], &v[87])
	v[90].Add(&v[38], &v[89])
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	v[93] = p.nodes[93].c
	v[94].Sub(&v[91], &v[92])
	v[95].Add(&v[38], &v[94])
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["edd5853c3bb41f03b3abd85149c08fd6666f85e9373846eae3b7ce6f7122f72d"] = chainFastShape122
}

func chainFastShape123(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[37] = p.nodes[37].c
	v[38].And(&v[27], &v[37])
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[41] = p.nodes[41].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[43].SetBytes(effectiveCaller(s))
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[45] = p.nodes[45].c
	v[46] = p.nodes[46].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[47] = p.nodes[47].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[50] = p.nodes[50].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[52].And(&v[28], &v[37])
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[55] = p.nodes[55].c
	v[56] = p.nodes[56].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	v[60] = p.nodes[60].c
	v[61].Sub(&v[58], &v[59])
	v[62].Add(&v[40], &v[61])
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["f1497674ecc847bddf94550be13b24868dea9656ea599b08c803365e369a902d"] = chainFastShape123
}

func chainFastShape124(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33].SetBytes(effectiveCaller(s))
	v[34] = p.nodes[34].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[37] = p.nodes[37].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[39] = p.nodes[39].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[41] = chainPure(OpCode(21), v[30], v[34])
	v[42] = p.nodes[42].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[43] = p.nodes[43].c
	v[44].And(&v[29], &v[43])
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[47] = p.nodes[47].c
	v[48].And(&v[46], &v[47])
	v[49] = p.nodes[49].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[51] = p.nodes[51].c
	v[52].And(&v[50], &v[51])
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["f24763fd44475b16b9c0b0eb426018ccffc3d59d7fbc0a7d22eb216a23530a50"] = chainFastShape124
}

func chainFastShape125(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38].And(&v[29], &v[36])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[41] = p.nodes[41].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[44] = p.nodes[44].c
	v[45].And(&v[43], &v[44])
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[47] = p.nodes[47].c
	v[48] = chainPure(OpCode(17), v[46], v[44])
	v[49] = chainPure(OpCode(21), v[48], v[37])
	v[50] = p.nodes[50].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[51] = p.nodes[51].c
	v[52] = p.nodes[52].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[54] = p.nodes[54].c
	v[55] = chainPure(OpCode(28), v[54], v[53])
	v[56] = p.nodes[56].c
	v[57] = chainPure(OpCode(20), v[55], v[56])
	v[58] = chainPure(OpCode(21), v[57], v[37])
	v[59] = p.nodes[59].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[60] = p.nodes[60].c
	v[61] = p.nodes[61].c
	v[62].And(&v[30], &v[36])
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[65].And(&v[44], &v[64])
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[67] = chainPure(OpCode(17), v[66], v[44])
	v[68] = chainPure(OpCode(21), v[67], v[37])
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[70] = chainPure(OpCode(28), v[54], v[69])
	v[71] = chainPure(OpCode(20), v[56], v[70])
	v[72] = chainPure(OpCode(21), v[71], v[37])
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	v[73] = p.nodes[73].c
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	v[75].Add(&v[40], &v[74])
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	v[77].Sub(&v[75], &v[76])
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["f380b3b81ef05d8654205418870b63778eae5be85171c3c05ce252c023efb8a6"] = chainFastShape125
}

func chainFastShape126(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37] = p.nodes[37].c
	v[38].And(&v[28], &v[37])
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[41] = p.nodes[41].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[44] = p.nodes[44].c
	v[45] = p.nodes[45].c
	v[46] = p.nodes[46].c
	v[47] = p.nodes[47].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[49].Add(&v[41], &v[48])
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[50] = p.nodes[50].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[51].Add(&v[40], &v[48])
	v[52] = p.nodes[52].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[53] = p.nodes[53].c
	v[54] = chainPure(OpCode(17), v[30], v[43])
	v[55] = chainPure(OpCode(21), v[54], v[35])
	v[56] = p.nodes[56].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[57].Sub(&v[43], &v[30])
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[58] = p.nodes[58].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[62] = p.nodes[62].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[64].Add(&v[41], &v[63])
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[65].Add(&v[40], &v[63])
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[66] = chainPure(OpCode(17), v[30], v[61])
	v[67] = chainPure(OpCode(21), v[66], v[35])
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[68].Sub(&v[61], &v[30])
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	v[69] = p.nodes[69].c
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	v[72] = p.nodes[72].c
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	v[74].Add(&v[41], &v[73])
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	v[75].Add(&v[40], &v[73])
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	v[76] = chainPure(OpCode(17), v[30], v[71])
	v[77] = chainPure(OpCode(21), v[76], v[35])
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	v[78].Sub(&v[71], &v[30])
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	v[79] = chainPure(OpCode(21), v[78], v[35])
	v[80] = chainPure(OpCode(21), v[79], v[35])
	v[81] = p.nodes[81].c
	code, err = p.actionFast(44, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 44 != 0, err
	}
	v[82] = p.nodes[82].c
	code, err = p.actionFast(45, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 45 != 0, err
	}
	code, err = p.actionFast(46, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 46 != 0, err
	}
	code, err = p.actionFast(47, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 47 != 0, err
	}
	code, err = p.actionFast(48, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 48 != 0, err
	}
	code, err = p.actionFast(49, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 49 != 0, err
	}
	code, err = p.actionFast(50, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 50 != 0, err
	}
	code, err = p.actionFast(51, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 51 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["f80136b81f60b21ca80b3a705d258636943c00d42f2423ab6a5a41583b4ad01a"] = chainFastShape126
}

func chainFastShape127(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35].And(&v[29], &v[33])
	v[36] = p.nodes[36].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	v[39].And(&v[28], &v[33])
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[40] = p.nodes[40].c
	v[41] = p.nodes[41].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[42] = p.nodes[42].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[45] = p.nodes[45].c
	v[46].And(&v[44], &v[45])
	v[47] = chainPure(OpCode(17), v[30], v[46])
	v[48] = chainPure(OpCode(21), v[47], v[34])
	v[49] = p.nodes[49].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[50] = p.nodes[50].c
	v[51] = p.nodes[51].c
	v[52] = p.nodes[52].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[54].And(&v[45], &v[53])
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[56] = p.nodes[56].c
	v[57] = chainPure(OpCode(17), v[55], v[45])
	v[58] = chainPure(OpCode(21), v[57], v[34])
	v[59] = p.nodes[59].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[60] = p.nodes[60].c
	v[61] = p.nodes[61].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[63] = p.nodes[63].c
	v[64] = chainPure(OpCode(28), v[63], v[62])
	v[65] = p.nodes[65].c
	v[66] = chainPure(OpCode(20), v[64], v[65])
	v[67] = chainPure(OpCode(21), v[66], v[34])
	v[68] = p.nodes[68].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[69] = p.nodes[69].c
	v[70] = p.nodes[70].c
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	v[73].And(&v[45], &v[72])
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[75] = chainPure(OpCode(17), v[74], v[45])
	v[76] = chainPure(OpCode(21), v[75], v[34])
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	v[78] = chainPure(OpCode(28), v[63], v[77])
	v[79] = chainPure(OpCode(20), v[65], v[78])
	v[80] = chainPure(OpCode(21), v[79], v[34])
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	v[81] = p.nodes[81].c
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	v[83].Add(&v[41], &v[82])
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	v[85].Sub(&v[83], &v[84])
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["f9b0d42fe67176a8341d4d6f02755c7ea6df85205f856f5f0e0116fe63740873"] = chainFastShape127
}

func chainFastShape128(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = chainPure(OpCode(16), v[30], v[33])
	v[36] = chainPure(OpCode(21), v[35], v[34])
	v[37] = p.nodes[37].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[38].SetBytes(effectiveCaller(s))
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[41] = p.nodes[41].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[44] = p.nodes[44].c
	v[45].And(&v[43], &v[44])
	v[46] = chainPure(OpCode(21), v[45], v[34])
	v[47] = p.nodes[47].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[49] = p.nodes[49].c
	v[50].And(&v[48], &v[49])
	v[51] = p.nodes[51].c
	v[52].Or(&v[50], &v[51])
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[53] = p.nodes[53].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[56] = p.nodes[56].c
	v[57] = p.nodes[57].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[59] = p.nodes[59].c
	v[60] = p.nodes[60].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[61] = p.nodes[61].c
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[63] = p.nodes[63].c
	v[64] = p.nodes[64].c
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[66] = chainPure(OpCode(20), v[55], v[65])
	v[67] = p.nodes[67].c
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	v[69].And(&v[49], &v[68])
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["fa663289ded7813dc734932c26c3ce808558975bd0dd85e2e7d4d5a20d489821"] = chainFastShape128
}

func chainFastShape129(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
	in := s.Stack.data
	v := e.chainValues
	var code int
	var err error
	if p.nodes[0].depth <= len(in) {
		v[0] = in[len(in)-p.nodes[0].depth]
	}
	if p.nodes[1].depth <= len(in) {
		v[1] = in[len(in)-p.nodes[1].depth]
	}
	if p.nodes[2].depth <= len(in) {
		v[2] = in[len(in)-p.nodes[2].depth]
	}
	if p.nodes[3].depth <= len(in) {
		v[3] = in[len(in)-p.nodes[3].depth]
	}
	if p.nodes[4].depth <= len(in) {
		v[4] = in[len(in)-p.nodes[4].depth]
	}
	if p.nodes[5].depth <= len(in) {
		v[5] = in[len(in)-p.nodes[5].depth]
	}
	if p.nodes[6].depth <= len(in) {
		v[6] = in[len(in)-p.nodes[6].depth]
	}
	if p.nodes[7].depth <= len(in) {
		v[7] = in[len(in)-p.nodes[7].depth]
	}
	if p.nodes[8].depth <= len(in) {
		v[8] = in[len(in)-p.nodes[8].depth]
	}
	if p.nodes[9].depth <= len(in) {
		v[9] = in[len(in)-p.nodes[9].depth]
	}
	if p.nodes[10].depth <= len(in) {
		v[10] = in[len(in)-p.nodes[10].depth]
	}
	if p.nodes[11].depth <= len(in) {
		v[11] = in[len(in)-p.nodes[11].depth]
	}
	if p.nodes[12].depth <= len(in) {
		v[12] = in[len(in)-p.nodes[12].depth]
	}
	if p.nodes[13].depth <= len(in) {
		v[13] = in[len(in)-p.nodes[13].depth]
	}
	if p.nodes[14].depth <= len(in) {
		v[14] = in[len(in)-p.nodes[14].depth]
	}
	if p.nodes[15].depth <= len(in) {
		v[15] = in[len(in)-p.nodes[15].depth]
	}
	if p.nodes[16].depth <= len(in) {
		v[16] = in[len(in)-p.nodes[16].depth]
	}
	if p.nodes[17].depth <= len(in) {
		v[17] = in[len(in)-p.nodes[17].depth]
	}
	if p.nodes[18].depth <= len(in) {
		v[18] = in[len(in)-p.nodes[18].depth]
	}
	if p.nodes[19].depth <= len(in) {
		v[19] = in[len(in)-p.nodes[19].depth]
	}
	if p.nodes[20].depth <= len(in) {
		v[20] = in[len(in)-p.nodes[20].depth]
	}
	if p.nodes[21].depth <= len(in) {
		v[21] = in[len(in)-p.nodes[21].depth]
	}
	if p.nodes[22].depth <= len(in) {
		v[22] = in[len(in)-p.nodes[22].depth]
	}
	if p.nodes[23].depth <= len(in) {
		v[23] = in[len(in)-p.nodes[23].depth]
	}
	if p.nodes[24].depth <= len(in) {
		v[24] = in[len(in)-p.nodes[24].depth]
	}
	if p.nodes[25].depth <= len(in) {
		v[25] = in[len(in)-p.nodes[25].depth]
	}
	if p.nodes[26].depth <= len(in) {
		v[26] = in[len(in)-p.nodes[26].depth]
	}
	if p.nodes[27].depth <= len(in) {
		v[27] = in[len(in)-p.nodes[27].depth]
	}
	if p.nodes[28].depth <= len(in) {
		v[28] = in[len(in)-p.nodes[28].depth]
	}
	if p.nodes[29].depth <= len(in) {
		v[29] = in[len(in)-p.nodes[29].depth]
	}
	if p.nodes[30].depth <= len(in) {
		v[30] = in[len(in)-p.nodes[30].depth]
	}
	if p.nodes[31].depth <= len(in) {
		v[31] = in[len(in)-p.nodes[31].depth]
	}
	v[32] = p.nodes[32].c
	v[33] = p.nodes[33].c
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	v[37].Div(&v[31], &v[36])
	v[38] = p.nodes[38].c
	v[39].And(&v[37], &v[38])
	v[40] = p.nodes[40].c
	v[41] = p.nodes[41].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[43] = chainPure(OpCode(17), v[28], v[42])
	v[44] = chainPure(OpCode(21), v[43], v[35])
	v[45] = p.nodes[45].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[46] = p.nodes[46].c
	v[47].And(&v[26], &v[46])
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[48] = p.nodes[48].c
	v[49] = p.nodes[49].c
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[50] = p.nodes[50].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[53] = p.nodes[53].c
	v[54] = p.nodes[54].c
	v[55].Div(&v[52], &v[54])
	v[56] = p.nodes[56].c
	v[57].And(&v[55], &v[56])
	v[58] = chainPure(OpCode(21), v[57], v[35])
	v[59] = p.nodes[59].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[60] = p.nodes[60].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[61].And(&v[27], &v[46])
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[64].Div(&v[63], &v[54])
	v[65].And(&v[56], &v[64])
	v[66] = chainPure(OpCode(21), v[65], v[35])
	v[67] = p.nodes[67].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[68] = p.nodes[68].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[69] = p.nodes[69].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[71] = p.nodes[71].c
	v[72] = p.nodes[72].c
	v[73] = p.nodes[73].c
	v[74].Div(&v[70], &v[73])
	v[75].And(&v[38], &v[74])
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[77] = p.nodes[77].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[80] = p.nodes[80].c
	v[81] = p.nodes[81].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[83] = chainPure(OpCode(17), v[82], v[76])
	v[84] = chainPure(OpCode(21), v[83], v[35])
	v[85] = p.nodes[85].c
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[86] = p.nodes[86].c
	v[87] = p.nodes[87].c
	v[88] = p.nodes[88].c
	v[89] = p.nodes[89].c
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[90] = chainPure(OpCode(21), v[30], v[35])
	v[91] = p.nodes[91].c
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[92] = p.nodes[92].c
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[94] = p.nodes[94].c
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	v[95].Add(&v[49], &v[93])
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[97].Sub(&v[95], &v[96])
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	v[98] = p.nodes[98].c
	v[99] = p.nodes[99].c
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	code, err = p.actionFast(31, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 31 != 0, err
	}
	v[102] = p.nodes[102].c
	v[103] = p.nodes[103].c
	code, err = p.actionFast(32, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 32 != 0, err
	}
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	v[105] = p.nodes[105].c
	v[106] = p.nodes[106].c
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	code, err = p.actionFast(41, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 41 != 0, err
	}
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	code, err = p.actionFast(44, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 44 != 0, err
	}
	code, err = p.actionFast(45, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 45 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["fb9db1064797f02d2cbcbe984908bef093581a9b750531932a231862f65ea36c"] = chainFastShape129
}
