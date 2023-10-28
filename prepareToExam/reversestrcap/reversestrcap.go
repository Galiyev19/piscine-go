package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	split := Split(args[0])
	res := ""
	for _, ch := range split {
		res += UpperWord(ch)
		res += " "
	}

	for i := 0; i < len(res); i++ {
		z01.PrintRune(rune(res[i]))
	}
}

func Split(s string) []string {
	var res []string
	index := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			res = append(res, s[index:i])
			index = i + 1
		}
	}
	res = append(res, s[index:])
	return res
}

func UpperWord(s string) string {
	word := []rune(s)
	res := ""
	for i, ch := range word {
		if i == len(s)-1 {
			if ch >= 'a' && ch <= 'z' {
				res += string(ch - 'a' + 'A')
			} else {
				res += string(ch)
			}
		} else if ch >= 'A' && ch <= 'Z' {
			res += string(ch + 'a' - 'A')
		} else {
			res += string(ch)
		}
	}
	return res
}
