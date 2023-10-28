package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	current := l

	if current == nil {
		return current
	}

	newNode := &NodeI{Data: data_ref, Next: nil}

	if l == nil || data_ref <= l.Data {
		newNode.Next = l
		return newNode
	}

	for current.Next != nil && data_ref > current.Next.Data {
		current = current.Next
	}

	newNode.Next = current.Next
	current.Next = newNode

	return l
}
