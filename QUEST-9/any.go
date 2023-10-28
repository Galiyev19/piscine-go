package piscine

func Any(f func(string) bool, a []string) bool {
	res := false
	for _, ch := range a {
		if f(ch) {
			res = true
		}
	}
	return res
}
