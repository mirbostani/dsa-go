package stack

import "fmt"

type StackWithDynamicArray struct {
	items []int // slice
}

func (s *StackWithDynamicArray) Size() int {
	return len(s.items)
}

func (s *StackWithDynamicArray) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *StackWithDynamicArray) Push(item int) {
	s.items = append(s.items, item)
}

func (s *StackWithDynamicArray) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("Pop from an empty stack")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

func (s *StackWithDynamicArray) Top() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("Peek from an empty stack")
	}
	item := s.items[len(s.items)-1]
	return item, nil
}

func (s *StackWithDynamicArray) Peek() (int, error) {
	return s.Top()
}
