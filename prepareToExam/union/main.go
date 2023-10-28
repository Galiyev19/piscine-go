package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	res := ""
	for _, ch1 := range args[0] {
		if !ContainChar(ch1, res) {
			res += string(ch1)
		}
	}

	for _, ch2 := range args[1] {
		if !ContainChar(ch2, res) {
			res += string(ch2)
		}
	}

	fmt.Println(res)
}

func ContainChar(char rune, str string) bool {
	for _, ch := range str {
		if char == ch {
			return true
		}
	}

	return false
}
