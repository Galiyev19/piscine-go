package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	runes1 := []rune(args[0])
	runes2 := []rune(args[1])
	count := 0
	res := ""
	for i := 0; i < len(runes1); i++ {
		for j := count; j < len(runes2); j++ {
			if runes1[i] == runes2[j] {
				res += string(runes1[i])
				j = len(runes2) - 1
			}
			count++
		}
	}

	if res == args[0] {
		z01.PrintRune('1')
		z01.PrintRune('\n')
	} else {
		z01.PrintRune('0')
		z01.PrintRune('\n')
	}
}
