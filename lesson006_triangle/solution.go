// https://codility.com/demo/results/trainingG7UPFX-GJA/

package solution

import "sort"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	if len(A) < 3 {
		return 0
	}

	sort.Ints(A)

	for i := 0; i < len(A)-2; i++ {
		if A[i+2] <= 0 || A[i+1] <= 0 || A[i] <= 0 {
			continue
		}

		if A[i]+A[i+1] > A[i+2] {
			return 1
		}
	}

	return 0
}
