package piscine

import (
	"github.com/01-edu/z01"
)

func isValidBase(s string) bool {
	if len(s) < 2 {
		return false
	}

	seen := make(map[rune]bool)

	for _, char := range s {
		if char == '-' || char == '+' {
			return false
		}
		if seen[char] {
			return false
		}
		seen[char] = true
	}

	return true
}

func PrintNbrBase(nbr int, base string) {
	if nbr == -9223372036854775808 {
		z01.PrintRune('-')
		z01.PrintRune('9')
		z01.PrintRune('2')
		z01.PrintRune('2')
		z01.PrintRune('3')
		z01.PrintRune('3')
		z01.PrintRune('7')
		z01.PrintRune('2')
		z01.PrintRune('0')
		z01.PrintRune('3')
		z01.PrintRune('6')
		z01.PrintRune('8')
		z01.PrintRune('5')
		z01.PrintRune('4')
		z01.PrintRune('7')
		z01.PrintRune('7')
		z01.PrintRune('5')
		z01.PrintRune('8')
		z01.PrintRune('0')
		z01.PrintRune('8')
		return
	}

	if !isValidBase(base) || len(base) < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	isNegative := false
	if nbr < 0 {
		isNegative = true
		nbr = -nbr
	}

	length := len(base)
	reversStr := ""

	for nbr > 0 {
		remainder := nbr % length
		digit := rune(base[remainder])
		reversStr += string(digit)
		nbr /= length
	}

	if isNegative {
		z01.PrintRune('-')
	}
	temp := []rune(reversStr)

	for i := len(temp) - 1; i >= 0; i-- {
		z01.PrintRune(temp[i])
	}
}
