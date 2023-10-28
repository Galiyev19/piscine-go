package piscine

func Unmatch(a []int) int {
	for _, v := range a {
		count := 0
		for _, el := range a {
			if v == el {
				count++
			}
		}
		if count%2 != 0 {
			return v
		}
	}
	return -1
}
