// https://codility.com/demo/results/trainingQBMUXY-ME8/
package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	n := len(A) + 1
	expected := (1 + n) * n / 2

	sum := 0
	for _, num := range A {
		sum += num
	}

	return expected - sum
}
