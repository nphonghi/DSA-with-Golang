package main

import "fmt"

func isSafe(board [][]int, row, col, n int) bool {
	for i := 0; i < col; i++ {
		if board[row][i] == 1 {
			return false
		}
	}

	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}

	for i, j := row, col; i < n && j >= 0; i, j = i+1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}

	return true
}

func solveNQueensUtil(board [][]int, col, n int) bool {
	if col >= n {
		return true
	}

	for i := 0; i < n; i++ {
		if isSafe(board, i, col, n) {
			board[i][col] = 1
			if solveNQueensUtil(board, col+1, n) {
				return true
			}
			board[i][col] = 0
		}
	}
	return false
}

func solveNQueens(n int) {
	board := make([][]int, n)
	for i := range board {
		board[i] = make([]int, n)
	}

	if solveNQueensUtil(board, 0, n) {
		printBoard(board)
	} else {
		fmt.Println("No solution exists")
	}
}

func printBoard(board [][]int) {
	for _, row := range board {
		for _, val := range row {
			if val == 1 {
				fmt.Print("Q ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}

func main() {
	n := 8
	solveNQueens(n)
}
