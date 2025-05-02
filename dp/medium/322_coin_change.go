package medium

import (
	"math"
)

func coinChange(coins []int, amount int) int {
	memo := make([]int, amount+1)

	for i := range memo {
		memo[i] = -1
	}

	// helper is a recursive function with memoization
	var helper func(int) int
	helper = func(remain int) int {
		if remain == 0 {
			return 0
		}
		if remain < 0 {
			return -1 // invalid case
		}
		if memo[remain] != -1 {
			return memo[remain] // return already computed value
		}

		var minCount = math.MaxInt32
		for _, coin := range coins {
			count := helper(remain - coin) // try using this coin
			if count == -1 {
				continue // skip invalid subproblem
			}
			minCount = min(minCount, count+1)
		}

		// Save the result in memo
		if minCount == math.MaxInt32 {
			return -1
		}
		memo[remain] = minCount
		return memo[remain]
	}

	return helper(amount)
}
