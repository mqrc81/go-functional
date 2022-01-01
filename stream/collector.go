package stream

type Collector[T any] func(stream stream[T]) []T

func (s *stream[T]) Collect(collector Collector[T]) []T {
	var collectorFunc func(stream stream[T]) []T = collector
	return collectorFunc(*s)
}

func ToSlice[T any]() Collector[T] {
	return func(stream stream[T]) []T {
		return stream.elements
	}
}

func (s *stream[T]) Count() int {
	return len(s.elements)
}
