package stream

func (s stream[T]) Filter(filterFunc func(element T) bool) stream[T] {
	s.applyFilter(filterFunc)
	return s
}

func (s stream[T]) Reverse() stream[T] {
	s.applyReverse()
	return s
}

func (s stream[T]) Limit(maxSize int) stream[T] {
	if len(s.elements) >= maxSize {
		s.elements = s.elements[:maxSize]
	}
	return s
}

func (s stream[T]) Reduce(maxSize int, filterFunc func(element T) bool) stream[T] {
	if len(s.elements) >= maxSize {
		s.applyFilter(filterFunc)
	}
	if len(s.elements) >= maxSize {
		s.elements = s.elements[:maxSize]
	}
	return s
}
