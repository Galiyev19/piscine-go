package piscine

func IsPrintable(s string) bool {
	for _, value := range s {
		if value >= 0 && value <= 31 {
			return false
		}
	}

	return true
}
