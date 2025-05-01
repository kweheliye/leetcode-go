package medium

import (
	"math"
)

// IN PROGRESS ADD MEMO ARRAY
func coinChange(coins []int, amount int) int {
	return helper(coins, amount)
}

func helper(coins []int, remain int) int {
	if remain == 0 {
		return 0
	}
	if remain < 0 {
		return -1
	}

	minCount := math.MaxInt32

	for _, coin := range coins {
		count := helper(coins, remain-coin)
		if count == -1 {
			continue
		}
		minCount = min(minCount, count+1)
	}

	if minCount == math.MaxInt32 {
		return -1
	}

	return minCount

}
