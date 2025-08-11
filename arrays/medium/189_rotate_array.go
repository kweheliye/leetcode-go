package medium

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n

	for i := 0; i < k; i++ {
		nums = append(nums[1:], nums[0])
	}

}
