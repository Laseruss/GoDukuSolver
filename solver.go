package main

func solve(board *Board, pos Position) bool {
	if pos.x == 8 && pos.y == 8 && isValidBoard(board, pos.y, pos.x, board[pos.y][pos.x]) {
		return true
	}

	for y, row := range board {
		for x, cell := range row {
			if cell == 0 {
				for k := 9; k > 0; k-- {
					board[y][x] = k
					if isValidBoard(board, y, x, board[y][x]) {
						if solve(board, Position{x: x, y: y}) {
							return true
						}
						board[y][x] = 0
					} else {
						board[y][x] = 0
					}
				}
				return false
			}
		}
	}
	return false
}
