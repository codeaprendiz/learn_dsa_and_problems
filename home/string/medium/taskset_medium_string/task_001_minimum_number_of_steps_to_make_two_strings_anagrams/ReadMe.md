# [Minimum Number of Steps to Make Two Strings Anagrams](https://leetcode.com/problems/minimum-number-of-steps-to-make-two-strings-anagram/description)

[leetcode.com » Minimum Number of Steps to Make Two Strings Anagrams](https://leetcode.com/problems/minimum-number-of-steps-to-make-two-strings-anagram/description)

> An Anagram of a string is a string that contains the same characters with a different (or the same) ordering.

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

## Complexity

Time Complexity : O(n)
