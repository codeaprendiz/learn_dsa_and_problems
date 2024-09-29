import unittest
from loops import printsquares

class TestPrintSquaresFunction(unittest.TestCase):
    
    def test_squares(self):
        self.assertEqual(printsquares(5), [0, 1, 4, 9, 16])
    
    def test_empty_case(self):
        self.assertEqual(printsquares(0), [])
    
    def test_single_value(self):
        self.assertEqual(printsquares(1), [0])
    
    def test_large_input(self):
        self.assertEqual(printsquares(10), [0, 1, 4, 9, 16, 25, 36, 49, 64, 81])

if __name__ == '__main__':
    unittest.main()

