package medium

func majorityElement(nums []int) []int {
	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}
	result := make([]int, 0)

	for num, count := range freq {
		if count > len(nums)/3 {
			result = append(result, num)
		}
	}

	return result
}
