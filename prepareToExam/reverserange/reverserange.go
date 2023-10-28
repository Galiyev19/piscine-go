package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	start := Atoi(args[0])
	end := Atoi(args[1])

	if start > end {
		start, end = end, start
	}
	for i := start; i <= end; i++ {
		PrintNbr(i)
		if i != start {
			z01.PrintRune(' ')
		}
	}
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	for _, ch := range Itoa(n) {
		z01.PrintRune(ch)
	}
}

func Atoi(s string) int {
	res := 0
	sign := 1
	count := 0

	for _, ch := range s {
		if ch == '-' || ch == '+' {
			count++

			if count > 1 {
				return 0
			}

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

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	res := ""
	isNegative := false

	if n < 0 {
		isNegative = true
		n = -n
	}

	for n > 0 {
		digit := n % 10
		res += string(digit + '0')
		n /= 10
	}

	if isNegative {
		res += "-"
	}

	revStr := ""

	for i := len(res) - 1; i >= 0; i-- {
		revStr += string(res[i])
	}

	return revStr
}
