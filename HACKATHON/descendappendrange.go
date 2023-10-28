package piscine

func DescendAppendRange(max, min int) []int {
	var res []int
	if min > max || max == 0 && min == 0 {
		return []int{}
	}
	for i := max; i > min; i-- {
		res = append(res, i)
	}

	return res
}
