package medium

func subarraySum(nums []int, k int) int {
	prefixSum := make(map[int]int, len(nums)+1)
	prefixSum[0] = 1

	currSum, total := 0, 0

	for i := 0; i < len(nums); i++ {
		currSum += nums[i]

		if val, exists := prefixSum[currSum-k]; exists {
			total += val
		}
		prefixSum[currSum]++

	}

	return total
}
