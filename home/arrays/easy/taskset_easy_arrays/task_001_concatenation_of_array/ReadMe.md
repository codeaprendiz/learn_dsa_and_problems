# Concatenation of Array

[leetcode.com » Concatenation of Array](https://leetcode.com/problems/concatenation-of-array)

## Solution

```bash
func getConcatenation(nums []int) []int {
    lenOfNums := len(nums)
    ans := make([]int, 2*(lenOfNums))
    for index, element := range nums {
        ans[index] = element
        ans[index+lenOfNums] = element
    }
    return ans
}
```

## Setup

```bash
$ go mod init concatenation_of_array
go: creating new go.mod: module concatenation_of_array
go: to add module requirements and sums:
        go mod tidy

$ go test

Input: 1,2,3,4,5, Expected: [1 2 3 4 5 1 2 3 4 5], Result: [1 2 3 4 5 1 2 3 4 5]    --------- Pass
Input: 2,1,3,5,3, Expected: [2 1 3 5 3 2 1 3 5 3], Result: [2 1 3 5 3 2 1 3 5 3]    --------- Pass
Overall Result
PASS
ok      concatenation_of_array  1.120s
```

## Time and Space Complexity

1. **Time Complexity: O(n)**
   - The function iterates over the array `nums` which has `n` elements. Each iteration involves constant-time operations. Hence, the time complexity is O(n), where n is the length of `nums`.

2. **Auxiliary Space: O(n)**
   - Auxiliary space refers to extra space or temporary space used by an algorithm, not including space taken by inputs. In this case, the auxiliary space is used by the `ans` array, which is twice the size of `nums`. Since `ans` has a length of `2n` and constants are dropped in Big O notation, the auxiliary space complexity is O(n).

3. **Space Complexity: O(n)**
   - Space complexity includes both the auxiliary space and the space used by inputs. Since the auxiliary space (O(n)) is larger than the space used for variables and indices, the overall space complexity is dominated by the auxiliary space, which is O(n).
