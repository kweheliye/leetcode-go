package easy

func firstPalindrome(words []string) string {
	for _, word := range words {
		if isPalindrome(word) {
			return word
		}
	}
	return ""
}
func isPalindrome(s string) bool {

	for left, right := 0, len(s)-1; left < right; {
		if s[left] != s[right] {
			return false
		}
		left, right = left+1, right-1
	}
	return true
}
