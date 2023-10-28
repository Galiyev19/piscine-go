package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	programmName := []rune(os.Args[0])

	for i := 2; i < len(programmName); i++ {
		z01.PrintRune(programmName[i])
	}
	z01.PrintRune('\n')
}
