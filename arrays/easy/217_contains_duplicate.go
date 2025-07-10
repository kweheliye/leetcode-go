package easy

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)

	for _, num := range nums {
		if _, exist := m[num]; exist {
			return true
		}
		m[num] = true
	}

	return false
}
