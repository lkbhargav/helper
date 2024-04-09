package helper

import "fmt"

type Stack struct {
	val []any
}

func NewStack() Stack {
	return Stack{
		val: []any{},
	}
}

func (s *Stack) Push(v any) {
	s.val = append(s.val, v)
}

func (s *Stack) Pop() any {
	toReturn := s.val[len(s.val)-1]
	s.val = s.val[:len(s.val)-1]
	return toReturn
}

func (s *Stack) Len() int {
	return len(s.val)
}

func (s *Stack) Print() {
	fmt.Println(s.val...)
}
