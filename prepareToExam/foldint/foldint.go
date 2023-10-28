package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	table := []int{1, 2, 3}
	ac := 93
	// FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	// FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac)
	// FoldInt(Mul, table, ac)
	// FoldInt(Sub, table, ac)
}

func FoldInt(f func(int, int) int, a []int, n int) {
	res := a[0]
	for i := 1; i < len(a); i++ {
		res = f(res, a[i])
	}
	res += n

	PrintDigits(res)
	z01.PrintRune('\n')
}

func Add(a, b int) int {
	return a + b
}

func Mul(a, b int) int {
	return a * b
}

func PrintDigits(num int) {
	if num >= 10 {
		PrintDigits(num / 10)
	}

	digit := num % 10
	runeDigit := '0' + rune(digit)
	z01.PrintRune(runeDigit)
}
