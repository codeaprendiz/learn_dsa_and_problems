package binary_gap

import "testing"

func TestSolution(t *testing.T) {
	testCases := map[int]int{
		1:          0,
		5:          1,
		6:          0,
		9:          2,
		11:         1,
		15:         0,
		16:         0,
		19:         2,
		20:         1,
		32:         0,
		42:         1,
		328:        2,
		1024:       0,
		1041:       5,
		1162:       3,
		51712:      2,
		561892:     3,
		66561:      9,
		6291457:    20,
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
			t.Errorf("Solution(%d) = %d; expected %d", input, result, expected)
		}
	}
}
