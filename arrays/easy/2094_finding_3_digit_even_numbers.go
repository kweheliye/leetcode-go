package easy

import "sort"

func findEvenNumbers(digits []int) []int {
	numSet := make(map[int]bool)
	n := len(digits)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if i == j || i == k || j == k {
					continue
				}
				num := digits[i]*100 + digits[j]*10 + digits[k]
				if num%2 == 0 && num >= 100 {
					numSet[num] = true
				}
			}
		}
	}

	res := make([]int, 0, len(numSet))

	for num := range numSet {
		res = append(res, num)
	}

	sort.Ints(res)
	return res
}
