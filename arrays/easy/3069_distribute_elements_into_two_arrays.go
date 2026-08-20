package easy

func resultArray(nums []int) []int {
	firstArray := make([]int, len(nums))
	secondArray := make([]int, len(nums))

	firstArray[0] = nums[0]
	secondArray[0] = nums[1]

	lastIndexFirstArray := 0
	lastIndexSecondArray := 0

	for i := 2; i < len(nums); i++ {

		if firstArray[lastIndexFirstArray] > secondArray[lastIndexSecondArray] {
			lastIndexFirstArray++
			firstArray[lastIndexFirstArray] = nums[i]
		} else {
			lastIndexSecondArray++
			secondArray[lastIndexSecondArray] = nums[i]
		}
	}

	for i := 0; i <= lastIndexSecondArray; i++ {
		lastIndexFirstArray++
		firstArray[lastIndexFirstArray] = secondArray[i]
	}

	return firstArray
}
