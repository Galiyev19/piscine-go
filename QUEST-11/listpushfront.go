package piscine

func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{Data: data}

	newNode.Next = l.Head
	l.Head = newNode

	if l.Tail == nil {
		l.Tail = newNode
	}
}
