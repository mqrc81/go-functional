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

func (s stream[T]) Reverse() stream[T] {
	if s.parallel {
		return s
	}
	var reversed []T
	for _, element := range s.elements {
		if s.terminate(&element) {
			prepend(&reversed, element)
		}
	}
	s.elements = reversed
	s.operations = nil
	return s
}

func (s stream[T]) Peek(peekFunc func(element T)) stream[T] {
	s.operations = append(s.operations, peekOperation[T]{
		peekFunc: peekFunc,
	})
	return s
}

func (s stream[T]) Parallel() stream[T] {
	s.parallel = true
	return s
}

func (s stream[T]) Ordered() stream[T] {
	s.parallel = false
	return s
}

func prepend[T any](slice *[]T, element T) {
	*slice = append(*slice, element)
	copy((*slice)[1:], *slice)
	(*slice)[0] = element
}
