package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	fmt.Println(LastWord(args[0]))
}

func LastWord(s string) string {
	var result []string
	startIndex := 0
	trimStr := Trim(s)
	for i, ch := range trimStr {
		if ch == ' ' {
			result = append(result, trimStr[startIndex:i])
			startIndex = i + 1
		}
	}

	result = append(result, trimStr[startIndex:])
	return result[len(result)-1]
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
