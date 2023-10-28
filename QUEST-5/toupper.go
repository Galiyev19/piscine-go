package piscine

func ToUpper(s string) string {
	temp := []rune(s)
	// fmt.Println(temp)
	for i := 0; i < len(temp); i++ {
		if s[i] >= 97 && s[i] <= 122 {
			temp[i] -= 32
		}
	}

	return string(temp)
}
