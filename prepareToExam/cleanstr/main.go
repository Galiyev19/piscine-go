package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	temp := Split2(args[0])
	temp1 := ""
	fmt.Println(temp)
	for _, ch := range temp {
		if ch != " " {
			temp1 += string(ch)
		}
		// temp1 += " "
	}

	Print(temp1)
}

func Split(s string) []string {
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

func Split2(s string) []string {
	var str []string
	startIndex := 0
	trimStr := Trim(s)
	for i := 0; i < len(trimStr); i++ {
		if trimStr[i] == ' ' {
			str = append(str, trimStr[startIndex:i+1])
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

func Print(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
