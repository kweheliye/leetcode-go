package medium

func lengthOfLongestSubstring(s string) int {
	seen := make(map[byte]bool)
	result, left := 0, 0

	for right := 0; right < len(s); right++ {

		for seen[s[right]] {
			delete(seen, s[left])
			left++
		}
		seen[s[right]] = true
		result = max(result, right-left+1)
	}

	return result
}
