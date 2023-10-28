package piscine

func Index(s string, toFind string) int {
	if len(s) < 0 {
		return 0
	}
	sentence := []rune(s)
	for i := 0; i < len(sentence); i++ {
		if s[i] == toFind[0] {
			return i
		}
	}

	return -1
}
