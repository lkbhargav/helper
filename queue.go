package helper

import "fmt"

type Queue[T comparable] struct {
	val []T
}

func NewQueue[T comparable]() Queue[T] {
	return Queue[T]{
		val: []T{},
	}
}

func (q *Queue[T]) EnQueue(v T) {
	q.val = append(q.val, v)
}

func (q *Queue[T]) DeQueue() T {
	toReturn := q.val[0]
	q.val = q.val[1:]

	return toReturn
}

func (q *Queue[T]) Len() int {
	return len(q.val)
}

func (q *Queue[T]) Print() {
	fmt.Printf("q.val: %v\n", q.val)
}
