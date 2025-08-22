package easy

import "container/list"

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

func islandPerimeterBFS(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				// Use container/list as queue
				queue := list.New()
				queue.PushBack([2]int{i, j})
				visited[i][j] = true
				perimeter := 0

				for queue.Len() > 0 {
					// Dequeue (poll)
					cell := queue.Remove(queue.Front()).([2]int)
					x, y := cell[0], cell[1]

					for _, dir := range directions {
						nx, ny := x+dir[0], y+dir[1]
						if nx < 0 || ny < 0 || nx >= rows || ny >= cols || grid[nx][ny] == 0 {
							perimeter++
						} else if !visited[nx][ny] {
							visited[nx][ny] = true
							queue.PushBack([2]int{nx, ny})
						}
					}
				}
				return perimeter
			}
		}
	}
	return 0
}
