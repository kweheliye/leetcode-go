package easy

import "strings"

func makeFancyString(s string) string {

	var ans strings.Builder

	ans.WriteByte(s[0])
	ans.WriteByte(s[1])

	for i := 2; i < len(s); i++ {
		if s[i] == s[i-1] || s[i] == s[i-2] {
			ans.WriteByte(s[i])
		}
	}

	return ans.String()
}
