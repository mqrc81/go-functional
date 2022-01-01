package stream

import (
	"sync"
)

var (
	actionApplier sync.WaitGroup
	sliceModifier sync.Mutex
)

func (s *stream[T]) Filter(filterFunc func(element T) bool) *stream[T] {
	s.applyFilter(filterFunc)
	return s
}

func (s *stream[T]) applyFilter(filterFunc func(element T) bool) {
	var toRemove []int
	actionApplier.Add(len(s.elements))
	for i, el := range s.elements {
		go checkFilter(el, i, filterFunc, &toRemove)
	}
	actionApplier.Wait()
	s.elements = removeFromSlice(s.elements, toRemove)
}

func checkFilter[T any](element T, index int, filterFunc func(element T) bool, toRemove *[]int) {
	defer actionApplier.Done()
	if !filterFunc(element) {
		sliceModifier.Lock()
		*toRemove = append(*toRemove, index)
		sliceModifier.Unlock()
	}
}

func removeFromSlice[T any](slice []T, remove []int) []T {
	lenSlice, lenRemove := len(slice), len(remove)
	for i := len(remove) - 1; i >= 0; i-- {
		slice[remove[i]] = slice[lenSlice-(lenRemove-i)]
	}
	return slice[:len(slice)-len(remove)]
}
