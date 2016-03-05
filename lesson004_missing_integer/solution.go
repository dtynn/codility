// https://codility.com/demo/results/training4RK7TT-F3Y/

package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	counts := map[int]int{}
	for _, num := range A {
		if num > 0 {
			counts[num] += 1
		}
	}

	for n := 1; n <= len(counts); n += 1 {
		if _, ok := counts[n]; !ok {
			return n
		}
	}

	return len(counts) + 1
}
