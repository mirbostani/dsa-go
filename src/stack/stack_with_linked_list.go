package stack

import (
	"errors"
)

type Node struct {
	data int
	next *Node
}

type StackWithLinkedList struct {
	head *Node
}

func (s *StackWithLinkedList) Size() int {
	size := 0
	currentNode := s.head
	for currentNode != nil {
		size++
		currentNode = currentNode.next
	}
	return size
}

func (s *StackWithLinkedList) IsEmpty() bool {
	return s.head == nil
}

func (s *StackWithLinkedList) Push(item int) {
	newNode := &Node{data: item, next: s.head}
	s.head = newNode
}

func (s *StackWithLinkedList) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Pop from an empty stack")
	}
	poppedNode := s.head
	s.head = s.head.next
	return poppedNode.data, nil
}

func (s *StackWithLinkedList) Top() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Peek from an empty stack")
	}

	return s.head.data, nil
}

func (s *StackWithLinkedList) Peek() (int, error) {
	return s.Top()
}
