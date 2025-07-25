package easy

//func maxSum(nums []int) int {
//	positiveNumSet := make(map[int]bool)
//	maxNum := nums[0]
//	for _, num := range nums {
//		if num > 0 {
//			positiveNumSet[num] = true
//		}
//		maxNum = max(maxNum, num)
//	}
//
//	if len(positiveNumSet) == 0 {
//		return maxNum
//	}
//
//	sum := 0
//	for num := range positiveNumSet {
//		sum += num
//	}
//	return sum
//}

func maxSum(nums []int) int {
	positiveNumsSet := make(map[int]bool)
	maxNum := nums[0]
	for _, num := range nums {
		if num > 0 {
			positiveNumsSet[num] = true
		}
		maxNum = max(maxNum, num)
	}

	if len(positiveNumsSet) == 0 {
		return maxNum
	}
	sum := 0
	for num := range positiveNumsSet {
		sum += num
	}
	return sum
}
