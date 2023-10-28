package piscine

func ListClear(l *List) {
	current := l.Head

	for current != nil {
		current = current.Next
		current = nil
	}

	l.Head, l.Tail = nil, nil
}
