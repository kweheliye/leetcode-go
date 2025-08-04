package medium

func totalFruit(fruits []int) int {
	basket := make(map[int]int)
	total, left := 0, 0

	for right := 0; right < len(fruits); right++ {
		basket[fruits[right]]++
		for len(basket) > 2 {
			basket[fruits[left]]--
			if basket[fruits[left]] == 0 {
				delete(basket, fruits[left])
			}
			left++
		}

		total = max(total, right-left+1)
	}
	return total
}
