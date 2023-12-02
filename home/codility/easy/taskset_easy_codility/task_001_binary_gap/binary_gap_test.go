package binary_gap

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	testCases := map[int]int{
		74901729:   4,
		805306373:  25,
		1073741825: 29,
		1376796946: 5,
		1610612737: 28,
		2147483647: 0,
	}

	for input, expected := range testCases {
		result := Solution(input)
		if result != expected {
			t.Errorf("\n\nSolution(%d) = %d; expected %d", input, result, expected)
		} else {
			fmt.Printf("\nInput : %v, Expected : %v, Result : %v", input, expected, result)
		}
	}
}
