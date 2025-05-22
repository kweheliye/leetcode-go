package hard

func firstMissingPositive(nums []int) int {
	n := len(nums)
	seen := make([]bool, n+1)

	for _, num := range nums {
		if num > 0 && num <= n {
			seen[num] = true
		}
	}

	for i := 1; i <= n; i++ {
		if !seen[i] {
			return i
		}
	}

	return n + 1
}
