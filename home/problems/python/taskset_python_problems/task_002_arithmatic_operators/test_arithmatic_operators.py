import unittest
from arithmatic_operators import checksum

class TestChecksumFunction(unittest.TestCase):
    
    def test_addition(self):
        self.assertEqual(checksum(3, 2)[0], 5)
    
    def test_subtraction(self):
        self.assertEqual(checksum(3, 2)[1], 1)
    
    def test_multiplication(self):
        self.assertEqual(checksum(3, 2)[2], 6)
    
    def test_negative_numbers(self):
        self.assertEqual(checksum(-3, 2), (-1, -5, -6))
    
    def test_zero_values(self):
        self.assertEqual(checksum(0, 0), (0, 0, 0))

if __name__ == '__main__':
    unittest.main()
