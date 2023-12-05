package find_words_containing_character

//package main

import (
	// "fmt"
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

func main() {

}
