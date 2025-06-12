package easy

import "math"

func maxAdjacentDistance(nums []int) int {
	n := len(nums)
	res := int(math.Abs(float64(nums[0]) - float64(nums[n-1])))

	for i := 0; i < n-1; i++ {
		res = int(math.Max(float64(res), float64(nums[i])-float64(nums[i+1])))
	}
	return res
}
