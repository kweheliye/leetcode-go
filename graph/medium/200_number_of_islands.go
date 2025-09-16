package medium

var rows int
var cols int
var directions [][]int

func numIslands(grid [][]byte) int {
	rows = len(grid)
	cols = len(grid[0])
	directions = [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	islands := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {

			if grid[r][c] == '1' {
				islands++
				dfs(r, c, grid)
			}
		}
	}

	return islands

}

func dfs(r int, c int, grid [][]byte) {
	if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == '0' {
		return
	}

	grid[r][c] = '0'
	for _, d := range directions {
		dfs(r+d[0], c+d[1], grid)
	}

}
