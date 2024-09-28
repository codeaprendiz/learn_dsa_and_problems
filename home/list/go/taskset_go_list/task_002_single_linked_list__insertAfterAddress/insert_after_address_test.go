package insert_after_address

import (
	"fmt"
	"reflect"
	"testing"
)

// TestInsertAfterLinkedList checks if a new element is correctly inserted after a specified node
func TestInsertAfterLinkedList(t *testing.T) {
	fmt.Println("\n\n-------------- Test Inserting After a Node in a LinkedList ----------------------------")

	testCases := []struct {
		name          string
		initialInputs []int
		insertAfter   int // Value after which to insert
		toInsert      int // Value to insert
		expected      []int
	}{
		{
			name:          "Insert 5 after first node (1)",
			initialInputs: []int{1, 2, 3, 4},
			insertAfter:   1,
			toInsert:      5,
			expected:      []int{1, 5, 2, 3, 4},
		},
		{
			name:          "Insert 6 after second node (2)",
			initialInputs: []int{1, 2, 3, 4},
			insertAfter:   2,
			toInsert:      6,
			expected:      []int{1, 2, 6, 3, 4},
		},
		{
			name:          "Insert 7 after non-existent node (8)",
			initialInputs: []int{1, 2, 3, 4},
			insertAfter:   8,
			toInsert:      7,
			expected:      []int{1, 2, 3, 4}, // No insertion should happen
		},
	}

	for _, tc := range testCases {
		ll := LinkedList{} // Create a new LinkedList for each test case

		// Populate the list using appendToList
		for _, input := range tc.initialInputs {
			ll.appendToList(input)
		}

		// Find the right node to insert after
		var nodeToInsertAfter *Node
		current := ll.head
		for current != nil && current.val != tc.insertAfter {
			current = current.next
		}
		nodeToInsertAfter = current

		fmt.Printf("\nList status : ")
		ll.displayList()

		// Perform the insertion
		ll.insertAfter(nodeToInsertAfter, tc.toInsert)

		// Convert the list to a slice and compare with the expected result
		result := linkedListToSlice(ll)
		fmt.Printf("      %s - Insert: %v After: %v, Expected: %v, Result: %v", tc.name, tc.toInsert, tc.insertAfter, tc.expected, result)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("\n\nTest failed for %s - got %v; want %v", tc.name, result, tc.expected)
		} else {
			fmt.Printf("    --------- Pass")
		}
	}
	fmt.Printf("\n\n\nOverall Result\n")
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
