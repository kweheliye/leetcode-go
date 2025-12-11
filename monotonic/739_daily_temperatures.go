package monotonic

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int, n)
	stack := []int{} // this will store indexes

	for i := 0; i < n; i++ {
		// while stack not empty AND current temp > temp at index on stack top
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			j := stack[len(stack)-1]     // top index
			stack = stack[:len(stack)-1] // pop
			result[j] = i - j
		}
		stack = append(stack, i) // push index
	}

	return result
}
