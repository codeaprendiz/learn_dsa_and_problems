import unittest

from linked_list import LinkedList

# Assuming the LinkedList and Node classes are defined here or imported from another module
class TestLinkedList(unittest.TestCase):
    
    def test_insert_at_begin_empty_list(self):
        # Test inserting into an empty linked list
        ll = LinkedList()
        ll.insert_at_begin(10)
        self.assertEqual(ll.head.data, 10)  # The head node's data should be 10
        self.assertIsNone(ll.head.next)  # Since there's only one node, its next should be None

    def test_insert_at_begin_single_element(self):
        # Test inserting into a list with one element
        ll = LinkedList()
        ll.insert_at_begin(10)
        ll.insert_at_begin(20)
        self.assertEqual(ll.head.data, 20)  # After insertion, the new head should have data 20
        self.assertEqual(ll.head.next.data, 10)  # The next node should have data 10
        self.assertIsNone(ll.head.next.next)  # The list should end after the second node

    def test_insert_at_begin_multiple_elements(self):
        # Test inserting into a list with multiple elements
        ll = LinkedList()
        ll.insert_at_begin(10)
        ll.insert_at_begin(20)
        ll.insert_at_begin(30)
        self.assertEqual(ll.head.data, 30)  # The head should now have data 30
        self.assertEqual(ll.head.next.data, 20)  # The next node should have data 20
        self.assertEqual(ll.head.next.next.data, 10)  # The last node should have data 10
        self.assertIsNone(ll.head.next.next.next)  # The last node's next should be None

    def test_insert_at_begin_null_values(self):
        # Test inserting null or None values
        ll = LinkedList()
        ll.insert_at_begin(None)
        self.assertIsNone(ll.head.data)  # The head's data should be None
        self.assertIsNone(ll.head.next)  # The list should contain only one node

    def test_insert_at_index_beginning(self):
        # Test inserting at the beginning (index 0)
        ll = LinkedList()
        ll.insert_at_index(10, 0)  # Insert at index 0
        self.assertEqual(ll.head.data, 10)  # The head node's data should be 10
        self.assertIsNone(ll.head.next)  # There should be only one node

    def test_insert_at_index_middle(self):
        # Test inserting in the middle of the list
        ll = LinkedList()
        ll.insert_at_begin(30)
        ll.insert_at_begin(20)
        ll.insert_at_begin(10)
        ll.insert_at_index(25, 2)  # Insert 25 at index 2
        self.assertEqual(ll.head.next.next.data, 25)  # The third node should have data 25
        self.assertEqual(ll.head.next.next.next.data, 30)  # The fourth node should have data 30
        self.assertIsNone(ll.head.next.next.next.next)  # The last node should point to None

    def test_insert_at_index_end(self):
        # Test inserting at the end of the list
        ll = LinkedList()
        ll.insert_at_begin(30)
        ll.insert_at_begin(20)
        ll.insert_at_begin(10)
        ll.insert_at_index(40, 3)  # Insert 40 at the end of the list (index 3)
        self.assertEqual(ll.head.next.next.next.data, 40)  # The last node should have data 40
        self.assertIsNone(ll.head.next.next.next.next)  # The last node should point to None

    def test_insert_at_invalid_index(self):
        # Test inserting at an invalid index (larger than the list size)
        ll = LinkedList()
        ll.insert_at_begin(30)
        ll.insert_at_begin(20)
        ll.insert_at_begin(10)
        with self.assertRaises(IndexError):  # Assuming we modify the code to raise IndexError
            ll.insert_at_index(40, 5)  # Invalid index, out of range

    def test_insert_at_empty_list(self):
        # Test inserting into an empty list at index 1 (invalid)
        ll = LinkedList()
        with self.assertRaises(IndexError):  # Assuming we modify the code to raise IndexError
            ll.insert_at_index(10, 1)  # Invalid index in an empty list

    def test_insert_at_index_single_element(self):
        # Test inserting in a list with one element
        ll = LinkedList()
        ll.insert_at_begin(10)  # Initial list with one element
        ll.insert_at_index(20, 1)  # Insert 20 at index 1 (after 10)
        self.assertEqual(ll.head.next.data, 20)  # Second node should have data 20
        self.assertIsNone(ll.head.next.next)  # The list ends after the second node

if __name__ == '__main__':
    unittest.main()
