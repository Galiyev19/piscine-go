package piscine

func Atoi(s string) int {
	var result int
	sign := 1
	count := 0
	digitFound := false

	for _, char := range s {

		if char == '+' || char == '-' {
			count++
		}

		if count > 1 {
			return 0
		}

		if char == '-' || char == '+' {
			if digitFound {
				return 0
			}
			if char == '-' {
				sign = -1
			}
		} else if char >= '0' && char <= '9' {
			digit := int(char - '0')
			result = result*10 + digit
			digitFound = true
		} else {
			return 0
		}
	}

	return sign * result
}
