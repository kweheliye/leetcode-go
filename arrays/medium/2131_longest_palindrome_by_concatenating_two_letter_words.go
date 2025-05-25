package medium

func longestPalindrome(words []string) int {
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
