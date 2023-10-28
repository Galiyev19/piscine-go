package piscine

func LoafOfBread(str string) string {
	if str == "" {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	var _str string
	j := 0
	for i := 0; i < len(str); i++ {
		if j < 5 && rune(str[i]) == ' ' {
			continue
		}
		if j == 5 {
			if i != len(str)-1 && str[i+1] == ' ' {
				continue
			}
			if i == len(str)-1 {
				break
			}
			_str += " "
			j = 0
			continue
		}
		_str += string(rune(str[i]))
		j++
	}
	_str += "\n"
	return _str
}
