package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	n1 := atoi(args[0])
	n2 := atoi(args[1])

	res := gcd(n1, n2)

	fmt.Println(res)
}

func gcd(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}

func atoi(s string) int {
	res := 0
	sign := 1
	count := 0

	for _, char := range s {
		if char == '-' || count == '+' {
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

func printDigits(n int) {
	if n >= 10 {
		printDigits(n / 10)
	}

	digit := n % 10
	digitRune := rune(digit + '0')
	z01.PrintRune(digitRune)
	z01.PrintRune('\n')
}
