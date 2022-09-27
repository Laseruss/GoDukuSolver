package main

func isValidBoard(board *Board, row, col, value int) bool {
	return validRow(board, row, value) &&
		validCol(board, col, value) &&
		validSubGroup(board, row, col, value)
}

func validRow(board *Board, row, value int) bool {
	sum := 0
	for _, v := range board[row] {
		if v == value {
			sum++
		}
	}
	if sum > 1 {
		return false
	}
	return true
}

func validCol(board *Board, col, value int) bool {
	sum := 0
	for row := 0; row < 9; row++ {
		if board[row][col] == value {
			sum++
		}
	}
	if sum > 1 {
		return false
	}
	return true
}

func validSubGroup(board *Board, row, col, value int) bool {
	sum := 0
	subRow := row / 3
	subCol := col / 3

	xStart := max(subCol*3, 0)
	yStart := max(subRow*3, 0)

	for row := yStart; row < yStart+3; row++ {
		for col := xStart; col < xStart+3; col++ {
			if board[row][col] == value {
				sum++
			}
		}
	}
	if sum > 1 {
		return false
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
