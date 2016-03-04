// https://codility.com/demo/results/training4QRJS5-9PF/
package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(X int, A []int) int {
	// write your code in Go 1.4
	falls := make([]int, X)

	for sec, pos := range A {
		if pos > X {
			continue
		}

		if falls[pos-1] == 0 || falls[pos-1] > sec {
			falls[pos-1] = sec
		}
	}

	maxSec := 0
	for posIdx, sec := range falls {
		if sec == 0 && A[0] != posIdx+1 {
			return -1
		}

		if sec > maxSec {
			maxSec = sec
		}
	}

	return maxSec
}
