package medium

import "math"

func minimumDeletions2091(nums []int) int {
	minVal, maxVal := math.MaxInt, math.MinInt
	minPosition, maxPosition := 0, 0

	for i, num := range nums {
		if num < minVal {
			minVal = num
			minPosition = i
		}

		if num > maxVal {
			maxVal = num
			maxPosition = i
		}

	}

	n := len(nums)

	// Remove both elements from the front.
	fromFront := max(minPosition, maxPosition) + 1

	// Remove both elements from the back.
	fromBack := n - min(minPosition, maxPosition)

	fromBothSides := min(
		(minPosition+1)+(n-maxPosition),
		(maxPosition+1)+(n-minPosition),
	)

	return min(fromFront, fromBack, fromBothSides)
}
