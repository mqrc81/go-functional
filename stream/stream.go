package stream

import (
	"go-functional/check"
)

// TODO stream
//  - Map (to T, string, bool, int)
//  - Sort

type Stream[T any] interface {
	Filter(filterFunc func(element T) bool) stream[T]
	Reduce(amount int) stream[T]
	Reverse() stream[T]

	Collect(Collector[T]) []T
	Count() int
	AnyMatch(matchFunc func(element T) bool) bool
	AllMatch(matchFunc func(element T) bool) bool
	NoneMatch(matchFunc func(element T) bool) bool
	ForEach(forEachFunc func(element T))
	Find(findFunc func(element T) bool) check.Check[T]

	Peek(func(elements T)) stream[T]
	Ordered() stream[T]
	Parallel() stream[T]
}

type stream[T any] struct {
	elements []T
	ordered  bool
}
