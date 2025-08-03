package medium

func majorityElementFrequencyCount(nums []int) []int {
	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}
	result := make([]int, 0)

	for num, count := range freq {
		if count > len(nums)/3 {
			result = append(result, num)
		}
	}

	return result
}

// Boyer-Moore Voting Algorithm
func majorityElementBoyerMooreVotingAlgorithm(nums []int) []int {
	candidate1, candidate2, count1, count2 := 0, 0, 0, 0

	for _, num := range nums {
		if candidate1 == num {
			count1++
		} else if candidate2 == num {
			count2++
		} else if count1 == 0 {
			candidate1 = num
			count1 = 1
		} else if count2 == 0 {
			candidate2 = num
			count2 = 1
		} else {
			count1--
			count2--
		}
	}

	count1, count2 = 0, 0
	for _, num := range nums {
		if candidate1 == num {
			count1++
		} else if candidate2 == num {
			count2++
		}
	}

	var result []int
	if count1 > len(nums)/3 {
		result = append(result, candidate1)
	}
	if count2 > len(nums)/3 {
		result = append(result, candidate2)
	}

	return result

}
