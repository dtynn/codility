// https://codility.com/demo/results/trainingDJFJ8K-R8E/
// 尝试证明 avg 最小的 slice 长度为2或者3

package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	avgs := make([]float64, len(A)-1)
	for i := 0; i < len(avgs)-1; i += 1 {
		avg2 := float64((A[i] + A[i+1])) / 2.0
		avg3 := float64((A[i] + A[i+1] + A[i+2])) / 3.0
		if avg2 < avg3 {
			avgs[i] = avg2
		} else {
			avgs[i] = avg3
		}
	}

	avgs[len(avgs)-1] = float64(A[len(avgs)-1]+A[len(avgs)]) / 2.0

	minAvg := avgs[0]
	minIdx := 0

	for idx, avg := range avgs {
		if avg < minAvg {
			minAvg = avg
			minIdx = idx
		}
	}

	return minIdx
}
