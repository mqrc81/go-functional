package stream

func Of[T any](values []T) *stream[T] {
	return &stream[T]{
		values:  values,
		actions: []action{},
	}
}

func From[T any](values ...T) *stream[T] {
	return &stream[T]{
		values:  values,
		actions: []action{},
	}
}

func Empty[T any]() *stream[T] {
	return &stream{
		actions: []action{},
	}
}
