package helper

type Set struct {
	set map[any]bool
}

func NewSet() Set {
	return Set{
		set: make(map[any]bool),
	}
}

func (s *Set) Insert(k any) {
	s.set[k] = true
}

func (s *Set) Find(k any) bool {
	return s.set[k]
}

func (s *Set) Len() int {
	return len(s.set)
}
