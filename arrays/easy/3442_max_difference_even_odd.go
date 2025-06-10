package easy

func maxDifference(s string) int {
	freq := [26]int{}

	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	maxEven, maxOdd := len(s), 0

	for _, value := range freq {
		if value == 0 {
			continue
		} else if value%2 == 0 {
			maxEven = min(maxEven, value)
		} else {
			maxOdd = max(maxOdd, value)
		}
	}
	return maxOdd - maxEven
}
