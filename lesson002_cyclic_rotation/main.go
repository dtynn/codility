package lesson2a

// https://codility.com/demo/results/trainingMSHKRD-KNY/
func Solution(A []int, K int) []int {
	if len(A) == 0 {
		return A
	}

	k := K % len(A)

	if k == 0 {
		return A
	}

	idx := len(A) - k
	return append(A[idx:], A[:idx]...)
}
