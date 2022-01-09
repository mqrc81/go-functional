package stream

import (
	"go-functional/check"
)

// TODO stream
//  - CONCURRENT

type Stream[T any] interface {
	Filter(filterFunc func(element T) bool) stream[T]
	Limit(maxSize int) stream[T]
	Reverse() stream[T]
	Sort(sortFunc func(element1 T, element2 T) bool) stream[T]
	Map(mapFunc func(element T) T) stream[T]
	MapToInt(mapFunc func(element T) int) stream[int]
	MapToString(mapFunc func(element T) string) stream[string]
	Peek(func(elements T)) stream[T]
	Ordered() stream[T]
	Parallel() stream[T]

	Collect() []T
	Count() int
	Concat(delimiter string) string
	AnyMatch(matchFunc func(element T) bool) bool
	AllMatch(matchFunc func(element T) bool) bool
	NoneMatch(matchFunc func(element T) bool) bool
	ForEach(forEachFunc func(element T))
	Find(findFunc func(element T) bool) check.Check[T]
	Fold(initialValue T, foldFunc func(value T, element T) T) T
	FoldToInt(initialValue int, foldFunc func(value int, element T) int) int
	FoldToString(initialValue string, foldFunc func(value string, element T) string) string
}

type stream[T any] struct {
	elements   []T
	operations []operation[T]
	parallel   bool
}
