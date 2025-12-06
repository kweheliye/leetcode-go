package medium

func maxOperations(nums []int, k int) int {
	pairs := 0

	freq := make(map[int]int)

	for _, num := range nums {
		if freq[k-num] > 0 {
			freq[k-num]--
			pairs++
		} else {
			freq[num]++
		}
	}
	return pairs
}
