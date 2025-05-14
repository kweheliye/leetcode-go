package medium

import (
	"container/list"
	"math"
)

func updateMatrix(mat [][]int) [][]int {
	rows, cols := len(mat), len(mat[0])
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	queue := list.New()
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if mat[i][j] == 0 {
				queue.PushBack([]int{i, j})
			} else {
				mat[i][j] = math.MaxInt32
			}
		}
	}

	for queue.Len() > 0 {
		cell := queue.Remove(queue.Front()).([]int)
		row, col := cell[0], cell[1]

		for _, direction := range directions {
			newRow, newCol := row+direction[0], col+direction[1]
			if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols {
				continue
			}
			if mat[newRow][newCol] > mat[cell[0]][cell[1]]+1 {
				mat[newRow][newCol] = mat[cell[0]][cell[1]] + 1
				queue.PushBack([]int{newRow, newCol})
			}
		}
	}

	return mat
}
