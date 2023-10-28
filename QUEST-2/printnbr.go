package piscine

import (
	"github.com/01-edu/z01"
)

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
