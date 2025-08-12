package medium

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	ans, left, right := 0, 0, len(people)-1

	for left <= right {
		if people[left]+people[right] <= limit {
			left++
		}
		right--
		ans++
	}

	return ans
}
