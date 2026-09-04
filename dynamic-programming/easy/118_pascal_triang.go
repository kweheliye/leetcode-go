package easy

func generate(numRows int) [][]int {
	triangle := make([][]int, numRows)

	triangle[0] = append(triangle[0], 1)

	for rowNum := 1; rowNum < numRows; rowNum++ {
		row := make([]int, 0)
		prevRow := triangle[rowNum-1]
		row = append(row, 1)
		for i := 1; i < rowNum; i++ {
			row = append(row, prevRow[i-1]+prevRow[i])
		}
		row = append(row, 1)
		triangle[rowNum] = row

	}

	return triangle

}
