package piscine

func StringToIntSlice(str string) []int {
	var arr []int

	for _, s := range str {
		arr = append(arr, int(rune(s)))
	}
	return arr
}
