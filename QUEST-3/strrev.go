package piscine

func StrRev(s string) string {
	reverseString := []byte(s)
	result := ""
	for i := len(reverseString) - 1; i >= 0; i-- {
		// z01.PrintRune(rune(reverseString[i]))
		result += string(rune(reverseString[i]))
	}

	return result
}
