package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	str := args[0]
	error := "No vowels"

	if fIsWovel(args[0][0]) {
		args[0] += "ay"
		for _, ch := range str {
			z01.PrintRune(ch)
		}
	} else if IsVowel(args[0]) {
		start := 0
		end := 0
		var res []string

		for i, s := range str {
			if IsVowel(string(s)) {
				res = append(res, str[i:])
				end = i
			}
		}
		res = append(res, str[start:end])
		temp := ""
		for _, ch := range res {
			temp += string(ch)
		}

		temp += "ay"
		fmt.Println(temp)

	} else {
		for _, ch := range error {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}

func IsVowel(s string) bool {
	for _, s := range s {
		if s == 'A' || s == 'E' || s == 'I' || s == 'O' || s == 'U' || s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' {
			return true
		}
	}
	return false
}

func fIsWovel(s byte) bool {
	if s == 'A' || s == 'E' || s == 'I' || s == 'O' || s == 'U' || s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' {
		return true
	}

	return false
}
