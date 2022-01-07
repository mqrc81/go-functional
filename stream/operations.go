package stream

type operation[T any] interface {
	apply(element *T, operations *[]operation[T], current int, ordered bool) bool
}

func applyIfHasNext[T any](element *T, operations *[]operation[T], current int, ordered bool) bool {
	if len(*operations)-1 > current {
		current++
		return (*operations)[current].apply(element, operations, current, ordered)
	}
	return true
}

type filterOperation[T any] struct {
	filterFunc func(element T) bool
}

func (o filterOperation[T]) apply(element *T, operations *[]operation[T], current int, ordered bool) bool {
	var filterFunc func(element T) bool = o.filterFunc
	if filterFunc(*element) {
		return applyIfHasNext(element, operations, current, ordered)
	}
	return false
}

type limitOperation[T any] struct {
	maxSize    int
	spotsTaken *int
}

func (o limitOperation[T]) apply(element *T, operations *[]operation[T], current int, ordered bool) bool {
	if *o.spotsTaken < o.maxSize {
		*o.spotsTaken++
		return applyIfHasNext(element, operations, current, ordered)
	}
	return false
}