package piscine

func ListSize(l *List) int {
	current := l.Head
	count := 0
	for current != nil {
		current = current.Next
		count++
	}
	return count
}
