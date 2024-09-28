package how_many_numbers_are_smaller_than_current_number

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSmallerNumbersThanCurrent(t *testing.T) {
	testCases := map[string]struct {
		input    []int
		expected []int
	}{
		"Case 1 :                    ": {input: []int{8, 1, 2, 2, 3}, expected: []int{4, 0, 1, 1, 3}},
		"Case 2 :                    ": {input: []int{6, 5, 4, 8}, expected: []int{2, 1, 0, 3}},
		"Case 3 : All Same           ": {input: []int{7, 7, 7, 7}, expected: []int{0, 0, 0, 0}},
		// Adding additional cases
		"Case 4 : Empty Array        ": {input: []int{}, expected: []int{}},
		"Case 5 : Single Element     ": {input: []int{5}, expected: []int{0}},
		"Case 6 : Descending Order   ": {input: []int{5, 4, 3, 2, 1}, expected: []int{4, 3, 2, 1, 0}},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result := smallerNumbersThanCurrent(tc.input)
			inputStr := fmt.Sprintf("%v", tc.input)
			expectedStr := fmt.Sprintf("%v", tc.expected)
			resultStr := fmt.Sprintf("%v", result)
			fmt.Printf("\nTest %s - Input: %s, Expected: %s, Result: %s", name, inputStr, expectedStr, resultStr)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("\n\nTest failed for %s - got %s; want %s", name, resultStr, expectedStr)
			} else {
				fmt.Printf("    --------- Pass")
			}
		})
	}
	fmt.Printf("\nOverall Result\n")
}
