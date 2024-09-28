package counting_bits

import (
	"fmt"
)

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

func main() {
	// Test the function with the example inputs
	fmt.Println(countBits(2))  // Output: [0, 1, 1]
	fmt.Println(countBits(5))  // Output: [0, 1, 1, 2, 1, 2]
	fmt.Println(countBits(10)) // Output: [0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2]
}
