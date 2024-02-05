package queue

import (
	"errors"
)

type QueueWithDynamicArray struct {
	items []int
}

func (q *QueueWithDynamicArray) Size() int {
	return len(q.items)
}

func (q *QueueWithDynamicArray) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *QueueWithDynamicArray) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *QueueWithDynamicArray) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Dequeue from an empty queue")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

func (q *QueueWithDynamicArray) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Peek from an empty queue")
	}
	return q.items[0], nil
}

func (q *QueueWithDynamicArray) Front() (int, error) {
	return q.Peek()
}
