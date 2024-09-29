import unittest
from balanced_brackets import isBalanced

class TestIsBalanced(unittest.TestCase):

    def test_balanced_brackets(self):
        # Test a balanced case with multiple types of brackets
        self.assertEqual(isBalanced("{[()]}"), "YES")

    def test_unbalanced_brackets(self):
        # Test an unbalanced case with mismatched brackets
        self.assertEqual(isBalanced("{[(])}"), "NO")

    def test_nested_balanced_brackets(self):
        # Test a more complex balanced case
        self.assertEqual(isBalanced("{{[[(())]]}}"), "YES")

    def test_empty_string(self):
        # Test an empty string, which should be considered balanced
        self.assertEqual(isBalanced(""), "YES")

    def test_unbalanced_missing_closing(self):
        # Test an unbalanced case with missing closing brackets
        self.assertEqual(isBalanced("{[("), "NO")

    def test_unbalanced_missing_opening(self):
        # Test an unbalanced case with missing opening brackets
        self.assertEqual(isBalanced(")]}"), "NO")

    def test_balanced_with_only_parentheses(self):
        # Test a balanced case with only parentheses
        self.assertEqual(isBalanced("()()()"), "YES")

    def test_unbalanced_with_incorrect_order(self):
        # Test an unbalanced case where closing brackets are in the wrong order
        self.assertEqual(isBalanced("([)]"), "NO")

    def test_balanced_with_single_type(self):
        # Test a balanced case with only one type of bracket
        self.assertEqual(isBalanced("[]{}()"), "YES")

if __name__ == '__main__':
    unittest.main()
