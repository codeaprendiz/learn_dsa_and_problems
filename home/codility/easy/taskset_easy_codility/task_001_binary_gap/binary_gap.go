package main

import (
	"fmt"
	"strconv"
)

func Solution(N int) int {
	// Convert N to binary string
	fmt.Println("Integer received : ", N)
	fmt.Println("int64(N) : ", int64(N))
	currentGap, maxGap := 0, 0
	// http://golang.org/pkg/strconv/#FormatInt
	BinaryN := strconv.FormatInt(int64(N), 2)

	fmt.Println("BinaryN : ", BinaryN)
	fmt.Printf("currentGap : %v, maxGap : %v\n", currentGap, maxGap)
	return maxGap
}

func main() {
	fmt.Println(Solution(1041))
}
