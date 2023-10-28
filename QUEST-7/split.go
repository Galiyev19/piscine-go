package piscine

func Split(s, sep string) []string {
	var strArr []string
	index := 0

	for i := 0; i < len(s); i++ {
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			strArr = append(strArr, s[index:i])
			index = i + len(sep)
			i += len(sep) - 1
		}
	}

	if index < len(s) {
		strArr = append(strArr, s[index:])
	}

	return strArr
}
