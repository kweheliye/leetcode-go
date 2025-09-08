package easy

import (
	"strconv"
	"strings"
)

func getNoZeroIntegers(n int) []int {
	for i := 0; i < n; i++ {
		b := n - i
		if !strings.Contains(strconv.Itoa(b), "0") && !strings.Contains(strconv.Itoa(i), "0") {
			return []int{i, b}
		}
	}
	return []int{}
}
