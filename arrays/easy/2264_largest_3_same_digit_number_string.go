package easy

func largestGoodInteger(num string) string {
	sameDigitNumbers := []string{
		"999", "888", "777", "666", "555", "444", "333", "222", "111", "000"}

	for _, sameDigitNumber := range sameDigitNumbers {
		if contains(sameDigitNumber, num) {
			return sameDigitNumber
		}
	}
	return ""
}

func contains(sameDigitNumber string, num string) bool {
	for index := 0; index <= len(num)-3; index++ {
		if num[index] == sameDigitNumber[0] &&
			num[index+1] == sameDigitNumber[1] &&
			num[index+2] == sameDigitNumber[2] {
			return true
		}
	}
	return false
}
