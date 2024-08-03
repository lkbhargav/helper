package helper

import "fmt"

type Stack[T comparable] struct {
	val []T
}

func NewStack[T comparable]() Stack[T] {
	return Stack[T]{
		val: []T{},
	}
}

func (s *Stack[T]) Push(v T) {
	s.val = append(s.val, v)
}

func (s *Stack[T]) Pop() T {
	toReturn := s.val[len(s.val)-1]
	s.val = s.val[:len(s.val)-1]
	return toReturn
}

func (s *Stack[T]) Len() int {
	return len(s.val)
}

func (s *Stack[T]) Print() {
	fmt.Printf("s.val: %v\n", s.val)
}
