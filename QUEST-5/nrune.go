package piscine

func NRune(s string, n int) rune {
	nth := []rune(s)

	if n <= 0 || n > len(s) {
		return 0
	}

	return nth[n-1]
}
