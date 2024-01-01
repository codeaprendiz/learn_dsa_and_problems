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

## Notes

1. In the function `findWordsContaining`, the second argument `x` is of type `byte` because individual characters in Go can be conveniently represented as bytes when dealing with ASCII characters. The `byte` type in Go is an alias for `uint8` and is commonly used to handle individual ASCII characters.
2. However, the `strings.Contains` function requires its second argument to be a string, not a single byte. This is why `x` needs to be converted to a string before being passed to `strings.Contains`.

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
