package piscine

func Map(f func(int) bool, a []int) []bool {
	arr := make([]bool, len(a))

	for i, num := range a {
		arr[i] = f(num)
	}

	return arr
}
