package twopointers

func sortColors(nums []int) {

	lowPointer, curr, highPointer := 0, 0, len(nums)-1

	for curr <= highPointer {
		if nums[curr] == 0 {
			nums[lowPointer], nums[curr] = nums[curr], nums[lowPointer]
			curr++
			lowPointer++
		} else if nums[curr] == 2 {
			nums[curr], nums[highPointer] = nums[highPointer], nums[curr]
			highPointer--
		} else {
			curr++
		}

	}
}
