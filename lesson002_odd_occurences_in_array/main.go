package lesson2b

// https://codility.com/demo/results/trainingPHJ4PJ-NTP/
func Solution(A []int) int {
	// write your code in Go 1.4
	countMap := map[int]int{}

	for _, num := range A {
		countMap[num] += 1
	}

	for n, count := range countMap {
		if count == 1 {
			return n
		}
	}

	return 0
}
