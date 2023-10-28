package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	res := true
	res1 := true

	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) < 0 {
			res = false
		}
	}

	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			res1 = false
		}
	}

	return res || res1
}
