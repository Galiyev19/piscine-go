package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	operator := getOperator(args)

	sign := 0
	if operator == "+" {
		sign = 0
	} else if operator == "-" {
		sign = 1
	} else if operator == "*" {
		sign = 2
	} else if operator == "/" {
		sign = 3
	} else if operator == "%" {
		sign = 4
	} else {
		return
	}

	for _, ch := range args[0] {
		if ch > '1' && ch < '9' {
			return
		}
	}

	for _, ch := range args[2] {
		if ch > '1' && ch < '9' {
			return
		}
	}

	leftNbr := Atoi(args[0])
	rightNbr := Atoi(args[2])

	if leftNbr == 9223372036854775807 || rightNbr == 9223372036854775807 {
		return
	}

	errorDivider := "No division by 0"
	errorModule := "No modulo by 0"

	if rightNbr == 0 || leftNbr == 0 {
		if operator == "/" {
			for _, ch := range errorDivider {
				z01.PrintRune(rune(ch))
			}
			z01.PrintRune('\n')
			return
		}

		if operator == "%" {
			for _, ch := range errorModule {
				z01.PrintRune(rune(ch))
			}
			z01.PrintRune('\n')
			return
		}
	}

	// res := leftNbr + rightNbr
	arrayFunction := []func(int, int) int{plus, minus, multiply, divide, module}
	res := apply(arrayFunction[sign], leftNbr, rightNbr)

	if res >= 9223372036854775806 || res <= -9223372036854775806 {
		return
	} else {
		PrintNbr(res)
		z01.PrintRune('\n')
	}
}

func plus(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	return a / b
}

func module(a, b int) int {
	return a % b
}

func apply(f func(a, b int) int, a, b int) int {
	res := f(a, b)
	return res
}

func getOperator(args []string) string {
	operator := ""
	for _, oprt := range args {
		if oprt == "+" || oprt == "-" || oprt == "*" || oprt == "/" || oprt == "%" {
			operator += oprt
		}
	}
	return operator
}

func printDigits(num int) {
	if num >= 10 {
		printDigits(num / 10)
	}

	digit := num % 10
	runeDigit := '0' + rune(digit)
	z01.PrintRune(runeDigit)
}

func PrintNbr(num int) {
	if num == 0 {
		z01.PrintRune('0')
		return
	}

	if num < 0 {
		z01.PrintRune('-')
		// num = -(num + 1)
		num = -num

	}

	if num == -9223372036854775808 {
		num = 922337203685477580
		printDigits(num)
		z01.PrintRune('8')
		return
	}

	printDigits(num)
}

func Atoi(s string) int {
	var result int
	sign := 1
	count := 0
	digitFound := false

	for _, char := range s {

		if char == '+' || char == '-' {
			count++
		}

		if count > 1 {
			return 0
		}

		if char == '-' || char == '+' {
			if digitFound {
				return 0
			}
			if char == '-' {
				sign = -1
			}
		} else if char >= '0' && char <= '9' {
			digit := int(char - '0')
			result = result*10 + digit
			digitFound = true
		} else {
			return 0
		}
	}

	return sign * result
}

func PrintErr() {
	error := "No division by 0"

	for _, ch := range error {
		z01.PrintRune(rune(ch))
	}
	z01.PrintRune('\n')
}
