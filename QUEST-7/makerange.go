package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}

	size := max - min
	arr := make([]int, size)

	if min == 0 {
		arr = make([]int, 1)
		return arr
	}

	for i := range arr {
		arr[i] = min + i
	}

	return arr
}
