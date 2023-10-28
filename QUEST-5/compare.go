package piscine

func Compare(a, b string) int {
	res := 1
	if a < b {
		res = -1
	}

	if a == b {
		res = 0
	}

	return res
}
