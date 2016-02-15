package main

import (
	"fmt"
)

func Solution(N int) int {
	str := fmt.Sprintf("%b", N)

	byteOne := '1'

	maxGapLen := 0
	gapLen := 0

	gapStarted := false

	for _, s := range str {
		if s == byteOne {
			if gapLen > maxGapLen {
				maxGapLen = gapLen
			}

			gapLen = 0

			gapStarted = true
		} else {
			if gapStarted {
				gapLen = gapLen + 1
			}
		}
	}

	return maxGapLen
}

func main() {
	n := 74901729
	Solution(n)
	fmt.Println(fmt.Sprintf("%b", n))
	fmt.Println(Solution(n))
}
