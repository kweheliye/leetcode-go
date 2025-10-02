package easy

func numWaterBottles(numBottles int, numExchange int) int {

	consumedBottles := 0

	for numBottles >= numExchange {
		consumedBottles += numExchange
		numBottles -= numExchange
		numBottles++
	}

	return consumedBottles + numBottles

}
