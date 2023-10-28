package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	s1 := args[0]
	s2 := args[1]

	res := ""

	for _, ch1 := range s1 {
		for _, ch2 := range s2 {
			if ch1 == ch2 && !ContainChar(ch1, res) {
				res += string(ch1)
			}
		}
	}

	for _, ch := range res {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

func ContainChar(char rune, s string) bool {
	for _, ch := range s {
		if char == ch {
			return true
		}
	}

	return false
}
