package src

import "sync"

type Queue[T any] struct {
	Items []T
	mtex  sync.Mutex
}

func (q *Queue[T]) Enqueue(item T) {
	q.mtex.Lock()
	defer q.mtex.Unlock()
	q.Items = append(q.Items, item)
}

func (q *Queue[T]) Dequeue() T {
	q.mtex.Lock()
	defer q.mtex.Unlock()
	val := q.Items[len(q.Items)-1]
	q.Items = q.Items[1:]
	return val
}
