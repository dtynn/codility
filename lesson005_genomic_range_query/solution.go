// https://codility.com/demo/results/trainingTD3JF2-EH9/

package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(S string, P []int, Q []int) []int {
	// write your code in Go 1.4
	factors := map[rune]int{
		'A': 1,
		'C': 2,
		'G': 3,
		'T': 4,
	}

	counts := make([][5]int, len(S)+1)

	for idx, one := range S {
		factor := factors[one]
		counts[idx+1][1] = counts[idx][1]
		counts[idx+1][2] = counts[idx][2]
		counts[idx+1][3] = counts[idx][3]
		counts[idx+1][4] = counts[idx][4]
		counts[idx+1][factor] += 1
	}

	M := len(P)
	result := make([]int, M)
	for i := 0; i < M; i += 1 {
		if counts[Q[i]+1][1]-counts[P[i]][1] > 0 {
			result[i] = 1
			continue
		}

		if counts[Q[i]+1][2]-counts[P[i]][2] > 0 {
			result[i] = 2
			continue
		}

		if counts[Q[i]+1][3]-counts[P[i]][3] > 0 {
			result[i] = 3
			continue
		}

		result[i] = 4
	}

	return result
}
