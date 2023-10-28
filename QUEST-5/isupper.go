package piscine

func IsUpper(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= 0 && s[i] < 65 || s[i] >= 91 && s[i] <= 127 {
			return false
		}
	}
	return true
}
