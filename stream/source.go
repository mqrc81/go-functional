package stream

import (
	"constraints"
	"fmt"
)

func Of[T any](elements []T, options ...flag) stream[T] {
	parallel := fromflags(options)
	return stream[T]{
		elements: elements,
		parallel: parallel,
	}
}

func OfElements[T any](elements ...T) stream[T] {
	return stream[T]{
		elements: elements,
	}
}

func Empty[T any]() stream[T] {
	return stream[T]{}
}

func OfMapKeys[T comparable, V any](aMap map[T]V, options ...flag) stream[T] {
	parallel := fromflags(options)
	keys := make([]T, len(aMap))
	i := 0
	for k := range aMap {
		keys[i] = k
		i++
	}
	return stream[T]{
		elements: keys,
		parallel: parallel,
	}
}

func OfMapVals[K comparable, T any](aMap map[K]T, options ...flag) stream[T] {
	parallel := fromflags(options)
	vals := make([]T, len(aMap))
	i := 0
	for _, v := range aMap {
		vals[i] = v
		i++
	}
	return stream[T]{
		elements: vals,
		parallel: parallel,
	}
}

func OfRange[T constraints.Integer](start T, end T, options ...flag) stream[T] {
	diff := (end + 1) - start
	if diff < 0 {
		return Empty[T]()
	}

	parallel := fromflags(options)
	s := stream[T]{
		elements: make([]T, diff),
		parallel: parallel,
	}
	for num := start; num <= end; num++ {
		s.elements[num-start] = num
	}
	return s
}

type flag int

const (
	_ flag = iota

	// Parallel enhances performance by executing all operations concurrently,
	// thus disregarding the initial order. It is disabled by default.
	Parallel
)

func fromflags(options []flag) (parallel bool) {
	for _, option := range options {
		switch option {
		case Parallel:
			parallel = true
		default:
			panic(fmt.Sprint(option, " is not a valid flag argument"))
		}
	}
	return parallel
}
