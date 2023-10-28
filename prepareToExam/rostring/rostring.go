package main

import (
	"fmt"
	"os"
)

func main() {
	str := os.Args[1:]
	temp := Spilt(str[0])
	fWord := temp[0]
	res := ""
	for i := 1; i < len(temp); i++ {
		if temp[i] != " " {
			res += string(temp[i])
		}
	}

	fmt.Println(res + " " + fWord)
}

func Spilt(s string) []string {
	var result []string
	startIndex := 0
	for i, ch := range s {
		if ch == ' ' {
			result = append(result, s[startIndex:i+1])
			startIndex = i + 1
		}
	}
	result = append(result, s[startIndex:])

	return result
}
