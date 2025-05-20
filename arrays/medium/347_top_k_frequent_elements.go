package medium

func topKFrequent(nums []int, k int) []int {
	if k == len(nums) {
		return nums
	}

	//Step 1 : Count the frequency of each element
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	//Step 2: Create buckets where index is frequency and value is the element
	freq := make([][]int, len(nums)+1)
	for num, freqCount := range count {
		freq[freqCount] = append(freq[freqCount], num)
	}

	res := make([]int, k)
	index := 0
	for i := len(freq) - 1; i >= 0; i-- {
		for _, num := range freq[i] {
			if index < k {
				res[index] = num
				index++
			}
			if index == k {
				break
			}
		}
	}

	return res

}
