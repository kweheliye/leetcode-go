package medium

func maximumUniqueSubarray(nums []int) int {
	set := make(map[int]bool)
	sum, maxSum, left := 0, 0, 0

	for right := 0; right < len(nums); right++ {
		for set[nums[right]] {
			delete(set, nums[left])
			sum -= nums[left]
			left++
		}
		set[nums[right]] = true
		sum += nums[right]
		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}
