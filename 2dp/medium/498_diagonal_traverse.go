package medium

import "slices"

func findDiagonalOrder(matrix [][]int) []int {
	totalRows, totalCols := len(matrix), len(matrix[0])

	result := make([]int, totalRows*totalCols)
	var intermediate []int
	k := 0

	for diagonal := 0; diagonal < totalRows+totalCols; diagonal++ {
		intermediate = nil
		var startRow, startCol int

		if diagonal < totalCols {
			startRow, startCol = 0, diagonal
		} else {
			startRow, startCol = diagonal-totalCols+1, totalCols-1
		}

		for startRow < totalRows && startCol > -1 {
			intermediate = append(intermediate, matrix[startRow][startCol])
			startRow++
			startCol--
		}

		if diagonal%2 == 0 {
			slices.Reverse(intermediate)
		}

		for _, val := range intermediate {
			result[k] = val
			k++
		}

	}
	return result
}
