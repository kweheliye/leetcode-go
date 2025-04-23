package easy

func countLargestGroup(n int) int {
	hashMap := make(map[int]int)
	var maxValue int = 0

	for i := 1; i <= n; i++ {
		key, i0 := 0, i
		for i0 != 0 {
			key += i0 % 10
			i0 /= 10
		}
		hashMap[key]++
		maxValue = max(maxValue, hashMap[key])
	}
	count := 0
	for _, val := range hashMap {
		if val == maxValue {
			count++
		}
	}
	return count
}
