package piscine

import (
	"github.com/01-edu/z01"
)

const N = 8

var (
	board  [N][N]bool
	result [92]string // There are 92 solutions for the 8 queens puzzle
	count  int
)

func EightQueens() {
	count = 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			board[i][j] = false
		}
	}
	solve(0)
	// Sort the result array manually
	for i := 0; i < count-1; i++ {
		for j := i + 1; j < count; j++ {
			if result[i] > result[j] {
				result[i], result[j] = result[j], result[i]
			}
		}
	}
	for i := 0; i < count; i++ {
		printSolution(result[i])
	}
}

func isSafe(row, col int) bool {
	for i := 0; i < col; i++ {
		if board[row][i] {
			return false
		}
	}
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] {
			return false
		}
	}
	for i, j := row, col; i < N && j >= 0; i, j = i+1, j-1 {
		if board[i][j] {
			return false
		}
	}
	return true
}

func solve(col int) {
	if col == N {
		var solution string
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if board[i][j] {
					solution += string(49 + j)
				}
			}
		}
		result[count] = solution
		count++
		return
	}
	for i := 0; i < N; i++ {
		if isSafe(i, col) {
			board[i][col] = true
			solve(col + 1)
			board[i][col] = false
		}
	}
}

func printSolution(solution string) {
	for _, runeValue := range solution {
		z01.PrintRune(runeValue)
	}
	z01.PrintRune(10)
}
