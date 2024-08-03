package helper

type Set[T comparable] struct {
	set map[T]bool
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{
		set: make(map[T]bool),
	}
}

func (s *Set[T]) Insert(k T) {
	s.set[k] = true
}

func (s *Set[T]) Find(k T) bool {
	return s.set[k]
}

func (s *Set[T]) Len() int {
	return len(s.set)
}
