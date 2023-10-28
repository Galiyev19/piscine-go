package piscine

func BasicAtoi2(s string) int {
	result := 0

	for i, value := range s {
		if rune(value) >= rune(65) && rune(value) <= rune(123) || rune(value) >= rune(32) && rune(value) <= rune(47) {
			return 0
		}
		digit := int(value - '0')

		result += digit

		if i != len(s)-1 {
			result *= 10
		}
	}

	return result
}
