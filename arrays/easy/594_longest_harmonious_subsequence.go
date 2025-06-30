package easy

func findLHS(nums []int) int {

	m := make(map[int]int)
	for _, val := range nums {
		m[val]++
	}

	res := 0
	for key := range m {
		if _, exist := m[key+1]; exist {
			res = max(res, m[key]+m[key+1])
		}
	}

	return res
}
