package easy

func containsNearbyDuplicate(nums []int, k int) bool {
	window := make(map[int]bool)

	left := 0

	for right := 0; right < len(nums); right++ {
		if right-left > k {
			delete(window, nums[left])
			left++
		}

		if _, exist := window[nums[right]]; exist {
			return true
		}

		window[nums[right]] = true

	}

	return false
}
