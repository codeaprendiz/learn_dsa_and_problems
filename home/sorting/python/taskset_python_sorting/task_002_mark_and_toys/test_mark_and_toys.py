import unittest
from mark_and_toys import maximumToys

class TestMaximumToys(unittest.TestCase):

    def test_case_1(self):
        # Basic test case
        prices = [1, 12, 5, 111, 200, 1000, 10]
        k = 50
        result = maximumToys(prices, k)
        self.assertEqual(result, 4)  # The maximum number of toys we can buy is 4 (1, 12, 5, 10)

    def test_case_2(self):
        # Test case where budget is enough to buy all toys
        prices = [1, 2, 3, 4, 5]
        k = 15
        result = maximumToys(prices, k)
        self.assertEqual(result, 5)  # We can buy all the toys

    def test_case_3(self):
        # Test case where the budget is only enough for one toy
        prices = [100, 200, 300]
        k = 150
        result = maximumToys(prices, k)
        self.assertEqual(result, 1)  # We can only buy the first toy (100)

    def test_case_4(self):
        # Test case where the budget is not enough to buy any toy
        prices = [100, 200, 300]
        k = 50
        result = maximumToys(prices, k)
        self.assertEqual(result, 0)  # We can't buy any toys

    def test_case_5(self):
        # Test case with larger numbers
        prices = [10, 20, 30, 40, 50]
        k = 100
        result = maximumToys(prices, k)
        self.assertEqual(result, 4)  # We can buy 4 toys (10, 20, 30, 40)

    def test_case_6(self):
        # Test case with edge case where prices are equal
        prices = [10, 10, 10, 10]
        k = 25
        result = maximumToys(prices, k)
        self.assertEqual(result, 2)  # We can buy only two toys

    def test_case_7(self):
        # Test case with mixed prices and k limiting to small toys
        prices = [1, 1, 1, 1, 10]
        k = 5
        result = maximumToys(prices, k)
        self.assertEqual(result, 4)  # We can buy 4 toys (1, 1, 1, 1)

if __name__ == '__main__':
    unittest.main()
