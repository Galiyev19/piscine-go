package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	splitArr := Split(args[0])
	fmt.Println(splitArr)
	res := ""
	for i := len(splitArr) - 1; i >= 0; i-- {
		res += splitArr[i]
		if i != 0 {
			res += " "
		}
	}

	for _, ch := range res {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

func Split(s string) []string {
	var result []string
	startIndex := 0

	for i, ch := range s {
		if ch == ' ' {
			result = append(result, s[startIndex:i])
			startIndex = i + 1
		}
	}

	result = append(result, s[startIndex:])
	return result
}
