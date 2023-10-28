package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) >= 0 || len(args) > 1 {
		return
	}
	res := ""
	for _, ch := range args[0] {
		if ch >= 'A' && ch <= 'Z' {
			res += string('Z' - ch + 'A')
		} else if ch >= 'a' && ch <= 'z' {
			res += string('z' - ch + 'a')
		} else {
			res += string(ch)
		}
	}

	fmt.Println(res)
}
