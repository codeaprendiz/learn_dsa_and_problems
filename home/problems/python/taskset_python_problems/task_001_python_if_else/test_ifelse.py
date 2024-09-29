import unittest
from ifelse import check_weird

class TestWeirdFunction(unittest.TestCase):
    
    def test_odd_number(self):
        self.assertEqual(check_weird(3), "Weird")
    
    def test_even_in_2_to_5(self):
        self.assertEqual(check_weird(4), "Not Weird")
    
    def test_even_in_6_to_20(self):
        self.assertEqual(check_weird(10), "Weird")
    
    def test_even_greater_than_20(self):
        self.assertEqual(check_weird(22), "Not Weird")
    
    def test_edge_case_2(self):
        self.assertEqual(check_weird(2), "Not Weird")
    
    def test_edge_case_20(self):
        self.assertEqual(check_weird(20), "Weird")
    
    def test_edge_case_21(self):
        self.assertEqual(check_weird(21), "Weird")

if __name__ == '__main__':
    unittest.main()
