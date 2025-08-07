package medium

type NumMatrix struct {
	data [][]int
}

func ConstructorBruteForce(matrix [][]int) NumMatrix {
	return NumMatrix{data: matrix}
}

func ConstructorOneDimensionalPrefixSum(matrix [][]int) NumMatrix {
	m := len(matrix)
	n := len(matrix[0])
	data := make([][]int, m)
	for i := 0; i < m; i++ {
		data[i] = make([]int, n+1)
		for j := 0; j < n; j++ {
			data[i][j+1] = data[i][j] + matrix[i][j]
		}
	}
	return NumMatrix{data: data}
}

// SumRegionBruteForce Brute force Version
func (nMatrix *NumMatrix) SumRegionBruteForce(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for r := row1; r <= row2; r++ {
		for c := col1; c <= col2; c++ {
			sum += nMatrix.data[r][c]
		}
	}
	return sum
}

// SumRegionOneDimensionalPrefixSum SumRegionBruteForce One Dimensional Prefix Sum Version
func (nMatrix *NumMatrix) SumRegionOneDimensionalPrefixSum(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for r := row1; r <= row2; r++ {
		if col1 > 0 {
			sum += nMatrix.data[r][col2] - nMatrix.data[r][col1-1]
		} else {
			sum += nMatrix.data[r][col2]
		}

	}
	return sum
}
