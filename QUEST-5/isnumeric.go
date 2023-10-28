package piscine

func IsNumeric(s string) bool {
	for _, value := range s {
		if value >= 0 && value < 48 || value >= 58 && value <= 127 {
			return false
		}
	}

	return true
}
