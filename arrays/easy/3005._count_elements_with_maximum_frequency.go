package easy

func maxFrequencyElements(nums []int) int {

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
