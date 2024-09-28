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

// TestDeleteNodeFromListWithData checks if nodes are correctly deleted from the doubly linked list
func TestDeleteNodeFromListWithData(t *testing.T) {
	fmt.Println("\n\n-------------- Test Deletion From a Doubly LinkedList ----------------------------")

	testCases := []struct {
		name           string
		initialInputs  []int
		deleteValue    int
		expectedValues []int // Expected values in order after deletion
	}{
		{
			name:           "Delete from middle of list",
			initialInputs:  []int{1, 2, 3, 4},
			deleteValue:    3,
			expectedValues: []int{1, 2, 4},
		},
		{
			name:           "Delete head of list",
			initialInputs:  []int{1, 2, 3, 4},
			deleteValue:    1,
			expectedValues: []int{2, 3, 4},
		},
		{
			name:           "Delete tail of list",
			initialInputs:  []int{1, 2, 3, 4},
			deleteValue:    4,
			expectedValues: []int{1, 2, 3},
		},
		{
			name:           "Delete non-existent value",
			initialInputs:  []int{1, 2, 3, 4},
			deleteValue:    5,
			expectedValues: []int{1, 2, 3, 4},
		},
	}

	for _, tc := range testCases {
		ll := LinkedList{}

		// Populate the list
		for _, val := range tc.initialInputs {
			ll.appendNodeToListWithData(val)
		}

		// Print the list status before deletion
		ll.displayList()

		// Delete the node
		ll.deleteNodeFromListWithData(tc.deleteValue)

		// Convert list to a slice for comparison
		result := linkedListToSlice(ll)
		fmt.Printf("%s - Delete: %v, Expected: %v, Result: %v", tc.name, tc.deleteValue, tc.expectedValues, result)
		if !reflect.DeepEqual(result, tc.expectedValues) {
			t.Errorf("\n\nTest failed for %s - expected %v, got %v", tc.name, tc.expectedValues, result)
		} else {
			fmt.Printf("    --------- Pass")
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\nOverall Result\n")
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

		// Print the list status before and after appending
		fmt.Printf("Before Appending...")
		ll.displayList()

		// Append elements to the list
		for _, input := range tc.inputs {
			ll.appendNodeToListWithData(input)
		}

		// Convert the list to a slice and compare with the expected result
		result := linkedListToSlice(ll)
		fmt.Printf("%s - Inputs: %v, Expected: %v, Result: %v", tc.name, tc.inputs, tc.expected, result)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("\n\nTest failed for %s - got %v; want %v", tc.name, result, tc.expected)
		} else {
			fmt.Printf("    --------- Pass")
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\nOverall Result\n")
}

// TestPrependToListDoublyLinkedList checks if elements are correctly prepended to the doubly linked list
// TestPrependToListDoublyLinkedList checks if elements are correctly prepended to the doubly linked list
func TestPrependToListDoublyLinkedList(t *testing.T) {
	fmt.Println("\n\n-------------- Test Prepending To a Doubly LinkedList ----------------------------")

	testCases := []struct {
		name     string
		inputs   []int
		expected []int // Expected values in order after appending
	}{
		{
			name:     "Prepend single element",
			inputs:   []int{1},
			expected: []int{1},
		},
		{
			name:     "Prepend multiple elements",
			inputs:   []int{3, 2, 1},
			expected: []int{1, 2, 3},
		},
	}

	for _, tc := range testCases {
		ll := LinkedList{}

		// Print the list status before appending
		fmt.Println("\nBefore Prepending...")
		ll.displayList()

		// Append elements to the list
		for _, input := range tc.inputs {
			ll.prependNodeToListWithData(input)
		}

		// Convert list to a slice for comparison
		result := linkedListToSlice(ll)
		fmt.Printf("%s - Inputs: %v, Expected: %v, Result: %v", tc.name, tc.inputs, tc.expected, result)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("    --------- Test failed for %s - expected %v, got %v", tc.name, tc.expected, result)
		} else {
			fmt.Printf("    --------- Pass")
		}
		fmt.Println()
	}
	fmt.Println("\nOverall Result")
}
