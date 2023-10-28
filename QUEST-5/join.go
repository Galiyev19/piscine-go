package piscine

func Join(strs []string, sep string) string {
	res := ""
	// res += strs[0]
	// res += sep

	for index, str := range strs {
		if index != len(strs)-1 {
			str += sep
		}
		res += str
	}

	return res
}
