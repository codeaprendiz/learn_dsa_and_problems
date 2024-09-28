package counting_bits

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestCountBits(t *testing.T) {
	testCases := map[int][]int{
		2:  {0, 1, 1},
		5:  {0, 1, 1, 2, 1, 2},
		10: {0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2},
	}

	for input, expected := range testCases {
		result := countBits(input)
		binaryN := strconv.FormatInt(int64(input), 2)
		fmt.Printf("\nInput : %v, Expected : %v, Result : %v, binaryN : %v ", input, expected, result, binaryN)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("\n\ncountBits(%d) = %v; expected %v", input, result, expected)
		} else {
			fmt.Printf("    --------- Pass")
		}
	}
	fmt.Printf("\nOverall Result\n")
}
