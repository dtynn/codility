// https://codility.com/demo/results/trainingZGQYSA-WE6/
package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	countZero := 0
	countPairs := 0

	for _, num := range A {
		if num == 0 {
			countZero += 1
		} else {
			countPairs += countZero
		}

		if countPairs > 1000000000 {
			return -1
		}
	}

	return countPairs
}
