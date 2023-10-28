package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	// fmt.Println(args)
	if len(args) <= 0 {
		return
	}

	res := ""

	for _, ch := range args[0] {
		if ch >= 'A' && ch <= 'Z' {
			if ch >= 'M' {
				ch -= 13
			} else {
				ch += 13
			}
		} else if ch >= 'a' && ch <= 'z' {
			if ch >= 'm' {
				ch -= 13
			} else {
				ch += 13
			}
		}
		res += string(ch)

	}

	for _, ch := range res {
		z01.PrintRune(ch)
	}

	z01.PrintRune('\n')
}
