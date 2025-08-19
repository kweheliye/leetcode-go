package medium

func zeroFilledSubarray(nums []int) int64 {
	ans, countSubArray := int64(0), 0
	for _, val := range nums {
		if val == 0 {
			countSubArray++
		} else {
			countSubArray = 0
		}
		ans += int64(countSubArray)
	}

	return ans
}
