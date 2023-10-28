package main

import "github.com/01-edu/z01"

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}

func ReduceInt(a []int, f func(int, int) int) {
	result := a[0]

	for i := 1; i < len(a); i++ {
		result = f(result, a[i])
	}

	PrintDigits(result)
	z01.PrintRune('\n')
}

func PrintDigits(num int) {
	if num >= 10 {
		PrintDigits(num / 10)
	}

	digit := num % 10
	runeDigit := '0' + rune(digit)
	z01.PrintRune(runeDigit)
}
