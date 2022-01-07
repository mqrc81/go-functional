package stream

import (
	"go-functional/check"
)

// TODO streamOld
//  - Map (to T, string, bool, int)
//  - Sort
//  - Distinct

type Stream[T any] interface {
	Filter(filterFunc func(element T) bool) stream[T]
	Limit(maxSize int) stream[T]
	Reverse() stream[T]
	Peek(func(elements T)) stream[T]
	Ordered() stream[T]
	Parallel() stream[T]

	Collect() []T
	Count() int
	AnyMatch(matchFunc func(element T) bool) bool
	AllMatch(matchFunc func(element T) bool) bool
	NoneMatch(matchFunc func(element T) bool) bool
	Any() bool
	ForEach(forEachFunc func(element T))
	Find(findFunc func(element T) bool) check.Check[T]
	Fold(initialValue T, foldFunc func(value T, element T) T) T
	FoldInt(initialValue int, foldFunc func(value int, element T) int) int
	FoldString(initialValue string, foldFunc func(value string, element T) string) string
}

type stream[T any] struct {
	elements   []T
	operations []operation[T]
	ordered    bool
}
