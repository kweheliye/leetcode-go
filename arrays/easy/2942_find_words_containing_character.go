package easy

import "strings"

func findWordsContaining(words []string, x byte) []int {
	result := make([]int, 0)

	for i, word := range words {

		if strings.ContainsRune(word, rune(x)) {
			result = append(result, i)
		}

	}
	return result
}
