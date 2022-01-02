package stream

type Stream[T any] interface {
	Filter(filterFunc func(element T) bool) *stream[T]
	Reverse() *stream[T]

	Collect(Collector[T]) []T
	Count() int
	AnyMatch(matchFunc func(element T) bool) bool
	AllMatch(matchFunc func(element T) bool) bool

	Flag(flag flag) *stream[T]
}

type stream[T any] struct {
	elements []T
	ordered  bool
}
