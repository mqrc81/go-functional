package stream

type SliceCollector[T any] func(stream stream[T]) []T

func (s *stream[T]) Collect(collector SliceCollector[T]) []T {
	var collectorFunc func(stream stream[T]) []T = collector
	return collectorFunc(*s)
}

func ToSlice[T any]() SliceCollector[T] {
	return func(stream stream[T]) []T {
		applySequence(&stream.elements, &stream.sequence)
		return stream.elements
	}
}

func applySequence[T any](elements *[]T, sequence *[]action[T]) {
	for _, action := range *sequence {
		switch action.actionType {
		case filter:
			applyFilter(elements, action.filterFunc)
		}
	}
}

func applyFilter[T any](elements *[]T, filterFunc func(el T) bool) {
	var toRemove []int
	for i, el := range *elements {
		if !filterFunc(el) {
			toRemove = append(toRemove, i)
		}
	}
	*elements = removeFromSlice(*elements, toRemove)
}

func removeFromSlice[T any](slice []T, remove []int) []T {
	lenSlice, lenRemove := len(slice), len(remove)
	for i := len(remove) - 1; i >= 0; i-- {
		slice[remove[i]] = slice[lenSlice-(lenRemove-i)]
	}
	return slice[:len(slice)-len(remove)]
}
