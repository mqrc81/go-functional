package stream

import (
	"sync"
)

var wg sync.WaitGroup

type SliceCollector[T any] func(stream stream[T]) []T

func (s *stream[T]) Collect(collector SliceCollector[T]) []T {
	var collectorFunc func(stream stream[T]) []T = collector
	return collectorFunc(*s)
}

func ToSlice[T any]() SliceCollector[T] {
	return func(stream stream[T]) []T {
		return applySequence(&stream.elements, &stream.sequence)
		// wg.Add(len(stream.elements))
		// for _, el := range stream.elements {
		// 	go applyActions(&el, stream.sequence)
		// }
		// wg.Wait()
	}
}

func applySequence[T any](elements *[]T, sequence *[]action[T]) []T {
	for _, action := range *sequence {
		switch action.actionType {
		case filter:
			*elements = applyFilter(elements, action.filterFunc)
		}
	}
	return *elements
}

func applyFilter[T any](elements *[]T, filterFunc func(el T) bool) []T {
	var toRemove []int
	for i, el := range *elements {
		if !filterFunc(el) {
			toRemove = append(toRemove, i)
		}
	}
	return removeFromSlice(*elements, toRemove)
}

func removeFromSlice[T any](slice []T, remove []int) []T { // 2, 5, 13
	for i := len(remove) - 1; i >= 0; i-- {
		slice[remove[i]] = slice[len(slice)-(len(remove)-i)]
	}
	return slice[:len(slice)-len(remove)]
}
