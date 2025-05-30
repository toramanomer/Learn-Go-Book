package main

import (
	"fmt"
)

type IntOrFloat interface {
	int | uint | uintptr |
		int8 | int16 | int32 | int64 |
		uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

func Double[T IntOrFloat](v T) T {
	return v * v
}

type Printable interface {
	fmt.Stringer
	~int | ~float64
}

func Print[T Printable](p T) {
	fmt.Println(p)
}

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

type LinkedList[T comparable] struct {
	start *Node[T]
	end   *Node[T]
}

func (l *LinkedList[T]) Add(val T) {
	if l == nil {
		l = &LinkedList[T]{}
	}
	newNode := &Node[T]{value: val}

	if l.start == nil {
		l.start = newNode
		l.end = newNode
		return
	}
	l.end.next = newNode
	l.end = l.end.next
}

func (l *LinkedList[T]) Insert(value T, index int) {
	newNode := &Node[T]{value: value}
	if l.start == nil {
		l.start = newNode
		l.end = newNode
		return
	}

	if index <= 0 {
		newNode.next = l.end
		l.start = newNode
		return
	}

	curNode := l.start

	for i := 1; i < index; i++ {
		if curNode.next == nil {
			curNode.next = newNode
			l.end = curNode.next
			return
		}
		curNode = curNode.next
	}
	newNode.next = curNode.next
	curNode.next = newNode
	if l.end == curNode {
		l.end = newNode
	}

}

func (l *LinkedList[T]) Index(element T) int {
	return -1
}

func main() {
	var linkedList *LinkedList[int]
	fmt.Println(linkedList)
	linkedList.Add(1)
	fmt.Println(linkedList)
}
