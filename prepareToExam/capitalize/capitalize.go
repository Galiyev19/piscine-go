package main

import (
	"fmt"
	piscine "piscine/QUEST-5"
)

func main() {
	fmt.Println(piscine.Capitalize("Hello! How are you? How+are+things+4you?"))
}

func Capitalize(s string) string {
	upperToNext := false
	result := make([]rune, len(s))
	for i, ch := range s {
		if ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' {
			if upperToNext {
				result[i] = ch - 'a' - 'A'
				upperToNext = false
			} else {
				result[i] = ch + 'a' - 'A'
			}
		} else {
			result[i] = ch
			upperToNext = true
		}
	}

	return string(result)
}
