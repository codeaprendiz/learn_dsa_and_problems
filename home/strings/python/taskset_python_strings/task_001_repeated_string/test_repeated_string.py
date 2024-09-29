import unittest
from repeated_string import repeatedString

class TestRepeatedString(unittest.TestCase):

    def test_case_1(self):
        # Basic test case where the string is repeated exactly
        s = "abcac"
        n = 10
        result = repeatedString(s, n)
        self.assertEqual(result, 4)  # 'a' appears 4 times in the first 10 characters

    def test_case_2(self):
        # Case where string is just "a"
        s = "a"
        n = 1000000000000
        result = repeatedString(s, n)
        self.assertEqual(result, 1000000000000)  # 'a' appears n times

    def test_case_3(self):
        # Case where the string has no 'a'
        s = "bcdef"
        n = 12
        result = repeatedString(s, n)
        self.assertEqual(result, 0)  # No 'a' in the string

    def test_case_4(self):
        # Case where string length is exactly the same as n
        s = "abc"
        n = 3
        result = repeatedString(s, n)
        self.assertEqual(result, 1)  # 'a' appears once in the full string

    def test_case_5(self):
        # Case where 'a' appears multiple times
        s = "aba"
        n = 10
        result = repeatedString(s, n)
        self.assertEqual(result, 7)  # 'a' appears 7 times in the first 10 characters

    def test_case_6(self):
        # Case where string is repeated many times with remainder
        s = "aab"
        n = 882787
        result = repeatedString(s, n)
        self.assertEqual(result, 588525)  # Verify complex cases with large n

if __name__ == '__main__':
    unittest.main()
