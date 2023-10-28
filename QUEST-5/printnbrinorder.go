package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n <= 0 {
		z01.PrintRune('0')
	}

	var digits [10]int

	var count int

	for n > 0 {
		digit := n % 10
		digits[digit]++

		n /= 10
		count++
	}

	for i := 0; i < 10; i++ {
		for digits[i] > 0 {
			z01.PrintRune(rune('0' + i))
			digits[i]--
		}
	}
}
