package stream

type Stream[T any] struct {
	slice    *[]T
	sequence *[]action
	// TODO
}

type action struct {
	// TODO
}

func Of[T any](slice []T) *Stream[T] {
	return &Stream[T]{
		&slice,
		&[]action{},
	}
}
