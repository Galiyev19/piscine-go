package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	res := ""

	for _, ch := range args[0] {
		if ch >= 'A' && ch <= 'Z' {
			res += string(rune(ch) + 32)
		} else if ch >= 'a' && ch <= 'z' {
			res += string(rune(ch) - 32)
		} else {
			res += string(ch)
		}
	}

	for _, ch := range res {
		z01.PrintRune(ch)
	}
}
