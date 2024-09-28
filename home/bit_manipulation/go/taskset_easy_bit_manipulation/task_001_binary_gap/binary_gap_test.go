package binary_gap

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	testCases := map[int]int{
		9:           2, // Binary: 1001 has a binary gap of 2
		8:           0, // 8 in binary is "1000", There are not any adjacent pairs of 1's in the binary representation of 8, so we return 0.
		7778742049:  4, // Binary: 111001111101001100010111100100001 has a binary gap of 4
		12586269025: 4, // Binary: 1011101110001100110011100101100001 has a binary gap of 4
	}

	for input, expected := range testCases {
		result := Solution(input)
		binaryN := strconv.FormatInt(int64(input), 2)
		fmt.Printf("\nInput : %v, Expected : %v, Result : %v, binaryN : %v ", input, expected, result, binaryN)
		if result != expected {
			t.Errorf("\n\nSolution(%d) = %d; expected %d", input, result, expected)
		} else {
			fmt.Printf("    --------- Pass")
		}
	}
	fmt.Printf("\nOverall Result\n")
}
