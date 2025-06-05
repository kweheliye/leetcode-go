package easy

import "strings"

func addStrings(num1 string, num2 string) string {
	res := strings.Builder{}
	carry := 0

	p1, p2 := len(num1)-1, len(num2)-1
	for p1 >= 0 || p2 >= 0 {
		x1, x2 := 0, 0

		if p1 >= 0 {
			x1 = int(num1[p1]) - '0'
		}

		if p2 >= 0 {
			x2 = int(num2[p2]) - '0'
		}
		value := (x1 + x2 + carry) % 10
		carry = (x1 + x2 + carry) / 10
		res.WriteByte(byte(value + '0'))

		p1--
		p2--
	}

	if carry != 0 {
		res.WriteByte(byte(carry + '0'))
	}

	return reverse(res.String())

}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
