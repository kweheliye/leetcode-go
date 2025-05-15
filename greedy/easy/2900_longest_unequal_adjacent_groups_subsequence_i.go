package easy

func getLongestSubsequence(words []string, groups []int) []string {
	var result []string

	result = append(result, words[0])

	for i := 1; i < len(groups); i++ {

		if groups[i] != groups[i-1] {
			result = append(result, words[i])
		}
	}

	return result

}
