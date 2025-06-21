package medium

func minimumDeletions(word string, k int) int {
	freqMap := make(map[rune]int)
	for _, char := range word {
		freqMap[char]++
	}

	var ans = len(word)
	for _, freqBase := range freqMap {
		var deletions = 0
		for _, freq := range freqMap {
			if freq < freqBase {
				deletions += freq
			} else if freq > freqBase+k { //we need to trim it down
				deletions += freq - freqBase - k
			}
		}
		ans = min(ans, deletions)
	}
	return ans
}
