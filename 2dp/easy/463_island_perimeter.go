package easy

func islandPerimeter(grid [][]int) int {

	rows := len(grid)
	cols := len(grid[0])
	result, up, down, left, right := 0, 0, 0, 0, 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				if r == 0 {
					up = 0
				} else {
					up = grid[r-1][c]
				}

				if c == 0 {
					left = 0
				} else {
					left = grid[r][c-1]
				}

				if r == rows-1 {
					down = 0
				} else {
					down = grid[r+1][c]
				}

				if c == cols-1 {
					right = 0
				} else {
					right = grid[r][c+1]
				}
				result += 4 - (up + down + left + right)
			}
		}
	}

	return result

}
