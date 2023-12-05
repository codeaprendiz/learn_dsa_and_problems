package concatenation_of_array

// package main

import (
// "fmt"
)

func getConcatenation(nums []int) []int {
	lenOfNums := len(nums)
	ans := make([]int, 2*(lenOfNums))
	for index, element := range nums {
		ans[index] = element
		ans[index+lenOfNums] = element
	}
	return ans
}

func main() {
	// firstArray := []int{1, 2, 1}
	// ans := getConcatenation(firstArray)
	// fmt.Println(ans)
}
