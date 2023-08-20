package ds

import (
	"errors"
	"slices"
)

var ErrQueueEmpty = errors.New("queue is empty")

type SliceQueue[T any] struct {
	queue []T
}

func NewSliceQueue[T any]() *SliceQueue[T] {
	return &SliceQueue[T]{
		queue: make([]T, 0),
	}
}

func (q *SliceQueue[T]) Enqueue(val T) {
	q.queue = append(q.queue, val)
}

func (q *SliceQueue[T]) Dequeue() (T, error) {
	if len(q.queue) == 0 {
		return *new(T), ErrQueueEmpty
	}
	ret := q.queue[0]
	q.queue = slices.Delete(q.queue, 0, 1)
	return ret, nil
}

func (q *SliceQueue[T]) Peek() (T, error) {
	if len(q.queue) == 0 {
		return *new(T), ErrQueueEmpty
	}
	return q.queue[0], nil
}

func (q *SliceQueue[T]) Empty() bool {
	return len(q.queue) == 0
}
