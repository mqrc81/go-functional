package stream

import (
	"go-functional/check"
)

type Collector[T any] func(stream stream[T]) []T

func (s stream[T]) Collect(collector Collector[T]) []T {
	var collectorFunc func(stream stream[T]) []T = collector
	return collectorFunc(s)
}

func ToSlice[T any]() Collector[T] {
	return func(stream stream[T]) []T {
		return stream.elements
	}
}

func (s *stream[T]) Count() int {
	return len(s.elements)
}

func (s stream[T]) AnyMatch(matchFunc func(element T) bool) bool {
	for _, el := range s.elements {
		if matchFunc(el) {
			return true
		}
	}
	return false
}

func (s stream[T]) AllMatch(matchFunc func(element T) bool) bool {
	for _, el := range s.elements {
		if !matchFunc(el) {
			return false
		}
	}
	return true
}

func (s stream[T]) NoneMatch(matchFunc func(element T) bool) bool {
	for _, el := range s.elements {
		if matchFunc(el) {
			return false
		}
	}
	return true
}

func (s stream[T]) ForEach(forEachFunc func(element T)) {
	for _, el := range s.elements {
		forEachFunc(el)
	}
}

func (s stream[T]) Find(matchFunc func(element T) bool) check.Check[T] {
	for _, element := range s.elements {
		if matchFunc(element) {
			return check.Of[T](element, check.Valid)
		}
	}
	return check.Empty[T]()
}
