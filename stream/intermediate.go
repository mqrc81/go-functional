package stream

func (s stream[T]) Filter(filterFunc func(element T) bool) stream[T] {
	s.operations = append(s.operations, filterOperation[T]{
		filterFunc: filterFunc,
	})
	return s
}

func (s stream[T]) Limit(maxSize int) stream[T] {
	s.operations = append(s.operations, limitOperation[T]{
		maxSize:    maxSize,
		spotsTaken: new(int),
	})
	return s
}

func (s stream[T]) Peek(peekFunc func(element T)) stream[T] {
	s.operations = append(s.operations, peekOperation[T]{
		peekFunc: peekFunc,
	})
	return s
}

func (s stream[T]) Ordered() stream[T] {
	s.ordered = true
	return s
}

func (s stream[T]) Parallel() stream[T] {
	s.ordered = false
	return s
}
