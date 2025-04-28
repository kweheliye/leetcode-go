package medium

func countCompleteSubarrays(nums []int) int {

	//Step 1 count total distinct elements
	unique := make(map[int]bool)
	for _, num := range nums {
		unique[num] = true
	}
	totalDistinct := len(unique)

	//Step 2 sliding window logic
	left := 0
	count := 0
	freqMap := make(map[int]int)
	n := len(nums)

	for right := 0; right < n; right++ {
		freqMap[nums[right]]++

		for len(freqMap) == totalDistinct {
			count += n - right

			//shrink the window
			freqMap[nums[left]]--
			if freqMap[nums[left]] == 0 {
				delete(freqMap, nums[left])
			}
			left++
		}
	}

	return count

}
