package medium

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
	ans := make([]int, len(spells))
	m := len(potions)
	sort.Ints(potions)

	for i, spell := range spells {
		j := binarySearch(potions, spell, success)
		ans[i] = m - j
	}
	return ans
}

func binarySearch(arr []int, spell int, target int64) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2
		if int64(arr[mid]*spell) >= target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}
