# [Minimum Number of Steps to Make Two Strings Anagrams](https://leetcode.com/problems/minimum-number-of-steps-to-make-two-strings-anagram/description)

[leetcode.com » Minimum Number of Steps to Make Two Strings Anagrams](https://leetcode.com/problems/minimum-number-of-steps-to-make-two-strings-anagram/description)

> Revision Count : 1

An Anagram of a string is a string that contains the same characters with a different (or the same) ordering.

<br>

## Solution

```go
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
```

<br>

### Why use `rune`?

In Go, a rune represents a Unicode code point. It's essentially an alias for int32, and it's used to handle individual characters in a string, especially when dealing with Unicode characters that may occupy more than one byte.

1. **Unicode Support**: In Go, strings are byte sequences. A `rune` handles Unicode characters, which may be more than one byte, ensuring correct processing of a wide range of text, including non-ASCII characters.

2. **Correct String Iteration**: When iterating over a string with `for...range` in Go, the iteration variable is a `rune`. This approach accurately handles multi-byte Unicode characters, ensuring each character, regardless of its byte size, is processed correctly.

3. **Best Practice**: Using `rune` is recommended in Go for character handling, even for ASCII-only strings, as it makes the code more versatile and robust for different types of text data.

<br>

## Skills

String, Hash Table, Counting

```bash
$ go mod init minimum_number_of_steps_to_make_two_strings_anagrams            
go: creating new go.mod: module minimum_number_of_steps_to_make_two_strings_anagrams
go: to add module requirements and sums:
        go mod tidy

$ go test

Example 1 - s: bab, t: aba, Expected: 1, Result: 1    --------- Pass
Example 2 - s: anagram, t: mangaar, Expected: 0, Result: 0    --------- Pass
Example 3 - s: listen, t: silent, Expected: 0, Result: 0    --------- Pass
Example 4 - s: abc, t: def, Expected: 3, Result: 3    --------- Pass
Overall Result
PASS
ok      minimum_number_of_steps_to_make_two_strings_anagrams    0.348s
```

<br>

## Dry Run

Let's perform a dry run of the `minStepsToMakeAnagram` function with the input strings `s = "bab"` and `t = "aba"`:

1. **Initialize Frequency Maps**:
   - Create two maps, `freqS` and `freqT`, to store the frequency of each character in `s` and `t`, respectively.

2. **Count Frequency in `s` (`"bab")`**:
   - Iterate over each character in `s`:
     - `freqS['b']++`: After first 'b', `freqS` becomes `{'b': 1}`.
     - `freqS['a']++`: After 'a', `freqS` becomes `{'b': 1, 'a': 1}`.
     - `freqS['b']++`: After second 'b', `freqS` becomes `{'b': 2, 'a': 1}`.

3. **Count Frequency in `t` (`"aba")`**:
   - Iterate over each character in `t`:
     - `freqT['a']++`: After first 'a', `freqT` becomes `{'a': 1}`.
     - `freqT['b']++`: After 'b', `freqT` becomes `{'a': 1, 'b': 1}`.
     - `freqT['a']++`: After second 'a', `freqT` becomes `{'a': 2, 'b': 1}`.

4. **Calculate Steps Needed for Anagram**:
   - Initialize `steps` to 0.
   - Iterate over each character in `freqS` and compare with `freqT`:
     - For character 'b' in `freqS`, `freqS['b'] = 2` and `freqT['b'] = 1`. Since `freqT['b'] < freqS['b']`, we add `2 - 1 = 1` to `steps`.
     - For character 'a' in `freqS`, `freqS['a'] = 1` and `freqT['a'] = 2`. `freqT['a']` is not less than `freqS['a']`, so no change in `steps`.
   - `steps` is now `1`.

5. **Return Result**:
   - The function returns `steps`, which is `1`.

<br>

## Complexity

**Time Complexity ( O(n) ):**

1. Iterates over each character in strings `s` and `t`, each of length up to \( n \).
2. Frequency calculation for each string is linear with respect to its length.
3. Therefore, total time complexity is \( O(n) \), where \( n \) is the length of the longer string.

**Space Complexity ( O(n) ):**

1. Uses two maps for frequency counts, size limited by number of unique characters, leading to a constant auxiliary space of \( O(1) \).
2. However, including the input size, the space complexity is dependent on the length of strings `s` and `t`.
3. Combining input and auxiliary space, the total space complexity is \( O(n) \), where \( n \) is the length of the longer string.
