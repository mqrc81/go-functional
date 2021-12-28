package option

type Option[T any] struct {
	value *T
	// TODO
}

func Of[T any](value T) Option[T] {
	return Option[T]{
		&value,
	}
}
