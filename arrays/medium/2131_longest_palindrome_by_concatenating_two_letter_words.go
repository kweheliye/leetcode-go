package medium

// A Two-Dimensional Array Approach
func longestPalindromeV2(words []string) int {
	const alphabetSize = 26
	ans := 0
	freq := [alphabetSize][alphabetSize]int{}

	for _, word := range words {
		a := word[0] - 'a'
		b := word[1] - 'a'
		if freq[b][a] > 0 {
			freq[b][a]--
			ans += 4
		} else {
			freq[a][b]++
		}
	}

	for i := 0; i < alphabetSize; i++ {
		if freq[i][i] > 0 {
			return ans + 2
		}
	}

	return ans
}

// A Hash Map Approach
func longestPalindromeV1(words []string) int {
	count := make(map[string]int)
	for _, word := range words {
		count[word]++
	}

	var res int
	var central bool
	for key, value := range count {

		if key[0] == key[1] {
			if value%2 == 0 {
				res += value
			} else {
				res += value - 1
				central = true
			}
		} else if key[0] < key[1] {
			reversedWord := string(key[1]) + string(key[0])
			if val, ok := count[reversedWord]; ok {
				res += 2 * min(value, val)
			}
		}
	}
	if central {
		res++
	}
	return 2 * res
}
