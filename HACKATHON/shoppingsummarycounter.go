package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	arr := []string{}

	word := ""

	for _, v := range str {
		if v == ' ' {
			arr = append(arr, string(word))
			word = ""
		} else {
			word += string(v)
		}
	}

	arr = append(arr, string(word))
	res := make(map[string]int)

	for _, word := range arr {
		res[word]++
	}
	return res
}
