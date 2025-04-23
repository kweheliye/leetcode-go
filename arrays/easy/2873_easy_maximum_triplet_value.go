package easy

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
	var res int64 = 0
	var n int = len(nums)
	leftMax := make([]int, n)
	rightMax := make([]int, n)

	// Build leftMax: Maximum from index 0 to i
	leftMax[0] = nums[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], nums[i])
	}

	// Build rightMax: Maximum from index i+1 to n-1
	rightMax[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], nums[i]) // Fixed: Use nums[i] instead of nums[i+1]
	}

	// Compute maximum triplet value
	for j := 1; j < n-1; j++ {
		res = max(res, int64(leftMax[j-1]-nums[j])*int64(rightMax[j+1]))
	}

	return res
}
