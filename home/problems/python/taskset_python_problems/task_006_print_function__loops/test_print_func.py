import unittest
from print_func import printfunc

class TestPrintFunc(unittest.TestCase):

    def test_small_number(self):
        self.assertEqual(printfunc(3), [1, 2, 3])

    def test_large_number(self):
        self.assertEqual(printfunc(5), [1, 2, 3, 4, 5])

    def test_edge_case_zero(self):
        self.assertEqual(printfunc(0), [])

    def test_edge_case_one(self):
        self.assertEqual(printfunc(1), [1])

if __name__ == '__main__':
    unittest.main()
