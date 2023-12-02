package main

import (
	"fmt"
	"strconv"
)

func Solution(N int) int {
	// Convert N to binary string
	// fmt.Println("Integer received : ", N)
	// fmt.Println("int64(N) : ", int64(N))
	currentGap, maxGap := 0, 0
	// http://golang.org/pkg/strconv/#FormatInt
	BinaryN := strconv.FormatInt(int64(N), 2)
	// fmt.Println("BinaryN : ", BinaryN)

	// fmt.Printf("\n\n")
	for _, bit := range BinaryN {
		// fmt.Println("Current bit : ", string(bit))
		// fmt.Printf("currentGap : %v, maxGap : %v\n", currentGap, maxGap)
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

func main() {
	fmt.Println(Solution(1041))
}
