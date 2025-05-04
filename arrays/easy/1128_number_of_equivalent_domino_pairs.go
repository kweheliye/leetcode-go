package easy

func numEquivDominoPairs(dominoes [][]int) int {
	count := make(map[int]int)

	for _, domino := range dominoes {
		a := max(domino[0], domino[1])
		b := min(domino[1], domino[0])

		key := a*10 + b
		count[key]++
	}

	totalPairs := 0

	for _, n := range count {
		if n > 1 {
			totalPairs += n * (n - 1) / 2
		}
	}

	return totalPairs
}
