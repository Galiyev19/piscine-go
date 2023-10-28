package main

import (
	"fmt"
)

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{Data: data}

	newNode.Next = l.Head
	l.Head = newNode

	if l.Tail == nil {
		l.Tail = newNode
	}
}

func ListSize(l *List) int {
	current := l.Head
	if current == nil {
		return 0
	}
	count := 0

	for current != nil {
		current = current.Next
		count++
	}

	return count
}

func main() {
	link := &List{}

	ListPushFront(link, "Hello")
	ListPushFront(link, "2")
	ListPushFront(link, "you")
	ListPushFront(link, "man")

	fmt.Println(ListSize(link))
}
