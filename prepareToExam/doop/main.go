package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	num1 := atoi(args[0])
	num2 := atoi(args[2])
	operator := args[1]
	sign := 0

	errorDivide := "No division by 0"
	errorModule := "No modulo by 0"

	if operator == "/" && num2 == 0 {
		for _, ch := range errorDivide {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
		return
	} else if operator == "%" && num2 == 0 {
		for _, ch := range errorModule {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
		return
	}

	if args[1] == "+" {
		sign = 0
	} else if args[1] == "-" {
		sign = 1
	} else if args[1] == "*" {
		sign = 2
	} else if args[1] == "/" {
		sign = 3
	} else if args[1] == "%" {
		sign = 4
	} else {
		return
	}

	funcArr := []func(int, int) int{plus, minus, multi, divide, module}
	result := apply(funcArr[sign], num1, num2)

	PrintDigis(result)
}

func plus(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

func multi(a, b int) int {
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

func atoi(s string) int {
	res := 0
	sign := 1
	count := 0

	for _, char := range s {
		if char == '-' || char == '+' {
			count++

			if count > 1 {
				return 0
			}

			if count == '-' {
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

func PrintDigis(num int) {
	if num >= 10 {
		PrintDigis(num / 10)
	}

	digit := num % 10
	digitRune := '0' + rune(digit)
	z01.PrintRune(digitRune)
	z01.PrintRune('\n')
}
