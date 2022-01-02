package stream

func (s *stream[T]) Peek(peekFunc func(element T)) *stream[T] {
	for _, element := range s.elements {
		peekFunc(element)
	}
	return s
}

func (s *stream[T]) Ordered() *stream[T] {
	s.ordered = true
	return s
}

func (s *stream[T]) Parallel() *stream[T] {
	s.ordered = false
	return s
}
