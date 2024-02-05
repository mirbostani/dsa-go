package queue

import (
	"errors"
)

type QueueWithFixedSizeArray struct {
	items   []int
	maxSize int
	rear    int
	front   int
}

func NewQueueWithFixedSizeArray(maxSize int) *QueueWithFixedSizeArray {
	return &QueueWithFixedSizeArray{
		items:   make([]int, maxSize),
		maxSize: maxSize,
		rear:    -1,
		front:   -1,
	}
}

func (q *QueueWithFixedSizeArray) IsEmpty() bool {
	return q.front == -1
}

func (q *QueueWithFixedSizeArray) IsFull() bool {
	return q.rear == q.maxSize-1
}

func (q *QueueWithFixedSizeArray) Enqueue(item int) error {
	if q.IsFull() {
		return errors.New("Enqueue to a full queue")
	}

	// First enqueue operation
	if q.front == -1 {
		q.front = 0
	}

	q.rear++
	q.items[q.rear] = item
	return nil
}

func (q *QueueWithFixedSizeArray) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Dequeue from an empty queue")
	}

	dequeuedItem := q.items[q.front]

	// Prevents dequeue if exceeds the number of enqueue
	if q.front >= q.rear {
		q.front = -1
		q.rear = -1
	} else {
		q.front++
	}

	return dequeuedItem, nil
}

func (q *QueueWithFixedSizeArray) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Peek from an empty queue")
	}

	return q.items[q.front], nil
}

func (q *QueueWithFixedSizeArray) Front() (int, error) {
	return q.Peek()
}
