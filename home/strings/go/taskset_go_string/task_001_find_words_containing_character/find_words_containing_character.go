package find_words_containing_character

import (
	"strings"
)

func findWordsContaining(words []string, x byte) []int {
	var indices []int

	for i, word := range words {
		if strings.Contains(word, string(x)) {
			indices = append(indices, i)
		}
	}

	return indices
}
