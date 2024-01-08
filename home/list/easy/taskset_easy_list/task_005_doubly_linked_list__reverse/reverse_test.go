package main

import (
	"fmt"
	"reflect"
	"testing"
)

// TestReverseLinkedList checks if the linked list is correctly reverse
func TestReverseLinkedList(t *testing.T) {
	testCases := []struct {
		name     string
		inputs   []int
		expected []int // Expected values in reverse order
	}{
		{
			name:     "Reverse single element list",
			inputs:   []int{1},
			expected: []int{1},
		},
		{
			name:     "Reverse multiple element list",
			inputs:   []int{1, 2, 3},
			expected: []int{3, 2, 1},
		},
		{
			name:     "Reverse empty list",
			inputs:   []int{},
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		ll := LinkedList{}

		// Populate the list
		for _, val := range tc.inputs {
			ll.appendNodeToListWithData(val)
		}

		// Reverse the list
		ll.reverse()

		// Convert list to a slice for comparison
		result := linkedListToSlice(ll)

		// Check for empty list edge case
		if len(tc.inputs) == 0 && len(result) == 0 {
			fmt.Printf("%s - Inputs: %v, Expected: %v, Result: %v    --------- Pass\n", tc.name, tc.inputs, tc.expected, result)
			continue
		}

		// Regular comparison
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("    --------- Test failed for %s - expected %v, got %v\n", tc.name, tc.expected, result)
		} else {
			fmt.Printf("%s - Inputs: %v, Expected: %v, Result: %v    --------- Pass\n", tc.name, tc.inputs, tc.expected, result)
		}
	}
	fmt.Println("\nOverall Result")
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
