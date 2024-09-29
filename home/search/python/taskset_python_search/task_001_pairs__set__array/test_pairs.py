import unittest
from pairs import pairs

class TestPairsFunction(unittest.TestCase):

    def test_case_1(self):
        # Basic test case
        k = 2
        arr = [1, 5, 3, 4, 2]
        result = pairs(k, arr)
        self.assertEqual(result, 3)  # Pairs: (1, 3), (3, 5), (2, 4)

    def test_case_2(self):
        # Test case where k = 1 and all consecutive pairs
        k = 1
        arr = [1, 2, 3, 4]
        result = pairs(k, arr)
        self.assertEqual(result, 3)  # Pairs: (1, 2), (2, 3), (3, 4)

    def test_case_3(self):
        # Test case where k = 3 and no valid pairs
        k = 3
        arr = [1, 2, 3]
        result = pairs(k, arr)
        self.assertEqual(result, 0)  # No pairs

    def test_case_4(self):
        # Test case where k is larger than the largest element difference
        k = 10
        arr = [1, 2, 3, 4, 5]
        result = pairs(k, arr)
        self.assertEqual(result, 0)  # No valid pairs because k is too large

    def test_case_5(self):
        # Test case with larger k
        k = 5
        arr = [10, 15, 20, 25]
        result = pairs(k, arr)
        self.assertEqual(result, 3)  # Pairs: (10, 15), (15, 20), (20, 25)

    def test_case_6(self):
        # Test case with negative numbers
        k = 3
        arr = [-10, -7, -4, -1, 2]
        result = pairs(k, arr)
        self.assertEqual(result, 4)  # Pairs: (-10, -7), (-7, -4), (-4, -1), (-1, 2)

    def test_case_7(self):
        # Test case with a single element
        k = 1
        arr = [1]
        result = pairs(k, arr)
        self.assertEqual(result, 0)  # No pairs possible in a single-element array

if __name__ == '__main__':
    unittest.main()
