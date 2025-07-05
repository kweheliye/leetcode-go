package easy

func findLucky(arr []int) int {
	m := make(map[int]int)
	for _, val := range arr {
		m[val]++
	}

	largest := -1
	for key, val := range m {
		if key == val {
			largest = max(largest, key)
		}
	}

	return largest
}
