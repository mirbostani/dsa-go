package stack

import (
	"errors"
)

type StackWithFixedSizeArray struct {
	items   []int
	maxSize int
	head    int
}

func NewStackWithFixedSizeArray(maxSize int) *StackWithFixedSizeArray {
	return &StackWithFixedSizeArray{
		items:   make([]int, maxSize),
		maxSize: maxSize,
		head:    -1,
	}
}

func (s *StackWithFixedSizeArray) Size() int {
	return s.head + 1
}

func (s *StackWithFixedSizeArray) IsEmpty() bool {
	return s.head == -1
}

func (s *StackWithFixedSizeArray) IsFull() bool {
	return s.head == s.maxSize-1
}

func (s *StackWithFixedSizeArray) Push(item int) error {
	if s.IsFull() {
		return errors.New("Push to a full stack")
	}
	s.head++
	s.items[s.head] = item
	return nil
}

func (s *StackWithFixedSizeArray) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Pop from an empty stack")
	}

	item := s.items[s.head]
	s.head--
	return item, nil
}

func (s *StackWithFixedSizeArray) Top() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Peek from an empty stack")
	}
	item := s.items[s.head]
	return item, nil
}
