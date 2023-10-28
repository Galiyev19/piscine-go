package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	num, _ := strconv.Atoi(arg[0])
	if num <= 0 || num >= 4000 {
		Print("ERROR: cannot convert to roman digit")
		return
	}
	fmt.Println(RomanNumbers(num))
}

func RomanNumbers(num int) string {
	var result string

	romanSymbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	romanValues := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	for i := 0; i < len(romanSymbols); i++ {
		for num >= romanValues[i] {
			result += romanSymbols[i]
			num -= romanValues[i]
		}
	}
	return result
}

func Print(err string) {
	for _, ch := range err {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
