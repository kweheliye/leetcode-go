package medium

import "sort"

func divideArray(nums []int, k int) [][]int {
	sort.Ints(nums)

	res := make([][]int, len(nums)/3)

	for i := 0; i < len(nums); i += 3 {
		if nums[i+2]-nums[i] > k {
			return [][]int{}
		}
		res[i/3] = append(res[i/3], nums[i], nums[i+1], nums[i+2])
	}
	return res

}
