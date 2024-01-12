package ds

import (
	"sync"
)

type Queue[T any] struct {
	mu    sync.Mutex
	items []T
}

func NewQueue[T any](initItems []T) *Queue[T] {
	return &Queue[T]{items: initItems}
}

func (q *Queue[T]) Enqueue(item T) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() T {
	q.mu.Lock()
	defer q.mu.Unlock()
	l := len(q.items)
	if l <= 0 {
		panic("Queue already empty")
	}
	item := q.items[0]
	q.items = q.items[1:l]
	return item
}

func (q *Queue[T]) Len() int {
	return len(q.items)
}
