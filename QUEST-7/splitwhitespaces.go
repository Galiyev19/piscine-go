package piscine

func countWord(s string) int {
	size := 0
	inWord := false
	for i := 0; i < len(s); i++ {
		if s[i] >= 33 && s[i] <= 127 || s[i] == '\\' {
			if !inWord {
				size++
				inWord = true
			}
		} else {
			inWord = false
		}
	}

	return size
}

func SplitWhiteSpaces(s string) []string {
	size := countWord(s)
	// fmt.Println(size)
	strArr := make([]string, size)

	startIndex := 0
	wordIndex := 0
	inWord := false

	for i, char := range s {
		if char == ' ' {
			if inWord {
				endIndex := i
				strArr[wordIndex] = s[startIndex:endIndex]
				wordIndex++
				inWord = false
			}
		} else {
			if !inWord {
				startIndex = i
				inWord = true
			}
		}
	}

	if inWord && wordIndex < size {
		strArr[wordIndex] = s[startIndex:]
	}

	return strArr
}
