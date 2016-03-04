// https://codility.com/demo/results/trainingMJ7KKU-WAD/
package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	N := len(A)
	count := map[int]int{}

	for _, num := range A {
		if num > N {
			return 0
		}

		count[num] += 1
		if count[num] > 1 {
			return 0
		}
	}

	return 1
}
