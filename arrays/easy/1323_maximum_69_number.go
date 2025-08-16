package easy

import "strconv"

func maximum69Number(num int) int {
	str := strconv.Itoa(num)
	bytes := []byte(str)

	for i := 0; i < len(bytes); i++ {

		if bytes[i] == '6' {
			bytes[i] = '9'
			break
		}
	}

	result, _ := strconv.Atoi(string(bytes))
	return result
}
