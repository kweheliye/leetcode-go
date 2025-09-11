package medium

import (
	"strings"
)

func sortVowels(s string) string {
	count := make([]int, 128)

	// Count vowels
	for _, c := range s {
		if isVowel(c) {
			count[c-'A']++
		}
	}

	sortedVowel := "AEIOUaeiou"
	var ans strings.Builder
	j := 0

	for _, c := range s {
		if !isVowel(c) {
			ans.WriteRune(c)
		} else {
			for count[rune(sortedVowel[j])-'A'] == 0 {
				j++
			}
			ans.WriteByte(sortedVowel[j])
			count[rune(sortedVowel[j])-'A']--
		}
	}

	return ans.String()
}

func isVowel(c rune) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
		c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}
