package piscine

func AtoiBase(s string, base string) int {
	length := len(base)
	res := 0

	if !isValidBase(base) {
		return res
	}

	for _, str := range s {
		digit := string(str)
		index := 0

		for i, ch := range base {
			if string(ch) == digit {
				index = i
				break
			}
		}
		res = res*length + index
	}

	return res
}
