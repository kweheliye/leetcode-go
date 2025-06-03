package medium

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Map := make(map[rune]int)
	for _, char := range s1 {
		s1Map[char]++
	}

	for i := 0; i <= len(s2)-len(s1); i++ {
		s2Map := make(map[rune]int)

		for j := 0; j < len(s1); j++ {
			s2Map[rune(s2[i+j])]++
		}
		if matches(s1Map, s2Map) {
			return true
		}
	}

	return false

}

func matches(s1Map, s2Map map[rune]int) bool {
	for key, val := range s1Map {
		if val-s2Map[key] != 0 {
			return false
		}
	}
	return true
}
