package piscine

func JumpOver(str string) string {
	res := ""

	index := 2
	for i := 0; i < len(str); i++ {
		if i+index < len(str) {
			i = i + index
		} else {
			break
		}
		res += string(str[i])
	}

	if len(str) < 3 {
		res = "\n"
		return res
	} else {
		return res + "\n"
	}
}
