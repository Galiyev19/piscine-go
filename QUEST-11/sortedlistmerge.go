package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	// Создаем фиктивный узел, чтобы облегчить код
	dummy := &NodeI{}
	current := dummy

	// Пока есть узлы в обоих списках
	for n1 != nil && n2 != nil {
		if n1.Data < n2.Data {
			// Если значение в первом списке меньше, добавляем его
			current.Next = n1
			n1 = n1.Next
		} else {
			// Иначе добавляем значение из второго списка
			current.Next = n2
			n2 = n2.Next
		}
		current = current.Next
	}

	// Если остались узлы только в первом списке, добавляем их
	if n1 != nil {
		current.Next = n1
	}

	// Если остались узлы только во втором списке, добавляем их
	if n2 != nil {
		current.Next = n2
	}

	return dummy.Next // Возвращаем голову отсортированного списка (после фиктивного узла)
}
