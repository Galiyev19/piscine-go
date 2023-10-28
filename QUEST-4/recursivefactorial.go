package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb >= 21 {
		return 0
	}
	if nb == 0 || nb == 1 {
		return 1
	}

	// maxInt := 9223372036854775807

	if nb != 1 {
		nb = nb * RecursiveFactorial(nb-1)
	}

	if nb < 0 {
		return 0
	}

	if nb == 0 {
		return 1
	}

	return nb
}
