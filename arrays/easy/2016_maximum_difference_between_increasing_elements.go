package easy

func maximumDifference(nums []int) int {
	ans := -1
	for i, preMin := 1, nums[0]; i < len(nums); i++ {
		if nums[i] > preMin {
			ans = max(ans, nums[i]-preMin)
		} else {
			preMin = nums[i]
		}
	}
	return ans
}
