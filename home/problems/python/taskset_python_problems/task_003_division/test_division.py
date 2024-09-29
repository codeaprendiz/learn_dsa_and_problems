import unittest
from division import division

class TestDivisionFunction(unittest.TestCase):
    
    def test_integer_division(self):
        self.assertEqual(division(10, 3)[0], 3)
    
    def test_float_division(self):
        self.assertAlmostEqual(division(10, 3)[1], 10/3, places=6)
    
    def test_division_by_one(self):
        self.assertEqual(division(10, 1), (10, 10.0))
    
    def test_division_by_negative(self):
        self.assertEqual(division(10, -2), (-5, -5.0))
    
    def test_zero_dividend(self):
        self.assertEqual(division(0, 5), (0, 0.0))

    def test_zero_division(self):
        with self.assertRaises(ZeroDivisionError):
            division(10, 0)

if __name__ == '__main__':
    unittest.main()

