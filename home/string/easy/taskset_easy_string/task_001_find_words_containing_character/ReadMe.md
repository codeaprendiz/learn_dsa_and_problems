# Find Words Containing A Character

[leetcode.com » Find words containing a character](https://leetcode.com/problems/find-words-containing-character)

## Solution

```go
func findWordsContaining(words []string, x byte) []int {
    var indices []int
    for i, word := range words {
        if strings.Contains(word, string(x)) {
            indices = append(indices, i)
        }
    }
    return indices
}
```

## Setup

```bash
$ go mod init find_words_containing_character 
go: creating new go.mod: module find_words_containing_character
go: to add module requirements and sums:
        go mod tidy

$ go test

Example 1 - Input: [test code], Char: 'e', Expected: [0 1], Result: [0 1]    --------- Pass
Example 2 - Input: [abc bcd aaaa cbc], Char: 'a', Expected: [0 2], Result: [0 2]    --------- Pass
Example 3 - Input: [abc bcd aaaa cbc], Char: 'z', Expected: [], Result: []    --------- Pass
Overall Result
PASS
ok      find_words_containing_character 0.293s
```
