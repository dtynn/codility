// https://codility.com/demo/results/trainingSNQSHJ-DJT/
// http://stackoverflow.com/questions/4801242/algorithm-to-calculate-number-of-intersecting-discs

package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	starts := make([]int, len(A))
	stops := make([]int, len(A))

	for p, l := range A {
		start := p - l
		if start < 0 {
			start = 0
		}

		stop := p + l
		if stop > len(A)-1 {
			stop = len(A) - 1
		}

		starts[start] += 1
		stops[stop] += 1
	}

	count := 0
	activated := 0

	for i := 0; i < len(A); i++ {
		if starts[i] > 0 {
			count += activated*starts[i] + starts[i]*(starts[i]-1)/2
			if count > 10000000 {
				return -1
			}
			activated += starts[i]
		}

		activated -= stops[i]
	}

	return count
}
