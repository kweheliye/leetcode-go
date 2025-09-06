package hard

func trap(height []int) int {
	n := len(height)
	maxLeft, maxRight := make([]int, n), make([]int, n)

	maxLeft[0] = 0
	for i := 1; i < n; i++ {
		maxLeft[i] = max(maxLeft[i-1], height[i-1])
	}

	maxRight[n-1] = 0
	for i := n - 2; i >= 0; i-- {
		maxRight[i] = max(maxRight[i+1], height[i+1])
	}

	res := 0

	for i := 0; i < n; i++ {
		water := min(maxLeft[i], maxRight[i]) - height[i]
		if water > 0 {
			res += water
		}
	}

	return res

}
