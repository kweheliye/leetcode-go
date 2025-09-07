package easy

func sumZero(n int) []int {
	res := make([]int, n)
	index := 0

	for i := 1; i <= n/2; i++ {
		res[index] = i
		res[index+1] = -i
		index += 2
	}

	return res
}
