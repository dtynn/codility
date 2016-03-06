// https://codility.com/demo/results/trainingW8MECT-S42/

package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A int, B int, K int) int {
	// write your code in Go 1.4
	count := 0
	startN := A/K + 1
	if A%K == 0 {
		count += 1
	}

	endN := B / K
	return count + (endN - startN) + 1
}
