package how_many_numbers_are_smaller_than_current_number

// package main

import (
	// "fmt"
	"sort"
)

func smallerNumbersThanCurrent(nums []int) []int {
	n := len(nums)
	sortedNums := make([]int, n) // https://pkg.go.dev/builtin#make
	copy(sortedNums, nums)       // https://pkg.go.dev/builtin#copy
	sort.Ints(sortedNums)        // https://pkg.go.dev/sort#example-Ints

	// Use map to store smallest index for each number
	countMap := make(map[int]int)          // map of key int and value as int
	for index, value := range sortedNums { // Note the key of Map is actually the value in sortedNums, so for [ 0 2 ], we create countMap[0]=0, countMap[2]=1
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

// func main() {
// 	nums1 := []int{8, 1, 2, 2, 3}
// 	fmt.Println(smallerNumbersThanCurrent(nums1)) // [4,0,1,1,3]

// 	nums2 := []int{6, 5, 4, 8}
// 	fmt.Println(smallerNumbersThanCurrent(nums2)) // [2,1,0,3]

// 	nums3 := []int{7, 7, 7, 7}
// 	fmt.Println(smallerNumbersThanCurrent(nums3)) // [0,0,0,0]
// }
