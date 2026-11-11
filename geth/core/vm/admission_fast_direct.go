// Code generated from effect summaries. DO NOT EDIT.
package vm

func admissionFastShape0(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[35] = p.nodes[35].c
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
	v[41].Div(&v[37], &v[40])
	v[42] = p.nodes[42].c
	v[43].And(&v[41], &v[42])
	v[44] = chainPure(OpCode(21), v[43], v[32])
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
	v[47] = p.nodes[47].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[49] = p.nodes[49].c
	v[50] = p.nodes[50].c
	v[51] = p.nodes[51].c
	v[52].Div(&v[48], &v[51])
	v[53] = p.nodes[53].c
	v[54].And(&v[52], &v[53])
	v[55] = p.nodes[55].c
	v[56] = p.nodes[56].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[58] = p.nodes[58].c
	v[59] = p.nodes[59].c
	v[60] = p.nodes[60].c
	v[61] = p.nodes[61].c
	v[62].And(&v[27], &v[61])
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
	v[65] = p.nodes[65].c
	v[66] = p.nodes[66].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[68] = chainPure(OpCode(17), v[67], v[57])
	v[69] = chainPure(OpCode(21), v[68], v[32])
	v[70] = p.nodes[70].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[71] = p.nodes[71].c
	v[72] = p.nodes[72].c
	v[73] = p.nodes[73].c
	v[74] = p.nodes[74].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[75] = chainPure(OpCode(21), v[30], v[32])
	v[76] = p.nodes[76].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[77] = p.nodes[77].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[79] = p.nodes[79].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[80].Add(&v[34], &v[78])
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[82].Sub(&v[80], &v[81])
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[83] = p.nodes[83].c
	v[84] = p.nodes[84].c
	v[85].And(&v[26], &v[61])
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
	v[88] = p.nodes[88].c
	v[89] = p.nodes[89].c
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
	v[91] = p.nodes[91].c
	v[92] = p.nodes[92].c
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
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["1946b796f184fe072c4e8e1b1e17fd558338ee63c33a81cd42a747196a90a3c4"] = admissionFastShape0
}

func admissionFastShape1(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[32].SetBytes(effectiveCaller(s))
	v[33] = p.nodes[33].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[34] = p.nodes[34].c
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
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[40] = chainPure(OpCode(21), v[31], v[33])
	v[41] = p.nodes[41].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[42] = p.nodes[42].c
	v[43].And(&v[30], &v[42])
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
	v[46] = p.nodes[46].c
	v[47].And(&v[45], &v[46])
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
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
	v[50] = p.nodes[50].c
	v[51].And(&v[49], &v[50])
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["2749fd0a9f54990acaed2dd8066b54caa30a983df9dd0e45b9721f5125c526a4"] = admissionFastShape1
}

func admissionFastShape2(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[35] = p.nodes[35].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[37] = p.nodes[37].c
	v[38] = p.nodes[38].c
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	v[41].And(&v[25], &v[40])
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
	v[43].And(&v[24], &v[40])
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
	v[46] = chainPure(OpCode(17), v[27], v[45])
	v[47] = chainPure(OpCode(21), v[46], v[32])
	v[48] = p.nodes[48].c
	v[49] = p.nodes[49].c
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
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[51] = chainPure(OpCode(17), v[27], v[50])
	v[52] = p.nodes[52].c
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
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["3fb42a1e47fb9873a2477184f9dbe389d66ba9f363ded358040292b6737a4c22"] = admissionFastShape2
}

func admissionFastShape3(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[32].SetBytes(effectiveCaller(s))
	v[33] = p.nodes[33].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[34] = p.nodes[34].c
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
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[42] = p.nodes[42].c
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
	v[45] = chainPure(OpCode(17), v[41], v[44])
	v[46] = chainPure(OpCode(21), v[45], v[33])
	v[47] = p.nodes[47].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[48] = p.nodes[48].c
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
	v[51] = p.nodes[51].c
	v[52].And(&v[50], &v[51])
	v[53] = chainPure(OpCode(21), v[52], v[33])
	v[54] = p.nodes[54].c
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
	v[56] = p.nodes[56].c
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
	v[58] = p.nodes[58].c
	v[59] = p.nodes[59].c
	v[60] = p.nodes[60].c
	v[61].And(&v[30], &v[60])
	v[62] = p.nodes[62].c
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[63] = p.nodes[63].c
	v[64] = p.nodes[64].c
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	v[66] = p.nodes[66].c
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
	v[68] = p.nodes[68].c
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
	v[74] = p.nodes[74].c
	v[75] = p.nodes[75].c
	v[76] = p.nodes[76].c
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	v[77].Add(&v[35], &v[73])
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	v[79].Sub(&v[77], &v[78])
	code, err = p.actionFast(40, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 40 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["4b9059781b113c3ed677d97260063e5d79df76dfc5ec5ff755c5737e308ea760"] = admissionFastShape3
}

func admissionFastShape4(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[32].SetBytes(effectiveCaller(s))
	v[33] = p.nodes[33].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[34] = p.nodes[34].c
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
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[39] = p.nodes[39].c
	v[40].And(&v[38], &v[39])
	v[41] = p.nodes[41].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[42] = p.nodes[42].c
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	v[45].And(&v[32], &v[44])
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
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[49] = chainPure(OpCode(28), v[39], v[48])
	v[50] = p.nodes[50].c
	v[51] = chainPure(OpCode(20), v[49], v[50])
	v[52] = chainPure(OpCode(21), v[51], v[33])
	v[53] = p.nodes[53].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[54] = p.nodes[54].c
	v[55].And(&v[29], &v[44])
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
	v[58] = chainPure(OpCode(28), v[39], v[57])
	v[59] = chainPure(OpCode(20), v[50], v[58])
	v[60] = chainPure(OpCode(21), v[59], v[33])
	v[61] = p.nodes[61].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[62] = p.nodes[62].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[63] = chainPure(OpCode(17), v[30], v[33])
	v[64] = p.nodes[64].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
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
	v[68] = chainPure(OpCode(17), v[30], v[67])
	v[69] = chainPure(OpCode(21), v[68], v[33])
	v[70] = p.nodes[70].c
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[71] = p.nodes[71].c
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
	v[76] = p.nodes[76].c
	v[77] = p.nodes[77].c
	v[78] = p.nodes[78].c
	v[79] = p.nodes[79].c
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
	v[81] = p.nodes[81].c
	v[82].And(&v[80], &v[81])
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	v[84] = p.nodes[84].c
	v[85] = chainPure(OpCode(17), v[83], v[81])
	v[86] = chainPure(OpCode(21), v[85], v[33])
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
	v[90] = chainPure(OpCode(28), v[39], v[89])
	v[91] = chainPure(OpCode(20), v[50], v[90])
	v[92] = chainPure(OpCode(21), v[91], v[33])
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
	v[99] = p.nodes[99].c
	v[100].Sub(&v[97], &v[98])
	v[101].Add(&v[35], &v[100])
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
	v[104] = p.nodes[104].c
	v[105].Sub(&v[102], &v[103])
	v[106].Add(&v[35], &v[105])
	code, err = p.actionFast(52, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 52 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["4c13b60d42355f77467424d0cde7814cca81db6a52bc750f3e70c7797e0c0254"] = admissionFastShape4
}

func admissionFastShape5(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[35] = p.nodes[35].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[38] = p.nodes[38].c
	v[39].And(&v[37], &v[38])
	v[40] = chainPure(OpCode(21), v[39], v[32])
	v[41] = chainPure(OpCode(21), v[40], v[32])
	v[42] = p.nodes[42].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[43] = p.nodes[43].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[44] = p.nodes[44].c
	v[45] = p.nodes[45].c
	v[46] = p.nodes[46].c
	v[47] = p.nodes[47].c
	v[48] = p.nodes[48].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[49] = chainPure(OpCode(21), v[30], v[32])
	v[50] = p.nodes[50].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[51] = p.nodes[51].c
	v[52] = p.nodes[52].c
	v[53] = p.nodes[53].c
	v[54].And(&v[27], &v[53])
	v[55] = p.nodes[55].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[57] = p.nodes[57].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[58].Add(&v[34], &v[56])
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[60].Sub(&v[58], &v[59])
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[61] = p.nodes[61].c
	v[62] = p.nodes[62].c
	v[63].And(&v[26], &v[53])
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
	v[73] = p.nodes[73].c
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
	chainFastDirect["4cdd4ae2e20aeefa279ab2cf21a0e62b1aa31d5840b99b4a928feac82fa67e53"] = admissionFastShape5
}

func admissionFastShape6(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[35] = p.nodes[35].c
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
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[42] = p.nodes[42].c
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	v[45] = p.nodes[45].c
	v[46].And(&v[28], &v[45])
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
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[53] = p.nodes[53].c
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
	v[59] = p.nodes[59].c
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
	v[63] = p.nodes[63].c
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
	v[66] = p.nodes[66].c
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
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["4f1ae1ff12268c3fa272db2f33343d98a0afed1c257628db0cd3a6aadab62682"] = admissionFastShape6
}

func admissionFastShape7(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[34].And(&v[29], &v[32])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[37] = p.nodes[37].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[39].SetBytes(effectiveCaller(s))
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
	v[42] = p.nodes[42].c
	v[43] = chainPure(OpCode(20), v[41], v[42])
	v[44] = p.nodes[44].c
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
	v[51] = p.nodes[51].c
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
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[54].And(&v[30], &v[32])
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
	v[62] = p.nodes[62].c
	v[63].Sub(&v[60], &v[61])
	v[64].Add(&v[36], &v[63])
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["507cf7b433540ba6768869402269d11e39ab291fc1007549df3dde9d89b09f17"] = admissionFastShape7
}

func admissionFastShape8(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[35] = p.nodes[35].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[37].SetBytes(effectiveCaller(s))
	v[38].And(&v[29], &v[37])
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[39] = p.nodes[39].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[42].And(&v[39], &v[41])
	v[43] = chainPure(OpCode(21), v[42], v[32])
	v[44] = p.nodes[44].c
	v[45] = p.nodes[45].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[46] = p.nodes[46].c
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
	v[51] = p.nodes[51].c
	v[52].Sub(&v[50], &v[51])
	v[53] = p.nodes[53].c
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[54] = p.nodes[54].c
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
	v[57] = p.nodes[57].c
	v[58] = p.nodes[58].c
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	v[59] = p.nodes[59].c
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
	v[62] = p.nodes[62].c
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
	v[65].And(&v[29], &v[30])
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
	v[68] = p.nodes[68].c
	code, err = p.actionFast(45, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 45 != 0, err
	}
	v[70] = p.nodes[70].c
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
	chainFastDirect["5da1f755815b9fbc0d439f801ab9c3e16a85419db75ef612df660fb21660ae4d"] = admissionFastShape8
}

func admissionFastShape9(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[32].SetBytes(effectiveCaller(s))
	v[33] = p.nodes[33].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[34] = p.nodes[34].c
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
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[39] = chainPure(OpCode(17), v[28], v[38])
	v[40] = chainPure(OpCode(21), v[39], v[33])
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
	v[44] = p.nodes[44].c
	v[45] = p.nodes[45].c
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
	v[49] = p.nodes[49].c
	v[50] = p.nodes[50].c
	v[51] = p.nodes[51].c
	v[52].And(&v[27], &v[51])
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
	v[56] = p.nodes[56].c
	v[57].And(&v[55], &v[56])
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[59] = p.nodes[59].c
	v[60] = chainPure(OpCode(17), v[58], v[56])
	v[61] = chainPure(OpCode(21), v[60], v[33])
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
	v[70] = chainPure(OpCode(21), v[69], v[33])
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
	v[77] = p.nodes[77].c
	v[78].Sub(&v[75], &v[76])
	v[79].Add(&v[35], &v[78])
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
	v[82] = p.nodes[82].c
	v[83].Sub(&v[80], &v[81])
	v[84].Add(&v[35], &v[83])
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["61d57e039c304e3ba044f766bfdb8d981942f3b33c2111ea3843fc850cfec76b"] = admissionFastShape9
}

func admissionFastShape10(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[34].And(&v[27], &v[32])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[37] = p.nodes[37].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[39].SetBytes(effectiveCaller(s))
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
	v[41] = p.nodes[41].c
	v[42] = p.nodes[42].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[43] = p.nodes[43].c
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
	v[46] = p.nodes[46].c
	v[47] = p.nodes[47].c
	v[48] = p.nodes[48].c
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
	v[50].And(&v[28], &v[32])
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
	v[53] = p.nodes[53].c
	v[54] = p.nodes[54].c
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
	v[58] = p.nodes[58].c
	v[59].Sub(&v[56], &v[57])
	v[60].Add(&v[36], &v[59])
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["67818edc0e443721b6acdb117e04847eed52c7685466c14dd22eb96078485f44"] = admissionFastShape10
}

func admissionFastShape11(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[32].SetBytes(effectiveCaller(s))
	v[33] = p.nodes[33].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[34] = p.nodes[34].c
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
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[39] = p.nodes[39].c
	v[40].And(&v[38], &v[39])
	v[41] = chainPure(OpCode(21), v[40], v[33])
	v[42] = p.nodes[42].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[43] = p.nodes[43].c
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
	v[47] = p.nodes[47].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
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
	v[52] = chainPure(OpCode(17), v[48], v[51])
	v[53] = chainPure(OpCode(21), v[52], v[33])
	v[54] = p.nodes[54].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[55] = p.nodes[55].c
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
	v[58].And(&v[39], &v[57])
	v[59] = chainPure(OpCode(21), v[58], v[33])
	v[60] = p.nodes[60].c
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
	v[62] = p.nodes[62].c
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
	v[64] = p.nodes[64].c
	v[65] = p.nodes[65].c
	v[66] = p.nodes[66].c
	v[67].And(&v[30], &v[66])
	v[68] = p.nodes[68].c
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	v[69] = p.nodes[69].c
	v[70] = p.nodes[70].c
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
	v[74] = p.nodes[74].c
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
	v[77] = p.nodes[77].c
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
	v[80] = p.nodes[80].c
	v[81] = p.nodes[81].c
	v[82] = p.nodes[82].c
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	v[83].Add(&v[35], &v[79])
	code, err = p.actionFast(44, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 44 != 0, err
	}
	v[85].Sub(&v[83], &v[84])
	code, err = p.actionFast(45, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 45 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["69783b09f406c9a971f31fcc432ea8502a580c2dab773ae59cdc8b830a3975da"] = admissionFastShape11
}

func admissionFastShape12(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[35] = p.nodes[35].c
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
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[43].Add(&v[35], &v[42])
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[44] = p.nodes[44].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[45].Add(&v[34], &v[42])
	v[46] = p.nodes[46].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[47] = p.nodes[47].c
	v[48] = chainPure(OpCode(17), v[30], v[37])
	v[49] = chainPure(OpCode(21), v[48], v[32])
	v[50] = p.nodes[50].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[51].Sub(&v[37], &v[30])
	v[52] = p.nodes[52].c
	v[53] = p.nodes[53].c
	v[54] = p.nodes[54].c
	v[55] = p.nodes[55].c
	v[56].And(&v[28], &v[55])
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
	v[58] = p.nodes[58].c
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
	v[62] = p.nodes[62].c
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[64].Add(&v[35], &v[63])
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[65].Add(&v[34], &v[63])
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[66] = chainPure(OpCode(17), v[30], v[61])
	v[67] = chainPure(OpCode(21), v[66], v[32])
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[68].Sub(&v[61], &v[30])
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
	v[69] = p.nodes[69].c
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
	v[72] = p.nodes[72].c
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	v[74].Add(&v[35], &v[73])
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	v[75].Add(&v[34], &v[73])
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	v[76] = chainPure(OpCode(17), v[30], v[71])
	v[77] = chainPure(OpCode(21), v[76], v[32])
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	v[78].Sub(&v[71], &v[30])
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
	v[79] = chainPure(OpCode(21), v[78], v[32])
	v[80] = chainPure(OpCode(21), v[79], v[32])
	v[81] = p.nodes[81].c
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	v[82] = p.nodes[82].c
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
	code, err = p.actionFast(49, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 49 != 0, err
	}
	code, err = p.actionFast(50, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 50 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["767f81b6ccc084d0f0bac79a326ee45398d0307ae2c1dfc1ff102daaa22aa0cc"] = admissionFastShape12
}

func admissionFastShape13(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[32].SetBytes(effectiveCaller(s))
	v[33] = p.nodes[33].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[34] = p.nodes[34].c
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
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[39] = p.nodes[39].c
	v[40].And(&v[38], &v[39])
	v[41] = p.nodes[41].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[42] = p.nodes[42].c
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
	v[45].And(&v[39], &v[44])
	v[46] = chainPure(OpCode(21), v[45], v[33])
	v[47] = p.nodes[47].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[48] = p.nodes[48].c
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
	v[51] = p.nodes[51].c
	v[52] = p.nodes[52].c
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
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[57] = chainPure(OpCode(17), v[53], v[56])
	v[58] = chainPure(OpCode(21), v[57], v[33])
	v[59] = p.nodes[59].c
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
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[63].And(&v[39], &v[62])
	v[64] = chainPure(OpCode(21), v[63], v[33])
	v[65] = p.nodes[65].c
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
	v[67] = p.nodes[67].c
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
	v[69] = p.nodes[69].c
	v[70] = p.nodes[70].c
	v[71] = p.nodes[71].c
	v[72].And(&v[30], &v[71])
	v[73] = p.nodes[73].c
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	v[74] = p.nodes[74].c
	v[75] = p.nodes[75].c
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	v[77] = p.nodes[77].c
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
	v[79] = p.nodes[79].c
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
	v[82] = p.nodes[82].c
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
	code, err = p.actionFast(47, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 47 != 0, err
	}
	v[85] = p.nodes[85].c
	v[86] = p.nodes[86].c
	v[87] = p.nodes[87].c
	code, err = p.actionFast(48, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 48 != 0, err
	}
	v[88].Add(&v[35], &v[84])
	code, err = p.actionFast(49, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 49 != 0, err
	}
	v[90].Sub(&v[88], &v[89])
	code, err = p.actionFast(50, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 50 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["79f21006eb8b0b7bae5e0279798c6a6cd5420c11b065b40eb652d9f6ed016c49"] = admissionFastShape13
}

func admissionFastShape14(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[32].SetBytes(effectiveCaller(s))
	v[33] = p.nodes[33].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[34] = p.nodes[34].c
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
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	v[41].And(&v[38], &v[40])
	v[42] = p.nodes[42].c
	v[43].Or(&v[41], &v[42])
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
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[52] = p.nodes[52].c
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
	v[55] = p.nodes[55].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[57] = chainPure(OpCode(20), v[46], v[56])
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
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[60].And(&v[40], &v[59])
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["7f29430681af74c45a51cef46e96ac6b2a9b486a10b2c8bc9fb8894a9f4ab21d"] = admissionFastShape14
}

func admissionFastShape15(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[35] = p.nodes[35].c
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
	v[41].Div(&v[37], &v[40])
	v[42] = p.nodes[42].c
	v[43].And(&v[41], &v[42])
	v[44] = chainPure(OpCode(21), v[43], v[32])
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
	v[47] = p.nodes[47].c
	v[48] = p.nodes[48].c
	v[49] = p.nodes[49].c
	v[50].And(&v[27], &v[49])
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
	v[53].Div(&v[52], &v[40])
	v[54].And(&v[42], &v[53])
	v[55] = chainPure(OpCode(21), v[54], v[32])
	v[56] = p.nodes[56].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[57] = p.nodes[57].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[58] = p.nodes[58].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[60] = p.nodes[60].c
	v[61] = p.nodes[61].c
	v[62] = p.nodes[62].c
	v[63].Div(&v[59], &v[62])
	v[64] = p.nodes[64].c
	v[65].And(&v[63], &v[64])
	v[66] = p.nodes[66].c
	v[67] = p.nodes[67].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[69] = p.nodes[69].c
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
	v[72] = p.nodes[72].c
	v[73] = p.nodes[73].c
	code, err = p.actionFast(18, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 18 != 0, err
	}
	v[75] = chainPure(OpCode(17), v[74], v[68])
	v[76] = chainPure(OpCode(21), v[75], v[32])
	v[77] = p.nodes[77].c
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[78] = p.nodes[78].c
	v[79] = p.nodes[79].c
	v[80] = p.nodes[80].c
	v[81] = p.nodes[81].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[82] = chainPure(OpCode(21), v[30], v[32])
	v[83] = p.nodes[83].c
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[84] = p.nodes[84].c
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[86] = p.nodes[86].c
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[87].Add(&v[34], &v[85])
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[89].Sub(&v[87], &v[88])
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	v[90] = p.nodes[90].c
	v[91] = p.nodes[91].c
	v[92].And(&v[26], &v[49])
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
	v[95] = p.nodes[95].c
	v[96] = p.nodes[96].c
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
	v[98] = p.nodes[98].c
	v[99] = p.nodes[99].c
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
	code, err = p.actionFast(43, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 43 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["87776607645cce24c14bd1546528db516dd1f3c32eb5e03a01e0c3e394ef8451"] = admissionFastShape15
}

func admissionFastShape16(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[32].SetBytes(effectiveCaller(s))
	v[33] = p.nodes[33].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[34] = p.nodes[34].c
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
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[39] = p.nodes[39].c
	v[40].And(&v[38], &v[39])
	v[41] = chainPure(OpCode(21), v[40], v[33])
	v[42] = p.nodes[42].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[43] = p.nodes[43].c
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
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[49] = p.nodes[49].c
	v[50] = p.nodes[50].c
	v[51] = p.nodes[51].c
	v[52].And(&v[29], &v[51])
	v[53] = p.nodes[53].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[54] = p.nodes[54].c
	v[55] = p.nodes[55].c
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
	v[59] = p.nodes[59].c
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
	v[62] = p.nodes[62].c
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
	v[65] = p.nodes[65].c
	v[66] = p.nodes[66].c
	v[67] = p.nodes[67].c
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	v[68].Add(&v[35], &v[64])
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	v[70].Sub(&v[68], &v[69])
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["8d3cf756963c11097ef90f35e1198a3cd57cc0d70654edf491222d55f7a265f5"] = admissionFastShape16
}

func admissionFastShape17(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[34].And(&v[25], &v[32])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[37] = p.nodes[37].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[39].SetBytes(effectiveCaller(s))
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
	v[42] = chainPure(OpCode(17), v[27], v[41])
	v[43] = chainPure(OpCode(21), v[42], v[33])
	v[44] = p.nodes[44].c
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
	v[48].And(&v[26], &v[32])
	v[49] = p.nodes[49].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[50] = p.nodes[50].c
	v[51] = p.nodes[51].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[52] = p.nodes[52].c
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
	v[55] = p.nodes[55].c
	v[56].And(&v[54], &v[55])
	v[57] = chainPure(OpCode(17), v[27], v[56])
	v[58] = chainPure(OpCode(21), v[57], v[33])
	v[59] = p.nodes[59].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[60] = p.nodes[60].c
	v[61] = p.nodes[61].c
	v[62] = p.nodes[62].c
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
	v[64].And(&v[55], &v[63])
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[66] = p.nodes[66].c
	v[67] = chainPure(OpCode(17), v[65], v[55])
	v[68] = chainPure(OpCode(21), v[67], v[33])
	v[69] = p.nodes[69].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[70] = p.nodes[70].c
	v[71] = p.nodes[71].c
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
	v[73] = p.nodes[73].c
	v[74] = chainPure(OpCode(28), v[73], v[72])
	v[75] = p.nodes[75].c
	v[76] = chainPure(OpCode(20), v[74], v[75])
	v[77] = chainPure(OpCode(21), v[76], v[33])
	v[78] = p.nodes[78].c
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
	v[79] = p.nodes[79].c
	v[80] = p.nodes[80].c
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
	v[83].And(&v[55], &v[82])
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	v[85] = chainPure(OpCode(17), v[84], v[55])
	v[86] = chainPure(OpCode(21), v[85], v[33])
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
	v[88] = chainPure(OpCode(28), v[73], v[87])
	v[89] = chainPure(OpCode(20), v[75], v[88])
	v[90] = chainPure(OpCode(21), v[89], v[33])
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
	v[91] = p.nodes[91].c
	code, err = p.actionFast(45, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 45 != 0, err
	}
	code, err = p.actionFast(46, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 46 != 0, err
	}
	v[93].Add(&v[36], &v[92])
	code, err = p.actionFast(47, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 47 != 0, err
	}
	v[95].Sub(&v[93], &v[94])
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
	v[97] = p.nodes[97].c
	v[98] = p.nodes[98].c
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
	code, err = p.actionFast(62, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 62 != 0, err
	}
	code, err = p.actionFast(63, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 63 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["8fc4ac512f8b4d08be0aa8d9d17dbb0b64bd500932f011d732a881ee07a23d57"] = admissionFastShape17
}

func admissionFastShape18(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[34].And(&v[29], &v[32])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[37] = p.nodes[37].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[40] = p.nodes[40].c
	v[41] = p.nodes[41].c
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
	v[44].And(&v[30], &v[32])
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
	v[47] = p.nodes[47].c
	v[48] = p.nodes[48].c
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
	v[53].Sub(&v[50], &v[51])
	v[54].Add(&v[36], &v[53])
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["90380cd12b71148223664821cd3b4216f7883036335f527e411341e186fae81a"] = admissionFastShape18
}

func admissionFastShape19(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[34].And(&v[28], &v[32])
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[35] = p.nodes[35].c
	v[36] = p.nodes[36].c
	code, err = p.actionFast(1, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 1 != 0, err
	}
	v[37] = p.nodes[37].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[39].SetBytes(effectiveCaller(s))
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
	v[42] = p.nodes[42].c
	v[43] = p.nodes[43].c
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
	v[46] = p.nodes[46].c
	v[47] = p.nodes[47].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[48] = p.nodes[48].c
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
	v[53].And(&v[29], &v[32])
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
	v[56] = p.nodes[56].c
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
	v[61] = p.nodes[61].c
	v[62].Sub(&v[59], &v[60])
	v[63].Add(&v[36], &v[62])
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["961bb7dac3bbf8cce02b53cef5b62e052382adc0566716b4c676e5a8ac7a2e84"] = admissionFastShape19
}

func admissionFastShape20(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[32].SetBytes(effectiveCaller(s))
	v[33] = p.nodes[33].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[34] = p.nodes[34].c
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
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[39] = p.nodes[39].c
	v[40].And(&v[38], &v[39])
	v[41] = chainPure(OpCode(21), v[40], v[33])
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
	v[45].And(&v[43], &v[44])
	v[46] = p.nodes[46].c
	v[47].Or(&v[45], &v[46])
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
	v[54] = p.nodes[54].c
	v[55] = p.nodes[55].c
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
	v[58] = p.nodes[58].c
	v[59] = p.nodes[59].c
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[61] = chainPure(OpCode(20), v[50], v[60])
	v[62] = p.nodes[62].c
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
	v[64].And(&v[44], &v[63])
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["9989441a18dbfdb7c19dae9ac0c690b382e48106cf8e10b1fc153d9e195ef8d6"] = admissionFastShape20
}

func admissionFastShape21(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[34] = p.nodes[34].c
	v[35] = p.nodes[35].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[38].And(&v[34], &v[37])
	v[39] = chainPure(OpCode(21), v[38], v[32])
	v[40] = p.nodes[40].c
	v[41] = p.nodes[41].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[42] = p.nodes[42].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[44].SetBytes(effectiveCaller(s))
	v[45].And(&v[28], &v[44])
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
	v[49] = p.nodes[49].c
	v[50].Sub(&v[48], &v[49])
	v[51] = p.nodes[51].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[52] = p.nodes[52].c
	v[53] = p.nodes[53].c
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
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	v[55] = p.nodes[55].c
	v[56] = p.nodes[56].c
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	v[57] = p.nodes[57].c
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
	v[60] = p.nodes[60].c
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
	v[63].And(&v[28], &v[29])
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
	v[66] = p.nodes[66].c
	code, err = p.actionFast(42, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 42 != 0, err
	}
	v[68] = p.nodes[68].c
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
	chainFastDirect["9b327d568887b44f9526896a704374053654552f5a7eaa43519251b8d6c8baf4"] = admissionFastShape21
}

func admissionFastShape22(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[35] = p.nodes[35].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[37].SetBytes(effectiveCaller(s))
	v[38].And(&v[30], &v[37])
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
	v[42] = p.nodes[42].c
	v[43].Sub(&v[41], &v[42])
	v[44] = p.nodes[44].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[45] = p.nodes[45].c
	v[46] = p.nodes[46].c
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
	v[48] = p.nodes[48].c
	v[49] = p.nodes[49].c
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[50] = p.nodes[50].c
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
	v[53] = p.nodes[53].c
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
	v[56].And(&v[30], &v[31])
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
	v[59] = p.nodes[59].c
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	v[61] = p.nodes[61].c
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
	chainFastDirect["9bf7242e0c5834d117710032b9f8c6eaa094e0a89a49c77cb9c70c7adb3335ba"] = admissionFastShape22
}

func admissionFastShape23(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[35] = p.nodes[35].c
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
	v[39] = p.nodes[39].c
	v[40] = chainPure(OpCode(17), v[27], v[38])
	v[41] = chainPure(OpCode(21), v[40], v[32])
	v[42] = p.nodes[42].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[46] = chainPure(OpCode(16), v[45], v[35])
	v[47] = chainPure(OpCode(21), v[46], v[32])
	v[48] = p.nodes[48].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[49].Add(&v[28], &v[34])
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[51] = chainPure(OpCode(16), v[50], v[35])
	v[52] = chainPure(OpCode(21), v[51], v[32])
	v[53] = p.nodes[53].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[55].Add(&v[34], &v[49])
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[57] = p.nodes[57].c
	v[58] = chainPure(OpCode(20), v[54], v[57])
	v[59] = chainPure(OpCode(21), v[58], v[32])
	v[60] = p.nodes[60].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[61] = p.nodes[61].c
	v[62] = p.nodes[62].c
	v[63] = p.nodes[63].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[64] = p.nodes[64].c
	v[65] = p.nodes[65].c
	v[66] = p.nodes[66].c
	v[67] = p.nodes[67].c
	v[68].And(&v[25], &v[67])
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[69] = p.nodes[69].c
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
	v[72] = p.nodes[72].c
	v[73] = p.nodes[73].c
	v[74] = p.nodes[74].c
	v[75] = p.nodes[75].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[77].Add(&v[35], &v[76])
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[78] = p.nodes[78].c
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[79].Add(&v[34], &v[76])
	v[80] = p.nodes[80].c
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[81] = p.nodes[81].c
	v[82] = chainPure(OpCode(17), v[27], v[71])
	v[83] = chainPure(OpCode(21), v[82], v[32])
	v[84] = p.nodes[84].c
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[85].Sub(&v[71], &v[27])
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
	v[89] = p.nodes[89].c
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	v[91].Add(&v[35], &v[90])
	code, err = p.actionFast(36, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 36 != 0, err
	}
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	v[92].Add(&v[34], &v[90])
	code, err = p.actionFast(38, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 38 != 0, err
	}
	v[93] = chainPure(OpCode(17), v[27], v[88])
	v[94] = chainPure(OpCode(21), v[93], v[32])
	code, err = p.actionFast(39, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 39 != 0, err
	}
	v[95].Sub(&v[88], &v[27])
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
	v[96] = p.nodes[96].c
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
	v[101].Add(&v[35], &v[100])
	code, err = p.actionFast(51, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 51 != 0, err
	}
	code, err = p.actionFast(52, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 52 != 0, err
	}
	v[102].Add(&v[34], &v[100])
	code, err = p.actionFast(53, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 53 != 0, err
	}
	v[103] = chainPure(OpCode(17), v[27], v[98])
	v[104] = chainPure(OpCode(21), v[103], v[32])
	code, err = p.actionFast(54, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 54 != 0, err
	}
	v[105].Sub(&v[98], &v[27])
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
	v[106] = chainPure(OpCode(21), v[105], v[32])
	v[107] = chainPure(OpCode(21), v[106], v[32])
	v[108] = p.nodes[108].c
	code, err = p.actionFast(59, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 59 != 0, err
	}
	v[109] = p.nodes[109].c
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
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["a2c30b47295681491550c2e1bf7059e61f56ca9b8359269f75f1d315c6f44773"] = admissionFastShape23
}

func admissionFastShape24(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[35] = p.nodes[35].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[37] = p.nodes[37].c
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
	v[41] = p.nodes[41].c
	v[42] = p.nodes[42].c
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[46].Add(&v[35], &v[45])
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[47] = p.nodes[47].c
	code, err = p.actionFast(12, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 12 != 0, err
	}
	v[48].Add(&v[34], &v[45])
	v[49] = p.nodes[49].c
	code, err = p.actionFast(13, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 13 != 0, err
	}
	v[50] = p.nodes[50].c
	v[51] = chainPure(OpCode(17), v[29], v[40])
	v[52] = chainPure(OpCode(21), v[51], v[32])
	v[53] = p.nodes[53].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[54].Sub(&v[40], &v[29])
	v[55] = p.nodes[55].c
	v[56] = p.nodes[56].c
	v[57] = p.nodes[57].c
	v[58] = p.nodes[58].c
	v[59].And(&v[27], &v[58])
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
	v[62] = p.nodes[62].c
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
	v[65] = p.nodes[65].c
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	v[67].Add(&v[35], &v[66])
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	v[68].Add(&v[34], &v[66])
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	v[69] = chainPure(OpCode(17), v[29], v[64])
	v[70] = chainPure(OpCode(21), v[69], v[32])
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	v[71].Sub(&v[64], &v[29])
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
	v[72] = chainPure(OpCode(21), v[71], v[32])
	v[73] = chainPure(OpCode(21), v[72], v[32])
	v[74] = p.nodes[74].c
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
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["ab771b4cf29c5bb44e1adc722ebe4d7eacb79120002fe7dba32fe252690ab9c0"] = admissionFastShape24
}

func admissionFastShape25(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[35] = p.nodes[35].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[37].SetBytes(effectiveCaller(s))
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
	v[40] = p.nodes[40].c
	v[41] = p.nodes[41].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	v[45] = p.nodes[45].c
	v[46] = p.nodes[46].c
	v[47].And(&v[27], &v[46])
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
	v[50] = p.nodes[50].c
	v[51] = p.nodes[51].c
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
	v[54] = p.nodes[54].c
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
	v[56].And(&v[28], &v[46])
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
	v[59] = p.nodes[59].c
	v[60] = p.nodes[60].c
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
	v[64] = p.nodes[64].c
	v[65].Sub(&v[62], &v[63])
	v[66].Add(&v[34], &v[65])
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["afa03c7331b0b33192808d3ba555edb03bdc3d0f2eaf7f92c6a46ed9de92c5f3"] = admissionFastShape25
}

func admissionFastShape26(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[35] = p.nodes[35].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[38] = p.nodes[38].c
	v[39].And(&v[37], &v[38])
	v[40] = chainPure(OpCode(21), v[39], v[32])
	v[41] = p.nodes[41].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[42] = p.nodes[42].c
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	v[45] = p.nodes[45].c
	v[46].And(&v[29], &v[45])
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
	v[50].And(&v[48], &v[49])
	v[51].Or(&v[42], &v[50])
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[52].SetBytes(effectiveCaller(s))
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[53] = p.nodes[53].c
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
	v[56].And(&v[38], &v[55])
	v[57] = p.nodes[57].c
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
	v[61] = chainPure(OpCode(17), v[30], v[60])
	v[62] = chainPure(OpCode(21), v[61], v[32])
	v[63] = p.nodes[63].c
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
	v[65] = p.nodes[65].c
	v[66] = p.nodes[66].c
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
	v[69].And(&v[49], &v[68])
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["b00c42215570cad51d6d15754b0cbe10a18a8920b6abc7e2049fee9e6bf6769b"] = admissionFastShape26
}

func admissionFastShape27(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[35] = p.nodes[35].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[38] = p.nodes[38].c
	v[39].And(&v[37], &v[38])
	v[40] = chainPure(OpCode(21), v[39], v[32])
	v[41] = chainPure(OpCode(21), v[40], v[32])
	v[42] = p.nodes[42].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	v[45] = p.nodes[45].c
	v[46] = p.nodes[46].c
	v[47] = p.nodes[47].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[48] = chainPure(OpCode(21), v[30], v[32])
	v[49] = p.nodes[49].c
	code, err = p.actionFast(6, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 6 != 0, err
	}
	v[50] = p.nodes[50].c
	v[51] = p.nodes[51].c
	v[52] = p.nodes[52].c
	v[53].And(&v[27], &v[52])
	v[54] = p.nodes[54].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[56] = p.nodes[56].c
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[57].Add(&v[34], &v[55])
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[59].Sub(&v[57], &v[58])
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[60] = p.nodes[60].c
	v[61] = p.nodes[61].c
	v[62].And(&v[26], &v[52])
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
	v[66] = p.nodes[66].c
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
	v[68] = p.nodes[68].c
	v[69] = p.nodes[69].c
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
	v[72] = p.nodes[72].c
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
	chainFastDirect["b99d6ab34c4987c7ae9d26405b73de2c23dd0b64261c6b276e450f9f3343779c"] = admissionFastShape27
}

func admissionFastShape28(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[32].SetBytes(effectiveCaller(s))
	v[33] = p.nodes[33].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[34] = p.nodes[34].c
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
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[39] = chainPure(OpCode(17), v[31], v[38])
	v[40] = chainPure(OpCode(21), v[39], v[33])
	v[41] = p.nodes[41].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[42] = p.nodes[42].c
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
	v[46].And(&v[44], &v[45])
	v[47] = chainPure(OpCode(21), v[46], v[33])
	v[48] = p.nodes[48].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
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
	v[56] = p.nodes[56].c
	v[57] = p.nodes[57].c
	v[58].And(&v[28], &v[57])
	v[59] = p.nodes[59].c
	code, err = p.actionFast(19, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 19 != 0, err
	}
	v[60] = p.nodes[60].c
	v[61] = p.nodes[61].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[63] = p.nodes[63].c
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
	v[65] = p.nodes[65].c
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
	v[68] = p.nodes[68].c
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
	v[71] = p.nodes[71].c
	v[72] = p.nodes[72].c
	v[73] = p.nodes[73].c
	code, err = p.actionFast(33, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 33 != 0, err
	}
	v[74].Add(&v[35], &v[70])
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	v[76].Sub(&v[74], &v[75])
	code, err = p.actionFast(35, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 35 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["b9b9377f4b3b735aa7cb89dddf3c9d492f11580b72293143e01b2d86f36d0b1e"] = admissionFastShape28
}

func admissionFastShape29(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[32].SetBytes(effectiveCaller(s))
	v[33] = p.nodes[33].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[34] = p.nodes[34].c
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
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
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
	v[42] = p.nodes[42].c
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	v[45].And(&v[29], &v[44])
	v[46] = p.nodes[46].c
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
	v[50] = p.nodes[50].c
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
	v[55] = p.nodes[55].c
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
	v[58] = p.nodes[58].c
	v[59] = p.nodes[59].c
	v[60] = p.nodes[60].c
	code, err = p.actionFast(23, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 23 != 0, err
	}
	v[61].Add(&v[35], &v[57])
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[63].Sub(&v[61], &v[62])
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["c3d27171b5e6cad528a01f7222fa91f1960885f66270dbf0cef024c760075356"] = admissionFastShape29
}

func admissionFastShape30(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[32].SetBytes(effectiveCaller(s))
	v[33] = p.nodes[33].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[34] = p.nodes[34].c
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
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[39] = p.nodes[39].c
	v[40].And(&v[38], &v[39])
	v[41] = p.nodes[41].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[42] = p.nodes[42].c
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	v[45].And(&v[32], &v[44])
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
	code, err = p.actionFast(8, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 8 != 0, err
	}
	v[49] = chainPure(OpCode(28), v[39], v[48])
	v[50] = p.nodes[50].c
	v[51] = chainPure(OpCode(20), v[49], v[50])
	v[52] = chainPure(OpCode(21), v[51], v[33])
	v[53] = p.nodes[53].c
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
	v[57] = p.nodes[57].c
	v[58].And(&v[56], &v[57])
	v[59] = chainPure(OpCode(17), v[31], v[33])
	v[60] = p.nodes[60].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[61] = chainPure(OpCode(16), v[58], v[31])
	v[62] = chainPure(OpCode(21), v[61], v[33])
	v[63] = p.nodes[63].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[64] = p.nodes[64].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[66] = p.nodes[66].c
	v[67] = p.nodes[67].c
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
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
	v[72] = p.nodes[72].c
	v[73] = chainPure(OpCode(17), v[71], v[57])
	v[74] = chainPure(OpCode(21), v[73], v[33])
	v[75] = p.nodes[75].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	v[76] = p.nodes[76].c
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
	v[78] = chainPure(OpCode(28), v[39], v[77])
	v[79] = chainPure(OpCode(20), v[50], v[78])
	v[80] = chainPure(OpCode(21), v[79], v[33])
	v[81] = p.nodes[81].c
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
	v[84] = p.nodes[84].c
	v[85].Sub(&v[82], &v[83])
	v[86].Add(&v[35], &v[85])
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
	v[89] = p.nodes[89].c
	v[90].Sub(&v[87], &v[88])
	v[91].Add(&v[35], &v[90])
	code, err = p.actionFast(37, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 37 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["c696aeae91614310ce0058d773c79829c3ac16fc5360b54fe67ef9c0dfc4435a"] = admissionFastShape30
}

func admissionFastShape31(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[32].SetBytes(effectiveCaller(s))
	v[33] = p.nodes[33].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[34] = p.nodes[34].c
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
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[39] = chainPure(OpCode(17), v[31], v[38])
	v[40] = chainPure(OpCode(21), v[39], v[33])
	v[41] = p.nodes[41].c
	code, err = p.actionFast(4, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 4 != 0, err
	}
	v[42] = p.nodes[42].c
	v[43] = chainPure(OpCode(17), v[31], v[42])
	v[44] = chainPure(OpCode(21), v[43], v[33])
	v[45] = p.nodes[45].c
	code, err = p.actionFast(5, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 5 != 0, err
	}
	v[46] = p.nodes[46].c
	v[47] = chainPure(OpCode(16), v[31], v[46])
	v[48] = chainPure(OpCode(21), v[47], v[33])
	v[49] = p.nodes[49].c
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
	v[53] = p.nodes[53].c
	v[54].And(&v[52], &v[53])
	v[55] = chainPure(OpCode(21), v[54], v[33])
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
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[58] = p.nodes[58].c
	v[59].And(&v[57], &v[58])
	v[60] = p.nodes[60].c
	v[61].Or(&v[59], &v[60])
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
	v[64] = p.nodes[64].c
	code, err = p.actionFast(20, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 20 != 0, err
	}
	code, err = p.actionFast(21, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 21 != 0, err
	}
	v[66] = p.nodes[66].c
	v[67] = p.nodes[67].c
	code, err = p.actionFast(22, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 22 != 0, err
	}
	v[68] = p.nodes[68].c
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
	v[70] = p.nodes[70].c
	v[71] = p.nodes[71].c
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	v[73] = chainPure(OpCode(20), v[62], v[72])
	v[74] = p.nodes[74].c
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
	v[76].And(&v[58], &v[75])
	code, err = p.actionFast(34, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 34 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["ce76cfd1d69f398698a61f1f239f286c16d4086c50c4e4bca5fd616756cd8baf"] = admissionFastShape31
}

func admissionFastShape32(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[35] = p.nodes[35].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[37].SetBytes(effectiveCaller(s))
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
	v[39] = p.nodes[39].c
	v[40] = p.nodes[40].c
	v[41] = p.nodes[41].c
	v[42] = p.nodes[42].c
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	v[45].And(&v[26], &v[44])
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
	v[51].And(&v[27], &v[44])
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
	v[54] = p.nodes[54].c
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
	v[59] = p.nodes[59].c
	v[60].Sub(&v[57], &v[58])
	v[61].Add(&v[34], &v[60])
	code, err = p.actionFast(27, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 27 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["e6a8ec3e1155aa8ee9483a50449f0ab832563d1a7e28c6fe98f482c86f4913cf"] = admissionFastShape32
}

func admissionFastShape33(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[32].SetBytes(effectiveCaller(s))
	v[33] = p.nodes[33].c
	code, err = p.actionFast(0, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 0 != 0, err
	}
	v[34] = p.nodes[34].c
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
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[39] = chainPure(OpCode(17), v[28], v[38])
	v[40] = p.nodes[40].c
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
	v[42] = p.nodes[42].c
	v[43] = p.nodes[43].c
	code, err = p.actionFast(9, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 9 != 0, err
	}
	v[45] = chainPure(OpCode(21), v[44], v[33])
	v[46] = p.nodes[46].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[47] = p.nodes[47].c
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[49] = p.nodes[49].c
	v[50].Div(&v[48], &v[49])
	v[51] = p.nodes[51].c
	v[52].And(&v[50], &v[51])
	v[53] = chainPure(OpCode(21), v[52], v[33])
	v[54] = chainPure(OpCode(21), v[53], v[33])
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
	v[57] = p.nodes[57].c
	v[58] = p.nodes[58].c
	code, err = p.actionFast(14, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 14 != 0, err
	}
	v[60] = p.nodes[60].c
	v[61].Div(&v[59], &v[60])
	v[62].And(&v[51], &v[61])
	v[63] = chainPure(OpCode(21), v[62], v[33])
	v[64] = p.nodes[64].c
	code, err = p.actionFast(15, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 15 != 0, err
	}
	v[65] = chainPure(OpCode(21), v[63], v[33])
	v[66] = p.nodes[66].c
	code, err = p.actionFast(16, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 16 != 0, err
	}
	v[67] = p.nodes[67].c
	v[68] = p.nodes[68].c
	v[69] = p.nodes[69].c
	v[70] = p.nodes[70].c
	v[71].And(&v[32], &v[70])
	code, err = p.actionFast(17, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 17 != 0, err
	}
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
	v[75].And(&v[27], &v[70])
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
	v[78].And(&v[51], &v[74])
	v[79].And(&v[51], &v[77])
	v[80] = chainPure(OpCode(21), v[78], v[33])
	v[81] = p.nodes[81].c
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	v[82] = p.nodes[82].c
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	v[83] = chainPure(OpCode(21), v[80], v[33])
	v[84] = p.nodes[84].c
	code, err = p.actionFast(26, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 26 != 0, err
	}
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
	v[89] = p.nodes[89].c
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
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["f1440f62eff86063a28a68f8dea998e82d50769801cac64743a0608832585558"] = admissionFastShape33
}

func admissionFastShape34(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[35] = p.nodes[35].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	code, err = p.actionFast(3, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 3 != 0, err
	}
	v[37] = p.nodes[37].c
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
	v[41] = p.nodes[41].c
	code, err = p.actionFast(10, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 10 != 0, err
	}
	v[42] = p.nodes[42].c
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	v[45] = p.nodes[45].c
	v[46].And(&v[27], &v[45])
	code, err = p.actionFast(11, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 11 != 0, err
	}
	v[47] = p.nodes[47].c
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
	v[51] = p.nodes[51].c
	v[52] = p.nodes[52].c
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
	code, err = p.actionFast(24, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 24 != 0, err
	}
	code, err = p.actionFast(25, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 25 != 0, err
	}
	v[55] = p.nodes[55].c
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
	v[58] = p.nodes[58].c
	code, err = p.actionFast(29, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 29 != 0, err
	}
	code, err = p.actionFast(30, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 30 != 0, err
	}
	v[59] = p.nodes[59].c
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
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["f87b2ca166c7f9825fb2479fa95fe5813fd1ef512e02ac3503549b1e9c5aa99c"] = admissionFastShape34
}

func admissionFastShape35(p *effectChain, pc *uint64, s *ScopeContext, e *EVM) (bool, error) {
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
	v[35] = p.nodes[35].c
	code, err = p.actionFast(2, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 2 != 0, err
	}
	v[37].SetBytes(effectiveCaller(s))
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
	v[40] = p.nodes[40].c
	v[41] = chainPure(OpCode(20), v[39], v[40])
	v[42] = p.nodes[42].c
	code, err = p.actionFast(7, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 7 != 0, err
	}
	v[43] = p.nodes[43].c
	v[44] = p.nodes[44].c
	v[45] = p.nodes[45].c
	v[46] = p.nodes[46].c
	v[47] = p.nodes[47].c
	v[48] = p.nodes[48].c
	v[49].And(&v[28], &v[48])
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
	v[55].And(&v[29], &v[48])
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
	v[58] = p.nodes[58].c
	v[59] = p.nodes[59].c
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
	v[63] = p.nodes[63].c
	v[64].Sub(&v[61], &v[62])
	v[65].Add(&v[34], &v[64])
	code, err = p.actionFast(28, pc, s, e, in, v)
	if code != 0 {
		return *pc != p.entry || 28 != 0, err
	}
	p.finish(pc, s, in, v)
	e.chainProgress = len(p.pcs)
	return true, nil
}
func init() {
	chainFastDirect["f9a92416170ab2a82fb2cd0b24a6dbeee2279466766e5f84b077306ec65b4140"] = admissionFastShape35
}
