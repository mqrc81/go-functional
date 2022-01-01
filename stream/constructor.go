package stream

func Of[T any](values []T) *stream[T] {
	return &stream[T]{
		elements: values,
	}
}

func From[T any](values ...T) *stream[T] {
	return &stream[T]{
		elements: values,
	}
}

func Empty[T any]() *stream[T] {
	return &stream[T]{}
}
