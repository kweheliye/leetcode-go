package medium

import (
	"math"
)

// helper is a recursive function with memoization
func helper(remain int, coins []int, memo []int) int {
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
		count := helper(remain-coin, coins, memo) // try using this coin
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

func coinChangeTopDown(coins []int, amount int) int {
	memo := make([]int, amount+1)

	for i := range memo {
		memo[i] = -1
	}

	return helper(amount, coins, memo)
}

func coinChangeBottomUp(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for amnt := 1; amnt <= amount; amnt++ {
		for _, coin := range coins {
			if coin <= amnt {
				dp[amnt] = min(dp[amnt], dp[amnt-coin]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}
