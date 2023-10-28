package piscine

func Concat(str1 string, str2 string) string {
	str1 += str2
	temp := []rune(str1)
	res := temp
	return string(res)
}
