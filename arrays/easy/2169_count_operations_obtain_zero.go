package easy

func countOperations(num1 int, num2 int) int {
	res := 0

	for num1 != 0 && num2 != 0 {
		if num1 >= num2 {
			res += num1 / num2
			num1 = num1 % num2
		} else {
			res += num2 / num1
			num2 = num2 % num1
		}

	}
	return res
}
