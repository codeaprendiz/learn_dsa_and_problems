import unittest
from bubble_sort import bubble_sort

class TestBubbleSort(unittest.TestCase):
    
    def test_sorted_array(self):
        # Test already sorted array
        arr = [1, 2, 3, 4, 5]
        bubble_sort(arr)
        self.assertEqual(arr, [1, 2, 3, 4, 5])

    def test_reverse_sorted_array(self):
        # Test reverse sorted array
        arr = [5, 4, 3, 2, 1]
        bubble_sort(arr)
        self.assertEqual(arr, [1, 2, 3, 4, 5])
    
    def test_random_unsorted_array(self):
        # Test random unsorted array
        arr = [64, 34, 25, 12, 22, 11, 90]
        bubble_sort(arr)
        self.assertEqual(arr, [11, 12, 22, 25, 34, 64, 90])
    
    def test_array_with_duplicates(self):
        # Test array with duplicate elements
        arr = [3, 5, 2, 3, 8, 1, 2]
        bubble_sort(arr)
        self.assertEqual(arr, [1, 2, 2, 3, 3, 5, 8])

    def test_single_element_array(self):
        # Test array with a single element
        arr = [1]
        bubble_sort(arr)
        self.assertEqual(arr, [1])

    def test_empty_array(self):
        # Test an empty array
        arr = []
        bubble_sort(arr)
        self.assertEqual(arr, [])

if __name__ == '__main__':
    unittest.main()
