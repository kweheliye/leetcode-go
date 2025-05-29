package medium

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)

	for _, num := range nums {
		numSet[num] = true
	}

	longestStreak := 0

	for num := range numSet {

		if !numSet[num-1] {
			currentStreak := 1
			currentNum := num + 1
			for numSet[currentNum] {
				currentStreak++
				currentNum++
			}
			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		}
	}
	return longestStreak
}
