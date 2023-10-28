package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	start := Atoi(args[0])
	end := Atoi(args[1])
	// fmt.Println(Atoi("-123"))

	for i := start; i <= end; i++ {
		PrintNbr(i)
		if i != end {
			z01.PrintRune(' ')
		}
	}
}

func PrintNbr(num int) {
	for _, digit := range Itoa(num) {
		z01.PrintRune(digit)
	}
}

func Atoi(s string) int {
	result := 0
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
			result = result*10 + digit
		} else {
			return 0
		}
	}

	return result * sign
}

func Itoa(num int) string {
	if num == 0 {
		return "0"
	}

	res := ""
	isNegative := false

	if num < 0 {
		isNegative = true
		num = -num
	}

	for num > 0 {
		digit := num % 10
		res += string('0' + digit)
		num /= 10
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
