package stream

import (
	"sort"
	"sync"
)

var (
	actionApplier sync.WaitGroup
	sliceModifier sync.Mutex
)

func (s *stream[T]) applyFilter(filterFunc func(element T) bool) {
	if s.ordered {
		applyFilterOrdered(s, filterFunc)
	} else {
		applyFilterParallel(s, filterFunc)
	}
}

func applyFilterOrdered[T any](s *stream[T], filterFunc func(element T) bool) {
	var newElements []T
	for i, el := range s.elements {
		checkFilterOrdered(el, i, filterFunc, &newElements)
	}
	s.elements = newElements
}

func checkFilterOrdered[T any](element T, index int, filterFunc func(element T) bool, newElements *[]T) {
	if filterFunc(element) {
		*newElements = append(*newElements, element)
	}
}

func applyFilterParallel[T any](s *stream[T], filterFunc func(element T) bool) {
	var toRemove []int
	actionApplier.Add(len(s.elements))
	for i, el := range s.elements {
		go checkFilterParallel(el, i, filterFunc, &toRemove)
	}
	actionApplier.Wait()
	s.elements = removeFromSlice(s.elements, toRemove)
}

func checkFilterParallel[T any](element T, index int, filterFunc func(element T) bool, toRemove *[]int) {
	defer actionApplier.Done()
	if !filterFunc(element) {
		sliceModifier.Lock()
		*toRemove = append(*toRemove, index)
		sliceModifier.Unlock()
	}
}

func removeFromSlice[T any](slice []T, toRremove []int) []T {
	sort.Slice(toRremove, func(i, j int) bool {
		return toRremove[i] > toRremove[j]
	})
	lenSlice, lenRemove := len(slice), len(toRremove)
	for i, val := range toRremove {
		slice[val] = slice[lenSlice-(lenRemove-i)]
	}
	return slice[:lenSlice-lenRemove]
}

func (s *stream[T]) applyReverse() {
	for i, j := 0, len(s.elements)-1; i < j; i, j = i+1, j-1 {
		s.elements[i], s.elements[j] = s.elements[j], s.elements[i]
	}
}
