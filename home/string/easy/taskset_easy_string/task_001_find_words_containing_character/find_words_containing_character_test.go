package find_words_containing_character

// package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindWordsContaining(t *testing.T) {
	testCases := map[string]struct {
		words    []string
		x        byte
		expected []int
	}{
		"Example 1": {
			words:    []string{"test", "code"},
			x:        'e',
			expected: []int{0, 1},
		},
		"Example 2": {
			words:    []string{"abc", "bcd", "aaaa", "cbc"},
			x:        'a',
			expected: []int{0, 2},
		},
		"Example 3": {
			words:    []string{"abc", "bcd", "aaaa", "cbc"},
			x:        'z',
			expected: []int{},
		},
	}

	for inputDesc, tc := range testCases {
		result := findWordsContaining(tc.words, tc.x)
		fmt.Printf("\n%s - Input: %v, Char: %q, Expected: %v, Result: %v", inputDesc, tc.words, tc.x, tc.expected, result)
		if (len(result) == 0 && len(tc.expected) == 0) || reflect.DeepEqual(result, tc.expected) {
			fmt.Printf("    --------- Pass")
		} else {
			t.Errorf("\n\nTest %s failed: findWordsContaining(%v, %q) = %v; expected %v", inputDesc, tc.words, tc.x, result, tc.expected)
		}
	}
	fmt.Printf("\nOverall Result\n")
}
