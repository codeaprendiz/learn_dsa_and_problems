package binary_gap

// package main

import (
	// "fmt"
	"strconv"
)

func Solution(N int) int {
	currentGap, maxGap := 0, 0
	BinaryN := strconv.FormatInt(int64(N), 2) // http://golang.org/pkg/strconv/#FormatInt

	for _, bit := range BinaryN {
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
	/*
		In order to test by following
			$ go run binary_go.go
		Make the following changes
			import "fmt"
			package main
	*/

	// fmt.Println(Solution(9))
	// fmt.Println(Solution(7778742049))
	// fmt.Println(Solution(12586269025))

}
