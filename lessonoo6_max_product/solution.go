// https://codility.com/demo/results/trainingRJ4VV4-GWU/

package solution

import "sort"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	if len(A) == 3 {
		return A[0] * A[1] * A[2]
	}

	sort.Ints(A)

	max := maxProduct(A)
	for i := 1; i < len(A)-2; i++ {
		if maxp := maxProduct(A[i:]); maxp > max {
			max = maxp
		}
	}

	return max
}

func maxProduct(A []int) int {
	if A[0] < 0 && A[1] < 0 {
		return A[0] * A[1] * A[len(A)-1]
	} else {
		return A[0] * A[1] * A[2]
	}
}
