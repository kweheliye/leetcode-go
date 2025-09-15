package easy

import (
	"strings"
)

func canBeTypedWords(text string, brokenLetters string) int {

	count := 0

	words := strings.Split(text, " ")

	for _, word := range words {
		canType := true
		for _, c := range word {
			if strings.ContainsRune(brokenLetters, c) {
				canType = false
				break
			}
		}
		if canType {
			count++
		}
	}

	return count
}
