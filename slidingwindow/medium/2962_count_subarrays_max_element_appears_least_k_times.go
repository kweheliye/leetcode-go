package medium

func countSubarrays(nums []int, k int) int64 {
	ans := int64(0)

	maxElement := nums[0]
	for _, num := range nums {
		if num > maxElement {
			maxElement = num
		}
	}
	maxElementInWindow := 0
	left := 0
	n := len(nums)

	for right := 0; right < n; right++ {
		if nums[right] == maxElement {
			maxElementInWindow++
		}

		for maxElementInWindow == k {
			ans += int64(n - right)
			if nums[left] == maxElement {
				maxElementInWindow--
			}
			left++
		}
	}
	return ans
}
