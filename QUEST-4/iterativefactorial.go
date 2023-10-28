package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}

	factorial := 1
	maxInt := 9223372036854775807

	for i := 1; i <= nb; i++ {

		if factorial > maxInt/i {
			return 0
		}
		factorial *= i
	}

	return factorial
}
