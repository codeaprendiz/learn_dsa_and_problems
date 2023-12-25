package minimum_number_of_steps_to_make_two_strings_anagrams

import (
// "fmt"
)

func minStepsToMakeAnagram(s string, t string) int {
	freqS := make(map[rune]int)
	freqT := make(map[rune]int)

	for _, ch := range s {
		freqS[ch]++
	}

	for _, ch := range t {
		freqT[ch]++
	}

	steps := 0

	for ch, count := range freqS {
		if freqT[ch] < count {
			steps += count - freqT[ch]
		}
	}

	return steps
}

// func main() {
// 	fmt.Println(minStepsToMakeAnagram("bab", "aba"))         // Output: 1
// 	fmt.Println(minStepsToMakeAnagram("anagram", "mangaar")) // Output: 0
// }
