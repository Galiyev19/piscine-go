package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	num := Atoi(args[0])
	res := 0

	for i := 2; i <= num; i++ {
		if IsPrime(i) {
			res += i
		}
	}
	PrintDigits(res)
}

func PrintDigits(n int) {
	if n >= 10 {
		PrintDigits(n / 10)
	}

	digit := n % 10
	digitRune := digit + '0'
	z01.PrintRune(rune(digitRune))
}

func IsPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func Atoi(s string) int {
	result := 0
	sign := 1
	count := 0

	for _, ch := range s {
		if ch == '-' || ch == '+' {
			count++
		}
		if count > 1 {
			return 0
		}

		if ch == '-' {
			sign = -1
		} else if ch >= '0' && ch <= '9' {
			digit := int(ch - '0')
			result = result*10 + digit
		} else {
			return 0
		}
	}
	return result * sign
}
