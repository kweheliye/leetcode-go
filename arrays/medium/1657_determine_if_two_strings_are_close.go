package medium

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	word1Map := make(map[rune]int)
	word2Map := make(map[rune]int)
	for _, char := range word1 {
		word1Map[char]++
	}
	for _, char := range word2 {
		word2Map[char]++
	}

	// Check if both maps have the same set of keys
	if len(word1Map) != len(word2Map) {
		return false
	}
	for char := range word1Map {
		if _, exists := word2Map[char]; !exists {
			return false
		}
	}

	// Check if the frequency distribution is the same
	freqMap1 := make(map[int]int)
	freqMap2 := make(map[int]int)
	for _, count := range word1Map {
		freqMap1[count]++
	}
	for _, count := range word2Map {
		freqMap2[count]++
	}

	for freq, count := range freqMap1 {
		if freqMap2[freq] != count {
			return false
		}
	}

	return true
}
