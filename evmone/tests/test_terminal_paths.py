import unittest
from terminal_paths import exit_only_in_loop,loop_intervals

class TerminalPathsTest(unittest.TestCase):
    def test_exit_only_loop(self):
        # JUMPDEST; PUSH1 0; JUMP; STOP; RETURN; REVERT
        code=bytes.fromhex('5b60005600f3fd')
        self.assertEqual(loop_intervals(code),[(0,3)])
        self.assertTrue(exit_only_in_loop(code,0,[5,6]))
        self.assertFalse(exit_only_in_loop(code,0,[0,6]))
        self.assertFalse(exit_only_in_loop(code,5,[6]))
        self.assertFalse(exit_only_in_loop(code,0,[]))
    def test_not_data_or_dynamic_jump(self):
        self.assertEqual(loop_intervals(bytes.fromhex('605b600056')),[])
        self.assertEqual(loop_intervals(bytes.fromhex('605b600156')),[])
        self.assertEqual(loop_intervals(bytes.fromhex('5b60008056')),[])
    def test_forward_jump_and_truncated_push(self):
        self.assertEqual(loop_intervals(bytes.fromhex('6003565b00')),[])
        self.assertEqual(loop_intervals(bytes.fromhex('5b7f0056')),[])

if __name__=='__main__':unittest.main()
