package medium

import "sort"

func partitionArray(nums []int, k int) int {
	sort.Ints(nums)
	ans := 1
	rec := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i]-rec > k {
			ans++
			rec = nums[i]
		}
	}

	return ans

}
