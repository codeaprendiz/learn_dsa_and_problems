package main

import (
	"fmt"
	"strconv"
)

func Solution(N int) int {
	// Convert N to binary string
	binaryN := strconv.FormatInt(int64(N), 2)
	fmt.Println(binaryN)

	currentGap, maxGap := -1, 0
	// Iterate through each character in the binary string
	for _, bit := range binaryN {
		if bit == 1 {
			// Update maxGap if currentGap is larger
			if currentGap > maxGap {
				maxGap = currentGap
			}
			// Reset curentGap to 0 after finding bit 1
			currentGap = 0
		} else if currentGap >= 0 {
			// Only count gaps after the first one has been found
			currentGap += 1
		}
	}

	return maxGap
}

func main() {
	Solution(1041)
}
