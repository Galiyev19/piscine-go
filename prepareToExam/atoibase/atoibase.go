package main

import (
	"fmt"
)

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}

func AtoiBase(s string, base string) int {
	length := len(base)
	res := 0

	if !IsValidBase(base) {
		return res
	}

	for _, str := range s {
		digit := string(str)
		index := 0

		for i, char := range base {
			if string(char) == digit {
				index = i
				break
			}
		}
		res = res*length + index
	}

	return res
}

func IsValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	uniqueChars := make(map[rune]bool)

	for _, ch := range base {
		if uniqueChars[ch] || ch == '-' || ch == '+' {
			return false
		}
		uniqueChars[ch] = true
	}

	return true
}
