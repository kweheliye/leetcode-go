package easy

import "strings"

func divideString(s string, k int, fill byte) []string {

	res := make([]string, 0)

	n := len(s)
	curr := 0

	for curr < n {
		end := min(curr+k, n)
		res = append(res, s[curr:end])
		curr += k
	}

	last := res[len(res)-1]

	if len(last) < k {
		last += strings.Repeat(string(fill), k-len(last))
		
	}

	return res
}
