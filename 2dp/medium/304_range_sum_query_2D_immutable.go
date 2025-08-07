package medium

type NumMatrix struct {
	data [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	return NumMatrix{data: matrix}
}

func (nMatrix *NumMatrix) SumRegionBruteForce(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for r := row1; r <= row2; r++ {
		for c := col1; c <= col2; c++ {
			sum += nMatrix.data[r][c]
		}
	}
	return sum
}
