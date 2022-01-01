package stream

// TODO stream:
//  - everything

type Stream[T any] interface {
	Collect(SliceCollector[T]) []T
}

type stream[T any] struct {
	elements []T
	sequence []action[T]
}

type action[T any] struct {
	actionType actionType
	filterFunc func(el T) bool
}

type actionType int

const (
	filter actionType = iota
	mapToT
	mapToString
	mapToInt
	mapToBool
	peek
)
