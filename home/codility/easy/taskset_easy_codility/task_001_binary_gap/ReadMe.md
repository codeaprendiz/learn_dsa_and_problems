# Binary Gap

[app.codility.com » BinaryGap](https://app.codility.com/c/run/trainingPYBXG8-TKK/)

## Question

A binary gap within a positive integer N is any maximal sequence of consecutive zeros that is surrounded by ones at both ends in the binary representation of N.

For example, number 9 has binary representation 1001 and contains a binary gap of length 2. The number 529 has binary representation 1000010001 and contains two binary gaps: one of length 4 and one of length 3. The number 20 has binary representation 10100 and contains one binary gap of length 1. The number 15 has binary representation 1111 and has no binary gaps. The number 32 has binary representation 100000 and has no binary gaps.

Write a function:

```go
func Solution(N int) int
```

that, given a positive integer N, returns the length of its longest binary gap. The function should return 0 if N doesn't contain a binary gap.

For example, given N = 1041 the function should return 5, because N has binary representation 10000010001 and so its longest binary gap is of length 5. Given N = 32 the function should return 0, because N has binary representation '100000' and thus no binary gaps.

Write an efficient algorithm for the following assumptions:

> N is an integer within the range [1..2,147,483,647].

## Approach

To solve this problem, we can convert the integer `N` into its binary representation and then iterate through the binary string to find the longest sequence of zeros that is flanked by ones. Here's a step-by-step approach:

1. **Convert to Binary**: Convert the integer `N` to its binary representation.
2. **Iterate Through the Binary String**: Go through each character in the binary string.
3. **Count Zeros**: When a '1' is encountered, start counting zeros until another '1' is found.
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