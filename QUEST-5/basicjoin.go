package piscine

func BasicJoin(elems []string) string {
	res := ""

	for _, value := range elems {
		res += value
	}

	return res
}
