package queue

import (
	"testing"

	testingutil "github.com/mirbostani/dsa-go/src/utils"
)

func TestQueueWithDynamicArrayInit(t *testing.T) {
	t.Log("Should create a simple queue with 3 elements")

	queue := &QueueWithDynamicArray{}
	queue.Enqueue(3)
	queue.Enqueue(2)
	queue.Enqueue(1)

	testingutil.AssertEqual(t, queue.Size(), 3)
}

func TestQueueWithDynamicArrayEnqueueDequeue(t *testing.T) {
	t.Log("Should enqueue and dequeue following FIFO principle")

	queue := &QueueWithDynamicArray{}
	queue.Enqueue(3)
	queue.Enqueue(2)
	queue.Enqueue(1)

	val, err := queue.Dequeue()
	testingutil.AssertEqual(t, val, 3)
	testingutil.AssertIsNil(t, err)

	val, err = queue.Dequeue()
	testingutil.AssertEqual(t, val, 2)
	testingutil.AssertIsNil(t, err)

	val, err = queue.Dequeue()
	testingutil.AssertEqual(t, val, 1)
	testingutil.AssertIsNil(t, err)

	testingutil.AssertTrue(t, queue.IsEmpty())
}

func TestQueueWithDynamicArrayPeek(t *testing.T) {
	t.Log("Should peek the front element")

	queue := &QueueWithDynamicArray{}
	queue.Enqueue(3)
	queue.Enqueue(2)
	queue.Enqueue(1)

	val, err := queue.Peek()
	testingutil.AssertEqual(t, val, 3)
	testingutil.AssertIsNil(t, err)

	testingutil.AssertEqual(t, queue.Size(), 3)
}

func TestQueueWithFixedSizeArrayInit(t *testing.T) {
	t.Log("Should create a queue with fixed-size array including 3 elements")

	maxSize := 3
	queue := NewQueueWithFixedSizeArray(maxSize)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	testingutil.AssertTrue(t, queue.IsFull())
	testingutil.AssertEqual(t, queue.front, 0)
	testingutil.AssertEqual(t, queue.rear, maxSize-1)
}

func TestQueueWithFixedSizeArrayEnqueueDequeue(t *testing.T) {
	t.Log("Should enqueue and dequeue following FIFO principle")

	maxSize := 3
	queue := NewQueueWithFixedSizeArray(maxSize)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	val, err := queue.Dequeue()
	testingutil.AssertEqual(t, val, 1)
	testingutil.AssertIsNil(t, err)

	val, err = queue.Dequeue()
	testingutil.AssertEqual(t, val, 2)
	testingutil.AssertIsNil(t, err)

	val, err = queue.Dequeue()
	testingutil.AssertEqual(t, val, 3)
	testingutil.AssertIsNil(t, err)

	testingutil.AssertTrue(t, queue.IsEmpty())
	testingutil.AssertEqual(t, queue.front, -1)
	testingutil.AssertEqual(t, queue.rear, -1)
}

func TestQueueWithFixedSizeArrayPeek(t *testing.T) {
	t.Log("Should peek the front element")

	maxSize := 3
	queue := NewQueueWithFixedSizeArray(maxSize)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	val, err := queue.Peek()
	testingutil.AssertEqual(t, val, 1)
	testingutil.AssertIsNil(t, err)
}

func TestQueueWithFixedSizeArrayDequeueMoreThanEnqueue(t *testing.T) {
	t.Log("Should prevent dequeue if it exceeds the number of enqueue")

	maxSize := 5
	queue := NewQueueWithFixedSizeArray(maxSize)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	val, err := queue.Dequeue()
	testingutil.AssertEqual(t, val, 1)
	testingutil.AssertIsNil(t, err)

	val, err = queue.Dequeue()
	testingutil.AssertEqual(t, val, 2)
	testingutil.AssertIsNil(t, err)

	val, err = queue.Dequeue()
	testingutil.AssertEqual(t, val, 3)
	testingutil.AssertIsNil(t, err)

	_, err = queue.Dequeue()
	testingutil.AssertIsNotNil(t, err)
}
