package easy

func numOfUnplacedFruitsV1(fruits []int, baskets []int) int {
	ans := len(fruits)
	for _, fruit := range fruits {
		for i := 0; i < len(baskets); i++ {
			if fruit <= baskets[i] {
				baskets[i] = 0
				ans--
				break
			}
		}
	}
	return ans
}

func numOfUnplacedFruitsV2(fruits []int, baskets []int) int {
	used := make([]bool, len(fruits))
	unplaced := 0

	for i := 0; i < len(fruits); i++ {
		placed := false
		for j := 0; j < len(baskets); j++ {
			if !used[j] && fruits[i] <= baskets[j] {
				used[j] = true
				placed = true
				break
			}
		}
		if !placed {
			unplaced++
		}

	}
	return unplaced
}
