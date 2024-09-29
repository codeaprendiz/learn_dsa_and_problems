import unittest
from leap import is_leap

class TestIsLeapFunction(unittest.TestCase):

    def test_divisible_by_400(self):
        self.assertTrue(is_leap(2000))  # 2000 is a leap year
    
    def test_divisible_by_100_not_400(self):
        self.assertFalse(is_leap(1900))  # 1900 is not a leap year
    
    def test_divisible_by_4_not_100(self):
        self.assertTrue(is_leap(2024))  # 2024 is a leap year
    
    def test_not_divisible_by_4(self):
        self.assertFalse(is_leap(2023))  # 2023 is not a leap year
    
    def test_edge_case_year_0(self):
        self.assertTrue(is_leap(0))  # Year 0 is considered a leap year

if __name__ == '__main__':
    unittest.main()

