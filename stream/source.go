package stream

import (
	"constraints"
)

func Of[T any](elements []T, options ...flag) stream[T] {
	ordered := fromFlags(options)
	return stream[T]{
		elements: elements,
		ordered:  ordered,
	}
}

func From[T any](elements ...T) stream[T] {
	return stream[T]{
		elements: elements,
	}
}

func Empty[T any]() stream[T] {
	return stream[T]{}
}

func OfKeys[T comparable, V any](aMap map[T]V, options ...flag) stream[T] {
	ordered := fromFlags(options)
	keys := make([]T, len(aMap))
	i := 0
	for k := range aMap {
		keys[i] = k
		i++
	}
	return stream[T]{
		elements: keys,
		ordered:  ordered,
	}
}

func OfVals[K comparable, T any](aMap map[K]T, options ...flag) stream[T] {
	ordered := fromFlags(options)
	vals := make([]T, len(aMap))
	i := 0
	for _, v := range aMap {
		vals[i] = v
		i++
	}
	return stream[T]{
		elements: vals,
		ordered:  ordered,
	}
}

func OfRange[T constraints.Integer](start T, end T, options ...flag) stream[T] {
	diff := (end + 1) - start
	if diff < 0 {
		return Empty[T]()
	}

	ordered := fromFlags(options)
	s := stream[T]{
		elements: make([]T, diff),
		ordered:  ordered,
	}
	for num := start; num <= end; num++ {
		s.elements[num-start] = num
	}
	return s
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
