package main

import (
	"os"

	"github.com/01-edu/z01"
)

func BasicAtoi(s string) int {
	result := 0

	for i, value := range s {
		digit := int(value - '0')

		result += digit

		if i != len(s)-1 {
			result *= 10
		}
	}

	return result
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	for i := 0; i < len(args); i++ {
		// fmt.Println(arg)
		num := BasicAtoi(args[i])

		if args[0] == "--upper" {
			if num > 26 {
				z01.PrintRune(' ')
			} else {
				z01.PrintRune(rune(64 + num))
			}
		} else {
			if num > 26 {
				z01.PrintRune(' ')
			} else {
				z01.PrintRune(rune(96 + num))
			}
		}
	}
	z01.PrintRune('\n')
}
