package piscine

func AppendRange(min, max int) []int {
	var arr []int

	if min >= max {
		return nil
	}

	if min == 0 {
		arr = append(arr, 0)
		return arr
	}

	for i := min - 1; i < max-1; i++ {
		arr = append(arr, i+1)
	}

	return arr
}
