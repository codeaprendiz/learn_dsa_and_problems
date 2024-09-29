# Binary Gap

[leetcode.com » BinaryGap](https://leetcode.com/problems/binary-gap/)

> Revision Count : 1

<br>

## Solution

```go
func Solution(N int) int {
    currentGap, maxGap := 0, 0
    BinaryN := strconv.FormatInt(int64(N), 2) // http://golang.org/pkg/strconv/#FormatInt

    for _, bit := range BinaryN {
        if bit == '1' {
            if currentGap > maxGap {
                maxGap = currentGap
            }
            currentGap = 0
        } else {
            currentGap = currentGap + 1
        }
    }
    return maxGap
}
```

<br>

## Setup

```bash
$ go mod init binary_gap       
go: creating new go.mod: module binary_gap
go: to add module requirements and sums:
        go mod tidy

$ go test

Input : 9, Expected : 2, Result : 2, binaryN : 1001     --------- Pass
Input : 7778742049, Expected : 4, Result : 4, binaryN : 111001111101001100010111100100001     --------- Pass
Input : 12586269025, Expected : 4, Result : 4, binaryN : 1011101110001100110011100101100001     --------- Pass
Overall Result
PASS
ok      binary_gap      0.471s
```

<br>

## Approach

To solve this problem, we can convert the integer `N` into its binary representation and then iterate through the binary string to find the longest sequence of zeros that is flanked by ones. Here's a step-by-step approach:

1. **Convert to Binary**: Convert the integer `N` to its binary representation.
2. **Iterate Through the Binary String**: Go through each character in the binary string.
3. **Count Zeros**: When a `1` is encountered, start counting zeros until another `1` is found.
4. **Update the Maximum Gap**: If the current gap (number of zeros between two ones) is greater than the previous maximum, update the maximum.
5. **Return the Maximum Gap**: After iterating through the entire string, return the maximum gap found.

This function should work efficiently for the given range of integers `[1..2,147,483,647]`. The time complexity is O(log N) since the binary representation of a number `N` has a length proportional to the logarithm of `N`, and the space complexity is also O(log N) due to the storage of the binary representation.

<br>

### Why use `binaryN := strconv.FormatInt(int64(N), 2)`

In Go, the `strconv.FormatInt` function requires an argument of type `int64`. If your variable `N` is declared as an `int`, which could be either `int32` or `int64` depending on the system architecture (32-bit or 64-bit), it's good practice to explicitly convert it to `int64` to ensure compatibility with the `FormatInt` function.

This conversion ensures that your program behaves consistently across different platforms, as `int` in Go is architecture dependent.
An `int` in Go can be 32 or 64 bits depending on the platform (32-bit systems use 32 bits for `int`, and 64-bit systems use 64 bits). To avoid any compatibility issues and to make sure the function works as expected on all platforms, the integer `N` is explicitly converted to `int64`. This conversion ensures that the `strconv.FormatInt` function receives the correctly typed argument regardless of the platform.

- On a 32-bit system, `int` is equivalent to `int32`.
- On a 64-bit system, `int` is equivalent to `int64`.

By explicitly converting `N` to `int64`, you make sure that the `FormatInt` function receives the correct type regardless of the underlying architecture, thereby making your code more portable.

If `N` is already an `int64`, then the explicit conversion is not necessary, but it does not harm and makes the code more explicit and possibly clearer.

<br>

### Why us string(bit) to print Binary

In Go, when you iterate over a string with a range loop, each character (or more precisely, each rune) is represented by its Unicode code point. The Unicode code points for the characters '0' and `1` are 48 and 49, respectively, in decimal. That's why you see 48 when you print a '0' bit and 49 when you print a `1` bit. 
If you want to print the character itself ('0' or `1`) rather than its Unicode code point, you need to convert the rune back into a character. This can be done simply by using the string() function to convert the rune to a string

<br>

## Time And Space Complexity

**Time Complexity Analysis**:

1. The time complexity of the `Solution` function is primarily determined by the loop that iterates over the binary representation of the input integer `N`. The length of this binary representation is proportional to the logarithm of `N`, often denoted as `O(log N)`. This is because the number of digits in a binary representation of a number grows logarithmically with the number itself.
2. Inside the loop, the operations performed (comparison, assignment) are constant time operations, i.e., their execution time does not depend on the size of the input. Therefore, the overall time complexity of the function is `O(log N)`, where `N` is the value of the input integer.

**Space Complexity Analysis**:

1. The space complexity of the `Solution` function includes both the space used by the input and the auxiliary space used by the function. The input space is the space taken by the integer `N`, which is a constant and typically does not count towards space complexity. The most significant contributor to space complexity here is the binary string `BinaryN`, whose length is proportional to `O(log N)`.
2. Other variables used in the function, such as `currentGap` and `maxGap`, require a constant amount of space. Therefore, these do not significantly impact the overall space complexity. As a result, the total space complexity of the function is dominated by the length of the binary string, leading to a space complexity of `O(log N)`. This reflects the space needed to store the binary representation of `N`.