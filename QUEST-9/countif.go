package piscine

func CountIf(f func(string) bool, tab []string) int {
	index := 0

	for _, ch := range tab {
		if f(ch) {
			index += 1
			// break
		}
	}

	return index
}
