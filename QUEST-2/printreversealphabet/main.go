package main

import "github.com/01-edu/z01"

func main() {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	runes := []rune(alphabet)

	for i := len(runes) - 1; i >= 0; i-- {
		z01.PrintRune(runes[i])
	}
	z01.PrintRune('\n')
}
