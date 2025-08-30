package medium

func isValidSudoku(board [][]byte) bool {
	N := 9

	rows := make([]map[byte]bool, N)
	cols := make([]map[byte]bool, N)
	boxes := make([]map[byte]bool, N)

	for i := 0; i < N; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			cell := board[r][c]
			if cell == '.' {
				continue
			}
			if rows[r][cell] || cols[c][cell] || boxes[r/3*3+c/3][cell] {
				return false
			}

			rows[r][cell] = true
			cols[c][cell] = true
			boxes[r/3*3+c/3][cell] = true
		}
	}

	return true
}
