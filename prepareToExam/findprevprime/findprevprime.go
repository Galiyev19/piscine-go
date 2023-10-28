package main

import "fmt"

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}

func FindPrevPrime(nb int) int {
	prime := IsPrime(nb)
	if prime {
		return nb
	}

	prev := nb - 1

	for {
		if IsPrime(prev) {
			return prev
		}
		prev--
	}
}

func IsPrime(num int) bool {
	if num < 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}
