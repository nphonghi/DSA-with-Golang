package main

import "fmt"

func isValid(board [][]int, row, col, num int) bool {
	for x := 0; x < 9; x++ {
		if board[row][x] == num || board[x][col] == num || board[3*(row/3)+x/3][3*(col/3)+x%3] == num {
			return false
		}
	}
	return true
}

func solveSudoku(board [][]int) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == 0 {
				for num := 1; num <= 9; num++ {
					if isValid(board, row, col, num) {
						board[row][col] = num
						if solveSudoku(board) {
							return true
						}
						board[row][col] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

func printBoard(board [][]int) {
	for _, row := range board {
		for _, num := range row {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}

func main() {
	board := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	if solveSudoku(board) {
		printBoard(board)
	} else {
		fmt.Println("No solution exists")
	}
}
