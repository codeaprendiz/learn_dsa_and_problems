package main

import (
	"fmt"
	"strconv"
)

func Solution(N int) int {
	// Convert N to binary string
	binaryN := strconv.FormatInt(int64(N), 2)
	fmt.Println(binaryN)

	currentGap, maxGap := 0, 0
	// Iterate through each character in the binary string
	for _, bit := range binaryN {
		fmt.Printf("currentGap : %v, maxGap : %v, bit : %v\n", currentGap, maxGap, bit)
		fmt.Println(bit)
		if bit == '1' {
			fmt.Println("Bit is 1")
			// Update maxGap if currentGap is larger
			if currentGap > maxGap {
				fmt.Println("currentGap > maxGap")
				maxGap = currentGap
			}
			// Reset curentGap to 0 after finding bit 1
			currentGap = 0
		} else if bit == '0' {
			// Only count gaps after the first one has been found
			fmt.Println("currentGap >= 0 OR bit is 0")
			currentGap += 1
		} else {
			fmt.Println("Default")
		}
	}

	return maxGap
}

func main() {
	fmt.Println(Solution(1041))
}
