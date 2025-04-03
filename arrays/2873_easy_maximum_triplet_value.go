package arrays

func MaximumTripletValueBruteForce(nums []int) int64 {
	var n int = len(nums)
	var res int64 = 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				res = max(res, int64(nums[i]-nums[j])*int64(nums[k]))
			}
		}
	}

	return res
}

func MaximumTripletValuePrefixSuffixArray(nums []int) int64 {
	var n int = len(nums)
	var res int64 = 0
	left := nums[0]

	for j := 1; j < n; j++ {
		left = max(left, nums[j])

		for k := j + 1; k < n; k++ {
			res = max(res, int64(left-nums[j])*int64(nums[k]))
		}

	}

	return res
}
