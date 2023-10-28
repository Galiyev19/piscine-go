package piscine

func Rot14(s string) string {
	result := ""
	for _, str := range s {
		if str >= 'a' && str <= 'z' {
			if str >= 'm' {
				str -= 12
			} else {
				str += 14
			}
		} else if str >= 'A' && str <= 'Z' {
			if str >= 'M' {
				str -= 12
			} else {
				str += 14
			}
		}
		result += string(str)
	}

	return result
}
