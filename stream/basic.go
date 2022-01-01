package stream

func (s *stream[T]) Filter(filterFunc func(element T) bool) *stream[T] {
	s.applyFilter(filterFunc)
	return s
}

func (s *stream[T]) Reverse() *stream[T] {
	s.applyReverse()
	return s
}
