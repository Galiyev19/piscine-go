package piscine

func Compact(ptr *[]string) int {
	add := 0
	var arr []string

	for _, ch := range *ptr {
		if ch != "" {
			arr = append(arr, ch)
			add++
		}
	}
	*ptr = arr
	return add
}
