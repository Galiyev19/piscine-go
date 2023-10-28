package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	num := Atoi(args[0])

	fmt.Println(IsPowerOf2(num))
}

func IsPowerOf2(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

func Atoi(s string) int {
	res := 0
	count := 0
	sign := 1
	for _, ch := range s {
		if ch == '-' || ch == '+' {
			count++
		}

		if count > 1 {
			return 0
		}

		if ch == '-' || ch == '+' {
			sign = -1
		} else if ch >= '0' && ch <= '9' {
			digit := int(ch - '0')
			res = res*10 + digit
		} else {
			return 0
		}

	}

	return res * sign
}
