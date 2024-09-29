import unittest
from runner_up_score import find_runner_up_score  # Make sure to import the function from the appropriate module

class TestFindRunnerUpScore(unittest.TestCase):

    def test_case_1(self):
        # Normal case
        n = 5
        arr = [2, 3, 6, 6, 5]
        expected_output = 5
        self.assertEqual(find_runner_up_score(n, arr), expected_output)

    def test_case_2(self):
        # Case where max score appears multiple times
        n = 6
        arr = [1, 4, 4, 4, 2, 3]
        expected_output = 3
        self.assertEqual(find_runner_up_score(n, arr), expected_output)

    def test_case_3(self):
        # Case where all elements are the same
        n = 4
        arr = [5, 5, 5, 5]
        with self.assertRaises(ValueError):
            find_runner_up_score(n, arr)

    def test_case_4(self):
        # Minimum number of participants
        n = 2
        arr = [1, 2]
        expected_output = 1
        self.assertEqual(find_runner_up_score(n, arr), expected_output)

    def test_case_5(self):
        # Case with negative numbers
        n = 6
        arr = [-1, -2, -3, -4, -5, -6]
        expected_output = -2
        self.assertEqual(find_runner_up_score(n, arr), expected_output)

if __name__ == '__main__':
    unittest.main()
