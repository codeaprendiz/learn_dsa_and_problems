package minimum_number_of_steps_to_make_two_strings_anagrams

import (
// "fmt"
)

func minStepsToMakeAnagram(s string, t string) int {
	freqS := make(map[rune]int) // This line initializes a map named freqS, where each key is of type rune (representing a character in the string s), and each value is an int that counts the frequency of the corresponding character in s.
	freqT := make(map[rune]int) // So key is of type rune and value is of type int

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
