package main

import "github.com/01-edu/z01"

func main() {
	const a = "abcdefghijklmnopqrstuvwxyz"
	for _, ch := range a {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
