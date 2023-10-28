package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	trimStr := Trim(args[0])
	var str []string
	index := 0
	for i, ch := range trimStr {
		if ch == ' ' {
			str = append(str, trimStr[index:i])
		}
	}
	Print(str)
}

func Print(s []string) {
	for _, ch := range s {
		for _, b := range ch {
			z01.PrintRune(b)
		}
	}
	z01.PrintRune('\n')
}

func Trim(s string) string {
	start, end := 0, len(s)

	for start < end && s[start] == ' ' {
		start++
	}

	for start < end && s[end-1] == ' ' {
		end--
	}

	return s[start:end]
}
