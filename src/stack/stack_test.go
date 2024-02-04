package stack

import "testing"

func AssertEqual(t *testing.T, actual, expected interface{}) {
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func AssertIsNil(t *testing.T, actual interface{}) {
	if actual != nil {
		t.Errorf("Expected nil, but got %v", actual)
	}
}

func AssertIsNotNil(t *testing.T, actual interface{}) {
	if actual == nil {
		t.Errorf("Expected not nil, but got nil")
	}
}

func AssertTrue(t *testing.T, actual interface{}) {
	if actual != true {
		t.Errorf("Expected true, but got %v", actual)
	}
}

func TestStackWithDynamicArrayInit(t *testing.T) {
	stack := StackWithDynamicArray{}
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)

	size := stack.Size()
	AssertEqual(t, size, 3)
}

func TestStackWithDynamicArrayPushAndPop(t *testing.T) {
	stack := StackWithDynamicArray{}
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)

	val, err := stack.Pop()
	AssertEqual(t, val, 1)
	AssertIsNil(t, err)

	val, err = stack.Pop()
	AssertEqual(t, val, 2)
	AssertIsNil(t, err)

	val, err = stack.Pop()
	AssertEqual(t, val, 3)
	AssertIsNil(t, err)

	AssertTrue(t, stack.IsEmpty())
}

func TestStackWithDynamicArrayPeek(t *testing.T) {
	stack := StackWithDynamicArray{}
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)

	val, err := stack.Peek()
	AssertEqual(t, val, 1)
	AssertIsNil(t, err)
	AssertEqual(t, stack.Size(), 3)
}

func TestStackWithLinkedListInit(t *testing.T) {
	stack := &StackWithLinkedList{}
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)

	AssertEqual(t, stack.Size(), 3)
}

func TestStackWithLinkedListPushAndPop(t *testing.T) {
	stack := &StackWithLinkedList{}
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)

	val, err := stack.Pop()
	AssertEqual(t, val, 1)
	AssertIsNil(t, err)

	val, err = stack.Pop()
	AssertEqual(t, val, 2)
	AssertIsNil(t, err)

	val, err = stack.Pop()
	AssertEqual(t, val, 3)
	AssertIsNil(t, err)

	AssertTrue(t, stack.IsEmpty())
}

func TestStackWithLinkedListPeek(t *testing.T) {
	stack := &StackWithLinkedList{}
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)

	val, err := stack.Peek()
	AssertEqual(t, val, 1)
	AssertIsNil(t, err)
	AssertEqual(t, stack.Size(), 3)
}

func TestStackWithFixedSizeArrayInit(t *testing.T) {
	stack := NewStackWithFixedSizeArray(10)
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)

	AssertEqual(t, stack.Size(), 3)
	AssertEqual(t, stack.maxSize, 10)
}

func TestStackWithFixedSizeArrayPushAndPop(t *testing.T) {
	stack := NewStackWithFixedSizeArray(10)
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)

	val, err := stack.Pop()
	AssertEqual(t, val, 1)
	AssertIsNil(t, err)

	val, err = stack.Pop()
	AssertEqual(t, val, 2)
	AssertIsNil(t, err)

	val, err = stack.Pop()
	AssertEqual(t, val, 3)
	AssertIsNil(t, err)

	AssertTrue(t, stack.IsEmpty())
}

func TestStackWithFixedSizeArrayFull(t *testing.T) {
	stack := NewStackWithFixedSizeArray(3)
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)

	AssertTrue(t, stack.IsFull())

	err := stack.Push(0)
	AssertIsNotNil(t, err)
}

func TestStackWithFixedSizeArrayEmpty(t *testing.T) {
	stack := NewStackWithFixedSizeArray(3)

	AssertTrue(t, stack.IsEmpty())

	_, err := stack.Pop()
	AssertIsNotNil(t, err)
}
