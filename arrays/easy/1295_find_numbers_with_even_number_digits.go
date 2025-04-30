package easy

import "strconv"

func findNumbersV1(nums []int) int {
	count := 0
	for _, num := range nums {
		length := len(strconv.Itoa(num))

		if length%2 == 0 {
			count++
		}
	}
	return count
}

func findNumbersV2(nums []int) int {
	count := 0
	for _, num := range nums {
		if hasEvenDigits(num) {
			count++
		}
	}
	return count
}

func hasEvenDigits(num int) bool {
	digitCount := 0
	for num > 0 {
		digitCount++
		num /= 10
	}
	return digitCount%2 == 0
}
