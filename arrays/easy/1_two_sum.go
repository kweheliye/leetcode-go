package easy

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if j, ok := m[complement]; ok {
			return []int{j, i}
		}
		m[nums[i]] = i

	}
	return []int{}
}
