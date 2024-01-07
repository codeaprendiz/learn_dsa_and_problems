package main

import (
	"fmt"
	"reflect"
	"testing"
)

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

// TestAppendToListWithDataDoublyLinkedList checks if elements are correctly appended to the doubly linked list
func TestAppendToListWithDataDoublyLinkedList(t *testing.T) {
	fmt.Println("\n\n-------------- Test Appending To a Doubly LinkedList ----------------------------")

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
			ll.appendNodeToListWithData(input)
		}

		// Print the list status before and after appending
		fmt.Printf("List status : ")
		ll.displayList()

		// Convert the list to a slice and compare with the expected result
		result := linkedListToSlice(ll)
		fmt.Printf("      %s - Inputs: %v, Expected: %v, Result: %v", tc.name, tc.inputs, tc.expected, result)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("\n\nTest failed for %s - got %v; want %v", tc.name, result, tc.expected)
		} else {
			fmt.Printf("    --------- Pass")
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\nOverall Result\n")
}
