# How Many Numbers Are Smaller Than the Current Number

[leetcode.com » How Many Numbers Are Smaller Than the Current Number](https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number)

> Revised : 1

<br>

## Solution

```go
func smallerNumbersThanCurrent(nums []int) []int {
    n := len(nums)
    sortedNums := make([]int, n)    // https://pkg.go.dev/builtin#make
    copy(sortedNums, nums) // https://pkg.go.dev/builtin#copy
    sort.Ints(sortedNums)  // https://pkg.go.dev/sort#example-Ints

    // Use map to store smallest index for each number
    countMap := make(map[int]int) // map of key int and value as int
    for index, value := range sortedNums {    // Note the key of Map is actually the value in sortedNums, so for [ 0 2 ], we create countMap[0]=0, countMap[2]=1
        if _, found := countMap[value]; !found { // If there is repeation, like 0 1 1, then our countMap[1] should have 1, the second occurance should fail.
            countMap[value] = index // Note: index here is also the number of values that are smaller than the number at this index. For 0th index number, there is no number smaller than this and likewise.
        }
    }

    result := make([]int, n)
    for index, value := range nums { // So now for every nums[index], i need to return how many numbers are smaller than this, which is actually stored in countMap[nums[index]]
        result[index] = countMap[value]
    }

    return result
}
```

<br>

## Setup

```bash
$ go mod init how_many_numbers_are_smaller_than_current_number                  
go: creating new go.mod: module how_many_numbers_are_smaller_than_current_number
go: to add module requirements and sums:
        go mod tidy


$ go test

Test Case 1 :                     - Input: [8 1 2 2 3], Expected: [4 0 1 1 3], Result: [4 0 1 1 3]    --------- Pass
Test Case 2 :                     - Input: [6 5 4 8], Expected: [2 1 0 3], Result: [2 1 0 3]    --------- Pass
Test Case 3 : All Same            - Input: [7 7 7 7], Expected: [0 0 0 0], Result: [0 0 0 0]    --------- Pass
Test Case 4 : Empty Array         - Input: [], Expected: [], Result: []    --------- Pass
Test Case 5 : Single Element      - Input: [5], Expected: [0], Result: [0]    --------- Pass
Test Case 6 : Descending Order    - Input: [5 4 3 2 1], Expected: [4 3 2 1 0], Result: [4 3 2 1 0]    --------- Pass
Overall Result
PASS
ok      how_many_numbers_are_smaller_than_current_number        0.425s
```

<br>

## Dry Run

Let's do a dry run of the `smallerNumbersThanCurrent` function with the input `nums1 := []int{8, 1, 2, 2, 3}`:

1. **Initialization**:
   - `n := len(nums)` sets `n` to 5, as there are 5 elements in `nums`.
   - `sortedNums` is created and made a copy of `nums`, so `sortedNums = [8, 1, 2, 2, 3]`.

2. **Sorting**:
   - `sort.Ints(sortedNums)` sorts the array, so `sortedNums` becomes `[1, 2, 2, 3, 8]`.

3. **Creating and Populating `countMap`**:
   - `countMap` is created to store the smallest index for each unique number in `sortedNums`.
   - Looping through `sortedNums`:
     - When `index = 0`, `value = 1`, `countMap[1] = 0`.
     - When `index = 1`, `value = 2`, `countMap[2] = 1`.
     - When `index = 2`, `value = 2` again, but `countMap[2]` is already set, so no change.
     - When `index = 3`, `value = 3`, `countMap[3] = 3`.
     - When `index = 4`, `value = 8`, `countMap[8] = 4`.
   - Now, `countMap` is `{1: 0, 2: 1, 3: 3, 8: 4}`.

4. **Creating the `result` Array**:
   - Loop through the original `nums` array and use `countMap` to get the count of smaller numbers.
   - For `nums[0] = 8`, `result[0] = countMap[8] = 4`.
   - For `nums[1] = 1`, `result[1] = countMap[1] = 0`.
   - For `nums[2] = 2`, `result[2] = countMap[2] = 1`.
   - For `nums[3] = 2`, `result[3] = countMap[2] = 1`.
   - For `nums[4] = 3`, `result[4] = countMap[3] = 3`.
   - Now, `result` is `[4, 0, 1, 1, 3]`.

5. **Return**:
   - The function returns the `result` array, which is `[4, 0, 1, 1, 3]`.

This dry run demonstrates how the function processes the input array to determine how many numbers are smaller than each number in it. The key part of this approach is using sorting and a map to efficiently count smaller numbers.

<br>

## Time Complexity

The time complexity of the solution can be analyzed based on the major operations performed in the `smallerNumbersThanCurrent` function:

1. **Copying and Sorting the Array**:
   - Copying the array takes O(n) time, where n is the number of elements in `nums`.
   - Sorting the array using `sort.Ints` has a time complexity of O(n log n).

2. **Populating the `countMap`**:
   - Iterating through the sorted array and populating the map takes O(n) time. In the worst case, each element is unique, and we insert each one into the map.

3. **Filling the Result Array**:
   - Iterating over the original array and filling the result array using values from `countMap` also takes O(n) time.

Therefore, the overall time complexity of the function is dominated by the sorting step, which is O(n log n). The other steps (copying the array, populating the map, and filling the result array) add linear time complexity, but they do not change the overall time complexity. Hence, the total time complexity of the solution is O(n log n).