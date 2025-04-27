package easy

func countSubarraysV1(nums []int) int {
	second, first, count := 0, 0, 0

	for _, num := range nums {
		third := num
		sum := first + third

		if sum*2 == second {
			count++
		}

		temp := second
		second = third
		first = temp
	}

	return count
}

func countSubarraysV2(nums []int) int {
	ans := 0
	n := len(nums)

	for i := 1; i < n-1; i++ {
		if nums[i] == nums[i-1]+nums[i+1] {
			ans++
		}
	}
	return ans
}
