package medium

func minSum(nums1 []int, nums2 []int) int64 {
	sum1, sum2 := int64(0), int64(0)
	zero1, zero2 := 0, 0

	for _, num := range nums1 {
		sum1 += int64(num)

		if num == 0 {
			zero1++
			sum1++
		}
	}

	for _, num := range nums2 {
		sum2 += int64(num)

		if num == 0 {
			zero2++
			sum2++
		}
	}

	if zero1 == 0 && sum2 > sum1 || zero2 == 0 && sum1 > sum2 {
		return -1
	}
	
	return max(sum1, sum2)

}
