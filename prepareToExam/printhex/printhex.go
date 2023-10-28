package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		Print("ERROR")
		return
	}

	if len(args) == 0 {
		return
	}

	num := atoi(args[0])

	if num < 0 {
		return
	}
	res := decimalTo16Base(num)

	Print(res)
}

func Print(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

func decimalTo16Base(decimal int) string {
	hexDigits := "0123456789abcdef"
	length := len(hexDigits)

	var result string
	for decimal > 0 {
		remainder := decimal % length
		result = string(hexDigits[remainder]) + result
		decimal /= length
	}

	return result
}

func atoi(s string) int {
	res := 0
	sign := 1
	count := 0

	for _, ch := range s {
		if ch == '-' || ch == '+' {
			count++
		}

		if count > 1 {
			return 0
		}

		if ch == '-' || count == '+' {
			if ch == '-' {
				sign = -1
			}
		} else if ch >= '0' && ch <= '9' {
			digit := int(ch - '0')
			res = res*10 + digit
		} else {
			return 0
		}
	}

	return res * sign
}
