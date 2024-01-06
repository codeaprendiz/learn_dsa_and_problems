package has_cycle

import (
	"fmt"
	"testing"
)

// TestHasCycle checks if the cycle detection works correctly
func TestHasCycle(t *testing.T) {
	fmt.Println("\n\n-------------- Test Cycle Detection in a LinkedList ----------------------------")

	testCases := []struct {
		name      string
		inputs    []int
		makeCycle bool // Indicates whether to create a cycle in the list
		expected  bool
	}{
		{
			name:      "No cycle",
			inputs:    []int{1, 2, 3, 4},
			makeCycle: false,
			expected:  false,
		},
		{
			name:      "Cycle exists",
			inputs:    []int{1, 2, 3, 4},
			makeCycle: true,
			expected:  true,
		},
	}

	for _, tc := range testCases {
		ll := LinkedList{}

		// Populate the list using appendToList
		for _, input := range tc.inputs {
			ll.appendToList(input)
		}

		// Print the list before making a cycle
		fmt.Printf("\nList before making a cycle (%s): ", tc.name)
		ll.displayList()

		// Create a cycle if required
		if tc.makeCycle {
			// Creating a cycle: Pointing the next of the last node to the head
			lastNode := ll.head
			for lastNode.next != nil {
				lastNode = lastNode.next
			}
			lastNode.next = ll.head // Creating a cycle
		}

		// Test cycle detection
		result := ll.hasCycle()
		fmt.Printf("      %s - MakeCycle: %v, Expected: %v, Result: %v", tc.name, tc.makeCycle, tc.expected, result)
		if result != tc.expected {
			t.Errorf("\n\nTest failed for %s - expected %v, got %v", tc.name, tc.expected, result)
		} else {
			fmt.Printf("    --------- Pass")
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n\n\nOverall Result\n")
}
