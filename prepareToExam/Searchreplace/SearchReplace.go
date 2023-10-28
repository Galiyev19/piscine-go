package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	str := args[0]
	fletter := args[1]
	sletter := args[2]
	result := ""
	for _, ch := range str {
		if string(ch) == fletter {
			result += sletter
		} else {
			result += string(ch)
		}
	}

	for _, ch := range result {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
