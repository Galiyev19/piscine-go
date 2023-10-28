package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	num := atoi(args[0])

	for i := 1; i <= num; i++ {
		res := num * i
		z01.PrintRune(rune(i + '0'))
		z01.PrintRune(' ')
		z01.PrintRune('x')
		z01.PrintRune(' ')
		z01.PrintRune(rune(num + '0'))
		z01.PrintRune(' ')
		z01.PrintRune('=')
		z01.PrintRune(' ')
		PrintDigis(res)
		z01.PrintRune('\n')
	}
}

func atoi(str string) int {
	res := 0
	count := 0
	sign := 1

	for _, char := range str {
		if char == '-' || char == '+' {
			count++

			if count > 1 {
				return 0
			}

			if char == '-' {
				sign = -1
			}
		} else if char >= '0' && char <= '9' {
			digit := int(char - '0')
			res = res*10 + digit
		} else {
			return 0
		}
	}

	return res * sign
}

func PrintDigis(n int) {
	if n >= 10 {
		PrintDigis(n / 10)
	}

	digit := n % 10
	ditgitRune := digit + '0'
	z01.PrintRune(rune(ditgitRune))
}
