package stack

import (
	"testing"

	testingutil "github.com/mirbostani/dsa-go/src/utils"
)

func TestStackWithDynamicArrayInit(t *testing.T) {
	stack := StackWithDynamicArray{}
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)

	size := stack.Size()
	testingutil.AssertEqual(t, size, 3)
}

func TestStackWithDynamicArrayPushAndPop(t *testing.T) {
	stack := StackWithDynamicArray{}
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)

	val, err := stack.Pop()
	testingutil.AssertEqual(t, val, 1)
	testingutil.AssertIsNil(t, err)

	val, err = stack.Pop()
	testingutil.AssertEqual(t, val, 2)
	testingutil.AssertIsNil(t, err)

	val, err = stack.Pop()
	testingutil.AssertEqual(t, val, 3)
	testingutil.AssertIsNil(t, err)

	testingutil.AssertTrue(t, stack.IsEmpty())
}

func TestStackWithDynamicArrayPeek(t *testing.T) {
	stack := StackWithDynamicArray{}
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)

	val, err := stack.Peek()
	testingutil.AssertEqual(t, val, 1)
	testingutil.AssertIsNil(t, err)
	testingutil.AssertEqual(t, stack.Size(), 3)
}

func TestStackWithLinkedListInit(t *testing.T) {
	stack := &StackWithLinkedList{}
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)

	testingutil.AssertEqual(t, stack.Size(), 3)
}

func TestStackWithLinkedListPushAndPop(t *testing.T) {
	stack := &StackWithLinkedList{}
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)

	val, err := stack.Pop()
	testingutil.AssertEqual(t, val, 1)
	testingutil.AssertIsNil(t, err)

	val, err = stack.Pop()
	testingutil.AssertEqual(t, val, 2)
	testingutil.AssertIsNil(t, err)

	val, err = stack.Pop()
	testingutil.AssertEqual(t, val, 3)
	testingutil.AssertIsNil(t, err)

	testingutil.AssertTrue(t, stack.IsEmpty())
}

func TestStackWithLinkedListPeek(t *testing.T) {
	stack := &StackWithLinkedList{}
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)

	val, err := stack.Peek()
	testingutil.AssertEqual(t, val, 1)
	testingutil.AssertIsNil(t, err)
	testingutil.AssertEqual(t, stack.Size(), 3)
}

func TestStackWithFixedSizeArrayInit(t *testing.T) {
	stack := NewStackWithFixedSizeArray(10)
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)

	testingutil.AssertEqual(t, stack.Size(), 3)
	testingutil.AssertEqual(t, stack.maxSize, 10)
}

func TestStackWithFixedSizeArrayPushAndPop(t *testing.T) {
	stack := NewStackWithFixedSizeArray(10)
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)

	val, err := stack.Pop()
	testingutil.AssertEqual(t, val, 1)
	testingutil.AssertIsNil(t, err)

	val, err = stack.Pop()
	testingutil.AssertEqual(t, val, 2)
	testingutil.AssertIsNil(t, err)

	val, err = stack.Pop()
	testingutil.AssertEqual(t, val, 3)
	testingutil.AssertIsNil(t, err)

	testingutil.AssertTrue(t, stack.IsEmpty())
}

func TestStackWithFixedSizeArrayFull(t *testing.T) {
	stack := NewStackWithFixedSizeArray(3)
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)

	testingutil.AssertTrue(t, stack.IsFull())

	err := stack.Push(0)
	testingutil.AssertIsNotNil(t, err)
}

func TestStackWithFixedSizeArrayEmpty(t *testing.T) {
	stack := NewStackWithFixedSizeArray(3)

	testingutil.AssertTrue(t, stack.IsEmpty())

	_, err := stack.Pop()
	testingutil.AssertIsNotNil(t, err)
}
