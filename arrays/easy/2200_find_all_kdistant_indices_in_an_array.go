package easy

func findKDistantIndices(nums []int, key int, k int) []int {
	var res []int

	n := len(nums)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if nums[i] == key && abs(i-j) <= k {
				res = append(res, i, j)
			}
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
