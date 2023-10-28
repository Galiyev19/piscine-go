package main

import "github.com/01-edu/z01"

func main() {
	// digits := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	digits := "0123456789"
	for _, value := range digits {
		z01.PrintRune(value)
	}
	z01.PrintRune('\n')
}
