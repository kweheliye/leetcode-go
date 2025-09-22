package easy

// One pass , Count Frequency and Max Frequency in same loop
func maxFrequencyElementsV1(nums []int) int {

	freqMap := make(map[int]int)
	maxFreq, totalFreq := 0, 0

	for _, num := range nums {
		freq := freqMap[num] + 1
		freqMap[num] = freq

		if freq > maxFreq {
			maxFreq = freq
			totalFreq = freq
		} else if freq == maxFreq {
			totalFreq += freq
		}
	}

	return totalFreq
}

// Count Frequency and Max Frequency in seprate loop
func maxFrequencyElementsV2(nums []int) int {

	freqMap := make(map[int]int)

	for _, num := range nums {
		freqMap[num]++
	}

	maxFreq := 0

	for _, val := range freqMap {
		if val > maxFreq {
			maxFreq = val
		}
	}

	totalFreq := 0

	for _, val := range freqMap {
		if val == maxFreq {
			totalFreq++
		}
	}

	return totalFreq * maxFreq
}
