package how_many_numbers_are_smaller_than_current_number

// package main

import (
	// "fmt"
	"sort"
)

func smallerNumbersThanCurrent(nums []int) []int {
	n := len(nums)
	sortedNums := make([]int, n)
	copy(sortedNums, nums)
	sort.Ints(sortedNums)

	// Use map to store smallest index for each number
	countMap := make(map[int]int)
	for index, value := range sortedNums {
		if _, found := countMap[value]; !found {
			countMap[value] = index
		}
	}

	result := make([]int, n)
	for index, value := range nums {
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
