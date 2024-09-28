package main

import (
	"fmt"
	"reflect"
	"testing"
)

// Helper function to convert BST to a slice (in-order traversal)
func bstToSlice(bst *BST) []int {
	var result []int
	inorderTraversalForSlice(bst.Root, &result)
	return result
}

// Helper function for in-order traversal to a slice
func inorderTraversalForSlice(root *Node, result *[]int) {
	if root != nil {
		inorderTraversalForSlice(root.Left, result)
		*result = append(*result, root.Key)
		inorderTraversalForSlice(root.Right, result)
	}
}

// TestBSTInsert checks if keys are correctly inserted into the BST
func TestBSTInsert(t *testing.T) {
	fmt.Println("\n\n-------------- Test Binary Search Tree Insertion ----------------------------")

	testCases := []struct {
		name           string
		initialInputs  []int
		insertValue    int
		expectedValues []int // Expected values in order after insertion
	}{
		{
			name:           "Insert into an empty tree",
			initialInputs:  []int{},
			insertValue:    10,
			expectedValues: []int{10},
		},
		{
			name:           "Insert smaller value",
			initialInputs:  []int{20, 30, 40},
			insertValue:    15,
			expectedValues: []int{15, 20, 30, 40},
		},
		{
			name:           "Insert larger value",
			initialInputs:  []int{20, 30, 40},
			insertValue:    35,
			expectedValues: []int{20, 30, 35, 40},
		},
		{
			name:           "Insert equal value",
			initialInputs:  []int{20, 30, 40},
			insertValue:    30,
			expectedValues: []int{20, 30, 40}, // Duplicates are ignored
		},
	}

	for _, tc := range testCases {
		bst := BST{Root: nil}

		// Populate the BST
		for _, val := range tc.initialInputs {
			bst.Insert(val)
		}

		// Print the BST status before insertion
		fmt.Printf("Before Insertion: %v\n", bstToSlice(&bst))

		// Insert the value
		bst.Insert(tc.insertValue)

		// Convert the BST to a slice for comparison
		result := bstToSlice(&bst)
		fmt.Printf("%s - Insert: %v, Expected: %v, Result: %v", tc.name, tc.insertValue, tc.expectedValues, result)
		if !reflect.DeepEqual(result, tc.expectedValues) {
			t.Errorf("\n\nTest failed for %s - expected %v, got %v", tc.name, tc.expectedValues, result)
		} else {
			fmt.Printf("    --------- Pass")
		}
		fmt.Printf("\n")
	}
}

// TestPreOrderTraversal checks if the BST's pre-order traversal is correct
func TestPreOrderTraversal(t *testing.T) {
	fmt.Println("\n\n-------------- Test Binary Search Tree PreOrderTraversal ----------------------------")
	testCases := []struct {
		name     string
		inputs   []int
		expected []int // Expected values in pre-order traversal
	}{
		{
			name:     "Single element",
			inputs:   []int{10},
			expected: []int{10},
		},
		{
			name:     "Multiple elements",
			inputs:   []int{25, 14, 30, 8, 18, 26, 40},
			expected: []int{25, 14, 8, 18, 30, 26, 40},
		},
		{
			name:     "Empty tree",
			inputs:   []int{},
			expected: []int{},
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		bst := BST{Root: nil}

		// Inserting keys into BST
		for _, value := range tc.inputs {
			bst.Insert(value)
		}

		// Perform pre-order traversal and capture the output
		result := bst.PreOrderTraversal()

		// Compare expected and actual results
		if len(tc.expected) == 0 && len(result) == 0 || reflect.DeepEqual(result, tc.expected) {
			fmt.Printf("%s - Inputs: %v, Expected: %v, Result: %v    --------- Pass\n", tc.name, tc.inputs, tc.expected, result)
		} else {
			t.Errorf("Test failed for %s - expected %v, got %v", tc.name, tc.expected, result)
		}
	}
}

// TestPostOrderTraversal checks if the BST's post-order traversal is correct
func TestPostOrderTraversal(t *testing.T) {
	fmt.Println("\n\n-------------- Test Binary Search Tree PostOrderTraversal ----------------------------")
	testCases := []struct {
		name     string
		inputs   []int
		expected []int // Expected values in post-order traversal
	}{
		{
			name:     "Single element",
			inputs:   []int{10},
			expected: []int{10},
		},
		{
			name:     "Multiple elements",
			inputs:   []int{25, 14, 30, 8, 18, 26, 40},
			expected: []int{8, 18, 14, 26, 40, 30, 25},
		},
		{
			name:     "Empty tree",
			inputs:   []int{},
			expected: []int{},
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		bst := BST{Root: nil}

		// Inserting keys into BST
		for _, value := range tc.inputs {
			bst.Insert(value)
		}

		// Perform post-order traversal and capture the output
		result := bst.PostOrderTraversal()

		// Compare expected and actual results
		if len(tc.expected) == 0 && len(result) == 0 || reflect.DeepEqual(result, tc.expected) {
			fmt.Printf("%s - Inputs: %v, Expected: %v, Result: %v    --------- Pass\n", tc.name, tc.inputs, tc.expected, result)
		} else {
			t.Errorf("Test failed for %s - expected %v, got %v", tc.name, tc.expected, result)
		}
	}
}

// TestInOrderTraversal checks if the BST's in-order traversal is correct
func TestInOrderTraversal(t *testing.T) {
	fmt.Println("\n\n-------------- Test Binary Search Tree InorderTraversal ----------------------------")
	testCases := []struct {
		name     string
		inputs   []int
		expected []int // Expected values in in-order traversal
	}{
		{
			name:     "Single element",
			inputs:   []int{10},
			expected: []int{10},
		},
		{
			name:     "Multiple elements",
			inputs:   []int{25, 14, 30, 8, 18, 26, 40},
			expected: []int{8, 14, 18, 25, 26, 30, 40},
		},
		{
			name:     "Empty tree",
			inputs:   []int{},
			expected: []int{},
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		bst := BST{Root: nil}

		// Inserting keys into BST
		for _, value := range tc.inputs {
			bst.Insert(value)
		}

		// Perform in-order traversal and capture the output
		result := bst.InOrderTraversal()

		// Compare expected and actual results
		if len(tc.expected) == 0 && len(result) == 0 || reflect.DeepEqual(result, tc.expected) {
			fmt.Printf("%s - Inputs: %v, Expected: %v, Result: %v    --------- Pass\n", tc.name, tc.inputs, tc.expected, result)
		} else {
			t.Errorf("Test failed for %s - expected %v, got %v", tc.name, tc.expected, result)
		}
	}
	fmt.Println("\nOverall Result")
}
