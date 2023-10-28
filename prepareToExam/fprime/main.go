package main

import "fmt"

func main() {
	num := 42
	factors := FPrime(num)

	if len(factors) > 0 {
		fmt.Print(factors)
		fmt.Print("\n")
	}
}

func FPrime(n int) string {
	res := ""

	for i := 2; i <= n; i++ {
		if IsPrime(i) && n%i == 0 {
			res += string(i + '0')
			n /= i
			if n > 1 {
				res += "*"
			}
		}
	}
	return res
}

func IsPrime(n int) bool {
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
