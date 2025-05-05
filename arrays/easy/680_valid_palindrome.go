package easy

func validPalindrome(s string) bool {
	p2 := len(s) - 1
	p1 := 0
	for p1 < p2 {
		if s[p1] != s[p2] {
			return check(p1+1, p2, s) || check(p1, p2-1, s)
		}

		p1++
		p2--
	}

	return true
}

func check(left, right int, s string) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}
