package dfs_iterative

func lexicalOrder(n int) []int {
	lexicographicalNumbers := make([]int, 0)
	currentNumber := 1

	for i := 0; i < n; i++ {
		lexicographicalNumbers = append(lexicographicalNumbers, currentNumber)
		if currentNumber*10 < n {
			currentNumber *= 10
		} else {
			for currentNumber%10 == 9 || currentNumber >= n {
				currentNumber /= 10
			}
			currentNumber++
		}
	}
	return lexicographicalNumbers
}
