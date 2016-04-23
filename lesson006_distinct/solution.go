// https://codility.com/demo/results/training42DP6R-NDC/

package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	m := map[int]int{}
	for _, num := range A {
		m[num] += 1
	}

	return len(m)
}
