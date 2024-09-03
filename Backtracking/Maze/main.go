package main

import "fmt"

func isSafe(maze [][]int, x, y int) bool {
	return x >= 0 && x < len(maze) && y >= 0 && y < len(maze[0]) && maze[x][y] == 1
}

func solveMazeUtil(maze [][]int, x, y int, sol [][]int) bool {
	if x == len(maze)-1 && y == len(maze[0])-1 && maze[x][y] == 1 {
		sol[x][y] = 1
		return true
	}

	if isSafe(maze, x, y) {
		sol[x][y] = 1

		if solveMazeUtil(maze, x+1, y, sol) {
			return true
		}
		if solveMazeUtil(maze, x, y+1, sol) {
			return true
		}

		sol[x][y] = 0
		return false
	}
	return false
}

func solveMaze(maze [][]int) {
	sol := make([][]int, len(maze))
	for i := range sol {
		sol[i] = make([]int, len(maze[0]))
	}

	if solveMazeUtil(maze, 0, 0, sol) {
		printSolution(sol)
	} else {
		fmt.Println("No solution exists")
	}
}

func printSolution(sol [][]int) {
	for _, row := range sol {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}

func main() {
	maze := [][]int{
		{1, 0, 0, 0},
		{1, 1, 0, 1},
		{0, 1, 0, 0},
		{1, 1, 1, 1},
	}

	solveMaze(maze)
}
