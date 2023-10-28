package main

import "fmt"

func main() {
	fmt.Println(Itoa(123))
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	res := ""
	isNegative := false
	if n < 0 {
		isNegative = true
		n = -n
	}
	for n > 0 {
		digit := n % 10
		res = string('0' + digit)
		n /= 10
	}

	if isNegative {
		res += "-"
	}

	revStr := ""

	for i := len(res) - 1; i >= 0; i-- {
		revStr += string(res[i])
	}

	return revStr
}
