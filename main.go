package main

func main() {
	board := "061800200\n590070061\n070004095\n900000020\n105030709\n030000004\n840700050\n710080042\n009001370"
	sudoku := parseInput(board)
	printBoard(sudoku)

	solve(&sudoku, Position{x: 0, y: 0})
	printBoard(sudoku)
}
