package piscine

func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}

	result := make([]rune, len(s))
	nextToUpper := true

	for i, char := range s {
		isAlphanumeric := (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')

		if isAlphanumeric {
			if nextToUpper {
				// Capitalize the first letter of the word
				if char >= 'a' && char <= 'z' {
					char -= 'a' - 'A'
				}
				result[i] = char
				nextToUpper = false
			} else {
				// Convert to lowercase for the rest of the word
				if char >= 'A' && char <= 'Z' {
					char += 'a' - 'A'
				}
				result[i] = char
			}
		} else {
			// Non-alphanumeric character, preserve as is
			result[i] = char
			nextToUpper = true
		}
	}

	return string(result)
}
