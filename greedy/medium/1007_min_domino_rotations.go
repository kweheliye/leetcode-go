package medium

func minDominoRotations(tops []int, bottoms []int) int {
	n := len(tops)

	rotations := check(tops[0], tops, bottoms, n)
	if rotations != -1 {
		return rotations
	} else {
		return check(bottoms[0], tops, bottoms, n)
	}

}

func check(x int, tops []int, bottoms []int, n int) int {
	rotationsA, rotationsB := 0, 0

	for i := 0; i < n; i++ {
		if tops[i] != x && bottoms[i] != x {
			return -1
		} else if tops[i] != x {
			rotationsA++
		} else if bottoms[i] != x {
			rotationsB++
		}
	}

	return min(rotationsA, rotationsB)
}
