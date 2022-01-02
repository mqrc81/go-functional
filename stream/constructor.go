package stream

func Of[T any](values []T, options ...flag) *stream[T] {
	ordered := fromFlags(options)
	return &stream[T]{
		elements: values,
		ordered:  ordered,
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

type flag int

const (
	_ flag = iota

	// Ordered ensures that the order is preserved.
	// It is disabled by default.
	Ordered
)

func fromFlags(options []flag) (ordered bool) {
	for _, option := range options {
		switch option {
		case Ordered:
			ordered = true
		}
	}
	return ordered
}
