package easy

func romanToInt(s string) int {
	values := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sum := 0
	n := len(s)

	for i := 0; i < n; i++ {
		currValue := values[rune(s[i])]

		if i+1 < n && currValue < values[rune(s[i+1])] {
			sum -= currValue
		} else {
			sum += currValue
		}

	}
	return sum
}
