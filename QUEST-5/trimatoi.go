package piscine

func atoi(s string) int {
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
				continue
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

func TrimAtoi(s string) int {
	res := 0
	str := []rune(s)
	temp := ""

	for _, value := range str {
		if value >= '0' && value <= '9' || value == '-' {
			temp += string(value)
		}
	}

	res = atoi(temp)
	return res
}
