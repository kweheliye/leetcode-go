package medium

import "math"

func getMaxInt(ribbons []int) int {
	maxInt := math.MinInt32
	for i := 0; i < len(ribbons); i++ {
		if ribbons[i] > maxInt {
			maxInt = ribbons[i]
		}
	}
	return maxInt
}

func numOfBlocks(ribbons []int, m int) int {
	count := 0
	for i := 0; i < len(ribbons); i++ {
		count += ribbons[i] / m
	}
	return count
}

func maxLength(ribbons []int, k int) int {
	left, result := 1, 0
	right := getMaxInt(ribbons)

	for left <= right {
		mid := left + (right-left)/2
		if numOfBlocks(ribbons, mid) >= k {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}
