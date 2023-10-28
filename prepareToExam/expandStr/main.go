package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) > 1 || len(args) <= 0 {
		return
	}

	strArr := Spilt(args[0])
	str := ""
	space := " "

	for i := 0; i < len(strArr); i++ {
		str += strArr[i]
		if i < len(strArr)-1 {
			str += space
		}

	}

	for _, ch := range str {
		z01.PrintRune(ch)
	}
}

func Spilt(s string) []string {
	var str []string
	startIndex := 0
	trimStr := Trim(s)
	for i := 0; i < len(trimStr); i++ {
		if trimStr[i] == ' ' {
			str = append(str, trimStr[startIndex:i])
			startIndex = i + 1
		}
	}
	str = append(str, trimStr[startIndex:])
	return str
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
