package easy

import "strings"

func convertToTitle(columnNumber int) string {

	builder := strings.Builder{}

	for columnNumber > 0 {
		columnNumber--
		rem := columnNumber % 26
		builder.WriteByte(byte(rem + 'A'))
		columnNumber /= 26

	}
	return reverseString(builder.String())

}

func reverseString(s string) string {
	builder := strings.Builder{}

	for i := len(s) - 1; i >= 0; i-- {
		builder.WriteByte(s[i])
	}
	return builder.String()
}
