import unittest
from sherlock_and_anagrams import sherlockAndAnagrams

class TestSherlockAndAnagrams(unittest.TestCase):

    def test_case_1(self):
        # Basic test case
        s = "abba"
        result = sherlockAndAnagrams(s)
        self.assertEqual(result, 4)  # Expected 4 pairs of anagrams

    def test_case_2(self):
        # Test case with no anagrams
        s = "abcd"
        result = sherlockAndAnagrams(s)
        self.assertEqual(result, 0)  # No anagrams

    def test_case_3(self):
        # Test case with all same characters
        s = "kkkk"
        result = sherlockAndAnagrams(s)
        self.assertEqual(result, 10)  # Expected 10 pairs of anagrams

    def test_case_4(self):
        # Test case with multiple possible anagrams
        s = "ifailuhkqq"
        result = sherlockAndAnagrams(s)
        self.assertEqual(result, 3)  # Expected 3 pairs of anagrams

    def test_case_5(self):
        # Edge case: single character string
        s = "a"
        result = sherlockAndAnagrams(s)
        self.assertEqual(result, 0)  # No anagrams with a single character

    def test_case_6(self):
        # Edge case: empty string
        s = ""
        result = sherlockAndAnagrams(s)
        self.assertEqual(result, 0)  # No anagrams in an empty string

    def test_case_7(self):
        # Test case with repeated letters
        s = "mom"
        result = sherlockAndAnagrams(s)
        self.assertEqual(result, 2)  # Expected 2 pairs of anagrams

if __name__ == '__main__':
    unittest.main()
