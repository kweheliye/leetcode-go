package medium

func longestSubarray(nums []int) int {
	left, longestWindow, zeroCount := 0, 0, 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroCount++
		}
		for zeroCount > 1 {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		longestWindow = max(longestWindow, right-left)
	}

	return longestWindow

}
