package piscine

func ListMerge(l1 *List, l2 *List) {
	list1 := l1.Head
	list2 := l2.Head

	if list1 == nil {
		l1.Head = list2
		l1.Tail = l2.Tail
	} else if list2 != nil {
		for list1.Next != nil {
			list1 = list1.Next
		}

		list1.Next = list2

		l1.Tail = l2.Tail
	}

	l2.Head = nil
	l2.Tail = nil
}
