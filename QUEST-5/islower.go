package piscine

func IsLower(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= 0 && s[i] < 90 || s[i] >= 123 && s[i] <= 127 {
			return false
		}
	}
	return true
}
