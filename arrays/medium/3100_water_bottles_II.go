package medium

func maxBottlesDrunk(numBottles int, numExchange int) int {
	ans := numBottles
	empty := numBottles

	for empty >= numExchange {
		ans++
		empty -= numExchange
		empty++
		numExchange++
	}
	return ans
}
