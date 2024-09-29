import unittest
from list_comprehension import list_comprehension

# Unit test class
class TestListComprehension(unittest.TestCase):
    
    def test_case_1(self):
        # Test case where x=1, y=1, z=1, n=2
        x, y, z, n = 1, 1, 1, 2
        expected_output = [[0, 0, 0], [0, 0, 1], [0, 1, 0], [1, 0, 0], [1, 1, 1]]
        self.assertEqual(list_comprehension(x, y, z, n), expected_output)
    
    def test_case_2(self):
        # Test case where x=0, y=0, z=0, n=0 (minimal input)
        x, y, z, n = 0, 0, 0, 0
        expected_output = []
        self.assertEqual(list_comprehension(x, y, z, n), expected_output)

    def test_case_3(self):
        # Test case where x=1, y=1, z=1, n=0 (no filtering, all coordinates included)
        x, y, z, n = 1, 1, 1, 0
        expected_output = [[0, 0, 1], [0, 1, 0], [0, 1, 1], [1, 0, 0], [1, 0, 1], [1, 1, 0], [1, 1, 1]]
        self.assertEqual(list_comprehension(x, y, z, n), expected_output)

    def test_case_4(self):
        # Test case where x=2, y=2, z=2, n=3
        x, y, z, n = 2, 2, 2, 3
 
        expected_output = [[0, 0, 0], [0, 0, 1], [0, 0, 2], [0, 1, 0], [0, 1, 1], [0, 2, 0], [0, 2, 2], [1, 0, 0], [1, 0, 1], [1, 1, 0], [1, 1, 2], [1, 2, 1], [1, 2, 2], [2, 0, 0], [2, 0, 2], [2, 1, 1], [2, 1, 2], [2, 2, 0], [2, 2, 1], [2, 2, 2]]
        self.assertEqual(list_comprehension(x, y, z, n), expected_output)
    
    def test_case_5(self):
        # Test case with large inputs to check performance
        x, y, z, n = 2, 2, 2, 6
        expected_output = [[0, 0, 0], [0, 0, 1], [0, 0, 2], [0, 1, 0], [0, 1, 1], [0, 1, 2], [0, 2, 0], [0, 2, 1], [0, 2, 2], [1, 0, 0], [1, 0, 1], [1, 0, 2], [1, 1, 0], [1, 1, 1], [1, 1, 2], [1, 2, 0], [1, 2, 1], [1, 2, 2], [2, 0, 0], [2, 0, 1], [2, 0, 2], [2, 1, 0], [2, 1, 1], [2, 1, 2], [2, 2, 0], [2, 2, 1]]
        self.assertEqual(list_comprehension(x, y, z, n), expected_output)

if __name__ == '__main__':
    unittest.main()
