package piscine

func BasicAtoi(s string) int {
	result := 0

	for i, value := range s {
		digit := int(value - '0')

		result += digit

		if i != len(s)-1 {
			result *= 10
		}
	}

	return result
}
