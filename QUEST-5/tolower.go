package piscine

func ToLower(s string) string {
	res := []rune(s)

	for i := 0; i < len(res); i++ {
		if s[i] >= 65 && s[i] <= 90 {
			res[i] += 32
		}
	}

	return string(res)
}
