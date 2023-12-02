# Binary Gap

[leetcode.com » BinaryGap](https://leetcode.com/problems/binary-gap/)

- [Binary Gap](#binary-gap)
  - [Setup](#setup)
  - [Approach](#approach)
    - [Why use `binaryN := strconv.FormatInt(int64(N), 2)`](#why-use-binaryn--strconvformatintint64n-2)
    - [Why us string(bit) to print Binary](#why-us-stringbit-to-print-binary)

## Setup

```bash
$ go mod init binary_gap       
go: creating new go.mod: module binary_gap
go: to add module requirements and sums:
        go mod tidy

$ go test
Input : 805306373, Expected : 25, Result : 25
Input : 1073741825, Expected : 29, Result : 29
Input : 1376796946, Expected : 5, Result : 5
Input : 1610612737, Expected : 28, Result : 28
Input : 2147483647, Expected : 0, Result : 0
Input : 74901729, Expected : 4, Result : 4

Overall Result
PASS
ok      binary_gap      1.030s 
```

## Approach

To solve this problem, we can convert the integer `N` into its binary representation and then iterate through the binary string to find the longest sequence of zeros that is flanked by ones. Here's a step-by-step approach:

1. **Convert to Binary**: Convert the integer `N` to its binary representation.
2. **Iterate Through the Binary String**: Go through each character in the binary string.
3. **Count Zeros**: When a `1` is encountered, start counting zeros until another `1` is found.
4. **Update the Maximum Gap**: If the current gap (number of zeros between two ones) is greater than the previous maximum, update the maximum.
5. **Return the Maximum Gap**: After iterating through the entire string, return the maximum gap found.

This function should work efficiently for the given range of integers `[1..2,147,483,647]`. The time complexity is O(log N) since the binary representation of a number `N` has a length proportional to the logarithm of `N`, and the space complexity is also O(log N) due to the storage of the binary representation.

### Why use `binaryN := strconv.FormatInt(int64(N), 2)`

In Go, the `strconv.FormatInt` function requires an argument of type `int64`. If your variable `N` is declared as an `int`, which could be either `int32` or `int64` depending on the system architecture (32-bit or 64-bit), it's good practice to explicitly convert it to `int64` to ensure compatibility with the `FormatInt` function.

This conversion ensures that your program behaves consistently across different platforms, as `int` in Go is architecture dependent:

- On a 32-bit system, `int` is equivalent to `int32`.
- On a 64-bit system, `int` is equivalent to `int64`.

By explicitly converting `N` to `int64`, you make sure that the `FormatInt` function receives the correct type regardless of the underlying architecture, thereby making your code more portable. 

If `N` is already an `int64`, then the explicit conversion is not necessary, but it does not harm and makes the code more explicit and possibly clearer.

### Why us string(bit) to print Binary

In Go, when you iterate over a string with a range loop, each character (or more precisely, each rune) is represented by its Unicode code point. The Unicode code points for the characters '0' and `1` are 48 and 49, respectively, in decimal. That's why you see 48 when you print a '0' bit and 49 when you print a `1` bit. 
If you want to print the character itself ('0' or `1`) rather than its Unicode code point, you need to convert the rune back into a character. This can be done simply by using the string() function to convert the rune to a string
