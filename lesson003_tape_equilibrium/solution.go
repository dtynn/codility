// https://codility.com/demo/results/trainingD9MJK3-VRN/
package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	sums := make([]int, len(A))
	all := 0
	for i, num := range A {
		sums[i] = all
		all += num
	}

	minDiff := all - sums[1] - sums[1]
	if minDiff < 0 {
		minDiff = -minDiff
	}

	for idx := 2; idx < len(sums); idx += 1 {
		diff := all - sums[idx] - sums[idx]
		if diff < 0 {
			diff = -diff
		}

		if diff < minDiff {
			minDiff = diff
		}
	}

	return minDiff
}
