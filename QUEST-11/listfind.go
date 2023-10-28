package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	current := l.Head
	val := &ref
	for current != nil {
		if comp(current, current.Next) {
		}
		current = current.Next
	}

	return val
}
