package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Board [9][9]int
type Position struct {
	x int
	y int
}

func printBoard(board Board) {
	fmt.Println("+-------+-------+-------+")
	for row := 0; row < 9; row++ {
		fmt.Print("| ")
		for col := 0; col < 9; col++ {
			if col == 3 || col == 6 {
				fmt.Print("| ")
			}
			fmt.Printf("%d ", board[row][col])
			if col == 8 {
				fmt.Print("|")
			}
		}
		if row == 2 || row == 5 || row == 8 {
			fmt.Println("\n+-------+-------+-------+")
		} else {
			fmt.Println()
		}
	}
}

func parseInput(input string) Board {
	board := Board{}
	scanner := bufio.NewScanner(strings.NewReader(input))

	for row := 0; row < 9; row++ {
		scanner.Scan()
		for col, c := range scanner.Text() {
			i1, _ := strconv.Atoi(string(c))
			board[row][col] = i1
		}
	}
	return board
}
