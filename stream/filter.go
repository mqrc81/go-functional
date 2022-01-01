package stream

func (s *stream[T]) Filter(method func(element T) bool) *stream[T] {
	s.sequence = append(s.sequence, action[T]{
		actionType: filter,
		filterFunc: method,
	})
	return s
}
