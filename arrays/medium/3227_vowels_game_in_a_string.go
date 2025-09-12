package medium

func doesAliceWin(s string) bool {
	for _, c := range s {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			return true
		}
	}
	return false
}
