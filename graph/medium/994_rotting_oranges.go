package medium

func orangesRotting(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	time := 0
	fresh := 0
	queue := [][]int{}

	// Count fresh oranges and collect rotten ones
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				fresh++
			}
			if grid[r][c] == 2 {
				queue = append(queue, []int{r, c})
			}
		}
	}

	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	// BFS
	for len(queue) > 0 && fresh > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cell := queue[0]
			queue = queue[1:] // dequeue
			row, col := cell[0], cell[1]

			for _, dir := range directions {
				r, c := row+dir[0], col+dir[1]
				if r >= 0 && r < rows && c >= 0 && c < cols && grid[r][c] == 1 {
					grid[r][c] = 2
					fresh--
					queue = append(queue, []int{r, c})
				}
			}
		}
		time++
	}

	if fresh > 0 {
		return -1
	}
	return time
}
