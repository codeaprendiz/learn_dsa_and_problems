import unittest
from all_substrings import generate_substrings_with_comprehension, generate_substrings_with_loops

class TestSubstringGeneration(unittest.TestCase):

    def test_generate_substrings_with_comprehension(self):
        # Test case 1: Regular string
        test_str = "abc"
        expected_output = ['a', 'ab', 'abc', 'b', 'bc', 'c']
        result = generate_substrings_with_comprehension(test_str)
        self.assertEqual(result, expected_output)

        # Test case 2: Single character
        test_str = "a"
        expected_output = ['a']
        result = generate_substrings_with_comprehension(test_str)
        self.assertEqual(result, expected_output)

        # Test case 3: Empty string
        test_str = ""
        expected_output = []
        result = generate_substrings_with_comprehension(test_str)
        self.assertEqual(result, expected_output)

    def test_generate_substrings_with_loops(self):
        # Test case 1: Regular string
        test_str = "abc"
        expected_output = ['a', 'ab', 'abc', 'b', 'bc', 'c']
        result = generate_substrings_with_loops(test_str)
        self.assertEqual(result, expected_output)

        # Test case 2: Single character
        test_str = "a"
        expected_output = ['a']
        result = generate_substrings_with_loops(test_str)
        self.assertEqual(result, expected_output)

        # Test case 3: Empty string
        test_str = ""
        expected_output = []
        result = generate_substrings_with_loops(test_str)
        self.assertEqual(result, expected_output)


if __name__ == '__main__':
    unittest.main()
