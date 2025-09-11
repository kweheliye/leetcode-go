package sorting

func countingSort(array []int) []int {
	output := make([]int, len(array))
	maxElement := array[0]

	// Find the largest element of the array
	for _, num := range array {
		if num > maxElement {
			maxElement = num
		}
	}
	count := make([]int, maxElement+1)

	// Store the count of each element
	for _, num := range array {
		count[num]++
	}

	// Store the cumulative count of each array
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	// Find the index of each element of the original array in count array, and
	// place the elements in output array
	for i := len(array) - 1; i >= 0; i-- {
		output[count[array[i]]-1] = array[i]
		count[array[i]]--
	}

	// Copy the sorted elements into original array
	for i := 0; i < len(array); i++ {
		array[i] = output[i]
	}
	return array
}
