package stream

func Of[T any](values []T) *stream[T] {
	return &stream[T]{
		elements: values,
		sequence: []action[T]{},
	}
}

func From[T any](values ...T) *stream[T] {
	return &stream[T]{
		elements: values,
		sequence: []action[T]{},
	}
}

func Empty[T any]() *stream[T] {
	return &stream[T]{
		sequence: []action[T]{},
	}
}
