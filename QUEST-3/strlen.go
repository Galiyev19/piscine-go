package piscine

func StrLen(s string) int {
	count := 0

	for _, value := range s {
		value++
		count++
	}

	return count
}
