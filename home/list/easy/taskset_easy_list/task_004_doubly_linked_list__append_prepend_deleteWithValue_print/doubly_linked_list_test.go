package main

import (
	"fmt"
	"reflect"
	"testing"
)

// TestAppendToListDoublyLinkedList checks if elements are correctly appended to the doubly linked list
func TestAppendToListDoublyLinkedList(t *testing.T) {
	testCases := []struct {
		name     string
		inputs   []int
		expected []int // Expected values in the order they should appear in the list
	}{
		{
			name:     "Append single element",
			inputs:   []int{1},
			expected: []int{1},
		},
		{
			name:     "Append multiple elements",
			inputs:   []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
	}

	for _, tc := range testCases {
		ll := LinkedList{} // Create a new LinkedList for each test case

		// Append elements to the list
		for _, input := range tc.inputs {
			ll.appendToList(input)
		}

		// Convert the list to a slice and compare with the expected result
		result := linkedListToSlice(ll)
		fmt.Printf("Test '%s':\n", tc.name)
		fmt.Printf("    Expected: %v\n", tc.expected)
		fmt.Printf("    Result:   %v\n", result)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Test '%s' failed: expected %v, got %v", tc.name, tc.expected, result)
		}
	}
}

// Helper function to convert doubly linked list to a slice (forward traversal)
func linkedListToSlice(ll LinkedList) []int {
	var slice []int
	current := ll.head
	for current != nil {
		slice = append(slice, current.val)
		current = current.next
	}
	return slice
}
