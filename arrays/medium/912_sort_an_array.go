package medium

func sortArray(nums []int) []int {

	if len(nums) < 2 {
		return nums
	}

	mid := len(nums) / 2
	left := make([]int, mid)
	right := make([]int, len(nums)-mid)

	for i := 0; i < mid; i++ {
		left[i] = nums[i]
	}

	for i := mid; i < len(nums); i++ {
		right[i-mid] = nums[i]
	}

	return merge(sortArray(left), sortArray(right))
}

func merge(left []int, right []int) []int {

	leftLen := len(left)
	rightLen := len(right)
	result := make([]int, leftLen+rightLen)

	leftIndex, rightIndex, resultIndex := 0, 0, 0

	for leftIndex < leftLen && rightIndex < rightLen {
		if left[leftIndex] < right[rightIndex] {
			result[resultIndex] = left[leftIndex]
			leftIndex++
		} else {
			result[resultIndex] = right[rightIndex]
			rightIndex++
		}
		resultIndex++
	}

	// Copy remaining elements from left array
	for leftIndex < leftLen {
		result[resultIndex] = left[leftIndex]
		leftIndex++
		resultIndex++
	}

	// Copy remaining elements from right array
	for rightIndex < rightLen {
		result[resultIndex] = right[rightIndex]
		rightIndex++
		resultIndex++
	}

	return result
}
