package easy

func buildArray(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		nums[i] += 1000 * (nums[nums[i]] % 1000)
	}

	for i := 0; i < n; i++ {
		nums[i] /= 1000
	}

	return nums
}
