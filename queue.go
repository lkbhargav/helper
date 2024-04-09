package helper

import "fmt"

type Queue struct {
	val []any
}

func NewQueue() Queue {
	return Queue{
		val: []any{},
	}
}

func (q *Queue) EnQueue(v any) {
	q.val = append(q.val, v)
}

func (q *Queue) DeQueue() any {
	toReturn := q.val[0]
	q.val = q.val[1:]

	return toReturn
}

func (q *Queue) Len() int {
	return len(q.val)
}

func (q *Queue) Print() {
	fmt.Println(q.val...)
}
