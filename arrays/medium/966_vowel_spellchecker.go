package medium

import "strings"

var exactWordsSet map[string]bool
var caseInsensitiveMap map[string]string
var vowelErrorMap map[string]string

func spellchecker(wordlist []string, queries []string) []string {

	exactWordsSet = make(map[string]bool)
	caseInsensitiveMap = make(map[string]string)
	vowelErrorMap = make(map[string]string)

	for _, word := range wordlist {
		exactWordsSet[word] = true

		lowerCaseWord := strings.ToLower(word)
		if _, exist := caseInsensitiveMap[lowerCaseWord]; !exist {
			caseInsensitiveMap[lowerCaseWord] = word
		}

		devoweledWord := devowel(lowerCaseWord)
		if _, exist := vowelErrorMap[devoweledWord]; !exist {
			vowelErrorMap[devoweledWord] = word
		}
	}

	var ans []string

	for _, query := range queries {
		ans = append(ans, solve(query))
	}
	return ans

}

func solve(query string) string {
	if _, exist := exactWordsSet[query]; exist {
		return query
	}

	lowerCaseWord := strings.ToLower(query)

	if word, exist := caseInsensitiveMap[lowerCaseWord]; exist {
		return word
	}

	devoweledWord := devowel(lowerCaseWord)
	if word, exist := vowelErrorMap[devoweledWord]; exist {
		return word
	}

	return ""

}

func devowel(word string) string {
	ans := strings.Builder{}

	for _, c := range word {
		if isVowelByte(byte(c)) {
			ans.WriteByte('*')
		} else {
			ans.WriteByte(byte(c))
		}
	}
	return ans.String()
}
