package concatenation_of_array

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetConcatenation(t *testing.T) {
	testCases := map[string][]int{
		"1,2,3,4,5": {1, 2, 3, 4, 5},
		"2,1,3,5,3": {2, 1, 3, 5, 3},
	}

	for inputDesc, input := range testCases {
		expected := append(input, input...)
		result := getConcatenation(input)
		fmt.Printf("\nInput: %v, Expected: %v, Result: %v", inputDesc, expected, result)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("\n\ngetConcatenation(%v) = %v; expected %v", input, result, expected)
		} else {
			fmt.Printf("    --------- Pass")
		}
	}
	fmt.Printf("\nOverall Result\n")
}
