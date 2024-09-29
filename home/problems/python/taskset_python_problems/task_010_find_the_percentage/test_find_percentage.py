import unittest
from find_percentage import avg

class TestAvgFunction(unittest.TestCase):
    
    def test_case_1(self):
        # Test case where marks are [67.0, 68.0, 69.0]
        marks = [67.0, 68.0, 69.0]
        self.assertEqual(avg(marks), "68.00")
    
    def test_case_2(self):
        # Test case with marks being all the same
        marks = [100.0, 100.0, 100.0]
        self.assertEqual(avg(marks), "100.00")
    
    def test_case_3(self):
        # Test case with different marks
        marks = [70.0, 98.0, 63.0]
        self.assertEqual(avg(marks), "77.00")
    
    def test_case_4(self):
        # Test case with marks in the decimal range
        marks = [52.5, 56.5, 60.0]
        self.assertEqual(avg(marks), "56.33")
    
    def test_case_5(self):
        # Test case with only one mark
        marks = [75.0]
        self.assertEqual(avg(marks), "75.00")

if __name__ == '__main__':
    unittest.main()
