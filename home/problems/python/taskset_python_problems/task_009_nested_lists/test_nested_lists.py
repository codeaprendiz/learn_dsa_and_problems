import unittest
from io import StringIO
from nested_lists import students_with_second_lowest_grade
import sys


# Test cases for the function
class TestSecondLowest(unittest.TestCase):

    def setUp(self):
        # Save the original stdout (console output)
        self.held_output = StringIO()
        sys.stdout = self.held_output

    def tearDown(self):
        # Reset stdout
        sys.stdout = sys.__stdout__

    def test_case_1(self):
        # Normal case with two students having the second lowest score
        student_list = [['Harry', 37.21], ['Berry', 37.21], ['Tina', 37.2], ['Akriti', 41], ['Harsh', 39]]
        students_with_second_lowest_grade(student_list)
        self.assertEqual(self.held_output.getvalue().strip(), "Berry\nHarry")

    def test_case_2(self):
        # Case with more than two students having the second lowest score
        student_list = [['Harry', 37.21], ['Berry', 37.21], ['Tina', 37.2], ['Akriti', 41], ['Harsh', 37.21]]
        students_with_second_lowest_grade(student_list)
        self.assertEqual(self.held_output.getvalue().strip(), "Berry\nHarry\nHarsh")

    def test_case_3(self):
        # All students have the same score
        student_list = [['Harry', 37.21], ['Berry', 37.21], ['Tina', 37.21]]
        students_with_second_lowest_grade(student_list)
        self.assertEqual(self.held_output.getvalue().strip(), "")

    def test_case_4(self):
        # Only two students with different scores
        student_list = [['Harry', 37.21], ['Berry', 37.22]]
        students_with_second_lowest_grade(student_list)
        self.assertEqual(self.held_output.getvalue().strip(), "Berry")

    def test_case_5(self):
        # Large input case
        student_list = [
            ['Harry', 50], ['Berry', 50], ['Tina', 45], ['Akriti', 45],
            ['Harsh', 45], ['George', 60], ['Emily', 60], ['John', 55]
        ]
        students_with_second_lowest_grade(student_list)
        self.assertEqual(self.held_output.getvalue().strip(), "Berry\nHarry")

if __name__ == '__main__':
    unittest.main()
