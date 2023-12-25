package minimum_number_of_steps_to_make_two_strings_anagrams

import (
	"fmt"
	"testing"
)

func TestMinStepsToMakeAnagram(t *testing.T) {
	testCases := map[string]struct {
		s        string
		t        string
		expected int
	}{
		"Example 1": {
			s:        "bab",
			t:        "aba",
			expected: 1,
		},
		"Example 2": {
			s:        "anagram",
			t:        "mangaar",
			expected: 0,
		},
		"Example 3": {
			s:        "listen",
			t:        "silent",
			expected: 0,
		},
		"Example 4": {
			s:        "abc",
			t:        "def",
			expected: 3,
		},
	}

	for inputDesc, tc := range testCases {
		result := minStepsToMakeAnagram(tc.s, tc.t)
		fmt.Printf("\n%s - s: %v, t: %v, Expected: %v, Result: %v", inputDesc, tc.s, tc.t, tc.expected, result)
		if result != tc.expected {
			t.Errorf("\n\nTest failed for %s - got %v; want %v", inputDesc, result, tc.expected)
		} else {
			fmt.Printf("    --------- Pass")
		}
	}
	fmt.Printf("\nOverall Result\n")
}
