# Counting Bits

[leetcode.com » BinaryGap](https://leetcode.com/problems/counting-bits)

- [Counting Bits](#counting-bits)
  - [Solution](#solution)
    - [Explanation of Changes](#explanation-of-changes)
  - [Explaination](#explaination)
    - [Dynamic Programming Approach](#dynamic-programming-approach)
    - [Example Calculation for `( n = 5 )`](#example-calculation-for--n--5-)
    - [Final Result](#final-result)
  - [Unit Tests](#unit-tests)

## Solution

```go
func countBits(n int) []int {
  // Initialize an array to store the result, with length n + 1
  ans := make([]int, n+1)

  // Iterate over each number from 0 to n
  for i := 1; i <= n; i++ {
    if i%2 == 0 {
      // Even number: ans[i] = ans[i / 2]
      ans[i] = ans[i/2]
    } else {
      // Odd number: ans[i] = ans[i / 2] + 1
      ans[i] = ans[i/2] + 1
    }
  }

  return ans
}
```

### Explanation of Changes

1. **Initialization**: Create a slice `ans` of length `n + 1` to store the results.
2. **Iteration**: Loop from `1` to `n` and calculate the number of 1 bits for each number `i`:
   - For even numbers (`i % 2 == 0`): `ans[i] = ans[i / 2]`
   - For odd numbers (`i % 2 != 0`): `ans[i] = ans[i / 2] + 1`

By using this approach, we ensure that the program computes the result in linear time, `( O(n) )`, by leveraging previously computed results. This updated program adheres to the optimized dynamic programming approach.

## Explaination

Let's break down the process step-by-step using the same dynamic programming approach to calculate the number of 1 bits in the binary representation of numbers from 0 to `( n )`.

### Dynamic Programming Approach

1. **Base Cases**:
   - For `result[0]`, the count of 1 bits is 0 because the binary representation of 0 is `0`.
   - For `result[1]`, the count of 1 bits is 1 because the binary representation of 1 is `1`.

2. **Even Numbers**:
   - For any even number `( x )`, the count of 1 bits is the same as the count of 1 bits in `( x / 2 )`.
   - This is because shifting an even number to the right by one bit (i.e., dividing by 2) simply removes the trailing zero without affecting the number of 1 bits.
   - Formula: `result[x] = result[x / 2]`

3. **Odd Numbers**:
   - For any odd number `( x )`, the count of 1 bits is the count of 1 bits in `( x / 2 )` plus 1 (for the least significant bit, which is 1).
   - This is because shifting an odd number to the right by one bit (i.e., dividing by 2) reduces the number but the least significant bit (which is 1) is lost.
   - Formula: `result[x] = result[x / 2] + 1`

### Example Calculation for `( n = 5 )`

- **Initialization**:
  `[ result = [0, 0, 0, 0, 0, 0]`

- **Iteration**:
  - `x = 0`:
    - `result[0] = 0`
    - `result = [0, 0, 0, 0, 0, 0]`

  - `x = 1`:
    - `result[1] = 1`
    - `[result = [0, 1, 0, 0, 0, 0]`

  - `x = 2` (even):
    - `result[2] = result[2 / 2] = result[1] = 1`
    - `result = [0, 1, 1, 0, 0, 0]`

  - `x = 3` (odd):
    - `result[3] = result[3 / 2] + 1 = result[1] + 1 = 1 + 1 = 2`
    - `result = [0, 1, 1, 2, 0, 0]`

  - `x = 4` (even):
    - `result[4] = result[4 / 2] = result[2] = 1`
    - `result = [0, 1, 1, 2, 1, 0]`

  - `x = 5` (odd):
    - `result[5] = result[5 / 2] + 1 = result[2] + 1 = 1 + 1 = 2`
    - `result = [0, 1, 1, 2, 1, 2]`

### Final Result

The final `result` array, which contains the number of 1 bits for each integer from 0 to 5, is:

`result = [0, 1, 1, 2, 1, 2]`

This approach uses dynamic programming to build up the solution from known results, ensuring an efficient computation with a time complexity of `O(n)`.

## Unit Tests

```bash
go mod init counting_bits
```

```bash
go mod tidy
```

```bash
go test
```

Output

```bash
$ go test    

Input : 2, Expected : [0 1 1], Result : [0 1 1], binaryN : 10     --------- Pass
Input : 5, Expected : [0 1 1 2 1 2], Result : [0 1 1 2 1 2], binaryN : 101     --------- Pass
Input : 10, Expected : [0 1 1 2 1 2 2 3 1 2 2], Result : [0 1 1 2 1 2 2 3 1 2 2], binaryN : 1010     --------- Pass
Overall Result
PASS
ok      counting_bits   0.741s
```
