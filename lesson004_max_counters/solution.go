// https://codility.com/demo/results/trainingW8SH5N-BTU/

package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int, A []int) []int {
	// write your code in Go 1.4
	setAllTo := 0
	maxAdded := 0

	counters := make([]int, N)
	zeros := make([]int, N)
	for _, num := range A {
		if num >= 1 && num <= N {
			counters[num-1] = counters[num-1] + 1
			if counters[num-1] > maxAdded {
				maxAdded = counters[num-1]
			}
		} else {
			if maxAdded != 0 {
				copy(counters, zeros)
				setAllTo = setAllTo + maxAdded
				maxAdded = 0
			}
		}
	}

	for i := 0; i < N; i += 1 {
		counters[i] += setAllTo
	}

	return counters
}
