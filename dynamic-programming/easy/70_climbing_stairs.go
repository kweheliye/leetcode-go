package easy

// Memoization, Top Down
func climbStairsTopDown(n int) int {
	memo := make([]int, n+1)

	var helper func(int, []int) int
	helper = func(k int, m []int) int {
		if k <= 1 {
			return 1
		}

		if m[k] > 0 {
			return m[k]
		}

		result := helper(k-1, m) + helper(k-2, m)
		m[k] = result
		return result
	}

	return helper(n, memo)
}

// BottomUp
func climbStairsBottomUp(n int) int {
	memo := make([]int, n+1)
	memo[0], memo[1] = 1, 1

	for i := 2; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[n]
}
