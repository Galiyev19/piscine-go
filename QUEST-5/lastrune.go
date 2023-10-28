package piscine

func LastRune(s string) rune {
	res := []rune(s)

	index := 0
	for i := 0; i < len(res); i++ {
		if i == len(res)-1 {
			index = i
		}
	}
	return res[index]
}
