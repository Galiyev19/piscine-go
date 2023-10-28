package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	str := os.Args[1:]

	if !isValid(str) {
		PrintError()
	} else {
		sudoku := make([][]int, 9)

		for i := 0; i < 9; i++ {
			sudoku[i] = make([]int, 9)
		}

		for row := 0; row < 9; row++ {
			for col := 0; col < 9; col++ {
				sudoku[row][col] = RuneToInt(rune(str[row][col]))
			}
		}

		if SudokuFirstCheck(sudoku) {
			if SudokuSolv(sudoku) {
				for _, a := range sudoku {
					for i, ch := range a {
						z01.PrintRune(rune(ch + '0'))
						if i != 8 {
							z01.PrintRune(' ')
						}
					}
					z01.PrintRune('\n')
				}
			} else {
				PrintError()
			}
		} else {
			PrintError()
		}
	}
}

func SudokuFirstCheck(sudoku [][]int) bool {
	num := 0
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			num = sudoku[row][col]
			if num != 0 {
				if CheckColumn(sudoku, row, col, num) || CheckRow(sudoku, row,
					col, num) || CheckSquare(sudoku, row, col, num) {
					return false
				}
			}
		}
	}
	return true
}

func SudokuSolv(sudoku [][]int) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			// Check positions with 0
			if sudoku[row][col] == 0 {
				// Iterate numbers from 1 to 9
				for num := 1; num <= 9; num++ {
					// Check if duplicates exist
					if !CheckColumn(sudoku, row, col, num) && !CheckRow(sudoku, row, col, num) && !CheckSquare(sudoku, row, col, num) {
						// Assign number po position [row][col]
						sudoku[row][col] = num
						// Recurse SudokuSolv with new data with new assigned number
						if SudokuSolv(sudoku) {
							return true
						}
						// Kill current SudokuSolv
						sudoku[row][col] = 0
					}
				}
				// If cant put number kill Solution
				return false
			}
		}
	}
	return true
}

func CheckRow(sudoku [][]int, row, col, number int) bool {
	for column := 0; column < 9; column++ {
		if col == column {
			continue
		}
		if sudoku[row][column] == number {
			return true
		}
	}
	return false
}

// Check exist number by column in sudoku and return bool
func CheckColumn(sudoku [][]int, row, column, number int) bool {
	for i := 0; i < 9; i++ {
		if i == row {
			continue
		}
		if sudoku[i][column] == number {
			return true
		}
	}
	return false
}

func CheckSquare(sudoku [][]int, row, column, number int) bool {
	x := row
	y := column
	row = row - row%3
	column = column - column%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if x == i+row && y == j+column {
				continue
			}
			if number == sudoku[i+row][j+column] {
				return true
			}
		}
	}
	return false
}

func PrintError() {
	error := "Erorr"

	for _, ch := range error {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

func isValid(s []string) bool {
	if len(s) != 9 {
		return false
	}

	for _, str := range s {
		for _, ch := range str {
			if ch >= '0' && ch <= '9' || ch == '.' {
			} else {
				return false
			}
		}
	}
	return true
}

func RuneToInt(r rune) int {
	counter := 0

	if r == '.' {
		return 0
	}

	for i := '1'; i <= r; i++ {
		counter++
	}

	return counter
}
