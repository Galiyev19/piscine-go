package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argsWithoutProg := os.Args[1:]

	for _, arg := range argsWithoutProg {
		for _, char := range arg {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
