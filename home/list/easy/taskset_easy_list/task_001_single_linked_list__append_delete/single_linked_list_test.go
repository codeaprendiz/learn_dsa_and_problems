package single_linked_list

import (
	"fmt"
	"reflect"
	"testing"
)

// TestDeleteLinkedList checks if elements are correctly deleted from the list
func TestDeleteLinkedList(t *testing.T) {
	ll := LinkedList{}
	// Populate the linked list
	inputs := []int{1, 2, 3, 4, 5}
	for _, input := range inputs {
		ll.appendToList(input)
	}

	// Define test cases for deletion
	testCases := []struct {
		name     string
		toDelete int
		expected []int
	}{
		{
			name:     "Delete head (1)",
			toDelete: 1,
			expected: []int{2, 3, 4, 5},
		},
		{
			name:     "Delete middle (3)",
			toDelete: 3,
			expected: []int{2, 4, 5},
		},
		{
			name:     "Delete tail (5)",
			toDelete: 5,
			expected: []int{2, 4},
		},
		{
			name:     "Delete non-existent (6)",
			toDelete: 6,
			expected: []int{2, 4},
		},
	}

	// Execute test cases
	for _, tc := range testCases {
		fmt.Printf("\nList status : ")
		ll.displayList()
		ll.deleteFromList(tc.toDelete)
		result := linkedListToSlice(ll)
		fmt.Printf("          %s - Delete: %v, Expected: %v, Result: %v", tc.name, tc.toDelete, tc.expected, result)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("\n\nTest failed for %s - got %v; want %v", tc.name, result, tc.expected)
		} else {
			fmt.Printf("    --------- Pass")
		}
	}
	fmt.Printf("\n\n")
}

// TestAppendLinkedList checks if a single element is correctly appended to the list
func TestAppendLinkedList(t *testing.T) {
	testCases := []struct {
		name     string
		inputs   []int
		expected []int
	}{
		{
			name:     "Append 1",
			inputs:   []int{1},
			expected: []int{1},
		},
		{
			name:     "Append 1, 2",
			inputs:   []int{1, 2},
			expected: []int{1, 2},
		},
		{
			name:     "Append 1, 2, 3",
			inputs:   []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
	}

	for _, tc := range testCases {
		ll := LinkedList{} // Create a new LinkedList for each test case
		for _, input := range tc.inputs {
			ll.appendToList(input) // Append each element in the inputs slice
		}

		result := linkedListToSlice(ll)
		fmt.Printf("\nList status : ")
		ll.displayList()
		fmt.Printf("      %s - Inputs: %v, Expected: %v, Result: %v", tc.name, tc.inputs, tc.expected, result)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("\n\nTest failed for %s - got %v; want %v", tc.name, result, tc.expected)
		} else {
			fmt.Printf("    --------- Pass")
		}
	}
	fmt.Printf("\nOverall Result\n")
}

// TestPrependLinkedList checks if elements are correctly prepended to the list
func TestPrependLinkedList(t *testing.T) {
	testCases := []struct {
		name     string
		inputs   []int
		expected []int
	}{
		{
			name:     "Prepend 1",
			inputs:   []int{1},
			expected: []int{1},
		},
		{
			name:     "Prepend 2, 1",
			inputs:   []int{2, 1},
			expected: []int{1, 2},
		},
		{
			name:     "Prepend 3, 2, 1",
			inputs:   []int{3, 2, 1},
			expected: []int{1, 2, 3},
		},
	}

	for _, tc := range testCases {
		ll := LinkedList{} // Create a new LinkedList for each test case
		for _, input := range tc.inputs {
			ll.prependToList(input) // Prepend each element in the inputs slice
		}

		result := linkedListToSlice(ll)
		fmt.Printf("\nList status : ")
		ll.displayList()
		fmt.Printf("      %s - Inputs: %v, Expected: %v, Result: %v", tc.name, tc.inputs, tc.expected, result)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("\n\nTest failed for %s - got %v; want %v", tc.name, result, tc.expected)
		} else {
			fmt.Printf("    --------- Pass")
		}
	}
	fmt.Printf("\nOverall Result\n")
}

// Helper function to convert linked list to a slice
func linkedListToSlice(ll LinkedList) []int {
	var slice []int
	current := ll.head
	for current != nil {
		slice = append(slice, current.val)
		current = current.next
	}
	return slice
}
